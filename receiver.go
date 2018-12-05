package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

type Mail struct {
	Name string
	Mail string
}

type EmailMessage  struct {
	Users []User
	Event string
}

type ConfirmationQueue struct {
	Type    string
	Message string
	Event string
}

type User struct {
	ID        int
	Name      string
	Firstname string
	Lastname  string
	Email     string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func receiveMails() {
	conn, err := amqp.Dial(rabbitServer)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		emailRequestQueue, // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {

			var mails EmailMessage
			errMail := json.Unmarshal(d.Body, &mails)
			if errMail != nil {
				panic(errMail)
			}
			f, err := os.OpenFile("emailLog", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				log.Fatal(err)
			}
			var UsersList []User
			UsersList = mails.Users
			event := mails.Event
			//defer to close when you're done with it, not because you think it's idiomatic!
			defer f.Close()
			for c := range UsersList {
				user := UsersList[c]
				//set output of logs to f
				log.SetOutput(f)
				log.Println("mails sent to " + user.Name + " Mail : " + user.Email +" Event : " +event)
				fmt.Println("mails sent to " + user.Name + " Mail : " + user.Email +" Event : " +event)
			}

			log.Println("INFO : Sent successfully ")

			documentJSON, _ := json.Marshal(ConfirmationQueue{Type: "successfully", Message: " Sent successfully "})
			sendMailMessage(documentJSON)

			d.Ack(false)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	forever := make(chan bool)
	<-forever
}
