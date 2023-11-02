package v1

import (
	"errors"
	"fmt"
	"gin-dubbogo-consumer/constant"
	dto "gin-dubbogo-consumer/dto"
	"gin-dubbogo-consumer/global"
	"gin-dubbogo-consumer/service"
	"github.com/gin-gonic/gin"
	"log"
)

var imageRepoService = service.NewIImageRepoService()
var dockerService = service.NewIDockerService()

func (b *BaseApi) ListContainer(c *gin.Context) {
	list, err := service.NewIDockerService().List()
	if err != nil {
		NewResponse(c).Fail(constant.CodeErrInternalServer, constant.ErrTypeInternalServer)
		//return
	}
	NewResponse(c).Success(map[string]interface{}{"list": list}).Json()
	//helper.SuccessWithData(c, list)
}

func (this *BaseApi) SearchContainer(c *gin.Context) {
	//return func(c *gin.Context) {
	//defer this.Base.Catch(NewResponse(c))
	var req dto.PageContainer
	fmt.Println("1111111111111111111")
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/
	total, list, err := service.NewIDockerService().Page(req)

	fmt.Println(total)
	fmt.Println(list)
	fmt.Println(err)
	if err != nil {
		//NewResponse(c).error(  constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		//return
	}

	NewResponse(c).Success(dto.PageResult{
		Rows:  list,
		Total: total,
	}).Json()
	//helper.SuccessWithData(c, dto.PageResult{
	//	Items: list,
	//	Total: total,
	//})

}

//	func (h *BaseApi) ContainerWsSsh(c *gin.Context) {
//		//return func(c *gin.Context) {
//		wsConn, err := wsSsh.UpGrader.Upgrade(c.Writer, c.Request, nil)
//		if err != nil {
//			global.LOG.Errorf("gin context http handler failed, err: %v", err)
//			return
//		}
//		defer wsConn.Close()
//
//		containerID := c.Query("containerid")
//		command := c.Query("command")
//		user := c.Query("user")
//		if len(command) == 0 || len(containerID) == 0 {
//			if wsSsh.WshandleError(wsConn, errors.New("error param of command or containerID")) {
//				return
//			}
//		}
//		cols, err := strconv.Atoi(c.DefaultQuery("cols", "80"))
//		if wsSsh.WshandleError(wsConn, errors.WithMessage(err, "invalid param cols in request")) {
//			return
//		}
//		rows, err := strconv.Atoi(c.DefaultQuery("rows", "40"))
//		if wsSsh.WshandleError(wsConn, errors.WithMessage(err, "invalid param rows in request")) {
//			return
//		}
//
//		cmds := []string{"exec", containerID, command}
//		if len(user) != 0 {
//			cmds = []string{"exec", "-u", user, containerID, command}
//		}
//		if cmd.CheckIllegal(user, containerID, command) {
//			if wsSsh.WshandleError(wsConn, errors.New("  The command contains illegal characters.")) {
//				return
//			}
//		}
//		stdout, err := cmd.ExecWithCheck("docker", cmds...)
//		if wsSsh.WshandleError(wsConn, errors.WithMessage(err, stdout)) {
//			return
//		}
//
//		commands := fmt.Sprintf("docker exec -it %s %s", containerID, command)
//		if len(user) != 0 {
//			commands = fmt.Sprintf("docker exec -it -u %s %s %s", user, containerID, command)
//		}
//		pidMap := wsSsh.LoadMapFromDockerTop(containerID)
//		slave, err := terminal.NewCommand(commands)
//		if wsSsh.WshandleError(wsConn, err) {
//			return
//		}
//		defer wsSsh.KillBash(containerID, command, pidMap)
//		defer slave.Close()
//
//		tty, err := terminal.NewLocalWsSession(cols, rows, wsConn, slave)
//		if wsSsh.WshandleError(wsConn, err) {
//			return
//		}
//
//		quitChan := make(chan bool, 3)
//		tty.Start(quitChan)
//		go slave.Wait(quitChan)
//
//		<-quitChan
//
//		global.LOG.Info("websocket finished")
//		if wsSsh.WshandleError(wsConn, err) {
//			return
//		}
//
// }
func (b *BaseApi) LoadContainerLog(c *gin.Context) {
	var req dto.OperationWithNameAndType
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	content := service.NewIDockerService().LoadContainerLogs(req)
	NewResponse(c).Success(map[string]interface{}{"page": content}).Json()
}
func (b *BaseApi) ContainerLogs(c *gin.Context) {
	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalln("gin context http handler failed, err: %v", err)
		return
	}
	defer wsConn.Close()

	container := c.Query("container")
	since := c.Query("since")
	follow := c.Query("follow") == "true"
	tail := c.Query("tail")

	if err := service.NewIDockerService().ContainerLogs(wsConn, container, since, tail, follow); err != nil {
		_ = wsConn.WriteMessage(1, []byte(err.Error()))
		return
	}
}

