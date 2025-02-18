### `grok()` {#fn-grok}

函数原型：`grok(input=required, pattern=required, trim_space=optional)`

函数说明：通过 `pattern` 提取文本串 `input` 中的内容。

参数:

- `input`：待提取文本，可以是原始文本（`_`）或经过初次提取之后的某个 `key`
- `pattern`: grok 表达式，表达式中支持指定 key 的数据类型：bool, float, int, string，默认为 string
- `trim_space`: 删除提取出的字符中的空白首尾字符，默认值为 true

```python
grok(_, pattern)    # 直接使用输入的文本作为原始数据
grok(key, pattern)  # 对之前已经提取出来的某个 key，做再次 grok
```

示例:

```python
# 待处理数据: "12/01/2021 21:13:14.123"

# pipline脚本
add_pattern("_second", "(?:(?:[0-5]?[0-9]|60)(?:[:.,][0-9]+)?)")
add_pattern("_minute", "(?:[0-5][0-9])")
add_pattern("_hour", "(?:2[0123]|[01]?[0-9])")
add_pattern("time", "([^0-9]?)%{_hour:hour:string}:%{_minute:minute:int}(?::%{_second:second:float})([^0-9]?)
grok(_, "%{DATE_US:date} %{time}")

# 处理结果
{
  "date": "12/01/2021",
  "hour": "21",
  "message": "12/01/2021 21:13:14.123",
  "minute": 13,
  "second": 14.123
}
```
