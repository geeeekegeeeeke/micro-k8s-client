package controllers

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
	"os"
	"path/filepath"
	"time"
)

type ClientController struct {
	Base *BaseController
}

/*
# 添加
namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	for _, namespace := range namespaceList.Items {
		fmt.Println("k8s namespace")
		fmt.Printf("%s\t%s\t%s\t\n",
			namespace.Name,
			namespace.CreationTimestamp,
			namespace.Status.Phase,
		)
	}
*/
func (c ClientController) Getnode(context *gin.Context) {
	fmt.Println("liuyucaho  router list get ")
	var kubeconfig *string
	if home, _ := os.Getwd(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "conf", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	/*client, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}*/
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	nodeList, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, node := range nodeList.Items {
		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t,\n",
			node.Name,
			node.Status.Phase,
			node.Status.Addresses,
			node.Status.NodeInfo.OSImage,
			node.Status.NodeInfo.KubeletVersion,
			node.Status.NodeInfo.OperatingSystem,
			node.Status.NodeInfo.Architecture,
		)
	}

}

func (c ClientController) Getnodeinfo(context *gin.Context) {
	fmt.Println("liuyucaho  router list get ")
	var kubeconfig *string
	if home, _ := os.Getwd(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "conf", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	//client, err := dynamic.NewForConfig(config)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	/* liyuchao */
	var nodeName string
	nodeRel, err := clientset.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %s \n", nodeRel.Name)
	fmt.Printf("CreateTime: %s \n", nodeRel.CreationTimestamp)
	fmt.Printf("NowTime: %s \n", nodeRel.Status.Conditions[0].LastHeartbeatTime)
	fmt.Printf("kernelVersion: %s \n", nodeRel.Status.NodeInfo.KernelVersion)
	fmt.Printf("SystemOs: %s \n", nodeRel.Status.NodeInfo.OSImage)
	fmt.Printf("Cpu: %s \n", nodeRel.Status.Capacity.Cpu())
	fmt.Printf("docker: %s \n", nodeRel.Status.NodeInfo.ContainerRuntimeVersion)
	// fmt.Printf("Status: %s \n", nodeRel.Status.Conditions[len(nodes.Items[0].Status.Conditions)-1].Type)
	fmt.Printf("Status: %s \n", nodeRel.Status.Conditions[len(nodeRel.Status.Conditions)-1].Type)
	fmt.Printf("Mem: %s \n", nodeRel.Status.Allocatable.Memory().String())

	//获取NODE
	/*fmt.Println("####### 获取node ######")
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, nds := range nodes.Items {
		fmt.Printf("NodeName: %s\n", nds.Name)
	}

	//获取 指定NODE 的详细信息
	fmt.Println("\n ####### node详细信息 ######")
	nodeName := "k8s-master2"
	nodeRel, err := clientset.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %s \n", nodeRel.Name)
	fmt.Printf("CreateTime: %s \n", nodeRel.CreationTimestamp)
	fmt.Printf("NowTime: %s \n", nodeRel.Status.Conditions[0].LastHeartbeatTime)
	fmt.Printf("kernelVersion: %s \n", nodeRel.Status.NodeInfo.KernelVersion)
	fmt.Printf("SystemOs: %s \n", nodeRel.Status.NodeInfo.OSImage)
	fmt.Printf("Cpu: %s \n", nodeRel.Status.Capacity.Cpu())
	fmt.Printf("docker: %s \n", nodeRel.Status.NodeInfo.ContainerRuntimeVersion)
	// fmt.Printf("Status: %s \n", nodeRel.Status.Conditions[len(nodes.Items[0].Status.Conditions)-1].Type)
	fmt.Printf("Status: %s \n", nodeRel.Status.Conditions[len(nodeRel.Status.Conditions)-1].Type)
	fmt.Printf("Mem: %s \n", nodeRel.Status.Allocatable.Memory().String())
	*/
}

func (this *ClientController) OperateDeploy(*gin.Context) {
	fmt.Println("liuyucaho  router list get ")
	var kubeconfig *string
	if home, _ := os.Getwd(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "conf", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

	deployment := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name": "demo-deployment",
			},
			"spec": map[string]interface{}{
				"replicas": 2,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "demo",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app": "demo",
						},
					},

					"spec": map[string]interface{}{
						"containers": []map[string]interface{}{
							{
								"name":  "web",
								"image": "nginx:1.12",
								"ports": []map[string]interface{}{
									{
										"name":          "http",
										"protocol":      "TCP",
										"containerPort": 80,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deployment
	fmt.Println("Creating deployment...")
	result, err := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).Create(deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetName())

	// Update Deployment

	fmt.Println("Updating deployment...")
	//    You have two options to Update() this Deployment:
	//
	//    1. Modify the "deployment" variable and call: Update(deployment).
	//       This works like the "kubectl replace" command and it overwrites/loses changes
	//       made by other clients between you Create() and Update() the object.
	//    2. Modify the "result" returned by Get() and retry Update(result) until
	//       you no longer get a conflict error. This way, you can preserve changes made
	//       by other clients between Create() and Update(). This is implemented below
	//			 using the retry utility package included with client-go. (RECOMMENDED)
	//
	// More Info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).Get("demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("failed to get latest version of Deployment: %v", getErr))
		}

		// update replicas to 1
		if err := unstructured.SetNestedField(result.Object, int64(1), "spec", "replicas"); err != nil {
			panic(fmt.Errorf("failed to set replica value: %v", err))
		}

		// extract spec containers
		containers, found, err := unstructured.NestedSlice(result.Object, "spec", "template", "spec", "containers")
		if err != nil || !found || containers == nil {
			panic(fmt.Errorf("deployment containers not found or error in spec: %v", err))
		}

		// update container[0] image
		if err := unstructured.SetNestedField(containers[0].(map[string]interface{}), "nginx:1.13", "image"); err != nil {
			panic(err)
		}
		if err := unstructured.SetNestedField(result.Object, containers, "spec", "template", "spec", "containers"); err != nil {
			panic(err)
		}

		_, updateErr := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).Update(result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("update failed: %v", retryErr))
	}
	fmt.Println("Updated deployment...")

	// List Deployments

	fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
	list, err := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		replicas, found, err := unstructured.NestedInt64(d.Object, "spec", "replicas")
		if err != nil || !found {
			fmt.Printf("Replicas not found for deployment %s: error=%s", d.GetName(), err)
			continue
		}
		fmt.Printf(" * %s (%d replicas)\n", d.GetName(), replicas)
	}

	// Delete Deployment

	fmt.Println("Deleting deployment...")
	//deletePolicy := metav1.DeletePropagationForeground
	/*deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}*/
	/*if err := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).Delete( "demo-deployment", deleteOptions); err != nil {
		panic(err)
	}*/

	fmt.Println("Deleted deployment.")
}

