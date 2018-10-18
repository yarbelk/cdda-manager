// +build windows

package launcher

import (
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
	return cmd.Run()
}
