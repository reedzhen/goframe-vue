package gfres

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/golang-migrate/migrate/v4/source"
	"io"
	"os"
)

func init() {
	source.Register("gfres", &GfRes{})
}

type GfRes struct {
	path       string
	migrations *source.Migrations
}

func New(path string) (source.Driver, error) {
	ms := source.NewMigrations()
	files := gres.ScanDirFile(path, "*")
	for _, v := range files {
		m, err := source.DefaultParse(v.FileInfo().Name())
		if err != nil {
			continue
		}
		if !ms.Append(m) {
			return nil, errors.New("文件添加失败")
		}
	}
	return &GfRes{
		path:       path,
		migrations: ms,
	}, nil
}

func (g *GfRes) Open(url string) (source.Driver, error) {
	return nil, errors.New("open() cannot be called on the gres passthrough driver")
}

func (g *GfRes) Close() error {
	return nil
}

func (g *GfRes) First() (version uint, err error) {
	if v, ok := g.migrations.First(); !ok {
		return 0, &os.PathError{Op: "first", Path: g.path, Err: os.ErrNotExist}
	} else {
		return v, nil
	}
}

func (g *GfRes) Prev(version uint) (prevVersion uint, err error) {
	if v, ok := g.migrations.Prev(version); !ok {
		return 0, &os.PathError{Op: fmt.Sprintf("prev for version %v", version), Path: g.path, Err: os.ErrNotExist}
	} else {
		return v, nil
	}
}

func (g *GfRes) Next(version uint) (nextVersion uint, err error) {
	if v, ok := g.migrations.Next(version); !ok {
		return 0, &os.PathError{Op: fmt.Sprintf("next for version %v", version), Path: g.path, Err: os.ErrNotExist}
	} else {
		return v, nil
	}
}

func (g *GfRes) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := g.migrations.Up(version); ok {
		body := gres.GetContent(g.path + "/" + m.Raw)

		return io.NopCloser(bytes.NewReader(body)), m.Identifier, nil
	}
	return nil, "", &os.PathError{Op: fmt.Sprintf("read version %v", version), Path: g.path, Err: os.ErrNotExist}
}

func (g *GfRes) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := g.migrations.Down(version); ok {
		body := gres.GetContent(g.path + "/" + m.Raw)

		return io.NopCloser(bytes.NewReader(body)), m.Identifier, nil
	}
	return nil, "", &os.PathError{Op: fmt.Sprintf("read version %v", version), Path: g.path, Err: os.ErrNotExist}
}
