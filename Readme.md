## 目录分层
* api: 注册接口，接受参数，检验参数，返回数据
* server: 处理主要业务逻辑
* model: 数据操作层
* config: 放置项目的配置文件（路径固定不可修改）

## 配置项
| key | 例子                | 说明                                              |
|-----|-------------------|-------------------------------------------------|
| n-ip | 127.0.0.1         | nacos服务器地址                                      |
| n-port | 8848              | nacos服务端口号                                      |
|n-path | /nacos            | nacos访问路径                                       |
|data-id | user-rpc          | 数据ID（建议：使用服务名加服务协议作为id）                         |
|namespace | 2206A             | 部门命名空间                                          |
|group | rpc               | 数据分组（按协议分租，rpc服务使用“rpc”作为分组名，api服务使用“api”作为分组名） |
|timeout-ms | 5000              | 客户端访问nacos超时时间                                  |
|log-dir | ./tmp/nacos/log   | nacos日志存放位置                                     |
|cache-dir | ./tmp/nacos/cache | nacos缓存存放位                                      |
|log-level | debug             | 日志等级                                            |
