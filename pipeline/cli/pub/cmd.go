package pub

import (
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/naveego/api/client"
	"github.com/naveego/api/types/pipeline"
	"github.com/spf13/cobra"
)

const apiClientVersion = "4.0.0-alpha"

var (
	// TypeName holds the name of the connector being used in this package
	TypeName = "none"

	targetURL         string
	apiClient         *client.Client
	apitoken          string
	verbose           bool
	log               *logrus.Entry
	publisherInstance pipeline.PublisherInstance
)

var RootCmd = &cobra.Command{
	Use:   "",
	Short: "Publishes data to Naveego Pipeline",
	RunE:  nil,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}

		bearerStr := fmt.Sprintf("Bearer %s", apitoken)

		apiClientHeaders := map[string]string{
			"Authorization": bearerStr,
		}

		targetURL = strings.TrimSpace(targetURL)

		var err error
		apiClient, err = client.NewClient(targetURL, apiClientVersion, apiClientHeaders)
		if err != nil {
			return err
		}

		publisherID := args[0]
		publisherInstance, err = apiClient.GetPublisherInstance(publisherID)
		if err != nil {
			return err
		}

		log = logrus.WithField("publisher_id", publisherID)
		return nil
	},
}

func init() {
	RootCmd.SilenceUsage = true
	RootCmd.PersistentFlags().StringVarP(&targetURL, "api", "a", "", "The url for the pipeline api")
	RootCmd.PersistentFlags().StringVarP(&apitoken, "token", "t", "", "The API token to use for authentication")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Turn on verbose logging")
}

// Execute is the main entry command for the package.  It creates a
// Cobra root command, adds all sub-commands, then executes them.
func Execute() error {

	addCommands()

	err := RootCmd.Execute()
	return err
}

func addCommands() {
	RootCmd.AddCommand(shapesCmd)
	RootCmd.AddCommand(publishCmd)
	RootCmd.AddCommand(runCmd)
}
