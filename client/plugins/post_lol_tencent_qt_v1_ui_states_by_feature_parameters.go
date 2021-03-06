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

// NewPostLolTencentQtV1UIStatesByFeatureParams creates a new PostLolTencentQtV1UIStatesByFeatureParams object
// with the default values initialized.
func NewPostLolTencentQtV1UIStatesByFeatureParams() *PostLolTencentQtV1UIStatesByFeatureParams {
	var ()
	return &PostLolTencentQtV1UIStatesByFeatureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolTencentQtV1UIStatesByFeatureParamsWithTimeout creates a new PostLolTencentQtV1UIStatesByFeatureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolTencentQtV1UIStatesByFeatureParamsWithTimeout(timeout time.Duration) *PostLolTencentQtV1UIStatesByFeatureParams {
	var ()
	return &PostLolTencentQtV1UIStatesByFeatureParams{

		timeout: timeout,
	}
}

// NewPostLolTencentQtV1UIStatesByFeatureParamsWithContext creates a new PostLolTencentQtV1UIStatesByFeatureParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolTencentQtV1UIStatesByFeatureParamsWithContext(ctx context.Context) *PostLolTencentQtV1UIStatesByFeatureParams {
	var ()
	return &PostLolTencentQtV1UIStatesByFeatureParams{

		Context: ctx,
	}
}

// NewPostLolTencentQtV1UIStatesByFeatureParamsWithHTTPClient creates a new PostLolTencentQtV1UIStatesByFeatureParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolTencentQtV1UIStatesByFeatureParamsWithHTTPClient(client *http.Client) *PostLolTencentQtV1UIStatesByFeatureParams {
	var ()
	return &PostLolTencentQtV1UIStatesByFeatureParams{
		HTTPClient: client,
	}
}

/*PostLolTencentQtV1UIStatesByFeatureParams contains all the parameters to send to the API endpoint
for the post lol tencent qt v1 Ui states by feature operation typically these are written to a http.Request
*/
type PostLolTencentQtV1UIStatesByFeatureParams struct {

	/*Feature*/
	Feature string
	/*State*/
	State *models.LolTencentQtTencentQTNotification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol tencent qt v1 Ui states by feature params
func (o *PostLolTencentQtV1UIStatesByFeatureParams) WithTimeout(timeout time.Duration) *PostLolTencentQtV1UIStatesByFeatureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol tencent qt v1 Ui states by feature params
func (o *PostLolTencentQtV1UIStatesByFeatureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol tencent qt v1 Ui states by feature params
func (o *PostLolTencentQtV1UIStatesByFeatureParams) WithContext(ctx context.Context) *PostLolTencentQtV1UIStatesByFeatureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol tencent qt v1 Ui states by feature params
func (o *PostLolTencentQtV1UIStatesByFeatureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol tencent qt v1 Ui states by feature params
func (o *PostLolTencentQtV1UIStatesByFeatureParams) WithHTTPClient(client *http.Client) *PostLolTencentQtV1UIStatesByFeatureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol tencent qt v1 Ui states by feature params
func (o *PostLolTencentQtV1UIStatesByFeatureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeature adds the feature to the post lol tencent qt v1 Ui states by feature params
func (o *PostLolTencentQtV1UIStatesByFeatureParams) WithFeature(feature string) *PostLolTencentQtV1UIStatesByFeatureParams {
	o.SetFeature(feature)
	return o
}

// SetFeature adds the feature to the post lol tencent qt v1 Ui states by feature params
func (o *PostLolTencentQtV1UIStatesByFeatureParams) SetFeature(feature string) {
	o.Feature = feature
}

// WithState adds the state to the post lol tencent qt v1 Ui states by feature params
func (o *PostLolTencentQtV1UIStatesByFeatureParams) WithState(state *models.LolTencentQtTencentQTNotification) *PostLolTencentQtV1UIStatesByFeatureParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the post lol tencent qt v1 Ui states by feature params
func (o *PostLolTencentQtV1UIStatesByFeatureParams) SetState(state *models.LolTencentQtTencentQTNotification) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolTencentQtV1UIStatesByFeatureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param feature
	if err := r.SetPathParam("feature", o.Feature); err != nil {
		return err
	}

	if o.State != nil {
		if err := r.SetBodyParam(o.State); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
