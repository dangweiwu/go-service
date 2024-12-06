package query

import (
	"fmt"
	"github.com/dangweiwu/microkit/db/mysqlx"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/pkg/model"
	"gorm.io/gorm"
	"strconv"
	"testing"
	"time"
)

type Demo struct {
	model.Model
	Name  string
	ActAt time.Time `gorm:"type:timestamp"`
	EndAt int
	DayAt time.Time
}

func TestQuery(t *testing.T) {
	cli, err := mysqlx.NewClient(mysqlx.Config{
		User:     "root",
		Password: "a12346",
		Host:     "127.0.0.1:13306",
		DbName:   "goservice",
		LogLevel: 4,
	})

	if err != nil {
		t.Error(err)
		return
	}
	SearCh(cli, t)

	PageQuery(cli, t)

	PageQuery2(cli, t)
}

func PageQuery(cli *gorm.DB, t *testing.T) {
	pos := []Demo{}
	q := NewQuery(cli, &DbGiner{}).SetTable(&Demo{})

	q.WhereLike([]string{"name", "end_at"}).
		WhereRangeTimestamp("created_at").
		Order()
	r, err := q.PageFind(&pos)
	t.Log(r, r.Page, r.Data)
	t.Log(err)

}

func PageQuery2(cli *gorm.DB, t *testing.T) {
	pos := []Demo{}
	q := NewQuery(cli, &DbGiner{}).SetTable(&Demo{})

	q.Where([]string{"name", "end_at"}).
		WhereLike([]string{"name", "end_at"}).
		WhereRangeTimestamp("created_at").
		Order()
	r, err := q.PageFind(&pos)
	t.Log(r, r.Page, r.Data)
	t.Log(err)

}

func Create(cli *gorm.DB, t *testing.T) {
	cli.Migrator().AutoMigrate(&Demo{})
	t.Log(time.Now())
	for i := 0; i < 1000; i++ {
		po := &Demo{
			Name:  "name_" + strconv.Itoa(i),
			ActAt: time.Now(),
			EndAt: int(time.Now().Unix()),
			DayAt: time.Now(),
		}
		err := cli.Create(po)
		if err != nil {
			t.Log("no:", i)
			t.Error(err)
		}
	}
	t.Log(time.Now())
}

func SearCh(cli *gorm.DB, t *testing.T) {
	po := &Demo{}
	pos := []Demo{}

	q := NewQuery(cli, &DbGiner{})
	q.Db = q.Db.Model(po)
	q.WhereRangeTimestamp("act_at")
	q.Db.Find(&pos)
	fmt.Println(len(pos), pos[0])

	q2 := NewQuery(cli, &DbGinerTime{})
	q2.WhereRangeDate("act_at")
	q2.Db.Find(&pos)
	fmt.Println(len(pos))

}

type DbGiner struct {
	ginx.EmptyGinx
}

func (this DbGiner) Query(key string) (value string) {
	fmt.Println("key:", key)
	if key == "start" {
		return strconv.Itoa(int(time.Now().Add(-time.Hour * 10).Unix()))
	} else if key == "end" {
		return strconv.Itoa(int(time.Now().Unix()))
	} else if key == "name" {
		return "name_1"
	} else if key == "end_at" {
		return "1733451952"
	}
	return ""
}

func (this DbGiner) ShouldBindQuery(obj any) error {
	page, ok := obj.(*Page)
	if !ok {
		return fmt.Errorf("obj is not of type *Page")
	}
	fmt.Println("shouldBindQUeyr")
	page.Current = 1
	page.Limit = 50
	return nil
}

type DbGinerTime struct {
	ginx.EmptyGinx
}

func (this DbGinerTime) Query(key string) (value string) {
	fmt.Println("key:", key)
	if key == "start" {
		return time.Now().Add(-time.Hour * 10).Format(time.DateTime)
	} else if key == "end" {
		return time.Now().Format(time.DateTime)
	} else if key == "name" {
		return "name_1"
	} else if key == "end_at" {
		return "1733451952"
	}
	return ""
}
func (this DbGinerTime) ShouldBindQuery(obj any) error {
	page, ok := obj.(*Page)
	if !ok {
		return fmt.Errorf("obj is not of type *Page")
	}
	fmt.Println("shouldBindQUeyr")
	page.Current = 0
	page.Limit = 50
	return nil
}
