// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/karmada-io/karmada/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ResourceExploringWebhookConfigurations returns a ResourceExploringWebhookConfigurationInformer.
	ResourceExploringWebhookConfigurations() ResourceExploringWebhookConfigurationInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ResourceExploringWebhookConfigurations returns a ResourceExploringWebhookConfigurationInformer.
func (v *version) ResourceExploringWebhookConfigurations() ResourceExploringWebhookConfigurationInformer {
	return &resourceExploringWebhookConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
