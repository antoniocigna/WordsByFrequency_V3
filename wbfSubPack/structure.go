package wbfSubPack


//---------------

type separPrefStruct struct {
	sLenPref  int 
	sPrefix   string 
	sPrefTran string 	
}
//---
type lemmaStruct struct {
	leLemma    string        // ls_lemma_einstellen einstellen
	
	leNumWords int 
	leFromIxLW  int 
	leToIxLW    int  
	leTran      string 
	leLevel     string  
	lePara      string  
	leExample   string  
	ls_lemma_ix_stellen  int	
	ls_lemma_stellen     string
	ls_pref_ein          string
	ls_pref_tran         string 
	ls_lemma_einStellenList []int 
} 
//-------------------------------

type wordLemmaPairStruct struct {
	lWordCod string 
	lWord2   string 
	lLemma   string
	lIxLemma int
	lIx_einStellenList []int 
} 
//---

//---------------
type lemmaWordStruct struct {
	lw_lemmaCod string 	
	lw_lemma2   string 	
	lw_prefix   string
	lw_word     string 
	lw_ixLemma    int
	lw_ixWordUnFr int
	lw_origLemma string
}

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
type wDictStruct struct {
	dWord  		string 
	dIxWuFreq 	int 
	dLemmaL     []string 
	dTranL      []string 	
} 

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
type paraStruct struct {
	p_lemma    string	
	p_ixLemma  int 
	p_level    string
	p_para     string
	p_example  string 
}	
//--------------------------------
