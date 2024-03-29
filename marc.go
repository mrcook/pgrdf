package pgrdf

// MarcRelator representing a MARC Relator code, e.g. `aut`, `edt`, etc.
type MarcRelator string

// All MARC codes for Relators.
// For a full description of each code visit:
// https://www.loc.gov/marc/relators/relaterm.html
const (
	// RoleAut is not found in any current RDF, but is used to distinguish authors
	// in the pgrdf.Ebook, and for marshaling to an RDF as a `dcterms:creator`.
	RoleAut MarcRelator = "aut" // Author

	RoleAbr MarcRelator = "abr" // Abridger
	RoleAcp MarcRelator = "acp" // Art copyist
	RoleAct MarcRelator = "act" // Actor
	RoleAdi MarcRelator = "adi" // Art director
	RoleAdp MarcRelator = "adp" // Adapter
	RoleAft MarcRelator = "aft" // Author of afterword, colophon, etc.
	RoleAnl MarcRelator = "anl" // Analyst
	RoleAnm MarcRelator = "anm" // Animator
	RoleAnn MarcRelator = "ann" // Annotator
	RoleAnt MarcRelator = "ant" // Bibliographic antecedent
	RoleApe MarcRelator = "ape" // Appellee
	RoleApl MarcRelator = "apl" // Appellant
	RoleApp MarcRelator = "app" // Applicant
	RoleAqt MarcRelator = "aqt" // Author in quotations or text abstracts
	RoleArc MarcRelator = "arc" // Architect
	RoleArd MarcRelator = "ard" // Artistic director
	RoleArr MarcRelator = "arr" // Arranger
	RoleArt MarcRelator = "art" // Artist
	RoleAsg MarcRelator = "asg" // Assignee
	RoleAsn MarcRelator = "asn" // Associated name
	RoleAto MarcRelator = "ato" // Autographer
	RoleAtt MarcRelator = "att" // Attributed name
	RoleAuc MarcRelator = "auc" // Auctioneer
	RoleAud MarcRelator = "aud" // Author of dialog
	RoleAui MarcRelator = "aui" // Author of introduction, etc.
	RoleAus MarcRelator = "aus" // Screenwriter
	RoleBdd MarcRelator = "bdd" // Binding designer
	RoleBjd MarcRelator = "bjd" // Bookjacket designer
	RoleBkd MarcRelator = "bkd" // Book designer
	RoleBkp MarcRelator = "bkp" // Book producer
	RoleBlw MarcRelator = "blw" // Blurb writer
	RoleBnd MarcRelator = "bnd" // Binder
	RoleBpd MarcRelator = "bpd" // Bookplate designer
	RoleBrd MarcRelator = "brd" // Broadcaster
	RoleBrl MarcRelator = "brl" // Braille embosser
	RoleBsl MarcRelator = "bsl" // Bookseller
	RoleCas MarcRelator = "cas" // Caster
	RoleCcp MarcRelator = "ccp" // Conceptor
	RoleChr MarcRelator = "chr" // Choreographer
	RoleCli MarcRelator = "cli" // Client
	RoleCll MarcRelator = "cll" // Calligrapher
	RoleClr MarcRelator = "clr" // Colorist
	RoleClt MarcRelator = "clt" // Collotyper
	RoleCmm MarcRelator = "cmm" // Commentator
	RoleCmp MarcRelator = "cmp" // Composer
	RoleCmt MarcRelator = "cmt" // Compositor
	RoleCnd MarcRelator = "cnd" // Conductor
	RoleCng MarcRelator = "cng" // Cinematographer
	RoleCns MarcRelator = "cns" // Censor
	RoleCoe MarcRelator = "coe" // Contestant-appellee
	RoleCol MarcRelator = "col" // Collector
	RoleCom MarcRelator = "com" // Compiler
	RoleCon MarcRelator = "con" // Conservator
	RoleCor MarcRelator = "cor" // Collection registrar
	RoleCos MarcRelator = "cos" // Contestant
	RoleCot MarcRelator = "cot" // Contestant-appellant
	RoleCou MarcRelator = "cou" // Court governed
	RoleCov MarcRelator = "cov" // Cover designer
	RoleCpc MarcRelator = "cpc" // Copyright claimant
	RoleCpe MarcRelator = "cpe" // Complainant-appellee
	RoleCph MarcRelator = "cph" // Copyright holder
	RoleCpl MarcRelator = "cpl" // Complainant
	RoleCpt MarcRelator = "cpt" // Complainant-appellant
	RoleCre MarcRelator = "cre" // Creator
	RoleCrp MarcRelator = "crp" // Correspondent
	RoleCrr MarcRelator = "crr" // Corrector
	RoleCrt MarcRelator = "crt" // Court reporter
	RoleCsl MarcRelator = "csl" // Consultant
	RoleCsp MarcRelator = "csp" // Consultant to a project
	RoleCst MarcRelator = "cst" // Costume designer
	RoleCtb MarcRelator = "ctb" // Contributor
	RoleCte MarcRelator = "cte" // Contestee-appellee
	RoleCtg MarcRelator = "ctg" // Cartographer
	RoleCtr MarcRelator = "ctr" // Contractor
	RoleCts MarcRelator = "cts" // Contestee
	RoleCtt MarcRelator = "ctt" // Contestee-appellant
	RoleCur MarcRelator = "cur" // Curator
	RoleCwt MarcRelator = "cwt" // Commentator for written text
	RoleDbp MarcRelator = "dbp" // Distribution place
	RoleDfd MarcRelator = "dfd" // Defendant
	RoleDfe MarcRelator = "dfe" // Defendant-appellee
	RoleDft MarcRelator = "dft" // Defendant-appellant
	RoleDgc MarcRelator = "dgc" // Degree committee member
	RoleDgg MarcRelator = "dgg" // Degree granting institution
	RoleDgs MarcRelator = "dgs" // Degree supervisor
	RoleDis MarcRelator = "dis" // Dissertant
	RoleDln MarcRelator = "dln" // Delineator
	RoleDnc MarcRelator = "dnc" // Dancer
	RoleDnr MarcRelator = "dnr" // Donor
	RoleDpc MarcRelator = "dpc" // Depicted
	RoleDpt MarcRelator = "dpt" // Depositor
	RoleDrm MarcRelator = "drm" // Draftsman
	RoleDrt MarcRelator = "drt" // Director
	RoleDsr MarcRelator = "dsr" // Designer
	RoleDst MarcRelator = "dst" // Distributor
	RoleDtc MarcRelator = "dtc" // Data contributor
	RoleDte MarcRelator = "dte" // Dedicatee
	RoleDtm MarcRelator = "dtm" // Data manager
	RoleDto MarcRelator = "dto" // Dedicator
	RoleDub MarcRelator = "dub" // Dubious author
	RoleEdc MarcRelator = "edc" // Editor of compilation
	RoleEdm MarcRelator = "edm" // Editor of moving image work
	RoleEdt MarcRelator = "edt" // Editor
	RoleEgr MarcRelator = "egr" // Engraver
	RoleElg MarcRelator = "elg" // Electrician
	RoleElt MarcRelator = "elt" // Electrotyper
	RoleEng MarcRelator = "eng" // Engineer
	RoleEnj MarcRelator = "enj" // Enacting jurisdiction
	RoleEtr MarcRelator = "etr" // Etcher
	RoleEvp MarcRelator = "evp" // Event place
	RoleExp MarcRelator = "exp" // Expert
	RoleFac MarcRelator = "fac" // Facsimilist
	RoleFds MarcRelator = "fds" // Film distributor
	RoleFld MarcRelator = "fld" // Field director
	RoleFlm MarcRelator = "flm" // Film editor
	RoleFmd MarcRelator = "fmd" // Film director
	RoleFmk MarcRelator = "fmk" // Filmmaker
	RoleFmo MarcRelator = "fmo" // Former owner
	RoleFmp MarcRelator = "fmp" // Film producer
	RoleFnd MarcRelator = "fnd" // Funder
	RoleFpy MarcRelator = "fpy" // First party
	RoleFrg MarcRelator = "frg" // Forger
	RoleGis MarcRelator = "gis" // Geographic information specialist
	RoleHis MarcRelator = "his" // Host institution
	RoleHnr MarcRelator = "hnr" // Honoree
	RoleHst MarcRelator = "hst" // Host
	RoleIll MarcRelator = "ill" // Illustrator
	RoleIlu MarcRelator = "ilu" // Illuminator
	RoleIns MarcRelator = "ins" // Inscriber
	RoleInv MarcRelator = "inv" // Inventor
	RoleIsb MarcRelator = "isb" // Issuing body
	RoleItr MarcRelator = "itr" // Instrumentalist
	RoleIve MarcRelator = "ive" // Interviewee
	RoleIvr MarcRelator = "ivr" // Interviewer
	RoleJud MarcRelator = "jud" // Judge
	RoleJug MarcRelator = "jug" // Jurisdiction governed
	RoleLbr MarcRelator = "lbr" // Laboratory
	RoleLbt MarcRelator = "lbt" // Librettist
	RoleLdr MarcRelator = "ldr" // Laboratory director
	RoleLed MarcRelator = "led" // Lead
	RoleLee MarcRelator = "lee" // Libelee-appellee
	RoleLel MarcRelator = "lel" // Libelee
	RoleLen MarcRelator = "len" // Lender
	RoleLet MarcRelator = "let" // Libelee-appellant
	RoleLgd MarcRelator = "lgd" // Lighting designer
	RoleLie MarcRelator = "lie" // Libelant-appellee
	RoleLil MarcRelator = "lil" // Libelant
	RoleLit MarcRelator = "lit" // Libelant-appellant
	RoleLsa MarcRelator = "lsa" // Landscape architect
	RoleLse MarcRelator = "lse" // Licensee
	RoleLso MarcRelator = "lso" // Licensor
	RoleLtg MarcRelator = "ltg" // Lithographer
	RoleLyr MarcRelator = "lyr" // Lyricist
	RoleMcp MarcRelator = "mcp" // Music copyist
	RoleMdc MarcRelator = "mdc" // Metadata contact
	RoleMed MarcRelator = "med" // Medium
	RoleMfp MarcRelator = "mfp" // Manufacture place
	RoleMfr MarcRelator = "mfr" // Manufacturer
	RoleMod MarcRelator = "mod" // Moderator
	RoleMon MarcRelator = "mon" // Monitor
	RoleMrb MarcRelator = "mrb" // Marbler
	RoleMrk MarcRelator = "mrk" // Markup editor
	RoleMsd MarcRelator = "msd" // Musical director
	RoleMte MarcRelator = "mte" // Metal-engraver
	RoleMtk MarcRelator = "mtk" // Minute taker
	RoleMus MarcRelator = "mus" // Musician
	RoleNrt MarcRelator = "nrt" // Narrator
	RoleOpn MarcRelator = "opn" // Opponent
	RoleOrg MarcRelator = "org" // Originator
	RoleOrm MarcRelator = "orm" // Organizer
	RoleOsp MarcRelator = "osp" // Onscreen presenter
	RoleOth MarcRelator = "oth" // Other
	RoleOwn MarcRelator = "own" // Owner
	RolePad MarcRelator = "pad" // Place of address
	RolePan MarcRelator = "pan" // Panelist
	RolePat MarcRelator = "pat" // Patron
	RolePbd MarcRelator = "pbd" // Publishing director
	RolePbl MarcRelator = "pbl" // Publisher
	RolePdr MarcRelator = "pdr" // Project director
	RolePfr MarcRelator = "pfr" // Proofreader
	RolePht MarcRelator = "pht" // Photographer
	RolePlt MarcRelator = "plt" // Platemaker
	RolePma MarcRelator = "pma" // Permitting agency
	RolePmn MarcRelator = "pmn" // Production manager
	RolePop MarcRelator = "pop" // Printer of plates
	RolePpm MarcRelator = "ppm" // Papermaker
	RolePpt MarcRelator = "ppt" // Puppeteer
	RolePra MarcRelator = "pra" // Praeses
	RolePrc MarcRelator = "prc" // Process contact
	RolePrd MarcRelator = "prd" // Production personnel
	RolePre MarcRelator = "pre" // Presenter
	RolePrf MarcRelator = "prf" // Performer
	RolePrg MarcRelator = "prg" // Programmer
	RolePrm MarcRelator = "prm" // Printmaker
	RolePrn MarcRelator = "prn" // Production company
	RolePro MarcRelator = "pro" // Producer
	RolePrp MarcRelator = "prp" // Production place
	RolePrs MarcRelator = "prs" // Production designer
	RolePrt MarcRelator = "prt" // Printer
	RolePrv MarcRelator = "prv" // Provider
	RolePta MarcRelator = "pta" // Patent applicant
	RolePte MarcRelator = "pte" // Plaintiff-appellee
	RolePtf MarcRelator = "ptf" // Plaintiff
	RolePth MarcRelator = "pth" // Patent holder
	RolePtt MarcRelator = "ptt" // Plaintiff-appellant
	RolePup MarcRelator = "pup" // Publication place
	RoleRbr MarcRelator = "rbr" // Rubricator
	RoleRcd MarcRelator = "rcd" // Recordist
	RoleRce MarcRelator = "rce" // Recording engineer
	RoleRcp MarcRelator = "rcp" // Addressee
	RoleRdd MarcRelator = "rdd" // Radio director
	RoleRed MarcRelator = "red" // Redaktor
	RoleRen MarcRelator = "ren" // Renderer
	RoleRes MarcRelator = "res" // Researcher
	RoleRev MarcRelator = "rev" // Reviewer
	RoleRpc MarcRelator = "rpc" // Radio producer
	RoleRps MarcRelator = "rps" // Repository
	RoleRpt MarcRelator = "rpt" // Reporter
	RoleRpy MarcRelator = "rpy" // Responsible party
	RoleRse MarcRelator = "rse" // Respondent-appellee
	RoleRsg MarcRelator = "rsg" // Restager
	RoleRsp MarcRelator = "rsp" // Respondent
	RoleRsr MarcRelator = "rsr" // Restorationist
	RoleRst MarcRelator = "rst" // Respondent-appellant
	RoleRth MarcRelator = "rth" // Research team head
	RoleRtm MarcRelator = "rtm" // Research team member
	RoleSad MarcRelator = "sad" // Scientific advisor
	RoleSce MarcRelator = "sce" // Scenarist
	RoleScl MarcRelator = "scl" // Sculptor
	RoleScr MarcRelator = "scr" // Scribe
	RoleSds MarcRelator = "sds" // Sound designer
	RoleSec MarcRelator = "sec" // Secretary
	RoleSgd MarcRelator = "sgd" // Stage director
	RoleSgn MarcRelator = "sgn" // Signer
	RoleSht MarcRelator = "sht" // Supporting host
	RoleSll MarcRelator = "sll" // Seller
	RoleSng MarcRelator = "sng" // Singer
	RoleSpk MarcRelator = "spk" // Speaker
	RoleSpn MarcRelator = "spn" // Sponsor
	RoleSpy MarcRelator = "spy" // Second party
	RoleSrv MarcRelator = "srv" // Surveyor
	RoleStd MarcRelator = "std" // Set designer
	RoleStg MarcRelator = "stg" // Setting
	RoleStl MarcRelator = "stl" // Storyteller
	RoleStm MarcRelator = "stm" // Stage manager
	RoleStn MarcRelator = "stn" // Standards body
	RoleStr MarcRelator = "str" // Stereotyper
	RoleTcd MarcRelator = "tcd" // Technical director
	RoleTch MarcRelator = "tch" // Teacher
	RoleThs MarcRelator = "ths" // Thesis advisor
	RoleTld MarcRelator = "tld" // Television director
	RoleTlp MarcRelator = "tlp" // Television producer
	RoleTrc MarcRelator = "trc" // Transcriber
	RoleTrl MarcRelator = "trl" // Translator
	RoleTyd MarcRelator = "tyd" // Type designer
	RoleTyg MarcRelator = "tyg" // Typographer
	RoleUvp MarcRelator = "uvp" // University place
	RoleVac MarcRelator = "vac" // Voice actor
	RoleVdg MarcRelator = "vdg" // Videographer
	RoleWac MarcRelator = "wac" // Writer of added commentary
	RoleWal MarcRelator = "wal" // Writer of added lyrics
	RoleWam MarcRelator = "wam" // Writer of accompanying material
	RoleWat MarcRelator = "wat" // Writer of added text
	RoleWdc MarcRelator = "wdc" // Woodcutter
	RoleWde MarcRelator = "wde" // Wood engraver
	RoleWin MarcRelator = "win" // Writer of introduction
	RoleWit MarcRelator = "wit" // Witness
	RoleWpr MarcRelator = "wpr" // Writer of preface
	RoleWst MarcRelator = "wst" // Writer of supplementary textual content

	/*
	 * NOTE: the following have been discontinued and should not be used.
	 */
	RoleClb MarcRelator = "clb" // Collaborator (discontinued)
	RoleGrt MarcRelator = "grt" // Graphic technician (discontinued)
	RoleVoc MarcRelator = "voc" // Vocalist (discontinued)

	// RoleUnk (`unk`) is only found in a few Project Gutenberg RDF documents,
	// but is not listed in the official MARC relators.
	RoleUnk MarcRelator = "unk"
)
