package v1

import (
	"fmt"
	"gin-dubbogo-consumer/constant"
	"gin-dubbogo-consumer/dto"
	"strconv"

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
	var req dto.ClusterUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err).Json()
		return
	}

	if err := k8sService.Update(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err).Json()
		return
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()
}

// @Router /delete/:id [delete]
func (b *BaseApi) Delete(c *gin.Context) {
	id, ok := c.Params.Get("id")
	fmt.Println("--------------------id1---------------------", id)
	fmt.Println(ok)
	uintID, err := strconv.ParseUint(id, 10, 64)
	fmt.Println("id")
	fmt.Println(id)
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err).Json()
		return
	}
	if err := k8sService.Delete(uint(uintID)); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err).Json()
		return
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()

}
func (b *BaseApi) GetInfo(c *gin.Context) {
	//var req request.AddrRuleUpdate
	fmt.Println("-------------------------id1------------------")
	id := c.Param("id")
	int, err := strconv.Atoi(id)
	fmt.Println("id")
	fmt.Println(id)
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err).Json()
		return
	}
	var dtoClusterInfo dto.ClusterInfo
	if dtoClusterInfo, err = k8sService.Get(uint(int)); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err).Json()
		return
	}
	NewResponse(c).Success(map[string]interface{}{"clustr": dtoClusterInfo}).Json()
}
func (b *BaseApi) List(c *gin.Context) {
	var dtoClusterInfo []dto.ClusterInfo
	//var err errors.New("")
	dtoClusterInfo, err := k8sService.List(dto.ClusterSearch{})
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err).Json()
		return
	}
	NewResponse(c).Success(map[string]interface{}{"cluster": dtoClusterInfo}).Json()

}
