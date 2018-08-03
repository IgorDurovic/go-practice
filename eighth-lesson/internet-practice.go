package main

// ACCESSING DATA FROM WEBSITES ON THE INTERNET
// USING THE GO STANDARD LIBRARY

import ("fmt"
		"net/http"
		"io/ioutil")

// The raw data that's received from a website
// needs to be deserialized and interpreted so that
// it can be used in a meaningful way

func main() {
	// use underscore when you don't intend to assign
	// a value to a variable. In this case the second
	// return value is the status code of the get request
	//
	// The site map contains references to the different
	// topics on the washington post wesbite
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")

	// Retrieve the body of the response. The second output
	// is an error if one occurs (this is a common pattern in go)
	bytes, _ := ioutil.ReadAll(resp.Body)

	// Interpret the bytes as a string
	string_body := string(bytes)

	// Close the response stream to free up resources
	resp.Body.Close()

	fmt.Println("The response from wapo is:\n", string_body)

	// The response body interpreted as a string is xml from the
	// url above.
}