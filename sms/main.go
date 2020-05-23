package main

import (
	"fmt"
	"encoding/json"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"sms/requestbody"
	"sms/message"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		//"bootstrap.servers": "kafka-1:19092,kafka-2:29092,kafka-3:39092",
		"bootstrap.servers": "localhost:19092,localhost:29092,localhost:39092",
		"group.id":          "sms",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"foo", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		
		data:=string(msg.Value);

		var body requestbody.RequestBody;

		json.Unmarshal([]byte(data), &body)
		
		phone:=body.Phone;
		transaction_id:=body.Transaction_id

		fmt.Println("\n\n");

		if err == nil {
			
			errorr:= sms.Send(phone,transaction_id)
			fmt.Printf("sms sent",errorr, phone)
			
		} else {

			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()

}