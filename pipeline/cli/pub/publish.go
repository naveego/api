package pub

import (
	"github.com/naveego/api/pipeline/publisher"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publishes data to the Naveego pipeline API",
	RunE:  runPublish,
}

func runPublish(cmd *cobra.Command, args []string) error {
	pubFactory, err := publisher.GetFactory(TypeName)
	if err != nil {
		return err
	}

	ctx := publisher.Context{
		Logger:            log,
		PublisherInstance: publisherInstance,
	}

	transport := publisher.NewDataTransport(apiURL, apitoken, log)
	p := pubFactory()
	p.Publish(ctx, transport)
	return nil
}
