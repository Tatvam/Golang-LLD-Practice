package main

import "./logger"

func main() {
	var logger = logger.GetInstance()
	logger.Warn("Something fishy is going on")
	logger.Info("Everything is going on perfect")
	logger.Error("Alert!! Something bad happened")

}
