// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// StreamingWindow Number of recent input _vectors_ concatenated together to make up a full input _pattern_ presented to the model for inference.
//
// Let `featureCount` be the configured number of features. The model consumes data sequentially in steps of size `featureCount`. Each time it receives `featureCount` data values, `featureCount` input values are consumed and concatenated together to form an input _vector_. This input vector is then concatenated with zero or more past input vectors to form an input _pattern_. The input _pattern_ is the true data vector inferenced by the model at each step. Configuring the `streamingWindow` greater than 1 allows a model to identify patterns in vectors that change over time.
//
// If monitoring a single timeseries signal, the model should be configured with just one feature. In that case the input vector has length 1, and `streamingWindow` determines the length of a moving window over past samples which is the input pattern to the model for each new sample. For example, a model configured with one feature and a `streamingWindow` of 25 will concatenate together and inference the 25 most recent data values for each new value consumed.
//
// If monitoring instantaneous readings from several sensors jointly, each sensor should be associated with one feature in the configuration. In this case `streamingWindow` is usually set to 1 so that the input pattern is just the current vector of readings. For example, a model configured with 5 features and a `streamingWindow` of 1 will consume 5 values at a time and inference those 5 values as a pattern of length 5. If the `streamingWindow` were 2, the model would still consume 5 values at a time, but its input pattern would contain the last 10 samples.
//
// swagger:model streamingWindow
type StreamingWindow uint16

// Validate validates this streaming window
func (m StreamingWindow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinimumUint("", "body", uint64(m), 1, false); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this streaming window based on context it is used
func (m StreamingWindow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
