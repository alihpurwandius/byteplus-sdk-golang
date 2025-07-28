package cdn

import (
	"testing"

	"github.com/alihpurwandius/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func UntagResources(t *testing.T) {
	resp, err := DefaultInstance.UntagResources(&cdn.UntagResourcesRequest{
		ResourceType: "Domain",
		ResourceIds:  []string{operateDomain},
		TagKeys:      []string{"k1"},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
