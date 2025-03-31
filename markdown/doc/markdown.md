### 1. 批量删除文件

1. route definition

- Url: /api/v1/batch
- Method: DELETE
- Request: `BatchDeleteRequest`
- Response: `BatchDeleteResponse`

2. request definition



```golang
type BatchDeleteRequest struct {
	Filenames []string `json:"filenames"`
}
```


3. response definition



```golang
type BatchDeleteResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	SuccessCount int `json:"successCount"`
	FailedCount int `json:"failedCount"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 2. 获取文件内容

1. route definition

- Url: /api/v1/file/:filename
- Method: GET
- Request: `DownloadResponse`
- Response: `DownloadResponse`

2. request definition



```golang
type DownloadResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Content string `json:"content,optional"` // 小文件直接返回内容
	DownloadURL string `json:"downloadUrl"` // 预签名的S3下载URL
	File MarkdownFile `json:"file,optional"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type MarkdownFile struct {
	Id int64 `json:"id"`
	Filename string `json:"filename"`
	Size int64 `json:"size"`
	UploadTime int64 `json:"uploadTime"`
	UpdateTime int64 `json:"updateTime"`
	OwnerId int64 `json:"ownerId"`
	DownloadURL string `json:"downloadUrl"` // 预签名的 S3 URL
	Hash string `json:"hash"`
	IsPublic bool `json:"isPublic"`
	IsDeleted bool `json:"isDeleted"`
	Tags []string `json:"tags,optional"`
	Path string `json:"path"` // S3对象的完整路径
}
```


3. response definition



```golang
type DownloadResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Content string `json:"content,optional"` // 小文件直接返回内容
	DownloadURL string `json:"downloadUrl"` // 预签名的S3下载URL
	File MarkdownFile `json:"file,optional"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type MarkdownFile struct {
	Id int64 `json:"id"`
	Filename string `json:"filename"`
	Size int64 `json:"size"`
	UploadTime int64 `json:"uploadTime"`
	UpdateTime int64 `json:"updateTime"`
	OwnerId int64 `json:"ownerId"`
	DownloadURL string `json:"downloadUrl"` // 预签名的 S3 URL
	Hash string `json:"hash"`
	IsPublic bool `json:"isPublic"`
	IsDeleted bool `json:"isDeleted"`
	Tags []string `json:"tags,optional"`
	Path string `json:"path"` // S3对象的完整路径
}
```

### 3. 删除文件

1. route definition

- Url: /api/v1/file/:filename
- Method: DELETE
- Request: `-`
- Response: `DeleteResponse`

2. request definition



3. response definition



```golang
type DeleteResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 4. 更新文件信息

1. route definition

- Url: /api/v1/file/:filename
- Method: PUT
- Request: `UpdateFileRequest`
- Response: `UpdateFileResponse`

2. request definition



```golang
type UpdateFileRequest struct {
	IsPublic bool `json:"isPublicoptional"`
	Content string `json:"contentoptional"`
	Tags []string `json:"tagsoptional"`
	Filename string `json:"filenameoptional"` // 重命名
	Path string `json:"pathoptional"` // 移动文件到其他路径
}
```


3. response definition



```golang
type UpdateFileResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	File MarkdownFile `json:"file"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type MarkdownFile struct {
	Id int64 `json:"id"`
	Filename string `json:"filename"`
	Size int64 `json:"size"`
	UploadTime int64 `json:"uploadTime"`
	UpdateTime int64 `json:"updateTime"`
	OwnerId int64 `json:"ownerId"`
	DownloadURL string `json:"downloadUrl"` // 预签名的 S3 URL
	Hash string `json:"hash"`
	IsPublic bool `json:"isPublic"`
	IsDeleted bool `json:"isDeleted"`
	Tags []string `json:"tags,optional"`
	Path string `json:"path"` // S3对象的完整路径
}
```

### 5. 获取当前用户的文件列表

1. route definition

- Url: /api/v1/files
- Method: GET
- Request: `ListRequest`
- Response: `ListResponse`

2. request definition



```golang
type ListRequest struct {
	Page int `form:"pageoptionaldefault=1"`
	PageSize int `form:"pageSizeoptionaldefault=20"`
	Path string `form:"pathoptionaldefault=/"` // S3路径前缀，替代文件夹ID
	SortBy string `form:"sortByoptionaldefault=updateTime"`
	Order string `form:"orderoptionaldefault=desc"`
	Recursive bool `form:"recursiveoptionaldefault=false"` // 是否递归列出子目录
}
```


3. response definition



