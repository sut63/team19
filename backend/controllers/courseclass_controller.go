package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team19/app/ent"
	"github.com/team19/app/ent/classdate"
	"github.com/team19/app/ent/classroom"
	"github.com/team19/app/ent/classtime"
	"github.com/team19/app/ent/courseclass"
	"github.com/team19/app/ent/instructorinfo"
	"github.com/team19/app/ent/subject"
)

// CourseclassController defines the struct for the courseclass controller
type CourseclassController struct {
	client *ent.Client
	router gin.IRouter
}

// Courseclass defines the struct for the courseclass
type Courseclass struct {
	TableCode  string
	Day        int
	Time       int
	Instructor int
	Subject    int
	Room       int
}

// CreateCourseclass handles POST requests for adding courseclass entities
// @Summary Create courseclass
// @Description Create courseclass
// @ID create-courseclass
// @Accept   json
// @Produce  json
// @Param courseclass body ent.Courseclass true "Courseclass entity"
// @Success 200 {object} ent.Courseclass
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courseclasss [post]
func (ctl *CourseclassController) CreateCourseclass(c *gin.Context) {
	obj := Courseclass{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "courseclass binding failed",
		})
		return
	}

	d, err := ctl.client.Classdate.
		Query().
		Where(classdate.IDEQ(int(obj.Day))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Classdate not found",
		})
		return
	}

	ir, err := ctl.client.Classtime.
		Query().
		Where(classtime.IDEQ(int(obj.Time))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Classtime not found",
		})
		return
	}

	cr, err := ctl.client.Classroom.
		Query().
		Where(classroom.IDEQ(int(obj.Room))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Classroom not found",
		})
		return
	}

	t, err := ctl.client.InstructorInfo.
		Query().
		Where(instructorinfo.IDEQ(int(obj.Instructor))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "InstructorInfo not found",
		})
		return
	}

	a, err := ctl.client.Subject.
		Query().
		Where(subject.IDEQ(int(obj.Subject))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Subject not found",
		})
		return
	}

	u, err := ctl.client.Courseclass.
		Create().
		SetTablecode(obj.TableCode).
		SetClassdate(d).
		SetClasstime(ir).
		SetInstructorInfo(t).
		SetSubject(a).
		SetClassroom(cr).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   u,
	})
}

// GetCourseclass handles GET requests to retrieve a courseclass entity
// @Summary Get a courseclass entity by ID
// @Description get courseclass by ID
// @ID get-courseclass
// @Produce  json
// @Param id path int true "Courseclass ID"
// @Success 200 {array} ent.Courseclass
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courseclasss/{id} [get]
func (ctl *CourseclassController) GetCourseclass(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Courseclass.
		Query().
		WithClassdate().
		WithClasstime().
		WithInstructorInfo().
		WithSubject().
		WithClassroom().
		Where(courseclass.IDEQ(int(id))).
		All(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListCourseclass handles request to get a list of courseclass entities
// @Summary List courseclass entities
// @Description list courseclass entities
// @ID list-courseclass
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Courseclass
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courseclasss [get]
func (ctl *CourseclassController) ListCourseclass(c *gin.Context) {
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

	courseclasss, err := ctl.client.Courseclass.
		Query().
		WithClassdate().
		WithClasstime().
		WithInstructorInfo().
		WithSubject().
		WithClassroom().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, courseclasss)
}

// DeleteCourseclass handles DELETE requests to delete a courseclass entity
// @Summary Delete a courseclass entity by ID
// @Description get courseclass by ID
// @ID delete-courseclass
// @Produce  json
// @Param id path int true "Courseclass ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courseclasss/{id} [delete]
func (ctl *CourseclassController) DeleteCourseclass(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Courseclass.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateCourseclass handles PUT requests to update a courseclass entity
// @Summary Update a courseclass entity by ID
// @Description update courseclass by ID
// @ID update-courseclass
// @Accept   json
// @Produce  json
// @Param id path int true "Courseclass ID"
// @Param courseclass body ent.Courseclass true "Courseclass entity"
// @Success 200 {object} ent.Courseclass
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courseclasss/{id} [put]
func (ctl *CourseclassController) UpdateCourseclass(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Courseclass{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "courseclass binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.Courseclass.
		UpdateOneID(int(id)).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, u)
}

// NewCourseclassController creates and registers handles for the courseclass controller
func NewCourseclassController(router gin.IRouter, client *ent.Client) *CourseclassController {
	uc := &CourseclassController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *CourseclassController) register() {
	courseclasss := ctl.router.Group("/courseclasss")

	courseclasss.GET("", ctl.ListCourseclass)

	// CRUD
	courseclasss.POST("", ctl.CreateCourseclass)
	courseclasss.GET(":id", ctl.GetCourseclass)
	courseclasss.PUT(":id", ctl.UpdateCourseclass)
	courseclasss.DELETE(":id", ctl.DeleteCourseclass)
}
