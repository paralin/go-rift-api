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

// NewDeleteRiotMessagingServiceV1SessionParams creates a new DeleteRiotMessagingServiceV1SessionParams object
// with the default values initialized.
func NewDeleteRiotMessagingServiceV1SessionParams() *DeleteRiotMessagingServiceV1SessionParams {

	return &DeleteRiotMessagingServiceV1SessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRiotMessagingServiceV1SessionParamsWithTimeout creates a new DeleteRiotMessagingServiceV1SessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRiotMessagingServiceV1SessionParamsWithTimeout(timeout time.Duration) *DeleteRiotMessagingServiceV1SessionParams {

	return &DeleteRiotMessagingServiceV1SessionParams{

		timeout: timeout,
	}
}

// NewDeleteRiotMessagingServiceV1SessionParamsWithContext creates a new DeleteRiotMessagingServiceV1SessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRiotMessagingServiceV1SessionParamsWithContext(ctx context.Context) *DeleteRiotMessagingServiceV1SessionParams {

	return &DeleteRiotMessagingServiceV1SessionParams{

		Context: ctx,
	}
}

// NewDeleteRiotMessagingServiceV1SessionParamsWithHTTPClient creates a new DeleteRiotMessagingServiceV1SessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRiotMessagingServiceV1SessionParamsWithHTTPClient(client *http.Client) *DeleteRiotMessagingServiceV1SessionParams {

	return &DeleteRiotMessagingServiceV1SessionParams{
		HTTPClient: client,
	}
}

/*DeleteRiotMessagingServiceV1SessionParams contains all the parameters to send to the API endpoint
for the delete riot messaging service v1 session operation typically these are written to a http.Request
*/
type DeleteRiotMessagingServiceV1SessionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete riot messaging service v1 session params
func (o *DeleteRiotMessagingServiceV1SessionParams) WithTimeout(timeout time.Duration) *DeleteRiotMessagingServiceV1SessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete riot messaging service v1 session params
func (o *DeleteRiotMessagingServiceV1SessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete riot messaging service v1 session params
func (o *DeleteRiotMessagingServiceV1SessionParams) WithContext(ctx context.Context) *DeleteRiotMessagingServiceV1SessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete riot messaging service v1 session params
func (o *DeleteRiotMessagingServiceV1SessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete riot messaging service v1 session params
func (o *DeleteRiotMessagingServiceV1SessionParams) WithHTTPClient(client *http.Client) *DeleteRiotMessagingServiceV1SessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete riot messaging service v1 session params
func (o *DeleteRiotMessagingServiceV1SessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRiotMessagingServiceV1SessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}