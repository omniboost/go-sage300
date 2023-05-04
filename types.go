package sage300

import (
	"github.com/omniboost/go-sage300/omitempty"
)

type GLAccounts []GLAccount

type GLAccount struct {
	UnformattedAccount             string        `json:"UnformattedAccount"`
	DateCreated                    Time          `json:"DateCreated"`
	Description                    string        `json:"Description"`
	AccountType                    string        `json:"AccountType"`
	NormalBalanceDRCR              string        `json:"NormalBalanceDRCR"`
	Status                         string        `json:"Status"`
	PostToAccount                  string        `json:"PostToAccount"`
	QuantitiesAllowed              string        `json:"QuantitiesAllowed"`
	UnitOfMeasure                  string        `json:"UnitOfMeasure"`
	AllocationsAllowed             string        `json:"AllocationsAllowed"`
	AllocOffsetAccount             string        `json:"AllocOffsetAccount"`
	AllocSourceType                string        `json:"AllocSourceType"`
	Multicurrency                  string        `json:"Multicurrency"`
	SpecificCurrency               string        `json:"SpecificCurrency"`
	ControlAccount                 string        `json:"ControlAccount"`
	AllocationPercentTotal         float64       `json:"AllocationPercentTotal"`
	StructureCode                  string        `json:"StructureCode"`
	YearLastClosed                 int           `json:"YearLastClosed"`
	AccountNumber                  string        `json:"AccountNumber"`
	AccountSegmentCode1            string        `json:"AccountSegmentCode1"`
	AccountSegmentCode2            string        `json:"AccountSegmentCode2"`
	AccountSegmentCode3            string        `json:"AccountSegmentCode3"`
	AccountSegmentCode4            string        `json:"AccountSegmentCode4"`
	AccountSegmentCode5            string        `json:"AccountSegmentCode5"`
	AccountSegmentCode6            string        `json:"AccountSegmentCode6"`
	AccountSegmentCode7            string        `json:"AccountSegmentCode7"`
	AccountSegmentCode8            string        `json:"AccountSegmentCode8"`
	AccountSegmentCode9            string        `json:"AccountSegmentCode9"`
	AccountSegmentCode10           string        `json:"AccountSegmentCode10"`
	SegmentCode                    string        `json:"SegmentCode"`
	PostToSegmentID                string        `json:"PostToSegmentID"`
	DefaultCurrencyCode            string        `json:"DefaultCurrencyCode"`
	NumberOfOptionalFields         int           `json:"NumberOfOptionalFields"`
	TransactionOptionalFields      int           `json:"TransactionOptionalFields"`
	ProcessSwitch                  string        `json:"ProcessSwitch"`
	AccountGroupCode               string        `json:"AccountGroupCode"`
	AccountGroupSortCode           string        `json:"AccountGroupSortCode"`
	RollupSwitch                   string        `json:"RollupSwitch"`
	RollupMemberNonRollup          string        `json:"RollupMemberNonRollup"`
	AcctSelectorFromAccount        string        `json:"AcctSelectorFromAccount"`
	AcctSelectorToAccount          string        `json:"AcctSelectorToAccount"`
	AcctSelectorFromSeg1Val        string        `json:"AcctSelectorFromSeg1Val"`
	AcctSelectorToSeg1Val          string        `json:"AcctSelectorToSeg1Val"`
	AcctSelectorFromSeg2Val        string        `json:"AcctSelectorFromSeg2Val"`
	AcctSelectorToSeg2Val          string        `json:"AcctSelectorToSeg2Val"`
	AcctSelectorFromSeg3Val        string        `json:"AcctSelectorFromSeg3Val"`
	AcctSelectorToSeg3Val          string        `json:"AcctSelectorToSeg3Val"`
	AcctSelectorFromSeg4Val        string        `json:"AcctSelectorFromSeg4Val"`
	AcctSelectorToSeg4Val          string        `json:"AcctSelectorToSeg4Val"`
	AcctSelectorFromSeg5Val        string        `json:"AcctSelectorFromSeg5Val"`
	AcctSelectorToSeg5Val          string        `json:"AcctSelectorToSeg5Val"`
	AcctSelectorFromSeg6Val        string        `json:"AcctSelectorFromSeg6Val"`
	AcctSelectorToSeg6Val          string        `json:"AcctSelectorToSeg6Val"`
	AcctSelectorFromSeg7Val        string        `json:"AcctSelectorFromSeg7Val"`
	AcctSelectorToSeg7Val          string        `json:"AcctSelectorToSeg7Val"`
	AcctSelectorFromSeg8Val        string        `json:"AcctSelectorFromSeg8Val"`
	AcctSelectorToSeg8Val          string        `json:"AcctSelectorToSeg8Val"`
	AcctSelectorFromSeg9Val        string        `json:"AcctSelectorFromSeg9Val"`
	AcctSelectorToSeg9Val          string        `json:"AcctSelectorToSeg9Val"`
	AcctSelectorFromSeg10Val       string        `json:"AcctSelectorFromSeg10Val"`
	AcctSelectorToSeg10Val         string        `json:"AcctSelectorToSeg10Val"`
	AcctSelectorOptionalField1     string        `json:"AcctSelectorOptionalField1"`
	AcctSelectorOptionalField2     string        `json:"AcctSelectorOptionalField2"`
	AcctSelectorOptionalField3     string        `json:"AcctSelectorOptionalField3"`
	AcctSelectorFromOpt1Val        string        `json:"AcctSelectorFromOpt1Val"`
	AcctSelectorToOpt1Val          string        `json:"AcctSelectorToOpt1Val"`
	AcctSelectorFromOpt2Val        string        `json:"AcctSelectorFromOpt2Val"`
	AcctSelectorToOpt2Val          string        `json:"AcctSelectorToOpt2Val"`
	AcctSelectorFromOpt3Val        string        `json:"AcctSelectorFromOpt3Val"`
	AcctSelectorToOpt3Val          string        `json:"AcctSelectorToOpt3Val"`
	AcctSelectorAcctType           int           `json:"AcctSelectorAcctType"`
	AcctSelectorFromAcctGroup      string        `json:"AcctSelectorFromAcctGroup"`
	AcctSelectorToAcctGroup        string        `json:"AcctSelectorToAcctGroup"`
	AcctSelectorQuantityAllowed    int           `json:"AcctSelectorQuantityAllowed"`
	NumberOfcreatedpreviewaccounts int           `json:"NumberOfcreatedpreviewaccounts"`
	RollupMembersInOtherAccounts   int           `json:"RollupMembersInOtherAccounts"`
	AccountAllocationInstructions  []interface{} `json:"AccountAllocationInstructions"`
	AccountValidCurrencies         []interface{} `json:"AccountValidCurrencies"`
	ControlAccountSubledgers       []interface{} `json:"ControlAccountSubledgers"`
	AccountOptionalFields          []struct {
		UnformattedAccount       string      `json:"UnformattedAccount"`
		OptionalField            string      `json:"OptionalField"`
		Value                    string      `json:"Value"`
		AccountOptionalFieldType string      `json:"AccountOptionalFieldType"`
		Length                   int         `json:"Length"`
		Decimals                 int         `json:"Decimals"`
		AllowBlank               bool        `json:"AllowBlank"`
		Validate                 bool        `json:"Validate"`
		ValueSet                 string      `json:"ValueSet"`
		TypedValueFieldIndex     int         `json:"TypedValueFieldIndex"`
		TextValue                string      `json:"TextValue"`
		AmountValue              float64     `json:"AmountValue"`
		NumberValue              int         `json:"NumberValue"`
		IntegerValue             int         `json:"IntegerValue"`
		YesNoValue               bool        `json:"YesNoValue"`
		DateValue                interface{} `json:"DateValue"`
		TimeValue                Time        `json:"TimeValue"`
		OptionalFieldDescription string      `json:"OptionalFieldDescription"`
		ValueDescription         string      `json:"ValueDescription"`
		UpdateOperation          string      `json:"UpdateOperation"`
	} `json:"AccountOptionalFields"`
	AccountTransactionOptionalFields []interface{} `json:"AccountTransactionOptionalFields"`
	RollupGroups                     []interface{} `json:"RollupGroups"`
	RollupMemberPreview              []interface{} `json:"RollupMemberPreview"`
	UpdateOperation                  string        `json:"UpdateOperation"`
}

