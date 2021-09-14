// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetEdgeNodeReader is a Reader for the GetEdgeNode structure.
type GetEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEdgeNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEdgeNodeOK creates a GetEdgeNodeOK with default headers values
func NewGetEdgeNodeOK() *GetEdgeNodeOK {
	return &GetEdgeNodeOK{}
}

/* GetEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetEdgeNodeOK struct {
	Payload *swagger_models.DeviceConfig
}

func (o *GetEdgeNodeOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] getEdgeNodeOK  %+v", 200, o.Payload)
}
func (o *GetEdgeNodeOK) GetPayload() *swagger_models.DeviceConfig {
	return o.Payload
}

func (o *GetEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.DeviceConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeUnauthorized creates a GetEdgeNodeUnauthorized with default headers values
func NewGetEdgeNodeUnauthorized() *GetEdgeNodeUnauthorized {
	return &GetEdgeNodeUnauthorized{}
}

/* GetEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetEdgeNodeUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] getEdgeNodeUnauthorized  %+v", 401, o.Payload)
}
func (o *GetEdgeNodeUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeForbidden creates a GetEdgeNodeForbidden with default headers values
func NewGetEdgeNodeForbidden() *GetEdgeNodeForbidden {
	return &GetEdgeNodeForbidden{}
}

/* GetEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetEdgeNodeForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] getEdgeNodeForbidden  %+v", 403, o.Payload)
}
func (o *GetEdgeNodeForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeNotFound creates a GetEdgeNodeNotFound with default headers values
func NewGetEdgeNodeNotFound() *GetEdgeNodeNotFound {
	return &GetEdgeNodeNotFound{}
}

/* GetEdgeNodeNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetEdgeNodeNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] getEdgeNodeNotFound  %+v", 404, o.Payload)
}
func (o *GetEdgeNodeNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeInternalServerError creates a GetEdgeNodeInternalServerError with default headers values
func NewGetEdgeNodeInternalServerError() *GetEdgeNodeInternalServerError {
	return &GetEdgeNodeInternalServerError{}
}

/* GetEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetEdgeNodeInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] getEdgeNodeInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEdgeNodeInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeGatewayTimeout creates a GetEdgeNodeGatewayTimeout with default headers values
func NewGetEdgeNodeGatewayTimeout() *GetEdgeNodeGatewayTimeout {
	return &GetEdgeNodeGatewayTimeout{}
}

/* GetEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetEdgeNodeGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] getEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetEdgeNodeGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}