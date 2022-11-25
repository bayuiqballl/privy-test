package handlers

import (
	"database/sql"
	"net/http"
	"privy-test/entities"
	"privy-test/presenters"
	"privy-test/usecases"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type handlerHttpServer struct {
	app usecases.IUsecasesInterface
}

func NewHttpHandler(app usecases.IUsecasesInterface) handlerHttpServer {
	return handlerHttpServer{app: app}
}

func (h *handlerHttpServer) PostCake(c *gin.Context) {
	var req entities.CakeRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	if req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Title Cannot empty",
		})
		return
	}

	// validate image
	fileLogo, err := presenters.GetFileExtensionFromUrl(req.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	extLogo := strings.Split(presenters.FileExtentionLogo, ",")
	isValidLogo := presenters.ContainString(extLogo, fileLogo)
	if !isValidLogo {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "not format image",
		})
		return
	}

	err = h.app.CreateCake(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "created",
		"message": "success",
	})
}

func (h *handlerHttpServer) GetAllCakes(c *gin.Context) {
	sorting, _ := c.GetQuery("sortBy")

	res, err := h.app.GetAllCakes(sorting)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *handlerHttpServer) GetCakeDetailByID(c *gin.Context) {
	ID := c.Param("id")
	idInt, _ := strconv.Atoi(ID)

	res, err := h.app.GetCakeByID(idInt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "cake not found",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *handlerHttpServer) UpdateCakeByID(c *gin.Context) {
	var req entities.CakeRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	ID := c.Param("id")
	idInt, _ := strconv.Atoi(ID)

	if req.Image != "" {
		// validate image
		fileLogo, err := presenters.GetFileExtensionFromUrl(req.Image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		extLogo := strings.Split(presenters.FileExtentionLogo, ",")
		isValidLogo := presenters.ContainString(extLogo, fileLogo)
		if !isValidLogo {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "not format image",
			})
			return
		}
	}

	err = h.app.UpdateCakeByID(idInt, req)
	if err == presenters.ErrIdNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "updated",
		"message": "success",
	})
}

func (h *handlerHttpServer) DeleteCakeByID(c *gin.Context) {
	ID := c.Param("id")
	idInt, _ := strconv.Atoi(ID)

	err := h.app.DeleteCakeByID(idInt)
	if err == presenters.ErrIdNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}
