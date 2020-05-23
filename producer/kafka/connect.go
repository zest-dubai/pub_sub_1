package kafka

import (
	"time"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

var writer *kafka.Writer

func Configure(kafkaBrokerUrls []string, clientId string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientId,
	}

	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	w = kafka.NewWriter(config)
	writer = w
	return w, nil
}



func KafkaConnect()(writer *kafka.Writer, errr error){

	kafkaBrokersUrls:=[]string{"kafka-1:19092","kafka-2:29092" , "kafka-3:39092" }
	var clientId string ="first_consumer"
	var foo string = "foo"

	var w,err = Configure(kafkaBrokersUrls,clientId,foo);
	return w, err;
	
}