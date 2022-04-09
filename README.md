# Simple REST API in Go

This repository shows a basic `REST API` with endpoints showcasing different simple things you can do.

## Endpoints

* `/hello`: When using a `GET` request it responds with "Hello to you too!" and when using a `POST` request it responds with a `JSON` object like this:
```javascript
{
    "endpoint":        "hello",
    "function":        "postHello",
    "what_did_i_send": body
}
```
The `body` is a replica of the `JSON` the client attaches to the request.
* `/print/{what_to_print}`: Sends the `{what_to_print}` parameter to the client as plain text
* `/system`: Responds a `JSON` object with information about the system that the server runs (probably your computer).
* `/request-info/{params}`: Responds with a `JSON` object with general information about the request that was sent.

## How to run the REST API?

Start by getting this code repository by either using `clone` or `fork` from `git` or go to `Code` and then `Download ZIP` and extract the repository somewhere on your computer.

It is required that you have `Go` installed on your computer (find it [here](https://go.dev/dl/)). Furthermore, you need to have the packages needed and you install them by going to your folder with the code and inside a terminal (also called `Command prompt` or `Command line`) and type in the following:
```
go mod tidy
```

To start the `REST API` you can compile it using `go build` or type in the following in the terminal/command prompt:
```
go run main.go
```

And the API is running and waiting for your requests.

## How do I use it?

You make requests to the API using whatever tool or language you like. Two easy ways is the user friendly [Postman](https://www.postman.com/downloads/) and the nerd friendly [Curl](https://curl.se/download.html). To call the `/hello` endpoint with `Curl` you type this inside a terminal/command prompt:
```
curl localhost:5000/hello
```

## But it doesn't work!!?
If you have any diffuculty making the code work or anything you didn't understand, feel free to create an `issue` and ask you question.