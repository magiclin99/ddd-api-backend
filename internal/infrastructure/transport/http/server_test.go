package http

import (
	"dddapib/internal/domain/service"
	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServer(t *testing.T) {

	// just make sure no panic when creating server
	srv := NewServer(&service.Service{
		TaskService: nil,
	})
	assert.NotNil(t, srv)
}
