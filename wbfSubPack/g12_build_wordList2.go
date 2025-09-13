package wbfSubPack

	import (
		"fmt"
		"strings"
		//"strconv"
		"sort"
		//"regexp"
		//"encoding/hex"
	)

//----------------------------------------
func g12_checkTheWord( word0 string ) string {
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
	return wor; 
}
//-------------------------------------------

func g12_add_totRow_and_indexLemmaPair() {
	
	g12_add_totRow_and_indexLemmaPairUNO_NUOVO()
	//g12_add_totRow_and_indexLemmaPairUNO()
	g12_add_totRow_and_indexLemmaPairDUE()

} // end of add_totRow_and_indexLemmaPair(

//------------------------------------------------
var errori=0
var numPrefPresi=0;
var numPrefTestati=0
//-----------------------------

//-----------------------------
func g12_add_totRow_and_indexLemmaPairUNO_NUOVO() {	
	fmt.Println("   func ", green("add_totRow_and_indexLemmaPair") )
	/*
	each element of wordSliceAlpha contains a word (the same word may be in several rows) 
	the number of repetition of a word (totRow) is put in its element  ( later will be put in each row that contain it) 
		eg.  one 3, one 3, one 3, two 4, two 4, two 4, two 4	
	*/
	/**
	totR  := 0	
	
	
	pre_wSwSelRow := SEL_NO_EXTR_ROW 
	//----------------
	tot_extrRow:=0
	lastIx:=0
	**/
	// alla fine dello slice c'è la parola LAST_WORD  che rende non necessaria la gestione di fine file  
	
	//fmt.Println(" len(listAllLemmaFromFile) = ", len(listAllLemmaFromFile) )
	fmt.Println(" len(lemmaSlice) = ", len(lemmaSlice) )
	maxNumWApp:= len(wordSliceAlpha)
	//wordSliceAlphaToApp = make([]wordStruct, 0, maxNumWApp)
	
	wordAlphaPlusPref   = make([]wordStruct, 0, maxNumWApp)
	
	
	g12_updateWordAlpha_with_ix_lemma()     //  crea nuove word  con lemma = prefix + lemma originale    wordAlphaPlusPref[]
	
	
} // end of g12_add_totRow_and_indexLemmaPairUNO_NUOVO


//-----------------------------------------------------------------
func printLemmaList( ixList []int ) string {
	str:=""
	for _,ixL:= range ixList {
		str += " " + lemmaSlice[ixL].leLemma
	}
	return str
}
//-----------------------------------------

