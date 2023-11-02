package main

import (
	"fmt"
	"log"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// 获取 Kubernetes 配置文件路径
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}

	// 创建 Kubernetes 客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// 定义 Helm 安装命令
	helmCommand := []string{"helm", "install", "my-release", "stable/my-chart"}

	// 定义执行命令的容器名称和命名空间
	containerName := "helm-exec-container"
	namespace := "default"

	// 创建执行命令的请求
	req := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name("my-pod").
		Namespace(namespace).
		SubResource("exec").
		VersionedParams(&v1.PodExecOptions{
			Container: containerName,
			Command:   helmCommand,
			Stdin:     false,
			Stdout:    true,
			Stderr:    true,
		}, v1.ParameterCodec)

	// 执行命令并获取输出
	executor, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		log.Fatalf("Failed to create executor: %v", err)
	}
	var stdout, stderr []byte
	err = executor.Stream(remotecommand.StreamOptions{
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

	// 输出命令执行结果
	fmt.Printf("stdout: %s\n", stdout)
	fmt.Printf("stderr: %s\n", stderr)
}
