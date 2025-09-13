package wbfSubPack

	import (
		"fmt"
		"os"	
		"strconv"
		"sort"
	)
//--------------------------------------------------------

func g06_bind_go_passToJs_getIxRowFromGroup( rowGrIndex int,   html_rowGroup_beginNum int, html_rowGroup_numRows int, js_function string)  {

	//fmt.Println("func ", green("g06_bind_go_passToJs_getIxRowFromGroup"), "(rowGrIndex=", rowGrIndex, ", html_rowGroup_beginNum=",html_rowGroup_beginNum, 
	//	", html_rowGroup_numRows=", html_rowGroup_numRows ); 
		
	if rowGrIndex < 0 { return }
	
	rG := lista_gruppiSelectRow[ rowGrIndex ]	
	
	

	//fmt.Println("    g06_bind_go_passToJs_getIxRowFromGroup",  " lista_gruppiSelectRow[", rowGrIndex, "] = ", rG) 
	

	ixRowBeg := html_rowGroup_beginNum + rG.rG_firstIxRowOfGr - 1 	 
	ixRowEnd := ixRowBeg + html_rowGroup_numRows - 1
	if ixRowEnd > rG.rG_lastIxRowOfGr { 
		ixRowEnd = rG.rG_lastIxRowOfGr  
		html_rowGroup_numRows = 1 + ixRowEnd - ixRowBeg
	}

	last_rG_ixSelGrOption = rG.rG_ixSelGrOption	
	last_rG_group         = rG.rG_group	 
	last_rG_firstIxRowOfGr= rG.rG_firstIxRowOfGr
	last_rG_lastIxRowOfGr = rG.rG_lastIxRowOfGr 	   	
	last_ixRowBeg         = ixRowBeg	 
	last_ixRowEnd         = ixRowEnd	
	last_html_rowGroup_index_gr = rowGrIndex	
	last_html_rowGroup_beginNum = html_rowGroup_beginNum	// from the beginning of the group ( starting from 1 )
	last_html_rowGroup_numRows  = html_rowGroup_numRows	
	
	write_lastValueSets()
	
	outS1:= fmt.Sprintf( "inp,%d,%d,%d,gr,%d,%s,%d,%d,ixr,%d,%d, %s",
				rowGrIndex, html_rowGroup_beginNum, html_rowGroup_numRows,
				rG.rG_ixSelGrOption, 
				rG.rG_group, 
				rG.rG_firstIxRowOfGr, 
				rG.rG_lastIxRowOfGr,
				ixRowBeg, ixRowEnd, 	
				inputTextRowSlice[  rG.rG_firstIxRowOfGr ].rRow1   )
	
	//fmt.Println("    g06_bind_go_passToJs_getIxRowFromGroup"  , " run write_lastValueSets ", "\n\t--> ", js_function , "(", outS1 ); 
	
	go_exec_js_function( js_function, outS1 )	
	
} // end of bind_go_passToJs_updateRowGroup
//------------------------------------------
/***
func  g06_bind_go_passToJs_getIxRowFromGroup (rowGrIndex= 1 , html_rowGroup_beginNum= 1 , html_rowGroup_numRows= 100
    g06_bind_go_passToJs_getIxRowFromGroup  run write_lastValueSets
        -->  js_go_gotIxRowFromGroup ( inp,1,1,7,gr,1,2,7,13,ixr,7,13, file: prova2.csv
func  g06_bind_go_passToJs_getIxRowFromGroup (rowGrIndex= 5 , html_rowGroup_beginNum= 1 , html_rowGroup_numRows= 100
    g06_bind_go_passToJs_getIxRowFromGroup  run write_lastValueSets
        -->  js_go_gotIxRowFromGroup ( inp,5,1,7,gr,0,1,0,6,ixr,0,6, file: prova.csv

   1° rowGrIndex OK;  1,7 (3°,4°) sono quelli precedenti da html;  5° .rgGroup ERRORE e seguenti errati 
**/
//-----------------------------
func write_lastValueSets() {
	
	//fmt.Println( green("write_lastValueSets") ) 
	
	outS1:= fmt.Sprint(
		"rG_ixSelGrOption="  + strconv.Itoa( last_rG_ixSelGrOption ) + ", " ,
		"rG_group="          + last_rG_group                         + ", " ,  
		"rG_firstIxRowOfGr=" + strconv.Itoa( last_rG_firstIxRowOfGr) + ", " ,     
		"rG_lastIxRowOfGr="  + strconv.Itoa( last_rG_lastIxRowOfGr ) + ", \n" ,      

		"html_rowGroup_index_gr=" + strconv.Itoa( last_html_rowGroup_index_gr ) + ", " ,
		"html_rowGroup_beginNum=" + strconv.Itoa( last_html_rowGroup_beginNum ) + ", " ,
		"html_rowGroup_numRows="  + strconv.Itoa( last_html_rowGroup_numRows  ) + ", \n" ,	
		
		"ixRowBeg="     + strconv.Itoa( last_ixRowBeg  ) + ", " , 
		"ixRowEnd="     + strconv.Itoa( last_ixRowEnd  ) + ", \n" , 
		
		"w_fromWord="             + strconv.Itoa( last_word_fromWord ) +", " ,	
		"w_numWords="             + strconv.Itoa( last_word_numWords ) +", \n" ,	

		"sel_extrRow="            + last_sel_extrRow )
	outS2:= []string{ outS1}	
	writeList(FOLDER_INPUT_OUTPUT  + string(os.PathSeparator) + FILE_last_mainpage_values2, outS2 )		
	
	//fmt.Println( green("write_lastValueSets"), " in file ",  FILE_last_mainpage_values2    ) 
	
} // end of 
//----------------------------------------------

