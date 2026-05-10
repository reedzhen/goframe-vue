package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "goframe-vben/addons"
	"goframe-vben/internal/boot"
	"goframe-vben/internal/cmd"
	_ "goframe-vben/internal/library/gftenant"
	_ "goframe-vben/internal/logic"
	_ "goframe-vben/internal/packed"
)

func main() {
	ctx := gctx.GetInitCtx()
	boot.Init(ctx)
	cmd.Main.Run(ctx)
}
