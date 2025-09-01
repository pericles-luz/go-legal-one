package legal_one_test

import (
	"encoding/json"
	"testing"

	"github.com/pericles-luz/go-legal-one/pkg/legal_one"
	"github.com/stretchr/testify/require"
)

func TestCaimResponseShouldParse(t *testing.T) {
	source := `{
  "id": 1,
  "natureId": 1,
  "claim": {
    "Id": 1,
    "Description": "string"
  },
  "claimObjects": [
    {
      "Id": 1,
      "ClaimObject": {
        "Id": 1,
        "Description": "string"
      }
    }
  ],
  "claimReasons": [
    {
      "Id": 1,
      "ClaimReason": {
        "Id": 1,
        "Description": "string"
      }
    }
  ],
  "contingency": "Active",
  "claimDate": "2025-08-31T14:43:18.611Z",
  "judgmentDate": "2025-08-31T14:43:18.611Z",
  "remarks": "string"
}`

	var result legal_one.ClaimResponse
	require.NoError(t, json.Unmarshal([]byte(source), &result))
}
