package biz

import (
	"github.com/gogoclouds/gogo/_examples/internal/app/admin/dao"
	"github.com/gogoclouds/gogo/_examples/internal/app/admin/model"
	"github.com/gogoclouds/gogo/g"
	"github.com/gogoclouds/gogo/web/r"
)

type sysUserServiceImpl struct{}

func (svc *sysUserServiceImpl) PageList(query model.PageQuery) (*r.PageResp[model.SysUser], *g.Error) {
	list, err := dao.SysUser.PageList(query)
	return list, err
}

func (svc *sysUserServiceImpl) Create(user model.SysUser) *g.Error {
	//TODO implement me
	panic("implement me")
}

func (svc *sysUserServiceImpl) Update(user model.SysUser) *g.Error {
	//TODO implement me
	panic("implement me")
}

func (svc *sysUserServiceImpl) Delete(id int) *g.Error {
	//TODO implement me
	panic("implement me")
}