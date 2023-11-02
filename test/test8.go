package main

//
//import (
//	"fmt"
//	"log"
//	"os"
//	"path/filepath"
//
//	"helm.sh/helm/v3/pkg/action"
//	"helm.sh/helm/v3/pkg/cli"
//	"k8s.io/client-go/tools/clientcmd"
//)
//
//func main() {
//	// 获取 Kubernetes 配置文件路径
//	kubeconfigPath := filepath.Join(
//		os.Getenv("HOME"), ".kube", "config",
//	)
//
//	// 创建 Action 配置
//	settings := cli.New()
//
//	// 创建 Helm Action 配置
//	actionConfig := new(action.Configuration)
//
//	// 初始化 Action 配置
//	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
//		log.Fatalf("Failed to initialize Helm action configuration: %v", err)
//	}
//
//	// 创建 Helm 客户端
//	client := action.NewInstall(actionConfig)
//
//	// 设置 Chart 参数
//	client.ChartPathOptions.RepoURL = "https://charts.helm.sh/stable"
//	client.ChartPathOptions.Version = "1.2.3"
//
//	// 设置 Release 参数
//	client.ReleaseName = "my-release"
//	client.Namespace = "default"
//
//	// 执行 Helm 安装操作
//	_, err := client.Run()
//	if err != nil {
//		log.Fatalf("Failed to install Helm chart: %v", err)
//	}
//
//	fmt.Println("Helm chart installed successfully")
//}
