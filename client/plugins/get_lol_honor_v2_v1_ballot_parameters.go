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

// NewGetLolHonorV2V1BallotParams creates a new GetLolHonorV2V1BallotParams object
// with the default values initialized.
func NewGetLolHonorV2V1BallotParams() *GetLolHonorV2V1BallotParams {

	return &GetLolHonorV2V1BallotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolHonorV2V1BallotParamsWithTimeout creates a new GetLolHonorV2V1BallotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolHonorV2V1BallotParamsWithTimeout(timeout time.Duration) *GetLolHonorV2V1BallotParams {

	return &GetLolHonorV2V1BallotParams{

		timeout: timeout,
	}
}

// NewGetLolHonorV2V1BallotParamsWithContext creates a new GetLolHonorV2V1BallotParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolHonorV2V1BallotParamsWithContext(ctx context.Context) *GetLolHonorV2V1BallotParams {

	return &GetLolHonorV2V1BallotParams{

		Context: ctx,
	}
}

// NewGetLolHonorV2V1BallotParamsWithHTTPClient creates a new GetLolHonorV2V1BallotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolHonorV2V1BallotParamsWithHTTPClient(client *http.Client) *GetLolHonorV2V1BallotParams {

	return &GetLolHonorV2V1BallotParams{
		HTTPClient: client,
	}
}

/*GetLolHonorV2V1BallotParams contains all the parameters to send to the API endpoint
for the get lol honor v2 v1 ballot operation typically these are written to a http.Request
*/
type GetLolHonorV2V1BallotParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol honor v2 v1 ballot params
func (o *GetLolHonorV2V1BallotParams) WithTimeout(timeout time.Duration) *GetLolHonorV2V1BallotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol honor v2 v1 ballot params
func (o *GetLolHonorV2V1BallotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol honor v2 v1 ballot params
func (o *GetLolHonorV2V1BallotParams) WithContext(ctx context.Context) *GetLolHonorV2V1BallotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol honor v2 v1 ballot params
func (o *GetLolHonorV2V1BallotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol honor v2 v1 ballot params
func (o *GetLolHonorV2V1BallotParams) WithHTTPClient(client *http.Client) *GetLolHonorV2V1BallotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol honor v2 v1 ballot params
func (o *GetLolHonorV2V1BallotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolHonorV2V1BallotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}