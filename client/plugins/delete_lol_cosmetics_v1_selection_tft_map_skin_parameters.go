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

// NewDeleteLolCosmeticsV1SelectionTftMapSkinParams creates a new DeleteLolCosmeticsV1SelectionTftMapSkinParams object
// with the default values initialized.
func NewDeleteLolCosmeticsV1SelectionTftMapSkinParams() *DeleteLolCosmeticsV1SelectionTftMapSkinParams {

	return &DeleteLolCosmeticsV1SelectionTftMapSkinParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolCosmeticsV1SelectionTftMapSkinParamsWithTimeout creates a new DeleteLolCosmeticsV1SelectionTftMapSkinParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolCosmeticsV1SelectionTftMapSkinParamsWithTimeout(timeout time.Duration) *DeleteLolCosmeticsV1SelectionTftMapSkinParams {

	return &DeleteLolCosmeticsV1SelectionTftMapSkinParams{

		timeout: timeout,
	}
}

// NewDeleteLolCosmeticsV1SelectionTftMapSkinParamsWithContext creates a new DeleteLolCosmeticsV1SelectionTftMapSkinParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolCosmeticsV1SelectionTftMapSkinParamsWithContext(ctx context.Context) *DeleteLolCosmeticsV1SelectionTftMapSkinParams {

	return &DeleteLolCosmeticsV1SelectionTftMapSkinParams{

		Context: ctx,
	}
}

// NewDeleteLolCosmeticsV1SelectionTftMapSkinParamsWithHTTPClient creates a new DeleteLolCosmeticsV1SelectionTftMapSkinParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolCosmeticsV1SelectionTftMapSkinParamsWithHTTPClient(client *http.Client) *DeleteLolCosmeticsV1SelectionTftMapSkinParams {

	return &DeleteLolCosmeticsV1SelectionTftMapSkinParams{
		HTTPClient: client,
	}
}

/*DeleteLolCosmeticsV1SelectionTftMapSkinParams contains all the parameters to send to the API endpoint
for the delete lol cosmetics v1 selection tft map skin operation typically these are written to a http.Request
*/
type DeleteLolCosmeticsV1SelectionTftMapSkinParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol cosmetics v1 selection tft map skin params
func (o *DeleteLolCosmeticsV1SelectionTftMapSkinParams) WithTimeout(timeout time.Duration) *DeleteLolCosmeticsV1SelectionTftMapSkinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol cosmetics v1 selection tft map skin params
func (o *DeleteLolCosmeticsV1SelectionTftMapSkinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol cosmetics v1 selection tft map skin params
func (o *DeleteLolCosmeticsV1SelectionTftMapSkinParams) WithContext(ctx context.Context) *DeleteLolCosmeticsV1SelectionTftMapSkinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol cosmetics v1 selection tft map skin params
func (o *DeleteLolCosmeticsV1SelectionTftMapSkinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol cosmetics v1 selection tft map skin params
func (o *DeleteLolCosmeticsV1SelectionTftMapSkinParams) WithHTTPClient(client *http.Client) *DeleteLolCosmeticsV1SelectionTftMapSkinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol cosmetics v1 selection tft map skin params
func (o *DeleteLolCosmeticsV1SelectionTftMapSkinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolCosmeticsV1SelectionTftMapSkinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}