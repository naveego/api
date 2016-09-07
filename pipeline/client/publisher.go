package client

import (
	"encoding/json"
	"fmt"

	"github.com/naveego/api/types/pipeline"
)

func (cli *Client) GetPublisherInstance(publisherId string) (pipeline.PublisherInstance, error) {
	var publisher pipeline.PublisherInstance

	publisherURL := fmt.Sprintf("/publisher/%s", publisherId)

	resp, err := cli.sendRequest("GET", publisherURL, nil, nil)
	if err != nil {
		return publisher, err
	}

	err = json.NewDecoder(resp.body).Decode(&publisher)
	if err != nil {
		return publisher, err
	}

	return publisher, nil
}
