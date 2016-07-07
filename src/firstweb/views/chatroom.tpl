<html>
<head>
    <script src="//cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
</head>
<body>
    <script type="text/javascript">
        var sock = null;
        window.onload = function() {
            
            var wsuri = "ws://www.lovepangchao.com/chat?username="+$('#uname').html();
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
                var data = JSON.parse(event.data);
                console.log(data);
                $("#chatbox").append("<li><b>" + data.User + "</b>: " + data.Content + "</li>");
            }
            
             
        }
        function send() {
                var msg = $('#message').val();
                sock.send(msg);
                $('#message').val("");
                //$("#chatbox").append("<li><b>I said</b>: " + msg + "</li>");
            }; 
    </script>
    <h1>许个愿吧，胖超！</h1>
    <form>          
        <p>
            <label id="username" value=""><span id="uname">{{.Username}}</span></label> 
            <input id="message" type="text" value="">
        </p>
    </form>
    <button onclick="send();">Send Message</button>
    <div class="container">
    <h3>ChatHistory:</h3>
    <ul id="chatbox">
        <li>welcome!</li>
    </ul>
</div>
</body>
</html>