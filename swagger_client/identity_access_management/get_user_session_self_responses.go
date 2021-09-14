// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetUserSessionSelfReader is a Reader for the GetUserSessionSelf structure.
type GetUserSessionSelfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserSessionSelfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserSessionSelfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserSessionSelfUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserSessionSelfForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserSessionSelfInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserSessionSelfGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserSessionSelfOK creates a GetUserSessionSelfOK with default headers values
func NewGetUserSessionSelfOK() *GetUserSessionSelfOK {
	return &GetUserSessionSelfOK{}
}

/* GetUserSessionSelfOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetUserSessionSelfOK struct {
	Payload *swagger_models.AAAFrontendSessionDetailsResponse
}

func (o *GetUserSessionSelfOK) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] getUserSessionSelfOK  %+v", 200, o.Payload)
}
func (o *GetUserSessionSelfOK) GetPayload() *swagger_models.AAAFrontendSessionDetailsResponse {
	return o.Payload
}

func (o *GetUserSessionSelfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AAAFrontendSessionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSessionSelfUnauthorized creates a GetUserSessionSelfUnauthorized with default headers values
func NewGetUserSessionSelfUnauthorized() *GetUserSessionSelfUnauthorized {
	return &GetUserSessionSelfUnauthorized{}
}

/* GetUserSessionSelfUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetUserSessionSelfUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetUserSessionSelfUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] getUserSessionSelfUnauthorized  %+v", 401, o.Payload)
}
func (o *GetUserSessionSelfUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetUserSessionSelfUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSessionSelfForbidden creates a GetUserSessionSelfForbidden with default headers values
func NewGetUserSessionSelfForbidden() *GetUserSessionSelfForbidden {
	return &GetUserSessionSelfForbidden{}
}

/* GetUserSessionSelfForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetUserSessionSelfForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetUserSessionSelfForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] getUserSessionSelfForbidden  %+v", 403, o.Payload)
}
func (o *GetUserSessionSelfForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetUserSessionSelfForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSessionSelfInternalServerError creates a GetUserSessionSelfInternalServerError with default headers values
func NewGetUserSessionSelfInternalServerError() *GetUserSessionSelfInternalServerError {
	return &GetUserSessionSelfInternalServerError{}
}

/* GetUserSessionSelfInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetUserSessionSelfInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetUserSessionSelfInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] getUserSessionSelfInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUserSessionSelfInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetUserSessionSelfInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSessionSelfGatewayTimeout creates a GetUserSessionSelfGatewayTimeout with default headers values
func NewGetUserSessionSelfGatewayTimeout() *GetUserSessionSelfGatewayTimeout {
	return &GetUserSessionSelfGatewayTimeout{}
}

/* GetUserSessionSelfGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetUserSessionSelfGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetUserSessionSelfGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] getUserSessionSelfGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetUserSessionSelfGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetUserSessionSelfGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}