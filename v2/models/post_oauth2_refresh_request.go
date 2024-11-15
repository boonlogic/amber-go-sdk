// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostOauth2RefreshRequest post oauth2 refresh request
//
// swagger:model postOauth2RefreshRequest
type PostOauth2RefreshRequest struct {

	// Refresh token used to acquire new access token
	// Required: true
	RefreshToken *string `json:"refreshToken"`
}

// Validate validates this post oauth2 refresh request
func (m *PostOauth2RefreshRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRefreshToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostOauth2RefreshRequest) validateRefreshToken(formats strfmt.Registry) error {

	if err := validate.Required("refreshToken", "body", m.RefreshToken); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post oauth2 refresh request based on context it is used
func (m *PostOauth2RefreshRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostOauth2RefreshRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostOauth2RefreshRequest) UnmarshalBinary(b []byte) error {
	var res PostOauth2RefreshRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}