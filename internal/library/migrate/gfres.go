package migrate

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/golang-migrate/migrate/v4/source"
)

type gfRes struct {
	path       string
	migrations *source.Migrations
}

func newGfResourceSource(path string) (source.Driver, error) {
	ms := source.NewMigrations()
	files := gres.ScanDirFile(path, "*")
	for _, file := range files {
		migration, err := source.DefaultParse(file.FileInfo().Name())
		if err != nil {
			continue
		}
		if !ms.Append(migration) {
			return nil, gerror.New("迁移文件添加失败")
		}
	}
	return &gfRes{
		path:       path,
		migrations: ms,
	}, nil
}

// Open 禁止通过 url 打开 gf resource，启动时直接传入 source 实例。
func (s *gfRes) Open(url string) (source.Driver, error) {
	return nil, gerror.New("gf resource 迁移源不支持 Open")
}

// Close 关闭迁移源。
func (s *gfRes) Close() error {
	return nil
}

// First 返回第一个迁移版本。
func (s *gfRes) First() (version uint, err error) {
	if v, ok := s.migrations.First(); ok {
		return v, nil
	}
	return 0, &os.PathError{Op: "first", Path: s.path, Err: os.ErrNotExist}
}

// Prev 返回上一个迁移版本。
func (s *gfRes) Prev(version uint) (prevVersion uint, err error) {
	if v, ok := s.migrations.Prev(version); ok {
		return v, nil
	}
	return 0, &os.PathError{Op: fmt.Sprintf("prev for version %v", version), Path: s.path, Err: os.ErrNotExist}
}

// Next 返回下一个迁移版本。
func (s *gfRes) Next(version uint) (nextVersion uint, err error) {
	if v, ok := s.migrations.Next(version); ok {
		return v, nil
	}
	return 0, &os.PathError{Op: fmt.Sprintf("next for version %v", version), Path: s.path, Err: os.ErrNotExist}
}

// ReadUp 读取指定版本的 up 脚本。
func (s *gfRes) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := s.migrations.Up(version); ok {
		body := gres.GetContent(s.path + "/" + m.Raw)
		return io.NopCloser(bytes.NewReader(body)), m.Identifier, nil
	}
	return nil, "", &os.PathError{Op: fmt.Sprintf("read up version %v", version), Path: s.path, Err: os.ErrNotExist}
}

// ReadDown 读取指定版本的 down 脚本。
func (s *gfRes) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := s.migrations.Down(version); ok {
		body := gres.GetContent(s.path + "/" + m.Raw)
		return io.NopCloser(bytes.NewReader(body)), m.Identifier, nil
	}
	return nil, "", &os.PathError{Op: fmt.Sprintf("read down version %v", version), Path: s.path, Err: os.ErrNotExist}
}
