package validator

import (
	"github.com/advoretsky/gqlparser/v2/ast"

	//nolint:revive // Validator rules each use dot imports for convenience.
	. "github.com/advoretsky/gqlparser/v2/validator"
)

func init() {
	AddRule("UniqueDirectivesPerLocation", func(observers *Events, addError AddErrFunc) {
		observers.OnDirectiveList(func(walker *Walker, directives []*ast.Directive) {
			seen := map[string]bool{}

			for _, dir := range directives {
				if dir.Name != "repeatable" && seen[dir.Name] {
					addError(
						Message(`The directive "@%s" can only be used once at this location.`, dir.Name),
						At(dir.Position),
					)
				}
				seen[dir.Name] = true
			}
		})
	})
}
