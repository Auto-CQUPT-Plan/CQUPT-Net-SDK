package CQUPT_Net_SDK

import (
	"github.com/Auto-CQUPT-Plan/CQUPT-Net-SDK/core/models"
)

// NetLogin 登录校园网
func (r *SDK) NetLogin(data *BasicData) (resp *AuthResult, err error) {
	req := &models.LoginData{
		StuID:    data.StuID,
		Password: data.Password,
		UA:       data.UA,
		Isp:      data.Isp,
		IPAddr:   data.IPAddr,
		MACAddr:  data.MACAddr,
	}

	res, err := r.spyder.NetLogin(req)
	if err != nil {
		return nil, err
	}

	resp = &AuthResult{
		Result:  res.Result,
		Msg:     res.Msg,
		RetCode: res.RetCode,
	}

	return resp, nil
}

// NetLogout 校园网注销
func (r *SDK) NetLogout(data *BasicData) (resp *UnbindResult, err error) {
	req := &models.LoginData{
		StuID:    data.StuID,
		Password: data.Password,
		UA:       data.UA,
		Isp:      data.Isp,
		IPAddr:   data.IPAddr,
		MACAddr:  data.MACAddr,
	}

	res, err := r.spyder.NetLogout(req)
	if err != nil {
		return nil, err
	}

	resp = &UnbindResult{
		Result: res.Result,
		Msg:    res.Msg,
	}

	return resp, nil
}

// NetChecker 登录状态检测
func (r *SDK) NetChecker(data *BasicData) (resp *CheckerResult, err error) {
	req := &models.LoginData{
		StuID:    data.StuID,
		Password: data.Password,
		UA:       data.UA,
		Isp:      data.Isp,
		IPAddr:   data.IPAddr,
		MACAddr:  data.MACAddr,
	}

	res, err := r.spyder.NetChecker(req)
	if err != nil {
		return nil, err
	}

	resp = &CheckerResult{
		Result: res.Result,
		Time:   res.Time,
		Msg:    res.Msg,
	}

	return resp, nil
}

// GetNetServiceInfo 获取上网信息
func (r *SDK) GetNetServiceInfo(data *BasicData) (resp *NetInfoS, err error) {
	req := &models.LoginData{
		StuID:    data.StuID,
		Password: data.Password,
		UA:       data.UA,
		Isp:      data.Isp,
		IPAddr:   data.IPAddr,
		MACAddr:  data.MACAddr,
	}

	res, err := r.spyder.GetNetServiceInfo(req)
	if err != nil {
		return nil, err
	}

	resp = &NetInfoS{
		UserName:      res.UserName,
		UsedTime:      res.UsedTime,
		UsedFlow:      res.UsedFlow,
		Balance:       res.Balance,
		Account:       res.Account,
		Status:        res.Status,
		Package:       res.Package,
		PackageDetail: res.PackageDetail,
		BillingType:   res.BillingType,
		BillingCycle:  res.BillingCycle,
	}

	return resp, nil
}

// GetLoginHistory 获取上网记录
func (r *SDK) GetLoginHistory(data *BasicData) (resp []*LoginHistoryRecordS, err error) {
	req := &models.LoginData{
		StuID:    data.StuID,
		Password: data.Password,
		UA:       data.UA,
		Isp:      data.Isp,
		IPAddr:   data.IPAddr,
		MACAddr:  data.MACAddr,
	}

	res, err := r.spyder.GetLoginHistory(req)
	if err != nil {
		return nil, err
	}

	resp = make([]*LoginHistoryRecordS, 0, len(res))
	for _, item := range res {
		resp = append(
			resp,
			&LoginHistoryRecordS{
				LoginTime:    item.LoginTime,
				LogoutTime:   item.LogoutTime,
				IP:           item.IP,
				MAC:          item.MAC,
				UseTime:      item.UseTime,
				UseFlow:      item.UseFlow,
				BillingType:  item.BillingType,
				ChargeAmount: item.ChargeAmount,
				HostName:     item.HostName,
				TerminalType: item.TerminalType,
			},
		)
	}

	return resp, nil
}

func (r *SDK) GetOnlineList(data *BasicData) (resp []*OnlineRecord, err error) {
	req := &models.LoginData{
		StuID:    data.StuID,
		Password: data.Password,
		UA:       data.UA,
		Isp:      data.Isp,
		IPAddr:   data.IPAddr,
		MACAddr:  data.MACAddr,
	}

	res, err := r.spyder.GetOnlineList(req)
	if err != nil {
		return nil, err
	}

	resp = make([]*OnlineRecord, 0, len(res))
	for _, item := range res {
		resp = append(
			resp,
			&OnlineRecord{
				Brasid:       item.Brasid,
				DownFlow:     item.UpFlow,
				HostName:     item.HostName,
				IP:           item.IP,
				LoginTime:    item.LoginTime,
				MAC:          item.MAC,
				SessionId:    item.SessionId,
				TerminalType: item.TerminalType,
				UpFlow:       item.UpFlow,
				UseTime:      item.UseTime,
				UserId:       item.UserId,
			},
		)
	}

	return resp, nil
}

func (r *SDK) GetDefaultMAC() string {
	return "000000000000"
}
