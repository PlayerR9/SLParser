package pkg

import (
	"strings"

	"golang.org/x/exp/slices"
)

// Rule is a rule of the grammar.
type Rule struct {
	// lhs is the left-hand side of the rule.
	lhs string

	// rhss are the right-hand sides of the rule.
	rhss []string
}

// String implements the fmt.Stringer interface.
//
// Format:
//
//	RHS(n) RHS(n-1) ... RHS(1) -> LHS .
func (r *Rule) String() string {
	values := make([]string, len(r.rhss))
	copy(values, r.rhss)
	values = append(values, "->")
	values = append(values, r.lhs)
	values = append(values, ".")

	return strings.Join(values, " ")
}

// NewRule is a constructor for a Rule.
//
// Parameters:
//   - lhs: The left-hand side of the rule.
//   - rhss: The right-hand sides of the rule.
//
// Returns:
//   - *Rule: The created rule.
//
// Returns nil iff the rhss is empty.
func NewRule(lhs string, rhss []string, is_reverse bool) *Rule {
	if len(rhss) == 0 {
		return nil
	}

	if !is_reverse {
		slices.Reverse(rhss)
	}

	return &Rule{
		lhs:  lhs,
		rhss: rhss,
	}
}

// GetLhs returns the left-hand side of the rule.
//
// Returns:
//   - string: The left-hand side of the rule.
func (r *Rule) GetLhs() string {
	return r.lhs
}

// GetIndicesOfRhs returns the ocurrence indices of the rhs in the rule.
//
// Parameters:
//   - rhs: The right-hand side to search.
//
// Returns:
//   - []int: The indices of the rhs in the rule.
func (r *Rule) GetIndicesOfRhs(rhs string) []int {
	var indices []int

	for i := 0; i < len(r.rhss); i++ {
		if r.rhss[i] == rhs {
			indices = append(indices, i)
		}
	}

	return indices
}

// GetRhss returns the right-hand sides of the rule.
//
// Returns:
//   - []string: The right-hand sides of the rule.
func (r *Rule) GetRhss() []string {
	return r.rhss
}

// Size returns the number of right-hand sides of the rule.
//
// Returns:
//   - int: The "size" of the rule.
func (r *Rule) Size() int {
	return len(r.rhss)
}

// GetSymbols returns the symbols of the rule.
//
// Returns:
//   - []string: The symbols of the rule.
//
// The symbols are unique and sorted.
func (r *Rule) GetSymbols() []string {
	var symbols []string

	for _, rhs := range r.rhss {
		pos, ok := slices.BinarySearch(symbols, rhs)
		if !ok {
			symbols = slices.Insert(symbols, pos, rhs)
		}
	}

	pos, ok := slices.BinarySearch(symbols, r.lhs)
	if !ok {
		symbols = slices.Insert(symbols, pos, r.lhs)
	}

	return symbols
}

// GetRuleTempl returns the template of the rule.
//
// Parameters:
//   - pkg_name: The name of the package.
//   - tt_name: The name of the token type.
//
// Returns:
//   - string: The template of the rule.
func (r *Rule) GetRuleTempl(pkg_name, tt_name string) string {
	var builder strings.Builder

	builder.WriteString(pkg_name)
	builder.WriteString(".NewRule(")
	builder.WriteString(r.lhs)
	builder.WriteString(", []")
	builder.WriteString(tt_name)
	builder.WriteString("{")
	builder.WriteString(strings.Join(r.rhss, ", "))
	builder.WriteString("})")

	return builder.String()
}

// StringOriginal returns the original string of the rule.
//
// Returns:
//   - string: The original string of the rule.
func (r *Rule) StringOriginal() string {
	rhss := make([]string, 0, len(r.rhss))

	for i := len(r.rhss) - 1; i >= 0; i-- {
		rhss = append(rhss, r.rhss[i])
	}

	var builder strings.Builder

	builder.WriteString(r.lhs)
	builder.WriteString(" : ")
	builder.WriteString(strings.Join(rhss, " "))
	builder.WriteString(" .")

	return builder.String()
}
