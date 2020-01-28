// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostLolPreEndOfGameV1CompleteBySequenceEventNameParams creates a new PostLolPreEndOfGameV1CompleteBySequenceEventNameParams object
// with the default values initialized.
func NewPostLolPreEndOfGameV1CompleteBySequenceEventNameParams() *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams {
	var ()
	return &PostLolPreEndOfGameV1CompleteBySequenceEventNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolPreEndOfGameV1CompleteBySequenceEventNameParamsWithTimeout creates a new PostLolPreEndOfGameV1CompleteBySequenceEventNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolPreEndOfGameV1CompleteBySequenceEventNameParamsWithTimeout(timeout time.Duration) *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams {
	var ()
	return &PostLolPreEndOfGameV1CompleteBySequenceEventNameParams{

		timeout: timeout,
	}
}

// NewPostLolPreEndOfGameV1CompleteBySequenceEventNameParamsWithContext creates a new PostLolPreEndOfGameV1CompleteBySequenceEventNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolPreEndOfGameV1CompleteBySequenceEventNameParamsWithContext(ctx context.Context) *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams {
	var ()
	return &PostLolPreEndOfGameV1CompleteBySequenceEventNameParams{

		Context: ctx,
	}
}

// NewPostLolPreEndOfGameV1CompleteBySequenceEventNameParamsWithHTTPClient creates a new PostLolPreEndOfGameV1CompleteBySequenceEventNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolPreEndOfGameV1CompleteBySequenceEventNameParamsWithHTTPClient(client *http.Client) *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams {
	var ()
	return &PostLolPreEndOfGameV1CompleteBySequenceEventNameParams{
		HTTPClient: client,
	}
}

/*PostLolPreEndOfGameV1CompleteBySequenceEventNameParams contains all the parameters to send to the API endpoint
for the post lol pre end of game v1 complete by sequence event name operation typically these are written to a http.Request
*/
type PostLolPreEndOfGameV1CompleteBySequenceEventNameParams struct {

	/*SequenceEventName*/
	SequenceEventName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol pre end of game v1 complete by sequence event name params
func (o *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams) WithTimeout(timeout time.Duration) *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol pre end of game v1 complete by sequence event name params
func (o *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol pre end of game v1 complete by sequence event name params
func (o *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams) WithContext(ctx context.Context) *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol pre end of game v1 complete by sequence event name params
func (o *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol pre end of game v1 complete by sequence event name params
func (o *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams) WithHTTPClient(client *http.Client) *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol pre end of game v1 complete by sequence event name params
func (o *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSequenceEventName adds the sequenceEventName to the post lol pre end of game v1 complete by sequence event name params
func (o *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams) WithSequenceEventName(sequenceEventName string) *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams {
	o.SetSequenceEventName(sequenceEventName)
	return o
}

// SetSequenceEventName adds the sequenceEventName to the post lol pre end of game v1 complete by sequence event name params
func (o *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams) SetSequenceEventName(sequenceEventName string) {
	o.SequenceEventName = sequenceEventName
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolPreEndOfGameV1CompleteBySequenceEventNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sequenceEventName
	if err := r.SetPathParam("sequenceEventName", o.SequenceEventName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}