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

// NewPutLolEmailVerificationV1EmailParams creates a new PutLolEmailVerificationV1EmailParams object
// with the default values initialized.
func NewPutLolEmailVerificationV1EmailParams() *PutLolEmailVerificationV1EmailParams {
	var ()
	return &PutLolEmailVerificationV1EmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLolEmailVerificationV1EmailParamsWithTimeout creates a new PutLolEmailVerificationV1EmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLolEmailVerificationV1EmailParamsWithTimeout(timeout time.Duration) *PutLolEmailVerificationV1EmailParams {
	var ()
	return &PutLolEmailVerificationV1EmailParams{

		timeout: timeout,
	}
}

// NewPutLolEmailVerificationV1EmailParamsWithContext creates a new PutLolEmailVerificationV1EmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLolEmailVerificationV1EmailParamsWithContext(ctx context.Context) *PutLolEmailVerificationV1EmailParams {
	var ()
	return &PutLolEmailVerificationV1EmailParams{

		Context: ctx,
	}
}

// NewPutLolEmailVerificationV1EmailParamsWithHTTPClient creates a new PutLolEmailVerificationV1EmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLolEmailVerificationV1EmailParamsWithHTTPClient(client *http.Client) *PutLolEmailVerificationV1EmailParams {
	var ()
	return &PutLolEmailVerificationV1EmailParams{
		HTTPClient: client,
	}
}

/*PutLolEmailVerificationV1EmailParams contains all the parameters to send to the API endpoint
for the put lol email verification v1 email operation typically these are written to a http.Request
*/
type PutLolEmailVerificationV1EmailParams struct {

	/*EmailUpdate*/
	EmailUpdate *models.LolEmailVerificationEmailUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lol email verification v1 email params
func (o *PutLolEmailVerificationV1EmailParams) WithTimeout(timeout time.Duration) *PutLolEmailVerificationV1EmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lol email verification v1 email params
func (o *PutLolEmailVerificationV1EmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lol email verification v1 email params
func (o *PutLolEmailVerificationV1EmailParams) WithContext(ctx context.Context) *PutLolEmailVerificationV1EmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lol email verification v1 email params
func (o *PutLolEmailVerificationV1EmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lol email verification v1 email params
func (o *PutLolEmailVerificationV1EmailParams) WithHTTPClient(client *http.Client) *PutLolEmailVerificationV1EmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lol email verification v1 email params
func (o *PutLolEmailVerificationV1EmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailUpdate adds the emailUpdate to the put lol email verification v1 email params
func (o *PutLolEmailVerificationV1EmailParams) WithEmailUpdate(emailUpdate *models.LolEmailVerificationEmailUpdate) *PutLolEmailVerificationV1EmailParams {
	o.SetEmailUpdate(emailUpdate)
	return o
}

// SetEmailUpdate adds the emailUpdate to the put lol email verification v1 email params
func (o *PutLolEmailVerificationV1EmailParams) SetEmailUpdate(emailUpdate *models.LolEmailVerificationEmailUpdate) {
	o.EmailUpdate = emailUpdate
}

// WriteToRequest writes these params to a swagger request
func (o *PutLolEmailVerificationV1EmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EmailUpdate != nil {
		if err := r.SetBodyParam(o.EmailUpdate); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
