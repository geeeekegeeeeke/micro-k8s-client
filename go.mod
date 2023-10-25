module gin-dubbogo-consumer

go 1.14

require (
	github.com/Azure/go-autorest/autorest v0.11.18 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.614
	github.com/apache/dubbo-go v1.5.4
	github.com/apache/dubbo-go-hessian2 v1.7.0
	github.com/boombuler/barcode v1.0.0 // indirect
	github.com/elazarl/goproxy v0.0.0-20180725130230-947c36da3153 // indirect
	github.com/emicklei/go-restful v2.9.5+incompatible // indirect
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible // indirect
	github.com/garyburd/redigo v1.6.2 // indirect
	github.com/getkin/kin-openapi v0.76.0 // indirect
	github.com/gin-gonic/gin v1.9.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/onsi/ginkgo/v2 v2.13.0 // indirect
	github.com/onsi/gomega v1.28.0 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	github.com/tuotoo/qrcode v0.0.0-20190222102259-ac9c44189bf2 // indirect
	github.com/willf/bitset v1.1.11 // indirect
	gitlab.stagingvip.net/publicGroup/public v0.0.0-20201029110713-ec938beba922
	google.golang.org/protobuf v1.31.0 // indirect
	k8s.io/api v0.26.1
	k8s.io/apiextensions-apiserver v0.26.1
	k8s.io/apimachinery v0.26.1
	k8s.io/apiserver v0.26.1
	k8s.io/cli-runtime v0.26.1
	k8s.io/client-go v0.26.1
	//k8s.io/client-go v0.24.1
	k8s.io/code-generator v0.26.1
	k8s.io/component-base v0.26.1
	k8s.io/klog/v2 v2.90.0
	k8s.io/kube-openapi v0.0.0-20230224204730-66828de6f33b
	k8s.io/kubectl v0.26.1
	//k8s.io/kubernetes v1.24.13
	k8s.io/metrics v0.26.1
	k8s.io/utils v0.0.0-20230202215443-34013725500c
	sigs.k8s.io/application v0.8.4-0.20201016185654-c8e2959e57a0
	sigs.k8s.io/controller-runtime v0.14.4
	sigs.k8s.io/controller-tools v0.11.1
	//sigs.k8s.io/kubefed v0.0.0-20230207032540-cdda80892665
	sigs.k8s.io/kustomize/api v0.12.1
	sigs.k8s.io/kustomize/kyaml v0.13.9
	sigs.k8s.io/yaml v1.3.0
)

//k8s.io/api v0.26.0
//	k8s.io/apiextensions-apiserver v0.26.0
//	k8s.io/apimachinery v0.26.0
//	k8s.io/apiserver v0.24.1
//	k8s.io/cli-runtime v0.24.1
//	k8s.io/client-go v12.0.0+incompatible
//	k8s.io/code-generator v0.26.0
//	k8s.io/component-base v0.24.1
//	k8s.io/cri-api v0.23.1
//	k8s.io/kubernetes v1.23.12
//	sigs.k8s.io/controller-runtime v0.12.1
//	sigs.k8s.io/yaml v1.3.0
//replace gitlab.stagingvip.net/publicGroup/public v0.0.0-20201029110713-ec938beba922 => D:\\go-project\\dubbo-go-demo-comsumer\\doc\\gitlab.stagingvip.net\\publicGroup
//replace gitlab.stagingvip.net/publicGroup/public v0.0.0-20201029110713-ec938beba922 => D:\go-project\dubbo-go-demo-comsumer\doc\gitlab.stagingvip.net\publicGroup
replace (
	gitlab.stagingvip.net/publicGroup/public v0.0.0-20201029110713-ec938beba922 => ./doc/gitlab.stagingvip.net/publicGroup/public
	k8s.io/api => k8s.io/api v0.26.1
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.26.1
	k8s.io/apimachinery => k8s.io/apimachinery v0.26.1
	k8s.io/apiserver => k8s.io/apiserver v0.26.1
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.26.1
	k8s.io/client-go => k8s.io/client-go v0.26.1
	k8s.io/code-generator => k8s.io/code-generator v0.26.1
	k8s.io/component-base => k8s.io/component-base v0.26.1
	k8s.io/gengo => k8s.io/gengo v0.0.0-20220902162205-c0856e24416d
	k8s.io/klog/v2 => k8s.io/klog/v2 v2.90.0
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20230224204730-66828de6f33b
	k8s.io/kubectl => k8s.io/kubectl v0.26.1
	k8s.io/metrics => k8s.io/metrics v0.26.1
	k8s.io/utils => k8s.io/utils v0.0.0-20230202215443-34013725500c

)
