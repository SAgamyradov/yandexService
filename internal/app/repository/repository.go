package repository

type URLRepository interface {
	GenerateShortURL(longURL string) (string, error)
	GetLongURL(shortURL string) (string, error)
}
