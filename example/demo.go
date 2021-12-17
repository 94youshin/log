package main

import "github.com/youshintop/log"

func main() {
	opts := &log.Options{
		Level:        "debug",
		Format:       "console",
		EnableColor:  false,
		EnableCaller: true,

		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"}}

	defer log.Flush()
	_, err := log.New(opts)
	if err != nil {
		panic(err)
	}

	log.Debug("This is a debug log message.")
	log.Debugf("This is a %s log message.", "debugf")
	log.Debugw("This is a debugw log message", log.String("id", "001"))

	log.Info("This is a info log message.")
	log.Warn("This is a warn log message.")

	log.Error("This is a error log message.")

	log.Panic("This is a panic log message.")
	log.Fatal("This is a fatal log message.")

}
