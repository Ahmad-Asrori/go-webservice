## Go RESTful

#### Native
URL Endpoints
```text
GET 128.199.110.107:3001/auth

JSON Request
{
    "user_name": "APP-Mobile",
    "password": "APM-Password"
}

JSON Response
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7InVzZXJfbmFtZSI6IkFQUC1Nb2JpbGUiLCJhY3RpdmUiOnRydWUsInJvbGUiOnsicm9sZV9uYW1lIjoiQVBQIn0sInBlcm1pc3Npb25zIjpbIkdFVCIsIlBPU1QiLCJQVVQiLCJERUxFVEUiXX0sImV4cCI6MTU5NDgwNDEwN30.lRkBAR-4b0sLYn4YkuqJb5bbLWVmSZKnFfpSYAIjMkE"
}
```

```text
POST 128.199.110.107:3001/account/insert

JSON Request
{
    "first_name": "ahmad",
    "last_name": "asrori",
    "email": "ahmad.asrori08@gmail.com",
    "password": "p123"
}

JSON Response
{
    "message": "success"
}
```

```text
GET 128.199.110.107:3001/account

JSON Request
{
    "account_id": "BCGEEYRHKBW"
}

JSON Response
{
	"account_id": "BCGEEYRHKBW",
	"first_name": "ahmad",
	"last_name": "asrori",
	"email": "ahmad.asrori08@gmail.com"
}
```

```text
GET 128.199.110.107:3001/account/all

JSON Response
[
	{
		"account_id": "BCGEGKDHKBW",
		"first_name": "surya",
		"last_name": "dinata",
		"email": "surya.dinata@gmail.com"
	},
	{
		"account_id": "BRKDHHURIFY",
		"first_name": "ahmad",
		"last_name": "asrori",
		"email": "ahmad.asrori08@gmail.com"
	}
]
```


```text
PUT 128.199.110.107:3001/account/update

JSON Request
{
    "account_id": "BCGEEYRHKBW",
    "first_name": "ahmad",
    "last_name": "asrori"
}

JSON Response
{
    "message": "success"
}
```

```text
DELETE 128.199.110.107:3001/account/delete

JSON Request
{
    "account_id": "BCGEEYRHKBW",
}

JSON Response
{
    "message": "success"
}
```

NB : khusus untuk RESTful Webservice yang dibuat dengan native,
user APP-Mobile tidak dapat melakukan DELETE HTTP Request ke URL
_128.199.110.107:3001/account/delete_. karena authorization telah diatur
sedemikian rupa.
untuk authorization dapat diatur sesuai kebutuhan, ada berbagai macam
jenis authorization seperti ACL, RBAC, namun terimakasih untuk casbin karena
mereka memberikan jenis baru authorization yaitu REST.
berikut adalah tabel authorization untuk setiap user.

User        | Password      | Resource  | Method        
------------|---------------|-----------|------------------------
APP-Mobile  | APM-Password  | /*        | GET, POST, PUT
APP-Web     | APW-Password  | /*        | GET, POST, PUT, DELETE
APP-Desktop | APD-Password  | /*        | GET, POST, PUT, DELETE

#### Framework
URL Endpoints
```text
GET 128.199.110.107:3000/authentication

JSON Request
{
    "user_name": "APP-Mobile",
    "password": "$2y$12$aXdvFMVniT1NLGH3X0JTr.pcBy/8yuXI0ykIT87Ixq6M357gZep0S"
}

JSON Response
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7InVzZXJfbmFtZSI6IkFQUC1Nb2JpbGUiLCJhY3RpdmUiOnRydWUsInJvbGUiOnsicm9sZV9uYW1lIjoiQVBQIn0sInBlcm1pc3Npb25zIjpbIkdFVCIsIlBPU1QiLCJQVVQiLCJERUxFVEUiXX0sImV4cCI6MTU5NDgwNDEwN30.lRkBAR-4b0sLYn4YkuqJb5bbLWVmSZKnFfpSYAIjMkE"
}
```

```text
POST 128.199.110.107:3000/mobile/user

JSON Request
{
    "first_name": "ahmad",
    "last_name": "asrori",
    "email": "ahmad.asrori08@gmail.com",
    "password": "p123"
}

JSON Response
{
    "message": "success"
}
```

```text
GET 128.199.110.107:3000/mobile/user

JSON Request
{
    "account_id": "BCGEEYRHKBW"
}

JSON Response
{
	"account_id": "BCGEEYRHKBW",
	"first_name": "ahmad",
	"last_name": "asrori",
	"email": "ahmad.asrori08@gmail.com"
}
```

```text
GET 128.199.110.107:3000/mobile/users

JSON Response
[
	{
		"account_id": "BCGEGKDHKBW",
		"first_name": "surya",
		"last_name": "dinata",
		"email": "surya.dinata@gmail.com"
	},
	{
		"account_id": "BRKDHHURIFY",
		"first_name": "ahmad",
		"last_name": "asrori",
		"email": "ahmad.asrori08@gmail.com"
	}
]
```


```text
PUT 128.199.110.107:3000/mobile/user

JSON Request
{
    "account_id": "BCGEEYRHKBW",
    "first_name": "ahmad",
    "last_name": "asrori"
}

JSON Response
{
    "message": "success"
}
```

```text
DELETE 128.199.110.107:3001/mobile/user

JSON Request
{
    "account_id": "BCGEEYRHKBW",
}

JSON Response
{
    "message": "success"
}
```
