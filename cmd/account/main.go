package main

import (
	"flag"
	"fmt"
)

var configFile = flag.String("f", "account.yml", "set config file which viper will loading.")

func main() {
	flag.Parse()
	app, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(app)
	// if err := app.Start(); err != nil {
	// 	panic(err)
	// }

	// app.AwaitSignal()
}
