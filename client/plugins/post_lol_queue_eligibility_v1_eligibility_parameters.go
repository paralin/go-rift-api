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

// NewPostLolQueueEligibilityV1EligibilityParams creates a new PostLolQueueEligibilityV1EligibilityParams object
// with the default values initialized.
func NewPostLolQueueEligibilityV1EligibilityParams() *PostLolQueueEligibilityV1EligibilityParams {
	var ()
	return &PostLolQueueEligibilityV1EligibilityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolQueueEligibilityV1EligibilityParamsWithTimeout creates a new PostLolQueueEligibilityV1EligibilityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolQueueEligibilityV1EligibilityParamsWithTimeout(timeout time.Duration) *PostLolQueueEligibilityV1EligibilityParams {
	var ()
	return &PostLolQueueEligibilityV1EligibilityParams{

		timeout: timeout,
	}
}

// NewPostLolQueueEligibilityV1EligibilityParamsWithContext creates a new PostLolQueueEligibilityV1EligibilityParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolQueueEligibilityV1EligibilityParamsWithContext(ctx context.Context) *PostLolQueueEligibilityV1EligibilityParams {
	var ()
	return &PostLolQueueEligibilityV1EligibilityParams{

		Context: ctx,
	}
}

// NewPostLolQueueEligibilityV1EligibilityParamsWithHTTPClient creates a new PostLolQueueEligibilityV1EligibilityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolQueueEligibilityV1EligibilityParamsWithHTTPClient(client *http.Client) *PostLolQueueEligibilityV1EligibilityParams {
	var ()
	return &PostLolQueueEligibilityV1EligibilityParams{
		HTTPClient: client,
	}
}

/*PostLolQueueEligibilityV1EligibilityParams contains all the parameters to send to the API endpoint
for the post lol queue eligibility v1 eligibility operation typically these are written to a http.Request
*/
type PostLolQueueEligibilityV1EligibilityParams struct {

	/*EligibilityQueryParam*/
	EligibilityQueryParam *models.LolQueueEligibilityEligibilityQueryParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol queue eligibility v1 eligibility params
func (o *PostLolQueueEligibilityV1EligibilityParams) WithTimeout(timeout time.Duration) *PostLolQueueEligibilityV1EligibilityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol queue eligibility v1 eligibility params
func (o *PostLolQueueEligibilityV1EligibilityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol queue eligibility v1 eligibility params
func (o *PostLolQueueEligibilityV1EligibilityParams) WithContext(ctx context.Context) *PostLolQueueEligibilityV1EligibilityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol queue eligibility v1 eligibility params
func (o *PostLolQueueEligibilityV1EligibilityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol queue eligibility v1 eligibility params
func (o *PostLolQueueEligibilityV1EligibilityParams) WithHTTPClient(client *http.Client) *PostLolQueueEligibilityV1EligibilityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol queue eligibility v1 eligibility params
func (o *PostLolQueueEligibilityV1EligibilityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEligibilityQueryParam adds the eligibilityQueryParam to the post lol queue eligibility v1 eligibility params
func (o *PostLolQueueEligibilityV1EligibilityParams) WithEligibilityQueryParam(eligibilityQueryParam *models.LolQueueEligibilityEligibilityQueryParams) *PostLolQueueEligibilityV1EligibilityParams {
	o.SetEligibilityQueryParam(eligibilityQueryParam)
	return o
}

// SetEligibilityQueryParam adds the eligibilityQueryParam to the post lol queue eligibility v1 eligibility params
func (o *PostLolQueueEligibilityV1EligibilityParams) SetEligibilityQueryParam(eligibilityQueryParam *models.LolQueueEligibilityEligibilityQueryParams) {
	o.EligibilityQueryParam = eligibilityQueryParam
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolQueueEligibilityV1EligibilityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EligibilityQueryParam != nil {
		if err := r.SetBodyParam(o.EligibilityQueryParam); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
