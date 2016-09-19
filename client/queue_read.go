package client

import (
	"encoding/json"
	"fmt"

	"github.com/naveego/api/types/queue"
)

type queueMessagesResponse struct {
	Data []queue.Message `json:"data"`
}

func (cli *Client) ReadQueueMessages(queueID string) ([]queue.Message, error) {
	messages := []queue.Message{}

	resourceURL := fmt.Sprintf("/queues/%s/messages", queueID)

	r, err := cli.sendRequest("GET", resourceURL, nil, nil)
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
