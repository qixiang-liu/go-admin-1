package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
)

func (h *Handler) Operation(ctx *context.Context) {
	id := ctx.Query("__id")
	if !h.OperationHandler(config.Get().Url("/operation/"+id), ctx) {
		errMsg := "not found"
		if ctx.Headers(constant.PjaxHeader) == "" && ctx.Method() == "POST" {
			response.BadRequest(ctx, errMsg)
		} else {
			response.Alert(ctx, errMsg, errMsg, errMsg, h.conn)
		}
		return
	}
}
