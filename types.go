package accountviewnet

import (
	"strings"
)

type Fields []Field

type Field string

func (ff Fields) MarshalSchema() string {
	fields := make([]string, len(ff))
	for i, f := range ff {
		fields[i] = string(f)
	}
	return strings.Join(fields, ",")
}

type BusinessObjects []BusinessObject

type BusinessObject struct {
	Code          string `json:"Code"`
	Description   string `json:"Description"`
	InsertAllowed bool   `json:"InsertAllowed"`
	UpdateAllowed bool   `json:"UpdateAllowed"`
	DeleteAllowed bool   `json:"DeleteAllowed"`
}

type Companies []Company

type Company struct {
	ID          string `json:"Id"`
	Description string `json:"Description"`
	Code        string `json:"Code"`
}

type VAT struct {
	AcctMemo  string   `json:"ACCT_MEMO"`  // Rekeningnummer BTW verlegd
	CalcCum   bool     `json:"CALC_CUM"`   // BTW berekenen over BTW
	CalcType  float64  `json:"CALC_TYPE"`  // BTW-rekenwijze
	CngDate   DateTime `json:"CNG_DATE"`   // Wijzigingsdatum
	CngNr     int      `json:"CNG_NR"`     // Wijzignummer
	CngUsr    string   `json:"CNG_USR"`    // Medewerker laatste wijziging
	CogsType  bool     `json:"COGS_TYPE"`  // BTW toevoegen aan kostprijs
	FrmCode   string   `json:"FRM_CODE"`   // Rubr.code
	FrmMemo   string   `json:"FRM_MEMO"`   // Gekoppelde rubriek BTW verlegd
	InpDate   DateTime `json:"INP_DATE"`   // Invoerdatum
	InpUsr    string   `json:"INP_USR"`    // Invoermedewerker
	IsAbc     bool     `json:"IS_ABC"`     // ABC-transactie
	IsReverse bool     `json:"IS_REVERSE"` // BTW verlegd
	VATAcct   string   `json:"VAT_ACCT"`   // Rekeningnr
	VATCode   string   `json:"VAT_CODE"`   // BTW-code
	VATLink   string   `json:"VAT_LINK"`   // Gekoppelde code
	VATName   string   `json:"VAT_NAME"`   // Omschrijving BTW-code
	VATPct    float64  `json:"VAT_PCT"`    // BTW-percentage
	VATType   float64  `json:"VAT_TYPE"`   // BTW-type (numeriek)

	ExtKey int `json:"EXT_KEY"`
}
