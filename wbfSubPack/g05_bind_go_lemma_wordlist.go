package wbfSubPack

	import (		
		"fmt"
		"strconv"
		"strings"		
		//"regexp"
		//"sort"
	)
//--------------------------------------------------------
func g05_bind_go_passToJs_lemmaWordList(lemmaToFind0 string, inpMaxWordLemma int, js_function string)   {
	sw_oneLemma:= true
	//listA := getListLemmaIndex( strings.ToLower(strings.TrimSpace( lemmaToFind0)), sw_oneLemma, inpMaxWordLemma) 
	
	//fmt.Println("g05_bind_go_passToJs_lemmaWordList --> g34_getListDirectLemmaIndex(", lemmaToFind0, " sw_oneLemmma=", sw_oneLemma)
	
	listA := g34_getListDirectLemmaIndex( strings.ToLower(strings.TrimSpace( lemmaToFind0)), sw_oneLemma, inpMaxWordLemma) 
	
	outS1:=""
	sw_oneLemma = true;
	
	for _, ix2:= range listA {
		//fmt.Println("LEMMA lemmaSlice[", ix2, "]=", lemmaSlice[ix2] ) 
		rowW := g05_build_one_lemma_row_word( ix2, lemmaSlice[ix2] , len(listA) , sw_oneLemma) 
		outS1 += rowW 	
	}	
	go_exec_js_function( js_function, outS1 ) 	
	
} // end of bind_go_passToJs_lemmaWordList 
//------------------------


//-----------------------------------------------------------
func containIntList( list1 []int, s int) bool{
	for _,tar:= range list1 {
		if tar == s { return true}
	} 
	return false 
}

//------------------------------------------------------


func g05_lemma_word_to_row( myLem1_name string, myLem2_name string, lw_lemma2 string, xWordAlpha wordUnAlphaStruct) string  {
		
	separ1:= ";."
	
		ix2 := -1
		for x1, oneLemmaName := range xWordAlpha.uLemmaL {
			if oneLemmaName == lw_lemma2 {
				ix2=x1; 
				break
			}  	
		}
		if ix2 < 0 { return ""}
		lastLemma := lemmaSlice[ ix2 ]
	
		ixL1 := xWordAlpha.uIxLemmaL[ix2]
				
		return xWordAlpha.uWordSeq + separ1 + xWordAlpha.uWord2 + separ1 + 
			"ix" + separ1 + 
			strconv.Itoa(xWordAlpha.uIxUnW_fr)   + separ1 + strconv.Itoa(xWordAlpha.uTotRow)  + separ1 + 
			xWordAlpha.uLemmaL[ix2]              + separ1 + 
			lemmaSlice[ixL1].leTran               + separ1 +   
			separ1 +  
			lastLemma.lePara                     + separ1 +  
			lastLemma.leExample                  + separ1 +  			
			strconv.Itoa(xWordAlpha.uTotExtrRow) + separ1 +  	
			xWordAlpha.uLearnedYN                + separ1 + 						
			"ixLemma" + separ1 + fmt.Sprint( xWordAlpha.uIxLemmaL[ix2] ) + separ1 + 	
			endOfLine 			
			 
} // end of lemma_word_to_row 

