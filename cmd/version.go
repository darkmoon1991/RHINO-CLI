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
	"fmt"
	"path/filepath"

	rhinojob "github.com/OpenRHINO/RHINO-Operator/api/v1alpha1"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// printKubernetesVersion prints the version of Kubernetes installed on the local machine
func printKubernetesVersion() (string, error) {
	// Load the Kubernetes configuration from file
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		return "", fmt.Errorf("kubeconfig file not found, please use --config to specify the absolute path")
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Get the Kubernetes server version
	version, err := clientset.Discovery().ServerVersion()
	if err != nil {
		panic(err)
	}
	// Print the version information
	return version.String(), nil
}

// printRhinojobVersion prints the version of Rhino installed on the local machine
func printRhinojobVersion() (string, error) {
	rhinojobVersion := rhinojob.GroupVersion.Version
	if len(rhinojobVersion) == 0 {
		return "", fmt.Errorf("neither openmpi nor mpich found")
	}
	return rhinojob.GroupVersion.Version, nil
}

// NewVersionCommand creates a new version command
func NewVersionCommand() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version ",
		Short: "Print the version of Rhino and MPI installed on the local machine",
		RunE: func(cmd *cobra.Command, args []string) error {
			rhinojobVersion, err := printRhinojobVersion()
			if err != nil {
				return err
			}
			fmt.Printf("OpenRHINOJob %s\n", rhinojobVersion)
			k8sVersion, err := printKubernetesVersion()
			if err != nil {
				return err
			}
			fmt.Printf("Kubernetes %s\n", k8sVersion)
			return nil
		},
	}
	return versionCmd
}
