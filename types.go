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

type Ledger struct {
	AcctAlloc string   `json:"ACCT_ALLOC"` // Saldo leegboeken naar
	AcctBal   float64  `json:"ACCT_BAL"`   // Rekeningsaldo
	AcctCost  string   `json:"ACCT_COST"`  // Kostprijsrekening
	AcctLink  float64  `json:"ACCT_LINK"`  // Type verzamelrekening
	AcctName  string   `json:"ACCT_NAME"`  // Omschrijving rekening
	AcctNr    string   `json:"ACCT_NR"`    // Rekeningnr
	AcctProj  bool     `json:"ACCT_PROJ"`  // Projectrekening
	AcctType  float64  `json:"ACCT_TYPE"`  // Rekeningtype
	ActPy     float64  `json:"ACT_PY"`     // Saldo vorig jaar
	BalBal    float64  `json:"BAL_BAL"`    // Balans
	BalPl     float64  `json:"BAL_PL"`     // V&W
	CltBlk    bool     `json:"CLT_BLK"`    // Geblokkeerd voor cliënt
	CngDate   DateTime `json:"CNG_DATE"`   // Wijzigingsdatum
	CngNr     int      `json:"CNG_NR"`     // Wijzignummer
	CngUsr    string   `json:"CNG_USR"`    // Medewerker laatste wijziging
	ConAcct   string   `json:"CON_ACCT"`   // Consolidatierekening statutair
	ConAcct2  string   `json:"CON_ACCT2"`  // Consolidatierekening intern
	CorBal    float64  `json:"COR_BAL"`    // Saldo VJP's
	CorCamt   float64  `json:"COR_CAMT"`   // Creditbedrag correcties
	CorDamt   float64  `json:"COR_DAMT"`   // Debetbedrag correcties
	CostCode  string   `json:"COST_CODE"`  // Kostenplaatscode
	CostMand  bool     `json:"COST_MAND"`  // Kostenplaatscode verplicht
	CoAcct    bool     `json:"CO_ACCT"`    // Tussenrekening
	CrAmt     float64  `json:"CR_AMT"`     // Creditbedrag
	CrAmtmin  float64  `json:"CR_AMTMIN"`  // Negatief creditbedrag
	CurCamt   float64  `json:"CUR_CAMT"`   // Valutabedrag credit
	CurCode   string   `json:"CUR_CODE"`   // Valutacode
	CurDamt   float64  `json:"CUR_DAMT"`   // Valutabedrag debet
	DistCode  string   `json:"DIST_CODE"`  // Verdeelsleutel
	DossDesc  string   `json:"DOSS_DESC"`  // Notitie
	DossRef   string   `json:"DOSS_REF"`   // Dossierreferentie
	DrAmt     float64  `json:"DR_AMT"`     // Debetbedrag
	EfsChk    string   `json:"EFS_CHK"`    // Element van te controleren taxonomie
	EfsLedg   string   `json:"EFS_LEDG"`   // Element van administratietaxonomie
	FaAcct    bool     `json:"FA_ACCT"`    // Activumrekening
	FaSell    bool     `json:"FA_SELL"`    // Activumverk.rek
	InpDate   DateTime `json:"INP_DATE"`   // Invoerdatum
	InpUsr    string   `json:"INP_USR"`    // Invoermedewerker
	JrnlQty   bool     `json:"JRNL_QTY"`   // Aantal journaalregels controleren
	LdgStat   string   `json:"LDG_STAT"`   // Rekeningstatus
	NormStat  float64  `json:"NORM_STAT"`  // Normaal saldo
	NotDash   bool     `json:"NOT_DASH"`   // Niet opnemen in dashboardgrafieken
	PartMand  float64  `json:"PART_MAND"`  // Vennoot verplicht
	PcCode    string   `json:"PC_CODE"`    // Kostensoortcode
	PreCor    float64  `json:"PRE_COR"`    // Voor VJP's
	PreCorc   float64  `json:"PRE_CORC"`   // Voor VJP's credit
	PreCord   float64  `json:"PRE_CORD"`   // Voor VJP's debet
	ProjExp   bool     `json:"PROJ_EXP"`   // Projectkosten/-opbrengsten declarabel
	ProjObli  bool     `json:"PROJ_OBLI"`  // Projectverplichtingen
	ProAttr   string   `json:"PRO_ATTR"`   // Gekoppelde projectdimensies
	QtyMand   float64  `json:"QTY_MAND"`   // Aantal verplicht
	RecBlk    bool     `json:"REC_BLK"`    // Geblokkeerd
	RecSys    bool     `json:"REC_SYS"`    // Systeemrecord
	RevAcct   bool     `json:"REV_ACCT"`   // Valuta herwaarderen
	Rubrieken string   `json:"RUBRIEKEN"`  // Rubrieken
	SlCode1   string   `json:"SL_CODE1"`   // Subadministratie 1
	SlCode2   string   `json:"SL_CODE2"`   // Subadministratie 2
	SlCode3   string   `json:"SL_CODE3"`   // Subadministratie 3
	SplitPa   bool     `json:"SPLIT_PA"`   // Per vennoot uitsplitsen
	StatOrig  string   `json:"STAT_ORIG"`  // Oorspronkelijke rekeningstatus
	StatSell  bool     `json:"STAT_SELL"`  // Omzetrekening
	SumRpt    bool     `json:"SUM_RPT"`    // Op grootboekkrt.rapp samenvatten
	TransNr   int      `json:"TRANS_NR"`   // Aantal mutaties
	VATCode   string   `json:"VAT_CODE"`   // BTW-code
	VATMand   bool     `json:"VAT_MAND"`   // BTW-code verplicht
	WfAllow   bool     `json:"WF_ALLOW"`   // Rekening geschikt voor gebruik in IFA
	WgtMand   float64  `json:"WGT_MAND"`   // Gewicht verplicht
	WkrTrn    bool     `json:"WKR_TRN"`    // Mutatie ten laste van vrije ruimte WKR
	XbrlSpec  bool     `json:"XBRL_SPEC"`  // XBRL-specificatie

	ExtKey int `json:"EXT_KEY"`
}

