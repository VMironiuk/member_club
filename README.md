[![Go](https://github.com/VMironiuk/member_club/actions/workflows/CI.yml/badge.svg)](https://github.com/VMironiuk/member_club/actions/workflows/CI.yml)

# Install dependencies:
- Gin:
```
$ go get -u github.com/gin-gonic/gin
```
- Official CORS gin's middleware
```
$ go get -u github.com/gin-contrib/cors
```

# Start server:
- From the preject's root directory run the command below:
```
$ go run main.go
```

# Test server with the client:
- Just run *member_club/client/index.html* in a browser