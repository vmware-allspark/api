// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for PeerAuthentication
func (this *PeerAuthentication) MarshalJSON() ([]byte, error) {
	str, err := PeerAuthenticationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PeerAuthentication
func (this *PeerAuthentication) UnmarshalJSON(b []byte) error {
	return PeerAuthenticationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for PeerAuthentication_MutualTLS
func (this *PeerAuthentication_MutualTLS) MarshalJSON() ([]byte, error) {
	str, err := PeerAuthenticationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PeerAuthentication_MutualTLS
func (this *PeerAuthentication_MutualTLS) UnmarshalJSON(b []byte) error {
	return PeerAuthenticationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	PeerAuthenticationMarshaler   = &jsonpb.Marshaler{}
	PeerAuthenticationUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)
