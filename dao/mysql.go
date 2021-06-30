package dao

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gcfg.v1"
	"main/model"
)

var db *sql.DB

type Mysql struct {
	Username string
	Password string
	Url string
}

type Conf struct {
	Mysql Mysql
}

func ConnectMysql(){
	var conf = Conf{}
	err := gcfg.ReadFileInto(&conf, "database.conf")
	if err != nil{
		fmt.Println("Database Config File database.conf open failed!")
	}
	username := conf.Mysql.Username
	password := conf.Mysql.Password
	url := conf.Mysql.Url

	dsn := username + ":" + password + "@tcp(" + url + ")/iotsystem?charset=utf8"
	db, _ = sql.Open("mysql", dsn)
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil{
		fmt.Println("Database open failed!")
		return
	}
	fmt.Println("Database connect success!")
}

/*func FindDeviceById(clientId string) model.Device {
	var info model.Device
	sqlstr := "select * from device where clientId = ?"
	rows, err := db.Query(sqlstr, clientId)
	if rows.Next(){

	}
}*/

func CheckDeviceExist(clientId string)  bool{
	infoList := FindDeviceAll()
	for i:=0; i<len(infoList); i++{
		if infoList[i].ClientId == clientId{
			return true
		}
	}
	return false
}

func FindDeviceAll() []model.Device{
	var infoList []model.Device
	sqlstr := "select * from device"
	rows, err := db.Query(sqlstr)
	defer rows.Close()
	if err != nil{
		fmt.Println(err)
		return infoList
	}
	for rows.Next(){
		var info model.Device
		err := rows.Scan(&info.ClientId, &info.Name)
		if err != nil{
			fmt.Println(err)
			continue
		}
		infoList = append(infoList, info)
	}
	return infoList
}

func FindDeviceInfoByID(clientId string) []model.DeviceInfo{
	var infoList []model.DeviceInfo
	sqlstr := "select * from deviceInfo where clientId = ?"
	rows, err := db.Query(sqlstr, clientId)
	defer rows.Close()
	if err != nil{
		fmt.Println(err)
		return infoList
	}
	for rows.Next(){
		var info model.DeviceInfo
		err := rows.Scan(&info.ClientId, &info.Info, &info.Value, &info.Alert, &info.Lng, &info.Lat, &info.Timestamp)
		if err != nil{
			fmt.Println(err)
			continue
		}
		infoList = append(infoList, info)
	}
	return infoList
}

func FindDeviceInfoByIDAndCnt(info model.SearchDeviceInfo)  []model.DeviceInfo{
	var infoList []model.DeviceInfo
	sqlstr := "select * from deviceInfo where clientId = ? order by timestamp desc limit ?"
	rows, err := db.Query(sqlstr, info.ClientId, info.Cnt)
	defer rows.Close()
	if err != nil{
		fmt.Println(err)
		return infoList
	}
	for rows.Next(){
		var info model.DeviceInfo
		err := rows.Scan(&info.ClientId, &info.Info, &info.Value, &info.Alert, &info.Lng, &info.Lat, &info.Timestamp)
		if err != nil{
			fmt.Println(err)
			continue
		}
		infoList = append(infoList, info)
	}
	return infoList
}

func InsertDeviceInfo(info model.DeviceInfo)  {
	if CheckDeviceExist(info.ClientId) == false{
		sqlstr := "insert into device values(?,?)"
		_, err := db.Exec(sqlstr, info.ClientId, info.ClientId)
		if err != nil{
			fmt.Println(err)
			return
		}
	}
	sqlstr := "insert into deviceInfo values(?,?,?,?,?,?,?)"
	_, err := db.Exec(sqlstr, info.ClientId, info.Info, info.Value, info.Alert, info.Lng, info.Lat, info.Timestamp)
	if err != nil{
		fmt.Println(err)
		return
	}
	/*var infoList = FindDiveceInfoByID(info.ClientId)
	sort.Slice(infoList, func(i, j int) bool {
		return infoList[i].Timestamp > infoList[j].Timestamp
	})
	for i := 501; i<len(infoList); i++{
		sqlstr := "delete from deviceInfo where clientId = ? and timestamp = ?"
		_, err := db.Exec(sqlstr, info.ClientId, infoList[i].Timestamp)
		if err != nil{
			fmt.Println(err)
		}
	}*/
}

func FindUserByNameAndPassword(info model.LoginInfo)error{
	sqlstr := "select * from user where username = ? and password = ?"
	rows, err := db.Query(sqlstr, info.Username, info.Password)
	defer rows.Close()
	if err != nil{
		return err
	}else if rows.Next(){
		return nil
	}else {
		return errors.New("Username or password missed!")
	}
}

func InsertUser(info model.RegisterInfo) error {
	sqlstr := "insert into user values(?,?,?)"
	_, err := db.Exec(sqlstr, info.Username, info.Password, info.Email)
	return err
}

func ClearDeviceInfo() error{
	sqlstr := "set SQL_SAFE_UPDATES = 0"//取消数据库安全模式
	_, err := db.Exec(sqlstr)
	if err != nil{
		return err
	}
	sqlstr = "delete from deviceInfo"
	_, err = db.Exec(sqlstr)
	return err
}

func ModifyDeviceNameById(info model.ModifyDeviceNameInfo) error{
	sqlstr := "update device set name = ? where clientId = ?"
	_, err := db.Exec(sqlstr, info.Name, info.ClientId)
	return err
}

func GetOnlineDeviceNum() (int,error){
	sqlstr := "select count(*) from device"
	rows, err := db.Query(sqlstr)
	defer rows.Close()
	if err != nil{
		return 0, err
	}
	var num = 0
	if !rows.Next(){
		return 0, errors.New("select count from database failed!")
	}
	err = rows.Scan(&num)
	return num, err
}

func GetTotalInfo() (int,error){
	sqlstr := "select count(*) from deviceInfo"
	rows, err := db.Query(sqlstr)
	defer rows.Close()
	if err != nil{
		return 0, err
	}
	var num = 0
	if !rows.Next(){
		return 0, errors.New("select count from database failed!")
	}
	err = rows.Scan(&num)
	return num, err
}

func GetAlertInfo() (int,error){
	sqlstr := "select count(*) from deviceInfo where alert = 1"
	rows, err := db.Query(sqlstr)
	defer rows.Close()
	if err != nil{
		return 0, err
	}
	var num = 0
	if !rows.Next(){
		return 0, errors.New("select count from database failed!")
	}
	err = rows.Scan(&num)
	return num, err
}
