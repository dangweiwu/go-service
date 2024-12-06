---
theme: orange
---
```run
window.SetConfig(
    "http://127.0.0.1",
    {"Authorization":""}
)
```
# 

- 版本
# 1 用户管理

系统用户管理 增删查改

## 1.1 创建用户

> 基础信息

- **PATH: /api/admin**
- **METHOD: POST**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|
>  FORM 参数 

| 参数 | 说明 | 类型 | 校验 | 示例 | 备注 |
| --- | --- | -- | -- | -- |  --  |
|account|账号|string|required|admin||
|phone|手机号|string|max=11|123456789||
|name|姓名|string|max=100|张三||
|status|状态|string|oneof=0 1|1|0 无效 1 有效|
|password|密码|string|max=100,required|123456||
|memo|备注|string|max=300|||
|email|Email|string|omitempty,email|||
|is_super_admin|是否超级管理员|string|oneof=0 1|1|0:否 1:是|
|role|角色ID|string|max=100|||

```button
var req = {

    Url:"/api/admin",
    Method:"POST",
    Header:{
        Authorization:"",
    },
    Form:{
        account:"admin2",
        phone:"123456789",
        name:"张三",
        status:"1",
        password:"123456",
        memo:"",
        email:"",
        is_super_admin:"1",
        role:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||string|ok|成功|



---