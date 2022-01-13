// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostFeatureConfig post feature config
//
// swagger:model postFeatureConfig
type PostFeatureConfig struct {

	// label associated with feature
	Label *string `json:"label,omitempty"`

	// corresponding maximum value
	MaxVal *float32 `json:"maxVal,omitempty"`

	// the value that should be considered the minimum value for this feature. This can be set to a value larger than the actual min if you want to treat all value less than that as the same (for instance, to keep a noise spike from having undue influence in the clustering
	MinVal *float32 `json:"minVal,omitempty"`
}

// Validate validates this post feature config
func (m *PostFeatureConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post feature config based on context it is used
func (m *PostFeatureConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostFeatureConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostFeatureConfig) UnmarshalBinary(b []byte) error {
	var res PostFeatureConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}