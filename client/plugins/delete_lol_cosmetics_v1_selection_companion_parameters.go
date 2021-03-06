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

// NewDeleteLolCosmeticsV1SelectionCompanionParams creates a new DeleteLolCosmeticsV1SelectionCompanionParams object
// with the default values initialized.
func NewDeleteLolCosmeticsV1SelectionCompanionParams() *DeleteLolCosmeticsV1SelectionCompanionParams {

	return &DeleteLolCosmeticsV1SelectionCompanionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolCosmeticsV1SelectionCompanionParamsWithTimeout creates a new DeleteLolCosmeticsV1SelectionCompanionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolCosmeticsV1SelectionCompanionParamsWithTimeout(timeout time.Duration) *DeleteLolCosmeticsV1SelectionCompanionParams {

	return &DeleteLolCosmeticsV1SelectionCompanionParams{

		timeout: timeout,
	}
}

// NewDeleteLolCosmeticsV1SelectionCompanionParamsWithContext creates a new DeleteLolCosmeticsV1SelectionCompanionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolCosmeticsV1SelectionCompanionParamsWithContext(ctx context.Context) *DeleteLolCosmeticsV1SelectionCompanionParams {

	return &DeleteLolCosmeticsV1SelectionCompanionParams{

		Context: ctx,
	}
}

// NewDeleteLolCosmeticsV1SelectionCompanionParamsWithHTTPClient creates a new DeleteLolCosmeticsV1SelectionCompanionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolCosmeticsV1SelectionCompanionParamsWithHTTPClient(client *http.Client) *DeleteLolCosmeticsV1SelectionCompanionParams {

	return &DeleteLolCosmeticsV1SelectionCompanionParams{
		HTTPClient: client,
	}
}

/*DeleteLolCosmeticsV1SelectionCompanionParams contains all the parameters to send to the API endpoint
for the delete lol cosmetics v1 selection companion operation typically these are written to a http.Request
*/
type DeleteLolCosmeticsV1SelectionCompanionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol cosmetics v1 selection companion params
func (o *DeleteLolCosmeticsV1SelectionCompanionParams) WithTimeout(timeout time.Duration) *DeleteLolCosmeticsV1SelectionCompanionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol cosmetics v1 selection companion params
func (o *DeleteLolCosmeticsV1SelectionCompanionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol cosmetics v1 selection companion params
func (o *DeleteLolCosmeticsV1SelectionCompanionParams) WithContext(ctx context.Context) *DeleteLolCosmeticsV1SelectionCompanionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol cosmetics v1 selection companion params
func (o *DeleteLolCosmeticsV1SelectionCompanionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol cosmetics v1 selection companion params
func (o *DeleteLolCosmeticsV1SelectionCompanionParams) WithHTTPClient(client *http.Client) *DeleteLolCosmeticsV1SelectionCompanionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol cosmetics v1 selection companion params
func (o *DeleteLolCosmeticsV1SelectionCompanionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolCosmeticsV1SelectionCompanionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
