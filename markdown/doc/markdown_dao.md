# markdown_dao

## markdown 数据访问对象层介绍

用户上传的markdown文件实际存储在S3服务中，使用 PostgreSQL 存用户与 markdown 文件的元数据关系。

> 根据这个 MarkdownFile 构建 markdown 文件存储元数据

```go
// Markdown 文件结构 - 保留必要字段
type MarkdownFile {
	Id          int64    `json:"id"`
	Filename    string   `json:"filename"`
	Size        int64    `json:"size"`
	UploadTime  int64    `json:"uploadTime"`
	UpdateTime  int64    `json:"updateTime"`
	OwnerId     int64    `json:"ownerId"`
	DownloadURL string   `json:"downloadUrl"` // 预签名的 S3 URL
	Hash        string   `json:"hash"`
	IsPublic    bool     `json:"isPublic"`
	IsDeleted   bool     `json:"isDeleted"`
	Tags        []string `json:"tags,optional"`
	Path        string   `json:"path"` // S3对象的完整路径
}
```

> 设计


