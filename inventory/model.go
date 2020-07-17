package inventory

type Entry struct {
	ID         string   `json:"id" bson:"id"`
	Name       string   `json:"name" bson:"name"`
	Category   string   `json:"category" bson:"category"`
	Tags       []string `json:"tags" bson:"tags"`
	ImageURL   string   `json:"image_url" bson:"image_url"`
	CreatedAt  int64    `json:"created_at" bson:"created_at"`
	ModifiedAt int64    `json:"modified_at" bson:"modified_at"`
}
