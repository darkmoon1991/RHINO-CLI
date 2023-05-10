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
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type VersionOptions struct {
	kubeconfig string
}

const RHINOCLIENTVERSION = "v0.2.0"
const RHINOCRDNAME = "rhinojobs.openrhino.org"

// NewVersionCommand creates a new version command
func NewVersionCommand() *cobra.Command {
	versionOptions := &VersionOptions{}
	versionCmd := &cobra.Command{
		Use:       "version",
		Short:     "Print the version of RhinoClient and kubernetes installed on the local machine,and the version of RhinoServer ",
		RunE:      versionOptions.RunVersionCommand,
		ValidArgs: []string{"kubeconfig"},
	}
	versionCmd.Flags().StringVar(&versionOptions.kubeconfig, "kubeconfig", "", "the path of the kubeconfig file")

	return versionCmd
}

// getKubernetesVersion returns the version of Kubernetes installed on the local machine
func (v *VersionOptions) getKubernetesVersion() (string, error) {

	//build kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", v.kubeconfig)
	if err != nil {
		return "", fmt.Errorf("error building kubeconfig: %s", err.Error())
	}
	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", fmt.Errorf("error creating clientset: %s", err.Error())
	}
	// Get the Kubernetes server version
	version, err := clientset.Discovery().ServerVersion()
	if err != nil {
		fmt.Println("Error getting server version:", err.Error())
		return "", err
	}
	// Print the version information
	return version.String(), nil
}

// getRhinoServerVersion prints the version of Server Rhino
func (v *VersionOptions) getRhinoServerVersion() (string, error) {

	//build kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", v.kubeconfig)
	if err != nil {
		return "", fmt.Errorf("error building kubeconfig: %s", err.Error())
	}

	// 创建 API Extension 客户端
	apiExtClient, err := apiextv1.NewForConfig(config)
	if err != nil {
		return "", fmt.Errorf("error creating apiExtClient: %s", err.Error())
	}
	// 获取指定 CRD 的 Group 和 Version
	crd, err := apiExtClient.CustomResourceDefinitions().Get(context.TODO(), RHINOCRDNAME, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("error getting crd: %s", err.Error())
	}
	// concatenate all the available versions of RhinoJob Interfaces into a string
	var RhinoServerVersion []string
	for _, v := range crd.Spec.Versions {
		RhinoServerVersion = append(RhinoServerVersion, v.Name)
	}
	RhinoServerVersionStr := strings.Join(RhinoServerVersion, ", ")

	return RhinoServerVersionStr, nil
}

// RunVersionCommand runs the version command
func (v *VersionOptions) RunVersionCommand(cmd *cobra.Command, args []string) error {
	if v.kubeconfig == "" {
		homeDir, err := homedir.Dir()
		if err != nil {
			return fmt.Errorf("could not get home directory: %s,please use --kubeconfig to specify the absolute path", err)
		}
		kubeconfigPath := filepath.Join(homeDir, ".kube", "config")
		_, err = os.Stat(kubeconfigPath)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("kubeconfig file not found at %s, please use --kubeconfig to specify the absolute path", kubeconfigPath)
			}
			return fmt.Errorf("error checking kubeconfig file at %s: %s", kubeconfigPath, err)
		}
		v.kubeconfig = kubeconfigPath
	}

	// Print the version of Kubernetes
	kubernetesVersion, err := v.getKubernetesVersion()
	if err != nil {
		return err
	}
	fmt.Println("Kubernetes version:", kubernetesVersion)

	// Print the version of RhinoServer
	rhinoServerVersion, err := v.getRhinoServerVersion()
	if err != nil {
		return err
	}
	fmt.Println("RhinoServer version:", rhinoServerVersion)

	// Print the version of RhinoClient
	fmt.Println("RhinoClient version:", RHINOCLIENTVERSION)
	return nil
}
