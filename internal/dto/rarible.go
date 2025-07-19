package dto

type CollectionTrait struct {
	CollectionID string          `json:"collectionId" validate:"required"`
	Properties   []TraitProperty `json:"properties" validate:"required"`
}

type TraitProperty struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