/* 容器监控*/
func (h *BaseApi) ContainerStats(c *gin.Context) {
	containerID, ok := c.Params.Get("id")
	if !ok {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, errors.New("error container id in path"))
		return
	}

	result, err := service.NewIDockerService().ContainerStats(containerID)
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(result).Json()
	//helper.SuccessWithData(c, result)
}

func (h *BaseApi) ContainerCreate(c *gin.Context) {
	//return func(c *gin.Context) {
	var req dto.ContainerOperate
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*	if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/
	if err := service.NewIDockerService().ContainerCreate(req); err != nil {
		log.Fatalln(err)
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
	}
	log.Println("-------------------------")
	NewResponse(c).Success(nil).Json()
}

func (h *BaseApi) ContainerUpgrade(c *gin.Context) {
	//return func(c *gin.Context) {

	var req dto.ContainerUpgrade
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/
	if err := service.NewIDockerService().ContainerUpgrade(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()
}

// }
func (h *BaseApi) ContainerPrune(c *gin.Context) {
	//return func(c *gin.Context) {
	var req dto.ContainerPrune
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)

	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)

	}*/
	report, err := service.NewIDockerService().Prune(req)
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)

	}
	//helper.SuccessWithData(c, report)
	NewResponse(c).Success(map[string]interface{}{"report": report}).Json()
}

// }
func (h *BaseApi) CleanContainerLog(c *gin.Context) {
	//return func(c *gin.Context) {
	var req dto.OperationWithName
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/
	if err := service.NewIDockerService().ContainerLogClean(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()
}

// }
func (h *BaseApi) ContainerOperation(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) ContainerOperation(c *gin.Context) {
	var req dto.ContainerOperation
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*	if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
	}*/
	if err := service.NewIDockerService().ContainerOperation(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()
}

//}

func (h *BaseApi) Inspect(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) Inspect(c *gin.Context) {
	var req dto.InspectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/

	result, err := service.NewIDockerService().Inspect(req)
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(result).Json()
}

// }
func (h *BaseApi) ContainerUpdate(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) ContainerUpdate(c *gin.Context) {
	var req dto.ContainerOperate
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/
	if err := service.NewIDockerService().ContainerUpdate(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()
}

// }
func (h *BaseApi) ContainerInfo(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) ContainerInfo(c *gin.Context) {
	var req dto.OperationWithName
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/
	data, err := service.NewIDockerService().ContainerInfo(req)
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(data).Json()

}

func (h *BaseApi) LoadResouceLimit(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) LoadResouceLimit(c *gin.Context) {
	data, err := service.NewIDockerService().LoadResouceLimit()
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(data).Json()

}
func (h *BaseApi) pconfigContainerLogs(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) ContainerLogs(c *gin.Context) {
	/*wsConn, err := wsSsh.UpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		//global.LOG.Errorf("gin context http handler failed, err: %v", err)
		return
	}
	defer wsConn.Close()

	container := c.Query("container")
	since := c.Query("since")
	follow := c.Query("follow") == "true"
	tail := c.Query("tail")

	if err := service.NewIDockerService().ContainerLogs(wsConn, container, since, tail, follow); err != nil {
		_ = wsConn.WriteMessage(1, []byte(err.Error()))
		return
	}*/
}

// }
func (h *BaseApi) ListImage(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) ListImage(c *gin.Context) {
	list, err := service.NewIImageService().List()
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(list).Json()

}

// }
func (h *BaseApi) SearchImage(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) SearchImage(c *gin.Context) {
	var req dto.SearchWithPage
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/

	total, list, err := service.NewIImageService().Page(req)
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(dto.PageResult{
		Rows:  list,
		Total: total,
	}).Json()

	/*helper.SuccessWithData(c, dto.PageResult{
		Items: list,
		Total: total,
	})*/
}

func (h *BaseApi) ImageBuild(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) ImageBuild(c *gin.Context) {
	var req dto.ImageBuild
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/

	log, err := service.NewIImageService().ImageBuild(req)
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}

	NewResponse(c).Success(log).Json()
}

func (h *BaseApi) ImagePull(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) ImagePull(c *gin.Context) {
	var req dto.ImagePull
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/

	logPath, err := service.NewIImageService().ImagePull(req)
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}

	NewResponse(c).Success(logPath).Json()

}
func (h *BaseApi) ImagePush(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) ImagePush(c *gin.Context) {
	var req dto.ImagePush
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/

	logPath, err := service.NewIImageService().ImagePush(req)
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}

	NewResponse(c).Success(logPath).Json()

}
func (h *BaseApi) ImageRemove(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) ImageRemove(c *gin.Context) {
	var req dto.BatchDelete
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/

	if err := service.NewIImageService().ImageRemove(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(nil).Json()

}
func (h *BaseApi) ImageSave(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) ImageSave(c *gin.Context) {
	var req dto.ImageSave
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err).Json()
		//return
	}
	/*	if err := global.VALID.Struct(req); err != nil {
			NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
			return
		}
	*/
	if err := service.NewIImageService().ImageSave(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err).Json()
		return
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()

}
func (h *BaseApi) ImageTag(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) ImageTag(c *gin.Context) {
	var req dto.ImageTag
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*	if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/

	if err := service.NewIImageService().ImageTag(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()
}

func (h *BaseApi) ImageLoad(c *gin.Context) {
	//return func(c *gin.Context) {
	//func (b *BaseApi) (c *gin.Context) {
	var req dto.ImageLoad
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/

	if err := service.NewIImageService().ImageLoad(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()

}

//

// @Tags Container Image-repo
// @Summary Page image repos
// @Description 获取镜像仓库列表分页
// @Accept json
// @Param request body dto.SearchWithPage true "request"
// @Produce json
// @Success 200 {object} dto.PageResult
// @Security ApiKeyAuth
// @Router /containers/repo/search [post]
func (b *BaseApi) SearchRepo(c *gin.Context) {
	var req dto.SearchWithPage
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	total, list, err := imageRepoService.Page(req)
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}

	NewResponse(c).Success(dto.PageResult{
		Rows:  list,
		Total: total,
	}).Json()
}

// @Tags Container Image-repo
// @Summary List image repos
// @Description 获取镜像仓库列表
// @Produce json
// @Success 200 {array} dto.ImageRepoOption
// @Security ApiKeyAuth
// @Router /containers/repo [get]
func (b *BaseApi) ListRepo(c *gin.Context) {
	list, err := imageRepoService.List()
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}

	NewResponse(c).Success(list).Json()
}

