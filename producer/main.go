package main

import(
	"fmt"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"main/kafka"
	"main/requestbody"
)


func main(){

	router := gin.Default()
	
	kafka.KafkaConnect();

	router.POST("/api/v1/produce", func(c *gin.Context) {
	  
		var body requestbody.RequestBody
		c.BindJSON(&body)
          
        if (body.Message_body=="successfull") {
			
			c.JSON(200, gin.H{
			"status":  "Transaction Successful",
			"email":body.Email,	
			"phone":body.Phone,
			})

			data,_:=json.Marshal(body);

			err := kafka.Push(context.Background(), nil, []byte(data))
			fmt.Println("the error in producer is ",err);
             

		}else{
			    c.JSON(200, gin.H{
				"Status":"Your transaction failed",
			  })
		}


	})

	router.Run(":8080")
}
