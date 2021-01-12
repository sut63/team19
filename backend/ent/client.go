// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/team19/app/ent/migrate"

	"github.com/team19/app/ent/course"
	"github.com/team19/app/ent/degree"
	"github.com/team19/app/ent/department"
	"github.com/team19/app/ent/instructorinfo"
	"github.com/team19/app/ent/instructorroom"
	"github.com/team19/app/ent/subject"
	"github.com/team19/app/ent/title"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Course is the client for interacting with the Course builders.
	Course *CourseClient
	// Degree is the client for interacting with the Degree builders.
	Degree *DegreeClient
	// Department is the client for interacting with the Department builders.
	Department *DepartmentClient
	// InstructorInfo is the client for interacting with the InstructorInfo builders.
	InstructorInfo *InstructorInfoClient
	// InstructorRoom is the client for interacting with the InstructorRoom builders.
	InstructorRoom *InstructorRoomClient
	// Subject is the client for interacting with the Subject builders.
	Subject *SubjectClient
	// Title is the client for interacting with the Title builders.
	Title *TitleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Course = NewCourseClient(c.config)
	c.Degree = NewDegreeClient(c.config)
	c.Department = NewDepartmentClient(c.config)
	c.InstructorInfo = NewInstructorInfoClient(c.config)
	c.InstructorRoom = NewInstructorRoomClient(c.config)
	c.Subject = NewSubjectClient(c.config)
	c.Title = NewTitleClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Course:         NewCourseClient(cfg),
		Degree:         NewDegreeClient(cfg),
		Department:     NewDepartmentClient(cfg),
		InstructorInfo: NewInstructorInfoClient(cfg),
		InstructorRoom: NewInstructorRoomClient(cfg),
		Subject:        NewSubjectClient(cfg),
		Title:          NewTitleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:         cfg,
		Course:         NewCourseClient(cfg),
		Degree:         NewDegreeClient(cfg),
		Department:     NewDepartmentClient(cfg),
		InstructorInfo: NewInstructorInfoClient(cfg),
		InstructorRoom: NewInstructorRoomClient(cfg),
		Subject:        NewSubjectClient(cfg),
		Title:          NewTitleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Course.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Course.Use(hooks...)
	c.Degree.Use(hooks...)
	c.Department.Use(hooks...)
	c.InstructorInfo.Use(hooks...)
	c.InstructorRoom.Use(hooks...)
	c.Subject.Use(hooks...)
	c.Title.Use(hooks...)
}

// CourseClient is a client for the Course schema.
type CourseClient struct {
	config
}

