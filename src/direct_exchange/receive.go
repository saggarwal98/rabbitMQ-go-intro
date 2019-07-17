package main
import (
	"log"
    "github.com/streadway/amqp"
	// "os"
	"time"
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


	
	
	//Declaring a queue
	q, err := ch.QueueDeclare(
		"queue_2", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")


	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")




	//declaring binding
	err = ch.QueueBind(
		q.Name, // queue name
		"rkey1",     // routing key
		"exchange_2", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")
	



	//Consumer consuming messages from queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")  
	
	
	
	forever := make(chan bool)


	//go routine for displaying values
	go func() {
		counter:=1
		for d := range msgs {
			log.Printf("%d : %s",counter, d.Body)
			time.Sleep(3 * time.Second)
			counter++
		}
		log.Printf("Done")
	}()  
	log.Printf("To exit press CTRL+C")

	<-forever
}