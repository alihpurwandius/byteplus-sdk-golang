package live_v20200801_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/alihpurwandius/byteplus-sdk-golang/base"
	live_v20200801 "github.com/alihpurwandius/byteplus-sdk-golang/service/live/v20200801"
)

func Test_ListStorageSpace(t *testing.T) {
	instance := live_v20200801.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &live_v20200801.ListStorageSpaceBody{}

	resp, err := instance.ListStorageSpace(param)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
