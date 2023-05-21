package service

import (
	"context"

	pb "comment-main/api/accountService/v1"
)

type AccountServiceService struct {
	pb.UnimplementedAccountServiceServer
}

func NewAccountServiceService() *AccountServiceService {
	return &AccountServiceService{}
}

func (s *AccountServiceService) CreateAccountService(ctx context.Context, req *pb.CreateAccountServiceRequest) (*pb.CreateAccountServiceReply, error) {
	return &pb.CreateAccountServiceReply{}, nil
}
func (s *AccountServiceService) UpdateAccountService(ctx context.Context, req *pb.UpdateAccountServiceRequest) (*pb.UpdateAccountServiceReply, error) {
	return &pb.UpdateAccountServiceReply{}, nil
}
func (s *AccountServiceService) DeleteAccountService(ctx context.Context, req *pb.DeleteAccountServiceRequest) (*pb.DeleteAccountServiceReply, error) {
	return &pb.DeleteAccountServiceReply{}, nil
}
func (s *AccountServiceService) GetAccountService(ctx context.Context, req *pb.GetAccountServiceRequest) (*pb.GetAccountServiceReply, error) {
	return &pb.GetAccountServiceReply{}, nil
}
func (s *AccountServiceService) ListAccountService(ctx context.Context, req *pb.ListAccountServiceRequest) (*pb.ListAccountServiceReply, error) {
	return &pb.ListAccountServiceReply{}, nil
}
