package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team19/app/ent"
	"github.com/team19/app/ent/classdate"
)

// ClassdateController defines the struct for the classdate controller
type ClassdateController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateClassdate handles POST requests for adding classdate entities
// @Summary Create classdate
// @Description Create classdate
// @ID create-classdate
// @Accept   json
// @Produce  json
// @Param classdate body ent.Classdate true "Classdate entity"
// @Success 200 {object} ent.Classdate
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /classdates [post]
func (ctl *ClassdateController) CreateClassdate(c *gin.Context) {
	obj := ent.Classdate{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "classdate binding failed",
		})
		return
	}

	d, err := ctl.client.Classdate.
		Create().
		SetDAY(obj.DAY).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, d)
}

// GetClassdate handles GET requests to retrieve a classdate entity
// @Summary Get a classdate entity by ID
// @Description get classdate by ID
// @ID get-classdate
// @Produce  json
// @Param id path int true "Classdate ID"
// @Success 200 {object} ent.Classdate
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /classdates/{id} [get]
func (ctl *ClassdateController) GetClassdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	cl, err := ctl.client.Classdate.
		Query().
		Where(classdate.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cl)
}

// ListClassdate handles request to get a list of classdate entities
// @Summary List classdate entities
// @Description list classdate entities
// @ID list-classdate
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Classdate
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /classdates [get]
func (ctl *ClassdateController) ListClassdate(c *gin.Context) {
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

	classdates, err := ctl.client.Classdate.
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

	c.JSON(200, classdates)
}

// NewClassdateController creates and registers handles for the classdate controller
func NewClassdateController(router gin.IRouter, client *ent.Client) *ClassdateController {
	rc := &ClassdateController{
		client: client,
		router: router,
	}

	rc.register()

	return rc

}

func (ctl *ClassdateController) register() {
	classdates := ctl.router.Group("/classdates")

	classdates.POST("", ctl.CreateClassdate)
	classdates.GET(":id", ctl.GetClassdate)
	classdates.GET("", ctl.ListClassdate)

}
