package cmd

import (
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yarbelk/cdda-manager/build"
	"github.com/yarbelk/cdda-manager/lib/downloader"
)

// listCommand prints out available versions
var downloadCommand = &cobra.Command{
	Use:   "download",
	Short: "Download the thing",
	Long:  `Download the specified file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var av []downloader.AvailableVersion
		if build.BuildV == build.Unsupported {
			log.Printf("%s:%s is an unsupported build", runtime.GOOS, runtime.GOARCH)
			os.Exit(1)
		}
		if len(args) != 1 {
			cmd.Printf("Need a filename")
		}
		tgt := args[0]
		if strings.HasPrefix(tgt, "Tiles") {
			av, _ = downloader.ListGameVersions(build.BuildV, downloader.Tiles, url)
		} else if strings.HasPrefix(tgt, "Cusrses") {
			av, _ = downloader.ListGameVersions(build.BuildV, downloader.Curses, url)
		}
		for _, a := range av {
			if a.Name == tgt {
				if err := downloader.GetPackage(a, os.Getenv("HOME")); err != nil {
					panic(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCommand)
}
