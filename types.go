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

type Transact struct {
	AcctNr    string  `json:"ACCT_NR"`    // Rekeningnr
	ActCost   float64 `json:"ACT_COST"`   // Werkelijke kosten
	ActProc   float64 `json:"ACT_PROC"`   // Werkelijke opbrengsten
	AdmCode   string  `json:"ADM_CODE"`   // Admin.code
	Amount    float64 `json:"AMOUNT"`     // Mutatiebedrag
	BaseAmt   float64 `json:"BASE_AMT"`   // Basisbedrag BTW
	ChkID     string  `json:"CHK_ID"`     // Afletternummer
	CnID      string  `json:"CN_ID"`      // Unieke consolidatiecode
	CogsAmt   float64 `json:"COGS_AMT"`   // Kostprijs
	CogsQty   float64 `json:"COGS_QTY"`   // Kostprijs aantal
	CostCode  string  `json:"COST_CODE"`  // Kostenplaatscode
	CstAmt    float64 `json:"CST_AMT"`    // Bedr geboekt
	CurAmt    float64 `json:"CUR_AMT"`    // Valutabedrag
	CurCode   string  `json:"CUR_CODE"`   // Valutacode
	DeclAmt   float64 `json:"DECL_AMT"`   // Bedrag declarabel
	DistCode  string  `json:"DIST_CODE"`  // Verdeelsleutel
	DistQty   float64 `json:"DIST_QTY"`   // Aantal
	DjCode    string  `json:"DJ_CODE"`    // Dagboekcode
	ExchRate  float64 `json:"EXCH_RATE"`  // Valutakoers
	GrekType  string  `json:"GREK_TYPE"`  // Transactietype G-rekening
	HasDoc    bool    `json:"HAS_DOC"`    // Documentkoppeling aanwezig
	InvDesc   string  `json:"INV_DESC"`   // Factuuromschrijving transactie
	InvNr     string  `json:"INV_NR"`     // Document-/factuurnummer
	ItemDesc  string  `json:"ITEM_DESC"`  // Omschrijving transactie
	ItemQty   float64 `json:"ITEM_QTY"`   // Aantal (agr)
	LinkID    string  `json:"LINK_ID"`    // Koppeling projectjournaal
	MndPer    int     `json:"MND_PER"`    // Maanden
	ObliNr    string  `json:"OBLI_NR"`    // Verplichtingnummer
	PageNr    string  `json:"PAGE_NR"`    // Bladzijdenummer (teken)
	PageNrv   float64 `json:"PAGE_NRV"`   // Bladzijdenummer
	PartCode  string  `json:"PART_CODE"`  // Vennoot
	PartRng   string  `json:"PART_RNG"`   // Vennoot als dimensie
	PaCode    string  `json:"PA_CODE"`    // Vennoot als dimensie
	PcCode    string  `json:"PC_CODE"`    // Kostensoortcode
	PcDim     string  `json:"PC_DIM"`     // Kostensoort als dimensie
	Period    int     `json:"PERIOD"`     // Periode
	PerCalc   int     `json:"PER_CALC"`   // Periode berekend
	PerProj   float64 `json:"PER_PROJ"`   // Projectperiode
	ProjCode  string  `json:"PROJ_CODE"`  // Projectcode
	QtrPer    int     `json:"QTR_PER"`    // Kwartalen
	RecID     string  `json:"REC_ID"`     // Unieke code
	ResID     string  `json:"RES_ID"`     // Bronidentificatie
	RevAmt    float64 `json:"REV_AMT"`    // Omgedraaid bedrag
	RngCode   string  `json:"RNG_CODE"`   // Vennotenadministratie
	SellAmt   float64 `json:"SELL_AMT"`   // Bedrag te factureren
	SellQty   float64 `json:"SELL_QTY"`   // Aant te fact
	SrcID     string  `json:"SRC_ID"`     // Identificatie bron
	SrcType   float64 `json:"SRC_TYPE"`   // Brontype
	SubNr     string  `json:"SUB_NR"`     // Debiteur-/crediteurnummer
	TrnDate   Date    `json:"TRN_DATE"`   // Mutatiedatum
	TrnDesc   string  `json:"TRN_DESC"`   // Omschrijving mutatie
	TrnQty    float64 `json:"TRN_QTY"`    // Aantal (agr)
	TrnWgt    float64 `json:"TRN_WGT"`    // Gewicht
	VATAmt    float64 `json:"VAT_AMT"`    // BTW-bedrag
	VATCode   string  `json:"VAT_CODE"`   // BTW-code
	VoucherNr string  `json:"VOUCHER_NR"` // Boekstuknummer
	WkrTrn    bool    `json:"WKR_TRN"`    // Mutatie ten laste van vrije ruimte WKR
	XmlData   string  `json:"XML_DATA"`   // Gegevens in XML-formaat

	ExtKey     int `json:"EXT_KEY"`
	TransactID int `json:"TRANSACT_Id"`
}