func (this *ClientController) CreateDeploy(*gin.Context) {
	fmt.Println("liuyucaho  router list get ")
	var kubeconfig *string
	if home, _ := os.Getwd(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "conf", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

	deployment := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name": "demo-deployment",
			},
			"spec": map[string]interface{}{
				"replicas": 2,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "demo",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app": "demo",
						},
					},

					"spec": map[string]interface{}{
						"containers": []map[string]interface{}{
							{
								"name":  "web",
								"image": "nginx:1.12",
								"ports": []map[string]interface{}{
									{
										"name":          "http",
										"protocol":      "TCP",
										"containerPort": 80,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deployment
	fmt.Println("Creating deployment...")
	result, err := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).Create(deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetName())

}

func (this *ClientController) UpdateDeploy(*gin.Context) {
	fmt.Println("liuyucaho  router list get ")
	var kubeconfig *string
	if home, _ := os.Getwd(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "conf", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

	// Update Deployment

	fmt.Println("Updating deployment...")
	//    You have two options to Update() this Deployment:
	//
	//    1. Modify the "deployment" variable and call: Update(deployment).
	//       This works like the "kubectl replace" command and it overwrites/loses changes
	//       made by other clients between you Create() and Update() the object.
	//    2. Modify the "result" returned by Get() and retry Update(result) until
	//       you no longer get a conflict error. This way, you can preserve changes made
	//       by other clients between Create() and Update(). This is implemented below
	//			 using the retry utility package included with client-go. (RECOMMENDED)
	//
	// More Info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).Get("demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("failed to get latest version of Deployment: %v", getErr))
		}

		// update replicas to 1
		if err := unstructured.SetNestedField(result.Object, int64(1), "spec", "replicas"); err != nil {
			panic(fmt.Errorf("failed to set replica value: %v", err))
		}

		// extract spec containers
		containers, found, err := unstructured.NestedSlice(result.Object, "spec", "template", "spec", "containers")
		if err != nil || !found || containers == nil {
			panic(fmt.Errorf("deployment containers not found or error in spec: %v", err))
		}

		// update container[0] image
		if err := unstructured.SetNestedField(containers[0].(map[string]interface{}), "nginx:1.13", "image"); err != nil {
			panic(err)
		}
		if err := unstructured.SetNestedField(result.Object, containers, "spec", "template", "spec", "containers"); err != nil {
			panic(err)
		}

		_, updateErr := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).Update(result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("update failed: %v", retryErr))
	}
	fmt.Println("Updated deployment...")
}

func (this *ClientController) ListDeploy(*gin.Context) {
	fmt.Println("liuyucaho  router list get ")
	var kubeconfig *string
	if home, _ := os.Getwd(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "conf", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

	// List Deployments
	prompt()
	fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
	list, err := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		replicas, found, err := unstructured.NestedInt64(d.Object, "spec", "replicas")
		if err != nil || !found {
			fmt.Printf("Replicas not found for deployment %s: error=%s", d.GetName(), err)
			continue
		}
		fmt.Printf(" * %s (%d replicas)\n", d.GetName(), replicas)
	}

	// Delete Deployment
	prompt()
	fmt.Println("Deleting deployment...")
	//deletePolicy := metav1.DeletePropagationForeground
	/*deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}*/
	/*if err := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).Delete( "demo-deployment", deleteOptions); err != nil {
		panic(err)
	}*/

	fmt.Println("Deleted deployment.")
}

//func (this *ClientController) DeleteDeploy(*gin.Context) {
//	fmt.Println("liuyucaho  router list get ")
//	var kubeconfig *string
//	if home, _ := os.Getwd(); home != "" {
//		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "conf", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
//	} else {
//		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
//	}
//	flag.Parse()
//
//	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
//	if err != nil {
//		panic(err)
//	}
//	client, err := dynamic.NewForConfig(config)
//	if err != nil {
//		panic(err)
//	}
//
//	deploymentRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
//
//	// Delete Deployment
//	prompt()
//	fmt.Println("Deleting deployment...")
//	/*deletePolicy := metav1.DeletePropagationForeground
//	deleteOptions := metav1.DeleteOptions{
//		PropagationPolicy: &deletePolicy,
//	}
//	if err := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).Delete( "demo-deployment",*deleteOptions); err != nil {
//		panic(err)
//	}*/
//
//	fmt.Println("Deleted deployment.")
//}

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()

}

