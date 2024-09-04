package wbfSubPack

import (  
	"fmt"
    "strings"
	"strconv"
    //"sort"
)
//------------------------------------------------

func read_lastValueSets2() {
	bytesPerRow:= 10
    lineD := rowListFromFile( FOLDER_INPUT_OUTPUT, FILE_last_mainpage_values2, "last run values", "read_lastValueSets2", bytesPerRow)  
	
	fmt.Println("read_lastValueSets2 lineD=", lineD)
	
	var dat="";
	if sw_stop == false { 
		for v1:=0; v1 < len(lineD); v1++ { 
			dat = dat + lineD[v1] + " ";  
		}  
	} else {
		sw_stop = false
	}		
	getMainPageLastVal2( dat + ",,,,,,,,,,,," )
	rwS:= ""
	if last_rG_firstIxRowOfGr < len(inputTextRowSlice) {
		rwS = inputTextRowSlice[ last_rG_firstIxRowOfGr ].rRow1  		
	}
	outS1:= fmt.Sprintf( "%d,%s,%d,%d,html,%d,%d,%d,ix,%d,%d,w,%d,%d,%s, :row=,%s", 
				last_rG_ixSelGrOption, 
				last_rG_group,
				last_rG_firstIxRowOfGr ,     
				last_rG_lastIxRowOfGr  ,      
				
				last_html_rowGroup_index_gr, 
				last_html_rowGroup_beginNum, 
				last_html_rowGroup_numRows,  
    
				last_ixRowBeg,     
				last_ixRowEnd,     
				
				last_word_fromWord, 	
				last_word_numWords,		
				
				last_sel_extrRow, 	
				rwS) 
	if outS1 == "" {
		fmt.Println( red("read_lastValueSets2"), " outS1 empty =" + outS1, " \n" + " sw_stop=", sw_stop, "  dat=" , dat)
		return
	}				
	go_exec_js_function( "js_go_valueFromLastRun", outS1 )	
	
	
} // end read_lastValueSets2()

//----------------------------

func getMainPageLastVal2(valueStr string ) {

	//fmt.Println("getMainPageLastVal2(valueStr=", valueStr)
	
	col1:= strings.Split( valueStr, ",") 
	for _, ele1:= range(col1) {
		el1       := strings.Split( ele1, "=")
		if len(el1) != 2 {continue} 
		name1 := strings.TrimSpace( el1[0] ) 
		var1  := strings.TrimSpace( el1[1] )		
		if ((name1 == "") || (var1  == "")) { continue } 
		//fmt.Println("  \tgetMainPageLastVal2  name1=", name1,  " \t var1=", var1)
		switch( name1 ) {
			case "rG_ixSelGrOption"  :  last_rG_ixSelGrOption , _ = strconv.Atoi( var1 )
			case "rG_group"          :  last_rG_group             = var1 
			
			case "rG_firstIxRowOfGr" :  last_rG_firstIxRowOfGr, _ = strconv.Atoi( var1 )     
			case "rG_lastIxRowOfGr"  :  last_rG_lastIxRowOfGr , _ = strconv.Atoi( var1 )    		
			
			case "html_rowGroup_index_gr" :  last_html_rowGroup_index_gr, _ = strconv.Atoi( var1 )
			case "html_rowGroup_beginNum" :  last_html_rowGroup_beginNum, _ = strconv.Atoi( var1 )
			case "html_rowGroup_numRows"  :  last_html_rowGroup_numRows , _ = strconv.Atoi( var1 )	
			
			case "ixRowBeg"               :  last_ixRowBeg              , _ = strconv.Atoi( var1 )
			case "ixRowEnd"               :  last_ixRowEnd              , _ = strconv.Atoi( var1 )
			
			case "w_fromWord"        :  last_word_fromWord    , _ = strconv.Atoi( var1 )	
			case "w_numWords"        :  last_word_numWords    , _ = strconv.Atoi( var1 )	
			
			case "sel_extrRow"       :  last_sel_extrRow          = var1  
		}	
	} 	
	
	fmt.Println( "\n" + green("last_mainpage_values2"),  
		    "\n\t", green("rG_ixSelGrOption"  ), last_rG_ixSelGrOption ,
			"\n\t", green("rG_group"          ), last_rG_group         ,
			
			"\n\t", green("rG_firstIxRowOfGr" ), last_rG_firstIxRowOfGr ,
			"\n\t", green("rG_lastIxRowOfGr"  ), last_rG_lastIxRowOfGr  ,
			
			"\n\t", green("html_rowGroup_index_gr" ), last_html_rowGroup_index_gr, 
			"\n\t", green("html_rowGroup_beginNum" ), last_html_rowGroup_beginNum, 
			"\n\t", green("html_rowGroup_numRows"  ), last_html_rowGroup_numRows , 				
			
			"\n\t", green("ixRowBeg"          ), last_ixRowBeg         ,         
			"\n\t", green("ixRowEnd"          ), last_ixRowEnd         ,
			
			"\n\t", green("w_fromWord"        ), last_word_fromWord ,	
			"\n\t", green("w_numWords"        ), last_word_numWords ,	
			
			"\n\t", green("sel_extrRow"       ), last_sel_extrRow  )
	
	
} // end of getMainPageLastVal2()
//-----------------------------------------------