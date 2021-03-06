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

// NewDeleteLolLoginV1SessionParams creates a new DeleteLolLoginV1SessionParams object
// with the default values initialized.
func NewDeleteLolLoginV1SessionParams() *DeleteLolLoginV1SessionParams {

	return &DeleteLolLoginV1SessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolLoginV1SessionParamsWithTimeout creates a new DeleteLolLoginV1SessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolLoginV1SessionParamsWithTimeout(timeout time.Duration) *DeleteLolLoginV1SessionParams {

	return &DeleteLolLoginV1SessionParams{

		timeout: timeout,
	}
}

// NewDeleteLolLoginV1SessionParamsWithContext creates a new DeleteLolLoginV1SessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolLoginV1SessionParamsWithContext(ctx context.Context) *DeleteLolLoginV1SessionParams {

	return &DeleteLolLoginV1SessionParams{

		Context: ctx,
	}
}

// NewDeleteLolLoginV1SessionParamsWithHTTPClient creates a new DeleteLolLoginV1SessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolLoginV1SessionParamsWithHTTPClient(client *http.Client) *DeleteLolLoginV1SessionParams {

	return &DeleteLolLoginV1SessionParams{
		HTTPClient: client,
	}
}

/*DeleteLolLoginV1SessionParams contains all the parameters to send to the API endpoint
for the delete lol login v1 session operation typically these are written to a http.Request
*/
type DeleteLolLoginV1SessionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol login v1 session params
func (o *DeleteLolLoginV1SessionParams) WithTimeout(timeout time.Duration) *DeleteLolLoginV1SessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol login v1 session params
func (o *DeleteLolLoginV1SessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol login v1 session params
func (o *DeleteLolLoginV1SessionParams) WithContext(ctx context.Context) *DeleteLolLoginV1SessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol login v1 session params
func (o *DeleteLolLoginV1SessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol login v1 session params
func (o *DeleteLolLoginV1SessionParams) WithHTTPClient(client *http.Client) *DeleteLolLoginV1SessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol login v1 session params
func (o *DeleteLolLoginV1SessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolLoginV1SessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