type TaxRates []TaxRate

type TaxRate struct {
	TaxAuthority    string      `json:"TaxAuthority"`
	TransactionType string      `json:"TransactionType"`
	BuyerClass      int         `json:"BuyerClass"`
	ItemRate1       float64     `json:"ItemRate1"`
	ItemRate2       float64     `json:"ItemRate2"`
	ItemRate3       float64     `json:"ItemRate3"`
	ItemRate4       float64     `json:"ItemRate4"`
	ItemRate5       float64     `json:"ItemRate5"`
	ItemRate6       float64     `json:"ItemRate6"`
	ItemRate7       float64     `json:"ItemRate7"`
	ItemRate8       float64     `json:"ItemRate8"`
	ItemRate9       float64     `json:"ItemRate9"`
	ItemRate10      float64     `json:"ItemRate10"`
	LastMaintained  interface{} `json:"LastMaintained"`
	UpdateOperation string      `json:"UpdateOperation"`
}

type TaxGroups []TaxGroup

type TaxGroup struct {
	TaxGroupKey          string `json:"TaxGroupKey"`
	TransactionType      string `json:"TransactionType"`
	Description          string `json:"Description"`
	TaxReportingCurrency string `json:"TaxReportingCurrency"`
	TaxAuthority1        string `json:"TaxAuthority1"`
	TaxAuthority2        string `json:"TaxAuthority2"`
	TaxAuthority3        string `json:"TaxAuthority3"`
	TaxAuthority4        string `json:"TaxAuthority4"`
	TaxAuthority5        string `json:"TaxAuthority5"`
	Taxable1             bool   `json:"Taxable1"`
	Taxable2             bool   `json:"Taxable2"`
	Taxable3             bool   `json:"Taxable3"`
	Taxable4             bool   `json:"Taxable4"`
	Taxable5             bool   `json:"Taxable5"`
	TaxCalculationMethod string `json:"TaxCalculationMethod"`
	LastMaintained       Time   `json:"LastMaintained"`
	Surtax1              bool   `json:"Surtax1"`
	Surtax2              bool   `json:"Surtax2"`
	Surtax3              bool   `json:"Surtax3"`
	Surtax4              bool   `json:"Surtax4"`
	Surtax5              bool   `json:"Surtax5"`
	SurtaxOnAuthority1   string `json:"SurtaxOnAuthority1"`
	SurtaxOnAuthority2   string `json:"SurtaxOnAuthority2"`
	SurtaxOnAuthority3   string `json:"SurtaxOnAuthority3"`
	SurtaxOnAuthority4   string `json:"SurtaxOnAuthority4"`
	SurtaxOnAuthority5   string `json:"SurtaxOnAuthority5"`
	TaxReportingRateType string `json:"TaxReportingRateType"`
	UpdateOperation      string `json:"UpdateOperation"`
}

