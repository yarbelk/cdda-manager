// +build windows,386

package build

// BuildV is what is supported by this arch; this is going to break because reasons
// windows runs both amd64 and 386.  People expect 1 binary.
const BuildV = Windows
