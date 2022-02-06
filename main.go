package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/2559065/user/domain/repository"
	service2 "github.com/2559065/user/domain/service"
	"github.com/micro/go-micro/v2"
	"github.com/jinzhu/gorm"
	"github.com/2559065/user/handler"
	pb "github.com/2559065/user/proto/user"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	// 初始化服务
	srv.Init()

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
