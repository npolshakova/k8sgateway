// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/extensions/filters/http/csrf/v3/csrf.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *CsrfPolicy) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CsrfPolicy)
	if !ok {
		that2, ok := that.(CsrfPolicy)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetFilterEnabled()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFilterEnabled()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFilterEnabled(), target.GetFilterEnabled()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetShadowEnabled()).(equality.Equalizer); ok {
		if !h.Equal(target.GetShadowEnabled()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetShadowEnabled(), target.GetShadowEnabled()) {
			return false
		}
	}

	if len(m.GetAdditionalOrigins()) != len(target.GetAdditionalOrigins()) {
		return false
	}
	for idx, v := range m.GetAdditionalOrigins() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAdditionalOrigins()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAdditionalOrigins()[idx]) {
				return false
			}
		}

	}

	return true
}
