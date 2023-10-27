// Code generated by goctl. DO NOT EDIT.
package types

type NoticeRequest struct {
	JKey        string `json:"j_key" desc:"通知key"`
	Cycle       string `json:"cycle" desc:"执行周期"`
	Hour        int8   `json:"hour" desc:"小时"`
	Minute      int8   `json:"minute" desc:"分钟"`
	LineId      string `json:"line_id"`
	LineName    string `json:"line_name"`
	LineFromTo  string `json:"line_from_to"`
	StationNum  string `json:"station_num" desc:"线路序号"`
	StationId   string `json:"station_id"`
	StationName string `json:"station_name"`
	StartTime   string `json:"start_time" desc:"开始时间"`
	EndTime     string `json:"end_time" desc:"结束时间"`
	NoticeTime  int8   `json:"notice_time" desc:"提醒次数"`
}

type NoticeResponse struct {
}

type SearchRequest struct {
	Linename string `json:"linename" desc:"线路名称"`
}

type SearchResponse struct {
	List []LineList `json:"list" desc:"公交列表"`
}

type LineList struct {
	Linename   string `json:"linename" desc:"线路名称"`
	LineFromTo string `json:"line_from_to" desc:"线路始发站"`
	Lineid     string `json:"lineid" desc:"线路ID"`
}

type LineRequest struct {
	Lineid string `json:"lineid" desc:"线路ID"`
}

type LineResponse struct {
	Linename    string     `json:"linename" desc:"线路名称"`
	Lineto      string     `json:"lineto" desc:"目的地"`
	Lineid      string     `json:"lineid" desc:"线路ID"`
	Start       string     `json:"start" desc:"首班车时间"`
	End         string     `json:"end" desc:"末班车时间"`
	Directionid string     `json:"directionid" desc:"反向公交车编号"`
	Timetable   string     `json:"timetable" desc:"时刻表"`
	List        []LineInfo `json:"list" desc:"公交详情"`
}

type LineInfo struct {
	Stationnum    string   `json:"stationnum" desc:"线路序号"`
	Stationname   string   `json:"stationname" desc:"线路名称"`
	Stationid     string   `json:"stationid" desc:"站台编号"`
	Stationdetail []string `json:"stationdetail" desc:"站点进站公交列表"`
}

type HomeResp struct {
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type SendEMailRequest struct {
	Email string `json:"email"`
}

type UserLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SendEMailUserResponse struct {
	Code int `json:"code"`
}

type UserResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRequest struct {
	Name     string `json:"name,optional"`
	Page     int    `json:"page,optional"`
	PageSize int    `json:"page_size,optional"`
}

type UserListResponse struct {
	List  []User `json:"users"`
	Total int64  `json:"total"`
}

type User struct {
	Id          int    `json:"id"`
	Identity    string `json:"identity"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	NowVolume   int    `json:"now_volume"`
	TotalVolume int    `json:"total_volume"`
}

type UserDetailRequest struct {
	Identity string `json:"identity"`
}

type UserDetailResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GeeRequest struct {
	Uuid string `json:"uuid,optional"`
}

type GeeResponse struct {
	Challenge  string `json:"challenge"`
	Gt         string `json:"gt"`
	NewCaptcha bool   `json:"new_captcha"`
	Success    int    `json:"success"`
}
