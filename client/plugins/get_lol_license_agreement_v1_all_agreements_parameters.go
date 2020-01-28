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

// NewGetLolLicenseAgreementV1AllAgreementsParams creates a new GetLolLicenseAgreementV1AllAgreementsParams object
// with the default values initialized.
func NewGetLolLicenseAgreementV1AllAgreementsParams() *GetLolLicenseAgreementV1AllAgreementsParams {

	return &GetLolLicenseAgreementV1AllAgreementsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLicenseAgreementV1AllAgreementsParamsWithTimeout creates a new GetLolLicenseAgreementV1AllAgreementsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLicenseAgreementV1AllAgreementsParamsWithTimeout(timeout time.Duration) *GetLolLicenseAgreementV1AllAgreementsParams {

	return &GetLolLicenseAgreementV1AllAgreementsParams{

		timeout: timeout,
	}
}

// NewGetLolLicenseAgreementV1AllAgreementsParamsWithContext creates a new GetLolLicenseAgreementV1AllAgreementsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLicenseAgreementV1AllAgreementsParamsWithContext(ctx context.Context) *GetLolLicenseAgreementV1AllAgreementsParams {

	return &GetLolLicenseAgreementV1AllAgreementsParams{

		Context: ctx,
	}
}

// NewGetLolLicenseAgreementV1AllAgreementsParamsWithHTTPClient creates a new GetLolLicenseAgreementV1AllAgreementsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLicenseAgreementV1AllAgreementsParamsWithHTTPClient(client *http.Client) *GetLolLicenseAgreementV1AllAgreementsParams {

	return &GetLolLicenseAgreementV1AllAgreementsParams{
		HTTPClient: client,
	}
}

/*GetLolLicenseAgreementV1AllAgreementsParams contains all the parameters to send to the API endpoint
for the get lol license agreement v1 all agreements operation typically these are written to a http.Request
*/
type GetLolLicenseAgreementV1AllAgreementsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol license agreement v1 all agreements params
func (o *GetLolLicenseAgreementV1AllAgreementsParams) WithTimeout(timeout time.Duration) *GetLolLicenseAgreementV1AllAgreementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol license agreement v1 all agreements params
func (o *GetLolLicenseAgreementV1AllAgreementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol license agreement v1 all agreements params
func (o *GetLolLicenseAgreementV1AllAgreementsParams) WithContext(ctx context.Context) *GetLolLicenseAgreementV1AllAgreementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol license agreement v1 all agreements params
func (o *GetLolLicenseAgreementV1AllAgreementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol license agreement v1 all agreements params
func (o *GetLolLicenseAgreementV1AllAgreementsParams) WithHTTPClient(client *http.Client) *GetLolLicenseAgreementV1AllAgreementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol license agreement v1 all agreements params
func (o *GetLolLicenseAgreementV1AllAgreementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLicenseAgreementV1AllAgreementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}