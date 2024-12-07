# File Server API Documentation

## Base URL
```
http://localhost:8080
```

## Upload Endpoints

### 1. Upload File
Upload a file using multipart/form-data.

**Endpoint:** `POST /upload`

**Content-Type:** `multipart/form-data`

**Request Body:**
- `file`: File to upload (form-data)

**Response:**
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

**Error Response:**
```json
{
    "isOk": false,
    "error": "Failed to read file"
}
```

### 2. Upload Base64 File
Upload a file using base64 encoded data.

**Endpoint:** `POST /upload/base64`

**Content-Type:** `application/json`

**Request Body:**
```json
{
    "fileData": "base64EncodedString",
    "filename": "example.jpg"
}
```

**Response:**
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

**Error Response:**
```json
{
    "isOk": false,
    "error": "Invalid base64 data"
}
```

## Download Endpoints

### 1. Download File
Download a file directly.

**Endpoint:** `GET /download/:type/:filename`

**Parameters:**
- `type`: File type (images, videos, others)
- `filename`: Name of the file to download

**Response:**
- File will be downloaded directly
- Content-Type will be set based on file type
- Content-Disposition header will be set for attachment

**Error Response:**
```json
{
    "isOk": false,
    "error": "File not found"
}
```

### 2. Download Base64 File
Download a file with base64 encoding.

**Endpoint:** `GET /download/base64/:type/:filename`

**Parameters:**
- `type`: File type (images, videos, others)
- `filename`: Name of the file to download

**Response:**
- File data will be sent with appropriate Content-Type header
- Content-Disposition header will be set for attachment

**Error Response:**
```json
{
    "isOk": false,
    "error": "File not found"
}
```

## File Types
The server automatically categorizes files into the following types:
- `images`: .jpg, .jpeg, .png, .gif
- `videos`: .mp4, .avi, .mov, .wmv
- `others`: Any other file type

## Additional Information

### CORS Configuration
- Allows all origins
- Allowed Methods: GET, POST, OPTIONS
- Allowed Headers: Origin, Content-Type, Accept

### File Size Limits
- No file size limit is set (`MaxMultipartMemory = 0`)

### Storage Location
- Files are stored in the `./uploads` directory
- Subdirectories are automatically created for different file types:
  - `./uploads/images`
  - `./uploads/videos`
  - `./uploads/others`

### File Permissions
- Linux: 755 for directories, 644 for files
- Windows: 666 for both directories and files

### Error Handling
All endpoints return standardized error responses in the following format:
```json
{
    "isOk": false,
    "error": "Error message description"
}
```