type PostedTransactions []PostedTransaction

type PostedTransaction struct {
	AccountNumber               string  `json:"AccountNumber"`
	FiscalYear                  string  `json:"FiscalYear"`
	FiscalPeriod                string  `json:"FiscalPeriod"`
	SourceCurrencyCode          string  `json:"SourceCurrencyCode"`
	SourceLedgerCode            string  `json:"SourceLedgerCode"`
	SourceTypeCode              string  `json:"SourceTypeCode"`
	PostingSequenceNumber       int     `json:"PostingSequenceNumber"`
	DetailCount                 int     `json:"DetailCount"`
	JournalDate                 Time    `json:"JournalDate"`
	BatchNumber                 string  `json:"BatchNumber"`
	JournalEntryNumber          string  `json:"JournalEntryNumber"`
	JournalTransactionNumber    int     `json:"JournalTransactionNumber"`
	ConsolidationOccurredOnPost string  `json:"ConsolidationOccurredOnPost"`
	CompanyID                   string  `json:"CompanyID"`
	JournalDetailDescription    string  `json:"JournalDetailDescription"`
	JournalDetailReference      string  `json:"JournalDetailReference"`
	JournalTransactionAmount    float64 `json:"JournalTransactionAmount"`
	JournalTransactionQuantity  float64 `json:"JournalTransactionQuantity"`
	NbrOfSourceCurrencyDecimals string  `json:"NbrOfSourceCurrencyDecimals"`
	SourceCurrencyAmount        float64 `json:"SourceCurrencyAmount"`
	HomeCurrencyCode            string  `json:"HomeCurrencyCode"`
	CurrencyRateTableType       string  `json:"CurrencyRateTableType"`
	SCURNCODE                   string  `json:"SCURNCODE"`
	DateOfCurrencyRateSelected  Time    `json:"DateOfCurrencyRateSelected"`
	CurrencyRateForConversion   int     `json:"CurrencyRateForConversion"`
	CurrencyRateSpreadAllowed   float64 `json:"CurrencyRateSpreadAllowed"`
	CodeForRateDateMatching     string  `json:"CodeForRateDateMatching"`
	CurrencyRateOperator        string  `json:"CurrencyRateOperator"`
	DrillDownType               int     `json:"DrillDownType"`
	DrillDownLinkNumber         int     `json:"DrillDownLinkNumber"`
	DrillDownApplicationSource  string  `json:"DrillDownApplicationSource"`
	ReportCurrencyAmount        float64 `json:"ReportCurrencyAmount"`
	NumberOfOptionalFields      int     `json:"NumberOfOptionalFields"`
	DocumentDate                Time    `json:"DocumentDate"`
	TaxReportingStatus          int     `json:"TaxReportingStatus"`
	TaxAuthority                string  `json:"TaxAuthority"`
	TaxAccountType              string  `json:"TaxAccountType"`
	UpdateOperation             string  `json:"UpdateOperation"`
}

type PostingJournalDetails []PostingJournalDetail

type PostingJournalDetail struct {
	Postingsequencenumber       int     `json:"Postingsequencenumber,omitempty"`
	BatchNumber                 string  `json:"BatchNumber,omitempty"`
	Journalentrynumber          string  `json:"Journalentrynumber,omitempty"`
	Journaltransactionnumber    int     `json:"Journaltransactionnumber,omitempty"`
	Journaldate                 Time    `json:"Journaldate,omitempty"`
	Fiscalyear                  string  `json:"Fiscalyear,omitempty"`
	Fiscalperiod                string  `json:"Fiscalperiod,omitempty"`
	SourceLedgerCode            string  `json:"SourceLedgerCode,omitempty"`
	SourceTypeCode              string  `json:"SourceTypeCode,omitempty"`
	ConsolidationoccurredOnpost string  `json:"ConsolidationoccurredOnpost,omitempty"`
	AccountNumber               string  `json:"AccountNumber,omitempty"`
	CompanyID                   string  `json:"CompanyID,omitempty"`
	Journaldetaildescription    string  `json:"Journaldetaildescription,omitempty"`
	Journaldetailreference      string  `json:"Journaldetailreference,omitempty"`
	Journaltransactionamount    float64 `json:"Journaltransactionamount,omitempty"`
	Journaltransactionquantity  float64 `json:"Journaltransactionquantity,omitempty"`
	NbrOfsourcecurrencydecimals string  `json:"NbrOfsourcecurrencydecimals,omitempty"`
	Sourcecurrencyamount        float64 `json:"Sourcecurrencyamount,omitempty"`
	Homecurrencycode            string  `json:"Homecurrencycode,omitempty"`
	Currencyratetabletype       string  `json:"Currencyratetabletype,omitempty"`
	Sourcecurrencycode          string  `json:"Sourcecurrencycode,omitempty"`
	DateOfcurrencyrateselected  Time    `json:"DateOfcurrencyrateselected,omitempty"`
	CurrencyrateForconversion   int     `json:"CurrencyrateForconversion,omitempty"`
	Currencyratespreadallowed   float64 `json:"Currencyratespreadallowed,omitempty"`
	CodeForratedatematching     string  `json:"CodeForratedatematching,omitempty"`
	Currencyrateoperator        string  `json:"Currencyrateoperator,omitempty"`
	Printedstatuscode           string  `json:"Printedstatuscode,omitempty"`
	PostingDate                 Time    `json:"PostingDate,omitempty"`
	Reportcurrencyamount        float64 `json:"Reportcurrencyamount,omitempty"`
	NumberOfOptionalFields      int     `json:"NumberOfOptionalFields,omitempty"`
	Originator                  string  `json:"Originator,omitempty"`
	AutoReversal                string  `json:"AutoReversal,omitempty"`
	Destination                 string  `json:"Destination,omitempty"`
	RouteNumber                 int     `json:"RouteNumber,omitempty"`
	DocumentDate                Time    `json:"DocumentDate,omitempty"`
	TaxAuthority                string  `json:"TaxAuthority,omitempty"`
	TaxAccountType              string  `json:"TaxAccountType,omitempty"`
	UpdateOperation             string  `json:"UpdateOperation,omitempty"`
}

