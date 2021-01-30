# mkpasswd
随机密码生成

# 示例
```
# 默认生成包含大小写字母和数字的8位长度密码
$ mkpasswd
dT4ne8j6

```

```
# 使用 -l 参数指定生成密码长度
$ mkpasswd -l 10
CXvjdWj9qZ
```

```
# 使用 -u 参数指定生成的密码无重复字符
$ mkpasswd -l 10 -u
XbN5Um7cMI
```

```
# 使用 -s 参数指定生成的密码样式，如果检测到包含大写字母、小写字母、数字、特殊字符则会在生成密码时使用这类字符
$ mkpasswd -s b1
qe30lflv

$ mkpasswd -s a3, -u
3u/5^wh@
```
