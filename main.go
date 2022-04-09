package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
)

/*
	The main function starts the entire program. It starts by creating a new
	router that makes sure that each request gets handled by the correct function.
*/
func main() {
	router := mux.NewRouter()

	/*
		Each router.HandleFunc method handles a route and attaches a function to a
		route (url path) that takes care of these requests. Here we attach /hello
		to the hello function below.

		The .Methods() part ensures that a function will only apply to certain
		http methods (e.g. GET, POST, PUT and DELETE)
	*/
	router.HandleFunc("/hello", hello).Methods("GET")

	// You can have a different function to handle POST request to the same path
	router.HandleFunc("/hello", postHello).Methods("GET")

	/*
		A path can contain dynamic parameters which can either contain anything or a certain
		pattern. This parameter can contain anything.
	*/
	router.HandleFunc("/print/{what_to_print}", print).Methods("GET")

	// The function name doesn't have to be the same as the path name
	router.HandleFunc("/system", getSystemInfo).Methods("GET")

	router.HandleFunc("/request-info/{params}", requestInfo)

	port := "5000"

	fmt.Printf("Running on http://localhost:%s\n", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	/*
		To print a simple text string to the client, we use Fprint from the fmt package.
		It requires some sort of writer, where we in this case use a http responsewriter w.
	*/
	fmt.Fprint(w, "Hello to you too!")
}

func postHello(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	json.NewDecoder(r.Body).Decode(&body)

	/*
		A map is created to store multiple key-value pairs. Here it stores key value pairs of
		different type through the use of the interface data type. In this endpoint two strings and a map is used as values.

		A map can be compared to a dictionary in Python, object in JavaScript or a HashMap in C++, Java and C#.
	*/
	output := map[string]interface{}{
		"endpoint":        "hello",
		"function":        "postHello",
		"what_did_i_send": body,
	}

	/*
		To deliver the data to the user in JSON format, a json encoder is used where
		the output variable is encoded using a new json encoder based on the
		http.ResponseWriter w (it acts as a channel to write through).
	*/
	json.NewEncoder(w).Encode(output)
}

func print(w http.ResponseWriter, r *http.Request) {
	/*
		To fetch dynamic url parameters, use mux.Vars and use the *http.Request parameter as input.
		mux.Vars(r) returns a key-value pair of all potentiel url parameters in the path.
		In this case there's only one and it's called what_to_print.

		An example of returning the full mux.Vars output as JSON can be found in the requestInfo function.
	*/
	text_to_print := mux.Vars(r)["what_to_print"]
	fmt.Fprint(w, text_to_print)
}

func getSystemInfo(w http.ResponseWriter, r *http.Request) {
	/*
		This endpoint responds with system info
		from the runtime environment using the runtime package. This gives info
		about the servers' system info - not the client.
	*/
	system_info := map[string]string{
		"operating_system":    runtime.GOOS,
		"system_architecture": runtime.GOARCH,
	}

	/*
		To deliver the data to the user in JSON format, a json encoder is used where
		the system_info variable is encoded using a new json encoder based on the
		http.ResponseWriter w (it acts as a channel to write through).
	*/
	json.NewEncoder(w).Encode(system_info)
}

func requestInfo(w http.ResponseWriter, r *http.Request) {
	/*
		This example gets you the most important things to get from a request through a web
		service (API) and sends the data encoded in JSON format.
	*/
	request_info := map[string]interface{}{
		"dynamic_url_parameters": mux.Vars(r),
		"path":                   r.URL.Path,
		"query_parameters":       r.URL.Query(),
		"http_method":            r.Method,
		"host":                   r.Host,
		"headers":                r.Header,
	}
	json.NewEncoder(w).Encode(request_info)
}
