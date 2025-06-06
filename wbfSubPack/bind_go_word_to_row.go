package wbfSubPack

	import (
		"fmt"
		"strings"		
		"strconv"
		"sort"
		"slices"
		"regexp"
	)
//--------------------------------------------------------
func add_wordCombinations( wordL1 []string, wordL2 []string )  []string {
	wor3List:= []string{} 
	for _, wor1 := range wordL1 {	
		if wor1 == "" {continue}
		for _, wor2 := range wordL2 {
			if wor2 == "" {continue}
			wor3List = append(wor3List, wor1 + wor2  ) ;
			wor3List = append(wor3List, wor2 + wor1  ) ;
			wor3List = append(wor3List, wor1 + "s" + wor2  ) ;
			wor3List = append(wor3List, wor2 + "s" + wor1  ) ;			
		}	
	}	
	newWordList:= []string{}
	for _, wor1 := range wor3List {
		word1 := checkTheWord( wor1 ) ;
		if word1 == "" { continue }					
		wordCod:= seqCode( word1 )	
		_, ixT:= lookForWordInUniqueAlpha( wordCod)	
		if (ixT >= 0) { newWordList = append(newWordList, strings.ToLower(word1))	}
	}	
	return newWordList
	
} // end of add_wordCombinations  


//--------------------------------------------------------
func getRowIndexFromWordIndex( wordA00 []string, swComb bool) ( string, string, []int) {
	
	var maxNumRow = 100 
	
	listRowIndices := make([]int,0, maxNumRow)
	listIxRR := make([]int,0, maxNumRow)
	
	var listWords, listLemma string
	//-----------------
	for _, wor1 := range wordA00 {
		word1 := checkTheWord( wor1 ) ;
		if word1 == "" { continue }					
		wordCod:= seqCode( word1 )	
		
		ixF, ixT:= lookForWordInUniqueAlpha( wordCod)	
		if (ixT < 0) { 
			if swComb == false {
				listWords += " " +  word1
				listLemma += word1 + "|||\n" 
			}
			continue 
		}
		ixWord:= -1 	
		for ix:= ixF; ix <= ixT; ix++ {
			xWordA :=  uniqueWordByAlpha[ix] 			
			if xWordA.uWordSeq != wordCod { continue } // get only the required word (might be several entries of the same word) and then the list of lemmas of this word 
			ixWord = xWordA.uIxUnW			
			if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}	
			if ixWord < 0 { continue }
			listRowIndices = getRowIndicesFromIxFreqWord(ixWord, maxNumRow)
			if len(listRowIndices) < 1 {continue} 
			listIxRR = append(listIxRR, listRowIndices...)
		} // end for ix 
		//------------
		if ixWord < 0 { 
			if swComb == false {
				listWords += " " +  word1
				listLemma += word1 + "|||\n" 
			}
			continue 
		} 
		xWordF  := uniqueWordByFreq[ixWord]   
		listWords += " " +  xWordF.uWord2
		//--------
		for z:=0; z < len(xWordF.uLemmaL); z++  {
			ixL1:= xWordF.uIxLemmaL[z]
			LeS := lemmaSlice[ixL1]
			newL:= xWordF.uLemmaL[z]
			if newL != LeS.leLemma { continue}  // error 			
			if len(newL) > 1 { if newL[0:1] == "?" { newL = ""} }
			newT:= LeS.leTran
			newP:= xWordF.uPara[z]			
			if newP == "" { newP = newL}
			if z == 0 {	listLemma += xWordF.uWord2} 
			listLemma += "|" + newP + "|" + newT + "\n" 
		} // end for z	
	} // end of range wordA00		
	
	return listWords, listLemma, listIxRR	
		
} // end of getRowIndexFromWordIndex

//----------------------------------------

func PROVAbind_go_passToJs_thisWordRowList( aWord string,  maxNumRow int, js_function string) {  
	
	//  lista tutte le frasi che contengono le parole con lemma della parola cercata 
	//fmt.Println("bind_go_passToJs_thisWordLemmaWordRowList() 1  aWord=", aWord )
  
	var outS1 string;
	
	//---------------------------------------
	var listWords_str string;
	var listLemma_str string
	var listIxRR []int
	
	listWords_str, listLemma_str, listIxRR = getRowIndexFromWordIndex( []string{aWord}, false )
	fmt.Println("listWords_str=", 	listWords_str)
	fmt.Println("listLemma_str=", 	listLemma_str)
	
	sort.Ints(listIxRR) 
	
	listWords_pref := ""
	listWords := ""
	hd_tr := ""
	
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
	
	if (listWords_pref != "") {
		listWords += PREF_MARKER + listWords_pref 
	}
	header:= "<HEADER>\n" + "<WORD>" + aWord + ",L:" + strings.TrimSpace(listWords) + "</WORD>"
	//header:= "<HEADER>\n" + "<WORD>" + aWord + "</WORD>"
	header += hd_tr   // 14giugno
	header += "</HEADER> \n"
	
	go_exec_js_function( js_function, header + outS1 ); 	
				
} // end of PROVAbind_go_passToJs_thisWordRowList

