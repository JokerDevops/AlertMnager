/**
 * @Author Jinxiao Zhang · Az · JokerDevops
 * @Description //TODO
 * @Date 20:39 2023/4/15
 **/

package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type Hello struct {
}

func (p Hello) PostHello(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	c.String(200, buf.String())
}
