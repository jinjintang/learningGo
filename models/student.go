package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	 "strconv"  

	"github.com/astaxie/beego/orm"
)

type Student struct {
	Id              string    `orm:"column(USER_ID);pk"`
	USERNAME        string  `orm:"column(USER_NAME);size(50);null" description:"姓名"`
	PHONENUMBER     string  `orm:"column(PHONE_NUMBER);size(11);null" description:"电话"`
	EMAIL           string  `orm:"column(EMAIL);size(100);null" description:"邮箱"`
	COMPANY         string  `orm:"column(COMPANY);size(100);null" description:"公司"`
	SEX             int     `orm:"column(SEX);null" description:"0男1女"`
	ADRESS          string  `orm:"column(ADRESS);size(100);null" description:"地址"`
	WEIXIN          string  `orm:"column(WEI_XIN);size(30);null" description:"微信"`
	MARRYD          int     `orm:"column(MARRYD);null" description:"0未婚1已婚"`
	PHOTOPATH       string  `orm:"column(PHOTO_PATH);size(50);null" description:"相片路径"`
	USEROPENID      string  `orm:"column(USER_OPENID);size(50);null" description:"openid"`
	AREAID          string  `orm:"column(AREA_ID);size(36);null" description:"所属区域id"`
	AREANAME        string  `orm:"column(AREA_NAME);size(20);null" description:"区域名称"`
	FROZENASSETS    float64 `orm:"column(FROZEN_ASSETS);null" description:"冻结资产"`
	AVAILABLEASSETS float64 `orm:"column(AVAILABLE_ASSETS);null" description:"可用资产"`
	STATE           int     `orm:"column(STATE);null" description:"状态0老师待激活1学员2分部老师3激活"`
	SHENFEN         string  `orm:"column(SHENFEN);size(18);null" description:"身份证"`
	XIANJIN         float64 `orm:"column(XIANJIN);null" description:"现金"`
	LIKES           string  `orm:"column(LIKES);size(500);null" description:"爱好"`
	ZISHU           string  `orm:"column(ZISHU);null" description:"申请者自述"`
}

func (t *Student) TableName() string {
	return "student"
}

func init() {
	orm.RegisterModel(new(Student))
}

// AddStudent insert a new Student into database and returns
// last inserted Id on success.
func AddStudent(m *Student) (id string, err error) {
	
	o := orm.NewOrm()
	
	sid, serr := o.Insert(m)
	id=strconv.FormatInt(sid,10)
    err=serr
	return
}

// GetStudentById retrieves Student by Id. Returns error if
// Id doesn't exist
func GetStudentById(id string) (v *Student, err error) {
	o := orm.NewOrm()
	v = &Student{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllStudent retrieves all Student matches certain condition. Returns empty list if
// no records exist
func GetAllStudent(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Student))
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

	var l []Student
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

// UpdateStudent updates Student by Id and returns error if
// the record to be updated doesn't exist
func UpdateStudentById(m *Student) (err error) {
	o := orm.NewOrm()
	v := Student{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteStudent deletes Student by Id and returns error if
// the record to be deleted doesn't exist
func DeleteStudent(id string) (err error) {
	o := orm.NewOrm()
	v := Student{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Student{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
