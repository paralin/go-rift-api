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

// NewDeleteRiotMessagingServiceV1ConnectParams creates a new DeleteRiotMessagingServiceV1ConnectParams object
// with the default values initialized.
func NewDeleteRiotMessagingServiceV1ConnectParams() *DeleteRiotMessagingServiceV1ConnectParams {

	return &DeleteRiotMessagingServiceV1ConnectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRiotMessagingServiceV1ConnectParamsWithTimeout creates a new DeleteRiotMessagingServiceV1ConnectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRiotMessagingServiceV1ConnectParamsWithTimeout(timeout time.Duration) *DeleteRiotMessagingServiceV1ConnectParams {

	return &DeleteRiotMessagingServiceV1ConnectParams{

		timeout: timeout,
	}
}

// NewDeleteRiotMessagingServiceV1ConnectParamsWithContext creates a new DeleteRiotMessagingServiceV1ConnectParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRiotMessagingServiceV1ConnectParamsWithContext(ctx context.Context) *DeleteRiotMessagingServiceV1ConnectParams {

	return &DeleteRiotMessagingServiceV1ConnectParams{

		Context: ctx,
	}
}

// NewDeleteRiotMessagingServiceV1ConnectParamsWithHTTPClient creates a new DeleteRiotMessagingServiceV1ConnectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRiotMessagingServiceV1ConnectParamsWithHTTPClient(client *http.Client) *DeleteRiotMessagingServiceV1ConnectParams {

	return &DeleteRiotMessagingServiceV1ConnectParams{
		HTTPClient: client,
	}
}

/*DeleteRiotMessagingServiceV1ConnectParams contains all the parameters to send to the API endpoint
for the delete riot messaging service v1 connect operation typically these are written to a http.Request
*/
type DeleteRiotMessagingServiceV1ConnectParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete riot messaging service v1 connect params
func (o *DeleteRiotMessagingServiceV1ConnectParams) WithTimeout(timeout time.Duration) *DeleteRiotMessagingServiceV1ConnectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete riot messaging service v1 connect params
func (o *DeleteRiotMessagingServiceV1ConnectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete riot messaging service v1 connect params
func (o *DeleteRiotMessagingServiceV1ConnectParams) WithContext(ctx context.Context) *DeleteRiotMessagingServiceV1ConnectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete riot messaging service v1 connect params
func (o *DeleteRiotMessagingServiceV1ConnectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete riot messaging service v1 connect params
func (o *DeleteRiotMessagingServiceV1ConnectParams) WithHTTPClient(client *http.Client) *DeleteRiotMessagingServiceV1ConnectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete riot messaging service v1 connect params
func (o *DeleteRiotMessagingServiceV1ConnectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRiotMessagingServiceV1ConnectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}