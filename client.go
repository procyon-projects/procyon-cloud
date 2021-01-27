package cloud

type ServiceInstance interface {
	GetInstanceId() string
	GetServiceId() string
	GetHost() string
	GetPort() int
	IsSecure() bool
	GetMetadata() map[string]string
	GetScheme() string
}

type DiscoveryClient interface {
	Description() string
	GetServiceInstances(serviceId string) []ServiceInstance
	GetServices() []string
}
