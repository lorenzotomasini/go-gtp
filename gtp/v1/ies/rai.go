// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"

	"github.com/wmnsk/go-gtp/gtp/utils"
)

// NewRouteingAreaIdentity creates a new RouteingAreaIdentity IE.
func NewRouteingAreaIdentity(mcc, mnc string, lac uint16, rac uint8) *IE {
	mc, err := utils.StrToSwappedBytes(mcc, "f")
	if err != nil {
		return nil
	}
	mn, err := utils.StrToSwappedBytes(mnc, "f")
	if err != nil {
		return nil
	}

	rai := New(
		RouteingAreaIdentity,
		make([]byte, 6),
	)
	copy(rai.Payload[0:2], mc)
	rai.Payload[2] = mn[0]
	binary.BigEndian.PutUint16(rai.Payload[3:5], lac)
	rai.Payload[5] = rac

	return rai
}

// RouteingAreaIdentity returns RouteingAreaIdentity value if type matches.
func (i *IE) RouteingAreaIdentity() []byte {
	if i.Type != RouteingAreaIdentity {
		return nil
	}
	return i.Payload
}