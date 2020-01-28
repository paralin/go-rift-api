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

	models "github.com/paralin/go-rift-api/models"
)

// NewPutLolPerksV1PagesByIDParams creates a new PutLolPerksV1PagesByIDParams object
// with the default values initialized.
func NewPutLolPerksV1PagesByIDParams() *PutLolPerksV1PagesByIDParams {
	var ()
	return &PutLolPerksV1PagesByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLolPerksV1PagesByIDParamsWithTimeout creates a new PutLolPerksV1PagesByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLolPerksV1PagesByIDParamsWithTimeout(timeout time.Duration) *PutLolPerksV1PagesByIDParams {
	var ()
	return &PutLolPerksV1PagesByIDParams{

		timeout: timeout,
	}
}

// NewPutLolPerksV1PagesByIDParamsWithContext creates a new PutLolPerksV1PagesByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLolPerksV1PagesByIDParamsWithContext(ctx context.Context) *PutLolPerksV1PagesByIDParams {
	var ()
	return &PutLolPerksV1PagesByIDParams{

		Context: ctx,
	}
}

// NewPutLolPerksV1PagesByIDParamsWithHTTPClient creates a new PutLolPerksV1PagesByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLolPerksV1PagesByIDParamsWithHTTPClient(client *http.Client) *PutLolPerksV1PagesByIDParams {
	var ()
	return &PutLolPerksV1PagesByIDParams{
		HTTPClient: client,
	}
}

/*PutLolPerksV1PagesByIDParams contains all the parameters to send to the API endpoint
for the put lol perks v1 pages by Id operation typically these are written to a http.Request
*/
type PutLolPerksV1PagesByIDParams struct {

	/*ID*/
	ID int32
	/*Page*/
	Page *models.LolPerksPerkPageResource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lol perks v1 pages by Id params
func (o *PutLolPerksV1PagesByIDParams) WithTimeout(timeout time.Duration) *PutLolPerksV1PagesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lol perks v1 pages by Id params
func (o *PutLolPerksV1PagesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lol perks v1 pages by Id params
func (o *PutLolPerksV1PagesByIDParams) WithContext(ctx context.Context) *PutLolPerksV1PagesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lol perks v1 pages by Id params
func (o *PutLolPerksV1PagesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lol perks v1 pages by Id params
func (o *PutLolPerksV1PagesByIDParams) WithHTTPClient(client *http.Client) *PutLolPerksV1PagesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lol perks v1 pages by Id params
func (o *PutLolPerksV1PagesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put lol perks v1 pages by Id params
func (o *PutLolPerksV1PagesByIDParams) WithID(id int32) *PutLolPerksV1PagesByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put lol perks v1 pages by Id params
func (o *PutLolPerksV1PagesByIDParams) SetID(id int32) {
	o.ID = id
}

// WithPage adds the page to the put lol perks v1 pages by Id params
func (o *PutLolPerksV1PagesByIDParams) WithPage(page *models.LolPerksPerkPageResource) *PutLolPerksV1PagesByIDParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the put lol perks v1 pages by Id params
func (o *PutLolPerksV1PagesByIDParams) SetPage(page *models.LolPerksPerkPageResource) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *PutLolPerksV1PagesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Page != nil {
		if err := r.SetBodyParam(o.Page); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
