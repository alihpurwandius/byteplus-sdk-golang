package billing

import (
	"github.com/alihpurwandius/byteplus-sdk-golang/service/billing"
)

var (
	ak              = "<Your AK>"
	sk              = "<Your SK>"
	DefaultInstance = billing.DefaultInstance
)

func init() {
	DefaultInstance.Client.SetAccessKey(ak)
	DefaultInstance.Client.SetSecretKey(sk)
}
