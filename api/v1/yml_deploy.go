package v1

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/restmapper"
	"log"
	"net/http"
)

/*
type BaseApi struct {
	Base *BaseController
}*/

func YmlDeployUtil(fileName string) (string, error) {
	//home := GetHomePath()
	nameSpace := metav1.NamespaceDefault
	/*var kubeconfig *string
	if home, _ := os.Getwd(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "conf", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Failed to build config:", err)
		return "", err
	}
	//clientset, err = kubernetes.NewForConfig(config)
	//client, err := dynamic.NewForConfig(config)

	// 创建一个k8s客户端
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("%v", err)
		return "", err
	}
	dd, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}*/

	filebytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	decoder := yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader(filebytes), 100)
	for {
		var rawObj runtime.RawExtension
		if err = decoder.Decode(&rawObj); err != nil {
			break
		}

		obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
		unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			log.Fatal(err)
		}

		unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}

		gr, err := restmapper.GetAPIGroupResources(clientset.Discovery())
		if err != nil {
			log.Fatal(err)
		}

		mapper := restmapper.NewDiscoveryRESTMapper(gr)
		mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
		if err != nil {
			log.Fatal(err)
		}

		var dri dynamic.ResourceInterface
		if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
			if unstructuredObj.GetNamespace() == "" {
				unstructuredObj.SetNamespace(nameSpace)
			}
			dri = dynamicClient.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())
		} else {
			dri = dynamicClient.Resource(mapping.Resource)
		}

		obj2, err := dri.Create(context.Background(), unstructuredObj, metav1.CreateOptions{})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s/%s created", obj2.GetKind(), obj2.GetName())
	}
	return "", nil
}

/*
	/*var kubeconfig *string
	if home, _ := os.Getwd(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "conf", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Failed to build config:", err)
		return "", err
	}
	//clientset, err = kubernetes.NewForConfig(config)
	//client, err := dynamic.NewForConfig(config)

	// 创建一个k8s客户端
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("%v", err)
		return "", err
	}
	dd, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}*/
//defer this.Base.Catch(NewResponse(c))
//form, _ := c.MultipartForm()
//file, _ := form.File["file"]
//_ = []byte(file)
//yml, _ := ioutil.ReadAll(c.Request.Body)*/
func (this *BaseApi) YmlDeploy(c *gin.Context) {
	//home := GetHomePath()
	nameSpace := metav1.NamespaceDefault
	fmt.Println(c.Query("pipelineId"))
	//fmt.Println(c.Query("yml"))
	yml := c.Query("yml")
	filebytes := []byte(yml)
	decoder := yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader(filebytes), 100)
	for {
		var rawObj runtime.RawExtension
		if err := decoder.Decode(&rawObj); err != nil {
			break
		}

		obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
		unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			log.Fatal(err)
		}

		unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}

		gr, err := restmapper.GetAPIGroupResources(clientset.Discovery())
		if err != nil {
			log.Fatal(err)
		}

		mapper := restmapper.NewDiscoveryRESTMapper(gr)
		mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
		if err != nil {
			log.Fatal(err)
		}

		var dri dynamic.ResourceInterface
		if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
			if unstructuredObj.GetNamespace() == "" {
				unstructuredObj.SetNamespace(nameSpace)
			}
			dri = dynamicClient.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())
		} else {
			dri = dynamicClient.Resource(mapping.Resource)
		}

		obj2, err := dri.Create(context.Background(), unstructuredObj, metav1.CreateOptions{})
		//time.Sleep(10 * time.Second)
		//_ = dri.Delete(context.Background(), "ikos", metav1.DeleteOptions{})
		if err != nil {
			log.Fatal(err)
		}
		create(c.Query("pipelineId"))
		fmt.Printf("%s/%s created", obj2.GetKind(), obj2.GetName())
		req("http://192.168.1.224:8000", c.Query("pipelineId"))
	}

	NewResponse(c).Success(map[string]interface{}{}).Json()
}
func create(content string) (string, error) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "write-file-pod",
			Namespace: "default",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{{
				Name:  "write-file-container",
				Image: "alpine",
				Command: []string{
					"sh",
					"-c",
					fmt.Sprintf("echo '%s' > /mnt/ClientFolder/pipeline.txt", content),
				},
				VolumeMounts: []corev1.VolumeMount{{
					Name:      "share-folder",
					MountPath: "/mnt/ClientFolder",
				}},
			}},
			Volumes: []corev1.Volume{{
				Name: "share-folder",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/mnt/ClientFolder",
					},
				},
			}},
		},
	}

	pod, err := clientset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Pod created successfully.")

	// 等待 Pod 完成
	// 这里可以根据你的需求添加适当的等待时间或轮询逻辑，以确保 Pod 完成并文件写入成功。

	// 删除 Pod
	deletePolicy := metav1.DeletePropagationForeground
	if err := clientset.CoreV1().Pods("default").Delete(context.TODO(), pod.Name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err.Error())
	}

	fmt.Println("Pod deleted successfully.")

	//在上述代码中，我们创建了一个 Pod，使用 Alpine 镜像运行一个简单的 shell 命令来将 "hello world" 写入到 `/mnt/ShareFolder/file.txt` 文件中。然后，我们等待 Pod 完成并成功写入文件后，将其删除。
	return "", nil
}

func req(ipport, pipelineId string) {
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 创建 GET 请求
	req, err := http.NewRequest("GET", ipport+"/api/v1/loopScantestIkosInfo/"+pipelineId, nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 打印响应内容
	fmt.Println("响应内容:", string(body))
}
