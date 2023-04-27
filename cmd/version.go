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
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// rhinoVersion is the version of Rhino
var (
	rhinoVersion string = "v0.2.0"
)

// printMpiVersion prints the version of MPI installed on the local machine
func printMpiVersion() (string, error) {
	var output []byte
	var err error
	var versions []string
	if _, err = exec.LookPath("ompi_info"); err == nil {
		output, err = exec.Command("ompi_info", "--version").Output()
		if err == nil {
			versions = append(versions, strings.Split(string(output), "\n")[0])
		}
	}
	if _, err = exec.LookPath("mpichversion"); err == nil {
		output, err = exec.Command("mpichversion").Output()
		if err == nil {
			versions = append(versions, strings.Split(string(output), "\n")[0])
		}
	}
	if len(versions) == 0 {
		return "", fmt.Errorf("neither openmpi nor mpich found")
	}
	return strings.Join(versions, "\n"), nil
}

// NewVersionCommand creates a new version command
func NewVersionCommand() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version ",
		Short: "Print the version of Rhino and MPI installed on the local machine",
		RunE: func(cmd *cobra.Command, args []string) error {
			mpiVersion, err := printMpiVersion()
			if err != nil {
				return err
			}
			fmt.Printf("OpenRHINO %s\n%s\n", rhinoVersion, mpiVersion)
			return nil
		},
	}
	return versionCmd
}
