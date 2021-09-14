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
)

// NewGetLatestImageVersionParams creates a new GetLatestImageVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLatestImageVersionParams() *GetLatestImageVersionParams {
	return &GetLatestImageVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLatestImageVersionParamsWithTimeout creates a new GetLatestImageVersionParams object
// with the ability to set a timeout on a request.
func NewGetLatestImageVersionParamsWithTimeout(timeout time.Duration) *GetLatestImageVersionParams {
	return &GetLatestImageVersionParams{
		timeout: timeout,
	}
}

// NewGetLatestImageVersionParamsWithContext creates a new GetLatestImageVersionParams object
// with the ability to set a context for a request.
func NewGetLatestImageVersionParamsWithContext(ctx context.Context) *GetLatestImageVersionParams {
	return &GetLatestImageVersionParams{
		Context: ctx,
	}
}

// NewGetLatestImageVersionParamsWithHTTPClient creates a new GetLatestImageVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLatestImageVersionParamsWithHTTPClient(client *http.Client) *GetLatestImageVersionParams {
	return &GetLatestImageVersionParams{
		HTTPClient: client,
	}
}

/* GetLatestImageVersionParams contains all the parameters to send to the API endpoint
   for the get latest image version operation.

   Typically these are written to a http.Request.
*/
type GetLatestImageVersionParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// EnterpriseID.
	EnterpriseID *string

	/* ImageArch.

	   Image architecture to be matched : 'AMD64' or 'ARM64'.
	*/
	ImageArch string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get latest image version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLatestImageVersionParams) WithDefaults() *GetLatestImageVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get latest image version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLatestImageVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get latest image version params
func (o *GetLatestImageVersionParams) WithTimeout(timeout time.Duration) *GetLatestImageVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get latest image version params
func (o *GetLatestImageVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get latest image version params
func (o *GetLatestImageVersionParams) WithContext(ctx context.Context) *GetLatestImageVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get latest image version params
func (o *GetLatestImageVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get latest image version params
func (o *GetLatestImageVersionParams) WithHTTPClient(client *http.Client) *GetLatestImageVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get latest image version params
func (o *GetLatestImageVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get latest image version params
func (o *GetLatestImageVersionParams) WithXRequestID(xRequestID *string) *GetLatestImageVersionParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get latest image version params
func (o *GetLatestImageVersionParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEnterpriseID adds the enterpriseID to the get latest image version params
func (o *GetLatestImageVersionParams) WithEnterpriseID(enterpriseID *string) *GetLatestImageVersionParams {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the get latest image version params
func (o *GetLatestImageVersionParams) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithImageArch adds the imageArch to the get latest image version params
func (o *GetLatestImageVersionParams) WithImageArch(imageArch string) *GetLatestImageVersionParams {
	o.SetImageArch(imageArch)
	return o
}

// SetImageArch adds the imageArch to the get latest image version params
func (o *GetLatestImageVersionParams) SetImageArch(imageArch string) {
	o.ImageArch = imageArch
}

// WriteToRequest writes these params to a swagger request
func (o *GetLatestImageVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param imageArch
	if err := r.SetPathParam("imageArch", o.ImageArch); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}