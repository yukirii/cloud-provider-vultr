package vultr

import (
	"context"
	"fmt"
	"strconv"

	vultr "github.com/JamesClonk/vultr/lib"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	cloudprovider "k8s.io/cloud-provider"
	"sigs.k8s.io/cluster-api/controllers/noderefutil"
)

type instances struct {
	vultrClient *vultr.Client
}

func newInstances(client *vultr.Client) cloudprovider.Instances {
	return &instances{
		vultrClient: client,
	}
}

// NodeAddresses returns the addresses of the specified instance.
func (i *instances) NodeAddresses(ctx context.Context, name types.NodeName) ([]v1.NodeAddress, error) {
	server, err := findServerByName(i.vultrClient, string(name))
	if err != nil {
		return nil, err
	}

	return []v1.NodeAddress{
		v1.NodeAddress{Type: v1.NodeExternalIP, Address: server.MainIP},
	}, nil
}

// NodeAddressesByProviderID returns the addresses of the specified instance.
// The instance is specified using the providerID of the node. The
// ProviderID is a unique identifier of the node. This will not be called
// from the node whose nodeaddresses are being queried. i.e. local metadata
// services cannot be used in this method to obtain nodeaddresses
func (i *instances) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
	pid, err := noderefutil.NewProviderID(providerID)
	if err != nil {
		return nil, err
	}

	server, err := i.vultrClient.GetServer(pid.ID())
	if err != nil {
		return nil, err
	}

	return []v1.NodeAddress{
		v1.NodeAddress{Type: v1.NodeExternalIP, Address: server.MainIP},
	}, nil
}

// InstanceID returns the cloud provider ID of the node with the specified NodeName.
// Note that if the instance does not exist, we must return ("", cloudprovider.InstanceNotFound)
// cloudprovider.InstanceNotFound should NOT be returned for instances that exist but are stopped/sleeping
func (i *instances) InstanceID(ctx context.Context, nodeName types.NodeName) (string, error) {
	server, err := findServerByName(i.vultrClient, string(nodeName))
	if err != nil {
		return "", err
	}

	optional := ""
	segments := strconv.Itoa(server.RegionID)
	providerID := fmt.Sprintf("%s/%s/%s", optional, segments, server.ID)

	return providerID, nil
}

// InstanceType returns the type of the specified instance.
func (i *instances) InstanceType(ctx context.Context, name types.NodeName) (string, error) {
	server, err := findServerByName(i.vultrClient, string(name))
	if err != nil {
		return "", err
	}

	return strconv.Itoa(server.PlanID), nil
}

// InstanceTypeByProviderID returns the type of the specified instance.
func (i *instances) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
	pid, err := noderefutil.NewProviderID(providerID)
	if err != nil {
		return "", err
	}

	server, err := i.vultrClient.GetServer(pid.ID())
	if err != nil {
		return "", err
	}

	return strconv.Itoa(server.PlanID), nil
}

// AddSSHKeyToAllInstances adds an SSH public key as a legal identity for all instances
// expected format for the key is standard ssh-keygen format: <protocol> <blob>
func (i *instances) AddSSHKeyToAllInstances(ctx context.Context, user string, keyData []byte) error {
	return cloudprovider.NotImplemented
}

// CurrentNodeName returns the name of the node we are currently running on
// On most clouds (e.g. GCE) this is the hostname, so we provide the hostname
func (i *instances) CurrentNodeName(ctx context.Context, hostname string) (types.NodeName, error) {
	return types.NodeName(hostname), nil
}

// InstanceExistsByProviderID returns true if the instance for the given provider exists.
// If false is returned with no error, the instance will be immediately deleted by the cloud controller manager.
// This method should still return true for instances that exist but are stopped/sleeping.
func (i *instances) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
	pid, err := noderefutil.NewProviderID(providerID)
	if err != nil {
		return false, err
	}

	_, err = i.vultrClient.GetServer(pid.ID())
	if err != nil {
		// The server does not exists.
		if err.Error() == "Invalid server." {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// InstanceShutdownByProviderID returns true if the instance is shutdown in cloudprovider
func (i *instances) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
	pid, err := noderefutil.NewProviderID(providerID)
	if err != nil {
		return false, err
	}

	server, err := i.vultrClient.GetServer(pid.ID())
	if err != nil {
		return false, err
	}

	return server.PowerStatus == "running", nil
}

// findServerByName finds a Vultr Server instance by name.
func findServerByName(vultrClient *vultr.Client, name string) (*vultr.Server, error) {
	servers, err := vultrClient.GetServers()
	if err != nil {
		return nil, err
	}

	for _, server := range servers {
		if server.Name == name {
			return &server, nil
		}
	}

	return nil, cloudprovider.InstanceNotFound
}
