package pool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockQueueTaskPool_Start(t *testing.T) {

	testCases := []struct {
		name    string
		entity  interface{}
		wantErr error
	}{
		{
			name:    "正常启动",
			wantErr: nil,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := NewBlockQueueTaskPool(1, 1)
			err := b.Start()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
