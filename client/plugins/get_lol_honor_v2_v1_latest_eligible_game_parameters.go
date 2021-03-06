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

// NewGetLolHonorV2V1LatestEligibleGameParams creates a new GetLolHonorV2V1LatestEligibleGameParams object
// with the default values initialized.
func NewGetLolHonorV2V1LatestEligibleGameParams() *GetLolHonorV2V1LatestEligibleGameParams {

	return &GetLolHonorV2V1LatestEligibleGameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolHonorV2V1LatestEligibleGameParamsWithTimeout creates a new GetLolHonorV2V1LatestEligibleGameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolHonorV2V1LatestEligibleGameParamsWithTimeout(timeout time.Duration) *GetLolHonorV2V1LatestEligibleGameParams {

	return &GetLolHonorV2V1LatestEligibleGameParams{

		timeout: timeout,
	}
}

// NewGetLolHonorV2V1LatestEligibleGameParamsWithContext creates a new GetLolHonorV2V1LatestEligibleGameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolHonorV2V1LatestEligibleGameParamsWithContext(ctx context.Context) *GetLolHonorV2V1LatestEligibleGameParams {

	return &GetLolHonorV2V1LatestEligibleGameParams{

		Context: ctx,
	}
}

// NewGetLolHonorV2V1LatestEligibleGameParamsWithHTTPClient creates a new GetLolHonorV2V1LatestEligibleGameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolHonorV2V1LatestEligibleGameParamsWithHTTPClient(client *http.Client) *GetLolHonorV2V1LatestEligibleGameParams {

	return &GetLolHonorV2V1LatestEligibleGameParams{
		HTTPClient: client,
	}
}

/*GetLolHonorV2V1LatestEligibleGameParams contains all the parameters to send to the API endpoint
for the get lol honor v2 v1 latest eligible game operation typically these are written to a http.Request
*/
type GetLolHonorV2V1LatestEligibleGameParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol honor v2 v1 latest eligible game params
func (o *GetLolHonorV2V1LatestEligibleGameParams) WithTimeout(timeout time.Duration) *GetLolHonorV2V1LatestEligibleGameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol honor v2 v1 latest eligible game params
func (o *GetLolHonorV2V1LatestEligibleGameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol honor v2 v1 latest eligible game params
func (o *GetLolHonorV2V1LatestEligibleGameParams) WithContext(ctx context.Context) *GetLolHonorV2V1LatestEligibleGameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol honor v2 v1 latest eligible game params
func (o *GetLolHonorV2V1LatestEligibleGameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol honor v2 v1 latest eligible game params
func (o *GetLolHonorV2V1LatestEligibleGameParams) WithHTTPClient(client *http.Client) *GetLolHonorV2V1LatestEligibleGameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol honor v2 v1 latest eligible game params
func (o *GetLolHonorV2V1LatestEligibleGameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolHonorV2V1LatestEligibleGameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
