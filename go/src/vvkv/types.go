package vvkv

/*
ListResultItem use
*/
type ListResultItem struct {
	Name         string `json:"name"`
	LastModified string `json:"lastModified"`
	Etag         string `json:"etag"`         // etag contains ", e.g.: "5B3C1A2E053D763E1B002CC607C5A0FE"
	Type         string `json:"type"`         // type, e.g.: Normal
	Size         int    `json:"size"`         // size, e.g.: 344606
	StorageClass string `json:"storageClass"` // "Standard" | "IA" | "Archive"
	Owner        struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
	} `json:"owner"`
}

/*
Meta use
*/
type Meta struct {
	codetype string
	isURL    bool
}

/*
URLType use
*/
type URLType struct {
	URL string `json:"url"`
}
