package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Advertise struct {
	Id               int       `orm:"column(ADVERTISE_ID);auto"`
	ADCATEGORY       int       `orm:"column(ADCATEGORY);null" description:"0内部课程1外部2商品"`
	ADVERTISECOVER   string    `orm:"column(ADVERTISE_COVER);size(200);null" description:"广告图片"`
	ADVERTISEURL     string    `orm:"column(ADVERTISE_URL);size(200);null" description:"链接"`
	ISMAIN           int       `orm:"column(ISMAIN);null" description:"0不显示首页1显示首页"`
	ADVERTISECONTENT string    `orm:"column(ADVERTISE_CONTENT);size(200);null" description:"广告标题"`
	ADSTATE          int       `orm:"column(ADSTATE);null" description:"0草稿1提交2分部已审核3总部已审核4停用"`
	AREAID           string    `orm:"column(AREA_ID);size(36);null" description:"区域主键"`
	AREANAME         string    `orm:"column(AREA_NAME);size(20);null" description:"区域名称"`
	PRICE            float64   `orm:"column(PRICE);null" description:"单价"`
	RELEASETIME      time.Time `orm:"column(RELEASETIME);type(timestamp);auto_now_add" description:"发布时间"`
	RELEASEUSERID    string    `orm:"column(RELEASEUSERID);size(36);null" description:"发布用户id"`
	RELEASENAME      string    `orm:"column(RELEASENAME);size(30);null" description:"发布用户姓名"`
	RELEASEOPENID    string    `orm:"column(RELEASEOPENID);size(50);null" description:"发布用户OPENID"`
	FCHECKUSERID     string    `orm:"column(F_CHECK_USERID);size(36);null" description:"分部审核人id"`
	FCHECKUSERNAME   string    `orm:"column(F_CHECK_USERNAME);size(50);null" description:"分部审核人姓名(冗余)"`
	FCHECKOPENID     string    `orm:"column(F_CHECK_OPENID);size(50);null" description:"分部审核人openid(冗余)"`
	FCHECKTIME       time.Time `orm:"column(F_CHECK_TIME);type(timestamp);null" description:"分部审核时间"`
	ZCHECKUSERID     string    `orm:"column(Z_CHECK_USERID);size(36);null" description:"总部审核人id"`
	ZCHECKUSERNAME   string    `orm:"column(Z_CHECK_USERNAME);size(50);null" description:"总部审核人姓名"`
	ZCHECKOPENID     string    `orm:"column(Z_CHECK_OPENID);size(50);null" description:"总部审核人openid"`
	ZCHECKTIME       time.Time `orm:"column(Z_CHECK_TIME);type(timestamp);null" description:"总部审核时间"`
}

func (t *Advertise) TableName() string {
	return "advertise"
}

func init() {
	orm.RegisterModel(new(Advertise))
}

// AddAdvertise insert a new Advertise into database and returns
// last inserted Id on success.
func AddAdvertise(m *Advertise) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAdvertiseById retrieves Advertise by Id. Returns error if
// Id doesn't exist
func GetAdvertiseById(id int) (v *Advertise, err error) {
	o := orm.NewOrm()
	v = &Advertise{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAdvertise retrieves all Advertise matches certain condition. Returns empty list if
// no records exist
func GetAllAdvertise(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Advertise))
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

	var l []Advertise
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

// UpdateAdvertise updates Advertise by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdvertiseById(m *Advertise) (err error) {
	o := orm.NewOrm()
	v := Advertise{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAdvertise deletes Advertise by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAdvertise(id int) (err error) {
	o := orm.NewOrm()
	v := Advertise{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Advertise{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
