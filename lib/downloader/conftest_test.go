package downloader_test

import (
	"io"
	"net/http"

	"github.com/yarbelk/cdda-manager/lib/downloader"
)

const LinuxTilesPage = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<html>
 <head>
  <title>Index of /cataclysm/jenkins-latest/Linux_x64/Tiles</title>
 </head>
 <body>
<h1>Index of /cataclysm/jenkins-latest/Linux_x64/Tiles</h1>
  <table>
   <tr><th valign="top"><img src="/icons/blank.gif" alt="[ICO]"></th><th><a href="?C=N;O=D">Name</a></th><th><a href="?C=M;O=A">Last modified</a></th><th><a href="?C=S;O=A">Size</a></th><th><a href="?C=D;O=A">Description</a></th></tr>
   <tr><th colspan="5"><hr></th></tr>
<tr><td valign="top"><img src="/icons/back.gif" alt="[PARENTDIR]"></td><td><a href="/cataclysm/jenkins-latest/Linux_x64/">Parent Directory</a></td><td>&nbsp;</td><td align="right">  - </td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8015.tar.gz">cataclysmdda-0.C-8015.tar.gz</a></td><td align="right">2018-10-05 02:26  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8016.tar.gz">cataclysmdda-0.C-8016.tar.gz</a></td><td align="right">2018-10-05 03:13  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8017.tar.gz">cataclysmdda-0.C-8017.tar.gz</a></td><td align="right">2018-10-05 05:15  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8018.tar.gz">cataclysmdda-0.C-8018.tar.gz</a></td><td align="right">2018-10-05 06:35  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8019.tar.gz">cataclysmdda-0.C-8019.tar.gz</a></td><td align="right">2018-10-05 12:55  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8020.tar.gz">cataclysmdda-0.C-8020.tar.gz</a></td><td align="right">2018-10-06 03:23  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8021.tar.gz">cataclysmdda-0.C-8021.tar.gz</a></td><td align="right">2018-10-06 05:56  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8022.tar.gz">cataclysmdda-0.C-8022.tar.gz</a></td><td align="right">2018-10-06 06:28  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8023.tar.gz">cataclysmdda-0.C-8023.tar.gz</a></td><td align="right">2018-10-06 15:32  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8024.tar.gz">cataclysmdda-0.C-8024.tar.gz</a></td><td align="right">2018-10-07 06:08  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8025.tar.gz">cataclysmdda-0.C-8025.tar.gz</a></td><td align="right">2018-10-07 18:42  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8026.tar.gz">cataclysmdda-0.C-8026.tar.gz</a></td><td align="right">2018-10-08 06:09  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8027.tar.gz">cataclysmdda-0.C-8027.tar.gz</a></td><td align="right">2018-10-08 23:34  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8028.tar.gz">cataclysmdda-0.C-8028.tar.gz</a></td><td align="right">2018-10-09 04:24  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8029.tar.gz">cataclysmdda-0.C-8029.tar.gz</a></td><td align="right">2018-10-09 05:33  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8030.tar.gz">cataclysmdda-0.C-8030.tar.gz</a></td><td align="right">2018-10-09 20:12  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8031.tar.gz">cataclysmdda-0.C-8031.tar.gz</a></td><td align="right">2018-10-10 06:50  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8032.tar.gz">cataclysmdda-0.C-8032.tar.gz</a></td><td align="right">2018-10-10 09:50  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8033.tar.gz">cataclysmdda-0.C-8033.tar.gz</a></td><td align="right">2018-10-10 13:27  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8034.tar.gz">cataclysmdda-0.C-8034.tar.gz</a></td><td align="right">2018-10-10 14:34  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8035.tar.gz">cataclysmdda-0.C-8035.tar.gz</a></td><td align="right">2018-10-11 04:33  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8036.tar.gz">cataclysmdda-0.C-8036.tar.gz</a></td><td align="right">2018-10-11 05:20  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8037.tar.gz">cataclysmdda-0.C-8037.tar.gz</a></td><td align="right">2018-10-11 07:29  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8038.tar.gz">cataclysmdda-0.C-8038.tar.gz</a></td><td align="right">2018-10-11 09:13  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8039.tar.gz">cataclysmdda-0.C-8039.tar.gz</a></td><td align="right">2018-10-11 10:39  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8040.tar.gz">cataclysmdda-0.C-8040.tar.gz</a></td><td align="right">2018-10-11 20:03  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
<tr><td valign="top"><img src="/icons/compressed.gif" alt="[   ]"></td><td><a href="cataclysmdda-0.C-8041.tar.gz">cataclysmdda-0.C-8041.tar.gz</a></td><td align="right">2018-10-12 00:02  </td><td align="right"> 63M</td><td>&nbsp;</td></tr>
   <tr><th colspan="5"><hr></th></tr>
</table>
<address>Apache/2.4.18 (Ubuntu) Server at dev.narc.ro Port 80</address>
</body></html>`

var LinuxTileFilenames []downloader.AvailableVersion = []downloader.AvailableVersion{
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8015.tar.gz", "Tiles cataclysmdda-0.C-8015"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8016.tar.gz", "Tiles cataclysmdda-0.C-8016"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8017.tar.gz", "Tiles cataclysmdda-0.C-8017"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8018.tar.gz", "Tiles cataclysmdda-0.C-8018"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8019.tar.gz", "Tiles cataclysmdda-0.C-8019"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8020.tar.gz", "Tiles cataclysmdda-0.C-8020"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8021.tar.gz", "Tiles cataclysmdda-0.C-8021"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8022.tar.gz", "Tiles cataclysmdda-0.C-8022"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8023.tar.gz", "Tiles cataclysmdda-0.C-8023"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8024.tar.gz", "Tiles cataclysmdda-0.C-8024"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8025.tar.gz", "Tiles cataclysmdda-0.C-8025"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8026.tar.gz", "Tiles cataclysmdda-0.C-8026"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8027.tar.gz", "Tiles cataclysmdda-0.C-8027"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8028.tar.gz", "Tiles cataclysmdda-0.C-8028"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8029.tar.gz", "Tiles cataclysmdda-0.C-8029"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8030.tar.gz", "Tiles cataclysmdda-0.C-8030"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8031.tar.gz", "Tiles cataclysmdda-0.C-8031"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8032.tar.gz", "Tiles cataclysmdda-0.C-8032"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8033.tar.gz", "Tiles cataclysmdda-0.C-8033"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8034.tar.gz", "Tiles cataclysmdda-0.C-8034"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8035.tar.gz", "Tiles cataclysmdda-0.C-8035"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8036.tar.gz", "Tiles cataclysmdda-0.C-8036"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8037.tar.gz", "Tiles cataclysmdda-0.C-8037"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8038.tar.gz", "Tiles cataclysmdda-0.C-8038"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8039.tar.gz", "Tiles cataclysmdda-0.C-8039"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8040.tar.gz", "Tiles cataclysmdda-0.C-8040"},
	{"http://dev.narc.ro/cataclysm/jenkins-latest/Linux_x64/Tiles/cataclysmdda-0.C-8041.tar.gz", "Tiles cataclysmdda-0.C-8041"},
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
