
# MongoDB Timeseries Tester 

## 개요 
Timeseries 를 테스트하기 위한 테스터 

## 사용법 
1. MONOGDB_URI 환경변수에 접속할 mongodb 접속 경로 명시

```shell
export MONGODB_URI="mongodb://my:password@localhost:27017/?replicaset=prod&maxPoolSize=8192&retryWrites=false"
```

2. 해당 바이너리 빌드
```shell
$ make
```

3. 프로그램 수행 
```shell
./timeseries_linux 
```

## 옵션 

```shell
Usage of ./timeseries_osx:
  -bulk int
    	number of batch count (default 10)
  -collection string
    	Collection to insert  (default "timeseries")
  -database string
    	database to insert  (default "timeseries")
  -iteration int
    	number of test set  (default 1000)
  -metafield-count int
    	metafield count (default 3)
  -metric-count int
    	metric field count (default 5)
  -sensor int
    	# of sensors (default 10)
``` 

* bulk : 한번에 모아서 쓸 갯수. iteration / bulk 는 무리수가 되면 안됨. 
* collection : collection to write 
* database : database to write 
* iteration : 수행할 iteration. 이 횟수만큼 초단위로 event 를 발생 
* metafield-count : metafield 의 갯수 
* metric-count : metric 의 갯수 
* sensor : 센서수. 각 센서는 1초당 하나의 이벤트를 생성 

소스 