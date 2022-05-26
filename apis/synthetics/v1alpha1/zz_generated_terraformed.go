/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Test
func (mg *Test) GetTerraformResourceType() string {
	return "datadog_synthetics_test"
}

// GetConnectionDetailsMapping for this Test
func (tr *Test) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"api_step[*].request_basicauth[*].access_key": "spec.forProvider.apiStep[*].requestBasicauth[*].accessKeySecretRef", "api_step[*].request_basicauth[*].password": "spec.forProvider.apiStep[*].requestBasicauth[*].passwordSecretRef", "api_step[*].request_basicauth[*].secret_key": "spec.forProvider.apiStep[*].requestBasicauth[*].secretKeySecretRef", "api_step[*].request_client_certificate[*].cert[*].content": "spec.forProvider.apiStep[*].requestClientCertificate[*].cert[*].contentSecretRef", "api_step[*].request_client_certificate[*].key[*].content": "spec.forProvider.apiStep[*].requestClientCertificate[*].key[*].contentSecretRef", "request_basicauth[*].access_key": "spec.forProvider.requestBasicauth[*].accessKeySecretRef", "request_basicauth[*].password": "spec.forProvider.requestBasicauth[*].passwordSecretRef", "request_basicauth[*].secret_key": "spec.forProvider.requestBasicauth[*].secretKeySecretRef", "request_client_certificate[*].cert[*].content": "spec.forProvider.requestClientCertificate[*].cert[*].contentSecretRef", "request_client_certificate[*].key[*].content": "spec.forProvider.requestClientCertificate[*].key[*].contentSecretRef"}
}

// GetObservation of this Test
func (tr *Test) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Test
func (tr *Test) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Test
func (tr *Test) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Test
func (tr *Test) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Test
func (tr *Test) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Test using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Test) LateInitialize(attrs []byte) (bool, error) {
	params := &TestParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Test) GetTerraformSchemaVersion() int {
	return 0
}
