package network_test

import (
	"testing"

	"github.com/shelepuginivan/hakutest/internal/pkg/network"
	"github.com/stretchr/testify/assert"
)

func TestGetLocalIP(t *testing.T) {
	ip, err := network.GetLocalIP()
	assert.NotEmpty(t, ip)
	assert.NoError(t, err)
}

func TestIsLocalIP(t *testing.T) {
	assert.True(t, network.IsLocalIP("127.0.0.1"))

	localIP, _ := network.GetLocalIP()
	assert.True(t, network.IsLocalIP(localIP))
}
