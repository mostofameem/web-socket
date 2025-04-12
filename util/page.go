package util

type Page struct {
	// Items are the list items
	Items interface{} `json:"items"`

	// ItemsPerPage is the number of items per page
	ItemsPerPage int `json:"itemsPerPage"`

	// PageNumber is the current page number
	PageNumber int `json:"pageNumber"`

	// TotalItems is the total number of items found
	//TotalItems int `json:"totalItems"`

	// TotalPages is the total number of pages available
	//TotalPages int `json:"totalPages"`
}
