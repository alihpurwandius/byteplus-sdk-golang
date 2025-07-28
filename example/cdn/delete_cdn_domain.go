package cdn

import (
	"testing"

	"github.com/alihpurwandius/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DeleteCdnDomain(t *testing.T) {
	resp, err := DefaultInstance.DeleteCdnDomain(&cdn.DeleteCdnDomainRequest{Domain: operateDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
