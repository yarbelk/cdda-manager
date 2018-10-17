package list

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/yarbelk/cdda-manager/lib/downloader"
)

const url = "http://dev.narc.ro/cataclysm/jenkins-latest/"

// ListCommand prints out available versions
var ListCommand = &cobra.Command{
	Use:   "list",
	Short: "Print all the available versions for your OS",
	Long:  `Print all the available versions for your OS`,
	Run: func(cmd *cobra.Command, args []string) {
		var os downloader.Build
		if runtime.GOARCH == "amd64" && runtime.GOOS == "linux" {
			os = downloader.LinuxX64
		} else {
			panic("stupid os")
		}

		avTiled, _ := downloader.ListGameVersions(os, downloader.Tiles, url)
		avCurses, _ := downloader.ListGameVersions(os, downloader.Tiles, url)
		av := append(avTiled, avCurses...)
		for _, a := range av {
			fmt.Println(a)
		}
	},
}
