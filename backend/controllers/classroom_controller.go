package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team19/app/ent"
	"github.com/team19/app/ent/classroom"
)

// ClassroomController defines the struct for the classroom controller
type ClassroomController struct {
	client *ent.Client
	router gin.IRouter
}

// Classroom defines the struct for the classroom
type Classroom struct {
	ROOM     string

}

// CreateClassroom handles POST requests for adding classroom entities
// @Summary Create classroom
// @Description Create classroom
// @ID create-classroom
// @Accept   json
// @Produce  json
// @Param classroom body ent.Classroom true "Classroom entity"
// @Success 200 {object} ent.Classroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /classrooms [post]
func (ctl *ClassroomController) CreateClassroom(c *gin.Context) {
	obj := Classroom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "classroom binding failed",
		})
		return
	}

	d, err := ctl.client.Classroom.
		Create().
		SetROOM(obj.ROOM).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, d)
}

// GetClassroom handles GET requests to retrieve a classroom entity
// @Summary Get a classroom entity by ID
// @Description get classroom by ID
// @ID get-classroom
// @Produce  json
// @Param id path int true "Classroom ID"
// @Success 200 {object} ent.Classroom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /classrooms/{id} [get]
func (ctl *ClassroomController) GetClassroom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	cl, err := ctl.client.Classroom.
		Query().
		Where(classroom.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cl)
}

// ListClassroom handles request to get a list of classroom entities
// @Summary List classroom entities
// @Description list classroom entities
// @ID list-classroom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Classroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /classrooms [get]
func (ctl *ClassroomController) ListClassroom(c *gin.Context) {
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

	classrooms, err := ctl.client.Classroom.
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

	c.JSON(200, classrooms)
}

// NewClassroomController creates and registers handles for the classroom controller
func NewClassroomController(router gin.IRouter, client *ent.Client) *ClassroomController {
	rc := &ClassroomController{
		client: client,
		router: router,
	}

	rc.register()

	return rc

}

func (ctl *ClassroomController) register() {
	classrooms := ctl.router.Group("/classrooms")

	classrooms.POST("", ctl.CreateClassroom)
	classrooms.GET(":id", ctl.GetClassroom)
	classrooms.GET("", ctl.ListClassroom)

}
