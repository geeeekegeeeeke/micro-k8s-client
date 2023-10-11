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
)

type StorageController struct {
	Base *BaseController
}

func (this *StorageController) ListPersistent(c *gin.Context) {
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
	pvs, err := clientset.CoreV1().PersistentVolumes().List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get persistent volumes:", err)
	}
	for i, pv := range pvs.Items {
		fmt.Printf("[%d] %s\n", i, pv.GetName())
	}
}
func (this *StorageController) ListPersistentVol(c *gin.Context) {
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
	pvcs, err := clientset.CoreV1().PersistentVolumeClaims("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get persistent volume claim:", err)
	}
	for i, pvc := range pvcs.Items {
		fmt.Printf("[%d] %s\n", i, pvc.GetName())
	}
}
