package cdn

import (
	"testing"

	"github.com/alihpurwandius/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeEdgeTopNrtData(t *testing.T) {
	resp, err := DefaultInstance.DescribeEdgeTopNrtData(&cdn.DescribeEdgeTopNrtDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
		Domain:    &exampleDomain,
		Item:      "region",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}
