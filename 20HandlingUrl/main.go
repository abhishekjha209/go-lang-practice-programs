package main

import (
	"fmt"
	"net/url"
)

// test link
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl) // Entire URL

	//parsing
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)	// protocol scheme: https
	// fmt.Println(result.Host)	// host or host:port  lco.dev:3000
	// fmt.Println(result.Path)	// path (relative paths may omit leading slash) /learn
	// fmt.Println(result.Port())	// func for port number  3000

	fmt.Println(result.RawQuery) // coursename=reactjs&paymentid=ghbj456ghb

	qparams := result.Query() // returns map.
	fmt.Println(qparams)
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}

/*
type URL struct {
	Scheme      string
	Opaque      string    // encoded opaque data
	User        *Userinfo // username and password information
	Host        string    // host or host:port
	Path        string    // path (relative paths may omit leading slash)
	RawPath     string    // encoded path hint (see EscapedPath method)
	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	RawQuery    string    // encoded query values, without '?'
	Fragment    string    // fragment for references, without '#'
	RawFragment string    // encoded fragment hint (see EscapedFragment method)
}


A URL represents a parsed URL. The general form represented is:

[scheme:][//[userinfo@]host][/]path[?query][#fragment]

*/
