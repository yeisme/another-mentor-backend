# data 层

## 结构

应用服务层 <-> Repository层 <-> DAO层 <-> Model层 <-> 数据库

```bash
.data/
├── dao 基础数据操作（原子性操作）
├── model 数据模型（对应数据库表结构）
├── database 数据库配置
└── repository 业务仓储（组合多个DAO操作）
```

## 各层职责

### 1. model 层（数据模型）

- 定义数据库表映射结构
- 生成基础 CURD 方法
- 维护数据库字段与结构体映射关系

### 2. dao 层（数据访问对象）

- 实现跨表基础操作
- 封装复杂SQL查询
- 处理数据缓存逻辑

### 3. database 数据库配置

### 3. repository 层（业务仓储）

- 组合多个DAO操作
- 处理业务级事务
- 实现领域驱动设计中的仓储模式
