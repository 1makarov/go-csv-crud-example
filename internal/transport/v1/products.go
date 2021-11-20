package v1

import (
	"fmt"
	"github.com/1makarov/go-csv-crud-example/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initProductsRouter(api *gin.RouterGroup) {
	products := api.Group("/products")
	{
		products.GET("/", h.get)
		products.DELETE("/:id", h.delete)
		products.PUT("/:id", h.update)
		products.POST("/", h.create)
	}
}

// @Summary create
// @Tags products
// @ID create-product
// @Param input body types.CreateInput true " "
// @Success 200 {integer} integer 1
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/products/ [post]
func (h *Handler) create(c *gin.Context) {
	var input types.CreateInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	id, err := h.services.Products.Create(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, map[string]int64{
		"id": id,
	})
}

// @Summary get csv
// @Tags products
// @ID get-csv-products
// @Success 200 "OK"
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/products/ [get]
func (h *Handler) get(c *gin.Context) {
	products, err := h.services.Products.Get(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	headers := map[string]string{
		"Content-Disposition": `attachment; filename="csv.txt"`,
	}

	c.DataFromReader(http.StatusOK, -1, "text/html; charset=UTF-8", products, headers)
}

// @Summary delete
// @Tags products
// @ID delete-product
// @Param id path integer true " "
// @Success 200 "OK"
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/products/{id} [delete]
func (h *Handler) delete(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("empty id"))
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err = h.services.Products.Delete(c, int64(i)); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary update
// @Tags products
// @ID update-product
// @Param id path integer true " "
// @Param input body types.UpdateInput true " "
// @Success 200 "OK"
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/products/{id} [put]
func (h *Handler) update(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("empty id"))
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var input types.UpdateInput

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err = h.services.Products.Update(c, int64(i), input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusOK)
}
