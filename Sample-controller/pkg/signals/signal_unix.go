package signals

import (
	"os"
	"syscall"
)

var shutdownSignals = []os.Signal{os.Interupt, syscall.SIGTERM}