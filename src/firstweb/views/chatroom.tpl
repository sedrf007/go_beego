<html>
<head></head>
<body>
    <script type="text/javascript">
        var sock = null;
        var wsuri = "ws://www.lovepangchao.com/chat";

        window.onload = function() {

            console.log("onload");

            sock = new WebSocket(wsuri);

            sock.onopen = function() {
                console.log("connected to " + wsuri);
            }

            sock.onclose = function(e) {
                console.log("connection closed (" + e.code + ")");
            }

            sock.onmessage = function(e) {
                console.log("message received: " + e.data);
            }
        };

        function send() {
            var msg = document.getElementById('message').value;
            sock.send(msg);
            $('#message').val("");
        };
    </script>
    <h1>许个愿吧，胖超！</h1>
    <form>
        <p>
            Message: <input id="message" type="text" value="">
        </p>
    </form>
    <button onclick="send();">Send Message</button>
</body>
</html>