type Daily struct {
	AcctPuw  string   `json:"ACCT_PUW"`  // Rekening betalingen onderweg
	AcctRuw  string   `json:"ACCT_RUW"`  // Rekening ontvangsten onderweg
	BnkAcct  string   `json:"BNK_ACCT"`  // Bankrekening dagboek
	BnkAcctv string   `json:"BNK_ACCTV"` // Bankrekening dagboek
	BnkCode  string   `json:"BNK_CODE"`  // Bankcode
	CheckNr  string   `json:"CHECK_NR"`  // Laatst gebruikte cheque-nummer
	ChkBal   bool     `json:"CHK_BAL"`   // Bladzijdesaldi controleren
	ChkDif   bool     `json:"CHK_DIF"`   // Verschillen controleren
	CltBlk   bool     `json:"CLT_BLK"`   // Geblokkeerd voor cliënt
	CngBlk   bool     `json:"CNG_BLK"`   // Handmatige wijzigingen geblokkeerd
	CngDate  DateTime `json:"CNG_DATE"`  // Wijzigingsdatum
	CngNr    int      `json:"CNG_NR"`    // Wijzignummer
	CngUsr   string   `json:"CNG_USR"`   // Medewerker laatste wijziging
	CntCode  string   `json:"CNT_CODE"`  // ISO-landcode
	CredLim  float64  `json:"CRED_LIM"`  // Kredietlimiet bankrekening
	CurCode  string   `json:"CUR_CODE"`  // Valutacode
	DjpBlk   bool     `json:"DJP_BLK"`   // Bladzijde automatisch blokkeren
	DjAcct   string   `json:"DJ_ACCT"`   // Gekoppelde rekening
	DjCode   string   `json:"DJ_CODE"`   // Dagboekcode
	DjComb   bool     `json:"DJ_COMB"`   // Verdicht boeken
	DjExt    bool     `json:"DJ_EXT"`    // Extern dagboek
	DjName   string   `json:"DJ_NAME"`   // Dagboeknaam
	DjOwner  float64  `json:"DJ_OWNER"`  // Dagboekeigenaar
	DjRecon  bool     `json:"DJ_RECON"`  // Bankaflettering
	DjType   float64  `json:"DJ_TYPE"`   // Type dagboek (numeriek)
	DjTypev  string   `json:"DJ_TYPEV"`  // Type dagboek
	ExchStat float64  `json:"EXCH_STAT"` // Uitwisselingsstatus
	GrekBank bool     `json:"GREK_BANK"` // Geblokkeerde bankrekening
	GrossFlg bool     `json:"GROSS_FLG"` // Bruto-invoer
	HdrDesc  int      `json:"HDR_DESC"`  // Omschrijving dagboekbladzijde
	IbanNr   string   `json:"IBAN_NR"`   // IBAN
	InpDate  DateTime `json:"INP_DATE"`  // Invoerdatum
	InpUsr   string   `json:"INP_USR"`   // Invoermedewerker
	LayCheck string   `json:"LAY_CHECK"` // Layout cheque
	LayRemit string   `json:"LAY_REMIT"` // Layout cheque-begeleiding
	LineDesc int      `json:"LINE_DESC"` // Omschrijving dagboekregels
	LstStmt  string   `json:"LST_STMT"`  // Laatst ingelezen bankafschrift
	NumAuto  bool     `json:"NUM_AUTO"`  // Document-/factuurnummers automatisch uitdelen
	NumGrp   float64  `json:"NUM_GRP"`   // Document-/factuurnummergroep
	PaCode   string   `json:"PA_CODE"`   // Vennoot
	PdDesc   int      `json:"PD_DESC"`   // Betalingsreferentie
	SendAppr bool     `json:"SEND_APPR"` // Nieuwe bladzijden verplicht in de workflow sturen
	SysVal   bool     `json:"SYS_VAL"`   // Systeemvalidaties

	ExtKey int `json:"EXT_KEY"`
}

