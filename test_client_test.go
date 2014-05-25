package ghostman

import (
	"testing"
)

func TestTestClient(t *testing.T) {
	client := NewTestClient()

	from := "sender@test.com"
	to := "receiver@test.com"
	replyTo := "reply-to@test.com"
	subject := "Test Subject!"
	body := "Test body!"

	client.Message().
		From(from).
		To(to).
		ReplyTo(replyTo).
		Subject(subject).
		Body(body).
		Send()

	if len(client.Sent()) != 1 {
		t.Fatalf("Got incorrect sent queue length: %d", len(client.Sent()))
	}

	message := client.Sent()[0]

	if message.From != from {
		t.Fatalf("Got incorrect from: %s", message.From)
	} else if message.To != to {
		t.Fatalf("Got incorrect to: %s", message.To)
	} else if message.ReplyTo != replyTo {
		t.Fatalf("Got incorrect reply to: %s", message.ReplyTo)
	} else if message.Subject != subject {
		t.Fatalf("Got incorrect subject: %s", message.Subject)
	} else if message.Body != body {
		t.Fatalf("Got incorrect body: %s", message.Body)
	}
}
