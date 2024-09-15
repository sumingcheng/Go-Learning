## 使用须知

### 安装wire
```bash
go install github.com/google/wire/cmd/wire@latest
```

### 生成wire代码
```bash
cd di
wire
```

### go 代理源
```bash
go env -w GOPROXY=https://goproxy.cn,direct
```