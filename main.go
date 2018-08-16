package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = "shomu2"
	app.Usage = "fight the loneliness!"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
