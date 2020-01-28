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

// NewPostLolAccountVerificationV1VerifyParams creates a new PostLolAccountVerificationV1VerifyParams object
// with the default values initialized.
func NewPostLolAccountVerificationV1VerifyParams() *PostLolAccountVerificationV1VerifyParams {
	var ()
	return &PostLolAccountVerificationV1VerifyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolAccountVerificationV1VerifyParamsWithTimeout creates a new PostLolAccountVerificationV1VerifyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolAccountVerificationV1VerifyParamsWithTimeout(timeout time.Duration) *PostLolAccountVerificationV1VerifyParams {
	var ()
	return &PostLolAccountVerificationV1VerifyParams{

		timeout: timeout,
	}
}

// NewPostLolAccountVerificationV1VerifyParamsWithContext creates a new PostLolAccountVerificationV1VerifyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolAccountVerificationV1VerifyParamsWithContext(ctx context.Context) *PostLolAccountVerificationV1VerifyParams {
	var ()
	return &PostLolAccountVerificationV1VerifyParams{

		Context: ctx,
	}
}

// NewPostLolAccountVerificationV1VerifyParamsWithHTTPClient creates a new PostLolAccountVerificationV1VerifyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolAccountVerificationV1VerifyParamsWithHTTPClient(client *http.Client) *PostLolAccountVerificationV1VerifyParams {
	var ()
	return &PostLolAccountVerificationV1VerifyParams{
		HTTPClient: client,
	}
}

/*PostLolAccountVerificationV1VerifyParams contains all the parameters to send to the API endpoint
for the post lol account verification v1 verify operation typically these are written to a http.Request
*/
type PostLolAccountVerificationV1VerifyParams struct {

	/*VerifyRequest*/
	VerifyRequest *models.LolAccountVerificationVerifyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol account verification v1 verify params
func (o *PostLolAccountVerificationV1VerifyParams) WithTimeout(timeout time.Duration) *PostLolAccountVerificationV1VerifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol account verification v1 verify params
func (o *PostLolAccountVerificationV1VerifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol account verification v1 verify params
func (o *PostLolAccountVerificationV1VerifyParams) WithContext(ctx context.Context) *PostLolAccountVerificationV1VerifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol account verification v1 verify params
func (o *PostLolAccountVerificationV1VerifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol account verification v1 verify params
func (o *PostLolAccountVerificationV1VerifyParams) WithHTTPClient(client *http.Client) *PostLolAccountVerificationV1VerifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol account verification v1 verify params
func (o *PostLolAccountVerificationV1VerifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVerifyRequest adds the verifyRequest to the post lol account verification v1 verify params
func (o *PostLolAccountVerificationV1VerifyParams) WithVerifyRequest(verifyRequest *models.LolAccountVerificationVerifyRequest) *PostLolAccountVerificationV1VerifyParams {
	o.SetVerifyRequest(verifyRequest)
	return o
}

// SetVerifyRequest adds the verifyRequest to the post lol account verification v1 verify params
func (o *PostLolAccountVerificationV1VerifyParams) SetVerifyRequest(verifyRequest *models.LolAccountVerificationVerifyRequest) {
	o.VerifyRequest = verifyRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolAccountVerificationV1VerifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.VerifyRequest != nil {
		if err := r.SetBodyParam(o.VerifyRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
