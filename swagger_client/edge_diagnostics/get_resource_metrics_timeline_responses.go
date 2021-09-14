// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetResourceMetricsTimelineReader is a Reader for the GetResourceMetricsTimeline structure.
type GetResourceMetricsTimelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceMetricsTimelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceMetricsTimelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResourceMetricsTimelineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetResourceMetricsTimelineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResourceMetricsTimelineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResourceMetricsTimelineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetResourceMetricsTimelineGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourceMetricsTimelineOK creates a GetResourceMetricsTimelineOK with default headers values
func NewGetResourceMetricsTimelineOK() *GetResourceMetricsTimelineOK {
	return &GetResourceMetricsTimelineOK{}
}

/* GetResourceMetricsTimelineOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetResourceMetricsTimelineOK struct {
	Payload *swagger_models.MetricQueryResponse
}

func (o *GetResourceMetricsTimelineOK) Error() string {
	return fmt.Sprintf("[GET /v1/events/timeSeries/{mType}][%d] getResourceMetricsTimelineOK  %+v", 200, o.Payload)
}
func (o *GetResourceMetricsTimelineOK) GetPayload() *swagger_models.MetricQueryResponse {
	return o.Payload
}

func (o *GetResourceMetricsTimelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.MetricQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceMetricsTimelineBadRequest creates a GetResourceMetricsTimelineBadRequest with default headers values
func NewGetResourceMetricsTimelineBadRequest() *GetResourceMetricsTimelineBadRequest {
	return &GetResourceMetricsTimelineBadRequest{}
}

/* GetResourceMetricsTimelineBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type GetResourceMetricsTimelineBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceMetricsTimelineBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/events/timeSeries/{mType}][%d] getResourceMetricsTimelineBadRequest  %+v", 400, o.Payload)
}
func (o *GetResourceMetricsTimelineBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceMetricsTimelineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceMetricsTimelineUnauthorized creates a GetResourceMetricsTimelineUnauthorized with default headers values
func NewGetResourceMetricsTimelineUnauthorized() *GetResourceMetricsTimelineUnauthorized {
	return &GetResourceMetricsTimelineUnauthorized{}
}

/* GetResourceMetricsTimelineUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetResourceMetricsTimelineUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceMetricsTimelineUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/events/timeSeries/{mType}][%d] getResourceMetricsTimelineUnauthorized  %+v", 401, o.Payload)
}
func (o *GetResourceMetricsTimelineUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceMetricsTimelineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceMetricsTimelineForbidden creates a GetResourceMetricsTimelineForbidden with default headers values
func NewGetResourceMetricsTimelineForbidden() *GetResourceMetricsTimelineForbidden {
	return &GetResourceMetricsTimelineForbidden{}
}

/* GetResourceMetricsTimelineForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-app level access permission for the operation or does not have access scope to the project.
*/
type GetResourceMetricsTimelineForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceMetricsTimelineForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/events/timeSeries/{mType}][%d] getResourceMetricsTimelineForbidden  %+v", 403, o.Payload)
}
func (o *GetResourceMetricsTimelineForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceMetricsTimelineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceMetricsTimelineInternalServerError creates a GetResourceMetricsTimelineInternalServerError with default headers values
func NewGetResourceMetricsTimelineInternalServerError() *GetResourceMetricsTimelineInternalServerError {
	return &GetResourceMetricsTimelineInternalServerError{}
}

/* GetResourceMetricsTimelineInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetResourceMetricsTimelineInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceMetricsTimelineInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/events/timeSeries/{mType}][%d] getResourceMetricsTimelineInternalServerError  %+v", 500, o.Payload)
}
func (o *GetResourceMetricsTimelineInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceMetricsTimelineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceMetricsTimelineGatewayTimeout creates a GetResourceMetricsTimelineGatewayTimeout with default headers values
func NewGetResourceMetricsTimelineGatewayTimeout() *GetResourceMetricsTimelineGatewayTimeout {
	return &GetResourceMetricsTimelineGatewayTimeout{}
}

/* GetResourceMetricsTimelineGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetResourceMetricsTimelineGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceMetricsTimelineGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/events/timeSeries/{mType}][%d] getResourceMetricsTimelineGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetResourceMetricsTimelineGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceMetricsTimelineGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}