```golang
type ListResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Total int64 `json:"total"`
	Files []MarkdownFile `json:"files"`
	Paths []string `json:"paths,optional"` // 子路径列表
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 6. 获取当前用户公开文件列表

1. route definition

- Url: /api/v1/files/public
- Method: GET
- Request: `ListRequest`
- Response: `ListResponse`

2. request definition



```golang
type ListRequest struct {
	Page int `form:"pageoptionaldefault=1"`
	PageSize int `form:"pageSizeoptionaldefault=20"`
	Path string `form:"pathoptionaldefault=/"` // S3路径前缀，替代文件夹ID
	SortBy string `form:"sortByoptionaldefault=updateTime"`
	Order string `form:"orderoptionaldefault=desc"`
	Recursive bool `form:"recursiveoptionaldefault=false"` // 是否递归列出子目录
}
```


3. response definition



```golang
type ListResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Total int64 `json:"total"`
	Files []MarkdownFile `json:"files"`
	Paths []string `json:"paths,optional"` // 子路径列表
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 7. 获取当前用户回收站文件

1. route definition

- Url: /api/v1/files/trash
- Method: GET
- Request: `ListRequest`
- Response: `ListResponse`

2. request definition



```golang
type ListRequest struct {
	Page int `form:"pageoptionaldefault=1"`
	PageSize int `form:"pageSizeoptionaldefault=20"`
	Path string `form:"pathoptionaldefault=/"` // S3路径前缀，替代文件夹ID
	SortBy string `form:"sortByoptionaldefault=updateTime"`
	Order string `form:"orderoptionaldefault=desc"`
	Recursive bool `form:"recursiveoptionaldefault=false"` // 是否递归列出子目录
}
```


3. response definition



```golang
type ListResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Total int64 `json:"total"`
	Files []MarkdownFile `json:"files"`
	Paths []string `json:"paths,optional"` // 子路径列表
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 8. 清空回收站

1. route definition

- Url: /api/v1/files/trash
- Method: DELETE
- Request: `-`
- Response: `DeleteResponse`

2. request definition



3. response definition



```golang
type DeleteResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 9. 获取预签名URL

1. route definition

- Url: /api/v1/presigned-url
- Method: POST
- Request: `PresignedUrlRequest`
- Response: `PresignedUrlResponse`

2. request definition



```golang
type PresignedUrlRequest struct {
	Filename string `json:"filename"`
	Operation string `json:"operation"` // upload, download
	ExpireIn int64 `json:"expireInoptionaldefault=3600"` // 默认1小时
}
```


3. response definition



```golang
type PresignedUrlResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	URL string `json:"url"`
	ExpireAt int64 `json:"expireAt"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 10. 恢复已删除文件

1. route definition

- Url: /api/v1/recover
- Method: POST
- Request: `RecoverFileRequest`
- Response: `RecoverFileResponse`

2. request definition



```golang
type RecoverFileRequest struct {
	Filename string `json:"filename"`
}
```


3. response definition



```golang
type RecoverFileResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 11. 搜索文件

1. route definition

- Url: /api/v1/search
- Method: GET
- Request: `SearchRequest`
- Response: `SearchResponse`

2. request definition



```golang
type SearchRequest struct {
	Keyword string `form:"keyword"`
	IsPublic bool `form:"isPublicoptional"`
	Page int `form:"pageoptionaldefault=1"`
	PageSize int `form:"pageSizeoptionaldefault=10"`
	Tags []string `form:"tagsoptional"`
	Path string `form:"pathoptional"` // 在指定路径下搜索
}
```


3. response definition



```golang
type SearchResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Total int64 `json:"total"`
	Files []MarkdownFile `json:"files"`
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 12. 分享文件

1. route definition

- Url: /api/v1/share
- Method: POST
- Request: `ShareRequest`
- Response: `ShareResponse`

2. request definition



```golang
type ShareRequest struct {
	Filename string `json:"filename"`
	ExpireIn int64 `json:"expireInoptionaldefault=86400"` // 默认24小时
	RequirePassword bool `json:"requirePasswordoptional"`
	Password string `json:"passwordoptional"`
}
```


3. response definition



```golang
type ShareResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	ShareURL string `json:"shareUrl"`
	ExpireAt int64 `json:"expireAt"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 13. 获取文件统计信息

1. route definition

- Url: /api/v1/stats
- Method: GET
- Request: `-`
- Response: `StatsResponse`

2. request definition



3. response definition



```golang
type StatsResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	TotalFiles int64 `json:"totalFiles"`
	TotalSize int64 `json:"totalSize"`
	PublicFiles int64 `json:"publicFiles"`
	PrivateFiles int64 `json:"privateFiles"`
	DeletedFiles int64 `json:"deletedFiles"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 14. 获取用户所有标签

1. route definition

- Url: /api/v1/tags
- Method: GET
- Request: `-`
- Response: `TagsResponse`

2. request definition



3. response definition



```golang
type TagsResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Tags []string `json:"tags"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 15. 上传文件

1. route definition

- Url: /api/v1/upload
- Method: POST
- Request: `UploadRequest`
- Response: `UploadResponse`

2. request definition



```golang
type UploadRequest struct {
	File []byte `form:"file"`
	IsPublic bool `form:"isPublicoptionaldefault=false"`
	Tags []string `form:"tagsoptional"`
	Filename string `form:"filenameoptional"`
	Path string `form:"pathoptionaldefault=/"` // S3路径，替代文件夹ID
}
```


