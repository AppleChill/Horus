# Horus
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags -w main.go

## 條件
1. MA1 > MA2 > MA3
2. MV1 > MV2 > MV3
3. D天內，每日的成交量不低於X張

## config.yml 參數設定
1. MA1: k線1
2. MA2: k線2
3. MA3: k線3
4. MV1: 成交量1
5. MV2: 成交量2
6. MV3: 成交量3
7. MAX: 查詢最大天數
8. DAYS: 條件三的D
9. LOT: 條件三的X