package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

type ServiceController struct {
	Base *BaseController
}

func (this *ServiceController) DeployService(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	// 创建Service对象
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-service",
			Namespace: "default",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": "my-app",
			},
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					Port:       8080,
					TargetPort: intstr.FromInt(8080),
				},
			},
		},
	}

	// 部署Service
	services, err := clientset.CoreV1().Services("default").Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("Failed to deploy service:", err)
		return
	}

	fmt.Println("Service deployed successfully.")

	NewResponse(c).Success(map[string]interface{}{"svc": services}).Json()
}
func (this *ServiceController) ListService(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))

	// https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1
	services, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get services:", err)
	}
	for i, svc := range services.Items {
		fmt.Printf("[%d] %s\n", i, svc.GetName())
	}
	NewResponse(c).Success(map[string]interface{}{"svc": services}).Json()
}
func (this *ServiceController) ListServiceAccount(c *gin.Context) {
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
	serviceAccounts, err := clientset.CoreV1().ServiceAccounts("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get service accounts:", err)
	}
	for i, sa := range serviceAccounts.Items {
		fmt.Printf("[%d] %s\n", i, sa.GetName())
	}
	NewResponse(c).Success(map[string]interface{}{"svcac": serviceAccounts}).Json()

}
