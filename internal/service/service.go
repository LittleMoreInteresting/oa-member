package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "oa-member/api/member/v1"
	"oa-member/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMemberServiceService)

type MemberServiceService struct {
	pb.UnimplementedMemberServiceServer

	member *biz.MemberUserCase
	log    *log.Helper
}
