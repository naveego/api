package sub

import (
	"github.com/naveego/api/pipeline/subscriber"
	"github.com/naveego/api/types/pipeline"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	subscriberID  string
	subscriberRef pipeline.SubscriberInstance
)

func init() {
	subscribeCmd.Flags().StringVar(&subscriberID, "subscriberid", "", "The ID of the subscriber")
}

var subscribeCmd = &cobra.Command{
	Use:     "subscribe",
	Short:   "Subscribes to the data from the Naveego Pipeline API",
	PreRunE: runPreSubscribe,
	RunE:    runSubscribe,
}

func runPreSubscribe(cmd *cobra.Command, args []string) error {
	var err error
	subscriberRef, err = apiClient.GetSubscriber(subscriberID)
	if err != nil {
		logrus.Warn("Error Fetching Subscriber From API: ", err)
	}
	log = logrus.WithFields(logrus.Fields{
		"repository": subscriberRef.Repository,
		"pipeline": map[string]interface{}{
			"subscriber_id": subscriberID,
		},
	})
	return err
}

func runSubscribe(cmd *cobra.Command, args []string) error {
	subFactory, err := subscriber.GetFactory(TypeName)
	if err != nil {
		return err
	}

	s := subFactory()
	ctx := subscriber.Context{
		Logger:     log,
		Subscriber: subscriberRef,
	}
	if initer, ok := s.(subscriber.Initer); ok {
		log.Info("Initializing Subscriber")
		log.Debugf("Subscriber Settings: %v", subscriberRef.Settings)
		initer.Init(ctx)
	}

	ctx.Logger = log

	log.Debugf("Setting Up Stream Reader: %s %s", subscriberRef.StreamEndpoint, subscriberRef.InputStream)

	streamReader, err := subscriber.NewStreamReader(subscriberRef.StreamEndpoint, subscriberRef.InputStream, "")
	if err != nil {
		log.Error("Error creating stream reader: ", err)
		return err
	}

	for msg := range streamReader.Messages() {
		shapeInfo := subscriber.GenerateShapeInfo(subscriberRef.Shapes, msg.DataPoint)

		if shapeInfo.HasChanges() {
			err := apiClient.UpdateSubscriber(subscriberRef)
			if err != nil {
				log.Warn("Could not save subscriber changes to API", err)
			}
		}

		s.Receive(ctx, shapeInfo, msg.DataPoint)
	}
	return nil
}
