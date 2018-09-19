package ledlib

type LedObject interface {
	DidDetach()
	Draw(canvas LedCanvas)
}
