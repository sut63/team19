// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team19/app/ent/courseclass"
	"github.com/team19/app/ent/department"
	"github.com/team19/app/ent/instructorinfo"
	"github.com/team19/app/ent/instructorroom"
	"github.com/team19/app/ent/predicate"
	"github.com/team19/app/ent/title"
)

// InstructorInfoQuery is the builder for querying InstructorInfo entities.
type InstructorInfoQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.InstructorInfo
	// eager-loading edges.
	withTitle          *TitleQuery
	withInstructorroom *InstructorRoomQuery
	withDepartment     *DepartmentQuery
	withCourseclasses  *CourseclassQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (iiq *InstructorInfoQuery) Where(ps ...predicate.InstructorInfo) *InstructorInfoQuery {
	iiq.predicates = append(iiq.predicates, ps...)
	return iiq
}

// Limit adds a limit step to the query.
func (iiq *InstructorInfoQuery) Limit(limit int) *InstructorInfoQuery {
	iiq.limit = &limit
	return iiq
}

// Offset adds an offset step to the query.
func (iiq *InstructorInfoQuery) Offset(offset int) *InstructorInfoQuery {
	iiq.offset = &offset
	return iiq
}

// Order adds an order step to the query.
func (iiq *InstructorInfoQuery) Order(o ...OrderFunc) *InstructorInfoQuery {
	iiq.order = append(iiq.order, o...)
	return iiq
}

