package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Vipcard struct {
	Id              int       `orm:"column(cardid);auto"`
	Userid          int       `orm:"column(userid)"`
	Type            int8      `orm:"column(type)"`
	Cardtitle       string    `orm:"column(cardtitle);size(64);null"`
	Cardtype        string    `orm:"column(cardtype);size(255);null"`
	Cardno          string    `orm:"column(cardno);size(64);null"`
	Cardcode        string    `orm:"column(cardcode);size(64);null"`
	Cardname        string    `orm:"column(cardname);size(64);null"`
	Cardphone       string    `orm:"column(cardphone);size(16);null"`
	Carddescription string    `orm:"column(carddescription);size(255);null"`
	Cardbackgroud   string    `orm:"column(cardbackgroud);size(255);null"`
	Categoryid      int       `orm:"column(categoryid);null"`
	Usage           int8      `orm:"column(usage)"`
	Discount        int8      `orm:"column(discount);null"`
	Over            int       `orm:"column(over);null"`
	Bonus           int       `orm:"column(bonus);null"`
	Vipprice        int       `orm:"column(vipprice);null"`
	Starttime       time.Time `orm:"column(starttime);type(datetime);null"`
	Expiretime      time.Time `orm:"column(expiretime);type(datetime);null"`
	Shared          int8      `orm:"column(shared)"`
	Sharedname      string    `orm:"column(sharedname);size(64);null"`
	Sharedphone     string    `orm:"column(sharedphone);size(16);null"`
	Like            int       `orm:"column(like)"`
	Dislike         int       `orm:"column(dislike)"`
	Tags            string    `orm:"column(tags);size(255);null"`
	Createtime      time.Time `orm:"column(createtime);type(datetime)"`
	Lastupdatetime  time.Time `orm:"column(lastupdatetime);type(timestamp)"`
}

func init() {
	orm.RegisterModel(new(Vipcard))
}

// AddVipcard insert a new Vipcard into database and returns
// last inserted Id on success.
func AddVipcard(m *Vipcard) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetVipcardById retrieves Vipcard by Id. Returns error if
// Id doesn't exist
func GetVipcardById(id int) (v *Vipcard, err error) {
	o := orm.NewOrm()
	v = &Vipcard{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllVipcard retrieves all Vipcard matches certain condition. Returns empty list if
// no records exist
func GetAllVipcard(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Vipcard))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []Vipcard
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
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

// UpdateVipcard updates Vipcard by Id and returns error if
// the record to be updated doesn't exist
func UpdateVipcardById(m *Vipcard) (err error) {
	o := orm.NewOrm()
	v := Vipcard{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteVipcard deletes Vipcard by Id and returns error if
// the record to be deleted doesn't exist
func DeleteVipcard(id int) (err error) {
	o := orm.NewOrm()
	v := Vipcard{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Vipcard{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
