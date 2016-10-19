package client

import (
	"encoding/json"

	"github.com/naveego/api/types/pipeline"
)

func (cli *Client) GetSubscriber(subscriberID string) (pipeline.SubscriberInstance, error) {
	var subscriber pipeline.SubscriberInstance
	resp, err := cli.get("/subscribers/"+subscriberID, nil)
	if err != nil {
		return subscriber, err
	}

	err = json.NewDecoder(resp.body).Decode(&subscriber)
	if err != nil {
		return subscriber, err
	}

	return subscriber, nil
}

func (cli *Client) UpdateSubscriber(subscriber pipeline.SubscriberInstance) error {
	_, err := cli.put("/pipeline/subscribers/"+subscriber.ID, subscriber, nil)
	if err != nil {
		return err
	}

	return nil
}
