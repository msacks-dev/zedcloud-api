// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

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

// NewGetHardwareBrandByNameParams creates a new GetHardwareBrandByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHardwareBrandByNameParams() *GetHardwareBrandByNameParams {
	return &GetHardwareBrandByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHardwareBrandByNameParamsWithTimeout creates a new GetHardwareBrandByNameParams object
// with the ability to set a timeout on a request.
func NewGetHardwareBrandByNameParamsWithTimeout(timeout time.Duration) *GetHardwareBrandByNameParams {
	return &GetHardwareBrandByNameParams{
		timeout: timeout,
	}
}

// NewGetHardwareBrandByNameParamsWithContext creates a new GetHardwareBrandByNameParams object
// with the ability to set a context for a request.
func NewGetHardwareBrandByNameParamsWithContext(ctx context.Context) *GetHardwareBrandByNameParams {
	return &GetHardwareBrandByNameParams{
		Context: ctx,
	}
}

// NewGetHardwareBrandByNameParamsWithHTTPClient creates a new GetHardwareBrandByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHardwareBrandByNameParamsWithHTTPClient(client *http.Client) *GetHardwareBrandByNameParams {
	return &GetHardwareBrandByNameParams{
		HTTPClient: client,
	}
}

/* GetHardwareBrandByNameParams contains all the parameters to send to the API endpoint
   for the get hardware brand by name operation.

   Typically these are written to a http.Request.
*/
type GetHardwareBrandByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* EnterpriseID.

	   deprecated field: EnterpriseId
	*/
	EnterpriseID *string

	/* Name.

	   user defined sys brand name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get hardware brand by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHardwareBrandByNameParams) WithDefaults() *GetHardwareBrandByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get hardware brand by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHardwareBrandByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get hardware brand by name params
func (o *GetHardwareBrandByNameParams) WithTimeout(timeout time.Duration) *GetHardwareBrandByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hardware brand by name params
func (o *GetHardwareBrandByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hardware brand by name params
func (o *GetHardwareBrandByNameParams) WithContext(ctx context.Context) *GetHardwareBrandByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hardware brand by name params
func (o *GetHardwareBrandByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hardware brand by name params
func (o *GetHardwareBrandByNameParams) WithHTTPClient(client *http.Client) *GetHardwareBrandByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hardware brand by name params
func (o *GetHardwareBrandByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get hardware brand by name params
func (o *GetHardwareBrandByNameParams) WithXRequestID(xRequestID *string) *GetHardwareBrandByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get hardware brand by name params
func (o *GetHardwareBrandByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEnterpriseID adds the enterpriseID to the get hardware brand by name params
func (o *GetHardwareBrandByNameParams) WithEnterpriseID(enterpriseID *string) *GetHardwareBrandByNameParams {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the get hardware brand by name params
func (o *GetHardwareBrandByNameParams) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithName adds the name to the get hardware brand by name params
func (o *GetHardwareBrandByNameParams) WithName(name string) *GetHardwareBrandByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get hardware brand by name params
func (o *GetHardwareBrandByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetHardwareBrandByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EnterpriseID != nil {

		// query param enterpriseId
		var qrEnterpriseID string

		if o.EnterpriseID != nil {
			qrEnterpriseID = *o.EnterpriseID
		}
		qEnterpriseID := qrEnterpriseID
		if qEnterpriseID != "" {

			if err := r.SetQueryParam("enterpriseId", qEnterpriseID); err != nil {
				return err
			}
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