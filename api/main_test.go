package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// set gin to test mode to make tests logs cleaner to read 
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
