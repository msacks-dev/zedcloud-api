// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_diagnostics

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

// NewGetDeviceTwinConfigParams creates a new GetDeviceTwinConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceTwinConfigParams() *GetDeviceTwinConfigParams {
	return &GetDeviceTwinConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceTwinConfigParamsWithTimeout creates a new GetDeviceTwinConfigParams object
// with the ability to set a timeout on a request.
func NewGetDeviceTwinConfigParamsWithTimeout(timeout time.Duration) *GetDeviceTwinConfigParams {
	return &GetDeviceTwinConfigParams{
		timeout: timeout,
	}
}

// NewGetDeviceTwinConfigParamsWithContext creates a new GetDeviceTwinConfigParams object
// with the ability to set a context for a request.
func NewGetDeviceTwinConfigParamsWithContext(ctx context.Context) *GetDeviceTwinConfigParams {
	return &GetDeviceTwinConfigParams{
		Context: ctx,
	}
}

// NewGetDeviceTwinConfigParamsWithHTTPClient creates a new GetDeviceTwinConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceTwinConfigParamsWithHTTPClient(client *http.Client) *GetDeviceTwinConfigParams {
	return &GetDeviceTwinConfigParams{
		HTTPClient: client,
	}
}

/* GetDeviceTwinConfigParams contains all the parameters to send to the API endpoint
   for the get device twin config operation.

   Typically these are written to a http.Request.
*/
type GetDeviceTwinConfigParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   system generated unique id for a device
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device twin config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceTwinConfigParams) WithDefaults() *GetDeviceTwinConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device twin config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceTwinConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device twin config params
func (o *GetDeviceTwinConfigParams) WithTimeout(timeout time.Duration) *GetDeviceTwinConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device twin config params
func (o *GetDeviceTwinConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device twin config params
func (o *GetDeviceTwinConfigParams) WithContext(ctx context.Context) *GetDeviceTwinConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device twin config params
func (o *GetDeviceTwinConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device twin config params
func (o *GetDeviceTwinConfigParams) WithHTTPClient(client *http.Client) *GetDeviceTwinConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device twin config params
func (o *GetDeviceTwinConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get device twin config params
func (o *GetDeviceTwinConfigParams) WithXRequestID(xRequestID *string) *GetDeviceTwinConfigParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get device twin config params
func (o *GetDeviceTwinConfigParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the get device twin config params
func (o *GetDeviceTwinConfigParams) WithID(id string) *GetDeviceTwinConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device twin config params
func (o *GetDeviceTwinConfigParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceTwinConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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