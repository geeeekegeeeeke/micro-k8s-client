package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

type IngressController struct {
	Base *BaseController
}

func (this *IngressController) ListIngress(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	// kubernetesの設定ファイルのパスを組み立てる
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	// BuildConfigFromFlags is a helper function that builds configs from a master url or
	// a kubeconfig filepath.
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// NewForConfig creates a new Clientset for the given config.
	// https://godoc.org/k8s.io/client-go/kubernetes#NewForConfig
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1
	ingresses, err := clientset.ExtensionsV1beta1().Ingresses("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get ingress:", err)
	}
	for i, ingress := range ingresses.Items {
		fmt.Printf("[%d] %s\n", i, ingress.GetName())
	}
}
func (this *IngressController) ListIngressInfo(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	// kubernetesの設定ファイルのパスを組み立てる
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	// BuildConfigFromFlags is a helper function that builds configs from a master url or
	// a kubeconfig filepath.
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// NewForConfig creates a new Clientset for the given config.
	// https://godoc.org/k8s.io/client-go/kubernetes#NewForConfig
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1
	ingress, err := clientset.ExtensionsV1beta1().Ingresses("jx").Get("docker-registry", metav1.GetOptions{})
	if err != nil {
		log.Fatalln("failed to get ingresses:", err)
	}

	fmt.Println(reflect.TypeOf(ingress)) // *v1beta1.Ingress
	fmt.Println(ingress)
	fmt.Println(ingress.ObjectMeta.Name) // docker-registry
}
