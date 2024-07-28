package api

import (
	"gitlab.ozon.dev/berkinv/homework/internal/models"
	"gitlab.ozon.dev/berkinv/homework/pkg/api/proto/pvz/v1/pvz/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Converter interface {
	PackageToDomain(req *pvz.AddPackageRequest) models.PackageUnit
	ChangePackageToDomain(req *pvz.ChangePackageRequest) models.ChangePackage
	ReceiveOrderDeliverToDomain(req *pvz.RecieveDeliverRequest) models.ReceiveOrderDeliver
	RefundOrderDeliverToDomain(req *pvz.RefundDeliverRequest) models.RefundOrderDeliver
	ReceiveOrderUserToDomain(req *pvz.ReceiveOrderUserRequest) []models.ReceiveOrderUser
	OrderListToDomain(req *pvz.OrderListRequest) models.OrderList
	RefundUserToDomain(req *pvz.RefundUserRequest) models.RefundUser
	OrderListResponseToDomain(list []models.DataUnit) *pvz.OrderListResponse
	RefundListResponseToDomain(list []models.DataUnit, req *pvz.RefundListRequest) *pvz.RefundListResponse
}

func PackageToDomain(req *pvz.AddPackageRequest) models.PackageUnit {
	return models.PackageUnit{
		PackageName: req.GetPackageName(),
		PackageCost: req.GetPackageCost(),
		LowerMass:   req.GetLowerMass(),
		UpperMass:   req.GetUpperMass(),
	}
}

func ChangePackageToDomain(req *pvz.ChangePackageRequest) models.ChangePackage {
	return models.ChangePackage{
		IdOrder:   req.GetIdOrder(),
		IdPackage: req.GetIdPackage(),
	}
}

func ReceiveOrderDeliverToDomain(req *pvz.RecieveDeliverRequest) models.ReceiveOrderDeliver {
	return models.ReceiveOrderDeliver{
		IdOrder:   req.GetIdOrder(),
		IdUser:    req.GetIdUser(),
		IdPackage: req.GetIdPackage(),
		DeadLine:  int(req.GetDeadline()),
		Mass:      req.GetMass(),
	}
}

func RefundOrderDeliverToDomain(req *pvz.RefundDeliverRequest) models.RefundOrderDeliver {

	return models.RefundOrderDeliver{
		IdOrder: req.GetIdOrder(),
	}
}

func ReceiveOrderUserToDomain(req *pvz.ReceiveOrderUserRequest) []models.ReceiveOrderUser {
	temp := req.GetIdOrder()
	ans := make([]models.ReceiveOrderUser, len(temp))
	for i, v := range temp {
		ans[i] = models.ReceiveOrderUser{IdOrder: v}
	}
	return ans
}

func OrderListToDomain(req *pvz.OrderListRequest) models.OrderList {
	return models.OrderList{
		IdUser: req.GetIdUser(),
	}
}

func RefundUserToDomain(req *pvz.RefundUserRequest) models.RefundUser {
	return models.RefundUser{
		IdOrder: req.GetIdOrder(),
		IdUser:  req.GetIdUser(),
	}
}

func OrderListResponseToDomain(list []models.DataUnit) *pvz.OrderListResponse {
	ans := make([]*pvz.DataUnit, 0, len(list))
	for _, v := range list {
		ans = append(ans, &pvz.DataUnit{
			IdUser:        v.IdUser,
			IdOrder:       v.IdOrder,
			IdPackage:     v.IdPackage,
			ReceivedDate:  timestamppb.New(v.ReceivedDate),
			DeliveredDate: timestamppb.New(v.DeliveredDate),
			RefundDate:    timestamppb.New(v.RefundDate),
			Mass:          v.Mass,
		})
	}
	return &pvz.OrderListResponse{
		Data: ans,
	}
}

func RefundListResponseToDomain(list []models.DataUnit, req *pvz.RefundListRequest) *pvz.RefundListResponse {
	ans := make([]*pvz.DataUnit, 0, min(len(list), int(req.GetNum())))
	for _, v := range list {
		ans = append(ans, &pvz.DataUnit{
			IdUser:        v.IdUser,
			IdOrder:       v.IdOrder,
			IdPackage:     v.IdPackage,
			ReceivedDate:  timestamppb.New(v.ReceivedDate),
			DeliveredDate: timestamppb.New(v.DeliveredDate),
			RefundDate:    timestamppb.New(v.RefundDate),
			Mass:          v.Mass,
		})
	}
	return &pvz.RefundListResponse{
		Data: ans,
	}
}
