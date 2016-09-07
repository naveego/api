package subscriber

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/naveego/api/types/pipeline"
)

type StreamReader interface {
	DataPoints() <-chan pipeline.DataPoint

	Close() error
}

type defaultStreamReader struct {
	dataPoints        chan pipeline.DataPoint
	consumer          sarama.Consumer
	partitionConsumer sarama.PartitionConsumer
}

func NewStreamReader(addr, stream string) (StreamReader, error) {

	// Bufferred channel to hold incoming datapoints
	dataPoints := make(chan pipeline.DataPoint, 100)

	consumer, err := sarama.NewConsumer(strings.Split(addr, ","), nil)
	if err != nil {
		return nil, err
	}

	partConsumer, err := consumer.ConsumePartition(stream, 0, sarama.OffsetOldest)
	if err != nil {
		return nil, err
	}

	reader := &defaultStreamReader{
		dataPoints:        dataPoints,
		consumer:          consumer,
		partitionConsumer: partConsumer,
	}

	go reader.messageLoop()

	return reader, nil

}

func (sr *defaultStreamReader) DataPoints() <-chan pipeline.DataPoint {
	return sr.dataPoints
}

func (sr *defaultStreamReader) Close() error {
	return sr.consumer.Close()
}

func (sr *defaultStreamReader) messageLoop() {
	for msg := range sr.partitionConsumer.Messages() {
		var dataPoint pipeline.DataPoint

		ioutil.WriteFile("Message.json", msg.Value, 777)

		err := json.Unmarshal(msg.Value, &dataPoint)
		if err != nil {
			panic(err)
		}

		sr.dataPoints <- dataPoint
	}
}
