package client

import (
	"encoding/json"
	"fmt"

	"github.com/naveego/api/types/pipeline"
)

func (cli *Client) GetSubscriber(subscriberID string) (pipeline.SubscriberInstance, error) {
	var subscriber pipeline.SubscriberInstance

	resourceURL := fmt.Sprintf("/subscriber/%s", subscriberID)

	resp, err := cli.sendRequest("GET", resourceURL, nil, nil)
	if err != nil {
		return subscriber, err
	}

	err = json.NewDecoder(resp.body).Decode(&subscriber)
	if err != nil {
		return subscriber, err
	}

	return subscriber, nil
}
