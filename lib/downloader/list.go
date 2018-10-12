package downloader

import (
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
)

type Option int
type Build int

const (
	Tiles Option = iota
	Curses
)

// only 64 bit for now
const (
	Linux Build = iota
	LinuxX64
	Darwin
	Windows
	WindowsX64
)

type AvailableVersion struct {
	Target string
	Name   string
}

func ListGameVersions(os Build, option Option, base string) ([]AvailableVersion, error) {
	client := http.Client{Transport: http.DefaultTransport}
	var name string
	url, _ := url.Parse(base)
	switch os {
	case Linux:
		url.Path = filepath.Join(url.Path, "/Linux")
	case LinuxX64:
		url.Path = filepath.Join(url.Path, "/Linux_x64")
	case Darwin:
		url.Path = filepath.Join(url.Path, "/OSX")
	case Windows:
		url.Path = filepath.Join(url.Path, "/Windows")
	case WindowsX64:
		url.Path = filepath.Join(url.Path, "/Windows_x64")
	}

	switch option {
	case Tiles:
		url.Path = filepath.Join(url.Path, "/Tiles")
		name = "Tiles "
	case Curses:
		url.Path = filepath.Join(url.Path, "/Curses")
	}

	response, _ := client.Get(url.String())
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	av := make([]AvailableVersion, 0)
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		if alt, ok := s.Find("img").Attr("alt"); ok && alt == "[   ]" {
			a, _ := s.Find("a").Attr("href")
			filename := getFilename(a, os)
			av = append(av, AvailableVersion{
				Target: url.String() + "/" + a,
				Name:   name + filename,
			})
		}

	})

	return av, nil
}

func getFilename(a string, os Build) string {
	switch os {
	case Linux, LinuxX64:
		return a[:len(a)-len(".tar.gz")]
	case Darwin:
		return a[:len(a)-len(".dmg")]
	case WindowsX64, Windows:
		return a[:len(a)-len(".zip")]
	default:
		return a
	}
}
