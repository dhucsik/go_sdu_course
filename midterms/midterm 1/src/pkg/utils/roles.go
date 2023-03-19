package utils

import (
	"fmt"
	"module/pkg/repository"
)

func VerifyRole(role string) (string, error) {
	switch role {
	case repository.AdminRoleName:
		// Nothing to do, verified successfully
	case repository.SellerRoleName:
		// Nothing to do, verified successfully
	case repository.UserRoleName:
		// Nothing to do, verified successfully
	default:
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}
