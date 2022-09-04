// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/greboid/puzzad/ent/access"
	"github.com/greboid/puzzad/ent/adventure"
	"github.com/greboid/puzzad/ent/team"
)

// Access is the model entity for the Access schema.
type Access struct {
	config `json:"-"`
	// Status holds the value of the "status" field.
	Status access.Status `json:"status,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// TeamID holds the value of the "team_id" field.
	TeamID int `json:"team_id,omitempty"`
	// AdventureID holds the value of the "adventure_id" field.
	AdventureID int `json:"adventure_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccessQuery when eager-loading is set.
	Edges AccessEdges `json:"edges"`
}

// AccessEdges holds the relations/edges for other nodes in the graph.
type AccessEdges struct {
	// Team holds the value of the team edge.
	Team *Team `json:"team,omitempty"`
	// Adventures holds the value of the adventures edge.
	Adventures *Adventure `json:"adventures,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccessEdges) TeamOrErr() (*Team, error) {
	if e.loadedTypes[0] {
		if e.Team == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: team.Label}
		}
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// AdventuresOrErr returns the Adventures value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccessEdges) AdventuresOrErr() (*Adventure, error) {
	if e.loadedTypes[1] {
		if e.Adventures == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: adventure.Label}
		}
		return e.Adventures, nil
	}
	return nil, &NotLoadedError{edge: "adventures"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Access) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case access.FieldTeamID, access.FieldAdventureID:
			values[i] = new(sql.NullInt64)
		case access.FieldStatus, access.FieldCode:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Access", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Access fields.
func (a *Access) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case access.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				a.Status = access.Status(value.String)
			}
		case access.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				a.Code = value.String
			}
		case access.FieldTeamID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field team_id", values[i])
			} else if value.Valid {
				a.TeamID = int(value.Int64)
			}
		case access.FieldAdventureID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field adventure_id", values[i])
			} else if value.Valid {
				a.AdventureID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTeam queries the "team" edge of the Access entity.
func (a *Access) QueryTeam() *TeamQuery {
	return (&AccessClient{config: a.config}).QueryTeam(a)
}

// QueryAdventures queries the "adventures" edge of the Access entity.
func (a *Access) QueryAdventures() *AdventureQuery {
	return (&AccessClient{config: a.config}).QueryAdventures(a)
}

// Update returns a builder for updating this Access.
// Note that you need to call Access.Unwrap() before calling this method if this Access
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Access) Update() *AccessUpdateOne {
	return (&AccessClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Access entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Access) Unwrap() *Access {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Access is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Access) String() string {
	var builder strings.Builder
	builder.WriteString("Access(")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", a.Status))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(a.Code)
	builder.WriteString(", ")
	builder.WriteString("team_id=")
	builder.WriteString(fmt.Sprintf("%v", a.TeamID))
	builder.WriteString(", ")
	builder.WriteString("adventure_id=")
	builder.WriteString(fmt.Sprintf("%v", a.AdventureID))
	builder.WriteByte(')')
	return builder.String()
}

// Accesses is a parsable slice of Access.
type Accesses []*Access

func (a Accesses) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
