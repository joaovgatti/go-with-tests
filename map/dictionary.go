package _map

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrWordExists = DictionaryErr("could not find the world you were looking for")
	ErrNotFound =  DictionaryErr("could not find the key in the map")
	ErrWordDoesNotExists = DictionaryErr("could not update key because it does not exists")
)

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok{
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error{
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value

	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error{
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d,key)

}

