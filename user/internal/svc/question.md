# svc 当前架构问题

```go
userRepo := repository.NewUserRepository()
userModel, err := userRepo.GetUserByID(queryUserId)
```

我发现我每个logic代码中如果涉及数据库交互，每个Logic独立创建Repository

- 每次调用都创建新实例，可能造成资源浪费
- 如果Repository初始化有开销（比如连接池设置），会影响性能
- 不便于统一配置和单元测试

> 如果通过 ServiceContext注入Repository，会解决这个问题

```go
type ServiceContext struct {
    Config config.Config
    // 添加Repository字段
    UserRepo repository.UserRepository
    // 其他依赖...
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config: c,
        // 初始化一次Repository
        UserRepo: repository.NewUserRepository(),
        // 其他初始化...
    }
}
```

- Repository只创建一次，避免资源浪费
- 便于统一管理配置
- 清晰的依赖关系，更易于测试
- 符合依赖注入的设计原则
- 符合go-zero框架的最佳实践

> 导致的新问题？

repository层代码过多，可能需要对repository层进行封装
