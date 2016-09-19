package client

import (
	"encoding/json"

	"github.com/naveego/api/types/pipeline"
)

func (cli *Client) GetSubscriber(subscriberID string) (pipeline.SubscriberInstance, error) {
	var subscriber pipeline.SubscriberInstance
	resp, err := cli.get("/pipeline/subscribers/"+subscriberID, nil)
	if err != nil {
		return subscriber, err
	}

	err = json.NewDecoder(resp.body).Decode(&subscriber)
	if err != nil {
		return subscriber, err
	}

	return subscriber, nil
}
