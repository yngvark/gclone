// Package execute can run our program.
package execute

import (
	"os/exec"
	"path"

	"github.com/yngvark.com/clonerepo/pkg/testhelper/build_executable"
)

// CloneRepo runs the clonerepo command.
func CloneRepo(arg ...string) *exec.Cmd {
	return exec.Command(cloneRepoCmd(), arg...) //nolint:gosec // This should be secure
}

func cloneRepoCmd() string {
	return path.Join(build_executable.ProjectRoot(), "clonerepo")
}
