-- 创建markdown_files表
CREATE TABLE IF NOT EXISTS markdown_files (
    id BIGSERIAL PRIMARY KEY,
    filename VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL,
    upload_time TIMESTAMP NOT NULL,
    update_time TIMESTAMP NOT NULL,
    owner_id BIGINT NOT NULL,
    hash VARCHAR(64) NOT NULL,
    is_public BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    tags JSONB DEFAULT '[]'::JSONB,
    path VARCHAR(512) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 创建索引
CREATE INDEX idx_markdown_files_filename ON markdown_files(filename);
CREATE INDEX idx_markdown_files_upload_time ON markdown_files(upload_time);
CREATE INDEX idx_markdown_files_update_time ON markdown_files(update_time);
CREATE INDEX idx_markdown_files_owner_id ON markdown_files(owner_id);
CREATE UNIQUE INDEX idx_markdown_files_hash ON markdown_files(hash);
CREATE INDEX idx_markdown_files_is_public ON markdown_files(is_public);
CREATE INDEX idx_markdown_files_is_deleted ON markdown_files(is_deleted);
CREATE INDEX idx_markdown_files_path ON markdown_files(path);
CREATE INDEX idx_markdown_files_deleted_at ON markdown_files(deleted_at);

-- 添加注释
COMMENT ON TABLE markdown_files IS 'Markdown文件元数据表';
COMMENT ON COLUMN markdown_files.id IS '主键ID';
COMMENT ON COLUMN markdown_files.filename IS '文件名';
COMMENT ON COLUMN markdown_files.size IS '文件大小(字节)';
COMMENT ON COLUMN markdown_files.upload_time IS '上传时间';
COMMENT ON COLUMN markdown_files.update_time IS '更新时间';
COMMENT ON COLUMN markdown_files.owner_id IS '所有者ID';
COMMENT ON COLUMN markdown_files.hash IS '文件哈希值';
COMMENT ON COLUMN markdown_files.is_public IS '是否公开';
COMMENT ON COLUMN markdown_files.is_deleted IS '是否删除(移入回收站)';
COMMENT ON COLUMN markdown_files.tags IS '标签(JSON数组格式)';
COMMENT ON COLUMN markdown_files.path IS '文件路径';
COMMENT ON COLUMN markdown_files.created_at IS '记录创建时间';
COMMENT ON COLUMN markdown_files.updated_at IS '记录更新时间';
COMMENT ON COLUMN markdown_files.deleted_at IS '记录软删除时间';
