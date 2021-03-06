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

// NewPostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams creates a new PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams object
// with the default values initialized.
func NewPostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams() *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams {
	var ()
	return &PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParamsWithTimeout creates a new PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParamsWithTimeout(timeout time.Duration) *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams {
	var ()
	return &PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams{

		timeout: timeout,
	}
}

// NewPostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParamsWithContext creates a new PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParamsWithContext(ctx context.Context) *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams {
	var ()
	return &PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams{

		Context: ctx,
	}
}

// NewPostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParamsWithHTTPClient creates a new PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParamsWithHTTPClient(client *http.Client) *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams {
	var ()
	return &PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams{
		HTTPClient: client,
	}
}

/*PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams contains all the parameters to send to the API endpoint
for the post lol purchase widget v1 purchasable items by inventory type operation typically these are written to a http.Request
*/
type PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams struct {

	/*InventoryType*/
	InventoryType string
	/*ItemIds*/
	ItemIds []int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol purchase widget v1 purchasable items by inventory type params
func (o *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams) WithTimeout(timeout time.Duration) *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol purchase widget v1 purchasable items by inventory type params
func (o *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol purchase widget v1 purchasable items by inventory type params
func (o *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams) WithContext(ctx context.Context) *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol purchase widget v1 purchasable items by inventory type params
func (o *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol purchase widget v1 purchasable items by inventory type params
func (o *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams) WithHTTPClient(client *http.Client) *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol purchase widget v1 purchasable items by inventory type params
func (o *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInventoryType adds the inventoryType to the post lol purchase widget v1 purchasable items by inventory type params
func (o *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams) WithInventoryType(inventoryType string) *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams {
	o.SetInventoryType(inventoryType)
	return o
}

// SetInventoryType adds the inventoryType to the post lol purchase widget v1 purchasable items by inventory type params
func (o *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams) SetInventoryType(inventoryType string) {
	o.InventoryType = inventoryType
}

// WithItemIds adds the itemIds to the post lol purchase widget v1 purchasable items by inventory type params
func (o *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams) WithItemIds(itemIds []int64) *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams {
	o.SetItemIds(itemIds)
	return o
}

// SetItemIds adds the itemIds to the post lol purchase widget v1 purchasable items by inventory type params
func (o *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams) SetItemIds(itemIds []int64) {
	o.ItemIds = itemIds
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolPurchaseWidgetV1PurchasableItemsByInventoryTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param inventoryType
	if err := r.SetPathParam("inventoryType", o.InventoryType); err != nil {
		return err
	}

	if o.ItemIds != nil {
		if err := r.SetBodyParam(o.ItemIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
