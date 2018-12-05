package main

var rabbitServer string
var emailRequestQueue string
var emailResponseQueue string

func main() {

	//RabbitMq server receiver
	rabbitServer = "amqp://guest:guest@localhost:5672/"
	emailRequestQueue = "emailRequestQueue"
	//RabbitMq server receiver
	emailResponseQueue = "emailResponseQueue"

	//
	receiveMails()
}
