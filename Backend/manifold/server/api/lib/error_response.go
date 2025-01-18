package lib

import (
	"encoding/json"
	"fmt"
	t "gitlab.com/manifold555112/manifold/server/api/types"
)

func ApiErroResponse(errText *string, errCode *t.ErrorCode, schema interface{}) (responseBody string, statusCode int) {
	bresponse, uerr := json.MarshalIndent(
		t.ErrorResponse{
			Error:          errText,
			ErrorCode:      errCode,
			ExpectedSchema: schema,
		},
		"",
		"  ",
	)

	if uerr != nil {
		return fmt.Sprintf(`{"error_text": "%v", "error_code":"%v"}`, errText, errCode), 418
	}
	return string(bresponse), 400
}
