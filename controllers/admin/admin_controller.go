package admin

import (
	"github.com/gin-gonic/gin"
	"mitchvillanueva.com/pulseid/domain/admin"
	"mitchvillanueva.com/pulseid/services"
	"log"
	"net/http"
)

// GetUser godoc
// @Summary Retrieves admin token based on username and password
// @Accept json
// @Produce json
// @Param admin body admin.Admin true "Admin Data"
// @Success 200 {object} admin.RequestToken
// @Router /login/ [post]
func Login(c *gin.Context) {
	request := &admin.Admin{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest,err)
		return
	}

	admin, err := services.AdminService.GenerateToken(*request)
	if err != nil {
		c.JSON(http.StatusBadRequest,err)
		return
	}
	c.JSON(http.StatusOK, admin)
}

// GetUser godoc
// @Summary Generatate a client invitation token
// @Produce json
// @Success 200 {object} client.Client
// @Router /requesttoken/ [post]
// @Security ApiKeyAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func RequestToken(c *gin.Context) {

	client, err := services.ClientService.GenerateToken()

	if err != nil {
		c.JSON(http.StatusBadRequest,err)
		return
	}

	c.JSON(http.StatusOK, client)
}