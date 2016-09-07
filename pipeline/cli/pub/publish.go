package pub

import (
	"github.com/Sirupsen/logrus"
	"github.com/naveego/api/pipeline/publisher"
	"github.com/naveego/api/types/pipeline"
	"github.com/spf13/cobra"
)

var (
	publisherID  string
	publisherRef pipeline.RepositoryPublisher
)

var publishCmd = &cobra.Command{
	Use:     "publish",
	Short:   "Publishes data to the Naveego pipeline API",
	PreRunE: runPrePublish,
	RunE:    runPublish,
}

func init() {
	publishCmd.PersistentFlags().StringVarP(&publisherID, "publisherid", "p", "", "The url for the pipeline api")
}

func runPrePublish(cmd *cobra.Command, args []string) error {
	var err error
	publisherRef, err = apiClient.GetPublisher(publisherID)
	log = logrus.WithField("repository", publisherRef.Repository)
	return err
}

func runPublish(cmd *cobra.Command, args []string) error {
	pubFactory, err := publisher.GetFactory(TypeName)
	if err != nil {
		return err
	}

	ctx := publisher.Context{
		Logger:       log,
		PublisherRef: publisherRef,
	}
	transport := publisher.NewDataTransport(targetURL, apitoken, log)
	p := pubFactory()
	p.Publish(ctx, transport)
	return nil
}
