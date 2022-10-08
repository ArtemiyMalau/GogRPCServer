package domain

type Document struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Hash string `json:"hash" bson:"hash"`
}

type InsertDocumentDTO struct {
	Url string `json:"url"`
}
