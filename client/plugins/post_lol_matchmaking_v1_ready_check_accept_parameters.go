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

// NewPostLolMatchmakingV1ReadyCheckAcceptParams creates a new PostLolMatchmakingV1ReadyCheckAcceptParams object
// with the default values initialized.
func NewPostLolMatchmakingV1ReadyCheckAcceptParams() *PostLolMatchmakingV1ReadyCheckAcceptParams {

	return &PostLolMatchmakingV1ReadyCheckAcceptParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolMatchmakingV1ReadyCheckAcceptParamsWithTimeout creates a new PostLolMatchmakingV1ReadyCheckAcceptParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolMatchmakingV1ReadyCheckAcceptParamsWithTimeout(timeout time.Duration) *PostLolMatchmakingV1ReadyCheckAcceptParams {

	return &PostLolMatchmakingV1ReadyCheckAcceptParams{

		timeout: timeout,
	}
}

// NewPostLolMatchmakingV1ReadyCheckAcceptParamsWithContext creates a new PostLolMatchmakingV1ReadyCheckAcceptParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolMatchmakingV1ReadyCheckAcceptParamsWithContext(ctx context.Context) *PostLolMatchmakingV1ReadyCheckAcceptParams {

	return &PostLolMatchmakingV1ReadyCheckAcceptParams{

		Context: ctx,
	}
}

// NewPostLolMatchmakingV1ReadyCheckAcceptParamsWithHTTPClient creates a new PostLolMatchmakingV1ReadyCheckAcceptParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolMatchmakingV1ReadyCheckAcceptParamsWithHTTPClient(client *http.Client) *PostLolMatchmakingV1ReadyCheckAcceptParams {

	return &PostLolMatchmakingV1ReadyCheckAcceptParams{
		HTTPClient: client,
	}
}

/*PostLolMatchmakingV1ReadyCheckAcceptParams contains all the parameters to send to the API endpoint
for the post lol matchmaking v1 ready check accept operation typically these are written to a http.Request
*/
type PostLolMatchmakingV1ReadyCheckAcceptParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol matchmaking v1 ready check accept params
func (o *PostLolMatchmakingV1ReadyCheckAcceptParams) WithTimeout(timeout time.Duration) *PostLolMatchmakingV1ReadyCheckAcceptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol matchmaking v1 ready check accept params
func (o *PostLolMatchmakingV1ReadyCheckAcceptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol matchmaking v1 ready check accept params
func (o *PostLolMatchmakingV1ReadyCheckAcceptParams) WithContext(ctx context.Context) *PostLolMatchmakingV1ReadyCheckAcceptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol matchmaking v1 ready check accept params
func (o *PostLolMatchmakingV1ReadyCheckAcceptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol matchmaking v1 ready check accept params
func (o *PostLolMatchmakingV1ReadyCheckAcceptParams) WithHTTPClient(client *http.Client) *PostLolMatchmakingV1ReadyCheckAcceptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol matchmaking v1 ready check accept params
func (o *PostLolMatchmakingV1ReadyCheckAcceptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolMatchmakingV1ReadyCheckAcceptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}