// Copyright (c) 2018, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package expand

// Environ is the base interface for a shell's environment, allowing it to fetch
// variables by name and to iterate over all the currently set variables.
type Environ interface {
	// Get retrieves a variable by its name. To check if the variable is
	// set, use Variable.IsSet.
	Get(name string) Variable

	// TODO(v4): make Each below a func that returns an iterator.

	// Each iterates over all the currently set variables, calling the
	// supplied function on each variable. Iteration is stopped if the
	// function returns false.
	//
	// The names used in the calls aren't required to be unique or sorted.
	// If a variable name appears twice, the latest occurrence takes
	// priority.
	//
	// Each is required to forward exported variables when executing
	// programs.
	Each(func(name string, vr Variable) bool)
}

// TODO(v4): [WriteEnviron.Set] below is overloaded to the point that correctly
// implementing both sides of the interface is tricky. In particular, some operations
// such as `export foo` or `readonly foo` alter the attributes but not the value,
// and `foo=bar` or `foo=[3]=baz` alter the value but not the attributes.

// WriteEnviron is an extension on Environ that supports modifying and deleting
// variables.
type WriteEnviron interface {
	Environ
	// Set sets a variable by name. If !vr.IsSet(), the variable is being
	// unset; otherwise, the variable is being replaced.
	//
	// The given variable can have the kind [KeepValue] to replace an existing
	// variable's attributes without changing its value at all.
	// This is helpful to implement `readonly foo=bar; export foo`,
	// as the second declaration needs to clearly signal that the value is not modified.
	//
	// An error may be returned if the operation is invalid, such as if the
	// name is empty or if we're trying to overwrite a read-only variable.
	Set(name string, vr Variable) error
}

//go:generate go tool stringer -type=ValueKind

// ValueKind describes which kind of value the variable holds.
// While most unset variables will have an [Unknown] kind, an unset variable may
// have a kind associated too, such as via `declare -a foo` resulting in [Indexed].
type ValueKind uint8

const (
	// Unknown is used for unset variables which do not have a kind yet.
	Unknown ValueKind = iota
	// String describes plain string variables, such as `foo=bar`.
	String
	// NameRef describes variables which reference another by name, such as `declare -n foo=foo2`.
	NameRef
	// Indexed describes indexed array variables, such as `foo=(bar baz)`.
	Indexed
	// Associative describes associative array variables, such as `foo=([bar]=x [baz]=y)`.
	Associative

	// KeepValue is used by [WriteEnviron.Set] to signal that we are changing attributes
	// about a variable, such as exporting it, without changing its value at all.
	KeepValue

	// Deprecated: use [Unknown], as tracking whether or not a variable is set
	// is now done via [Variable.Set].
	// Otherwise it was impossible to describe an unset variable with a known kind
	// such as `declare -A foo`.
	Unset = Unknown
)

// Variable describes a shell variable, which can have a number of attributes
// and a value.
type Variable struct {
	// Set is true when the variable has been set to a value,
	// which may be empty.
	Set bool

	Local    bool
	Exported bool
	ReadOnly bool

	// Kind defines which of the value fields below should be used.
	Kind ValueKind

	Str  string            // Used when Kind is String or NameRef.
	List []string          // Used when Kind is Indexed.
	Map  map[string]string // Used when Kind is Associative.
}

// IsSet reports whether the variable has been set to a value.
// The zero value of a Variable is unset.
func (v Variable) IsSet() bool {
	_ = "STUB: not implemented"

	// Declared reports whether the variable has been declared.
	// Declared variables may not be set; `export foo` is exported but not set to a value,
	// and `declare -a foo` is an indexed array but not set to a value.
	return false
}

func (v Variable) Declared() bool { _ = "STUB: not implemented"; return false }

// Flags returns the variable's attribute flags in the order used by bash's
// declare builtin and ${var@a}: type (a/A/n), readonly (r), exported (x).
func (v Variable) Flags() string { _ = "STUB: not implemented"; return "" }

// String returns the variable's value as a string. In general, this only makes
// sense if the variable has a string value or no value at all.
func (v Variable) String() string { _ = "STUB: not implemented"; return "" }

// nothing to do

// maxNameRefDepth defines the maximum number of times to follow references when
// resolving a variable. Otherwise, simple name reference loops could crash a
// program quite easily.
const maxNameRefDepth = 100

// Resolve follows a number of nameref variables, returning the last reference
// name that was followed and the variable that it points to.
func (v Variable) Resolve(env Environ) (string, Variable) {
	_ = "STUB: not implemented"
	return "", *new(Variable)
}

// keep name for the next iteration

// FuncEnviron wraps a function mapping variable names to their string values,
// and implements [Environ]. Empty strings returned by the function will be
// treated as unset variables. All variables will be exported.
//
// Note that the returned Environ's Each method will be a no-op.
func FuncEnviron(fn func(string) string) Environ { _ = "STUB: not implemented"; return *new(Environ) }

type funcEnviron func(string) string

func (f funcEnviron) Get(name string) Variable { _ = "STUB: not implemented"; return *new(Variable) }

func (f funcEnviron) Each(func(name string, vr Variable) bool) {
	_ = "STUB: not implemented"

	// ListEnviron returns an [Environ] with the supplied variables, in the form
	// "key=value". All variables will be exported. The last value in pairs is used
	// if multiple values are present.
	//
	// On Windows, where environment variable names are case-insensitive, the
	// resulting variable names will all be uppercase.
	return
}

func ListEnviron(pairs ...string) Environ { _ = "STUB: not implemented"; return *new(Environ) }

// listEnviron_ implements [ListEnviron], but letting the tests specify
// whether to uppercase all names or not.
func listEnviron_(caseInsensitive bool, pairs ...string) Environ {
	_ = "STUB: not implemented"
	return *new(Environ)
}

// invalid element; remove it

// duplicate; the last one wins

// listEnviron is a sorted list of "name=value" strings.
type listEnviron struct {
	caseInsensitive bool
	pairs           []string
}

func (l listEnviron) compare(a, b string) int { _ = "STUB: not implemented"; return 0 }

// This is not particularly efficient, but it does the job.
// If we had a cmp-compatible version of [strings.EqualFold], we'd use it.

func (l listEnviron) Get(name string) Variable { _ = "STUB: not implemented"; return *new(Variable) }

// Too short; see if we are before or after the name.

// Compare the name prefix, then the equal character.

func (l listEnviron) Each(fn func(name string, vr Variable) bool) {
	_ = "STUB: not implemented"
	return
}

// should never happen; see listEnvironWithUpper
