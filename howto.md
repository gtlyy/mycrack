# 如何使用远程仓库

## 0. 重新初始化项目模块
```
go mod init
```

## 1. 获取库代码
go get github.com/gtlyy/mycoin@v0.0.2

## 2. 检查模块路径一致性
go list -m all | grep mycoin
## 应输出：github.com/gtlyy/mycoin v0.0.2

## 3. 检查远程标签
git ls-remote --tags https://github.com/gtlyy/mycoin
## 应输出：0881e894bc4580df531fcd1707eb0eb71b1c0641	refs/tags/v0.0.2

## 4. 检查远程Go模块版本
go list -m github.com/gtlyy/mycoin
## 应输出：github.com/gtlyy/mycoin v0.0.2
