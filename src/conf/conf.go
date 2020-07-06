package conf

import (
	"giligili/cache"
	"giligili/model"
	"giligili/util"
	"github.com/gin-gonic/gin"
	"io"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
