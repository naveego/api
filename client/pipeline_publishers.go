package client

import (
	"encoding/json"

	"github.com/naveego/api/types/pipeline"
)

func (cli *Client) GetPublisherInstance(publisherID string) (pipeline.PublisherInstance, error) {
	var publisher pipeline.PublisherInstance
	resp, err := cli.get("/publishers/"+publisherID, nil)
	if err != nil {
		return publisher, err
	}

	err = json.NewDecoder(resp.body).Decode(&publisher)
	if err != nil {
		return publisher, err
	}

	return publisher, nil
}
