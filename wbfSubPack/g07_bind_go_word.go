package wbfSubPack

	import (
		"fmt"
		"strings"		
		"strconv"
		//"regexp"
		"sort"
	)
//--------------------------------------------------------
const PREF_MARKER = " :PREF: "
//-------------------------------------------

func g07_bind_go_passToJs_wordList( isChange_extrRow bool, fromWord int, numWords int, onlyThisLevel string, 
						sel_extrRow string, sel_toBeLearned string,  js_function string) {
					
		//fmt.Println( cyan("bind_go_passToJs_wordList"), "isChange_extrRow=", isChange_extrRow, " sel_extrRow=", sel_extrRow, " sel_toBeLearned=", sel_toBeLearned, 
		//				" from=",  fromWord, " numWords=", numWords )
		
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
				fmt.Println("error in bind_go_passToJs_wordList () numberOfUniqueWords=", numberOfUniqueWords)
				go_exec_js_function( js_function, ""); 	
				return
			}
		}
		//---------------------------
		sw_tobeLearnedOnly := ( sel_toBeLearned == "toBeLearned" )                //     ( 0 = 'allWords'   1 = 'toBeLearned' )
		//fmt.Println(red("bind_go_passToJs_wordList"), 	" sel_toBeLearned=", sel_toBeLearned, "  sw_tobeLearnedOnly=", sw_tobeLearnedOnly)
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
				fmt.Println("bind_go_passToJs_wordList () sel_extrRow != last_run_extrRow call build_and_elab_word_list() ")
				
				build_and_elab_word_list()
				
				fmt.Println("bind_go_passToJs_wordList () end of build_and_elab_word_list() ")
				
			}	
			//fmt.Println("bind_go_passToJs_wordList () return ")	
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
				
		//var xWordF wordUnAlphaStruct;  
		var outS1 string = ""; 
		//var row11 string;
		//var numNoTran = 0
		
		//fmt.Println("bind_go_passToJs_wordList () 0  from1 m=", from1 ) 
		
		onlyIfExtr := true 
		
		
		if from1 == 1 { from1=0; }
		numOut:=0
		//----------------------------
		for _,uWordFreq := range uniqueWordByFreq { 
			ixAl:= uWordFreq.fuIxUnW_al
			xWordAlpha := uniqueWordByAlpha[ixAl] 
			
			if sw_tobeLearnedOnly {
				if xWordAlpha.uLearnedYN == LEARNED_YES {  				
					continue
				} 
			}	
		
			//fmt.Println("call 3  loop  i=", i, "   ", uniqueWordByFreq[i]);
			
			sw, rowW := word_to_row("", onlyIfExtr, onlyThisLevel, xWordAlpha,-1 )  
			if sw {	
				numOut++
				if (numOut < from1) {continue}   // July7, 2025  
				//if (numOut < 20) {fmt.Println("     rowW=", rowW)  }	
				outS1 += rowW 						
				if numOut >= numWords { 
					break
				}
			}
		}	
		
		go_exec_js_function( js_function, outS1 ); 	
		
}  // end of bind_go_passToJs_wordList	

//---------------------------------------------------------
func print_wordIx( w1 wordUnAlphaStruct ) {
	fmt.Println("\t", cyan(w1.uWord2), " ix UniqueByFreq=",w1.uIxUnW_fr, " byAlfa=",w1.uIxUnW_al, " totR=", w1.uTotRow, " totExR=", w1.uTotExtrRow,
		 " uIxFromWord_al=", w1.uIxFromWord_al, " uSwSelRowG=", w1.uSwSelRowG, " uSwSelRowR=", w1.uSwSelRowR, " uLearnedYN=", w1.uLearnedYN, 
		 " uIxLemmaL=", w1.uIxLemmaL, " uLemmaL=", w1.uIxLemmaL) 
	
	/**		
		type wordUnAlphaStruct struct {
			uWordSeq    string	
		    uWord2      string	
			uIxUnW      int            // index of this word in the uniqueWordByFreq	
			uIxUnW_al   int            // index of this word in the uniqueWordByAlpha 	
			uTotRow     int 
			uTotExtrRow int
		    uIxWordFreq int            // index of this word in the wordSliceFreq	
			uSwSelRowG  int
			uSwSelRowR   int  
			uLearnedYN   string         // y n ( ie.yes,I learned / not yet  
			//uKnow_yes_ctr int 
			//uKnow_no_ctr  int         // a value > 0  means that this is a word that I don't know, ie. it's to be learned   
			uIxLemmaL  []int  
			uLemmaL    []string       // list of lemma 
			//uTranL     []string       // list of translation    
			uLevel     []string  
			uPara      []string  
			uExample   []string  		
		}
	**/ 	
} // end of print_wordIx



