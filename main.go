package main

import (
	//"uber_eats/cmd/uberSync"
	"uber_eats/cmd/uberSync"
	"uber_eats/config"
	"uber_eats/pkg/dbparser"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	logger := uberSync.SetupLogger()
	logger.Warnln("Hello Uber Eats")
	//cfg, _ := json.Marshal(config.CONFIG)

	err := config.LoadConfig()

	if err != nil {
		logger.Fatalf("Error loading config: %v", err)
	}
	logger.Warnln(config.CONFIG)
	dbp := dbparser.NewDbParser(config.CONFIG.DBConfig, *logger)
	dbp.Parse()
	//	https://devs.upcitemdb.com/ for lookup
}
