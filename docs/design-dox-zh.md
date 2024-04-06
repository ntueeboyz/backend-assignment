# 設計文件
在 Dacard 作業使用 Gin 框架作為後端服務，並使用 Redis 作為資料庫。

## Go-Gin 框架
在作業中我選擇了效能比較好的 HTTP 框架 Gin-Gonic。

### 理由
- Go-Gin 的效能比較好。它是基於 httprouter 構建的，可以比其他框架快，有機會達到每秒處理 10,000 個請求的目標。
- Gin 除了高效能以外，同時使用上比較方便，讓開發和維護程式碼變得容易。
- Gin 擁有龐大的使用群眾，所以能比較容易找到資源來解決開方上的問題。

## Redis 資料庫
選擇 Redis 作為主要的資料庫，是因為可以快速存取廣告資料。

### 理由
- Redis 通常將全部的資料儲存在記憶體中，所以能提供了低延遲和高吞吐量，能達到快速存取數據需求。這也使達到每秒處理 10,000 個請求的目標變得更加容易。
- 雖然 Redis 主要要將資料存在記憶體中，但 Redis 也提供了將資料存在硬碟選擇，可以確保資料在服務的重啟之間不會丟失。

## 作業架構
```
├─ app
│   ├─ main.go
│   ├─ handler.go
│   ├─ model.go
│   ├─ service.go
│   └─ service_test.go
│
├─ conf
│   └─ redis.conf
│
├─ Dockerfile
├─ docker-compose.yml
├─ go.mod
└─ go.sum
```

1. **控制器(Controller)：** `handler.go` 基本上是處理傳入請求並向客戶端返回響應的控制器。
2. **服務(Service)：**服務提供處理資料的基本功能，例如篩選。這個作業中的服務是指 service.go。
3. **模型(Model)：** 模型指的是用於從資料庫存儲和檢索的資料結構。

**其他：**  
1. 由於項目中只有兩個 API，資料結構沒有像其專案有很好的分類。
2. 單元測試可能未涵蓋所有情境。