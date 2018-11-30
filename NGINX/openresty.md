## NGINX+LUA+REDIS 高性能服务器搭建
> 对于传统的服务端程序 （PHP, FASTCGI等）,大多都是通过产生一个请求，有一个进程与之相对应，请求处理结束后，进程销毁释放，所以一些语言都通过常驻进程，线程，线程池等降低资源开销。即使是资源占用最小线程，当并发量超过1K(单台)的时候，操作系统的处理就会开始明显下降，因为有太多的CPU时间都消耗在系统的上下文切换。因此，对于一些**性能要求比较高,并发量较大**的需求，就需要一套高性能的服务区去支撑。<br>

### NGINX+LUA
nginx的lua-nginx-module模块将lua嵌入到nginx，让nginx高效的执行lua脚本，**高并发,非阻塞**的处理各种请求。<br>
Lua内建协程，这样就可以很好的将一步回调转换成顺序调用的形式。ngx_lua在lua中进行IO操作都会委托给Nginx事件模型，实现非阻塞调用。<br>
每个NginxWorker进程持有一个Lua解释器或者LuaJIT实例，被这个Worker处理的所有请求共享这个实例。每个请求的Context会被Lua轻量级的协程分割，从而保证各个请求是独立的。Ngx_lua 采用“one-conroutine-pre-request”的处理模型，对于每个用户请求，ngx_lua 会唤醒用于执行用户代码处理请求，当请求处理完成这个写成会被销毁。每个协程都有一个独立的全局环境（变量空间），继承于全局共享的、只读的“comman data”。所以，被用户代码注入全局空间的任何变量都不会影响其他请求的处理，并且这些变量在请求处理完成后会被释放，这样就保证所有的用户代码都运行在一个“sandbox”（沙箱），这个沙箱与请求具有相同的生命周期。 得益于Lua协程的支持，ngx_lua在处理10000个并发请求时只需要很少的内存。根据测试，ngx_lua处理每个请求只需要2KB的内存，如果使用LuaJIT则会更少。所以ngx_lua非常适合用于实现可扩展的、高并发的服务。

### OPENRESTY 安装
```
sudo yum-config-manager --add-repo https://openresty.org/yum/cn/centos/OpenResty.repo
sudo yum install openresty
```
没有 `yum-config-manager` 命令 执行 `sudo yum -y install yum-utils` 安装

执行 `openresty -h` 成功

### REDIS 安装
```
yum install epel-release
yum install redis

// 修改redis 配置
redis-service path/redis.conf &

// 测试是否安装成功
redis-cli
```
### 初试-HelloWorld
```
// 根目录下 创建项目目录
mkdir ~/openresty ~/openresty/logs/ ~/openresty/conf/

// conf 目录 设置 nginx配置
worker_processes  2;        #nginx worker 数量 核数2
error_log logs/error.log;   #指定错误日志文件路径
pid        logs/openresty.pid;
#access_log logs/access.log; #info  日志
events {
    worker_connections 10240;
}

http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    log_format  main1 '[$time_local] $host $remote_addr "$request_uri" '
                '$status "$bytes_sent" "$request_time" "$upstream_response_time" '
                '"$http_user_agent" "$http_referer" "$http_x_forwarded_for" "$http_cookie" "$uid_set" "$uid_got"';
    
    lua_package_path "lua/?.lua;;";
    lua_package_cpath "lua/?.so;;";
    
    #access_log  logs/access.log  main;
    server {
        access_log  logs/access.log  main1;
        #监听端口，若你的6699端口已经被占用，则需要修改
        listen 6699;

        location ~ ^/api/([-_a-zA-Z0-9/]+) {
            # 准入阶段 参数校验
            #access_by_lua_file lua/access_check.lua;

            # 逻辑
            content_by_lua_file lua/redis/repos.lua;
        }
        location / {
            default_type text/html;

            content_by_lua '
                ngx.say("HelloWorld") // helloworld
            ';
        }
    }
}

```

