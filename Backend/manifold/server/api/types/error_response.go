package types

type ErrorResponse struct {
	Error          *string     `json:"error_text"`
	ErrorCode      *ErrorCode  `json:"error_code"`
	EntityId       *string     `json:"entity_id,omitempty"`
	ExpectedSchema interface{} `json:"expected_schema"`
}
