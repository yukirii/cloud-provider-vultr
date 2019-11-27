package vultr

import (
	"io"
	"os"

	vultr "github.com/JamesClonk/vultr/lib"
	cloudprovider "k8s.io/cloud-provider"
)

const (
	// ProviderName is name of CCM provider
	ProviderName string = "vultr"

	// ControllerName is name of CCM for logging
	ControllerName string = "vultr-cloud-controller-manager"
)

type cloud struct {
	vultrClient *vultr.Client

	instances *cloudprovider.Instances
}

func newCloud(configReader io.Reader) (cloudprovider.Interface, error) {
	apiKey := os.Getenv("VULTR_API_KEY")
	client := vultr.NewClient(apiKey, nil)

	return &cloud{
		vultrClient: client,
		instances:   newInstances(client),
	}, nil
}

func init() {
	cloudprovider.RegisterCloudProvider(ProviderName, func(config io.Reader) (cloudprovider.Interface, error) {
		return newCloud(config)
	})
}

// Initialize provides the cloud with a kubernetes client builder and may spawn goroutines
// to perform housekeeping or run custom controllers specific to the cloud provider.
// Any tasks started here should be cleaned up when the stop channel closes.
func (c *cloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
}

// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
func (c *cloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	// not supported
	return nil, false
}

// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
func (c *cloud) Instances() (cloudprovider.Instances, bool) {
	// not supported
	return nil, false
}

// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
func (c *cloud) Zones() (cloudprovider.Zones, bool) {
	// not supported
	return nil, false
}

// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
func (c *cloud) Clusters() (cloudprovider.Clusters, bool) {
	// not supported
	return nil, false
}

// Routes returns a routes interface along with whether the interface is supported.
func (c *cloud) Routes() (cloudprovider.Routes, bool) {
	// not supported
	return nil, false
}

// ProviderName returns the cloud provider ID.
func (c *cloud) ProviderName() string {
	return ProviderName
}

// HasClusterID returns true if a ClusterID is required and set
func (c *cloud) HasClusterID() bool {
	return true
}
