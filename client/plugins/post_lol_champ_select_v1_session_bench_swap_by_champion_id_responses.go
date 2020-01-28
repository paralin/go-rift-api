// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolChampSelectV1SessionBenchSwapByChampionIDReader is a Reader for the PostLolChampSelectV1SessionBenchSwapByChampionID structure.
type PostLolChampSelectV1SessionBenchSwapByChampionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolChampSelectV1SessionBenchSwapByChampionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolChampSelectV1SessionBenchSwapByChampionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolChampSelectV1SessionBenchSwapByChampionIDOK creates a PostLolChampSelectV1SessionBenchSwapByChampionIDOK with default headers values
func NewPostLolChampSelectV1SessionBenchSwapByChampionIDOK() *PostLolChampSelectV1SessionBenchSwapByChampionIDOK {
	return &PostLolChampSelectV1SessionBenchSwapByChampionIDOK{}
}

/*PostLolChampSelectV1SessionBenchSwapByChampionIDOK handles this case with default header values.

Successful response
*/
type PostLolChampSelectV1SessionBenchSwapByChampionIDOK struct {
	Payload interface{}
}

func (o *PostLolChampSelectV1SessionBenchSwapByChampionIDOK) Error() string {
	return fmt.Sprintf("[POST /lol-champ-select/v1/session/bench/swap/{championId}][%d] postLolChampSelectV1SessionBenchSwapByChampionIdOK  %+v", 200, o.Payload)
}

func (o *PostLolChampSelectV1SessionBenchSwapByChampionIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolChampSelectV1SessionBenchSwapByChampionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
