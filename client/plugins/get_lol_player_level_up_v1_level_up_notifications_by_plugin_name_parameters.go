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

// NewGetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams creates a new GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams object
// with the default values initialized.
func NewGetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams() *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	var ()
	return &GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParamsWithTimeout creates a new GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParamsWithTimeout(timeout time.Duration) *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	var ()
	return &GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams{

		timeout: timeout,
	}
}

// NewGetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParamsWithContext creates a new GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParamsWithContext(ctx context.Context) *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	var ()
	return &GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams{

		Context: ctx,
	}
}

// NewGetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParamsWithHTTPClient creates a new GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParamsWithHTTPClient(client *http.Client) *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	var ()
	return &GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams{
		HTTPClient: client,
	}
}

/*GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams contains all the parameters to send to the API endpoint
for the get lol player level up v1 level up notifications by plugin name operation typically these are written to a http.Request
*/
type GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams struct {

	/*PluginName*/
	PluginName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol player level up v1 level up notifications by plugin name params
func (o *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) WithTimeout(timeout time.Duration) *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol player level up v1 level up notifications by plugin name params
func (o *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol player level up v1 level up notifications by plugin name params
func (o *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) WithContext(ctx context.Context) *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol player level up v1 level up notifications by plugin name params
func (o *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol player level up v1 level up notifications by plugin name params
func (o *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) WithHTTPClient(client *http.Client) *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol player level up v1 level up notifications by plugin name params
func (o *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPluginName adds the pluginName to the get lol player level up v1 level up notifications by plugin name params
func (o *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) WithPluginName(pluginName string) *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	o.SetPluginName(pluginName)
	return o
}

// SetPluginName adds the pluginName to the get lol player level up v1 level up notifications by plugin name params
func (o *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) SetPluginName(pluginName string) {
	o.PluginName = pluginName
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pluginName
	if err := r.SetPathParam("pluginName", o.PluginName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
