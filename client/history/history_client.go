// Code generated by go-swagger; DO NOT EDIT.

package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new history API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for history API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetHistory(params *GetHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetHistoryOK, error)

	GetHistoryNumberOfWeeks(params *GetHistoryNumberOfWeeksParams, authInfo runtime.ClientAuthInfoWriter) (*GetHistoryNumberOfWeeksOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetHistory retrieves all history for the past one week note as of 10 3 0 device42 now has a more robust audit log available at api api 1 0 auditlogs

  Get History
*/
func (a *Client) GetHistory(params *GetHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHistory",
		Method:             "GET",
		PathPattern:        "/api/1.0/history/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetHistoryNumberOfWeeks retrieves history for specified number of weeks in the past

  Get History by # of Weeks
*/
func (a *Client) GetHistoryNumberOfWeeks(params *GetHistoryNumberOfWeeksParams, authInfo runtime.ClientAuthInfoWriter) (*GetHistoryNumberOfWeeksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHistoryNumberOfWeeksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHistoryNumber_of_weeks",
		Method:             "GET",
		PathPattern:        "/api/1.0/history/<number_of_weeks>/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetHistoryNumberOfWeeksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHistoryNumberOfWeeksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getHistoryNumber_of_weeks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
