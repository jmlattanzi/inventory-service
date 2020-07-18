package json

import (
	"encoding/json"

	"github.com/jmlattanzi/rex/services/inventory-service/inventory"
	"github.com/pkg/errors"
)

type Entry struct{}

func (e *Entry) Decode(input []byte) (*inventory.Entry, error) {
	entry := &inventory.Entry{}
	if err := json.Unmarshal(input, entry); err != nil {
		return nil, errors.Wrap(err, "serializer.Entry.Decode")
	}
	return entry, nil
}

func (e *Entry) Encode(entry *inventory.Entry) ([]byte, error) {
	raw, err := json.Marshal(entry)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Entry.Encode")
	}
	return raw, nil
}

func (e *Entry) EncodeMultiple(entries []*inventory.Entry) ([]byte, error) {
	raw, err := json.Marshal(entries)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Entry.EncodeMultiple")
	}
	return raw, nil
}
