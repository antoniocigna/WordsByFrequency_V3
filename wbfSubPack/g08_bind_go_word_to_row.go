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
		word0 := g12_checkTheWord( wor1 ) 			
		if word0 == "" { continue }	
		word2   := stdCode( word0 )    // elimina umlaut ed eszet  
		wordCod := word2 
		
		_, ixT:= lookForWordInUniqueAlpha( wordCod,0)	
		if (ixT >= 0) { newWordList = append(newWordList, strings.ToLower(word2))	}
	}	
	return newWordList
	
} // end of add_wordCombinations  

//----------------------------------------------------

func g08_getRowIndexFromWordIndex( wordA00 []string, swComb bool, maxNumRow int) ( string, string, []int) {
	
	listRowIndices := make([]int,0, maxNumRow)
	listIxRR := make([]int,0, maxNumRow)
	
	var listWords, listLemma string
	
	var xWordAlpha wordUnAlphaStruct
	//------------------------------
	for _, wor1 := range wordA00 {
		word0 := g12_checkTheWord( wor1 ) 			
		if word0 == "" { continue }	
		word2   := stdCode( word0 )    // elimina umlaut ed eszet  
		wordCod := word2
			
		ixF, ixT:= lookForWordInUniqueAlpha( wordCod,0)	
		if (ixT < 0) { 
			if swComb == false {
				listWords += " " +  word2
				listLemma += word2 + "|||\n" 
			}
			continue 
		}
		ixWord:= -1 	
		for ix:= ixF; ix <= ixT; ix++ {
			xWordAlpha =  uniqueWordByAlpha[ix] 			
			if xWordAlpha.uWord2 != wordCod { continue } // get only the required word (might be several entries of the same word) and then the list of lemmas of this word 
			ixWord = xWordAlpha.uIxUnW_al			
			if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}	
			if ixWord < 0 { continue }
			listRowIndices = g07_getRowIndicesFromIxAlphaWord(ixWord, maxNumRow) 
			if len(listRowIndices) < 1 {continue} 
			listIxRR = append(listIxRR, listRowIndices...)
		} // end for ix 
		//------------
		if ixWord < 0 { 
			if swComb == false {
				listWords += " " +  word2
				listLemma += word2 + "|||\n" 
			}
			continue 
		} 
		//------------------
		listWords += " " + xWordAlpha.uWord2
		//--------
		for z:=0; z < len(xWordAlpha.uLemmaL); z++  {
			ixL1:= xWordAlpha.uIxLemmaL[z]
			LeS := lemmaSlice[ixL1]
			newL:= xWordAlpha.uLemmaL[z]
			if newL != LeS.leLemma { continue}  // error 			
			if len(newL) > 1 { if newL[0:1] == "?" { newL = ""} }
			newT:= LeS.leTran
			//newP:= xWordAlpha.uPara[z]
			newP:= LeS.lePara	
			if newP == "" { newP = newL}
			if z == 0 {	listLemma += xWordAlpha.uWord2} 
			listLemma += "|" + newP + "|" + newT + "\n" 
		} // end for z	
	} // end of range wordA00		
	
	return listWords, listLemma, listIxRR	
		
} // end of g08_getRowIndexFromWordIndex


//--------------------------------------------------------

