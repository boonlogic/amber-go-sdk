// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// VersionNumber version number
//
// swagger:model VersionNumber
type VersionNumber uint16

// Validate validates this version number
func (m VersionNumber) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this version number based on context it is used
func (m VersionNumber) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}