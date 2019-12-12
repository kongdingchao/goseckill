package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/*
drop table `Seckill`;
CREATE TABLE `seckill` (
  `SeckillId` int(10) NOT NULL AUTO_INCREMENT,
  `ComName` varchar(64) NOT NULL,
  `ComNumber` int(10) unsigned DEFAULT '0',
  `TimeStart` datetime DEFAULT NULL,
  `TimeEnd` datetime DEFAULT NULL,
  `TimeCreate` datetime DEFAULT NULL,
  PRIMARY KEY (`SeckillId`)
);
*/
//秒杀商品信息
type Seckill struct {
	SeckillId  int       `orm:"column(SeckillId);auto"` //秒杀序号
	ComName    string    `orm:"column(Name);size(64);default()"`
	ComNumber  int       `orm:"column(ComNumber);default(0)"`
	TimeStart  time.Time `orm:"column(TimeStart);default(0)"`  //秒杀开始时间
	TimeEnd    time.Time `orm:"column(TimeEnd);default(0)"`    //秒杀截止时间
	TimeCreate time.Time `orm:"column(TimeCreate);default(0)"` //秒杀添加时间
}

func NewSeckill() *Seckill {
	return &Seckill{}
}

func GetSeckill() Seckill {

	o := orm.NewOrm()
	seckill := Seckill{ComName: "app", ComNumber: 1,
		TimeStart:  time.Date(2020, time.Month(2), 23, 1, 30, 30, 0, time.Local),
		TimeEnd:    time.Date(2020, time.Month(2), 23, 10, 30, 30, 0, time.Local),
		TimeCreate: time.Now(),
	}

	// insert
	_, err := o.Insert(&seckill)
	checkErr(err)

	// update
	seckill.ComNumber = 2
	_, err = o.Update(&seckill)
	checkErr(err)

	// read one
	u := Seckill{SeckillId: seckill.SeckillId}
	err = o.Read(&u)
	checkErr(err)

	// delete
	_, err = o.Delete(&u)
	checkErr(err)

	return seckill
}

func InsertAndGetSeckill() Seckill {
	var insertItem = Seckill{}
	insertItem.ComName = "iphone 8s"
	insertItem.ComNumber = 0
	insertItem.TimeStart = time.Date(2020, time.Month(2), 23, 1, 30, 30, 0, time.Local)
	insertItem.TimeEnd = insertItem.TimeStart.AddDate(0, 0, 1)
	insertItem.TimeCreate = time.Now()

	db, err := sql.Open("mysql", "root:123456@/seckill?charset=utf8")
	checkErr(err)
	tx, err := db.Begin()
	checkErr(err)
	defer func() {
		if err != nil {
			tx.Rollback()
			fmt.Println("tx.Rollback()")
		} else {
			tx.Commit()
			fmt.Println("tx.Commit()")
		}
		db.Close()
	}()

	//insert
	stmt, err := db.Prepare("insert into seckill set ComName=?, ComNumber=?, TimeStart=?, TimeEnd=?,TimeCreate=?")
	checkErr(err)
	result, err := stmt.Exec(insertItem.ComName, insertItem.ComNumber, insertItem.TimeStart, insertItem.TimeEnd, insertItem.TimeCreate)
	checkErr(err)
	lastInsertId, err := result.LastInsertId()
	checkErr(err)
	fmt.Printf("insert ok!%v\n", lastInsertId)

	//update
	stmt, err = db.Prepare("update seckill set ComNumber=? where seckillId=?")
	checkErr(err)
	result, err = stmt.Exec(1, lastInsertId)
	checkErr(err)

	rowsAffected, err := result.RowsAffected()
	checkErr(err)
	fmt.Printf("update ok!%v\n", rowsAffected)

	//select
	row := db.QueryRow("SELECT SeckillId, ComName, ComNumber, TimeStart, TimeEnd,TimeCreate FROM seckill where SeckillId=?", lastInsertId)
	checkErr(err)

	var selectResult = Seckill{}
	var timeStart string
	var timeEnd string
	var timeCreate string
	row.Scan(&selectResult.SeckillId, &selectResult.ComName, &selectResult.ComNumber, &timeStart,
		&timeEnd, &timeCreate)
	selectResult.TimeStart, _ = time.Parse("2019-01-01 01:02:33.001", timeStart)
	selectResult.TimeEnd, _ = time.Parse("2019-01-01 01:02:33.001", timeEnd)
	selectResult.TimeCreate, _ = time.Parse("2019-01-01 01:02:33.001", timeCreate)
	fmt.Printf("select ok!%v\n", selectResult)

	//delete
	/*stmt, err = db.Prepare("delete from seckill where seckillId=?")
	checkErr(err)
	result, err = stmt.Exec(1, lastInsertId)
	checkErr(err)

	rowsAffected, err = result.RowsAffected()
	checkErr(err)
	fmt.Printf("delete ok!%v\n", rowsAffected)*/

	return selectResult
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
