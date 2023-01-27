package fuck

import (
	"path/filepath"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	filename := filepath.Base(c.Param("filename"))

	c.File("/home/glint/code/go/src/suberbutmoreworse/shid/" + filename)
}