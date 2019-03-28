# What's this?

This is a simple exercise that is meant to help teach the principles of testing
code.

It was deliberately not built in a TDD manner, *this is not production code*.

## How to start

The code is not easy to test in its current form, you'll need to refactor
a bit to get started.

## Testing manually

This is a simple Go command, you will need a Pushover API token and User ID to
test this.

```shell
$ export PUSHOVER_TOKEN=atxxxxxxxxxxxxxxxxxxxxxxxxxxxx
$ export NOTIFICATION_USER=ufxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
$ go run main.go
```

The `testing.sh` script will simulate a GitHub hook notification with an
appropriate payload.

```shell
$ ./testing.sh
```

## Things to think about

 - Why is the code hard to test?
 - Why is it harder to refactor?
 - What should we do if Pushover is down?
 - What if we wanted to notify more than one user?
 - What if we wanted multiple notification methods, i.e. AWS SNS and Pushover?

## Things to try

  * Apply the ["Standard Go Project Layout"](https://github.com/golang-standards/project-layout)
  * Adopt a command-line [framework](https://github.com/spf13/cobra)
  * Consider testing with an [HTTP mocking framework](https://github.com/h2non/gock) - is this easier than `net/http/httptest` ?
