package downloader_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/yarbelk/cdda-manager/lib/downloader"
)

func TestGetNightlyList(t *testing.T) {
	var body nopCloser = nopCloser{
		bytes.NewBufferString(LinuxTilesPage),
	}
	var roundTripFunction RoundTripFunc = func(r *http.Request) *http.Response {
		resp := http.Response{
			Status:        http.StatusText(http.StatusOK),
			StatusCode:    http.StatusOK,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Body:          body,
			ContentLength: 7323,
		}
		return &resp
	}

	http.DefaultTransport = roundTripFunction

	os := downloader.LinuxX64
	option := downloader.Tiles
	base := "http://dev.narc.ro/cataclysm/jenkins-latest/"

	listFiles, err := downloader.ListGameVersions(os, option, base)

	if err != nil {
		t.Logf("Got error: %+v", err)
		t.FailNow()
	}

	if len(listFiles) != len(LinuxTileFilenames) {
		t.Logf("Didn't get any data back")
		t.Fail()
	}

	for i, v := range listFiles {
		if v != LinuxTileFilenames[i] {
			t.Logf("didn't get expected filenames: got:\n %+v want: %+v", v, LinuxTileFilenames[i])
			t.Fail()
		}
	}
}
