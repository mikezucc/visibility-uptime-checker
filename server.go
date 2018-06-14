package main

import (
  "encoding/json"
  "errors"
  "fmt"
  "log"
  "net/http"
  "time"
  "strings"
  "sync"

  "github.com/googollee/go-socket.io"
)

/** https://gist.github.com/harshavardhana/327e0577c4fed9211f65 */
const (
	day  = time.Minute * 60 * 24
	year = 365 * day
)

type ServerStatus struct {
  Code string
}

/** ZERO SCOPE IVARS */
var cache CACHE
var server *socketio.Server

var serverStatus string

func duration(d time.Duration) string {
	if d < day {
		return d.String()
	}

	var b strings.Builder

	if d >= year {
		years := d / year
		fmt.Fprintf(&b, "%dy", years)
		d -= years * year
	}

	days := d / day
	d -= days * day
	fmt.Fprintf(&b, "%dd%s", days, d)

	return b.String()
}

func main() {
  var err error
  defer func() {
      if (recover() != nil) {
          err = errors.New("panic occurred")
      }
  }()

  serverStatus = "200"

  cache = CACHE{cached_results: loadRecords()}
  cached_json, err := json.Marshal(cache.cached_results)
  fmt.Println(string(cached_json))

  namespace_notification_root_domain := "nnrd0"
  event_status_update := "esu0"
  event_status_result := "esr0"
  // server_uptime_update := "suu0"
  server_cache_burst := "scb0"

  // server_uptime_start := time.Now().Round(0).Add(-(3600 + 60 + 45) * time.Second)

  server_loc, _ := socketio.NewServer(nil)
  server = server_loc

  server.On("connection", func(so socketio.Socket) {
		fmt.Println("[SOCKETIO] NEW connected to client %s", so.Id())

		so.Join(namespace_notification_root_domain)
		so.Emit("server-ack-connect", "")

    so.On(event_status_update, func(data string) {
      fmt.Println("[SOCKETIO] EVENT ")

      result := getAPITestResult(data) //APITestResult
      // send result to web page
      fmt.Println(result)
      results_json, err := json.Marshal(result)
      fmt.Println(string(results_json))
      if err != nil {
        fmt.Println("[SOCKETIO] failed encode api result: " + err.Error())
      }
      so.Emit(event_status_result, string(results_json))
      recordAPIResult(result)
    })

    cached_json, err := json.Marshal(cache.cached_results)
    if err == nil {
      so.Emit(server_cache_burst, string(cached_json))
    } else {
      fmt.Println("[SOCKETIO] Cache encode failed: " + err.Error())
    }

    /** Default Socket.io callbacks */
    so.On("disconnection", func() {
      log.Println("[SOCKETIO] DISCONNECT ")
    })
	})

  /** Routes */
  http.Handle("/socket.io/", server)
  http.Handle("/", http.FileServer(http.Dir("./public")))

  log.Println("[SOCKETIO] Serving all :3008...")
  go func() {
    fmt.Println("[status] server start goroutine")
    log.Fatal(http.ListenAndServe(":3008", nil))
    server.On("error", func(so socketio.Socket, err error) {
      log.Println("[SOCKETIO] error:", err)
      if err == nil {
        serverStatus = "500"
      } else {
        serverStatus = "200"
      }
    })
    fmt.Println("[status] servers fine")
  }()

  fmt.Println("[status]  slideways")

	ticker := time.NewTicker(2000 * time.Millisecond)
	go func() {
    fmt.Println("[statuis]  server TICK goroutine")
    if serverStatus == "500" {
      fmt.Println("[status] HTTP server failed to start, returning tick early")
      return
    }
		for timeidx := range ticker.C {
      fmt.Println(timeidx)
			for endpointName, _ := range cache.cached_results.Endpoints {
        result := getAPITestResult(endpointName) //APITestResult
        // send result to web page
        fmt.Println(result)
        results_json, err := json.Marshal(result)
        fmt.Println(string(results_json))
        if err != nil {
          fmt.Println("[SOCKETIO] failed encode api result: " + err.Error())
        }
        recordAPIResult(result)
      }
      cached_json, err := json.Marshal(cache.cached_results)
      if err == nil {
			  server.BroadcastTo(namespace_notification_root_domain, server_cache_burst, string(cached_json))
      }
		}
  }()

  var wg sync.WaitGroup
  wg.Add(1)
  wg.Wait()

}
