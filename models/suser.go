package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
 "strconv"  
	"github.com/astaxie/beego/orm"
)

type Suser struct {
	Id          string    `orm:"column(SYSTEM_USER_ID);pk"`
	NAME        string `orm:"column(NAME);size(10);null" description:"姓名"`
	OPENID      string `orm:"column(OPENID);size(100);null" description:"openid"`
	WEIXIN      string `orm:"column(WEIXIN);size(30);null" description:"微信"`
	PHONENUMBER string `orm:"column(PHONENUMBER);size(11);null" description:"电话"`
	PHOTOPATH   string `orm:"column(PHOTOPATH);size(50);null" description:"头像路径"`
	EMAIL       string `orm:"column(EMAIL);size(30);null" description:"邮箱"`
	SEX         int    `orm:"column(SEX);null" description:"0男1女"`
	AREAID      string `orm:"column(AREA_ID);size(36);null" description:"所属区域id"`
	AREANAME    string `orm:"column(AREA_NAME);size(20);null" description:"所属区域"`
	PASSWORD    string `orm:"column(PASSWORD);size(50);null" description:"密码"`
	STATE       int    `orm:"column(STATE);null" description:"状态"`
}

func (t *Suser) TableName() string {
	return "suser"
}

func init() {
	orm.RegisterModel(new(Suser))
}

// AddSuser insert a new Suser into database and returns
// last inserted Id on success.
func AddSuser(m *Suser) (id string, err error) {
	o := orm.NewOrm()
	
	sid, serr := o.Insert(m)
	id=strconv.FormatInt(sid,10)
    err=serr
	return
}

// GetSuserById retrieves Suser by Id. Returns error if
// Id doesn't exist
func GetSuserById(id string) (v *Suser, err error) {
	o := orm.NewOrm()
	v = &Suser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSuser retrieves all Suser matches certain condition. Returns empty list if
// no records exist
func GetAllSuser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Suser))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Suser
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSuser updates Suser by Id and returns error if
// the record to be updated doesn't exist
func UpdateSuserById(m *Suser) (err error) {
	o := orm.NewOrm()
	v := Suser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSuser deletes Suser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSuser(id string) (err error) {
	o := orm.NewOrm()
	v := Suser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Suser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
