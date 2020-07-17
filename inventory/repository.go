package inventory

type InventoryRepository interface {
	Get(id string) (*Entry, error)
	Post(entry *Entry) error
	Update(entry *Entry) error
	Delete(id string) error
}
