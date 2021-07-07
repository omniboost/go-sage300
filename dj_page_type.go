package accountviewnet

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type DjPage struct {
	AcctNr    string   `json:"ACCT_NR"`    // Rekeningnr
	AcctPost  bool     `json:"ACCT_POST"`  // Direct op grootboekrek boeken
	AcctUpd   bool     `json:"ACCT_UPD"`   // Wijziging accountant
	ApprSeq   float64  `json:"APPR_SEQ"`   // Approval-volgummer
	ApprState int      `json:"APPR_STATE"` // Workflowstatus
	AuthDate  DateTime `json:"AUTH_DATE"`  // Laatste controledatum
	AuthNr    string   `json:"AUTH_NR"`    // Autorisatienr
	AuthUsr   string   `json:"AUTH_USR"`   // Laatste controlegebruiker
	BlkAmt    float64  `json:"BLK_AMT"`    // Bedrag G-rekening
	BlkCur    float64  `json:"BLK_CUR"`    // Bedrag G-rekening VV
	BvAmt     float64  `json:"BV_AMT"`     // Bedrag UWV
	BvCur     float64  `json:"BV_CUR"`     // Bedrag UWV VV
	CardName  string   `json:"CARD_NAME"`  // Creditcardnaam
	CardNr    string   `json:"CARD_NR"`    // Creditcardnr
	CheckAmt  float64  `json:"CHECK_AMT"`  // Cheque-bedrag
	CheckPrn  bool     `json:"CHECK_PRN"`  // Cheque afgedrukt
	CheckWd   string   `json:"CHECK_WD"`   // Cheque-bedrag voluit
	CltUpd    bool     `json:"CLT_UPD"`    // Wijziging cliënt
	CngDate   DateTime `json:"CNG_DATE"`   // Wijzigingsdatum
	CngNr     int      `json:"CNG_NR"`     // Wijzignummer
	CngUsr    string   `json:"CNG_USR"`    // Medewerker laatste wijziging
	Comment1  string   `json:"COMMENT1"`   // Notitie
	CostCode  string   `json:"COST_CODE"`  // Kostenplaatscode
	CredDays  float64  `json:"CRED_DAYS"`  // Betalingstermijn in dagen
	CurAmt    float64  `json:"CUR_AMT"`    // Totaal valuta
	CurCode   string   `json:"CUR_CODE"`   // Valutacode
	CurVAT    float64  `json:"CUR_VAT"`    // BTW-valutatotaal
	DdAct     bool     `json:"DD_ACT"`     // Automatische incasso
	DiscCode  string   `json:"DISC_CODE"`  // Betalingsconditie
	DiscDays  float64  `json:"DISC_DAYS"`  // Kortingstermijn in dagen
	DjlRate   float64  `json:"DJL_RATE"`   // Valutakoers vreemde valutadagboek
	DjCode    string   `json:"DJ_CODE"`    // Dagboekcode
	DocID     string   `json:"DOC_ID"`     // Documentcode
	DocType   int      `json:"DOC_TYPE"`   // Documenttype
	EaVoucher string   `json:"EA_VOUCHER"` // ID van het boekstuk in eAccounting
	EbankID   string   `json:"EBANK_ID"`   // Bankkoppeling-identificatie
	EbankStat int      `json:"EBANK_STAT"` // Status bankkoppeling
	ExchRate  float64  `json:"EXCH_RATE"`  // Valutakoers
	ExpCli    bool     `json:"EXP_CLI"`    // Export naar cliënt
	ExpDate   Date     `json:"EXP_DATE"`   // Geldig tot
	ExpDjp    bool     `json:"EXP_DJP"`    // Geëxporteerde dagboekbladzijden
	ExpSys    bool     `json:"EXP_SYS"`    // Oude waarde van sys_gen
	ExtDesc   string   `json:"EXT_DESC"`   // Betalingsreferentie
	FirstWf   string   `json:"FIRST_WF"`   // Eerste goedkeurder in de workflow
	GrossFlg  bool     `json:"GROSS_FLG"`  // Basisbedrag incl BTW weergeven
	HdrDesc   string   `json:"HDR_DESC"`   // Omschrijving bladzijde
	ImpCode   string   `json:"IMP_CODE"`   // Dagboekbericht
	InpDate   DateTime `json:"INP_DATE"`   // Invoerdatum
	InpUsr    string   `json:"INP_USR"`    // Invoermedewerker
	InvNr     string   `json:"INV_NR"`     // Document-/factuurnummer
	LbAmt     float64  `json:"LB_AMT"`     // Bedrag loonbelasting
	LbCur     float64  `json:"LB_CUR"`     // Bedrag loonbelasting VV
	MandB2b   bool     `json:"MAND_B2B"`   // Incassovorm: zakelijk
	MandDate  Date     `json:"MAND_DATE"`  // Datum ondertekening
	MandID    string   `json:"MAND_ID"`    // Kenmerk incassomachtiging
	MandSeq   int      `json:"MAND_SEQ"`   // Incassomachtiging
	NettAmt   float64  `json:"NETT_AMT"`   // Netto bladsaldo
	NoPost    bool     `json:"NO_POST"`    // Bladzijde niet boeken
	OwnCngDt  DateTime `json:"OWN_CNG_DT"` // Wijzigingsdatum eigenaar
	PageAmt   float64  `json:"PAGE_AMT"`   // Bladsaldo
	PageAmtt  float64  `json:"PAGE_AMTT"`  // Eindsaldo
	PageEamt  float64  `json:"PAGE_EAMT"`  // Eindsaldo
	PageNr    string   `json:"PAGE_NR"`    // Bladzijdenummer (teken)
	PageNrd   string   `json:"PAGE_NRD"`   // Bladzijdenummer (omschrijving)
	PageNrv   float64  `json:"PAGE_NRV"`   // Bladzijdenummer
	PageVAT   float64  `json:"PAGE_VAT"`   // BTW-bedrag bladzijde
	PayinID   string   `json:"PAYIN_ID"`   // Pay-in ID
	PayID     string   `json:"PAY_ID"`     // Betalingscode
	PayType   string   `json:"PAY_TYPE"`   // Betaalwijze
	PdAmt     float64  `json:"PD_AMT"`     // Bedrag betalingskorting
	PdPct     float64  `json:"PD_PCT"`     // Percentage betalingskorting
	Period    float64  `json:"PERIOD"`     // Periode
	PrevAmt   float64  `json:"PREV_AMT"`   // Beginsaldo bladzijde
	PrnCheck  bool     `json:"PRN_CHECK"`  // Cheque afdrukken
	PrnRa     bool     `json:"PRN_RA"`     // Cheque-begeleiding afdrukken
	ProjCode  string   `json:"PROJ_CODE"`  // Projectcode
	RecBlk    bool     `json:"REC_BLK"`    // Geblokkeerd
	StmtDate  Date     `json:"STMT_DATE"`  // Datum afschrift
	StmtLine  bool     `json:"STMT_LINE"`  // Bankafschriftregel
	StmtNr    int      `json:"STMT_NR"`    // Bankafschrift
	SubNr     string   `json:"SUB_NR"`     // Debiteur/Crediteur
	SysGen    bool     `json:"SYS_GEN"`    // Bladzijde gegenereerd door systeem
	SysVal    bool     `json:"SYS_VAL"`    // Systeemvalidaties
	TgtBal    float64  `json:"TGT_BAL"`    // Controlesaldo
	TrnDate   Date     `json:"TRN_DATE"`   // Mutatiedatum
	UsrName   string   `json:"USR_NAME"`   // Laatst gewijzigd door
	VexchRate float64  `json:"VEXCH_RATE"` // Valutakoers (import)

	ExtKey   int `json:"EXT_KEY"`
	DjPageID int `json:"DJ_PAGE_Id"`
}

func (djPage DjPage) BusinessObject() string {
	return "DJ2"
}

func (djPage DjPage) Table() string {
	return "DJ_PAGE"
}

func (djPage DjPage) Fields() []string {
	return []string{
		"PAGE_NR",
		"DJ_CODE",
		"HDR_DESC",
	}
}

func (djPage DjPage) Values() ([]interface{}, error) {
	fields := djPage.Fields()
	values := make([]interface{}, len(fields))
	b, err := json.Marshal(djPage)
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

func (djPage DjPage) ToAccountviewDataPostRequest(client *Client, lines []DjLine) (AccountviewDataPostRequest, error) {
	children := make([]BusinessObjectInterface, len(lines))
	for i, v := range lines {
		children[i] = BusinessObjectInterface(v)
	}
	return BusinessObjectToAccountviewDataPostRequest(client, djPage, children)
}
