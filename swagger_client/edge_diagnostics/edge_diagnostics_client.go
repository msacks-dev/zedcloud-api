// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new edge diagnostics API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for edge diagnostics API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDeviceTwinConfig(params *GetDeviceTwinConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTwinConfigOK, error)

	GetDeviceTwinConfigByName(params *GetDeviceTwinConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTwinConfigByNameOK, error)

	GetDeviceTwinNextConfig(params *GetDeviceTwinNextConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTwinNextConfigOK, error)

	GetDeviceTwinNextConfigByName(params *GetDeviceTwinNextConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTwinNextConfigByNameOK, error)

	GetDeviceTwinOfflineConfigByName(params *GetDeviceTwinOfflineConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTwinOfflineConfigByNameOK, error)

	GetDeviceTwinOfflineNextConfig(params *GetDeviceTwinOfflineNextConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTwinOfflineNextConfigOK, error)

	GetEventsTimeline(params *GetEventsTimelineParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEventsTimelineOK, error)

	GetResourceMetricsTimeline(params *GetResourceMetricsTimelineParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetResourceMetricsTimelineOK, error)

	GetResourceMetricsTimeline2(params *GetResourceMetricsTimeline2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetResourceMetricsTimeline2OK, error)

	RegenDeviceConfig(params *RegenDeviceConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegenDeviceConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDeviceTwinConfig gets current device twin configuration

  Get currentnext Device twin configuration for the edge node. Edge node has read this configuration when it queried Cloud controller last time.
*/
func (a *Client) GetDeviceTwinConfig(params *GetDeviceTwinConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTwinConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceTwinConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeviceTwinConfig",
		Method:             "GET",
		PathPattern:        "/v1/devices/id/{id}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceTwinConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceTwinConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDeviceTwinConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeviceTwinConfigByName gets current device twin configuration

  Get currentnext Device twin configuration for the edge node. Edge node has read this configuration when it queried Cloud controller last time.
*/
func (a *Client) GetDeviceTwinConfigByName(params *GetDeviceTwinConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTwinConfigByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceTwinConfigByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeviceTwinConfigByName",
		Method:             "GET",
		PathPattern:        "/v1/devices/name/{name}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceTwinConfigByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceTwinConfigByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDeviceTwinConfigByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeviceTwinNextConfig gets next device twin configuration

  Get next Device twin configuration for the edge node. Edge node will get this configuration when it queries Cloud controller next time.
*/
func (a *Client) GetDeviceTwinNextConfig(params *GetDeviceTwinNextConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTwinNextConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceTwinNextConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeviceTwinNextConfig",
		Method:             "GET",
		PathPattern:        "/v1/devices/id/{id}/config/next",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceTwinNextConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceTwinNextConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDeviceTwinNextConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeviceTwinNextConfigByName gets next device twin configuration

  Get next Device twin configuration for the edge node. Edge node will get this configuration when it queries Cloud controller next time.
*/
func (a *Client) GetDeviceTwinNextConfigByName(params *GetDeviceTwinNextConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTwinNextConfigByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceTwinNextConfigByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeviceTwinNextConfigByName",
		Method:             "GET",
		PathPattern:        "/v1/devices/name/{name}/config/next",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceTwinNextConfigByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceTwinNextConfigByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDeviceTwinNextConfigByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeviceTwinOfflineConfigByName gets offline device twin configuration

  Get offline Device twin configuration for the edge node. Edge node will get this configuration when it queries Cloud controller next time.
*/
func (a *Client) GetDeviceTwinOfflineConfigByName(params *GetDeviceTwinOfflineConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTwinOfflineConfigByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceTwinOfflineConfigByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeviceTwinOfflineConfigByName",
		Method:             "GET",
		PathPattern:        "/v1/devices/name/{name}/config/offline",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceTwinOfflineConfigByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceTwinOfflineConfigByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDeviceTwinOfflineConfigByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeviceTwinOfflineNextConfig gets offline device twin configuration

  Get offline Device twin configuration for the edge node. Edge node will get this configuration when it queries Cloud controller next time.
*/
func (a *Client) GetDeviceTwinOfflineNextConfig(params *GetDeviceTwinOfflineNextConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTwinOfflineNextConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceTwinOfflineNextConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeviceTwinOfflineNextConfig",
		Method:             "GET",
		PathPattern:        "/v1/devices/id/{id}/config/offline",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceTwinOfflineNextConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceTwinOfflineNextConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDeviceTwinOfflineNextConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEventsTimeline gets events timeline

  Get aggregated events timeline
*/
func (a *Client) GetEventsTimeline(params *GetEventsTimelineParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEventsTimelineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventsTimelineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEventsTimeline",
		Method:             "GET",
		PathPattern:        "/v1/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEventsTimelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEventsTimelineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEventsTimeline: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetResourceMetricsTimeline gets resource usage timeline

  Get the aggregated resource usage timeline as reported by the edge nodes and edge application instances.
*/
func (a *Client) GetResourceMetricsTimeline(params *GetResourceMetricsTimelineParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetResourceMetricsTimelineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceMetricsTimelineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetResourceMetricsTimeline",
		Method:             "GET",
		PathPattern:        "/v1/events/timeSeries/{mType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourceMetricsTimelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetResourceMetricsTimelineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetResourceMetricsTimeline: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetResourceMetricsTimeline2 gets resource usage timeline

  Get the aggregated resource usage timeline as reported by the edge nodes and edge application instances.
*/
func (a *Client) GetResourceMetricsTimeline2(params *GetResourceMetricsTimeline2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetResourceMetricsTimeline2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceMetricsTimeline2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetResourceMetricsTimeline2",
		Method:             "GET",
		PathPattern:        "/v1/timeSeries/{mType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourceMetricsTimeline2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetResourceMetricsTimeline2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetResourceMetricsTimeline2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RegenDeviceConfig res generate device configuration

  Re-generate the device configuration. Edge node will get this configuration when it queries Cloud controller next time.
*/
func (a *Client) RegenDeviceConfig(params *RegenDeviceConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegenDeviceConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegenDeviceConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RegenDeviceConfig",
		Method:             "PUT",
		PathPattern:        "/v1/devices/id/{id}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RegenDeviceConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RegenDeviceConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RegenDeviceConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}