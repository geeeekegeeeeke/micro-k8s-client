package controllers

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 *  控制层初始化文件
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package controllers
 * @description
 *
 * 说明:
 *
 *  1.封装Json响应对象
 *    //成功响应(成功数据)
 * 	 //NewResponse(c).Success(map[string]interface{}{"key": "hello"}).Json()
 *   //失败响应（异常码，异常信息)
 *	 NewResponse(c).Fail(1000, "Business exception ").Json()
 *   //支持xml,yaml
 *	 NewResponse(c).Fail(1000, "Business exception ").Xml()
 *	 NewResponse(c).Fail(1000, "Business exception").Yaml()
 *
 */
import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/staging/src/k8s.io/client-go/dynamic"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Response struct {
	c        *gin.Context
	statusOk int
	Code     int                    `json:"code"`
	Message  string                 `json:"message"`
	Data     map[string]interface{} `json:"data"`
}

func NewResponse(c *gin.Context) *Response {
	return &Response{c: c, statusOk: http.StatusOK}
}

func (this *Response) GetStatusOk() int {
	return this.statusOk
}

func (this *Response) Success(data map[string]interface{}) *Response {
	this.Code = 200
	this.Message = "success"
	this.Data = data
	return this
}
func (this *Response) Fail(code int, message string) *Response {
	this.Code = -1
	this.Message = message
	this.Data = make(map[string]interface{})
	return this
}

func (this *Response) Json() {
	this.c.JSON(this.statusOk, this)
}

func (this *Response) Xml() {
	this.c.XML(this.statusOk, gin.H{
		"message": this.Message,
		"Code":    this.Code,
	})
}

func (this *Response) Yaml() {
	this.c.YAML(this.statusOk, gin.H{
		"message": this.Message,
		"Code":    this.Code,
	})

}

var (
	clientset     *kubernetes.Clientset
	dynamicClient dynamic.Interface
)

func init() {
	fmt.Println("liuyucaho  router list get ")
	var kubeconfig *string
	if home, _ := os.Getwd(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "conf", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Failed to build config:", err)
		return
	}
	clientset, err = kubernetes.NewForConfig(config)
	dynamicClient, err = dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println("Failed to create clientset:", err)
		return
	}
}

func GetClientset() *kubernetes.Clientset {
	return clientset
}
