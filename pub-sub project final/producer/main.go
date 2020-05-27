package main

import(
	"fmt"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"producer/kafka"
	"producer/requestbody"
)


func main(){

	router := gin.Default()
	
	router.POST("/api/v1/produce", func(c *gin.Context) {
	  
		var body requestbody.RequestBody
		c.BindJSON(&body)
          
        if (body.Message_body=="successfull") {
			
			c.JSON(200, gin.H{
			"status":  "Transaction Successful",
			"email":body.Email,	
			"phone":body.Phone,
		    })

			topic:=body.Topic_name;
			
			data,_:=json.Marshal(body);
			
			fmt.Println(topic);

			 _,er:=kafka.KafkaConnect(topic);
			 fmt.Println("connect error",er);

				
			err := kafka.Push(context.Background(), nil, []byte(data))
			
			fmt.Println("push error ",err);
             

		}else{
			    c.JSON(200, gin.H{
				"Status":"Your transaction failed",
			  })
		}


	})

	router.Run(":8080")
}
