# APIs Basic Concepts

## HTTP Protocol

The base of the actual APIs. A protocol is a steps by steps series that are being executed to complete a goal in a standard way.  
With the protocol we are only standardizing the steps we are doing.

The idea is to standardize request and responses.

### Request

sample only

```shell
GET /index.html HTTP/1.1
Host: www.example.com
Referer: www.google.com
User-Agent: Mozilla/5.0 (X11; Linux x64; rv:45.0)
```

### Response

sample only

```shell
HTTP/1.1 200 OK
Date: Fri, 31 Dec 2019 24:59:59 GTM
Content-Type: text/html

<html lang='en'>
<head>
```

### RESTFul Architecture

API means (**Application Programming Interfaces**) and means that are interfaces which allow us tu communicate with a resource. RESTFul (**Representational State Transfer**) and is an standard architecture in order to communicate in an unique way with different resources.
- Everything that is moving via internet is a resource
- Every resource is represented like a format
- Every resource must have an unique identifier (URI)
- Resources must be accessed via standard methods (GET, POST, PATCH, PUT, DELETE, OPTIONS, HEAD)
- The resources can have multiple representations (XML, HTTP, plain text)
- Stateless communication (Every new request to the server is a new complete request. The only way to persist request / information.. we can use **Cookies** and **Tokens**).

### Formats

#### Plain Text

```text
Andres W, 26, (Bogota)
```

#### XML

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<person>
    <name>
        Andres W
    </name>
    <age>
        26
    </age>
    <location>
        Bogota
    </location>
</person>
```

#### JSON

JavaScript Object Notation

```json
{
  "name": "Andres W",
  "age": 26,
  "location": "Bogota"
}
```

With those formats we can allow client -server communication.

### HTTP Verbs

The verbs let us to say to the server which is the kind of action we want to do

- GET (request a representation of the resource)
- POST (submits an entity into the specified resource)
- PUT (updates a representation of the resource)
- PATCH (applies partial modification of the resource)
- DELETE (delete a resource)
- OPTIONS (requests permitted communication options for a given URL or server. This can be used to test the allowed HTTP methods for a request.)
- HEAD (requests the metadata of a resource in the form of headers that the server would have sent if the GET method was used instead)

### HTTP Status Codes

Are an standard codes that let identify the actual status of the response from our request:

- 1XX (Information responses)
- 2XX (Success responses)
- 3XX (Redirection response)
- 4XX (Client error responses)
- 5XX (Server error responses)
