//go:build ruleguard
// +build ruleguard

package rules

import "github.com/quasilyte/go-ruleguard/dsl"

func boolFunctionNaming(m dsl.Matcher) {
	m.Match(`func $name($*params) bool { $*body }`).
		Where(!m["name"].Text.Matches(`^(Is).*`)).
		Report("bool function name should start with 'Is'").
		Suggest(`func Is$name $params bool { $body }`)
}
