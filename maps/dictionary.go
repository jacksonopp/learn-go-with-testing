package main

type Dictionary map[string]string

const (
	ErrNotFound                    = DictionaryErr("could not find the word you are looking for")
	ErrWordExists                  = DictionaryErr("word already exists")
	ErrWordDoesNotExist            = DictionaryErr("word does not exist")
	ErrCannotDeleteNonExistantWord = DictionaryErr("cannot delete word that does not exist in dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	// If the error isn't found, add it
	case ErrNotFound:
		d[word] = definition
		// If the word already exists, tell the user
	case nil:
		return ErrWordExists
		// Otherwise return nil
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrCannotDeleteNonExistantWord
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
