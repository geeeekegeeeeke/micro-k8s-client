package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kubernetes/kompose/client"
	"io/ioutil"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/dynamic"
	"log"
	"os"
)

type DcokerDeployController struct {
	Base *BaseController
}

func (this *DcokerDeployController) DeployAppByDocker(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	// 定义Deployment对象
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "webapp-deployment",
			Namespace: "default",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "webapp",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "webapp",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "webapp-container",
							Image: "your-webapp-image:tag",
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}
	// 部署Deployment
	_, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("Failed to deploy webapp deployment:", err)
		return
	}
	// 定义Service对象
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "webapp-service",
			Namespace: "default",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": "webapp",
			},
			Ports: []corev1.ServicePort{
				{
					Protocol: corev1.ProtocolTCP,
					Port:     80,
					TargetPort: intstr.IntOrString{
						IntVal: 8080,
					},
				},
			},
			Type: corev1.ServiceTypeLoadBalancer,
		},
	}

	// 部署Service
	_, err = clientset.CoreV1().Services("default").Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("Failed to deploy webapp service:", err)
		return
	}
	fmt.Println("Web application deployed successfully.")
	NewResponse(c).Success(map[string]interface{}{}).Json()
}
func (this *DcokerDeployController) DeployAppByDockerCompose(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	/* 保存定义好的docker-compose描述 保存在数据库和远程集群的同步和区别以及限制*/
	// 定义 Docker Compose 文件的路径

	NewResponse(c).Success(map[string]interface{}{}).Json()

}

/* 命令行转换为 yaml*/

/* go代码引入kubect客户端执行 yaml*/

func (this *DcokerDeployController) DeployAppByHelm(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	// 定义Deployment对象

	fmt.Println("Web application deployed successfully.")
	NewResponse(c).Success(map[string]interface{}{}).Json()
}
func (this *DcokerDeployController) DeployAppBydirectCompose(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	// 定义Deployment对象
	// 使用 LoadFile 函数加载 Docker Compose 文件
	// 获取 Docker Compose 文件的路径

	//client, err := NewClient(WithErrorOnWarning())
	//assert.Check(t, is.Equal(err, nil))
	/*clientCovert, err := client.NewClient(client.WithErrorOnWarning())
	home, _ := os.Getwd()
	home = filepath.Join(home, "conf/testdata/docker-compose.yaml", "")
	fmt.Println("========================", home)
	filename := "ymlyml.yml"
	_, err = clientCovert.Convert(client.ConvertOptions{
		OutFile: filename,
		//ToStdout: true,
		InputFiles: []string{
			home,
		},
	})
	fmt.Println(err)*/

	deploymentYAML, err := ioutil.ReadFile("ymlyml.yml")

	do := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(corev1.NamespaceDefault).
		SubResource("apply").
		Body(deploymentYAML).Do(context.TODO())

	//result, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, v1.CreateOptions{})
	if err != nil {
		log.Fatalln("failed to apply Tomcat deployment YAML: %v", err)
	}
	//clientset.AppsV1().Deployments().
	get, err := do.Get()
	//do.
	fmt.Println("do ", get)
	fmt.Println("Web application deployed successfully.")

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
	ct, err := createCrontabWithYaml(dynamicClient, "default", createData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s %s %s %s\n", ct.Namespace, ct.Name, ct.Spec.CronSpec, ct.Spec.Image)

	NewResponse(c).Success(map[string]interface{}{}).Json()
}
func (this *DcokerDeployController) DeployAppComposeParam(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	// 定义Deployment对象
	// 使用 LoadFile 函数加载 Docker Compose 文件
	// 获取 Docker Compose 文件的路径

	//client, err := NewClient(WithErrorOnWarning())
	//assert.Check(t, is.Equal(err, nil))
	// 定义输出文件名
	outputFile := "docker-compose.yml"

	// 将字符串写入文件
	fmt.Println(c.Query("compose"))
	err := ioutil.WriteFile(outputFile, []byte(c.Query("compose")), 0644)
	if err != nil {
		fmt.Printf("failed to write file: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Docker Compose YAML has been written to %s\n", outputFile)

	clientCovert, err := client.NewClient(client.WithErrorOnWarning())
	//home, _ := os.Getwd()
	//home = filepath.Join(home, "conf/testdata/docker-compose.yaml", "")
	//fmt.Println("========================", home)
	filename := "liuyucaho.yml"
	_, err = clientCovert.Convert(client.ConvertOptions{
		OutFile: filename,
		//ToStdout: true,
		InputFiles: []string{
			outputFile,
		},
	})
	fmt.Println(err)
	YmlDeployUtil(filename)
	NewResponse(c).Success(map[string]interface{}{}).Json()
}
func createCrontabWithYaml(client dynamic.Interface, namespace string, yamlData string) (*Crontab, error) {
	decoder := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	obj := &unstructured.Unstructured{}
	file, err2 := ioutil.ReadFile("ymlyml.yml")
	if err2 != nil {
		log.Fatalln(err2)
	}
	if _, _, err := decoder.Decode([]byte(file), nil, obj); err != nil {
		return nil, err
	}

	utd, err := client.Resource(gvr).Namespace(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	data, err := utd.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var ct Crontab
	if err := json.Unmarshal(data, &ct); err != nil {
		return nil, err
	}
	return &ct, nil
}

var gvr = schema.GroupVersionResource{
	Group:    "stable.example.com",
	Version:  "v1",
	Resource: "crontabs",
}

type CrontabSpec struct {
	CronSpec string `json:"cronSpec"`
	Image    string `json:"image"`
}

type Crontab struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CrontabSpec `json:"spec,omitempty"`
}

type CrontabList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Crontab `json:"items"`
}
