package sms

import (
  "strings"
  "math/rand"
  "time"
  "net/http"
  "net/url"
  "encoding/json"
)

func Send(phone string, transaction_id string)(err string )  {
  // Set account keys & information
  accountSid := "AC856d0fdd276d3cea5b974ba80c7ec225"
  authToken := "8644e8fc80be8deb872fec06fb4b4154"
  urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

  // Create possible message bodies
  quotes := [7]string{"I urge you to please notice when you are happy, and exclaim or murmur or think at some point, 'If this isn't nice, I don't know what is.'",
                      "Peculiar travel suggestions are dancing lessons from God.",
                      "There's only one rule that I know of, babies—God damn it, you've got to be kind.",
                      "Many people need desperately to receive this message: 'I feel and think much as you do, care about many of the things you care about, although most people do not care about them. You are not alone.'",
                      "That is my principal objection to life, I think: It's too easy, when alive, to make perfectly horrible mistakes.",
                      "So it goes.",
                      "We must be careful about what we pretend to be."}

  // Set up rand
  rand.Seed(time.Now().Unix())

  // Pack up the data for our message
  msgData := url.Values{}
  msgData.Set("To","+919837681431")
  msgData.Set("From","+12513201344")
  msgData.Set("Body",quotes[rand.Intn(len(quotes))])
  msgDataReader := *strings.NewReader(msgData.Encode())

  // Create HTTP request client
  client := &http.Client{}
  req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
  req.SetBasicAuth(accountSid, authToken)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  // Make HTTP POST request and return message SID
  resp, _ := client.Do(req)
  if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
    var data map[string]interface{}
    decoder := json.NewDecoder(resp.Body)
    err := decoder.Decode(&data)
    if (err == nil) {

     err= nil

    }
  } else {
      err="error"
  }
   
  return err


}