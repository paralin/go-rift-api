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

// NewGetLolGameflowV1GameflowPhaseParams creates a new GetLolGameflowV1GameflowPhaseParams object
// with the default values initialized.
func NewGetLolGameflowV1GameflowPhaseParams() *GetLolGameflowV1GameflowPhaseParams {

	return &GetLolGameflowV1GameflowPhaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolGameflowV1GameflowPhaseParamsWithTimeout creates a new GetLolGameflowV1GameflowPhaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolGameflowV1GameflowPhaseParamsWithTimeout(timeout time.Duration) *GetLolGameflowV1GameflowPhaseParams {

	return &GetLolGameflowV1GameflowPhaseParams{

		timeout: timeout,
	}
}

// NewGetLolGameflowV1GameflowPhaseParamsWithContext creates a new GetLolGameflowV1GameflowPhaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolGameflowV1GameflowPhaseParamsWithContext(ctx context.Context) *GetLolGameflowV1GameflowPhaseParams {

	return &GetLolGameflowV1GameflowPhaseParams{

		Context: ctx,
	}
}

// NewGetLolGameflowV1GameflowPhaseParamsWithHTTPClient creates a new GetLolGameflowV1GameflowPhaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolGameflowV1GameflowPhaseParamsWithHTTPClient(client *http.Client) *GetLolGameflowV1GameflowPhaseParams {

	return &GetLolGameflowV1GameflowPhaseParams{
		HTTPClient: client,
	}
}

/*GetLolGameflowV1GameflowPhaseParams contains all the parameters to send to the API endpoint
for the get lol gameflow v1 gameflow phase operation typically these are written to a http.Request
*/
type GetLolGameflowV1GameflowPhaseParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol gameflow v1 gameflow phase params
func (o *GetLolGameflowV1GameflowPhaseParams) WithTimeout(timeout time.Duration) *GetLolGameflowV1GameflowPhaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol gameflow v1 gameflow phase params
func (o *GetLolGameflowV1GameflowPhaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol gameflow v1 gameflow phase params
func (o *GetLolGameflowV1GameflowPhaseParams) WithContext(ctx context.Context) *GetLolGameflowV1GameflowPhaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol gameflow v1 gameflow phase params
func (o *GetLolGameflowV1GameflowPhaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol gameflow v1 gameflow phase params
func (o *GetLolGameflowV1GameflowPhaseParams) WithHTTPClient(client *http.Client) *GetLolGameflowV1GameflowPhaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol gameflow v1 gameflow phase params
func (o *GetLolGameflowV1GameflowPhaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolGameflowV1GameflowPhaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
