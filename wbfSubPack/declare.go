package wbfSubPack

//-----------------------------------------
const wSep = "§";                         // used to separe word in a list 
const endOfLine = ";;\n"
var lemma_para_list = make([]paraStruct, 0, 0 )   	
var lemmaNotFoundList = make([]string,0,100)
//----------------------
var level_other = "Oth"
var list_level_str = " " + level_other + " A0 A1 A2 B1 B2 C1 C2 "
var list_level          = make([]string,0, 0) 
var only_level_numWords = make([]int,   0, 0)  
var perc_level          = make([]int,   0, 0)  
//-------------------------------------------
var separRow   = "[\r\n.;:?!]";
//var separRowFalse = "[\r\n]";
var separWord = "["     +
			"\r\n.;?!" + 
			"\t"       +
			" "        + 
			",:"       +
            "|"        +
			"¿"        +	
			"°"        + 
			"¡"        + 
			"\\"       + 
			"_"        +
			"\\+"      +
			"\\*"      +
			"()<>"     + 
			"\\]\\["   +
			"`{}«»‘’‚‛“””„‟" +		
			"\""       + 
			"'"        + 	
			"\\/"      +  
			wSep       + 
			"]" ; 	
//-------------------------------------
var fseq = "z§" ; 
var lenFseq = len(fseq) 
//--------------------------
//var apiceInverso = `40`  //  in windows:  tasto Alt + 96 (tastierino numerico)
var lastPerc=0; 
var prevPerc = -1; 
var sw_list_Word_if_in_ExtrRow = false    // se true, lista soltanto le parole che si trovano su righe estratte  
//-----------------
//---------------------------------------------

var last_written_dict_rowFile string = "";         //  filename of the last written dictionary row file   

//------------------------------
const BASE_rowGrNum = 100000;  // centomila,  1 e 5 zeri


var sw_rewrite_wordLemma_dict bool = true

//---
var sw_write_numbered_text  bool = false
var fname_out_numbered_text string = ""  
var fPrefix_out_numbered_text string = ""  
var outNumberBegin       int  = 0    //   write numbered text beginning with this number    
//---
var sw_read_numbered_text   bool = false
var fname_inp_numbered_text string = ""  
var fPrefix_inp_numbered_text string = ""  
//-----------------------

//var outWordLemmaDict_file = "newWordLemmaDict.txt";
//var outLemmaTranDict_file = "newLemmaTranDict.txt"
//var outLemmaNotFound      = "newLemmaNotFound.txt"
//var words_to_learn_file   = "ixWordsToLearn.txt"    

//var outRowTran_file  = "newRowDict.txt"
var sw_ignore_newLine bool = false
var sw_nl_only        bool = false
//------------------------------------

const LAST_WORD = "\xff\xff\xff"

//------
var parameter_path_html string  = "WBF_html_jsw"
//------------------



var separRowList = make([]string,0,0)		


var msgLevelStat = "" 

var sw_PRINT_TIME bool = false;                 // in caso di durata abnorme, usa true per vedere dove impiega più tempo

var wordSliceAlpha = make([]wordStruct, 0, 0)  
var wordSliceFreq  = make([]wordStruct, 0, 0)  


var uniqueWordByFreq  []wordIxStruct;   // elenco delle parole in ordine di frequenza
var uniqueWordByAlpha []wordIxStruct;   // elenco delle parole in ordine alphabetico


var dictLemmaTran []lemmaTranStruct ;  // dictionary lemma translation  


var wordStatistic_tx =  make( []statStruct, 101, 101);	
var inputTextRowSlice       []rowStruct;
var isUsedArray    []bool

//var rowLineIxList     []rowIxStruct

var lista_gruppiSelectRow  = make([]  rowGroupStruct, 0, 100)  

var countNumLines  bool = false 
var maxNumLinesToWrite = 0
var lemmaFormat   = "word lemma"
var sw_lemma_word = false 

//var rowArrayByFreq []rowStruct;
//-------------------------------

var dictionaryWord = make([]wDictStruct,0,0);  
var numberOfDictLines =0;  
//var dictionaryRow  = make([]rDictStruct,0,0);  
var numberOfRowDictLines =0;  
var prevRunLanguage = ""; 
//var prevRunListFile = "" 
//-----------

var numberOfRows  = 0; 
var numberOfWords = 0; 
var numberERRORS=0
var numberOfUniqueWords = 0; 
var showReadFile string = ""
//--------------------------------------------------------
//var scrX, scrY int = getScreenXY();


var sw_begin_ended = false     
var sw_HTML_ready  = false     
//--------------------------------

//---------------------------------

var last_rG_ixSelGrOption   	int = 0 
var last_rG_group           	string = ""  
var last_rG_firstIxRowOfGr  	int = 0     
var last_rG_lastIxRowOfGr   	int = 0       
var	last_ixRowBeg           	int = 0	 
var	last_ixRowEnd           	int = 0	

var last_html_rowGroup_index_gr int = 0
var last_html_rowGroup_beginNum int = 0 
var last_html_rowGroup_numRows  int = 0
var last_word_fromWord          int = 0 
var last_word_numWords          int = 0
var last_sel_extrRow            string = "" 	
//--------------------------
var separPrefList = make([]separPrefStruct, 0, 200) 
//--------------------------
	
var last_run_extrRow = ""	
//---------------