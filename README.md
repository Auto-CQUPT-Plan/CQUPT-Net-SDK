# CQUPT Net SDK

> **本项目为 *Auto CQUPT Plan* 的一部分 **

![images](./images/cover.png)

## 1. 项目简介

`CQUPT Net SDK` 是由 `Golang` 编写的 **重庆邮电大学校园网SDK**，实现了校园网的**登录**，**登出**，**登录记录查询**，**在线设备查询**等功能。

## 2. 使用示例

> [!note]
>
> 值得注意的是登录时不一定要在路由器上执行此SDK的方法，只需在路由器局域网内任意一台设备上执行即可。

### 2.1 基础配置

```go
package CQUPT_Net_SDK

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/bytedance/sonic"
)

func genBasicData() *BasicData {
	return &BasicData{
		StuID:    "统一认证码", // 统一认证码
		Password: "密码", // 密码
		UA:       "desktop", // 设备类型，支持 desktop，phone，pad
		Isp:      "cmcc", // 运营商, 支持 telecom，cmcc，unicom，xyw
		IPAddr:   "IP", // 设备IP
		MACAddr:  "000000000000", // 路由器MAC，全部置0的话AuthServer会自动识别
	}
}
```

| 字段     | 备注       | 取值                                               |
| -------- | ---------- | -------------------------------------------------- |
| StuID    | 统一认证码 | 统一认证码                                         |
| Password | 密码       | 密码                                               |
| UA       | 设备类型   | desktop(电脑)，phone(手机)，pad(平板)              |
| Isp      | 运营商     | telecom(电信)，cmcc(移动)，unicom(联通)，xyw(教师) |
| IPAddr   | 设备IP     | 设备IP                                             |
| MACAddr  | 路由器MAC  | 路由器MAC，全部置0的话AuthServer会自动识别         |

### 2.2 校园网登录

**单元测试代码：**

```go
func TestSDK_NetLogin(t *testing.T) {
	sdk := NewSDK()

	data := genBasicData()

	res, err := sdk.NetLogin(data)
	if err != nil {
		t.Error(err)
	}

	dis, _ := sonic.Marshal(res)

	var out bytes.Buffer
	_ = json.Indent(&out, dis, "", "  ")

	t.Logf("JSON: %v", out.String())
	t.Logf("Raw: %v", res)
}
```

### 2.3 校园网登出

**单元测试代码：**

```go
func TestSDK_NetLogout(t *testing.T) {
	sdk := NewSDK()

	data := genBasicData()

	res, err := sdk.NetLogout(data)
	if err != nil {
		t.Error(err)
	}

	dis, _ := sonic.Marshal(res)

	var out bytes.Buffer
	_ = json.Indent(&out, dis, "", "  ")

	t.Logf("JSON: %v", out.String())
	t.Logf("Raw: %v", res)
}
```

### 2.4 检测认证状态

**单元测试代码：**

```go
func TestSDK_NetChecker(t *testing.T) {
	sdk := NewSDK()

	data := genBasicData()

	res, err := sdk.NetChecker(data)
	if err != nil {
		t.Error(err)
	}

	dis, _ := sonic.Marshal(res)

	var out bytes.Buffer
	_ = json.Indent(&out, dis, "", "  ")

	t.Logf("JSON: %v", out.String())
	t.Logf("Raw: %v", res)
}
```

### 2.5 获取上网服务信息

**单元测试代码：**

```go
func TestSDK_GetNetServiceInfo(t *testing.T) {
	sdk := NewSDK()

	data := genBasicData()

	res, err := sdk.GetNetServiceInfo(data)
	if err != nil {
		t.Error(err)
	}

	dis, _ := sonic.Marshal(res)

	var out bytes.Buffer
	_ = json.Indent(&out, dis, "", "  ")

	t.Logf("JSON: %v", out.String())
	t.Logf("Raw: %v", res)
}
```

### 2.6 获取登录记录

**单元测试代码：**

```go
func TestSDK_GetLoginHistory(t *testing.T) {
	sdk := NewSDK()

	data := genBasicData()

	res, err := sdk.GetLoginHistory(data)
	if err != nil {
		t.Error(err)
	}

	dis, _ := sonic.Marshal(res)

	var out bytes.Buffer
	_ = json.Indent(&out, dis, "", "  ")

	t.Logf("JSON: %v", out.String())
	t.Logf("Raw: %v", res)
}
```

### 2.7 获取当前登录列表

**单元测试代码：**

```go
func TestSDK_GetOnlineList(t *testing.T) {
	sdk := NewSDK()

	data := genBasicData()

	res, err := sdk.GetOnlineList(data)
	if err != nil {
		t.Error(err)
	}

	dis, _ := sonic.Marshal(res)

	var out bytes.Buffer
	_ = json.Indent(&out, dis, "", "  ")

	t.Logf("JSON: %v", out.String())
	t.Logf("Raw: %v", res)
}
```

> [!warning]
>
> **此SDK仅供学习使用，请勿用于非法用途，否则造成的危害后果自负**
