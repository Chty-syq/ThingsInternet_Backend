package model

type LoginInfo struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type RegisterInfo struct{
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Email		string	`json:"email"`
}

type Device struct {
	ClientId	string 	`json:"clientId"`
	Name		string 	`json:"name"`
}

type DeviceInfo struct {
	ClientId	string 	`json:"clientId"`
	Info 		string 	`json:"info"`
	Value 		int 	`json:"value"`
	Alert 		int 	`json:"alert"`
	Lng 		float64 `json:"lng"`
	Lat 		float64 `json:"lat"`
	Timestamp 	int64 	`json:"timestamp"`
}

type ModifyDeviceNameInfo struct {
	Token		string	`json:"token"`
	ClientId	string 	`json:"clientId"`
	Name 		string 	`json:"name"`
}

type OneDeviceInfo struct {
	ClientId	string 	`json:"clientId"`
	AlertNum	int		`json:"alertNum"`
	TotalNum	int		`json:"totalNum"`
}

type SearchDeviceInfo struct {
	ClientId	string 	`json:"clientId"`
	Token		string	`json:"token"`
	Cnt			string 	`json:"cnt"`
}

type TokenOnly struct {
	Token		string	`json:"token"`
}