//-----------------------------------------------------------
func g05_build_one_lemma_row_word_ManyLemma( ixLemS int, myLem1 lemmaStruct, totNumLemmas int, sw_oneLemma bool ) string {
	separ1:= ";."

	ixW:= -1
	xWordAlpha:= wordUnAlphaStruct{}	
	
	var zeroW wordUnAlphaStruct 
	zeroW.uWordSeq = ""
	zeroW.uWord2 = ""
	zeroW.uIxUnW_al   = -1 	
	zeroW.uIxUnW_fr   = -1
	zeroW.uTotRow     = 0 
	zeroW.uTotExtrRow = 0
	zeroW.uIxFromWord_al = -1	
	zeroW.uSwSelRowG  = 0
	zeroW.uSwSelRowR  = 0  
	zeroW.uLearnedYN  = "y" 	
	
	//fmt.Println("\n1 g05_build_one_lemma_row_word_ManyLemma ", myLem1)
	totAllWordRows:=0
	if len( myLem1.leUnWord_al_IxList) > 0 {
		ixW = myLem1.leUnWord_al_IxList[0]
		
		lenW :=len(uniqueWordByAlpha)
		for _,ixW2:= range  myLem1.leUnWord_al_IxList{
			if ((ixW2 >=0) && (ixW2 < lenW)) {
				totAllWordRows += uniqueWordByAlpha[ixW2].uTotRow
			}
		}		
	}
	
	if ixW >= 0 {
			xWordAlpha = uniqueWordByAlpha[ixW] 
	} else {
			xWordAlpha = zeroW
			//xWordAlpha.uWordSeq = wPair.lWord2
			//xWordAlpha.uWord2   = wPair.lWord2
			xWordAlpha.uIxLemmaL= append(xWordAlpha.uIxLemmaL, myLem1.leFromIxLW) 
			xWordAlpha.uLemmaL  = append(xWordAlpha.uLemmaL,   myLem1.leLemma   )  
	}	
	
	//xWordAlpha:= uniqueWordByAlpha[ixW] 			
	//fmt.Println("2            g05_build_one_lemma_row_word_ManyLemma ixW=", ixW, " uniqueWordByAlpha[ixW] = xWordAlpha =", xWordAlpha)
	
	ix200 := -1
	for x1, oneLemmaName := range xWordAlpha.uLemmaL {
		if oneLemmaName == myLem1.leLemma {
			ix200=x1; 
			break
		}  	
	}
	var ixAlLem string 
	if ix200 < 0 {ixAlLem="-1"} else { ixAlLem = strconv.Itoa(xWordAlpha.uIxLemmaL[ix200]) }
	
	
	/**
	if ix200 < 0 {
		fmt.Println( red("errore "), " g05_build_one_lemma_row_word_ManyLemma ", " lemma=", myLem1.leLemma, 
			" contiene Word=",xWordAlpha.uWord2, " che ha una lista dei lemma ", xWordAlpha.uLemmaL, " che non contiene il lemma ", myLem1.leLemma)   
		return ""	
	}
	//-----------------------
	fmt.Println("            g05_build_one_lemma_row_word_ManyLemma ixW=", ixW, " uniqueWordByAlpha[ixW] = xWordAlpha =", xWordAlpha, " xWordAlpha.uLemmaL=", xWordAlpha.uLemmaL)
	
	//if ix2 < 0 {return ""}	
	ixL1:=-1
	
	if ix200 < 0 {
		ixL1 = -1
	} else {
		ixL1 = xWordAlpha.uIxLemmaL[ix200]		
		if ((ixL1 >= 0) && (ixL1 < len(lemmaSlice) )) {
			if lemmaSlice[ixL1].leLemma == myLem1.leLemma {ixL1 = -1} 
		}  
	}
	if (ixL1 < 0) {
		fmt.Println( red("errore"), " g05_build_one_lemma_row_word_ManyLemma ",  " lemma=", myLem1.leLemma, "  è diverso dal lemma puntato dalla parola puntata dal lemma")
		return ""
	} 
	fmt.Println("            g05_build_one_lemma_row_word_ManyLemma indice lemma ix200 =", ix200, " ixL1=", ixL1, " ixLemS=", ixLemS, 
		" lemmaSlice[ixL1].leLemma=", lemmaSlice[ixL1].leLemma ,
		" myLem1=", myLem1)
	***/
	/***many
				
		return xWordAlpha.uWordSeq + separ1 + xWordAlpha.uWord2 + separ1 + 
			"ix" + separ1 + 
			strconv.Itoa(xWordAlpha.uIxUnW_fr)   + separ1 + strconv.Itoa(xWordAlpha.uTotRow)  + separ1 + 
			xWordAlpha.uLemmaL[ix2]              + separ1 + 
			lemmaSlice[ixL1].leTran               + separ1 +   
			separ1 +  
			lastLemma.lePara                     + separ1 +  
	****/
	//------------------------
	rowW:= "" + separ1 + "" + separ1 + 
		"ix"   + separ1 + 
		"0"    + separ1 + 
		strconv.Itoa(totAllWordRows)    + separ1 + 
		myLem1.leLemma     				+ separ1 + 
		myLem1.leTran   	   			+ separ1 +      
		separ1 +  
		myLem1.lePara                  	+ separ1 +  
		myLem1.leExample               	+ separ1 +  
		"0" + separ1 +  	
		xWordAlpha.uLearnedYN             + separ1 + 						
		"ixLemma" + separ1 + ixAlLem + separ1 + 	
		endOfLine 	
	
	return rowW		
	
} // end of g05_build_one_lemma_row_word_ManyLemma

