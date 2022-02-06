package main

import (
	"fmt"
	"github.com/2559065/common"
	"github.com/2559065/user/domain/repository"
	service2 "github.com/2559065/user/domain/service"
	"github.com/2559065/user/handler"
	pb "github.com/2559065/user/proto/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

var (
	//qps = os.Getenv("QPS")
	QPS = 1000
)

func main() {
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"localhost:8500",
		}
	})
	// 3.jaeger 链路追踪
	t, io, err := common.NewTracer("go.micro.service.order", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 创建数据库连接
	db, err := gorm.Open("mysql", "root:root@/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	db.SingularTable(true)

	//只执行一次,数据表初始化
	//rp := repository.NewUserRepository(db)
	//rp.InitTable()

	// 创建服务实例
	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))

	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		//暴露的服务地址
		micro.Address(":9087"),
		//添加consul 注册中心
		micro.Registry(consul),
		//添加链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)
	// 初始化服务
	srv.Init()

	// Register handler
	err = pb.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
