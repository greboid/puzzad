// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/greboid/puzzad/ent/adventure"
	"github.com/greboid/puzzad/ent/game"
	"github.com/greboid/puzzad/ent/predicate"
	"github.com/greboid/puzzad/ent/puzzle"
	"github.com/greboid/puzzad/ent/user"
)

// GameUpdate is the builder for updating Game entities.
type GameUpdate struct {
	config
	hooks    []Hook
	mutation *GameMutation
}

// Where appends a list predicates to the GameUpdate builder.
func (gu *GameUpdate) Where(ps ...predicate.Game) *GameUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetStatus sets the "status" field.
func (gu *GameUpdate) SetStatus(ga game.Status) *GameUpdate {
	gu.mutation.SetStatus(ga)
	return gu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (gu *GameUpdate) SetNillableStatus(ga *game.Status) *GameUpdate {
	if ga != nil {
		gu.SetStatus(*ga)
	}
	return gu
}

// SetCode sets the "code" field.
func (gu *GameUpdate) SetCode(s string) *GameUpdate {
	gu.mutation.SetCode(s)
	return gu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (gu *GameUpdate) SetNillableCode(s *string) *GameUpdate {
	if s != nil {
		gu.SetCode(*s)
	}
	return gu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gu *GameUpdate) SetUserID(id int) *GameUpdate {
	gu.mutation.SetUserID(id)
	return gu
}

// SetUser sets the "user" edge to the User entity.
func (gu *GameUpdate) SetUser(u *User) *GameUpdate {
	return gu.SetUserID(u.ID)
}

// SetAdventureID sets the "adventure" edge to the Adventure entity by ID.
func (gu *GameUpdate) SetAdventureID(id int) *GameUpdate {
	gu.mutation.SetAdventureID(id)
	return gu
}

// SetAdventure sets the "adventure" edge to the Adventure entity.
func (gu *GameUpdate) SetAdventure(a *Adventure) *GameUpdate {
	return gu.SetAdventureID(a.ID)
}

// SetCurrentPuzzleID sets the "current_puzzle" edge to the Puzzle entity by ID.
func (gu *GameUpdate) SetCurrentPuzzleID(id int) *GameUpdate {
	gu.mutation.SetCurrentPuzzleID(id)
	return gu
}

// SetCurrentPuzzle sets the "current_puzzle" edge to the Puzzle entity.
func (gu *GameUpdate) SetCurrentPuzzle(p *Puzzle) *GameUpdate {
	return gu.SetCurrentPuzzleID(p.ID)
}

// Mutation returns the GameMutation object of the builder.
func (gu *GameUpdate) Mutation() *GameMutation {
	return gu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (gu *GameUpdate) ClearUser() *GameUpdate {
	gu.mutation.ClearUser()
	return gu
}

// ClearAdventure clears the "adventure" edge to the Adventure entity.
func (gu *GameUpdate) ClearAdventure() *GameUpdate {
	gu.mutation.ClearAdventure()
	return gu
}

// ClearCurrentPuzzle clears the "current_puzzle" edge to the Puzzle entity.
func (gu *GameUpdate) ClearCurrentPuzzle() *GameUpdate {
	gu.mutation.ClearCurrentPuzzle()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GameUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		if err = gu.check(); err != nil {
			return 0, err
		}
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gu.check(); err != nil {
				return 0, err
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GameUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GameUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GameUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GameUpdate) check() error {
	if v, ok := gu.mutation.Status(); ok {
		if err := game.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Game.status": %w`, err)}
		}
	}
	if v, ok := gu.mutation.Code(); ok {
		if err := game.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Game.code": %w`, err)}
		}
	}
	if _, ok := gu.mutation.UserID(); gu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.user"`)
	}
	if _, ok := gu.mutation.AdventureID(); gu.mutation.AdventureCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.adventure"`)
	}
	if _, ok := gu.mutation.CurrentPuzzleID(); gu.mutation.CurrentPuzzleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.current_puzzle"`)
	}
	return nil
}

