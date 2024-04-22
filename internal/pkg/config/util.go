package config

import "github.com/gin-gonic/gin"

// ServerModeMap is a mapping server mode and its name.
var ServerModeMap = map[string]string{
	gin.ReleaseMode: "Release",
	gin.DebugMode:   "Debug",
	gin.TestMode:    "Test",
}
