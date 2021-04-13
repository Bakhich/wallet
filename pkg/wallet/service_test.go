package wallet

import (
	"testing"
)

func TestService_RegisterAccount_success(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("+992000000001")

	account, err := svc.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_FindAccoundByIdmethod_notFound(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("+992000000001")

	account, err := svc.FindAccountByID(2)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}
