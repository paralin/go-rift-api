// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/paralin/go-rift-api/models"
)

// GetLolGamhsV1ProductsByProductIDByPuuidMatchesReader is a Reader for the GetLolGamhsV1ProductsByProductIDByPuuidMatches structure.
type GetLolGamhsV1ProductsByProductIDByPuuidMatchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolGamhsV1ProductsByProductIDByPuuidMatchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolGamhsV1ProductsByProductIDByPuuidMatchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolGamhsV1ProductsByProductIDByPuuidMatchesOK creates a GetLolGamhsV1ProductsByProductIDByPuuidMatchesOK with default headers values
func NewGetLolGamhsV1ProductsByProductIDByPuuidMatchesOK() *GetLolGamhsV1ProductsByProductIDByPuuidMatchesOK {
	return &GetLolGamhsV1ProductsByProductIDByPuuidMatchesOK{}
}

/*GetLolGamhsV1ProductsByProductIDByPuuidMatchesOK handles this case with default header values.

Successful response
*/
type GetLolGamhsV1ProductsByProductIDByPuuidMatchesOK struct {
	Payload *models.LolGamhsMatchHistoryList
}

func (o *GetLolGamhsV1ProductsByProductIDByPuuidMatchesOK) Error() string {
	return fmt.Sprintf("[GET /lol-gamhs/v1/products/{productId}/{puuid}/matches][%d] getLolGamhsV1ProductsByProductIdByPuuidMatchesOK  %+v", 200, o.Payload)
}

func (o *GetLolGamhsV1ProductsByProductIDByPuuidMatchesOK) GetPayload() *models.LolGamhsMatchHistoryList {
	return o.Payload
}

func (o *GetLolGamhsV1ProductsByProductIDByPuuidMatchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolGamhsMatchHistoryList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}