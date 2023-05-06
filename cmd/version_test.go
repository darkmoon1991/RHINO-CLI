// generate a test for version
// Path: cmd/version_test.go
// Compare this snippet from cmd/root_test.go:
//   - you may not use this file except in compliance with the License.
//   - You may obtain a copy of the License at
//     *
//   - http://www.apache.org/licenses/LICENSE-2.0
//     *
//   - Unless required by applicable law or agreed to in writing, software
//   - distributed under the License is distributed on an "AS IS" BASIS,
//   - WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   - See the License for the specific language governing permissions and
//   - limitations under the License.
//     */
package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	// change work directory to ${workspaceFolder}
	cwd, err := os.Getwd()
	assert.Equal(t, nil, err, "test command version failed: %s", errorMessage(err))
	if strings.HasSuffix(cwd, "cmd") {
		os.Chdir("..")
	}
	rootCmd := NewRootCommand()
	rootCmd.SetArgs([]string{"version"})
	err = rootCmd.Execute()
	assert.Equal(t, nil, err, "test command version failed: %s", errorMessage(err))

	rootCmd.SetArgs([]string{"version", "--kubeconfig", "/home/zhao/.kube/config"})
	err = rootCmd.Execute()
	assert.Error(t, err, "test command version failed: %s", errorMessage(err))
}
