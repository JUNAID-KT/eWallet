package app

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

// formatter adds default fields to each log entry.
type formatter struct {
	fields log.Fields
	lf     log.Formatter
}

// Format satisfies the logrus.Formatter interface.
func (f *formatter) Format(e *log.Entry) ([]byte, error) {
	for k, v := range f.fields {
		e.Data[k] = v
	}
	return f.lf.Format(e)
}

func InitializeLogger() {
	log.SetLevel(getLogLevel(Config.LogLevel))

	if os.Getenv("Environment") == "production" {
		log.SetFormatter(&formatter{
			fields: log.Fields{
				// Put your default fields here
				"service": "eWallet",
			},
			lf: &log.JSONFormatter{
				PrettyPrint:      true,
				DisableTimestamp: false,
			},
		})
	} else {
		log.SetFormatter(&formatter{
			fields: log.Fields{
				// Put your default fields here
				"service": "eWallet",
			},
			lf: &log.TextFormatter{
				// Put your usual options here
				DisableColors:             false,
				FullTimestamp:             true,
				ForceColors:               true,
				EnvironmentOverrideColors: true,
			},
		})
	}
}

//getLogLevel: Returns the corresponding Logrus logLevel based on the input from the config.toml file
func getLogLevel(logLevel int) log.Level {
	var level log.Level

	switch logLevel {
	case 0:
		level = log.PanicLevel
	case 1:
		level = log.FatalLevel
	case 2:
		level = log.ErrorLevel
	case 3:
		level = log.WarnLevel
	case 4:
		level = log.InfoLevel
	case 5:
		level = log.DebugLevel
	case 6:
		level = log.TraceLevel
	}
	return level
}
