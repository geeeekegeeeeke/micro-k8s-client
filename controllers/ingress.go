package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"reflect"
)

type IngressController struct {
	Base *BaseController
}

func (this *IngressController) ListIngress(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))

	// https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1
	ingresses, err := clientset.ExtensionsV1beta1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get ingress:", err)
	}
	for i, ingress := range ingresses.Items {
		fmt.Printf("[%d] %s\n", i, ingress.GetName())
	}
}
func (this *IngressController) ListIngressInfo(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	// https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1
	ingress, err := clientset.ExtensionsV1beta1().Ingresses("jx").Get(context.TODO(), "docker-registry", metav1.GetOptions{})
	if err != nil {
		log.Fatalln("failed to get ingresses:", err)
	}

	fmt.Println(reflect.TypeOf(ingress)) // *v1beta1.Ingress
	fmt.Println(ingress)
	fmt.Println(ingress.ObjectMeta.Name) // docker-registry
}
