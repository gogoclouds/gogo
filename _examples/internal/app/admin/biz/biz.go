package biz

import (
	"github.com/gogoclouds/gogo/_examples/internal/app/admin/model"
	"github.com/gogoclouds/gogo/g"
)

var SysUser g.IBase[model.SysUser, model.PageQuery] = new(sysUserServiceImpl)