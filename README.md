# GO-SP2
小型訂單系統（Go + Gin + GORM + Redis + Docker）

- 實作使用者註冊/登入，使用 JWT 與 bcrypt 確保安全性
- 設計商品與訂單資料模型，支援 RESTful CRUD 與關聯查詢
- 整合 Redis Pub/Sub 機制，模擬訂單成立後的非同步通知
- 使用 Goroutine 處理後台訂單狀態異動，並記錄 log 日誌
- 專案完整容器化，透過 Docker Compose 管理 API、DB、Redis 等服務
- 使用 GitHub Actions 實作自動化部署，專案可一鍵上線至 Render
