package logger

import "os"

type Options struct {
	EncoderJson bool
	LogFile     string
	Development bool
}

var DefaultOptions = Options{
	EncoderJson: false,
	LogFile:     "",
	Development: false,
}

// init from env
func init() {
	env, ok := os.LookupEnv("LOG_OUTPUT_FILE")
	if ok {
		DefaultOptions.LogFile = env
	}
}
