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

// NewGetLolHonorV2V1RecognitionParams creates a new GetLolHonorV2V1RecognitionParams object
// with the default values initialized.
func NewGetLolHonorV2V1RecognitionParams() *GetLolHonorV2V1RecognitionParams {

	return &GetLolHonorV2V1RecognitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolHonorV2V1RecognitionParamsWithTimeout creates a new GetLolHonorV2V1RecognitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolHonorV2V1RecognitionParamsWithTimeout(timeout time.Duration) *GetLolHonorV2V1RecognitionParams {

	return &GetLolHonorV2V1RecognitionParams{

		timeout: timeout,
	}
}

// NewGetLolHonorV2V1RecognitionParamsWithContext creates a new GetLolHonorV2V1RecognitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolHonorV2V1RecognitionParamsWithContext(ctx context.Context) *GetLolHonorV2V1RecognitionParams {

	return &GetLolHonorV2V1RecognitionParams{

		Context: ctx,
	}
}

// NewGetLolHonorV2V1RecognitionParamsWithHTTPClient creates a new GetLolHonorV2V1RecognitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolHonorV2V1RecognitionParamsWithHTTPClient(client *http.Client) *GetLolHonorV2V1RecognitionParams {

	return &GetLolHonorV2V1RecognitionParams{
		HTTPClient: client,
	}
}

/*GetLolHonorV2V1RecognitionParams contains all the parameters to send to the API endpoint
for the get lol honor v2 v1 recognition operation typically these are written to a http.Request
*/
type GetLolHonorV2V1RecognitionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol honor v2 v1 recognition params
func (o *GetLolHonorV2V1RecognitionParams) WithTimeout(timeout time.Duration) *GetLolHonorV2V1RecognitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol honor v2 v1 recognition params
func (o *GetLolHonorV2V1RecognitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol honor v2 v1 recognition params
func (o *GetLolHonorV2V1RecognitionParams) WithContext(ctx context.Context) *GetLolHonorV2V1RecognitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol honor v2 v1 recognition params
func (o *GetLolHonorV2V1RecognitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol honor v2 v1 recognition params
func (o *GetLolHonorV2V1RecognitionParams) WithHTTPClient(client *http.Client) *GetLolHonorV2V1RecognitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol honor v2 v1 recognition params
func (o *GetLolHonorV2V1RecognitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolHonorV2V1RecognitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}