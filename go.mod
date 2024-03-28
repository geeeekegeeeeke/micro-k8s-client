module micro-k8s-client

go 1.20

require (
	github.com/apache/dubbo-go v1.5.4
	github.com/apache/dubbo-go-hessian2 v1.7.0
	github.com/boombuler/barcode v1.0.1
	github.com/creack/pty v1.1.18
	github.com/docker/docker v24.0.7+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/garyburd/redigo v1.6.2
	github.com/gin-contrib/i18n v0.0.1
	github.com/gin-gonic/gin v1.9.1
	github.com/glebarez/sqlite v1.9.0
	github.com/go-gormigrate/gormigrate/v2 v2.1.1
	github.com/go-playground/validator/v10 v10.14.1
	github.com/gophercloud/gophercloud v0.3.0
	github.com/gorilla/websocket v1.5.0
	github.com/helm/helm v2.17.0+incompatible
	github.com/jinzhu/copier v0.3.5
	github.com/kubernetes/kompose v1.31.2
	github.com/mozillazg/go-pinyin v0.20.0
	github.com/nicksnyder/go-i18n/v2 v2.1.2
	github.com/opencontainers/image-spec v1.1.0-rc5
	github.com/pkg/errors v0.9.1
	github.com/shirou/gopsutil v2.19.12+incompatible
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/viper v1.16.0
	github.com/tealeg/xlsx v1.0.5
	github.com/tuotoo/qrcode v0.0.0-20190222102259-ac9c44189bf2
	gitlab.stagingvip.net/publicGroup/public v0.0.0-20201029110713-ec938beba922
	golang.org/x/crypto v0.14.0
	gorm.io/gorm v1.25.5
	k8s.io/api v0.28.2
	k8s.io/apimachinery v0.28.2
	k8s.io/client-go v0.26.2
//k8s.io/helm v2.17.0+incompatible // indirect
)

require google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect

require (
	//github.com/1Panel-dev/1Panel v1.7.4 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/BurntSushi/toml v1.2.0 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/Workiva/go-datastructures v1.0.50 // indirect
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5 // indirect
	github.com/alibaba/sentinel-golang v0.6.2 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1755 // indirect
	github.com/apache/dubbo-getty v1.3.10 // indirect
	github.com/buger/jsonparser v0.0.0-20181115193947-bf1c66bbce23 // indirect
	github.com/bytedance/sonic v1.10.0-rc3 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/compose-spec/compose-go v1.20.0 // indirect
	github.com/containerd/containerd v1.7.7 // indirect
	github.com/creasty/defaults v1.3.0 // indirect
	github.com/cyphar/filepath-securejoin v0.2.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/distribution/reference v0.5.0 // indirect
	github.com/docker/distribution v2.8.2+incompatible // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dubbogo/gost v1.9.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/elazarl/goproxy v0.0.0-20191011121108-aa519ddbe484 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/fsouza/go-dockerclient v1.10.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/glebarez/go-sqlite v1.21.2 // indirect
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/jonboulle/clockwork v0.3.1-0.20230117163003-a89700cec744 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f // indirect
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/maruel/rs v0.0.0-20150922171536-2c81c4312fe4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-shellwords v1.0.12 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/patternmatcher v0.6.0 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/moby/sys/sequential v0.5.0 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/nacos-group/nacos-sdk-go v1.0.0 // indirect
	github.com/novln/docker-parser v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/runc v1.1.7 // indirect
	github.com/openshift/api v3.9.0+incompatible // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.9 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b // indirect
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/cobra v1.7.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/toolkits/concurrent v0.0.0-20150624120057-a4371d70e3e3 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/goleak v1.2.1 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/arch v0.4.0 // indirect
	golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/oauth2 v0.11.0 // indirect
	golang.org/x/sync v0.4.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/term v0.15.0 // indirect
	golang.org/x/text v0.13.0
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.13.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/helm v2.17.0+incompatible // indirect
	k8s.io/klog/v2 v2.100.1 // indirect
	k8s.io/utils v0.0.0-20230406110748-d93618cff8a2 // indirect
	modernc.org/libc v1.22.5 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.5.0 // indirect
	modernc.org/sqlite v1.23.1 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
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
	//k8s.io/api => k8s.io/api v0.26.1
	//k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.26.1
	//k8s.io/apimachinery => k8s.io/apimachinery v0.26.1
	//k8s.io/apiserver => k8s.io/apiserver v0.26.1
	//k8s.io/cli-runtime => k8s.io/cli-runtime v0.26.1
	//k8s.io/client-go => k8s.io/client-go v0.26.1
	//k8s.io/code-generator => k8s.io/code-generator v0.26.1
	//k8s.io/component-base => k8s.io/component-base v0.26.1
	//k8s.io/gengo => k8s.io/gengo v0.0.0-20220902162205-c0856e24416d
	//k8s.io/klog/v2 => k8s.io/klog/v2 v2.90.0
	//k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20230224204730-66828de6f33b
	//k8s.io/kubectl => k8s.io/kubectl v0.26.1
	//k8s.io/metrics => k8s.io/metrics v0.26.1
	//k8s.io/utils => k8s.io/utils v0.0.0-20230202215443-34013725500c
	github.com/cucumber/godog => github.com/cucumber/godog v0.12.6
	gitlab.stagingvip.net/publicGroup/public v0.0.0-20201029110713-ec938beba922 => ./doc/gitlab.stagingvip.net/publicGroup/public
	golang.org/x/net => golang.org/x/net v0.17.0
	k8s.io/api => k8s.io/api v0.22.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.22.4
	k8s.io/apiserver => k8s.io/apiserver v0.22.4
	k8s.io/client-go => k8s.io/client-go v0.22.4
)
