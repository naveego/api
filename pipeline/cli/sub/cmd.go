package sub

import (
	"fmt"
	"strings"

	"github.com/naveego/api/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const apiClientVersion = "4.0.0-alpha"

var (
	// TypeName holds the name of the connector being used in this package
	TypeName = "none"

	targetURL  string
	repository string
	apiClient  *client.Client
	apitoken   string
	log        *logrus.Entry
	verbose    bool
)

var RootCmd = &cobra.Command{
	Use:   "",
	Short: "Executes subscriber commands",
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

		return nil
	},
}

func init() {
	RootCmd.SilenceUsage = true
	RootCmd.PersistentFlags().StringVarP(&targetURL, "url", "u", "", "The url for the pipeline api")
	RootCmd.PersistentFlags().StringVarP(&apitoken, "token", "t", "", "The API token to use for authentication")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
}

// Execute is the main entry command for the package.  It creates a
// Cobra root command, adds all sub-commands, then executes them.
func Execute() error {

	addCommands()

	err := RootCmd.Execute()
	return err
}

func addCommands() {
	RootCmd.AddCommand(subscribeCmd)
}
