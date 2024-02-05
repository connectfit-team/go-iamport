package types

type SchedulePayemntRequest struct {
	CustomerUid    string                  `json:"customer_uid,omitempty"`
	CheckingAmount int32                   `json:"checking_amount,omitempty"`
	CardNumber     string                  `json:"card_number,omitempty"`
	Expiry         string                  `json:"expiry,omitempty"`
	Birth          string                  `json:"birth,omitempty"`
	Pwd_2Digit     string                  `json:"pwd_2digit,omitempty"`
	Pg             string                  `json:"pg,omitempty"`
	Schedules      []*PaymentScheduleParam `json:"schedules,omitempty"`
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
	NoticeUrl     string `json:"notice_url,omitempty"`
}
