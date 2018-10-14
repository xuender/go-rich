<a name="0.1.11"></a>
## [0.1.11](https://github.com/xuender/go-rich/compare/v0.1.10...v0.1.11) (2018-10-14)


### Bug Fixes

* **index:** 页面不缓存 ([a910b96](https://github.com/xuender/go-rich/commit/a910b96)), closes [#45](https://github.com/xuender/go-rich/issues/45)
* **trades:** 订单页面没选择客户也会查询客户信息 ([fe99537](https://github.com/xuender/go-rich/commit/fe99537))


### Features

* **console:** 优化控制台输出，显示新用户初始帐号 ([17c2538](https://github.com/xuender/go-rich/commit/17c2538)), closes [#46](https://github.com/xuender/go-rich/issues/46) [#47](https://github.com/xuender/go-rich/issues/47)



<a name="0.1.10"></a>
## [0.1.10](https://github.com/xuender/go-rich/compare/v0.1.9...v0.1.10) (2018-10-13)


### Bug Fixes

* **console:** 控制台显示手机操作地址二维码 ([f59914b](https://github.com/xuender/go-rich/commit/f59914b)), closes [#39](https://github.com/xuender/go-rich/issues/39)
* **items:** 修复商品选择页面标签显示错误 ([2896578](https://github.com/xuender/go-rich/commit/2896578)), closes [#42](https://github.com/xuender/go-rich/issues/42)
* **obj:** 所有对象新增修改前Trim掉空格 ([9d4473a](https://github.com/xuender/go-rich/commit/9d4473a)), closes [#43](https://github.com/xuender/go-rich/issues/43)
* **windows:** rich.bat 换行符错误 ([84726e3](https://github.com/xuender/go-rich/commit/84726e3)), closes [#41](https://github.com/xuender/go-rich/issues/41)


### Features

* **api:** 规范API返回的状态码 ([0952047](https://github.com/xuender/go-rich/commit/0952047)), closes [#36](https://github.com/xuender/go-rich/issues/36)
* **bak:** 增加-z参数，备份数据库 ([7453779](https://github.com/xuender/go-rich/commit/7453779)), closes [#40](https://github.com/xuender/go-rich/issues/40)
* **console:** 控制台显示服务器IP及端口号 ([d9d424a](https://github.com/xuender/go-rich/commit/d9d424a)), closes [#44](https://github.com/xuender/go-rich/issues/44)
* **login:** 用户初始密码设置成版本号 ([d4be93d](https://github.com/xuender/go-rich/commit/d4be93d)), closes [#38](https://github.com/xuender/go-rich/issues/38)
* **login:** 登录时根据新用户状态显示帐号密码 ([e3d4a22](https://github.com/xuender/go-rich/commit/e3d4a22)), closes [#37](https://github.com/xuender/go-rich/issues/37)



<a name="0.1.9"></a>
## [0.1.9](https://github.com/xuender/go-rich/compare/v0.1.8...v0.1.9) (2018-10-12)


### Bug Fixes

* **交易, 商品:** 交易页面新增商品页面有循环依赖 ([0edbb00](https://github.com/xuender/go-rich/commit/0edbb00)), closes [#35](https://github.com/xuender/go-rich/issues/35)


### Features

* **安全:** 新用户引导页面,新的登录界面 ([06aae04](https://github.com/xuender/go-rich/commit/06aae04)), closes [#33](https://github.com/xuender/go-rich/issues/33)
* **设置:** 增加关于页面,显示名称,版本,简介 ([e237c38](https://github.com/xuender/go-rich/commit/e237c38)), closes [#34](https://github.com/xuender/go-rich/issues/34)



<a name="0.1.8"></a>
## [0.1.8](https://github.com/xuender/go-rich/compare/5958129...v0.1.8) (2018-10-11)


### Bug Fixes

* **订单:** 修复显示订单中客户和商品时会发起多次调用的错误 ([b0f09df](https://github.com/xuender/go-rich/commit/b0f09df)), closes [#31](https://github.com/xuender/go-rich/issues/31)


### Features

* **客户,商品:** 首字母多音字时生成多个单字标签 ([a9cf5b5](https://github.com/xuender/go-rich/commit/a9cf5b5)), closes [#32](https://github.com/xuender/go-rich/issues/32)
* **客户和商品:** 删除客户或商品后,历史订单还可以看到当时的客户和购买的商品 ([5958129](https://github.com/xuender/go-rich/commit/5958129)), closes [#29](https://github.com/xuender/go-rich/issues/29) [#30](https://github.com/xuender/go-rich/issues/30)



