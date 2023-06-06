package tools

import (
	"encoding/hex"
	"fmt"
)

func CheckPasswordValidity(hashInDb string, userInput string) error {
	hashedUserInput := hex.EncodeToString(HashPassword(userInput))
	for i := range hashInDb {
		if  hashInDb[i] != hashedUserInput[i] {
		return fmt.Errorf("incorrect password")
		}
	} 
	return nil
}
