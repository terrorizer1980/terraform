package refactoring

import (
	"fmt"

	"github.com/hashicorp/terraform/internal/addrs"
	"github.com/hashicorp/terraform/internal/configs"
	"github.com/hashicorp/terraform/internal/tfdiags"
)

// ValidateMove checks whether the given from and to addresses from a single
// "moved" block are compatible with one another and with the objects they
// refer to in the configuration.
//
// ValidateMove expects to recieve addresses resulting from calling
// addrs.UnifyMoveEndpoints and thus requires that both "from" and "to"
// always have the same dynamic address type. If not, this function will
// panic.
//
// This function is limited to validating a single "moved" block in isolation.
// It doesn't consider the interactions between different "moved" blocks
// that might have overlapping addresses.
func ValidateMove(from, to addrs.AbsMoveable, srcRange tfdiags.SourceRange, fullConfig *configs.Config) tfdiags.Diagnostics {
	switch from := from.(type) {
	case addrs.AbsResourceInstance:
	default:
		panic(fmt.Sprintf("unsupported AbsMovable type %T", from))
	}
}
