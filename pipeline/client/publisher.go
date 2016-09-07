package client

import (
	"encoding/json"
	"fmt"

	"github.com/naveego/api/types/pipeline"
)

func (cli *Client) GetPublisher(publisherId string) (pipeline.RepositoryPublisher, error) {
	var publisher pipeline.RepositoryPublisher

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
