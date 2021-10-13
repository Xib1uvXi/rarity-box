package appbackend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinServer struct {
	srv *srv
}

func NewGinServer(srv *srv) *GinServer {
	return &GinServer{srv: srv}
}

func (s *GinServer) Summoners(c *gin.Context) {
	address := c.Param("address")

	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"err_msg": "address is empty",
		})
		return
	}

	result, err := s.srv.Summoners(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"err_msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

func (s *GinServer) Tasks(c *gin.Context) {
	address := c.Param("address")

	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"err_msg": "address is empty",
		})
		return
	}

	result, err := s.srv.Tasks(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"err_msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

type IsOperatorReq struct {
	Address string
}

func (s *GinServer) IsOperator(c *gin.Context) {
	var req IsOperatorReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"err_msg": err.Error(),
		})
		return
	}

	result, err := s.srv.IsOperator(req.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"err_msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}