func OLDg12_manage_one_word_of_manyRowsUNO(ix1 int, ix2 int, preW string, pre_wSwSelRow int, tot_extrRow int, totR int, lastIx int) {
	
	
	var wordSliceAlphaToApp = make([]wordStruct, 0, 0)   // serve proseguire la compilazione senza errori, questa funzione deve essere eliminata 
	
	swNoLemma:= false 
	/*
		var NO_LEMMA_WORD  = "_"
		var NO_LEMMA_LEMMA = "_"
		var NO_LEMMA_INDEX = 0
	*/
	
	ixLemmaPairFoundList := lookForAllLemmas( preW, lastIx)   // indici lemma per questo gruppo di word 
	if len(ixLemmaPairFoundList) < 1 {	
		swNoLemma = true 
		ixLemmaPairFoundList = append(ixLemmaPairFoundList, NO_LEMMA_INDEX)
	} else {
		if ixLemmaPairFoundList[0] < 0 { 
			swNoLemma = true 
			ixLemmaPairFoundList[0] = NO_LEMMA_INDEX
		}
	}	
		
	//fmt.Println("add_... lemma... word=", preW, " 	ixLemmaPairFoundList=", ixLemmaPairFoundList,   " for ix=", ix1, " to ix2=", ix2) 
	lemmaIndex:= make([]int,0,100)
	ixAF:=0
	
	//-------------------
	for i2 := ix1; i2 < ix2;i2++ {
		 wS22:= wordSliceAlpha[i2]
		 var wS33 wordStruct
		 
		 //wS22.wIxLemmaPair = make([]int, len( ixLemmaPairFoundList ) )
		 //copy( wS22.wIxLemmaPair, ixLemmaPairFoundList)
		 
		 
		 //fmt.Println("add_... lemma...  wordSliceAlpha[",i2,"].wIxLemmaPair=", wS22.wIxLemmaPair
		 //------------------------------
		 var le_Lemma string 
		 //if len(wS22.wListPref) > 0 {
		for _,ixLe:= range ixLemmaPairFoundList {  
			 lePe:= wordLemmaPair[ixLe]
			 if swNoLemma {
				 if ixLe != NO_LEMMA_INDEX {
					 fmt.Println(red("ERRORE g12_manage_one_word_of_manyRowsUNO "), " ixLe=", ixLe , " not equal to NO_LEMMA_INDEX=", NO_LEMMA_INDEX)
					 continue
				 }
				 le_Lemma = lePe.lLemma 
			 } else {
				 if wS22.wWord2 != lePe.lWord2 { 
					continue } // error   
				 le_Lemma = lePe.lLemma
			 }
			 
			 //ixAF = lePe.lIxLemma
			 ixAF = lePe.lIxLemma; //  binarySearch_string(listAllLemmaFromFile, le_Lemma) 
			 if ixAF >=0 { 
				lemmaIndex = append(lemmaIndex, ixAF)
				wS22.wIxLemmaList = append(wS22.wIxLemmaList, ixAF) 
			 }		
			 
			 if lePe.lIxLemma != ixAF {  
					//fmt.Println(red(" lePe.lIxLemma="), lePe.lIxLemma, " diversa da " ixAF=", ixAF); 
					errori++ 
				}
			 
			 if ixLe == NO_LEMMA_INDEX {continue}
			 //fmt.Println("               wRord2=", wS22.wWord2, " lemma=", le_Lemma)	
			 pref3 := strings.Fields(wS22.wListPref) 
			 if len(pref3) < 1 {continue}
			 
			 for _, onePref:= range pref3 {
				 numPrefTestati++
				 newLemma:= onePref + le_Lemma 
				 //fmt.Println("              1 word2=", wS22.wWord2, " lemma=", le_Lemma, " newLemma =", newLemma)	
				 ixAF = binarySearch_string(listAllLemmaFromFile, newLemma) 
				 if ixAF < 0 { continue}
				 numPrefPresi++
				 lemmaIndex = append(lemmaIndex, ixAF)
				 wS33 = wS22
				 wS33.wWord2   =  wS22.wWord2   + " ... " + onePref 
				 wS33.wIxLemmaList = []int{} 
				 wS33.wIxLemmaList = append( wS33.wIxLemmaList, ixAF)
				 wS33.wSwSelRowG    = pre_wSwSelRow; 	// se esiste almeno un richiamo a una riga estratta ( wSwSelRowR)allora questo segnale è ripetuto come wSwSelRowG
				 wS33.wTotExtrRow   = tot_extrRow 
				 wS33.wTotRow       = totR;   // se una parola è ripetuta 3 volte, ad ogni parola è associato 3  		
				 wordSliceAlphaToApp = append(wordSliceAlphaToApp, wS33)
				
				 //fmt.Println("              word2=", wS22.wWord2 + " ... " + onePref, " lemma=", le_Lemma, " newLemma =", newLemma)	
				 //PRElemmaSlice = append(PRElemmaSlice, leV )
			 }
		 }
		 // }
		 //---------------------------	
		 wS22.wSwSelRowG    = pre_wSwSelRow; 	// se esiste almeno un richiamo a una riga estratta ( wSwSelRowR)allora questo segnale è ripetuto come wSwSelRowG
		 wS22.wTotExtrRow   = tot_extrRow 
		 wS22.wTotRow       = totR;   // se una parola è ripetuta 3 volte, ad ogni parola è associato 3  		
		 wordSliceAlpha[i2] = wS22
		 //fmt.Println("                       wS22:   ",  wS22.wWord2 , " ", wS22.wIxLemmaList )
	}
    //-------------
} // end of OLDg12_manage_one_word_of_manyRowsUNO
//------------------------	


