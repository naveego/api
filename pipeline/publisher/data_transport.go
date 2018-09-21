package publisher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/naveego/api/types/pipeline"
	"github.com/sirupsen/logrus"
)

// DataTransport defines an API for sending pipeline data points
// to a destination.
type DataTransport interface {

	// Send will transport the data points to the respective
	// destination.  If anything causes the data points to not
	// reach their destination successfully an error is returned.
	Send(dataPoints []pipeline.DataPoint) error
}

type defaultTransport struct {
	pipelineURL string
	apiToken    string
	repository  string
	source      string
	httpClient  *http.Client
	log         *logrus.Entry
}

// NewDataTransport creates a new instance of the default data transport
// for delivering data points to the pipeline.
func NewDataTransport(pipelineURL, apiToken string, log *logrus.Entry) DataTransport {
	return &defaultTransport{
		pipelineURL: pipelineURL,
		apiToken:    apiToken,
		httpClient:  &http.Client{},
		log:         log,
	}
}

func (dt *defaultTransport) Send(dataPoints []pipeline.DataPoint) error {

	publishURL := fmt.Sprintf("%s/publish", dt.pipelineURL)

	messageBytes, err := json.Marshal(&dataPoints)
	if err != nil {
		return err
	}

	dt.log.Debugf("Publishing data points to %s", publishURL)

	reqBody := bytes.NewReader(messageBytes)
	req, err := http.NewRequest("POST", publishURL, reqBody)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+dt.apiToken)

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return fmt.Errorf("The API returned HTTP Status %d", resp.StatusCode)
	}

	return nil

}
