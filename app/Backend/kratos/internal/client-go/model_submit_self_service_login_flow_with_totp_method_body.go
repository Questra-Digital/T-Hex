// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version:
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceLoginFlowWithTotpMethodBody struct for SubmitSelfServiceLoginFlowWithTotpMethodBody
type SubmitSelfServiceLoginFlowWithTotpMethodBody struct {
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method should be set to \"totp\" when logging in using the TOTP strategy.
	Method string `json:"method"`
	// The TOTP code.
	TotpCode string `json:"totp_code"`
}

// NewSubmitSelfServiceLoginFlowWithTotpMethodBody instantiates a new SubmitSelfServiceLoginFlowWithTotpMethodBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceLoginFlowWithTotpMethodBody(method string, totpCode string) *SubmitSelfServiceLoginFlowWithTotpMethodBody {
	this := SubmitSelfServiceLoginFlowWithTotpMethodBody{}
	this.Method = method
	this.TotpCode = totpCode
	return &this
}

// NewSubmitSelfServiceLoginFlowWithTotpMethodBodyWithDefaults instantiates a new SubmitSelfServiceLoginFlowWithTotpMethodBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceLoginFlowWithTotpMethodBodyWithDefaults() *SubmitSelfServiceLoginFlowWithTotpMethodBody {
	this := SubmitSelfServiceLoginFlowWithTotpMethodBody{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceLoginFlowWithTotpMethodBody) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithTotpMethodBody) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceLoginFlowWithTotpMethodBody) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceLoginFlowWithTotpMethodBody) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *SubmitSelfServiceLoginFlowWithTotpMethodBody) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithTotpMethodBody) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitSelfServiceLoginFlowWithTotpMethodBody) SetMethod(v string) {
	o.Method = v
}

// GetTotpCode returns the TotpCode field value
func (o *SubmitSelfServiceLoginFlowWithTotpMethodBody) GetTotpCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotpCode
}

// GetTotpCodeOk returns a tuple with the TotpCode field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithTotpMethodBody) GetTotpCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotpCode, true
}

// SetTotpCode sets field value
func (o *SubmitSelfServiceLoginFlowWithTotpMethodBody) SetTotpCode(v string) {
	o.TotpCode = v
}

func (o SubmitSelfServiceLoginFlowWithTotpMethodBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["totp_code"] = o.TotpCode
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceLoginFlowWithTotpMethodBody struct {
	value *SubmitSelfServiceLoginFlowWithTotpMethodBody
	isSet bool
}

func (v NullableSubmitSelfServiceLoginFlowWithTotpMethodBody) Get() *SubmitSelfServiceLoginFlowWithTotpMethodBody {
	return v.value
}

func (v *NullableSubmitSelfServiceLoginFlowWithTotpMethodBody) Set(val *SubmitSelfServiceLoginFlowWithTotpMethodBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceLoginFlowWithTotpMethodBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceLoginFlowWithTotpMethodBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceLoginFlowWithTotpMethodBody(val *SubmitSelfServiceLoginFlowWithTotpMethodBody) *NullableSubmitSelfServiceLoginFlowWithTotpMethodBody {
	return &NullableSubmitSelfServiceLoginFlowWithTotpMethodBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceLoginFlowWithTotpMethodBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceLoginFlowWithTotpMethodBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
