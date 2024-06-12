/*
 * Copyright 2023 RHINO Team
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	rhinojob "github.com/OpenRHINO/RHINO-Operator/api/v1alpha2"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/dynamic"
)

type RunOptions struct {
	parallel   int
	timeToLive int
	dataPath   string
	dataServer string
	funcName   string

	//the fields in v1alpha2 API
	memoryAllocationMode string
  	memoryAllocationSize int

	kubeconfig string
	namespace  string
}

func NewRunCommand() *cobra.Command {
	runOpts := &RunOptions{}
	runCmd := &cobra.Command{
		Use:   "run [image]",
		Short: "Submit and run a RHINO job",
		Long:  "\nSubmit an MPI function/project and run it as a RHINO job",
		Example: `  rhino run hello:v1.0 --namespace user_space
  rhino run foo/matmul:v2.1 --np 4 -- arg1 arg2 
  rhino run mpi/testbench -n 32 -t 800 --server 10.0.0.7 --dir /mnt -- --in=/data/file --out=/data/out`,
		RunE: runOpts.run,
	}

	runCmd.Flags().StringVar(&runOpts.dataServer, "server", "", "IP address of an NFS server")
	runCmd.Flags().StringVar(&runOpts.dataPath, "dir", "", "a directory in the NFS server, to store data and shared with all the MPI processes")
	runCmd.MarkFlagsRequiredTogether("server", "dir")
	runCmd.Flags().IntVar(&runOpts.parallel, "np", 1, "the number of MPI processes")
	runCmd.Flags().IntVarP(&runOpts.timeToLive, "ttl", "t", 600, "Time To Live (seconds). The RHINO job will be deleted after this time, whether it is completed or not.")
	runCmd.Flags().StringVar(&runOpts.memoryAllocationMode, "mem-mode", "FixedPerCoreMemory", "the memory allocation mode of the RHINO job, choose from [FixedTotalMemory, FixedPerCoreMemory]")
	runCmd.Flags().IntVar(&runOpts.memoryAllocationSize, "mem-size", 2, "the memory allocation size of the RHINO job, in GB")
	runCmd.Flags().StringVarP(&runOpts.namespace, "namespace", "n", "", "the namespace of the RHINO job")
	runCmd.Flags().StringVar(&runOpts.kubeconfig, "kubeconfig", "", "the path of the kubeconfig file")

	return runCmd
}

func (r *RunOptions) run(cmd *cobra.Command, args []string) error {
	// Check the arguments
	if len(args) == 0 {
		cmd.Help()
		return nil
	}
	r.funcName = getFuncName(args[0])
	if r.parallel < 1 {
		return fmt.Errorf("the number of MPI processes (--np) must be greater than 0")
	}
	if r.timeToLive < 0 {
		return fmt.Errorf("the time to live (--ttl) must be greater than or equal to 0")
	}
	if r.memoryAllocationMode != "FixedTotalMemory" && r.memoryAllocationMode != "FixedPerCoreMemory" {
		return fmt.Errorf("the memory allocation mode (--mem-mode) must be either FixedTotalMemory or FixedPerCoreMemory")
	}
	if r.memoryAllocationSize < 1 {
		return fmt.Errorf("the memory allocation size (--mem-size) must be greater than or equal to 1")
	}

	var err error
	r.kubeconfig, err = getKubeconfigPath(r.kubeconfig)
	if err != nil {
		return fmt.Errorf("%v, please set the kubeconfig path by --kubeconfig", err)
	}

	dynamicClient, currentNamespace, err := buildFromKubeconfig(r.kubeconfig)
	if err != nil {
		return err
	}
	if r.namespace == "" {
		r.namespace = *currentNamespace
	}

	// Create a RHINO job
	_, err = r.runRhinoJob(dynamicClient, args)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("failed to create a RHINO job")
	}
	fmt.Println("RhinoJob", r.funcName, "created")
	return nil
}

func (r *RunOptions) printYAML(args []string) (yamlFile string) {
	yamlFile = `apiVersion: openrhino.org/v1alpha2
kind: RhinoJob
metadata:
  labels:
    app.kubernetes.io/name: rhinojob 
    app.kubernetes.io/instance: "`
	yamlFile += r.funcName + `"
    app.kubernetes.io/part-of: rhino-operator
    app.kubernetes.io/managed-by: kustomize
    app.kubernetes.io/created-by: rhino-operator
  name: "`
	yamlFile += r.funcName + `"
spec:
  image: "`
	yamlFile += args[0] + `"
  ttl: `
	yamlFile += strconv.Itoa(r.timeToLive) + `
  parallelism: `
	yamlFile += strconv.Itoa(r.parallel) + ` 
  memoryAllocationMode: `
  	yamlFile += r.memoryAllocationMode + `
  memoryAllocationSize: `
    yamlFile += strconv.Itoa(r.memoryAllocationSize) + `
  appExec: "/app/mpi-func"`
	if len(args) > 1 {
		yamlFile += `
  appArgs: [`
		for i := 1; i < len(args); i++ {
			yamlFile += `"` + args[i] + `", `
		}
		yamlFile += `]`
	}
	if len(r.dataServer) != 0 {
		yamlFile += `
  dataServer: "` + r.dataServer + `"`
		yamlFile += `
  dataPath: "` + r.dataPath + `"`
	}
	return yamlFile
}

func (r *RunOptions) runRhinoJob(client dynamic.Interface, args []string) (*rhinojob.RhinoJobList, error) {
	decoder := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	obj := &unstructured.Unstructured{}
	rhinojobYAML := r.printYAML(args)
	if _, _, err := decoder.Decode([]byte(rhinojobYAML), nil, obj); err != nil {
		return nil, err
	}
	createdRhinoJob, err := client.Resource(RhinoJobGVR).Namespace(r.namespace).Create(context.TODO(), obj, metav1.CreateOptions{})

	if err != nil {
		return nil, err
	}
	data, err := createdRhinoJob.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var rj rhinojob.RhinoJobList
	if err := json.Unmarshal(data, &rj); err != nil {
		return nil, err
	}
	return &rj, nil
}