//------------------------------

func g12_add_totRow_and_indexLemmaPairDUE() {	

	fmt.Println("   func ", green("add_totRow_and_indexLemmaPairDUE") )
	
	/**
	fmt.Println("    nuove word ", len(wordSliceAlphaToApp)  )
	for _, ww:= range wordSliceAlphaToApp{
		fmt.Println("      aggiunto ", ww)
	} 
	**/
	
	//wordSliceAlpha = append(wordSliceAlpha, wordSliceAlphaToApp...)	
	//wordSliceAlphaToApp = make([]wordStruct, 0, 0)  // rilascio lo spazio 
	
	wordSliceAlpha = append(wordSliceAlpha, wordAlphaPlusPref...)
	
	wordAlphaPlusPref = make([]wordStruct, 0, 0)  // rilascio lo spazio  
	
	//-----------------------------------------------------------	
	sort.Slice(wordSliceAlpha, func(i, j int) bool {
		if wordSliceAlpha[i].wWord2 != wordSliceAlpha[j].wWord2 {
			return wordSliceAlpha[i].wWord2 < wordSliceAlpha[j].wWord2  
		} else {
			if wordSliceAlpha[i].wWord0 != wordSliceAlpha[j].wWord0 {
				return wordSliceAlpha[i].wWord0 < wordSliceAlpha[j].wWord0  
			} else {			
				return wordSliceAlpha[i].wNfile < wordSliceAlpha[j].wNfile          // nFile ascending order (eg.   0 before 1 ) 
			}
		}		
	})
	//------------------------------	
	//=====================================
	totR  := 0		
	pre_wSwSelRow := SEL_NO_EXTR_ROW 
	//----------------
	tot_extrRow:=0
	lastIx:=0
	// alla fine dello slice c'è la parola LAST_WORD  che rende non necessaria la gestione di fine file  	
	fmt.Println(" len(listAllLemmaFromFile) = ", len(listAllLemmaFromFile) )
	
	ix1   := 0
	if len(wordSliceAlpha) < 1 {return }
	preW  := wordSliceAlpha[0].wWord2;	
	//--------------------------------------
	for ix2, wS1 := range wordSliceAlpha {
		
		if (wS1.wWord2 != preW) {			
			g12_manage_one_word_of_manyRowsDUE(ix1, ix2, preW, pre_wSwSelRow , tot_extrRow, totR, lastIx)			
			pre_wSwSelRow = SEL_NO_EXTR_ROW 
			totR = 0
			tot_extrRow = 0
			ix1  = ix2; 
			preW = wS1.wWord2; 
		} 
		
		if wS1.wSwSelRowR == SEL_EXTR_ROW {   // se almeno uno è "estratto", tutti lo sono 
			pre_wSwSelRow = SEL_EXTR_ROW 
			tot_extrRow++	
			//if (wS1.wWord2 == "schrift") { fmt.Println("ANTO addTotRowToWord  ",  wS1 , " SEL_EXTR_ROW=", SEL_EXTR_ROW, "  tot_extrRow=", tot_extrRow) }  //
		} 
		totR++;     	
	}	
	//------	
	
} // end of add_totRow_and_indexLemmaPairDUE
//------------------------------------


//-----------------------------------------------------------------