func (gu *GameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   game.Table,
			Columns: game.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: game.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: game.FieldStatus,
		})
	}
	if value, ok := gu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldCode,
		})
	}
	if gu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.UserTable,
			Columns: []string{game.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.UserTable,
			Columns: []string{game.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.AdventureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   game.AdventureTable,
			Columns: []string{game.AdventureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adventure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.AdventureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   game.AdventureTable,
			Columns: []string{game.AdventureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adventure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.CurrentPuzzleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   game.CurrentPuzzleTable,
			Columns: []string{game.CurrentPuzzleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: puzzle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.CurrentPuzzleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   game.CurrentPuzzleTable,
			Columns: []string{game.CurrentPuzzleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: puzzle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GameUpdateOne is the builder for updating a single Game entity.
type GameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GameMutation
}

// SetStatus sets the "status" field.
func (guo *GameUpdateOne) SetStatus(ga game.Status) *GameUpdateOne {
	guo.mutation.SetStatus(ga)
	return guo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableStatus(ga *game.Status) *GameUpdateOne {
	if ga != nil {
		guo.SetStatus(*ga)
	}
	return guo
}

// SetCode sets the "code" field.
func (guo *GameUpdateOne) SetCode(s string) *GameUpdateOne {
	guo.mutation.SetCode(s)
	return guo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableCode(s *string) *GameUpdateOne {
	if s != nil {
		guo.SetCode(*s)
	}
	return guo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (guo *GameUpdateOne) SetUserID(id int) *GameUpdateOne {
	guo.mutation.SetUserID(id)
	return guo
}

// SetUser sets the "user" edge to the User entity.
func (guo *GameUpdateOne) SetUser(u *User) *GameUpdateOne {
	return guo.SetUserID(u.ID)
}

// SetAdventureID sets the "adventure" edge to the Adventure entity by ID.
func (guo *GameUpdateOne) SetAdventureID(id int) *GameUpdateOne {
	guo.mutation.SetAdventureID(id)
	return guo
}

// SetAdventure sets the "adventure" edge to the Adventure entity.
func (guo *GameUpdateOne) SetAdventure(a *Adventure) *GameUpdateOne {
	return guo.SetAdventureID(a.ID)
}

// SetCurrentPuzzleID sets the "current_puzzle" edge to the Puzzle entity by ID.
func (guo *GameUpdateOne) SetCurrentPuzzleID(id int) *GameUpdateOne {
	guo.mutation.SetCurrentPuzzleID(id)
	return guo
}

// SetCurrentPuzzle sets the "current_puzzle" edge to the Puzzle entity.
func (guo *GameUpdateOne) SetCurrentPuzzle(p *Puzzle) *GameUpdateOne {
	return guo.SetCurrentPuzzleID(p.ID)
}

// Mutation returns the GameMutation object of the builder.
func (guo *GameUpdateOne) Mutation() *GameMutation {
	return guo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (guo *GameUpdateOne) ClearUser() *GameUpdateOne {
	guo.mutation.ClearUser()
	return guo
}

// ClearAdventure clears the "adventure" edge to the Adventure entity.
func (guo *GameUpdateOne) ClearAdventure() *GameUpdateOne {
	guo.mutation.ClearAdventure()
	return guo
}

// ClearCurrentPuzzle clears the "current_puzzle" edge to the Puzzle entity.
func (guo *GameUpdateOne) ClearCurrentPuzzle() *GameUpdateOne {
	guo.mutation.ClearCurrentPuzzle()
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GameUpdateOne) Select(field string, fields ...string) *GameUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Game entity.
func (guo *GameUpdateOne) Save(ctx context.Context) (*Game, error) {
	var (
		err  error
		node *Game
	)
	if len(guo.hooks) == 0 {
		if err = guo.check(); err != nil {
			return nil, err
		}
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = guo.check(); err != nil {
				return nil, err
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, guo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Game)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GameMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GameUpdateOne) SaveX(ctx context.Context) *Game {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GameUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GameUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GameUpdateOne) check() error {
	if v, ok := guo.mutation.Status(); ok {
		if err := game.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Game.status": %w`, err)}
		}
	}
	if v, ok := guo.mutation.Code(); ok {
		if err := game.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Game.code": %w`, err)}
		}
	}
	if _, ok := guo.mutation.UserID(); guo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.user"`)
	}
	if _, ok := guo.mutation.AdventureID(); guo.mutation.AdventureCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.adventure"`)
	}
	if _, ok := guo.mutation.CurrentPuzzleID(); guo.mutation.CurrentPuzzleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.current_puzzle"`)
	}
	return nil
}

func (guo *GameUpdateOne) sqlSave(ctx context.Context) (_node *Game, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   game.Table,
			Columns: game.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: game.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Game.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, game.FieldID)
		for _, f := range fields {
			if !game.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != game.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: game.FieldStatus,
		})
	}
	if value, ok := guo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldCode,
		})
	}
	if guo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.UserTable,
			Columns: []string{game.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.UserTable,
			Columns: []string{game.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.AdventureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   game.AdventureTable,
			Columns: []string{game.AdventureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adventure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.AdventureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   game.AdventureTable,
			Columns: []string{game.AdventureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adventure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.CurrentPuzzleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   game.CurrentPuzzleTable,
			Columns: []string{game.CurrentPuzzleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: puzzle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.CurrentPuzzleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   game.CurrentPuzzleTable,
			Columns: []string{game.CurrentPuzzleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: puzzle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Game{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
