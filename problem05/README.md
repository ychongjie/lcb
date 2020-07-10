请改进系统设计，允许⽤户更换⼿机时保留账号。

UserCenter 中增加更换绑定手机的接口, redis 中数据模型不变, 处理逻辑如下:

1. 判断 token 是否正确, 用户是否已登录
2. 在 phones 这个 Hash 中增加 新手机-用户id 的键值对
3. 修改该用户的用户信息对象 Hash 中的 phone 字段
4. 在 phones 这个 Hash 中删除 旧手机-用户id 的键值对