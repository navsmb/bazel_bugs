package repro

import "github.com/bazel-contrib/bazel-gazelle/v2/language"

func NewV2() language.Language {
	return new(lang)
}

type lang struct{}

func (lang) Name() string { return "repro" }
