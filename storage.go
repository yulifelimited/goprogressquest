package main

// Storage defines a storage medium. It could be anything that implements this interface.
type Storage interface {
	Save(Character) error
	Load(ID string) (Character, error)
	Update(Character) error
}
