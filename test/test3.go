package main

//
//import (
//	"context"
//	"flag"
//	"fmt"
//	"io/ioutil"
//	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
//	"k8s.io/apimachinery/pkg/runtime"
//	"k8s.io/apimachinery/pkg/runtime/schema"
//	"k8s.io/apimachinery/pkg/runtime/serializer"
//	"k8s.io/client-go/discovery"
//	"k8s.io/client-go/dynamic"
//	"k8s.io/client-go/tools/clientcmd"
//	"log"
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
//	// 解析 YAML 文件为运行时对象
//	decode := serializer.NewCodecFactory(runtime.NewScheme()).UniversalDeserializer().Decode
//	obj, _, err := decode([]byte(yamlContent), nil, nil)
//	if err != nil {
//		panic(fmt.Errorf("failed to decode YAML: %v", err))
//	}
//
//	// 将对象转换为 Unstructured 对象
//	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
//	if err != nil {
//		panic(fmt.Errorf("failed to convert object to Unstructured: %v", err))
//	}
//
//	// 创建资源
//	createdObj, err := clientset.Resource(getGVR(obj)).Namespace(namespace).Create(context.TODO(), &unstructured.Unstructured{
//		Object: unstructuredObj,
//	}, v1.CreateOptions{})
//	if err != nil {
//		panic(fmt.Errorf("failed to create resource: %v", err))
//	}
//
//	fmt.Printf("Resource created: %s/%s\n", createdObj)
//}
//
///*
//// 获取资源的 GroupVersionKind
//
//	func getResourceMapping(obj runtime.Object) schema.GroupVersionResource {
//		gvk := obj.GetObjectKind().GroupVersionKind()
//		return gvk
//	}
//*/
//func getGVR(discoveryClient discovery.DiscoveryInterface, group, version, kind string) (schema.GroupVersionResource, error) {
//	resourceList, err := discoveryClient.ServerPreferredResources()
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	for _, resource := range resourceList {
//		for _, apiResource := range resource.APIResources {
//			if apiResource.Kind == kind && apiResource.Group == group && apiResource.Version == version {
//				versionResource := schema.GroupVersionResource{
//					Group:    apiResource.Group,
//					Version:  apiResource.Version,
//					Resource: apiResource.Name,
//				}
//				return versionResource, nil
//				//return fmt.Sprintf("%s.%s/%s", apiResource.Group, apiResource.Version, apiResource.Name), nil
//			}
//		}
//	}
//
//	var gvr = schema.GroupVersionResource{
//		Group:    "stable.example.com",
//		Version:  "v1",
//		Resource: "crontabs",
//	}
//	return gvr, nil
//	//return , nil
//}
