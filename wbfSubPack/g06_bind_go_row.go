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
	var oneP rowPriorStruct
	
	rowPriorityList = make([]rowPriorStruct, 0, len(inputTextRowSlice) )
	lenT:= len(inputTextRowSlice)
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
		oneP.rP_ixGroup      = rline.rIxGroup
		rowPriorityList   = append(rowPriorityList, oneP)
		
	} // end for ixRR
	//---------------------------------------
	sort.Slice(rowPriorityList, func(i, j int) bool {
		if rowPriorityList[i].rP_numWords != rowPriorityList[j].rP_numWords {
			return rowPriorityList[i].rP_numWords < rowPriorityList[j].rP_numWords      // number of words:  ascending order  ( prima le righe con meno parole
		} else {
			if rowPriorityList[i].rP_wordFreqAvg != rowPriorityList[j].rP_wordFreqAvg {
				return rowPriorityList[i].rP_wordFreqAvg > rowPriorityList[j].rP_wordFreqAvg //average of word frequency: descending descending order (prima le più frequenti) 
			} else {
				return rowPriorityList[i].rP_index < rowPriorityList[j].rP_index            //  index in the rows in the text: ascending order  
			}
		}
	})
	//-----------
	fmt.Println( green("rowPriorityList "), len(rowPriorityList), " righe")
	//------------------------
	// set priority to the text row
	for pp, oneP := range rowPriorityList {
		ixRR := oneP.rP_index
		if ixRR >= lenT { continue }   // error
		inputTextRowSlice[ixRR].rPriority = pp
	}
	
} // end of g06_setAvgWordFreqInRow		
//------------------------------