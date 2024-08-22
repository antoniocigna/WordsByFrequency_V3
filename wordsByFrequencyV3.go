/*
The purpose of this application is to help learn the most used words first.
1) I extract all the words of the text and put them in descending order of use ( first the most used words)
2) for each word I list the sentences that contain it, but these can also be thousands, which ones to propose?
3) If I assume that the most used words must be studied firstly, I assume that those with greater frequency have all already been learned 
and those with less frequency are yet to be learned.
4) From what has been hypothesized above, I assume that the most interesting sentences to propose are those in which the number of unknown words is as low as possible
5) Concluding, for each word with frequency F I propose the sentences in increasing order of the number of words the frequency of which is lower than F

=============================
program logic:
1) split the text into rows
2) extract all the words of each row
3) associate to each word the number of lines that contain it (the frequency) and the index of these lines
4) list the words in reverse order of frequency of use (number of lines containing the word)
5) to each line associate the list of frequencies of its words
6) for each word to each line that contains it, associate the number of words with a lower frequency than this one
7) at each request a word to study, list some or all of the sentences that contain it, starting with those with the fewest unknown words (i.e. with frequency < of this word)
*/


package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"runtime"
    "strings"
	"strconv"
	"encoding/hex"
    "regexp"
    "sort"
    "bufio"		
    "io"
    "github.com/zserge/lorca"	
	"github.com/lxn/win"	
)
//--
//  ZSERGE LORCA  vedi C:\Users\cigna\go\pkg\mod\github.com\zserge\lorca@v0.1.10
//--
//----------------------------------

//var apiceInverso = `40`

var lastPerc=0; 
var prevPerc = -1; 
var sw_list_Word_if_in_ExtrRow = false    // se true, lista soltanto le parole che si trovano su righe estratte  

//---------------
var	html_path         string = "WBF_html_jsw"
var file_inputControl string = "inputControl.txt";     // file control   ( file names and switches) 

var main_input_text_file string = ""              // input text file ( this is the real input, it is unnumbered, it's read when read_numbered_text = false)  
 

//------------------------------------------------
var outFile          = "outFileProva.csv" 
//--------------------------------------------
const FOLDER_dictionary string = "myDictionary"

const FOLDER_INPUT          string = "INPUT"
const FOLDER_OUTPUT         string = "OUTPUT"
const FOLDER_INPUT_OUTPUT   string = "INPUT_OUTPUT"
const FOLDER_IO_lastTRAN    string = "INPUT_OUTPUT/lastTRAN"
  
const FOLDER_I_lemma        string = "INPUT/inputLemma"
const FOLDER_I_paradigma    string = "INPUT/inputParadigma"

const FOLDER_O_ARCHIVE        string = "OUTPUT/archive"
const FOLDER_O_arc_TRAN_rows  string = "OUTPUT/archive/arc_TRAN_ROWS"
const FOLDER_O_arc_TRAN_words string = "OUTPUT/archive/arc_TRAN_words"
const FOLDER_O_arc_TO_learn   string = "OUTPUT/archive/arc_TO_learn"

const FOLDER_inpNumText string = "inpNumTEXT"
const FOLDER_outNumText string = "outNumTEXT"
const FOLDER_outTRAN    string = "outTRAN"

//------------------------------------------------
const FILE_last_updated_dict_rows  string = "lastUpdated_dict_tran_rows.txt"	
const FILE_last_updated_dict_words string = "lastUpdated_dict_tran_words.txt" 
const FILE_outWordLemmaDict  string = "newWordLemmaDict.txt";
//const FILE_ outLemmaTranDict  string = "newLemmaTranDict.txt"
const FILE_outLemmaNotFound  string = "newLemmaNotFound.txt"
const FILE_words_to_learn    string = "wordsToLearn.txt"   

const FILE_last_mainpage_values2 string = "last_mainpage_values2.txt" // to transfer values between runs
const FILE_inpParadigma      string = "myParadigma.csv" ; 
const FILE_inpLemma_word_lemma   string = "myWordLemmaFile_fmtWordLemma.txt"  
const FILE_inpLemma_lemma_word   string = "myWordLemmaFile_fmtLemmaWord.txt"  
const FILE_language          string = "language.txt"           // lang and voice file    
//---------------------------------------------

var last_written_dict_rowFile string = "";         //  filename of the last written dictionary row file   

//------------------------------
const BASE_rowGrNum = 100000;  // centomila,  1 e 5 zeri


var sw_rewrite_wordLemma_dict bool = true
//var sw_rewrite_word_dict bool = false
//var sw_rewrite_row_dict  bool = false
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
const wSep = "§";                         // used to separe word in a list 
const endOfLine = ";;\n"
//const LAST_WORD0 = "zzz";
const LAST_WORD = "\xff\xff\xff"

const LAST_WORD_FREQ = 999999999 
//---------------------
var righe      =  []string{} 
//------------
//------
var parameter_path_html string  = "WBF_html_jsw"
//------------------

type rowStruct struct {
	rIdRow       string
	rixGroup     int      // indice del gruppo 
	rixBaseGroup int      // posizione del row nel gruppo ( si inzia dal num.1 )  
    rRow1        string
	rNfile1      int  
	rSwExtra     bool 
    rNumWords    int       // number of words in the row 
	rListIxUnF   []int     // for each word in the row, index of the word in the uniqueWordByFreq  
	rListFreq    []int     // for each word in the row, how many times the woird is used in all the text    
	rTran1       string 
}

type rowGroupStruct struct {
	rG_ixSelGrOption  int 
	rG_group          string 
	rG_firstIxRowOfGr int 
	rG_lastIxRowOfGr  int	
} 

type rowIxStruct struct {
	ixR_id        string  // group + "_" + string(100000 + num)
	ixR_id_gr     string
	ixR_id_num    int  
    ixR_ix        int	
	ixR_ix_last   int 
}
const SEL_EXTR_ROW    = 1; 
const SEL_NO_EXTR_ROW = 2; 

//--
type wordStruct struct {       // a word is repeated several time one for each row containing it  
	wWordCod  string
    wWord2    string
	wIxUniq   int               // index of uniqueWordByFreq      
	wNfile    int 
	wSwSelRowG int
	wSwSelRowR int
    wIxRow    int       
	wIxPosRow int 
	wTotRow   int              // number of rows 
	wTotExtrRow int            // number of extracted rows 
	wTotMinRow int
	wTotWrdRow int 
}
//--
type wordIxStruct struct {
	uWordCod    string	
    uWord2      string	
	uIxUnW      int            // index of this word in the uniqueWordByFreq	
	uIxUnW_al   int            // index of this word in the uniqueWordByAlpha 	
	uTotRow     int 
	uTotExtrRow int
    uIxWordFreq int            // index of this word in the wordSliceFreq	
	uSwSelRowG  int
	uSwSelRowR   int  
	uKnow_yes_ctr int 
	uKnow_no_ctr  int         // a value > 0  means that this is a word that I don't know, ie. it's to be learned   
	uIxLemmaL  []int  
	uLemmaL    []string       // list of lemma 
	//uTranL     []string       // list of translation    
	uLevel     []string  
	uPara      []string  
	uExample   []string  
}	
//---
type lemmaStruct struct {
	leLemma    string
	leNumWords int 
	leFromIxLW  int 
	leToIxLW    int  
	leTran      string 
	leLevel     string  
	lePara      string  
	leExample   string  	
} 
//-------------------------------
type wordLemmaPairStruct struct {
	lWordCod string 
	lWord2   string 
	lLemma   string
	lIxLemma int
} 
//---
type wDictStruct struct {
	dWord  		string 
	dIxWuFreq 	int 
	dLemmaL     []string 
	dTranL      []string 	
} 
//---
/**
type rDictStruct struct {
	rdIxRow  int 
	rdTran   string 
} 
**/
//---------------
type statStruct struct {
	uniqueWords  int 
	uniquePerc   int 
	totWords  int
	totPerc   int 
}
//--------------------------
var lastNumDict = 0;   
type lemmaTranStruct struct {
	dL_lemmaCod  string 
	dL_lemma2    string 
	dL_numDict   int	
	dL_tran      string         //
	
} 
//---------------
type lemmaWordStruct struct {
	lw_lemmaCod string 	
	lw_lemma2   string 	
	lw_word     string 
	lw_ixLemma    int
	lw_ixWordUnFr int
}
//------------
type paraStruct struct {
	p_lemma    string	
	p_ixLemma  int 
	p_level    string
	p_para     string
	p_example  string 
}	
//-----------------
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
var separRowList = make([]string,0,0)		


var msgLevelStat = "" 

var sw_PRINT_TIME bool = false;                 // in caso di durata abnorme, usa true per vedere dove impiega più tempo

var wordSliceAlpha = make([]wordStruct, 0, 0)  
var wordSliceFreq  = make([]wordStruct, 0, 0)  


var uniqueWordByFreq  []wordIxStruct;   // elenco delle parole in ordine di frequenza
var uniqueWordByAlpha []wordIxStruct;   // elenco delle parole in ordine alphabetico

var newWordLemmaPair    [] wordLemmaPairStruct // all word-lemma pair 


var dictLemmaTran []lemmaTranStruct ;  // dictionary lemma translation  
var lemma_word_ix []lemmaWordStruct  

//var wordStatistic_un =  make( []statStruct, 101, 101);	
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

var lemmaSlice       [] lemmaStruct         // lemma , translation 

var wordLemmaPair    [] wordLemmaPairStruct // all word-lemma pair  
var numLemmaDict int =0 
//var result_row1  string; 
//var result_word1 string;
//var result_word2 string;  
var numberOfRows  = 0; 
var numberOfWords = 0; 
var numberERRORS=0
var numberOfUniqueWords = 0; 
var showReadFile string = ""
//--------------------------------------------------------
var scrX, scrY int = getScreenXY();

var sw_stop bool = false
var errorMSG = ""; 
var sw_begin_ended = false     
var sw_HTML_ready  = false     
//--------------------------------

var lemma_para_list = make([]paraStruct, 0, 0 )   	
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
//---
var fseq = "z§" ; 
var lenFseq = len(fseq) 
//--------------------------
//--------------------------------
func check(e error) {
    if e != nil {
        panic(e)
    }
}

//-------------------------------
func getPgmArgs( key0, key1 , key2 , key3, key4 string) (string, string, bool, int, string) {  
	
	//  getPgmArgs("-html", "-input" , "-countNumLines" ,  "-maxNumLinesToWrite")	

	args1    :=  os.Args[1:]		
	
	
	var val0, val1, val2, val3, val4 string
	for a:=0; a < (len(args1)-1); a++ {
		switch args1[a] {
			case key0 :   val0 = args1[a+1]
			case key1 :   val1 = args1[a+1]
			case key2 :   val2 = args1[a+1]
			case key3 :   val3 = args1[a+1]
			case key4 :   val4 = args1[a+1]
		}
	}  
	var isCount = false;
	if strings.TrimSpace(val2) == "true" {
		isCount = true
	}
	var num=0; 
	num, err := strconv.Atoi( strings.TrimSpace(val3) )
	if err != nil {
		num=0
	}

	//fmt.Println("args=", args1,  " val0=", val0, " val1=", val1, " val2=", val2 , " val3=", val3, " num=", num, " val4=", val4)   
	
	return val0, val1, isCount, num, val4
	
} // end of getPgmArgs

//============================================

var ui, err = lorca.New("", "", scrX, scrY); // crea ambiente html e javascript  // if height and width set to more then maximum (eg. 2000, 2000), it seems it works  

//============================================
func main() {
	fmt.Println(  red("WordsByFrequence") )
	
	fmt.Println( "\ncolori:", red("rosso"), green("verde"), yellow("giallo"),  magenta("magenta"), cyan("ciano") , "\n"  )  
	
	val0, val1, val2, val3, val4 := getPgmArgs("-html",  "-input" , "-countNumLines" , "-maxNumLinesToWrite", "-lemmaformat")	
	countNumLines      = val2
	maxNumLinesToWrite = val3
	lemmaFormat        = val4
	
	sw_lemma_word = (lemmaFormat == "lemma-word")
	
	if val0 != "" { parameter_path_html = strings.ReplaceAll(val0,"\\","/")  } 
	if val1 != "" { file_inputControl   = strings.ReplaceAll(val1,"\\","/")  }  	
	
	
	
	fmt.Println("\n"+ "parameter_path_html =" + parameter_path_html + "\n" +  "input = " + file_inputControl )
	fmt.Println("countNumLines       = " ,  countNumLines)
	//fmt.Println("maxNumLinesToWrite = " , maxNumLinesToWrite) 
	
	fmt.Println("lemmaformat = ", lemmaFormat, " sw_lemma_word =", sw_lemma_word)
	fmt.Println("\n----------------\n")
	
	//------------------------------
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	//  err := lorca.New("", "", 480, 320, args...) moved out of main so that ui is available outside main()
	if err != nil {
		fmt.Println( red( "errore in lorca "), err )  //  //log.Fatal(err)
	}
	defer ui.Close()
	
	get_all_binds()  //  binds inside are executed asynchronously after calling from js function (html/js are ready) 
	
	begin_GO_HTML_Talk();  // this function is  executed firstily before html/js is ready  
	
	// the following in main() is executed at the end when the browser is close 
	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
		case <-sigc:
		case <-ui.Done():
	}
	fmt.Println("exiting") // log.Println("exiting...")
}
// =================================

//---------------------------------------
/**         ***  COLORS: got from https // www.dolthub.com/blog/2024-02-23-colors-in-golang/ ***
    var Reset   = "\033[0m" 
	var Red     = "\033[31m" 
	var Green   = "\033[32m" 
	var Yellow  = "\033[33m" 
	var Blue    = "\033[34m" 
	var Magenta = "\033[35m" 
	var Cyan    = "\033[36m" 
	var Gray    = "\033[37m" 
	var White   = "\033[97m"
**/
func red(     str1 string) string { return "\033[31m" + str1 +  "\033[0m" }
func green(   str1 string) string { return "\033[32m" + str1 +  "\033[0m" }
func yellow(  str1 string) string { return "\033[33m" + str1 +  "\033[0m" }
//func blue(    str1 string) string { return "\033[34m" + str1 +  "\033[0m" }
func magenta( str1 string) string { return "\033[35m" + str1 +  "\033[0m" }
func cyan(    str1 string) string { return "\033[36m" + str1 +  "\033[0m" }
//func gray(    str1 string) string { return "\033[37m" + str1 +  "\033[0m" }
//func white(   str1 string) string { return "\033[97m" + str1 +  "\033[0m" }

//---------------------------		

func getInt(x string) int {	
	y1, e1 := strconv.Atoi( x ) 
	if e1 == nil { 
		return y1
	} 
	y2, e2 := strconv.Atoi(  "0"+strings.TrimSpace(x)  ) 
	if e2 == nil {
		return y2
	} else {
		fmt.Println("error in getInt(",x,") ", e2) 
	}
	return 0
}

//----------------------------

