package translator

type Translator interface {
	Translate(source string, target string, q string) (*string, error)
}
