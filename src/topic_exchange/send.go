package main
import (
  "log"
  "github.com/streadway/amqp"
  "os"
//   "fmt"
//   "strings"
)
func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}

func main(){

	
	//creating connection to rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	
	
	//creating channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()


	// //Declaring a queue
	// q, err := ch.QueueDeclare(
	// 	"durable_queue_exchange", // name
	// 	true,   // durable
	// 	false,   // delete when unused
	// 	false,   // exclusive
	// 	false,   // no-wait
	// 	nil,     // arguments
	// )
	// failOnError(err, "Failed to declare a queue")



	//getting input from command line arguments
	arg:=os.Args
	body:=""
	for i:=0; i < len(arg[2:]); i++ {
		body=body+" "+arg[i+2]
	}
	


	//defining custom exchange
	err = ch.ExchangeDeclare(
		"exchange_3",   // name
		"topic", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)



	//declaring binding
	// err = ch.QueueBind(
	// 	q.Name, // queue name
	// 	"",     // routing key
	// 	"custom_exchange", // exchange
	// 	false,
	// 	nil,
	// )
	// failOnError(err, "Failed to bind a queue")



	//Sending data to Queue
	err = ch.Publish(
		"exchange_3",     // exchange
		arg[1], // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,
		    ContentType: "text/plain",
		    Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	

	forever := make(chan bool)
	<-forever
}