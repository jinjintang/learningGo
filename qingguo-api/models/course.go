package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Course struct {
	Id             int       `orm:"column(COURSE_ID);auto"`
	COURSENAME     string    `orm:"column(COURSE_NAME);size(50);null" description:"课程名"`
	COURSECOVER    string    `orm:"column(COURSE_COVER);size(200);null" description:"课程首页展示图片路径"`
	COURSEABSTRACT string    `orm:"column(COURSE_ABSTRACT);size(100);null" description:"课程摘要"`
	COURSEPRICE    float64   `orm:"column(COURSE_PRICE);null" description:"课程单价"`
	COURSECONTENT  string    `orm:"column(COURSE_CONTENT);null" description:"课程内容"`
	COURSEREMARK   string    `orm:"column(COURSE_REMARK);size(500);null" description:"课程备注"`
	OPENNUMBER     int       `orm:"column(OPENNUMBER);null" description:"开课需要人数"`
	SORTNUM        int       `orm:"column(SORTNUM);null" description:"排序号"`
	STATE          int       `orm:"column(STATE);null" description:"0草稿1提交2分部已审核3总部已审核4停用"`
	RELEASEUSERID  string    `orm:"column(RELEASEUSERID);size(36);null" description:"发布用户ID"`
	RELEASENAME    string    `orm:"column(RELEASENAME);size(10);null" description:"发布用户姓名(冗余)"`
	RELEASEOPENID  string    `orm:"column(RELEASEOPENID);size(100);null" description:"发布用户OPENID(冗余)"`
	RELEASETIME    time.Time `orm:"column(RELEASETIME);type(timestamp);auto_now_add" description:"发布时间"`
	AREAID         string    `orm:"column(AREA_ID);size(36);null" description:"区域ID"`
	AREANAME       string    `orm:"column(AREA_NAME);size(20);null" description:"区域名称（冗余）"`
	CATEGORYID     string    `orm:"column(CATEGORY_ID);size(36);null" description:"类别主键"`
	CATEGORYNAME   string    `orm:"column(CATEGORY_NAME);size(20);null" description:"类别名称（冗余）"`
	RANKID         string    `orm:"column(RANK_ID);size(36);null" description:"等级主键"`
	RANKNAME       string    `orm:"column(RANK_NAME);size(20);null" description:"等级名称(冗余)"`
	FCHECKUSERID   string    `orm:"column(F_CHECK_USERID);size(36);null" description:"分部审核人id"`
	FCHECKUSERNAME string    `orm:"column(F_CHECK_USERNAME);size(50);null" description:"分部审核人姓名(冗余)"`
	FCHECKOPENID   string    `orm:"column(F_CHECK_OPENID);size(100);null" description:"分部审核人openid(冗余)"`
	ZCHECKUSERID   string    `orm:"column(Z_CHECK_USERID);size(36);null" description:"总部审核人id"`
	ZCHECKUSERNAME string    `orm:"column(Z_CHECK_USERNAME);size(50);null" description:"总部审核人姓名"`
	ZCHECKOPENID   string    `orm:"column(Z_CHECK_OPENID);size(100);null" description:"总部审核人openid"`
	FCHECKTIME     time.Time `orm:"column(F_CHECK_TIME);type(timestamp);null" description:"分部审核时间"`
	ZCHECKTIME     time.Time `orm:"column(Z_CHECK_TIME);type(timestamp);null" description:"总部审核时间"`
	RECOMMEND      int       `orm:"column(RECOMMEND);null" description:"是否是精华"`
}

func (t *Course) TableName() string {
	return "course"
}

func init() {
	orm.RegisterModel(new(Course))
}

// AddCourse insert a new Course into database and returns
// last inserted Id on success.
func AddCourse(m *Course) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCourseById retrieves Course by Id. Returns error if
// Id doesn't exist
func GetCourseById(id int) (v *Course, err error) {
	o := orm.NewOrm()
	v = &Course{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCourse retrieves all Course matches certain condition. Returns empty list if
// no records exist
func GetAllCourse(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Course))
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

	var l []Course
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

// UpdateCourse updates Course by Id and returns error if
// the record to be updated doesn't exist
func UpdateCourseById(m *Course) (err error) {
	o := orm.NewOrm()
	v := Course{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCourse deletes Course by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCourse(id int) (err error) {
	o := orm.NewOrm()
	v := Course{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Course{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
