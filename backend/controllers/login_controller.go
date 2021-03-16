package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/team19/app/ent"
	"github.com/team19/app/ent/instructorinfo"
)

// LoginController defines the struct for the login controller
type LoginController struct {
	client *ent.Client
	router gin.IRouter
}

// CheckLogin handles GET requests to retrieve a instructorinfo entity
// @Summary Get a instructorinfo entity by Email and Password
// @Description get email by Email
// @Description get password by Password
// @ID check-login
// @Produce  json
// @Param email query string false "InstructorInfo Email"
// @Param password query string false "InstructorInfo Password"
// @Success 200 {array} ent.InstructorInfo
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /logins [get]
func (ctl *LoginController) CheckLogin(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")

	u, err := ctl.client.InstructorInfo.
		Query().
		WithDepartment().
		WithInstructorroom().
		WithTitle().
		Where(instructorinfo.EMAILContains(email),instructorinfo.PASSWORDContains(password)).
		All(context.Background())

	

	if err != nil {
		c.JSON(404, gin.H{
			"status": false,
			"error":  "Account not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data" : u,
	})
}

// NewLoginController creates and registers handles for the login controller
func NewLoginController(router gin.IRouter, client *ent.Client) *LoginController {
	uc := &LoginController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *LoginController) register() {
	login := ctl.router.Group("/logins")
	login.GET("", ctl.CheckLogin)
}