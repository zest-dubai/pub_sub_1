package main

import (
	"fmt"
	"sync"
	"encoding/json"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"consumer/mail"
	"consumer/sms"
	"consumer/requestbody"
)

func parse_data(data string)(string,string){

	var body requestbody.RequestBody;
	json.Unmarshal([]byte(data), &body);
	phone:=body.Phone;
	email:=body.Email
	
	return email,phone;
	
}


func main() {

    email_consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
	    "bootstrap.servers": "localhost:19092",
		//"bootstrap.servers": "kafka-1:19092,kafka-2:29092,kafka-3:39092",
		"group.id":          "email",
		"auto.offset.reset": "earliest",
		"enable.auto.commit": false,

	})

	if err != nil {
		panic(err)
	}

	sms_consumer,er := kafka.NewConsumer(&kafka.ConfigMap{
	    "bootstrap.servers": "localhost:19092",
		//"bootstrap.servers": "kafka-1:19092,kafka-2:29092,kafka-3:39092",
		"group.id":          "sms",
		"auto.offset.reset": "earliest",
		"enable.auto.commit": false,

	})

	if er != nil {
		panic(er)
	}

	email_consumer.SubscribeTopics([]string{"foo", "^aRegex.*[Tt]opic"}, nil)
	sms_consumer.SubscribeTopics([]string{"foo", "^aRegex.*[Tt]opic"}, nil)

	var wg sync.WaitGroup
	
	fmt.Println("running");

		wg.Add(1)
        go func(){
			   
			defer wg.Done();
			for{
               
				msg,e:=email_consumer.ReadMessage(-1);

				if e == nil {

					email,_:=parse_data(string(msg.Value));

					errorr:= mail.Send(email);	
					
					if(errorr== nil ){
						email_consumer.Commit();
						fmt.Printf("email sent ");

					}else{
					fmt.Printf("email not sent ");
					}	
					
				} else {

					fmt.Printf("email consumer error")
				}

			}

		}()
	   

		wg.Add(1)
		go func(){
			defer wg.Done();

			for{
			    
				message,ee:=sms_consumer.ReadMessage(-1);
				
				if ee == nil {

					_,phone:=parse_data(string(message.Value));

					errorr:= sms.Send(phone);	
					
					if(errorr == nil ){
						sms_consumer.Commit();
						fmt.Printf("sms sent ");

					}else{
											
					fmt.Printf("\n sms not sent \n");
					}	
					
				} else {

					fmt.Printf("sms consumer error\n")
				}  
				
			}

		} ()


		wg.Wait();

	email_consumer.Close();
	sms_consumer.Close();


}
