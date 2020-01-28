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

// NewDeleteRecofrienderV2ContactsParams creates a new DeleteRecofrienderV2ContactsParams object
// with the default values initialized.
func NewDeleteRecofrienderV2ContactsParams() *DeleteRecofrienderV2ContactsParams {

	return &DeleteRecofrienderV2ContactsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRecofrienderV2ContactsParamsWithTimeout creates a new DeleteRecofrienderV2ContactsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRecofrienderV2ContactsParamsWithTimeout(timeout time.Duration) *DeleteRecofrienderV2ContactsParams {

	return &DeleteRecofrienderV2ContactsParams{

		timeout: timeout,
	}
}

// NewDeleteRecofrienderV2ContactsParamsWithContext creates a new DeleteRecofrienderV2ContactsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRecofrienderV2ContactsParamsWithContext(ctx context.Context) *DeleteRecofrienderV2ContactsParams {

	return &DeleteRecofrienderV2ContactsParams{

		Context: ctx,
	}
}

// NewDeleteRecofrienderV2ContactsParamsWithHTTPClient creates a new DeleteRecofrienderV2ContactsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRecofrienderV2ContactsParamsWithHTTPClient(client *http.Client) *DeleteRecofrienderV2ContactsParams {

	return &DeleteRecofrienderV2ContactsParams{
		HTTPClient: client,
	}
}

/*DeleteRecofrienderV2ContactsParams contains all the parameters to send to the API endpoint
for the delete recofriender v2 contacts operation typically these are written to a http.Request
*/
type DeleteRecofrienderV2ContactsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete recofriender v2 contacts params
func (o *DeleteRecofrienderV2ContactsParams) WithTimeout(timeout time.Duration) *DeleteRecofrienderV2ContactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete recofriender v2 contacts params
func (o *DeleteRecofrienderV2ContactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete recofriender v2 contacts params
func (o *DeleteRecofrienderV2ContactsParams) WithContext(ctx context.Context) *DeleteRecofrienderV2ContactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete recofriender v2 contacts params
func (o *DeleteRecofrienderV2ContactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete recofriender v2 contacts params
func (o *DeleteRecofrienderV2ContactsParams) WithHTTPClient(client *http.Client) *DeleteRecofrienderV2ContactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete recofriender v2 contacts params
func (o *DeleteRecofrienderV2ContactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRecofrienderV2ContactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
