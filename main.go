package main

import (
	"fmt"
	"os"
)

func main() {
	dbn, ok := os.LookupEnv("EXCHANGEGO_DB_DSN")
	if !ok {
		fmt.Println("Not ok")
	}
	fmt.Println(dbn)
}
