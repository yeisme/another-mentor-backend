# middleware 问题

1. adminauthmiddleware 和 userauthmiddleware 仅有用户身份判读部分不同，能否通过输入参数统一为同样的代码，在 servicecontext.go 中，通过输入不同参数 UserAuth:  middleware.NewUserAuthMiddleware(c).Handle 实现
