package ghostman

func NewTestClient() *TestClient {
	return &TestClient{}
}

type TestClient struct {
	sentMessages []sentMessage
}

func (t *TestClient) Message() *Message {
	return &Message{client: t}
}

func (t *TestClient) Sent() []sentMessage {
	return t.sentMessages
}

func (t *TestClient) Send(m *Message) error {
	sent := sentMessage{
		From:    m.from,
		To:      m.to,
		ReplyTo: m.replyTo,
		Subject: m.subject,
		Body:    m.body,
	}

	t.sentMessages = append(t.sentMessages, sent)
	return nil
}

type sentMessage struct {
	From    string
	ReplyTo string
	To      string
	Subject string
	Body    string
}