// QueryTitle chains the current query on the title edge.
func (iiq *InstructorInfoQuery) QueryTitle() *TitleQuery {
	query := &TitleQuery{config: iiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instructorinfo.Table, instructorinfo.FieldID, iiq.sqlQuery()),
			sqlgraph.To(title.Table, title.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, instructorinfo.TitleTable, instructorinfo.TitleColumn),
		)
		fromU = sqlgraph.SetNeighbors(iiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInstructorroom chains the current query on the instructorroom edge.
func (iiq *InstructorInfoQuery) QueryInstructorroom() *InstructorRoomQuery {
	query := &InstructorRoomQuery{config: iiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instructorinfo.Table, instructorinfo.FieldID, iiq.sqlQuery()),
			sqlgraph.To(instructorroom.Table, instructorroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, instructorinfo.InstructorroomTable, instructorinfo.InstructorroomColumn),
		)
		fromU = sqlgraph.SetNeighbors(iiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDepartment chains the current query on the department edge.
func (iiq *InstructorInfoQuery) QueryDepartment() *DepartmentQuery {
	query := &DepartmentQuery{config: iiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instructorinfo.Table, instructorinfo.FieldID, iiq.sqlQuery()),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, instructorinfo.DepartmentTable, instructorinfo.DepartmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(iiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCourseclasses chains the current query on the courseclasses edge.
func (iiq *InstructorInfoQuery) QueryCourseclasses() *CourseclassQuery {
	query := &CourseclassQuery{config: iiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instructorinfo.Table, instructorinfo.FieldID, iiq.sqlQuery()),
			sqlgraph.To(courseclass.Table, courseclass.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, instructorinfo.CourseclassesTable, instructorinfo.CourseclassesColumn),
		)
		fromU = sqlgraph.SetNeighbors(iiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first InstructorInfo entity in the query. Returns *NotFoundError when no instructorinfo was found.
func (iiq *InstructorInfoQuery) First(ctx context.Context) (*InstructorInfo, error) {
	iis, err := iiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(iis) == 0 {
		return nil, &NotFoundError{instructorinfo.Label}
	}
	return iis[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iiq *InstructorInfoQuery) FirstX(ctx context.Context) *InstructorInfo {
	ii, err := iiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ii
}

// FirstID returns the first InstructorInfo id in the query. Returns *NotFoundError when no id was found.
func (iiq *InstructorInfoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{instructorinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (iiq *InstructorInfoQuery) FirstXID(ctx context.Context) int {
	id, err := iiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only InstructorInfo entity in the query, returns an error if not exactly one entity was returned.
func (iiq *InstructorInfoQuery) Only(ctx context.Context) (*InstructorInfo, error) {
	iis, err := iiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(iis) {
	case 1:
		return iis[0], nil
	case 0:
		return nil, &NotFoundError{instructorinfo.Label}
	default:
		return nil, &NotSingularError{instructorinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iiq *InstructorInfoQuery) OnlyX(ctx context.Context) *InstructorInfo {
	ii, err := iiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ii
}

// OnlyID returns the only InstructorInfo id in the query, returns an error if not exactly one id was returned.
func (iiq *InstructorInfoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{instructorinfo.Label}
	default:
		err = &NotSingularError{instructorinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iiq *InstructorInfoQuery) OnlyIDX(ctx context.Context) int {
	id, err := iiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InstructorInfos.
func (iiq *InstructorInfoQuery) All(ctx context.Context) ([]*InstructorInfo, error) {
	if err := iiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iiq *InstructorInfoQuery) AllX(ctx context.Context) []*InstructorInfo {
	iis, err := iiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return iis
}

// IDs executes the query and returns a list of InstructorInfo ids.
func (iiq *InstructorInfoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := iiq.Select(instructorinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iiq *InstructorInfoQuery) IDsX(ctx context.Context) []int {
	ids, err := iiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iiq *InstructorInfoQuery) Count(ctx context.Context) (int, error) {
	if err := iiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iiq *InstructorInfoQuery) CountX(ctx context.Context) int {
	count, err := iiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iiq *InstructorInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := iiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iiq *InstructorInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := iiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iiq *InstructorInfoQuery) Clone() *InstructorInfoQuery {
	return &InstructorInfoQuery{
		config:     iiq.config,
		limit:      iiq.limit,
		offset:     iiq.offset,
		order:      append([]OrderFunc{}, iiq.order...),
		unique:     append([]string{}, iiq.unique...),
		predicates: append([]predicate.InstructorInfo{}, iiq.predicates...),
		// clone intermediate query.
		sql:  iiq.sql.Clone(),
		path: iiq.path,
	}
}

//  WithTitle tells the query-builder to eager-loads the nodes that are connected to
// the "title" edge. The optional arguments used to configure the query builder of the edge.
func (iiq *InstructorInfoQuery) WithTitle(opts ...func(*TitleQuery)) *InstructorInfoQuery {
	query := &TitleQuery{config: iiq.config}
	for _, opt := range opts {
		opt(query)
	}
	iiq.withTitle = query
	return iiq
}

//  WithInstructorroom tells the query-builder to eager-loads the nodes that are connected to
// the "instructorroom" edge. The optional arguments used to configure the query builder of the edge.
func (iiq *InstructorInfoQuery) WithInstructorroom(opts ...func(*InstructorRoomQuery)) *InstructorInfoQuery {
	query := &InstructorRoomQuery{config: iiq.config}
	for _, opt := range opts {
		opt(query)
	}
	iiq.withInstructorroom = query
	return iiq
}

//  WithDepartment tells the query-builder to eager-loads the nodes that are connected to
// the "department" edge. The optional arguments used to configure the query builder of the edge.
func (iiq *InstructorInfoQuery) WithDepartment(opts ...func(*DepartmentQuery)) *InstructorInfoQuery {
	query := &DepartmentQuery{config: iiq.config}
	for _, opt := range opts {
		opt(query)
	}
	iiq.withDepartment = query
	return iiq
}

//  WithCourseclasses tells the query-builder to eager-loads the nodes that are connected to
// the "courseclasses" edge. The optional arguments used to configure the query builder of the edge.
func (iiq *InstructorInfoQuery) WithCourseclasses(opts ...func(*CourseclassQuery)) *InstructorInfoQuery {
	query := &CourseclassQuery{config: iiq.config}
	for _, opt := range opts {
		opt(query)
	}
	iiq.withCourseclasses = query
	return iiq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NAME string `json:"NAME,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.InstructorInfo.Query().
//		GroupBy(instructorinfo.FieldNAME).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (iiq *InstructorInfoQuery) GroupBy(field string, fields ...string) *InstructorInfoGroupBy {
	group := &InstructorInfoGroupBy{config: iiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iiq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		NAME string `json:"NAME,omitempty"`
//	}
//
//	client.InstructorInfo.Query().
//		Select(instructorinfo.FieldNAME).
//		Scan(ctx, &v)
//
func (iiq *InstructorInfoQuery) Select(field string, fields ...string) *InstructorInfoSelect {
	selector := &InstructorInfoSelect{config: iiq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iiq.sqlQuery(), nil
	}
	return selector
}

func (iiq *InstructorInfoQuery) prepareQuery(ctx context.Context) error {
	if iiq.path != nil {
		prev, err := iiq.path(ctx)
		if err != nil {
			return err
		}
		iiq.sql = prev
	}
	return nil
}

func (iiq *InstructorInfoQuery) sqlAll(ctx context.Context) ([]*InstructorInfo, error) {
	var (
		nodes       = []*InstructorInfo{}
		withFKs     = iiq.withFKs
		_spec       = iiq.querySpec()
		loadedTypes = [4]bool{
			iiq.withTitle != nil,
			iiq.withInstructorroom != nil,
			iiq.withDepartment != nil,
			iiq.withCourseclasses != nil,
		}
	)
	if iiq.withTitle != nil || iiq.withInstructorroom != nil || iiq.withDepartment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, instructorinfo.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &InstructorInfo{config: iiq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, iiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := iiq.withTitle; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*InstructorInfo)
		for i := range nodes {
			if fk := nodes[i].title_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(title.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "title_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Title = n
			}
		}
	}

	if query := iiq.withInstructorroom; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*InstructorInfo)
		for i := range nodes {
			if fk := nodes[i].instructorroom_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(instructorroom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "instructorroom_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Instructorroom = n
			}
		}
	}

	if query := iiq.withDepartment; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*InstructorInfo)
		for i := range nodes {
			if fk := nodes[i].department_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(department.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "department_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Department = n
			}
		}
	}

	if query := iiq.withCourseclasses; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*InstructorInfo)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Courseclass(func(s *sql.Selector) {
			s.Where(sql.InValues(instructorinfo.CourseclassesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.InstructorInfo_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "InstructorInfo_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "InstructorInfo_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Courseclasses = append(node.Edges.Courseclasses, n)
		}
	}

	return nodes, nil
}

func (iiq *InstructorInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iiq.querySpec()
	return sqlgraph.CountNodes(ctx, iiq.driver, _spec)
}

func (iiq *InstructorInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (iiq *InstructorInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instructorinfo.Table,
			Columns: instructorinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instructorinfo.FieldID,
			},
		},
		From:   iiq.sql,
		Unique: true,
	}
	if ps := iiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iiq *InstructorInfoQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(iiq.driver.Dialect())
	t1 := builder.Table(instructorinfo.Table)
	selector := builder.Select(t1.Columns(instructorinfo.Columns...)...).From(t1)
	if iiq.sql != nil {
		selector = iiq.sql
		selector.Select(selector.Columns(instructorinfo.Columns...)...)
	}
	for _, p := range iiq.predicates {
		p(selector)
	}
	for _, p := range iiq.order {
		p(selector)
	}
	if offset := iiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InstructorInfoGroupBy is the builder for group-by InstructorInfo entities.
type InstructorInfoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (iigb *InstructorInfoGroupBy) Aggregate(fns ...AggregateFunc) *InstructorInfoGroupBy {
	iigb.fns = append(iigb.fns, fns...)
	return iigb
}

// Scan applies the group-by query and scan the result into the given value.
func (iigb *InstructorInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := iigb.path(ctx)
	if err != nil {
		return err
	}
	iigb.sql = query
	return iigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (iigb *InstructorInfoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := iigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (iigb *InstructorInfoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(iigb.fields) > 1 {
		return nil, errors.New("ent: InstructorInfoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := iigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (iigb *InstructorInfoGroupBy) StringsX(ctx context.Context) []string {
	v, err := iigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (iigb *InstructorInfoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = iigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instructorinfo.Label}
	default:
		err = fmt.Errorf("ent: InstructorInfoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (iigb *InstructorInfoGroupBy) StringX(ctx context.Context) string {
	v, err := iigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (iigb *InstructorInfoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(iigb.fields) > 1 {
		return nil, errors.New("ent: InstructorInfoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := iigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (iigb *InstructorInfoGroupBy) IntsX(ctx context.Context) []int {
	v, err := iigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (iigb *InstructorInfoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = iigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instructorinfo.Label}
	default:
		err = fmt.Errorf("ent: InstructorInfoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (iigb *InstructorInfoGroupBy) IntX(ctx context.Context) int {
	v, err := iigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (iigb *InstructorInfoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(iigb.fields) > 1 {
		return nil, errors.New("ent: InstructorInfoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := iigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (iigb *InstructorInfoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := iigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (iigb *InstructorInfoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = iigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instructorinfo.Label}
	default:
		err = fmt.Errorf("ent: InstructorInfoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (iigb *InstructorInfoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := iigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (iigb *InstructorInfoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(iigb.fields) > 1 {
		return nil, errors.New("ent: InstructorInfoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := iigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (iigb *InstructorInfoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := iigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (iigb *InstructorInfoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = iigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instructorinfo.Label}
	default:
		err = fmt.Errorf("ent: InstructorInfoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (iigb *InstructorInfoGroupBy) BoolX(ctx context.Context) bool {
	v, err := iigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (iigb *InstructorInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := iigb.sqlQuery().Query()
	if err := iigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (iigb *InstructorInfoGroupBy) sqlQuery() *sql.Selector {
	selector := iigb.sql
	columns := make([]string, 0, len(iigb.fields)+len(iigb.fns))
	columns = append(columns, iigb.fields...)
	for _, fn := range iigb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(iigb.fields...)
}

// InstructorInfoSelect is the builder for select fields of InstructorInfo entities.
type InstructorInfoSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (iis *InstructorInfoSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := iis.path(ctx)
	if err != nil {
		return err
	}
	iis.sql = query
	return iis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (iis *InstructorInfoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := iis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (iis *InstructorInfoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(iis.fields) > 1 {
		return nil, errors.New("ent: InstructorInfoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := iis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (iis *InstructorInfoSelect) StringsX(ctx context.Context) []string {
	v, err := iis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (iis *InstructorInfoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = iis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instructorinfo.Label}
	default:
		err = fmt.Errorf("ent: InstructorInfoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (iis *InstructorInfoSelect) StringX(ctx context.Context) string {
	v, err := iis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (iis *InstructorInfoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(iis.fields) > 1 {
		return nil, errors.New("ent: InstructorInfoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := iis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (iis *InstructorInfoSelect) IntsX(ctx context.Context) []int {
	v, err := iis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (iis *InstructorInfoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = iis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instructorinfo.Label}
	default:
		err = fmt.Errorf("ent: InstructorInfoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (iis *InstructorInfoSelect) IntX(ctx context.Context) int {
	v, err := iis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (iis *InstructorInfoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(iis.fields) > 1 {
		return nil, errors.New("ent: InstructorInfoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := iis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (iis *InstructorInfoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := iis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (iis *InstructorInfoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = iis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instructorinfo.Label}
	default:
		err = fmt.Errorf("ent: InstructorInfoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (iis *InstructorInfoSelect) Float64X(ctx context.Context) float64 {
	v, err := iis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (iis *InstructorInfoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(iis.fields) > 1 {
		return nil, errors.New("ent: InstructorInfoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := iis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (iis *InstructorInfoSelect) BoolsX(ctx context.Context) []bool {
	v, err := iis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (iis *InstructorInfoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = iis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instructorinfo.Label}
	default:
		err = fmt.Errorf("ent: InstructorInfoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (iis *InstructorInfoSelect) BoolX(ctx context.Context) bool {
	v, err := iis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (iis *InstructorInfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := iis.sqlQuery().Query()
	if err := iis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (iis *InstructorInfoSelect) sqlQuery() sql.Querier {
	selector := iis.sql
	selector.Select(selector.Columns(iis.fields...)...)
	return selector
}
