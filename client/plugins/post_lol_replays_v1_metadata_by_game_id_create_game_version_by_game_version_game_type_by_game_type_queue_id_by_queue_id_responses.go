// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueIDReader is a Reader for the PostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueID structure.
type PostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueIDNoContent creates a PostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueIDNoContent with default headers values
func NewPostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueIDNoContent() *PostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueIDNoContent {
	return &PostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueIDNoContent{}
}

/*PostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueIDNoContent handles this case with default header values.

No content
*/
type PostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueIDNoContent struct {
}

func (o *PostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueIDNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-replays/v1/metadata/{gameId}/create/gameVersion/{gameVersion}/gameType/{gameType}/queueId/{queueId}][%d] postLolReplaysV1MetadataByGameIdCreateGameVersionByGameVersionGameTypeByGameTypeQueueIdByQueueIdNoContent ", 204)
}

func (o *PostLolReplaysV1MetadataByGameIDCreateGameVersionByGameVersionGameTypeByGameTypeQueueIDByQueueIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}