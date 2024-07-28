package api

import (
	"context"

	"gitlab.ozon.dev/berkinv/homework/internal/module"
	"gitlab.ozon.dev/berkinv/homework/pkg/api/proto/pvz/v1/pvz/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PvzService struct {
	module.Module
	PvzServ pvz.PvzServer
	Converter
}
type Implementation interface {
	AddPackage(_ context.Context, req *pvz.AddPackageRequest) (*emptypb.Empty, error)
	ReceiveOrderDeliver(_ context.Context, req *pvz.RecieveDeliverRequest) (*emptypb.Empty, error)
	RefundDeliver(_ context.Context, req *pvz.RefundDeliverRequest) (*emptypb.Empty, error)
	ReceiveOrderUser(_ context.Context, req *pvz.ReceiveOrderUserRequest) (*emptypb.Empty, error)
	OrderList(_ context.Context, req *pvz.OrderListRequest) (*pvz.OrderListResponse, error)
	RefundUser(_ context.Context, req *pvz.RefundUserRequest) (*emptypb.Empty, error)
	ChangePackage(_ context.Context, req *pvz.ChangePackageRequest) (*emptypb.Empty, error)
}

func (pvz *PvzService) AddPackage(_ context.Context, req *pvz.AddPackageRequest) (*emptypb.Empty, error) {
	if errValidate := req.ValidateAll(); errValidate != nil {
		return nil, status.Error(codes.InvalidArgument, errValidate.Error())
	}
	pack := pvz.PackageToDomain(req)
	err := pvz.Module.AddPackage(pack)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (pvz *PvzService) ReceiveOrderDeliver(_ context.Context, req *pvz.RecieveDeliverRequest) (*emptypb.Empty, error) {
	err := pvz.Module.ReceiveOrderDeliver(pvz.ReceiveOrderDeliverToDomain(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (pvz *PvzService) RefundDeliver(_ context.Context, req *pvz.RefundDeliverRequest) (*emptypb.Empty, error) {
	err := pvz.Module.RefundDeliver(pvz.RefundOrderDeliverToDomain(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (pvz *PvzService) ReceiveOrderUser(_ context.Context, req *pvz.ReceiveOrderUserRequest) (*emptypb.Empty, error) {
	err := pvz.Module.ReceiveOrderUser(pvz.ReceiveOrderUserToDomain(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (pvz *PvzService) OrderList(_ context.Context, req *pvz.OrderListRequest) (*pvz.OrderListResponse, error) {
	temp, err := pvz.Module.OrderList(pvz.OrderListToDomain(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	ans := pvz.OrderListResponseToDomain(temp)
	return ans, nil
}
func (pvz *PvzService) RefundUser(_ context.Context, req *pvz.RefundUserRequest) (*emptypb.Empty, error) {
	err := pvz.Module.RefundUser(pvz.RefundUserToDomain(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (pvz *PvzService) RefundList(_ context.Context, req *pvz.RefundListRequest) (*pvz.RefundListResponse, error) {
	temp, err := pvz.Module.RefundList()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return pvz.RefundListResponseToDomain(temp, req), nil
}
func (pvz *PvzService) ChangePackage(_ context.Context, req *pvz.ChangePackageRequest) (*emptypb.Empty, error) {
	err := pvz.Module.ChangePackage(pvz.ChangePackageToDomain(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
