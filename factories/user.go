package factories

import (
	"Gintest/src/entities"
	"Gintest/src/utilities"
)

var password = "test"

func GetUser() entities.User {

	hash_utility := utilities.Hash{}
	hashedPwd, _ := hash_utility.Make(password)

	return entities.User{ID: 1, Email: "example@giftpakc.ai", Password: hashedPwd}
}

func GetUserCredential() entities.UserCredential {

	return entities.UserCredential{Email: "example@giftpakc.ai", Password: password}
}

func GetUserRegisteration() entities.UserRegisteration {

	return entities.UserRegisteration{Email: "example@giftpakc.ai", Password: password}
}

func GetUserUpdation() entities.UserUpdation {

	return entities.UserUpdation{}
}
