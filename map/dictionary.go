package dictionary

import (
	"errors"
)



var (
	//ErrNotFound if the word is not found
	ErrNotFound   = errors.New("could not find the word you were looking for")
	//ErrWordExists word already exist
	ErrWordExists = errors.New("cannot add word because it already exists")
	//ErrWordDoesNotExist word does not exist
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)
//Dictionary type
type Dictionary map[string]string
//DictionaryErr type
type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}
//Search for a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }

    return definition, nil
}
//Add a word in the dictionary
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

//Update a word in the dictionary
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

//Delete a word in the dictionary
func (d Dictionary) Delete(word string) {
    delete(d, word)
}