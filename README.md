# tools description
# 功能

```
1. util 基础工具（各种功能函数集合）, aes、base64、hmac、http、ip、md5、uuid、util
2. reids, 注册使用，封装了常用的 redis功能集
3. db, 注册使用，提供了drivers map方式，支持多DB
4. log, 提供标准日志集合，可输出标准日志，业务流水日志，grpc流水日志，包含，输入输出，耗时，错误等元素
5. errors, 提供标准错误定义，并包含通用错误定义，-1，1004 ，1005等，后续业务只需关注业务码，支持错误码转码操作
6. context, 建立输入输出标准，支持标准输入输出转换
7. prometheus，封装初始化和写入prometheus等函数。
```

```
2021-08-24 add by andre
提供高性能协程池 pool

```

```
2021-09-09 add by andre
提供双循环链表数据结构，双循环列表支持快速查询
```