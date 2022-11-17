package engine

import (
	"context"
	"world_generator_processor_service/core/engine"
	"world_generator_processor_service/internal/languages/entities"

	"github.com/olehmushka/golang-toolkit/either"
	engineLang "github.com/olehmushka/world-generator-engine/language"
	"go.uber.org/fx"
)

type adapter struct {
	engine engine.Engine
}

func New(eng engine.Engine) Adapter {
	return &adapter{
		engine: eng,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (a *adapter) LoadNativeFamilies(context.Context) chan either.Either[[]*entities.Family] {
	ch := make(chan either.Either[[]*entities.Family])
	go func() {
		for familiesChunk := range a.engine.LoadLanguageFamilies() {
			if err := familiesChunk.Err; err != nil {
				ch <- either.Either[[]*entities.Family]{Err: err}
				return
			}
			fs := make([]*entities.Family, len(familiesChunk.Value))
			for i := range fs {
				fs[i] = &entities.Family{
					Slug:   familiesChunk.Value[i],
					Origin: entities.NativeOrigin,
				}
			}
			ch <- either.Either[[]*entities.Family]{Value: fs}
		}
		close(ch)
	}()

	return ch
}

func deserializeRawSubfamily(in *engineLang.Subfamily) *entities.RawSubfamily {
	if in == nil {
		return nil
	}

	return &entities.RawSubfamily{
		Slug:              in.Slug,
		FamilySlug:        in.FamilySlug,
		ExtendedSubfamily: deserializeRawSubfamily(in.ExtendedSubfamily),
	}
}

func deserializeSubfamily(in *engineLang.Subfamily) *entities.Subfamily {
	if in == nil {
		return nil
	}

	return &entities.Subfamily{
		Slug:              in.Slug,
		FamilySlug:        in.FamilySlug,
		Origin:            entities.NativeOrigin,
		ExtendedSubfamily: deserializeRawSubfamily(in.ExtendedSubfamily),
	}
}

func (a *adapter) LoadNativeSubfamilies(context.Context) chan either.Either[[]*entities.Subfamily] {
	ch := make(chan either.Either[[]*entities.Subfamily])
	go func() {
		for subfamiliesChunk := range a.engine.LoadLanguageSubfamilies() {
			if err := subfamiliesChunk.Err; err != nil {
				ch <- either.Either[[]*entities.Subfamily]{Err: err}
				return
			}
			sfs := make([]*entities.Subfamily, len(subfamiliesChunk.Value))
			for i := range sfs {
				sfs[i] = deserializeSubfamily(subfamiliesChunk.Value[i])
			}
			ch <- either.Either[[]*entities.Subfamily]{Value: sfs}
		}
		close(ch)
	}()

	return ch
}

func (a *adapter) LoadNativeLanguages(context.Context) chan either.Either[*entities.Language] {
	ch := make(chan either.Either[*entities.Language])
	go func() {
		for languagesChunk := range a.engine.LoadLanguages() {
			if err := languagesChunk.Err; err != nil {
				ch <- either.Either[*entities.Language]{Err: err}
				return
			}
			ch <- either.Either[*entities.Language]{Value: &entities.Language{
				Slug:           languagesChunk.Value.Slug,
				FamilySlug:     languagesChunk.Value.Subfamily.FamilySlug,
				SubfamilySlug:  languagesChunk.Value.Subfamily.Slug,
				WordbaseSlug:   languagesChunk.Value.Wordbase.Slug,
				FemaleOwnNames: languagesChunk.Value.Wordbase.FemaleOwnNames,
				MaleOwnNames:   languagesChunk.Value.Wordbase.MaleOwnNames,
				Words:          languagesChunk.Value.Wordbase.Words,
				Min:            languagesChunk.Value.Wordbase.Min,
				Max:            languagesChunk.Value.Wordbase.Max,
				Dupl:           languagesChunk.Value.Wordbase.Dupl,
				M:              languagesChunk.Value.Wordbase.M,
				Origin:         entities.NativeOrigin,
			}}
		}
		close(ch)
	}()

	return ch
}
