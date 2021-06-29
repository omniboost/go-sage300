package accountviewnet

import "time"

type Tickets []Ticket

type Ticket struct {
	AccountClientNr    string             `json:"accountClientNr"`
	AccountName        string             `json:"accountName"`
	TotalPrice         float64            `json:"totalPrice"`
	TotalQuantity      float64            `json:"totalQuantity"`
	TotalToPay         float64            `json:"totalToPay"`
	TicketNr           int                `json:"ticketNr"`
	PcNr               int                `json:"pcNr"`
	Name               string             `json:"name"`
	CreateDT           time.Time          `json:"createDT"`
	BillDT             time.Time          `json:"billDT"`
	Prints             int                `json:"prints"`
	Seal               string             `json:"seal"`
	ReservationNr      string             `json:"reservationNr"`
	Info               string             `json:"info"`
	Footer             string             `json:"footer"`
	CheckID            int                `json:"checkId"`
	Orders             Orders             `json:"orders"`
	CreatedBy          User               `json:"createdBy"`
	CreatedByName      string             `json:"createdByName"`
	PC                 PC                 `json:"pc"`
	PcName             string             `json:"pcName"`
	Center             Center             `json:"center"`
	CenterName         string             `json:"centerName"`
	User               User               `json:"user"`
	CompressedLines    CompressedLines    `json:"compressedLines"`
	CompressedPayments CompressedPayments `json:"compressedPayments"`
	ShiftNr            int                `json:"shiftNr"`
	Key                string             `json:"key"`
	InReview           bool               `json:"inReview,omitempty"`
	ZNumber            int                `json:"zNumber,omitempty"`
	TableKey           string             `json:"tableKey,omitempty"`
	TableNr            int                `json:"tableNr,omitempty"`
}

type Orders []Order

type Order struct {
	TicketKey   string     `json:"ticketKey"`
	TotalPrice  float64    `json:"totalPrice"`
	ActionValue string     `json:"actionValue"`
	User        User       `json:"user"`
	UserName    string     `json:"userName"`
	CreateDT    time.Time  `json:"createDT"`
	Lines       OrderLines `json:"lines"`
	Payments    Payments   `json:"payments"`
	Seal        string     `json:"seal"`
	StartDT     time.Time  `json:"startDT"`
	Pc          PC         `json:"pc"`
	PcKey       string     `json:"pcKey"`
	PcName      string     `json:"pcName"`
	Center      Center     `json:"center"`
	CenterName  string     `json:"centerName"`
	Key         string     `json:"key"`
}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	NickName  string `json:"nickName"`
	FullName  string `json:"fullName"`
	Key       string `json:"key"`
}

type OrderLines []OrderLine

type OrderLine struct {
	ProductKey       string    `json:"productKey"`
	CourseKey        string    `json:"courseKey"`
	TotalPriceIncVat float64   `json:"totalPriceIncVat"`
	TotalPriceExVat  float64   `json:"totalPriceExVat"`
	TotalVat         float64   `json:"totalVat"`
	VatPercentage    float64   `json:"vatPercentage"`
	VatCode          int       `json:"vatCode"`
	PageKey          string    `json:"pageKey"`
	MenuID           string    `json:"menuId"`
	ProductName      string    `json:"productName"`
	Options          string    `json:"options"`
	Memo             string    `json:"memo"`
	Quantity         float64   `json:"quantity"`
	Units            int       `json:"units"`
	Price            float64   `json:"price"`
	CreateDT         time.Time `json:"createDT"`
	Key              string    `json:"key"`
}

type Payments []Payment

type Payment struct {
	Name             string    `json:"name"`
	TotalPrice       float64   `json:"totalPrice"`
	TransactionID    string    `json:"transactionId"`
	TransactionError string    `json:"transactionError"`
	Quantity         float64   `json:"quantity"`
	Units            int       `json:"units"`
	Price            float64   `json:"price"`
	CreateDT         time.Time `json:"createDT"`
	Key              string    `json:"key"`
}

type Center struct {
	Name       string `json:"name"`
	CenterNr   string `json:"centerNr"`
	GroupLevel int    `json:"groupLevel"`
	LeftNr     int    `json:"leftNr"`
	RightNr    int    `json:"rightNr"`
	Key        string `json:"key"`
}

type PC struct {
	PcNr       int    `json:"pcNr"`
	Name       string `json:"name"`
	PcGroupKey string `json:"pcGroupKey"`
	Key        string `json:"key"`
}

type CompressedLines []CompressedLine

type CompressedLine struct {
	AddOns      []interface{} `json:"addOns"`
	MenuParts   []interface{} `json:"menuParts"`
	ProductName string        `json:"productName"`
	ProductKey  string        `json:"productKey"`
	Quantity    float64       `json:"quantity"`
	Course      struct {
		CourseNr  int    `json:"courseNr"`
		Name      string `json:"name"`
		ShortName string `json:"shortName"`
		Key       string `json:"key"`
	} `json:"course"`
	Units          int     `json:"units"`
	UnitPrice      float64 `json:"unitPrice"`
	Memo           string  `json:"memo"`
	TotalPrice     float64 `json:"totalPrice"`
	TotalNetPrice  float64 `json:"totalNetPrice"`
	Committed      bool    `json:"committed"`
	MenuID         string  `json:"menuId"`
	TransactionID  string  `json:"transactionId"`
	TerminalID     string  `json:"terminalId"`
	PageKey        string  `json:"pageKey"`
	PromoCondition string  `json:"promoCondition"`
	Margin         float64 `json:"margin"`
}

type CompressedPayments []CompressedPayment

