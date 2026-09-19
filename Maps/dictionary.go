package main

type Dictionary map[string]string

const (
	ErrNotFound              = DictionaryErr("could not find the word you were looking for")
	ErrWordExists            = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist      = DictionaryErr("cant perform the following action on word as it does not exist")
	ErrDeleteNonExistingWord = DictionaryErr("you can not delete the key that does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrDeleteNonExistingWord
	case nil:
		delete(d, key)
	default:
		return err
	}
	return nil
}
func (d Dictionary) Update(key, val string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = val
	default:
		return err
	}
	return nil
}

func (d Dictionary) Search(s string) (string, error) {
	definitions, ok := d[s]
	if !ok {
		return "", ErrNotFound
	}
	return definitions, nil
}

func (d Dictionary) Add(key, val string) error {

	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = val
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
