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