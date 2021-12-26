package constants

type BillType uint

const (
	BT_INCOME   BillType = 0 // 收入
	BT_PAY      BillType = 1 // 支出
	BT_TRANSFER BillType = 2 // 内部转账
)

type BillAccountType uint

const (
	BAT_CASH  BillAccountType = 0 // 现金
	BAT_CARD  BillAccountType = 1 // 银行卡
	BAT_IDEAL BillAccountType = 2 // 虚拟货币
)
