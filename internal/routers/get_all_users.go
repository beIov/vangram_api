package routers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (r *Route) getAllUsers(c *gin.Context) {
	list, err := r.service.GetAllUsers(c)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}
