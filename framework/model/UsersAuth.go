package model

import (
	"fmt"
	"log"
)

type Users struct {
	UserId			int				`json:"user_id"`
	UserName 		string			`json:"user_name"`
	Password		string			`json:"password"`
	Active 			bool			`json:"active"`
	Role			Role			`json:"role"`
	Permissions		[]Permission	`json:"permissions"`
}

type Role struct {
	RoleName		string		`json:"role_name"`
}

type Permission struct {
	PermissionsName	string		`json:"permission_name"`
}

type UsersDTO struct {
	UserName 		string			`json:"user_name"`
	Active 			bool			`json:"active"`
	Roles			Role			`json:"role"`
	Permissions		[]Permission	`json:"permissions"`
}

func GetAuthUser(username *string, password *string) (*UsersDTO, error) {
	var authUser UsersDTO
	result, err := Db.Query("SELECT u.username, u.active, r.role_name, p.permissions_name FROM users u " +
								"INNER JOIN user_role ur on u.user_id = ur.user_id " +
								"INNER JOIN roles r on ur.role_id = r.role_id " +
								"INNER JOIN role_permission rp on r.role_id = rp.role_id " +
								"INNER JOIN permissions p on rp.permission_id = p.permissions_id " +
								"WHERE u.username = ? and u.password = ? and u.active = 1;", *username, *password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var permissions []Permission
	for result.Next() {
		var rawPermission Permission

		err = result.Scan(&authUser.UserName, &authUser.Active, &authUser.Roles.RoleName, &rawPermission.PermissionsName)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		permissions = append(permissions, rawPermission)
	}

	authUser.Permissions = permissions
	return &authUser, nil
}