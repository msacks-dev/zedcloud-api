// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package resource_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetResourceGroupEventsByNameReader is a Reader for the GetResourceGroupEventsByName structure.
type GetResourceGroupEventsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceGroupEventsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceGroupEventsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceGroupEventsByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResourceGroupEventsByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourceGroupEventsByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResourceGroupEventsByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetResourceGroupEventsByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourceGroupEventsByNameOK creates a GetResourceGroupEventsByNameOK with default headers values
func NewGetResourceGroupEventsByNameOK() *GetResourceGroupEventsByNameOK {
	return &GetResourceGroupEventsByNameOK{}
}

/* GetResourceGroupEventsByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetResourceGroupEventsByNameOK struct {
	Payload *swagger_models.EventQueryResponse
}

func (o *GetResourceGroupEventsByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{objname}/events][%d] getResourceGroupEventsByNameOK  %+v", 200, o.Payload)
}
func (o *GetResourceGroupEventsByNameOK) GetPayload() *swagger_models.EventQueryResponse {
	return o.Payload
}

func (o *GetResourceGroupEventsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.EventQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupEventsByNameUnauthorized creates a GetResourceGroupEventsByNameUnauthorized with default headers values
func NewGetResourceGroupEventsByNameUnauthorized() *GetResourceGroupEventsByNameUnauthorized {
	return &GetResourceGroupEventsByNameUnauthorized{}
}

/* GetResourceGroupEventsByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetResourceGroupEventsByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupEventsByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{objname}/events][%d] getResourceGroupEventsByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *GetResourceGroupEventsByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupEventsByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupEventsByNameForbidden creates a GetResourceGroupEventsByNameForbidden with default headers values
func NewGetResourceGroupEventsByNameForbidden() *GetResourceGroupEventsByNameForbidden {
	return &GetResourceGroupEventsByNameForbidden{}
}

/* GetResourceGroupEventsByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetResourceGroupEventsByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupEventsByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{objname}/events][%d] getResourceGroupEventsByNameForbidden  %+v", 403, o.Payload)
}
func (o *GetResourceGroupEventsByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupEventsByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupEventsByNameNotFound creates a GetResourceGroupEventsByNameNotFound with default headers values
func NewGetResourceGroupEventsByNameNotFound() *GetResourceGroupEventsByNameNotFound {
	return &GetResourceGroupEventsByNameNotFound{}
}

/* GetResourceGroupEventsByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetResourceGroupEventsByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupEventsByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{objname}/events][%d] getResourceGroupEventsByNameNotFound  %+v", 404, o.Payload)
}
func (o *GetResourceGroupEventsByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupEventsByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupEventsByNameInternalServerError creates a GetResourceGroupEventsByNameInternalServerError with default headers values
func NewGetResourceGroupEventsByNameInternalServerError() *GetResourceGroupEventsByNameInternalServerError {
	return &GetResourceGroupEventsByNameInternalServerError{}
}

/* GetResourceGroupEventsByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetResourceGroupEventsByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupEventsByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{objname}/events][%d] getResourceGroupEventsByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *GetResourceGroupEventsByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupEventsByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupEventsByNameGatewayTimeout creates a GetResourceGroupEventsByNameGatewayTimeout with default headers values
func NewGetResourceGroupEventsByNameGatewayTimeout() *GetResourceGroupEventsByNameGatewayTimeout {
	return &GetResourceGroupEventsByNameGatewayTimeout{}
}

/* GetResourceGroupEventsByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetResourceGroupEventsByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupEventsByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{objname}/events][%d] getResourceGroupEventsByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetResourceGroupEventsByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupEventsByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}