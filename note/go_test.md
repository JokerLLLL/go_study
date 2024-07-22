## gcc 报错

`go test  ./... -cover -race -gcflags=-l`

报错：
```text
exec: "gcc": executable file not found in %PATH%解决办法
```

windowns安装gcc：
https://www.cnblogs.com/xpcloud/p/13269999.html

## build pii-services

`docker build -t pii .`

```DOCKFILE
# Version: 0.6
FROM docker-reg.uco.com/golang:1.15

WORKDIR /app/dest/pii-service
COPY . /app/dest/pii-service
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
RUN go mod download
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && cd /app/dest/pii-service && /usr/local/go/bin/go build .
RUN mkdir /var/log/pii-service
CMD [ "YouWillBeFiredIfUseThisPWInProd", "10000", "12", "/tmp" ]
ENTRYPOINT ["/app/dest/pii-service/piiserver"]

```

## 预跑 pii-services

`go run main.go crypticserver.go dummyDriver.go localDriver.go vaultDriver.go platformDriver.go 10000 10000 15`

## 起环境 pii-service

```cmd 
docker run 
-v /home/zouyun/pii-service2:/app/dest/pii-service (volumn有问题？)
-p 8082:8081 
-d --name=pii_zy2 
-e "IGNORE_SIGN=true" 
-e "ENABLE_VAULT=true" 
-e "VAULT_HOST=http://app11.uco.com:32774" 
-e "VAULT_ROLE_ID=d1bb22c4-c3c0-2a2f-33c3-472885d0e04a" 
-e "VAULT_SECRET_ID=c6e30dab-ceb3-340d-df13-abc9814cbb30" 
-e "SENTRY_DSN=https://c44521c365334a4fa1ad11f8ecda0bb3@sentry.f.uco.com/16" 
-e "MOCK_CMS=false" 
pii sh -c "go run main.go crypticserver.go vaultDriver.go dummyDriver.go localDriver.go platformDriver.go 10000 10000" 
```

```
docker run 
-p 8082:8081 
-d --name=pii_zy2 
-e "IGNORE_SIGN=true" 
-e "ENABLE_VAULT=true" 
-e "VAULT_HOST=http://app11.uco.com:32774" 
-e "VAULT_ROLE_ID=d1bb22c4-c3c0-2a2f-33c3-472885d0e04a" 
-e "VAULT_SECRET_ID=c6e30dab-ceb3-340d-df13-abc9814cbb30" 
-e "SENTRY_DSN=https://c44521c365334a4fa1ad11f8ecda0bb3@sentry.f.uco.com/16" 
-e "MOCK_CMS=false" 
pii
```

## unitest

```cmd
-- 单元测试
go test  ./... -cover -race -gcflags=-l  $*

-- 本地调试

$  go run crypticserver.go localDriver.go dummyDriver.go vaultDriver.go platformDriver.go main.go --cacheSize=100001 --logMode=15 --logPath="D:\demo\pii-service\logs"

recognize PlatformSalt config

-- 服务器起服务

docker build -t pii-stress .

docker run -d \
	--name=pii-stress \
        -p 9090:8081 \
	-v /home/chenchaoliang/pii-log:/var/log/pii-service \
	-v /home/chenchaoliang/pii-security:/app/dest/pii-service/security \
        -e "IGNORE_SIGN=true" \
        --rm \
        pii-stress

-- 压测工具采用：wrk
https://github.com/wg/wrk

./wrk -c1000 -t12 -d10s http://10.2.3.145:9090/info
./wrk -c1000 -t12 -d10s --script=post.lua  'https://pii-stress.test.uco.com/decryptPlatform?timeStamp=1&appKey=1&v=1.0&sign=1'
./wrk -c1000 -t12 -d10s --script=post.lua  'https://pii-stress.test.uco.com/decryptPlatform?timeStamp=1&appKey=1&v=1.0&sign=1'

```


### 本地调试

sign设置成TRUE

go run crypticserver.go localDriver.go dummyDriver.go vaultDriver.go platformDriver.go main.go --cacheSize=100001 --logMode=15 --logPath="D:\tmp"