package verbosity

//go:generate stringer -type=Verbosity -linecomment

type Verbosity int

const (
	Quiet Verbosity = iota - 1 // quiet
	Norm                       // norm
	Info                       // info
	Debug                      // debug
	Trace                      // trace
)
