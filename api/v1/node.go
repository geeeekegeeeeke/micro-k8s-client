package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"micro-k8s-client/filter"
)

/*
	type BaseApi struct {
		Base *BaseController
	}
*/
func (this *BaseApi) ListNode(c *gin.Context) {
	//defer this.Base.Catch(NewResponse(c))
	/*
		fmt.Println("liuyucaho  routers list get ")
		var kubeconfig *string
		if home, _ := os.Getwd(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, "conf", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		// NewForConfig creates a new Clientset for the given config.
		// https://godoc.org/k8s.io/client-go/kubernetes#NewForConfig
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatal(err)
		}
	*/
	// https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get nodes:", err)
	}

	for i, node := range nodes.Items {
		fmt.Printf("[%d] %s\n", i, node.GetName())
	}
	NewResponse(c).Success(map[string]interface{}{"node": nodes}).Json()
}

func (this *BaseApi) GetNodeInfo(c *gin.Context) {
	//defer this.Base.Catch(NewResponse(c))
	//获取 指定NODE 的详细信息
	fmt.Println("\n ####### node详细信息 ######")
	info := filter.NewNodeFilter(c).NodeInfo()
	fmt.Println("\n ####### node详细信息 ######", info)
	fmt.Println(len(info))

	nodeName := info["name"]
	/*func GetNode(clientset kubernetes.Interface, name string) (*v1.Node, error) {
	    return clientset.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})
	}
	*/
	nodeRel, err := clientset.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
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
	NewResponse(c).Success(map[string]interface{}{"nodeInfo": nodeRel}).Json()
}
