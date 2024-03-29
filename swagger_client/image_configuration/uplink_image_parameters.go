// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package image_configuration

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

	"github.com/zededa/zedcloud-api/swagger_models"
)

// NewUplinkImageParams creates a new UplinkImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUplinkImageParams() *UplinkImageParams {
	return &UplinkImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUplinkImageParamsWithTimeout creates a new UplinkImageParams object
// with the ability to set a timeout on a request.
func NewUplinkImageParamsWithTimeout(timeout time.Duration) *UplinkImageParams {
	return &UplinkImageParams{
		timeout: timeout,
	}
}

// NewUplinkImageParamsWithContext creates a new UplinkImageParams object
// with the ability to set a context for a request.
func NewUplinkImageParamsWithContext(ctx context.Context) *UplinkImageParams {
	return &UplinkImageParams{
		Context: ctx,
	}
}

// NewUplinkImageParamsWithHTTPClient creates a new UplinkImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewUplinkImageParamsWithHTTPClient(client *http.Client) *UplinkImageParams {
	return &UplinkImageParams{
		HTTPClient: client,
	}
}

/* UplinkImageParams contains all the parameters to send to the API endpoint
   for the uplink image operation.

   Typically these are written to a http.Request.
*/
type UplinkImageParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.ImageConfig

	/* Name.

	   User defined name of the image, unique across the enterprise. Once image is created, name can’t be changed.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the uplink image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UplinkImageParams) WithDefaults() *UplinkImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the uplink image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UplinkImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the uplink image params
func (o *UplinkImageParams) WithTimeout(timeout time.Duration) *UplinkImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the uplink image params
func (o *UplinkImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the uplink image params
func (o *UplinkImageParams) WithContext(ctx context.Context) *UplinkImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the uplink image params
func (o *UplinkImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the uplink image params
func (o *UplinkImageParams) WithHTTPClient(client *http.Client) *UplinkImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the uplink image params
func (o *UplinkImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the uplink image params
func (o *UplinkImageParams) WithXRequestID(xRequestID *string) *UplinkImageParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the uplink image params
func (o *UplinkImageParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the uplink image params
func (o *UplinkImageParams) WithBody(body *swagger_models.ImageConfig) *UplinkImageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the uplink image params
func (o *UplinkImageParams) SetBody(body *swagger_models.ImageConfig) {
	o.Body = body
}

// WithName adds the name to the uplink image params
func (o *UplinkImageParams) WithName(name string) *UplinkImageParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the uplink image params
func (o *UplinkImageParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UplinkImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
