# Pub_Sub
## Installation
First of all, you need to install docker, docker-compose, kafka and go in your respective machine. Download the given folder at \\\. 

## Run
Open a new terminal and create a topic named 'foo'with replication-factor 2 and no. of partitions 4.
```bash
  docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic foo --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181
```

Build your dockerfile of your producer as
```bash
  docker build -t producer
```
to create its image named 'producer'

Build your dockerfile of your consumer as
```bash
  docker build -t consumer
```
to create its image named 'consumer'

Open the command prompt in same folder you have downloaded the docker-compose.yml file and type and hit enter.
```bash
  MY_IP=your-ip docker-compose up 
```


