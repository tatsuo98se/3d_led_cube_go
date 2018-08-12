package ledlib

type LedCanvas struct{
	stop bool
	Name string
}

func NewLedCanvas() *LedCanvas {
	u := new(LedCanvas)
	u.Name = "aaa"
    return u
}