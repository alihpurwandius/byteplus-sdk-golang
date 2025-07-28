package cdn

import (
	"testing"

	"github.com/alihpurwandius/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeOriginTopNrtData(t *testing.T) {
	resp, err := DefaultInstance.DescribeOriginTopNrtData(&cdn.DescribeOriginTopNrtDataRequest{
		Metric: "flux",
		Item:   "domain",
		Domain: &exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}
