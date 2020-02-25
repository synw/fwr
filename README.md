# Flutter web rebuilder

A development server for Flutter web

- **Automatic rebuild** and reload in browser on changes
- **Multi devices** autoreload

## Usage

Grab the binary (or compile with Go) and put it in the path

Add this to `web/index.html` in your project:

```html
  <script type="text/javascript">
    const ip = "192.168.1.2";
    (function () {
      var conn = new WebSocket("ws://" + ip + ":8042/ws");
      conn.onmessage = function (evt) {
        window.location.reload();
      }
    })();
  </script>
```

`ip` is the ip of the machine running the server. Just set it to `localhost` if the same machine that runs the server is used to view the page. Use the ip of the machine running the server to see the page on other devices from the local network

Cd to your Flutter web project folder and run:

```
fwr
```

This will rebuild the project for the web on any change in lib/ or web/ and reload the page in all the opened browsers

Go to `http://192.168.1.2:8085` to see the page