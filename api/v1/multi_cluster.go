package v1

//
//import (
//	"context"
//	"errors"
//	"fmt"
//	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/client-go/kubernetes"
//	"k8s.io/client-go/rest"
//	"k8s.io/client-go/tools/clientcmd"
//	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
//)
//
//func main() {
//	// 从文件中加载Kubernetes配置
//	cfg, err := clientcmd.LoadFromFile("./config")
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// 初始化一个map，用于存储所有集群的*rest.Config
//	configs := make(map[string]*rest.Config, 0)
//	var i int32
//	// 遍历所有的Context，获取每个集群的*rest.Config
//	for name := range cfg.Contexts {
//		// 构建Config对象
//		i++
//		restcfg, err := clientcmd.BuildConfigFromKubeconfigGetter("", func() (*clientcmdapi.Config, error) {
//			cfg.CurrentContext = name
//			return cfg.DeepCopy(), nil
//		})
//		if err != nil {
//			panic(err.Error())
//		}
//
//		// 将Config对象存储到map中
//		configs[name] = restcfg
//	}
//	fmt.Println("========I===================", i)
//	fmt.Println("========configs===================", len(configs))
//	// 遍历所有的*rest.Config，获取每个集群中的Namespace列表
//	for _, cluster := range configs {
//		// 通过*rest.Config创建Kubernetes客户端对象
//		cli, err := kubernetes.NewForConfig(cluster)
//		if err != nil {
//			panic(err.Error())
//		}
//
//		// 获取Namespace列表
//		nslist, err := cli.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})
//		if err != nil {
//			panic(err.Error())
//		}
//
//		// 遍历所有的Namespace，打印出名称
//		for _, ns := range nslist.Items {
//			fmt.Println(ns.Name)
//		}
//	}
//}
//
///*
//func InitClient(clusterName string) (*kubernetes.Clientset, *rest.Config, error) {
//	//数据库取出集群信息
//	master, kubeconfig, err := GetClusterInfo(clusterName)
//	if err != nil {
//		fmt.Println("get db for cluster kubeconfig error. %v ", err)
//		return nil, nil, err
//	}
//	kubeconfigJson, err := yaml.YAMLToJSON([]byte(kubeconfig))
//	if err != nil {
//		fmt.Println("yaml to json err")
//	}
//	configV1 := clientcmdapi.Config{}
//	err = json.Unmarshal(kubeconfigJson, &configV1)
//	if err != nil {
//		fmt.Println("json unmarshal kubeconfig error. %v ", err)
//		return nil, nil, err
//	}
//	// 切换匹配的版本
//	//configObject, err := clientcmdlatest.Scheme.ConvertToVersion(&configV1, clientcmdapi.SchemeGroupVersion)
//	configObject, err := "", nil
//	if err != nil {
//		fmt.Println("ConvertToVersion error. %v ", err)
//		return nil, nil, err
//	}
//	//configInternal := configObject.(*clientcmdapi.Config)
//
//	// 实例化配置信息
//	clientConfig, err := clientcmd.NewDefaultClientConfig(*configInternal, &clientcmd.ConfigOverrides{
//		ClusterDefaults: clientcmdapi.Cluster{Server: master},
//	}).ClientConfig()
//
//	if err != nil {
//		fmt.Println("build client config error. %v ", err)
//		return nil, nil, err
//	}
//	//clientConfig.QPS = defaultQPS
//	//clientConfig.Burst = defaultBurst
//	// 实例化客户端
//	clientSet, err := kubernetes.NewForConfig(clientConfig)
//
//	if err != nil {
//		fmt.Println("(%s) kubernetes.NewForConfig(%v) error.%v", master, err, clientConfig)
//		return nil, nil, err
//	}
//	return clientSet, clientConfig, nil
//
//}*/
//
///*
//func GetOutClusterClient(name string) (*K8sClient, error) {
//	clientSet, config, err := InitClient(name)
//	if err != nil {
//		return nil, err
//	}
//	return &K8sClient{Clientset: clientSet, Config: config}, nil
//}*/
//
//func GetClusterInfo(name string) (string, string, error) {
//	err := errors.New("sqlite crud")
//	if err != nil {
//		return "nil", "nil", err
//	}
//	return "nil", "nil", err
//}
