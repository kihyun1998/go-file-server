# go-file-server

## API Documents

[API Documentation](api-docs.md)



## 개발 시 실행 방법

```bash
go run .
```

## Linux 용 build 하는 법

```bash
# 명령 프롬프트(cmd)에서:
GOOS=linux GOARCH=amd64 go build -o go-file-server

# PowerShell에서:
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o go-file-server
```

## 다시 윈도우로 되돌리기

```bash
$env:GOOS="windows";
```

## Linux에서 백그라운드 실행하는 법

```bash
nohup ./go-file-server &
```

```bash
nohup ./go-file-server > /dev/null 2>&1 &
```

## 정상 종료 하는 방법

```bash
ps aux | grep go-file-server

# 정상 종료
kill -15 [PID번호]

# 강제 종료 (필요시)
kill -9 [PID번호]
```