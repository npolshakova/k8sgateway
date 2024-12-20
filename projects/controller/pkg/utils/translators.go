package utils

import (
	"fmt"
	"strings"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Danger any label changes here likely requires more than a drop in replacement string.
// Any selectors likely still have to take the union of the old string name in order to facilitate orderly upgrades.

const (
	// ProxyTypeKey is the label key applied to Proxies generated by the k8sgateway translator
	ProxyTypeKey = "created_by"

	// GlooEdgeProxyValue is the label value for ProxyTypeKey applied to Proxy CRs
	// that have been generated from k8sgateway Gateway API resources
	GlooEdgeProxyValue = "gloo-gateway"

	// GatewayApiProxyValue is the label value for ProxyTypeKey applied to Proxy CRs
	// that have been generated from Kubernetes Gateway API resources
	GatewayApiProxyValue = "gloo-kube-gateway-api"

	// KnativeProxyValue is the label value applied to Proxies generated by the Gloo Knative translator
	KnativeProxyValue = "gloo-knative"

	// IngressProxyValue is the label value applied to Proxies generated by the Gloo Ingress translator
	IngressProxyValue = "gloo-ingress"

	// GatewayNamespaceKey is the label key applied to a Proxy CR
	// The value represents the namespace of the Gateway resource that generated it
	// This will only exist on Proxy CRs that have been generated from Kubernetes Gateway CRs
	// because those Proxies are always persisted in the writeNamespace, but we need to maintain
	// a reference to where the originating Gateway was defined
	GatewayNamespaceKey = "gateway_namespace"

	// ProxySyncId is an annotation used to associate a proxy translation with the status proxyReport based on the sync count
	ProxySyncId = "proxy_sync_id"
)

func GetTranslatorSelectorExpression(translators ...string) string {
	return fmt.Sprintf("%s in (%s)", ProxyTypeKey, strings.Join(translators, ", "))
}

func GetTranslatorValue(meta *core.Metadata) string {
	return meta.GetLabels()[ProxyTypeKey]
}
