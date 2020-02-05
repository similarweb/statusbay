package visibility

// SetupLogging is set the application log level and ship the logs outside
func SetupLogging(logLevel, gelfAddress, mode string) {

	SetLoggingLevel(logLevel)
	if gelfAddress != "" {
		ShipLogging(mode, gelfAddress)
	}

}
