package main

import (
  "errors"
  "fmt"
  "log"
  "net/http"
  "time"
  "strings"

  "github.com/googollee/go-socket.io"
)

/** https://gist.github.com/harshavardhana/327e0577c4fed9211f65 */
const (
	day  = time.Minute * 60 * 24
	year = 365 * day
)

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

  root_endpoint := "http://ec2-13-56-158-108.us-west-1.compute.amazonaws.com:3003"
  namespace_notification_root_domain := "nnrd0"
  event_status_update := "esu0"
  event_status_result := "ess0"
  server_uptime_update := "suu0"

  server_uptime_start := time.Now().Round(0).Add(-(3600 + 60 + 45) * time.Second)

  server, _ := socketio.NewServer(nil)

  server.On("connection", func(so socketio.Socket) {
		fmt.Println("[SOCKETIO] NEW connected to client %s", so.Id())

		so.Join(namespace_notification_root_domain)
		so.Emit("server-ack-connect", "")

    so.On(event_status_update, func(data string) {
      fmt.Println("[SOCKETIO] EVENT ")

      result := getAPITestResult(root_endpoint)
      // send result to web page
      result_string := fmt.Sprintf("%#v", result)
      so.Emit(event_status_result, result_string)

      up_time := time.Since(server_uptime_start)
      up_time_string := duration(up_time)
      so.Emit(server_uptime_update, up_time_string)
    })

    /** Default Socket.io callbacks */
    so.On("disconnection", func() {
      log.Println("[SOCKETIO] DISCONNECT ")
    })
	})


  /** Routes */
  http.Handle("/socket.io/", server)
  http.Handle("/", http.FileServer(http.Dir("./public")))

  log.Println("[SOCKETIO] Serving all :3003...")
  log.Fatal(http.ListenAndServe(":3003", nil))
	server.On("error", func(so socketio.Socket, err error) {
  	log.Println("[SOCKETIO] error:", err)
  })

}
