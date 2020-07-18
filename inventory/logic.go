package inventory

import (
	"errors"
	"fmt"
	"time"

	"github.com/teris-io/shortid"

	errs "github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrEntryNotFound = errors.New("Entry not found")
	ErrEntryInvalid  = errors.New("Entry invalid")
)

type inventoryService struct {
	inventoryRepo InventoryRepository
}

func NewInventoryService(inventoryRepo InventoryRepository) InventoryService {
	return &inventoryService{
		inventoryRepo,
	}
}

func (i *inventoryService) Get(id string) (*Entry, error) {
	return i.inventoryRepo.Get(id)
}

func (i *inventoryService) Post(entry *Entry) error {
	fmt.Println(i)
	if err := validate.Validate(entry); err != nil {
		return errs.Wrap(err, "service.Entry.Post")
	}
	entry.ID = shortid.MustGenerate()
	entry.CreatedAt = time.Now().UTC().Unix()
	entry.ModifiedAt = time.Now().UTC().Unix()
	return i.inventoryRepo.Post(entry)
}

func (i *inventoryService) Update(entry *Entry) error {
	entry.ModifiedAt = time.Now().UTC().Unix()
	return i.inventoryRepo.Update(entry)
}

func (i *inventoryService) Delete(id string) error {
	return i.inventoryRepo.Delete(id)
}

func (i *inventoryService) GetAll() ([]*Entry, error) {
	return i.inventoryRepo.GetAll()
}
