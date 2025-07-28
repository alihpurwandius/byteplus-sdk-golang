package cdn

import (
	"testing"

	"github.com/alihpurwandius/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeCdnOriginData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnOriginData(&cdn.DescribeCdnOriginDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}
