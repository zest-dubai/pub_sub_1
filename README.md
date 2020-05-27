# Pub_Sub (refer to the 'pub-sub project final' directory)
## Installation
First of all, you need to install docker, docker-compose and go in your system. 

## Setup
Setup the docker-compose file with kafka and zookeeper, i.e. remove the producer and the consumer image setup from the docker-compose file provided in the repository(as the images have not been built yet)

Open the terminal, move to the working directory and run
```bash
  MY_IP=your-ip docker-compose up 
```
Now our zookeeper and kafka brokers have been created.

The next step would be to create the topic.
We are creating a topic named "foo" with 4 partitions and replication factor as 2.

Open a termial and run
```bash
  docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic foo --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181
```
Now move to the "producer" directory and build the dockerfile of your producer as
```bash
  docker build -t producer .
```
to create its image named 'producer'


Move to the "consumer" directory and build the dockerfile of your consumer as
```bash
  docker build -t consumer .
```
to create its image named 'consumer'

Now add the docker images of the producer and the consumer to the docker-compose file as it is in the repository.

Open the terminal and run the docker-compose file the same way you did the last time. This time with the producer and consumer images added to it.
```bash
  MY_IP=your-ip docker-compose up 
```
All the necessary setup is completed.

Use Postman to send the API requests to the producer service which would then publish the message to kafka, and the consumer will read it to send email and sms to the provided mail id and mobile number respectively.

