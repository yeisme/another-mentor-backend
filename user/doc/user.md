### 1. 用户列表

1. route definition

- Url: /api/v1/admin/users
- Method: GET
- Request: `UserListRequest`
- Response: `UserListResponse`

2. request definition



```golang
type UserListRequest struct {
	Page int `form:"page,default=1"`
	PageSize int `form:"pageSize,default=20"`
	Keyword string `form:"keyword,optional"`
	Status int `form:"status,optional"`
}
```


3. response definition



```golang
type UserListResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Total int64 `json:"total"`
	List []User `json:"list"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 2. 用户登录

1. route definition

- Url: /api/v1/user/login
- Method: POST
- Request: `LoginRequest`
- Response: `LoginResponse`

2. request definition



```golang
type LoginRequest struct {
	Account string `json:"account" validate:"required"` // 用户名或邮箱
	Password string `json:"password" validate:"required"`
}
```


3. response definition



```golang
type LoginResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Token string `json:"token,optional"`
	ExpiresIn int64 `json:"expiresIn,optional"`
	User User `json:"user,optional"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type User struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Phone string `json:"phone,optional"`
	Avatar string `json:"avatar,optional"`
	Nickname string `json:"nickname,optional"`
	Introduction string `json:"introduction,optional"`
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
	Status int `json:"status"`
}
```

### 3. 用户注册

1. route definition

- Url: /api/v1/user/register
- Method: POST
- Request: `RegisterRequest`
- Response: `RegisterResponse`

2. request definition



```golang
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Email string `json:"email" validate:"required,email"`
	Code string `json:"code" validate:"required"`
}
```


3. response definition



```golang
type RegisterResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	UserId int64 `json:"userId,optional"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 4. 重置密码

1. route definition

- Url: /api/v1/user/resetPassword
- Method: POST
- Request: `ResetPasswordRequest`
- Response: `ResetPasswordResponse`

2. request definition



```golang
type ResetPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
	Code string `json:"code" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}
```


3. response definition



```golang
type ResetPasswordResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 5. 发送验证码

1. route definition

- Url: /api/v1/user/sendVerifyCode
- Method: POST
- Request: `SendVerifyCodeRequest`
- Response: `SendVerifyCodeResponse`

2. request definition



```golang
type SendVerifyCodeRequest struct {
	Email string `json:"email" validate:"required,email"`
	Type string `json:"type" validate:"required,oneof=register reset"`
}
```


3. response definition



```golang
type SendVerifyCodeResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 6. 获取用户信息

1. route definition

- Url: /api/v1/user/info
- Method: GET
- Request: `GetUserInfoRequest`
- Response: `GetUserInfoResponse`

2. request definition



```golang
type GetUserInfoRequest struct {
	UserId int64 `form:"userId,optional"` // 为空时返回当前登录用户信息
}
```


3. response definition



```golang
type GetUserInfoResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	User User `json:"user,optional"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type User struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Phone string `json:"phone,optional"`
	Avatar string `json:"avatar,optional"`
	Nickname string `json:"nickname,optional"`
	Introduction string `json:"introduction,optional"`
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
	Status int `json:"status"`
}
```

### 7. 更新用户信息

1. route definition

- Url: /api/v1/user/info
- Method: PUT
- Request: `UpdateUserInfoRequest`
- Response: `UpdateUserInfoResponse`

2. request definition



```golang
type UpdateUserInfoRequest struct {
	Nickname string `json:"nickname,optional"`
	Avatar string `json:"avatar,optional"`
	Introduction string `json:"introduction,optional"`
	Phone string `json:"phone,optional"`
}
```


3. response definition



```golang
type UpdateUserInfoResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 8. 修改密码

1. route definition

- Url: /api/v1/user/password
- Method: PUT
- Request: `ChangePasswordRequest`
- Response: `ChangePasswordResponse`

2. request definition



```golang
type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,min=6,max=20"`
}
```


3. response definition



```golang
type ChangePasswordResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

