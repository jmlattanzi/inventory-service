package inventory

// type Entry struct {
// 	ID         string   `json:"id" bson:"id"`
// 	Name       string   `json:"name" bson:"name"`
// 	Category   string   `json:"category" bson:"category"`
// 	Tags       []string `json:"tags" bson:"tags"`
// 	ImageURL   string   `json:"image_url" bson:"image_url"`
// 	CreatedAt  int64    `json:"created_at" bson:"created_at"`
// 	ModifiedAt int64    `json:"modified_at" bson:"modified_at"`
// }

type Entry struct {
	ID             string   `json:"id" bson:"id"`
	Publisher      string   `json:"publisher" bson:"publisher"`
	Title          string   `json:"title" bson:"title"`
	Issue          string   `json:"issue" bson:"issue"`
	Condition      string   `json:"condition" bson:"condition"`
	CoverPrice     float32  `json:"cover_price" bson:"cover_price"`
	Quantity       int      `json:"quantity" bson:"quantity"`
	Total          float32  `json:"total" bson:"total"`
	CoverVariation string   `json:"cover_variation" bson:"cover_variation"`
	Tags           []string `json:"tags" bson:"tags"`
	Category       string   `json:"category" bson:"category"`
	ImageURL       string   `json:"image_url" bson:"image_url"`
	CreatedAt      int64    `json:"created_at" bson:"created_at"`
	ModifiedAt     int64    `json:"modified_at" bson:"modified_at"`
}
