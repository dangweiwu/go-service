package query

/*
@Time : 2021/7/28 21:42
@Author : dang
@desz:
基于 mysql gorm 检索
支持 模糊和精确查询
支持 分页查询
支持 排序 order by
使用说明

page := NewQuery(ctx,rule,po,pos)
1. 修改where 一旦定义则默认失效
page.SetWhere(func(db *gorm.Db){
	if start,has:=page.QData["start"];has{
		delete data.QData["start"]
		db = db.Where("created_at > ?",start)
	}
	if end,has := page.QData["end"];has{
		delete data.QData["end"]
		db = db.Where("created_at < ?",end)
	}
	return page.Where(db)
})

2. 修改Order 一旦定义则默认失效
前端 传递 od=created_at desc,id
page.SetOrder(func(db *gorm.Db){
	return db.Order("id desc")
})

3. page 失效
page.ClearPage()
*/
import (
	// "admin/internal/app_pkg/mysql"

	"go-service/internal/apiservice/pkg/ginx"
	"gorm.io/gorm"
	"strconv"
	"time"
)

const (
	ORDER = "ord"
	START = "start"
	END   = "end"
)

// @doc | query.QueryResult
type QueryPageResult struct {
	Page *Page       `json:"page" doc:"|d 分页数据 |c 参考Page定义"`
	Data interface{} `json:"data" doc:"|d 数据 |c 参考Data定义"`
}

type QueryListResult struct {
	Data interface{}
}

type Query struct {
	Db *gorm.DB
	g  ginx.Giner
}

func NewQuery(db *gorm.DB, g ginx.Giner) *Query {
	return &Query{Db: db, g: g}
}

func (this *Query) SetTable(po interface{}) *Query {
	this.Db = this.Db.Model(po)
	return this
}

// where 相关操作
func (this *Query) WhereLike(likeQuery []string) *Query {
	_lq := map[string]string{}

	for _, v := range likeQuery {
		r := this.g.Query(v)
		if len(r) != 0 {
			_lq[v] = r
		}
	}

	if len(_lq) != 0 {
		for k, v := range _lq {
			this.Db = this.Db.Where(k+" like ?", "%"+v+"%")
		}
	}
	return this
}
func (this *Query) Where(query []string) *Query {
	_q := map[string]string{}
	for _, v := range query {
		r := this.g.Query(v)
		if len(r) != 0 {
			_q[v] = r
		}
	}

	if len(_q) != 0 {
		for k, v := range _q {
			this.Db = this.Db.Where(k+" = ?", v)
		}
	}

	return this
}

func (this *Query) WhereRangeTimestamp(dbTimeName string) *Query {
	start := this.g.Query(START)
	end := this.g.Query(END)
	if len(start) != 0 {
		if s, err := strconv.ParseInt(start, 10, 64); err == nil {
			this.Db = this.Db.Where(dbTimeName+" >= ?", time.Unix(s, 0))
		}
	}
	if len(end) != 0 {
		if s, err := strconv.ParseInt(end, 10, 64); err == nil {
			this.Db = this.Db.Where(dbTimeName+" <= ?", time.Unix(s, 0))
		}
	}
	return this
}

func (this *Query) WhereRangeDate(dbTimeName string) *Query {
	start := this.g.Query(START)
	end := this.g.Query(END)
	if len(start) != 0 {
		if t, err := time.Parse(time.DateTime, start); err == nil {
			this.Db = this.Db.Where(dbTimeName+" >= ?", t)
		}
	}

	if len(end) != 0 {
		if t, err := time.Parse(time.DateTime, end); err == nil {
			this.Db = this.Db.Where(dbTimeName+" <= ?", t)
		}
	}
	return this
}

func (this *Query) Order() *Query {
	od := this.g.Query(ORDER)
	if len(od) != 0 {
		this.Db = this.Db.Order(od)
	} else {
		this.Db = this.Db.Order("created_at desc")
	}
	return this
}

func (this *Query) Select(s []string) *Query {
	this.Db = this.Db.Select(s)
	return this
}

func (this *Query) PageFind(pos interface{}) (*QueryPageResult, error) {
	page, err := ParsePage(this.g)
	if err != nil {
		return nil, err
	}

	if r := this.Db.Count(&(page.Total)); r.Error != nil {
		return nil, r.Error
	}
	offset := (page.Current - 1) * page.Limit

	if r := this.Db.Offset(offset).Limit(page.Limit).Find(pos); r.Error != nil {
		return nil, r.Error
	}
	return &QueryPageResult{page, pos}, nil
}

func (this *Query) ListFind(pos interface{}) (*QueryListResult, error) {
	if r := this.Db.Find(pos); r.Error != nil {
		return nil, r.Error
	}
	return &QueryListResult{pos}, nil
}
