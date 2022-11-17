package entities

type Origin string

func (o Origin) IsZero() bool {
	return o == ""
}

const (
	NativeOrigin Origin = "native_origin"
	CustomOrigin Origin = "custom_origin"
)
