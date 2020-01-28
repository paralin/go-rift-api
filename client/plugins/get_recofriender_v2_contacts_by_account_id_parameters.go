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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRecofrienderV2ContactsByAccountIDParams creates a new GetRecofrienderV2ContactsByAccountIDParams object
// with the default values initialized.
func NewGetRecofrienderV2ContactsByAccountIDParams() *GetRecofrienderV2ContactsByAccountIDParams {
	var ()
	return &GetRecofrienderV2ContactsByAccountIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecofrienderV2ContactsByAccountIDParamsWithTimeout creates a new GetRecofrienderV2ContactsByAccountIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecofrienderV2ContactsByAccountIDParamsWithTimeout(timeout time.Duration) *GetRecofrienderV2ContactsByAccountIDParams {
	var ()
	return &GetRecofrienderV2ContactsByAccountIDParams{

		timeout: timeout,
	}
}

// NewGetRecofrienderV2ContactsByAccountIDParamsWithContext creates a new GetRecofrienderV2ContactsByAccountIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecofrienderV2ContactsByAccountIDParamsWithContext(ctx context.Context) *GetRecofrienderV2ContactsByAccountIDParams {
	var ()
	return &GetRecofrienderV2ContactsByAccountIDParams{

		Context: ctx,
	}
}

// NewGetRecofrienderV2ContactsByAccountIDParamsWithHTTPClient creates a new GetRecofrienderV2ContactsByAccountIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecofrienderV2ContactsByAccountIDParamsWithHTTPClient(client *http.Client) *GetRecofrienderV2ContactsByAccountIDParams {
	var ()
	return &GetRecofrienderV2ContactsByAccountIDParams{
		HTTPClient: client,
	}
}

/*GetRecofrienderV2ContactsByAccountIDParams contains all the parameters to send to the API endpoint
for the get recofriender v2 contacts by account Id operation typically these are written to a http.Request
*/
type GetRecofrienderV2ContactsByAccountIDParams struct {

	/*AccountID*/
	AccountID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recofriender v2 contacts by account Id params
func (o *GetRecofrienderV2ContactsByAccountIDParams) WithTimeout(timeout time.Duration) *GetRecofrienderV2ContactsByAccountIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recofriender v2 contacts by account Id params
func (o *GetRecofrienderV2ContactsByAccountIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recofriender v2 contacts by account Id params
func (o *GetRecofrienderV2ContactsByAccountIDParams) WithContext(ctx context.Context) *GetRecofrienderV2ContactsByAccountIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recofriender v2 contacts by account Id params
func (o *GetRecofrienderV2ContactsByAccountIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recofriender v2 contacts by account Id params
func (o *GetRecofrienderV2ContactsByAccountIDParams) WithHTTPClient(client *http.Client) *GetRecofrienderV2ContactsByAccountIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recofriender v2 contacts by account Id params
func (o *GetRecofrienderV2ContactsByAccountIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get recofriender v2 contacts by account Id params
func (o *GetRecofrienderV2ContactsByAccountIDParams) WithAccountID(accountID int64) *GetRecofrienderV2ContactsByAccountIDParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get recofriender v2 contacts by account Id params
func (o *GetRecofrienderV2ContactsByAccountIDParams) SetAccountID(accountID int64) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecofrienderV2ContactsByAccountIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt64(o.AccountID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}