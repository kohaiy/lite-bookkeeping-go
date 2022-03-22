package constants

import "time"

var UserOauthExpiredAt = time.Minute * 5

type BillType uint

const (
	BT_INCOME   BillType = 1 // 收入
	BT_PAY      BillType = 2 // 支出
	BT_TRANSFER BillType = 3 // 内部转账
)

type BillAccountType uint

const (
	BAT_CASH  BillAccountType = 1 // 现金
	BAT_CARD  BillAccountType = 2 // 银行卡
	BAT_IDEAL BillAccountType = 3 // 虚拟货币
)

type OauthType uint

const (
	OT_UNIAUTH OauthType = 1
)
