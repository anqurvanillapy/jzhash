# jzhash

**暨珠哈希算法**, 用于举办活动前确定好日期后, 随机地生成活动地点,
并保证所有成员结果唯一, 并且加入不可预测因素 (活动当日才能获得输出). 适用于在
*暨南大学珠海校区饮食服务中心二楼* 举办的活动.

## 安装

- 直接下载 [Releases](https://github.com/anqurvanillapy/jzhash/releases)
中的二进制文件
- 或从源码编译安装

```bash
$ go get -u github.com/anqurvanillapy/jzhash
```

## 使用

- 如果是源码编译安装得来, 则直接使用命令

```bash
$ jzhash
```

- 如果是 Linux/macOS 用户, 可以将文件移动到 `$PATH` 含有的路径下使用
- 如果是 Windows 用户, 除了将文件移动到环境变量含有的路径下, 还可以编写一个 BAT
文件方便鼠标操作

```bat
rem 先保证二进制文件路径正确
jzhash.exe
pause
```

- 可能的输出结果

```
1,2
```

- 则从食堂外侧直接上食堂二楼的楼梯口, 将离楼梯口最近的边角餐桌作为原点, `1`
代表向前数第 1 行, `2` 代表向左数第 2 列.

## 依赖

- 二进制文件跨平台无依赖
- 开发者需要 Go 1.8+

## 原理

参考了 [Geohashing](https://www.xkcd.com/426/) 算法.  具体的操作是将格式为
`2006-01-02` 的 *日期* 和活动当日 [腾讯控股](http://gu.qq.com/hk00700/gp)
的 *开盘价* 用 `-` 号连接, 进行 SHA512 算法获得字节数组,
从中截断后两部分转换为食堂餐桌对应的行数和列数.

## License

MIT
