package main

/*
项目规范：
1.项目命名规范。
	小写和连接线组成
	例如：go-tutorial、go-web-server、redis-client 等。

2.标准项目布局
my-project/
├── cmd/                    # 可执行文件目录
│   ├── api/               # API 服务入口
│   │   └── main.go
│   └── cli/               # 命令行工具入口
│       └── main.go
├── internal/              # 私有代码（外部无法导入）
│   ├── config/           # 配置处理
│   ├── handler/          # HTTP 处理器
│   ├── service/          # 业务逻辑
│   └── repository/       # 数据访问层
├── pkg/                  # 公共库代码（可被外部导入）
│   ├── database/         # 数据库相关
│   ├── middleware/       # 中间件
│   └── utils/            # 工具函数
├── api/                  # API 定义文件
│   └── v1/
│       └── user.proto    # Protobuf 文件
├── web/                  # Web 静态资源
├── scripts/              # 构建和部署脚本
├── configs/              # 配置文件
├── deployments/          # 部署配置
├── docs/                 # 文档
├── go.mod
├── go.sum
├── Makefile
└── README.md

3.包命名规范
// 小写字母，简洁明了
package http
package json
package time

// 避免无意义名称
package lib          // 不好
package common       // 不好
package util         // 不好

// 使用有意义的单数名称
package user         // 好

4. 文件命名规范
# 小写字母，下划线分隔
user_service.go
database_handler.go
config_loader.go

# 测试文件
user_service_test.go
config_loader_test.go

# 基准测试
user_service_bench_test.go

5. 代码组织规范
// user_service.go 示例
package service

import (
    "context"
    "fmt"
    
    "github.com/my-project/internal/domain"
    "github.com/my-project/internal/repository"
)

// 类型定义
type UserService struct {
    repo repository.UserRepository
}

// 构造函数
func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

// 导出方法（首字母大写）
func (s *UserService) CreateUser(ctx context.Context, user *domain.User) error {
    // 业务逻辑
    if err := s.validateUser(user); err != nil {
        return err
    }
    return s.repo.Create(ctx, user)
}

// 私有方法（首字母小写）
func (s *UserService) validateUser(user *domain.User) error {
    if user.Name == "" {
        return fmt.Errorf("user name is required")
    }
    return nil
}

6. 测试
// 单元测试 (Unit Tests) 一般用于功能测试，验证代码是否正确
func TestFunction(t *testing.T)

// 基准测试 (Benchmark Tests) 一般用于性能测试
func BenchmarkFunction(b *testing.B)

// 示例测试 (Example Tests) 一般用于文档示例
func ExampleFunction()

// 模糊测试 (Fuzz Tests) - Go 1.18+
func FuzzFunction(f *testing.F)
*/
func main() {

}