3. response definition



```golang
type UploadResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	URL string `json:"url"`
	File MarkdownFile `json:"file"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type MarkdownFile struct {
	Id int64 `json:"id"`
	Filename string `json:"filename"`
	Size int64 `json:"size"`
	UploadTime int64 `json:"uploadTime"`
	UpdateTime int64 `json:"updateTime"`
	OwnerId int64 `json:"ownerId"`
	DownloadURL string `json:"downloadUrl"` // 预签名的 S3 URL
	Hash string `json:"hash"`
	IsPublic bool `json:"isPublic"`
	IsDeleted bool `json:"isDeleted"`
	Tags []string `json:"tags,optional"`
	Path string `json:"path"` // S3对象的完整路径
}
```

### 16. 获取公开文件

1. route definition

- Url: /api/v1/public/file/:hash
- Method: GET
- Request: `DownloadResponse`
- Response: `-`

2. request definition



```golang
type DownloadResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Content string `json:"content,optional"` // 小文件直接返回内容
	DownloadURL string `json:"downloadUrl"` // 预签名的S3下载URL
	File MarkdownFile `json:"file,optional"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type MarkdownFile struct {
	Id int64 `json:"id"`
	Filename string `json:"filename"`
	Size int64 `json:"size"`
	UploadTime int64 `json:"uploadTime"`
	UpdateTime int64 `json:"updateTime"`
	OwnerId int64 `json:"ownerId"`
	DownloadURL string `json:"downloadUrl"` // 预签名的 S3 URL
	Hash string `json:"hash"`
	IsPublic bool `json:"isPublic"`
	IsDeleted bool `json:"isDeleted"`
	Tags []string `json:"tags,optional"`
	Path string `json:"path"` // S3对象的完整路径
}
```


3. response definition


### 17. 浏览公开文件列表

1. route definition

- Url: /api/v1/public/files
- Method: GET
- Request: `ListRequest`
- Response: `ListResponse`

2. request definition



```golang
type ListRequest struct {
	Page int `form:"pageoptionaldefault=1"`
	PageSize int `form:"pageSizeoptionaldefault=20"`
	Path string `form:"pathoptionaldefault=/"` // S3路径前缀，替代文件夹ID
	SortBy string `form:"sortByoptionaldefault=updateTime"`
	Order string `form:"orderoptionaldefault=desc"`
	Recursive bool `form:"recursiveoptionaldefault=false"` // 是否递归列出子目录
}
```


3. response definition



```golang
type ListResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Total int64 `json:"total"`
	Files []MarkdownFile `json:"files"`
	Paths []string `json:"paths,optional"` // 子路径列表
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 18. 搜索公开文件

1. route definition

- Url: /api/v1/public/search
- Method: GET
- Request: `SearchRequest`
- Response: `SearchResponse`

2. request definition



```golang
type SearchRequest struct {
	Keyword string `form:"keyword"`
	IsPublic bool `form:"isPublicoptional"`
	Page int `form:"pageoptionaldefault=1"`
	PageSize int `form:"pageSizeoptionaldefault=10"`
	Tags []string `form:"tagsoptional"`
	Path string `form:"pathoptional"` // 在指定路径下搜索
}
```


3. response definition



```golang
type SearchResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Total int64 `json:"total"`
	Files []MarkdownFile `json:"files"`
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

### 19. 访问分享链接

1. route definition

- Url: /api/v1/public/share/access
- Method: POST
- Request: `ShareAccessRequest`
- Response: `DownloadResponse`

2. request definition



```golang
type ShareAccessRequest struct {
	ShareCode string `form:"shareCode"`
	Password string `form:"passwordoptional"`
}
```


3. response definition



```golang
type DownloadResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Content string `json:"content,optional"` // 小文件直接返回内容
	DownloadURL string `json:"downloadUrl"` // 预签名的S3下载URL
	File MarkdownFile `json:"file,optional"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type MarkdownFile struct {
	Id int64 `json:"id"`
	Filename string `json:"filename"`
	Size int64 `json:"size"`
	UploadTime int64 `json:"uploadTime"`
	UpdateTime int64 `json:"updateTime"`
	OwnerId int64 `json:"ownerId"`
	DownloadURL string `json:"downloadUrl"` // 预签名的 S3 URL
	Hash string `json:"hash"`
	IsPublic bool `json:"isPublic"`
	IsDeleted bool `json:"isDeleted"`
	Tags []string `json:"tags,optional"`
	Path string `json:"path"` // S3对象的完整路径
}
```

### 20. 热门标签

1. route definition

- Url: /api/v1/public/tags/popular
- Method: GET
- Request: `-`
- Response: `TagsResponse`

2. request definition



3. response definition



```golang
type TagsResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Tags []string `json:"tags"`
}

type BaseResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

