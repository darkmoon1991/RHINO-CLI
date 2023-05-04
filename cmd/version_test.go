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

import "testing"

// generate a test for version, and test the version command is rhino version，output is OpenRHINOJob v1alpha1\nKubernetes v1.25.4
func TestVersion(t *testing.T) {
	out, err := execShellCmd("rhino", []string{"version"})
	if err != nil {
		t.Errorf("execShellCmd error: %v", err)
	}
	if out != "OpenRHINOJob v1alpha1\nKubernetes v1.25.4\n" {
		t.Errorf("execShellCmd error: %v", err)
	}
}
