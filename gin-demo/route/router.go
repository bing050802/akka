package route

import (
	"github.com/akkagao/akka/gin-demo/handler"
	"github.com/gin-gonic/gin"
)

/**
初始化路由
*/
func RouterInit(r *gin.Engine) {
	// 设置路由组
	testGroup := r.Group("/test")
	{
		/**
		获取restFul 参数
		curl GET http://localhost:8080/test/getuser/1345
		*/
		testGroup.GET("/getuser/:id", handler.GetUserById)

		/**
		获取？后面的参数
		curl GET http://localhost:8080/test/queryStrParameter?firstname=123&lastname=ggg
		*/
		testGroup.GET("/queryStrParameter", handler.QueryStrParameter)

		/**
		获取form表单提交参数
		curl -X POST http://localhost:8080/test/form_post -F name=gggg -F password=aaa
		*/
		testGroup.POST("/form_post", handler.FormPostParameter)

		/**
		上传文件
		  curl -X POST http://localhost:8080/test/uploadfile \
		  -F "file=@/Users/crazywolf/myapp/canal/canal.deployer-1.0.25.tar.gz" \
		  -H "Content-Type: multipart/form-data"
		*/
		testGroup.POST("/uploadfile", handler.Uploadfile)
	}

}
