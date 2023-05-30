package tools

import (
	"fmt"
)

func CheckPasswordValidity(hashInDb string, userInput string) error {
	
	hashedUserInput := HashPassword(userInput)

	for i := range hashInDb {
		if hashInDb[i] == hashedUserInput[i] {
			return nil
		} else if  hashInDb[i] != hashedUserInput[i] {
		return fmt.Errorf("incorrect password")
	}
} 
return nil
}