// NewCourseClient returns a client for the Course from the given config.
func NewCourseClient(c config) *CourseClient {
	return &CourseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `course.Hooks(f(g(h())))`.
func (c *CourseClient) Use(hooks ...Hook) {
	c.hooks.Course = append(c.hooks.Course, hooks...)
}

// Create returns a create builder for Course.
func (c *CourseClient) Create() *CourseCreate {
	mutation := newCourseMutation(c.config, OpCreate)
	return &CourseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Course.
func (c *CourseClient) Update() *CourseUpdate {
	mutation := newCourseMutation(c.config, OpUpdate)
	return &CourseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CourseClient) UpdateOne(co *Course) *CourseUpdateOne {
	mutation := newCourseMutation(c.config, OpUpdateOne, withCourse(co))
	return &CourseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CourseClient) UpdateOneID(id int) *CourseUpdateOne {
	mutation := newCourseMutation(c.config, OpUpdateOne, withCourseID(id))
	return &CourseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Course.
func (c *CourseClient) Delete() *CourseDelete {
	mutation := newCourseMutation(c.config, OpDelete)
	return &CourseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CourseClient) DeleteOne(co *Course) *CourseDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CourseClient) DeleteOneID(id int) *CourseDeleteOne {
	builder := c.Delete().Where(course.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CourseDeleteOne{builder}
}

// Create returns a query builder for Course.
func (c *CourseClient) Query() *CourseQuery {
	return &CourseQuery{config: c.config}
}

// Get returns a Course entity by its id.
func (c *CourseClient) Get(ctx context.Context, id int) (*Course, error) {
	return c.Query().Where(course.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CourseClient) GetX(ctx context.Context, id int) *Course {
	co, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return co
}

// QueryInstructorInfoID queries the InstructorInfo_id edge of a Course.
func (c *CourseClient) QueryInstructorInfoID(co *Course) *InstructorInfoQuery {
	query := &InstructorInfoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(course.Table, course.FieldID, id),
			sqlgraph.To(instructorinfo.Table, instructorinfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, course.InstructorInfoIDTable, course.InstructorInfoIDColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDepartmentID queries the Department_id edge of a Course.
func (c *CourseClient) QueryDepartmentID(co *Course) *DepartmentQuery {
	query := &DepartmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(course.Table, course.FieldID, id),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, course.DepartmentIDTable, course.DepartmentIDColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDegreeID queries the Degree_id edge of a Course.
func (c *CourseClient) QueryDegreeID(co *Course) *DegreeQuery {
	query := &DegreeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(course.Table, course.FieldID, id),
			sqlgraph.To(degree.Table, degree.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, course.DegreeIDTable, course.DegreeIDColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySubjectID queries the Subject_id edge of a Course.
func (c *CourseClient) QuerySubjectID(co *Course) *SubjectQuery {
	query := &SubjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(course.Table, course.FieldID, id),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, course.SubjectIDTable, course.SubjectIDColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CourseClient) Hooks() []Hook {
	return c.hooks.Course
}

// DegreeClient is a client for the Degree schema.
type DegreeClient struct {
	config
}

// NewDegreeClient returns a client for the Degree from the given config.
func NewDegreeClient(c config) *DegreeClient {
	return &DegreeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `degree.Hooks(f(g(h())))`.
func (c *DegreeClient) Use(hooks ...Hook) {
	c.hooks.Degree = append(c.hooks.Degree, hooks...)
}

// Create returns a create builder for Degree.
func (c *DegreeClient) Create() *DegreeCreate {
	mutation := newDegreeMutation(c.config, OpCreate)
	return &DegreeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Degree.
func (c *DegreeClient) Update() *DegreeUpdate {
	mutation := newDegreeMutation(c.config, OpUpdate)
	return &DegreeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DegreeClient) UpdateOne(d *Degree) *DegreeUpdateOne {
	mutation := newDegreeMutation(c.config, OpUpdateOne, withDegree(d))
	return &DegreeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DegreeClient) UpdateOneID(id int) *DegreeUpdateOne {
	mutation := newDegreeMutation(c.config, OpUpdateOne, withDegreeID(id))
	return &DegreeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Degree.
func (c *DegreeClient) Delete() *DegreeDelete {
	mutation := newDegreeMutation(c.config, OpDelete)
	return &DegreeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DegreeClient) DeleteOne(d *Degree) *DegreeDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DegreeClient) DeleteOneID(id int) *DegreeDeleteOne {
	builder := c.Delete().Where(degree.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DegreeDeleteOne{builder}
}

// Create returns a query builder for Degree.
func (c *DegreeClient) Query() *DegreeQuery {
	return &DegreeQuery{config: c.config}
}

// Get returns a Degree entity by its id.
func (c *DegreeClient) Get(ctx context.Context, id int) (*Degree, error) {
	return c.Query().Where(degree.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DegreeClient) GetX(ctx context.Context, id int) *Degree {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryDegree queries the degree edge of a Degree.
func (c *DegreeClient) QueryDegree(d *Degree) *CourseQuery {
	query := &CourseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(degree.Table, degree.FieldID, id),
			sqlgraph.To(course.Table, course.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, degree.DegreeTable, degree.DegreeColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DegreeClient) Hooks() []Hook {
	return c.hooks.Degree
}

// DepartmentClient is a client for the Department schema.
type DepartmentClient struct {
	config
}

// NewDepartmentClient returns a client for the Department from the given config.
func NewDepartmentClient(c config) *DepartmentClient {
	return &DepartmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `department.Hooks(f(g(h())))`.
func (c *DepartmentClient) Use(hooks ...Hook) {
	c.hooks.Department = append(c.hooks.Department, hooks...)
}

// Create returns a create builder for Department.
func (c *DepartmentClient) Create() *DepartmentCreate {
	mutation := newDepartmentMutation(c.config, OpCreate)
	return &DepartmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Department.
func (c *DepartmentClient) Update() *DepartmentUpdate {
	mutation := newDepartmentMutation(c.config, OpUpdate)
	return &DepartmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DepartmentClient) UpdateOne(d *Department) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartment(d))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DepartmentClient) UpdateOneID(id int) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartmentID(id))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Department.
func (c *DepartmentClient) Delete() *DepartmentDelete {
	mutation := newDepartmentMutation(c.config, OpDelete)
	return &DepartmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DepartmentClient) DeleteOne(d *Department) *DepartmentDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DepartmentClient) DeleteOneID(id int) *DepartmentDeleteOne {
	builder := c.Delete().Where(department.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DepartmentDeleteOne{builder}
}

// Create returns a query builder for Department.
func (c *DepartmentClient) Query() *DepartmentQuery {
	return &DepartmentQuery{config: c.config}
}

// Get returns a Department entity by its id.
func (c *DepartmentClient) Get(ctx context.Context, id int) (*Department, error) {
	return c.Query().Where(department.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DepartmentClient) GetX(ctx context.Context, id int) *Department {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryInstructorinfos queries the instructorinfos edge of a Department.
func (c *DepartmentClient) QueryInstructorinfos(d *Department) *InstructorInfoQuery {
	query := &InstructorInfoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, id),
			sqlgraph.To(instructorinfo.Table, instructorinfo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.InstructorinfosTable, department.InstructorinfosColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDepartment queries the department edge of a Department.
func (c *DepartmentClient) QueryDepartment(d *Department) *CourseQuery {
	query := &CourseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, id),
			sqlgraph.To(course.Table, course.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.DepartmentTable, department.DepartmentColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DepartmentClient) Hooks() []Hook {
	return c.hooks.Department
}

// InstructorInfoClient is a client for the InstructorInfo schema.
type InstructorInfoClient struct {
	config
}

// NewInstructorInfoClient returns a client for the InstructorInfo from the given config.
func NewInstructorInfoClient(c config) *InstructorInfoClient {
	return &InstructorInfoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `instructorinfo.Hooks(f(g(h())))`.
func (c *InstructorInfoClient) Use(hooks ...Hook) {
	c.hooks.InstructorInfo = append(c.hooks.InstructorInfo, hooks...)
}

// Create returns a create builder for InstructorInfo.
func (c *InstructorInfoClient) Create() *InstructorInfoCreate {
	mutation := newInstructorInfoMutation(c.config, OpCreate)
	return &InstructorInfoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for InstructorInfo.
func (c *InstructorInfoClient) Update() *InstructorInfoUpdate {
	mutation := newInstructorInfoMutation(c.config, OpUpdate)
	return &InstructorInfoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *InstructorInfoClient) UpdateOne(ii *InstructorInfo) *InstructorInfoUpdateOne {
	mutation := newInstructorInfoMutation(c.config, OpUpdateOne, withInstructorInfo(ii))
	return &InstructorInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *InstructorInfoClient) UpdateOneID(id int) *InstructorInfoUpdateOne {
	mutation := newInstructorInfoMutation(c.config, OpUpdateOne, withInstructorInfoID(id))
	return &InstructorInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for InstructorInfo.
func (c *InstructorInfoClient) Delete() *InstructorInfoDelete {
	mutation := newInstructorInfoMutation(c.config, OpDelete)
	return &InstructorInfoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *InstructorInfoClient) DeleteOne(ii *InstructorInfo) *InstructorInfoDeleteOne {
	return c.DeleteOneID(ii.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *InstructorInfoClient) DeleteOneID(id int) *InstructorInfoDeleteOne {
	builder := c.Delete().Where(instructorinfo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &InstructorInfoDeleteOne{builder}
}

// Create returns a query builder for InstructorInfo.
func (c *InstructorInfoClient) Query() *InstructorInfoQuery {
	return &InstructorInfoQuery{config: c.config}
}

// Get returns a InstructorInfo entity by its id.
func (c *InstructorInfoClient) Get(ctx context.Context, id int) (*InstructorInfo, error) {
	return c.Query().Where(instructorinfo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InstructorInfoClient) GetX(ctx context.Context, id int) *InstructorInfo {
	ii, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ii
}

// QueryTitle queries the title edge of a InstructorInfo.
func (c *InstructorInfoClient) QueryTitle(ii *InstructorInfo) *TitleQuery {
	query := &TitleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ii.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(instructorinfo.Table, instructorinfo.FieldID, id),
			sqlgraph.To(title.Table, title.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, instructorinfo.TitleTable, instructorinfo.TitleColumn),
		)
		fromV = sqlgraph.Neighbors(ii.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryInstructorroom queries the instructorroom edge of a InstructorInfo.
func (c *InstructorInfoClient) QueryInstructorroom(ii *InstructorInfo) *InstructorRoomQuery {
	query := &InstructorRoomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ii.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(instructorinfo.Table, instructorinfo.FieldID, id),
			sqlgraph.To(instructorroom.Table, instructorroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, instructorinfo.InstructorroomTable, instructorinfo.InstructorroomColumn),
		)
		fromV = sqlgraph.Neighbors(ii.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDepartment queries the department edge of a InstructorInfo.
func (c *InstructorInfoClient) QueryDepartment(ii *InstructorInfo) *DepartmentQuery {
	query := &DepartmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ii.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(instructorinfo.Table, instructorinfo.FieldID, id),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, instructorinfo.DepartmentTable, instructorinfo.DepartmentColumn),
		)
		fromV = sqlgraph.Neighbors(ii.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryInstructor queries the instructor edge of a InstructorInfo.
func (c *InstructorInfoClient) QueryInstructor(ii *InstructorInfo) *CourseQuery {
	query := &CourseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ii.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(instructorinfo.Table, instructorinfo.FieldID, id),
			sqlgraph.To(course.Table, course.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, instructorinfo.InstructorTable, instructorinfo.InstructorColumn),
		)
		fromV = sqlgraph.Neighbors(ii.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *InstructorInfoClient) Hooks() []Hook {
	return c.hooks.InstructorInfo
}

// InstructorRoomClient is a client for the InstructorRoom schema.
type InstructorRoomClient struct {
	config
}

// NewInstructorRoomClient returns a client for the InstructorRoom from the given config.
func NewInstructorRoomClient(c config) *InstructorRoomClient {
	return &InstructorRoomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `instructorroom.Hooks(f(g(h())))`.
func (c *InstructorRoomClient) Use(hooks ...Hook) {
	c.hooks.InstructorRoom = append(c.hooks.InstructorRoom, hooks...)
}

// Create returns a create builder for InstructorRoom.
func (c *InstructorRoomClient) Create() *InstructorRoomCreate {
	mutation := newInstructorRoomMutation(c.config, OpCreate)
	return &InstructorRoomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for InstructorRoom.
func (c *InstructorRoomClient) Update() *InstructorRoomUpdate {
	mutation := newInstructorRoomMutation(c.config, OpUpdate)
	return &InstructorRoomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *InstructorRoomClient) UpdateOne(ir *InstructorRoom) *InstructorRoomUpdateOne {
	mutation := newInstructorRoomMutation(c.config, OpUpdateOne, withInstructorRoom(ir))
	return &InstructorRoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *InstructorRoomClient) UpdateOneID(id int) *InstructorRoomUpdateOne {
	mutation := newInstructorRoomMutation(c.config, OpUpdateOne, withInstructorRoomID(id))
	return &InstructorRoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for InstructorRoom.
func (c *InstructorRoomClient) Delete() *InstructorRoomDelete {
	mutation := newInstructorRoomMutation(c.config, OpDelete)
	return &InstructorRoomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *InstructorRoomClient) DeleteOne(ir *InstructorRoom) *InstructorRoomDeleteOne {
	return c.DeleteOneID(ir.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *InstructorRoomClient) DeleteOneID(id int) *InstructorRoomDeleteOne {
	builder := c.Delete().Where(instructorroom.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &InstructorRoomDeleteOne{builder}
}

// Create returns a query builder for InstructorRoom.
func (c *InstructorRoomClient) Query() *InstructorRoomQuery {
	return &InstructorRoomQuery{config: c.config}
}

// Get returns a InstructorRoom entity by its id.
func (c *InstructorRoomClient) Get(ctx context.Context, id int) (*InstructorRoom, error) {
	return c.Query().Where(instructorroom.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InstructorRoomClient) GetX(ctx context.Context, id int) *InstructorRoom {
	ir, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ir
}

// QueryInstructorinfos queries the instructorinfos edge of a InstructorRoom.
func (c *InstructorRoomClient) QueryInstructorinfos(ir *InstructorRoom) *InstructorInfoQuery {
	query := &InstructorInfoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ir.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(instructorroom.Table, instructorroom.FieldID, id),
			sqlgraph.To(instructorinfo.Table, instructorinfo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, instructorroom.InstructorinfosTable, instructorroom.InstructorinfosColumn),
		)
		fromV = sqlgraph.Neighbors(ir.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *InstructorRoomClient) Hooks() []Hook {
	return c.hooks.InstructorRoom
}

// SubjectClient is a client for the Subject schema.
type SubjectClient struct {
	config
}

// NewSubjectClient returns a client for the Subject from the given config.
func NewSubjectClient(c config) *SubjectClient {
	return &SubjectClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `subject.Hooks(f(g(h())))`.
func (c *SubjectClient) Use(hooks ...Hook) {
	c.hooks.Subject = append(c.hooks.Subject, hooks...)
}

// Create returns a create builder for Subject.
func (c *SubjectClient) Create() *SubjectCreate {
	mutation := newSubjectMutation(c.config, OpCreate)
	return &SubjectCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Subject.
func (c *SubjectClient) Update() *SubjectUpdate {
	mutation := newSubjectMutation(c.config, OpUpdate)
	return &SubjectUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubjectClient) UpdateOne(s *Subject) *SubjectUpdateOne {
	mutation := newSubjectMutation(c.config, OpUpdateOne, withSubject(s))
	return &SubjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SubjectClient) UpdateOneID(id int) *SubjectUpdateOne {
	mutation := newSubjectMutation(c.config, OpUpdateOne, withSubjectID(id))
	return &SubjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Subject.
func (c *SubjectClient) Delete() *SubjectDelete {
	mutation := newSubjectMutation(c.config, OpDelete)
	return &SubjectDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SubjectClient) DeleteOne(s *Subject) *SubjectDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SubjectClient) DeleteOneID(id int) *SubjectDeleteOne {
	builder := c.Delete().Where(subject.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SubjectDeleteOne{builder}
}

// Create returns a query builder for Subject.
func (c *SubjectClient) Query() *SubjectQuery {
	return &SubjectQuery{config: c.config}
}

// Get returns a Subject entity by its id.
func (c *SubjectClient) Get(ctx context.Context, id int) (*Subject, error) {
	return c.Query().Where(subject.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubjectClient) GetX(ctx context.Context, id int) *Subject {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// QuerySubject queries the subject edge of a Subject.
func (c *SubjectClient) QuerySubject(s *Subject) *CourseQuery {
	query := &CourseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(subject.Table, subject.FieldID, id),
			sqlgraph.To(course.Table, course.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, subject.SubjectTable, subject.SubjectColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SubjectClient) Hooks() []Hook {
	return c.hooks.Subject
}

// TitleClient is a client for the Title schema.
type TitleClient struct {
	config
}

// NewTitleClient returns a client for the Title from the given config.
func NewTitleClient(c config) *TitleClient {
	return &TitleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `title.Hooks(f(g(h())))`.
func (c *TitleClient) Use(hooks ...Hook) {
	c.hooks.Title = append(c.hooks.Title, hooks...)
}

// Create returns a create builder for Title.
func (c *TitleClient) Create() *TitleCreate {
	mutation := newTitleMutation(c.config, OpCreate)
	return &TitleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Title.
func (c *TitleClient) Update() *TitleUpdate {
	mutation := newTitleMutation(c.config, OpUpdate)
	return &TitleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TitleClient) UpdateOne(t *Title) *TitleUpdateOne {
	mutation := newTitleMutation(c.config, OpUpdateOne, withTitle(t))
	return &TitleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TitleClient) UpdateOneID(id int) *TitleUpdateOne {
	mutation := newTitleMutation(c.config, OpUpdateOne, withTitleID(id))
	return &TitleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Title.
func (c *TitleClient) Delete() *TitleDelete {
	mutation := newTitleMutation(c.config, OpDelete)
	return &TitleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TitleClient) DeleteOne(t *Title) *TitleDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TitleClient) DeleteOneID(id int) *TitleDeleteOne {
	builder := c.Delete().Where(title.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TitleDeleteOne{builder}
}

// Create returns a query builder for Title.
func (c *TitleClient) Query() *TitleQuery {
	return &TitleQuery{config: c.config}
}

// Get returns a Title entity by its id.
func (c *TitleClient) Get(ctx context.Context, id int) (*Title, error) {
	return c.Query().Where(title.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TitleClient) GetX(ctx context.Context, id int) *Title {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// QueryInstructorinfos queries the instructorinfos edge of a Title.
func (c *TitleClient) QueryInstructorinfos(t *Title) *InstructorInfoQuery {
	query := &InstructorInfoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(title.Table, title.FieldID, id),
			sqlgraph.To(instructorinfo.Table, instructorinfo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, title.InstructorinfosTable, title.InstructorinfosColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TitleClient) Hooks() []Hook {
	return c.hooks.Title
}