//-------------------------------------------
func get_all_binds() {
		fmt.Println("func get_all_binds")
		
		/***
		// A simple way to know when UI is ready (uses body.onload event in HTML/JS)
		ui.Bind("goStart", func() { 
				fmt.Println("Bind goStart   sw_begin_ended=", sw_begin_ended); 
				sw_HTML_ready = true 
				if sw_begin_ended {
					bind_goStart("1get_all_binds")
				}	
			} )
		**/	
		//--------------
		ui.Bind("go_passToJs_html_is_ready", func( msg1 string,  js_function string) { 		
				bind_go_passToJs_html_is_ready( msg1,  js_function)  })
		//--------------
		/**
		ui.Bind("go_passToJs_html_and_go_ready", func( msg1 string,  js_function string) { 	
			bind_go_passToJs_html_and_go_ready( msg1,  js_function)  })
		**/
		//--------------
		//----------------------------
		ui.Bind("go_passToJs_wordList", func( isChange bool, s_fromWord string, s_numWords string, sel_level string, 
															sel_extrRow string, sel_toBeLearned string,  js_function string) { 		
				bind_go_passToJs_wordList( isChange, getInt(s_fromWord), getInt(s_numWords), sel_level, sel_extrRow, sel_toBeLearned,  js_function)  })
		//--------------------------------------		
		ui.Bind("go_passToJs_prefixWordList", func(s_numWords string, wordPrefix string, js_function string) {
				bind_go_passToJs_prefixWordList( getInt(s_numWords), wordPrefix, js_function) } )
		//--------------------------------------		
		ui.Bind("go_passToJs_betweenWordList", func( s_maxNumWords string, fromWordPref string, toWordPref string, js_function string) {
				bind_go_passToJs_betweenWordList( getInt(s_maxNumWords), fromWordPref, toWordPref, js_function) } ) 		 			
		//---------------------------------------
		ui.Bind("go_passToJs_lemmaWordList", func(lemma string, inpMaxWordLemma string, js_function string) {
				bind_go_passToJs_lemmaWordList(lemma, getInt( inpMaxWordLemma),  js_function) } ) 
		//---------------------------------------	
		ui.Bind("go_passToJs_getRowsByIxWord", func( sIxWord string, max_num_row4wordS string, js_function string) {
				bind_go_passToJs_getRowsByIxWord(  getInt(sIxWord), getInt(max_num_row4wordS), js_function) } ) 
		//-----------------------------------
		ui.Bind("go_passToJs_getRowsByIxLemma", func( sIxLemma string, max_num_row4lemmaS string, js_function string) {
				bind_go_passToJs_getRowsByIxLemma(  getInt(sIxLemma), getInt(max_num_row4lemmaS), js_function) } ) 
		//-----------------------------------
		
		ui.Bind("go_passToJs_getWordByIndex2", func( s_ixWord string, swOnlyThisWordRows bool, s_maxNumRow string,  js_function string) {
				bind_go_passToJs_getWordByIndex2(   getInt(s_ixWord), swOnlyThisWordRows,      getInt(s_maxNumRow), js_function) } ) 
		//---------------------------------------
		ui.Bind("go_passToJs_thisWordRowList", func( aWord string, s_maxNumRow string, js_function string) {				
				bind_go_passToJs_thisWordRowList( aWord, getInt(s_maxNumRow), js_function) } )
		//---------------------------------------
		ui.Bind("go_passToJs_rowList", func(  s_inpBegRow string,   s_maxNumRow string, js_function string) {	
				bind_go_passToJs_rowList(     getInt(s_inpBegRow), getInt(s_maxNumRow), js_function) } )
		//---------------------------------------
	
		ui.Bind("go_passToJs_rowWordList", func( numIdOut string, s_ixRR string, js_function string ) {
				bind_go_passToJs_rowWordList(numIdOut, getInt(s_ixRR), js_function) } ) 	
		//---------------------------------------	
		ui.Bind("go_write_lang_dictionary", func( langAndVoiceName string) {
			bind_go_write_lang_dictionary( langAndVoiceName ) } ) 
		//---------------------------------------	
		ui.Bind("go_write_word_dictionary", func( listGoWords string) {
			bind_go_write_word_dictionary( listGoWords ) } )  				
		//---------------------------------------			
		ui.Bind("go_write_row_dictionary", func( listGoRows string) {
			bind_go_write_row_dictionary( listGoRows ) } )  
		//---------------------
		ui.Bind("go_passToJs_word_known", func( s_ixWord string,  s_yesNo string,  s_knowCtr string, js_function string) {
			bind_go_passToJs_word_known(       getInt(s_ixWord), getInt(s_yesNo), getInt(s_knowCtr), js_function)  } )  
			
		//----------------------------------------------------------------
		/**
		ui.Bind("go_passToJs_read_wordsToLearn", func( js_function string) {
			bind_go_passToJs_read_wordsToLearn(js_function)  } )  
		**/
		//----------------------------------------------------------------
		ui.Bind("go_passToJs_write_WordsToLearn", func( js_function string) {
			bind_go_passToJs_write_WordsToLearn(js_function)  } )  	
			
		//-----------------------------------------
		/***
		ui.Bind("go_passToJs_updateRowGroup", func( s_index string,  s_inpBegRow string,  s_inpNumRow string,  js_function string) {	
			bind_go_passToJs_updateRowGroup(       getInt(s_index), getInt(s_inpBegRow), getInt(s_inpNumRow),  js_function)   } ) 
		**/
		//-------------------------------------------------------------------
		
		ui.Bind("go_passToJs_getIxRowFromGroup", func( s_rowGrIndex string,  s_html_rowGroup_beginNum string, s_html_rowGroup_numRows string, js_function string) {		
			bind_go_passToJs_getIxRowFromGroup( getInt(s_rowGrIndex),  getInt( s_html_rowGroup_beginNum), getInt( s_html_rowGroup_numRows),  js_function)  } ) 
		
		//----------------------------------------------------------------
}  
//---------------------------------------------
func bind_go_passToJs_html_is_ready( msg1 string,  js_function string) {
	fmt.Println("\n", "go func bind_go_passToJs_html_is_ready () " , "\n\t msg from html: ", msg1 )  
	
	fmt.Println("XXXXXXXXXX   ", green("html has been loaded"), "   XXXXXXXXXXXX")
	
	begin() 
	
	fmt.Println( green( "\n\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n" + 
		"xxxxxxxxxxxxxx you can use the tool xxxxxxxxxxxxxxxxxx\n"  + 
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n\n"  ) ) 
		
	//prova_js_function_treValori();
	
	//go_exec_js_function( js_function, msg1); 	
	
} // end of bind_go_passToJs_html_is_ready
 
//---------------------------------------------
/**
func bind_go_passToJs_html_and_go_ready( msg1 string,  js_function string) {
	fmt.Println("\n", "go func  bind_go_passToJs_html_and_go_ready() " , "\n\t msg from html: ", msg1 )  
	
	fmt.Println("XXXXXXXXXX   PRONTO  XXXXXXXXXXXX")
	
	begin() 
	
} // end of bind_go_passToJs_html_and_go_ready
***/
//-------------------------------------------------------- 


//-------------------------------------
func showErrMsg(errorMSG0 string, msg1 string, func1 string) {
	errorMSG = strings.ReplaceAll( errorMSG0, "\n"," ") 
	fmt.Println(msg1, " func ", func1)
	go_exec_js_function( "js_go_setError", errorMSG ); 	
}
//--------------------------------
func showErrMsg2(errorMSG0 string, msg1 string) {
	errorMSG = strings.ReplaceAll( errorMSG0, "\n"," ") 
	fmt.Println(msg1)
	go_exec_js_function( "js_go_setError", errorMSG ); 	
}
//------------------------------
func showInfoRead( fileName string, startStop string) {
		
	msg1_Js := "file " + fileName + " " + startStop 
 	msg1:=`<br>file &nbsp;&nbsp;<span style="color:blue;" >` + fileName + `</span>`	+ 				
				`<br><span style="font-size:0.7em;color:black;">`	+ startStop	+ `</span>` 		
				
	showErrMsg2(msg1, msg1_Js)	
}
//------------------------------------------------
func isNumber(s string) bool {
    for _, v := range s {
        if v < '0' || v > '9' {
            return false
        }
    }
    return true
}

//-------------------------------------------

func bind_go_passToJs_wordList( isChange_extrRow bool, fromWord int, numWords int, onlyThisLevel string, 
						sel_extrRow string, sel_toBeLearned string,  js_function string) {
					
		fmt.Println( cyan("bind_go_passToJs_wordList"), "isChange_extrRow=", isChange_extrRow, " sel_extrRow=", sel_extrRow, " sel_toBeLearned=", sel_toBeLearned, 
						" from=",  fromWord, " numWords=", numWords )
		var from1, to1 int; 
		from1 = fromWord; //   - 1; 
		if (from1 < 1) {
			if from1 < 0 {
				go_exec_js_function( js_function, ""); 	
				return
			}
			from1=1;
		}
		if (from1 >= numberOfUniqueWords) {
			from1 = numberOfUniqueWords - 1;	
			if from1 < 0 {
				fmt.Println("error in bind_go_passToJs_wordList() numberOfUniqueWords=", numberOfUniqueWords)
				go_exec_js_function( js_function, ""); 	
				return
			}
		}
		//---------------------------
		sw_tobeLearnedOnly := ( sel_toBeLearned == "toBeLearned" )                //     ( 0 = 'allWords'   1 = 'toBeLearned' )
			
		//----------------
		if (isChange_extrRow) {			
			//write_lastValueSets_wordList( fromWord, numWords, onlyThisLevel, sel_extrRow, isAlpha)
					
			last_word_fromWord = fromWord 
			last_word_numWords = numWords 
			last_sel_extrRow   = sel_extrRow	
			write_lastValueSets()	
			
			if (sel_extrRow != last_run_extrRow) {
				fmt.Println("on the word list button,  an option has been changed from \"" + last_run_extrRow + "\" to \"" + sel_extrRow + "\", this causes a rebuild of wordlist data")
				last_run_extrRow = sel_extrRow
				fmt.Println("bind_go_passToJs_wordList() sel_extrRow != last_run_extrRow call build_and_elab_word_list() ")
				
				build_and_elab_word_list()
				
				fmt.Println("bind_go_passToJs_wordList() end of build_and_elab_word_list() ")
				
			}	
			//fmt.Println("bind_go_passToJs_wordList() return ")	
			//return // this func has been called only to set the mainPage values 
		}
		//--------------------		

		//write_lastValueSets_wordList( fromWord, numWords, onlyThisLevel, sel_extrRow, isAlpha)
		
		last_word_fromWord = fromWord 
		last_word_numWords = numWords 
		last_sel_extrRow   = sel_extrRow	
		write_lastValueSets()	
		
		
		//to1 = numWords + from1; 
		if (to1 > numberOfUniqueWords)   {to1 = numberOfUniqueWords;}	 
				
		//var xWordF wordIxStruct;  
		var outS1 string = ""; 
		//var row11 string;
		//var numNoTran = 0
		
		//fmt.Println("bind_go_passToJs_wordList() 0  from1 m=", from1 ) 
		
		onlyIfExtr := true 
		
		
		if from1 == 1 { from1=0; }
		numOut:=0
		for i:= from1; i < numberOfUniqueWords; i++ { 
			if sw_tobeLearnedOnly {
				if uniqueWordByFreq[i].uKnow_no_ctr < 1 { 
					continue
				} 
			}	
			//fmt.Println("call 3  loop  i=", i, "   ", uniqueWordByFreq[i]);
			
			sw, rowW := word_to_row("", onlyIfExtr, onlyThisLevel, uniqueWordByFreq[i] )  
			if sw {	
				outS1 += rowW 						  	
				numOut++
				if numOut >= numWords { 
					break
				}
			}
		}	
		
		go_exec_js_function( js_function, outS1 ); 	
		
}  // end of bind_go_passToJs_wordList	

//---------------------------------------------------------

//------------------------------------------------------

func word_to_row(onlyThisLemma string, onlyIfExtr bool, onlyThisLevel string, xWordF2 wordIxStruct) (bool, string)  {
	
	//fmt.Println("\nXXXXXXX word_to_row ( onlyThisLevel=" + onlyThisLevel + 
	//  	"<==   onlyThisLevel=" + level_other +  "<==" + "  sw_list_Word_if_in_ExtrRow=" , sw_list_Word_if_in_ExtrRow); 
	
	sw:= true
	
	if onlyIfExtr {
		if (sw_list_Word_if_in_ExtrRow) {
			if (xWordF2.uSwSelRowG == SEL_NO_EXTR_ROW) {
				sw = false
				//fmt.Println("word to row " , xWordF2.uWord2, " da ignorare")
				return sw, ""
			} 
		}
	}
	//fmt.Println("word to row " , xWordF2.uWord2, " \t\t\t XXXXXXXXXXXX (xWordF2.uTotExtrRow =" ,xWordF2.uTotExtrRow , " xxxxxxxxxxxxxxxxxxx   accettato")
	//-----------
	
	if onlyThisLemma != "" {
		ix2 := -1
		for x1, oneLemma := range xWordF2.uLemmaL {
			if oneLemma == onlyThisLemma {
				ix2=x1; 
				break
			}  	
		}
		if ix2 >=0 {
			return sw, xWordF2.uWordCod + ";." + xWordF2.uWord2 + ";." + 
				"ix" + ";." + 
				strconv.Itoa(xWordF2.uIxUnW) + ";." + strconv.Itoa(xWordF2.uTotRow)  + ";." + 
				xWordF2.uLemmaL[ix2]              + ";." + 
				tranFromIxLemma( xWordF2, ix2)    + ";." +  
				xWordF2.uLevel[ix2]               + ";." +  
				xWordF2.uPara[ix2]                + ";." +  
				xWordF2.uExample[ix2]             + ";." +  
				strconv.Itoa(xWordF2.uTotExtrRow) + ";." +  					
				strconv.Itoa(xWordF2.uKnow_yes_ctr)  + ";." + strconv.Itoa(xWordF2.uKnow_no_ctr)  + ";." + 				
				"ixLemma" + ";." + fmt.Sprint( xWordF2.uIxLemmaL[ix2] ) + ";." + 	
				endOfLine 		
		}
	}
	
	//---------------------------	
	
	return sw, xWordF2.uWordCod + ";." + xWordF2.uWord2 + ";." + 
				"ix" + ";." + 
				strconv.Itoa(xWordF2.uIxUnW) + ";." + strconv.Itoa(xWordF2.uTotRow)  + ";." +
				fmt.Sprint( strings.Join(xWordF2.uLemmaL,  wSep)  ) + ";." + 
				//fmt.Sprint( strings.Join(xWordF2.uTranL,   wSep)  ) + ";." +  
				listStringLemmaSlice_Tran(xWordF2) + ";." +  
				fmt.Sprint( strings.Join(xWordF2.uLevel,   wSep)  ) + ";." +  
				fmt.Sprint( strings.Join(xWordF2.uPara,    wSep)  ) + ";." +  
				fmt.Sprint( strings.Join(xWordF2.uExample, wSep)  ) + ";." +  
				strconv.Itoa(xWordF2.uTotExtrRow) + ";." +  					
				strconv.Itoa(xWordF2.uKnow_yes_ctr)  + ";." + strconv.Itoa(xWordF2.uKnow_no_ctr)  + ";." + 				
				"ixLemma" + ";." + intSliceToString( xWordF2.uIxLemmaL,wSep )  + ";." + 		
				endOfLine 
				 
} // end of word_to_row 
//-----------------------------------
func intSliceToString(mySlice []int, wSep string) string {
	output := ""
	for _, v := range mySlice {
		output += (strconv.Itoa(v) + wSep)
	}
	return output
}
//------------------------------------------------------

func bind_go_passToJs_prefixWordList( numWords int, wordPrefix string, js_function string) {
	
	bind_go_passToJs_betweenWordList( numWords, wordPrefix, wordPrefix, js_function) 
			
} // end of bind_go_passToJs_prefixWordList

//-----------------------------------------

func bind_go_passToJs_betweenWordList( maxNumWords int, fromWordPref string, toWordPref string, js_function string) {
	
	/**
	var uno = LAST_WORD0
	var due = LAST_WORD
	fmt.Println( "(uno=", uno , " > ", " due=",  due, ") = ",  ( uno > due ) );     
	**/
	var onlyThisLevel string = "any" ; // "A0"  // questo deve arrivare da parametro  
	var outS1 string; 	
	
	manyWordInputList := regexp.MustCompile(separWord).Split( fromWordPref, -1);  // split into words 
	
	lenManyWord := len(manyWordInputList)
	if lenManyWord < 1 {
		go_exec_js_function( js_function, "");
		return	
	}
	if lenManyWord > 1 {
		toWordPref = ""
	}	
	
	//fmt.Println( "bind_go_passToJs_betweenWordList() fromWordPref=>" + fromWordPref  + "<== manyWordInputList len=",  lenManyWord, " words=",  manyWordInputList )    
	
	//-----------------------------
	for _, oneWord := range(manyWordInputList) {
		
		
		
		fromWord   := strings.ToLower(strings.TrimSpace( oneWord ));  
		if fromWord == "" { continue }
		
		//fmt.Println( "    bind_go_passToJs_betweenWordList() oneWord=", oneWord)
		
		toWordPref := fromWord 		
		
		fromWordCod:= newCode(fromWord)		
		from1, _:= lookForWordInUniqueAlpha( fromWordCod)	
		
		toWordPref2:= toWordPref + LAST_WORD
		toWord := strings.ToLower(strings.TrimSpace( toWordPref2 ));  
		toWordCod:= newCode(toWord)		
		
		_,   to2:= lookForWordInUniqueAlpha( toWordCod)			
		
		//--
		if from1 < 2 { from1=0} 
		if to2  >=  len(uniqueWordByAlpha) { to2 =  len(uniqueWordByAlpha) - 1 } 
		
		toWordCod += LAST_WORD
		num1:=0
		
		onlyIfExtr := false 
		
		//fmt.Println( "    1 bind_go_passToJs_betweenWordList() oneWord=", oneWord,    " from1=", from1, "  to2=", to2)	
		
		swNone:=true 
		for i:=from1; i <=to2; i++ {		
			sw, rowW := word_to_row("", onlyIfExtr, onlyThisLevel,  uniqueWordByAlpha[i] )  	  
			
			//fmt.Println( "      2 bind_go_passToJs_betweenWordList()  sw=", sw, " rowW=", rowW )	
			//if sw == false { continue }
			
			j1 := strings.Index(rowW, ";.") 
			if (sw) {
				if rowW[0:j1] < fromWordCod { continue}	
				if rowW[0:j1] > toWordCod { break}	
			}
			swNone = false 
			
			//fmt.Println( "      3 bind_go_passToJs_betweenWordList()  rowW[0:j1]=",  rowW[0:j1] )	
			
			outS1 += rowW 	
			num1++
			if num1 >= maxNumWords { break }
		}
		if swNone {
			rowW:= notFoundWord_row( fromWordCod, fromWord)
			
			//j1 := strings.Index(rowW, ";.")	
			//fmt.Println( "      3 bind_go_passToJs_betweenWordList()  rowW[0:j1]=",  rowW[0:j1] )				
			
			outS1 += rowW 	
			num1++
			if num1 >= maxNumWords { break }
		}
			
	} // end of for oneword range
	//-----------
	
	go_exec_js_function( js_function, outS1 ); 		
	
			
} // end of bind_go_passToJs_betweenWordList
//------------------

func notFoundWord_row( fromWordCod string, fromWord string ) string {
	/**
					xWordF2.uWordCod + ";." + xWordF2.uWord2 + ";." + 
					strconv.Itoa(xWordF2.uIxUnW) + ";." + strconv.Itoa(xWordF2.uTotRow)  + ";." + 
					fmt.Sprint( strings.Join(xWordF2.uLemmaL,  wSep)  ) + ";." + 
					fmt.Sprint( strings.Join(xWordF2.uTranL,   wSep)  ) + ";." +  
					fmt.Sprint( strings.Join(xWordF2.uLevel,   wSep)  ) + ";." +  
					fmt.Sprint( strings.Join(xWordF2.uPara,    wSep)  ) + ";." +  
					fmt.Sprint( strings.Join(xWordF2.uExample, wSep)  ) + ";." +  
					strconv.Itoa(xWordF2.uTotExtrRow) + ";." +  					
					strconv.Itoa(xWordF2.uKnow_yes_ctr)  + ";." + strconv.Itoa(xWordF2.uKnow_no_ctr)  + ";." + 
					endOfLine 
	**/
	return fromWordCod + ";." + fromWord + ";." + 
					strconv.Itoa(0) + ";." + strconv.Itoa(0)  + ";." + 
					fmt.Sprint( fromWord ) + ";." + 
					fmt.Sprint( "_word_not_found_" ) + ";." +  
					fmt.Sprint( "" ) + ";." +  
					fmt.Sprint( "" ) + ";." +  
					fmt.Sprint( "" ) + ";." +  
					strconv.Itoa(0) + ";." +  					
					strconv.Itoa(0) + ";." + strconv.Itoa(0)  + ";." + 
					endOfLine 
					
} // end of notFoundWord_row

//---------------------------------------------------


//---------------------------------------------- 
func bind_go_passToJs_lemmaWordList(lemmaToFind0 string, inpMaxWordLemma int, js_function string)   {

		lemmaCod:= newCode( lemmaToFind0 )
		var onlyThisLevel string = "any" ; // "A0"  // questo deve arrivare da parametro  
		onlyIfExtr := false 
		outS1 := "" 
				
		fromIxX, _ := lookForLemmaWord( lemmaCod )
		
		fromIx:= fromIxX
		for k:= fromIxX; k >= 0; k-- {
			if lemma_word_ix[k].lw_lemmaCod < lemmaCod { break }
			fromIx = k
		}
		
		numO:= 0
		
		if fromIx == 1 { fromIx=0; }
	
		for k:= fromIx; k < len(lemma_word_ix); k++ {
		
			if lemma_word_ix[k].lw_lemmaCod == lemmaCod {		
				ix := lemma_word_ix[k].lw_ixWordUnFr 	
				//outS1 += word_to_row ( onlyThisLevel,  uniqueWordByFreq[ix] ) 	
				sw, rowW := word_to_row(lemmaToFind0, onlyIfExtr, onlyThisLevel, uniqueWordByFreq[ix] )  
				if sw {	
					numO++
					if numO > inpMaxWordLemma { break } 
					outS1 += rowW 
				}  					
			} else {
				if lemma_word_ix[k].lw_lemmaCod > lemmaCod { break }
			}
		} 
		if len(outS1)< 3 {
			outS1 = "NONE," + lemmaToFind0; 			
			//fmt.Println(" bind_go_passToJs_lemmaWordList() ", outS1) 
		} 	
		
		go_exec_js_function( js_function, outS1 ); 		
				
} // end of bind_go_passToJs_lemmaWordList

//-----------------------------------------------------------

//-----------------------------------------------------------

func bind_go_passToJs_getWordByIndex( ixWord int, maxNumRow int, js_function string) {
		
		if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}	
		
		var xWordF     = uniqueWordByFreq[ixWord]  	
		aWord:= xWordF.uWord2; 
		
		swOnlyThisWordRows := true;
		
		OLD_bind_go_passToJs_thisWordRowList( aWord, swOnlyThisWordRows, maxNumRow, js_function)
		
} // end of bind_go_passToJs_getWordByIndex

//-----------------------------------------------------------
func bind_go_passToJs_getRowsByIxWord( ixWord int, maxNumRow int, js_function string) {
	if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}	
	
	listIxRR       := make([]int,0, maxNumRow)
	
	listIxRR = getRowIndicesFromIxFreqWord(ixWord, maxNumRow)
	var outS1 string;
	//--------------		
	hd_tr := ""; 
	preL:= ""
	preW:=""
	listWords := ""
	
	xWordF  := uniqueWordByFreq[ixWord]   
	listWords += " " +  xWordF.uWord2
	
	aWord:= xWordF.uWord2; 
	
	//-------------
	for z:=0; z < len(xWordF.uLemmaL); z++  {
		ixL1:= xWordF.uIxLemmaL[z]
		LeS := lemmaSlice[ixL1]
		newL:= xWordF.uLemmaL[z]
		if newL != LeS.leLemma {
			continue;  // error 
		}			
		newT:= LeS.leTran
		newP:= xWordF.uPara[z]
		if newL == preL { 
			newL=""
			newT=""
			if preW != xWordF.uWord2 { 
				hd_tr += xWordF.uWord2 + " ";  
				preW = xWordF.uWord2
			} 
		} else { 				
			if newP == "" {
				newP = newL;
			} 				
			preL = newL 			
			if hd_tr != "" { hd_tr += "\n" 	}  // 14giugno 
			
			hd_tr += " :lemma="  + newP + " :tran=" + newT  + " :wordsInLemma=" +	xWordF.uWord2 + " "
			//hd_tr += " :lemma="  + newP + " :tran=" + newT + " "
			
			preW =  xWordF.uWord2
		} 			
	} // end for z
	 
	//-------
	
	hd_tr += "\n"
	
	sort.Ints(listIxRR) 
		
	preIxRR_2:= 999999999 
	ixRR_2:=0
	ixRR  :=0
	
	nOut:=0
	new_rIdRow :=""

	for n1:= 0; n1 < len(listIxRR); n1++  {
		ixRR_2 = listIxRR[n1]
		if (ixRR_2 == preIxRR_2) { continue;} 
		preIxRR_2 = ixRR_2; 
		ixRR = ixRR_2 % 100000; 
			
		if ixRR >= numberOfRows { continue;} // actually there  must be some error here 		
		rline := inputTextRowSlice[ixRR]
		rowX := cleanRow(rline.rRow1)	
			
		if ((rowX =="") || (rowX == LAST_WORD)) { 
			continue 
		}		
		
		if rline.rixGroup < 0 { 
			new_rIdRow = "- " + strconv.Itoa( rline.rixBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		} else {
			new_rIdRow = lista_gruppiSelectRow[ rline.rixGroup ].rG_group + " " + strconv.Itoa( rline.rixBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		}	
		outS1 += "<br>" + strconv.Itoa( SEL_EXTR_ROW ) + "|" + new_rIdRow   + "|" + strconv.Itoa( ixRR) + "|"   + rowX + "|" + rline.rTran1; 
		
		nOut++
		if (nOut >= maxNumRow) {
			break;
		}
		
	} 	// end for n1
	
	header:= "<HEADER>\n" + "<WORD>" + aWord + ",L:" + strings.TrimSpace(listWords) + "</WORD>"
	//header:= "<HEADER>\n" + "<WORD>" + aWord + "</WORD>"
	header += hd_tr   // 14giugno
	header += "</HEADER> \n"
	
	go_exec_js_function( js_function, header + outS1 ); 					
	
} // end of  bind_go_passToJs_getRowsByIxWord  

//-----------------------------------------------------------
func bind_go_passToJs_getRowsByIxLemma( ixLemma int, max_num_row4lemma int, js_function string) {

		/***
		type lemmaStruct struct {
			leLemma    string
			leNumWords int 
			leFromIxLW  int 
			leToIxLW    int  
			leTran      string 
			leLevel     string  
			lePara      string  
			leExample   string  	
		} 
		//-------------------------------
		type wordLemmaPairStruct struct {
			lWordCod string 
			lWord2   string 
			lLemma   string
			lIxLemma int
		} ù//------------
		type wordLemmaPairStruct struct {
			lWordCod string 
			lWord2   string 
			lLemma   string
			lIxLemma int
		} 
		//--
		type lemmaWordStruct struct {
			lw_lemmaCod string 	
			lw_lemma2   string 	
			lw_word     string 
			lw_ixLemma    int
			lw_ixWordUnFr int
		}
		//---
		***/
		
		leS:= lemmaSlice[ixLemma]; 
		
		lemmaToFind0:= leS.leLemma;
		lemmaTran   := leS.leTran
		
		fmt.Println( green("bind_go_passToJs_getRowsByIxLemma"), " ixLemma=", ixLemma,  " lemmaSlice[]=", leS)
		
		for f:= leS.leFromIxLW; f <= leS.leToIxLW; f++ {
			
			wordLW := lemma_word_ix[f]
		
			fmt.Println("\t" , " lemma=", wordLW.lw_lemma2, " word=", wordLW.lw_word, " ixLemma=", wordLW.lw_ixLemma, " ixFreq=", wordLW.lw_ixWordUnFr) 
			
			//???? bind_go_passToJs_thisWordRowList( aWord string,  maxNumRow int, js_function string) { 
		
		}  
		//---------------------
		
		lemmaCod:= newCode( lemmaToFind0 )
		
		
		
				
		listIxRR        := make([]int,0, max_num_row4lemma)	
		listIxRR_temp   := make([]int,0, max_num_row4lemma)			
		
		outS1 :=""
			//--------------		
			hd_tr := ""; 
			
			listWords := ""
			//------------------
	
		totRR:=0 
		//------------------------------
		for kLe:= leS.leFromIxLW; kLe <= leS.leToIxLW; kLe++ {
		
			if lemma_word_ix[kLe].lw_lemmaCod != lemmaCod { continue}		// error 
			
			ixWord := lemma_word_ix[kLe].lw_ixWordUnFr 			        // indice word da lemma 
			xWordF := uniqueWordByFreq[ixWord]

			listWords += " " +  xWordF.uWord2	
			hd_tr += " :lemma="  + lemmaToFind0 + " :tran=" + lemmaTran + " :wordsInLemma=" +	xWordF.uWord2 + " "
			
			listIxRR_temp = getRowIndicesFromIxFreqWord(ixWord,  max_num_row4lemma)   // indici row da word 
			for _,lis:= range listIxRR_temp {
				totRR++
				if totRR > max_num_row4lemma { break }
				listIxRR = append(listIxRR, lis) 
			} 
			if totRR > max_num_row4lemma { break }
		} // end for kLe 
		
		
		 
	//-------
	
	hd_tr += "\n"
	
	sort.Ints(listIxRR) 
		
	preIxRR_2:= 999999999 
	ixRR_2:=0
	ixRR  :=0
	
	nOut:=0
	new_rIdRow :=""

	for n1:= 0; n1 < len(listIxRR); n1++  {
		ixRR_2 = listIxRR[n1]
		if (ixRR_2 == preIxRR_2) { continue;} 
		preIxRR_2 = ixRR_2; 
		ixRR = ixRR_2 % 100000; 
			
		if ixRR >= numberOfRows { continue;} // actually there  must be some error here 		
		rline := inputTextRowSlice[ixRR]
		rowX := cleanRow(rline.rRow1)	
			
		if ((rowX =="") || (rowX == LAST_WORD)) { 
			continue 
		}		
		
		if rline.rixGroup < 0 { 
			new_rIdRow = "- " + strconv.Itoa( rline.rixBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		} else {
			new_rIdRow = lista_gruppiSelectRow[ rline.rixGroup ].rG_group + " " + strconv.Itoa( rline.rixBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		}	
		outS1 += "<br>" + strconv.Itoa( SEL_EXTR_ROW ) + "|" + new_rIdRow   + "|" + strconv.Itoa( ixRR) + "|"   + rowX + "|" + rline.rTran1; 
		
		nOut++
		if (nOut >=  max_num_row4lemma) {
			break;
		}
		
	} 	// end for n1
	
	aWord:=  strings.TrimSpace(listWords)
	
	header:= "<HEADER>\n" + "<WORD>" + aWord + ",L:" + strings.TrimSpace(listWords) + "</WORD>"
	//header:= "<HEADER>\n" + "<WORD>" + aWord + "</WORD>"
	header += hd_tr   // 14giugno
	header += "</HEADER> \n"
	
	go_exec_js_function( js_function, header + outS1 ); 					
		
		go_exec_js_function( js_function, header + outS1 ); 	
		
} // end of  bind_go_passToJs_getRowsByIxLemma  

//--------------------
func bind_go_passToJs_getWordByIndex2( ixWord int, swOnlyThisWordRows bool, maxNumRow int, js_function string) {
		
		if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}	
		
		var xWordF     = uniqueWordByFreq[ixWord]  	
		aWord:= xWordF.uWord2; 
		
		OLD_bind_go_passToJs_thisWordRowList( aWord, swOnlyThisWordRows, maxNumRow, js_function)
		
} // end of bind_go_passToJs_getWordByIndex2



//----------------------------------------------------

//----------------------------------------------------
func getRowIndicesFromIxFreqWord(ixWord int, maxNumRow int) []int {

		var xWordF     = uniqueWordByFreq[ixWord]   	

		var ixFromList = xWordF.uIxWordFreq                     // index of this word in the wordSliceFreq	( una word per ogni row ) 
		var ixToList   = ixFromList + xWordF.uTotRow;
		var maxTo1     = ixFromList + maxNumRow; 		
		
		if ixToList > maxTo1        { ixToList = maxTo1; }
		if ixToList > numberOfWords { ixToList = numberOfWords; }

		listIxRR := make([]int,0, (ixToList - ixFromList) )
		
		for ix4:=ixFromList; ix4 < ixToList; ix4++ {
			wS1 := wordSliceFreq[ix4] 
			//  .wIxRow           indice del row che contiene la parola 	
			//  .wSwSelRowR	      1 or 2: 1 SEL_EXTR_ROW, 2 SEL_NO_EXTR_ROW  ( serve per dare la precedenza alle righe che appartengono al brano selezionato    
			listIxRR = append( listIxRR, wS1.wSwSelRowR * 100000 +  wS1.wIxRow)  	
		} 	
		
		sort.Ints(listIxRR) 
		
		return listIxRR 
		
} // end of getRowIndicesFromIxFreqWord



//---------------------------------------------
func bind_go_passToJs_thisWordRowList( aWord string,  maxNumRow int, js_function string) {  
	
	//  lista tutte le frasi che contengono le parole con lemma della parola cercata 
	//fmt.Println("bind_go_passToJs_thisWordLemmaWordRowList() 1  aWord=", aWord )
  
	var outS1 string;
	
	//---------------------------------------

	var xWordA, xWordF wordIxStruct;  		
	wordCod:= newCode( aWord )		
	
	ixF, ixT:= lookForWordInUniqueAlpha( wordCod)	
	
	//fmt.Println("   lookForWordInUniqueAlpha(  wordCod=", wordCod,   " ixF=", ixF, " ixT=", ixT) 
	if (ixT < 0) {
			outS1 += "NONE," + aWord  
			go_exec_js_function( js_function, outS1 ); 	
	} 
	//----
	listIxRR := make([]int,0, maxNumRow)
	listRowIndices := make([]int,0, maxNumRow)
	
	ixWord:= -1 
	for ix:= ixF; ix <= ixT; ix++ {
		xWordA =  uniqueWordByAlpha[ix] 
		
		if xWordA.uWordCod != wordCod { continue }            // get only the required word (might be several entries of the same word) and then the list of lemmas of this word 
		
		ixWord = xWordA.uIxUnW			
		if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}	
		
		listRowIndices = getRowIndicesFromIxFreqWord(ixWord, maxNumRow)

		for _, ind1:= range listRowIndices {
			listIxRR = append(listIxRR, ind1 )
		}			
	}
	if ixWord < 0 { return }
	
	//--------------
		
	hd_tr := ""; 
	preL:= ""
	preW:=""
	listWords := ""
	
	xWordF  = uniqueWordByFreq[ixWord]   
	listWords += " " +  xWordF.uWord2
	
	//-------------
	for z:=0; z < len(xWordF.uLemmaL); z++  {
		ixL1:= xWordF.uIxLemmaL[z]
		LeS := lemmaSlice[ixL1]
		newL:= xWordF.uLemmaL[z]
		if newL != LeS.leLemma {
			continue;  // error 
		}			
		newT:= LeS.leTran
		newP:= xWordF.uPara[z]
		if newL == preL { 
			newL=""
			newT=""
			if preW != xWordF.uWord2 { 
				hd_tr += xWordF.uWord2 + " ";  
				preW = xWordF.uWord2
			} 
		} else { 				
			if newP == "" {
				newP = newL;
			} 				
			preL = newL 			
			if hd_tr != "" { hd_tr += "\n" 	}  // 14giugno 
			
			hd_tr += " :lemma="  + newP + " :tran=" + newT  + " :wordsInLemma=" +	xWordF.uWord2 + " "
			//hd_tr += " :lemma="  + newP + " :tran=" + newT + " "
			
			preW =  xWordF.uWord2
		} 			
	} // end for z
	 
	//-------
	
	hd_tr += "\n"
	
	sort.Ints(listIxRR) 
		
	preIxRR_2:= 999999999 
	ixRR_2:=0
	ixRR  :=0
	
	nOut:=0
	new_rIdRow :=""

	for n1:= 0; n1 < len(listIxRR); n1++  {
		ixRR_2 = listIxRR[n1]
		if (ixRR_2 == preIxRR_2) { continue;} 
		preIxRR_2 = ixRR_2; 
		ixRR = ixRR_2 % 100000; 
			
		if ixRR >= numberOfRows { continue;} // actually there  must be some error here 		
		rline := inputTextRowSlice[ixRR]
		rowX := cleanRow(rline.rRow1)	
			
		if ((rowX =="") || (rowX == LAST_WORD)) { 
			continue 
		}		
		
		if rline.rixGroup < 0 { 
			new_rIdRow = "- " + strconv.Itoa( rline.rixBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		} else {
			new_rIdRow = lista_gruppiSelectRow[ rline.rixGroup ].rG_group + " " + strconv.Itoa( rline.rixBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		}	
		outS1 += "<br>" + strconv.Itoa( SEL_EXTR_ROW ) + "|" + new_rIdRow   + "|" + strconv.Itoa( ixRR) + "|"   + rowX + "|" + rline.rTran1; 
		
		nOut++
		if (nOut >= maxNumRow) {
			break;
		}
		
	} 	// end for n1
	
	header:= "<HEADER>\n" + "<WORD>" + aWord + ",L:" + strings.TrimSpace(listWords) + "</WORD>"
	//header:= "<HEADER>\n" + "<WORD>" + aWord + "</WORD>"
	header += hd_tr   // 14giugno
	header += "</HEADER> \n"
	
	go_exec_js_function( js_function, header + outS1 ); 	
				
} // end of bind_go_passToJs_thisWordRowList

//-----------------------------------------------------------
//----------------------------------------------

func OLD_bind_go_passToJs_thisWordRowList( aWord string, swOnlyThisWordRows bool, maxNumRow int, js_function string) {    // NEW 
	
	//  lista tutte le frasi che contengono le parole con lemma della parola cercata 
	//fmt.Println("bind_go_passToJs_thisWordLemmaWordRowList() 1  aWord=", aWord )
    
	//  1) from the word get all lemma 
	// 	2) from each lemma get all words   ( unless swOnlyThisWordRows  
	//  3) from each word all row
	//------------------
	
	var outS1 string;
	
	//---------------------------------------
	//  1) from the word get all lemma 
	//-------	
	var xWordF wordIxStruct;  		
	wordCod:= newCode( aWord )		
	ixF, ixT:= lookForWordInUniqueAlpha( wordCod)	
	
	//fmt.Println("   lookForWordInUniqueAlpha(  wordCod=", wordCod,   " ixF=", ixF, " ixT=", ixT) 
	if (ixT < 0) {
			outS1 += "NONE," + aWord  
			go_exec_js_function( js_function, outS1 ); 	
	} 
	//----
	lemmaList1:= make([]string,0,10)
	
	for ix:= ixF; ix <= ixT; ix++ {
		xWordF =  uniqueWordByAlpha[ix] 
		
		if xWordF.uWordCod != wordCod { continue }            // get only the required word (might be several entries of the same word) and then the list of lemmas of this word 
		
		/**
			if xWordF.uWord2 != aWord { continue; }
			if swOnlyThisWordRows {
				if xWordF.uWordCod != wordCod { continue }            // get only the required word 
		}
		**/
		ixWord := xWordF.uIxUnW			
		if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}		
		xWordF     = uniqueWordByFreq[ixWord] 
		// get the list of lemma of this word 
		for z:=0; z < len(xWordF.uLemmaL); z++  {
			lemmaList1 = append(lemmaList1, newCode( xWordF.uLemmaL[z]) ) 
		}			
	}
	if len(lemmaList1) < 1 { 
			outS1 += "NONE," + aWord  
			go_exec_js_function( js_function, outS1 ); 	
	} 
	//----   
	sort.Strings(lemmaList1)
	
	//fmt.Println("bind_go_passToJs_thisWordLemmaWordRowList() 2  lemmaList1=", lemmaList1 )
	
	//-----------------
	//  2) from each lemma get all words 
	//
	preL:= ""
	lemWordList2 := make([]lemmaWordStruct, 0, 10)  
	var wL2 lemmaWordStruct;
	
	//contaAnto3:=0	
	
	for _, lemma1Cod:= range( lemmaList1 ) {
	
	
		if lemma1Cod == preL { continue }
		preL = lemma1Cod	
		fromIxX, _ := lookForLemmaWord( lemma1Cod )		// se non trova cerca lemma1Cod fino al punto.
		fromIx:= fromIxX
		for k:= fromIxX; k >= 0; k-- {
			if lemma_word_ix[k].lw_lemmaCod < lemma1Cod { break }
			fromIx = k
		}
		for k:= fromIx; k < len(lemma_word_ix); k++ {
			if lemma_word_ix[k].lw_lemmaCod != lemma1Cod { continue	}   	//  se non trova cerca lemma1Cod fino al punto.
			wL2 = lemma_word_ix[k]	
			if swOnlyThisWordRows { 
				if aWord != wL2.lw_word { continue;}                // get only the required word  
				/**
				if aWord == "schrift" {
					fmt.Println(" antocontAnto3 lemma ", wL2) 
					contaAnto3++
				}
				***/
			}
			wL2.lw_word = newCode( wL2.lw_word )
			lemWordList2 = append( lemWordList2, wL2 )
		} 		
	}	
	//----
	//fmt.Println("anto lemma contaAnto3=",	contaAnto3) 
	
	//fmt.Println("bind_go_passToJs_thisWordLemmaWordRowList() 9 call  fun_wordListToRowList_head()")
	
	fun_wordListToRowList_head( aWord, lemWordList2, maxNumRow, js_function) 	
				
} // end of OLDbind_go_passToJs_thisWordRowList

//-----------------------------------------------------------

func fun_wordListToRowList_head(aWord string, lemmaList []lemmaWordStruct, maxNumRow int, js_function string) {
	//swAnto:= (aWord=="schrift") 
	//if swAnto {  fmt.Println("\nfun_wordListToRowList_head(aWord=", aWord, " lemmaList=", lemmaList)	}
	
	numberOfRows = len(inputTextRowSlice)
	
	maxNumRow0 := maxNumRow;	
	listWords:= ""
	//hd_tr := "<TABLE>\n"	
	hd_tr := ""; 
	
		preL :=""	
		preW :=""
		// line2:=""   // 14giugno
		
	listIxRR := make([]int,0, 110)
	
	for _, lemmaX := range(lemmaList) { 	
		ixWord:= lemmaX.lw_ixWordUnFr
		var xWordF     = uniqueWordByFreq[ixWord]   
		
		var ixFromList = xWordF.uIxWordFreq 
		var ixToList   = ixFromList + xWordF.uTotRow;
		var maxTo1     = ixFromList + maxNumRow; 		
		
		if ixToList > maxTo1        { ixToList = maxTo1; }
		if ixToList > numberOfWords { ixToList = numberOfWords; }		
		
		listWords += " " +  xWordF.uWord2
		for z:=0; z < len(xWordF.uLemmaL); z++  {
			ixL1:= xWordF.uIxLemmaL[z]
			LeS := lemmaSlice[ixL1]
			newL:= xWordF.uLemmaL[z]
			if newL != LeS.leLemma {
				continue;  // error 
			}
			
			if newL != lemmaX.lw_lemma2 { continue }   //     newCode   ( sto cercando frasi doppie 
			//newT:= xWordF.uTranL[z]		//	anto1  .uTranL
			newT:= LeS.leTran
			//newP:= xWordF.uPara[z]
			if newL == preL { 
				newL=""
				newT=""
				if preW != xWordF.uWord2 { 
					hd_tr += xWordF.uWord2 + " ";  
					preW = xWordF.uWord2
				} 
			} else { 
				/**
				if newP == "" {
				    newP = newL;
					// line2 = "&nbsp;&nbsp;&nbsp;&nbsp;"  	// 14giugno
				} else {
					// line2 = "<br>"  // 14giugno 
				}
				***/
				preL = newL 
				
				
				if hd_tr != "" { hd_tr += "\n" 	}  // 14giugno 
				hd_tr += " :lemma="  + newL + " :tran=" + newT  + " :wordsInLemma=" +	xWordF.uWord2 + " " // 14giugno 
				preW =  xWordF.uWord2
			} 
			
		}

		list3:= fun_wordListToRowList_dett(ixWord, maxNumRow0)	
		for n1:=0; n1 < len(list3); n1++ {
			listIxRR  = append(listIxRR, list3[n1] )
			//if (swAnto && xWordF.uWord2== "schrift")  { fmt.Println( "anto5 list3=", list3[n1] ) }  
		}			
				
	}// end of for lemma 
	//-------
	
	//hd_tr += "</td></tr>\n"    // 14 giugno
	hd_tr += "\n"
	
	sort.Ints(listIxRR) 
	
	//fmt.Println( "xxxxxxxxxxxxxx  len(listIxRR)=",   len(listIxRR) )
	
	preIxRR_2:= 999999999 
	ixRR_2:=0
	ixRR  :=0
	//nFF   :=0 
	outS1:= ""
	nOut:=0
	new_rIdRow :=""
	numMioRow:=0
	
	for n1:= 0; n1 < len(listIxRR); n1++  {
		ixRR_2 = listIxRR[n1]
		
		//fmt.Println( "xxxxxxxxxxxxxx  n1=", n1 , " ixRR_2=", ixRR_2," preIxRR_2=", preIxRR_2)
		
		if (ixRR_2 == preIxRR_2) { continue;} 
		preIxRR_2 = ixRR_2; 
		ixRR = ixRR_2 % 100000; 
		
		//nFF  = ( ixRR_2 - ixRR ) / 100000 
		
		//fmt.Println( "xxxxxxxxxxxxxx  ixRR=", ixRR,  " numberOfRows=",  numberOfRows)
		
		if ixRR >= numberOfRows { continue;} // actually there  must be some error here 		
		rline := inputTextRowSlice[ixRR]
		rowX := cleanRow(rline.rRow1)	
		
		//if swAnto {numMioRow++; fmt.Println( "xxxxxxxxxxxxxx   ixRR_2=", ixRR_2," ixRR=" , ixRR, " maxNumRow=", maxNumRow, " rowX=", rowX) } 
		
		if ((rowX =="") || (rowX == LAST_WORD)) { 
			continue 
		}		
		
		if rline.rixGroup < 0 { 
			new_rIdRow = "- " + strconv.Itoa( rline.rixBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		} else {
			new_rIdRow = lista_gruppiSelectRow[ rline.rixGroup ].rG_group + " " + strconv.Itoa( rline.rixBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		}	
		outS1 += "<br>" + strconv.Itoa( SEL_EXTR_ROW ) + "|" + new_rIdRow   + "|" + strconv.Itoa( ixRR) + "|"   + rowX + "|" + rline.rTran1; 
	  //outS1 += "<br>" + strconv.Itoa( SEL_EXTR_ROW ) + "|" + rline.rIdRow + "|" + strconv.Itoa( ixRR) + "|"   + rowX + "|" + rline.rTran1; 
		
		//fmt.Println( "  outS=" , outS1); 
		nOut++
		if (nOut >= maxNumRow) {
			break;
		}
		
	} 	// end for n1
	
	fmt.Println( "anto6  numMioRow=", numMioRow); 
	
	header:= "<HEADER>\n" + "<WORD>" + aWord + ",L:" + strings.TrimSpace(listWords) + "</WORD>"
	//header += "<TABLE style=\"padding: 0 2em;border:1px solid black;\">\n" + hd_tr + "</TABLE>\n"  // 14 giugno
	header += hd_tr   // 14giugno
	header += "</HEADER> \n"
	
	go_exec_js_function( js_function, header + outS1 ); 	
	
} // end of fun_wordListToRowList_head		

//-------------------------------
func fun_wordListToRowList_dett( ixWord int, maxNumRow int) []int {
		
		var xWordF     = uniqueWordByFreq[ixWord]   
		
		var ixFromList = xWordF.uIxWordFreq 
		var ixToList   = ixFromList + xWordF.uTotRow;
		var maxTo1     = ixFromList + maxNumRow; 		
		
		if ixToList > maxTo1        { ixToList = maxTo1; }
		if ixToList > numberOfWords { ixToList = numberOfWords; }		
				
		/*		
		here are scanned all the rows containing the word required (with index ixWord) 
		for each line totMinRow is the number of words in the line with lower reference than the required word (ie. not studied yet)	
		*/
		
		listIxRR := make([]int,0, ixToList-ixFromList)
		if (ixFromList < 1) { ixFromList=0;}
		for n1 := ixFromList; n1 < ixToList; n1++  {
			wS1 := wordSliceFreq[n1] 
			//fmt.Println("fun_wordListToRowList_dett() wS1=", wS1);
			listIxRR = append( listIxRR, wS1.wSwSelRowR * 100000 +  wS1.wIxRow) 			
		} 	
		return listIxRR

} // end fun_wordListToRowList


//-----------------------------------------------------

func bind_go_passToJs_rowList(inpBegRow int, maxNumRow int, js_function string) {
	// lista tutte le frasi richieste ( numero della prima frase, numero di frasi) 
	var ixFromList = inpBegRow 
	
		
	var outS1 string;
	numOut:=0 
	
	new_rIdRow :=""
	
	for ixRR := ixFromList; ixRR < len(inputTextRowSlice); ixRR++  {
		rline := inputTextRowSlice[ixRR]
		
		rowX := cleanRow(rline.rRow1)	
		
		if ((rowX =="") || (rowX == LAST_WORD)) { 
			continue 
		}		
		numOut++
		if (numOut > maxNumRow)  { break }
	
		if rline.rixGroup < 0 { 
			new_rIdRow = "- " + strconv.Itoa( rline.rixBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		} else {
			new_rIdRow = lista_gruppiSelectRow[ rline.rixGroup ].rG_group + " " + strconv.Itoa( rline.rixBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		}	
		
		outS1 += "<br>" + strconv.Itoa( SEL_EXTR_ROW ) + "|" + new_rIdRow + "|" + strconv.Itoa( ixRR) + "|"   + rowX + "|" + rline.rTran1; 
	} 
	
	go_exec_js_function( js_function, outS1 ); 	
			
} // end of bind_go_passToJs_rowList

//---------------------
func tranFromIxLemma( xWordF wordIxStruct, ix int) string {
	ixL1 := xWordF.uIxLemmaL[ix]
	return lemmaSlice[ixL1].leTran
}
//-------------------------------
func listStringLemmaSlice_Tran( xWordF wordIxStruct) string {
	listS:=""
	for _, ixL1:= range xWordF.uIxLemmaL { 
		listS += lemmaSlice[ixL1].leTran + wSep
	}		
	return listS
}

//------------------
func bind_go_passToJs_rowWordList(numIdOut string, ixRR int, js_function string) {
	//  lista di tutte le parole di una frase	
	
	if ixRR >= len(inputTextRowSlice) { ixRR = len(inputTextRowSlice) - 1 }
	
	rowX := inputTextRowSlice[ixRR]

	outS1:= numIdOut + "," + strconv.Itoa(ixRR) + "," + strconv.Itoa(rowX.rNumWords) +"," + strconv.Itoa( len(rowX.rListIxUnF)) + endOfLine
	
	for w:=0; w < len(rowX.rListIxUnF); w++ {  
		if rowX.rListFreq[w] < 1 {continue}     // the entry is allocated, but unused  
		ixWord := rowX.rListIxUnF[w] 
		if ixWord < 0 { continue}
		xWordF := uniqueWordByFreq[ixWord] 
		 
		// ??anto2 uTranL
		row11 := xWordF.uWord2 + "," + strconv.Itoa(xWordF.uIxUnW) + "," + 
			strconv.Itoa(xWordF.uTotRow)  + ";" + 
			fmt.Sprint( strings.Join(xWordF.uLemmaL, wSep) ) + 
			//";" + fmt.Sprint( strings.Join(xWordF.uTranL, wSep)) + 
			";" + listStringLemmaSlice_Tran(xWordF) +
			endOfLine 
		
		outS1 += row11; 
		
	}	

	go_exec_js_function( js_function, outS1 ); 				

} // end of bind_go_passToJs_rowWordList 

//---------------------------------------------------------------
func bind_go_passToJs_word_known(ixWord int, yesNo int, knowCtr int, js_function string) {

	if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}	
		
	var xWordF = uniqueWordByFreq[ixWord]  
	aWord:= xWordF.uWord2; 
	ixFreq := xWordF.uIxUnW     
	ixAlpha:= xWordF.uIxUnW_al   
	if ixFreq != ixWord { fmt.Println("ERRORE in bind_go_passToJs_word_known() ixWord not equal to ixFreq ", aWord, " ixWord=", ixWord, " ixFreq=", ixFreq )   }
	
	if yesNo == 0 {	
		uniqueWordByFreq[ixWord].uKnow_yes_ctr = knowCtr
	} else {
		uniqueWordByFreq[ixWord].uKnow_no_ctr  = knowCtr 
	} 
	
	uniqueWordByAlpha[ixAlpha].uKnow_yes_ctr = uniqueWordByFreq[ixWord].uKnow_yes_ctr 
	uniqueWordByAlpha[ixAlpha].uKnow_no_ctr  = uniqueWordByFreq[ixWord].uKnow_no_ctr 
	
	
	outS1 := fmt.Sprint("ixWord=", ixWord, " ", aWord, ", \t yesNo=", yesNo, ", knownYes=", xWordF.uKnow_yes_ctr,  ", knownNo=", xWordF.uKnow_no_ctr )  

	go_exec_js_function( js_function, outS1 ); 	
	
} // end of bind_go_passToJs_word_known

//-----------------------------------------------------------

func bind_go_write_lang_dictionary( langAndVoiceName  string) { 
		
		outFileName		:= FOLDER_INPUT_OUTPUT + string(os.PathSeparator) + FILE_language;  		
		
		//fmt.Println("bind_go_write_lang_dictionary file=" + outFileName +"\n\t",  langAndVoiceName );  	
		
		if langAndVoiceName[0:9] != "language=" {
			fmt.Println("bind_go_write_lang_dictionary() ERROR =>" +  langAndVoiceName[0:9] + "<==");  
			return 
		}	
		
		f, err := os.Create( outFileName )
		check(err)
		defer f.Close();

		_, err = f.WriteString( langAndVoiceName )
		check(err)

		f.Sync()
		
} // end of bind_go_lang_word_dictionary 	

//-----------------------------------------------------------
func bind_go_write_word_dictionary( listGoWords string) { 
		
		// listGoWords = list of NEW translated words
		
		//fmt.Println("bind_go_write_word_dictionary ",  listGoWords  );  		
		
	
		
		if len(listGoWords) < 1 { return }
		if len(listGoWords) > 9 {
			if listGoWords[0:9] == "language=" { return }
		}
			
		lemmaTranStr := ""; //  "__" + outFileName + "\n" + "_lemma	_traduzione"
		lastNumDict++; 
		lemmaTranStr += split_ALL_word_dict_row( listGoWords )	
		
		sort_lemmaTran2();  // sort_lemmaX2 in write_word_dict... utilizzabili già in questo run 	
		
		/**------------------
		f, err := os.Create( outFileName )
		check(err)
		defer f.Close();

		_, err = f.WriteString( lemmaTranStr );  
		check(err) 

		f.Sync()
		**/
		
		//----------------------
		
		//sort_lemmaTran();  // sort_lemma2 in write_word_dict... utilizzabili già in questo run 
		
		rewrite_LemmaTranDict_file() 
		
		
} // end of bind_go_write_word_dictionary 	

//-----------------------------------------------------------

func split_one_word_dict_row( row1 string ) (string, int, []string, []string) {

	// eg. einem;14; ein§einem§einer	
	
	lemmaLis := make( []string,0,0 )
	tranLis  := make( []string,0,0 )
	
	if row1 == "" { return "",	-1,	lemmaLis, tranLis }                           
	
	//row1:= strings.ReplaceAll( strings.ReplaceAll( row0, ....  parentesi quadre 
	
	var field = strings.Split( row1, ";");     // eg. einem;  14; ein§einem§einer; uno§uno§uno
											   //	 field[0]; 1;               2; 3 
	
	lemmaLis = strings.Split(field[2],wSep);   // eg. ein , einem,   einer       
	tranLis  = strings.Split(field[3],wSep);   //     uno , uno,     uno  
	
	ix1, err := strconv.Atoi( field[1] )
	if err != nil { return "",	-1,	lemmaLis, tranLis }                           //error
	 
	return field[0],ix1, lemmaLis, tranLis     // eg. return einem; 14, [ein , einem,   einer], [uno , uno,     uno] 
		
} // end of	split_one_Word_Dict_row(	

//-------------------------------

func split_ALL_word_dict_row(  strRows string) string {
	fmt.Println( "ANTONIO xxxxxxxxxxxxxxxxxxxxxxxxxxxx  split_ALL_word_dict_row( strRows=", strRows); 
	// eg. einem;14 ; ein einem einer ;  a uno uno;	  ==> word ; ix : list of lemmas ; list of translations	
	
	lemmaTranStr := ""
	
	lines := strings.Split( strRows, "\n");	
	
	var ele1 lemmaTranStruct     
	/**
	//--------------------------
		type lemmaTranStruct struct {
			dL_lemmaCod  string 
			dL_lemma2    string  
			dL_numDict   int	  
			dL_tran      string
		} 
	**/
	
	lenIx1:= len(uniqueWordByFreq)
	
	
	//------------  solo per controllare----	
	/**
	for ix1:=0;  ix1 < 28; ix1++ {   
		fmt.Println( "ANTONIO error7 check inizio ix1=", ix1, " uniqueWordByFreq[ix1]=" , uniqueWordByFreq[ix1] )
	}
	**/	
	/***
	fmt.Println( "ANTONIO xxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	for ix1:=(lenIx1-5);  ix1 < lenIx1; ix1++ {   
		fmt.Println( "ANTONIO error7 check fine   ix1=", ix1, " uniqueWordByFreq[ix1].ixUnW ="  , uniqueWordByFreq[ix1].ixUnW )
	}
	***/
	
	//-------------- fine solo per controllare-----
	
	for z:=0;  z < len(lines); z++ {   
		
		//fmt.Println("split_ALL_word_dict_row() 1 ", lines[z] )	
		
		_, ix1, lemmaLis,tranLis := split_one_word_dict_row( lines[z] )
		if ix1 < 0 { continue }		
		
		if ix1 >= lenIx1 { 
			fmt.Println("error7 1 len(uniqueWordByFreq)=", len(uniqueWordByFreq), " ix1=", ix1 ,  " lines[z=", z, "]=", lines[z] )
			continue 
		}
		//---------------
		//fmt.Println("   ix1=", ix1, " \tlemmaLis=", lemmaLis, "\ttranLis=", tranLis )  
		//---------------
		if uniqueWordByFreq[ix1].uIxUnW != ix1 {	
			fmt.Println("error7 2 len(uniqueWordByFreq)=", len(uniqueWordByFreq), " ix1=", ix1 ,  " lines[z=", z, "]=", lines[z] )
			continue 
		}  // error	
		
		//swUpdList[ix1] = true 
		//numUpd++  		
		ixAlfa := uniqueWordByFreq[ix1].uIxUnW_al  	
		
		len1:= len(uniqueWordByFreq[ix1].uLemmaL)
		
		//fmt.Println("\tsplit_ALL_word_dict_row() 2 ", " ix1=", ix1, " len1=", len1, ", len(lemmaLis)=", len(lemmaLis), ",     len(tranLis)=",  len(tranLis) )
		if len1 != len(lemmaLis) { 
			fmt.Println("split_ALL_word_dict_row() 1 ", lines[z] )
			for mio1:=0; mio1 < len1; mio1++ {
				fmt.Println("\tsplit_ALL_word_dict_row() 2.1 ", "  uniqueWordByFreq[ix1].uLemmaL[", mio1,"] =>" + uniqueWordByFreq[ix1].uLemmaL[mio1] + "<==")
			}   
		}   
		if len1 != len(lemmaLis) { continue }               // error 
		if len1 != len(tranLis)  { continue }               // error 
		
		//---------------------
		oneW := uniqueWordByFreq[ix1]
		for m:=0; m < len1; m++ {
			mLemm := strings.TrimSpace( lemmaLis[m] )
			if mLemm == oneW.uLemmaL[m] {
				ixLe:= oneW.uIxLemmaL[m]
				lemmaSlice[ixLe].leTran = strings.TrimSpace( tranLis[m] )	
			} else {
				for m2:=0; m2 < len(oneW.uLemmaL); m2++ {
					if mLemm == oneW.uLemmaL[m2] {
						ixLe:= oneW.uIxLemmaL[m]
						lemmaSlice[ixLe].leTran = strings.TrimSpace( tranLis[m2] )	
						break;
					}
				} 
			}
		} 
		//--------------------
		for m:=0; m < len1; m++ {
			mLemm := strings.TrimSpace( lemmaLis[m] )
			mTran := strings.TrimSpace( tranLis[m] 	)	
			//uniqueWordByFreq[ix1].uTranL[m]      = mTran 		//??anto3 uTranL    già fatto con lemmaSlice[]
			//uniqueWordByAlpha[ixAlfa].uTranL[m]  = mTran 		
			uniqueWordByAlpha[ixAlfa].uLemmaL[m] = mLemm	
			
			uniqueWordByAlpha[ixAlfa].uKnow_yes_ctr = 0 
			uniqueWordByAlpha[ixAlfa].uKnow_no_ctr  = 0 
			
			lemmaTranStr += "\n" + mLemm + "|" + mTran 	
			
			//fmt.Println("\tsplit_ALL_word_dict_row() 3 lemma n.", m, " lemma="+ mLemm + " trad.=" + mTran 	)
			
			ele1.dL_lemmaCod = newCode(mLemm) 
			ele1.dL_lemma2 = mLemm 
			ele1.dL_numDict = lastNumDict
			ele1.dL_tran  = mTran           ////cigna1_1 
			if mTran != "" { 		
				dictLemmaTran = append( dictLemmaTran, ele1 ) 
			}
		}	
	} // end of z 
	
	return lemmaTranStr
	
} // end of split_ALL_word_dict_row(

//------------------------------------------
/***
	//---------------------  
		oneW := uniqueWordByFreq[ix1]
		for m:=0; m < len1; m++ {
			mLemm := strings.TrimSpace( lemmaLis[m] )
			if mLemm == oneW.uLemmaL[m] {
				ixLe:= oneW.uIxLemmaL[m]
				lemmaSlice[ixLe] = strings.TrimSpace( tranLis[m] )	
			} else {
				for m2:=0; m2 < len(oneW.uLemmaL); m2++ {
					if mLemm == oneW.uLemmaL[m2] {
						ixLe:= oneW.uIxLemmaL[m]
						lemmaSlice[ixLe] = strings.TrimSpace( tranLis[m2] )	
						break;
					}
				} 
			}
		}
		//--------------------

***/
func TOGLIupdate_words_tranWWW( newTranWordStr string) string {
	fmt.Println("GO update_words_tranWWW newTranWordStr=", newTranWordStr);

	// eg. 	seinem;17;sein§seine;essere§il suo     // word; ixUnique...;  lemmas list separati da §; translations list separ. da §]  
	if newTranWordStr == "" {return ""}
	//												listNewTranWords += "\n" + word1 + "," + ix1 + "," + wordTran ;   // new line for dictionary 
	wordTranList := strings.Split(newTranWordStr, "\n") 	
	
	var lemma1 string; var ixS string //  var wordTran string
	var ixAlfa int
	
	
	//thisDictList :=  make( []wDictStruct,  len(wordTranList),  len(wordTranList) )
	//var oneDict wDictStruct
	
	lenU := len( uniqueWordByFreq)  
	
	swUpdList:= make([]bool,  lenU, lenU)
	numUpd:=0
	for i:= 0; i < len(wordTranList); i++ {
		//sw := (strings.Index( wordTranList[i], "sein") >= 0 )
		//if sw {fmt.Println( "\n" + "wordTranList[",i,"]=" , wordTranList[i] ) }
		
		if wordTranList[i] == "" { continue }
		col0:= strings.Split( wordTranList[i], ";")	
		if len(col0) != 2 { continue}  
		col1:= strings.Split( col0[0], ",")
		if len(col1) != 2 { continue}  
		
		lemma1   = strings.TrimSpace( col1[0] )
		ixS      = col1[1]
		//wordTran = strings.TrimSpace( col0[1] )  // ??anto
		ix1, err := strconv.Atoi(ixS)
		if err != nil { continue} 		
		
		if uniqueWordByFreq[ix1].uIxUnW != ix1 { continue } 
		
		swUpdList[ix1] = true 
		numUpd++  
		
		ixAlfa = uniqueWordByFreq[ix1].uIxUnW_al   		
		lemmaLis := uniqueWordByFreq[ix1].uLemmaL
		
		for m:=0; m < len(lemmaLis); m++ {
			if lemma1 == lemmaLis[m] {
				//uniqueWordByFreq[ix1].uTranL[m]      = wordTran 		//???anto5 .uTranL		( tutta la func. da togliere )
				//uniqueWordByAlpha[ixAlfa].uTranL[m]  = wordTran 			//???anto5 .uTranL		( tutta la func. da togliere )		
				uniqueWordByAlpha[ixAlfa].uLemmaL[m] = lemma1
				//if ix1 == 17 {  fmt.Println("\t\tm=", m, ", tran[]=", uniqueWordByFreq[ix1].uTranL[m]  ) }
				break	
			}
		}	
		
		//if ix1 == 17 {  fmt.Println( ix1, " ", swUpdList[ix1] , " uniqueWordByFreq[ix1]=",  uniqueWordByFreq[ix1] )   	}
		
	} // end of for i 
	//-------

	//fmt.Println("aggiornati ",  numUpd , " uniqueWordByFreq"  , " lenU=" , lenU)
	
	var newStr = ""
	
	for ix1:= 0; ix1 < lenU; ix1++ {
		if swUpdList[ix1] == false { continue }
		
		//  fmt.Println( "XXXXXXXXXXXXXXXX  ix1=", ix1, " ", swUpdList[ix1] , " uniqueWordByFreq[ix1]=",  uniqueWordByFreq[ix1] )   	
		  
		sU := uniqueWordByFreq[ix1]	
		lemmaLis := sU.uLemmaL
		//tranLis  := sU.uTranL              // ??anto6   		( tutta la func. da togliere )
		/**
		if ix1 == 17 {
			fmt.Println( " uniqueWordByFreq[ix1]=",  uniqueWordByFreq[ix1] )
			fmt.Println("\tlemmaL =",uniqueWordByFreq[ix1].uLemmaL, "\n\ttranL=",  uniqueWordByFreq[ix1].uTranL ) 
		}
		**/
		
		lemmS:= ""; tranS:=""
		for m:=0; m < len(lemmaLis); m++ {
			lemmS += "," +lemmaLis[m]
			//tranS += "," +tranLis[m]  //??anto6
			//if ix1 == 17 {fmt.Println( " lemmS=" + lemmS, "  \t ", " tranS=", tranS) }
		}	
		lemmS += " " ; tranS += " "
		newStr += sU.uWord2 + ";" + strconv.Itoa(ix1) + ";" +
				strings.TrimSpace( lemmS[1:] ) + ";" + strings.TrimSpace(tranS[1:]) + ";" + "\n"	
		//if ix1==17 {  fmt.Println("newSTR=" , newStr) }		
	}
	//fmt.Println( "\nupdate words \n", newStr, "\n---------------------------") 
	return newStr
	
} // end of TOGLIupdate_words_tranWWW

//-----------------------------------------------------------

//-------------------------------------
func bind_go_write_row_dictionary( listGoRows string) {
		/**
		2;Primo capitolo
		4;Gustav Aschenbach o von Aschenbach, come ha fatto sin dai suoi cinquant'anni
		5;compleanno, era il suo nome ufficiale, era l'uno
		**/
		
		//fmt.Println("bind_go_write_row_dictionary ()  esegue  update_rowTransLATION () ")
		
		update_rowTranslation(   strings.Split(listGoRows,"\n") ) ;
		
		//fmt.Println("bind_go_write_row_dictionary ()  esegue  writeTextRowSlice () ")
		
		writeTextRowSlice()  // after some new translated rows  the dictR file is rewritten  

} // end of  bind_go_write_row_dictionary 	

//--------------------------------------------
func writeTextRowSlice() {		

		nout:=0  
		lines:= make([]string, 0, 10+len( inputTextRowSlice) )
		
		for z:=0; z < len( inputTextRowSlice); z++ {
			rS1 := inputTextRowSlice[z]
			
			lines = append(lines, rS1.rIdRow + "|O|" + rS1.rRow1 )  
			
			if rS1.rTran1 != "" {
				lines = append(lines, rS1.rIdRow + "|T|" + rS1.rTran1)  
			}
			
			//fmt.Println("writeTextRowSlice () z=", z, "  lines[]= ", lines[ len(lines)-1 ] )
			nout++
		}  
		
		// riscrivi tutte le righe con text origine e traduzione  (ogni file ha data e ora nel nome, bisogna leggere l'ultimo, e ogni tanto cancellare i file vecchi)	
		
		currentTime := time.Now()	

		last_written_dict_rowFile = "dictR" + currentTime.Format("20060102150405") + ".txt"			
		//-----------------------
		outF1 		:= FOLDER_O_arc_TRAN_rows + string(os.PathSeparator)    		
		outFileName := outF1 + last_written_dict_rowFile				
		fmt.Println("wrote ", nout , " lines on ", outFileName ) 	
		writeList( outFileName, lines )	
		//-----------------
		outF12 		:= FOLDER_IO_lastTRAN + string(os.PathSeparator)    		
		outFileName2 := outF12 + FILE_last_updated_dict_rows				
		fmt.Println("wrote ", nout , " lines on ", outFileName2 ) 	
		writeList( outFileName2, lines )	
		//-----------------
		
		
} // end of bind_go_write_row_dictionary 	


//-----------------------------------------------

//--------------------------------------------
func writeUnaTantumNuovoFile() {		

		nout:=0  
		lines:= make([]string, 0, 10+len( inputTextRowSlice) )
		
		
		var nuovoId string; 
		
		var newCode string; 
		var preCode string 
		ctr:=0
		preCode=""
		for z:=0; z < len( inputTextRowSlice); z++ {
			rS1 := inputTextRowSlice[z]
			if (z >=0    ) && (z<=  5807) { newCode="1"} 
			if (z >= 5808) && (z<= 13464) { newCode="2"} 	
			if (z >=13465) && (z<= 14870) { newCode="3"} 
			if (z >=14871) && (z<= 22620) { newCode="4"} 
			if (z >=22621) && (z<= 25883) { newCode="5"} 
			if (z >=25884) && (z<= 35992) { newCode="6"} 
			if (z > 35992){newCode="07"}
			if (newCode != preCode) { preCode=newCode; ctr=0}
			ctr++
			nuovoId = newCode + "_" + strconv.Itoa(ctr) 
			if rS1.rTran1 == "" {
				lines = append(lines, nuovoId + "|O|" + rS1.rRow1 + "\n")  
			} else {
				lines = append(lines, nuovoId + "|O|" + rS1.rRow1 + "\n" + nuovoId + "|T|" + rS1.rTran1 + "\n")  
			}
			nout++
		}  
		
		// riscrivi tutte le righe con text origine e traduzione  (ogni file ha data e ora nel nome, bisogna leggere l'ultimo, e ogni tanto cancellare i file vecchi)	
		
		currentTime := time.Now()	

		last_written_dict_rowFile = "NUOVO_RINUMERATO_dictR" + currentTime.Format("20060102150405") + ".txt"			
		
		outF1 		:= FOLDER_OUTPUT + string(os.PathSeparator)    		
		outFileName := outF1 + last_written_dict_rowFile		
		
		fmt.Println("wrote ", nout , " lines on ", outFileName )  
	
		writeList( outFileName, lines )		
		
} // end of bind_go_write_row_dictionary 	

//---------
//--------------------

func update_rowTranslation( rowDictRow [] string)  {
	/*
	rowDictRow is filled in javascript, each line: rowIndex;row Translation   
	*/
	var len1 = len(inputTextRowSlice)
	
	//fmt.Println("update_rowTraslation() len(rowDictRow) =",len(rowDictRow), " = ", rowDictRow)
	
	//ion() len(rowDictRow) =  [ 2 206(2_206 211|211|1. Introduzione Preistoria 2 207(2_207 212|212|La prei
	
	for z:=0; z < len(rowDictRow); z++ {  		
		/**
		2;Primo capitolo
		**/
		row1dict := rowDictRow[z] 	
		if row1dict == "" { continue }	
		rowCols	 := strings.Split(row1dict, "|")
		
		//fmt.Println("\trowCols= len=", len(rowCols), " rowcols=",    rowCols) 
		
		if len(rowCols) < 3 { 
			fmt.Println("update_rowTraslation() len(rowDictRow) z=", z, " row1dict=", row1dict, " len(rowCols)=", len(rowCols))
			continue 
		}
		
		idS   := rowCols[0]   
		ixS   := rowCols[1] 
			
		tranS := rowCols[2]
		
		//fmt.Println("\tidS=", idS, "  ixS=", ixS,  " tranS=", tranS ) 
		
	
		ixRow, err := strconv.Atoi(  strings.TrimSpace( ixS) )		
		
		//fmt.Println("\tixRow=", ixRow, " err=", err)
		
		if err != nil {
			return 
		}
		if ixRow >= len1 { return }  // error 	

		//fmt.Println("\tinputTextRowSlice[ixRow].rIdRow = ", inputTextRowSlice[ixRow].rIdRow )
		
		if strings.Index( idS, inputTextRowSlice[ixRow].rIdRow) < 0  {
			// error
			fmt.Println("error func update_rowTranslation () inputTextRowSlice[ixRow=", ixRow, "].rIdRow = ", inputTextRowSlice[ixRow].rIdRow,   " idS=",  idS )
			continue
		}  
		inputTextRowSlice[ixRow].rTran1 = strings.TrimSpace(tranS ) 
		
		//fmt.Println("update_rowTraslation() row1dict=", row1dict, " \n\t ixRow=", ixRow, " inputTextRowSlice[ixRow]=", inputTextRowSlice[ixRow] )
		
	} // end for z
	
} // end of update_rowTranslation

//---------------------------------------------------------

func endBegin(wh string) {
	//fmt.Println("func endBegin (", wh,")")
	if sw_stop { 
		fmt.Println("\nXXXXXXXX  error found XXXXXXXXXXXXXX\n"); 
	}	
	sw_begin_ended = true 		
}
//---------------
func begin_GO_HTML_Talk() { 	
	fmt.Println("func begin_GO_HTML_Talk"); 
	setHtmlEnv();	
}
//---------------
func begin() { 	
	fmt.Println("func begin"); 
	
	read_all_files() 
	
	if sw_stop { endBegin("1"); return }
	
	//-------------------------------------
	
	build_and_elab_word_list()
	
	if sw_stop { endBegin("2"); return }	

	//stat_useWord();	

	if sw_stop { endBegin("3"); return }	
	
	if sw_rewrite_wordLemma_dict { rewrite_word_lemma_dictionary() }
	
	if sw_stop { endBegin("4"); return }	
		
	if sw_stop { endBegin("5"); return }	
		
	/**
	if sw_read_numbered_text {
		bind_go_passToJs_read_wordsToLearn("js_go_file_words_to_learn_read") 	
	}
	**/		
	
	read_wordsToLearn()
	
	
	endBegin("6")
	
	//writeUnaTantumNuovoFile()
	
	//-------------------------------------------
	
	numberOfRows = len(inputTextRowSlice)
	
	buildStatistics()
		
	mainNum := strconv.Itoa(numberOfUniqueWords) +";" + strconv.Itoa(numberOfWords) + ";" + strconv.Itoa(numberOfRows) +	
		   ";))"	
		   //";" + "level " + msgLevelStat + "))" 		
		
	go_exec_js_function("js_go_showReadFile", mainNum + showReadFile);  
			
	if sw_stop { 
				fmt.Println("UI is ready ( run stopped because of some error)")
	} else {
		//go_exec_js_function("js_go_ready", prevRunLanguage + ":mainpage_value=" + last_mainpage_valueString)  // +"<file>" + prevRunListFile); 
		
		go_exec_js_function("js_go_ready", prevRunLanguage )
		//log.Println("UI is ready")
		fmt.Println("UI is ready")
	}	
	fmt.Println("\nEND of begin \n") 
	fmt.Println(cyan("\nREADY"), "\n") 
	
}// end of begin	
 
//-------------------------------------
var last_run_extrRow = ""
//----------------------------------
func build_and_elab_word_list() {

	fmt.Println("\n", cyan("BUILD WORD LIST") )
	
	wordSliceAlpha    = nil 
	wordSliceFreq     = nil
	uniqueWordByFreq  = nil
	uniqueWordByAlpha = nil
	
	buildWordList() 	   	
	
	if sw_stop { endBegin("4"); return }
	
	elabWordList() 
	
} // end of build_and_elab_word_list()
 
//------------------------------------
func setHtmlEnv() {	
	fmt.Println("func setHtmlEnv:  start load html")
    // load file html 	
	
	html_path = getCompleteHtmlPath( parameter_path_html ) 
	            
	fmt.Println("path html        = " + html_path)
	
	ui.Load("file:///" + html_path + string(os.PathSeparator) + "wordsByFrequency.html" ); 
	
	fmt.Println("\n", "func setHtmlEnv: wait for html ( javascript function js_call_go()", "\n")  
	
} // end of setHtmlEnv
//--------------------------------------------------------
//-------------------------
func getCompleteHtmlPath( path_html string) string {
	
	//curDir    := "D:/ANTONIO/K_L_M_N/LINGUAGGI/GO/_WORDS_BY_FREQUENCE/WbF_prova1_input_piccolo
	 
	curDir, err := os.Getwd()
    if err != nil {
		fmt.Println("setHtmlEnv() 3 err=", err )
        //log.Fatal(err)
    }	
				
	fmt.Println("curDir           = " + curDir ); 
	
	curDirBack  := curDir
	k1:= strings.LastIndex(curDir, "/") 
	k2:= strings.LastIndex(curDir, "\\") 
	if k2 > k1 { k1 = k2 } 
	curDirBack = curDir[0:k1] 	
	
	var newPath string = ""
	if strings.Index(path_html,":") > 0 {
		newPath = path_html
	} else if path_html[0:2] == ".." {
		newPath = curDirBack  + path_html[2:] 
	} else {
		newPath = curDir + path_html
	}
	return newPath 
} 
//------------------------
//----------------------
func putFileError( msg1, inpFile string) {
	err1:= `document.getElementById("id_startwait").innerHTML = '<br><br> <span style="color:red;">§msg1§</span> <span style="color:blue;">§inpFile§</span>';` ; 		
	err1 = strings.ReplaceAll( err1, "§msg1§", msg1 ); 	 
	err1 = strings.ReplaceAll( err1, "§inpFile§", inpFile); 	
	ui.Eval( err1 );	
}   

//-----------------------------------

func test_all_folder() {	
	test_folder_exist( FOLDER_INPUT         ); if sw_stop { return } 	
	test_folder_exist( FOLDER_OUTPUT        ); if sw_stop { return } 	
	test_folder_exist( FOLDER_INPUT_OUTPUT  ); if sw_stop { return } 
	
	test_folder_exist( FOLDER_I_lemma )     ;  if sw_stop { return } 	
	test_folder_exist( FOLDER_I_paradigma)  ;  if sw_stop { return } 	
	test_folder_exist( FOLDER_IO_lastTRAN ) ;  if sw_stop { return } 	
	
	test_folder_exist( FOLDER_O_ARCHIVE        ) ;  if sw_stop { return } 	
	test_folder_exist( FOLDER_O_arc_TRAN_rows  ) ;  if sw_stop { return } 		
	test_folder_exist( FOLDER_O_arc_TRAN_words ) ;  if sw_stop { return } 	
	test_folder_exist( FOLDER_O_arc_TO_learn   ) ;  if sw_stop { return } 	
} // end of test_all_folders

//------------------
func getFileByteSize( path1 string,   fileName string) int {
	path2:=""
	if path1 != "" {
		path2 = path1 + string(os.PathSeparator) 
	} 	
	fileN := path2 + fileName 
	fileInfo, _ := os.Stat( fileN )  
	
	fmt.Println("getFileByteSize fileN=", fileN, " fileInfo = ", fileInfo) 
	if fileInfo == nil { return 0 }
	return int( fileInfo.Size() )
} // end of getFileByteSize

//--------------
func myOpenRead( path1 string,   fileName string,   descr string,  func1    string) (*os.File, int) {
	path2:="";
	path10:=""
	if path1 != "" {
		path10 = " in " + cyan(path1)
		path2 = path1 + string(os.PathSeparator) 
	} 	
	fileN := path2 + fileName 
	
	fmt.Println("\n" + yellow("open file"),  green(fileName) , path10 )
	
	sizeByte:= getFileByteSize(path1,fileName)
	readFile, err := os.Open( fileN )  
    if err == nil {				
		fmt.Println( "\t", "size: ", sizeByte, " bytes" )	
		return readFile, sizeByte
	}
	msg1_Js:= `il file "` + fileN + `" (` + descr + " " + func1 + ")" + " non esiste"
		
	errorMSG = `<br><br>il file ` + 
				`<span style="font-size:0.7em;color:black;">(`	+ descr + `)</span>` +
				`<br><span style="color:blue;" >` + fileName + `</span>`	+ 				
				`<br><span style="font-size:0.7em;color:red;">`	+ "non esiste" 	+ `</span>` +				
				`<br><span style="font-size:0.7em; color:black;">nella cartella ` + path2    + `</span>` 
				
	showErrMsg2(errorMSG, msg1_Js)	
	
	return readFile, 0		
	
} // end of myOpenRead


//----------------------------------

func rowListFromFile( path1 string, fileName string, descr string, func1 string, bytesPerRow int ) []string { 
	
	file, sizeByte := myOpenRead( path1, fileName, descr, func1 )  
	if file == nil {			
		sw_stop = true 
		return nil
	} 
	numEleMax:= int( sizeByte / bytesPerRow ); 
	if numEleMax < 10 {numEleMax=10}
	
	fmt.Println("    allocate for a maximum of ", numEleMax, " rows (assumed ", bytesPerRow, " bytes per row as average)" )
	
	retRowL := make( [] string, 0, numEleMax)	
	
	r := bufio.NewReader(file)
	for {
	  line, _, err := r.ReadLine()
	  if err != nil {
		if err == io.EOF {
			break
		}
		break
	  }	 
	  retRowL = append( retRowL, string(line) ) 	
	}
	defer file.Close()
	
	fmt.Println("letto file " , fileName, "  num lines=", len(retRowL) )
	
	return retRowL 
	
} // end of rowListFromFile	

//-------------------------------------
func read_all_files() { 

	//read_lastValueSets()
	
	read_control_file() 
	
	if sw_stop { endBegin("1"); return }
	
	test_all_folder()
	if sw_stop {return}
	
	read_lemma_file( FOLDER_I_lemma, FILE_inpLemma_word_lemma, FILE_inpLemma_lemma_word) 
	if sw_stop { return }
	
	read_ParadigmaFile( FOLDER_I_paradigma, FILE_inpParadigma ) ;		
	if sw_stop { return }
		
	read_dictLang_file( FOLDER_INPUT_OUTPUT, FILE_language );	 
	if sw_stop { return }
	

		read_dictLemmaTran_file( FOLDER_IO_lastTRAN, FILE_last_updated_dict_words ) 
		if sw_stop { return }	
		
		read_dictRow_Orig_and_Tran_file( FOLDER_IO_lastTRAN,  FILE_last_updated_dict_rows)
	
	if sw_stop { return }	
		
	read_lastValueSets2()
	
	if sw_stop { return }	
	
	//read_wordsToLearn()	
	
} // end of read_all_files

//-------------------------------------------------
//-----------------------------
func read_control_file() {

	bytesPerRow:= 10
    lineD := rowListFromFile( "", file_inputControl, "input control", "read_control_file", bytesPerRow)  
	if sw_stop { return }
	
	sw_nl_only  = false	
	
	
	trim1 := string(`"`)
	trim2 := string(`'`)
	
	//--------	
	
	for _, fline00:= range( lineD ) {
	
		if (fline00 == "") {continue}    // ignore zero length line 
		fline:= strings.Split(fline00, "//")[0]           // ignore all after // 
		fline = strings.Split(fline,   "/*")[0]           // ignore all after /* 
		fli  := strings.Split(fline,   "=")               //  dictionary_folder=folder of the dictionary files,   or file = filename  
		if len(fli) < 2 { continue} 
		varia1 := strings.ToLower( strings.TrimSpace(fli[0]) ) 
		
		value1 := strings.TrimSpace(fli[1]) 
		value1 = strings.Trim(value1, trim1) 
		value1 = strings.Trim(value1, trim2)
		
		
		//fmt.Println("nread file list " , fline) 
		rowArrayCap   := 0	
		wordSliceCap  := 0
		uniqueWordsCap:= 0
		//-------------------
		switch varia1 {
		
			case "write_numbered_text" :     //       = true 
				sw_write_numbered_text = (value1 == "true") 				
			case "read_numbered_text"  :     //         = false 
				sw_read_numbered_text = (value1 == "true") 
			case "out_numbered_text_fn_prefix"   :     //          = outNumText.txt
				fPrefix_out_numbered_text = strings.ReplaceAll(value1,"\\","/") 
			case "inp_numbered_text_fn_prefix"   :     //          = inpNumText.txt
				fPrefix_inp_numbered_text = strings.ReplaceAll(value1,"\\","/") 
			case "out_number_begin" :
				outNumberBegin, err = strconv.Atoi(value1)
				if err != nil { outNumberBegin = 0;}  	
				
				
			//------------
			case "max_num_lines" :
				rowArrayCap, err = strconv.Atoi(value1)
				if err != nil { rowArrayCap = 0;}  
				if rowArrayCap > 0 {
					inputTextRowSlice    = make( []rowStruct,   0, rowArrayCap) 					
					isUsedArray          = make( []bool       , 0, rowArrayCap)  
					//dictionaryRow        = make( []rDictStruct, 0, rowArrayCap)    
					fmt.Println("max_num_lines     :", rowArrayCap, " (inputTextRowSlice capacity)")  
				}
				
			case "max_num_words" : 
				wordSliceCap, err = strconv.Atoi(value1)
				if err != nil { wordSliceCap = 0;}  
				if wordSliceCap > 0 {
					wordSliceAlpha = make([]wordStruct, 0, wordSliceCap)   
					fmt.Println("max_num_words     :", wordSliceCap, " (wordSliceFreq capacity)")  
				}
				
			case "max_num_unique":
				uniqueWordsCap, err = strconv.Atoi(value1)
				if err != nil { uniqueWordsCap = 0;}  
				if uniqueWordsCap > 0 {
					uniqueWordByFreq    = make([]wordIxStruct, 0, uniqueWordsCap)  					
					dictionaryWord      = make([]wDictStruct,  0, uniqueWordsCap)  			 
					//uniqueWordByAlpha   = make([]wordIxStruct, 0, uniqueWordsCap)  
					fmt.Println("max_num_uniques   :", uniqueWordsCap, " (uniqueWordsByFreq capacity)")  
				}	
				
			case "text_split_ignore_newline" :           // if true, newLine Character (\n) are ignored and the text is split only by full stop or any of other character as .;:!?    
				value1 = strings.ToLower(value1)					
				fmt.Println("text_split_ignore_newline :", value1)  
				if value1 == "true" {   
					sw_ignore_newLine = true 
				}
				
			case "text_split_by_newline_only" :   		
				value1 = strings.ToLower(value1)					
				fmt.Println("text_split_by_newline_only :", value1)  
				if value1 == "true" {  
					sw_nl_only = true
				}  				
							
			case "main_text_file"  :
				main_input_text_file = value1 				
			
			
			case "rewrite_word_lemma_dictionary" :
				sw_rewrite_wordLemma_dict = (value1 == "true") 				
			
		}
		
    } // end for fline00 range
	//-----------------------------------
	
	fmt.Println("sw_write_numbered_text  = ", sw_write_numbered_text, 
				", sw_read_numbered_text   = ", sw_read_numbered_text )
				
	if sw_read_numbered_text {			
	} else {	
		if main_input_text_file == "" {
			sw_stop = true 	
			msg1_Js:= `in InputControl il parametro main_input_text_file è vuoto, ma sw_read_numbered_text = false` 			
			errorMSG = `<br><br><span style="font-size:0.7em;color:black;">` + msg1_Js + `</span>` 				
			showErrMsg2(errorMSG, msg1_Js)	
			return
		}
		// when main text file is read, the numbered file must be written  in any case	
		sw_write_numbered_text = true; 
	}
	if sw_write_numbered_text {			
		if fPrefix_out_numbered_text != "" {	
			fname_out_numbered_text = FOLDER_outNumText + "/" + fPrefix_out_numbered_text	
			fmt.Println("fname_out_numbered_text  =" + fname_out_numbered_text  + " first number to use=", outNumberBegin); 
		} else {
			//fmt.Println("required out numbered text file, but \"fname_out_numbered_text\" not specified in inputControl")  
		}
	}
	
	//---------------------
   
	
	if sw_ignore_newLine && sw_nl_only {	
		fmt.Println("text_split_ignore_newline = true and text_split_by_newline_only = true,  this is incompatible, both are ignored"  )    
	}
	
	if (sw_read_numbered_text) {
		fmt.Println("since " + "sw_read_numbered_text  = ", sw_read_numbered_text, ", then main text file is ignored")			
	} 
	
    
}  // end of read_control_file()
//----------------------------------------

func checkTheWord( word0 string ) string {
	// check word 
	//  return space if not valid
	//  return the same word if OK, sometime with the first character removed 
	//---
	var wor = strings.ToLower(  strings.TrimSpace( word0 ) )
	if ((wor == "") || ( wor == "&amp" )) { 
		return ""
	}
	if wor[0:1] < " "  {
		return ""
	}
	//--------------
	var j1 = strings.IndexAny(wor, "0123456789%|\\_*•-=^&~.,;?!\"'")
	if j1 >=0 {return ""}
	//------------
	var toRemove = "°¿¡€$£"
	var j2 = strings.IndexAny( wor, toRemove)
	if j2 >= 0 {
		// la parola "wor" contiene un carattere da rimuovere
		// se il primo carattere è tra quelli da rimuovere elimino il primo carattere e prendo il resto
		// se non è il primo, allora elimino tutta la parola
		//
		// voglio controllare soltanto il primo carattere,  
		// 		ma un carattere potrebbe essere più lungo di un byte, non posso usare substring  
		
		var wor2 = ""
		for i, letterR := range wor {
			var letter = string(letterR)
			//fmt.Println( "loop rune ", wor,  "  i=", i, " letter=", letter) 
			if (i == 0) { // test primo carattere 
				if strings.IndexAny( string(letter), toRemove) < 0 {
					return ""   // il primo è ok, ma gli altri no 
				} 
				continue
			}
			wor2 += letter 
			//fmt.Println( "loop rune ", wor,  "  i=", i, " letter=", letter,  " wor2=", wor2) 
		}
		wor = wor2
		if strings.IndexAny( wor, toRemove) >= 0 {
			// il carattere strano continua ad esserci, quindi elimino la parola 
			//fmt.Println(" ex loop ",  wor, "    ", "il carattere strano continua ad esserci, quindi elimino la parola " )	
			return ""    
		} 
	} 
	return stdCode(wor)
}
//---------------------------------------

//----------------------
func buildWordList() {
    /*
	write a line in wordSliceFreq and wordSliceAlpha  for each word in the row 
	*/
	
	if  len(inputTextRowSlice) < 1 { return }  
	
	var wS1 wordStruct;
	//numMio:=0
	
	
	numberOfWords=0; 
	nn:=0
		
	lastPerc = 10;
	
	//----
	delta1 := (37.0- float64(lastPerc) ) / float64( len( inputTextRowSlice) )  
	percX1 := float64( lastPerc )  
	//------------------------------------------
	
	//sw_list_Word_if_in_ExtrRow = (last_mainpage_val_sel_extRow == "extrRow")   
	sw_list_Word_if_in_ExtrRow = (last_sel_extrRow == "extrRow")  
	
	//fromN := last_mainpage_val_inpBegRow
	//toN   := last_mainpage_val_inpBegRow + last_mainpage_val_maxNumRow -1 
	
	fromN := last_ixRowBeg 
	toN   := last_ixRowEnd      
	
	thisR_SelRow := SEL_NO_EXTR_ROW
	
	fmt.Println("buildWordList() fromN=", fromN,  "  toN=", toN , " len=" ,len(inputTextRowSlice), " sw_list_Word_if_in_ExtrRow =", sw_list_Word_if_in_ExtrRow )
	
	//antoCtr_rowSchrift :=0 ;
	//antoCtr_wordSchrift:=0; 
	//---------------------
	for ixR, rS2 := range inputTextRowSlice {	//  for each text row 
		
		//fmt.Println( " cerca parole loop ixR=", ixR, "  ", rS2);
		
		row2   := rS2.rRow1;	
		if row2 == LAST_WORD { continue }
		if sw_HTML_ready {
			percX1 += delta1 
			if ixR == (1000 * int(ixR/1000)) {
				//fmt.Println("ixR=", ixR, " percX1=", int( percX1 ) )
				go_exec_js_function( "showProgress", strconv.Itoa( int( percX1 ) ) ) 	
			}
		}		
		wordA  := regexp.MustCompile(separWord).Split(row2, -1);  // split row into words 
		
		tot1:= len(wordA) 
		
		rS2.rNumWords  = tot1      // number of words in the row 
		rS2.rListIxUnF = make( []int, tot1, tot1 )	
		
		rS2.rListFreq  = make( []int, tot1, tot1 )	
		
		inputTextRowSlice[ixR] = rS2
	
        z:= -1;
		thisR_SelRow = SEL_NO_EXTR_ROW
		if (sw_list_Word_if_in_ExtrRow) {
				if ((ixR >= fromN) && (ixR <= toN)) {
					thisR_SelRow = SEL_EXTR_ROW
				} 
			}	
			
		//swTEST := ((ixR == 4130) || (ixR == 4192) || (ixR == 4196))  // 
		
		//if ixR==211 { fmt.Println( " cerca parole loop ixR=", ixR, "  ", row2, "\n\t", wordA,  "\n\t", inputTextRowSlice[ixR] )}
		
		//if swTEST {  fmt.Println( " row2=", row2); antoCtr_rowSchrift++; } 
		
		for _, wor1 := range wordA {
			wS1.wWord2 = checkTheWord( wor1 ) ;
			if wS1.wWord2 == "" { continue }					
			z++;
			nn++				
			wS1.wNfile    = rS2.rNfile1 
			wS1.wSwSelRowR= thisR_SelRow			
			wS1.wIxRow    = ixR   // index of row containing the word 
			wS1.wIxPosRow = z;    // position of the word in the row 
			wordSliceAlpha = append(wordSliceAlpha, wS1);	
			/**
			if swTEST {  
				fmt.Println( "\t ANTO ", "word2=", wS1.wWord2, " wIxRow =ixR=", strconv.Itoa( ixR), " wordSliceAlpha=", wS1) 
				antoCtr_wordSchrift++
			} else {
				if wS1.wWord2 == "schrift" {
					 fmt.Println( "ANTO2 row2=", row2, "\n\t ANTO2 ", "word2=", wS1.wWord2, " wIxRow =ixR=", strconv.Itoa( ixR), " wordSliceAlpha=", wS1)
					 antoCtr_rowSchrift++
					 antoCtr_wordSchrift++
				} 
			}
			***/
		}
	} // end of for_ixR 
	
	//fmt.Println("anto3 antoCtr_rowSchrift=", antoCtr_rowSchrift, " antoCtr_wordSchrift=", antoCtr_wordSchrift)
	
	//---------------------------------------
	numberOfWords = len(wordSliceAlpha); 

	fmt.Println("numberOfWords=", numberOfWords)
			
	fmt.Println("number of words in text lines ", numberOfWords);
	
	//----	
	sort.Slice(wordSliceAlpha, func(i, j int) bool {
		return wordSliceAlpha[i].wWord2 < wordSliceAlpha[j].wWord2            // word  ascending order (eg.   a before b ) 		
	})
	//------------------------------
	addCodedWordToWordSlice()
	//---------------------------------
	
	// now wordSliceAlpha is in order by coded word and actual word ( eg. both actual word "über"   and "ueber" have "uber" as coded word) 
		
	/**
	for g:=0; g < len( wordSliceAlpha ); g++ {
		fmt.Println( "ANTONIO2 alpha ", wordSliceAlpha[g] )
	}	
	fmt.Println( "ANTONIO2 alpha \n")
	**/
	
	addTotRowToWord()
	
	
	
	
} // end of buildWordList
//-----------------
func addCodedWordToWordSlice() {
	/*
	add a sortKeyWord to each word element
	*/
	preW     := "" 	
	preCoded := ""
	//----------------
	for i, wS1 := range wordSliceAlpha {
		if (wS1.wWord2 != preW) { 
			preW = wS1.wWord2	
			preCoded = newCode(preW)
		}
		wordSliceAlpha[i].wWordCod = preCoded;
	}	
	//----	
	sort.Slice(wordSliceAlpha, func(i, j int) bool {
		if wordSliceAlpha[i].wWordCod != wordSliceAlpha[j].wWordCod {
			return wordSliceAlpha[i].wWordCod < wordSliceAlpha[j].wWordCod            // word  ascending order (eg.   a before b ) 
		} else {
			if wordSliceAlpha[i].wWord2 != wordSliceAlpha[j].wWord2 {
				return wordSliceAlpha[i].wWord2 < wordSliceAlpha[j].wWord2  
			} else {
				return wordSliceAlpha[i].wNfile < wordSliceAlpha[j].wNfile          // nFile ascending order (eg.   0 before 1 ) 
			}
		}
	})
	//------------------------------
	
	
} // end of addCodedWordToWordSlice

//------------------------------------------------

func elabWordList() {
	
		
	elabWordAlpha_buildWordFreqList() 		
		
	build_uniqueWord_byFreqAlpha(); 
	
	putWordFrequenceInRowArray()

	//addRowTranslation() 
	
	build_lemma_word_ix()
	
	antoList_wordSchrift_anto()
	
	
} // end of elabWordList()
//-----------------

func antoList_wordSchrift_anto() {
 
	for n1:=0; n1 < 20; n1++ {
		WS:= uniqueWordByFreq[n1] 
		/*
			type wordIxStruct struct {
				uWordCod    string	
				uWord2      string	
				uIxUnW      int            // index of this word in the uniqueWordByFreq	
				uIxUnW_al   int            // index of this word in the uniqueWordByAlpha 	
				uTotRow     int 
				uTotExtrRow int
				uIxWordFreq int            // index of this word in the wordSliceFreq	
				uSwSelRowG  int
				uSwSelRowR   int  
				uKnow_yes_ctr int 
				uKnow_no_ctr  int         // a value > 0  means that this is a word that I don't know, ie. it's to be learned   
				uIxLemmaL  []int  
				uLemmaL    []string       // list of lemma 
				uTranL     []string       // list of translation    
				uLevel     []string  
				uPara      []string  
				uExample   []string  
			}	
			//---
			type lemmaStruct struct {
				leLemma    string
				leNumWords int 
				leFromIxLW  int 
				leToIxLW    int  
				leTran     string  
			} 
		*/
		fmt.Println("LISTA  uniqueWordByFreq ", n1, " \t ", WS.uWord2, " ixLemma=", WS.uIxLemmaL, " lemmaL=", WS.uIxLemmaL); // , " tran=",  WS.uTranL) 
		for _, ixLe:= range WS.uIxLemmaL {
			LE:= lemmaSlice[ixLe]
			fmt.Println("\t\t ixLemma=", ixLe,  " \t ", LE.leLemma, " \t ", LE.leNumWords ,"words ", "\t tran=", LE.leTran)
		}
	}	
} // end of antoConta_wordSchrift_anto


//----------------------------------------------

func addTotRowToWord() {
	/*
	each element of wordSliceAlpha contains a word (the same word may be in several rows) 
	the number of repetition of a word (totRow) is put in its element  ( later will be put in each row that contain it) 
		eg.  one 3, one 3, one 3, two 4, two 4, two 4, two 4	
	*/
	preW  := wordSliceAlpha[0].wWordCod;	
	totR  := 0	
	ix1   := 0
	ix2   :=-1
	pre_wSwSelRow := SEL_NO_EXTR_ROW 
	//----------------
	tot_extrRow:=0
	for i, wS1 := range wordSliceAlpha {
		
		//fmt.Println("ANTO addTot ... i=" , i , " => ", wS1 )
		
		if (wS1.wWordCod != preW) {
			ix2 = i; 
			for i2 := ix1; i2 < ix2;i2++ {
				 wordSliceAlpha[i2].wSwSelRowG = pre_wSwSelRow; 	// se esiste almeno un richiamo a una riga estratta ( wSwSelRowR)allora questo segnale è ripetuto come wSwSelRowG
				 wordSliceAlpha[i2].wTotExtrRow = tot_extrRow   
				 wordSliceAlpha[i2].wTotRow    = totR;   // se una parola è ripetuta 3 volte, ad ogni parola è associato 3  
				 //if wordSliceAlpha[i2].wWord2=="von" { if  tot_extrRow>0 { fmt.Println("addTotRowToWord () 1 extr ", wordSliceAlpha[i2].wWord2, " .uTotExtrRow =", tot_extrRow) }  }
				
				//if wordSliceAlpha[i2].wWord2 == "schrift" { fmt.Println("ANTO2 addTotRowToWord  ", wordSliceAlpha[i2], 
				//		" .wSwSelRowR=", wordSliceAlpha[i2].wSwSelRowR, "  tot_extrRow=", wordSliceAlpha[i2].wTotExtrRow)  }  
	
			}			
			pre_wSwSelRow = SEL_NO_EXTR_ROW 
			totR = 0
			tot_extrRow = 0
			ix1  = i; 
			preW = wS1.wWordCod; 
		} 
		
		if wS1.wSwSelRowR == SEL_EXTR_ROW {   // se almeno uno è "estratto", tutti lo sono 
			pre_wSwSelRow = SEL_EXTR_ROW 
			tot_extrRow++	
			//if (wS1.wWord2 == "schrift") { fmt.Println("ANTO addTotRowToWord  ",  wS1 , " SEL_EXTR_ROW=", SEL_EXTR_ROW, "  tot_extrRow=", tot_extrRow) }  //
		} 
		totR++;     	
	}
	ix2++; 
	for i2 := ix1; i2 < len(wordSliceAlpha);i2++ {
		wordSliceAlpha[i2].wTotRow   = totR; 
		wordSliceAlpha[i2].wTotExtrRow = tot_extrRow   
		wordSliceAlpha[i2].wSwSelRowG = pre_wSwSelRow;
		// if wordSliceAlpha[i2].wWord2=="von" { if  tot_extrRow>0 { fmt.Println("addTotRowToWord () 2 extr ", wordSliceAlpha[i2].wWord2, " .uTotExtrRow =", tot_extrRow) }  }
	
	}
	//------		
	/**
	for g:=0; g < len( wordSliceAlpha ); g++ {
		fmt.Println( "ANTONIO3 alpha ", wordSliceAlpha[g] )
	}	
	**/
	
} // end of addTotRowToWord 

//---------------------------------

func elabWordAlpha_buildWordFreqList() {
	/*
	put in each row of the inputTextRowSlice the number of its words  
	from wordSliceAlpha list obtain a new list by sorting it by occurrence of the words (totRow) 
	*/
	//preW  := ""; // wordSliceAlpha[0].word;	
	//ix:=0;
	//removeIxWord :=-1 
		/******************** rimosso il 15/11/2023
	
	for nn, wS1 := range wordSliceAlpha {	
		if wS1.wWordCod == LAST_WORD  { wordSliceAlpha[nn].totRow = LAST_WORD_FREQ }
		//fmt.Println("ANTO elab...FreqList() ", nn, "  alpha=",  wordSliceAlpha[nn]);  	
		//fmt.Println("ANTO elab...FreqList() nn=", nn, ", preW=" + preW + ",  wS1 ", wS1)  
		
		if (wS1.wWordCod != preW) {		
			removeIxWord = -1
			preW = wS1.wWord2; 
		}
		ix =  wS1.ixRow
		inputTextRowSlice[ix].numWords ++; 		// how many words contains the row (eg. the row "the cat is on the table"  contains 6 words --> .numWords = 6 
		if removeIxWord >= 0 {
			//fmt.Println("\t ANTO 2elab" , "  removeIxWord=" , removeIxWord , " wordSliceAlpha[ removeIxWord ].word=" + wordSliceAlpha[ removeIxWord ].word  )
			
		    if wS1.wWordCod == wordSliceAlpha[ removeIxWord ].wWordCod {
				wordSliceAlpha[ nn ].sw_ignore = true 
			}	
		}	
		//fmt.Println("\t ANTO XXXelab"  , "  wS1.sw_ignore = ", wS1.sw_ignore)  
	}
	*****************************/
	
	//----------------
	for nn, wS1 := range wordSliceAlpha {	
		if wS1.wWordCod == LAST_WORD  { wordSliceAlpha[nn].wTotRow = LAST_WORD_FREQ }	
		inputTextRowSlice[ wS1.wIxRow ].rNumWords ++; 		// number of words in a row (eg. the row "the cat is on the table"  contains 6 words --> .numWords = 6 
	}
	//------------
	// build WordList by frequence in the text
	// the slice is sorted in descending order of frequency   ( ie. firstly the most used)   
	//-----------------------------------
	
	wordSliceFreq  = make([]wordStruct, len(wordSliceAlpha),  len(wordSliceAlpha) ) 	 // la slice destinazione del 'copy' deve avere la stessa lunghezza di quella input  
	
	copy(wordSliceFreq , wordSliceAlpha);
	
	
	// le parole eguali si trovano in righe contigue perchè hanno la stessa frequenza
	
	sort.Slice(wordSliceFreq, func(i, j int) bool {
			if wordSliceFreq[i].wTotRow !=  wordSliceFreq[j].wTotRow {
			   return wordSliceFreq[i].wTotRow > wordSliceFreq[j].wTotRow        // totRow    descending order (how many row contain the word) 
			} else {
				if wordSliceFreq[i].wWordCod != wordSliceFreq[j].wWordCod {
					return wordSliceFreq[i].wWordCod < wordSliceFreq[j].wWordCod            // word  ascending order (eg.   a before b ) 
				} else {
					return wordSliceFreq[i].wWord2 < wordSliceFreq[j].wWord2  			
				}
			}
		})
		
} // end of elabWordAlpha_buildWordFreqList

//--------------------------------------

func putWordFrequenceInRowArray() {

	ix:=0;

	//-------------------------	
	/***
	//  in each element of inputTextRowSlice define an empty slice to contain the frequence of each of its words
	for k, _ := range inputTextRowSlice {	
		tot1 := inputTextRowSlice[k].rNumWords;
		inputTextRowSlice[k].rListIxUnF =  make( []int, tot1, tot1 )		
		inputTextRowSlice[k].rListFreq  =  make( []int, tot1, tot1 )			
	}
	***/
	//---------------------------------
	//  fill each row with the frequence of its words
	//-------------------	
	
	for _, wS1 := range wordSliceFreq {	
		//fmt.Println("ANTO putWordFrequenceInRowArray() wS1 ", wS1)  
		
		ix = wS1.wIxRow; 
		ixPos := wS1.wIxPosRow; 
		/****
		num2 := len(inputTextRowSlice[ix].rListFreq) 
		if (num2 <= ixPos) {		
			fmt.Println("error " , wS1.wWord2, " ix=" , ix  ," ixPos=",   ixPos, " row ", " num2=", num2, " tot1=",  
				inputTextRowSlice[ix].rNumWords, " list=" , inputTextRowSlice[ix].rListFreq , " " , inputTextRowSlice[ix].rRow1); 			
		}
		if ((ixPos<1) || (ix<1)) {
			fmt.Println( "ERRORe  putWordFrequenceInRowArray() nx=", nx, " WS1=", wS1, "\n\t", "ix=", ix, ", ixPos=", ixPos)  
		}
		***/
		if ixPos < inputTextRowSlice[ix].rNumWords { 
			inputTextRowSlice[ix].rListIxUnF[ixPos] = wS1.wIxUniq // index of the word in the uniqueWordByFreq  	
			inputTextRowSlice[ix].rListFreq[ ixPos] = wS1.wTotRow // for each word in the row  set its frequence of use (how many times the word is used in the whole text)  
		} else {
			fmt.Println("errore in func ", red( "putWordFrequenceInRowArray"), " row n.", ix, " word pos=", ixPos, " num words in row=",  
				inputTextRowSlice[ix].rNumWords, " word=", wS1.wWord2, " row=", inputTextRowSlice[ix].rRow1)
		}	
	}
	
	//---------------------------
} // end of putWordFrequenceInRowArray

//------------------------------------------------

func put_a_priority_to_the_row_of_each_word() {
	
	//fmt.Println("put_a_priority_to_the_row_of_each_word()")
	
	//the row importance is assigned by the number of its unknown words	
		
	for k, wS1 := range wordSliceFreq {	
		
		ix := wS1.wIxRow; 
		numMinor :=0 	
		wordFreq := wS1.wTotRow; 		
		for _, frw:= range   inputTextRowSlice[ix].rListFreq {
			if frw < wordFreq { numMinor++; }
		} 
		wordSliceFreq[k].wTotMinRow = numMinor;	             // number of words with frequency < this word   	
		wordSliceFreq[k].wTotWrdRow = inputTextRowSlice[ix].rNumWords   // number of words in this row    
	}		
	
	sortWordListByFreq_and_row_priority() 
		
} // end of put_a_priority

//--------------------------------------
func build_uniqueWord_byFreqAlpha() {
	
		
	put_a_priority_to_the_row_of_each_word() 	
	
	
	
	preW := ""
	numWordUn := 0
	numWordRi := 0	
	num_word:=0
	
	numWordUn_0 := 0
	numWordRi_0 := 0	
	num_word_0 :=0 
	//--------------------
		
	for _, wS1 := range wordSliceFreq {	
		num_word++
		//if wS1.sw_ignore == false { 
		num_word_0++
		//}
		
		if wS1.wWordCod != preW {
			preW = wS1.wWordCod;
			numWordUn += 1 
			numWordRi += wS1.wTotRow 
			//if wS1.sw_ignore == false { 
			numWordUn_0 += 1 
			numWordRi_0 += wS1.wTotRow 
			//}
		}  
	}
		//------------
	if num_word_0 != num_word {
		fmt.Println( "PAROLE SINGOLE File0= ", numWordUn_0, ", PAROLE Totale=", numWordRi_0,  "  numberOfWords=" , num_word_0 , " "  );
	}
	fmt.Println( "PAROLE SINGOLE tutti= ", numWordUn, ", PAROLE Totale=", numWordRi,  "  numberOfWords=" , num_word , "\n");
		//--
	//numberOfUniqueWords = numWordUn;
	numberOfUniqueWords = numWordUn_0;
	preW      = ""
	numWordUn = 0
	numWordRi = 0	
	
	percIx := 0; 	

	//result_word2 ="";
	
	var xWordF wordIxStruct;   
	var sS  statStruct;
	
	//numWordUn = -1
	numWordUn = 0
	
	//-----------------------------
	/**** removed 2023_11_15
	// remove elements to ignore  ( those of the files after the first )
	wrk := wordSliceFreq[:0]
	for _, wS1 := range wordSliceFreq {				
			if wS1.sw_ignore { continue }
			wrk = append( wrk, wS1) 
	}
	wordSliceFreq = wrk		
	**/
	//--------------------
	/**
	for n0, wS1 := range wordSliceFreq {	
		if n0 > 20 {fmt.Println("ANTONIO1 build_uniqueWord_byFreqAlpha() ", "... continua");  break;  }
		fmt.Println("ANTONIO1 build_uniqueWord_byFreqAlpha() ", wS1) 
	}
	***/
	//---------
	
	//fmt.Println( "build_uniqueWord_byFreqAlpha() loop 1 ");
	
	for n1, wS1 := range wordSliceFreq {	
		if wS1.wWordCod != preW {
			preW = wS1.wWordCod;
			
			if wS1.wTotRow >= LAST_WORD_FREQ {
				wS1.wTotRow = 0
			}
			xWordF.uWordCod  = wS1.wWordCod;
			xWordF.uWord2    = wS1.wWord2;
			xWordF.uTotRow   = wS1.wTotRow
			xWordF.uTotExtrRow = wS1.wTotExtrRow  
			xWordF.uSwSelRowR = wS1.wSwSelRowR 
			xWordF.uSwSelRowG = wS1.wSwSelRowG
			xWordF.uIxWordFreq = n1
			
			//if wS1.wWord2=="von" { if  wS1.wTotExtrRow>0 { fmt.Println("build_uniqueWord_byFreqAlpha() extr ", wS1.wWord2, " xWordF.uTotExtrRow =", xWordF.uTotExtrRow) }  }
			
			
			//xWordF.wTran = "" 
			xWordF.uIxUnW     = len(uniqueWordByFreq)  
			uniqueWordByFreq = append( uniqueWordByFreq, xWordF);  
			numWordUn += 1 
						
			numWordRi += wS1.wTotRow 
			percIx = int(numWordUn * 100 / numberOfUniqueWords); 
			
			//fmt.Println( "   n1=", n1, " \t ", xWordF.uWord2, " \t xWordF.uSwSelRow =" , xWordF.uSwSelRow)
			
			sS.uniqueWords = numWordUn 
			sS.totWords    = numWordRi
			sS.uniquePerc  = percIx 
			sS.totPerc     = int(numWordRi * 100 / numberOfWords);
			
			//if sS.totPerc > 100 {  fmt.Println("AN TONIO4 n1=", n1, " len(wordSliceFreq)=",  len(wordSliceFreq) , " wS1.wWord2=" + wS1.wWord2 + " wS1.totRow=", wS1.totRow, " numWordRi=", numWordRi ) }
			
			/**
			if strconv.Itoa(1000 + percIx)[3:] == "0" {  
				wordStatistic_un[percIx] = sS; 	
			}
			***/				
			if sS.totPerc <= 200 {   // esistono perc > 100%,  probabilmente c'è un errore di logica 
				if strconv.Itoa(1000 + sS.totPerc)[3:] == "0" {				
					wordStatistic_tx[sS.totPerc] = sS; 
				}
			}
			//fmt.Println("STAT. ", n1, " ", xWordF.word, " numWordUn=", numWordUn,  " numWordRi=", numWordRi, " percIx=", percIx, " ", sS.uniquePerc,  " sS.totPerc=" ,  sS.totPerc); 
		} 			
	}
	//---------
	
	highestValueByte, err := hex.DecodeString("ffff")   
	if err != nil { panic(err) }
	var highestValue = string( highestValueByte ) + "end_of_list"	
	xWordF.uWordCod = highestValue 
	xWordF.uWord2   = highestValue 	
	
	xWordF.uTotRow = 1 ; // the lowest frequency
	xWordF.uTotExtrRow = 0               
	xWordF.uIxWordFreq = len(uniqueWordByFreq)   
	xWordF.uIxUnW      = len(uniqueWordByFreq)  	
	xWordF.uKnow_yes_ctr = 0 
	xWordF.uKnow_no_ctr  = 0 	
	//xWordF.uTranL      = []string{ xWordF.uWord2 }        // ??anto8 .uTranL
	uniqueWordByFreq   = append( uniqueWordByFreq, xWordF);  
	
	//--------------------------
	
	addWordLemmaTranLevelParadigma()   
	
	add_ixWord_to_WordSliceFreq()
	
	
	//addWordTranslation()		
	
	//---------------------
	uniqueWordByAlpha = make([]wordIxStruct, len(uniqueWordByFreq),  len(uniqueWordByFreq))	 // la slice destinazione del 'copy' deve avere la stessa lunghezza di quella input  
	
	//stat_useWord();
	
	copy( uniqueWordByAlpha, uniqueWordByFreq); 
	
	
	//fmt.Println("build_uniqueWord_byFreqAlpha() PROVA len unique Freq() = ", len(uniqueWordByFreq)   )
	
	//------------
	sort.Slice(uniqueWordByAlpha, func(i, j int) bool {
		if uniqueWordByAlpha[i].uWordCod != uniqueWordByAlpha[j].uWordCod {
			return uniqueWordByAlpha[i].uWordCod < uniqueWordByAlpha[j].uWordCod            // word  ascending order (eg.   a before b ) 
		} else {
			return uniqueWordByAlpha[i].uWord2 < uniqueWordByAlpha[j].uWord2  			
		}
	})
	//---------

	//console( "\nlista uniqueWordByAlpha")	
	// update alpha index  // ixUnW,  ixUnW_al	
	
	//fmt.Println("\n build_uniqueWord_byFreqAlpha() PROVA len unique alpha() = ", len(uniqueWordByAlpha)   )
	
	for u:=0; u < len(uniqueWordByAlpha); u++ {
		f:= uniqueWordByAlpha[u].uIxUnW
		uniqueWordByFreq[f].uIxUnW_al  = u; 		
		uniqueWordByAlpha[u].uIxUnW_al = u
		
		//fmt.Println( "ANTONIO6 uniqueWordByAlpha ANTONIO prova u=", u , " \t unique = ", uniqueWordByAlpha[u].uWordCod , " \t ",   uniqueWordByAlpha[u].uWord2) 
				
	}
 	//console( "------------------\n")
	
	//fmt.Println("\n---------------- build_uniqueWord_byFreqAlpha() end  \n")
	

} // end of build_uniqueWord_byFreqAlpha

//-----------------------

func build_lemma_word_ix() {
	
	//	build a slice with all lemma with all words 
	//  lemma_word_ix  loaded in addWordLemmaTranLevelParadigma

	var LW lemmaWordStruct
	/***						
			type lemmaStruct struct {
				leLemma    string
				leNumWords int 
				leFromIxLW  int 
				leToIxLW    int  
				leTran     string  
			}
			//---------------
			type lemmaWordStruct struct {
				lw_lemmaCod string 	
				lw_lemma2   string 	
				lw_word     string 
				lw_ixLemma    int
				lw_ixWordUnFr int
			}
			//------------
	***/	
	//--------------------
	
	// order by lemma and word 	
	sort.Slice(lemma_word_ix, func(i, j int) bool {
			if  lemma_word_ix[i].lw_lemmaCod != lemma_word_ix[j].lw_lemmaCod {
				return lemma_word_ix[i].lw_lemmaCod < lemma_word_ix[j].lw_lemmaCod 	
			} else {
				if  lemma_word_ix[i].lw_lemma2 != lemma_word_ix[j].lw_lemma2 {
					return lemma_word_ix[i].lw_lemma2 < lemma_word_ix[j].lw_lemma2	
				} else {
					return lemma_word_ix[i].lw_word < lemma_word_ix[j].lw_word 
				}				
			}
		} )		
	fromIxWL:=-1
	toIxWL  :=-1	
	preIxLem:=-1 
	
	
	for z2:=0; z2 < len( lemma_word_ix); z2++ { 
		LW = lemma_word_ix[z2];
		if LW.lw_ixLemma != preIxLem  {	
			if preIxLem >=0 {
				lemmaSlice[preIxLem].leFromIxLW = fromIxWL
				lemmaSlice[preIxLem].leToIxLW   = toIxWL
			}
			preIxLem = LW.lw_ixLemma
			fromIxWL = z2
		} 
		toIxWL = z2	
	}
	if preIxLem >=0 {
				lemmaSlice[preIxLem].leFromIxLW = fromIxWL
				lemmaSlice[preIxLem].leToIxLW   = toIxWL
			}
	//---------------
	
	for z2:=0; z2 < len( lemma_word_ix); z2++ { 
		if z2 < 20 { fmt.Println(   "lemma_word_ix[", z2, "] = ", lemma_word_ix[z2]  ) } else { break }
		
	}
	//---------------
	
	for z2:=0; z2 < len( lemmaSlice); z2++ { 
		LE := lemmaSlice[z2];
		if LE.leLemma == "werkzeug" {
			fmt.Println("lemma = ", LE )
			for z3:= LE.leFromIxLW; z3 <= LE.leToIxLW; z3++ {
				LW = lemma_word_ix[z3]
				fmt.Println(" \t word = ", LW) 
			} 
			break;
		}
	}
	
} // end of build_lemma_word_ix  
 
//------------------------------------------
func OLDbuild_lemma_word_ix() {
	
	//	build a slice with all lemma with all words 
	/*
	lemmaWordStruct struct {lw_lemmaCod string, lw_lemma2 string, lw_word  string, lw_ixWordUnFr int
	*/
	lemma_word_ix = make([]lemmaWordStruct, 0, 0)  
	
	var LW lemmaWordStruct
	
	for z:=0; z < len(uniqueWordByFreq); z++ {	
		
		LW.lw_lemmaCod = ""
		LW.lw_lemma2 = ""
		LW.lw_word = ""
		LW.lw_ixWordUnFr = -1		
		LW.lw_word = uniqueWordByFreq[z].uWord2 
		LW.lw_ixWordUnFr   = uniqueWordByFreq[z].uIxUnW
		
		lemmaLis :=  uniqueWordByFreq[z].uLemmaL
		
		/**
		if ((LW.lw_word == "tun") || (LW.lw_word == "umwelt") ) {	
		
				fmt.Println("build_lemma_word_ix() z=", 
						"\t uniqueWordByFreq[", z, "] = ",  uniqueWordByFreq[z],  
							"\n\t(struct: uWordCod, uWord2, uIxUnW, uIxUnW_al, uTotRow,  uIxWordFreq,	uSwSelRowG,uSwSelRowR,uLemmaL,uTranL, uLevel, uPara,uExample ",  
							"\n\t wordSliceFreq[ uIxWordFreq ]  = " , wordSliceFreq[ uniqueWordByFreq[z].uIxWordFreq ]  ,
						"\n\t z=",z, "lemmaWordStruct="," LW.lw_ixWordUnFr=",LW.lw_ixWordUnFr, "  LW.lw_word=", LW.lw_word ) 						
		}
		**/
		
		for m:=0; m <  len(lemmaLis); m++ {
			LW.lw_lemma2   = lemmaLis[m]   
			LW.lw_lemmaCod = newCode( LW.lw_lemma2 )    
			lemma_word_ix = append( lemma_word_ix, LW )
		}	
	} 
	//--------------
	// order by lemma and word 	
	sort.Slice(lemma_word_ix, func(i, j int) bool {
			if  lemma_word_ix[i].lw_lemmaCod != lemma_word_ix[j].lw_lemmaCod {
				return lemma_word_ix[i].lw_lemmaCod < lemma_word_ix[j].lw_lemmaCod 	
			} else {
				if  lemma_word_ix[i].lw_lemma2 != lemma_word_ix[j].lw_lemma2 {
					return lemma_word_ix[i].lw_lemma2 < lemma_word_ix[j].lw_lemma2	
				} else {
					return lemma_word_ix[i].lw_word < lemma_word_ix[j].lw_word 
				}				
			}
		} )		
		
	/**
	for z2:=0; z2 < len( lemma_word_ix); z2++ { 
		LW = lemma_word_ix[z2];
		if ((LW.lw_word == "tun") || (LW.lw_word == "umwelt") ) {
				fmt.Println("build_lemma_word_ix() sortato ", LW); 
		}
	}	
	**/
	
} // end of OLDbuild_lemma_word_ix  
 

//--------------------------

func sortWordListByFreq_and_row_priority() {
	
	//fmt.Println(" sortWordListByFreq_and_row_priority()")
	
	sort.Slice(wordSliceFreq, func(i, j int) bool {
	
		if wordSliceFreq[i].wTotExtrRow !=  wordSliceFreq[j].wTotExtrRow {
		   return wordSliceFreq[i].wTotExtrRow > wordSliceFreq[j].wTotExtrRow  // wTotExtrRow  descending order ( number of extracted rows  if id_sel_2_extrRow option = id="extrRow", else =0) 	
		}
		
		if wordSliceFreq[i].wTotRow !=  wordSliceFreq[j].wTotRow {
		   return wordSliceFreq[i].wTotRow > wordSliceFreq[j].wTotRow         // totRow    descending order (how many rows contain the word)  	  		   
		}	
		
		if wordSliceFreq[i].wWordCod !=  wordSliceFreq[j].wWordCod {
		   return wordSliceFreq[i].wWordCod < wordSliceFreq[j].wWordCod            // word      ascending order	  		   
		}	
		if wordSliceFreq[i].wWord2 !=  wordSliceFreq[j].wWord2 {
		   return wordSliceFreq[i].wWord2 < wordSliceFreq[j].wWord2                // word      ascending order	  		   
		}				
		if wordSliceFreq[i].wTotMinRow !=  wordSliceFreq[j].wTotMinRow {
		   return wordSliceFreq[i].wTotMinRow < wordSliceFreq[j].wTotMinRow  // totMinRow ascending order	(how many words in the row are not yet learned) 
		}
		if wordSliceFreq[i].wSwSelRowR !=  wordSliceFreq[j].wSwSelRowR {
		   return wordSliceFreq[i].wSwSelRowR < wordSliceFreq[j].wSwSelRowR    // first the extracted row  
		}
		return wordSliceFreq[i].wIxRow < wordSliceFreq[j].wIxRow             // ixRow     ascending order	( first the rows which were first in the text)  		   
			
	})	
	
	/**
		fmt.Println( "\nANTONIO5 wordSliceFreq -------------------------")
	for g:=0; g < len( wordSliceFreq ); g++ {
		fmt.Println( "ANTONIO5 alpha ", wordSliceFreq[g] )
	}	
	**/
	
} // end of sortWordListByFreq_and_row_priority

//-----------------------------
/**
func print_rowArray( where string) {
		
	//result_row1 = ""; 
	
	for i, wR1 := range inputTextRowSlice {	
		if i > 10 {break;} 
		**
		strFreqList := arrayToString(wR1.listFreq, ",")
		
		strrow:= "ix="  + strconv.Itoa(i) + " w="   + strconv.Itoa(wR1.numWords) + 						
						" lf=" + strFreqList +
						" " + wR1.row1;
		**				
		fmt.Println( where , "  ",  wR1.row1);
		//result_row1 = result_row1 + "<br>" + strrow; 
	}

} // end of print_rowArray
**/

//----------------------------------------------------------------
/***
func prova_js_function_treValori() {
	js_function1 := "js_go_TreValori"
	jsInpFunction:= "funzioneJavaScript che ha chiamato il go" 
	//goFunc       := "build_func_go" 
	inpstr       := "stringa di prova"
	
	go_exec_js_function(js_function1 + "," + jsInpFunction, inpstr) 
	
}
***/
//----------------------------------------------------------------
func go_exec_js_function(js_function0 string, inpstr string) {
	var goFunc string 
 	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		goFunc = strings.ReplaceAll(details.Name(), "main.","")
	} else {
		goFunc=""
	}
	js_fun        := strings.Split( (js_function0 + ",,,,") ,",") 	
	js_function   := strings.TrimSpace( js_fun[0] )
	if js_function == ""  { return }
	jsInpFunction := strings.TrimSpace( js_fun[1] )
	
	js_parm:=""
	k1:= strings.Index(js_function, "(") 
	if k1 > 0 {
		js_parm     = strings.ReplaceAll(  js_function[k1+1:], ")","")			
		js_function = strings.TrimSpace(js_function[0:k1] )
	} 
	/**
	fmt.Println("  js_function=" + js_function)
	fmt.Println("      js_parm=" + js_parm)
	fmt.Println("jsInpFunction=" + jsInpFunction)
	fmt.Println("       goFunc=" + goFunc)
	**/
	
	/*
	This function executes a javascript eval command 
	which must execute a function by passing string constant to it. 
	Should this string contain some new line, e syntax error would occur in eval the statement.
	
	To avoid this kind of error, the string argument (inpstr) of the javascript function (js_function) 
	is forced to be always enclosed in back ticks trasforming it in "template literal".  
	Just in case back ticks and dollars are in the string, they are replaced by " "   	
	*/
	inpstr = strings.ReplaceAll( inpstr, "`", " "   ); 	   	 
	inpstr = strings.ReplaceAll( inpstr, "$", "&dollar;"); 
	
	evalStr := fmt.Sprintf( "%s(`%s`,`%s`,`%s`,`%s`);",  js_function, inpstr, js_parm, "js=" + jsInpFunction, "go=" + goFunc ) ; 
	
	//fmt.Println("evalStr=" + evalStr); 
	
	ui.Eval(evalStr)
	
} // end of go_exec_js_function
//----------------------------------------------------------------

func go_exec_js_functionX(js_function string, inpstr string ) {
	/*
	This function executes a javascript eval command 
	which must execute a function by passing string constant to it. 
	Should this string contain some new line, e syntax error would occur in eval the statement.
	
	To avoid this kind of error, the string argument (inpstr) of the javascript function (js_function) 
	is forced to be always enclosed in back ticks trasforming it in "template literal".  
	Just in case back ticks and dollars are in the string, they are replaced by their html symbols.   	
	*/
	inpstr = strings.ReplaceAll( inpstr, "`", " "   ); 	  
	//inpstr = strings.ReplaceAll( inpstr, "`", "&#96;"   ); 	 
	inpstr = strings.ReplaceAll( inpstr, "$", "&dollar;"); 
	
	evalStr := fmt.Sprintf( "%s(`%s`);",  js_function,  inpstr ) ; 
	
	ui.Eval(evalStr)
	
} // end of go_exec_js_function

//--------------------------------

func arrayToString(a []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
//----------------------------

func isThereNumber(s string) bool {
    for _, c := range s {
        if c >= '0' && c <= '9' {
            return true
        }
    }
    return false
}

//------------------------------
func buildStatistics() {		
		//var rows []string
		var result string = ""
		
		if len( only_level_numWords ) < 1 { return }
		
		/***	
		msgLevelStat = "" 		
		if percA0 > 0 {msgLevelStat += ", A0: " + strconv.Itoa(percA0) + "%" }
		if percA1 > 0 {msgLevelStat += ", A1: " + strconv.Itoa(percA1) + "%" }
		if percA2 > 0 {msgLevelStat += ", A2: " + strconv.Itoa(percA2) + "%" }
		if percB1 > 0 {msgLevelStat += ", B1: " + strconv.Itoa(percB1) + "%" }
		if percOth > 0 {msgLevelStat += ", Oth: " + strconv.Itoa(percOth) + "%" }
		if len(msgLevelStat) > 1 {msgLevelStat = msgLevelStat[2:] } 
		**/
		
		msgLevelStat = "" 
		for f:=1; f < len( only_level_numWords ) ; f++ {
			//if only_level_numWords[f] == 0 { continue }
			if perc_level[f] == 0 { continue }
			msgLevelStat += ", " + list_level[f] + ": " + strconv.Itoa( perc_level[f] ) + "%"  
		}	
		if only_level_numWords[0] > 0 {  
			msgLevelStat += ", " + list_level[0] + ": " + strconv.Itoa( perc_level[0] ) + "%"  
		}
		if len(msgLevelStat) > 1 {msgLevelStat = msgLevelStat[2:] } 

		result += "livello " + msgLevelStat //  + "..endLevel ";  
		
		for _, sS:= range wordStatistic_tx {	
			if sS.totWords == 0 { continue; }
			if sS.uniqueWords < 100 { continue}
			//fmt.Println( sS.uniqueWords , " words (",  sS.uniquePerc, "%), found ", 
			//	sS.totWords,  " times in the text(", sS.totPerc,"%)" ) 
			
			//result += "<br>" + fmt.Sprintln( sS.uniqueWords , " words (",  
			//	sS.uniquePerc, "%), make up ", sS.totPerc,"% of the text (", sS.totWords, " words)") 
			result += "<br>" + fmt.Sprintln( sS.uniqueWords, ",", sS.uniquePerc, ",", sS.totPerc, ",", sS.totWords) 	
		}  		
		result += "<br>" 
		go_exec_js_function("js_go_updateStatistics", result )		
	
}	
//-----------------------------------
func stat_useWord() {
	len1:=  len(uniqueWordByFreq)
	len2:= float64(len1)/100
	
	fmt.Println("len1=", len1, " ", len2) 	

	lisPerc := [29]float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9,1,2,3,4,5,6,7,8,9, 10,20,30,40,50,60,70,80,90,100}
	listIxPerc:= make([]int,0,40)
	for z:=0; z < len(lisPerc); z++ {
		per1 := lisPerc[z]
		per2 := int( float64(per1) * len2)
		listIxPerc = append( listIxPerc, per2 )
		//fmt.Println("stats ", per1, "% = num.Elem.",  per2 )  	
	}   
	
	lastTot:=0;
	ixNow:=0	
	for z:=0; z < len(listIxPerc); z++ {
		//from1 = ixNow
		ixNow =  listIxPerc[z]-1
		if ixNow< 0 { ixNow=0;}
		
		if uniqueWordByFreq[ixNow].uTotRow == lastTot { continue }
		
		//fmt.Println("stats ", lisPerc[z], "% = num.Elem.",  listIxPerc[z], " toIx=", ixNow,   
		//					" num.Rows per word=",uniqueWordByFreq[ixNow].totRow )
		if listIxPerc[z] >= 1 {	
			fmt.Println( "stats ",  prtFloat( lisPerc[z] , 5,1 ) ,"% = num.Elem.",prtInt( listIxPerc[z] , 5 )," sono usate ", prtInt( uniqueWordByFreq[ixNow].uTotRow, 5 ), 
					" o più volte  (",  prtFloat( lisPerc[z] , 5,1 ),"% delle parole non sono usate più di ", prtInt(  uniqueWordByFreq[ixNow].uTotRow, 5 ), " volte)")					
		}
		lastTot = uniqueWordByFreq[ixNow].uTotRow 	
	} 
	
} // end of stat_useWord
//------------------------------------

//--------------------------
func prtFloat( input float64, maxL int, dec1 int ) string {
	//  space character after % is the padding character which will be repeated by the value replacing the first *   
    //  first  * character is replaced by the difference between the maximum length and the actual length of the number converted to string 
    //  second * character is replaced by dec1 value ( how many decimal)    	
	return fmt.Sprintf("% *s%.*f", maxL-len( strconv.FormatFloat(input, 'f', 2, 64)), "", dec1,  input )
} 
//--------------------------
func prtFloat1( input float64, maxL int ) string {
  return fmt.Sprintf("% *s%.1f", maxL-len( strconv.FormatFloat(input, 'f', 2, 64)), "",  input )
} 
//---
func prtInt( input int, maxL int ) string {
  return fmt.Sprintf("% *s%d", maxL-len( strconv.Itoa(input)), "",  input )
}  
//-----------

//------------------------------------

func cleanRow(row0 string) string{

		row1 := strings.ReplaceAll( row0, "<br>", " "   ); 	// remove <br> because it's needed to split the lines to transmit  

		// if the row begins with a number remove this number
		row1 = strings.Trim(row1," \t")    // trim space and tab code
		k:=    strings.IndexAny(row1, " \t");	
		if (k <1) { return row1;}
		numS := row1[:k]	
		_, err := strconv.Atoi(numS)
		if err != nil { return row1;}  
		return strings.Trim(row1[k+1:]," \t"); 
}


//==============================================

func searchAllWordWithPrefixInAlphaList(  wordPref string) (int, int) {
	
	// get the indicies of the first and the last word beginning with the required prefix (-1,-1 if not found)  
	
	wordPref = strings.ToLower(strings.TrimSpace( wordPref));  
	wordCodPref:= newCode(wordPref)
	
	lenPref:= len(wordPref); 
	ixTo := -1; ixFrom:= -1;	
	
	if lenPref == 0 { return ixFrom, ixTo }
	
	ix1, ix2:= lookForWordInUniqueAlpha(wordCodPref)	
	
	/***
	fmt.Println("ANTONIO SEARCH ALPHA wordPref=" + wordPref + " wordCodPref=" +  wordCodPref + " ix1=", ix1, " ix2=", ix2) 
	for k:= ix1; k <= ix2; k++ {
		if k < 0 { continue}
		fmt.Println("ANTONIO SEARCH ALPHA k=", k , " ==>" , uniqueWordByAlpha[k])
	}	
	***/
	
	wA :=""
	spaceFill := "                                                          ";  
	//-----------
	for k:= ix1; k >= 0; k-- {
		wA =  uniqueWordByAlpha[k].uWordCod + spaceFill
		if wA[0:lenPref] < wordCodPref { break; }
		ixFrom = k; 
	}  
	
	if (ixFrom >=0) { ixTo = ixFrom; }  //  se ixFrom è valido, deve essere valido anche ixTo   
	
	for k:= ix2; k < numberOfUniqueWords; k++ {
		wA =  uniqueWordByAlpha[k].uWordCod + spaceFill  
		if wA[0:lenPref] > wordCodPref { break; }
		ixTo = k; 
		if (ixFrom < 0) {ixFrom = ixTo;}  //  se ixTo è valido, deve essere valido anche ixFrom   
	}  
	return ixFrom, ixTo 
	
} // end of searchAllWordWithPrefixInAlphaList

//--------------
func lookForWordInUniqueAlpha(wordCoded string) (int, int) {
	
	// find 2 indices of the 2 words nearest to the word to find 	
	
	low   := 0
	high  := numberOfUniqueWords - 1	
	maxIx := high; 
	//----
	for low <= high{
		median := (low + high) / 2
		if median >= len(uniqueWordByAlpha) {
			fmt.Println("errore in lookForWordInUniqueAlpha: median=", median , "     len(uniqueWordByAlpha)=" ,  len(uniqueWordByAlpha) )
		}
		if uniqueWordByAlpha[median].uWordCod < wordCoded {
			low = median + 1
		}else{
			high = median - 1
		}
	} 
	//---
	fromIx:= low; toIx := high; 
	if fromIx > toIx { fromIx = high; toIx = low;}
	if fromIx < 0 { fromIx=0} 
	if toIx  > maxIx { toIx = maxIx}
	return fromIx, toIx	

} // end of lookForWordInUniqueAlpha



//-----------------------------

func lookForLemmaWord(lemmaCode string) (int, int) {
	
	// find 2 indices of the 2 words nearest to the word to find 
	
	low   := 0
	high  := len(lemma_word_ix) - 1	
	maxIx := high; 
	
	//----
	for low <= high{
		median := (low + high) / 2
		if lemma_word_ix[median].lw_lemmaCod < lemmaCode {  
			low = median + 1
		}else{
			high = median - 1
		}
	} 
	//---
	fromIx:= low; toIx := high; 
	if fromIx > toIx { fromIx = high; toIx = low;}
	if fromIx < 0 { fromIx=0} 
	if toIx  > maxIx { toIx = maxIx}
	return fromIx, toIx	

} // end of lookForLemmaWord

//-----------------------------

func lookForLemma(lemmaTarg string) (int, int) {
	
	// find 2 indices of the 2 words nearest to the word to find 
	/**
				
			//---
			type lemmaStruct struct {
				leLemma     string
				leNumWords  int 
				leFromIxLW  int 
				leToIxLW    int  
				leTran      string 
				leLevel     string  
				lePara      string  
				leExample   string  	
			} 	
			var lemmaSlice       [] lemmaStruct         // lemma , translation 
	***/
	
	low   := 0
	high  := len(lemmaSlice) - 1	
	maxIx := high; 
	
	//----
	for low <= high{
		median := (low + high) / 2
		if lemmaSlice[median].leLemma < lemmaTarg {  
			low = median + 1
		}else{
			high = median - 1
		}
	} 
	//---
	fromIx:= low; toIx := high; 
	if fromIx > toIx { fromIx = high; toIx = low;}
	if fromIx < 0 { fromIx=0} 
	if toIx  > maxIx { toIx = maxIx}
	return fromIx, toIx	

} // end of lookForLemmaWord


//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

/**
//------------------------------------------------------
func testPreciseWord(target string ) {
	fmt.Println( "cerca ", target);
	var xWordF wordIxStruct;  
	ix1, ix2 := searchOneWordInAlphaList( target ) 
	
	if (ix2 < 0) {
		fmt.Println( target, "\t NOT FOUND "); 
	} else {
		for ix:= ix1; ix <= ix2; ix++ {
			xWordF =  uniqueWordByAlpha[ix] 
			fmt.Println( target, "\t found  ",  xWordF.uWord2, " ix=", ix,   " totRow=", xWordF.uTotRow,
				" ixwordFre=", xWordF.uIxWordFreq); 
		}
	} 
	fmt.Println( "-------------------------------------------------\n" ); 
} 
***/
//------------------------------------------------------
func testGenericWord(pref string ) {
	//fmt.Println( "cerca tutte le parole che iniziano con " + pref);	
	var xWordF wordIxStruct;  
	
	from1, to1 := searchAllWordWithPrefixInAlphaList( pref )
	if (to1 < 0) {
		fmt.Println( "nessuna parola che inizia con " , pref); 
	} else {	
		for i:=from1; i <=to1; i++ {	
			xWordF =  uniqueWordByAlpha[i] 
			fmt.Println( "trovato ", xWordF.uWord2, " ix=", i,   " totRow=", xWordF.uTotRow, " ixwordFre=", xWordF.uIxWordFreq); 
		}
	}
	fmt.Println( "-------------------------------------------------\n" ); 
	return 
}
//---------------------
func getScreenXY() (int, int) {
	
	// use ==>  var x, y int = getScreenXY();
	
	var width  int = int(win.GetSystemMetrics(win.SM_CXSCREEN));
	var height int = int(win.GetSystemMetrics(win.SM_CYSCREEN));
	if width == 0 || height == 0 {
		//fmt.Println( "errore" )
		return 2000, 2000; 
	}	
	width  = width  - 20;  // subtraction to make room for any decorations 
	height = height - 40;  // subtraction to make room for any decorations 
	
	return width, height
}
	
//=========================
func test_folder_exist( myDir string) {
    _, err := os.Open( myDir )
    if err != nil {
		msg0:= `la cartella <span style="color:blue;">` + myDir + "</span> non esiste"			
		msg1:= `la cartella ` + myDir + "</span> non esiste"				
			
		msg2:= "read_dictionary_folder"
		errorMSG = `<br><br> <span style="color:red;">` + msg0 + `</span>` +  
			`<br><span style="font-size:0.7em;">(func ` + msg2 	 + ")" + `</span>` 		
		showErrMsg(errorMSG, msg1, msg2 )	
		
		sw_stop = true 
		return		
    }
} // end of test_folder_exist	

//--------------------------------

func read_dictLang_file( path1 string, inpFile string) {
	bytesPerRow:= 10
    lineD := rowListFromFile( path1, inpFile, "scelta lingua e voce", "read_dictLang_file", bytesPerRow)  
	if sw_stop { return }
	
	lineZ := ""
	prevRunLanguage = ""	
	for z:=0; z< len(lineD); z++ { 
		lineZ = strings.TrimSpace(lineD[z]) 
		if lineZ == "" { continue }
		if lineZ[0:9] == "language=" {			
			prevRunLanguage = lineZ[9:] 			
		}
	}	
}// end of read_dictLang_file		

//---------------------
func OLDget_rowid(id string) (string, int) {
	// es.  10_5
	var id_pref string
	var id_num  int	
	k1:= strings.Index(id,"_")
	if k1 >= 0 { 
		id_pref   =  strings.TrimSpace(  id[:k1] ) 
		id_num, _ = strconv.Atoi( strings.TrimSpace( id[k1+1:] ) )
	} else {
		id_pref   ="00"  
		id_num, _ = strconv.Atoi( strings.TrimSpace( id ) )
	} 
	
	return id_pref, id_num 
} // end of OLDget_rowid	
//--------------------------------

func get_rowid2(id string) string {
	// es.  10_5
	var id_pref string
	k1:= strings.Index(id,"_")
	if k1 >= 0 { 
		id_pref   =  strings.TrimSpace(  id[:k1] ) 
	} else {
		id_pref   = id  
	} 	
	return id_pref
	
} // end of get_rowid2	
//--------------------------------

func read_dictRow_Orig_and_Tran_file( path1 string, inpRowFile string) {
	bytesPerRow:= 10
    lineD := rowListFromFile( path1, inpRowFile, "righe orig/tran", "read_dictRow_Orig_and_Tran_file", bytesPerRow)  
	if sw_stop { return }
	
	lineZ := ""
	/*		
	3|O|Ein Lebewesen tauscht aber auch Energie und Dinge, Stoffe mit seiner Umwelt aus.
	3|T|Ma un essere vivente scambia anche energia, cose e sostanze con il suo ambiente.
	*/
	
	
	prevRunLanguage = ""
	//prevRunListFile = ""
	//var rowDict rDictStruct
	var rS1 rowStruct
	var pre_id_key=""
	
	showReadFile = ""
	numLines:=0
	
	inputTextRowSlice = nil 
	isUsedArray       = nil 
	
	//rowLineIxList    = nil
	
	//var gRix rowIxStruct
	var indice = 0
	numAll :=0; numO :=0; numT:=0 ; numT_err:=0; num_oth :=0
	//--
	pre_id_group     := ""
	pre_id_keyFirst  := ""
	pre_id_keyLast   := ""
	pre_id_row       := ""
	id_group         := ""
	gruppi_option    := ""	
	
	num_O_ix   := 0 
	firstIxRow :=0
	lastIxRow  :=0
	//----------

	group_zero := "."; 
	
	lineZero := []string{ group_zero + "_0|O|"}  // empty row
	
	lineD = append( lineZero, lineD...)  // insertion of an empty row to avoid using index 0  
	ngr:=0
	
	var rG rowGroupStruct ; 
	
	//----
	for z:=0; z< len(lineD); z++ { 
		lineZ = strings.TrimSpace(lineD[z]) + "||||"	
		field  := strings.Split( lineZ, "|" )
		id_key := strings.TrimSpace( field[0] )
		ty     := strings.TrimSpace( field[1] ) 		
		row    := strings.TrimSpace( field[2] )  	 
		
		
		if ((row == "") || (id_key == "") || (ty=="")) { continue }
		
		//fmt.Println("carica riga z=", z, " lineZ=", lineZ, "\n\t fields: 0=", field[0], "  1=", field[1], " 2=", field[2], "  3=", field[3]   )
		
		id_group = get_rowid2(id_key)
		
		//fmt.Println("\t id_key=", id_key, " id_group=", id_group)
		
		if id_group != pre_id_group {
			if pre_id_group != "" {
				// elabora fine gruppo precedente {
				
				//fmt.Println("id_key=", pre_id_group, " from=", pre_id_keyFirst, " to=", pre_id_keyLast, " first_ixRow=", firstIxRow, " lastIxRow=", lastIxRow)  
				
				if pre_id_group != group_zero { 
					gruppi_option += "<option>" +  fmt.Sprintf("%s %s", pre_id_group,  pre_id_row)  + "</option>\n"
					rG.rG_ixSelGrOption   = ngr    // group number = index of group selection 
					rG.rG_group           = pre_id_group
					rG.rG_firstIxRowOfGr  = firstIxRow 
					rG.rG_lastIxRowOfGr   = lastIxRow						
					lista_gruppiSelectRow = append( lista_gruppiSelectRow, rG )    
					ngr++
					
					fmt.Println("Group=", pre_id_group, " key: from=", pre_id_keyFirst, " to=", pre_id_keyLast, " ixRow: first_ixRow=", firstIxRow, " lastIxRow=", lastIxRow)  
					
				}			
			}
			// elabora inizio nuovo gruppo 
			pre_id_group     = id_group 
			pre_id_keyFirst  = id_key
			pre_id_row       = row
			num_O_ix     = 0  
			if ty == "O" {
				num_O_ix = 0  
			} else {
				fmt.Println( red("errore "), " in read_dictRow_Orig_and_Tran_file" , " id_key=", pre_id_group, " from=", pre_id_keyFirst, " to=", pre_id_keyLast , 
					red("row type " + ty + " not equal O"), "\n\t", lineZ) 
				continue	
			}
			//newNum=0  ; // ?anto rinumera  
			pre_id_key  = id_key
			
			//fmt.Println("\nNUOVO Group=", id_group)
			
		}
		//--
		
		pre_id_keyLast   = id_key
		
		switch ty {
			case "O" :				
				pre_id_key  = id_key			
				numO++				
				num_O_ix++  
				rS1.rIdRow 		 = id_key
				rS1.rRow1  		 = row
				rS1.rTran1		 = ""
				rS1.rixGroup     = ngr      // indice del gruppo 
				rS1.rixBaseGroup = num_O_ix // posizione del row nel gruppo ( si inzia dal num.1 )  
		    case "T" : 
				if id_key != pre_id_key {  // error 
					numT_err++
					continue
				}		
				inputTextRowSlice[ numLines-1].rTran1  = row
				numT++
				continue 
			default  :  
				num_oth++
			    continue
		}
	
		numAll++
		numLines++
		inputTextRowSlice = append(inputTextRowSlice, rS1);	
		isUsedArray       = append(isUsedArray, false)  
		
		indice   = len(inputTextRowSlice)-1
		
		if num_O_ix==1 {
			firstIxRow = indice
		}  
		lastIxRow = indice
		
	} // end of for z 
	//-------------

	if pre_id_group != group_zero { 
		gruppi_option += "<option>" +  fmt.Sprintf("%s %s", pre_id_group,  pre_id_row)  + "</option>\n"
		rG.rG_ixSelGrOption   = ngr    // group number = index of group selection 
		rG.rG_group           = pre_id_group
		rG.rG_firstIxRowOfGr  = firstIxRow 
		rG.rG_lastIxRowOfGr   = lastIxRow						
		lista_gruppiSelectRow = append( lista_gruppiSelectRow, rG )    
		ngr++
		fmt.Println("Group=", pre_id_group, " key: from=", pre_id_keyFirst, " to=", pre_id_keyLast, " ixRow: first_ixRow=", firstIxRow, " lastIxRow=", lastIxRow)  		
	}
	//-------------
	
	go_exec_js_function( "js_go_build_rowGruppi", gruppi_option); 	
	
	//-------
	//fmt.Println("read_dictRow_Orig_and_Tran_file "  , " numAll=", numAll, " numO=", numO, " numT=", numT, " numT_err=", numT_err, " num_oth=", num_oth)
	/**
	for z:=0; z< len(inputTextRowSlice); z++ {
		fmt.Println("caricata inputTextRowSlice[",z,"]=", inputTextRowSlice[z] )
	}
	**/
	//-------------------
	
	showReadFile = showReadFile + strconv.Itoa(numLines) + "<file>" + inpRowFile + ";" ; 
	
	fmt.Println( "read ", len(lineD) , " lines of file ", inpRowFile, " which contains ",  numLines, " text rows" );  	
	
	//testRowGroup()  // TEST
	//-------------------
	
	
} // end of  read_dictRow_file

//------------------------
/***
func set_gr_ixLast(firstIxRow int, lastIxRow int) {
	for x:= firstIxRow; x <= lastIxRow; x++ {
		rowLineIxList[x].ixR_ix_last = lastIxRow
	}
} // end of set_gr_ixLast()
***/
//-------------------------------------
/***
func sortLineRowIx() {	
	
	sort.Slice( rowLineIxList, func(i, j int) bool {
		return rowLineIxList[i].ixR_id < rowLineIxList[j].ixR_id 
	})
	
} // end of sortLineRowIx
***/
//--------------------------------
/****
func listLineRowIx() {

	var grIx rowIxStruct
	for x:=0; x < len(rowLineIxList); x++ {
		if x > 50 { break}
		grIx = rowLineIxList[x]
		*****
		grIx.ixR_id 
		grIx.ixR_id_gr
		grIx.ixR_id_num
		grIx.ixR_ix
		grIx.ixR_ix_last   
		****		
		fmt.Println("groupix ",grIx.ixR_id, " \t= ", grIx.ixR_id_gr, " \t ", grIx.ixR_id_num, " \t ", grIx.ixR_ix,  " \t ", grIx.ixR_ix_last)
	}    

} // end of listLineRowIx
*****/

//----------------------------------

func bind_go_passToJs_getIxRowFromGroup( rowGrIndex int,   html_rowGroup_beginNum int, html_rowGroup_numRows int, js_function string)  {

	fmt.Println( green("bind_go_passToJs_getIxRowFromGroup"), 
		"()  rowGrIndex=", rowGrIndex, ",  html_rowGroup_beginNum=", html_rowGroup_beginNum, ",  html_rowGroup_numRows=", html_rowGroup_numRows	) 
	
	if rowGrIndex < 0 { return }
	
	/**
	if (sw_list_Word_if_in_ExtrRow) { 		
			fmt.Println("on the row list button \"inpBegRow\" or \"maxNumRow\" values has been changed, but since only extracted rows wanted, that causes a rebuild of wordlist data")			
			build_and_elab_word_list()
	}
	**/
	
	//gRix := rowLineIxList[ rowGrIndex ] 
	
	rG := lista_gruppiSelectRow[ rowGrIndex ]	


	ixRowBeg := html_rowGroup_beginNum + rG.rG_firstIxRowOfGr - 1 	 
	ixRowEnd := ixRowBeg + html_rowGroup_numRows - 1
	if ixRowEnd > rG.rG_lastIxRowOfGr { 
		ixRowEnd = rG.rG_lastIxRowOfGr  
		html_rowGroup_numRows = 1 + ixRowEnd - ixRowBeg
	}

	last_rG_ixSelGrOption = rG.rG_ixSelGrOption	
	last_rG_group         = rG.rG_group	 
	last_rG_firstIxRowOfGr= rG.rG_firstIxRowOfGr
	last_rG_lastIxRowOfGr = rG.rG_lastIxRowOfGr 	   	
	last_ixRowBeg         = ixRowBeg	 
	last_ixRowEnd         = ixRowEnd	
	last_html_rowGroup_index_gr = rowGrIndex	
	last_html_rowGroup_beginNum = html_rowGroup_beginNum	// from the beginning of the group ( starting from 1 )
	last_html_rowGroup_numRows  = html_rowGroup_numRows	
	//last_html_rowGroup_numRows  = 1 + last_rS_toIxRow - ixRowBeg - (html_rowGroup_beginNum - 1)
	
	//last_word_fromWord    = 	
	//last_word_numWords    = 
	//last_sel_extrRow      = 
	
	write_lastValueSets()
	
	outS1:= fmt.Sprintf( "inp,%d,%d,%d,gr,%d,%s,%d,%d,ixr,%d,%d, %s",
				rowGrIndex, html_rowGroup_beginNum, html_rowGroup_numRows,
				rG.rG_ixSelGrOption, 
				rG.rG_group, 
				rG.rG_firstIxRowOfGr, 
				rG.rG_lastIxRowOfGr,
				ixRowBeg, ixRowEnd, 	
				inputTextRowSlice[  rG.rG_firstIxRowOfGr ].rRow1   )
	
	go_exec_js_function( js_function, outS1 )	
	
} // end of bind_go_passToJs_updateRowGroup



//-----------------------------
func write_lastValueSets() {
	
	//fmt.Println( green("write_lastValueSets") ) 
	
	outS1:= fmt.Sprint(
		"rG_ixSelGrOption="  + strconv.Itoa( last_rG_ixSelGrOption ) + ", " ,
		"rG_group="          + last_rG_group                         + ", " ,  
		"rG_firstIxRowOfGr=" + strconv.Itoa( last_rG_firstIxRowOfGr) + ", " ,     
		"rG_lastIxRowOfGr="  + strconv.Itoa( last_rG_lastIxRowOfGr ) + ", \n" ,      

		"html_rowGroup_index_gr=" + strconv.Itoa( last_html_rowGroup_index_gr ) + ", " ,
		"html_rowGroup_beginNum=" + strconv.Itoa( last_html_rowGroup_beginNum ) + ", " ,
		"html_rowGroup_numRows="  + strconv.Itoa( last_html_rowGroup_numRows  ) + ", \n" ,	
		
		"ixRowBeg="     + strconv.Itoa( last_ixRowBeg  ) + ", " , 
		"ixRowEnd="     + strconv.Itoa( last_ixRowEnd  ) + ", \n" , 
		
		"w_fromWord="             + strconv.Itoa( last_word_fromWord ) +", " ,	
		"w_numWords="             + strconv.Itoa( last_word_numWords ) +", \n" ,	

		"sel_extrRow="            + last_sel_extrRow )
	outS2:= []string{ outS1}	
	writeList(FOLDER_INPUT_OUTPUT  + string(os.PathSeparator) + FILE_last_mainpage_values2, outS2 )		
	
	fmt.Println( green("write_lastValueSets"), " in file ",  FILE_last_mainpage_values2    ) 
	
} // end of 
//----------------------------------
/***
func bind_go_passToJs_updateRowGroup( rowGrIndex int, inpBegRow int, inpNumRow int,  js_function string)  {

	fmt.Println( " func ", green("bind_go_passToJs_updateRowGroup") , rowGrIndex, inpBegRow, inpNumRow,  js_function) 
	
	gRix := rowLineIxList[ rowGrIndex ] 
	
	fmt.Println(   " gruppo [", rowGrIndex, "]=",  gRix )

	groupId      := gRix.ixR_id_gr 
	gr_index     := rowGrIndex
	fromRowGrBeg := inpBegRow
	numRowGr     := inpNumRow
	rowIndex     := gRix.ixR_ix 
	rowLastIndex := gRix.ixR_ix_last 
	idRow1 := inputTextRowSlice[ rowIndex ].rIdRow  
	row1   := inputTextRowSlice[ rowIndex ].rRow1   
    
	outS1:= fmt.Sprintf("%s::%d::%d::%d::%d::%d::%s::%s", groupId, gr_index, fromRowGrBeg, numRowGr, rowIndex,rowLastIndex, idRow1, row1) 
	
	//fmt.Println("bind_go_passToJs_updateRowGroup    lancia " + red(js_function)+ "(" + outS1 +")" )
	
	go_exec_js_function( js_function, outS1 )	
	
} // end of bind_go_passToJs_updateRowGroup
***/

//-----------------------------
/***
func testRowGroup() {
	
	***
	grId="9";   grNum=0;  ix = getIxOfRowGroupId( grId, grNum);    fmt.Println("gr=" , grId, grNum, "  ix=", ix, " row=", inputTextRowSlice[ix])
	grId="9";   grNum=1;  ix = getIxOfRowGroupId( grId, grNum); fmt.Println("gr=" , grId, grNum, "  ix=", ix, " row=", inputTextRowSlice[ix])
	grId="9";   grNum=2;  ix = getIxOfRowGroupId( grId, grNum); fmt.Println("gr=" , grId, grNum, "  ix=", ix, " row=", inputTextRowSlice[ix])
	grId="10";   grNum=2;  ix = getIxOfRowGroupId( grId, grNum); fmt.Println("gr=" , grId, grNum, "  ix=", ix, " row=", inputTextRowSlice[ix])
	grId="10";   grNum=4;  ix = getIxOfRowGroupId( grId, grNum); fmt.Println("gr=" , grId, grNum, "  ix=", ix, " row=", inputTextRowSlice[ix])
	grId="11";   grNum=5;  ix = getIxOfRowGroupId( grId, grNum); fmt.Println("gr=" , grId, grNum, "  ix=", ix, " row=", inputTextRowSlice[ix])
	grId="12";   grNum=0;  ix = getIxOfRowGroupId( grId, grNum); fmt.Println("gr=" , grId, grNum, "  ix=", ix, " row=", inputTextRowSlice[ix])
	grId="12";   grNum=1;  ix = getIxOfRowGroupId( grId, grNum); fmt.Println("gr=" , grId, grNum, "  ix=", ix, " row=", inputTextRowSlice[ix])
	grId="12";   grNum=3;  ix = getIxOfRowGroupId( grId, grNum); fmt.Println("gr=" , grId, grNum, "  ix=", ix, " row=", inputTextRowSlice[ix])
	*****
	printGr("9", 0 ) 
	printGr("10", 4) 
	
	
} // end of testRowGroup
***/
//------------------
/***
func printGr(grId string, grNum int) {
	rowGrIndex:= getIxOfRowGroupId( grId, grNum); 
	if rowGrIndex < 0 { return }
	rowGix := rowLineIxList[ rowGrIndex ]
		
	fmt.Println("gr=" , rowGix.ixR_id, grNum, "  ix=", rowGix.ixR_ix, " row=", inputTextRowSlice[ rowGix.ixR_ix]  )
}
***/
//----------------------------------
/***
func getIxOfRowGroupId( rowGrPref0 string, rowGrNum int ) int {
	
	rowGrPref := strings.TrimSpace( rowGrPref0 )
	
	rowGrId := rowGrPref + "_" + strconv.Itoa( BASE_rowGrNum +  rowGrNum )
	
	****
		gRix.ixR_id          = id_pref + "_" + strconv.Itoa( 100000 + id_num )
		gRix.ixR_ix          = len(inputTextRowSlice)-1
        rowLineIxList       = append(rowLineIxList, gRix) 
	*****	
	
	fromIx, toIx := lookForRowGroup( rowGrId )
	
	if toIx < 0 { return -1 }
	if toIx >=  len(rowLineIxList) {toIx = len(rowLineIxList) -1 }
	
	for g:=fromIx; g <= toIx; g++ {
		if rowLineIxList[g].ixR_id == rowGrId {
			return g; // rowLineIxList[g].ixR_ix 
		}
	}
	// se non trovata la corrispondenza esatta, cerco  il gruppo esatto (qualunque numero entro i limiti) 
	minGrId:= rowGrPref + "_" + strconv.Itoa( BASE_rowGrNum )
	maxGrId:= rowGrPref + "_" + strconv.Itoa( 2*BASE_rowGrNum )
	
	for g:=fromIx; g <= toIx; g++ {
		if rowLineIxList[g].ixR_id < minGrId { continue  }
		if rowLineIxList[g].ixR_id > maxGrId { return -1 }
		return g; // rowLineIxList[g].ixR_ix  
	}	
	return -1 
	
} // end of  getIxOfRowGroupId
***/

//-----------------------------
//------------------------------------
/***
func lookForRowGroup( rowGr string) (int, int) {

	// find 2 indices of the 2 words nearest to the word to find 
	
	low   := 0
	high  := len(rowLineIxList) - 1	
	maxIx := high; 
	
	//----
	for low <= high{
		median := (low + high) / 2		  
		if rowLineIxList[median].ixR_id < rowGr {  		
			low = median + 1
		}else{
			high = median - 1
		}
	} 
	//---
	fromIx:= low; toIx := high; 
	if fromIx > toIx { fromIx = high; toIx = low;}
	if fromIx < 0 { fromIx=0} 
	if toIx  > maxIx { toIx = maxIx}
	
	return fromIx, toIx	

} // end of lookForRowGroup
****/
//---------------------------

func antoRinumera( row string) bool {  // ?anto rinumera
    rowU:= strings.TrimSpace(row)
	rowL := strings.ToLower( rowU );
	if strings.Index(rowL, "http:") >=0 {return false}
	if strings.Index(rowL, "https:") >=0 {return false}
	if rowL == "" { return false}    
	if rowL == "---" { return false } 
	if len(rowU) < 10 { 
		if _, err := strconv.Atoi(rowU); err == nil {
			return false ; //   looks like a number.\n", v)
		}
		if strings.Index(" I II III IV V VI VII VIII IX X XI XII XIII XIV XV XVI XVII XVIII XIX XX ", rowU) >0  {  return false}
		if rowU[0:1] == "[" {
			if rowU[len(rowU)-1:] == "]" { return false} 
		}
	}
	return true
} // end of antoRinumera //?anto

//---------------

func read_dictLemmaTran_file(path1 string, inpFile string) {
	bytesPerRow:= 10
    lineD := rowListFromFile( path1, inpFile, "traduzione lemma", "read_dictLemmaTran_file", bytesPerRow)  
	if sw_stop { return }
	
	// 	abnutzbarkeit vestibilità   ==>  lemma     \t traduzione                                                                   |    |  |        
	
	lineZ := ""
	
	var ele1 lemmaTranStruct       //  lemmaTranStruct: dL_lemmaCod string,  dL_lemma2 string, dL_tran string  
	
	//---------------
	cod1:= "" 	
	lastNumDict ++;
	//-----------------
	for z:=0; z< len(lineD); z++ { 
		
		lineZ = strings.TrimSpace(lineD[z]) 
		
		// eg. abnutzbarkeit \t	vestibilità     ==>  lemma \t translation  
		
		if lineZ == "" { continue }
		j1:= strings.Index(lineZ, "|")
		if j1 < 0 { continue }
		cod1 = lineZ[0:j1]
		ele1.dL_lemmaCod = newCode(cod1) 
		ele1.dL_lemma2   = strings.TrimSpace( cod1 )
		ele1.dL_numDict  = lastNumDict 
		ele1.dL_tran     = strings.TrimSpace( lineZ[j1+1:] )   ////cigna1_2
		/**
		if z < 20 {
			fmt.Println("ANTONIO CARICA dictL ele1=", ele1, "\n\t", " ele1.dL_lemmaCod=" + ele1.dL_lemmaCod + ", lemma2=" +  ele1.dL_lemma2 + ", numDict=", ele1.dL_numDict, ",  tran=", ele1.dL_tran); 
		}
		**/
		
		dictLemmaTran = append( dictLemmaTran, ele1 ) 	
		
	}	
	
	fmt.Println("read_dictLemmaTran_file", " len(lineD)=", len(lineD), " len( dictLemmaTran)=", len(dictLemmaTran) )
	
	sort_lemmaTran2();
	
	fmt.Println( len(dictLemmaTran) , "  lemma - translation elements of  dictLemmaTran" , "( input: ", inpFile, ")"  )   
	
} // end of read_dictLemmaTran_file  


//---------------------------------------

//-----------------------
func read_lemma_file( path1 string, inpLemmaFile_wordLemma, inpLemmaFile_lemmaWord string) {
	
	//showInfoRead( inpLemmaFile, " inizio lettura " )
	
	bytesPerRow:=10
	numLemmaDict=0; 
	
	var wordLemma1 wordLemmaPairStruct 
	//------
	file1_bytes := getFileByteSize(path1, inpLemmaFile_wordLemma)
	file2_bytes := getFileByteSize(path1, inpLemmaFile_lemmaWord)
	fmt.Println("file ", inpLemmaFile_wordLemma, "  ", file1_bytes , " bytes") 
	fmt.Println("file ", inpLemmaFile_lemmaWord, "  ", file2_bytes , " bytes") 
	numEleMax:= int(  (file1_bytes + file2_bytes) / bytesPerRow ); 
	if numEleMax < 10 {numEleMax=10}
	//----------------
    lineS:= rowListFromFile( path1, inpLemmaFile_wordLemma, "1assoc. word-lemma", "read_lemma_file", bytesPerRow)  		
	
	wordLemmaPairTMP := make( []wordLemmaPairStruct, 0, numEleMax)
	
	if (sw_stop == false) {	
		// read word lemma
		for z:=0; z< len(lineS); z++ { 
			lineZ0 := strings.TrimSpace(lineS[z])   //  format:     word   lemma		
			if lineZ0 == "" {continue}
			lineZ := strings.ReplaceAll( lineZ0, "\t" , " ")			
			
			cols:= strings.Fields( strings.ToLower( lineZ ) )   // Fields   split using whitespace,  treats consecutive whitespace characters as a single separator		
			if len(cols) < 2 { continue } 
			wordLemma1.lWord2   = stdCode( cols[0] ) 		
			wordLemma1.lLemma   = stdCode( cols[1] )	
			
			if len(wordLemma1.lLemma) < 1 { continue;  } 
			if ((wordLemma1.lLemma == "-") || (wordLemma1.lLemma[0:1] < "A")) { continue;  }   // ignore number  
						
			wordLemma1.lWordCod = newCode( wordLemma1.lWord2)
			wordLemma1.lIxLemma = -1
			
			if ((  wordLemma1.lWord2 == "cäsar") || (wordLemma1.lWord2 == "caesar") || (wordLemma1.lWord2 == "casar") ) { fmt.Println(green(" 00 carica Lemma "), " newCode=",wordLemma1.lWordCod       ) }
			
			wordLemmaPairTMP = append(wordLemmaPairTMP, wordLemma1 ) 
			numLemmaDict++		
		}
		fmt.Println(" read ", len(lineS), " input lemma: format word-lemma")
	}
	
	//----------------
    lineS = rowListFromFile( path1, inpLemmaFile_lemmaWord, "2assoc. lemma-word", "read_lemma_file", bytesPerRow)  		
	
	if (sw_stop == false) {	
		// read word lemma
		for z:=0; z< len(lineS); z++ { 
			lineZ0 := strings.TrimSpace(lineS[z])   //  format:     word   lemma		
			if lineZ0 == "" {continue}
			lineZ := strings.ReplaceAll( lineZ0, "\t" , " ")			
			
			cols:= strings.Fields( strings.ToLower( lineZ ) )   // Fields   split using whitespace,  treats consecutive whitespace characters as a single separator		
			if len(cols) < 2 { continue } 
			wordLemma1.lLemma   = stdCode( cols[0] )					
			wordLemma1.lWord2   = stdCode( cols[1] )
			
			if len(wordLemma1.lLemma) < 1 { continue;  } 
			if ((wordLemma1.lLemma == "-") || (wordLemma1.lLemma[0:1] < "A")) { continue;  } 		
			
			wordLemma1.lWordCod = newCode( wordLemma1.lWord2)
			wordLemma1.lIxLemma = -1
			
			if ((  wordLemma1.lWord2 == "cäsar") || (wordLemma1.lWord2 == "caesar") || (wordLemma1.lWord2 == "casar") ) { fmt.Println(green(" 00 carica Lemma "), " newCode=",wordLemma1.lWordCod       ) }
	
			
			wordLemmaPairTMP = append(wordLemmaPairTMP, wordLemma1 ) 
			numLemmaDict++		
		}
		fmt.Println(" read ", len(lineS), " input lemma: format lemma-word")
	}
	lineS = nil
	//---------------------------------------	
	//-----		
	fmt.Println( "lette " , numLemmaDict ,  " coppie word-lemma", "\n")
	//-----	
	// sort x lemma, word

	sort.Slice(wordLemmaPairTMP, func(i, j int) bool {
			if (wordLemmaPairTMP[i].lLemma != wordLemmaPairTMP[j].lLemma) {
				return wordLemmaPairTMP[i].lLemma < wordLemmaPairTMP[j].lLemma
			} else {
				return wordLemmaPairTMP[i].lWord2 < wordLemmaPairTMP[j].lWord2
			}
		} )	 	
	//-------------------------------------
	preLemS:=  ""
	lemS:= ""	
		
	wordLemmaPair = make( []wordLemmaPairStruct, 0, len(wordLemmaPairTMP)	)
	
	doppi:=0
 
	for nn1, lemX := range wordLemmaPairTMP {
		lemS = lemX.lLemma + " " + lemX.lWord2 
		if ((  lemX.lWord2 == "cäsar") || (lemX.lWord2 == "caesar") || (lemX.lWord2 == "casar") ) { fmt.Println(" 111 carica Lemma ", nn1,  " lemX=" , lemX) }
		if preLemS == lemS {
			doppi++
			continue
		}	
		preLemS = lemS
		wordLemmaPair = append( wordLemmaPair, lemX)
	}
	if doppi > 0 {
		fmt.Println(" scartate ", doppi, " entrate doppie in lemma - word ") 
	}
	numLemmaDict = len(wordLemmaPair)
	
	fmt.Println( "caricate " , numLemmaDict ,  " coppie word-lemma", "\n")
	
	//---------------------------
	wordLemmaPairTMP = nil; 
	//---------------------------
	/*
	//---
		type lemmaStruct struct {
			leLemma    string
			leNumWords int 
			leFromIxLW  int 
			leToIxLW    int  
			leTran     string  
		} 
		//-------------------------------
		type wordLemmaPairStruct struct {
			lWordCod string 
			lWord2   string 
			lLemma   string
			lIxLemma int
		} 
		//---
	*/
	numW:=0
	preLem:=""
	
	lemmaSlice = nil
	var leV lemmaStruct; 
	fromIx:=0; toIx:=0
	iixLem:=0	
	
	fmt.Println("  len(wordLemmaPair)=", len(wordLemmaPair) )
		
	for z:=0; z< len(wordLemmaPair); z++ {
		
		if (wordLemmaPair[z].lLemma != preLem) {
			if numW > 0 {
				//scrive lemma precedente 				
				leV.leLemma    = preLem
				leV.leNumWords = 0 
				leV.leTran     = ""
				lemmaSlice = append(lemmaSlice, leV ) 
				iixLem = len(lemmaSlice) -1 
				for h:=fromIx; h<= toIx; h++ {
					wordLemmaPair[h].lIxLemma = iixLem 
				}	
			}
			preLem = wordLemmaPair[z].lLemma			
			numW=0
			fromIx=z;  
		} 
		numW++
		toIx=z
	}	
	if numW > 0 {
				//scrive lemma precedente 				
				leV.leLemma    = preLem
				leV.leNumWords = 0 
				leV.leTran     = ""
				lemmaSlice = append(lemmaSlice, leV ); 
				iixLem = len(lemmaSlice) -1 
				for h:=fromIx; h<= toIx; h++ {
					wordLemmaPair[h].lIxLemma = iixLem 
				}			
	}
	
	//-------------------------------------
	
	//----------------------
	seq:=""; //swerr:=false
	for  _, lem := range lemmaSlice {
		if lem.leLemma < seq {
			fmt.Println(red("ERRORE lemmaSlice fuori sequenza "), " pre=", seq, "   new=", lem.leLemma )
			//swerr = true
			break;
		}
		seq = lem.leLemma
	}
	
	fmt.Println( green("lemmaSlice"), "  composto da ", len(lemmaSlice) , " elementi")    
	
	//if swerr == false { fmt.Println( green("lemmaSlice IN SEQUENZA")) 	}
	
	//---------------------------------
	/**
	fmt.Println( green("lemmaSlice[]"))
	for  il, lem := range lemmaSlice {
		if il > 10 { break }
		fmt.Printf("lemma index %d, %s, %d words, trans %s\n", il, lem.leLemma, lem.leNumWords, lem.leTran )    
	}
	**/
	/**
	fmt.Println( green("wordLemmaPair[]"))
	
	for  il2, lem2 := range wordLemmaPair {
		if il2 > 40 { break }
		
		fmt.Println(lem2.lWord2, "\t" , "lemma=",lem2.lLemma, "\t ixLemma=", lem2.lIxLemma,  " \t lemmaSlice[]=", lemmaSlice[  lem2.lIxLemma ].leLemma, " n.words with this lemma=",  lemmaSlice[  lem2.lIxLemma ].leNumWords)    
	}
	***/
	//--------------------------------
	// sort x word , lemma 
	sort.Slice(wordLemmaPair, func(i, j int) bool {
			if (wordLemmaPair[i].lWordCod != wordLemmaPair[j].lWordCod) {
				return wordLemmaPair[i].lWordCod < wordLemmaPair[j].lWordCod
			} else {
				if (wordLemmaPair[i].lWord2 != wordLemmaPair[j].lWord2) {
					return wordLemmaPair[i].lWord2 < wordLemmaPair[j].lWord2 
				} else {
					return wordLemmaPair[i].lLemma < wordLemmaPair[j].lLemma
				}
			}
		} )	 
	//------------------------	
	
	check_wordLemma_sameCode()
	
	//-------------------------------
} // end of  read_lemma_file

//-------------------------------------------

func check_wordLemma_sameCode() {
	fmt.Println( green("check_wordLemma_sameCode") , "()"  )
	// check same words  written in diffent way (eg. caesar   and  "cäsar")
	pre_wordCod := ""
	pre_word2   := ""	
	//pre_lemma   := ""
	pre_z := -1
	
	for z, wordPair := range wordLemmaPair {	
			if ((  wordPair.lWord2 == "cäsar") || (wordPair.lWord2 == "caesar") || (wordPair.lWord2 == "casar") ) { fmt.Println(" check 222 Lemma ", z,  " wordPair=" , wordPair) }
	
		if (wordPair.lWordCod != pre_wordCod) {
			pre_wordCod = wordPair.lWordCod 
			pre_word2   = wordPair.lWord2 
			//pre_lemma   = wordPair.lLemma 
			pre_z = z
			continue
		}
		if (wordLemmaPair[z].lWord2 == pre_word2) {
			continue
		}
		//--------
		fmt.Println( green("check_wordLemma_sameCode") )
		for x:= pre_z; x<= z; x++ {
			fmt.Println("\t", " wordLemmaPair[",x,"] = ", wordLemmaPair[x] )   
		} 
		
	}	
	//   ???antoX   if ((  lemX.lWord2 == "cäsar") || (lemX.lWord2 == "caesar")  || (lemX.lWord2 == "casar")) { fmt.Println(" 111 carica Lemma ", nn1,  " lemX=" , lemX) }
}

//---------------------------------

func addUnknowToLemma( lemma1 string) int {
	var leV lemmaStruct 
	leV.leLemma    = lemma1
	leV.leNumWords = 0; 
	leV.leTran     = ""
	lemmaSlice = append(lemmaSlice, leV ); 
	return len(lemmaSlice) -1 
} 

//-----------------------------------------
func OLDread_lemma_file( path1 string, inpLemmaFile string) {
	
	//showInfoRead( inpLemmaFile, " inizio lettura " )
	
	bytesPerRow:=10

    lineS:= rowListFromFile( path1, inpLemmaFile, "assoc. word-lemma", "read_lemma_file", bytesPerRow)  	
	
	//showInfoRead( inpLemmaFile, " fine lettura " )
	
	if sw_stop { return }
	
	var wordLemma1 wordLemmaPairStruct 
	numLemmaDict=0; 
	
	for z:=0; z< len(lineS); z++ { 
		lineZ0 := strings.TrimSpace(lineS[z]) 
		
		if lineZ0 == "" {continue}
		lineZ := strings.ReplaceAll( lineZ0, "\t" , " ")  // space as separator
		j1:= strings.Index(lineZ, " ")
		if j1 < 0 { continue }	
		wordLemma1.lWord2   = strings.TrimSpace(lineZ[0:j1])
		wordLemma1.lWordCod = newCode( wordLemma1.lWord2)
		wordLemma1.lLemma  = strings.TrimSpace( lineZ[j1+1:]) 
	
		if wordLemma1.lLemma == "-" { continue } 
		wordLemmaPair = append(wordLemmaPair, wordLemma1 ) 
		numLemmaDict++
	}
	//-----
		
	fmt.Println( "caricate " , numLemmaDict ,  " coppie word-lemma dal file ", inpLemmaFile, "\n")
	//-----	
	
	sort.Slice(wordLemmaPair, func(i, j int) bool {
			if (wordLemmaPair[i].lWordCod != wordLemmaPair[j].lWordCod) {
				return wordLemmaPair[i].lWordCod < wordLemmaPair[j].lWordCod
			} else {
				if (wordLemmaPair[i].lWord2 != wordLemmaPair[j].lWord2) {
					return wordLemmaPair[i].lWord2 < wordLemmaPair[j].lWord2 
				} else {
					return wordLemmaPair[i].lLemma < wordLemmaPair[j].lLemma
				}
			}
		} )	 

	
} // end of  OLDread_lemma_file

//---------------------------------

func sort_lemmaTran2() {
	
	if len(dictLemmaTran) < 1 { return }
	
	sort.Slice(dictLemmaTran, func(i, j int) bool {
			if (dictLemmaTran[i].dL_lemmaCod != dictLemmaTran[j].dL_lemmaCod) { 
				return dictLemmaTran[i].dL_lemmaCod < dictLemmaTran[j].dL_lemmaCod 
			} else {
				return dictLemmaTran[i].dL_numDict < dictLemmaTran[j].dL_numDict				
			}
		} )		
	//------------	
	var pre lemmaTranStruct
	pre = dictLemmaTran[0]
	
	for g2:=1; g2 < len(dictLemmaTran); g2++ {
		if (dictLemmaTran[g2].dL_lemmaCod == pre.dL_lemmaCod) {
			dictLemmaTran[g2 -1].dL_lemmaCod = LAST_WORD; 
		} 
		pre = dictLemmaTran[g2] 
	}
	//--------------------------
	sort.Slice(dictLemmaTran, func(i, j int) bool {
			if (dictLemmaTran[i].dL_lemmaCod != dictLemmaTran[j].dL_lemmaCod) { 
				return dictLemmaTran[i].dL_lemmaCod < dictLemmaTran[j].dL_lemmaCod 
			} else {
				return dictLemmaTran[i].dL_numDict < dictLemmaTran[j].dL_numDict				
			}
		} )		
	//--------------------------------	
	var numLin = 0
	for g2:=0; g2 < len(dictLemmaTran); g2++ {
		if (dictLemmaTran[g2].dL_lemmaCod == LAST_WORD) {
			numLin = g2;
			break; 	
		} 
	}
	if numLin > 0 {
		dictLemmaTran = dictLemmaTran[0:numLin+1]
	}
	
	
} // end of sort_lemmaTran2() 

//-------------------------

func fromLemmaTo3List( lemma string) (string, string, string) { 

		fromIx, toIx  := lookForAllParadigma( lemma ) 
		
		//fmt.Println( "         fromIx=", fromIx, "  toIx=", toIx )
		
		listLev := ""
		listPara:= ""
		listExam:= ""		
		for ix:= fromIx; ix <= toIx; ix++ {
			listLev  += "|" + lemma_para_list[ix].p_level   
			parad    := lemma_para_list[ix].p_para			
			if len(parad) > 2 { 				
				if parad[0:lenFseq] == fseq {
					parad = parad[lenFseq:] 
				}
			}
			listPara += "|" + parad  		 	
			listExam += "|" + lemma_para_list[ix].p_example 	
	 		//fmt.Println("\t ", ix, " listLev =", listLev)    
		}	
		if toIx >= fromIx {
			listLev = listLev[1:]
			listPara= listPara[1:]	
			listExam = listExam[1:]	
		}
		return listLev, listPara, listExam 
		
} // end of fromLemmaTo3List 		

//===========================================================================
func addWordLemmaTranLevelParadigma() {
	
	//fmt.Println("addWordLemmaTranLevelParadigma()" )
	
	newWordLemmaPair = make( []wordLemmaPairStruct, 0, len(uniqueWordByFreq)  )		
	var newWL wordLemmaPairStruct   // 	lWordCod string, lWord2 string , lLemma string
	var wP lemmaTranStruct  
		
	lemma_word_ix = make([]lemmaWordStruct, 0,  len(uniqueWordByFreq)  )  
	var LW lemmaWordStruct
	
	list1Level:= ""
	list1Para := ""
	list1Exam := ""	
	lemma3:= ""
	
	//var swMio bool = false
	
	//--------------------------------------
	for zz:=0; zz < len(uniqueWordByFreq); zz++ {
		wF:= uniqueWordByFreq[zz]
		
		
		swprova:= ((wF.uWord2 == "cäsar") || (wF.uWord2 == "caesar") || (wF.uWord2 == "casar")) 
		
		
		
		ixLemmaPairFoundList := lookForAllLemmas( wF.uWord2 ) // 
		
		if swprova { fmt.Println("1 loop unique x lemma ", wF.uWord2,  " ixLemmaPairFoundList=",  ixLemmaPairFoundList ) }

				/**
						//
						var lemmaSlice       [] lemmaStruct         // lemma , translation 
						var wordLemmaPair    [] wordLemmaPairStruct // all word-lemma pair  
						//---
						type lemmaStruct struct {
							leLemma    string
							leNumWords int 
							leFromIxLW  int 
							leToIxLW    int  
							leTran     string  
						} 
						//---------
						type lemmaWordStruct struct {
							lw_lemmaCod string 	
							lw_lemma2   string 	
							lw_word     string 
							lw_ixLemma    int
							lw_ixWordUnFr int
						}
						//-------------------------------
						type wordLemmaPairStruct struct {
							lWordCod string 
							lWord2   string 
							lLemma   string
							lIxLemma int
						} 
						//---
				**/ 
		nele := len(ixLemmaPairFoundList)
		
		if swprova { fmt.Println("2 loop unique x lemma  nele=", nele) }
		
		lis_ixLemma:=make( [] int,    0, nele )		
		lis_lemma := make( [] string, 0, nele )		
		lis_tran  := make( [] string, 0, nele )
		lis_level := make( [] string, 0, nele )
		lis_para  := make( [] string, 0, nele )
		lis_exam  := make( [] string, 0, nele )
		

		ixLemma:=-1
		numLerr:=0; maxNumLerr:=100; 
		for  _, ixLp := range ixLemmaPairFoundList { 
			
			if swprova { fmt.Println("\t3 loop unique x lemma  ixLemmaPairFoundList=", ixLp) }
			
			if numLerr > maxNumLerr { break}
			if ixLp < 0 {
				lemma3 = "?" + wF.uWord2	
				ixLemma = addUnknowToLemma(lemma3) 
			    if swprova { fmt.Println("\t4 loop unique x lemma  ixLemma=", ixLemma, " lemma3=", lemma3, " lemmaSlice[ixLemma]=", lemmaSlice[ixLemma]) }
			} else {
				newWL = wordLemmaPair[ixLp]
				if newWL.lWord2 != wF.uWord2 { // error 
					continue; 
				}
				ixLemma = newWL.lIxLemma 
				lemma3  = newWL.lLemma
			}
			lis_ixLemma = append( lis_ixLemma, ixLemma  )   
			lis_lemma   = append( lis_lemma,   lemma3   )  
			if (ixLemma < 0) {
					fmt.Println( red("errore1 "), " in addWordLemmaTranLevelParadigma" , " word=",wF.uWord2, " ixLp=", ixLp, " ixLemma=",ixLemma, 
					" lemma3=", lemma3, red("  ixLemma negativo") )
					numLerr++
					continue; 	
			}
			if lemmaSlice[ixLemma].leLemma != lemma3 {
				fmt.Println( red("errore2 "), " in addWordLemmaTranLevelParadigma" , " word=",wF.uWord2, " ixLp=", ixLp, " ixLemma=",ixLemma, 
					" lemma3=", lemma3," lemmaSlice[ixLemma].leLemma=",lemmaSlice[ixLemma].leLemma , red(" lemma non eguale") )  
				numLerr++	
				continue
			}	
			lemmaSlice[ixLemma].leNumWords++
		}
		if numLerr > 0 {
			fmt.Println( "trovati ", numLerr, " ", red("errori"), " in addWordLemmaTranLevelParadigma")
		}
		
		newWL.lWord2   = wF.uWord2
		newWL.lWordCod = newCode( wF.uWord2 )
		newWL.lLemma = ""
		newWL.lIxLemma = -1
		
		//swMio = (newWL.lWord2 == "erklären")   //antonio123
		
		// each word can have many lemmas 
		//                       each lemma can have many levels, paradigmas, translations ( they are separated by "|", eg "A1|A2|B1" ) 
		//---------------
		for  ixix, lem := range lis_lemma { 
			ixLemma2:= lis_ixLemma[ixix] 
			if ixLemma2 < 0 { continue }  // error 
			if lemmaSlice[ixLemma2].leLemma != lem { ixLemma2 = -1 }
			list1Level, list1Para, list1Exam = fromLemmaTo3List( lem )  // ogni list può contenere più elementi separati da |   ( lo stesso num.di elementi per tutte le liste: A1|A2,par1|par2, ex1|ex2 )  
			lis_level = append( lis_level, list1Level )
			lis_para  = append( lis_para , list1Para  )
			lis_exam  = append( lis_exam , list1Exam  )
			ixTra := lookForAllTran( lem ) 
			if ixTra >= 0 { 
				wP = dictLemmaTran[ixTra] 
				lis_tran = append( lis_tran, wP.dL_tran ) 		////cigna1	
				if ixLemma2 >=0 {	lemmaSlice[ixLemma2].leTran = wP.dL_tran } 
			} else {
				lis_tran = append( lis_tran, ""         ) 	
				if ixLemma2 >=0 {	lemmaSlice[ixLemma2].leTran = "" }  	
			}	
			if sw_rewrite_wordLemma_dict { 
				newWL.lLemma = lem 
				newWordLemmaPair = append( newWordLemmaPair, newWL ) 				
			}			
			//--			
			LW.lw_lemma2     = lem
			LW.lw_lemmaCod   = newCode( lem )			
			LW.lw_word       = wF.uWord2 
			LW.lw_ixLemma    = ixLemma2
			LW.lw_ixWordUnFr = wF.uIxUnW
			lemma_word_ix  = append( lemma_word_ix, LW )		
			
		} // end of for , lem 
		//-----------
		
		//if swMio {  fmt.Println("ANTONIO  word=", newWL.lWord2, " ixLemma=",lis_ixLemma,   " uLemmaL=" , lis_lemma, ",    tran=", lis_tran) }
		
		wF.uIxLemmaL= make( []int,     nele, nele )    
		wF.uLemmaL  = make( []string,  nele, nele )    
		//wF.uTranL   = make( []string,  nele, nele )        
		wF.uLevel   = make( []string,  nele, nele )   
		wF.uPara    = make( []string,  nele, nele )   
		wF.uExample = make( []string,  nele, nele )   
		
		
		copy( wF.uIxLemmaL, lis_ixLemma )		
		copy( wF.uLemmaL  , lis_lemma ) 
		//copy( wF.uTranL   , lis_tran  )  
		copy( wF.uLevel   , lis_level ) 
		copy( wF.uPara    , lis_para  ) 
		copy( wF.uExample , lis_exam  ) 
		uniqueWordByFreq[zz] = wF
		
		//if swMio {  fmt.Println("ANTONIO unique word=", newWL.lWord2, " wFixLemma=",wF.uIxLemmaL,   " wF uLemmaL=" ,wF.uLemmaL ) }  // , ",    tran=",  wF.uTranL ) }
	
		/**
		if ((wF.uWord2== "tun") || (wF.uWord2 == "umwelt") ) {
				fmt.Println("lookFromLemma( uniqueWordByFreq[",zz,"] = ", uniqueWordByFreq[zz]) 
		} 
		**/
		//if swMio {  fmt.Println("ANTONIO  2 wF=", wF) }
		
	}  // end of for zz 
	//-----------
	
	if len(lemmaNotFoundList) > 0 {
		outFile := FOLDER_OUTPUT +  string(os.PathSeparator) + FILE_outLemmaNotFound;
		writeList( outFile, lemmaNotFoundList )		
	}
	
	//countWordLemmaUse() 
	
} // end of addWordLemmaTranLevelParadigma

//-----------------------------------
func stat_level( lemmaLevel []string, numWords int) {	
	
	// get the first level of the first lemma 
	
	if len(lemmaLevel) < 1 { return }
	if numWords < 1 { return }
	
	level2 := strings.Split( lemmaLevel[0], "|" ) 
	if len(level2) < 1 { return }
	
	levelToText := level2[0]
	
	sw_oth:=true
	for m:=0; m < len(list_level); m++ {	
		if levelToText == list_level[m] {
			only_level_numWords[m] += numWords 
			sw_oth = false; 
			break
		} 
	} 
	if sw_oth {
		only_level_numWords[0] += numWords 
	}
	
}
//------
func add_ixWord_to_WordSliceFreq() {
	var ixFromList, ixToList int
	tot:=0
	for ixWord:=0; ixWord < len(uniqueWordByFreq); ixWord++ {		
		xWordF := uniqueWordByFreq[ixWord] 			
		stat_level( xWordF.uLevel, xWordF.uTotRow)		
		tot+=  xWordF.uTotRow
		ixFromList = xWordF.uIxWordFreq 
		ixToList   = ixFromList + xWordF.uTotRow;
		if ixToList > numberOfWords { ixToList = numberOfWords; }		
		for n1 := ixFromList; n1 < ixToList; n1++  {
			wordSliceFreq[n1].wIxUniq = ixWord 			
		} 	
	}  
	
	fmt.Println(" num. words tutte = ", tot) 
	
	/**
	percA0  =  only_A0 * 100 / tot ;     
	percA1  =  only_A1 * 100 / tot ;     
	percA2  =  only_A2 * 100 / tot ;     
	percB1  =  only_B1 * 100 / tot ;     
	percOth =  only_Ot* 100 / tot ;   
	
	fmt.Println(
		  " num. words A0 = ", only_A0,    " \t", percA0 , "%" , 
		"\n num. words A1 = ", only_A1,    " \t", percA1 , "%" ,       
		"\n num. words A2 = ", only_A2,    " \t", percA2 , "%" ,       
		"\n num. words B1 = ", only_B1,    " \t", percB1 , "%" ,        
		"\n num. words altro= ", only_Ot, " \t", percOth, "%"  ) 
	****/	
	
	for f:=1; f < len( only_level_numWords ) ; f++ {
		if only_level_numWords[f] == 0 { continue }
		perc_level[f] = only_level_numWords[f] * 100 / tot ;     
		fmt.Println(f, " num. words ","list_level[f]", list_level[f], " = " , "only_level_numWords[f]=", only_level_numWords[f],    " \t", perc_level[f] , "%" ) 
	}	

	if only_level_numWords[0] > 0 {  
		perc_level[0] = only_level_numWords[0] * 100 / tot ;     
		fmt.Println(" num. words ", list_level[0], " = ", only_level_numWords[0],    " \t", perc_level[0] , "%" ) 	
	}

	
} // end of add_index_toWordSliceFreq

//---------------
//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

func lookForAllTran ( lemma30 string ) int {
	lemma3Cod:= newCode(lemma30)	
	fromIxX, toIxX := lookForTranslation( lemma3Cod )
	if toIxX < 0 { return -1 }
	
	
	z:=-1
	fromIx:= fromIxX
	for k:= fromIxX; k >= 0; k-- {
		if dictLemmaTran[k].dL_lemmaCod == lemma3Cod {
			z=k
			break	
		}
		if dictLemmaTran[k].dL_lemmaCod < lemma3Cod { break }
		fromIx = k
	}
	if z < 0 {
		for k:= fromIx; k < len( dictLemmaTran); k++ {
			if dictLemmaTran[k].dL_lemmaCod == lemma3Cod {
				z=k
				break
			}
		} 
	}
	
	if z < 0 { return -1 }
	
	return z 
	
} // end of lookForAllTran

//-----------------------------

func lookForTranslation(lemmaToFindCod string) (int, int) {

	// find 2 indices of the 2 words nearest to the word to find 
	
	low   := 0
	high  := len(dictLemmaTran) - 1	
	maxIx := high; 
	
	//----
	for low <= high{
		median := (low + high) / 2
		if dictLemmaTran[median].dL_lemmaCod < lemmaToFindCod {  
			low = median + 1
		}else{
			high = median - 1
		}
	} 
	//---
	fromIx:= low; toIx := high; 
	if fromIx > toIx { fromIx = high; toIx = low;}
	if fromIx < 0 { fromIx=0} 
	if toIx  > maxIx { toIx = maxIx}
	
	return fromIx, toIx	

} // end of lookForTranslation


//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
//------------------------------------------
func lookForAllLemmas(  wordToFind string) []int {

	wordToFindCod:= newCode( wordToFind )	
	ixFoundList := lookForAllLemmas2(  wordToFindCod ) 
	if len(ixFoundList) > 0 {
		return ixFoundList
	}
	
	inp1:= strings.ReplaceAll( 
				strings.ReplaceAll( 
						strings.ReplaceAll( 
							strings.ReplaceAll( wordToFind, "ae","ä"),  
							"oe","ö"), 
						"ue","ü"), 		
				"ss","ß") 	
				
	wordToFindCod2 := newCode( inp1 )
	ixFoundList2 := lookForAllLemmas2(  wordToFindCod2 ) 
	if len(ixFoundList2) == 0 {
		ixFoundList2 = append( ixFoundList,  -1)	// if lemma is missing use the original word to find 
		lemmaNotFoundList = append( lemmaNotFoundList, wordToFind ) 
	} 	
	return ixFoundList2 
	
} // end of lookForAllLemmas
//------------------------------------------
func lookForAllLemmas2(  wordToFindCod string) []int {

	//wordToFindCod:= newCode( wordToFind )

	// get the index of a word in word-lemma dictionary (-1 if not found)  
	var ixFoundList = make( []int, 0,0) 
	
	if len(wordLemmaPair) == 0 { return ixFoundList}
	
	fromIxX, toIx := lookForWordLemmaPair(wordToFindCod)
	
	if toIx < 0 { return ixFoundList }
	
	fromIx:= fromIxX
	
	for k:= fromIxX; k >= 0; k-- {
		if wordLemmaPair[k].lWordCod < wordToFindCod { break }
		fromIx = k
	}
	for k:= fromIx; k < len(wordLemmaPair); k++ {
		if wordLemmaPair[k].lWordCod == wordToFindCod {
			ixFoundList = append( ixFoundList, k) ; //    wordLemmaPair[k].lLemma )	
		} else {
			if wordLemmaPair[k].lWordCod > wordToFindCod { break }
		}
	} 
			

	//fmt.Println("lookForLemma( ==>" + wordToFind + "<== lemmaList=" , lemmaList, "   numLemmaDict=" , numLemmaDict)
	
	return ixFoundList 	
	
} // end of lookForAllLemmas2

//-----------------------------

func lookForWordLemmaPair(wordToFindCod string) (int, int) {
	
	// find 2 indices of the 2 words nearest to the word to find 
	
	low   := 0
	high  := numLemmaDict - 1	
	maxIx := high; 
	
	if high < 1 { return -1, -1 } 
	
	//----
	for low <= high{
		median := (low + high) / 2
		if wordLemmaPair[median].lWordCod < wordToFindCod {  
			low = median + 1
		}else{
			high = median - 1
		}
	} 
	//---
	fromIx:= low; toIx := high; 
	if fromIx > toIx { fromIx = high; toIx = low;}
	
	if fromIx < 0 { fromIx=0} 
	
	if toIx  > maxIx { toIx = maxIx}
	return fromIx, toIx	

} // end of lookForWordLemmaPair
//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
//------------------------------------------
/****
func countWordLemmaUse() {
	
	ixList:= [9]int{ 1000, 2000, 3000, 4000,5000,6000,7000,8000,9000 }
	
	for i:=0; i < len( ixList); i++ {
		if ixList[i] >= len(uniqueWordByFreq) { countLemma( len(uniqueWordByFreq)-1 ) ; break}
		countLemma( ixList[ i] ) 
	}   
} 
//-----------------	
func countLemma(ixMax int) { 			
		
		accum:= make([]wordIxStruct , 0,0 )
		if ixMax >  len(uniqueWordByFreq) {  ixMax = len(uniqueWordByFreq) }
		for zz:=0; zz < ixMax; zz++ {	
			accum = append( accum, uniqueWordByFreq[zz]  ) 
		}
		//----
		sort.Slice(accum, func(i, j int) bool {
			return accum[i].wLemma < accum[j].wLemma
		})	
		//---
		preL := ""; numLem:=0;
		var oneR  wordIxStruct
		oneR = accum[0]
		preL = oneR.wLemma 
		for zz:=0; zz < ixMax; zz++ {	
			oneR = accum[zz]
			if oneR.wLemma != preL { 
				numLem++
				preL = oneR.wLemma
			} 
		}
		numLem++
		//fmt.Println( "CONTA LEMMA: ", ixMax, " parole,"  , numLem, " lemma")       
		
} // end of countLemma
***/
//------------------

func writeAllUsedRowsOfFile2() {
	
	var swWriteUsedRow = (maxNumLinesToWrite > 0)
	
	fmt.Println("maxNumLinesToWrite=" ,  maxNumLinesToWrite, " swWriteUsedRow =", swWriteUsedRow)
	
	var maxNum = 100
	if swWriteUsedRow { maxNum = maxNumLinesToWrite }
		
	var strOut = ""
	
	nOut  :=0
	nOut0 :=0
	nOut1 :=0 
	
	listMax := [10]int {10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	
	fmt.Println("\nwriteAllUsedRowsOfFile2() swWriteUsedRow=" , swWriteUsedRow , " countNumLines=" ,countNumLines ); 
	
	if  swWriteUsedRow == false {		
		if countNumLines == false { return }
		for z:= 0; z < len(listMax); z++ {   	
			findHowManyUsedRowsOfFile2( listMax[z] )
		}
		fmt.Println("\nxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n" )
		return
	}
			
	for ixWord:= 0; ixWord < len(uniqueWordByFreq); ixWord++ {
		var xWordF     = uniqueWordByFreq[ ixWord ]  		
		var ixFromList = xWordF.uIxWordFreq 
		var ixToList   = ixFromList + xWordF.uTotRow;
		var maxTo1     = ixFromList + maxNum 		
		if ixToList > maxTo1        { ixToList = maxTo1; }
		if ixToList > numberOfWords { ixToList = numberOfWords; }
		
		for n1 := ixFromList; n1 < ixToList; n1++  {
			wS1 := wordSliceFreq[n1] 
			ixRR:= wS1.wIxRow
			isUsedArray[ ixRR ] = true
			//rowX := inputTextRowSlice[ixRR].row1		
		}
	}	
	//----------------------
	nOut  =0
	nOut0 =0
	nOut1 =0 
	for n2:= 0; n2  < len(inputTextRowSlice); n2++ {
		if isUsedArray[n2] {
			nOut++
			if inputTextRowSlice[n2].rNfile1 == 0 {
				nOut0++
			} else { 
				nOut1++
				if swWriteUsedRow {
					strOut += inputTextRowSlice[n2].rRow1 + "\n"
				}
			}
		} 
	}
	fmt.Println( "\nxxxxxxxxxxxxxxxxxxxxxx\nnumber of rows used (maxNum=", maxNum, ")\nmain text file =\t", nOut0, "\nother text file =\t", nOut1, "\nall text files =\t", nOut, "\nxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n" )
	
	if swWriteUsedRow {		
		outFileName := "wrk_USED_ROWS.txt"
		
		f, err := os.Create( outFileName )
		check(err)
		defer f.Close();

		_, err = f.WriteString( strOut )
		check(err)
		fmt.Println(" file ", outFileName , " written\n") 
	}
	
} // end of writeAllUsedRowsOfFile2
//-------------------------------------------

func findHowManyUsedRowsOfFile2( maxNum int ) {
	
	nOut  :=0
	nOut0 :=0
	nOut1 :=0 
		
	for ixWord:= 0; ixWord < len(uniqueWordByFreq); ixWord++ {
		var xWordF     = uniqueWordByFreq[ ixWord ]  		
		var ixFromList = xWordF.uIxWordFreq 
		var ixToList   = ixFromList + xWordF.uTotRow;
		var maxTo1     = ixFromList + maxNum 		
		if ixToList > maxTo1        { ixToList = maxTo1; }
		if ixToList > numberOfWords { ixToList = numberOfWords; }
		
		for n1 := ixFromList; n1 < ixToList; n1++  {
			wS1 := wordSliceFreq[n1] 
			ixRR:= wS1.wIxRow
			isUsedArray[ ixRR ] = true
			//rowX := inputTextRowSlice[ixRR].row1		
		}
	}	
	//----------------------
		nOut  =0
		nOut0 =0
		nOut1 =0 
		for n2:= 0; n2  < len(inputTextRowSlice); n2++ {
			if isUsedArray[n2] {
				nOut++
				if inputTextRowSlice[n2].rNfile1 == 0 {
					nOut0++
				} else { 
					nOut1++
				}
			} 
		}
		fmt.Println( "\nxxxxxxxxxxxxxxxxxxxxxx\nnumber of rows used (maxNum=", maxNum, ")\nmain text file =\t", nOut0, 
			"\nother text file =\t", nOut1, "\nall text files =\t", nOut) 
		
		
} // end of findHowManyUsedRowsOfFile2	


//---------------------------

func rewrite_word_lemma_dictionary() {	
	
	//----------------------------------------------------
	sort.Slice(newWordLemmaPair, func(i, j int) bool {
			if (newWordLemmaPair[i].lWordCod != newWordLemmaPair[j].lWordCod) {
				return newWordLemmaPair[i].lWordCod < newWordLemmaPair[j].lWordCod
			} else {
				if (newWordLemmaPair[i].lWord2 != newWordLemmaPair[j].lWord2) {
					return newWordLemmaPair[i].lWord2 < newWordLemmaPair[j].lWord2 
				} else {
					return newWordLemmaPair[i].lLemma < newWordLemmaPair[j].lLemma
				}
			}
		} )		 	
	//------------			
	outFile := FOLDER_OUTPUT +  string(os.PathSeparator) + FILE_outWordLemmaDict ;		
	
	lines:= make([]string, 0, 10+len(newWordLemmaPair) )

	lines = append(lines,  "__" + outFile + "\n" + "_word _lemma ")
	
 	for z:=0; z < len( newWordLemmaPair); z++ {
		//lines = append(lines,  newWordLemmaPair[z].lWord2 + "|" + newWordLemmaPair[z].lLemma) 
		lines = append(lines,  newWordLemmaPair[z].lWord2 + " " + newWordLemmaPair[z].lLemma) 
	}  	
	
    writeList( outFile, lines )
	
	
} // end of rewrite_word_lemma_dictionary

//--------------------------------

func rewrite_LemmaTranDict_file() {
						
	//outFile := FOLDER_IO_lastTRAN +  string(os.PathSeparator) + FILE_ outLemmaTranDict;
	
	outFile := FOLDER_IO_lastTRAN  +  string(os.PathSeparator) + FILE_last_updated_dict_words 

	pkey := ""; key := ""
	
	lines:= make([]string, 0, 10+len(dictLemmaTran) )
	lines = append(lines,  "__" + outFile + "\n" + "_lemma	_traduzione")
	
	for z:=0; z < len(dictLemmaTran); z++ {
		pkey=key
		
		//if dictLemmaTran[z].dL_lemma2 == dictLemmaTran[z].dL_tran { continue }
		
		key = dictLemmaTran[z].dL_lemma2 + "|"  + dictLemmaTran[z].dL_tran  		////cigna1_3
		if pkey == key { continue}
		
		lines = append(lines, key ) 
	}
	writeList( outFile, lines )
	//--------------------

	currentTime := time.Now()		
	outF1 		:= FOLDER_O_arc_TRAN_words +  string(os.PathSeparator) + "dictL"  		
	outFile2 := outF1 + currentTime.Format("20060102150405") + ".txt"
	
	writeList( outFile2, lines )
	
	
} // end of rewrite_LemmaTranDict_file

//---------------------------
func read_wordsToLearn() {

	//swWrite:=false; outLearn:= make([]string,0,3000) 
	
	bytesPerRow:= 10
    lineD := rowListFromFile( FOLDER_INPUT_OUTPUT, FILE_words_to_learn, "words to learn", " bind_go_passToJs_read_wordsToLearn", bytesPerRow)  
	if sw_stop {  // this file might be missing
		sw_stop=false		
		return
	}
	
	nread:=0
	for z:=0; z< len(lineD); z++ { 
		fields:= strings.Split( lineD[z] ,"|") 
		
		/***
		if len(fields) == 4 {
			field01 := strings.TrimSpace(fields[0])[0:1]
			if z < 10 { fmt.Println( "1 toLearn ", fields) }
			if ((field01 >="0") && (field01 <="9")) { // si tratta della vecchia versione con indice nella prima posizione  
				fields = fields[1:]
				swWrite = true
			}	
			if z < 10 { fmt.Println( "2 toLearn ", fields) 	}
		}
		
		if swWrite {
			outLearn = append( outLearn, ( fields[0] + "|" +  fields[1] + "|" + fields[2] ) )   
		}
		****/
		
		if len(fields) < 3 { continue}
		 
		r_word2          := strings.TrimSpace( fields[0] ) 	
		known_yes_ctr, _ := strconv.Atoi( strings.TrimSpace( fields[1] ) ) 
		known_no_ctr , _ := strconv.Atoi( strings.TrimSpace( fields[2] ) ) 
		
		wordCod:= newCode( r_word2)		
	
		/**
			uWordCod    string	
			uWord2      string	
			uIxUnW      int            // index of this word in the uniqueWordByFreq	
			uIxUnW_al   int            // index of this word in the uniqueWordByAlpha 	
			uTotRow     int 
			uTotExtrRow int
			uIxWordFreq int            // index of this word in the wordSliceFreq	
			uSwSelRowG  int
			uSwSelRowR   int  
			uKnow_yes_ctr int 
			uKnow_no_ctr  int 
		**/
		
		// ignoro l'indice dell'input (potrebbero esserci state delle variazioni nella freq. delle parole) e lo ricalcolo ( ottengo in realtà un range di indici che dovrebbero coincidere)
		ixF, ixT:= lookForWordInUniqueAlpha( wordCod)		
		if (ixT < 0) {
				fmt.Println(red("error in" + "wordToLearn "), r_word2 , " not found in wordUniqueAlpha" ); 	
				continue
		}		
		for ixA:= ixF; ixA <= ixT; ixA++ {
			xWordA :=  uniqueWordByAlpha[ixA]
			if xWordA.uWordCod != wordCod {
				continue
			}			
			uniqueWordByAlpha[ixA].uKnow_yes_ctr = known_yes_ctr
			uniqueWordByAlpha[ixA].uKnow_no_ctr  = known_no_ctr
			if (ixA != xWordA.uIxUnW_al) {  
				fmt.Println( red("error in" + "wordToLearn "), r_word2 , " z=",z," fields=", fields, " ixA=", ixA, " xWordA.uIxUnW_al=",xWordA.uIxUnW_al);  
				continue
			}
			ix1 := xWordA.uIxUnW
			uniqueWordByFreq[ix1].uKnow_yes_ctr = known_yes_ctr
			uniqueWordByFreq[ix1].uKnow_no_ctr  = known_no_ctr 		

			//if nread < 10 { fmt.Println( "word to learn ", uniqueWordByFreq[ix1] , " no_ctr=", uniqueWordByFreq[ix1].uKnow_no_ctr) }
			
			nread++	
		}
	}
	
	/**
	if swWrite {
			writeList("NUOVO_wordToLearn.txt", outLearn) 	
	}
	**/
	
	fmt.Printf("letti %d parole da imparare  dal file %s\n", nread,  FILE_words_to_learn)
	
	//go_exec_js_function( js_function, outS1 ); 	
	
 } // end of read_words_to_learn_file 


//====================================================
func bind_go_passToJs_write_WordsToLearn(js_function string)  {
						
	outFile := FOLDER_INPUT_OUTPUT  + string(os.PathSeparator) + FILE_words_to_learn ;

	lines:= make([]string, 0, len(uniqueWordByFreq) )
	
	numOut:=0
	for i:= 0; i < len(uniqueWordByFreq); i++ { 
		var xWordF     = uniqueWordByFreq[i]
		if ( (xWordF.uKnow_yes_ctr < 1 ) && (xWordF.uKnow_no_ctr < 1 ))  { continue }
		lines = append( lines,   fmt.Sprintf( "%s|%d|%d",  xWordF.uWord2, xWordF.uKnow_yes_ctr,  xWordF.uKnow_no_ctr) ) 
		numOut++	
	}		
	writeList( outFile, lines )
	//--------
	currentTime := time.Now()	
	outF1 		:= FOLDER_O_arc_TO_learn + string(os.PathSeparator)  
	outFile2    := outF1 + "wordsToLearn_"  + currentTime.Format("20060102150405") + ".txt"	
	writeList( outFile2, lines )	
	//-----------------	
	
	if (js_function == "") { return }
	outS1:= fmt.Sprintf("scritti %d words to learn nel file %s", numOut,  FILE_words_to_learn)
	
	go_exec_js_function( js_function, outS1 ); 	
	
} // end of rewrite_word_to_learn_file

//----------------------------------------
func console( str1 string) {
	go_exec_js_function( "js_go_console", str1 ) 	
}
//-----------------------------------------------
func read_lastValueSets2() {
	bytesPerRow:= 10
    lineD := rowListFromFile( FOLDER_INPUT_OUTPUT, FILE_last_mainpage_values2, "last run values", "read_lastValueSets2", bytesPerRow)  
	
	fmt.Println("read_lastValueSets2 lineD=", lineD)
	
	var dat="";
	if sw_stop == false { 
		for v1:=0; v1 < len(lineD); v1++ { 
			dat = dat + lineD[v1] + " ";  
		}  
	} else {
		sw_stop = false
	}		
	getMainPageLastVal2( dat + ",,,,,,,,,,,," )
	rwS:= ""
	if last_rG_firstIxRowOfGr < len(inputTextRowSlice) {
		rwS = inputTextRowSlice[ last_rG_firstIxRowOfGr ].rRow1  		
	}
	outS1:= fmt.Sprintf( "%d,%s,%d,%d,html,%d,%d,%d,ix,%d,%d,w,%d,%d,%s, :row=,%s", 
				last_rG_ixSelGrOption, 
				last_rG_group,
				last_rG_firstIxRowOfGr ,     
				last_rG_lastIxRowOfGr  ,      
				
				last_html_rowGroup_index_gr, 
				last_html_rowGroup_beginNum, 
				last_html_rowGroup_numRows,  
    
				last_ixRowBeg,     
				last_ixRowEnd,     
				
				last_word_fromWord, 	
				last_word_numWords,		
				
				last_sel_extrRow, 	
				rwS) 
	if outS1 == "" {
		fmt.Println( red("read_lastValueSets2"), " outS1 empty =" + outS1, " \n" + " sw_stop=", sw_stop, "  dat=" , dat)
		return
	}				
	go_exec_js_function( "js_go_valueFromLastRun", outS1 )	
	
	
} // end read_lastValueSets2()

//----------------------------

func getMainPageLastVal2(valueStr string ) {

	//fmt.Println("getMainPageLastVal2(valueStr=", valueStr)
	
	col1:= strings.Split( valueStr, ",") 
	for _, ele1:= range(col1) {
		el1       := strings.Split( ele1, "=")
		if len(el1) != 2 {continue} 
		name1 := strings.TrimSpace( el1[0] ) 
		var1  := strings.TrimSpace( el1[1] )		
		if ((name1 == "") || (var1  == "")) { continue } 
		//fmt.Println("  \tgetMainPageLastVal2  name1=", name1,  " \t var1=", var1)
		switch( name1 ) {
			case "rG_ixSelGrOption"  :  last_rG_ixSelGrOption , _ = strconv.Atoi( var1 )
			case "rG_group"          :  last_rG_group             = var1 
			
			case "rG_firstIxRowOfGr" :  last_rG_firstIxRowOfGr, _ = strconv.Atoi( var1 )     
			case "rG_lastIxRowOfGr"  :  last_rG_lastIxRowOfGr , _ = strconv.Atoi( var1 )    		
			
			case "html_rowGroup_index_gr" :  last_html_rowGroup_index_gr, _ = strconv.Atoi( var1 )
			case "html_rowGroup_beginNum" :  last_html_rowGroup_beginNum, _ = strconv.Atoi( var1 )
			case "html_rowGroup_numRows"  :  last_html_rowGroup_numRows , _ = strconv.Atoi( var1 )	
			
			case "ixRowBeg"               :  last_ixRowBeg              , _ = strconv.Atoi( var1 )
			case "ixRowEnd"               :  last_ixRowEnd              , _ = strconv.Atoi( var1 )
			
			case "w_fromWord"        :  last_word_fromWord    , _ = strconv.Atoi( var1 )	
			case "w_numWords"        :  last_word_numWords    , _ = strconv.Atoi( var1 )	
			
			case "sel_extrRow"       :  last_sel_extrRow          = var1  
		}	
	} 	
	
	fmt.Println( "\n" + green("last_mainpage_values2"),  
		    "\n\t", green("rG_ixSelGrOption"  ), last_rG_ixSelGrOption ,
			"\n\t", green("rG_group"          ), last_rG_group         ,
			
			"\n\t", green("rG_firstIxRowOfGr" ), last_rG_firstIxRowOfGr ,
			"\n\t", green("rG_lastIxRowOfGr"  ), last_rG_lastIxRowOfGr  ,
			
			"\n\t", green("html_rowGroup_index_gr" ), last_html_rowGroup_index_gr, 
			"\n\t", green("html_rowGroup_beginNum" ), last_html_rowGroup_beginNum, 
			"\n\t", green("html_rowGroup_numRows"  ), last_html_rowGroup_numRows , 				
			
			"\n\t", green("ixRowBeg"          ), last_ixRowBeg         ,         
			"\n\t", green("ixRowEnd"          ), last_ixRowEnd         ,
			
			"\n\t", green("w_fromWord"        ), last_word_fromWord ,	
			"\n\t", green("w_numWords"        ), last_word_numWords ,	
			
			"\n\t", green("sel_extrRow"       ), last_sel_extrRow  )
	
	
} // end of getMainPageLastVal2()
//-----------------------------------------------

func read_ParadigmaFile( path1 string, inpFile string) {
	bytesPerRow:= 10
    righe := rowListFromFile( path1, inpFile, "paradigma", "read_ParadigmaFile", bytesPerRow)  
	if sw_stop { return }
	
	/*
		ab  | A2 | abholen, holt ab, hat abgeholt |Wann kann ich die Sachen bei dir abholen<br>Wir müssen noch meinen Bruder abholen  | 
	     0  | 1  |                 2              |                                  3                                                |                                                                  |    |  |        
	*/ 
	
	fmt.Println("\nletti paradigma file ", inpFile  + " " , len(righe) , " righe")   
	
	lemma_para_list = make([]paraStruct, 0, len(righe)+4 )   
	var wP paraStruct
	//var pP paraStruct
	var pkeyL, keyL string
	sk:=0
	//--------------
	for z1:=0; z1 < len(righe); z1++ {		
		col := strings.Split(righe[z1], "|") 	
		if len(col) < 4 { continue }	
		wP.p_lemma   = strings.TrimSpace( col[0] )
		wP.p_level   = strings.TrimSpace( col[1] ) 
		wP.p_para    = strings.TrimSpace( col[2] ) 	
		
		keyL = wP.p_lemma + "." + wP.p_level + "." + wP.p_para 
		if keyL == pkeyL { sk++; continue }
		pkeyL = keyL

		level1 := " " +  wP.p_level + " "
		if strings.Index(list_level_str, level1) < 0 { list_level_str += level1 } 	
		
		if wP.p_para != "" {
			//    0 = x48, A = x65, Z =x90,  a = x97
			ch1 := wP.p_para[0:1]
			if ch1 < "a" { 
				if ch1 < "A" || ch1 > "Z" {
					wP.p_para = fseq + wP.p_para  // per la chiave di sort,  serve per spostare la riga alla fine  se il paradigma inizia con ( [ o altro 
				}	 
			}
		} 
		wP.p_example = strings.TrimSpace( col[3] ) 	
		
		//if wP.p_lemma == pP.p_lemma &&  wP.p_level == pP.p_level { continue } 
		
		//pP = wP ;  
		lemma_para_list = append(lemma_para_list , wP ) 	
		
	} // end of for_z1	
	//------------
	fmt.Println("    scartate ", sk, " righe doppie, ", len( lemma_para_list ), " righe caricate in lemma_para_list")
	
	list_level = strings.Split(strings.TrimSpace(list_level_str), " ") 
	
	//fmt.Println("XXX livelli: string=>" + list_level_str + "<== \nlivelli=", list_level) 
	
	only_level_numWords = make([]int, len(list_level), len(list_level) )
	perc_level          = make([]int, len(list_level), len(list_level) )
	
	sort_lemmaPara() 
	
	numFound:=0; 
	
	for z2, wP2 := range lemma_para_list {		
		//if z2 < 10 {fmt.Println("XXX PARADIGMA: ", z2, " ",  wP2, "   wP2.p_lemma=" +wP2.p_lemma  )   }
		
		ix1, ix2 := lookForLemma(wP2.p_lemma)
		
		//if z2 < 10 {fmt.Println("     ix1=", ix1,  " ix2=", ix2) }
		
		for z3:=ix1; z3 <=ix2; z3++ {
			LEM:= lemmaSlice[z3]
			
			//if z2 < 10 {fmt.Println("     z3=", z3, " LEM.leLemma=", LEM.leLemma ) }
			
			if LEM.leLemma == wP2.p_lemma {
				lemma_para_list[z2].p_ixLemma = z3   // indice di lemmaSlice  				
				LEM.leLevel   = wP2.p_level   
				LEM.lePara    = wP2.p_para  
				LEM.leExample = wP2.p_example 
				lemmaSlice[z3] = LEM	
				numFound++	
				//if numFound < 10 { fmt.Println(" lemma_para_list ", wP2 , " \t XXX lemmaSlice[",z3,"] = ", lemmaSlice[z3] )   	}
			}
		} 
	}
	fmt.Println("\t\t", numFound , " lemma_para_list FOUND in lemmaSlice     ",  (len( lemma_para_list ) - numFound), " not found")       
	
} // end of read_ParadigmaFile

//-----------------------------------------------

func sort_lemmaPara() {

	sort.Slice(lemma_para_list, func(i, j int) bool {
		if lemma_para_list[i].p_lemma != lemma_para_list[j].p_lemma {
			return lemma_para_list[i].p_lemma < lemma_para_list[j].p_lemma         // lemma ascending order (eg.   a before b ) 
		} else {
			if lemma_para_list[i].p_level != lemma_para_list[j].p_level {
				return lemma_para_list[i].p_level < lemma_para_list[j].p_level     // level ascending order (eg.   A1 before A2 ) 
			} else {
				return lemma_para_list[i].p_para < lemma_para_list[j].p_para       // level ascending order (eg.   a  before b ) 
			}
		}	
	} )	

} // end of sort_lemmaPara() 

//--------------------

func lookForAllParadigma( lemma3 string ) (int, int) {
		
	fromIxX, toIxX := lookForParadigma(lemma3 )
	if toIxX < 0 { return -1, -99999 }
	
	
	minIx:= -1; maxIx := -1 

	fromIx:= fromIxX
	
	// get the smaller index  with the right lemma  
	
	for k:= fromIx; k >= 0; k-- {
		if lemma_para_list[k].p_lemma == lemma3 {
			minIx=k
		} else { 
			if lemma_para_list[k].p_lemma < lemma3 { break }
		}
	}
	
	// get the maximum index with the right lemma
	maxIx = minIx
	for k:= fromIx+1; k < len( lemma_para_list); k++ {
		if lemma_para_list[k].p_lemma == lemma3 {
			maxIx = k; 
			if minIx < 0 { minIx = k}
		}  else { 
			if lemma_para_list[k].p_lemma > lemma3 { break }
		}
	}	
	if maxIx < 0 { maxIx = -9999; minIx = -1 } 
	return minIx, maxIx
	
} // end of lookForAllParadigma

//-------------------------------------
func lookForParadigma(lemmaToFind string) (int, int) {

	// find 2 indices of the 2 words nearest to the word to find 
	
	low   := 0
	high  := len(lemma_para_list) - 1	
	maxIx := high; 
	
	//----
	for low <= high{
		median := (low + high) / 2
		if lemma_para_list[median].p_lemma < lemmaToFind {  
			low = median + 1
		}else{
			high = median - 1
		}
	} 
	//---
	fromIx:= low; toIx := high; 
	if fromIx > toIx { fromIx = high; toIx = low;}
	if fromIx < 0 { fromIx=0} 
	if toIx  > maxIx { toIx = maxIx}
	
	return fromIx, toIx	
	
} // end of lookForParadigma

//----------------------
func writeList( fileName string, lines []string)  {
	// create file
    f, err := os.Create( fileName )
    if err != nil {
        fmt.Println( red("error")," in writeList file=", fileName,"\n\t" , err ) //  log.Fatal(err)
    }
    // remember to close the file
    defer f.Close()

    // create new buffer
    buffer := bufio.NewWriter(f)

    for _, line := range lines {
        _, err := buffer.WriteString(line + "\n")
        if err != nil {
           fmt.Println( red("error"), " in buffer.WriteString file=", fileName,"\n\t" , err ) //log.Fatal(err)
        }
    }
    // flush buffered data to the file
    if err := buffer.Flush(); err != nil {
        fmt.Println( red("error"), " in buffer.Flush()cls file=", fileName,"\n\t" , err ) //  log.Fatal(err)
    }
} 
//----------------------------------------
func stdCode(inpCode string ) string {	
	CoerInp:= strings.ReplaceAll( 
					strings.ReplaceAll( 
						strings.ReplaceAll( 
							strings.ReplaceAll(inpCode, "ae","ä"),  
							"oe","ö"), 
						"ue","ü"),
					"ß","ss")  	           // non tutte le ss sono 	ß, ma è vero il contrario				
					
	return CoerInp  
}
//----------------------------
func newCode( inpCode string ) string {	

	//  pronto soltanto per il tedesco 

	/*
	1) serve soprattutto per mettere le parole in sequenza alfabetica coerente 
		( es. per il tedesco es.  ä, ö, ü, ß vicini rispettivamente ad a, o, u, ss)   
	2) a volte ä, ö, ü, ß sono scritti come ae, oe, ue, ss, in questo caso li sostituiamo con a, o, u, ss
    3) a volte eu dovrebbe rimanere tale (es. Treue), non ho modo di distinguere per cui la sequenza è falsata ue è prima di ua o di uz 
		l'alternativa potrebbe essere tradurre ü con ue, ma questo porterebbe alla sequenza errata  ue nel posto ua, ue, uz invece di u ua uz
	----------
	questo codice di sequenza   è usato per word e lemma
	la scrittura alternativa a quella ufficiale (es. ue invece di ü) può essere trovata in un testo.
	Il lemma si trova in un file "ufficiale"  quindi improbabile che venga usata la scrittura alternativa.

	a) nel lemma il newCode mi serve soltanto per correggere la sequenza
	b) nel word (che si trova nel testo analizzato) il new code mi serve per confrontare 
		però è probabile che un testo sia scritto o in codice alternativo o in modo normale, improbabile in entrambi modi. 
	c) si potrebbe pensare ad un switch da impostare 
	d) cosa devo confrontare?
		lemma ( non mi serve tradurre eventuale codice alternativo, però newCode mi serve per la sequenza )
			confronto per collegarlo a word e per trovare la traduzione
		word  ( confronto per assegnare il lemma )       	
	
	*/
	SQinp1:= strings.ReplaceAll( 
						strings.ReplaceAll( 
							strings.ReplaceAll(inpCode, "ae","a"),  
							"oe","o"), 
						"ue","u") 		
	SQinp2:= strings.ReplaceAll( 
					strings.ReplaceAll( 
						strings.ReplaceAll( 
							strings.ReplaceAll(SQinp1, "ä","a"),  
							"ö","o"), 
						"ü","u"), 
					"ß","ss")   					
	
	return SQinp2 + "." +  stdCode(inpCode)  
	
}// end of newCode					
//--------------------------------	
//----------------------------
func OLDnewCode( inpCode string ) string {	
	/*
	1) serve soprattutto per mettere le parole in sequenza alfabetica coerente 
		( es. per il tedesco es.  ä, ö, ü, ß vicini rispettivamente ad a, o, u, ss)   
	2) a volte ä, ö, ü, ß sono scritti come ae, oe, ue, ss, in questo caso li sostituiamo con a, o, u, ss
    3) a volte eu dovrebbe rimanere tale (es. Treue), non ho modo di distinguere per cui la sequenza è falsata ue è prima di ua o di uz 
		l'alternativa potrebbe essere tradurre ü con ue, ma questo porterebbe alla sequenza errata  ue nel posto ua, ue, uz invece di u ua uz
	----------
	questo codice di sequenza   è usato per word e lemma
	la scrittura alternativa a quella ufficiale (es. ue invece di ü) può essere trovata in un testo.
	Il lemma si trova in un file "ufficiale"  quindi improbabile che venga usata la scrittura alternativa.

	a) nel lemma il newCode mi serve soltanto per correggere la sequenza
	b) nel word (che si trova nel testo analizzato) il new code mi serve per confrontare 
		però è probabile che un testo sia scritto o in codice alternativo o in modo normale, improbabile in entrambi modi. 
	c) si potrebbe pensare ad un switch da impostare 
	d) cosa devo confrontare?
		lemma ( non mi serve tradurre eventuale codice alternativo, però newCode mi serve per la sequenza )
			confronto per collegarlo a word e per trovare la traduzione
		word  ( confronto per assegnare il lemma )       	
	
	*/
	inp1:= strings.ReplaceAll( 
						strings.ReplaceAll( 
							strings.ReplaceAll(inpCode, "ae","a"),  
							"oe","o"), 
						"ue","u") 		
	inp2:= strings.ReplaceAll( 
					strings.ReplaceAll( 
						strings.ReplaceAll( 
							strings.ReplaceAll(inp1, "ä","a"),  
							"ö","o"), 
						"ü","u"), 
					"ß","ss")   
	return inp2 + "." + inpCode  // 21Nov2023
	
}// end of OLDnewCode					
//--------------------------------	