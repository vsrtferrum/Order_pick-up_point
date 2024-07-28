package main

import (
	"context"
	"log"

	"gitlab.ozon.dev/berkinv/homework/pkg/api/proto/pvz/v1/pvz/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	target = "localhost:63342"
)

func main() {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pvz.NewPvzClient(conn)
	ctx := context.Background()
	err = listOrder(ctx, client, 2)
	if err != nil {
		log.Println(err)
	}
}

func listOrder(ctx context.Context, client pvz.PvzClient, idUser uint32) error {
	resp, err := client.OrderList(ctx, &pvz.OrderListRequest{
		IdUser: idUser,
	})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	for _, v := range resp.GetData() {
		log.Println(v)
	}
	return nil
}

func addPackage(ctx context.Context, client pvz.PvzClient, packageName string, packageCost uint32, lower_mass uint32, upper_mass uint32) error {
	_, err := client.AddPackage(ctx, &pvz.AddPackageRequest{
		PackageName: packageName,
		PackageCost: packageCost,
		LowerMass:   lower_mass,
		UpperMass:   upper_mass,
	})
	if err != nil {
		stat := status.Code(err)
		if stat == codes.InvalidArgument {
			return nil
		}
		return err
	}
	return nil
}

func changePackage(ctx context.Context, client pvz.PvzClient, idOrder uint32, idPackage uint32) error {
	_, err := client.ChangePackage(ctx, &pvz.ChangePackageRequest{
		IdOrder:   idOrder,
		IdPackage: idPackage,
	})
	if err != nil {
		stat := status.Code(err)
		if stat == codes.InvalidArgument {
			return nil
		}
		return err
	}
	return nil
}

func refundUser(ctx context.Context, client pvz.PvzClient, idOrder uint32, idUser uint32) error {
	_, err := client.RefundUser(ctx, &pvz.RefundUserRequest{
		IdOrder: idOrder,
		IdUser:  idUser,
	})
	if err != nil {
		stat := status.Code(err)
		if stat == codes.InvalidArgument {
			return nil
		}
		return err
	}
	return nil
}

func receiveOrderUser(ctx context.Context, client pvz.PvzClient, idOrder []uint32) error {
	_, err := client.ReceiveOrderUser(ctx, &pvz.ReceiveOrderUserRequest{
		IdOrder: idOrder,
	})
	if err != nil {
		stat := status.Code(err)
		if stat == codes.InvalidArgument {
			return nil
		}
		return err
	}
	return nil
}

func receiveOrderDeliver(ctx context.Context, client pvz.PvzClient, idOrder uint32, idUser uint32, deadLine int32, mass uint32) error {
	_, err := client.ReceiveOrderDeliver(ctx, &pvz.RecieveDeliverRequest{
		IdOrder:  idOrder,
		IdUser:   idUser,
		Deadline: deadLine,
		Mass:     mass,
	})
	if err != nil {
		stat := status.Code(err)
		if stat == codes.InvalidArgument {
			return nil
		}
		return err
	}
	return nil
}

func refundDeliver(ctx context.Context, client pvz.PvzClient, idOrder uint32) error {
	_, err := client.RefundDeliver(ctx, &pvz.RefundDeliverRequest{
		IdOrder: idOrder,
	})
	if err != nil {
		stat := status.Code(err)
		if stat == codes.InvalidArgument {
			return nil
		}
		return err
	}
	return nil
}
func listRefund(ctx context.Context, client pvz.PvzClient, num uint32) error {
	_, err := client.RefundList(ctx, &pvz.RefundListRequest{
		Num: num,
	})
	if err != nil {
		stat := status.Code(err)
		if stat == codes.InvalidArgument {
			return nil
		}
		return err
	}
	return nil
}
