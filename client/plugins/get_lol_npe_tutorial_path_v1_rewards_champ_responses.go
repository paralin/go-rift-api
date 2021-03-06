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

// GetLolNpeTutorialPathV1RewardsChampReader is a Reader for the GetLolNpeTutorialPathV1RewardsChamp structure.
type GetLolNpeTutorialPathV1RewardsChampReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolNpeTutorialPathV1RewardsChampReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolNpeTutorialPathV1RewardsChampOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolNpeTutorialPathV1RewardsChampOK creates a GetLolNpeTutorialPathV1RewardsChampOK with default headers values
func NewGetLolNpeTutorialPathV1RewardsChampOK() *GetLolNpeTutorialPathV1RewardsChampOK {
	return &GetLolNpeTutorialPathV1RewardsChampOK{}
}

/*GetLolNpeTutorialPathV1RewardsChampOK handles this case with default header values.

Successful response
*/
type GetLolNpeTutorialPathV1RewardsChampOK struct {
	Payload *models.LolNpeTutorialPathCollectionsChampion
}

func (o *GetLolNpeTutorialPathV1RewardsChampOK) Error() string {
	return fmt.Sprintf("[GET /lol-npe-tutorial-path/v1/rewards/champ][%d] getLolNpeTutorialPathV1RewardsChampOK  %+v", 200, o.Payload)
}

func (o *GetLolNpeTutorialPathV1RewardsChampOK) GetPayload() *models.LolNpeTutorialPathCollectionsChampion {
	return o.Payload
}

func (o *GetLolNpeTutorialPathV1RewardsChampOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolNpeTutorialPathCollectionsChampion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
