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
	
	fmt.Println("g05_bind_go_passToJs_lemmaWordList --> g34_getListDirectLemmaIndex(", lemmaToFind0, " sw_oneLemmma=", sw_oneLemma)
	
	listA := g34_getListDirectLemmaIndex( strings.ToLower(strings.TrimSpace( lemmaToFind0)), sw_oneLemma, inpMaxWordLemma) 
	
	outS1:=""
	for _, ix2:= range listA {
		//fmt.Println("LEMMA lemmaSlice[", ix2, "]=", lemmaSlice[ix2] ) 
		rowW := g05_build_one_lemma_row_word( lemmaSlice[ix2] , len(listA) ) 
		outS1 += rowW 	
	}	
	go_exec_js_function( js_function, outS1 ) 	
	
} // end of bind_go_passToJs_lemmaWordList 
//------------------------
func TOGLIbind_go_passToJs_lemmaWordList(lemmaToFind0 string, inpMaxWordLemma int, js_function string)   {
		/**
			lista tutte le parole col lemma indicato		
		**/
		//lemmaCod:= newCode( lemmaToFind0 )
		//var onlyThisLevel string = "any" ; // "A0"  // questo deve arrivare da parametro  
		//onlyIfExtr := false 
		
		outS1 := "" 
		
		fromIx, toIx:= lookForLemma(lemmaToFind0,0) // find the index of lemma in lemmaSlice 
		ixLe :=-1
		for k:= fromIx; k <= toIx; k++ {
			if lemmaSlice[k].leLemma == lemmaToFind0 { ixLe = k;  break }
		}
		if ixLe < 0 {
			go_exec_js_function( js_function,   "NONE," + lemmaToFind0 )
			return
		}
		
		//lem1_prefix      := ""
		//lem1_prefix_tran := ""
		myLem1_name      := lemmaToFind0
		myLem2_name      := ""
		myLem2_ix        := -1
		var myLem2 lemmaStruct
		myLem1 := lemmaSlice[ixLe]
		
		/*					
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
		*/
		
		
		//fmt.Println("bind_go_passToJs_lemmaWordList target lemma=", lemmaToFind0)
		//fmt.Println("\t", "               myLem1=", myLem1)   
		
		/**
		bind_go_passToJs_lemmaWordList target lemma= einstellen  trovato myLem= {einstellen 2  5678  5679 impostare    14893 stellen ein dentro, verso l'interno []}
					servono le parole con lemma einstellen, 
						ma anche quelle con lemma stellen:   es. stelle appartiene a stellen e insieme ad ein al lemma einstellen 
						ma poi le presento come se fossero tutte dei einstellen  			
		----
		bind_go_passToJs_lemmaWordList target lemma= stellen     trovato myLem= {stellen   21 20259 20279 fornire         -1                                     [718 1170 1362 4097 7175 7296 17303]}		
		
					servono solo le parole con lemma stellen, es. stellt, gestellt 
		**/
		/**
		myLem2_ix = myLem1.ls_lemma_ix_stellen
		if myLem2_ix >= 0 {
			myLem2_name = myLem1.ls_lemma_stellen
			//lem1_prefix = myLem1.ls_pref_ein
			//lem1_prefix_tran = myLem1.ls_pref_tran
			myLem2 = lemmaSlice[myLem2_ix]
			//fmt.Println("\t", " myLem2_name=", myLem2_name , "myLem2=", myLem2,    "  lem1_prefix =", lem1_prefix, " tran=", lem1_prefix_tran) 
		}
		**/
		//---------------------------------
		//  .leFromIxLW / .leToIxLW    are indicies of element in "lemma_word_ix"
		listaIxLemmaWordIx:= make([]int, 0, 200 ) 
		for k := myLem1.leFromIxLW; k <= myLem1.leToIxLW; k++ {
			listaIxLemmaWordIx = append(listaIxLemmaWordIx, k) 
		} 
		if myLem2_ix >= 0 {
			for k := myLem2.leFromIxLW; k <= myLem2.leToIxLW; k++ {				
				if containIntList(listaIxLemmaWordIx, k) == false {  
					listaIxLemmaWordIx = append(listaIxLemmaWordIx, k) 
				}
			} 
		}
		//fmt.Println(" listaIxLemmaWordIx=", listaIxLemmaWordIx)
		numO:= 0
		//---------------------------
		
		for _,k:= range listaIxLemmaWordIx {
			lw1:= lemma_word_ix[k]
			
			
			if ((lw1.lw_origLemma != myLem1_name) && (lw1.lw_origLemma != myLem2_name)) {
				continue
			} 
			
			ix := lemma_word_ix[k].lw_ixWordUnAl 		
			
				
			rowW := g05_lemma_word_to_row(myLem1_name, myLem2_name, lw1.lw_lemma2, uniqueWordByAlpha[ix] )  
				
				numO++
				if numO > inpMaxWordLemma { break } 
				outS1 += rowW 
			  	
			//fmt.Println("    rowW=", rowW)	
		} 
		
		if len(outS1)< 3 {
			outS1 = "NONE," + lemmaToFind0; 			
			//fmt.Println(" bind_go_passToJs_lemmaWordList() ", outS1) 
		} 	
		
		go_exec_js_function( js_function, outS1 ); 		
		
	
				
} // end of TOGLIbind_go_passToJs_lemmaWordList

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
func g05_build_one_lemma_row_word_ManyLemma( myLem1 lemmaStruct, totNumLemmas int) string {
	separ1:= ";."

	ixW:= 0
	
	if len( myLem1.leUnWord_al_IxList) > 0 {ixW = myLem1.leUnWord_al_IxList[0]}
	
	xWordAlpha:= uniqueWordByAlpha[ixW] 			
	ix2 := -1
	for x1, oneLemmaName := range xWordAlpha.uLemmaL {
		if oneLemmaName == myLem1.leLemma {
			ix2=x1; 
			break
		}  	
	}
	if ix2 < 0 {return ""}	
	
	ixL1 := xWordAlpha.uIxLemmaL[ix2]
	
	rowW:= "" + separ1 + "" + separ1 + 
		"ix"   + separ1 + 
		"0"    + separ1 + 
		"0"    + separ1 + 
		myLem1.leLemma     + separ1    + 
		lemmaSlice[ixL1].leTran        + separ1 +      
		separ1 +  
		myLem1.lePara                  + separ1 +  
		myLem1.leExample               + separ1 +  
		"0" + separ1 +  	
		xWordAlpha.uLearnedYN             + separ1 + 						
		"ixLemma" + separ1 + fmt.Sprint( xWordAlpha.uIxLemmaL[ix2] ) + separ1 + 	
		endOfLine 	
	
	return rowW		
	
} // end of g05_build_one_lemma_row_word_OnlyOneLemma
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
		lWordSeq string 
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

