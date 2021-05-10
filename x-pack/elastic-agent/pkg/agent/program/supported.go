// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

package program

import (
	"strings"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"
)

var Supported []Spec
var SupportedMap map[string]Spec

func init() {
	// Packed Files
	// spec/apm-server.yml
	// spec/endpoint.yml
	// spec/filebeat.yml
	// spec/fleet-server.yml
	// spec/heartbeat.yml
	// spec/metricbeat.yml
	// spec/osquerybeat.yml
	// spec/packetbeat.yml
	unpacked := packer.MustUnpack("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")
	SupportedMap = make(map[string]Spec)

	for f, v := range unpacked {
		s, err := NewSpecFromBytes(v)
		if err != nil {
			panic("Cannot read spec from " + f)
		}
		Supported = append(Supported, s)
		SupportedMap[strings.ToLower(s.Cmd)] = s
	}
}