package subscriber

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/Shopify/sarama"
	"github.com/Sirupsen/logrus"
	"github.com/naveego/api/types/pipeline"
	"github.com/wvanbergen/kafka/consumergroup"
	"github.com/wvanbergen/kazoo-go"
)

type StreamMessage struct {
	DataPoint  pipeline.DataPoint
	RawMessage interface{}
}

type StreamReader interface {
	Messages() <-chan StreamMessage
	CommitUpTo(message StreamMessage)
	Close() error
}

type defaultStreamReader struct {
	messages chan StreamMessage
	consumer *consumergroup.ConsumerGroup
	log      *logrus.Entry
}

func NewStreamReader(addr, stream, readerID string) (StreamReader, error) {
	return NewStreamReaderWithLogging(addr, stream, readerID, nil)
}

func NewStreamReaderWithLogging(addr, stream, readerID string, log *logrus.Entry) (StreamReader, error) {

	// Bufferred channel to hold incoming datapoints
	messages := make(chan StreamMessage, 100)

	config := consumergroup.NewConfig()
	config.Zookeeper.Timeout = 15 * time.Second
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 15 * time.Second

	var zkNodes []string
	zkNodes, config.Zookeeper.Chroot = kazoo.ParseConnectionString(addr)

	consumer, err := consumergroup.JoinConsumerGroup(readerID, []string{stream}, zkNodes, config)
	if err != nil {
		return nil, err
	}

	reader := &defaultStreamReader{
		messages: messages,
		consumer: consumer,
		log:      log,
	}

	go reader.messageLoop()

	if log != nil {
		go reader.errorLoop()
	}

	return reader, nil
}

func (sr *defaultStreamReader) Messages() <-chan StreamMessage {
	return sr.messages
}

func (sr *defaultStreamReader) CommitUpTo(message StreamMessage) {
	consumerMsg, ok := message.RawMessage.(*sarama.ConsumerMessage)
	if !ok {
		if sr.log != nil {
			sr.log.Error("Commit offset error: expected raw message of type *sarama.ConsumerMessage")
		}
		return
	}

	sr.consumer.CommitUpto(consumerMsg)
}

func (sr *defaultStreamReader) Close() error {
	return sr.consumer.Close()
}

func (sr *defaultStreamReader) errorLoop() {
	for err := range sr.consumer.Errors() {
		sr.log.Error("Error reading message from stream: ", err)
	}
}

func (sr *defaultStreamReader) messageLoop() {
	for msg := range sr.consumer.Messages() {
		var dataPoint pipeline.DataPoint

		ioutil.WriteFile("Message.json", msg.Value, 777)

		err := json.Unmarshal(msg.Value, &dataPoint)
		if err != nil {
			panic(err)
		}

		sr.messages <- StreamMessage{
			DataPoint:  dataPoint,
			RawMessage: msg,
		}
	}
}
