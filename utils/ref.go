package utils

func Ref[O any](o O) *O {
	return &o
}
