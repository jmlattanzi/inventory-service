package inventory

type InventoryRepository interface {
	Get(id string) (*Entry, error)
	Post(entry *Entry) error
	Update(entry *Entry, id string) error
	Delete(id string) error
	GetAll() ([]*Entry, error)
	GetCategory(category string) ([]*Entry, error)
}
