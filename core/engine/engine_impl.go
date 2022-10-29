package engine

import (
	"github.com/olehmushka/golang-toolkit/either"
	wordgenerator "github.com/olehmushka/word-generator"
	eng "github.com/olehmushka/world-generator-engine/engine"
	"github.com/olehmushka/world-generator-engine/language"
	"go.uber.org/fx"
)

type engine struct {
	driver        eng.Engine
	wordGenerator wordgenerator.Generator
}

func New() Engine {
	return &engine{driver: eng.New(wordgenerator.New())}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (e *engine) LoadLanguageFamilies() chan either.Either[[]string] {
	return e.driver.LoadLanguageFamilies()
}

func (e *engine) LoadLanguageSubfamilies() chan either.Either[[]*language.Subfamily] {
	return e.driver.LoadLanguageSubfamilies()
}

func (e *engine) LoadLanguages() chan either.Either[*language.Language] {
	return e.driver.LoadLanguages()
}

func (e *engine) GenerateWord(lang *language.Language) (string, error) {
	return e.driver.GenerateWord(lang)
}
