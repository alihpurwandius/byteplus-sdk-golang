package cdn

import (
	"testing"

	"github.com/alihpurwandius/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeEdgeNrtDataSummary(t *testing.T) {
	resp, err := DefaultInstance.DescribeEdgeNrtDataSummary(&cdn.DescribeEdgeNrtDataSummaryRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}