func (d PostingJournalDetail) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(d)
}

type GLJournalBatch struct {
	OdataContext                     string                       `json:"@odata.context,omitempty"`
	BatchNumber                      string                       `json:"BatchNumber,omitempty"`
	ActiveSwitch                     string                       `json:"ActiveSwitch,omitempty"`
	Description                      string                       `json:"Description,omitempty"`
	SourceLedger                     string                       `json:"SourceLedger,omitempty"`
	DateCreated                      Time                         `json:"DateCreated,omitempty"`
	DateLastEdited                   Time                         `json:"DateLastEdited,omitempty"`
	JournalBatchType                 string                       `json:"JournalBatchType,omitempty"`
	Status                           string                       `json:"Status,omitempty"`
	PostingSequence                  int                          `json:"PostingSequence,omitempty"`
	Debits                           float64                      `json:"Debits,omitempty"`
	Credits                          float64                      `json:"Credits,omitempty"`
	QuantityTotal                    float64                      `json:"QuantityTotal,omitempty"`
	NumberOfEntries                  int                          `json:"NumberOfEntries,omitempty"`
	NextEntryNumber                  int                          `json:"NextEntryNumber,omitempty"`
	NumberOfErrors                   int                          `json:"NumberOfErrors,omitempty"`
	OriginalStatus                   string                       `json:"OriginalStatus,omitempty"`
	Printed                          string                       `json:"Printed,omitempty"`
	ICTRelated                       string                       `json:"ICTRelated,omitempty"`
	RevaluationRecognizedBatch       string                       `json:"RevaluationRecognizedBatch,omitempty"`
	EligibleForEditSwitch            string                       `json:"EligibleForEditSwitch,omitempty"`
	EligibleForDeleteSwitch          string                       `json:"EligibleForDeleteSwitch,omitempty"`
	EligibleForPostingSwitch         string                       `json:"EligibleForPostingSwitch,omitempty"`
	EligibleForPrintingSwitch        string                       `json:"EligibleForPrintingSwitch,omitempty"`
	EligibleForProvisionalPostSwitch string                       `json:"EligibleForProvisionalPostSwitch,omitempty"`
	ReadyToPost                      string                       `json:"ReadyToPost,omitempty"`
	LockBatchSwitch                  string                       `json:"LockBatchSwitch,omitempty"`
	JournalHeaders                   GLJournalBatchJournalHeaders `json:"JournalHeaders,omitempty"`
	UpdateOperation                  string                       `json:"UpdateOperation,omitempty"`
}

func (b GLJournalBatch) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(b)
}

type JournalDetails []JournalDetail

// func (d JournalDetails) MarshalJSON() ([]byte, error) {
// 	return omitempty.MarshalJSON(d)
// }

