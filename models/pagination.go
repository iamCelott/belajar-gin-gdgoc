package models

type Pagination struct {
	TotalRecords int64 `json:"total_records"` // Total number of records
	TotalPages   int   `json:"total_pages"`   // Total number of pages
	CurrentPage  int   `json:"current_page"`  // Current page
	PageSize     int   `json:"page_size"`     // Number of records per page
}