//---------------------------------------------------------------
func g07_bind_go_passToJs_word_known2(ixWord int, yesNot_len1  string, js_function string) {
	
	if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}	
		
	xWordFreq := uniqueWordByFreq[ixWord]  
	ixAlpha   := xWordFreq.fuIxUnW_al
	
	xWordAlpha:= uniqueWordByAlpha[ixAlpha]	
	
	ixFreq    := xWordAlpha.uIxUnW_fr  
	if ixFreq != ixWord { fmt.Println("ERRORE in bind_go_passToJs_word_known2() ixWord not equal to ixFreq ",  xWordAlpha.uWord2, " ixWord=", ixWord, " ixFreq=", ixFreq )   }
	
	xWordAlpha.uLearnedYN = yesNot_len1 // uniqueWordByFreq[ixWord].uKnow_yes_ctr = knowCtr
		
	outS1 := fmt.Sprint("ixWord=", ixWord, " ", xWordAlpha.uWord2, ", \t yesNo=", yesNot_len1, " learned=", xWordAlpha.uLearnedYN  )  
	//fmt.Println(green("word_known "), outS1)
	go_exec_js_function( js_function, outS1 ); 	
	
} // end of bind_go_passToJs_word_known2

//---------------------------------------------------------------
func g07_bind_go_passToJs_word_known(ixWord int, yesNo int, knowCtr int, js_function string) {

	if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}	
		
	xWordFreq := uniqueWordByFreq[ixWord]  
	ixAlpha   := xWordFreq.fuIxUnW_al
	
	xWordAlpha:= uniqueWordByAlpha[ixAlpha]	
	
	ixFreq    := xWordAlpha.uIxUnW_fr  
	if ixFreq != ixWord { fmt.Println("ERRORE in bind_go_passToJs_word_known() ixWord not equal to ixFreq ",  xWordAlpha.uWord2, " ixWord=", ixWord, " ixFreq=", ixFreq )   }
	
	if yesNo == 0 {	
		xWordAlpha.uLearnedYN = LEARNED_YES  // uniqueWordByFreq[ixWord].uKnow_yes_ctr = knowCtr
	} else {
		xWordAlpha.uLearnedYN = LEARNED_NOT  // uniqueWordByFreq[ixWord].uKnow_no_ctr  = knowCtr 
	} 
	
	
	outS1 := fmt.Sprint("ixWord=", ixWord, " ", xWordAlpha.uWord2, ", \t yesNo=", yesNo, " learned=", xWordAlpha.uLearnedYN  )  
	fmt.Println(green("word_known "), outS1)
	go_exec_js_function( js_function, outS1 ); 	
	
} // end of bind_go_passToJs_word_known

//---------------------------------------------- 

