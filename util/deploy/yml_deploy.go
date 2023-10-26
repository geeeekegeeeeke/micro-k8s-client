package deploy

//func YmlDeployUtil(fileName string) (string, error) {
//	//home := GetHomePath()
//	nameSpace := metav1.NamespaceDefault
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
//		fmt.Println("Failed to build config:", err)
//		return "", err
//	}
//	//clientset, err = kubernetes.NewForConfig(config)
//	//client, err := dynamic.NewForConfig(config)
//
//	// 创建一个k8s客户端
//	clientSet, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		fmt.Printf("%v", err)
//		return "", err
//	}
//	dd, err := dynamic.NewForConfig(config)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	filebytes, err := ioutil.ReadFile(fileName)
//	if err != nil {
//		fmt.Printf("%v\n", err)
//	}
//
//	decoder := yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader(filebytes), 100)
//	for {
//		var rawObj runtime.RawExtension
//		if err = decoder.Decode(&rawObj); err != nil {
//			break
//		}
//
//		obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
//		unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}
//
//		gr, err := restmapper.GetAPIGroupResources(clientSet.Discovery())
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		mapper := restmapper.NewDiscoveryRESTMapper(gr)
//		mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		var dri dynamic.ResourceInterface
//		if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
//			if unstructuredObj.GetNamespace() == "" {
//				unstructuredObj.SetNamespace(nameSpace)
//			}
//			dri = dd.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())
//		} else {
//			dri = dd.Resource(mapping.Resource)
//		}
//
//		obj2, err := dri.Create(context.Background(), unstructuredObj, metav1.CreateOptions{})
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Printf("%s/%s created", obj2.GetKind(), obj2.GetName())
//	}
//	return "", nil
//}
