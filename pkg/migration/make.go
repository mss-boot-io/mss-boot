package migration

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/8/12 09:15:17
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/8/12 09:15:17
 */

import (
	"bytes"
	"embed"
	"gorm.io/gorm/schema"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
	"text/template"
	"time"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

//go:embed *.tpl
var FS embed.FS

var Migrate = &Migration{
	version: make(map[int]func(db *gorm.DB, version string) error),
}

type Migration struct {
	db      *gorm.DB
	version map[int]func(db *gorm.DB, version string) error
	mutex   sync.Mutex
}

func (e *Migration) GetDb() *gorm.DB {
	return e.db
}

func (e *Migration) SetDb(db *gorm.DB) {
	e.db = db
}

func (e *Migration) SetVersion(k int, f func(db *gorm.DB, version string) error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.version[k] = f
}

func (e *Migration) Migrate(ms ...schema.Tabler) {
	versions := make([]int, 0)
	for k := range e.version {
		versions = append(versions, k)
	}
	if !sort.IntsAreSorted(versions) {
		sort.Ints(versions)
	}
	var err error
	var count int64
	for _, v := range versions {
		if len(ms) == 0 {
			err = e.db.Table("mss_boot_migration").Where("version = ?", v).Count(&count).Error
		} else {
			err = e.db.Model(ms[0]).Where("version = ?", v).Count(&count).Error
		}
		if err != nil {
			log.Fatalf("get migration version error: %v", err)
		}
		if count > 0 {
			log.Println(count)
			count = 0
			continue
		}
		err = (e.version[v])(e.db, strconv.Itoa(v))
		if err != nil {
			log.Fatalf("migrate version %d error: %v", v, err)
		}
	}
}

func GetFilename(s string) int {
	s = filepath.Base(s)
	return cast.ToInt(s[:13])
}

func GenFile(system bool, path string) error {
	t1, err := template.ParseFS(FS, "migrate.tpl")
	if err != nil {
		slog.Error("parse template error", slog.Any("err", err))
		return err
	}
	m := map[string]string{}
	m["GenerateTime"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	m["Package"] = "custom"
	if system {
		m["Package"] = "system"
	}
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	if err != nil {
		slog.Error("execute template error", slog.Any("err", err))
		return err
	}
	if system {
		fileCreate(b1, filepath.Join(path, "system", m["GenerateTime"]+"_migrate.go"))
	} else {
		fileCreate(b1, filepath.Join(path, "custom", m["GenerateTime"]+"_migrate.go"))
	}
	return nil
}

func fileCreate(content bytes.Buffer, name string) {
	file, err := os.Create(name)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)
	if err != nil {
		log.Println(err)
	}
	_, err = file.WriteString(content.String())
	if err != nil {
		log.Println(err)
	}
}
