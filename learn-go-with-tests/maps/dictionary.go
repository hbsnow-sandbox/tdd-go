package main

// Dictionary 辞書
type Dictionary map[string]string

// Error
const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

// DictionaryErr 辞書のエラー
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search Mapからkeyでvalueを取り出す
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add 辞書に追加する
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
