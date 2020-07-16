package Model

import (
	_ "database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

type Users struct {
	UserId			int				`json:"user_id"`
	UserName 		string			`json:"user_name"`
	Password		string			`json:"password"`
	Active 			bool			`json:"active"`
	Roles			Roles			`json:"role"`
	Permissions		[]string		`json:"permissions"`
}

type Roles struct {
	RoleName		string		`json:"role_name"`
}

type Permissions struct {
	PermissionsName	string		`json:"permissions_name"`
}

type UsersDTO struct {
	UserName 		string			`json:"user_name"`
	Active 			bool			`json:"active"`
	Roles			Roles			`json:"role"`
	Permissions		[]string		`json:"permissions"`
}

func GetAuthUser(username *string, password *string) (*UsersDTO, error) {
	var authUser UsersDTO
	result, err := Db.Query("select u.username, u.active, r.role_name, p.permissions_name from users u join user_role ur on u.user_id = ur.user_id join roles r on ur.role_id = r.role_id join role_permission rp on r.role_id = rp.role_id join permissions p on rp.permission_id = p.permissions_id where u.username = ? and u.password = ? and u.active = 1;", *username, *password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var permissions []string
	for result.Next() {
		var rawPermission string
		err = result.Scan(&authUser.UserName, &authUser.Active, &authUser.Roles.RoleName, &rawPermission)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		permissions = append(permissions, rawPermission)
	}

	authUser.Permissions = permissions
	return &authUser, nil
}