package cdn

import (
	"testing"

	"github.com/alihpurwandius/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func SubmitBlockTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitBlockTask(&cdn.SubmitBlockTaskRequest{
		Urls: exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.TaskID)
}