//var range_gr_ixRow_seq = make([]int, 0, 0) 
//var range_gr_ixRow_pry = make([]int, 0, 0) 
//----------------------------------
func g06_setAvgWordFreqInRow() {
	/*	
	type rowPriorStruct struct {   // priority of rows in the text 	
			rP_numWords 	int    // number of words in row 
			rP_wordFreqAvg 	int	   // average of frequence of the words in the row		
			rP_index 		int    // index of row in inputTextRowSlice
			rP_ixGroup       int    // file number                          
	}
	*/
	
	
	//range_gr_ixRow_seq = make([]int, len(lista_gruppiSelectRow)  ) 
	//range_gr_ixRow_pry = make([]int, len(lista_gruppiSelectRow)  ) 	
	for ix1, rG:= range lista_gruppiSelectRow {
		//	range_gr_ixRow_seq[ix1] = rG.rG_firstIxRowOfGr		
		fmt.Println(  green("lista_gruppiSelectRow ["), ix1, "] = ", rG, ", rG.rG_firstIxRowOfGr=", rG.rG_firstIxRowOfGr  )	
	}
	
	//-------------
	rowPriorityList = make([]rowPriorStruct, 0, len(inputTextRowSlice) )
	
	lenT:= len(inputTextRowSlice)
	
	var oneP rowPriorStruct
	//range_gr_ixRow_seq := make([]int, 0, 100) 
	//range_gr_ixRow_pry := make([]int, 0, 100) 
	//pg:=-1; 
	
	//-----------------
	for ixRR := 0; ixRR < lenT; ixRR++  {
		rline := inputTextRowSlice[ixRR]
		avgWfreq:= 0;
		len1 := len( rline.rListFreq ) 
		if len1 > 0 {
			for _, fr:= range rline.rListFreq { 
				avgWfreq += fr
			}
			avgWfreq =  int( (float64( avgWfreq)) / (float64(len1) ) )
		}
		inputTextRowSlice[ixRR].rWordFreqAvg = avgWfreq
		
		oneP.rP_numWords 	= rline.rNumWords
		oneP.rP_wordFreqAvg = avgWfreq
		oneP.rP_index 		= ixRR
		oneP.rP_ixGroup     = rline.rIxGroup
		rowPriorityList   = append(rowPriorityList, oneP)
		
	} // end for ixRR
	//---------------------------------------
	sort.Slice(rowPriorityList, func(i, j int) bool {
		if rowPriorityList[i].rP_ixGroup != rowPriorityList[j].rP_ixGroup {
			return rowPriorityList[i].rP_ixGroup < rowPriorityList[j].rP_ixGroup        // indice gruppo:  ascending order
		} else {
			if rowPriorityList[i].rP_numWords != rowPriorityList[j].rP_numWords {
				return rowPriorityList[i].rP_numWords < rowPriorityList[j].rP_numWords      // number of words:  ascending order  ( prima le righe con meno parole
			} else {
				if rowPriorityList[i].rP_wordFreqAvg != rowPriorityList[j].rP_wordFreqAvg {
					return rowPriorityList[i].rP_wordFreqAvg > rowPriorityList[j].rP_wordFreqAvg //average of word frequency: descending descending order (prima le più frequenti) 
				} else {
					return rowPriorityList[i].rP_index < rowPriorityList[j].rP_index            //  index in the rows in the text: ascending order  
				}
			}
		}
	})
	//-----------
	fmt.Println( green("rowPriorityList "), len(rowPriorityList), " righe")
	//------------------------
	// set priority to the text row
	//pg:=-1
	for pp, oneP := range rowPriorityList {
		ixRR := oneP.rP_index
		if ixRR >= lenT { continue }   // error
		/**
		if oneP.rP_ixGroup > pg {
			pg = oneP.rP_ixGroup
			range_gr_ixRow_pry[pg] = pp 
		}	
		**/
		inputTextRowSlice[ixRR].rPriority = pp
		//if oneP.rP_ixGroup == 0 { fmt.Println( green("green "), pp, " priority ixRR=", pp, " ",inputTextRowSlice[ixRR].rRow1)  }
	} // end for pp 
	//-----------------------------
	//fmt.Println(green("range_gr_ixRow_seq = "), range_gr_ixRow_seq)
	//fmt.Println(green("range_gr_ixRow_pry = "), range_gr_ixRow_pry)
	
	
} // end of g06_setAvgWordFreqInRow		
//------------------------------


