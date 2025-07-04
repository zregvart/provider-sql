/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/crossplane-contrib/provider-sql/pkg/clients"
	"github.com/crossplane-contrib/provider-sql/pkg/clients/postgresql"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// Credentials required to authenticate to this provider.
	Credentials ProviderCredentials `json:"credentials"`
	// Defines the database name used to set up a connection to the provided
	// PostgreSQL instance. Same as PGDATABASE environment variable.
	// +kubebuilder:default="postgres"
	DefaultDatabase string `json:"defaultDatabase,omitempty"`
	// Defines the SSL mode used to set up a connection to the provided
	// PostgreSQL instance
	// +kubebuilder:validation:Enum=disable;require;verify-ca;verify-full
	// +kubebuilder:default=verify-full
	// +kubebuilder:validation:Optional
	SSLMode *string `json:"sslMode,omitempty"`
	// Path to the certificate used for client authentication
	// +kubebuilder:validation:Optional
	SSLCert *string `json:"sslCert,omitempty"`
	// Path to the key used for client authentication
	// +kubebuilder:validation:Optional
	SSLKey *string `json:"sslKey,omitempty"`
	// Path to the CA certificate(s) used for verifying the server certificate
	// +kubebuilder:validation:Optional
	SSLRootCert *string `json:"sslRootCert,omitempty"`
}

func (s ProviderConfigSpec) Options() postgresql.Options {
	return postgresql.Options{
		SSLMode:     clients.ToString(s.SSLMode),
		SSLCert:     clients.ToString(s.SSLCert),
		SSLKey:      clients.ToString(s.SSLKey),
		SSLRootCert: clients.ToString(s.SSLRootCert),
	}
}

const (
	// CredentialsSourcePostgreSQLConnectionSecret indicates that a provider
	// should acquire credentials from a connection secret written by a managed
	// resource that represents a PostgreSQL server.
	CredentialsSourcePostgreSQLConnectionSecret xpv1.CredentialsSource = "PostgreSQLConnectionSecret"
)

// ProviderCredentials required to authenticate.
type ProviderCredentials struct {
	// Source of the provider credentials.
	// +kubebuilder:validation:Enum=PostgreSQLConnectionSecret
	Source xpv1.CredentialsSource `json:"source"`

	// A CredentialsSecretRef is a reference to a PostgreSQL connection secret
	// that contains the credentials that must be used to connect to the
	// provider. +optional
	ConnectionSecretRef *xpv1.SecretReference `json:"connectionSecretRef,omitempty"`
}

// A ProviderConfigStatus reflects the observed state of a ProviderConfig.
type ProviderConfigStatus struct {
	xpv1.ProviderConfigStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A ProviderConfig configures a Template provider.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentialsSecretRef.name",priority=1
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,sql}
type ProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderConfigSpec   `json:"spec"`
	Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderConfigList contains a list of ProviderConfig.
type ProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfig `json:"items"`
}

// +kubebuilder:object:root=true

// A ProviderConfigUsage indicates that a resource is using a ProviderConfig.
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="CONFIG-NAME",type="string",JSONPath=".providerConfigRef.name"
// +kubebuilder:printcolumn:name="RESOURCE-KIND",type="string",JSONPath=".resourceRef.kind"
// +kubebuilder:printcolumn:name="RESOURCE-NAME",type="string",JSONPath=".resourceRef.name"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,sql}
type ProviderConfigUsage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	xpv1.ProviderConfigUsage `json:",inline"`
}

// +kubebuilder:object:root=true

// ProviderConfigUsageList contains a list of ProviderConfigUsage
type ProviderConfigUsageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfigUsage `json:"items"`
}