type AcctRec struct {
	AcctNr string `json:"ACCT_NR"` // Verzamelrekening
	SubNr  string `json:"SUB_NR"`  // Debiteurnummer

	ExtKey int `json:"EXT_KEY"`
}

type Contact struct {
	AcctBal   float64  `json:"ACCT_BAL"`   // Saldo
	AcctBal2  float64  `json:"ACCT_BAL2"`  // Saldo te betalen
	AcctLink  int      `json:"ACCT_LINK"`  // Type verzamelrekening
	AcctName  string   `json:"ACCT_NAME"`  // Bedrijfsnaam
	AcctNr    string   `json:"ACCT_NR"`    // Verzamelrekening
	AcctSnd   string   `json:"ACCT_SND"`   // Klank bedrijfsnaam
	Address1  string   `json:"ADDRESS1"`   // Adres
	AddrLine1 string   `json:"ADDR_LINE1"` // Adresregel 1
	AddrLine2 string   `json:"ADDR_LINE2"` // Adresregel 2
	AdminCont string   `json:"ADMIN_CONT"` // Administratiecontactpersoon
	AdminName string   `json:"ADMIN_NAME"` // Administratiecontactpersoon
	AdminPp   string   `json:"ADMIN_PP"`   // Contactpersooncode
	AdmDear   string   `json:"ADM_DEAR"`   // Aanhef administratiecontactpersoon
	AieiMand  bool     `json:"AIEI_MAND"`  // Altijd system-to-system
	AiemMand  bool     `json:"AIEM_MAND"`  // Altijd via e-mail AutoInvoice
	AiprMand  bool     `json:"AIPR_MAND"`  // Altijd printservice AutoInvoice
	AiNowf    bool     `json:"AI_NOWF"`    // Ink.fact boeken in dagb 'Ink.f, geen workflow'
	AmtDd     float64  `json:"AMT_DD"`     // Geselecteerd bedrag te ontvangen
	AmtDdmax  float64  `json:"AMT_DDMAX"`  // Maximumincassobedrag
	AmtPay    float64  `json:"AMT_PAY"`    // Geselecteerd bedrag te betalen
	AmtRec    float64  `json:"AMT_REC"`    // Te ontvangen
	AmtRest   float64  `json:"AMT_REST"`   // Restbedrag
	AmtRest2  float64  `json:"AMT_REST2"`  // Restbedrag te betalen
	ApxList   string   `json:"APX_LIST"`   // Prijslijst art.prijzen
	AuthEmp   string   `json:"AUTH_EMP"`   // Fiatteringsmedewerker
	AutoInv   bool     `json:"AUTO_INV"`   // Facturering via AutoInvoice
	BankAcc2  string   `json:"BANK_ACC2"`  // Bankrekening 2
	BankAcct  string   `json:"BANK_ACCT"`  // Bankrekening 1
	BankAcN1  string   `json:"BANK_AC_N1"` // Bankrekening 1 relatie
	BankAcN2  string   `json:"BANK_AC_N2"` // Bankrekening 2 relatie
	BankSplit bool     `json:"BANK_SPLIT"` // Gecombineerde betaling verdelen
	BlkAcct   string   `json:"BLK_ACCT"`   // Rekeningnummer G-rekening
	BlkBnk    string   `json:"BLK_BNK"`    // Bankcode G-rekening
	BlkDate   Date     `json:"BLK_DATE"`   // Datum bewijs betalingsgedrag onderaannemer
	BlkEinv   bool     `json:"BLK_EINV"`   // G-rekeninginfo in elektronische facturen toestaan
	BlkIban   string   `json:"BLK_IBAN"`   // IBAN G-rekening
	BlkPct    float64  `json:"BLK_PCT"`    // Percentage G-rekening
	BlkSp     bool     `json:"BLK_SP"`     // Geblokkeerd voor orders
	BnkClear  bool     `json:"BNK_CLEAR"`  // Postbankrekening zuiver
	BnkCode1  string   `json:"BNK_CODE1"`  // Bankcode, bankrekening 1
	BnkCode2  string   `json:"BNK_CODE2"`  // Bankcode, bankrekening 2
	BosCode   string   `json:"BOS_CODE"`   // AccountView Server-code
	BoxCity   string   `json:"BOX_CITY"`   // Plaats of plaats postadres
	BvNr      string   `json:"BV_NR"`      // Aansluitnummer UWV
	BvPct     float64  `json:"BV_PCT"`     // Percentage UWV
	BvSub     string   `json:"BV_SUB"`     // Crediteur UWV
	ChnAdress string   `json:"CHN_ADRESS"` // Postadres dealer
	ChnCity   string   `json:"CHN_CITY"`   // Plaats dealer
	ChnCmail  string   `json:"CHN_CMAIL"`  // E-mail hoofdcontactpersoon dealer
	ChnCont   string   `json:"CHN_CONT"`   // Hoofdcontactpersoon dealer
	ChnFax    string   `json:"CHN_FAX"`    // Fax dealer
	ChnMail   string   `json:"CHN_MAIL"`   // E-mail dealer
	ChnMemb   bool     `json:"CHN_MEMB"`   // Relatie is dealer
	ChnName   string   `json:"CHN_NAME"`   // Bedrijfsnaam dealer
	ChnTel    string   `json:"CHN_TEL"`    // Telefoon dealer
	ChnType   string   `json:"CHN_TYPE"`   // Dealerkanaal
	City      string   `json:"CITY"`       // Plaats
	CngDate   DateTime `json:"CNG_DATE"`   // Datum laatste wijz
	CngNr     int      `json:"CNG_NR"`     // Wijzignummer
	CngUsr    string   `json:"CNG_USR"`    // Gewijzigd door
	CntCode   string   `json:"CNT_CODE"`   // ISO-landcode
	CocCode   string   `json:"COC_CODE"`   // KvK-nummer
	Comment1  string   `json:"COMMENT1"`   // Notitie deb/cred
	Country   string   `json:"COUNTRY"`    // Land
	CpaCode   string   `json:"CPA_CODE"`   // Accountant
	CredBal   float64  `json:"CRED_BAL"`   // Kredietruimte
	CredLim   int      `json:"CRED_LIM"`   // Kredietlimiet
	CredShow  bool     `json:"CRED_SHOW"`  // Ook als crediteur tonen
	CrAmt     float64  `json:"CR_AMT"`     // Creditbedrag
	CurCode   string   `json:"CUR_CODE"`   // Valutacode
	CurSign   string   `json:"CUR_SIGN"`   // Valutasymbool
	DdAct     bool     `json:"DD_ACT"`     // Automatische incasso toegestaan
	DebNr     string   `json:"DEB_NR"`     // Extern debiteurnr
	DebNrv    string   `json:"DEB_NRV"`    // Extern debiteur-/crediteurnummer
	DebShow   bool     `json:"DEB_SHOW"`   // Ook als debiteur tonen
	DiscCode  string   `json:"DISC_CODE"`  // Code betalingsconditie
	DiscPct   float64  `json:"DISC_PCT"`   // Factuurkorting
	DjipAcor  bool     `json:"DJIP_ACOR"`  // Boekingsvoorstellen automatisch accepteren
	DjipAMax  float64  `json:"DJIP_A_MAX"` // Maximumacceptatiebedrag
	DjipAUm   bool     `json:"DJIP_A_UM"`  // Maximumacceptatiebedrag gebruiken
	DjipGrp   int      `json:"DJIP_GRP"`   // Voorstelregels groeperen
	DjipToPi  bool     `json:"DJIP_TO_PI"` // Boekingsvoorstellen als inkoopfacturen verwerken
	DjipWnote bool     `json:"DJIP_WNOTE"` // Geen waarschuwing voor notitie in boekingsvoorstel
	DrAmt     float64  `json:"DR_AMT"`     // Debetbedrag
	Eia       string   `json:"EIA"`        // Electronic Invoice Address
	ElCi      bool     `json:"EL_CI"`      // Baliefacturen per e-mail versturen
	ElDp      bool     `json:"EL_DP"`      // Aanbetalingen per e-mail versturen
	ElPo      bool     `json:"EL_PO"`      // Bestelbonnen per e-mail versturen
	ElPs      bool     `json:"EL_PS"`      // Pakbonnen per e-mail versturen
	ElRe      bool     `json:"EL_RE"`      // Aanmaningen per e-mail versturen
	ElSbs     bool     `json:"EL_SBS"`     // Abonnementen per e-mail versturen
	ElSi      bool     `json:"EL_SI"`      // Facturen per e-mail versturen
	ElSo      bool     `json:"EL_SO"`      // Orderbevestigingen per e-mail versturen
	EmpCat    string   `json:"EMP_CAT"`    // Medewerkerklasse
	EmpNr     string   `json:"EMP_NR"`     // Verkoper/inkoper
	FaxBus    string   `json:"FAX_BUS"`    // Fax zakelijk
	FirstWf   string   `json:"FIRST_WF"`   // Eerste goedkeurder in de workflow
	Giro      string   `json:"GIRO"`       // Postbankrekening
	GrayDate  DateTime `json:"GRAY_DATE"`  // Graydon-importdatum
	GrayInfo  string   `json:"GRAY_INFO"`  // Graydon-kredietinformatie
	GrayLim   int      `json:"GRAY_LIM"`   // Graydon-kredietlimiet
	GrayScore float64  `json:"GRAY_SCORE"` // Graydon-score
	HdqCode   string   `json:"HDQ_CODE"`   // Hoofdvestiging
	HdqMemb   bool     `json:"HDQ_MEMB"`   // Nevenvestiging
	IbanNr    string   `json:"IBAN_NR"`    // IBAN 1
	IbanNr2   string   `json:"IBAN_NR2"`   // IBAN 2
	IcatCode  string   `json:"ICAT_CODE"`  // IS categorie goederenstroom
	IcntCode  string   `json:"ICNT_CODE"`  // IS landcode
	ImmWf     bool     `json:"IMM_WF"`     // Facturen direct naar workflow sturen
	InpDate   DateTime `json:"INP_DATE"`   // Invoerdatum
	InpUsr    string   `json:"INP_USR"`    // Invoermedewerker
	InvAddres string   `json:"INV_ADDRES"` // Adres in landafhankelijke indeling (memo)
	InvAnnex  bool     `json:"INV_ANNEX"`  // Factuur en bijlagen samenvoegen
	IshpCode  string   `json:"ISHP_CODE"`  // IS vervoerswijze
	IsxfCode  string   `json:"ISXF_CODE"`  // IS statistische factorcode
	IsysCode  string   `json:"ISYS_CODE"`  // IS stelsel
	IsComp    bool     `json:"IS_COMP"`    // Relatie is mijn concurrent
	IsThird   bool     `json:"IS_THIRD"`   // Relatie is derde partij
	ItypeCodv string   `json:"ITYPE_CODV"` // IS transactiecode (lang)
	ItypCode  string   `json:"ITYP_CODE"`  // IS transactiecode
	KValDate  Date     `json:"K_VAL_DATE"` // Verificatiedatum KvK-nummer
	LbNr      string   `json:"LB_NR"`      // Loonbelastingnummer
	LbPct     float64  `json:"LB_PCT"`     // Percentage loonbelasting
	LbSub     string   `json:"LB_SUB"`     // Crediteur loonbelasting
	LngCode   string   `json:"LNG_CODE"`   // Taalcode
	LnkAcct   string   `json:"LNK_ACCT"`   // Standaardrekening
	MailBus   string   `json:"MAIL_BUS"`   // E-mail zakelijk
	MailOpt   int      `json:"MAIL_OPT"`   // Mailing
	MainCont  string   `json:"MAIN_CONT"`  // Contactpersoon
	MainPp    string   `json:"MAIN_PP"`    // Hoofdcontactpersoon
	MandB2b   bool     `json:"MAND_B2B"`   // Incassovorm: zakelijk
	MandDate  Date     `json:"MAND_DATE"`  // Datum ondertekening
	MandID    string   `json:"MAND_ID"`    // Kenmerk incassomachtiging
	McDear    string   `json:"MC_DEAR"`    // Aanhef hoofdcontactpersoon
	McMailb   string   `json:"MC_MAILB"`   // E-mail hoofdcontactpersoon
	MrnList   string   `json:"MRN_LIST"`   // Prijslijst partnerkorting
	NoAppr    bool     `json:"NO_APPR"`    // Uitsluiten van workflow
	OppCursol string   `json:"OPP_CURSOL"` // Pakketkeuze
	OppSrc    string   `json:"OPP_SRC"`    // Bron
	PayName   string   `json:"PAY_NAME"`   // Korte bedrijfsnaam
	PctList   string   `json:"PCT_LIST"`   // Prijslijst kortingspercentages
	PdPfx     string   `json:"PD_PFX"`     // Betalingskenmerk
	PgBlk     bool     `json:"PG_BLK"`     // Geblokkeerd voor autom betaling
	PolPrio   float64  `json:"POL_PRIO"`   // Uitleverprioriteit
	PostCode  string   `json:"POST_CODE"`  // Postcode
	PoBox     string   `json:"PO_BOX"`     // Postadres
	PoCity    string   `json:"PO_CITY"`    // Plaats postadres
	PoCode    string   `json:"PO_CODE"`    // Postcd postadres
	ProjCred  bool     `json:"PROJ_CRED"`  // Projectcrediteur
	QdCode    string   `json:"QD_CODE"`    // Prijslijstcode
	RecBlk    bool     `json:"REC_BLK"`    // Geblokkeerd
	RemAddr   string   `json:"REM_ADDR"`   // Adres crediteurenadministratie in landafh indeling
	RemBlk    bool     `json:"REM_BLK"`    // Geblokkeerd voor aanmaningen
	RevCat    string   `json:"REV_CAT"`    // Omzetcategorie
	RplChn    string   `json:"RPL_CHN"`    // Dealer van relatie
	RplGrp    string   `json:"RPL_GRP"`    // Relatiegroep
	RplInv    string   `json:"RPL_INV"`    // Factuurdebiteur/crediteur
	RplType   string   `json:"RPL_TYPE"`   // Relatietype
	RplUser   bool     `json:"RPL_USER"`   // Relatie is klant
	SaleAmt   float64  `json:"SALE_AMT"`   // Omzet huidig jaar
	SalePy    float64  `json:"SALE_PY"`    // Omzet vorig jaar
	ScanID    string   `json:"SCAN_ID"`    // Scanb-relatienr
	SgCode    string   `json:"SG_CODE"`    // Versie
	ShipComp  bool     `json:"SHIP_COMP"`  // Geen deelleveringen toegestaan
	ShipMth   string   `json:"SHIP_MTH"`   // Vervoersmethode
	ShowMempo bool     `json:"SHOW_MEMPO"` // Notitie tonen bij toevoegen inkooporder
	ShowMemso bool     `json:"SHOW_MEMSO"` // Notitie tonen bij toevoegen verkooporder
	SicCode   string   `json:"SIC_CODE"`   // Branchecode
	SoAnnex   bool     `json:"SO_ANNEX"`   // Orderbevestiging en bijlagen samenvoegen
	SpPostadr int      `json:"SP_POSTADR"` // Adres aflevering
	SrcCode   string   `json:"SRC_CODE"`   // Zoekcode bedrijf
	SrcPc     string   `json:"SRC_PC"`     // Postcode (snelzoeken)
	StExtra   string   `json:"ST_EXTRA"`   // Adres-toevoeging
	StNr      string   `json:"ST_NR"`      // Huisnummer
	StProp    string   `json:"ST_PROP"`    // Gebouw
	StRegion  string   `json:"ST_REGION"`  // Adres-regio
	StStreet  string   `json:"ST_STREET"`  // Straatnaam
	SubNr     string   `json:"SUB_NR"`     // Debiteur-/crediteurnummer
	SyncState int      `json:"SYNC_STATE"` // Synchronisatiestatus
	TelBus    string   `json:"TEL_BUS"`    // Telefoon zakelijk
	TelMob    string   `json:"TEL_MOB"`    // Telefoon mobiel
	TransNr   int      `json:"TRANS_NR"`   // Aantal mutaties
	VATCode   string   `json:"VAT_CODE"`   // BTW-code
	VATDate   Date     `json:"VAT_DATE"`   // Verificatiedatum
	VATNr     string   `json:"VAT_NR"`     // BTW-nummer
	VisAddres string   `json:"VIS_ADDRES"` // Bezoekadres in landafhankelijke indeling (memo)
	WfAuthemp string   `json:"WF_AUTHEMP"` // IFA-medewerker
	WfDirect  bool     `json:"WF_DIRECT"`  // Direct boeken op kostenrekeningen (IFA)
	WwwUrl    string   `json:"WWW_URL"`    // Website

	ExtKey int `json:"EXT_KEY"`
}

type Cost struct {
	CngDate   DateTime `json:"CNG_DATE"`   // Wijzigingsdatum
	CngNr     int      `json:"CNG_NR"`     // Wijzignummer
	CngUsr    string   `json:"CNG_USR"`    // Medewerker laatste wijziging
	CostCode  string   `json:"COST_CODE"`  // Kostenplaatscode
	CostDesc  string   `json:"COST_DESC"`  // Omschrijving kostenplaats
	CostType  float64  `json:"COST_TYPE"`  // Type kostenplaats
	DifAcct   string   `json:"DIF_ACCT"`   // Verschillenrek
	DistCode  string   `json:"DIST_CODE"`  // Verdeelsleutel
	InpDate   DateTime `json:"INP_DATE"`   // Invoerdatum
	InpUsr    string   `json:"INP_USR"`    // Invoermedewerker
	RepLevel  float64  `json:"REP_LEVEL"`  // Niveau
	WfAuthemp string   `json:"WF_AUTHEMP"` // IFA-medewerker

	ExtKey int `json:"EXT_KEY"`
}

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