func g12_manage_one_word_of_manyRowsDUE(ix1 int, ix2 int, preW string, pre_wSwSelRow int, tot_extrRow int, totR int, lastIx int) {
	
	//-------------------
	for i2 := ix1; i2 < ix2;i2++ {
		 wS22:= wordSliceAlpha[i2]
		 sort.Ints( wS22.wIxLemmaList )
		 wS22.wSwSelRowG    = pre_wSwSelRow; 	// se esiste almeno un richiamo a una riga estratta ( wSwSelRowR)allora questo segnale è ripetuto come wSwSelRowG
		 wS22.wTotExtrRow   = tot_extrRow 
		 wS22.wTotRow       = totR;   // se una parola è ripetuta 3 volte, ad ogni parola è associato 3  		
		 wordSliceAlpha[i2] = wS22		 
	}
    //-------------
	g12_buildUnique( ix1, wordSliceAlpha[ix1] )
	unq1:= len(uniqueWordByAlpha) - 1  	
	for i2 := ix1; i2 < ix2;i2++ {
		wordSliceAlpha[i2].wIxUniq_al = unq1
		wordSliceAlpha[i2].wIxThisWord_al = i2
	}
	
} // end of manage_one_word_of_manyRowsDUE

//--------------------------------------------
func g12_buildUnique( n1 int, wS1 wordStruct) {
	//lemmaIndex:= wS1.wIxLemmaList
	var xWordF wordUnAlphaStruct;   	
	if wS1.wTotRow >= LAST_WORD_FREQ {
		wS1.wTotRow = 0
	}
	xWordF.uWord2    	= wS1.wWord2;
	xWordF.uWord0    	= wS1.wWord0;
	xWordF.uTotRow   	= wS1.wTotRow
	xWordF.uTotExtrRow 	= wS1.wTotExtrRow  
	xWordF.uSwSelRowR 	= wS1.wSwSelRowR 
	xWordF.uSwSelRowG 	= wS1.wSwSelRowG
	xWordF.uIxFromWord_al = n1	
	xWordF.uLearnedYN  	= LEARNED_NOT 
	
	//if xWordF.uWord2 == "ich" { fmt.Println( green("1build_uniqueWord_byFreqAlpha"),  
	//		" freq:  learned= ", xWordF.uLearnedYN  ) } 
			
	//xWordF.wTran = "" 
	xWordF.uIxUnW_al  = len(uniqueWordByAlpha)  
	
	
	g13_addLemmaTranParadigmaToUniqueWord( xWordF , wS1.wIxLemmaList)  // append to uniqueWordByAlpha  
	
	//fmt.Println("STAT. ", n1, " ", xWordF.word, " numWordUn=", numWordUn,  " numWordRi=", numWordRi, " percIx=", percIx, " ", sS.uniquePerc,  " sS.totPerc=" ,  sS.totPerc); 
		 			
 } // end for buildUnique

//------------------------------------
/*
//--
type wordStruct struct {       // a word is repeated several time one for each row containing it  
    wWord2    string
	wWord0    string
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
//-------------------------------

type wordLemmaPairStruct struct {
	lWord2   	 string 
	lLemma   	 string
	lIxLemma  	 int
	lIxUnWord_al int
} 
//---
**/

type wordCandStruct struct {
	wc_word2     string
	wc_lemma     string 
	wc_ixLemma     []int 
	wc_ixWordAlpha int 
}	

var wordSliceAlphaCand = make([]wordCandStruct, 0, 0)  

