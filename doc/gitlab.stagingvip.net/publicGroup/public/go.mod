module cloudlive

go 1.14

require (
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.614
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef
	github.com/boombuler/barcode v1.0.0
	github.com/garyburd/redigo v1.6.2 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v7 v7.4.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.16
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742
	github.com/phachon/go-logger v0.0.0-20191215032019-86e4227f71ea
	github.com/tealeg/xlsx v1.0.5
	github.com/tuotoo/qrcode v0.0.0-20190222102259-ac9c44189bf2
	github.com/valyala/fasthttp v1.17.0
	github.com/willf/bitset v1.1.11 // indirect
	gitlab.stagingvip.net/publicGroup/public v0.0.0-20201029110713-ec938beba922
)

//replace gitlab.stagingvip.net/publicGroup/public => D:\workspace\gitlab.stagingvip.net\publicGroup\public
replace gitlab.stagingvip.net/publicGroup/public => D:/go-project/dubbo-go-demo-comsumer/doc/gitlab.stagingvip.net/publicGroup/public
