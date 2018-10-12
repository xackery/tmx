package model

var (
	isVerbose = false
)

// SetVerbose sets the client as verbose
func SetVerbose(value bool) {
	isVerbose = value
}

// IsVerbose returns if verbose mode is on
func IsVerbose() bool {
	return isVerbose
}
