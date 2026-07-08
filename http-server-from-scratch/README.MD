# http-server: A Minimal HTTP/1.1 Server in Go

`http-server` is a lightweight, custom-built HTTP/1.1 server written in Go. It uses Go's `net` package for listening and managing TCP connections, while implementing the full HTTP protocol parsing and request handling from the ground up.

## Features

- **Raw TCP Server** with connection management
- **Complete HTTP/1.1 Protocol Support**: Full parsing of HTTP methods, headers, and request bodies.
- **Custom Response Builder**: Clean and flexible `HTTPResponse` structure for constructing and sending HTTP responses.
- **Handler Interface & Minimal Router**: Simple but effective handler pattern with a lightweight router, inspired by Go's `http.ServeMux`.
- **Concurrent Request Handling**: Goroutine-based concurrency with per-connection model, bounded channels, and backpressure mechanisms.
- **Persistent Connections**: HTTP Keep-Alive support with request read timeout and idle timeout for efficient resource management.

## Package Structure

```md
app/
├── server/ # Low-level TCP server: listening, handling connections 
├── http/ # HTTP protocol: request, response, gzip 
└── services/ # Example services 
```

