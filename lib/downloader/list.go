package downloader

import (
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
	"github.com/yarbelk/cdda-manager/build"
)

// Option is for Tiled or Curses
type Option int

const (
	// Tiles is the graphical SDL based version
	Tiles Option = iota
	// Curses is the venerable but still cool ASCII version
	Curses
)

// AvailableVersion tells you about the version that ex
type AvailableVersion struct {
	Target string
	Name   string
}

func (av AvailableVersion) String() string {
	return av.Name
}

// ListGameVersions currently in the upstream repo
func ListGameVersions(os build.Build, option Option, base string) ([]AvailableVersion, error) {
	client := http.Client{Transport: http.DefaultTransport}
	var name string
	url, _ := url.Parse(base)
	switch build.BuildV {
	case build.LinuxX64:
		url.Path = filepath.Join(url.Path, "/Linux_x64")
	case build.Darwin:
		url.Path = filepath.Join(url.Path, "/OSX")
	case build.Windows:
		url.Path = filepath.Join(url.Path, "/Windows")
	case build.WindowsX64:
		url.Path = filepath.Join(url.Path, "/Windows_x64")
	default:
		panic("Unsupoorted")
	}

	switch option {
	case Tiles:
		url.Path = filepath.Join(url.Path, "/Tiles")
		name = "Tiles "
	case Curses:
		url.Path = filepath.Join(url.Path, "/Curses")
		name = "Curses "
	}

	response, _ := client.Get(url.String())
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	av := make([]AvailableVersion, 0)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// This is super fragile
		if href, ok := s.Attr("href"); ok && len(href) >= 12 && href[:12] == "cataclysmdda" {
			filename := getFilename(href, os)
			av = append(av, AvailableVersion{
				Target: url.String() + "/" + href,
				Name:   name + filename,
			})
		}

	})

	return av, nil
}

func getFilename(a string, os build.Build) string {
	switch build.BuildV {
	case build.LinuxX64:
		return a[:len(a)-len(".tar.gz")]
	case build.Darwin:
		return a[:len(a)-len(".dmg")]
	case build.WindowsX64, build.Windows:
		return a[:len(a)-len(".zip")]
	default:
		return a
	}
}
