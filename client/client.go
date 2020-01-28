package client

import (
	"fmt"

	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
)

// ConstructAPIClient constructs an API client from process information.
func ConstructAPIClient(port int, token string) (*Lol, error) {
	if port == 0 {
		return nil, errors.New("port must be set")
	}
	if token == "" {
		return nil, errors.New("token must be set")
	}

	transport := httptransport.New(
		fmt.Sprintf("127.0.0.1:%d", port),
		"/",
		[]string{"https"},
	)
	formats := strfmt.Default
	return New(transport, formats), nil
}
