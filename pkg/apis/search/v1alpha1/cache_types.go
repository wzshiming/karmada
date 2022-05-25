package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	policyv1alpha1 "github.com/karmada-io/karmada/pkg/apis/policy/v1alpha1"
)

const (
	// ResourceKindResourceRegistry is the name of the resource registry
	ResourceKindResourceRegistry = "ResourceRegistry"
	// ResourceSingularResourceRegistry is singular name of ResourceRegistry.
	ResourceSingularResourceRegistry = "resourceregistry"
	// ResourcePluralResourceRegistry is plural name of ResourceRegistry.
	ResourcePluralResourceRegistry = "resourceregistries"
	// ResourceNamespaceScopedResourceRegistry is the scope of the ResourceRegistry
	ResourceNamespaceScopedResourceRegistry = false
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ResourceRegistry defines a list of member cluster to be cached.
type ResourceRegistry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec represents the desired behavior of ResourceRegistry.
	Spec ResourceRegistrySpec `json:"spec,omitempty"`

	// Status represents the status of ResoruceRegistry.
	// +optional
	Status ResourceRegistryStatus `json:"status,omitempty"`
}

// ResourceRegistrySpec defines the desired state of ResourceRegistry.
type ResourceRegistrySpec struct {
	// TargetCluster is the cluster that the resource registry is targeting.
	// +required
	TargetCluster *policyv1alpha1.ClusterAffinity `json:"targetCluster"`

	// ResourceSelectors used to select resources.
	// +required
	ResourceSelectors []ResourceSelector `json:"resourceSelectors"`

	// StatusUpdatePeriodSeconds is the period to update the status of the resource.
	// default is 10s.
	// +optional
	StatusUpdatePeriodSeconds uint32 `json:"statusUpdatePeriodSeconds,omitempty"`
}

// ResourceSelector the resources will be selected.
type ResourceSelector struct {
	// APIVersion represents the API version of the target resources.
	// +required
	APIVersion string `json:"apiVersion"`

	// Kind represents the kind of the target resources.
	// +required
	Kind string `json:"kind"`

	// Namespace of the target resource.
	// Default is empty, which means all namespaces.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// ResourceRegistryStatus defines the observed state of ResourceRegistry
type ResourceRegistryStatus struct {
	// Conditions contain the different condition statuses.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:resource:scope="Cluster"
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ResourceRegistryList if a collection of ResourceRegistry.
type ResourceRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items holds a list of ResourceRegistry.
	Items []ResourceRegistry `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Search define a flag for resource search that do not have actual resources.
type Search struct {
	metav1.TypeMeta `json:",inline"`
}
