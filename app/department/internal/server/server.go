package server

import (
	"gbmerp/app/department/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	sc := make([]constant.ServerConfig, len(conf.GetNacosServers()))
	for i, n := range conf.GetNacosServers() {
		sc[i] = *constant.NewServerConfig(n.GetAddress(), n.GetPort())
	}

	cc := constant.ClientConfig{
		NamespaceId: conf.GetNacosClient().GetNamespace(),
		TimeoutMs:   conf.GetNacosClient().GetTimeout(),
	}

	client, err := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	})
	if err != nil {
		panic(err)
	}

	return nacos.New(client)
}
