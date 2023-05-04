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
	"bytes"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func TestVersionCommand() {
	cmd := NewVersionCommand()
	buf := new(bytes.Buffer)
	err := cmd.Execute()
	cmd.SetOut(buf)
	fmt.Println("buf0:", buf)

	if err == nil {
		fmt.Println("buf:", buf)
		output := buf.String()
		expected := "OpenRHINO v0.2.0\nOpen MPI v4.1.5"
		if !strings.Contains(output, expected) {
			fmt.Printf("expected output to contain %q, but got %q", expected, output)
		} else {
			fmt.Println("TestVersionCommand passed")
		}
	} else {
		fmt.Printf("TestVersionCommand failed")
	}

}
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "rhino",
		Short: "\nRHINO-CLI - Manage your OpenRHINO functions and jobs",
	}

	rootCmd.AddCommand(NewCreateCommand())
	rootCmd.AddCommand(NewBuildCommand())
	rootCmd.AddCommand(NewDeleteCommand())
	rootCmd.AddCommand(NewRunCommand())
	rootCmd.AddCommand(NewListCommand())
	rootCmd.AddCommand(NewDockerRunCommand())
	rootCmd.AddCommand(NewVersionCommand())
	return rootCmd
}
