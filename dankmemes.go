package main

import (
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "time"

  "go.uber.org/zap"
)

type APITestResult struct {
  url string
  time_elapsed time.Duration
  code int
}

func getAPITestResult(endpoint string) APITestResult {
  var err error
  defer func() {
      if (recover() != nil) {
          err = errors.New("panic occurred")
      }
  }()

  start_time := time.Now()
  resp, err := http.Get(endpoint)
  elapsed_time := time.Since(start_time)

  logger, _ := zap.NewProduction()
  defer logger.Sync()
  logger.Info("dankmemes-get-time",
    zap.String("url", endpoint),
    zap.Duration("time", elapsed_time),
    zap.Int("i", 0),
  )

  code := checkAndZap(err)

  body, err := ioutil.ReadAll(resp.Body)
  body_size := len(body)

  logger, _ = zap.NewProduction()
  logger.Info("failed to fetch URL",
    zap.String("url", endpoint),
    zap.Duration("time", elapsed_time),
    zap.Int("bs", body_size),
    zap.Int("i", 0),
  )
  return APITestResult{endpoint, elapsed_time, code}
}

func checkAndZap(err error) int {
  if err != nil {
    fmt.Println(err)
    logger, _ := zap.NewProduction()
    defer logger.Sync()
    logger.Info("dankmemes-checkAndZap",
      zap.String("code", "4xx"),
    )
    return 400
  }
  return 200
}

// MUSINGS -

// formulate a request to my web package
// this excepts on failed DNS Lookup
// for some reason references 127.0.0.1:53
// maybe thats the reason for the name of route 53 :smirk:

// so those encapsulating `#!` denotes script inserted
// functions
// this is useful because the whole point of using zap
// was to be FAST. aka avoid stack pushes, stack reads, etc
// by writing functions as you would normally and then
// placing them in-line before the compiler even takes a pass
// prevent scope changes.
