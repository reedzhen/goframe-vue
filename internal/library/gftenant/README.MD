# GFTENANT 是一个针对Goframe框架的多租户引擎

项目介绍:
- 目前引擎相对功能封装可能不太完善，后续会持续完善。 
- 支持 column模式和db模式

注意事项: 
- 全插件禁止使用 `g.DB()`
- 当前只支持 mysql

## 配置文件

### 1.gf gen dao 配置文件
```
  gen:
    dao:
      - link: "mysql:root:123456@tcp(127.0.0.1:3306)/tenant_c?loc=Local&parseTime=true"
        descriptionTag: true
        noModelComment: true
        jsonCase: "CamelLower"
        gJsonSupport: true
        group: "default"
        tables: "sys_base"

      - link: "mysql:root:123456@tcp(127.0.0.1:3306)/tenant_c?loc=Local&parseTime=true"
        descriptionTag: true
        noModelComment: true
        jsonCase: "CamelLower"
        gJsonSupport: true
        group: "tenant"
        tplDaoInternalPath: "./hack/tenant.tpl"
        tables: "sys_tenant"
```

### 2. 项目配置文件
```
database:
  default:
    link: "mysql:root:123456@tcp(127.0.0.1:3306)/tenant_c?charset=utf8&parseTime=True&loc=Local"
    debug: true
    maxIdle: 10
    maxOpen: 50
  tenant:
    link: "tenant_mysql:root:123456@tcp(127.0.0.1:3306)/tenant_c?charset=utf8&parseTime=True&loc=Local"
    debug: true
    maxIdle: 10
    maxOpen: 50
```

## 安装

```
go get -u github.com/gogf/gf-tenant
```

### 1.项目引入(初始化MYSQL驱动)
```
_ "gitee.com/cleango/gftenant" 
```

### 2. 配置文件

```
tenant:
  mode: "column"
```

### 3. 注入多租户标识

```
ctx = gftenant.SetTenant(ctx, "tenant_c")
```
### 4.当为db模式时，需要修改gf gen dao的模板

> 模版是项目目录中的tenant.tpl