type JournalDetail struct {
	BatchNumber                  string  `json:"BatchNumber,omitempty"`
	EntryNumber                  string  `json:"EntryNumber,omitempty"`
	TransactionNumber            string  `json:"TransactionNumber,omitempty"`
	Destination                  string  `json:"Destination,omitempty"`
	RouteNumber                  int     `json:"RouteNumber,omitempty"`
	AccountNumber                string  `json:"AccountNumber,omitempty"`
	CompanyID                    string  `json:"CompanyID,omitempty"`
	Amount                       float64 `json:"Amount,omitempty"`
	Quantity                     float64 `json:"Quantity,omitempty"`
	SourceCurrencyDecimals       string  `json:"SourceCurrencyDecimals,omitempty"`
	SourceCurrencyAmount         float64 `json:"SourceCurrencyAmount,omitempty"`
	HomeCurrency                 string  `json:"HomeCurrency,omitempty"`
	CurrencyRateTable            string  `json:"CurrencyRateTable,omitempty"`
	SourceCurrency               string  `json:"SourceCurrency,omitempty"`
	CurrencyRateDate             Time    `json:"CurrencyRateDate,omitempty"`
	CurrencyRate                 int     `json:"CurrencyRate,omitempty"`
	CurrencyRateSpread           float64 `json:"CurrencyRateSpread,omitempty"`
	CurrencyRateDateMatching     string  `json:"CurrencyRateDateMatching,omitempty"`
	CurrencyRateOperator         string  `json:"CurrencyRateOperator,omitempty"`
	Description                  string  `json:"Description,omitempty"`
	Reference                    string  `json:"Reference,omitempty"`
	JournalDate                  Time    `json:"JournalDate,omitempty"`
	SourceLedger                 string  `json:"SourceLedger,omitempty"`
	SourceType                   string  `json:"SourceType,omitempty"`
	Comment                      string  `json:"Comment,omitempty"`
	NumberOfOptionalFields       int     `json:"NumberOfOptionalFields,omitempty"`
	Processswitches              string  `json:"Processswitches,omitempty"`
	DestExists                   int     `json:"DestExists,omitempty"`
	DestDescription              string  `json:"DestDescription,omitempty"`
	DestStatus                   int     `json:"DestStatus,omitempty"`
	DestMulticurrencySwitch      int     `json:"DestMulticurrencySwitch,omitempty"`
	DestHomeCurrency             string  `json:"DestHomeCurrency,omitempty"`
	DestHomeCurrencyDecimals     int     `json:"DestHomeCurrencyDecimals,omitempty"`
	DestQtySwitch                int     `json:"DestQtySwitch,omitempty"`
	DestQtyDecimals              int     `json:"DestQtyDecimals,omitempty"`
	DestOptionalFlds             int     `json:"DestOptionalFlds,omitempty"`
	RouteExists                  int     `json:"RouteExists,omitempty"`
	RouteDescription             string  `json:"RouteDescription,omitempty"`
	RouteStatus                  int     `json:"RouteStatus,omitempty"`
	SourceCodeExists             int     `json:"SourceCodeExists,omitempty"`
	SourceCodeDescription        string  `json:"SourceCodeDescription,omitempty"`
	AcctExists                   int     `json:"AcctExists,omitempty"`
	AcctFormatted                string  `json:"AcctFormatted,omitempty"`
	AcctDescription              string  `json:"AcctDescription,omitempty"`
	AcctStatus                   int     `json:"AcctStatus,omitempty"`
	AcctMulticurrencySwitch      int     `json:"AcctMulticurrencySwitch,omitempty"`
	AcctSpecificCurrenciesSwitch int     `json:"AcctSpecificCurrenciesSwitch,omitempty"`
	AcctControlSwitch            int     `json:"AcctControlSwitch,omitempty"`
	AcctQtySwitch                int     `json:"AcctQtySwitch,omitempty"`
	AcctOptionalFlds             int     `json:"AcctOptionalFlds,omitempty"`
	AcctTransactionOptionalFlds  int     `json:"AcctTransactionOptionalFlds,omitempty"`
	AcctUnitOfMeasure            string  `json:"AcctUnitOfMeasure,omitempty"`
	TaxAuthority                 string  `json:"TaxAuthority,omitempty"`
	TaxAccountType               int     `json:"TaxAccountType,omitempty"`
	// JournalDetailOptionalFields  []struct {
	// 	BatchNumber                    string    `json:"BatchNumber,omitempty"`
	// 	EntryNumber                    string    `json:"EntryNumber,omitempty"`
	// 	TransactionNumber              string    `json:"TransactionNumber,omitempty"`
	// 	OptionalField                  string    `json:"OptionalField,omitempty"`
	// 	Value                          string    `json:"Value,omitempty"`
	// 	JournalDetailOptionalFieldType string    `json:"JournalDetailOptionalFieldType,omitempty"`
	// 	Length                         int       `json:"Length,omitempty"`
	// 	Decimals                       int       `json:"Decimals,omitempty"`
	// 	AllowBlank                     bool      `json:"AllowBlank,omitempty"`
	// 	Validate                       bool      `json:"Validate,omitempty"`
	// 	ValueSet                       string    `json:"ValueSet,omitempty"`
	// 	TypedValueFieldIndex           int       `json:"TypedValueFieldIndex,omitempty"`
	// 	TextValue                      string    `json:"TextValue,omitempty"`
	// 	AmountValue                    int       `json:"AmountValue,omitempty"`
	// 	NumberValue                    int       `json:"NumberValue,omitempty"`
	// 	IntegerValue                   int       `json:"IntegerValue,omitempty"`
	// 	YesNoValue                     bool      `json:"YesNoValue,omitempty"`
	// 	DateValue                      Time `json:"DateValue,omitempty"`
	// 	TimeValue                      Time `json:"TimeValue,omitempty"`
	// 	OptionalFieldDescription       string    `json:"OptionalFieldDescription,omitempty"`
	// 	ValueDescription               string    `json:"ValueDescription,omitempty"`
	// 	UpdateOperation                string    `json:"UpdateOperation,omitempty"`
	// 	Warnings                       []struct {
	// 		Message        string `json:"Message,omitempty"`
	// 		Priority       string `json:"Priority,omitempty"`
	// 		PriorityString string `json:"PriorityString,omitempty"`
	// 		Tag            struct {
	// 		} `json:"Tag,omitempty"`
	// 	} `json:"Warnings,omitempty"`
	// 	ETag           string `json:"ETag,omitempty"`
	// 	IsDeleted      bool   `json:"IsDeleted,omitempty"`
	// 	IsNewLine      bool   `json:"IsNewLine,omitempty"`
	// 	HasChanged     bool   `json:"HasChanged,omitempty"`
	// 	DisplayIndex   int    `json:"DisplayIndex,omitempty"`
	// 	ChangeSequence int    `json:"ChangeSequence,omitempty"`
	// 	PreviousKey    string `json:"PreviousKey,omitempty"`
	// 	Licenses       struct {
	// 	} `json:"Licenses,omitempty"`
	// 	IsSelected bool `json:"IsSelected,omitempty"`
	// } `json:"JournalDetailOptionalFields,omitempty"`
	UpdateOperation string `json:"UpdateOperation,omitempty"`
	Warnings        []struct {
		Message        string `json:"Message,omitempty"`
		Priority       string `json:"Priority,omitempty"`
		PriorityString string `json:"PriorityString,omitempty"`
		Tag            struct {
		} `json:"Tag,omitempty"`
	} `json:"Warnings,omitempty"`
	ETag           string `json:"ETag,omitempty"`
	IsDeleted      bool   `json:"IsDeleted,omitempty"`
	IsNewLine      bool   `json:"IsNewLine,omitempty"`
	HasChanged     bool   `json:"HasChanged,omitempty"`
	DisplayIndex   int    `json:"DisplayIndex,omitempty"`
	ChangeSequence int    `json:"ChangeSequence,omitempty"`
	PreviousKey    string `json:"PreviousKey,omitempty"`
	// Licenses       struct {
	// } `json:"Licenses,omitempty"`
	IsSelected                  bool          `json:"IsSelected,omitempty"`
	JournalDetailOptionalFields []interface{} `json:"JournalDetailOptionalFields,omitempty"`
}

