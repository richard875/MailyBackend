# Maily Backend

---

### This is the backend project for Maily. Currently, the client software is a macOS application and an Apple Mail extension powered by SwiftUI and MailKit.

### Go Air

```azure
$ alias air='/Users/{user}/go/bin/air'
$ air
```

### TODO

```azure
Links clicked
```

### Setup

```azure
go install github.com/cosmtrek/air@latest
go install github.com/swaggo/swag/cmd/swag@latest
```

### Swagger URL

```azure
http://localhost:8090/swagger/index.html
```

### IP Websites

```azure
https://ip-api.com/ Free with limited information
https://ipdata.co/ 1,500 request daliy with more information
```

### How to stop the cache server

```azure
sudo lsof -i :8090
kill -9 <PID>
```
###### Or
```azure
sudo kill -9 $(sudo lsof -t -i:8090)
```

### How to use Localtunnel

```azure
https://theboroer.github.io/localtunnel-www/

If you're the developer...
You and other visitors will only see this page from a standard web browser once per IP every 7 days.

Webhook, IPN, and other non-browser requests "should" be directly tunnelled to your localhost. If your webhook/ipn provider happens to send requests using a real browser user-agent header, those requests will unfortunately also be blocked / be forced to see this tunnel reminder page. FYI, this page returns a 401 HTTP Status.

Options to bypass this page:
Set and send a Bypass-Tunnel-Reminder request header (its value can be anything).
or, Set and send a custom / non-standard browser User-Agent request header.
```

## Build

```azure
go clean -cache -r
go build -a -tags netgo -ldflags \'-s -w\' -o app
```

## Reverse Proxy

### ngrok

```azure
ngrok http 8080
```

### localtunnel

```azure
lt --port 8090
```
