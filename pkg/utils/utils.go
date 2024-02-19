package utils

// "encoding/json": This package provides functions for encoding and decoding JSON data.
// "io": This package provides basic I/O interfaces.
// "net/http": This package provides HTTP client and server implementations.
import (
	"encoding/json"
	"io"
	"net/http"
)

// x as input parameters. x is an empty interface, meaning it can accept values of any type.
// parse JSON request bodies in HTTP requests.
func ParseBody(r *http.Request, x interface{}) {

	// io.ReadAll(r.Body): This reads the request body using io.ReadAll function, which reads from r.Body
	// until an error or EOF (end of file) occurs, and returns the data read.
	// read until it reaces end and gives error
	if body, err := io.ReadAll((r.Body)); err == nil {

		// if no error while reading , extract the json , return if error
		// decoding value from body to pointer(which is x interface int his case)
		// converts string to json object
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}

}

/*
json.Unmarshal([]byte(body), x): This attempts to unmarshal (decode) the JSON data read from the request body into the provided interface x.
The json.Unmarshal function converts JSON data into Go values based on the provided struct type or interface{}.
If the JSON data can be successfully unmarshaled into the provided interface x, the function completes successfully.
If there's an error during unmarshaling, the function returns without modifying x.
*/
