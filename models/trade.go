package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Trade struct {
	Id            int       `orm:"column(TRADE_ID);pk"`
	CATEGORY      int       `orm:"column(CATEGORY);null" description:"0表示购买课程1充值2购买商品"`
	TOTALFEE      float64   `orm:"column(TOTAL_FEE);null" description:"最初总金额"`
	PAYUSERID     string    `orm:"column(PAY_USER_ID);size(36);null" description:"付款用户id"`
	PAYUSERNAME   string    `orm:"column(PAY_USER_NAME);size(50);null" description:"付款用户姓名"`
	PAYUSEROPENID string    `orm:"column(PAY_USER_OPENID);size(100);null" description:"付款用户openid"`
	PAYTIME       time.Time `orm:"column(PAY_TIME);type(timestamp);auto_now_add" description:"付款时间"`
	STATE         int       `orm:"column(STATE);null" description:"0已支付1申请退款2驳回3部分退款4已退款5已完成不可更改6子订单已开班"`
	BANKTYPE      string    `orm:"column(BANK_TYPE);size(10);null" description:"银行缩写"`
	CASHFEE       float64   `orm:"column(CASH_FEE);null" description:"现金支付金额，除去红包什么的"`
	DEVICEINFO    string    `orm:"column(DEVICE_INFO);size(10);null" description:"支付设备信息，手机，网页等"`
	FEETYPE       string    `orm:"column(FEE_TYPE);size(10);null" description:"支付币种"`
	OUTTRADENO    string    `orm:"column(OUT_TRADE_NO);size(35);null" description:"内部交易订单号，<=32"`
	TIMEEND       string    `orm:"column(TIME_END);size(20);null" description:"支付订单完成时间"`
	TRADETYPE     string    `orm:"column(TRADE_TYPE);size(10);null" description:"支付方式，公众号，银行卡支付"`
	TRANSACTIONID string    `orm:"column(TRANSACTION_ID);size(35);null" description:"微信交易订单号"`
	AREAID        string    `orm:"column(AREA_ID);size(36);null" description:"区域id"`
	AREANAME      string    `orm:"column(AREA_NAME);size(20);null" description:"区域名称"`
	COUNTMONEY    float64   `orm:"column(COUNT_MONEY);null" description:"使用账户付款金额"`
	WEIXINMONEY   float64   `orm:"column(WEIXINMONEY);null" description:"微信付款金额"`
	FMONEY        float64   `orm:"column(F_MONEY);null" description:"最终付款金额"`
	XIANJINPAY    float64   `orm:"column(XIANJINPAY);null" description:"现金流支付"`
	NAME          string    `orm:"column(NAME);size(100);null" description:"课程名称"`
}

func (t *Trade) TableName() string {
	return "trade"
}

func init() {
	orm.RegisterModel(new(Trade))
}

// AddTrade insert a new Trade into database and returns
// last inserted Id on success.
func AddTrade(m *Trade) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTradeById retrieves Trade by Id. Returns error if
// Id doesn't exist
func GetTradeById(id int) (v *Trade, err error) {
	o := orm.NewOrm()
	v = &Trade{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTrade retrieves all Trade matches certain condition. Returns empty list if
// no records exist
func GetAllTrade(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Trade))
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

	var l []Trade
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

// UpdateTrade updates Trade by Id and returns error if
// the record to be updated doesn't exist
func UpdateTradeById(m *Trade) (err error) {
	o := orm.NewOrm()
	v := Trade{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTrade deletes Trade by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTrade(id int) (err error) {
	o := orm.NewOrm()
	v := Trade{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Trade{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
