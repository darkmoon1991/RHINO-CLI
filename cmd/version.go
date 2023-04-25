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

	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of MPI installed on the local machine",
		Run: func(cmd *cobra.Command, args []string) {
			output, err := exec.Command("mpirun", "--version").Output()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("MPI version: %s\n", string(output))
		},
	}
	return versionCmd
}
