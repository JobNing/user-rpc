package user

import (
	"context"
	"fmt"
	"github.com/JobNing/user-rpc/pb/user"
	"google.golang.org/grpc"
	"time"

	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

func client(ctx context.Context, hand func(c user.UserClient) error) error {
	//addr, err := config.GetServiceInstance(consts.SERVICE_NAME)
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Println("通过nacos获取地址：", addr)

	svcCfg := fmt.Sprintf(`{"loadBalancingPolicy":"%s"}`, "round_robin")
	conn, err := grpc.Dial("consul://60.204.199.43:8500/user-rpc?wait=14s&tag=public", grpc.WithInsecure(), grpc.WithDefaultServiceConfig(svcCfg))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	time.Sleep(29 * time.Second)

	//target := "consul://60.204.199.43:8500/user-rpc" // schema:[//authority/]host[:port]/service[?query] 参考文档：https://github.com/grpc/grpc/blob/master/doc/naming.md
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()
	//
	//conn, err := grpc.NewClient(
	//	target,
	//	grpc.WithTransportCredentials(insecure.NewCredentials()))

	////建立一个GRPC的链接
	//conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	//控制GRPC接口超时
	//_, cancel = context.WithTimeout(ctx, time.Second)
	//defer cancel()

	//延迟关闭链接
	defer conn.Close()

	//获取user client
	c := user.NewUserClient(conn)

	//将client 传递给闭包函数并执行闭包内的内容
	return hand(c)
}

func GetUser(ctx context.Context, id int64) (info *user.UserInfo, err error) {
	//通过client获取user Client
	err = client(ctx, func(c user.UserClient) error {
		//拿到client传递过来的user Client,调用 GetUser 接口
		res, err := c.GetUser(ctx, &user.GetUserRequest{
			ID: id,
		})
		if err != nil {
			return err
		}

		//处理返回值，将 user.GetUserResponse 内的UserInfo拆解出来赋值并返回
		info = res.GetInfo()
		return nil
	})
	return
}

func CreateUser(ctx context.Context, in *user.UserInfo) (info *user.UserInfo, err error) {
	//通过client获取user Client
	err = client(ctx, func(c user.UserClient) error {
		//拿到client传递过来的user Client,调用 GetUser 接口
		res, err := c.CreateUser(ctx, &user.CreateUserRequest{
			In: in,
		})
		if err != nil {
			return err
		}

		//处理返回值，将 user.GetUserResponse 内的UserInfo拆解出来赋值并返回
		info = res.GetInfo()
		return nil
	})
	return
}
