package ghostman

type Client interface {
	Message() *Message
	Send(*Message)
}

type Message struct {
	from    string
	replyTo string
	to      string
	subject string
	body    string
	client  Client
}

func (m *Message) From(address string) *Message {
	m.from = address
	return m
}

func (m *Message) ReplyTo(address string) *Message {
	m.replyTo = address
	return m
}

func (m *Message) To(address string) *Message {
	m.to = address
	return m
}

func (m *Message) Subject(subject string) *Message {
	m.subject = subject
	return m
}

func (m *Message) Body(body string) *Message {
	m.body = body
	return m
}

func (m *Message) Send() {
	m.client.Send(m)
}