func (this *ClientController) GetPod(c *gin.Context) {
	fmt.Println("liuyucaho  router list get ")
	var kubeconfig *string
	if home, _ := os.Getwd(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "conf", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	//client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	/// 使用k8s.io/client-go/kubernetes生成一个ClientSet的客户端
	// 客户端生成后，就可以使用这个客户端与k8s API server进行交互，如Create/Retrieve/Update/Delete Resource
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods(namespace string)返回 PodInterface
		// PodInterface 接口拥有操作 Pod 资源的方法，例如 Create、Update、Get、List 等方法
		// 注意：Pods() 方法中 namespace 不指定则获取 Cluster 所有 Pod 列表
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the k8s cluster\n", len(pods.Items))

		// 获取指定namespace中的pod列表信息
		namespace := "default"
		pods, err = clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("There are %d pods in the namespace %s\n", len(pods.Items), namespace)
		for _, pod := range pods.Items {
			fmt.Printf("Name: %s, Status: %s, CreateTime: %s\n", pod.ObjectMeta.Name, pod.Status.Phase, pod.ObjectMeta.CreationTimestamp)
		}
		fmt.Println()
		time.Sleep(3 * time.Second)
	}

}
func (this *ClientController) GetPodInfo(c *gin.Context) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		namespace := "default"
		pod := "example-xxxxx"
		_, err = clientset.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
				pod, namespace, statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		}

		time.Sleep(10 * time.Second)
	}
}
