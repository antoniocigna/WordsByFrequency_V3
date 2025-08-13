package wbfSubPack

	import (
		"fmt"
		"os"	
		"strconv"
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