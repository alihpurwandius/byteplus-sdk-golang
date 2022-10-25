package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeCertConfig(t *testing.T) {
	resp, err := DefaultInstance.DescribeCertConfig(&cdn.DescribeCertConfigRequest{
		CertId: cdn.GetStrPtr("cert-2b12dd79c3ef441ea1e58a09248d0fd6"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	rsp, _ := json.Marshal(resp)
	fmt.Printf("%+v", rsp)
}
