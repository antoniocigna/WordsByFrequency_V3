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

func bind_go_passToJs_wordList( isChange_extrRow bool, fromWord int, numWords int, onlyThisLevel string, 
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

//--------------------
func bind_go_passToJs_getWordByIndex2( ixWord int, swOnlyThisWordRows bool, maxNumRow int, js_function string) {
		
		if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}	
		
		var xWordF     = uniqueWordByFreq[ixWord]  	
		aWord:= xWordF.uWord2; 
		
		OLD_bind_go_passToJs_thisWordRowList( aWord, swOnlyThisWordRows, maxNumRow, js_function)
		
} // end of bind_go_passToJs_getWordByIndex2


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

//---------------------------------------------- 
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
	/***
	ls_lemma_ix_stellen  int	
	ls_lemma_stellen     string
	ls_pref_ein          string
	
	**/
	lemmaPrefList:=""
	//-------------
	for z:=0; z < len(xWordF.uLemmaL); z++  {
		ixL1:= xWordF.uIxLemmaL[z]
		LeS := lemmaSlice[ixL1]  
		lemmaPrefList +=  " " + LeS.ls_pref_ein		
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
	
	
	if lemmaPrefList != "" {
		listWords += PREF_MARKER + lemmaPrefList
	}
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
		
		
		/**
		fmt.Println( green("bind_go_passToJs_getRowsByIxLemma"), " ixLemma=", ixLemma,  " lemmaSlice[]=", leS)
		
		for f:= leS.leFromIxLW; f <= leS.leToIxLW; f++ {
			
			wordLW := lemma_word_ix[f]
		
			fmt.Println("\t" , " lemma=", wordLW.lw_lemma2, " word=", wordLW.lw_word, " ixLemma=", wordLW.lw_ixLemma, " ixFreq=", wordLW.lw_ixWordUnFr) 
			
			//???? bind_go_passToJs_thisWordRowList( aWord string,  maxNumRow int, js_function string) { 
		
		}  
		**/
		//---------------------
		
		lemmaCod:= newCode( lemmaToFind0 )
		
		
		
				
		listIxRR        := make([]int,0, max_num_row4lemma)	
		listIxRR_temp   := make([]int,0, max_num_row4lemma)			
		
		outS1 :=""
			//--------------		
			hd_tr := ""; 
			
			listWords := ""
			listWords_pref:=""
			//------------------
	
		totRR:=0 
		
		
		
		//------------------------------
		for kLe:= leS.leFromIxLW; kLe <= leS.leToIxLW; kLe++ {
		
			if lemma_word_ix[kLe].lw_lemmaCod != lemmaCod { continue}		// error 
			
			ixWord := lemma_word_ix[kLe].lw_ixWordUnFr 			        // indice word da lemma 
			xWordF := uniqueWordByFreq[ixWord]

			listWords += " " +  xWordF.uWord2	
			hd_tr += " :lemma="  + lemmaToFind0 
			if leS.ls_pref_ein != "" {	
				listWords_pref += " " + leS.ls_pref_ein
				hd_tr += " = " + leS.ls_pref_ein +"(" + leS.ls_pref_tran+ ")" + " + " +  leS.ls_lemma_stellen  
			}	
			
			hd_tr +=  " :tran=" + lemmaTran + " :wordsInLemma=" +	xWordF.uWord2 + " "
			
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
	
	if (listWords_pref != "") {
		listWords += PREF_MARKER + listWords_pref;   
	}	
	header:= "<HEADER>\n" + "<WORD>" + aWord + ",L:" + strings.TrimSpace(listWords) + "</WORD>"
	//header:= "<HEADER>\n" + "<WORD>" + aWord + "</WORD>"
	header += hd_tr   // 14giugno
	header += "</HEADER>\n"
	
	go_exec_js_function( js_function, header + outS1 ); 					
		
		go_exec_js_function( js_function, header + outS1 ); 	
		
} // end of  bind_go_passToJs_getRowsByIxLemma  


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
//---------------------------------------
