package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team19/app/ent"
	"github.com/team19/app/ent/classtime"
)

// ClasstimeController defines the struct for the classtime controller
type ClasstimeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateClasstime handles POST requests for adding classtime entities
// @Summary Create classtime
// @Description Create classtime
// @ID create-classtime
// @Accept   json
// @Produce  json
// @Param classtime body ent.Classtime true "Classtime entity"
// @Success 200 {object} ent.Classtime
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /classtimes [post]
func (ctl *ClasstimeController) CreateClasstime(c *gin.Context) {
	obj := ent.Classtime{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "classtime binding failed",
		})
		return
	}

	d, err := ctl.client.Classtime.
		Create().
		SetTIME(obj.TIME).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, d)
}

// GetClasstime handles GET requests to retrieve a classtime entity
// @Summary Get a classtime entity by ID
// @Description get classtime by ID
// @ID get-classtime
// @Produce  json
// @Param id path int true "Classtime ID"
// @Success 200 {object} ent.Classtime
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /classtimes/{id} [get]
func (ctl *ClasstimeController) GetClasstime(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	cl, err := ctl.client.Classtime.
		Query().
		Where(classtime.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cl)
}

// ListClasstime handles request to get a list of classtime entities
// @Summary List classtime entities
// @Description list classtime entities
// @ID list-classtime
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Classtime
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /classtimes [get]
func (ctl *ClasstimeController) ListClasstime(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	classtimes, err := ctl.client.Classtime.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, classtimes)
}

// NewClasstimeController creates and registers handles for the classtime controller
func NewClasstimeController(router gin.IRouter, client *ent.Client) *ClasstimeController {
	rc := &ClasstimeController{
		client: client,
		router: router,
	}

	rc.register()

	return rc

}

func (ctl *ClasstimeController) register() {
	classtimes := ctl.router.Group("/classtimes")

	classtimes.POST("", ctl.CreateClasstime)
	classtimes.GET(":id", ctl.GetClasstime)
	classtimes.GET("", ctl.ListClasstime)

}
