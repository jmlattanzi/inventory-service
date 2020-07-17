package inventory

type InventoryService interface {
	Get(id string) (*Entry, error)
	Post(entry *Entry) error
	Update(entry *Entry) error
	Delete(id string) error
}