func g07_bind_go_passToJs_getRowsByIxWord( ixWord int, maxNumRow int, js_function string) {
	
	fmt.Println("func ", green("g07_bind_go_passToJs_getRowsByIxWord "), " ixWord=", ixWord, 
		" uniqueWordByFreq[ixWord]=", uniqueWordByFreq[ixWord])
	var outS1 string;
	//--------------		
	hd_tr := ""; 
	preL:= ""
	preW:=""
	listWords := ""
	
	if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}	
		
	xWordFreq := uniqueWordByFreq[ixWord]  
	ixAlpha   := xWordFreq.fuIxUnW_al
	
	xWordAlpha:= uniqueWordByAlpha[ixAlpha]	
	
	ixFreq    := xWordAlpha.uIxUnW_fr  
	if ixFreq != ixWord { fmt.Println("ERRORE in bind_go_passToJs_word_known() ixWord not equal to ixFreq ",  xWordAlpha.uWord2, " ixWord=", ixWord, " ixFreq=", ixFreq )   }
	
	
	listWords += " " +  xWordAlpha.uWord2
	
	aWord:= xWordAlpha.uWord2; 
	
	listIxRR := g07_getRowIndicesFromWord(xWordAlpha, maxNumRow)
	
	
	/***
	ls_lemma_ix_stellen  int	
	ls_lemma_stellen     string
	ls_pref_ein          string
	
	**/
	//lemmaPrefList:=""
	//-------------
	for z:=0; z < len(xWordAlpha.uLemmaL); z++  {
		ixL1:= xWordAlpha.uIxLemmaL[z]
		LeS := lemmaSlice[ixL1]  
		//lemmaPrefList +=  " " + LeS.ls_pref_ein		
		newL:= xWordAlpha.uLemmaL[z]
		if newL != LeS.leLemma {
			continue;  // error 
		}			
		newT:= LeS.leTran
		//newP:= xWordAlpha.uPara[z]
		newP:= LeS.lePara
		if newL == preL { 
			newL=""
			newT=""
			if preW != xWordAlpha.uWord2 { 
				hd_tr += xWordAlpha.uWord2 + " ";  
				preW = xWordAlpha.uWord2
			} 
		} else { 				
			if newP == "" {
				newP = newL;
			} 				
			preL = newL 			
			if hd_tr != "" { hd_tr += "\n" 	}  // 14giugno 
			
			hd_tr += " :lemma="  + newP + " :tran=" + newT  + " :wordsInLemma=" +	xWordAlpha.uWord2 + " "
			//hd_tr += " :lemma="  + newP + " :tran=" + newT + " "
			
			preW =  xWordAlpha.uWord2
		} 			
	} // end for z
	
	/**
	if lemmaPrefList != "" {
		listWords += PREF_MARKER + lemmaPrefList
	}
	**/
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
func g07_bind_go_passToJs_getRowsByIxLemma( ixLemma int, max_num_row4lemma int, js_function string) {
	/*
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
	
	leS:= lemmaSlice[ixLemma]; 
	lemmaToFind0:= leS.leLemma;
	lemmaTran   := leS.leTran
	//-----------------------------
	listIxRR        := make([]int,0, max_num_row4lemma)	
	listIxRR_temp   := make([]int,0, max_num_row4lemma)				
	outS1 :=""
	//--------------		
	hd_tr := ""; 
	
	listWords := ""
	listWords_pref:=""
	//------------------	
	totRR:=0 		
	//---------
	for _, ixWord:= range leS.leUnWord_al_IxList {  
	
		xWordF := uniqueWordByAlpha[ixWord]

		listWords += " " +  xWordF.uWord2	
		hd_tr += " :lemma="  + lemmaToFind0 			
		hd_tr +=  " :tran=" + lemmaTran + " :wordsInLemma=" +	xWordF.uWord2 + " "
				
		listIxRR_temp = g07_getRowIndicesFromWord(xWordF,  max_num_row4lemma)   // indici row da word 
		listIxRR      = append(listIxRR, listIxRR_temp...)
		totRR         = len(listIxRR)
		if totRR > max_num_row4lemma { break }
	}
	//---------------------
	
	
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
	
	if (listWords_pref != "") {
		listWords += PREF_MARKER + listWords_pref;   
	}	
	header:= "<HEADER>\n" + "<WORD>" + aWord + ",L:" + strings.TrimSpace(listWords) + "</WORD>"
	//header:= "<HEADER>\n" + "<WORD>" + aWord + "</WORD>"
	header += hd_tr   // 14giugno
	header += "</HEADER>\n"
	
	go_exec_js_function( js_function, header + outS1 ); 	
		
} // end of  bind_go_passToJs_getRowsByIxLemma  

//----------------------------------------------------
func g07_getRowIndicesFromWord(xWordAlpha wordUnAlphaStruct, maxNumRow int) []int {
	
		var ixFromList = xWordAlpha.uIxFromWord_al                     // index of this word in the wordSliceAlpha	( una word per ogni row ) 
		var ixToList   = ixFromList + xWordAlpha.uTotRow;  // totRow è il numero di righe in cui si trova la parola ( qualcosa non va, dovrebbe essere egaule al numero di word   
		var maxTo1     = ixFromList + maxNumRow; 
		
		
		if ixToList > maxTo1        { ixToList = maxTo1; }
	
		if ixToList > len(wordSliceAlpha) { ixToList = len(wordSliceAlpha) }
		
		//fmt.Println("     g07_getRowIndicesFromWord  ",  " ixToList=", ixToList, " (ixToList - ixFromList)=", (ixToList - ixFromList) )

		listIxRR := make([]int,0, (ixToList - ixFromList) )
		
		for ix4:=ixFromList; ix4 < ixToList; ix4++ {
			wS1 := wordSliceAlpha[ix4] 
			//  .wIxRow           indice del row che contiene la parola 	
			//  .wSwSelRowR	      1 or 2: 1 SEL_EXTR_ROW, 2 SEL_NO_EXTR_ROW  ( serve per dare la precedenza alle righe che appartengono al brano selezionato    
			listIxRR = append( listIxRR, wS1.wSwSelRowR * 100000 +  wS1.wIxRow)  	
		} 	
		
		sort.Ints(listIxRR) 
		
		return listIxRR 
		
} // end of getRowIndicesFromWord

//----------------------------------------------------
func getRowIndicesFromIxAlphaWord(ixWord int, maxNumRow int) []int {
		fmt.Println("getRowIndicesFromIxAlphaWord ixWord=", ixWord, " len(uniqueWordByAlpha)=", len(uniqueWordByAlpha) )
			
	if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}			
	ixAlpha := ixWord;
	xWordAlpha:= uniqueWordByAlpha[ixAlpha]	
	
	ixFreq    := xWordAlpha.uIxUnW_fr  
	if ixFreq != ixWord { fmt.Println("ERRORE in bind_go_passToJs_word_known() ixWord not equal to ixFreq ",  xWordAlpha.uWord2, " ixWord=", ixWord, " ixFreq=", ixFreq )   }
	
		
		
		

		var ixFromList = xWordAlpha.uIxFromWord_al                     // index of this word in the wordSliceAlpha	( una word per ogni row ) 
		var ixToList   = ixFromList + xWordAlpha.uTotRow;
		var maxTo1     = ixFromList + maxNumRow; 		
		
		if ixToList > maxTo1        { ixToList = maxTo1; }
		if ixToList > len(wordSliceAlpha) { ixToList = len(wordSliceAlpha) }

		listIxRR := make([]int,0, (ixToList - ixFromList) )
		
		for ix4:=ixFromList; ix4 < ixToList; ix4++ {
			wS1 := wordSliceAlpha[ix4] 
			//  .wIxRow           indice del row che contiene la parola 	
			//  .wSwSelRowR	      1 or 2: 1 SEL_EXTR_ROW, 2 SEL_NO_EXTR_ROW  ( serve per dare la precedenza alle righe che appartengono al brano selezionato    
			listIxRR = append( listIxRR, wS1.wSwSelRowR * 100000 +  wS1.wIxRow)  	
		} 	
		
		sort.Ints(listIxRR) 
		
		return listIxRR 
		
} // end of getRowIndicesFromIxFreqWord

//----------------------------------------------------
func getRowIndicesFromIxFreqWord(ixWord int, maxNumRow int) []int {
		fmt.Println("getRowIndicesFromIxFreqWord ixWord=", ixWord, " len(uniqueWordByFreq)=", len(uniqueWordByFreq) )
			
	if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}	
		
	xWordFreq := uniqueWordByFreq[ixWord]  
	ixAlpha   := xWordFreq.fuIxUnW_al
	
	xWordAlpha:= uniqueWordByAlpha[ixAlpha]	
	
	ixFreq    := xWordAlpha.uIxUnW_fr  
	if ixFreq != ixWord { fmt.Println("ERRORE in bind_go_passToJs_word_known() ixWord not equal to ixFreq ",  xWordAlpha.uWord2, " ixWord=", ixWord, " ixFreq=", ixFreq )   }
	
		
		
		

		var ixFromList = xWordAlpha.uIxFromWord_al                     // index of this word in the wordSliceAlpha	( una word per ogni row ) 
		var ixToList   = ixFromList + xWordAlpha.uTotRow;
		var maxTo1     = ixFromList + maxNumRow; 		
		
		if ixToList > maxTo1        { ixToList = maxTo1; }
		if ixToList > len(wordSliceAlpha) { ixToList = len(wordSliceAlpha) }

		listIxRR := make([]int,0, (ixToList - ixFromList) )
		
		for ix4:=ixFromList; ix4 < ixToList; ix4++ {
			wS1 := wordSliceAlpha[ix4] 
			//  .wIxRow           indice del row che contiene la parola 	
			//  .wSwSelRowR	      1 or 2: 1 SEL_EXTR_ROW, 2 SEL_NO_EXTR_ROW  ( serve per dare la precedenza alle righe che appartengono al brano selezionato    
			listIxRR = append( listIxRR, wS1.wSwSelRowR * 100000 +  wS1.wIxRow)  	
		} 	
		
		sort.Ints(listIxRR) 
		
		return listIxRR 
		
} // end of getRowIndicesFromIxFreqWord
//---------------------------------------
