package cloud

import "net/url"

type ServiceInstance interface {
	GetInstanceId() string
	GetServiceId() string
	GetURL() url.URL
	GetScheme() string
	GetHost() string
	GetPort() int
	IsSecure() bool
	GetMetadata() map[string]string
}

type DiscoveryClient interface {
	GetDescription() string
	GetServiceInstances(serviceId string) []ServiceInstance
	GetServices() []string
}
