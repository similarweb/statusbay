package visibility

// SetupLogging is set the application log level and ship the logs outside
func SetupLogging(logLevel, mode string) {

	SetLoggingLevel(logLevel)
	ShipLogging(mode)

}
