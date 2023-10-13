package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

type StorageController struct {
	Base *BaseController
}

func (this *StorageController) ListPersistent(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	// https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1
	pvs, err := clientset.CoreV1().PersistentVolumes().List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get persistent volumes:", err)
	}
	for i, pv := range pvs.Items {
		fmt.Printf("[%d] %s\n", i, pv.GetName())
	}
	NewResponse(c).Success(map[string]interface{}{"pvs": pvs}).Json()

}
func (this *StorageController) ListPersistentVol(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))

	// https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1
	pvcs, err := clientset.CoreV1().PersistentVolumeClaims("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get persistent volume claim:", err)
	}
	for i, pvc := range pvcs.Items {
		fmt.Printf("[%d] %s\n", i, pvc.GetName())
	}
	NewResponse(c).Success(map[string]interface{}{"pvcVol": pvcs}).Json()
}
