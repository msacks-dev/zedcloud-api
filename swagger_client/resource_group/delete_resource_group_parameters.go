// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package resource_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteResourceGroupParams creates a new DeleteResourceGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteResourceGroupParams() *DeleteResourceGroupParams {
	return &DeleteResourceGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteResourceGroupParamsWithTimeout creates a new DeleteResourceGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteResourceGroupParamsWithTimeout(timeout time.Duration) *DeleteResourceGroupParams {
	return &DeleteResourceGroupParams{
		timeout: timeout,
	}
}

// NewDeleteResourceGroupParamsWithContext creates a new DeleteResourceGroupParams object
// with the ability to set a context for a request.
func NewDeleteResourceGroupParamsWithContext(ctx context.Context) *DeleteResourceGroupParams {
	return &DeleteResourceGroupParams{
		Context: ctx,
	}
}

// NewDeleteResourceGroupParamsWithHTTPClient creates a new DeleteResourceGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteResourceGroupParamsWithHTTPClient(client *http.Client) *DeleteResourceGroupParams {
	return &DeleteResourceGroupParams{
		HTTPClient: client,
	}
}

/* DeleteResourceGroupParams contains all the parameters to send to the API endpoint
   for the delete resource group operation.

   Typically these are written to a http.Request.
*/
type DeleteResourceGroupParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the resource group
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete resource group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteResourceGroupParams) WithDefaults() *DeleteResourceGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete resource group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteResourceGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete resource group params
func (o *DeleteResourceGroupParams) WithTimeout(timeout time.Duration) *DeleteResourceGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete resource group params
func (o *DeleteResourceGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete resource group params
func (o *DeleteResourceGroupParams) WithContext(ctx context.Context) *DeleteResourceGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete resource group params
func (o *DeleteResourceGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete resource group params
func (o *DeleteResourceGroupParams) WithHTTPClient(client *http.Client) *DeleteResourceGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete resource group params
func (o *DeleteResourceGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the delete resource group params
func (o *DeleteResourceGroupParams) WithXRequestID(xRequestID *string) *DeleteResourceGroupParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete resource group params
func (o *DeleteResourceGroupParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the delete resource group params
func (o *DeleteResourceGroupParams) WithID(id string) *DeleteResourceGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete resource group params
func (o *DeleteResourceGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteResourceGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}