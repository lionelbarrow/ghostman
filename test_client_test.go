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

func TestMessageString(t *testing.T) {
	client := NewTestClient()

	from := "sender@test.com"
	to := "receiver@test.com"
	replyTo := "reply-to@test.com"
	subject := "Test Subject!"
	body := "Test body!"

	message := client.Message().
		From(from).
		To(to).
		ReplyTo(replyTo).
		Subject(subject).
		Body(body)

	expected := "From: sender@test.com\nReply-To: reply-to@test.com\nTo: receiver@test.com\nSubject: Test Subject!\nMime-Version: 1.0\nContent-Type: text/plain; charset=utf-8\n\nTest body!\n\n"

	if message.String() != expected {
		t.Fatalf("Got incorrect message. Expected:\n%s\nReceived:\n%s", expected, message.String())
	}
}
