package wbfSubPack


//---------------

type separPrefStruct struct {
	sLenPref  int 
	sPrefix   string 
	sPrefTran string 	
}
//---
type lemmaStruct struct {
	leLemma    string    
	leLemmaOr  string    
	leNumWords int 
	leFromIxLW  int             // limite inferiore range indici a wordLemmaPair (in seq. di lemma)   wordLemmaPair_lemmaWordSeq[] 
	leToIxLW    int             // limite superiore range indici a wordLemmaPair (in seq. di lemma)   wordLemmaPair_lemmaWordSeq[]   
	leUnWord_al_IxList []int    // indice delle parole unique che puntano a questo lemma        
	leTran      string 
	lePara      string  
	leExample   string  
	leNumPara   int	
} 
//-------------------------------

type wordLemmaPairStruct struct {
	lWord2   	 string 
	lLemma   	 string
	lWord0   	 string 
	lLemmaOr   	 string
	lIxLemma  	 int
	lIxUnWord_al int
} 
//---

type rowStruct struct {
	rIdRow       string
	rIxGroup     int      // indice del gruppo 
	rIxBaseGroup int      // posizione del row nel gruppo ( si inzia dal num.1 )  
    rRow1        string
	rNfile1      int  
	rSwExtra     bool 
	rListIxUnF   []int     // for each word in the row, index of the word in the uniqueWordByFreq  
	rListFreq    []int     // for each word in the row, how many times the word is used in all the text  	
    rNumWords    int       // number of words in the row 	
	rWordFreqAvg int       // average of the frequency of use of the words in this row   
	rPriority    int       // lesser the number, the first to be learned (atually this is the index of row in the rowPriorityList    
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

type rowPriorStruct struct {   // priority of rows in the text 	
		rP_numWords 	int    // number of words in row 
		rP_wordFreqAvg 	int	   // average of frequence of the words in the row		
		rP_index 		int    // index of row in inputTextRowSlice
		rP_ixGroup      int    // group number                          
}

const SEL_EXTR_ROW    = 1; 
const SEL_NO_EXTR_ROW = 2; 

//--
type wordStruct struct {       // a word is repeated several time one for each row containing it 
    wWord2    string           // parola con ortografia standardizzata ( umlaut scompare aggiunta vocale es. ue  , eszet diventa ss) 
	wWord0    string           // parola originale con umlaut e eszet se ci sono
	wIxThisWord_al int 
	wIxUniq_al   int               // index of uniqueWordByFreq 	
	wIxUniq_fr   int               // index of uniqueWordByFreq 	
	wNfile    int 
	wSwSelRowG int
	wSwSelRowR int
    wIxRow    int       
	wIxPosRow int 
	wListPref string 
	wTotRow   int              // number of rows 
	wTotExtrRow int            // number of extracted rows 
	wTotMinRow int
	wTotWrdRow int 	
	wIxLemmaList []int
}
//--
type wordUnAlphaStruct struct {    // uniqueWordByAlpha
    uWord2      string		
	uWord0      string	
	uIxUnW_al   int            // index of this word in the uniqueWordByAlpha 	
	uIxUnW_fr   int            // index of this word in the uniqueWordByFreq	
	uTotRow     int 
	uTotExtrRow int
	uIxFromWord_al int          // index of this word in the wordSliceAlpha (first occurrence, last = uIxFromWordAl + uTotRow-1	
    //uIxWordFreq int            // index of this word in the wordSliceFreq	
	uSwSelRowG  int
	uSwSelRowR   int  
	uLearnedYN   string         // y n ( ie.yes,I learned / not yet  
	//uKnow_yes_ctr int 
	//uKnow_no_ctr  int         // a value > 0  means that this is a word that I don't know, ie. it's to be learned   
	uIxLemmaL  []int  
	uLemmaL    []string       // list of lemma 	
	//uPara      []string  
	//uExample   []string  
}	

//--
type wordUnFreqStruct struct {    // uniqueWordByFreq
    fuWord2      string		
	fuIxUnW_al   int            // index of this word in the uniqueWordByAlpha 	
	fuIxUnW_fr   int            // index of this word in the uniqueWordByFreq	
	fuTotRow     int 
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
	dL_lemma      string 
	dL_numDict    int	
	dL_tran       string    
} 
//---------------
type paraStruct struct {
	p_lemma    string	
	p_ixLemma  int 
	p_para     string
	p_example  string 
}	
//--------------------------------
