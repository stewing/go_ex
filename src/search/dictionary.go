package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("Could not find the word.")
	ErrWordExists       = DictionaryErr("The word already exists.")
	ErrWordDoesNotExist = DictionaryErr("The word doesn't eixst.")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	}

	return err
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
