package inventory

type InventorySerializer interface {
	Decode(input []byte) (*Entry, error)
	Encode(input *Entry) ([]byte, error)
	EncodeMultiple(input []*Entry) ([]byte, error)
}
