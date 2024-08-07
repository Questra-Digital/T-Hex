// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import "github.com/ory/herodot"

// JSON API Error Response
//
// The standard Ory JSON API error format.
//
// swagger:model errorGeneric
//
//nolint:deadcode,unused
//lint:ignore U1000 Used to generate Swagger and OpenAPI definitions
type errorGeneric struct {
	// Contains error details
	//
	// required: true
	Error herodot.DefaultError `json:"error"`
}

// Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is typically 201.
//
// swagger:response emptyResponse
//
//nolint:deadcode,unused
//lint:ignore U1000 Used to generate Swagger and OpenAPI definitions
type emptyResponse struct{}
