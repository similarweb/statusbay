package visibility

import (
	"strings"

	log "github.com/sirupsen/logrus"
	graylog "gopkg.in/gemnasium/logrus-graylog-hook.v2"
)

// ShipLogging configures hooks used to ship logs to ELK.
func ShipLogging(mode, address string) {

	loggingConfig := make(map[string]interface{})
	loggingConfig["application"] = "statusbay"
	loggingConfig["mode"] = mode
	hook := graylog.NewGraylogHook(address, loggingConfig)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.AddHook(hook)
}

// SetLoggingLevel sets the logging level to the specified string
func SetLoggingLevel(level string) {
	level = strings.ToLower(level)
	log.WithFields(log.Fields{"level": level}).Warn("setting logging level")
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn", "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.WithFields(log.Fields{"level": level}).Warn("Invalid log level, not setting")
	}
}
