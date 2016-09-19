package client

type ackMessage struct {
	MessageID int64 `json:"message_id"`
}

func (cli *Client) AcknowledgeQueueMessages(messageIDs []int64) error {

	acks := []ackMessage{}

	for _, ID := range messageIDs {
		acks = append(acks, ackMessage{
			MessageID: ID,
		})
	}

	_, err := cli.sendRequest("POST", "/queues/acknowledged", acks, nil)
	return err

}
