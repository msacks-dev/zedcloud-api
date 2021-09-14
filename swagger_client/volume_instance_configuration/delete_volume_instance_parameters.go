// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package volume_instance_configuration

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

// NewDeleteVolumeInstanceParams creates a new DeleteVolumeInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVolumeInstanceParams() *DeleteVolumeInstanceParams {
	return &DeleteVolumeInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVolumeInstanceParamsWithTimeout creates a new DeleteVolumeInstanceParams object
// with the ability to set a timeout on a request.
func NewDeleteVolumeInstanceParamsWithTimeout(timeout time.Duration) *DeleteVolumeInstanceParams {
	return &DeleteVolumeInstanceParams{
		timeout: timeout,
	}
}

// NewDeleteVolumeInstanceParamsWithContext creates a new DeleteVolumeInstanceParams object
// with the ability to set a context for a request.
func NewDeleteVolumeInstanceParamsWithContext(ctx context.Context) *DeleteVolumeInstanceParams {
	return &DeleteVolumeInstanceParams{
		Context: ctx,
	}
}

// NewDeleteVolumeInstanceParamsWithHTTPClient creates a new DeleteVolumeInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVolumeInstanceParamsWithHTTPClient(client *http.Client) *DeleteVolumeInstanceParams {
	return &DeleteVolumeInstanceParams{
		HTTPClient: client,
	}
}

/* DeleteVolumeInstanceParams contains all the parameters to send to the API endpoint
   for the delete volume instance operation.

   Typically these are written to a http.Request.
*/
type DeleteVolumeInstanceParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete volume instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVolumeInstanceParams) WithDefaults() *DeleteVolumeInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete volume instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVolumeInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete volume instance params
func (o *DeleteVolumeInstanceParams) WithTimeout(timeout time.Duration) *DeleteVolumeInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete volume instance params
func (o *DeleteVolumeInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete volume instance params
func (o *DeleteVolumeInstanceParams) WithContext(ctx context.Context) *DeleteVolumeInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete volume instance params
func (o *DeleteVolumeInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete volume instance params
func (o *DeleteVolumeInstanceParams) WithHTTPClient(client *http.Client) *DeleteVolumeInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete volume instance params
func (o *DeleteVolumeInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the delete volume instance params
func (o *DeleteVolumeInstanceParams) WithXRequestID(xRequestID *string) *DeleteVolumeInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete volume instance params
func (o *DeleteVolumeInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the delete volume instance params
func (o *DeleteVolumeInstanceParams) WithID(id string) *DeleteVolumeInstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete volume instance params
func (o *DeleteVolumeInstanceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVolumeInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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