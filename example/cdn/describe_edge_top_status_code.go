package cdn

import (
	"testing"

	"github.com/alihpurwandius/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeEdgeTopStatusCode(t *testing.T) {
	resp, err := DefaultInstance.DescribeEdgeTopStatusCode(&cdn.DescribeEdgeTopStatusCodeRequest{
		Metric: "status_5xx",
		Item:   "domain",
		Domain: &exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}
