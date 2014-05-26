package ghostman

import "fmt"

type Client interface {
	Message() *Message
	Send(*Message) error
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

func (m *Message) Send() error {
	return m.client.Send(m)
}

func (m *Message) String() string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s\n\n%s\n\n",
		fmt.Sprintf("From: %s", m.from),
		fmt.Sprintf("Reply-To: %s", m.replyTo),
		fmt.Sprintf("To: %s", m.to),
		fmt.Sprintf("Subject: %s", m.subject),
		"Mime-Version: 1.0",
		"Content-Type: text/plain; charset=utf-8",
		m.body,
	)
}