//--------------------------------------------------
/*
	lemma =  {gehen 8 27431 27438 [14 15 16 17 18 19 20 21] andare A1|A1|A1|A1|A1 gehen, geht, ging, ist gegangen|gehen, geht, ging, ist gegangen|gehen, geht, ging, ist gegangen|gehen, geht, ging, ist gegangen|gehen, geht, ging, ist gegangen Das geht nicht!|Ich muss zum Arzt gehen.|Ich weiß nicht, wie das geht.|Jetzt muss ich (aber) leider gehen.|Wie geht's? 5}
	 type lemmaStruct struct {
		leLemma    string        	
		leNumWords int 
		leFromIxLW  int 
		leToIxLW    int  
		leUnWord_al_IxList []int    // indice delle parole unique che puntano a questo lemma        
		leTran      string 
		leLevel     string  
		lePara      string  
		leExample   string  
		leNumPara   int	
	   } 		
	type wordLemmaPairStruct struct {
		lWord2 string 
		lWord2   string 
		lLemma   string
		lIxLemma  int
		lIxUnWord int
	} 
//---   
	   //--
type wordUnAlphaStruct struct {
	uWordSeq    string	
    uWord2      string		
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
	*/
//-------------------------------------------------

func g05_build_one_lemma_row_word_OnlyOneLemma(ix2 int, myLem1 lemmaStruct, totNumLemmas int, sw_oneLemma bool ) string {	
	// solo un lemma, dettaglio con tutte le parole
	
	//--------------------------------------------
	separ1:= ";."
	out1:= ""	
	//for _,ixW:= range myLem1.leUnWord_al_IxList {
	xWordAlpha:= wordUnAlphaStruct{}	
	
	var zeroW wordUnAlphaStruct 
	zeroW.uWordSeq = ""
	zeroW.uWord2 = ""
	zeroW.uIxUnW_al   = -1 	
	zeroW.uIxUnW_fr   = -1
	zeroW.uTotRow     = 0 
	zeroW.uTotExtrRow = 0
	zeroW.uIxFromWord_al = -1	
	zeroW.uSwSelRowG  = 0
	zeroW.uSwSelRowR  = 0  
	zeroW.uLearnedYN  = "y" 
	
	//fmt.Println("\n1 g05_build_one_lemma_row_word_OnlyOneLemma ", myLem1, " myLem1.leFromIxLW=", myLem1.leFromIxLW, " myLem1.leToIxLW=", myLem1.leToIxLW, " tran=", myLem1.leTran)
	
	//---------
	for g:=myLem1.leFromIxLW; g <= myLem1.leToIxLW; g++ {	
		wPair := wordLemmaPair_lemmaWordSeq[g]
		ixW := wPair.lIxUnWord_al 
		if ixW >= 0 {
			xWordAlpha = uniqueWordByAlpha[ixW] 
			//fmt.Println("   ", uniqueWordByAlpha[ixW] ) 
		} else {
			xWordAlpha = zeroW
			xWordAlpha.uWordSeq = wPair.lWord2
			xWordAlpha.uWord2   = wPair.lWord2
			xWordAlpha.uIxLemmaL= append(xWordAlpha.uIxLemmaL, wPair.lIxLemma ) 
			xWordAlpha.uLemmaL  = append(xWordAlpha.uLemmaL,   wPair.lLemma   )  
		}		
		
		ix2 := -1
		for x1, oneLemmaName := range xWordAlpha.uLemmaL {
			if oneLemmaName == myLem1.leLemma {
				ix2=x1; 
				break
			}  	
		}
		if ix2 < 0 {continue}
		
		ixL1 := xWordAlpha.uIxLemmaL[ix2]
		
		//fmt.Println("2    g05_build_one_lemma_row_word_OnlyOneLemma  g=", g, " lemma=", myLem1.leLemma, " ixW=", ixW, " ix2=", ix2, " ixL1=", ixL1 ); 
	
		
		rowW:= xWordAlpha.uWordSeq + separ1 + xWordAlpha.uWord2 + separ1 + 
			"ix" + separ1 + 
			strconv.Itoa(xWordAlpha.uIxUnW_fr)   + separ1 + 
			strconv.Itoa(xWordAlpha.uTotRow )    + separ1 + 
			xWordAlpha.uLemmaL[ix2]              + separ1 + 
			lemmaSlice[ixL1].leTran           + separ1 +  
			separ1 +  
			myLem1.lePara                     + separ1 +  
			myLem1.leExample                  + separ1 +  
			strconv.Itoa(xWordAlpha.uTotExtrRow) + separ1 +  	
			xWordAlpha.uLearnedYN                + separ1 + 						
			"ixLemma" + separ1 + fmt.Sprint( xWordAlpha.uIxLemmaL[ix2] ) + separ1 + 	
			endOfLine 	
		out1 += rowW	
		//fmt.Println(   "  g=", g, "  row=", rowW);
	}
	
	return out1	
	
} // end of g05_build_one_lemma_row_word_OnlyOneLemma

