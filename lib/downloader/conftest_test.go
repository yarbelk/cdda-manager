package downloader_test

import (
	"io"
	"net/http"

	"github.com/yarbelk/cdda-manager/lib/downloader"
)

const LinuxTilesPage = `
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<html>
 <head>
  <title>Index of /cataclysm/jenkins-latest/Linux_x64/Tiles</title>
 </head>
 <body>
<h1>Index of /cataclysm/jenkins-latest/Linux_x64/Tiles</h1>
<pre><img src="/icons/blank.gif" alt="Icon "> <a href="?C=N;O=A">Name</a>                                 <a href="?C=M;O=A">Last modified</a>      <a href="?C=S;O=A">Size</a>  <a href="?C=D;O=A">Description</a><hr><img src="/icons/back.gif" alt="[PARENTDIR]"> <a href="/cataclysm/jenkins-latest/Linux_x64/">Parent Directory</a>                                          -   
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8058.tar.gz">cataclysmdda-0.C-8058.tar.gz</a>         2018-10-17 16:06   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8057.tar.gz">cataclysmdda-0.C-8057.tar.gz</a>         2018-10-17 05:51   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8056.tar.gz">cataclysmdda-0.C-8056.tar.gz</a>         2018-10-16 23:26   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8055.tar.gz">cataclysmdda-0.C-8055.tar.gz</a>         2018-10-16 09:10   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8054.tar.gz">cataclysmdda-0.C-8054.tar.gz</a>         2018-10-16 06:19   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8053.tar.gz">cataclysmdda-0.C-8053.tar.gz</a>         2018-10-16 05:28   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8052.tar.gz">cataclysmdda-0.C-8052.tar.gz</a>         2018-10-16 00:05   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8051.tar.gz">cataclysmdda-0.C-8051.tar.gz</a>         2018-10-15 21:49   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8050.tar.gz">cataclysmdda-0.C-8050.tar.gz</a>         2018-10-15 06:24   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8049.tar.gz">cataclysmdda-0.C-8049.tar.gz</a>         2018-10-14 18:46   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8048.tar.gz">cataclysmdda-0.C-8048.tar.gz</a>         2018-10-14 07:32   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8047.tar.gz">cataclysmdda-0.C-8047.tar.gz</a>         2018-10-14 05:29   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8046.tar.gz">cataclysmdda-0.C-8046.tar.gz</a>         2018-10-13 10:25   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8045.tar.gz">cataclysmdda-0.C-8045.tar.gz</a>         2018-10-13 05:40   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8044.tar.gz">cataclysmdda-0.C-8044.tar.gz</a>         2018-10-12 13:39   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8043.tar.gz">cataclysmdda-0.C-8043.tar.gz</a>         2018-10-12 12:51   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8042.tar.gz">cataclysmdda-0.C-8042.tar.gz</a>         2018-10-12 05:33   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8041.tar.gz">cataclysmdda-0.C-8041.tar.gz</a>         2018-10-12 00:02   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8040.tar.gz">cataclysmdda-0.C-8040.tar.gz</a>         2018-10-11 20:03   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8039.tar.gz">cataclysmdda-0.C-8039.tar.gz</a>         2018-10-11 10:39   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8038.tar.gz">cataclysmdda-0.C-8038.tar.gz</a>         2018-10-11 09:13   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8037.tar.gz">cataclysmdda-0.C-8037.tar.gz</a>         2018-10-11 07:29   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8036.tar.gz">cataclysmdda-0.C-8036.tar.gz</a>         2018-10-11 05:20   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8035.tar.gz">cataclysmdda-0.C-8035.tar.gz</a>         2018-10-11 04:33   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8034.tar.gz">cataclysmdda-0.C-8034.tar.gz</a>         2018-10-10 14:34   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8033.tar.gz">cataclysmdda-0.C-8033.tar.gz</a>         2018-10-10 13:27   63M  
<img src="/icons/compressed.gif" alt="[   ]"> <a href="cataclysmdda-0.C-8032.tar.gz">cataclysmdda-0.C-8032.tar.gz</a>         2018-10-10 09:50   63M  
<hr></pre>
<address>Apache/2.4.18 (Ubuntu) Server at dev.narc.ro Port 80</address>
</body></html>
`

var LinuxTileFilenames []downloader.AvailableVersion = []downloader.AvailableVersion{
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8058.tar.gz", "Tiles cataclysmdda-0.C-8058"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8057.tar.gz", "Tiles cataclysmdda-0.C-8057"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8056.tar.gz", "Tiles cataclysmdda-0.C-8056"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8055.tar.gz", "Tiles cataclysmdda-0.C-8055"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8054.tar.gz", "Tiles cataclysmdda-0.C-8054"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8053.tar.gz", "Tiles cataclysmdda-0.C-8053"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8052.tar.gz", "Tiles cataclysmdda-0.C-8052"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8051.tar.gz", "Tiles cataclysmdda-0.C-8051"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8050.tar.gz", "Tiles cataclysmdda-0.C-8050"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8049.tar.gz", "Tiles cataclysmdda-0.C-8049"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8048.tar.gz", "Tiles cataclysmdda-0.C-8048"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8047.tar.gz", "Tiles cataclysmdda-0.C-8047"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8046.tar.gz", "Tiles cataclysmdda-0.C-8046"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8045.tar.gz", "Tiles cataclysmdda-0.C-8045"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8044.tar.gz", "Tiles cataclysmdda-0.C-8044"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8043.tar.gz", "Tiles cataclysmdda-0.C-8043"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8042.tar.gz", "Tiles cataclysmdda-0.C-8042"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8041.tar.gz", "Tiles cataclysmdda-0.C-8041"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8040.tar.gz", "Tiles cataclysmdda-0.C-8040"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8039.tar.gz", "Tiles cataclysmdda-0.C-8039"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8038.tar.gz", "Tiles cataclysmdda-0.C-8038"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8037.tar.gz", "Tiles cataclysmdda-0.C-8037"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8036.tar.gz", "Tiles cataclysmdda-0.C-8036"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8035.tar.gz", "Tiles cataclysmdda-0.C-8035"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8034.tar.gz", "Tiles cataclysmdda-0.C-8034"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8033.tar.gz", "Tiles cataclysmdda-0.C-8033"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8032.tar.gz", "Tiles cataclysmdda-0.C-8032"},
}

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() (err error) { return }
