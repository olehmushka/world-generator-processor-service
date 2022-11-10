package entities

import "github.com/olehmushka/world-generator-engine/influence"

type Influence string

const (
	StrongInfluence   Influence = Influence(influence.StrongInfluence)
	ModerateInfluence Influence = Influence(influence.ModerateInfluence)
	WeakInfluence     Influence = Influence(influence.WeakInfluence)
)
