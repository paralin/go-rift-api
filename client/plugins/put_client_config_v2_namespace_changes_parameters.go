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

// NewPutClientConfigV2NamespaceChangesParams creates a new PutClientConfigV2NamespaceChangesParams object
// with the default values initialized.
func NewPutClientConfigV2NamespaceChangesParams() *PutClientConfigV2NamespaceChangesParams {
	var ()
	return &PutClientConfigV2NamespaceChangesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutClientConfigV2NamespaceChangesParamsWithTimeout creates a new PutClientConfigV2NamespaceChangesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutClientConfigV2NamespaceChangesParamsWithTimeout(timeout time.Duration) *PutClientConfigV2NamespaceChangesParams {
	var ()
	return &PutClientConfigV2NamespaceChangesParams{

		timeout: timeout,
	}
}

// NewPutClientConfigV2NamespaceChangesParamsWithContext creates a new PutClientConfigV2NamespaceChangesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutClientConfigV2NamespaceChangesParamsWithContext(ctx context.Context) *PutClientConfigV2NamespaceChangesParams {
	var ()
	return &PutClientConfigV2NamespaceChangesParams{

		Context: ctx,
	}
}

// NewPutClientConfigV2NamespaceChangesParamsWithHTTPClient creates a new PutClientConfigV2NamespaceChangesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutClientConfigV2NamespaceChangesParamsWithHTTPClient(client *http.Client) *PutClientConfigV2NamespaceChangesParams {
	var ()
	return &PutClientConfigV2NamespaceChangesParams{
		HTTPClient: client,
	}
}

/*PutClientConfigV2NamespaceChangesParams contains all the parameters to send to the API endpoint
for the put client config v2 namespace changes operation typically these are written to a http.Request
*/
type PutClientConfigV2NamespaceChangesParams struct {

	/*Namespaces*/
	Namespaces *models.ClientConfigConfigNamespaceUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put client config v2 namespace changes params
func (o *PutClientConfigV2NamespaceChangesParams) WithTimeout(timeout time.Duration) *PutClientConfigV2NamespaceChangesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put client config v2 namespace changes params
func (o *PutClientConfigV2NamespaceChangesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put client config v2 namespace changes params
func (o *PutClientConfigV2NamespaceChangesParams) WithContext(ctx context.Context) *PutClientConfigV2NamespaceChangesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put client config v2 namespace changes params
func (o *PutClientConfigV2NamespaceChangesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put client config v2 namespace changes params
func (o *PutClientConfigV2NamespaceChangesParams) WithHTTPClient(client *http.Client) *PutClientConfigV2NamespaceChangesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put client config v2 namespace changes params
func (o *PutClientConfigV2NamespaceChangesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaces adds the namespaces to the put client config v2 namespace changes params
func (o *PutClientConfigV2NamespaceChangesParams) WithNamespaces(namespaces *models.ClientConfigConfigNamespaceUpdate) *PutClientConfigV2NamespaceChangesParams {
	o.SetNamespaces(namespaces)
	return o
}

// SetNamespaces adds the namespaces to the put client config v2 namespace changes params
func (o *PutClientConfigV2NamespaceChangesParams) SetNamespaces(namespaces *models.ClientConfigConfigNamespaceUpdate) {
	o.Namespaces = namespaces
}

// WriteToRequest writes these params to a swagger request
func (o *PutClientConfigV2NamespaceChangesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Namespaces != nil {
		if err := r.SetBodyParam(o.Namespaces); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}