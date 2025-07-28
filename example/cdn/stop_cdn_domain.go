package cdn

import (
	"testing"

	"github.com/alihpurwandius/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func StopCdnDomain(t *testing.T) {
	resp, err := DefaultInstance.StopCdnDomain(&cdn.StopCdnDomainRequest{Domain: operateDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
