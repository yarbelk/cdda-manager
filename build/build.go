package build

// Build is for OS/Arch
type Build int

// only 64 bit for now
const (
	LinuxX64 Build = iota
	Darwin
	Windows
	WindowsX64
	Unsupported
)
