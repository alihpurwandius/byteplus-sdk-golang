package cdn

import (
	"time"

	"github.com/alihpurwandius/byteplus-sdk-golang/service/cdn"
)

var (
	DefaultInstance = cdn.DefaultInstance
	ak              = ""
	sk              = ""
	operateDomain   = "operate.example.com"
	now             = time.Now()
	testStartTime   = now.Unix()
	testEndTime     = now.Add(time.Minute * 10).Unix()
	exampleDomain   = "example.com"
)

func init() {
	DefaultInstance.Client.SetAccessKey(ak)
	DefaultInstance.Client.SetSecretKey(sk)
}
