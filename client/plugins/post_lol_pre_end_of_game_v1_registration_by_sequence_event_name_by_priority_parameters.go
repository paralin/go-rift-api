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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams creates a new PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams object
// with the default values initialized.
func NewPostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams() *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams {
	var ()
	return &PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParamsWithTimeout creates a new PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParamsWithTimeout(timeout time.Duration) *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams {
	var ()
	return &PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams{

		timeout: timeout,
	}
}

// NewPostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParamsWithContext creates a new PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParamsWithContext(ctx context.Context) *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams {
	var ()
	return &PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams{

		Context: ctx,
	}
}

// NewPostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParamsWithHTTPClient creates a new PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParamsWithHTTPClient(client *http.Client) *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams {
	var ()
	return &PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams{
		HTTPClient: client,
	}
}

/*PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams contains all the parameters to send to the API endpoint
for the post lol pre end of game v1 registration by sequence event name by priority operation typically these are written to a http.Request
*/
type PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams struct {

	/*Priority*/
	Priority int32
	/*SequenceEventName*/
	SequenceEventName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol pre end of game v1 registration by sequence event name by priority params
func (o *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams) WithTimeout(timeout time.Duration) *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol pre end of game v1 registration by sequence event name by priority params
func (o *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol pre end of game v1 registration by sequence event name by priority params
func (o *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams) WithContext(ctx context.Context) *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol pre end of game v1 registration by sequence event name by priority params
func (o *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol pre end of game v1 registration by sequence event name by priority params
func (o *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams) WithHTTPClient(client *http.Client) *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol pre end of game v1 registration by sequence event name by priority params
func (o *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPriority adds the priority to the post lol pre end of game v1 registration by sequence event name by priority params
func (o *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams) WithPriority(priority int32) *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams {
	o.SetPriority(priority)
	return o
}

// SetPriority adds the priority to the post lol pre end of game v1 registration by sequence event name by priority params
func (o *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams) SetPriority(priority int32) {
	o.Priority = priority
}

// WithSequenceEventName adds the sequenceEventName to the post lol pre end of game v1 registration by sequence event name by priority params
func (o *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams) WithSequenceEventName(sequenceEventName string) *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams {
	o.SetSequenceEventName(sequenceEventName)
	return o
}

// SetSequenceEventName adds the sequenceEventName to the post lol pre end of game v1 registration by sequence event name by priority params
func (o *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams) SetSequenceEventName(sequenceEventName string) {
	o.SequenceEventName = sequenceEventName
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolPreEndOfGameV1RegistrationBySequenceEventNameByPriorityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param priority
	if err := r.SetPathParam("priority", swag.FormatInt32(o.Priority)); err != nil {
		return err
	}

	// path param sequenceEventName
	if err := r.SetPathParam("sequenceEventName", o.SequenceEventName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
