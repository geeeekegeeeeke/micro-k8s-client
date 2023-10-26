package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

var gvr = schema.GroupVersionResource{
	Group:    "",
	Version:  "v1",
	Resource: "deployments",
}

func createCrontabWithYaml(client dynamic.Interface, namespace string, yamlData string) error {
	decoder := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	obj := &unstructured.Unstructured{}
	if _, _, err := decoder.Decode([]byte(yamlData), nil, obj); err != nil {
		return err
	}
	fmt.Println("decoder-------------", decoder)

	fmt.Println("obj-----------", obj)
	utd, err := client.Resource(gvr).Namespace(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	data, err := utd.MarshalJSON()
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}

func main() {
	var kubeconfig *string
	if home, _ := os.Getwd(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "conf", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Failed to build config:", err)
		return
	}
	//clientset, err = kubernetes.NewForConfig(config)
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println("Failed to create clientset:", err)
		return
	}
	createData := `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-k8s
spec:
  replicas: 2
  selector:
    matchLabels:
      app: test-k8s
  template:
    metadata:
      labels:
        app: test-k8s
    spec:
      containers:
      - name: test-k8s 
        image: ccr.ccs.tencentyun.com/k8s-tutorial/test-k8s:v1 

`
	err = createCrontabWithYaml(client, "default", createData)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s %s %s %s\n", ct.Namespace, ct.Name, ct.Spec.CronSpec, ct.Spec.Image)
}
