package sms

import (
  "strings"
  "math/rand"
  "time"
  "net/http"
  "net/url"
  "encoding/json"
  "errors"
)

func Send(phone string, transaction_id string)(err error )  {

  accountSid := "AC48b2fc2591dcbdcf78e6543339cfb4c9"
  authToken := "975b63afb08c4027786adf9f56c52b24"
  urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

  quotes := "message api working"

  rand.Seed(time.Now().Unix())


  msgData := url.Values{}
  msgData.Set("To","+919837681431")
  msgData.Set("From","+12058431868")
  msgData.Set("Body",quotes)
  msgDataReader := *strings.NewReader(msgData.Encode())

  client := &http.Client{}
  req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
  req.SetBasicAuth(accountSid, authToken)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  resp, _ := client.Do(req)

  if (resp.StatusCode >= 200 && resp.StatusCode < 300) {

    var data map[string]interface{}
    decoder := json.NewDecoder(resp.Body)
    er := decoder.Decode(&data)

    if (er != nil) {

        return er;
    }

  } else {
    
    return errors.New("error");
    
  }
   
  return nil

}