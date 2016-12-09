package lifelog

import (
	"time"
)

type (
	CommonSystemFields struct {
		CreatedOn  time.Time
		ModifiedOn time.Time `json:"ModifiedOn,omitempty"`
	}

	ErrorResponse struct {
		Error string `json:"error"`
	}
)
