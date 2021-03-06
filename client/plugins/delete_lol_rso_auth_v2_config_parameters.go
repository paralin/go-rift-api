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

// NewDeleteLolRsoAuthV2ConfigParams creates a new DeleteLolRsoAuthV2ConfigParams object
// with the default values initialized.
func NewDeleteLolRsoAuthV2ConfigParams() *DeleteLolRsoAuthV2ConfigParams {

	return &DeleteLolRsoAuthV2ConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolRsoAuthV2ConfigParamsWithTimeout creates a new DeleteLolRsoAuthV2ConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolRsoAuthV2ConfigParamsWithTimeout(timeout time.Duration) *DeleteLolRsoAuthV2ConfigParams {

	return &DeleteLolRsoAuthV2ConfigParams{

		timeout: timeout,
	}
}

// NewDeleteLolRsoAuthV2ConfigParamsWithContext creates a new DeleteLolRsoAuthV2ConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolRsoAuthV2ConfigParamsWithContext(ctx context.Context) *DeleteLolRsoAuthV2ConfigParams {

	return &DeleteLolRsoAuthV2ConfigParams{

		Context: ctx,
	}
}

// NewDeleteLolRsoAuthV2ConfigParamsWithHTTPClient creates a new DeleteLolRsoAuthV2ConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolRsoAuthV2ConfigParamsWithHTTPClient(client *http.Client) *DeleteLolRsoAuthV2ConfigParams {

	return &DeleteLolRsoAuthV2ConfigParams{
		HTTPClient: client,
	}
}

/*DeleteLolRsoAuthV2ConfigParams contains all the parameters to send to the API endpoint
for the delete lol rso auth v2 config operation typically these are written to a http.Request
*/
type DeleteLolRsoAuthV2ConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol rso auth v2 config params
func (o *DeleteLolRsoAuthV2ConfigParams) WithTimeout(timeout time.Duration) *DeleteLolRsoAuthV2ConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol rso auth v2 config params
func (o *DeleteLolRsoAuthV2ConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol rso auth v2 config params
func (o *DeleteLolRsoAuthV2ConfigParams) WithContext(ctx context.Context) *DeleteLolRsoAuthV2ConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol rso auth v2 config params
func (o *DeleteLolRsoAuthV2ConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol rso auth v2 config params
func (o *DeleteLolRsoAuthV2ConfigParams) WithHTTPClient(client *http.Client) *DeleteLolRsoAuthV2ConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol rso auth v2 config params
func (o *DeleteLolRsoAuthV2ConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolRsoAuthV2ConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
