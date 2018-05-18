package main

import (
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "time"

  "go.uber.org/zap"
)

func getAPITestResult(endpoint string) HttpAPIResult {
  var err error
  defer func() {
      if (recover() != nil) {
          err = errors.New("panic occurred")
      }
  }()

  start_time := time.Now()
  resp, err := http.Get(endpoint) // should define http with Timeout
  elapsed_time := time.Since(start_time)

  logger, _ := zap.NewProduction()
  defer logger.Sync()
  logger.Info("api-get-time",
    zap.String("url", endpoint),
    zap.Duration("time", elapsed_time),
    zap.Int("i", 0),
  )

  code, body_len := SanitizeStatusCode(resp, endpoint, elapsed_time, err)

  api_result := HttpAPIResult{ResultEndpoint: endpoint,
                              ResultStatusCode: code,
                              ResultLen: body_len,
                              ResultTimePerformed: start_time.UnixNano(),
                              ResultRoundtrip: elapsed_time.Nanoseconds()}
  return api_result
}

// if DNS fails it will panic. thanks http?
func SanitizeStatusCode(resp *http.Response, endpoint string, elapsed_time time.Duration, err error) (int, int) {
  var body_size int
  logger, _ := zap.NewProduction()
  defer logger.Sync()
  if err != nil {
    fmt.Println(err)
    logger.Error("api-SanitizeStatusCode",
      zap.Int("code", 502),
    )
    return 502, 0
  } else {
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      body_size = len(body)
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
  return resp.StatusCode, body_size
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
