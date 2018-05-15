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
  logger.Info("api-get-time",
    zap.String("url", endpoint),
    zap.Duration("time", elapsed_time),
    zap.Int("i", 0),
  )

  code := SanitizeStatusCode(resp, endpoint, elapsed_time, err)

  return APITestResult{endpoint, elapsed_time, code}
}

// if DNS fails it will panic. thanks http?
func SanitizeStatusCode(resp *http.Response, endpoint string, elapsed_time time.Duration, err error) int {
  logger, _ := zap.NewProduction()
  defer logger.Sync()
  if err != nil {
    fmt.Println(err)
    logger.Error("api-SanitizeStatusCode",
      zap.Int("code", 502),
    )
    return 502
  } else {
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      body_size := len(body)
      logger.Info("Result",
        zap.String("url", endpoint),
        zap.Duration("time", elapsed_time),
        zap.Int("bs", body_size),
        zap.Int("i", 0),
      )
    } else {
      // query Max for good practices
    }
  }
  return resp.StatusCode
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
