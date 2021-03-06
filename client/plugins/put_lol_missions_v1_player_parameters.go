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

// NewPutLolMissionsV1PlayerParams creates a new PutLolMissionsV1PlayerParams object
// with the default values initialized.
func NewPutLolMissionsV1PlayerParams() *PutLolMissionsV1PlayerParams {
	var ()
	return &PutLolMissionsV1PlayerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLolMissionsV1PlayerParamsWithTimeout creates a new PutLolMissionsV1PlayerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLolMissionsV1PlayerParamsWithTimeout(timeout time.Duration) *PutLolMissionsV1PlayerParams {
	var ()
	return &PutLolMissionsV1PlayerParams{

		timeout: timeout,
	}
}

// NewPutLolMissionsV1PlayerParamsWithContext creates a new PutLolMissionsV1PlayerParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLolMissionsV1PlayerParamsWithContext(ctx context.Context) *PutLolMissionsV1PlayerParams {
	var ()
	return &PutLolMissionsV1PlayerParams{

		Context: ctx,
	}
}

// NewPutLolMissionsV1PlayerParamsWithHTTPClient creates a new PutLolMissionsV1PlayerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLolMissionsV1PlayerParamsWithHTTPClient(client *http.Client) *PutLolMissionsV1PlayerParams {
	var ()
	return &PutLolMissionsV1PlayerParams{
		HTTPClient: client,
	}
}

/*PutLolMissionsV1PlayerParams contains all the parameters to send to the API endpoint
for the put lol missions v1 player operation typically these are written to a http.Request
*/
type PutLolMissionsV1PlayerParams struct {

	/*Ids*/
	Ids *models.IdsDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lol missions v1 player params
func (o *PutLolMissionsV1PlayerParams) WithTimeout(timeout time.Duration) *PutLolMissionsV1PlayerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lol missions v1 player params
func (o *PutLolMissionsV1PlayerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lol missions v1 player params
func (o *PutLolMissionsV1PlayerParams) WithContext(ctx context.Context) *PutLolMissionsV1PlayerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lol missions v1 player params
func (o *PutLolMissionsV1PlayerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lol missions v1 player params
func (o *PutLolMissionsV1PlayerParams) WithHTTPClient(client *http.Client) *PutLolMissionsV1PlayerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lol missions v1 player params
func (o *PutLolMissionsV1PlayerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the put lol missions v1 player params
func (o *PutLolMissionsV1PlayerParams) WithIds(ids *models.IdsDTO) *PutLolMissionsV1PlayerParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the put lol missions v1 player params
func (o *PutLolMissionsV1PlayerParams) SetIds(ids *models.IdsDTO) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *PutLolMissionsV1PlayerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ids != nil {
		if err := r.SetBodyParam(o.Ids); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