func g05_build_one_lemma_row_word_OnlyOneLemma( myLem1 lemmaStruct, totNumLemmas int) string {	
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
	
	
	
	//---------
	for g:=myLem1.leFromIxLW; g <= myLem1.leToIxLW; g++ {	
		wPair := wordLemmaPair_lemmaWordSeq[g]
		ixW := wPair.lIxUnWord_al 
		if ixW >= 0 {
			xWordAlpha = uniqueWordByAlpha[ixW] 
			fmt.Println("   ", uniqueWordByAlpha[ixW] ) 
		} else {
			xWordAlpha = zeroW
			xWordAlpha.uWordSeq = wPair.lWordSeq
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
	}
	
	return out1	
	
} // end of g05_build_one_lemma_row_word_OnlyOneLemma
	
//-----------------------------------------------------------

func togli2g05_build_one_lemma_row_word_OnlyOneLemma( myLem1 lemmaStruct, totNumLemmas int) string {	
	// solo un lemma, dettaglio con tutte le parole
	
	fmt.Println("\n\n", "g05_build_one_lemma_row_word_OnlyOneLemma( myLem1 = ", myLem1)
	
	for g:=myLem1.leFromIxLW; g <= myLem1.leToIxLW; g++ {
		fmt.Println( green("    wordLemmaPair_lemmaWordSeq["), g, "] = ",  wordLemmaPair_lemmaWordSeq[g]  )
	}
	
	//--------------------------------------------
	separ1:= ";."
	out1:= ""	
	for _,ixW:= range myLem1.leUnWord_al_IxList {
		xWordAlpha:= uniqueWordByAlpha[ixW] 
		fmt.Println("   ", uniqueWordByAlpha[ixW] )  
		
		ix2 := -1
		for x1, oneLemmaName := range xWordAlpha.uLemmaL {
			if oneLemmaName == myLem1.leLemma {
				ix2=x1; 
				break
			}  	
		}
		if ix2 < 0 {continue}
		
		ixL1 := xWordAlpha.uIxLemmaL[ix2]
		//---
		rowW:= xWordAlpha.uWordSeq + separ1 + xWordAlpha.uWord2 + separ1 + 
			"ix" + separ1 + 
			strconv.Itoa(xWordAlpha.uIxUnW_fr)   + separ1 + 
			strconv.Itoa(xWordAlpha.uTotRow)     + separ1 + 
			xWordAlpha.uLemmaL[ix2]              + separ1 + 
			lemmaSlice[ixL1].leTran              + separ1 +  
			separ1 +  
			myLem1.lePara                     + separ1 +  
			myLem1.leExample                  + separ1 +  
			strconv.Itoa(xWordAlpha.uTotExtrRow) + separ1 +  	
			xWordAlpha.uLearnedYN                + separ1 + 						
			"ixLemma" + separ1 + fmt.Sprint( xWordAlpha.uIxLemmaL[ix2] ) + separ1 + 	
			endOfLine 	
		out1 += rowW	
	}
	
	return out1	
	
} // end of TOGLI2g05_build_one_lemma_row_word_OnlyOneLemma

//-----------------------------------------------------------
func g05_build_one_lemma_row_word( myLem1 lemmaStruct, totNumLemmas int) string {	

	fmt.Println( green("g05_build_one_lemma_row_word "), "lemma = ", myLem1, " totNumLemmas=", totNumLemmas  )
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
		return g05_build_one_lemma_row_word_ManyLemma(   myLem1, totNumLemmas)		
	} else {		
		return g05_build_one_lemma_row_word_OnlyOneLemma(myLem1, totNumLemmas)
	}
	

} // end of g05_build_one_lemma_row_word

