package totp

import (
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
)

func Generate(secret string) (string, error) {
	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", fmt.Errorf("failed to generate TOTP: %w", err)
	}
	return code, nil
}