//------------------------------------------------
func g12_updateWordAlpha_with_ix_lemma() {  

	var xWordF wordUnAlphaStruct; 
	
	xWordF.uWord2    	= ""
	xWordF.uTotRow   	= 0	
	//xWordF.uIxLemmaL  = []int 
	
	tempWordUniq := make([]wordUnAlphaStruct,0, len(wordSliceAlpha) )
	
	ix1:=0
	preW2:=""; preW0:=""
	// non is seve gestire l'ultimo elem = x'FF
	for ix2, wS1 := range wordSliceAlpha {		
		if (wS1.wWord2 != preW2) {	
			xWordF.uWord2   = preW2; 
			xWordF.uWord0   = preW0; 
			xWordF.uIxFromWord_al = ix1
			xWordF.uTotRow    = ix2-ix1	
			tempWordUniq = append(tempWordUniq, xWordF)						
			preW2 = wS1.wWord2 	
			preW0 = wS1.wWord0 	
			ix1 = ix2; 	
		} 
	}	
	
	g12_updateTempWordUniq( tempWordUniq, wordLemmaPair )
	
	wordSliceAlphaCand = make([]wordCandStruct, 0, len(wordSliceAlpha)  ) 
	
	
	var WC1 wordCandStruct
	//nn:=0
	fromIx:=0; toIx:=0
	for _, tU:= range tempWordUniq{
		//sw2:= (tU.uWord2 == "geht") 
		
		fromIx = tU.uIxFromWord_al;  
		toIx= fromIx + tU.uTotRow
		listIxLem := tU.uIxLemmaL
		
		//if sw2 {fmt.Println( magenta("g12 update "), tU.uWord2, " tU.uIxLemmaL=", tU.uIxLemmaL, 
		//	" wordSliceASlpha fromIx=", fromIx, " to Ix=", toIx ," tempWordUniq[", g,"]= ", tU)  }
		
		for x:= fromIx; x < toIx; x++ {
			wordSliceAlpha[x].wIxLemmaList = make([]int, len(listIxLem), len(listIxLem) )
			nc:= copy(wordSliceAlpha[x].wIxLemmaList, listIxLem)			
			if len(listIxLem) != len(wordSliceAlpha[x].wIxLemmaList) {
				fmt.Println( red("ERRORE g12 update "), "copy ", nc, " elem.", "wordSliceAlpha[x].wIxLemmaList=", wordSliceAlpha[x].wIxLemmaList,
					" listIxLem=", listIxLem, " tU.uIxLemmaL=", tU.uIxLemmaL)
				return	
			} else {
				//nn++
				/**
				if sw2 {
					fmt.Println( "\twordSliceAlpha[", x, "].wIxLemmaList=", wordSliceAlpha[x].wIxLemmaList,
					" listIxLem=[", listIxLem,"] " , wordSliceAlpha[x].wWord2)
					  for _, pp:= range wordSliceAlpha[x].wIxLemmaList { 	
						fmt.Println("\t\t lemma=", 	lemmaSlice[ pp ] )
					  }
				}
				**/
			}	
			//----
			wA:=wordSliceAlpha[x]
			
			//sw4:= ((wA.wWord2 == "geht") || (wA.wWord2 == "gingen") || (wA.wWord2 == "ging") )
			
			pref3 := strings.Fields(wA.wListPref) 
			
			//if sw4 { fmt.Println( green("g12_updateWordAlpha NUOVO "), wA.wWord2, " pref3: ", pref3 ) }
			
			if len(pref3) < 1 {continue}			 
			for _, onePref:= range pref3 {		
				WC1.wc_word2 = wA.wWord2   + " ... " + onePref 					
				WC1.wc_ixLemma = []int{}	
				WC1.wc_ixWordAlpha = x 				
				for _, ixLe3:= range listIxLem {
					WC1.wc_lemma = onePref + lemmaSlice[ixLe3].leLemma            // la sua esistenza è da testare  
					wordSliceAlphaCand = append(wordSliceAlphaCand, WC1) 					
					//if sw4 { fmt.Println( green("g12_updateWordAlpha NUOVO candidato "), WC1.wc_word2, " lemma: ", WC1.wc_lemma) }
				}			
			}

		} // end for x 	
	}  // end for g 
	//------------------------------------
	sort.Slice(wordSliceAlphaCand, func(i, j int) bool {	
		if wordSliceAlphaCand[i].wc_lemma != wordSliceAlphaCand[j].wc_lemma {
			return wordSliceAlphaCand[i].wc_lemma < wordSliceAlphaCand[j].wc_lemma 
		} else {
			if wordSliceAlphaCand[i].wc_word2 != wordSliceAlphaCand[j].wc_word2 {
				return wordSliceAlphaCand[i].wc_word2 < wordSliceAlphaCand[j].wc_word2 
			} else {
				return wordSliceAlphaCand[i].wc_ixWordAlpha < wordSliceAlphaCand[j].wc_ixWordAlpha 
			}
		}
	})
	//---------------
	/**
	for _, wC1:= range wordSliceAlphaCand {
		fmt.Println("   prima di vfylemma candiati ", wC1)
	}
	**/
	
	//-----------------------------------------	
	trovati := g12_update_wordCand_vfyLemma(wordSliceAlphaCand, lemmaSlice) 
	
	if trovati > 0 {
			trovati+=10
		wordAlphaPlusPref = make([]wordStruct, 0, trovati )
		for _, wC1:= range  wordSliceAlphaCand {
			if len(wC1.wc_ixLemma) < 1 {continue}  
			ixWordAlf := wC1.wc_ixWordAlpha 
			wS33 := wordSliceAlpha[ixWordAlf]
			wS33.wWord2   =  wC1.wc_word2 
			wS33.wIxLemmaList = []int{}
			wS33.wIxLemmaList = append(wS33.wIxLemmaList, wC1.wc_ixLemma...)
			wordAlphaPlusPref = append(wordAlphaPlusPref, wS33)
			//fmt.Println("    aggiunto ", wS33.wWord2, " \tlemma=", wS33.wIxLemmaList)
		}
		
	}
	fmt.Println( magenta("word aggiunte in  wordAlphaPlusPref "), len(wordAlphaPlusPref ) , " nuove parole")
	/**
	for _, wS33:= range wordAlphaPlusPref {
		fmt.Println("    aggiunto (NUOVO metodo ): ",  wS33.wWord2 , " lemma=", wS33.wIxLemmaList,
			printLemmaList(wS33.wIxLemmaList), " ", wS33	 )
	}
	**/
	
	wordSliceAlphaCand = make([]wordCandStruct, 0, 0 )  // serve per liberare lo spazio  
	
	
}// end of g12_updateWordAlpha_with_ix_lemma	
//-----------------------------------
func g12_updateTempWordUniq( tempWordUniq []wordUnAlphaStruct, wordLemmaPair []wordLemmaPairStruct) { 	
	/*
	aggiorna tempWordUniq con i dati di wordLemmaPair
	tutte le liste devono essere già in sequenza di codice (alfanumerico)
	*/
	//---------------------------------------
	var maxNumErr = 10;	
	
	numOutSeq := 0;
	len1 := len(tempWordUniq) 
	len2 := len(wordLemmaPair   )
	loopMax := len1 + len2 
	
	j1:=0; j2:=0
	min1:= ""; 
	type1:=0; 	
	//----------------
	t0 :=-1;
	preMin := ""
	index_ix1 := -1
	var cod1 , cod2  string;
	var pcod1, pcod2 string;
	numIgn2 :=0;
	//-----------------------
	
	for t0=0; t0 < loopMax; t0++ {
		if (j1 < len1) {cod1 = "1" + tempWordUniq[j1].uWord2        } else {cod1 = "9" }
		if (j2 < len2) {cod2 = "1" + wordLemmaPair[j2].lWord2   } else {cod2 = "9" }
		if (cod1 <= cod2) {
			min1 = cod1; type1=1; 
			if (cod1 < pcod1) {	_ = g14_outSeqErr("g12_updateTempWordUniq 1 ", type1, pcod1, cod1, numOutSeq, maxNumErr); return ; 	}
			pcod1 = cod1
		}  else {
			min1 = cod2; type1=2;
			if (cod2 < pcod2) {	_ = g14_outSeqErr("g12_updateTempWordUniq 2 ",  type1, pcod2, cod2, numOutSeq, maxNumErr); return ; 	}
			pcod2 = cod2
		}
		
		if (min1[0:1] == "9") {break}
		
		if (min1 > preMin) {
			if (preMin != "") { 
				//seqCodeChange(index_ix1, wordSliceAlpha)
			}			
			index_ix1 = -1
			preMin    = min1 
		} else {
			if (min1 < preMin) { 
				numOutSeq = g14_outSeqErr("g12_updateTempWordUniq 4 ",  type1, preMin, min1, numOutSeq, maxNumErr)
				if numOutSeq < 0 { return }
				continue; 					
			}
		}		
		if (type1 == 1) {
			index_ix1 = j1  
			j1++;
		} else {
			if (index_ix1 < 0) { 
				//fmt.Println("               coppia word lemma senza una word in testo, ignorato " , wordLemmaPair[j2] )
				//numIgn2++
			} else {
				tempWordUniq[index_ix1].uIxLemmaL = append( tempWordUniq[index_ix1].uIxLemmaL, wordLemmaPair[j2].lIxLemma  )
				/**
				if ((min1 == "1geht" ) || ( min1 ==  "1gehoren")) {
					vv:=  wordLemmaPair[j2].lIxLemma 
					fmt.Println(" wordLemmaPair[", j2, "] =", wordLemmaPair[j2], " tempWordUniq[",index_ix1,"]=", tempWordUniq[index_ix1], " lemmaSlice[", vv, "]=", lemmaSlice[ vv ])					
				} 
				**/
			}
			j2++;
			 
		}		
		
	} // end for t0
	//-----------------
	
	//if (preMin != "") {seqCodeChange(index_ix1, wordSliceAlpha)}
	g12_add_unknowLemma(tempWordUniq) 
	
	fmt.Println("  in update tempWordUniq ([]wordAlphaStruct), trovati ",    
		len(wordSliceAlpha), " righe in wordSliceAlpha, \n\t",  len(wordLemmaPair), " word_lemma pair letti, di cui ", (len(wordLemmaPair)-numIgn2), " usati" )


} // end of g14_update_tempWordUniq
//-----------------
func g12_add_unknowLemma(tempWordUniq []wordUnAlphaStruct) {
	//----------------
	/**
	var NO_LEMMA_WORD  = "aaalemmanotfound"
	var NO_LEMMA_LEMMA = "aaalemmanotfound"
	var NO_LEMMA_INDEX = 0
	**/
	noLem:=0
	for j3, wU:= range tempWordUniq{
		if len(wU.uIxLemmaL) < 1 { 
			tempWordUniq[j3].uIxLemmaL = append( tempWordUniq[j3].uIxLemmaL, NO_LEMMA_INDEX  )
			noLem++
			//fmt.Println("tempWordUniq[j3]=",tempWordUniq[j3], " lemma=", lemmaSlice[NO_LEMMA_INDEX] )  
		}  
	}
	if noLem > 0 {
		fmt.Println("  in update tempWordUniq ", noLem , " parole non trovate assegnate al lemma ", NO_LEMMA_LEMMA, "( con indice ",NO_LEMMA_INDEX,")" )	
	}
	
} // end of g12_add_unknowLemma
//-------------------------------------------
func g12_update_wordCand_vfyLemma( wordSliceAlphaCand []wordCandStruct, lemmaSlice []lemmaStruct) int {
	/*
		type wordCandStruct struct 
			wc_word2     string
			wc_lemma     string 
			wc_ixLemma    [] int 
			wc_ixWordAlpha int 
			//---
		type lemmaStruct struct {
			leLemma    string    
			leNumWords int 
			leFromIxLW  int             // limite inferiore range indici a wordLemmaPair (in seq. di lemma)   wordLemmaPair_lemmaWordSeq[] 
			leToIxLW    int             // limite superiore range indici a wordLemmaPair (in seq. di lemma)   wordLemmaPair_lemmaWordSeq[]   
			leUnWord_al_IxList []int    // indice delle parole unique che puntano a questo lemma        
			leTran      string 
			lePara      string  
			leExample   string  
			leNumPara   int	
		} 
	*/
	/*
	aggiorna tempWordUniq con i dati di wordLemmaPair
	tutte le liste devono essere già in sequenza di codice (alfanumerico)
	*/
	//---------------------------------------
	var maxNumErr = 10;	
	
	numOutSeq := 0;
	len1 := len(wordSliceAlphaCand ) 
	len2 := len(lemmaSlice         )
	loopMax := len1 + len2 
	
	j1:=0; j2:=0
	min1:= ""; 
	type1:=0; 	
	//----------------
	t0 :=-1;
	preMin := ""
	index_ix1 := -1
	var cod1 , cod2  string;
	var pcod1, pcod2 string;
	firstGrIx1:=0
	trovati:=0
	//-----------------------
	
	for t0=0; t0 < loopMax; t0++ {
		if (j1 < len1) {cod1 = "1" + wordSliceAlphaCand[j1].wc_lemma } else {cod1 = "9" }
		if (j2 < len2) {cod2 = "1" + lemmaSlice[j2].leLemma          } else {cod2 = "9" }
		if (cod1 <= cod2) {
			min1 = cod1; type1=1; 
			if (cod1 < pcod1) {	_ = g14_outSeqErr("g12_update_wordCand_vfyLemma 1 ", type1, pcod1, cod1, numOutSeq, maxNumErr); return -1; 	}
			pcod1 = cod1
		}  else {
			min1 = cod2; type1=2;
			if (cod2 < pcod2) {	_ = g14_outSeqErr("g12_update_wordCand_vfyLemma 2 ", type1, pcod2, cod2, numOutSeq, maxNumErr); return -1; 	}
			pcod2 = cod2
		}
		
		if (min1[0:1] == "9") {break}
		
		if (min1 > preMin) {
			if (preMin != "") { 
				//seqCodeChange(index_ix1, wordSliceAlpha)
			}			
			index_ix1 = -1
			preMin    = min1 
			firstGrIx1 = -1; 
		} else {
			if (min1 < preMin) { 
				numOutSeq = g14_outSeqErr( "g12_update_wordCand_vfyLemma 4 ", type1, preMin, min1, numOutSeq, maxNumErr)
				if numOutSeq < 0 { return -1 }
				continue; 					
			}
		}		
		if (type1 == 1) {
			index_ix1 = j1  
			if firstGrIx1 < 0 { firstGrIx1 = j1 } 
			//if min1[1:] == "aufgehen" {	fmt.Println("   vfy  newLemma da testare ", min1[1:],   " from firstGrIx1=", firstGrIx1,  " index_ix1=", index_ix1, " ", wordSliceAlphaCand[j1]) }
			j1++;
		} else {
			if (index_ix1 < 0) { 
				//fmt.Println("               coppia word lemma senza una word in testo, ignorato " , wordLemmaPair[j2] )
				//numIgn2++
			} else {
				if wordSliceAlphaCand[index_ix1].wc_lemma != lemmaSlice[j2].leLemma {
					fmt.Println( red("ERRORE in g12_update_wordCand_vfyLemma "), 
						" wordSliceAlphaCand[",index_ix1,"].wc_lemma =", wordSliceAlphaCand[index_ix1].wc_lemma , 
						" no eguale a  lemmaSlice[",j2,"].leLemma=", lemmaSlice[j2].leLemma)
					return -1					
				} 	
				for k:=firstGrIx1; k <= index_ix1; k++ {   	
					wordSliceAlphaCand[k].wc_ixLemma = append(wordSliceAlphaCand[k].wc_ixLemma, j2)  
					//if min1[1:] == "aufgehen" {	fmt.Println("                    ", k, " ", wordSliceAlphaCand[k] ) }	
				}	
				trovati++		
			}
			j2++;
			 
		}		
		
	} // end for t0
	//-----------------
	
	//if (preMin != "") {seqCodeChange(index_ix1, wordSliceAlpha)}
	
	fmt.Println("  in update word + prefisso: candidati =", len(  wordSliceAlphaCand ), 
		" lemma verificati validi=", trovati) 
	return trovati	
	
} // end of g12_update_wordCand_vfyLemma
//----------------------------------