# Flutter web rebuilder

A development server for Flutter web

- **Automatic rebuild** and reload in browser on changes
- **Multi devices** autoreload

## Usage

### Get the server

Grab the [Linux binary](https://github.com/synw/fwr/releases/download/0.2.0/fwr) or compile with Go, and put the executable in the path
or in your local Flutter project directory

### Setup the autoreload

Add this to `web/index.html` in your project:

```html
  <script type="text/javascript">
    const ip = "<server-ip>"; // The ip of the server: ex: 192.168.1.3 or localhost
    (function () {
      var conn = new WebSocket("ws://" + ip + ":8042/ws");
      conn.onmessage = function (evt) {
        window.location.reload();
      }
    })();
  </script>
```

`ip` is the ip of the machine running the server. Just set it to `localhost` if the same machine that runs the server is used to view the page. Use the ip of the machine running the server to see the page on other devices from the local network

### Run

Cd to your Flutter web project folder and run:

```
fwr
```

This will rebuild the project for the web on any change in lib/ or web/ and reload the page in all the opened browsers

Go to `http://<server-ip>:8085` to see the page