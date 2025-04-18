// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1

package types

type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type BatchDeleteRequest struct {
	Filenames []string `json:"filenames"`
}

type BatchDeleteResponse struct {
	BaseResponse
	SuccessCount int `json:"successCount"`
	FailedCount  int `json:"failedCount"`
}

type DeleteResponse struct {
	BaseResponse
}

type DownloadResponse struct {
	BaseResponse
	Content     string       `json:"content,optional"` // 小文件直接返回内容
	DownloadURL string       `json:"downloadUrl"`      // 预签名的S3下载URL
	File        MarkdownFile `json:"file,optional"`
}

type ErrorCode struct {
	Success       int `json:"success"`       // 200: 成功
	BadRequest    int `json:"badRequest"`    // 400: 请求参数错误
	Unauthorized  int `json:"unauthorized"`  // 401: 未授权
	Forbidden     int `json:"forbidden"`     // 403: 禁止访问
	NotFound      int `json:"notFound"`      // 404: 资源不存在
	FileTooLarge  int `json:"fileTooLarge"`  // 413: 文件过大
	InternalError int `json:"internalError"` // 500: 内部服务器错误
	StorageError  int `json:"storageError"`  // 507: 存储空间不足
}

type ListRequest struct {
	Page      int    `form:"pageoptionaldefault=1"`
	PageSize  int    `form:"pageSizeoptionaldefault=20"`
	Path      string `form:"pathoptionaldefault=/"` // S3路径前缀，替代文件夹ID
	SortBy    string `form:"sortByoptionaldefault=updateTime"`
	Order     string `form:"orderoptionaldefault=desc"`
	Recursive bool   `form:"recursiveoptionaldefault=false"` // 是否递归列出子目录
}

type ListResponse struct {
	BaseResponse
	Total    int64          `json:"total"`
	Files    []MarkdownFile `json:"files"`
	Paths    []string       `json:"paths,optional"` // 子路径列表
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
}

type MarkdownFile struct {
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

type PresignedUrlRequest struct {
	Filename  string `json:"filename"`
	Operation string `json:"operation"`                    // upload, download
	ExpireIn  int64  `json:"expireInoptionaldefault=3600"` // 默认1小时
}

type PresignedUrlResponse struct {
	BaseResponse
	URL      string `json:"url"`
	ExpireAt int64  `json:"expireAt"`
}

type RecoverFileRequest struct {
	Filename string `json:"filename"`
}

type RecoverFileResponse struct {
	BaseResponse
}

type SearchRequest struct {
	Keyword  string   `form:"keyword"`
	IsPublic bool     `form:"isPublicoptional"`
	Page     int      `form:"pageoptionaldefault=1"`
	PageSize int      `form:"pageSizeoptionaldefault=10"`
	Tags     []string `form:"tagsoptional"`
	Path     string   `form:"pathoptional"` // 在指定路径下搜索
}

type SearchResponse struct {
	BaseResponse
	Total    int64          `json:"total"`
	Files    []MarkdownFile `json:"files"`
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
}

type ShareAccessRequest struct {
	ShareCode string `form:"shareCode"`
	Password  string `form:"passwordoptional"`
}

type ShareRequest struct {
	Filename        string `json:"filename"`
	ExpireIn        int64  `json:"expireInoptionaldefault=86400"` // 默认24小时
	RequirePassword bool   `json:"requirePasswordoptional"`
	Password        string `json:"passwordoptional"`
}

type ShareResponse struct {
	BaseResponse
	ShareURL string `json:"shareUrl"`
	ExpireAt int64  `json:"expireAt"`
}

type StatsResponse struct {
	BaseResponse
	TotalFiles   int64 `json:"totalFiles"`
	TotalSize    int64 `json:"totalSize"`
	PublicFiles  int64 `json:"publicFiles"`
	PrivateFiles int64 `json:"privateFiles"`
	DeletedFiles int64 `json:"deletedFiles"`
}

type TagsResponse struct {
	BaseResponse
	Tags []string `json:"tags"`
}

type UpdateFileRequest struct {
	IsPublic bool     `json:"isPublicoptional"`
	Content  string   `json:"contentoptional"`
	Tags     []string `json:"tagsoptional"`
	Filename string   `json:"filenameoptional"` // 重命名
	Path     string   `json:"pathoptional"`     // 移动文件到其他路径
}

type UpdateFileResponse struct {
	BaseResponse
	File MarkdownFile `json:"file"`
}

type UploadRequest struct {
	File     []byte   `form:"file"`
	IsPublic bool     `form:"isPublicoptionaldefault=false"`
	Tags     []string `form:"tagsoptional"`
	Filename string   `form:"filenameoptional"`
	Path     string   `form:"pathoptionaldefault=/"` // S3路径，替代文件夹ID
}

type UploadResponse struct {
	BaseResponse
	URL  string       `json:"url"`
	File MarkdownFile `json:"file"`
}
