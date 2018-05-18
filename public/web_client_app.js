window.onload = function() {
  var k_SOCKET_ENDPOINT_PUBLIC_OSIRIS = "ws://" + window.location.hostname + ':3008';

  // essentially force web socket here
  var socket = io(k_SOCKET_ENDPOINT_PUBLIC_OSIRIS, {transports: ['websocket']});

  var namespace_notification_root_domain = "nnrd0"
  var event_status_update = "esu0"
  var event_status_result = "esr0"
  var server_uptime_update = "suu0"
  var server_cache_burst = "scb0"

  socket.on("server-ack-connect", function (data) {
    console.log("[SOCKET.IO] > server ack > ");
    console.log(data);
  });

  socket.on(event_status_result, function (data) {
    console.log(data);
    if (data.indexOf("200") != -1) {
      document.getElementById("success").style.display = "block";
      document.getElementById("success").innerHTML = data;
      document.getElementById("failure").style.display = "none";
    } else {
      document.getElementById("failure").style.display = "block";
      document.getElementById("failure").innerHTML = data;
      document.getElementById("success").style.display = "none";
    }
  })

  document.getElementById("endpoint").addEventListener("keyup", function(event) {
    // Cancel the default action, if needed
    event.preventDefault();
    e = event || window.event;
    // Number 13 is the "Enter" keyCode on the keyboard
    if (e.keyCode == 13)
    socket.emit(event_status_update, document.getElementById("endpoint").value)
  });

  socket.on(server_cache_burst, function (data) {
    console.log(json);
    var cache_results = JSON.parse(data);
    for (var result_idx in cached_results.Endpoints) {
      var result = cache_results.Endpoints[result_idx];
      var result_pre_format = JSON.stringify(result, null, "\t");
      var resultDiv = $('#success').clone();
      var rana = (Math.floor(Math.random() * 10000)).toString();
      resultDiv.attr('id',);
      var sessionMessage = '<pre id="session-message-"' + rana +' "style="word-wrap: break-word; white-space: pre-wrap; margin-bottom:0px;">'+ result_pre_format + '</pre>';
      resultDiv.html(sessionMessage)
      resultDiv.appendTo($('#status-container'))
      resultDiv.show()
    }
  })
}

/**
some sample copy pasta
logDiv = $('#template-session-div').clone();
tableBody = logDiv.find('#logs-console-tbody');
tableBody.attr('id', "logs-console-tbody-" + logContainerDivIdentifier);
logDiv.attr('id', logContainerDivIdentifier);
logDiv.addClass("rendered-done");
processedSessionMap[session_identifier] = logContainerDivIdentifier;
logDiv.appendTo('#template-session-container');
logDiv.attr("style", "left: "+ ((session_counter * 950)).toString() +"; display: table;  position:absolute; top: 0px; padding-left:20px; width:950px; display:inline-block; border-right: 2px solid gray;");
logDiv.show();
logRow.attr('id', "template-log-row-" + log_id);
logRow.find("[col='time']").html(printDateWithoutPipes());
logRow.find("[col='log']").html(logMessage);
*/
