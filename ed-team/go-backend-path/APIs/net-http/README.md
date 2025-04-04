# NET / HTTP Package in Go

This package allow us to create APIs and web applications. It's powerful. We can create http client and servers.

## Servers

Works as a router, create services and response to request. We can use ServeMux that compare HTTP requests with the list of URIs that will handle that requests.

## Handlers

Are functions that response to the request clients. Response headers and response body.

### Predefined Handlers

- FileServer (return static files)
- NotFoundHandler (i don't found a handler)
- RedirectHandler (redirection to other handler)
- StripPrefix (static files /public/ /private/ remove those prefix)
- TimeoutHandler (Timeout)

To run the server correctly must be run on the root project dir.

To create a handler:

```Go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", handlerHelloWorld)
	println("Server running on http://127.0.0.1:9080")
	http.ListenAndServe(":9080", nil)
}

func handlerHelloWorld(writer http.ResponseWriter, reader *http.Request) {
	fmt.Printf("Hello World from server %v",9)
}

```

so, with the `HandleFunc` we can define an endpoint and the required handler that will process our request. Every handler must receive a `responseWriter` and a `pointer to the request`.

## Request Struct

Struct `Http.Request` contains all the information of a client request server. We can found the following fields:

- URL: *url.URL
- Method: string
- Header: map[string][]string
- Body: io.ReadCloser
- Form/PostForm url.Values (form but not a json form)
- MultipartForm: *multipart.Form (files)

## Request Methods

These are the most common used methods:

```text
NewRequest(...): Creates a new request.
BasicAuth(): Gets user and password from a basic authentication (not recommended).
Context(): Gets the context of a request.
Cookie(...): Gets the cookie via cookie name.
Cookies(...): Gets all cookies.
FormFile(...): Gets the first file sent in the request.
FormValue(...): Gets the first value of a sent file via multipart-form.
ParseForm(): fulfill the fields from Form y PostForm.
ParseMultipartForm(...): fulfill the field MultipartForm.
```

## Response Struct

Struct response helps me to create the object to response a request.

- Status: string
- StatusCode: string
- Header: map[string][]string
- Body: io.ReadCloser

## ResponseWriter interface

but we use `ResponseWriter` interface. It's used by handlers to response, methods from the interface:

- Header(): write response headers
- Write(): write body response
- WriteHeader(): write status code response

## Server Struct

This struct allows us to customize server configurations. And that Struct will help with that.

some fields:

- Addr: string
- Handler: Handler
- ReadTimeout: time.Duration
- WriteTimeout: time.Duration
- MaxHeaderBytes: int
- TLSConfig: *int.Config
- ErrorLog: *log.Logger
 