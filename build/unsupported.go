// +build linux,!amd64
// +build !windows,!amd64 !windows,!386
// +build !darwin

package build

// BuildV is which build is supported by this OS/Arch.  This is for unsupported
const BuildV = Unsupported
