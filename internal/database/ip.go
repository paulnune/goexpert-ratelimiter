package database

type IpRequests struct {
	IP         string
	Qty        int
	BlockUntil int64
}

func NewRequest(ip string, qty int, block int64) IpRequests {
	return IpRequests{IP: ip, Qty: qty, BlockUntil: block}
}
