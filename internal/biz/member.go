package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Member struct {
	Id     int64
	Name   string
	Mobile string
	Email  string
}

type MemberRepo interface {
	CreateMember(ctx context.Context, m *Member) (*Member, error)
	UpdateMember(ctx context.Context, id int64, m *Member) (*Member, error)
	GetMember(ctx context.Context, id int64) (*Member, error)
	ListMember(ctx context.Context) ([]*Member, error)
	DeleteMember(ctx context.Context, id int64) error
}

type MemberUserCase struct {
	repo      MemberRepo
	logHelper *log.Helper
}

func NewMemberUserCase(repo MemberRepo, logger log.Logger) *MemberUserCase {
	return &MemberUserCase{repo: repo, logHelper: log.NewHelper(logger)}
}

func (us *MemberUserCase) Create(ctx context.Context, m *Member) (*Member, error) {
	return us.repo.CreateMember(ctx, m)
}

func (us *MemberUserCase) Get(ctx context.Context, id int64) (ms *Member, err error) {
	ms, err = us.repo.GetMember(ctx, id)
	return
}

func (us *MemberUserCase) Update(ctx context.Context, m *Member) (ms *Member, err error) {
	ms, err = us.repo.UpdateMember(ctx, m.Id, m)
	return
}

func (us *MemberUserCase) List(ctx context.Context) (ms []*Member, err error) {
	ms, err = us.repo.ListMember(ctx)
	if err != nil {
		return
	}
	return
}