func g06_bind_go_passToJs_rowList(indexGroup int,  inpBegRow int, maxNumRow int, selFrasiParole12 int, 
				js_function1R string, js_function2W string, js_parm string, caller string) {
	// lista tutte le frasi richieste ( numero della prima frase, numero di frasi) 
	
	// 	selFrasiParole12: 1 = list chosen rows,  2 = list words in the chosen rows  3 = list toBeLearned words in the chosen rows 
	//                    4 = list chosen rows sorted by priority 
	
	swPriorSort := false
	if selFrasiParole12 == 4 {
		swPriorSort = true
		selFrasiParole12 = 1
	}
	
	if inpBegRow == 1 {inpBegRow=0} 
	
	swLIST_WORD := (selFrasiParole12 != 1)    
	swLIST_TOLEARN := (selFrasiParole12 == 3)  // only to be learned words
	/**
	fmt.Println( green("g08_bind_go_passToJs_rowList ")," indexGroup =", indexGroup, " inpBegRow=", inpBegRow, 
			"  maxNumRow=", maxNumRow, " selFrasiParole12=", selFrasiParole12,
			green(" swPriorSort="), swPriorSort, " swLIST_WORD=", swLIST_WORD)
	**/
		
	var outS1 string;
	numOut:=0 
	
	new_rIdRow :=""
	wordIxFrList:= make([]int, 0, 10000)    //  list of index  of uniqueWordByFreq  of the words in the row 
	
	//fmt.Println("bind_go_passToJs_rowList ixFromList=", ixFromList, " swLIST_WORD=", swLIST_WORD ,  " selFrasiParole12=",  selFrasiParole12, " len=",  len(inputTextRowSlice) ) 
	//list swLIST_WORD= true  selFrasiParole12= 2  len= 7	
	//----------------------------------------------------------------
	
	type rowPr2 struct {   // priority of rows in the text 			
		p2_index 		int    // index of row in inputTextRowSlice
		p2_prio         int    // priority                          
	}
	var onePr rowPr2
	
	listIxRowPrio:= make([]rowPr2, 0, (maxNumRow + 10) )

	//-----------------------------
	/**
	for ix1, rG:= range lista_gruppiSelectRow {	
		fmt.Println(  green("lista_gruppiSelectRow ["), ix1, "] = ", rG, ", rG.rG_firstIxRowOfGr=", rG.rG_firstIxRowOfGr  )	
	}
	**/
	//-------------
	ixStartGroup := 0; 
	ixEndGroup   := len(inputTextRowSlice) 
	if indexGroup     < len(lista_gruppiSelectRow) { ixStartGroup = lista_gruppiSelectRow[indexGroup].rG_firstIxRowOfGr }
	if (indexGroup+1) < len(lista_gruppiSelectRow) { ixEndGroup   = lista_gruppiSelectRow[(indexGroup+1)].rG_firstIxRowOfGr }
	//fmt.Println( "=ixStartGroup=", ixStartGroup, " ixEndGroup=",  ixEndGroup)
	//--------------------------------
	minIx  := ixStartGroup + inpBegRow;
	maxOut := minIx        + maxNumRow
	
	if maxOut >= ixEndGroup {  maxOut = ixEndGroup }
	
	//fmt.Println( green("listIxRowPrio:"), minIx, " - ", maxOut)
	//----------------------
	if swPriorSort {
		for pp:=minIx; pp < maxOut; pp++ {			
			if rowPriorityList[pp].rP_ixGroup != indexGroup { continue }	
				onePr.p2_index = rowPriorityList[pp].rP_index
				onePr.p2_prio  = pp	
				listIxRowPrio  = append(listIxRowPrio, onePr) 			
		} // end for pp
	} else {
		for ixRR := minIx; ixRR < maxOut; ixRR++ {
			rline := inputTextRowSlice[ixRR]	
			if rline.rIxGroup != indexGroup { continue }
				onePr.p2_index = ixRR
				onePr.p2_prio  = rline.rPriority	
				listIxRowPrio  = append(listIxRowPrio, onePr) 				
		} // end for ixRR		
	} 
	//--------------------------
	//fmt.Println( green("listIxRowPrio:"), minIx, " - ", maxOut, " len(listIxRowPrio)=", len(listIxRowPrio) )
	//-------------------------------
	
	for _, onePr:= range listIxRowPrio {
		ixRR:= onePr.p2_index
		
		//if pp < 10 {fmt.Println( red(" out "), " pp=", pp , " onePr=", onePr, " ixRR=", ixRR ) }
		
		rline := inputTextRowSlice[ixRR]
		
		rowX := cleanRow(rline.rRow1)	
		
		//fmt.Println("bind_go_passToJs_rowList rline.rIxGroup=", rline.rIxGroup," rline=", rline )
		
		if ((rowX =="") || (rowX == LAST_WORD)) { 
			//fmt.Println("   bind_go_pa... 1")
			continue 
		}		
		
		if rline.rIxGroup < 0 { 
			new_rIdRow = "- " + strconv.Itoa( rline.rIxBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		} else {
			new_rIdRow = lista_gruppiSelectRow[ rline.rIxGroup ].rG_group + " " + strconv.Itoa( rline.rIxBaseGroup ) + "(" + rline.rIdRow +  " " + strconv.Itoa(ixRR) 
		}
		if swLIST_WORD {  
			//fmt.Println("leggi riga=", rowX,  " rNumWords=", rline.rNumWords, " .rListIxUnF=", rline.rListIxUnF, " freq=", rline.rListFreq)
			wordIxFrList = append(wordIxFrList, rline.rListIxUnF...)
		} else {	
			//if (pp < 10) {fmt.Println(      red(" out2 "), " ixRR=", ixRR , " ", rowX ) }
			outS1 += "<br>" + strconv.Itoa( SEL_EXTR_ROW ) + "|" + new_rIdRow + "|" + strconv.Itoa( ixRR) + "|"   + rowX + "|" + rline.rTran1 +
				"|" + strconv.Itoa(rline.rIxGroup); 
		}	
	} // end for ixRR
	//------------------------------
	if swLIST_WORD {
		sort.Ints(wordIxFrList)
		wordIxFrList = append(wordIxFrList, -1)  // serve per scrivere l'ultimo dell'elenco 
		   
		//fmt.Println("bind_go_passToJs_rowList  wordIxFrList=", wordIxFrList)
		
		preW:=-1
		numRw:=-999
		outS1 = ""
		//var highestValue = string( highestValueByte ) + "end_of_list"	
		//fmt.Println("bind_go_passToJs_rowList len(uniqueWordByFreq)=",  len(uniqueWordByFreq), " last word=",  uniqueWordByFreq[ len(uniqueWordByFreq)-1 ].uWord2 ) 
		//----------------
		/**
		for ix1,WF := range(uniqueWordByFreq) { 
			fmt.Println(" tutti    uniqueWordByFreq[",ix1, "]=", WF.uWord2) 
		}
		**/
		//------------------------------------
		swElab:=false; 
		var xWordAlpha wordUnAlphaStruct  
		
		for _,ixWF := range(wordIxFrList) {
			//fmt.Println("bind_go_passToJs_rowList  ixWF =", ixWF,  "  preW=", preW, "   numRw=", numRw )
			
			if ixWF == preW {
				numRw++
				continue
			}  	
			swElab = false
			if ((numRw > 0) && (preW >= 0)) { 
				swElab = true
				ixAl := uniqueWordByFreq[preW].fuIxUnW_al 
				xWordAlpha= uniqueWordByAlpha[ixAl] 
				//fmt.Println("preW=", preW, " WORD ", cyan(xWordAlpha.uWord2), " xWordAlpha.uLearnedYN=", xWordAlpha.uLearnedYN, "  swLIST_TOLEARN=", swLIST_TOLEARN)
				if swLIST_TOLEARN {
					if xWordAlpha.uLearnedYN == LEARNED_YES { swElab = false }
				}
			}
			if swElab {	
				//fmt.Println("    uniqueWordByFreq[",preW, "]=",cyan( uniqueWordByFreq[preW].uWord2 ) , " numRw=", numRw)     	
				sw, rowW := word_to_row("", false, "anyRow", xWordAlpha, numRw)  
				if sw {	
					numOut++
					//if (numOut < from1) {continue}   // July7, 2025  
					outS1 += "<br>" + rowW 						
					//if numOut >= numWords { 
					//	break
					//}
				}
			}
			preW = ixWF	
			numRw = 1
		} // end for ixWf range
		//fmt.Println("bind_go_passToJs_rowList  ", numOut, " words  --> js function=",js_function2W , "\n\toutS1=", outS1 )   
		if len(outS1) == 0 { outS1 = "<br>" }
		outS1 = sortWordToRowByFreq(outS1) 
		go_exec_js_functionPlus( js_function2W, outS1 , js_parm, caller); 			
	} else {	
		if len(outS1) == 0 { outS1 = "<br>" }
		go_exec_js_functionPlus( js_function1R, outS1 , js_parm, caller); 	
	}
			
} // end of bind_go_passToJs_rowList

//-------------------------------