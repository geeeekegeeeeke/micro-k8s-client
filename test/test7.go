package main

import (
	"fmt"
	"github.com/helm/helm/pkg/helm"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/helm/pkg/helm"
	//"k8s.io/helm/pkg/kube"
)

func main() {
	// 获取 Kubernetes 配置文件路径
	kubeconfigPath := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)

	// 创建 Kubernetes 客户端配置
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}

	// 创建 Kubernetes 客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// 创建 Helm 客户端
	helmClient := helm.NewClient(helm.Host("https://your-kubernetes-api-url"))

	// 创建 Kubernetes 配置
	kubeConfig := &kube.Config{
		APIHost:               config.Host,
		Username:              config.Username,
		Password:              config.Password,
		InsecureSkipTLSVerify: config.Insecure,
	}

	// 创建 Helm 命令
	cmd := []string{"install", "my-release", "stable/my-chart"}

	// 执行 Helm 命令
	res, err := helmClient.Run(cmd, kubeConfig, os.Stdout)
	if err != nil {
		log.Fatalf("Failed to run Helm command: %v", err)
	}

	fmt.Println("Helm command executed successfully")
	fmt.Println(res)
}
