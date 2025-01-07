// /////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gorm_gen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
// /////////////////////////////////////////////////////////

package user_repo

import (
    "fmt"
    "sync"
    "time"

    "github.com/haierkeys/obsidian-image-api-gateway/global"
    "github.com/haierkeys/obsidian-image-api-gateway/internal/model"
    "github.com/haierkeys/obsidian-image-api-gateway/pkg/timex"

    "github.com/pkg/errors"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
)

var once sync.Once

func Connection() *gorm.DB {
    dbDriver := global.DBEngine
    dbDriver.Config.NamingStrategy = schema.NamingStrategy{
        TablePrefix:   "pre_", // 表名前缀
        SingularTable: true,   // 使用单数表名
    }
    // 自动创建
    if global.Config.Database.AutoMigrate {
        once.Do(func() {
            dbDriver.AutoMigrate(User{})
        })
    }
    return dbDriver
}

func NewModel() *User {
    return new(User)
}

type userRepoQueryBuilder struct {
    order []string
    where []struct {
        prefix string
        value  interface{}
    }
    whereRaw []struct {
        query  string
        values []interface{}
    }
    limit  int
    offset int
}

func NewQueryBuilder() *userRepoQueryBuilder {
    return new(userRepoQueryBuilder)
}

func (qb *userRepoQueryBuilder) buildQuery() *gorm.DB {
    ret := Connection()
    for _, where := range qb.where {
        ret = ret.Where(where.prefix, where.value)
    }
    for _, where2 := range qb.whereRaw {
        ret = ret.Where(where2.query, where2.values...)
    }
    for _, order := range qb.order {
        ret = ret.Order(order)
    }
    ret = ret.Limit(qb.limit).Offset(qb.offset)
    return ret
}

func (t *User) Create() (id int64, err error) {
    t.CreatedAt = timex.Now()
    dbDriver := Connection()
    if err = dbDriver.Model(t).Create(t).Error; err != nil {
        return 0, errors.Wrap(err, "create err")
    }
    return t.Uid, nil
}

func (t *User) Save() (err error) {
    t.UpdatedAt = timex.Now()

    dbDriver := Connection()
    if err = dbDriver.Model(t).Save(t).Error; err != nil {
        return errors.Wrap(err, "update err")
    }
    return nil
}

func (qb *userRepoQueryBuilder) Updates(m map[string]interface{}) (err error) {

    dbDriver := Connection()
    dbDriver = dbDriver.Model(&User{})

    for _, where := range qb.where {
        dbDriver.Where(where.prefix, where.value)
    }

    if err = dbDriver.Updates(m).Error; err != nil {
        return errors.Wrap(err, "updates err")
    }
    return nil
}

// 自减
func (qb *userRepoQueryBuilder) Increment(column string, value int64) (err error) {

    dbDriver := Connection()
    dbDriver = dbDriver.Model(&User{})

    for _, where := range qb.where {
        dbDriver.Where(where.prefix, where.value)
    }

    if err = dbDriver.Update(column, gorm.Expr(column+" + ?", value)).Error; err != nil {
        return errors.Wrap(err, "increment err")
    }
    return nil
}

// 自增
func (qb *userRepoQueryBuilder) Decrement(column string, value int64) (err error) {

    dbDriver := Connection()
    dbDriver = dbDriver.Model(&User{})

    for _, where := range qb.where {
        dbDriver.Where(where.prefix, where.value)
    }

    if err = dbDriver.Update(column, gorm.Expr(column+" - ?", value)).Error; err != nil {
        return errors.Wrap(err, "decrement err")
    }
    return nil
}

func (qb *userRepoQueryBuilder) Delete() (err error) {

    dbDriver := Connection()
    for _, where := range qb.where {
        dbDriver = dbDriver.Where(where.prefix, where.value)
    }

    if err = dbDriver.Delete(&User{}).Error; err != nil {
        return errors.Wrap(err, "delete err")
    }
    return nil
}

func (qb *userRepoQueryBuilder) Count() (int64, error) {
    var c int64
    res := qb.buildQuery().Model(&User{}).Count(&c)
    if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
        c = 0
    }
    return c, res.Error
}

func (qb *userRepoQueryBuilder) First() (*User, error) {
    ret := &User{}
    res := qb.buildQuery().First(ret)
    if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
        ret = nil
    }
    return ret, res.Error
}

func (qb *userRepoQueryBuilder) Get() ([]*User, error) {
    return qb.QueryAll()
}

func (qb *userRepoQueryBuilder) QueryOne() (*User, error) {
    qb.limit = 1
    ret, err := qb.QueryAll()
    if len(ret) > 0 {
        return ret[0], err
    }
    return nil, err
}

func (qb *userRepoQueryBuilder) QueryAll() ([]*User, error) {
    var ret []*User
    err := qb.buildQuery().Find(&ret).Error
    return ret, err
}

func (qb *userRepoQueryBuilder) Limit(limit int) *userRepoQueryBuilder {
    qb.limit = limit
    return qb
}

func (qb *userRepoQueryBuilder) Offset(offset int) *userRepoQueryBuilder {
    qb.offset = offset
    return qb
}

