// +build !linux !cgo !pcre1,!pcre2

package fastregex

import "regexp"

/*
The fallback exposes the re2/regexp go implementaiton in the
cases where we can't compile with PCRE support
*/

const Version = "re2"

type compiledRegexp struct {
	re *regexp.Regexp
}

var _ CompiledRegexp = &compiledRegexp{}

func (s *compiledRegexp) CreateInstance() Regexp {
	return s.re
}

func buildRegexp(expr string, posix bool) (*regexp.Regexp, error) {
	if posix {
		return regexp.CompilePOSIX(expr)
	}
	return regexp.Compile(expr)
}

func Compile(expr string, posix bool) (CompiledRegexp, error) {
	re, err := buildRegexp(expr, posix)
	if err != nil {
		return nil, err
	}
	return &compiledRegexp{re}, nil
}

func MustCompile(expr string, posix bool) CompiledRegexp {
	if posix {
		return &compiledRegexp{regexp.MustCompilePOSIX(expr)}
	}
	return &compiledRegexp{regexp.MustCompile(expr)}
}