//-----------------------------

func TOGLIg05_build_one_lemma_row_word( myLem1 lemmaStruct) string {	

		//lem1_prefix      := ""
		//lem1_prefix_tran := ""

		myLem1_name      := myLem1.leLemma

		myLem2_name      := ""
		//myLem2_ix        := -1
		
		//var myLem2 lemmaStruct	
		
		fmt.Println( green("g05_build_one_lemma_row_word"), "   target lemma=", myLem1_name, " myLem1.leFromIxLW=",myLem1.leFromIxLW, " myLem1.leToIxLW=", myLem1.leToIxLW )
		
		/**
		myLem2_ix = myLem1.ls_lemma_ix_stellen
		if myLem2_ix >= 0 {
			myLem2_name = myLem1.ls_lemma_stellen
		}
		**/
		//---------------------------------
		//  .leFromIxLW / .leToIxLW    are indicies of element in "lemma_word_ix"
		
		listaIxLemmaWordIx:= make([]int, 0, 200 ) 
		
		for k := myLem1.leFromIxLW; k <= myLem1.leToIxLW; k++ {
			listaIxLemmaWordIx = append(listaIxLemmaWordIx, k) 
			break  //  only once 
		} 
		
		//fmt.Println(" listaIxLemmaWordIx=", listaIxLemmaWordIx)
		//numO:= 0
		//---------------------------
		
		for _,k:= range listaIxLemmaWordIx {
			lw1:= lemma_word_ix[k]
		
			
			ix := lemma_word_ix[k].lw_ixWordUnAl 		
			
			//fmt.Println("           word=", uniqueWordByFreq[ix].uWord2 )  
				
			rowW := g05_lemma_word_to_row(myLem1_name, myLem2_name, lw1.lw_lemma2, uniqueWordByAlpha[ix] )  
			
			//fmt.Println( "return rowW=", rowW)
			
			return  rowW // only once 	
		} 
	
	//fmt.Println( "return spazio")
	
	return ""
	
} // end of TOGLIg05_build_one_lemma_row_word		

//-------------------------------------------

