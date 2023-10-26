package main

//
//import (
//	"context"
//	"flag"
//	"fmt"
//	"io/ioutil"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/apimachinery/pkg/runtime/schema"
//	"k8s.io/client-go/dynamic"
//	//"k8s.io/client-go/restmapper"
//	"k8s.io/client-go/restmapper"
//	"path/filepath"
//
//	"k8s.io/apimachinery/pkg/runtime"
//	"k8s.io/apimachinery/pkg/runtime/serializer"
//	"k8s.io/client-go/tools/clientcmd"
//	"k8s.io/client-go/util/homedir"
//)
//
//func main() {
//	// 解析命令行参数获取 kubeconfig 文件路径和命名空间
//	var kubeconfig string
//	var namespace string
//	var yamlFile string
//	flag.StringVar(&kubeconfig, "kubeconfig", getDefaultKubeconfigPath(), "Path to the kubeconfig file")
//	flag.StringVar(&namespace, "namespace", "default", "Namespace for the resource")
//	flag.StringVar(&yamlFile, "yaml-file", "", "Path to the YAML file")
//	flag.Parse()
//
//	// 加载 kubeconfig 文件并创建 Kubernetes 客户端
//	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
//	if err != nil {
//		panic(fmt.Errorf("failed to build config: %v", err))
//	}
//	clientset, err := dynamic.NewForConfig(config)
//	if err != nil {
//		panic(fmt.Errorf("failed to create clientset: %v", err))
//	}
//
//	// 读取 YAML 文件内容
//	yamlContent, err := ioutil.ReadFile(yamlFile)
//	if err != nil {
//		panic(fmt.Errorf("failed to read YAML file: %v", err))
//	}
//
//	// 解析 YAML 文件为 Unstructured 对象
//	decode := serializer.NewCodecFactory(runtime.NewScheme()).UniversalDeserializer().Decode
//	obj, _, err := decode([]byte(yamlContent), nil, nil)
//	if err != nil {
//		panic(fmt.Errorf("failed to decode YAML: %v", err))
//	}
//
//	// 将 Unstructured 对象转换为 GVK（Group, Version, Kind）和 Namespace
//	gvk := obj.GetObjectKind().GroupVersionKind()
//	apiVersion, kind := gvk.ToAPIVersionAndKind()
//	// 获取 GVR（Group, Version, Resource）
//	mapper := restmapper.NewDeferredDiscoveryRESTMapper()
//	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
//	if err != nil {
//		panic(fmt.Errorf("failed to get REST mapping: %v", err))
//	}
//	gvr := mapping.Resource
//	// 创建资源
//	_, err = clientset.Resource(nil).Namespace(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
//	if err != nil {
//		panic(fmt.Errorf("failed to create resource: %v", err))
//	}
//
//	fmt.Printf("Resource created: %s/%s\n", apiVersion, kind)
//}
//
//// 获取默认的 kubeconfig 文件路径
//func getDefaultKubeconfigPath() string {
//	home := homedir.HomeDir()
//	return filepath.Join(home, ".kube", "config")
//}
//
//
