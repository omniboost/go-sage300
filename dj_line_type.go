package accountviewnet

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type DjLine struct {
	AccrPy    bool    `json:"ACCR_PY"`    // Posten vorig boekjaar
	AcctNr    string  `json:"ACCT_NR"`    // Rekeningnr
	AdmCode   string  `json:"ADM_CODE"`   // Admin.code
	Amount    float64 `json:"AMOUNT"`     // Mutatiebedrag
	BankDesc  string  `json:"BANK_DESC"`  // Omschrijving bankmutatie
	BaseAmt   float64 `json:"BASE_AMT"`   // Basisbedrag BTW
	BegPer    float64 `json:"BEG_PER"`    // Begin periode
	CnID      string  `json:"CN_ID"`      // Unieke consolidatiecode
	CostCode  string  `json:"COST_CODE"`  // Kostenplaatscode
	CredDays  float64 `json:"CRED_DAYS"`  // Betalingstermijn in dagen
	CurAmt    float64 `json:"CUR_AMT"`    // Valutabedrag
	CurCode   string  `json:"CUR_CODE"`   // Valutacode
	CurVAT    float64 `json:"CUR_VAT"`    // BTW-bedrag vreemde valuta
	DifAmt    float64 `json:"DIF_AMT"`    // Verschilbedrag
	DifCur    float64 `json:"DIF_CUR"`    // Valutaverschil
	DifType   float64 `json:"DIF_TYPE"`   // Type verschil
	DistCode  string  `json:"DIST_CODE"`  // Verdeelsleutel
	DistQty   float64 `json:"DIST_QTY"`   // Aantal
	DjlRate   float64 `json:"DJL_RATE"`   // Valutakoers dagboekregel
	DjCode    string  `json:"DJ_CODE"`    // Dagboekcode
	EbnkStat  int     `json:"EBNK_STAT"`  // Status bankkoppeling
	EndPer    float64 `json:"END_PER"`    // Eind periode
	ExchRate  float64 `json:"EXCH_RATE"`  // Valutakoers
	ExpID     string  `json:"EXP_ID"`     // Export-ID
	GrekType  string  `json:"GREK_TYPE"`  // Transactietype G-rekening
	HomeDif   float64 `json:"HOME_DIF"`   // Verschil in adm.valuta
	InvNr     string  `json:"INV_NR"`     // Document-/factuurnummer
	ObliNr    string  `json:"OBLI_NR"`    // Verplichtingnummer
	PageNr    string  `json:"PAGE_NR"`    // Bladzijdenummer (teken)
	PageNrv   float64 `json:"PAGE_NRV"`   // Bladzijdenummer
	PartCode  string  `json:"PART_CODE"`  // Vennoot
	PcCode    string  `json:"PC_CODE"`    // Kostensoortcode
	PdifVAT   float64 `json:"PDIF_VAT"`   // BTW-bedrag betalingskorting
	ProjCode  string  `json:"PROJ_CODE"`  // Projectcode
	RaAmt     float64 `json:"RA_AMT"`     // Bedrag cheque-begeleiding
	RaDisc    float64 `json:"RA_DISC"`    // Kortingsbedrag cheque-begeleiding
	RecID     string  `json:"REC_ID"`     // Unieke code
	RecOrd    int     `json:"REC_ORD"`    // Regelvolgorde
	SrcID     string  `json:"SRC_ID"`     // Identificatie bron
	SrcType   float64 `json:"SRC_TYPE"`   // Brontype
	SubNr     string  `json:"SUB_NR"`     // Debiteur-/crediteurnummer
	TrnDate   Date    `json:"TRN_DATE"`   // Mutatiedatum
	TrnDesc   string  `json:"TRN_DESC"`   // Omschrijving regel
	TrnQty    float64 `json:"TRN_QTY"`    // Aantal (agr)
	TrnWgt    float64 `json:"TRN_WGT"`    // Gewicht
	VATAmt    float64 `json:"VAT_AMT"`    // BTW-bedrag
	VATCode   string  `json:"VAT_CODE"`   // BTW-code
	VexchRate float64 `json:"VEXCH_RATE"` // Valutakoers (import)
	WkrTrn    bool    `json:"WKR_TRN"`    // Mutatie ten laste van vrije ruimte WKR

	ExtKey   int `json:"EXT_KEY"`
	DjPageID int `json:"DJ_PAGE_Id"`
}

func (djLine DjLine) BusinessObject() string {
	return ""
}

func (djLine DjLine) Table() string {
	return "DJ_PAGE"
}

func (djLine DjLine) Fields() []string {
	return []string{
		"ACCT_NR",
		"AMOUNT",
	}
}

func (djLine DjLine) Values() ([]interface{}, error) {
	fields := djLine.Fields()
	values := make([]interface{}, len(fields))
	b, err := json.Marshal(djLine)
	if err != nil {
		return values, errors.WithStack(err)
	}

	m := map[string]interface{}{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return values, errors.WithStack(err)
	}

	for i, f := range fields {
		values[i] = m[f]
	}

	return values, nil
}