// @Tags Container Image-repo
// @Summary Load repo status
// @Description 获取 docker 仓库状态
// @Accept json
// @Param request body dto.OperateByID true "request"
// @Produce json
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/repo/status [get]
func (b *BaseApi) CheckRepoStatus(c *gin.Context) {
	var req dto.OperateByID
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := imageRepoService.Login(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(nil).Json()
}

// @Tags Container Image-repo
// @Summary Create image repo
// @Description 创建镜像仓库
// @Accept json
// @Param request body dto.ImageRepoDelete true "request"
// @Produce json
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/repo [post]
// @x-panel-log {"bodyKeys":["name"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"创建镜像仓库 [name]","formatEN":"create image repo [name]"}
func (b *BaseApi) CreateRepo(c *gin.Context) {
	var req dto.ImageRepoCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	/*if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}*/
	if err := imageRepoService.Create(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(nil).Json()
}

// @Tags Container Image-repo
// @Summary Delete image repo
// @Description 删除镜像仓库
// @Accept json
// @Param request body dto.ImageRepoDelete true "request"
// @Produce json
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/repo/del [post]
// @x-panel-log {"bodyKeys":["ids"],"paramKeys":[],"BeforeFunctions":[{"input_column":"id","input_value":"ids","isList":true,"db":"image_repos","output_column":"name","output_value":"names"}],"formatZH":"删除镜像仓库 [names]","formatEN":"delete image repo [names]"}
func (b *BaseApi) DeleteRepo(c *gin.Context) {
	var req dto.ImageRepoDelete
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	if err := imageRepoService.BatchDelete(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(nil).Json()
}

// @Tags Container Image-repo
// @Summary Update image repo
// @Description 更新镜像仓库
// @Accept json
// @Param request body dto.ImageRepoUpdate true "request"
// @Produce json
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/repo/update [post]
// @x-panel-log {"bodyKeys":["id"],"paramKeys":[],"BeforeFunctions":[{"input_column":"id","input_value":"id","isList":false,"db":"image_repos","output_column":"name","output_value":"name"}],"formatZH":"更新镜像仓库 [name]","formatEN":"update image repo information [name]"}
func (b *BaseApi) UpdateRepo(c *gin.Context) {
	var req dto.ImageRepoUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		NewResponse(c).error(constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	if err := imageRepoService.Update(req); err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(nil).Json()
}

// @Tags Container Volume
// @Summary List volumes
// @Description 获取容器存储卷列表
// @Accept json
// @Produce json
// @Success 200 {array} dto.Options
// @Security ApiKeyAuth
// @Router /containers/volume [get]
func (b *BaseApi) ListVolume(c *gin.Context) {
	list, err := dockerService.ListVolume()
	if err != nil {
		NewResponse(c).error(constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	NewResponse(c).Success(list).Json()
}
