# File Server API Documentation

## Base URL
```
http://localhost:8080
```

## Endpoints

### Upload File
`POST /upload`

Upload a file to the server.

#### Request
- **Content-Type:** `multipart/form-data`
- **Form Field:** `file`

#### cURL Example
```bash
curl -X POST -F "file=@C:\path\to\your\image.jpg" http://localhost:8080/upload
```

#### Success Response (200 OK)
```json
{
    "isOk": true,
    "message": "File uploaded successfully",
    "data": {
        "path": "images/example.jpg",
        "filename": "example.jpg",
        "fileType": "images"
    }
}
```

#### Error Response (400/500)
```json
{
    "isOk": false,
    "error": "Failed to read file"
}
```

### Download File
`GET /download/{type}/{filename}`

Download a specific file from the server.

#### Parameters
| Name     | Type   | Description                          |
|----------|--------|--------------------------------------|
| type     | string | File type (images, videos, others)   |
| filename | string | Name of the file to download         |

#### cURL Example
```bash
curl -O http://localhost:8080/download/images/example.jpg
```

#### Browser Access
```
http://localhost:8080/download/images/example.jpg
```

#### Error Response (400/404)
```json
{
    "isOk": false,
    "error": "File not found"
}
```

## Supported File Types

### Images
- `.jpg`
- `.jpeg`
- `.png`
- `.gif`

### Videos
- `.mp4`
- `.avi`
- `.mov`
- `.wmv`

### Others
- All other file extensions are stored in the 'others' category

## Testing Tools
- **Postman**
- **cURL**
- Web Browser (for downloads)
- Any HTTP client that supports multipart/form-data
