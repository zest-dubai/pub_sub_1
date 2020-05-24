package mail

import (
	"net/smtp"
)


func Send(email string,transaction_id string)(er error) {
	
	from := "gaurav.sharma.civ17@itbhu.ac.in"
	password := "12345678"
	to := email	
	msg:="Subject: Transaction Successfull\n\n" +
	 	"Your transaction has been completed.\nTransaction ID:"+transaction_id;
	
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	return err;
	

}
