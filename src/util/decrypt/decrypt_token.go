package decrypt

import (
	"encoding/base64"
	"strings"

	"github.com/mohsanabbas/ticketing_utils-go/logger"
)

// Decrypt jwt token
func Decrypt(input string) []byte {
	b64data := input[strings.IndexByte(input, '.')+1:]
	decoded, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		logger.Error("Error: %v", err)
	}
	return decoded
}
