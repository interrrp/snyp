// Contains HTTP method helpers.

package route

// Method is an HTTP method.
type Method string

const (
	// MethodGET is the GET HTTP method.
	//
	// The GET HTTP method is commonly used to fetch
	// data from a specific resource.
	MethodGET Method = "GET"

	// MethodPOST is the POST HTTP method.
	//
	// POST is commonly used to create new resources.
	MethodPOST Method = "POST"

	// MethodPUT is the PUT HTTP method.
	//
	// PUT is commonly used to update resources.
	MethodPUT Method = "PUT"

	// MethodDELETE is the DELETE HTTP method.
	//
	// DELETE is commonly used to delete resources.
	MethodDELETE Method = "DELETE"
)
