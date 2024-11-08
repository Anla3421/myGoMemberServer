# Simple Member Server API

一個基於 Golang 實現的簡單會員服務系統 API，使用 Gin 框架開發。
API 文件：http://localhost:9000/swagger/index.html

## 功能特性

- 使用者註冊與認證
- JWT 身份驗證
- 會員資訊管理
- 登入紀錄
- Swagger API 文件
- gRPC 服務整合
- Redis 快取支援

## 技術架構

- Golang 1.23.1
- Gin Web Framework
- gRPC
- MySQL
- Redis
- Swagger
- JWT

## API 詳細說明

### 使用者管理

#### 1. 使用者註冊 (Registry)
- 端點：`POST /registry`
- 功能：
  - 確認帳號是否已存在
  - 密碼進行 MD5 加密
  - 將使用者資料存入資料庫
- 請求參數：
  ```json
  {
    "account": "使用者帳號",
    "password": "使用者密碼"
  }
  ```

#### 2. 使用者登入 (Login)
- 端點：`POST /login`
- 功能：
  - 驗證帳號密碼
  - 產生 JWT Token
  - 記錄登入時間（Unix timestamp）
- 請求參數：
  ```json
  {
    "account": "使用者帳號",
    "password": "使用者密碼"
  }
  ```
- 回應：包含 JWT Token

#### 3. 使用者登出 (Logout)
- 端點：`POST /logout`
- 功能：
  - 驗證 JWT Token
  - 刪除伺服器端的 JWT Token
- 請求標頭：
  - 需要包含有效的 JWT Token

#### 4. 修改密碼 (ChPWD)
- 端點：`POST /chpwd`
- 功能：
  - 驗證 JWT Token
  - 更新使用者密碼
  - 重新產生 JWT Token
- 請求參數：
  ```json
  {
    "jwt": "當前JWT Token",
    "newpassword": "新密碼"
  }
  ```

### 紀錄管理

#### 1. 刪除登入紀錄 (LogDel)
- 端點：`DELETE /logdel`
- 功能：刪除指定使用者的所有登入紀錄
- 請求參數：
  ```json
  {
    "account": "使用者帳號"
  }
  ```

#### 2. 查詢會員資訊 (Query)
- 端點：`GET /query/:account`
- 功能：
  - 查詢會員基本資料
  - 使用 Redis 快取提升查詢效能
- 回應：包含會員 ID、帳號和 JWT Token

#### 3. 查詢登入紀錄 (MemberLog)
- 端點：`GET /memberlog`
- 功能：查詢指定會員的所有登入時間記錄
- 查詢參數：
  - account: 使用者帳號

### API 文件
- 端點：`/swagger/index.html`
- 功能：提供互動式 API 文件介面
- 包含所有 API 的詳細說明和測試功能

## 專案結構