// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/greboid/puzzad/ent/adventure"
	"github.com/greboid/puzzad/ent/game"
	"github.com/greboid/puzzad/ent/puzzle"
	"github.com/greboid/puzzad/ent/user"
)

// Game is the model entity for the Game schema.
type Game struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status game.Status `json:"status,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GameQuery when eager-loading is set.
	Edges               GameEdges `json:"edges"`
	game_adventure      *int
	game_current_puzzle *int
	user_game           *int
}

// GameEdges holds the relations/edges for other nodes in the graph.
type GameEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Adventure holds the value of the adventure edge.
	Adventure *Adventure `json:"adventure,omitempty"`
	// CurrentPuzzle holds the value of the current_puzzle edge.
	CurrentPuzzle *Puzzle `json:"current_puzzle,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// AdventureOrErr returns the Adventure value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameEdges) AdventureOrErr() (*Adventure, error) {
	if e.loadedTypes[1] {
		if e.Adventure == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: adventure.Label}
		}
		return e.Adventure, nil
	}
	return nil, &NotLoadedError{edge: "adventure"}
}

// CurrentPuzzleOrErr returns the CurrentPuzzle value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameEdges) CurrentPuzzleOrErr() (*Puzzle, error) {
	if e.loadedTypes[2] {
		if e.CurrentPuzzle == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: puzzle.Label}
		}
		return e.CurrentPuzzle, nil
	}
	return nil, &NotLoadedError{edge: "current_puzzle"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Game) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case game.FieldID:
			values[i] = new(sql.NullInt64)
		case game.FieldStatus, game.FieldCode:
			values[i] = new(sql.NullString)
		case game.ForeignKeys[0]: // game_adventure
			values[i] = new(sql.NullInt64)
		case game.ForeignKeys[1]: // game_current_puzzle
			values[i] = new(sql.NullInt64)
		case game.ForeignKeys[2]: // user_game
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Game", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Game fields.
func (ga *Game) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case game.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ga.ID = int(value.Int64)
		case game.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ga.Status = game.Status(value.String)
			}
		case game.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				ga.Code = value.String
			}
		case game.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field game_adventure", value)
			} else if value.Valid {
				ga.game_adventure = new(int)
				*ga.game_adventure = int(value.Int64)
			}
		case game.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field game_current_puzzle", value)
			} else if value.Valid {
				ga.game_current_puzzle = new(int)
				*ga.game_current_puzzle = int(value.Int64)
			}
		case game.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_game", value)
			} else if value.Valid {
				ga.user_game = new(int)
				*ga.user_game = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Game entity.
func (ga *Game) QueryUser() *UserQuery {
	return (&GameClient{config: ga.config}).QueryUser(ga)
}

// QueryAdventure queries the "adventure" edge of the Game entity.
func (ga *Game) QueryAdventure() *AdventureQuery {
	return (&GameClient{config: ga.config}).QueryAdventure(ga)
}

// QueryCurrentPuzzle queries the "current_puzzle" edge of the Game entity.
func (ga *Game) QueryCurrentPuzzle() *PuzzleQuery {
	return (&GameClient{config: ga.config}).QueryCurrentPuzzle(ga)
}

// Update returns a builder for updating this Game.
// Note that you need to call Game.Unwrap() before calling this method if this Game
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *Game) Update() *GameUpdateOne {
	return (&GameClient{config: ga.config}).UpdateOne(ga)
}

// Unwrap unwraps the Game entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ga *Game) Unwrap() *Game {
	_tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("ent: Game is not a transactional entity")
	}
	ga.config.driver = _tx.drv
	return ga
}

// String implements the fmt.Stringer.
func (ga *Game) String() string {
	var builder strings.Builder
	builder.WriteString("Game(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ga.ID))
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ga.Status))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(ga.Code)
	builder.WriteByte(')')
	return builder.String()
}

// Games is a parsable slice of Game.
type Games []*Game

func (ga Games) config(cfg config) {
	for _i := range ga {
		ga[_i].config = cfg
	}
}
