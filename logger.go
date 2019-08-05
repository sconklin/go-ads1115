package ads

import logger "github.com/sconklin/go-logger"

// You can manage verbosity of log output
// in the package by changing last parameter value.
var logads = logger.NewPackageLogger("ads1115",
	// logger.DebugLevel,
	logger.InfoLevel,
)
