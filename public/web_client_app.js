window.onload = function() {
  var k_SOCKET_ENDPOINT_PUBLIC_OSIRIS = "ws://" + window.location.hostname + ':3003';

  // essentially force web socket here
  var socket = io(k_SOCKET_ENDPOINT_PUBLIC_OSIRIS, {transports: ['websocket']});

  var namespace_notification_root_domain = "nnrd0"
  var event_status_update = "esu0"
  var event_status_result = "esr0"
  var server_uptime_update = "suu0"

  socket.on("server-ack-connect", function (data) {
    console.log("[SOCKET.IO] > server ack > ");
    console.log(data);

    socket.emit(event_status_update, "")
  });

  socket.on(event_status_result, function (data) {
    console.log(data);
    if (data.indexOf("200") != -1) {
      document.getElementById("success").style.display = "block";
      document.getElementById("success").innerHTML = data;
    } else {
      document.getElementById("failure").style.display = "block";
      document.getElementById("failure").innerHTML = data;
    }
  })
}
