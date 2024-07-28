package cli

const (
	help              = "help"
	changePackagetype = "change"
	addPackage        = "add-package"
	acceptOrder       = "accept"
	setWorkersNum     = "set-workers"
	refundDeliver     = "refund_d"
	issueUser         = "issue"
	listOrders        = "list_o"
	refundUser        = "refund_u"
	listRefund        = "list_r"
	exit              = "exit"
)

type command struct {
	name        string
	description string
}
