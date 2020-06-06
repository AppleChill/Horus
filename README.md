# Horus
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags -w horus.go

## 條件
1. MA1 > MA2 > MA3
2. MV1 > MV2 > MV3
3. D天內，每日的成交量不低於X張
4. 最後Z天總和的成交量大於Z天前(C天)總和的成交量G%

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
10. Z: 條件四的Z
11. C: 條件四的C
12. G: 條件四的G
13. Role1: 條件1 true or false
14. Role2: 條件2 true or false
15. Role3: 條件3 true or false
16. Role4: 條件4 true or false
17. FILTER: 忽略最近的天數，不含假日