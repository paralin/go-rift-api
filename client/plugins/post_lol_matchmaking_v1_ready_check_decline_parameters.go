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

// NewPostLolMatchmakingV1ReadyCheckDeclineParams creates a new PostLolMatchmakingV1ReadyCheckDeclineParams object
// with the default values initialized.
func NewPostLolMatchmakingV1ReadyCheckDeclineParams() *PostLolMatchmakingV1ReadyCheckDeclineParams {

	return &PostLolMatchmakingV1ReadyCheckDeclineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolMatchmakingV1ReadyCheckDeclineParamsWithTimeout creates a new PostLolMatchmakingV1ReadyCheckDeclineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolMatchmakingV1ReadyCheckDeclineParamsWithTimeout(timeout time.Duration) *PostLolMatchmakingV1ReadyCheckDeclineParams {

	return &PostLolMatchmakingV1ReadyCheckDeclineParams{

		timeout: timeout,
	}
}

// NewPostLolMatchmakingV1ReadyCheckDeclineParamsWithContext creates a new PostLolMatchmakingV1ReadyCheckDeclineParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolMatchmakingV1ReadyCheckDeclineParamsWithContext(ctx context.Context) *PostLolMatchmakingV1ReadyCheckDeclineParams {

	return &PostLolMatchmakingV1ReadyCheckDeclineParams{

		Context: ctx,
	}
}

// NewPostLolMatchmakingV1ReadyCheckDeclineParamsWithHTTPClient creates a new PostLolMatchmakingV1ReadyCheckDeclineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolMatchmakingV1ReadyCheckDeclineParamsWithHTTPClient(client *http.Client) *PostLolMatchmakingV1ReadyCheckDeclineParams {

	return &PostLolMatchmakingV1ReadyCheckDeclineParams{
		HTTPClient: client,
	}
}

/*PostLolMatchmakingV1ReadyCheckDeclineParams contains all the parameters to send to the API endpoint
for the post lol matchmaking v1 ready check decline operation typically these are written to a http.Request
*/
type PostLolMatchmakingV1ReadyCheckDeclineParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol matchmaking v1 ready check decline params
func (o *PostLolMatchmakingV1ReadyCheckDeclineParams) WithTimeout(timeout time.Duration) *PostLolMatchmakingV1ReadyCheckDeclineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol matchmaking v1 ready check decline params
func (o *PostLolMatchmakingV1ReadyCheckDeclineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol matchmaking v1 ready check decline params
func (o *PostLolMatchmakingV1ReadyCheckDeclineParams) WithContext(ctx context.Context) *PostLolMatchmakingV1ReadyCheckDeclineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol matchmaking v1 ready check decline params
func (o *PostLolMatchmakingV1ReadyCheckDeclineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol matchmaking v1 ready check decline params
func (o *PostLolMatchmakingV1ReadyCheckDeclineParams) WithHTTPClient(client *http.Client) *PostLolMatchmakingV1ReadyCheckDeclineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol matchmaking v1 ready check decline params
func (o *PostLolMatchmakingV1ReadyCheckDeclineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolMatchmakingV1ReadyCheckDeclineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
