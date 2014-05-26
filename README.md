# Ghostman

Ghostman is an easy-to-use, easy-to-test email client for Go. It currently is mostly a wrapper around `sendmail`.

### Usage

```go
client := ghostman.NewClient("/usr/sbin/sendmail")

err := client.Message().
  From("sender@test.com").
  To("receiver@test.com").
  ReplyTo("reply-to@test.com").
  Subject("Hello there!").
  Body("Let's get coffee tomorrow.").
  Send()
```

Simple, right? If you want to test that other parts of your program send email correctly, use the provided `TestClient`.

```go
  client := ghostman.NewTestClient()

  DoSomething(client)

  if len(client.Sent()) != 1 {
    t.Fatal("Whoops, didn't send any email!")
  }
```

I'm going to start by only supporting plain-text messages, but if people are interested I'll add more and more sophistication to the wrapping layer. We could also add support for a wrapper around the SMTP standard library package as well.
