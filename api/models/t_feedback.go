package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TFeedback struct {
	Id              int       `orm:"column(id);auto"`
	UserId          string    `orm:"column(user_id);size(255);null"`
	Content         string    `orm:"column(content);size(1024);null"`
	Email           string    `orm:"column(email);size(128);null"`
	UserAgent       string    `orm:"column(user_agent);size(512);null"`
	Platform        string    `orm:"column(platform);size(64);null"`
	PlatformVersion string    `orm:"column(platform_version);size(64);null"`
	CarrierCode     string    `orm:"column(carrier_code);size(64);null"`
	CarrierName     string    `orm:"column(carrier_name);size(64);null"`
	DeviceName      string    `orm:"column(device_name);size(64);null"`
	DeviceId        string    `orm:"column(device_id);size(64);null"`
	SdkVersion      string    `orm:"column(sdk_version);size(64);null"`
	AppName         string    `orm:"column(app_name);size(64);null"`
	AppVersion      string    `orm:"column(app_version);size(64);null"`
	TimeZone        string    `orm:"column(time_zone);size(64);null"`
	Mobileb2version string    `orm:"column(mobileb2version);size(64);null"`
	SchoolName      string    `orm:"column(school_name);size(64);null"`
	SchoolPingUrl   string    `orm:"column(school_ping_url);size(64);null"`
	CreatedTime     time.Time `orm:"column(created_time);type(datetime);null"`
}

func (t *TFeedback) TableName() string {
	return "t_feedback"
}

func init() {
	orm.RegisterModel(new(TFeedback))
}

// AddTFeedback insert a new TFeedback into database and returns
// last inserted Id on success.
func AddTFeedback(m *TFeedback) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTFeedbackById retrieves TFeedback by Id. Returns error if
// Id doesn't exist
func GetTFeedbackById(id int) (v *TFeedback, err error) {
	o := orm.NewOrm()
	v = &TFeedback{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTFeedback retrieves all TFeedback matches certain condition. Returns empty list if
// no records exist
func GetAllTFeedback(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TFeedback))
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

	var l []TFeedback
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

// UpdateTFeedback updates TFeedback by Id and returns error if
// the record to be updated doesn't exist
func UpdateTFeedbackById(m *TFeedback) (err error) {
	o := orm.NewOrm()
	v := TFeedback{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTFeedback deletes TFeedback by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTFeedback(id int) (err error) {
	o := orm.NewOrm()
	v := TFeedback{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TFeedback{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
