package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team19/app/ent"
	"github.com/team19/app/ent/degree"
	"github.com/team19/app/ent/subject"
	"github.com/team19/app/ent/subjectsoffered"
	"github.com/team19/app/ent/term"
	"github.com/team19/app/ent/year"
)

// SubjectsOfferedController defines the struct for the SubjectsOffered controller
type SubjectsOfferedController struct {
	client *ent.Client
	router gin.IRouter
}

// SubjectsOffered defines the struct for the SubjectsOffered
type SubjectsOffered struct {
	AMOUNT  string
	STATUS  string
	Remain  string
	Subject int
	Degree  int
	Year    int
	Term    int
}

// CreateSubjectsOffered handles POST requests for adding SubjectsOffered entities
// @Summary Create SubjectsOffered
// @Description Create SubjectsOffered
// @ID create-SubjectsOffered
// @Accept   json
// @Produce  json
// @Param subjectsoffered body ent.SubjectsOffered true "SubjectsOffered entity"
// @Success 200 {object} ent.SubjectsOffered
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /SubjectsOffereds [post]
func (ctl *SubjectsOfferedController) CreateSubjectsOffered(c *gin.Context) {
	obj := SubjectsOffered{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "SubjectsOffered binding failed",
		})
		return
	}
	y, err := ctl.client.Year.
		Query().
		Where(year.IDEQ(int(obj.Year))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "year not found",
		})
		return
	}
	t, err := ctl.client.Term.
		Query().
		Where(term.IDEQ(int(obj.Term))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "year not found",
		})
		return
	}
	deg, err := ctl.client.Degree.
		Query().
		Where(degree.IDEQ(int(obj.Degree))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "educationlevel not found",
		})
		return
	}
	s, err := ctl.client.Subject.
		Query().
		Where(subject.IDEQ(int(obj.Subject))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "subjects not found",
		})
		return
	}

	so, err := ctl.client.SubjectsOffered.
		Create().
		SetAMOUNT(obj.AMOUNT).
		SetSTATUS(obj.STATUS).
		SetRemain(obj.Remain).
		SetSubject(s).
		SetDegree(deg).
		SetYear(y).
		SetTerm(t).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   so,
	})
}

// GetSubjectsOffered handles GET requests to retrieve a SubjectsOffered entity
// @Summary Get a SubjectsOffered entity by ID
// @Description get SubjectsOffered by ID
// @ID get-SubjectsOffered
// @Produce  json
// @Param id path int true "SubjectsOffered ID"
// @Success 200 {array} ent.SubjectsOffered
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /SubjectsOffereds/{id} [get]
func (ctl *SubjectsOfferedController) GetSubjectsOffered(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	so, err := ctl.client.SubjectsOffered.
		Query().
		WithSubject().
		WithYear().
		WithDegree().
		WithTerm().
		Where(subjectsoffered.IDEQ(int(id))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, so)
}

// ListSubjectsOffered handles request to get a list of SubjectsOffered entities
// @Summary List SubjectsOffered entities
// @Description list SubjectsOffered entities
// @ID list-SubjectsOffered
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.SubjectsOffered
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /SubjectsOffereds [get]
func (ctl *SubjectsOfferedController) ListSubjectsOffered(c *gin.Context) {
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

	subjectsoffereds, err := ctl.client.SubjectsOffered.
		Query().
		WithSubject().
		WithYear().
		WithDegree().
		WithTerm().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, subjectsoffereds)
}

// DeleteSubjectsOffered handles DELETE requests to delete a SubjectsOffered entity
// @Summary Delete a SubjectsOffered entity by ID
// @Description get SubjectsOffered by ID
// @ID delete-SubjectsOffered
// @Produce  json
// @Param id path int true "SubjectsOffered ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /SubjectsOffereds/{id} [delete]
func (ctl *SubjectsOfferedController) DeleteSubjectsOffered(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.SubjectsOffered.
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

// UpdateSubjectsOffered handles PUT requests to update a SubjectsOffered entity
// @Summary Update a SubjectsOffered entity by ID
// @Description update SubjectsOffered by ID
// @ID update-SubjectsOffered
// @Accept   json
// @Produce  json
// @Param id path int true "SubjectsOffered ID"
// @Param SubjectsOffered body ent.SubjectsOffered true "SubjectsOffered entity"
// @Success 200 {object} ent.SubjectsOffered
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /SubjectsOffereds/{id} [put]
func (ctl *SubjectsOfferedController) UpdateSubjectsOffered(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.SubjectsOffered{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "SubjectsOffered binding failed",
		})
		return
	}
	obj.ID = int(id)
	so, err := ctl.client.SubjectsOffered.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, so)
}

// NewSubjectsOfferedController creates and registers handles for the SubjectsOffered controller
func NewSubjectsOfferedController(router gin.IRouter, client *ent.Client) *SubjectsOfferedController {
	soc := &SubjectsOfferedController{
		client: client,
		router: router,
	}
	soc.register()
	return soc
}

// InitSubjectsOfferedController registers routes to the main engine
func (ctl *SubjectsOfferedController) register() {
	SubjectsOffereds := ctl.router.Group("/SubjectsOffereds")

	SubjectsOffereds.GET("", ctl.ListSubjectsOffered)

	// CRUD
	SubjectsOffereds.POST("", ctl.CreateSubjectsOffered)
	SubjectsOffereds.GET(":id", ctl.GetSubjectsOffered)
	SubjectsOffereds.PUT(":id", ctl.UpdateSubjectsOffered)
	SubjectsOffereds.DELETE(":id", ctl.DeleteSubjectsOffered)
}