//-----------------------------------------------------------
func g05_build_one_lemma_row_word( ix2 int, myLem1 lemmaStruct, totNumLemmas int, sw_oneLemma bool) string {	

	//fmt.Println("\n", green("g05_build_one_lemma_row_word "), "lemma = ", myLem1, " totNumLemmas=", totNumLemmas  )
	/*
	lemma =  {gehen 8 27431 27438 [14 15 16 17 18 19 20 21] andare A1|A1|A1|A1|A1 gehen, geht, ging, ist gegangen|gehen, geht, ging, ist gegangen|gehen, geht, ging, ist gegangen|gehen, geht, ging, ist gegangen|gehen, geht, ging, ist gegangen Das geht nicht!|Ich muss zum Arzt gehen.|Ich weiß nicht, wie das geht.|Jetzt muss ich (aber) leider gehen.|Wie geht's? 5}
	 type lemmaStruct struct {
		leLemma    string        	
		leNumWords int 
		leFromIxLW  int 
		leToIxLW    int  
		leUnWord_al_IxList []int    // indice delle parole unique che puntano a questo lemma        
		leTran      string 
		leLevel     string  
		lePara      string  
		leExample   string  
		leNumPara   int	
	   } 
	*/
	
	if (totNumLemmas > 1) {  // se più di un lemma scrivi solo una riga
		return g05_build_one_lemma_row_word_ManyLemma(   ix2,  myLem1, totNumLemmas, sw_oneLemma)		
	} else {		
		return g05_build_one_lemma_row_word_OnlyOneLemma(ix2, myLem1,  totNumLemmas, sw_oneLemma)
	}
	

} // end of g05_build_one_lemma_row_word

//-----------------------------