type CompHdr struct {
	AiAllowpr bool     `json:"AI_ALLOWPR"` // Facturen via printservice AutoInvoice mogelijk
	AiApi     string   `json:"AI_API"`     // UUID-sleutel onderneming AutoInvoice
	CngDate   DateTime `json:"CNG_DATE"`   // Wijzigingsdatum
	CngNr     int      `json:"CNG_NR"`     // Wijzignummer
	CngUsr    string   `json:"CNG_USR"`    // Medewerker laatste wijziging
	CompDesc  string   `json:"COMP_DESC"`  // Omschrijving onderneming
	CompID    string   `json:"COMP_ID"`    // Ondernemingcode
	CMdmStrt  DateTime `json:"C_MDM_STRT"` // Startmoment synchronisatie met Visma.net
	DimFo     string   `json:"DIM_FO"`     // Dimensies voor Financial Overview
	DjisMsg   string   `json:"DJIS_MSG"`   // Meldingen dagboekberichten
	DjisProcd DateTime `json:"DJIS_PROCD"` // Laatste verwerking Dagboekberichten
	EaVar     float64  `json:"EA_VAR"`     // eAccounting-uitvoering
	EvCkey    string   `json:"EV_CKEY"`    // Klant key (eVerbinding)
	EvCsec    string   `json:"EV_CSEC"`    // Klant geheim (eVerbinding)
	EvEntity  string   `json:"EV_ENTITY"`  // Bedrijfs-ID (eVerbinding)
	EvIntreq  string   `json:"EV_INTREQ"`  // Toegangsverzoekcode
	EvReqid   string   `json:"EV_REQID"`   // E-mailadres (eVerbinding)
	ExpDate   Date     `json:"EXP_DATE"`   // Verloopdatum
	FoID      string   `json:"FO_ID"`      // ID van de administratie in Visma Online
	GocompCng int      `json:"GOCOMP_CNG"` // Wijzignummer onderneming voor AV Go
	InpDate   DateTime `json:"INP_DATE"`   // Invoerdatum
	InpUsr    string   `json:"INP_USR"`    // Invoermedewerker
	IsAutoi   bool     `json:"IS_AUTOI"`   // AutoInvoice-onderneming
	IsConApp  bool     `json:"IS_CON_APP"` // AccountView Contact (alleen-lezen)
	IsEverb   bool     `json:"IS_EVERB"`   // eVerbinding-onderneming
	IsFo      bool     `json:"IS_FO"`      // Financial Overview-onderneming
	IsGo      bool     `json:"IS_GO"`      // AccountView Go
	IsPayl    bool     `json:"IS_PAYL"`    // Betaallinkdienst-onderneming
	IsScan    bool     `json:"IS_SCAN"`    // Scanboeken-onderneming
	IsVdc     bool     `json:"IS_VDC"`     // Beschikbaar in IFA
	KvkNr     string   `json:"KVK_NR"`     // Kamer van Koophandel-nummer
	LicCheck  string   `json:"LIC_CHECK"`  // Controlenummer cliëntlicentie
	LicCode   string   `json:"LIC_CODE"`   // Modulegegevens cliëntlicentie
	LicName   string   `json:"LIC_NAME"`   // Licentienaam cliëntlicentie
	LicSer    string   `json:"LIC_SER"`    // Serienummer cliëntlicentie
	LMdmStrt  DateTime `json:"L_MDM_STRT"` // Laatste correcte synchronisatie met Visma.net
	NmbrsID   string   `json:"NMBRS_ID"`   // Nmbrs-bedrijfscd
	OdpID     int      `json:"ODP_ID"`     // Visma-ID onderneming
	PaylSec   string   `json:"PAYL_SEC"`   // Betaallink-sleutel
	ResMdm    bool     `json:"RES_MDM"`    // Stamgegevens herstellen in MDM bij volgende sync
	SbDeptcd  string   `json:"SB_DEPTCD"`  // Vestigingscode
	SbSbsnr   string   `json:"SB_SBSNR"`   // YSB-nummer
	SbSpecs   string   `json:"SB_SPECS"`   // Specificatie
	SyncEnd   DateTime `json:"SYNC_END"`   // Einde laatste synchronisatie
	SyncMsg   string   `json:"SYNC_MSG"`   // Melding laatste synchronisatie
	SyncStat  int      `json:"SYNC_STAT"`  // Status laatste synchronisatie
	SyncStrt  DateTime `json:"SYNC_STRT"`  // Start laatste synchronisatie
	VATNr     string   `json:"VAT_NR"`     // Omzetbelastingnummer

	ExtKey    int      `json:"EXT_KEY"`
	CompHdrID int      `json:"COMP_HDR_Id"`
	CompLN    []CompLn `json:"COMP_LN"`
}

type CompLn struct {
}
