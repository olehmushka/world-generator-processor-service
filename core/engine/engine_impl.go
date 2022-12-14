package engine

import (
	"strings"

	"github.com/olehmushka/golang-toolkit/either"
	wordgenerator "github.com/olehmushka/word-generator"
	culture "github.com/olehmushka/world-generator-engine/culture"
	eng "github.com/olehmushka/world-generator-engine/engine"
	gender "github.com/olehmushka/world-generator-engine/gender"
	genderAcceptance "github.com/olehmushka/world-generator-engine/gender_acceptance"
	influence "github.com/olehmushka/world-generator-engine/influence"
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
	return e.driver.LoadLanguageFamilies(func(s string) string {
		for _, engineWord := range []string{"/engine", "/engine/", "engine/"} {
			s = strings.ReplaceAll(s, engineWord, "/")
		}
		return s
	})
}

func (e *engine) LoadLanguageSubfamilies() chan either.Either[[]*language.Subfamily] {
	return e.driver.LoadLanguageSubfamilies(func(s string) string {
		for _, engineWord := range []string{"/engine", "/engine/", "engine/"} {
			s = strings.Replace(s, engineWord, "/", -1)
		}
		return s
	})
}

func (e *engine) LoadLanguages() chan either.Either[*language.Language] {
	return e.driver.LoadLanguages(func(s string) string {
		for _, engineWord := range []string{"/engine", "/engine/", "engine/"} {
			s = strings.ReplaceAll(s, engineWord, "/")
		}
		return s
	})
}

func (e *engine) GenerateWord(lang *language.Language) (string, error) {
	return e.driver.GenerateWord(lang)
}

func (e *engine) LoadGenders() []gender.Sex {
	return e.driver.LoadGenders()
}

func (e *engine) LoadGenderAcceptances() []genderAcceptance.Acceptance {
	return e.driver.LoadGenderAcceptances()
}

func (e *engine) LoadInfluences() []influence.Influence {
	return e.driver.LoadInfluences()
}

func (e *engine) LoadCulturesBases() chan either.Either[[]string] {
	return e.driver.LoadCulturesBases(func(s string) string {
		for _, engineWord := range []string{"/engine", "/engine/", "engine/"} {
			s = strings.ReplaceAll(s, engineWord, "/")
		}
		return s
	})
}

func (e *engine) LoadCultureSubbases() chan either.Either[[]*culture.Subbase] {
	return e.driver.LoadCultureSubbases(func(s string) string {
		for _, engineWord := range []string{"/engine", "/engine/", "engine/"} {
			s = strings.ReplaceAll(s, engineWord, "/")
		}
		return s
	})
}

func (e *engine) LoadAllEthoses() chan either.Either[[]culture.Ethos] {
	return e.driver.LoadAllEthoses(func(s string) string {
		for _, engineWord := range []string{"/engine", "/engine/", "engine/"} {
			s = strings.ReplaceAll(s, engineWord, "/")
		}
		return s
	})
}

func (e *engine) LoadAllTraditions() chan either.Either[[]*culture.Tradition] {
	return e.driver.LoadAllTraditions(func(s string) string {
		for _, engineWord := range []string{"/engine", "/engine/", "engine/"} {
			s = strings.ReplaceAll(s, engineWord, "/")
		}
		return s
	})
}

func (e *engine) LoadAllParentRawCultures() chan either.Either[[]*culture.RawCulture] {
	return e.driver.LoadAllParentRawCultures(func(s string) string {
		for _, engineWord := range []string{"/engine", "/engine/", "engine/"} {
			s = strings.ReplaceAll(s, engineWord, "/")
		}
		return s
	})
}

func (e *engine) LoadAllParentCultures() chan either.Either[*culture.Culture] {
	return e.driver.LoadAllParentCultures(func(s string) string {
		for _, engineWord := range []string{"/engine", "/engine/", "engine/"} {
			s = strings.ReplaceAll(s, engineWord, "/")
		}
		return s
	})
}

func (e *engine) Generate(opts *culture.CreateCultureOpts, parents ...*culture.Culture) (*culture.Culture, error) {
	return e.driver.Generate(opts, parents...)
}
