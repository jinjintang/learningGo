package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CourseCategory struct {
	Id             string  `orm:"column(CATEGORY_ID);pk" description:"类别主键"`
	CATEGORYNAME   string `orm:"column(CATEGORY_NAME);size(20);null" description:"类别名称(801)"`
	CATEGORYREMARK string `orm:"column(CATEGORY_REMARK);size(500);null" description:"备注"`
	CATEGORYSTATE  int    `orm:"column(CATEGORY_STATE);null" description:"0表示停用1表示启用"`
	CATEGORYSORT   int    `orm:"column(CATEGORY_SORT)" description:"排序数"`
}

func (t *CourseCategory) TableName() string {
	return "course_category"
}

func init() {
	orm.RegisterModel(new(CourseCategory))
}

// AddCourseCategory insert a new CourseCategory into database and returns
// last inserted Id on success.
func AddCourseCategory(m *CourseCategory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCourseCategoryById retrieves CourseCategory by Id. Returns error if
// Id doesn't exist
func GetCourseCategoryById(id string) (v *CourseCategory, err error) {
	o := orm.NewOrm()
	v = &CourseCategory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCourseCategory retrieves all CourseCategory matches certain condition. Returns empty list if
// no records exist
func GetAllCourseCategory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CourseCategory))
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

	var l []CourseCategory
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

// UpdateCourseCategory updates CourseCategory by Id and returns error if
// the record to be updated doesn't exist
func UpdateCourseCategoryById(m *CourseCategory) (err error) {
	o := orm.NewOrm()
	v := CourseCategory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCourseCategory deletes CourseCategory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCourseCategory(id string) (err error) {
	o := orm.NewOrm()
	v := CourseCategory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CourseCategory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
