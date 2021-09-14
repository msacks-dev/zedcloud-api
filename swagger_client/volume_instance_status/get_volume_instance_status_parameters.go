// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package volume_instance_status

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

// NewGetVolumeInstanceStatusParams creates a new GetVolumeInstanceStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVolumeInstanceStatusParams() *GetVolumeInstanceStatusParams {
	return &GetVolumeInstanceStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVolumeInstanceStatusParamsWithTimeout creates a new GetVolumeInstanceStatusParams object
// with the ability to set a timeout on a request.
func NewGetVolumeInstanceStatusParamsWithTimeout(timeout time.Duration) *GetVolumeInstanceStatusParams {
	return &GetVolumeInstanceStatusParams{
		timeout: timeout,
	}
}

// NewGetVolumeInstanceStatusParamsWithContext creates a new GetVolumeInstanceStatusParams object
// with the ability to set a context for a request.
func NewGetVolumeInstanceStatusParamsWithContext(ctx context.Context) *GetVolumeInstanceStatusParams {
	return &GetVolumeInstanceStatusParams{
		Context: ctx,
	}
}

// NewGetVolumeInstanceStatusParamsWithHTTPClient creates a new GetVolumeInstanceStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVolumeInstanceStatusParamsWithHTTPClient(client *http.Client) *GetVolumeInstanceStatusParams {
	return &GetVolumeInstanceStatusParams{
		HTTPClient: client,
	}
}

/* GetVolumeInstanceStatusParams contains all the parameters to send to the API endpoint
   for the get volume instance status operation.

   Typically these are written to a http.Request.
*/
type GetVolumeInstanceStatusParams struct {

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

// WithDefaults hydrates default values in the get volume instance status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVolumeInstanceStatusParams) WithDefaults() *GetVolumeInstanceStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get volume instance status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVolumeInstanceStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get volume instance status params
func (o *GetVolumeInstanceStatusParams) WithTimeout(timeout time.Duration) *GetVolumeInstanceStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get volume instance status params
func (o *GetVolumeInstanceStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get volume instance status params
func (o *GetVolumeInstanceStatusParams) WithContext(ctx context.Context) *GetVolumeInstanceStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get volume instance status params
func (o *GetVolumeInstanceStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get volume instance status params
func (o *GetVolumeInstanceStatusParams) WithHTTPClient(client *http.Client) *GetVolumeInstanceStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get volume instance status params
func (o *GetVolumeInstanceStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get volume instance status params
func (o *GetVolumeInstanceStatusParams) WithXRequestID(xRequestID *string) *GetVolumeInstanceStatusParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get volume instance status params
func (o *GetVolumeInstanceStatusParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the get volume instance status params
func (o *GetVolumeInstanceStatusParams) WithID(id string) *GetVolumeInstanceStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get volume instance status params
func (o *GetVolumeInstanceStatusParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVolumeInstanceStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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