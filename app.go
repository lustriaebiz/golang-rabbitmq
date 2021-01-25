package main

import (
	"adb/config"
	"adb/services"
	"adb/utils"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
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
	http.HandleFunc("/publish", publish)

	err := http.ListenAndServe(":2400", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func publish(w http.ResponseWriter, r *http.Request) {
	var ch, err = config.RabbitMQ()

	if err != nil {
		fmt.Println(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueueGolang",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(q)

	if err != nil {
		fmt.Println(err)
	}

	err = ch.Publish(
		"",
		"TestQueueGolang",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("lorem ipsum"),
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	res := map[string]string{
		"status": "Successfully Published Message to Queue",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
	return
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
