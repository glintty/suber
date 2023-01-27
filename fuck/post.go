package fuck

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Post(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		print("amogus")
		return
	}

	if "deez" != c.Param("auth") {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fucker := uuid.NewString()

	c.SaveUploadedFile(file, "/home/glint/code/go/src/suber/shid/"+fucker+filepath.Ext(file.Filename))
	c.String(http.StatusAccepted, "i hope this isnt pornography! \nsuck my: localhost:8080"+fucker+filepath.Ext(file.Filename))

}
