package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ApplyController struct {
	Base *BaseController
}

func (this *ApplyController) DeployTomcatApp(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	// 设置命名空间和应用名称
	namespace := "default"
	appName := "tomcat"

	// 部署Tomcat应用程序
	err := deployTomcat(clientset, namespace, appName)
	if err != nil {
		log.Fatalf("Failed to deploy Tomcat: %v", err)
	}

	// 创建Service
	err = createService(clientset, namespace, appName)
	if err != nil {
		log.Fatalf("Failed to create service: %v", err)
	}

	// 等待Service的External IP分配
	fmt.Println("Waiting for External IP...")
	time.Sleep(30 * time.Second)

	// 获取Service的External IP
	service, err := clientset.CoreV1().Services(namespace).Get(context.TODO(), appName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Failed to get service: %v", err)
	}

	if len(service.Status.LoadBalancer.Ingress) > 0 {
		externalIP := service.Status.LoadBalancer.Ingress[0].IP
		fmt.Printf("Tomcat is accessible at http://%s\n", externalIP)
	} else {
		fmt.Println("External IP not available yet. Please try accessing Tomcat later.")
	}
	/*namespace := "default"
	serviceList, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
	    panic(err.Error())
	}

	for _, service := range serviceList.Items {
	    fmt.Printf("Service: %s\n", service.Name)
	    for _, ingress := range service.Status.LoadBalancer.Ingress {
	        fmt.Printf("IP: %s\n", ingress.IP)
	    }
	}*/
	NewResponse(c).Success(map[string]interface{}{}).Json()
}

// 部署Tomcat应用程序
func deployTomcat(clientset *kubernetes.Clientset, namespace, appName string) error {
	// 读取Tomcat的Deployment YAML文件
	home, _ := os.Getwd()
	confPath := filepath.Join(home, "conf", "tomcat-deployment.yaml")
	fmt.Println(home)
	deploymentYAML, err := ioutil.ReadFile(confPath)
	if err != nil {
		return fmt.Errorf("failed to read deployment YAML file: %v", err)
	}

	// 替换YAML文件中的占位符
	deploymentYAMLStr := string(deploymentYAML)
	deploymentYAMLStr = strings.ReplaceAll(deploymentYAMLStr, "{{NAMESPACE}}", namespace)
	deploymentYAMLStr = strings.ReplaceAll(deploymentYAMLStr, "{{APP_NAME}}", appName)

	// 创建Tomcat的Deployment
	_, err = clientset.AppsV1().Deployments(namespace).Create(context.TODO(), &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: appName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": appName,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": appName,
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            appName,
							Image:           "tomcat:9.0",
							ImagePullPolicy: v1.PullIfNotPresent,
							Ports: []v1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create Tomcat deployment: %v", err)
	}
	// 应用Tomcat的Deployment YAML文件
	_ = clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(namespace).
		Name(appName).
		SubResource("apply").
		Body([]byte(deploymentYAMLStr)).Do(context.TODO())

	//result, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, v1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to apply Tomcat deployment YAML: %v", err)
	}

	return nil
}

// 创建Service
func createService(clientset *kubernetes.Clientset, namespace, appName string) error {
	// 创建Service
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: appName,
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app": appName,
			},
			Ports: []v1.ServicePort{
				{
					Name:       "http",
					Port:       80,
					TargetPort: intstr.FromInt(8080),
				},
			},
			Type: v1.ServiceTypeNodePort,
		},
	}

	_, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create service: %v", err)
	}

	return nil

}
func (this *ApplyController) DeployApp(c *gin.Context) {
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
func int32Ptr(i int32) *int32 {
	return &i
}
func (this *ApplyController) Namespace(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))

	// https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get name space:", err)
	}
	for i, ns := range namespaces.Items {
		fmt.Printf("[%d] %s\n", i, ns.GetName())
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()
}
func (this *ApplyController) Secret(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))

	secrets, err := clientset.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get secret:", err)
	}
	for i, secret := range secrets.Items {
		fmt.Printf("[%d] %s\n", i, secret.GetName())
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()
}
func (this *ApplyController) SetSecret(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	data := make(map[string][]byte)
	data["user"] = []byte("admin")
	data["password"] = []byte("password")

	// https://godoc.org/k8s.io/api/core/v1
	// https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1#SecretInterface
	secrets, err := clientset.CoreV1().Secrets("default").Create(context.TODO(), &v1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind: "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "generic-secret",
			Namespace: "default",
		},
		Data: data,
	}, metav1.CreateOptions{})
	fmt.Println(secrets)
	if err != nil {
		panic(err)
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()
}
func (this *ApplyController) ConfigMap(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	// kubernetesの設定ファイルのパスを組み立てる
	configMaps, err := clientset.CoreV1().ConfigMaps("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get config map:", err)
	}
	for i, cm := range configMaps.Items {
		fmt.Printf("[%d] %s\n", i, cm.GetName())
	}
	NewResponse(c).Success(map[string]interface{}{}).Json()
}
