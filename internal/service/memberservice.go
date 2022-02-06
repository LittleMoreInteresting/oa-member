package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"oa-member/internal/biz"

	pb "oa-member/api/member/v1"
)

func NewMemberServiceService(member *biz.MemberUserCase, logger log.Logger) *MemberServiceService {
	return &MemberServiceService{member: member, log: log.NewHelper(logger)}
}

func (s *MemberServiceService) CreateMember(ctx context.Context, req *pb.CreateMemberRequest) (*pb.CreateMemberReply, error) {
	s.log.Infof("input data %v", req)
	errValidate := req.ValidateAll()
	if errValidate != nil {
		return nil, pb.ErrorParamsError("service : params error %s", errValidate)
	}
	m, err := s.member.Create(ctx, &biz.Member{
		Name:   req.Name,
		Mobile: req.Mobile,
		Email:  req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateMemberReply{Member: BizMember2Reply(m)}, err
}

func (s *MemberServiceService) UpdateMember(ctx context.Context, req *pb.UpdateMemberRequest) (*pb.UpdateMemberReply, error) {
	m, err := s.member.Update(ctx, &biz.Member{
		Id:     req.Id,
		Name:   req.Name,
		Mobile: req.Mobile,
		Email:  req.Email,
	})
	if err != nil {
		return nil, pb.ErrorDbActionError("service: db update error")
	}
	return &pb.UpdateMemberReply{Member: BizMember2Reply(m)}, err
}

func (s *MemberServiceService) DeleteMember(ctx context.Context, req *pb.DeleteMemberRequest) (*pb.DeleteMemberReply, error) {
	err := s.member.Delete(ctx, req.Id)
	if err != nil {
		return nil, pb.ErrorDbActionError("service: db delete error")
	}
	return &pb.DeleteMemberReply{}, nil
}

func (s *MemberServiceService) GetMember(ctx context.Context, req *pb.GetMemberRequest) (*pb.GetMemberReply, error) {
	member, err := s.member.Get(ctx, req.Id)
	if err != nil {
		return nil, pb.ErrorMemberNotFound("member not found for ID %v", req.Id)
	}
	return &pb.GetMemberReply{Member: BizMember2Reply(member)}, nil
}

func (s *MemberServiceService) ListMember(ctx context.Context, req *pb.ListMemberRequest) (*pb.ListMemberReply, error) {
	list, err := s.member.List(ctx)
	if err != nil {
		return nil, err
	}
	var result []*pb.Member
	for _, member := range list {
		result = append(result, BizMember2Reply(member))
	}

	return &pb.ListMemberReply{Results: result}, nil
}

func BizMember2Reply(m *biz.Member) *pb.Member {
	return &pb.Member{
		Id:     m.Id,
		Name:   m.Name,
		Mobile: m.Mobile,
		Email:  m.Email,
	}
}
