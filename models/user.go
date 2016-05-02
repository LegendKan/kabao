package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	//"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"kabao/common"
)

type User struct {
	Id             int       `orm:"column(userid);auto"`
	Username       string    `orm:"column(username);size(32);null"`
	Password       string    `orm:"column(password);size(64)"`
	Salt           string    `orm:"column(salt);size(32)"`
	Phone          string    `orm:"column(phone);size(11);null"`
	Email          string    `orm:"column(email);size(64);null"`
	Sex            int8      `orm:"column(sex)"`
	Avatar         string    `orm:"column(avatar);size(255);null"`
	Realname       string    `orm:"column(realname);size(64);null"`
	Slogan         string    `orm:"column(slogan);size(255);null"`
	Type           int8      `orm:"column(type)"`
	Status         int8      `orm:"column(status)"`
	Createtime     time.Time `orm:"column(createtime);type(datetime)"`
	Lastupdatetime time.Time `orm:"column(lastupdatetime);type(timestamp);auto_now"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func AddNewUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	v := &User{Phone: m.Phone}
	if err = o.Read(v, "Phone"); err == nil {
		err = ErrExist
		return
	}
	if err = o.Begin(); err != nil {
		return
	}
	id, err = o.Insert(m)
	if err == nil {
		t := &Token{Userid: int(id), Token: common.RandSeq(64)}
		_, err := AddToken(t)
		if err != nil {
			o.Rollback()
			return id, err
		} else {
			o.Commit()
			return id, nil
		}
	}
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetUserByPhone(phone string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Phone: phone}
	if err = o.Read(v, "Phone"); err == nil {
		return v, nil
	} else if err == orm.ErrNoRows {
		//logs.("No User with Phone %s", phone)
		return nil, ErrNoRows
	}
	return nil, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
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

	var l []User
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

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
