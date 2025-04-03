# Technical Task - Cron Expression Parser
This assesment is built using golang and provide the basic cron expression parsing.
feel free to have look to *_test.go files to get better undertanding of the usage.
Also don't forget to check `cmd/main.go` file to see example usage of parser.

# Requirements
- docker

>You need to install [docker](https://docs.docker.com/desktop/) in order to run the example and tests.

# Installation
```bash
$ git clone git@github.com:AhmadWaleed/cron-parser.git
$ cd cron-parser && chmod +x ./go.sh
```

# Run Example
```bash
$ ./go.sh run cmd/main.go
```
```go
// Example Output

// minute          0 15 30 45
// hour            0
// day of month    1 15
// month           1 2 3 4 5 6 7 8 9 10 11 12
// day of week     1 2 3 4 5
// command         /usr/bin/find
```

# Run Tests
```bash
$ ./go.sh test
```
# cron-parser
