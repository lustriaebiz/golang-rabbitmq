package main

import (
	"adb/config"
	"fmt"
	"log"
)

func main() {

	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")

}
