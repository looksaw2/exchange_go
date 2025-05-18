package config_test

import (
	"testing"

	"github.com/looksaw/exchange_go/config"
	"github.com/stretchr/testify/require"
)

func TestInitConfig(t *testing.T) {
	config := config.InitConfig()
	require.Equal(t, config.InfoConfig.Version, "v1")
	require.Equal(t, config.InfoConfig.Author, "looksaw")
	require.Equal(t, config.InfoConfig.Description, "This is The info about exchange_go")
}