//-----------------------------------------------------------
//--------------------------------------------------------

func bind_go_passToJs_thisWordRowList( aWord string,  maxNumRow int, js_function string) {  
	
	//  lista tutte le frasi che contengono le parole con lemma della parola cercata 
	//fmt.Println("bind_go_passToJs_thisWordLemmaWordRowList() 1  aWord=", aWord )
  
	var outS1 string;
	
	//---------------------------------------

	var xWordA, xWordF wordIxStruct;  		
	wordCod:= seqCode( aWord )		
	
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
		
		if xWordA.uWordSeq != wordCod { continue }            // get only the required word (might be several entries of the same word) and then the list of lemmas of this word 
		
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
	listWords_pref:=""
	xWordF  = uniqueWordByFreq[ixWord]   
	listWords += " " +  xWordF.uWord2
	/***
	leExample   string  
	ls_lemma_ix_stellen  int	
	ls_lemma_stellen     string
	ls_pref_ein          string
	ls_pref_tran         string 
	ls_lemma_einStellenList []int 
	***/
	newL2 :=""
	//-------------
	for z:=0; z < len(xWordF.uLemmaL); z++  {
		ixL1:= xWordF.uIxLemmaL[z]
		LeS := lemmaSlice[ixL1]
		newL:= xWordF.uLemmaL[z]
		if newL != LeS.leLemma {
			continue;  // error 
		}
		if LeS.ls_pref_ein == "" { 
			newL2=""
		} else {
			listWords_pref += " " + LeS.ls_pref_ein
			newL2= " = ??anto3 " + LeS.ls_pref_ein +"(" + LeS.ls_pref_tran+ ")" + " + " +  LeS.ls_lemma_stellen 
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
			
			hd_tr += " :lemma="  + newP + newL2 + " :tran=" + newT  + " :wordsInLemma=" +	xWordF.uWord2 + " "
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
	
	if (listWords_pref != "") {
		listWords += PREF_MARKER + listWords_pref 
	}
	header:= "<HEADER>\n" + "<WORD>" + aWord + ",L:" + strings.TrimSpace(listWords) + "</WORD>"
	//header:= "<HEADER>\n" + "<WORD>" + aWord + "</WORD>"
	header += hd_tr   // 14giugno
	header += "</HEADER> \n"
	
	go_exec_js_function( js_function, header + outS1 ); 	
				
} // end of bind_go_passToJs_thisWordRowList

//---------------------------------------------------------------------------

func bind_go_passToJs_someWordsRowList( aWordList1 string, aWordList2 string, maxNumRow int, js_function string) {  
	
	//  lista tutte le frasi che contengono le parole con lemma della parola cercata 
	//fmt.Println("bind_go_passToJs_thisWordLemmaWordRowList() 1  aWord=", aWord )
	swPrt:=false
	var maxNumRow2 = maxNumRow * 10; 
	var outS1 string;
	var wordA1 []string
	var wordA2 []string
	
	// se la seconda lista è vuota, sposta la prima lista sulla seconda e svuota la prima 
	//    le righe devono contenere una parole della lista2 e se presente anche una parola della lista1	
	if aWordList2 == "" {
		aWordList2 = aWordList1
		aWordList1 = ""
	}
	sw1:= (aWordList1 != "")
	sw2:= (aWordList2 != "") 
	if sw1 {wordA1 = regexp.MustCompile(separWord).Split(aWordList1, -1) } // split row into words 
	if sw2 {wordA2 = regexp.MustCompile(separWord).Split(aWordList2, -1) } // split row into words 
	
	wordA3:= []string{}
	
	/**
	if ( ( sw1 == false) && (len(wordA2) == 1) ) {
		bind_go_passToJs_thisWordRowList( wordA2[0], maxNumRow, js_function)  
		return
	}
	**/
		
	/*
	le parole in word1 sono in or: è sufficiente che una di queste sia presente  per estrarre la riga ==> estrae tutte le righe di tutte le parole 
	le parole in word2 sono in or: è sufficiente che una di queste sia presente  per estrarre la riga ==> estrae tutte le righe di tutte le parole 

	servono le righe che contengono almeno una parola della lista 1 ed almento 1 della lista2 
			==>  tutte le righe estratte nella prima lista che si trovano anche nella seconda lista
	*/
	
	listIxRR_L1    := make([]int,0, maxNumRow2)  // indici delle righe che contengono una parola della lista1
	listIxRR_L2    := make([]int,0, maxNumRow2)  // indici delle righe che contengono una parole della lista2
	listIxRR_L3    := make([]int,0, maxNumRow2)  // indici delle righe che contengono una parole della lista3
	listIxRR       := make([]int,0, maxNumRow2)  // indici delle righe che contengono una parola della lista1 e della lista2
	
	var listWords_str_L1 string; 
	var listWords_str_L2 string; 
	var listWords_str_L3 string; 
	var listLemmas_str_L1 string	
	var listLemmas_str_L2 string	
	var listLemmas_str_L3 string	
	//------------------
	// per ogni parola della lista1 estrae gli indici alle righe  
	if sw1 {
		listWords_str_L1, listLemmas_str_L1, listIxRR_L1 = getRowIndexFromWordIndex( wordA1 ,false)
		if swPrt { 
			fmt.Println("le parole listWords_str_L1 = ", listWords_str_L1)
			fmt.Println("le parole ", aWordList1, " si trovano in ", len(listIxRR_L1), " righe")
		}		
	} // end sw1 
	//-------------
	// per ogni parola della lista2 estrae gli indici alle righe  
	listWords_str_L2, listLemmas_str_L2, listIxRR_L2 = getRowIndexFromWordIndex( wordA2 , false)
	if swPrt {
		fmt.Println("le parole listWords_str_L2 = ", listWords_str_L2)	
		fmt.Println("le parole ", aWordList2, " si trovano in ", len(listIxRR_L2), " righe") 
	}	
	//-----------
	if sw1 {
		wordA3 = add_wordCombinations(wordA1, wordA2)	
		if len(wordA3) > 0 {
			listWords_str_L3, listLemmas_str_L3, listIxRR_L3 = getRowIndexFromWordIndex( wordA3 ,true)
			if swPrt {
				fmt.Println("sono state ottenute ", len(wordA3) , " parole combinando le parole di lista1 e lista2", 
					"\n\tqueste parole si trovano in ", len( listIxRR_L3 ), " righe")   
			}
		} 	
	}
	//---------------------
	if sw1 {	
		for _, ind1:= range listIxRR_L2 {
			// copia gli indici che si trovano anche nella lista1 
			//   questo significa che la riga puntata dall'indice, contiene almeno una parola della lista1 ed almeno una della lista2
			if slices.Contains(listIxRR_L1, ind1) { listIxRR = append(listIxRR, ind1 ) }
		}  
	} else {
		listIxRR = make([]int, len(listIxRR_L2), maxNumRow2)
		copy(listIxRR, listIxRR_L2)
	}
	if  len( listIxRR_L3 ) > 0 { listIxRR = append(listIxRR, listIxRR_L3...) }
	//----------------
	if swPrt { 
		fmt.Println( len(listIxRR),  "sono le righe che contengono almeno una parola in ", aWordList1, " ed almeno una parola in ", aWordList2) 
		if len(listIxRR) > maxNumRow { fmt.Println( "stampate soltanto le prime ", maxNumRow) }
	}
	
	//---------------------------------------
	
	sort.Ints(listIxRR) 
		
	preIxRR_2:= 999999999 
	ixRR_2:=0
	ixRR  :=0
	
	nOut:=0
	new_rIdRow :=""
	pre_rowX := ""
	
	for n1:= 0; n1 < len(listIxRR); n1++  {
		if n1 >= maxNumRow { break} 
		ixRR_2 = listIxRR[n1]
		if (ixRR_2 == preIxRR_2) { continue;} 
		preIxRR_2 = ixRR_2; 
		ixRR = ixRR_2 % 100000; 
			
		if ixRR >= numberOfRows { continue;} // actually there  must be some error here 		
		rline := inputTextRowSlice[ixRR]
		rowX := cleanRow(rline.rRow1)	
			
		if ((rowX =="") || (rowX == LAST_WORD) || (rowX == pre_rowX)) { 
			continue 
		}	
		pre_rowX = rowX 	
		
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
	//-------------------------
	header:= "<HEADER>\n" + "<WORD>"
	if (sw1 && sw2) { 
		header += aWordList1 + " " + aWordList2  + " " + listWords_str_L3 + "</WORD>\n"	
		header += "some:" + listLemmas_str_L1 + " + \n" + listLemmas_str_L2  
		if listLemmas_str_L3 != "" {header += "\n<hr>\n" + listLemmas_str_L3 }
	} else 			{ 
		header += aWordList2 + "</WORD>\n" 
		header += "some:" + listLemmas_str_L2  
	}
	header += "</HEADER> \n"
	
	go_exec_js_function( js_function, header + outS1 ); 		
				
} // end of bind_go_passToJs_someWordsRowList

//------------------------------------------------------

//----------------------------------------------
/**
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
	wordCod:= seqCode( aWord )		
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
		
		if xWordF.uWordSeq != wordCod { continue }            // get only the required word (might be several entries of the same word) and then the list of lemmas of this word 
		
		***
			if xWordF.uWord2 != aWord { continue; }
			if swOnlyThisWordRows {
				if xWordF.uWordSeq != wordCod { continue }            // get only the required word 
		}
		***
		ixWord := xWordF.uIxUnW			
		if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}		
		xWordF     = uniqueWordByFreq[ixWord] 
		// get the list of lemma of this word 
		for z:=0; z < len(xWordF.uLemmaL); z++  {
			lemmaList1 = append(lemmaList1, seqCode( xWordF.uLemmaL[z]) ) 
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
			if lemma_word_ix[k].lw_lemmaSeq < lemma1Cod { break }
			fromIx = k
		}
		for k:= fromIx; k < len(lemma_word_ix); k++ {
			if lemma_word_ix[k].lw_lemmaSeq != lemma1Cod { continue	}   	//  se non trova cerca lemma1Cod fino al punto.
			wL2 = lemma_word_ix[k]	
			if swOnlyThisWordRows { 
				if aWord != wL2.lw_word { continue;}                // get only the required word  
				**
				if aWord == "schrift" {
					fmt.Println(" antocontAnto3 lemma ", wL2) 
					contaAnto3++
				}
				***
			}
			wL2.lw_word = seqCode( wL2.lw_word )
			lemWordList2 = append( lemWordList2, wL2 )
		} 		
	}	
	//----
	//fmt.Println("anto lemma contaAnto3=",	contaAnto3) 
	
	//fmt.Println("bind_go_passToJs_thisWordLemmaWordRowList() 9 call  fun_wordListToRowList_head()")
	
	fun_wordListToRowList_head( aWord, lemWordList2, maxNumRow, js_function) 	
				
} // end of OLDbind_go_passToJs_thisWordRowList
**/
//-----------------------------------------------------------
//-----------------------------------------------------------

func fun_wordListToRowList_head(aWord string, lemmaList []lemmaWordStruct, maxNumRow int, js_function string) {
	//swAnto:= (aWord=="schrift") 
	//if swAnto {  fmt.Println("\nfun_wordListToRowList_head(aWord=", aWord, " lemmaList=", lemmaList)	}
	
	numberOfRows = len(inputTextRowSlice)
	
	maxNumRow0 := maxNumRow;	
	listWords:= ""
	listWords_pref:=""
	//hd_tr := "<TABLE>\n"	
	hd_tr := ""; 

	preL :=""	
	preW :=""
	
	// line2:=""   // 14giugno
		
	listIxRR := make([]int,0, 110)
	
	//-----------------
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
			newPref:=""
			if (LeS.ls_pref_ein == "") { 
				newPref = ""
			} else {
				listWords_pref += LeS.ls_pref_ein
				newPref = " = " + LeS.ls_pref_ein +"(" + LeS.ls_pref_tran+ ")" + " + " +  LeS.ls_lemma_stellen 
			}
			if newL != lemmaX.lw_lemma2 { continue }   //     seqCode   ( sto cercando frasi doppie 
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
				hd_tr += " :lemma="  + newL + newPref  + " :tran=" + newT  + " :wordsInLemma=" +	xWordF.uWord2 + " " // 14giugno 
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
	//numMioRow:=0
	
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
	
	
	if listWords_pref != "" {
		listWords += PREF_MARKER + listWords_pref
	}
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
		
		//pLemma:=  lemmaQuestionMark_remove( xWordF.uLemmaL )
		
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

//-------------------------------------------------

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


//------------------------------------
