// Code generated by go-swagger; DO NOT EDIT.

package passwords

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new passwords API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for passwords API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeletePassword(params *DeletePasswordParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePasswordOK, error)

	GetPassword(params *GetPasswordParams, authInfo runtime.ClientAuthInfoWriter) (*GetPasswordOK, error)

	PostCustomFields(params *PostCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomFieldsOK, error)

	PostUpdatePasswords(params *PostUpdatePasswordsParams, authInfo runtime.ClientAuthInfoWriter) (*PostUpdatePasswordsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeletePassword deletes password

  This API is used to delete the password with the password id supplied as the required argument. Note: You will only be able to delete the password if the supplied username has the correct permissions.
*/
func (a *Client) DeletePassword(params *DeletePasswordParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_Password",
		Method:             "DELETE",
		PathPattern:        "/api/1.0/passwords/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_Password: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPassword retrieves accounts usernames and passwords

  Get Usernames and Passwords
*/
func (a *Client) GetPassword(params *GetPasswordParams, authInfo runtime.ClientAuthInfoWriter) (*GetPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPassword",
		Method:             "GET",
		PathPattern:        "/api/1.0/passwords/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCustomFields Custom Fields
*/
func (a *Client) PostCustomFields(params *PostCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomFieldsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomFieldsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post_Custom_Fields",
		Method:             "POST",
		PathPattern:        "/api/1.0/custom_fields/password/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostCustomFieldsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCustomFieldsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_Custom_Fields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUpdatePasswords Create / Update Passwords. Use id if updating existing password - else, username and password are required.
*/
func (a *Client) PostUpdatePasswords(params *PostUpdatePasswordsParams, authInfo runtime.ClientAuthInfoWriter) (*PostUpdatePasswordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUpdatePasswordsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post_Update_Passwords",
		Method:             "POST",
		PathPattern:        "/api/1.0/passwords/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostUpdatePasswordsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUpdatePasswordsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_Update_Passwords: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
