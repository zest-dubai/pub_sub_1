# Pub_Sub
## Installation
First of all, you need to install docker, docker-compose, kafka and go in your respective machine. Download the given folder at \\\. 

## Run
Open the command prompt in same folder you have downloaded the docker-compose.yml file and type and hit enter.
```bash
  MY_IP=your-ip docker-compose up 
```
Open new window and create a topic manually named 'foo'with replication-factor 2 and no. of partitions 4.
```bash
  docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic foo --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181
```
Now build your dockerfile by using command
```bash
  docker build -t another_email .
```
followed by
```
  docker build -it another_email .
```