func (d JournalDetail) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(d)
}

type GLJournalBatchJournalHeaders []GLJournalBatchJournalHeader

type GLJournalBatchJournalHeader struct {
	BatchNumber                    string         `json:"BatchNumber,omitempty"`
	EntryNumber                    string         `json:"EntryNumber,omitempty"`
	Originator                     string         `json:"Originator,omitempty"`
	SourceLedger                   string         `json:"SourceLedger,omitempty"`
	SourceType                     string         `json:"SourceType,omitempty"`
	FiscalYear                     string         `json:"FiscalYear,omitempty"`
	FiscalPeriod                   string         `json:"FiscalPeriod,omitempty"`
	AutoReversal                   string         `json:"AutoReversal,omitempty"`
	Description                    string         `json:"Description,omitempty"`
	Debits                         float64        `json:"Debits,omitempty"`
	Credits                        float64        `json:"Credits,omitempty"`
	Quantity                       float64        `json:"Quantity,omitempty"`
	PostingDate                    Time           `json:"PostingDate,omitempty"`
	DrillDownType                  int            `json:"DrillDownType,omitempty"`
	DrillDownLinkNumber            int            `json:"DrillDownLinkNumber,omitempty"`
	DrillDownApplicationSource     string         `json:"DrillDownApplicationSource,omitempty"`
	OutOfBalanceBy                 float64        `json:"OutOfBalanceBy,omitempty"`
	SpecificReversalYear           string         `json:"SpecificReversalYear,omitempty"`
	SpecificReversalPeriod         string         `json:"SpecificReversalPeriod,omitempty"`
	ErrorBatch                     int            `json:"ErrorBatch,omitempty"`
	ErrorEntry                     int            `json:"ErrorEntry,omitempty"`
	NumberOfDetails                int            `json:"NumberOfDetails,omitempty"`
	Processswitches                string         `json:"Processswitches,omitempty"`
	EnteredBy                      string         `json:"EnteredBy,omitempty"`
	DocumentDate                   Time           `json:"DocumentDate,omitempty"`
	OrigExists                     int            `json:"OrigExists,omitempty"`
	OrigDescription                string         `json:"OrigDescription,omitempty"`
	OrigStatus                     int            `json:"OrigStatus,omitempty"`
	OrigMulticurrencySwitch        int            `json:"OrigMulticurrencySwitch,omitempty"`
	OrigHomeCurrency               string         `json:"OrigHomeCurrency,omitempty"`
	OrigHomeCurrencyDecimals       int            `json:"OrigHomeCurrencyDecimals,omitempty"`
	OrigQtySwitch                  int            `json:"OrigQtySwitch,omitempty"`
	OrigQtyDecimals                int            `json:"OrigQtyDecimals,omitempty"`
	SourceCodeExists               int            `json:"SourceCodeExists,omitempty"`
	SourceCodeDescription          string         `json:"SourceCodeDescription,omitempty"`
	TaxGroup                       string         `json:"TaxGroup,omitempty"`
	EntryType                      string         `json:"EntryType,omitempty"`
	DoNotCalculationTax            string         `json:"DoNotCalculationTax,omitempty"`
	CustomerVendorNumber           string         `json:"CustomerVendorNumber,omitempty"`
	DocumentType                   string         `json:"DocumentType,omitempty"`
	DocumentNumber                 string         `json:"DocumentNumber,omitempty"`
	TaxAuthority1                  string         `json:"TaxAuthority1,omitempty"`
	TaxAuthority2                  string         `json:"TaxAuthority2,omitempty"`
	TaxAuthority3                  string         `json:"TaxAuthority3,omitempty"`
	TaxAuthority4                  string         `json:"TaxAuthority4,omitempty"`
	TaxAuthority5                  string         `json:"TaxAuthority5,omitempty"`
	TaxVendorClass1                int            `json:"TaxVendorClass1,omitempty"`
	TaxVendorClass2                int            `json:"TaxVendorClass2,omitempty"`
	TaxVendorClass3                int            `json:"TaxVendorClass3,omitempty"`
	TaxVendorClass4                int            `json:"TaxVendorClass4,omitempty"`
	TaxVendorClass5                int            `json:"TaxVendorClass5,omitempty"`
	TaxItemClass1                  int            `json:"TaxItemClass1,omitempty"`
	TaxItemClass2                  int            `json:"TaxItemClass2,omitempty"`
	TaxItemClass3                  int            `json:"TaxItemClass3,omitempty"`
	TaxItemClass4                  int            `json:"TaxItemClass4,omitempty"`
	TaxItemClass5                  int            `json:"TaxItemClass5,omitempty"`
	TaxBaseAmount1                 float64        `json:"TaxBaseAmount1,omitempty"`
	TaxBaseAmount2                 float64        `json:"TaxBaseAmount2,omitempty"`
	TaxBaseAmount3                 float64        `json:"TaxBaseAmount3,omitempty"`
	TaxBaseAmount4                 float64        `json:"TaxBaseAmount4,omitempty"`
	TaxBaseAmount5                 float64        `json:"TaxBaseAmount5,omitempty"`
	TaxAmount1                     float64        `json:"TaxAmount1,omitempty"`
	TaxAmount2                     float64        `json:"TaxAmount2,omitempty"`
	TaxAmount3                     float64        `json:"TaxAmount3,omitempty"`
	TaxAmount4                     float64        `json:"TaxAmount4,omitempty"`
	TaxAmount5                     float64        `json:"TaxAmount5,omitempty"`
	TaxExpenseAmount1              float64        `json:"TaxExpenseAmount1,omitempty"`
	TaxExpenseAmount2              float64        `json:"TaxExpenseAmount2,omitempty"`
	TaxExpenseAmount3              float64        `json:"TaxExpenseAmount3,omitempty"`
	TaxExpenseAmount4              float64        `json:"TaxExpenseAmount4,omitempty"`
	TaxExpenseAmount5              float64        `json:"TaxExpenseAmount5,omitempty"`
	TaxRecoverableAmount1          float64        `json:"TaxRecoverableAmount1,omitempty"`
	TaxRecoverableAmount2          float64        `json:"TaxRecoverableAmount2,omitempty"`
	TaxRecoverableAmount3          float64        `json:"TaxRecoverableAmount3,omitempty"`
	TaxRecoverableAmount4          float64        `json:"TaxRecoverableAmount4,omitempty"`
	TaxRecoverableAmount5          float64        `json:"TaxRecoverableAmount5,omitempty"`
	TaxAllocatedAmount1            float64        `json:"TaxAllocatedAmount1,omitempty"`
	TaxAllocatedAmount2            float64        `json:"TaxAllocatedAmount2,omitempty"`
	TaxAllocatedAmount3            float64        `json:"TaxAllocatedAmount3,omitempty"`
	TaxAllocatedAmount4            float64        `json:"TaxAllocatedAmount4,omitempty"`
	TaxAllocatedAmount5            float64        `json:"TaxAllocatedAmount5,omitempty"`
	TaxReportingAmount1            float64        `json:"TaxReportingAmount1,omitempty"`
	TaxReportingAmount2            float64        `json:"TaxReportingAmount2,omitempty"`
	TaxReportingAmount3            float64        `json:"TaxReportingAmount3,omitempty"`
	TaxReportingAmount4            float64        `json:"TaxReportingAmount4,omitempty"`
	TaxReportingAmount5            float64        `json:"TaxReportingAmount5,omitempty"`
	TaxReportingExpensed1          float64        `json:"TaxReportingExpensed1,omitempty"`
	TaxReportingExpensed2          float64        `json:"TaxReportingExpensed2,omitempty"`
	TaxReportingExpensed3          float64        `json:"TaxReportingExpensed3,omitempty"`
	TaxReportingExpensed4          float64        `json:"TaxReportingExpensed4,omitempty"`
	TaxReportingExpensed5          float64        `json:"TaxReportingExpensed5,omitempty"`
	TaxReportingRecoverableAmount1 float64        `json:"TaxReportingRecoverableAmount1,omitempty"`
	TaxReportingRecoverableAmount2 float64        `json:"TaxReportingRecoverableAmount2,omitempty"`
	TaxReportingRecoverableAmount3 float64        `json:"TaxReportingRecoverableAmount3,omitempty"`
	TaxReportingRecoverableAmount4 float64        `json:"TaxReportingRecoverableAmount4,omitempty"`
	TaxReportingRecoverableAmount5 float64        `json:"TaxReportingRecoverableAmount5,omitempty"`
	TaxReportingAllocatedAmount1   float64        `json:"TaxReportingAllocatedAmount1,omitempty"`
	TaxReportingAllocatedAmount2   float64        `json:"TaxReportingAllocatedAmount2,omitempty"`
	TaxReportingAllocatedAmount3   float64        `json:"TaxReportingAllocatedAmount3,omitempty"`
	TaxReportingAllocatedAmount4   float64        `json:"TaxReportingAllocatedAmount4,omitempty"`
	TaxReportingAllocatedAmount5   float64        `json:"TaxReportingAllocatedAmount5,omitempty"`
	TaxReportingCurrencyCode       string         `json:"TaxReportingCurrencyCode,omitempty"`
	TaxReportingRate               float64        `json:"TaxReportingRate,omitempty"`
	TaxReportingRateType           string         `json:"TaxReportingRateType,omitempty"`
	TaxReportingRateDate           Time           `json:"TaxReportingRateDate,omitempty"`
	TaxReportingRateOperation      string         `json:"TaxReportingRateOperation,omitempty"`
	TaxExpenseAccount1             string         `json:"TaxExpenseAccount1,omitempty"`
	TaxExpenseAccount2             string         `json:"TaxExpenseAccount2,omitempty"`
	TaxExpenseAccount3             string         `json:"TaxExpenseAccount3,omitempty"`
	TaxExpenseAccount4             string         `json:"TaxExpenseAccount4,omitempty"`
	TaxExpenseAccount5             string         `json:"TaxExpenseAccount5,omitempty"`
	TaxRecoverableAccount1         string         `json:"TaxRecoverableAccount1,omitempty"`
	TaxRecoverableAccount2         string         `json:"TaxRecoverableAccount2,omitempty"`
	TaxRecoverableAccount3         string         `json:"TaxRecoverableAccount3,omitempty"`
	TaxRecoverableAccount4         string         `json:"TaxRecoverableAccount4,omitempty"`
	TaxRecoverableAccount5         string         `json:"TaxRecoverableAccount5,omitempty"`
	TaxRate1                       float64        `json:"TaxRate1,omitempty"`
	TaxRate2                       float64        `json:"TaxRate2,omitempty"`
	TaxRate3                       float64        `json:"TaxRate3,omitempty"`
	TaxRate4                       float64        `json:"TaxRate4,omitempty"`
	TaxRate5                       float64        `json:"TaxRate5,omitempty"`
	FunctionalTaxAmount1           float64        `json:"FunctionalTaxAmount1,omitempty"`
	FunctionalTaxAmount2           float64        `json:"FunctionalTaxAmount2,omitempty"`
	FunctionalTaxAmount3           float64        `json:"FunctionalTaxAmount3,omitempty"`
	FunctionalTaxAmount4           float64        `json:"FunctionalTaxAmount4,omitempty"`
	FunctionalTaxAmount5           float64        `json:"FunctionalTaxAmount5,omitempty"`
	FunctionalTaxBaseAmount1       float64        `json:"FunctionalTaxBaseAmount1,omitempty"`
	FunctionalTaxBaseAmount2       float64        `json:"FunctionalTaxBaseAmount2,omitempty"`
	FunctionalTaxBaseAmount3       float64        `json:"FunctionalTaxBaseAmount3,omitempty"`
	FunctionalTaxBaseAmount4       float64        `json:"FunctionalTaxBaseAmount4,omitempty"`
	FunctionalTaxBaseAmount5       float64        `json:"FunctionalTaxBaseAmount5,omitempty"`
	FunctionalExpensedAmount1      float64        `json:"FunctionalExpensedAmount1,omitempty"`
	FunctionalExpensedAmount2      float64        `json:"FunctionalExpensedAmount2,omitempty"`
	FunctionalExpensedAmount3      float64        `json:"FunctionalExpensedAmount3,omitempty"`
	FunctionalExpensedAmount4      float64        `json:"FunctionalExpensedAmount4,omitempty"`
	FunctionalExpensedAmount5      float64        `json:"FunctionalExpensedAmount5,omitempty"`
	FunctionalRecoverableAmount1   float64        `json:"FunctionalRecoverableAmount1,omitempty"`
	FunctionalRecoverableAmount2   float64        `json:"FunctionalRecoverableAmount2,omitempty"`
	FunctionalRecoverableAmount3   float64        `json:"FunctionalRecoverableAmount3,omitempty"`
	FunctionalRecoverableAmount4   float64        `json:"FunctionalRecoverableAmount4,omitempty"`
	FunctionalRecoverableAmount5   float64        `json:"FunctionalRecoverableAmount5,omitempty"`
	FunctionalAllocatedAmount1     float64        `json:"FunctionalAllocatedAmount1,omitempty"`
	FunctionalAllocatedAmount2     float64        `json:"FunctionalAllocatedAmount2,omitempty"`
	FunctionalAllocatedAmount3     float64        `json:"FunctionalAllocatedAmount3,omitempty"`
	FunctionalAllocatedAmount4     float64        `json:"FunctionalAllocatedAmount4,omitempty"`
	FunctionalAllocatedAmount5     float64        `json:"FunctionalAllocatedAmount5,omitempty"`
	InvoicetotalInTaxGroupCurre    float64        `json:"InvoicetotalInTaxGroupCurre,omitempty"`
	InvoicetotalInfunctionalCurr   float64        `json:"InvoicetotalInfunctionalCurr,omitempty"`
	TaxGroupDescription            string         `json:"TaxGroupDescription,omitempty"`
	TaxAuthorityDescription1       string         `json:"TaxAuthorityDescription1,omitempty"`
	TaxAuthorityDescription2       string         `json:"TaxAuthorityDescription2,omitempty"`
	TaxAuthorityDescription3       string         `json:"TaxAuthorityDescription3,omitempty"`
	TaxAuthorityDescription4       string         `json:"TaxAuthorityDescription4,omitempty"`
	TaxAuthorityDescription5       string         `json:"TaxAuthorityDescription5,omitempty"`
	TaxReportingCurrencyCodeDesc   string         `json:"TaxReportingCurrencyCodeDesc,omitempty"`
	RateTypeDescription            string         `json:"RateTypeDescription,omitempty"`
	CustomerVendorName             string         `json:"CustomerVendorName,omitempty"`
	JournalDetails                 JournalDetails `json:"JournalDetails,omitempty"`
	UpdateOperation                string         `json:"UpdateOperation,omitempty"`
	Warnings                       []struct {
		Message        string `json:"Message,omitempty"`
		Priority       string `json:"Priority,omitempty"`
		PriorityString string `json:"PriorityString,omitempty"`
		Tag            struct {
		} `json:"Tag,omitempty"`
	} `json:"Warnings,omitempty"`
	ETag           string `json:"ETag,omitempty"`
	IsDeleted      bool   `json:"IsDeleted,omitempty"`
	IsNewLine      bool   `json:"IsNewLine,omitempty"`
	HasChanged     bool   `json:"HasChanged,omitempty"`
	DisplayIndex   int    `json:"DisplayIndex,omitempty"`
	ChangeSequence int    `json:"ChangeSequence,omitempty"`
	PreviousKey    string `json:"PreviousKey,omitempty"`
	// Licenses       struct {
	// } `json:"Licenses,omitempty"`
	IsSelected bool `json:"IsSelected,omitempty"`
}

func (h GLJournalBatchJournalHeader) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(h)
}
