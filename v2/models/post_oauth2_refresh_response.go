// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostOauth2RefreshResponse post oauth2 refresh response
//
// swagger:model postOauth2RefreshResponse
type PostOauth2RefreshResponse struct {

	// Number of seconds before `idToken` becomes invalid.
	ExpiresIn string `json:"expiresIn,omitempty"`

	// Authorization token. Requests to all API endpoints except `POST /oauth2` must include a valid token in a header field called `Authorization` with value `Bearer ${idToken}`.
	IDToken string `json:"idToken,omitempty"`

	// Can be used to refresh the `idToken` when it is nearing expiration.
	RefreshToken string `json:"refreshToken,omitempty"`

	// Always `Bearer`.
	TokenType string `json:"tokenType,omitempty"`
}

// Validate validates this post oauth2 refresh response
func (m *PostOauth2RefreshResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post oauth2 refresh response based on context it is used
func (m *PostOauth2RefreshResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostOauth2RefreshResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostOauth2RefreshResponse) UnmarshalBinary(b []byte) error {
	var res PostOauth2RefreshResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
