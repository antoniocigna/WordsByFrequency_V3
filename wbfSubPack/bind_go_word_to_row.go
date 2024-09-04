package wbfSubPack

	import (
		"fmt"
		"strings"		
		"strconv"
		"sort"
	)
//--------------------------------------------------------

func bind_go_passToJs_thisWordRowList( aWord string,  maxNumRow int, js_function string) {  
	
	//  lista tutte le frasi che contengono le parole con lemma della parola cercata 
	//fmt.Println("bind_go_passToJs_thisWordLemmaWordRowList() 1  aWord=", aWord )
  
	var outS1 string;
	
	//---------------------------------------

	var xWordA, xWordF wordIxStruct;  		
	wordCod:= newCode( aWord )		
	
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
		
		if xWordA.uWordCod != wordCod { continue }            // get only the required word (might be several entries of the same word) and then the list of lemmas of this word 
		
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

//-----------------------------------------------------------
//----------------------------------------------

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
	wordCod:= newCode( aWord )		
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
		
		if xWordF.uWordCod != wordCod { continue }            // get only the required word (might be several entries of the same word) and then the list of lemmas of this word 
		
		/**
			if xWordF.uWord2 != aWord { continue; }
			if swOnlyThisWordRows {
				if xWordF.uWordCod != wordCod { continue }            // get only the required word 
		}
		**/
		ixWord := xWordF.uIxUnW			
		if ixWord >= numberOfUniqueWords {ixWord = numberOfUniqueWords - 1;}		
		xWordF     = uniqueWordByFreq[ixWord] 
		// get the list of lemma of this word 
		for z:=0; z < len(xWordF.uLemmaL); z++  {
			lemmaList1 = append(lemmaList1, newCode( xWordF.uLemmaL[z]) ) 
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
			if lemma_word_ix[k].lw_lemmaCod < lemma1Cod { break }
			fromIx = k
		}
		for k:= fromIx; k < len(lemma_word_ix); k++ {
			if lemma_word_ix[k].lw_lemmaCod != lemma1Cod { continue	}   	//  se non trova cerca lemma1Cod fino al punto.
			wL2 = lemma_word_ix[k]	
			if swOnlyThisWordRows { 
				if aWord != wL2.lw_word { continue;}                // get only the required word  
				/**
				if aWord == "schrift" {
					fmt.Println(" antocontAnto3 lemma ", wL2) 
					contaAnto3++
				}
				***/
			}
			wL2.lw_word = newCode( wL2.lw_word )
			lemWordList2 = append( lemWordList2, wL2 )
		} 		
	}	
	//----
	//fmt.Println("anto lemma contaAnto3=",	contaAnto3) 
	
	//fmt.Println("bind_go_passToJs_thisWordLemmaWordRowList() 9 call  fun_wordListToRowList_head()")
	
	fun_wordListToRowList_head( aWord, lemWordList2, maxNumRow, js_function) 	
				
} // end of OLDbind_go_passToJs_thisWordRowList

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
			if newL != lemmaX.lw_lemma2 { continue }   //     newCode   ( sto cercando frasi doppie 
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
		 
		// ??anto2 uTranL
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
