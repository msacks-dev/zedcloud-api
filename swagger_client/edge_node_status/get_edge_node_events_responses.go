// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_node_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetEdgeNodeEventsReader is a Reader for the GetEdgeNodeEvents structure.
type GetEdgeNodeEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEdgeNodeEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEdgeNodeEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEdgeNodeEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEdgeNodeEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEdgeNodeEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEdgeNodeEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetEdgeNodeEventsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEdgeNodeEventsOK creates a GetEdgeNodeEventsOK with default headers values
func NewGetEdgeNodeEventsOK() *GetEdgeNodeEventsOK {
	return &GetEdgeNodeEventsOK{}
}

/* GetEdgeNodeEventsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetEdgeNodeEventsOK struct {
	Payload *swagger_models.EventQueryResponse
}

func (o *GetEdgeNodeEventsOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] getEdgeNodeEventsOK  %+v", 200, o.Payload)
}
func (o *GetEdgeNodeEventsOK) GetPayload() *swagger_models.EventQueryResponse {
	return o.Payload
}

func (o *GetEdgeNodeEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.EventQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeEventsUnauthorized creates a GetEdgeNodeEventsUnauthorized with default headers values
func NewGetEdgeNodeEventsUnauthorized() *GetEdgeNodeEventsUnauthorized {
	return &GetEdgeNodeEventsUnauthorized{}
}

/* GetEdgeNodeEventsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetEdgeNodeEventsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] getEdgeNodeEventsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetEdgeNodeEventsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeEventsForbidden creates a GetEdgeNodeEventsForbidden with default headers values
func NewGetEdgeNodeEventsForbidden() *GetEdgeNodeEventsForbidden {
	return &GetEdgeNodeEventsForbidden{}
}

/* GetEdgeNodeEventsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetEdgeNodeEventsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] getEdgeNodeEventsForbidden  %+v", 403, o.Payload)
}
func (o *GetEdgeNodeEventsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeEventsNotFound creates a GetEdgeNodeEventsNotFound with default headers values
func NewGetEdgeNodeEventsNotFound() *GetEdgeNodeEventsNotFound {
	return &GetEdgeNodeEventsNotFound{}
}

/* GetEdgeNodeEventsNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetEdgeNodeEventsNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] getEdgeNodeEventsNotFound  %+v", 404, o.Payload)
}
func (o *GetEdgeNodeEventsNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeEventsInternalServerError creates a GetEdgeNodeEventsInternalServerError with default headers values
func NewGetEdgeNodeEventsInternalServerError() *GetEdgeNodeEventsInternalServerError {
	return &GetEdgeNodeEventsInternalServerError{}
}

/* GetEdgeNodeEventsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetEdgeNodeEventsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeEventsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] getEdgeNodeEventsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEdgeNodeEventsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeEventsGatewayTimeout creates a GetEdgeNodeEventsGatewayTimeout with default headers values
func NewGetEdgeNodeEventsGatewayTimeout() *GetEdgeNodeEventsGatewayTimeout {
	return &GetEdgeNodeEventsGatewayTimeout{}
}

/* GetEdgeNodeEventsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetEdgeNodeEventsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeEventsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] getEdgeNodeEventsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetEdgeNodeEventsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeEventsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}