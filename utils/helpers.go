package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func RespondError(c *gin.Context, code int, msg string) {
	c.IndentedJSON(code, gin.H{"message": msg})
}

func ContainsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
