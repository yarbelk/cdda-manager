package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/yarbelk/cdda-manager/build"
	"github.com/yarbelk/cdda-manager/lib/downloader"
)

const url = "http://dev.narc.ro/cataclysm/jenkins-latest/"

// listCommand prints out available versions
var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Print all the available versions for your OS",
	Long:  `Print all the available versions for your OS`,
	Run: func(cmd *cobra.Command, args []string) {
		if build.BuildV == build.Unsupported {
			log.Printf("%s:%s is an unsupported build", runtime.GOOS, runtime.GOARCH)
			os.Exit(1)
		}
		avTiled, _ := downloader.ListGameVersions(build.BuildV, downloader.Tiles, url)
		avCurses, _ := downloader.ListGameVersions(build.BuildV, downloader.Curses, url)
		av := append(avTiled, avCurses...)
		for _, a := range av {
			fmt.Println(a)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
