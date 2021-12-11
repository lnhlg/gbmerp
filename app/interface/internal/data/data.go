package data

import (
	"context"
	"github.com/go-redis/redis/extra/redisotel"
	"time"

	dpv1 "gbmerp/api/department/v1"
	duv1 "gbmerp/api/duty/v1"
	stv1 "gbmerp/api/staff/v1"
	usv1 "gbmerp/api/user/v1"
	"gbmerp/app/interface/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewUserServiceClient,
	NewStaffServiceClient,
	NewDutyServiceClient,
	NewDeptServiceClient,
	NewUserRepo,
)

type Data struct {
	usc usv1.UserClient
	stc stv1.StaffClient
	duc duv1.DutyClient
	dpc dpv1.DepartmentClient
	rdb *redis.Client
}

func NewData(
	conf *conf.Data,
	usc usv1.UserClient,
	stc stv1.StaffClient,
	duc duv1.DutyClient,
	dpc dpv1.DepartmentClient,
) *Data {
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})

	return &Data{
		usc: usc,
		stc: stc,
		duc: duc,
		dpc: dpc,
		rdb: rdb,
	}
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	sc := make([]constant.ServerConfig, len(conf.GetNacosServers()))
	for i, n := range conf.GetNacosServers() {
		sc[i] = *constant.NewServerConfig(n.GetAddress(), n.GetPort())
	}

	cc := constant.ClientConfig{
		NamespaceId:         conf.GetNacosClient().GetNamespace(),
		TimeoutMs:           conf.GetNacosClient().GetTimeout(),
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
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

func NewUserServiceClient(r registry.Discovery) usv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///gbmerp.user.grpc"),
		grpc.WithTimeout(time.Second*5),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}

	return usv1.NewUserClient(conn)
}

func NewStaffServiceClient(r registry.Discovery) stv1.StaffClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///gbmerp.staff.grpc"),
		grpc.WithTimeout(time.Second*5),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}

	return stv1.NewStaffClient(conn)
}

func NewDutyServiceClient(r registry.Discovery) duv1.DutyClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///gbmerp.duty.grpc"),
		grpc.WithTimeout(time.Second*5),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}

	return duv1.NewDutyClient(conn)
}

func NewDeptServiceClient(r registry.Discovery) dpv1.DepartmentClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///gbmerp.department.grpc"),
		grpc.WithTimeout(time.Second*5),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}

	return dpv1.NewDepartmentClient(conn)
}
