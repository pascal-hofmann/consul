// Code generated by protoc-json-shim. DO NOT EDIT.
package hcpv1

import (
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON is a custom marshaler for TelemetryState
func (this *TelemetryState) MarshalJSON() ([]byte, error) {
	str, err := TelemetryStateMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TelemetryState
func (this *TelemetryState) UnmarshalJSON(b []byte) error {
	return TelemetryStateUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for TelemetryState_MetricsState
func (this *TelemetryState_MetricsState) MarshalJSON() ([]byte, error) {
	str, err := TelemetryStateMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TelemetryState_MetricsState
func (this *TelemetryState_MetricsState) UnmarshalJSON(b []byte) error {
	return TelemetryStateUnmarshaler.Unmarshal(b, this)
}

var (
	TelemetryStateMarshaler   = &protojson.MarshalOptions{}
	TelemetryStateUnmarshaler = &protojson.UnmarshalOptions{DiscardUnknown: false}
)
