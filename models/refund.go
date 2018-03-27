package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Refund struct {
	Id               int       `orm:"column(REFUND_ID);pk"`
	TRADEID          string    `orm:"column(TRADE_ID);size(36);null" description:"交易订单ID"`
	TRADEDETAILID    string    `orm:"column(TRADE_DETAILID);size(36);null" description:"交易详情ID"`
	REFUNDUSERID     string    `orm:"column(REFUND_USERID);size(36);null" description:"申请退款用户id"`
	REFUNDUSERNAME   string    `orm:"column(REFUND_USERNAME);size(20);null" description:"申请退款用户姓名"`
	REFUNDUSEROPENID string    `orm:"column(REFUND_USEROPENID);size(100);null" description:"申请退款用户OPENID"`
	MONEY            float64   `orm:"column(MONEY);null" description:"退款金额"`
	STATE            int       `orm:"column(STATE);null" description:"0待退款1已审核2驳回"`
	REFUNDREASON     string    `orm:"column(REFUND_REASON);size(200);null" description:"退款理由"`
	APPLYTIME        time.Time `orm:"column(APPLY_TIME);type(timestamp);auto_now_add" description:"申请时间"`
	DEALTIME         time.Time `orm:"column(DEAL_TIME);type(timestamp);null" description:"处理时间"`
	DEALUSERID       string    `orm:"column(DEAL_USERID);size(36);null" description:"处理人ID"`
	DEALUSERNAME     string    `orm:"column(DEAL_USERNAME);size(20);null" description:"处理人姓名"`
	DEALUSEROPENID   string    `orm:"column(DEAL_USEROPENID);size(100);null" description:"处理人OPENID"`
	AREAID           string    `orm:"column(AREA_ID);size(36);null" description:"区域id"`
	AREANAME         string    `orm:"column(AREA_NAME);size(20);null" description:"区域名称"`
}

func (t *Refund) TableName() string {
	return "refund"
}

func init() {
	orm.RegisterModel(new(Refund))
}

// AddRefund insert a new Refund into database and returns
// last inserted Id on success.
func AddRefund(m *Refund) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRefundById retrieves Refund by Id. Returns error if
// Id doesn't exist
func GetRefundById(id int) (v *Refund, err error) {
	o := orm.NewOrm()
	v = &Refund{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRefund retrieves all Refund matches certain condition. Returns empty list if
// no records exist
func GetAllRefund(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Refund))
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

	var l []Refund
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

// UpdateRefund updates Refund by Id and returns error if
// the record to be updated doesn't exist
func UpdateRefundById(m *Refund) (err error) {
	o := orm.NewOrm()
	v := Refund{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRefund deletes Refund by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRefund(id int) (err error) {
	o := orm.NewOrm()
	v := Refund{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Refund{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
