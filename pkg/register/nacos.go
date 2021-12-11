package register

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type INacosServer interface {
	GetAddress() string
	GetPort() uint64
}

type iNacosClient interface {
	GetNamespace() string
	GetTimeout() uint64
}

func NacosRegister(
	nacosServers []INacosServer,
	nacosClient iNacosClient,
) (registry.Registrar, error) {
	sc := make([]constant.ServerConfig, len(nacosServers))
	for i, n := range nacosServers {
		sc[i] = *constant.NewServerConfig(n.GetAddress(), n.GetPort())
	}

	cc := constant.ClientConfig{
		NamespaceId: nacosClient.GetNamespace(),
		TimeoutMs:   nacosClient.GetTimeout(),
	}

	client, err := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	})
	if err != nil {
		return nil, err
	}

	return nacos.New(client), nil
}
