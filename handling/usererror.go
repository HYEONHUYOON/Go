package main

import (
	"fmt"
	"math"
)

// Error 로 사용될 수 있다.
type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}
	return nil
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 양수여야 합니다.")
	}
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := sqrt(-2)
	if err != nil {
		fmt.Printf("Error : %v\n", err)
	} else {
		fmt.Println(sqrt)
	}

	err2 := RegisterAccount("hhmykh234", "dbs001010!A")
	if err2 != nil {
		if errInfo, ok := err2.(PasswordError); ok {
			fmt.Printf("%v Len : %d RequireLen : %d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		} else {
			fmt.Println("회원가입 완료")
		}
	}
}
