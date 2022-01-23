package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"oa-member/internal/biz"
	"oa-member/internal/data/ent"
)

type memberRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewMemberRepo(data *Data, logger log.Logger) biz.MemberRepo {
	return &memberRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *memberRepo) CreateMember(ctx context.Context, m *biz.Member) (*biz.Member, error) {
	mb, err := r.data.db.Member.Create().
		SetName(m.Name).
		SetMobile(m.Mobile).
		SetEmail(m.Email).
		Save(ctx)
	if err != nil {
		return &biz.Member{}, err
	}
	return DbMember2BizMember(mb), nil
}
func (r *memberRepo) UpdateMember(ctx context.Context, id int64, m *biz.Member) (member *biz.Member, err error) {
	var md *ent.Member
	md, err = r.data.db.Member.Get(ctx, int(id))
	if err != nil {
		return
	}
	save, err := md.Update().
		SetName(m.Name).
		SetMobile(m.Mobile).
		SetEmail(m.Email).
		Save(ctx)
	if err != nil {
		return
	}
	member = DbMember2BizMember(save)
	return
}
func (r *memberRepo) GetMember(ctx context.Context, id int64) (*biz.Member, error) {
	member, err := r.data.db.Member.Get(ctx, int(id))
	if err != nil {
		return nil, err
	}
	return DbMember2BizMember(member), nil
}
func (r *memberRepo) ListMember(ctx context.Context) ([]*biz.Member, error) {
	members, err := r.data.db.Member.Query().
		Limit(10).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Member, 0)
	for _, member := range members {
		rv = append(rv, DbMember2BizMember(member))
	}
	return rv, nil
}
func (r *memberRepo) DeleteMember(ctx context.Context, id int64) error {
	return nil
}

func DbMember2BizMember(member *ent.Member) *biz.Member {
	return &biz.Member{Id: int64(member.ID),
		Name:   member.Name,
		Mobile: member.Mobile,
		Email:  member.Email,
	}
}
