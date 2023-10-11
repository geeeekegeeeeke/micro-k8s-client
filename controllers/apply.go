package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ApplyController struct {
	Base *BaseController
}

func (this *ApplyController) GetApply(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	//获取NODE
	fmt.Println("####### 获取node ######")
	/*nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, nds := range nodes.Items {
		fmt.Printf("NodeName: %s\n", nds.Name)
	}*/

	//获取 指定NODE 的详细信息
	fmt.Println("\n ####### node详细信息 ######")
	//nodeName := "k8s-master2"
	//nodeRel, err := clientset.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	//Response{}
	NewResponse(c).Success(map[string]interface{}{}).Json()
}
