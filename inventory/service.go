package inventory

type InventoryService interface {
	Get(id string) (*Entry, error)
	Post(entry *Entry) error
	Update(entry *Entry, id string) error
	Delete(id string) error
	GetAll() ([]*Entry, error)
	GetCategory(cat string) ([]*Entry, error)
}
