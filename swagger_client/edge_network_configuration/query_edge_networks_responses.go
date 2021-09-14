// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_network_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// QueryEdgeNetworksReader is a Reader for the QueryEdgeNetworks structure.
type QueryEdgeNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryEdgeNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryEdgeNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryEdgeNetworksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryEdgeNetworksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryEdgeNetworksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryEdgeNetworksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewQueryEdgeNetworksGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryEdgeNetworksOK creates a QueryEdgeNetworksOK with default headers values
func NewQueryEdgeNetworksOK() *QueryEdgeNetworksOK {
	return &QueryEdgeNetworksOK{}
}

/* QueryEdgeNetworksOK describes a response with status code 200, with default header values.

A successful response.
*/
type QueryEdgeNetworksOK struct {
	Payload *swagger_models.NetConfigList
}

func (o *QueryEdgeNetworksOK) Error() string {
	return fmt.Sprintf("[GET /v1/networks][%d] queryEdgeNetworksOK  %+v", 200, o.Payload)
}
func (o *QueryEdgeNetworksOK) GetPayload() *swagger_models.NetConfigList {
	return o.Payload
}

func (o *QueryEdgeNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.NetConfigList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeNetworksBadRequest creates a QueryEdgeNetworksBadRequest with default headers values
func NewQueryEdgeNetworksBadRequest() *QueryEdgeNetworksBadRequest {
	return &QueryEdgeNetworksBadRequest{}
}

/* QueryEdgeNetworksBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type QueryEdgeNetworksBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeNetworksBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/networks][%d] queryEdgeNetworksBadRequest  %+v", 400, o.Payload)
}
func (o *QueryEdgeNetworksBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeNetworksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeNetworksUnauthorized creates a QueryEdgeNetworksUnauthorized with default headers values
func NewQueryEdgeNetworksUnauthorized() *QueryEdgeNetworksUnauthorized {
	return &QueryEdgeNetworksUnauthorized{}
}

/* QueryEdgeNetworksUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type QueryEdgeNetworksUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeNetworksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/networks][%d] queryEdgeNetworksUnauthorized  %+v", 401, o.Payload)
}
func (o *QueryEdgeNetworksUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeNetworksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeNetworksForbidden creates a QueryEdgeNetworksForbidden with default headers values
func NewQueryEdgeNetworksForbidden() *QueryEdgeNetworksForbidden {
	return &QueryEdgeNetworksForbidden{}
}

/* QueryEdgeNetworksForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type QueryEdgeNetworksForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeNetworksForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/networks][%d] queryEdgeNetworksForbidden  %+v", 403, o.Payload)
}
func (o *QueryEdgeNetworksForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeNetworksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeNetworksInternalServerError creates a QueryEdgeNetworksInternalServerError with default headers values
func NewQueryEdgeNetworksInternalServerError() *QueryEdgeNetworksInternalServerError {
	return &QueryEdgeNetworksInternalServerError{}
}

/* QueryEdgeNetworksInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type QueryEdgeNetworksInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeNetworksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/networks][%d] queryEdgeNetworksInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryEdgeNetworksInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeNetworksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeNetworksGatewayTimeout creates a QueryEdgeNetworksGatewayTimeout with default headers values
func NewQueryEdgeNetworksGatewayTimeout() *QueryEdgeNetworksGatewayTimeout {
	return &QueryEdgeNetworksGatewayTimeout{}
}

/* QueryEdgeNetworksGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type QueryEdgeNetworksGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeNetworksGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/networks][%d] queryEdgeNetworksGatewayTimeout  %+v", 504, o.Payload)
}
func (o *QueryEdgeNetworksGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeNetworksGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}