func (qb *userRepoQueryBuilder) WhereRaw(query string, values ...interface{}) *userRepoQueryBuilder {
    vals := make([]interface{}, len(values))
    // nolint:S1001
    for i, v := range values {
        vals[i] = v
    }
    qb.whereRaw = append(qb.whereRaw, struct {
        query  string
        values []interface{}
    }{
        query,
        vals,
    })
    return qb
}

// ----------

func (qb *userRepoQueryBuilder) WhereUid(p model.Predicate, value int64) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "uid", p),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereUidIn(value []int64) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "uid", "IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereUidNotIn(value []int64) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "uid", "NOT IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) OrderByUid(asc bool) *userRepoQueryBuilder {
    order := "DESC"
    if asc {
        order = "ASC"
    }

    qb.order = append(qb.order, "`uid` "+order)
    return qb
}

func (qb *userRepoQueryBuilder) WhereEmail(p model.Predicate, value string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "email", p),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereEmailIn(value []string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "email", "IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereEmailNotIn(value []string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "email", "NOT IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) OrderByEmail(asc bool) *userRepoQueryBuilder {
    order := "DESC"
    if asc {
        order = "ASC"
    }

    qb.order = append(qb.order, "`email` "+order)
    return qb
}

func (qb *userRepoQueryBuilder) WhereUsername(p model.Predicate, value string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "username", p),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereUsernameIn(value []string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "username", "IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereUsernameNotIn(value []string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "username", "NOT IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) OrderByUsername(asc bool) *userRepoQueryBuilder {
    order := "DESC"
    if asc {
        order = "ASC"
    }

    qb.order = append(qb.order, "`username` "+order)
    return qb
}

func (qb *userRepoQueryBuilder) WherePassword(p model.Predicate, value string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "password", p),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WherePasswordIn(value []string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "password", "IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WherePasswordNotIn(value []string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "password", "NOT IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) OrderByPassword(asc bool) *userRepoQueryBuilder {
    order := "DESC"
    if asc {
        order = "ASC"
    }

    qb.order = append(qb.order, "`password` "+order)
    return qb
}

func (qb *userRepoQueryBuilder) WhereSalt(p model.Predicate, value string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "salt", p),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereSaltIn(value []string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "salt", "IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereSaltNotIn(value []string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "salt", "NOT IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) OrderBySalt(asc bool) *userRepoQueryBuilder {
    order := "DESC"
    if asc {
        order = "ASC"
    }

    qb.order = append(qb.order, "`salt` "+order)
    return qb
}

func (qb *userRepoQueryBuilder) WhereToken(p model.Predicate, value string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "token", p),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereTokenIn(value []string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "token", "IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereTokenNotIn(value []string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "token", "NOT IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) OrderByToken(asc bool) *userRepoQueryBuilder {
    order := "DESC"
    if asc {
        order = "ASC"
    }

    qb.order = append(qb.order, "`token` "+order)
    return qb
}

func (qb *userRepoQueryBuilder) WhereAvatar(p model.Predicate, value string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "avatar", p),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereAvatarIn(value []string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "avatar", "IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereAvatarNotIn(value []string) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "avatar", "NOT IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) OrderByAvatar(asc bool) *userRepoQueryBuilder {
    order := "DESC"
    if asc {
        order = "ASC"
    }

    qb.order = append(qb.order, "`avatar` "+order)
    return qb
}

func (qb *userRepoQueryBuilder) WhereIsDeleted(p model.Predicate, value int64) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "is_deleted", p),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereIsDeletedIn(value []int64) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereIsDeletedNotIn(value []int64) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) OrderByIsDeleted(asc bool) *userRepoQueryBuilder {
    order := "DESC"
    if asc {
        order = "ASC"
    }

    qb.order = append(qb.order, "`is_deleted` "+order)
    return qb
}

func (qb *userRepoQueryBuilder) WhereUpdatedAt(p model.Predicate, value time.Time) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "updated_at", p),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereUpdatedAtIn(value []time.Time) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "updated_at", "IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) OrderByUpdatedAt(asc bool) *userRepoQueryBuilder {
    order := "DESC"
    if asc {
        order = "ASC"
    }

    qb.order = append(qb.order, "`updated_at` "+order)
    return qb
}

func (qb *userRepoQueryBuilder) WhereCreatedAt(p model.Predicate, value time.Time) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "created_at", p),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereCreatedAtIn(value []time.Time) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "created_at", "IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) OrderByCreatedAt(asc bool) *userRepoQueryBuilder {
    order := "DESC"
    if asc {
        order = "ASC"
    }

    qb.order = append(qb.order, "`created_at` "+order)
    return qb
}

func (qb *userRepoQueryBuilder) WhereDeletedAt(p model.Predicate, value time.Time) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "deleted_at", p),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereDeletedAtIn(value []time.Time) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "deleted_at", "IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) WhereDeletedAtNotIn(value []time.Time) *userRepoQueryBuilder {
    qb.where = append(qb.where, struct {
        prefix string
        value  interface{}
    }{
        fmt.Sprintf("%v %v ?", "deleted_at", "NOT IN"),
        value,
    })
    return qb
}

func (qb *userRepoQueryBuilder) OrderByDeletedAt(asc bool) *userRepoQueryBuilder {
    order := "DESC"
    if asc {
        order = "ASC"
    }

    qb.order = append(qb.order, "`deleted_at` "+order)
    return qb
}
