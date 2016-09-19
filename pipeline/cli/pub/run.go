package pub

import (
	"os"
	"os/signal"
	"strings"

	"github.com/naveego/api/live"
	"github.com/naveego/api/pipeline/publisher"
	"github.com/robfig/cron"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a Naveego pipeline publisher",
	RunE:  runRun,
}

func runRun(cmd *cobra.Command, args []string) error {
	pubFactory, err := publisher.GetFactory(TypeName)
	if err != nil {
		return err
	}

	scheduler := cron.New()
	scheduler.Start()
	defer scheduler.Stop()

	host, _ := os.Hostname()
	var liveCli *live.Client

	if strings.HasPrefix(publisherInstance.LiveEndpoint, "tcp://") {
		tcpURL := publisherInstance.LiveEndpoint[6:]
		liveCli, err = live.NewTCPClient(tcpURL, publisherInstance.ID, host)
	} else if strings.HasPrefix(publisherInstance.LiveEndpoint, "ws://") || strings.HasPrefix(publisherInstance.LiveEndpoint, "wss://") {
		liveCli, err = live.NewWebSocketClient(publisherInstance.LiveEndpoint, publisherInstance.ID, host)
	}
	if err != nil {
		return err
	}

	go liveRead(liveCli)

	log.Infof("Scheduling publisher with schedule: %s", publisherInstance.Schedule)

	err = scheduler.AddFunc(publisherInstance.Schedule, func() {

		ctx := publisher.Context{
			Logger:            log,
			PublisherInstance: publisherInstance,
		}

		transport := publisher.NewDataTransport(publisherInstance.PublishEndpoint, apitoken, log)
		p := pubFactory()
		p.Publish(ctx, transport)

	})

	if err != nil {
		return err
	}

	log.Info("Successfully scheduled publisher")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs

	return nil
}

func liveRead(client *live.Client) {
	for msg := range client.Incoming() {
		log.Debugf("Received message: %s", string(msg.Content))
	}
}
