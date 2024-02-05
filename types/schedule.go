package types

type SchedulePayemntRequest struct {
	CustomerUid    string                  `protobuf:"bytes,1,opt,name=customer_uid,json=customerUid,proto3" json:"customer_uid,omitempty"`
	CheckingAmount int32                   `protobuf:"varint,2,opt,name=checking_amount,json=checkingAmount,proto3" json:"checking_amount,omitempty"`
	CardNumber     string                  `protobuf:"bytes,3,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	Expiry         string                  `protobuf:"bytes,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Birth          string                  `protobuf:"bytes,5,opt,name=birth,proto3" json:"birth,omitempty"`
	Pwd_2Digit     string                  `protobuf:"bytes,6,opt,name=pwd_2digit,json=pwd2digit,proto3" json:"pwd_2digit,omitempty"`
	Pg             string                  `protobuf:"bytes,7,opt,name=pg,proto3" json:"pg,omitempty"`
	Schedules      []*PaymentScheduleParam `protobuf:"bytes,8,rep,name=schedules,proto3" json:"schedules,omitempty"`
}

type PaymentScheduleParam struct {
	MerchantUid   string `json:"merchant_uid,omitempty"`
	ScheduleAt    int32  `json:"schedule_at,omitempty"`
	Amount        int32  `json:"amount,omitempty"`
	TaxFree       int32  `json:"tax_free,omitempty"`
	Name          string `json:"name,omitempty"`
	BuyerName     string `json:"buyer_name,omitempty"`
	BuyerEmail    string `json:"buyer_email,omitempty"`
	BuyerTel      string `json:"buyer_tel,omitempty"`
	BuyerAddr     string `json:"buyer_addr,omitempty"`
	BuyerPostcode string `json:"buyer_postcode,omitempty"`
}
