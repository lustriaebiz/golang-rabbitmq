package main

import (
	"adb/config"
	"adb/services"
	"adb/utils"
	"context"
	"fmt"
	"log"
	"net/http"
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

	http.HandleFunc("/checker-log", checkerLog)

	err := http.ListenAndServe(":2400", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func checkerLog(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		checker_log, err := services.CheckerLogAll(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, checker_log, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return

}
