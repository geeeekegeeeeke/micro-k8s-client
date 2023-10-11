package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

type ApplyController struct {
	Base *BaseController
}

func (this *ApplyController) GetApply(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	//获取NODE
	fmt.Println("####### 获取node ######")
	/*nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, nds := range nodes.Items {
		fmt.Printf("NodeName: %s\n", nds.Name)
	}*/

	//获取 指定NODE 的详细信息
	fmt.Println("\n ####### node详细信息 ######")
	//nodeName := "k8s-master2"
	//nodeRel, err := clientset.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	//Response{}
	NewResponse(c).Success(map[string]interface{}{}).Json()
}
func (this *ApplyController) Namespace(c *gin.Context) {
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
	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get name space:", err)
	}
	for i, ns := range namespaces.Items {
		fmt.Printf("[%d] %s\n", i, ns.GetName())
	}
}
func (this *ApplyController) Secret(c *gin.Context) {
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

	secrets, err := clientset.CoreV1().Secrets("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get secret:", err)
	}
	for i, secret := range secrets.Items {
		fmt.Printf("[%d] %s\n", i, secret.GetName())
	}
}
func (this *ApplyController) SetSecret(c *gin.Context) {
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

	data := make(map[string][]byte)
	data["user"] = []byte("admin")
	data["password"] = []byte("password")

	// https://godoc.org/k8s.io/api/core/v1
	// https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1#SecretInterface
	secrets, err := clientset.CoreV1().Secrets("default").Create(&v1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind: "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "generic-secret",
			Namespace: "default",
		},
		Data: data,
	})
	fmt.Println(secrets)
}
func (this *ApplyController) ConfigMap(c *gin.Context) {
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

	configMaps, err := clientset.CoreV1().ConfigMaps("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get config map:", err)
	}
	for i, cm := range configMaps.Items {
		fmt.Printf("[%d] %s\n", i, cm.GetName())
	}
}
