package main

import (
	"fmt"
	"encoding/json"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"email/mail"
	"email/requestbody"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
	    //"bootstrap.servers": "localhost:19092",
		"bootstrap.servers": "kafka-1:19092,kafka-2:29092,kafka-3:39092",
		"group.id":          "email",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(c);

	c.SubscribeTopics([]string{"foo", "^aRegex.*[Tt]opic"}, nil)

	for {

		msg, err := c.ReadMessage(-1)

		data:=string(msg.Value);

		var body requestbody.RequestBody;

		json.Unmarshal([]byte(data), &body)
		
		email:=body.Email;
		transaction_id:=body.Transaction_id

		fmt.Println("\n\n");

		if err == nil {

    		errorr:= mail.Send(email,transaction_id)		
			fmt.Printf("sent",errorr, email)
			
		} else {

			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()

}
