package repro

import (
	"flag"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/repo"
	"github.com/bazelbuild/bazel-gazelle/resolve"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

func NewLanguage() language.Language {
	return new(lang)
}

var _ language.Language = (*lang)(nil)

type lang struct{}

func (l *lang) RegisterFlags(fs *flag.FlagSet, cmd string, c *config.Config) {

}

func (l *lang) CheckFlags(fs *flag.FlagSet, c *config.Config) error {
	return nil
}

func (l *lang) Embeds(r *rule.Rule, from label.Label) []label.Label {
	return nil
}

func (*lang) Kinds() map[string]rule.KindInfo {
	return nil
}

func (l *lang) Loads() []rule.LoadInfo {
	return nil
}

func (l *lang) Fix(c *config.Config, f *rule.File) {

}

func (*lang) Name() string { return "repro" }

func (l *lang) Resolve(c *config.Config, ix *resolve.RuleIndex, rc *repo.RemoteCache, r *rule.Rule, rawImports interface{}, from label.Label) {

}

func (l *lang) Imports(c *config.Config, r *rule.Rule, f *rule.File) []resolve.ImportSpec {
	return nil
}

func (gl *lang) GenerateRules(args language.GenerateArgs) language.GenerateResult {

	return language.GenerateResult{}
}

func (l *lang) KnownDirectives() []string {
	return []string{}
}

func (l *lang) Configure(cfg *config.Config, rel string, f *rule.File) {
}
