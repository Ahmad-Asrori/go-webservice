package util

import (
	"fmt"
	"github.com/casbin/casbin"
)

func IsAuthorize(name string, resource string, method string) (bool, error) {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		fmt.Println(err)
	}

	name = "APP-Mobile"
	resource = "/"
	method = "GET"
	ok, err := e.Enforce(name, resource, method)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if ok {
		return true, nil
	}

	return false, nil
}
