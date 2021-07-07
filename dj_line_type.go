package accountviewnet

type DjLine struct {
	AccrPy   bool    `json:"ACCR_PY" field_type:"L"`   // Posten vorig boekjaar
	AcctNr   string  `json:"ACCT_NR" field_type:"C"`   // Rekeningnr
	AdmCode  string  `json:"ADM_CODE" field_type:"C"`  // Admin.code
	Amount   float64 `json:"AMOUNT" field_type:"N"`    // Mutatiebedrag
	BankDesc string  `json:"BANK_DESC" field_type:"M"` // Omschrijving bankmutatie
	BaseAmt  float64 `json:"BASE_AMT" field_type:"Y"`  // Basisbedrag BTW
	BegPer   float64 `json:"BEG_PER" field_type:"N"`   // Begin periode
	CnID     string  `json:"CN_ID" field_type:"C"`     // Unieke consolidatiecode
	CostCode string  `json:"COST_CODE" field_type:"C"` // Kostenplaatscode
	CredDays float64 `json:"CRED_DAYS" field_type:"N"` // Betalingstermijn in dagen
	CurAmt   float64 `json:"CUR_AMT" field_type:"Y"`   // Valutabedrag
	CurCode  string  `json:"CUR_CODE" field_type:"C"`  // Valutacode
	CurVAT   float64 `json:"CUR_VAT" field_type:"Y"`   // BTW-bedrag vreemde valuta
	DifAmt   float64 `json:"DIF_AMT" field_type:"Y"`   // Verschilbedrag
	DifCur   float64 `json:"DIF_CUR" field_type:"Y"`   // Valutaverschil
	DifType  float64 `json:"DIF_TYPE" field_type:"N"`  // Type verschil
	DistCode string  `json:"DIST_CODE" field_type:"C"` // Verdeelsleutel
	DistQty  float64 `json:"DIST_QTY" field_type:"N"`  // Aantal
	DjlRate  float64 `json:"DJL_RATE" field_type:"B"`  // Valutakoers dagboekregel
	DjCode   string  `json:"DJ_CODE" field_type:"C"`   // Dagboekcode
	EbnkStat int     `json:"EBNK_STAT" field_type:"I"` // Status bankkoppeling
	EndPer   float64 `json:"END_PER" field_type:"N"`   // Eind periode
	ExchRate float64 `json:"EXCH_RATE" field_type:"B"` // Valutakoers
	ExpID    string  `json:"EXP_ID" field_type:"C"`    // Export-ID
	GrekType string  `json:"GREK_TYPE" field_type:"C"` // Transactietype G-rekening
	HomeDif  float64 `json:"HOME_DIF" field_type:"Y"`  // Verschil in adm.valuta
	InvNr    string  `json:"INV_NR" field_type:"C"`    // Document-/factuurnummer
	ObliNr   string  `json:"OBLI_NR" field_type:"C"`   // Verplichtingnummer
	PageNr   string  `json:"PAGE_NR" field_type:"C"`   // Bladzijdenummer (teken)
	PageNrv  float64 `json:"PAGE_NRV" field_type:"N"`  // Bladzijdenummer
	PartCode string  `json:"PART_CODE" field_type:"C"` // Vennoot
	PcCode   string  `json:"PC_CODE" field_type:"C"`   // Kostensoortcode
	PdifVAT  float64 `json:"PDIF_VAT" field_type:"Y"`  // BTW-bedrag betalingskorting
	ProjCode string  `json:"PROJ_CODE" field_type:"C"` // Projectcode
	RaAmt    float64 `json:"RA_AMT" field_type:"Y"`    // Bedrag cheque-begeleiding
	RaDisc   float64 `json:"RA_DISC" field_type:"Y"`   // Kortingsbedrag cheque-begeleiding
	RecID    string  `json:"REC_ID" field_type:"C"`    // Unieke code
	RecOrd   int     `json:"REC_ORD" field_type:"I"`   // Regelvolgorde
	SrcID    string  `json:"SRC_ID" field_type:"C"`    // Identificatie bron
	SrcType  float64 `json:"SRC_TYPE" field_type:"N"`  // Brontype
	SubNr    string  `json:"SUB_NR" field_type:"C"`    // Debiteur-/crediteurnummer
	TrnDate  Date    `json:"TRN_DATE" field_type:"D"`  // Mutatiedatum
	TrnDesc  string  `json:"TRN_DESC" field_type:"C"`  // Omschrijving regel

	ExtKey   int `json:"EXT_KEY"`
	DjPageID int `json:"DJ_PAGE_Id"`
}

func (djLine DjLine) BusinessObject() string {
	return ""
}

func (djLine DjLine) Table() string {
	return "DJ_LINE"
}

func (djLine DjLine) Fields() []string {
	return []string{
		"AcctNr",
		"Amount",
		"RecID",
	}
}

func (djLine DjLine) Values() ([]interface{}, error) {
	return FieldsToValues(djLine, djLine.Fields())
}
