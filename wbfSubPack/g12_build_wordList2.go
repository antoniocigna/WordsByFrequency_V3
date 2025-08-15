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

//-----------------------------------------

func g12_add_totRow_and_indexLemmaPair() {

	g12_add_totRow_and_indexLemmaPairUNO()
	g12_add_totRow_and_indexLemmaPairDUE()

} // end of add_totRow_and_indexLemmaPair(

//------------------------------------------------

func g12_add_totRow_and_indexLemmaPairUNO() {	
	fmt.Println("   func ", green("add_totRow_and_indexLemmaPair") )
	/*
	each element of wordSliceAlpha contains a word (the same word may be in several rows) 
	the number of repetition of a word (totRow) is put in its element  ( later will be put in each row that contain it) 
		eg.  one 3, one 3, one 3, two 4, two 4, two 4, two 4	
	*/
	
	totR  := 0	
	
	
	pre_wSwSelRow := SEL_NO_EXTR_ROW 
	//----------------
	tot_extrRow:=0
	lastIx:=0
	// alla fine dello slice c'è la parola LAST_WORD  che rende non necessaria la gestione di fine file  
	
	//fmt.Println(" len(listAllLemmaFromFile) = ", len(listAllLemmaFromFile) )
	fmt.Println(" len(lemmaSlice) = ", len(lemmaSlice) )
	maxNumWApp:= len(wordSliceAlpha)
	wordSliceAlphaToApp = make([]wordStruct, 0, maxNumWApp)
	
	//wordSliceAlpha in ordine di wWordSeq, wWord2, wNfile
	ix1   := 0
	if len(wordSliceAlpha) < 1 {return }
	preW  := wordSliceAlpha[0].wWordSeq;	
	//--------------------------------------
	for ix2, wS1 := range wordSliceAlpha {
		
		if (wS1.wWordSeq != preW) {			
			g12_manage_one_word_of_manyRowsUNO(ix1, ix2, preW, pre_wSwSelRow , tot_extrRow, totR, lastIx)			
			pre_wSwSelRow = SEL_NO_EXTR_ROW 
			totR = 0
			tot_extrRow = 0
			ix1  = ix2; 
			preW = wS1.wWordSeq; 
		} 
		
		if wS1.wSwSelRowR == SEL_EXTR_ROW {   // se almeno uno è "estratto", tutti lo sono 
			pre_wSwSelRow = SEL_EXTR_ROW 
			tot_extrRow++	
			//if (wS1.wWord2 == "schrift") { fmt.Println("ANTO addTotRowToWord  ",  wS1 , " SEL_EXTR_ROW=", SEL_EXTR_ROW, "  tot_extrRow=", tot_extrRow) }  //
		} 
		totR++;     	
	}	
	//------		
	
	
} // end of add_totRow_and_indexLemmaPair

//-----------------------------------------------------------------

