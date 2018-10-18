// +build linux

package launcher

import (
	"bufio"
	"io"
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/yarbelk/cdda-manager/lib/game"
)

// LaunchGame with system configurations
func LaunchGame(gameVersion game.Interface) error {
	logger := logrus.WithFields(logrus.Fields{
		"executable": gameVersion.Executable(),
	})

	cmd := exec.Command(gameVersion.Executable())
	cmd.Dir = gameVersion.WorkingDir()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	multi := io.MultiReader(stdout, stderr)
	if err := cmd.Start(); err != nil {
		return err
	}
	in := bufio.NewScanner(multi)
	for in.Scan() {
		logger.Println(in.Text())
		if err = in.Err(); err != nil {
			return in.Err()
		}
	}
	return cmd.Wait()
}
