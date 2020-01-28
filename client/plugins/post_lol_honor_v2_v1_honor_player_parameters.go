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

	models "github.com/paralin/go-rift-api/models"
)

// NewPostLolHonorV2V1HonorPlayerParams creates a new PostLolHonorV2V1HonorPlayerParams object
// with the default values initialized.
func NewPostLolHonorV2V1HonorPlayerParams() *PostLolHonorV2V1HonorPlayerParams {
	var ()
	return &PostLolHonorV2V1HonorPlayerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolHonorV2V1HonorPlayerParamsWithTimeout creates a new PostLolHonorV2V1HonorPlayerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolHonorV2V1HonorPlayerParamsWithTimeout(timeout time.Duration) *PostLolHonorV2V1HonorPlayerParams {
	var ()
	return &PostLolHonorV2V1HonorPlayerParams{

		timeout: timeout,
	}
}

// NewPostLolHonorV2V1HonorPlayerParamsWithContext creates a new PostLolHonorV2V1HonorPlayerParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolHonorV2V1HonorPlayerParamsWithContext(ctx context.Context) *PostLolHonorV2V1HonorPlayerParams {
	var ()
	return &PostLolHonorV2V1HonorPlayerParams{

		Context: ctx,
	}
}

// NewPostLolHonorV2V1HonorPlayerParamsWithHTTPClient creates a new PostLolHonorV2V1HonorPlayerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolHonorV2V1HonorPlayerParamsWithHTTPClient(client *http.Client) *PostLolHonorV2V1HonorPlayerParams {
	var ()
	return &PostLolHonorV2V1HonorPlayerParams{
		HTTPClient: client,
	}
}

/*PostLolHonorV2V1HonorPlayerParams contains all the parameters to send to the API endpoint
for the post lol honor v2 v1 honor player operation typically these are written to a http.Request
*/
type PostLolHonorV2V1HonorPlayerParams struct {

	/*HonorPlayerRequest*/
	HonorPlayerRequest *models.LolHonorV2APIHonorPlayerServerRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol honor v2 v1 honor player params
func (o *PostLolHonorV2V1HonorPlayerParams) WithTimeout(timeout time.Duration) *PostLolHonorV2V1HonorPlayerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol honor v2 v1 honor player params
func (o *PostLolHonorV2V1HonorPlayerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol honor v2 v1 honor player params
func (o *PostLolHonorV2V1HonorPlayerParams) WithContext(ctx context.Context) *PostLolHonorV2V1HonorPlayerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol honor v2 v1 honor player params
func (o *PostLolHonorV2V1HonorPlayerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol honor v2 v1 honor player params
func (o *PostLolHonorV2V1HonorPlayerParams) WithHTTPClient(client *http.Client) *PostLolHonorV2V1HonorPlayerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol honor v2 v1 honor player params
func (o *PostLolHonorV2V1HonorPlayerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHonorPlayerRequest adds the honorPlayerRequest to the post lol honor v2 v1 honor player params
func (o *PostLolHonorV2V1HonorPlayerParams) WithHonorPlayerRequest(honorPlayerRequest *models.LolHonorV2APIHonorPlayerServerRequest) *PostLolHonorV2V1HonorPlayerParams {
	o.SetHonorPlayerRequest(honorPlayerRequest)
	return o
}

// SetHonorPlayerRequest adds the honorPlayerRequest to the post lol honor v2 v1 honor player params
func (o *PostLolHonorV2V1HonorPlayerParams) SetHonorPlayerRequest(honorPlayerRequest *models.LolHonorV2APIHonorPlayerServerRequest) {
	o.HonorPlayerRequest = honorPlayerRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolHonorV2V1HonorPlayerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HonorPlayerRequest != nil {
		if err := r.SetBodyParam(o.HonorPlayerRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
