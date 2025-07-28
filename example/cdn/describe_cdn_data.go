package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/alihpurwandius/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeCdnData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnData(&cdn.DescribeCdnDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}
