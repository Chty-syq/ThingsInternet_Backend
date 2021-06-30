package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/dao"
	"main/model"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func IndexPageHandler(context *gin.Context)  {
	context.HTML(http.StatusOK, "index.html", nil)
}
func LoginPageHandler(context *gin.Context)  {
	context.HTML(http.StatusOK, "login.html", nil)
}
func RegisterPageHandler(context *gin.Context)  {
	context.HTML(http.StatusOK, "register.html", nil)
}

func LoginHandler(context *gin.Context)  {
	var info model.LoginInfo
	context.BindJSON(&info)
	err := dao.FindUserByNameAndPassword(info)
	if err == nil{
		claims := &JWTClaims{
			Username: info.Username,
			Password: info.Password,
		}
		claims.IssuedAt = time.Now().Unix()
		claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
		token, err := getToken(claims)
		if err != nil{
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"token": token,
		})
	} else if strings.Contains(err.Error(), "missed"){
		context.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg": "username or password missed!",
		})
	} else{
		context.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg": err.Error(),
		})
	}
}

func RegisterHandler(context *gin.Context)  {
	var info model.RegisterInfo
	context.BindJSON(&info)
	var usernameLen = strings.Count(info.Username, "") - 1
	var passwordLen = strings.Count(info.Password, "") - 1
	fmt.Println(usernameLen)
	if(usernameLen < 6 || usernameLen > 18){
		context.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg": "Username length should be in 6~18!",
		})
		return
	}
	if(passwordLen < 6 || passwordLen > 18){
		context.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg": "Password length should be in 6~18!",
		})
		return
	}
	if(VerifyEmailFormat(info.Email) == false){
		context.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg": "Email format is not correct!",
		})
		return
	}
	err := dao.InsertUser(info)
	if err == nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success!",
		})
	}else if strings.Contains(err.Error(), "user.PRIMARY"){
		context.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg": "The username already exists!",
		})
	}else if strings.Contains(err.Error(), "user.email"){
		context.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg": "The email already exists!",
		})
	}else{
		context.JSON(http.StatusOK, err.Error())
	}
}

func ClearDeviceInfoHandler(context *gin.Context)  {
	var info model.TokenOnly
	err := context.BindJSON(&info)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 402,
			"msg": "Bad request form!",
		})
		return
	}
	err = verifyToken(info.Token)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 286,
			"msg": err.Error(),
		})
		return
	}
	err = dao.ClearDeviceInfo()
	if err != nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg": err.Error(),
		})
	}else{
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success!",
		})
	}
}

func ModifyDeviceNameHandler(context *gin.Context)  {
	var info model.ModifyDeviceNameInfo
	err := context.BindJSON(&info)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 402,
			"msg": "Bad request form!",
		})
		return
	}
	err = verifyToken(info.Token)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 286,
			"msg": err.Error(),
		})
		return
	}
	err = dao.ModifyDeviceNameById(info)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg": err.Error(),
		})
	}else{
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success!",
		})
	}
}

func GetOnlineDeviceNum(context *gin.Context)  {
	num, err := dao.GetOnlineDeviceNum()
	if err != nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 402,
			"msg": err.Error(),
		})
		return
	}else{
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": gin.H{
				"num": num,
			},
		})
		return
	}
}

func GetTotalInfo(context *gin.Context)  {
	num, err := dao.GetTotalInfo()
	if err != nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 402,
			"msg": err.Error(),
		})
		return
	}else{
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": gin.H{
				"num": num,
			},
		})
		return
	}
}

func GetAlertInfo(context *gin.Context)  {
	num, err := dao.GetAlertInfo()
	if err != nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 402,
			"msg": err.Error(),
		})
		return
	}else{
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": gin.H{
				"num": num,
			},
		})
		return
	}
}

func GetDeviceInfo(context *gin.Context)  {
	var dataList []model.OneDeviceInfo
	deviceList := dao.FindDeviceAll()
	for i:=0; i<len(deviceList); i++{
		infoList := dao.FindDeviceInfoByID(deviceList[i].ClientId)
		var alertNum = 0
		for j:=0; j<len(infoList); j++{
			alertNum += infoList[j].Alert
		}
		dataList = append(dataList, model.OneDeviceInfo{deviceList[i].Name, alertNum, len(infoList)})
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "success!",
		"data": dataList,
	})
}

func GetDeviceNameInfo(context *gin.Context)  {
	var info model.TokenOnly
	err := context.BindJSON(&info)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 402,
			"msg": "Bad request form!",
		})
		return
	}
	err = verifyToken(info.Token)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 286,
			"msg": err.Error(),
		})
		return
	}
	deviceList := dao.FindDeviceAll()
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "success!",
		"data": deviceList,
	})
}

func SearchDeviceInfo(context *gin.Context)  {
	var info model.SearchDeviceInfo
	err := context.BindJSON(&info)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 402,
			"msg": "Bad request form!",
		})
		return
	}
	err = verifyToken(info.Token)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{
			"code": 286,
			"msg": err.Error(),
		})
		return
	}
	infoList := dao.FindDeviceInfoByIDAndCnt(info)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "success!",
		"data": infoList,
	})
}