func g08_bind_go_passToJs_thisWordRowList( aWord string,  maxNumRow int, js_function string) {  
	
	//  lista tutte le frasi che contengono le parole con lemma della parola cercata 
	//fmt.Println("bind_go_passToJs_thisWordLemmaWordRowList() 1  aWord=", aWord )
  
	var outS1 string;
	
	//---------------------------------------

	var xWordA, xWordF wordUnAlphaStruct;  		
	wordCod:= aWord 		
	
	ixF, ixT:= lookForWordInUniqueAlpha( wordCod,0)	
	
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
		
		if xWordA.uWord2 != wordCod { continue }            // get only the required word (might be several entries of the same word) and then the list of lemmas of this word 
		
		ixWord = xWordA.uIxUnW_al			
		if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}	
		
		listRowIndices = g07_getRowIndicesFromIxAlphaWord(ixWord, maxNumRow)

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
	xWordF  = uniqueWordByAlpha[ixWord]   
	listWords += " " +  xWordF.uWord2
	
	newL2 :=""
	//-------------
	for z:=0; z < len(xWordF.uLemmaL); z++  {
		ixL1:= xWordF.uIxLemmaL[z]
		LeS := lemmaSlice[ixL1]
		newL:= xWordF.uLemmaL[z]
		if newL != LeS.leLemma {
			continue;  // error 
		}
		
		newT:= LeS.leTran
		
		newP:= LeS.lePara
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
		
		if rline.rIxGroup < 0 { 
			new_rIdRow = "- " + strconv.Itoa( rline.rIxBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		} else {
			new_rIdRow = lista_gruppiSelectRow[ rline.rIxGroup ].rG_group + " " + strconv.Itoa( rline.rIxBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
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
func g08_getWordList(aWordList0 string, maxNumList int) []string {
	
	wordA0 := regexp.MustCompile(separWord).Split(aWordList0, -1)  // split row into words 
	
	wordAlist := make([]string,0, 2*maxNumList) 
	for _, wor00 := range wordA0 {
		if len(wordAlist) >= maxNumList { return wordAlist }
		if strings.Index(wor00,"-") < 0 { 
			wordAlist = append(wordAlist, wor00)
			continue 
		}
		_, wordPrefixIndexList, wordSuffixIndexList := g04_get_word_row_list( maxNumList, wor00)   // get list of indices 
		for _, ixWord:= range wordPrefixIndexList {  if ixWord >= 0 { wordAlist = append(wordAlist, uniqueWordByAlpha[ixWord].uWord2 )  } }
		for _, ixWord:= range wordSuffixIndexList {  if ixWord >= 0 { wordAlist = append(wordAlist, uniqueWordByAlpha[ixWord].uWord2 )  } } 		
	} 	
	if len(wordAlist) <= maxNumList { return wordAlist }
	
	return wordAlist[0:maxNumList]
	
} // end of g08_getWordList

//-----------------------------------------------
func g08_bind_go_passToJs_someWordsRowList( aWordList1 string, aWordList2 string, maxNumRow int, js_function string) {  
	
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
	
	if sw1 {wordA1 = g08_getWordList(aWordList1, maxNumRow) } // split row into words  ( also  prefix and suffix as input )
	if sw2 {wordA2 = g08_getWordList(aWordList2, maxNumRow) } // split row into words  ( also  prefix and suffix as input )	
	
	//fmt.Println("_someWordsRowList wordA1=", strings.Join(wordA1," "), " sw1=", sw1, " sw2=", sw2,  " maxNumRow=", maxNumRow   )
	
	wordA3:= []string{}
		
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
		listWords_str_L1, listLemmas_str_L1, listIxRR_L1 = g08_getRowIndexFromWordIndex( wordA1 ,false, maxNumRow2)
		if swPrt { 
			fmt.Println("le parole listWords_str_L1 = ", listWords_str_L1)
			fmt.Println("le parole ", aWordList1, " si trovano in ", len(listIxRR_L1), " righe")
		}		
	} // end sw1 
	//-------------
	// per ogni parola della lista2 estrae gli indici alle righe  
	listWords_str_L2, listLemmas_str_L2, listIxRR_L2 = g08_getRowIndexFromWordIndex( wordA2 , false, maxNumRow2)
	if swPrt {
		fmt.Println("le parole listWords_str_L2 = ", listWords_str_L2)	
		fmt.Println("le parole ", aWordList2, " si trovano in ", len(listIxRR_L2), " righe") 
	}	
	//-----------
	if sw1 {
		wordA3 = add_wordCombinations(wordA1, wordA2)	
		if len(wordA3) > 0 {
			listWords_str_L3, listLemmas_str_L3, listIxRR_L3 = g08_getRowIndexFromWordIndex( wordA3 ,true, maxNumRow2)
			if swPrt {
				if len( listIxRR_L3 ) > 0 {	
					fmt.Println("sono state ottenute ", len(wordA3) , " parole combinando le parole di lista1 e lista2 (.es. A e B possono formare AB e BA, es. ein e steigen formano einsteigen e steigenein)", 
						"\n\tqueste parole si trovano in ", len( listIxRR_L3 ), " righe")   
				}	
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
		fmt.Println( len(listIxRR),  "sono le righe che contengono almeno una parola della lista (", aWordList1, ") ed almeno una della lista(", aWordList2, ")" ) 
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
		
		if rline.rIxGroup < 0 { 
			new_rIdRow = "- " + strconv.Itoa( rline.rIxBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		} else {
			new_rIdRow = lista_gruppiSelectRow[ rline.rIxGroup ].rG_group + " " + strconv.Itoa( rline.rIxBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
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
		header += strings.Join(wordA1, " ") + " " + strings.Join(wordA2, " ")  + " " + listWords_str_L3 + "</WORD>\n"	
		header += "some:" + listLemmas_str_L1 + " + \n" + listLemmas_str_L2  
		if listLemmas_str_L3 != "" {header += "\n<hr>\n" + listLemmas_str_L3 }
	} else 			{ 
		header += strings.Join(wordA2, " ") + "</WORD>\n" 
		header += "some:" + listLemmas_str_L2  
	}
	header += "</HEADER> \n"
	
	go_exec_js_function( js_function, header + outS1 ); 		
				
} // end of bind_go_passToJs_someWordsRowList

//-----------------------------------------------------

func listStringLemmaSlice_Tran( xWordAlpha wordUnAlphaStruct) string {
	listS:=""
	for _, ixL1:= range xWordAlpha.uIxLemmaL { 
		listS += lemmaSlice[ixL1].leTran + wSep
	}		
	return listS
}

//------------------
/**

g00_2structure.go => //---
g00_2structure.go => type lemmaStruct struct {
g00_2structure.go => 	leLemma    string    
g00_2structure.go => 	leNumWords int 
g00_2structure.go => 	leFromIxLW  int             // limite inferiore range indici a wordLemmaPair (in seq. di lemma)   wordLemmaPair_lemmaWordSeq[] 
g00_2structure.go => 	leToIxLW    int             // limite superiore range indici a wordLemmaPair (in seq. di lemma)   wordLemmaPair_lemmaWordSeq[]   
g00_2structure.go => 	leUnWord_al_IxList []int    // indice delle parole unique che puntano a questo lemma        
g00_2structure.go => 	leTran      string 
g00_2structure.go => 	lePara      string  
g00_2structure.go => 	leExample   string  
g00_2structure.go => 	leNumPara   int	
g00_2structure.go => } 
g00_2structure.go => //-------------------------------
g00_2structure.go => //--
g00_2structure.go => type wordUnAlphaStruct struct {    // uniqueWordByAlpha
g00_2structure.go => 	uWord0    string	
g00_2structure.go =>     uWord2      string		
g00_2structure.go => 	uIxUnW_al   int            // index of this word in the uniqueWordByAlpha 	
g00_2structure.go => 	uIxUnW_fr   int            // index of this word in the uniqueWordByFreq	
g00_2structure.go => 	uTotRow     int 
g00_2structure.go => 	uTotExtrRow int
g00_2structure.go => 	uIxFromWord_al int          // index of this word in the wordSliceAlpha (first occurrence, last = uIxFromWordAl + uTotRow-1	
g00_2structure.go =>     //uIxWordFreq int            // index of this word in the wordSliceFreq	
g00_2structure.go => 	uSwSelRowG  int
g00_2structure.go => 	uSwSelRowR   int  
g00_2structure.go => 	uLearnedYN   string         // y n ( ie.yes,I learned / not yet  
g00_2structure.go => 	//uKnow_yes_ctr int 
g00_2structure.go => 	//uKnow_no_ctr  int         // a value > 0  means that this is a word that I don't know, ie. it's to be learned   
g00_2structure.go => 	uIxLemmaL  []int  
g00_2structure.go => 	uLemmaL    []string       // list of lemma 	
g00_2structure.go => 	//uPara      []string  
g00_2structure.go => 	//uExample   []string  
g00_2structure.go => }	

****/
func g08_bind_go_passToJs_rowWordList(numIdOut string, ixRR int, js_function string) {
	//  lista di tutte le parole di una frase	
	
	if ixRR >= len(inputTextRowSlice) { ixRR = len(inputTextRowSlice) - 1 }
	
	rowX := inputTextRowSlice[ixRR]

	outS1:= numIdOut + "," + strconv.Itoa(ixRR) + "," + strconv.Itoa(rowX.rNumWords) +"," + strconv.Itoa( len(rowX.rListIxUnF)) + endOfLine
	
	for w:=0; w < len(rowX.rListIxUnF); w++ {  
		if rowX.rListFreq[w] < 1 {continue}     // the entry is allocated, but unused  
		ixWord := rowX.rListIxUnF[w] 
		if ixWord < 0 { continue}
		ixAl:= uniqueWordByFreq[ixWord].fuIxUnW_al 
		xWordAlpha := uniqueWordByAlpha[ ixAl ]	
		
		row11 := xWordAlpha.uWord2 + "," + strconv.Itoa(xWordAlpha.uIxUnW_fr) + "," + 
			strconv.Itoa(xWordAlpha.uTotRow)  + ";" ; 
		for _,ixLe:= range xWordAlpha.uIxLemmaL {
			lem := lemmaSlice[ixLe]  
			row11 += "[" + lem.leLemmaOr + ";" + lem.lePara + ";" + lem.leTran + ";" + lem.leExample + ";" + strconv.Itoa(lem.leNumWords) + "] " ; 
		}
		outS1 += row11 + endOfLine; 		
	}	

	go_exec_js_function( js_function, outS1 ); 				

} // end of bind_go_passToJs_rowWordList 

//------------------

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
