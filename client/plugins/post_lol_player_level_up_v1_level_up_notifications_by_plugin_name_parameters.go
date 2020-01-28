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

// NewPostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams creates a new PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams object
// with the default values initialized.
func NewPostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams() *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	var ()
	return &PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParamsWithTimeout creates a new PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParamsWithTimeout(timeout time.Duration) *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	var ()
	return &PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams{

		timeout: timeout,
	}
}

// NewPostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParamsWithContext creates a new PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParamsWithContext(ctx context.Context) *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	var ()
	return &PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams{

		Context: ctx,
	}
}

// NewPostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParamsWithHTTPClient creates a new PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParamsWithHTTPClient(client *http.Client) *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	var ()
	return &PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams{
		HTTPClient: client,
	}
}

/*PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams contains all the parameters to send to the API endpoint
for the post lol player level up v1 level up notifications by plugin name operation typically these are written to a http.Request
*/
type PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams struct {

	/*LevelUpEventAck*/
	LevelUpEventAck *models.LolPlayerLevelUpPlayerLevelUpEventAck
	/*PluginName*/
	PluginName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol player level up v1 level up notifications by plugin name params
func (o *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) WithTimeout(timeout time.Duration) *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol player level up v1 level up notifications by plugin name params
func (o *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol player level up v1 level up notifications by plugin name params
func (o *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) WithContext(ctx context.Context) *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol player level up v1 level up notifications by plugin name params
func (o *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol player level up v1 level up notifications by plugin name params
func (o *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) WithHTTPClient(client *http.Client) *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol player level up v1 level up notifications by plugin name params
func (o *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLevelUpEventAck adds the levelUpEventAck to the post lol player level up v1 level up notifications by plugin name params
func (o *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) WithLevelUpEventAck(levelUpEventAck *models.LolPlayerLevelUpPlayerLevelUpEventAck) *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	o.SetLevelUpEventAck(levelUpEventAck)
	return o
}

// SetLevelUpEventAck adds the levelUpEventAck to the post lol player level up v1 level up notifications by plugin name params
func (o *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) SetLevelUpEventAck(levelUpEventAck *models.LolPlayerLevelUpPlayerLevelUpEventAck) {
	o.LevelUpEventAck = levelUpEventAck
}

// WithPluginName adds the pluginName to the post lol player level up v1 level up notifications by plugin name params
func (o *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) WithPluginName(pluginName string) *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams {
	o.SetPluginName(pluginName)
	return o
}

// SetPluginName adds the pluginName to the post lol player level up v1 level up notifications by plugin name params
func (o *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) SetPluginName(pluginName string) {
	o.PluginName = pluginName
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolPlayerLevelUpV1LevelUpNotificationsByPluginNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LevelUpEventAck != nil {
		if err := r.SetBodyParam(o.LevelUpEventAck); err != nil {
			return err
		}
	}

	// path param pluginName
	if err := r.SetPathParam("pluginName", o.PluginName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