func g12_manage_one_word_of_manyRowsUNO(ix1 int, ix2 int, preW string, pre_wSwSelRow int, tot_extrRow int, totR int, lastIx int) {
	swNoLemma:= false 
	/*
		var NO_LEMMA_WORD  = "_"
		var NO_LEMMA_LEMMA = "_"
		var NO_LEMMA_INDEX = 0
	*/
	ixLemmaPairFoundList := lookForAllLemmas( preW, lastIx) 
	if len(ixLemmaPairFoundList) < 1 {	
		swNoLemma = true 
		ixLemmaPairFoundList = append(ixLemmaPairFoundList, NO_LEMMA_INDEX)
	} else {
		if ixLemmaPairFoundList[0] < 0 { 
			swNoLemma = true 
			ixLemmaPairFoundList[0] = NO_LEMMA_INDEX
		}
	}	
	
	//fmt.Println("g12_manage_one_word_of_manyRowsUNO  ", preW,  " ix1=", ix1, " ix2=", ix2, " totR=", totR,  "  ixLemmaPairFoundList=", ixLemmaPairFoundList);
	
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
			 ixAF = binarySearch_string(listAllLemmaFromFile, le_Lemma) 
			 if ixAF >=0 { 
				lemmaIndex = append(lemmaIndex, ixAF)
				wS22.wIxLemmaList = append(wS22.wIxLemmaList, ixAF) 
			 }		
			 
			 if ixLe == NO_LEMMA_INDEX {continue}
			 //fmt.Println("               wRord2=", wS22.wWord2, " lemma=", le_Lemma)	
			 pref3 := strings.Fields(wS22.wListPref) 
			 if len(pref3) < 1 {continue}
			 
			 for _, onePref:= range pref3 {
				 newLemma:= onePref + le_Lemma 
				 //fmt.Println("              1 word2=", wS22.wWord2, " lemma=", le_Lemma, " newLemma =", newLemma)	
				 ixAF = binarySearch_string(listAllLemmaFromFile, newLemma) 
				 if ixAF < 0 { continue}
				 lemmaIndex = append(lemmaIndex, ixAF)
				 wS33 = wS22
				 wS33.wWord2   =  wS22.wWord2   + " ... " + onePref 
				 wS33.wWordSeq =  wS22.wWordSeq + " ... " + onePref 
				 wS33.wIxLemmaList = append( wS33.wIxLemmaList, ixAF)
				 wS33.wSwSelRowG    = pre_wSwSelRow; 	// se esiste almeno un richiamo a una riga estratta ( wSwSelRowR)allora questo segnale è ripetuto come wSwSelRowG
				 wS33.wTotExtrRow   = tot_extrRow 
				 wS33.wTotRow       = totR;   // se una parola è ripetuta 3 volte, ad ogni parola è associato 3  		
				 wordSliceAlphaToApp = append(wordSliceAlphaToApp, wS33)
				 
				 //fmt.Println("                       wS33: ",  wS33.wWord2 , " ", wS33.wIxLemmaList )
				 
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
} // end of manage_one_word_of_manyRows
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
	
	wordSliceAlpha = append(wordSliceAlpha, wordSliceAlphaToApp...)
	
	wordSliceAlphaToApp = make([]wordStruct, 0, 0)  // rilascio lo spazio 
	//-----------------------------------------------------------	
	sort.Slice(wordSliceAlpha, func(i, j int) bool {
		if wordSliceAlpha[i].wWordSeq != wordSliceAlpha[j].wWordSeq {
			return wordSliceAlpha[i].wWordSeq < wordSliceAlpha[j].wWordSeq            // word  ascending order (eg.   a before b ) 
		} else {
			if wordSliceAlpha[i].wWord2 != wordSliceAlpha[j].wWord2 {
				return wordSliceAlpha[i].wWord2 < wordSliceAlpha[j].wWord2  
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
	preW  := wordSliceAlpha[0].wWordSeq;	
	//--------------------------------------
	for ix2, wS1 := range wordSliceAlpha {
		
		if (wS1.wWordSeq != preW) {			
			g12_manage_one_word_of_manyRowsDUE(ix1, ix2, preW, pre_wSwSelRow , tot_extrRow, totR, lastIx)			
			pre_wSwSelRow = SEL_NO_EXTR_ROW 
			totR = 0
			tot_extrRow = 0
			ix1  = ix2; 
			preW = wS1.wWordSeq; 
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
	xWordF.uWordSeq  	= wS1.wWordSeq;
	xWordF.uWord2    	= wS1.wWord2;
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
	
	/***
	if len(wS1.wListPref) > 0 {
		for _, unPref := range wS1.wListPref {
			xWordF.uWordSeq  = wS1.wWordSeq + "..." + unPref;
			xWordF.uWord2    = wS1.wWord2   + "..." + unPref;
			xWordF.uIxUnW_fr     = len(uniqueWordByFreq)  
			uniqueWordByFreq = append( uniqueWordByFreq, xWordF);  					
		}					
	} 
	***/
	
	g13_addLemmaTranParadigmaToUniqueWord( xWordF , wS1.wIxLemmaList)  // append to uniqueWordByAlpha  
	
	//fmt.Println("STAT. ", n1, " ", xWordF.word, " numWordUn=", numWordUn,  " numWordRi=", numWordRi, " percIx=", percIx, " ", sS.uniquePerc,  " sS.totPerc=" ,  sS.totPerc); 
		 			
 } // end for buildUnique

//------------------------------------