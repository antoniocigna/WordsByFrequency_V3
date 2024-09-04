package wbfSubPack

	import (
		"fmt"
		"os"	
		"strconv"
	)
//--------------------------------------------------------

func bind_go_passToJs_getIxRowFromGroup( rowGrIndex int,   html_rowGroup_beginNum int, html_rowGroup_numRows int, js_function string)  {

	fmt.Println( green("bind_go_passToJs_getIxRowFromGroup"), 
		"()  rowGrIndex=", rowGrIndex, ",  html_rowGroup_beginNum=", html_rowGroup_beginNum, ",  html_rowGroup_numRows=", html_rowGroup_numRows	) 
	
	if rowGrIndex < 0 { return }
	
	/**
	if (sw_list_Word_if_in_ExtrRow) { 		
			fmt.Println("on the row list button \"inpBegRow\" or \"maxNumRow\" values has been changed, but since only extracted rows wanted, that causes a rebuild of wordlist data")			
			build_and_elab_word_list()
	}
	**/
	
	//gRix := rowLineIxList[ rowGrIndex ] 
	
	rG := lista_gruppiSelectRow[ rowGrIndex ]	


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
	//last_html_rowGroup_numRows  = 1 + last_rS_toIxRow - ixRowBeg - (html_rowGroup_beginNum - 1)
	
	//last_word_fromWord    = 	
	//last_word_numWords    = 
	//last_sel_extrRow      = 
	
	write_lastValueSets()
	
	outS1:= fmt.Sprintf( "inp,%d,%d,%d,gr,%d,%s,%d,%d,ixr,%d,%d, %s",
				rowGrIndex, html_rowGroup_beginNum, html_rowGroup_numRows,
				rG.rG_ixSelGrOption, 
				rG.rG_group, 
				rG.rG_firstIxRowOfGr, 
				rG.rG_lastIxRowOfGr,
				ixRowBeg, ixRowEnd, 	
				inputTextRowSlice[  rG.rG_firstIxRowOfGr ].rRow1   )
	
	go_exec_js_function( js_function, outS1 )	
	
} // end of bind_go_passToJs_updateRowGroup



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