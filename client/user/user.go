package user

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
	"user-rpc/pb/user"
)

func client(ctx context.Context, hand func(c user.UserClient) error) error {
	//建立一个GRPC的链接
	conn, err := grpc.NewClient("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	//控制GRPC接口超时
	_, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

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
