/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RequestParameterObservation struct {
}

type RequestParameterParameters struct {

	// +kubebuilder:validation:Required
	RequestParameterKey *string `json:"requestParameterKey" tf:"request_parameter_key,omitempty"`

	// +kubebuilder:validation:Required
	Required *bool `json:"required" tf:"required,omitempty"`
}

type RouteObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouteParameters struct {

	// +crossplane:generate:reference:type=API
	// +crossplane:generate:reference:refFieldName=ApiIdRef
	// +crossplane:generate:reference:selectorFieldName=ApiIdSelector
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// +kubebuilder:validation:Optional
	APIKeyRequired *bool `json:"apiKeyRequired,omitempty" tf:"api_key_required,omitempty"`

	// +kubebuilder:validation:Optional
	ApiIdRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ApiIdSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AuthorizationScopes []*string `json:"authorizationScopes,omitempty" tf:"authorization_scopes,omitempty"`

	// +kubebuilder:validation:Optional
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// +kubebuilder:validation:Optional
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// +kubebuilder:validation:Optional
	ModelSelectionExpression *string `json:"modelSelectionExpression,omitempty" tf:"model_selection_expression,omitempty"`

	// +kubebuilder:validation:Optional
	OperationName *string `json:"operationName,omitempty" tf:"operation_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RequestModels map[string]*string `json:"requestModels,omitempty" tf:"request_models,omitempty"`

	// +kubebuilder:validation:Optional
	RequestParameter []RequestParameterParameters `json:"requestParameter,omitempty" tf:"request_parameter,omitempty"`

	// +kubebuilder:validation:Required
	RouteKey *string `json:"routeKey" tf:"route_key,omitempty"`

	// +kubebuilder:validation:Optional
	RouteResponseSelectionExpression *string `json:"routeResponseSelectionExpression,omitempty" tf:"route_response_selection_expression,omitempty"`

	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteParameters `json:"forProvider"`
}

// RouteStatus defines the observed state of Route.
type RouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route is the Schema for the Routes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteSpec   `json:"spec"`
	Status            RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Repository type metadata.
var (
	Route_Kind             = "Route"
	Route_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Route_Kind}.String()
	Route_KindAPIVersion   = Route_Kind + "." + CRDGroupVersion.String()
	Route_GroupVersionKind = CRDGroupVersion.WithKind(Route_Kind)
)

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
