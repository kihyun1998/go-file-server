# 파일 서버 API 명세서

## 1. Base64 파일 업로드

Base64로 인코딩된 파일을 서버에 업로드합니다.

### 요청 정보
- **URL**: `/upload/base64`
- **Method**: POST
- **Content-Type**: application/json

### 요청 본문
```json
{
    "fileData": "data:image/jpeg;base64,/9j/4AAQSkZJRg...",  // Base64로 인코딩된 파일 데이터
    "filename": "example.jpg"                                  // 원본 파일명
}
```

### 응답
#### 성공 응답 (200 OK)
```json
{
    "isOk": true,
    "message": "File uploaded successfully",
    "data": {
        "path": "images/example.jpg",      // 저장된 파일 경로
        "filename": "example.jpg",         // 저장된 파일명
        "fileType": "images"              // 파일 유형 (images, videos)
    }
}
```

#### 실패 응답
- **잘못된 요청 (400 Bad Request)**
```json
{
    "isOk": false,
    "error": "Invalid base64 data"
}
```

- **서버 오류 (500 Internal Server Error)**
```json
{
    "isOk": false,
    "error": "Failed to save file"
}
```

### 주의사항
- 지원되는 이미지 형식: jpg, jpeg, png, gif
- 지원되는 비디오 형식: mp4, avi, mov, wmv
- 파일명에는 안전한 문자만 허용됨
- Base64 데이터는 프리픽스(예: "data:image/jpeg;base64,")를 포함하거나 제외할 수 있음

## 2. 파일 다운로드

저장된 파일을 다운로드합니다.

### 요청 정보
- **URL**: `/download/:type/:filename`
- **Method**: GET
- **URL 파라미터**:
  - type: 파일 유형 (images, videos)
  - filename: 파일명

### 예시 URL
```
/download/images/example.jpg
/download/videos/example.mp4
```

### 응답
#### 성공 응답
- **Content-Type**: 파일 유형에 따라 자동 설정
  - 이미지: image/jpeg, image/png, image/gif
  - 비디오: video/mp4, video/x-msvideo
- **Content-Disposition**: attachment; filename=파일명

#### 실패 응답
- **잘못된 요청 (400 Bad Request)**
```json
{
    "isOk": false,
    "error": "Invalid file type"
}
```

- **파일 없음 (404 Not Found)**
```json
{
    "isOk": false,
    "error": "File not found"
}
```

## 3. 다운로드 페이지

파일 다운로드를 위한 웹 페이지를 제공합니다.

### 요청 정보
- **URL**: `/page/download/:filename`
- **Method**: GET
- **URL 파라미터**:
  - filename: 파일명

### 예시 URL
```
/page/download/example.jpg
```

### 응답
- **Content-Type**: text/html
- 파일 다운로드를 위한 HTML 페이지가 반환됩니다.

### 주의사항
- 모든 API 엔드포인트는 CORS가 활성화되어 있음
- 파일 크기 제한이 해제되어 있으나, 서버 리소스를 고려하여 적절한 크기의 파일만 업로드 권장
- 모든 파일 경로는 서버의 보안을 위해 검증됨