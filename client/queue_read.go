package client

import (
	"encoding/json"

	"github.com/naveego/api/types/queue"
)

type queueMessagesResponse struct {
	Data []queue.Message `json:"data"`
}

func (cli *Client) ReadQueueMessages(queueID string) ([]queue.Message, error) {
	messages := []queue.Message{}

	r, err := cli.get("/queues/"+queueID+"/messages", nil)
	if err != nil {
		return messages, err
	}

	qResp := queueMessagesResponse{}
	err = json.NewDecoder(r.body).Decode(&qResp)
	if err != nil {
		return messages, err
	}

	return qResp.Data, nil
}
