package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CourseAdvertise struct {
	Id          int    `orm:"column(COURSE_ADVERTISE_ID);auto"`
	COURSEID    int    `orm:"column(COURSE_ID);null" description:"课程id"`
	ADVERTISEID int    `orm:"column(ADVERTISE_ID);null" description:"广告id"`
	REMARK      string `orm:"column(REMARK);size(100);null" description:"备注"`
	STATE       int    `orm:"column(STATE);null" description:"0开启1停用"`
}

func (t *CourseAdvertise) TableName() string {
	return "course_advertise"
}

func init() {
	orm.RegisterModel(new(CourseAdvertise))
}

// AddCourseAdvertise insert a new CourseAdvertise into database and returns
// last inserted Id on success.
func AddCourseAdvertise(m *CourseAdvertise) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCourseAdvertiseById retrieves CourseAdvertise by Id. Returns error if
// Id doesn't exist
func GetCourseAdvertiseById(id int) (v *CourseAdvertise, err error) {
	o := orm.NewOrm()
	v = &CourseAdvertise{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCourseAdvertise retrieves all CourseAdvertise matches certain condition. Returns empty list if
// no records exist
func GetAllCourseAdvertise(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CourseAdvertise))
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

	var l []CourseAdvertise
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

// UpdateCourseAdvertise updates CourseAdvertise by Id and returns error if
// the record to be updated doesn't exist
func UpdateCourseAdvertiseById(m *CourseAdvertise) (err error) {
	o := orm.NewOrm()
	v := CourseAdvertise{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCourseAdvertise deletes CourseAdvertise by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCourseAdvertise(id int) (err error) {
	o := orm.NewOrm()
	v := CourseAdvertise{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CourseAdvertise{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
