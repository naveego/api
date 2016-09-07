package sub

import (
	"github.com/Sirupsen/logrus"
	"github.com/naveego/api/pipeline/subscriber"
	"github.com/naveego/api/types/pipeline"
	"github.com/spf13/cobra"
)

var (
	readerAddr    string
	subscriberID  string
	subscriberRef pipeline.RepositorySubscriber
)

func init() {
	subscribeCmd.Flags().StringVar(&readerAddr, "readeraddr", "127.0.0.1:9092", "The address for the reader to use to read data points")
	subscribeCmd.Flags().StringVar(&subscriberID, "subscriberid", "", "The ID of the subscriber")
}

var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribes to the data from the Naveego Pipeline API",
	RunE:  runSubscribe,
}

func runPreSubscribe(cmd *cobra.Command, args []string) error {
	var err error
	subscriberRef, err = apiClient.GetSubscriber(subscriberID)
	log = logrus.WithField("repository", subscriberRef.Repository)
	return err
}

func runSubscribe(cmd *cobra.Command, args []string) error {
	subFactory, err := subscriber.GetFactory(TypeName)
	if err != nil {
		return err
	}

	s := subFactory()
	ctx := subscriber.Context{}
	if initer, ok := s.(subscriber.Initer); ok {
		initer.Init(ctx)
	}

	streamReader, err := subscriber.NewStreamReader(readerAddr, ctx.Subscriber.InputStream)
	if err != nil {
		return err
	}

	for dataPoint := range streamReader.DataPoints() {
		s.Receive(ctx, dataPoint)
	}
	return err
}
