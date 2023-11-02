package v1

import (
	"fmt"
	"gin-dubbogo-consumer/constant"
	"gin-dubbogo-consumer/dto"
	//_ "gin-dubbogo-consumer/dto"
	//"github.com/1Panel-dev/1Panel/backend/app/api/v1/helper"

	//"gin-dubbogo-consumer/dto/request"
	"github.com/gin-gonic/gin"
)

func (b *BaseApi) Create(c *gin.Context) {
	var req dto.ClusterCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err).Json()
		return
	}

	if err := k8sService.Create(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err).Json()
		return
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()

}
func (b *BaseApi) Update(c *gin.Context) {
	//var req request.AddrRuleUpdate
	fmt.Println("hello world")
	/*if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/

}
func (b *BaseApi) Get(c *gin.Context) {
	//var req request.AddrRuleUpdate
	fmt.Println("hello world")
	/*if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/

}
func (b *BaseApi) List(c *gin.Context) {
	//var req request.AddrRuleUpdate
	fmt.Println("hello world")
	/*if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/

}
