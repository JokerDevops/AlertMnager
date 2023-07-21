/**
 * @Author Jinxiao Zhang · Az · JokerDevops
 * @Description //TODO
 * @Date 20:50 2023/4/15
 **/

package router

import (
	"AlertManager/controller"
	"AlertManager/pkg/middleWare/log"
	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.Default()
	r.Use(log.LogMiddleWare())
	r.POST("/", controller.Hello{}.PostHello)
	r.POST("/promethusAlert", controller.PrometheusController{}.PostPrometheusAlert)
	r.Run(":8080")
}
