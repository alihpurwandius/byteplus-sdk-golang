package live_v20230101_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/alihpurwandius/byteplus-sdk-golang/base"
	live_v20230101 "github.com/alihpurwandius/byteplus-sdk-golang/service/live/v20230101"
)

func Test_DisableDomain(t *testing.T) {
	instance := live_v20230101.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &live_v20230101.DisableDomainBody{}

	resp, err := instance.DisableDomain(param)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