### NGINX优化
worker_processes 设置 最好和cpu 逻辑核数一致 (避免上下文切换) 
```
# 查看CPU 逻辑核数
cat /proc/cpuinfo |grep "processor"|wc -l
```

### siege 压力测试
[安装](!https://centos.pkgs.org/6/repoforge-x86_64/siege-2.72-1.el6.rf.x86_64.rpm.html)
```
wget http://ftp.tu-chemnitz.de/pub/linux/dag/redhat/el6/en/x86_64/rpmforge/RPMS/siege-2.72-1.el6.rf.x86_64.rpm
sudo rpm -Uvh siege-2.72-1.el6.rf.x86_64.rpm
sudo yum install siege

# 压力测试
# -r 测试轮数
# -c 并发数  siege超过1000会报错
siege http://test.com/api/123 -r 20 -c 2000
```

### ab 压力测试
```
# 只安装 ab工具包
yum -y install httpd-tools

# 测试是否安装成功
ab -V

# 1W 并发 10W 请求
ab -n100000 -c10000 http://test.com/api/123
ab -r # 默认报错
```

并发超过1000 会出错

- helloworld接口

| 总次数 | 并发数 | 失败数 | 成功率 | 最大时长(ms) | 大概率时长(90%) |
| ---- | ---- | ---- | ---- | --- | --- |
|100000|100|0|100%| 17 | 5 | 
|100000|1000|0|100%| 3043 | 62 |
|100000|5000|0|100%| 3342 | 419 |
|100000|10000|0|100%| 4180 | 1490 |
|100000|20000|0|100%| 5118 | 1650 |

- 优化前结果

| 总次数 | 并发数 | 失败数 | 成功率 | 最大时长(ms) | 大概率时长(90%) |
| ---- | ---- | ---- | ---- | --- | --- |
|100000|100|0|100%| 29 | 12 | 
|100000|1000|0|100%| 8769 | 26 |
|100000|5000|0|100%| 15779 | 1577 |
|100000|10000|99|100%| 9231 | 3057 |
|100000|10000|105|100%| 8803 | 2416 |
|100000|10000|0|100%| 8388| 2343 |
|1000000|10000|1|100%| 17535| 2723 |
|100000|20000|0|100%| 21464 | 6463 |
|100000|20000|0|100%| 14434 | 6916 |
|1000000|20000|1|100%| 10494 | 8520 |

### 优化
- worker_connections 被打满<br>
[NGINX最大并发max_clients计算](!http://blog.51cto.com/liuqunying/1420556)<br>
NGINX作为http服务器时:<br>
max_clients = worker_process * worker_conections<br>
NGINX作为反向代理时:<br>
max_clients = worker_process * worker_conections / 4<br>
由以上计算得，原来最大clients = 2 * 10240 = 20480
在测试并发2W以上，会报连接数被打满。修改worker_connnections为20480,解决此问题。

- lua tcp socket connect timed out<br>
    <br>redis 连接超时<br>
    - 修改方案1-将超时时间设置为20s(原来2s)<br>
    结果: 超时操作明显减少
    - 修改方案2-[使用连接池](!https://www.the5fire.com/golang+redis-vs-nginx+lua+redis.html)<br>
    两W并发 100W请求下 连接池对比<br>
    
    |连接池数|超时数|
    | ---- | --- |
    |100|4434|
    |100|3777|
    |1000|4295|
    |1000|2579|
    |10000|2252|
    |10000|5106|
    发现，加大连接池数并没有减少超时，因此采用加大超时时间+设置100连接池(减少资源消耗)。<br>2W并发 100W请求 超时数为0.

- 优化后结果

| 总次数 | 并发数 | 失败数 | 成功率 | 最大时长(ms) | 大概率时长(90%) |
| ---- | ---- | ---- | ---- | --- | --- |
|165270|16527|0|100%|9051|4899|
|104490|10449|0|100%|16109|2515|
|90590|9059|0|100%|4759|2164|
|132460|13246|0|100%|8670|3210|
