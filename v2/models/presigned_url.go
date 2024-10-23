// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PresignedURL presigned URL
//
// swagger:model presignedURL
type PresignedURL struct {

	// feature name
	Name string `json:"name,omitempty"`

	// url used for downloading
	URL string `json:"url,omitempty"`
}

// Validate validates this presigned URL
func (m *PresignedURL) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this presigned URL based on context it is used
func (m *PresignedURL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PresignedURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PresignedURL) UnmarshalBinary(b []byte) error {
	var res PresignedURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
