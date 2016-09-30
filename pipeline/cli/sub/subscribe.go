package sub

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/naveego/api/pipeline/subscriber"
	"github.com/naveego/api/types/pipeline"
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

	streamReader, err := subscriber.NewStreamReader(subscriberRef.StreamEndpoint, subscriberRef.InputStream)
	if err != nil {
		log.Error("Error creating stream reader: ", err)
		return err
	}

	for dataPoint := range streamReader.DataPoints() {
		shapeInfo := generateShapeInfo(subscriberRef, dataPoint)

		if shapeInfo.HasChanges() {
			err := apiClient.UpdateSubscriber(subscriberRef)
			if err != nil {
				log.Warn("Could not save subscriber changes to API", err)
			}
		}

		s.Receive(ctx, shapeInfo, dataPoint)
	}
	return nil
}

// generateShapeInfo will determine the diffferences between an existing shape and the shape of a new
// data point.  If the new shape is a subset of the current shape it is not considered a change.  This
// is due to the fact that it does not represent a change that needs to be made in the storge system.
func generateShapeInfo(sub pipeline.SubscriberInstance, dataPoint pipeline.DataPoint) subscriber.ShapeInfo {
	shape := dataPoint.Shape

	// create the info
	info := subscriber.ShapeInfo{
		Shape: shape,
	}

	// Get the shape if we already know about it
	prevShape, ok := sub.Shapes[dataPoint.Entity]

	// If this shape does not exists previously then
	// we need to treat it as brand new
	if !ok {
		info.IsNew = true
		info.NewKeys = dataPoint.KeyNames
		info.HasNewProperties = true
		info.HasKeyChanges = true
	} else {

		// If the shape is exactly the same as the previous shape, or it is a subset of the previous shape
		// then there is no change.  We can just use the previous shape.
		if shape.PropertyHash != prevShape.PropertyHash && isSubsetOf(shape.Properties, prevShape.Properties) {
			info.Shape = prevShape
		}

		// Check the key names
		if !areSame(info.Shape.KeyNames, prevShape.KeyNames) {
			info.HasKeyChanges = true

			// Load the new keys
			info.NewKeys = []string{}
			for _, key := range info.Shape.KeyNames {
				info.NewKeys = append(info.NewKeys, key)
			}
		}
	}

	// Set the previous shape on the info
	info.PreviousShape = prevShape

	// Load any new properties
	info.NewProperties = subscriber.PropertiesAndTypes{}
	for _, prop := range info.Shape.Properties {
		if !contains(prevShape.Properties, prop) {
			p := strings.Split(prop, ":")
			info.NewProperties[p[0]] = p[1]
		}
	}

	info.HasNewProperties = (len(info.NewProperties) > 0)

	return info
}

// contains is a helper function to determine if a string slice
// contains a string value
func contains(a []string, v string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}

// areSame is a helper function that determines if two slices are
// the same.  Two slices are considered the same if they are the same
// length and contain equal values at the same indexes.
func areSame(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// isSubsetOf is a helper function that determines if one slice
// is a subset of another
func isSubsetOf(list []string, all []string) bool {
	for _, l := range list {
		if !contains(all, l) {
			return false
		}
	}
	return true
}
