package util

type (
	Util struct{}
)

func NewUtil() IUtil {
	return &Util{}
}
