// Code generated by "stringer -type ErrCode -linecomment"; DO NOT EDIT.

package global

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Success-200]
	_ = x[StartServerErr-101000]
	_ = x[SystemErr-101001]
	_ = x[ParamsError-101002]
	_ = x[ConnectMysqlErr-101003]
	_ = x[RequestOvertimeErr-101004]
	_ = x[SignErr-101005]
}

const (
	_ErrCode_name_0 = "Success"
	_ErrCode_name_1 = "启动服务异常系统异常参数异常，请检查连接数据库异常请求发起时间超时参数签名异常"
)

var (
	_ErrCode_index_1 = [...]uint8{0, 18, 30, 54, 75, 99, 117}
)

func (i ErrCode) String() string {
	switch {
	case i == 200:
		return _ErrCode_name_0
	case 101000 <= i && i <= 101005:
		i -= 101000
		return _ErrCode_name_1[_ErrCode_index_1[i]:_ErrCode_index_1[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}