type CompressedPayment struct {
	AddOns         []interface{} `json:"addOns"`
	MenuParts      []interface{} `json:"menuParts"`
	ProductName    string        `json:"productName"`
	ProductKey     string        `json:"productKey"`
	Quantity       float64       `json:"quantity"`
	Units          int           `json:"units"`
	UnitPrice      float64       `json:"unitPrice"`
	TotalPrice     float64       `json:"totalPrice"`
	TotalNetPrice  float64       `json:"totalNetPrice"`
	IsExSales      bool          `json:"isExSales"`
	Committed      bool          `json:"committed"`
	MenuID         string        `json:"menuId"`
	TransactionID  string        `json:"transactionId"`
	TerminalID     string        `json:"terminalId"`
	PromoCondition string        `json:"promoCondition"`
	Memo           string        `json:"memo,omitempty"`
}

type Products []Product

type Product struct {
	ProductType            string        `json:"productType"`
	GroupKey               string        `json:"groupKey"`
	ImageURL               interface{}   `json:"imageUrl"`
	PrepName               string        `json:"prepName,omitempty"`
	PreparationInfo        string        `json:"preparationInfo,omitempty"`
	AutoAddons             []interface{} `json:"autoAddons"`
	AutoWindows            []interface{} `json:"autoWindows"`
	Parts                  []interface{} `json:"parts"`
	Translations           []interface{} `json:"translations"`
	Prices                 Prices        `json:"prices"`
	DisplayName            string        `json:"displayName,omitempty"`
	ProdNr                 int           `json:"prodNr,omitempty"`
	Name                   string        `json:"name,omitempty"`
	ShortName              string        `json:"shortName,omitempty"`
	Info                   string        `json:"info,omitempty"`
	AllowDiscount          bool          `json:"allowDiscount,omitempty"`
	Price                  float64       `json:"price,omitempty"`
	ShowInSearch           bool          `json:"showInSearch,omitempty"`
	PrintX                 int           `json:"printX,omitempty"`
	PrepPrintAddons        bool          `json:"prepPrintAddons,omitempty"`
	Description            string        `json:"description,omitempty"`
	Key                    string        `json:"key"`
	AskMemo                bool          `json:"askMemo,omitempty"`
	KdsSummary             bool          `json:"kdsSummary,omitempty"`
	AskPrice               bool          `json:"askPrice,omitempty"`
	PrintMemo              bool          `json:"printMemo,omitempty"`
	PrepTotal              bool          `json:"prepTotal,omitempty"`
	AskUnit                bool          `json:"askUnit,omitempty"`
	UnitID                 int           `json:"unitId,omitempty"`
	UnitBase               int           `json:"unitBase,omitempty"`
	VipOnly                bool          `json:"vipOnly,omitempty"`
	ReadOnly               bool          `json:"readOnly,omitempty"`
	ShowEditor             bool          `json:"showEditor,omitempty"`
	IncompleteMenuOptionID int           `json:"incompleteMenuOptionId,omitempty"`
	AllInMenu              bool          `json:"allInMenu,omitempty"`
	TableRequired          bool          `json:"tableRequired,omitempty"`
	VoucherServiceName     string        `json:"voucherServiceName,omitempty"`
	PrintCoveredProducts   bool          `json:"printCoveredProducts,omitempty"`
	AutoCalcPrice          bool          `json:"autoCalcPrice,omitempty"`
	AskCourse              bool          `json:"askCourse,omitempty"`
	AskQty                 bool          `json:"askQty,omitempty"`
}

type Prices []Price

type Price struct {
	SeqNr int     `json:"seqNr"`
	Price float64 `json:"price"`
	Key   string  `json:"key"`
}

type ProductGroups []ProductGroup

type ProductGroup struct {
	Name       string `json:"name"`
	Data       string `json:"data"`
	GroupLevel int    `json:"groupLevel"`
	LeftNr     int    `json:"leftNr"`
	RightNr    int    `json:"rightNr"`
	Key        string `json:"key"`
	GroupNr    int    `json:"groupNr,omitempty"`
}

type Paymodes []Paymode

type Paymode struct {
	DefaultCash     bool    `json:"defaultCash,omitempty"`
	Drawer          bool    `json:"drawer,omitempty"`
	IsCash          bool    `json:"isCash,omitempty"`
	OpenDrawer      bool    `json:"openDrawer"`
	GroupKey        string  `json:"groupKey"`
	GroupLeftNr     int     `json:"groupLeftNr"`
	GroupRightNr    int     `json:"groupRightNr"`
	TerminalName    string  `json:"terminalName"`
	Footer          string  `json:"footer"`
	Rounding        float64 `json:"rounding,omitempty"`
	Name            string  `json:"name"`
	ShortName       string  `json:"shortName"`
	Info            string  `json:"info"`
	AskPrice        bool    `json:"askPrice,omitempty"`
	ShowInSearch    bool    `json:"showInSearch"`
	PrintX          int     `json:"printX"`
	PrintMemo       bool    `json:"printMemo"`
	Description     string  `json:"description"`
	DisplayName     string  `json:"displayName"`
	Key             string  `json:"key"`
	PaymodeTypeID   int     `json:"paymodeTypeId,omitempty"`
	ProdNr          int     `json:"prodNr,omitempty"`
	AskMemo         bool    `json:"askMemo,omitempty"`
	InvoiceTypeID   int     `json:"invoiceTypeId,omitempty"`
	PrintSig        bool    `json:"printSig,omitempty"`
	AccountRequired bool    `json:"accountRequired,omitempty"`
}

type PaymodeGroups []PaymodeGroup

type PaymodeGroup struct {
	Name       string `json:"name"`
	Data       string `json:"data"`
	GroupLevel int    `json:"groupLevel"`
	LeftNr     int    `json:"leftNr"`
	RightNr    int    `json:"rightNr"`
	Key        string `json:"key"`
	GroupNr    int    `json:"groupNr,omitempty"`
}
