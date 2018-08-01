package main

// GO PROVIDES A PRETTY SIMPLE INTERFACE FOR RUNNING A BASIC
// WEB SERVER. YOU CAN RUN THIS CODE AND TEST IT ON localhost:8000

import ("fmt"
		"net/http")

// the handler function for the root.
// passes in a writer object and the request
// that needs to be serviced.
//
// http.ResponseWriter and Request are not go types, they're
// custom types
func index_handler(writer http.ResponseWriter, request *http.Request) {
	// Fprintf formats based on the specified writer and outputs
	// to defined stream
	fmt.Fprintf(writer, "Handling request by printing to response writer")
}

func custom_handler(writer http.ResponseWriter, request *http.Request) {
	// you can also use backticks ` for multi line strings instead of just
	// adding regular strings
	fmt.Fprintf(writer, "Handling request to custom path by printing to" +
						"response writer (page)")
}

func structured_handler(writer http.ResponseWriter, request *http.Request) {
	output :=  `<h1>
					<b>
						LET THERE BE HTML/CSS
					</b>
				</h1>`
	fmt.Fprintf(writer, output)
}

func main() {
	// define handler func for a root path
	http.HandleFunc("/", index_handler)

	// define a handler func for a custom path
	http.HandleFunc("/custom_path", custom_handler)

	// you can even return structured html/css when handling
	// a request to a specific path
	http.HandleFunc("/structured", structured_handler)

	// define port to listen for requests on
	// and serve requests to clients through
	//
	// second parameter is for additional server
	// customization that I haven't looked into
	// yet. Pass in nil for now
	http.ListenAndServe(":8000", nil)
}