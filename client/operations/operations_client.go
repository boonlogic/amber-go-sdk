// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteSensor(params *DeleteSensorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSensorOK, error)

	GetAmberSummary(params *GetAmberSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAmberSummaryOK, error)

	GetConfig(params *GetConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConfigOK, error)

	GetPretrain(params *GetPretrainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPretrainOK, *GetPretrainAccepted, error)

	GetRootCause(params *GetRootCauseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRootCauseOK, error)

	GetSensor(params *GetSensorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSensorOK, error)

	GetSensors(params *GetSensorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSensorsOK, error)

	GetStatus(params *GetStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStatusOK, error)

	GetVersion(params *GetVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVersionOK, error)

	PostConfig(params *PostConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostConfigOK, error)

	PostOauth2(params *PostOauth2Params, opts ...ClientOption) (*PostOauth2OK, error)

	PostOutage(params *PostOutageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOutageOK, error)

	PostPretrain(params *PostPretrainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostPretrainOK, *PostPretrainAccepted, error)

	PostSensor(params *PostSensorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostSensorOK, error)

	PostStream(params *PostStreamParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostStreamOK, error)

	PutConfig(params *PutConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutConfigOK, error)

	PutSensor(params *PutSensorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutSensorOK, error)

	PutStream(params *PutStreamParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutStreamOK, *PutStreamAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteSensor deletes a sensor instance

  Deletes the sensor instance with the specified sensorId.
*/
func (a *Client) DeleteSensor(params *DeleteSensorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSensorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSensorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSensor",
		Method:             "DELETE",
		PathPattern:        "/sensor",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteSensorReader{formats: a.formats},
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
	success, ok := result.(*DeleteSensorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteSensor: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAmberSummary gets the JSON block of the amber image

  Returns the json block of the amber sensor
*/
func (a *Client) GetAmberSummary(params *GetAmberSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAmberSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAmberSummaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAmberSummary",
		Method:             "GET",
		PathPattern:        "/__summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAmberSummaryReader{formats: a.formats},
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
	success, ok := result.(*GetAmberSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAmberSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetConfig gets the current configuration of a sensor instance

  Returns the current configuration of the sensor instance specified.
*/
func (a *Client) GetConfig(params *GetConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getConfig",
		Method:             "GET",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConfigReader{formats: a.formats},
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
	success, ok := result.(*GetConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPretrain gets status of pretrain operation

  Get status of a sensor which is currently pretraining.
*/
func (a *Client) GetPretrain(params *GetPretrainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPretrainOK, *GetPretrainAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPretrainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPretrain",
		Method:             "GET",
		PathPattern:        "/pretrain",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPretrainReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetPretrainOK:
		return value, nil, nil
	case *GetPretrainAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for operations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRootCause gets root cause analysis information from a sensor

  Returns analytic information on the root cause for the clusters provided.
*/
func (a *Client) GetRootCause(params *GetRootCauseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRootCauseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRootCauseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRootCause",
		Method:             "GET",
		PathPattern:        "/rootCause",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRootCauseReader{formats: a.formats},
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
	success, ok := result.(*GetRootCauseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRootCause: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSensor gets basic information about a sensor instance

  Returns basic information about an existing sensor instance.
*/
func (a *Client) GetSensor(params *GetSensorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSensorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSensorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSensor",
		Method:             "GET",
		PathPattern:        "/sensor",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSensorReader{formats: a.formats},
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
	success, ok := result.(*GetSensorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSensor: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSensors lists all sensors for this user

  Returns a list of all current sensor instances for this user.
*/
func (a *Client) GetSensors(params *GetSensorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSensorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSensorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSensors",
		Method:             "GET",
		PathPattern:        "/sensors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSensorsReader{formats: a.formats},
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
	success, ok := result.(*GetSensorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSensors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStatus gets analytic information from a sensor

  Returns analytic information derived from data processed by a sensor thus far.
*/
func (a *Client) GetStatus(params *GetStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStatus",
		Method:             "GET",
		PathPattern:        "/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStatusReader{formats: a.formats},
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
	success, ok := result.(*GetStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVersion retrieves API version information
*/
func (a *Client) GetVersion(params *GetVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVersion",
		Method:             "GET",
		PathPattern:        "/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetVersionReader{formats: a.formats},
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
	success, ok := result.(*GetVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostConfig applies configuration to a sensor instance

  Applies the provided configuration to the sensor instance specified.
*/
func (a *Client) PostConfig(params *PostConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postConfig",
		Method:             "POST",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostConfigReader{formats: a.formats},
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
	success, ok := result.(*PostConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostOauth2 requests a bearer token using amber account credentials

  Requests a bearer token using Amber account credentials. The requested bearer token is returned as the "id-token" response attribute. This token is to be used for authenticating API requests throughout a usage session and expires after 60 minutes.
*/
func (a *Client) PostOauth2(params *PostOauth2Params, opts ...ClientOption) (*PostOauth2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOauth2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "postOauth2",
		Method:             "POST",
		PathPattern:        "/oauth2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostOauth2Reader{formats: a.formats},
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
	success, ok := result.(*PostOauth2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postOauth2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostOutage informs the server of an outage

  Clears the load buffer of streaming window and resets statistics. Returns stream status
*/
func (a *Client) PostOutage(params *PostOutageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOutageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOutageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postOutage",
		Method:             "POST",
		PathPattern:        "/outage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostOutageReader{formats: a.formats},
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
	success, ok := result.(*PostOutageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postOutage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostPretrain pretrains a sensor using historical data

  Pretrains a sensor. Ingoing data should be formatted as a simple string of comma-separated numbers with no spaces.
*/
func (a *Client) PostPretrain(params *PostPretrainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostPretrainOK, *PostPretrainAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPretrainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postPretrain",
		Method:             "POST",
		PathPattern:        "/pretrain",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPretrainReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostPretrainOK:
		return value, nil, nil
	case *PostPretrainAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for operations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostSensor creates a new a sensor instance

  Spawns a new sensor instance, returning its unique sensorId.
*/
func (a *Client) PostSensor(params *PostSensorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostSensorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSensorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postSensor",
		Method:             "POST",
		PathPattern:        "/sensor",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostSensorReader{formats: a.formats},
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
	success, ok := result.(*PostSensorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postSensor: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostStream streams data to a sensor

  Sends data to a sensor. Ingoing data should be formatted as a simple string of comma-separated numbers with no spaces.

The following analytic results are returned:
- state : sensor state as of this call (one of: "Buffering", "Autotuning", "Learning", "Monitoring", "Error")
- ID : array of cluster IDs. These correspond one-to-one with input samples, indicating the cluster to which each input pattern was assigned.
- SI : array of smoothed anomaly index values. These values correspond one-to-one with the input samples and range between 0 and 1000. Values closer to 0 represent input patterns which are ordinary given the data seen so far on this sensor. Values closer to 1000 represent novel patterns which are anomalous with respect to data seen before.
- AD : array of 0's and 1's as anomaly detection indicators. These correspond one-to-one with input samples and are produced by thresholding the smoothed anomaly index (SI). The threshold is determined automatically from the SI values. A value of 0 indicates that the SI has not exceeded the anomaly detection threshold. A value of 1 indicates it has, signaling an anomaly at the corresponding input sample.
- AH : array of anomaly history values. These values are a moving-window sum of the AD, giving the number of anomaly detections (1's) present in the AD signal over a "recent history" window whose length is the buffer size.
- AM : array of Amber Metric values. These are floating-point values between 0.0 and 1.0 indicating the extent to which the AH contains an unusually high number of anomalies in recent history. The values are derived statistically from a Poisson model, with values close to 0.0 signaling a lower, and values close to 1.0 signaling a higher, frequency of anomalies than usual.
- AW : array of Amber Warning Level values. This index is produced by thresholding the Amber Metric (AM) and takes on the values 0, 1 or 2 representing a discrete "warning level" for an asset based on the frequency of anomalies within recent history. 0 = normal, 1 = asset changing, 2 = asset critical. The default thresholds for the two warning levels are the standard statistical values of 0.95 (outlier, asset chaing) and 0.997 (extreme outlier, asset critical).
*/
func (a *Client) PostStream(params *PostStreamParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostStreamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStreamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postStream",
		Method:             "POST",
		PathPattern:        "/stream",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostStreamReader{formats: a.formats},
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
	success, ok := result.(*PostStreamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postStream: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutConfig updates configuration for a sensor instance

  Updates the configuration for the sensor instance specified.
*/
func (a *Client) PutConfig(params *PutConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putConfig",
		Method:             "PUT",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutConfigReader{formats: a.formats},
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
	success, ok := result.(*PutConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutSensor updates label for a sensor instance

  Changes the label of an existing sensor instance to the new label specified.
*/
func (a *Client) PutSensor(params *PutSensorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutSensorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutSensorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putSensor",
		Method:             "PUT",
		PathPattern:        "/sensor",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutSensorReader{formats: a.formats},
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
	success, ok := result.(*PutSensorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putSensor: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutStream streams data to a sensor fusion vector

  Update fusion vector with new values for the given features, and optionally submit to Amber. Analytic results returned are the same as POST /stream.
*/
func (a *Client) PutStream(params *PutStreamParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutStreamOK, *PutStreamAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutStreamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putStream",
		Method:             "PUT",
		PathPattern:        "/stream",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutStreamReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PutStreamOK:
		return value, nil, nil
	case *PutStreamAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for operations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
