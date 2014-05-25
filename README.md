# Ghostman

Ghostman is an easy-to-use, easy-to-test email client for Go.

### Usage

```go
client := ghostman.NewClient()

client.Message().
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

I'm going to start by only supporting plain-text messages delivered via `sendmail`, but if people are interested I'll add more and more sophistication to the wrapping layer and maybe an `smtp` package wrapper as well.
