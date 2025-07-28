package cdn

import (
	"testing"

	"github.com/alihpurwandius/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func SubmitUnblockTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitUnblockTask(&cdn.SubmitUnblockTaskRequest{
		Urls: exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.TaskID)
}
