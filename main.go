package main

import (
	"bash/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Start point.")

	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
