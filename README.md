# Simple Member Server API

一個基於 Golang 實現的簡單會員服務系統 API，使用 Gin 框架開發。此系統提供完整的會員管理功能，包括註冊、登入、登出、密碼修改以及會員資訊查詢等功能。

## 功能特點

* 使用者註冊與認證
* JWT 身份驗證
* 會員資訊管理
* 登入紀錄管理
* Swagger API 文件
* gRPC 服務整合
* Redis 快取支援

## 技術架構

* Golang 1.23.1
* Gin Web Framework
* gRPC
* MySQL 8.0+
* Redis 6.0+
* Swagger
* JWT
* Docker & Docker Compose

## API 端點說明

| 功能 | HTTP 方法 | URI | 描述 |
| --- | --- | --- | --- |
| 使用者註冊 | POST | /registry | 註冊新會員帳號 |
| 使用者登入 | POST | /login | 會員登入並獲取JWT |
| 使用者登出 | POST | /logout | 會員登出並刪除JWT |
| 修改密碼 | POST | /chpwd | 修改會員密碼 |
| 刪除登入紀錄 | DELETE | /logdel | 刪除指定會員的登入紀錄 |
| 查詢會員資訊 | GET | /query/:account | 查詢指定會員的基本資料 |
| 查詢登入紀錄 | GET | /memberlog | 查詢指定會員的登入時間記錄 |

## API 詳細說明

### 1. 使用者註冊 (Registry)
- **端點**：`POST /registry`
- **功能**：
  - 確認帳號是否已存在
  - 密碼進行 MD5 加密
  - 將使用者資料存入資料庫
- **請求參數**：
  ```json
  {
    "account": "使用者帳號",
    "password": "使用者密碼"
  }
  ```
- **回應**：
  ```json
  {
    "code": 200,
    "message": "註冊成功",
    "data": null
  }
  ```

### 2. 使用者登入 (Login)
- **端點**：`POST /login`
- **功能**：
  - 驗證帳號密碼
  - 產生 JWT Token
  - 記錄登入時間（Unix timestamp）
- **請求參數**：
  ```json
  {
    "account": "使用者帳號",
    "password": "使用者密碼"
  }
  ```
- **回應**：
  ```json
  {
    "code": 200,
    "message": "登入成功",
    "data": {
      "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
  }
  ```

### 3. 使用者登出 (Logout)
- **端點**：`POST /logout`
- **功能**：
  - 驗證 JWT Token
  - 刪除伺服器端的 JWT Token
- **請求標頭**：
  - 需要包含有效的 JWT Token
- **回應**：
  ```json
  {
    "code": 200,
    "message": "登出成功",
    "data": null
  }
  ```

### 4. 修改密碼 (ChPWD)
- **端點**：`POST /chpwd`
- **功能**：
  - 驗證 JWT Token
  - 更新使用者密碼
  - 重新產生 JWT Token
- **請求參數**：
  ```json
  {
    "jwt": "當前JWT Token",
    "newpassword": "新密碼"
  }
  ```
- **回應**：
  ```json
  {
    "code": 200,
    "message": "密碼修改成功",
    "data": {
      "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
  }
  ```

### 5. 刪除登入紀錄 (LogDel)
- **端點**：`DELETE /logdel`
- **功能**：刪除指定使用者的所有登入紀錄
- **請求參數**：
  ```json
  {
    "account": "使用者帳號"
  }
  ```
- **回應**：
  ```json
  {
    "code": 200,
    "message": "登入紀錄刪除成功",
    "data": null
  }
  ```

### 6. 查詢會員資訊 (Query)
- **端點**：`GET /query/:account`
- **功能**：
  - 查詢會員基本資料
  - 使用 Redis 快取提升查詢效能
- **回應**：
  ```json
  {
    "code": 200,
    "message": "查詢成功",
    "data": {
      "id": 1,
      "account": "使用者帳號",
      "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
  }
  ```

### 7. 查詢登入紀錄 (MemberLog)
- **端點**：`GET /memberlog`
- **功能**：查詢指定會員的所有登入時間記錄
- **查詢參數**：
  - account: 使用者帳號
- **回應**：
  ```json
  {
    "code": 200,
    "message": "查詢成功",
    "data": [
      {
        "account": "使用者帳號",
        "login_time": 1636600000
      },
      {
        "account": "使用者帳號",
        "login_time": 1636700000
      }
    ]
  }
  ```

## 會員資料結構

```json
{
  "id": 1,
  "account": "使用者帳號",
  "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

## 登入紀錄結構

```json
{
  "account": "使用者帳號",
  "login_time": 1636600000
}
```

## 系統流程
另可參考 `chart flow memberServer1.png` 及 `chart flow memberServer2.png`

### 會員管理流程 (POST 方法)

1. **註冊流程**：確認帳號不存在 → 密碼雜湊處理 → 資料存入資料庫
2. **登入流程**：驗證帳號密碼 → 生成 JWT → 記錄登入時間
3. **登出流程**：驗證 JWT → 刪除伺服器端 JWT
4. **密碼修改流程**：驗證 JWT → 更新密碼 → 生成新 JWT

### 資料查詢流程 (GET/DELETE 方法)

1. **會員資訊查詢**：查詢會員資料 → 使用 Redis 快取提升效能
2. **登入記錄查詢**：查詢指定會員的所有登入時間記錄
3. **登入記錄刪除**：刪除指定會員的所有登入記錄

## 運行項目

1. 建立並啟動 Docker 容器
   ```bash
   bash ./infrastructure/docker/go.sh
   ```

2. 啟動服務
   ```bash
   go run main.go
   ```

3. 訪問 Swagger API 文檔
   ```
   http://localhost:9000/swagger/index.html
   ```

