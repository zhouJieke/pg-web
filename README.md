# 基于Gin-Admin初始化模板
### 数据库：
1. Mysql
2. Redis 
### 已集成以下功能：
1. JsonWebToken（自动续约、验证）
2. 日志管理

### 待集成功能：
1. RBAC权限管理
2. 路由菜单
3. 租户管理

# 核心目录介绍
```azur
├── app
│ └── business # 业务层
├── bootstrap # 服务和日志
├── global # 全局访问常量 
├── model # Mysql模型
├── resource # 配置文件目录
├── router # 路由
├── runtime # 日志文件
└── utils # 工具包
```