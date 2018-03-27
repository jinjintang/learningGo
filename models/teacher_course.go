package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TeacherCourse struct {
	Id          int    `orm:"column(ID);auto"`
	TEACHERID   string `orm:"column(TEACHER_ID);size(36);null" description:"教师编号"`
	TEACHERNAME string `orm:"column(TEACHER_NAME);size(50);null" description:"教师姓名"`
	COURSEID    int    `orm:"column(COURSE_ID);null" description:"课程名"`
	COURSENAME  string `orm:"column(COURSE_NAME);size(100);null" description:"课程编号"`
}

func (t *TeacherCourse) TableName() string {
	return "teacher_course"
}

func init() {
	orm.RegisterModel(new(TeacherCourse))
}

// AddTeacherCourse insert a new TeacherCourse into database and returns
// last inserted Id on success.
func AddTeacherCourse(m *TeacherCourse) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTeacherCourseById retrieves TeacherCourse by Id. Returns error if
// Id doesn't exist
func GetTeacherCourseById(id int) (v *TeacherCourse, err error) {
	o := orm.NewOrm()
	v = &TeacherCourse{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTeacherCourse retrieves all TeacherCourse matches certain condition. Returns empty list if
// no records exist
func GetAllTeacherCourse(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TeacherCourse))
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

	var l []TeacherCourse
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

// UpdateTeacherCourse updates TeacherCourse by Id and returns error if
// the record to be updated doesn't exist
func UpdateTeacherCourseById(m *TeacherCourse) (err error) {
	o := orm.NewOrm()
	v := TeacherCourse{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTeacherCourse deletes TeacherCourse by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTeacherCourse(id int) (err error) {
	o := orm.NewOrm()
	v := TeacherCourse{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TeacherCourse{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
