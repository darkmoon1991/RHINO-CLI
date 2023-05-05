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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	cmdOutput, err := execShellCmd("rhino", []string{"version"})
	assert.Equal(t, nil, err, "test run failed: %s", errorMessage(err))
	// Check that the output is contained "kubenetes version"&&"RhinoServer version"&&"RhinoClient version"
	assert.Contains(t, cmdOutput, "Kubernetes version")
	assert.Contains(t, cmdOutput, "RhinoServer version")
	assert.Contains(t, cmdOutput, "RhinoClient version")
	//if the kubeconfig is not exist or error path, the output should be "error building kubeconfig"
	cmdOutput1, err1 := execShellCmd("rhino", []string{"version", "--kubeconfig", "/home/zhao/.kube/config"})
	assert.Error(t, err1, "test run failed: %s", errorMessage(err1))
	assert.Contains(t, cmdOutput1, "error building kubeconfig")
}
