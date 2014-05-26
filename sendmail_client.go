package ghostman

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func NewClient(location string) *SendmailClient {
	return &SendmailClient{
		location: location,
	}
}

type SendmailClient struct {
	location string
}

func (s *SendmailClient) Message() *Message {
	return &Message{client: s}
}

func (s *SendmailClient) Send(m *Message) error {
	var out bytes.Buffer
	args := []string{
		fmt.Sprintf("-f %s", m.replyTo),
		"-i",
		"-t",
	}

	cmd := exec.Command(s.location, args...)
	cmd.Stdin = strings.NewReader(m.String())
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	str := out.String()
	if str != "" {
		return fmt.Errorf("%s", out.String())
	}

	return nil
}
