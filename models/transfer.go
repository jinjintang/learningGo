package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Transfer struct {
	Id                 int       `orm:"column(TRANSFER_ID);pk"`
	TRANSFERUSERID     string    `orm:"column(TRANSFER_USERID);size(36);null" description:"转账用户id"`
	TRANSFERUSERNAME   string    `orm:"column(TRANSFER_USERNAME);size(20);null" description:"转账用户姓名"`
	TRANSFERUSEROPENID string    `orm:"column(TRANSFER_USEROPENID);size(50);null" description:"转账用户OPENID"`
	TRANSFERMONEY      float64   `orm:"column(TRANSFER_MONEY);null" description:"交易金额"`
	TRANSFERCOMMISSION float64   `orm:"column(TRANSFER_COMMISSION);null" description:"交易手续费"`
	TRANSFERCASH       float64   `orm:"column(TRANSFER_CASH);null" description:"实际到账金额"`
	TRANSFERTIME       time.Time `orm:"column(TRANSFER_TIME);type(timestamp);auto_now_add" description:"申请时间"`
	AREAID             string    `orm:"column(AREA_ID);size(36);null" description:"区域id"`
	AREANAME           string    `orm:"column(AREA_NAME);size(20);null" description:"区域名称"`
}

func (t *Transfer) TableName() string {
	return "transfer"
}

func init() {
	orm.RegisterModel(new(Transfer))
}

// AddTransfer insert a new Transfer into database and returns
// last inserted Id on success.
func AddTransfer(m *Transfer) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTransferById retrieves Transfer by Id. Returns error if
// Id doesn't exist
func GetTransferById(id int) (v *Transfer, err error) {
	o := orm.NewOrm()
	v = &Transfer{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTransfer retrieves all Transfer matches certain condition. Returns empty list if
// no records exist
func GetAllTransfer(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Transfer))
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

	var l []Transfer
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

// UpdateTransfer updates Transfer by Id and returns error if
// the record to be updated doesn't exist
func UpdateTransferById(m *Transfer) (err error) {
	o := orm.NewOrm()
	v := Transfer{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTransfer deletes Transfer by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTransfer(id int) (err error) {
	o := orm.NewOrm()
	v := Transfer{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Transfer{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
