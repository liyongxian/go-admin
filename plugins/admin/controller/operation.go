package controller

import (
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/connections"
	"github.com/chenhg5/go-admin/context"
	"github.com/gin-gonic/gin/json"
)

func RecordOperationLog(ctx *context.Context) {
	if user, ok := ctx.UserValue["user"].(auth.User); ok {
		var input []byte
		form := ctx.Request.MultipartForm
		if form != nil {
			input, _ = json.Marshal((*form).Value)
		}

		connections.GetConnection().Exec("insert into goadmin_operation_log (user_id, path, method, ip, input) values (?, ?, ?, ?, ?)", user.ID, ctx.Path(),
			ctx.Method(), ctx.LocalIP(), string(input))
	}
}
