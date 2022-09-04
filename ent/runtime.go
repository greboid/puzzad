// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/greboid/puzzad/ent/guess"
	"github.com/greboid/puzzad/ent/schema"
	"github.com/greboid/puzzad/ent/team"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accessFields := schema.Access{}.Fields()
	_ = accessFields
	guessMixin := schema.Guess{}.Mixin()
	guessMixinFields0 := guessMixin[0].Fields()
	_ = guessMixinFields0
	guessFields := schema.Guess{}.Fields()
	_ = guessFields
	// guessDescCreateTime is the schema descriptor for create_time field.
	guessDescCreateTime := guessMixinFields0[0].Descriptor()
	// guess.DefaultCreateTime holds the default value on creation for the create_time field.
	guess.DefaultCreateTime = guessDescCreateTime.Default.(func() time.Time)
	teamMixin := schema.Team{}.Mixin()
	teamMixinFields0 := teamMixin[0].Fields()
	_ = teamMixinFields0
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescCreateTime is the schema descriptor for create_time field.
	teamDescCreateTime := teamMixinFields0[0].Descriptor()
	// team.DefaultCreateTime holds the default value on creation for the create_time field.
	team.DefaultCreateTime = teamDescCreateTime.Default.(func() time.Time)
	// teamDescName is the schema descriptor for name field.
	teamDescName := teamFields[0].Descriptor()
	// team.NameValidator is a validator for the "name" field. It is called by the builders before save.
	team.NameValidator = func() func(string) error {
		validators := teamDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// teamDescCode is the schema descriptor for code field.
	teamDescCode := teamFields[1].Descriptor()
	// team.DefaultCode holds the default value on creation for the code field.
	team.DefaultCode = teamDescCode.Default.(func() string)
	// teamDescVerifyCode is the schema descriptor for verifyCode field.
	teamDescVerifyCode := teamFields[2].Descriptor()
	// team.DefaultVerifyCode holds the default value on creation for the verifyCode field.
	team.DefaultVerifyCode = teamDescVerifyCode.Default.(func() string)
}
