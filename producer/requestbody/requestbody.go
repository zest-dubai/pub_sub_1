package requestbody

type RequestBody struct{
		
	Request_id string `json: request_id`
	Topic_name string `json:topic_name`
	Message_body string `json:message_body`
	Transaction_id string `json:transaction_id`
	Email string `json:email`
	Phone string `json:phone`
	Customer_id string `json:Customer_id`	
	Key string `json:key`

}