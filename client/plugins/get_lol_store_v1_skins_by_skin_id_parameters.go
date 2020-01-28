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

// NewGetLolStoreV1SkinsBySkinIDParams creates a new GetLolStoreV1SkinsBySkinIDParams object
// with the default values initialized.
func NewGetLolStoreV1SkinsBySkinIDParams() *GetLolStoreV1SkinsBySkinIDParams {
	var ()
	return &GetLolStoreV1SkinsBySkinIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolStoreV1SkinsBySkinIDParamsWithTimeout creates a new GetLolStoreV1SkinsBySkinIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolStoreV1SkinsBySkinIDParamsWithTimeout(timeout time.Duration) *GetLolStoreV1SkinsBySkinIDParams {
	var ()
	return &GetLolStoreV1SkinsBySkinIDParams{

		timeout: timeout,
	}
}

// NewGetLolStoreV1SkinsBySkinIDParamsWithContext creates a new GetLolStoreV1SkinsBySkinIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolStoreV1SkinsBySkinIDParamsWithContext(ctx context.Context) *GetLolStoreV1SkinsBySkinIDParams {
	var ()
	return &GetLolStoreV1SkinsBySkinIDParams{

		Context: ctx,
	}
}

// NewGetLolStoreV1SkinsBySkinIDParamsWithHTTPClient creates a new GetLolStoreV1SkinsBySkinIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolStoreV1SkinsBySkinIDParamsWithHTTPClient(client *http.Client) *GetLolStoreV1SkinsBySkinIDParams {
	var ()
	return &GetLolStoreV1SkinsBySkinIDParams{
		HTTPClient: client,
	}
}

/*GetLolStoreV1SkinsBySkinIDParams contains all the parameters to send to the API endpoint
for the get lol store v1 skins by skin Id operation typically these are written to a http.Request
*/
type GetLolStoreV1SkinsBySkinIDParams struct {

	/*SkinID*/
	SkinID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol store v1 skins by skin Id params
func (o *GetLolStoreV1SkinsBySkinIDParams) WithTimeout(timeout time.Duration) *GetLolStoreV1SkinsBySkinIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol store v1 skins by skin Id params
func (o *GetLolStoreV1SkinsBySkinIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol store v1 skins by skin Id params
func (o *GetLolStoreV1SkinsBySkinIDParams) WithContext(ctx context.Context) *GetLolStoreV1SkinsBySkinIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol store v1 skins by skin Id params
func (o *GetLolStoreV1SkinsBySkinIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol store v1 skins by skin Id params
func (o *GetLolStoreV1SkinsBySkinIDParams) WithHTTPClient(client *http.Client) *GetLolStoreV1SkinsBySkinIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol store v1 skins by skin Id params
func (o *GetLolStoreV1SkinsBySkinIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkinID adds the skinID to the get lol store v1 skins by skin Id params
func (o *GetLolStoreV1SkinsBySkinIDParams) WithSkinID(skinID int32) *GetLolStoreV1SkinsBySkinIDParams {
	o.SetSkinID(skinID)
	return o
}

// SetSkinID adds the skinId to the get lol store v1 skins by skin Id params
func (o *GetLolStoreV1SkinsBySkinIDParams) SetSkinID(skinID int32) {
	o.SkinID = skinID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolStoreV1SkinsBySkinIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param skinId
	if err := r.SetPathParam("skinId", swag.FormatInt32(o.SkinID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
