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

// NewGetEventsTimelineParams creates a new GetEventsTimelineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEventsTimelineParams() *GetEventsTimelineParams {
	return &GetEventsTimelineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventsTimelineParamsWithTimeout creates a new GetEventsTimelineParams object
// with the ability to set a timeout on a request.
func NewGetEventsTimelineParamsWithTimeout(timeout time.Duration) *GetEventsTimelineParams {
	return &GetEventsTimelineParams{
		timeout: timeout,
	}
}

// NewGetEventsTimelineParamsWithContext creates a new GetEventsTimelineParams object
// with the ability to set a context for a request.
func NewGetEventsTimelineParamsWithContext(ctx context.Context) *GetEventsTimelineParams {
	return &GetEventsTimelineParams{
		Context: ctx,
	}
}

// NewGetEventsTimelineParamsWithHTTPClient creates a new GetEventsTimelineParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEventsTimelineParamsWithHTTPClient(client *http.Client) *GetEventsTimelineParams {
	return &GetEventsTimelineParams{
		HTTPClient: client,
	}
}

/* GetEventsTimelineParams contains all the parameters to send to the API endpoint
   for the get events timeline operation.

   Typically these are written to a http.Request.
*/
type GetEventsTimelineParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* EnterpriseID.

	   system generated unique id for an enterprise (deprecated)
	*/
	EnterpriseID *string

	/* Objid.

	   Object id
	*/
	Objid *string

	/* Objname.

	   Object name
	*/
	Objname *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get events timeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEventsTimelineParams) WithDefaults() *GetEventsTimelineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get events timeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEventsTimelineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get events timeline params
func (o *GetEventsTimelineParams) WithTimeout(timeout time.Duration) *GetEventsTimelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get events timeline params
func (o *GetEventsTimelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get events timeline params
func (o *GetEventsTimelineParams) WithContext(ctx context.Context) *GetEventsTimelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get events timeline params
func (o *GetEventsTimelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get events timeline params
func (o *GetEventsTimelineParams) WithHTTPClient(client *http.Client) *GetEventsTimelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get events timeline params
func (o *GetEventsTimelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get events timeline params
func (o *GetEventsTimelineParams) WithXRequestID(xRequestID *string) *GetEventsTimelineParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get events timeline params
func (o *GetEventsTimelineParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEnterpriseID adds the enterpriseID to the get events timeline params
func (o *GetEventsTimelineParams) WithEnterpriseID(enterpriseID *string) *GetEventsTimelineParams {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the get events timeline params
func (o *GetEventsTimelineParams) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithObjid adds the objid to the get events timeline params
func (o *GetEventsTimelineParams) WithObjid(objid *string) *GetEventsTimelineParams {
	o.SetObjid(objid)
	return o
}

// SetObjid adds the objid to the get events timeline params
func (o *GetEventsTimelineParams) SetObjid(objid *string) {
	o.Objid = objid
}

// WithObjname adds the objname to the get events timeline params
func (o *GetEventsTimelineParams) WithObjname(objname *string) *GetEventsTimelineParams {
	o.SetObjname(objname)
	return o
}

// SetObjname adds the objname to the get events timeline params
func (o *GetEventsTimelineParams) SetObjname(objname *string) {
	o.Objname = objname
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventsTimelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Objid != nil {

		// query param objid
		var qrObjid string

		if o.Objid != nil {
			qrObjid = *o.Objid
		}
		qObjid := qrObjid
		if qObjid != "" {

			if err := r.SetQueryParam("objid", qObjid); err != nil {
				return err
			}
		}
	}

	if o.Objname != nil {

		// query param objname
		var qrObjname string

		if o.Objname != nil {
			qrObjname = *o.Objname
		}
		qObjname := qrObjname
		if qObjname != "" {

			if err := r.SetQueryParam("objname", qObjname); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}