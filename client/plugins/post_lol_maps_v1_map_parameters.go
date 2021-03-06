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

// NewPostLolMapsV1MapParams creates a new PostLolMapsV1MapParams object
// with the default values initialized.
func NewPostLolMapsV1MapParams() *PostLolMapsV1MapParams {
	var ()
	return &PostLolMapsV1MapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolMapsV1MapParamsWithTimeout creates a new PostLolMapsV1MapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolMapsV1MapParamsWithTimeout(timeout time.Duration) *PostLolMapsV1MapParams {
	var ()
	return &PostLolMapsV1MapParams{

		timeout: timeout,
	}
}

// NewPostLolMapsV1MapParamsWithContext creates a new PostLolMapsV1MapParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolMapsV1MapParamsWithContext(ctx context.Context) *PostLolMapsV1MapParams {
	var ()
	return &PostLolMapsV1MapParams{

		Context: ctx,
	}
}

// NewPostLolMapsV1MapParamsWithHTTPClient creates a new PostLolMapsV1MapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolMapsV1MapParamsWithHTTPClient(client *http.Client) *PostLolMapsV1MapParams {
	var ()
	return &PostLolMapsV1MapParams{
		HTTPClient: client,
	}
}

/*PostLolMapsV1MapParams contains all the parameters to send to the API endpoint
for the post lol maps v1 map operation typically these are written to a http.Request
*/
type PostLolMapsV1MapParams struct {

	/*Map*/
	Map *models.LolMapsMaps

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol maps v1 map params
func (o *PostLolMapsV1MapParams) WithTimeout(timeout time.Duration) *PostLolMapsV1MapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol maps v1 map params
func (o *PostLolMapsV1MapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol maps v1 map params
func (o *PostLolMapsV1MapParams) WithContext(ctx context.Context) *PostLolMapsV1MapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol maps v1 map params
func (o *PostLolMapsV1MapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol maps v1 map params
func (o *PostLolMapsV1MapParams) WithHTTPClient(client *http.Client) *PostLolMapsV1MapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol maps v1 map params
func (o *PostLolMapsV1MapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMap adds the mapVar to the post lol maps v1 map params
func (o *PostLolMapsV1MapParams) WithMap(mapVar *models.LolMapsMaps) *PostLolMapsV1MapParams {
	o.SetMap(mapVar)
	return o
}

// SetMap adds the map to the post lol maps v1 map params
func (o *PostLolMapsV1MapParams) SetMap(mapVar *models.LolMapsMaps) {
	o.Map = mapVar
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolMapsV1MapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Map != nil {
		if err := r.SetBodyParam(o.Map); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
