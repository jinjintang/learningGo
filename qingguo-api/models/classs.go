package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
"strconv"  
	"github.com/astaxie/beego/orm"
)

type Classs struct {
	Id           string   `orm:"column(CLASS_ID);pk"`
	CLASSNAME    string `orm:"column(CLASS_NAME);size(20);null" description:"班级名字"`
	CLASSMONITOR string `orm:"column(CLASS_MONITOR);size(10);null" description:"班长"`
	CLASSNUMBER  int    `orm:"column(CLASS_NUMBER);null" description:"班级人数"`
	CLASSTIME    string `orm:"column(CLASS_TIME);size(30);null" description:"上课时间"`
	CLASSADDRESS string `orm:"column(CLASS_ADDRESS);size(30);null" description:"上课地点"`
	AREAID       string `orm:"column(AREA_ID);size(36);null" description:"区域ID"`
	AREANAME     string `orm:"column(AREA_NAME);size(20);null" description:"区域名称"`
	CLASSCISHU   int    `orm:"column(CLASS_CISHU);null" description:"次数"`
	TEACHERID    string `orm:"column(TEACHER_ID);size(36);null" description:"老师编号"`
	TEACHERNAME  string `orm:"column(TEACHER_NAME);size(36);null" description:"老师姓名"`
	CLASSSSTATE  int    `orm:"column(CLASSS_STATE);null" description:"班级状态0表示上课中1表示已结课2待审查"`
}

func (t *Classs) TableName() string {
	return "classs"
}

func init() {
	orm.RegisterModel(new(Classs))
}

// AddClasss insert a new Classs into database and returns
// last inserted Id on success.
func AddClasss(m *Classs) (id string, err error) {
	o := orm.NewOrm()
	sid, serr := o.Insert(m)
	id=strconv.FormatInt(sid,10)
    err=serr
	return
}

// GetClasssById retrieves Classs by Id. Returns error if
// Id doesn't exist
func GetClasssById(id string) (v *Classs, err error) {
	o := orm.NewOrm()
	v = &Classs{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllClasss retrieves all Classs matches certain condition. Returns empty list if
// no records exist
func GetAllClasss(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Classs))
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

	var l []Classs
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

// UpdateClasss updates Classs by Id and returns error if
// the record to be updated doesn't exist
func UpdateClasssById(m *Classs) (err error) {
	o := orm.NewOrm()
	v := Classs{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteClasss deletes Classs by Id and returns error if
// the record to be deleted doesn't exist
func DeleteClasss(id string) (err error) {
	o := orm.NewOrm()
	v := Classs{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Classs{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
