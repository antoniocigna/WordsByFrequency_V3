package wbfSubPack

	import (
		"fmt"
		"strings"		
		"strconv"
	)
//--------------------------------------------------------

//-------------------------------------------
func get_all_binds() {
		fmt.Println("func get_all_binds")
		
		/***
		// A simple way to know when UI is ready (uses body.onload event in HTML/JS)
		ui.Bind("goStart", func() { 
				fmt.Println("Bind goStart   sw_begin_ended=", sw_begin_ended); 
				sw_HTML_ready = true 
				if sw_begin_ended {
					bind_goStart("1get_all_binds")
				}	
			} )
		**/	
		//--------------
		ui.Bind("go_passToJs_html_is_ready", func( msg1 string,  js_function string) { 		
				bind_go_passToJs_html_is_ready( msg1,  js_function)  })
		//--------------
		/**
		ui.Bind("go_passToJs_html_and_go_ready", func( msg1 string,  js_function string) { 	
			bind_go_passToJs_html_and_go_ready( msg1,  js_function)  })
		**/
		//--------------
		//----------------------------
		ui.Bind("go_passToJs_wordList", func( isChange bool, s_fromWord string, s_numWords string, sel_level string, 
															sel_extrRow string, sel_toBeLearned string,  js_function string) { 		
				bind_go_passToJs_wordList( isChange, getInt(s_fromWord), getInt(s_numWords), sel_level, sel_extrRow, sel_toBeLearned,  js_function)  })
		//--------------------------------------		
		ui.Bind("go_passToJs_prefixWordList", func(s_numWords string, wordPrefix string, js_function string) {
				bind_go_passToJs_prefixWordList( getInt(s_numWords), wordPrefix, js_function) } )
		//--------------------------------------		
		ui.Bind("go_passToJs_betweenWordList", func( s_maxNumWords string, fromWordPref string, js_function string) {
				bind_go_passToJs_betweenWordList( getInt(s_maxNumWords), fromWordPref, js_function) } ) 	
	 	//--------------------------------------		
		ui.Bind("go_passToJs_betweenLemmaList", func( s_maxNumLemmas string, fromLemmaPref string, js_function string) {
				bind_go_passToJs_betweenLemmaList( getInt(s_maxNumLemmas), fromLemmaPref, js_function) } ) 		 			
			
		//---------------------------------------
		ui.Bind("go_passToJs_lemmaWordList", func(lemma string, inpMaxWordLemma string, js_function string) {
				bind_go_passToJs_lemmaWordList(lemma, getInt( inpMaxWordLemma),  js_function) } ) 
		//---------------------------------------	
		ui.Bind("go_passToJs_getRowsByIxWord", func( sIxWord string, max_num_row4wordS string, js_function string) {
				bind_go_passToJs_getRowsByIxWord(  getInt(sIxWord), getInt(max_num_row4wordS), js_function) } ) 
		//-----------------------------------
		ui.Bind("go_passToJs_getRowsByIxLemma", func( sIxLemma string, max_num_row4lemmaS string, js_function string) {
				bind_go_passToJs_getRowsByIxLemma(  getInt(sIxLemma), getInt(max_num_row4lemmaS), js_function) } ) 
		//-----------------------------------
		
		ui.Bind("go_passToJs_getWordByIndex2", func( s_ixWord string, swOnlyThisWordRows bool, s_maxNumRow string,  js_function string) {
				bind_go_passToJs_getWordByIndex2(   getInt(s_ixWord), swOnlyThisWordRows,      getInt(s_maxNumRow), js_function) } ) 
		//---------------------------------------
		ui.Bind("go_passToJs_thisWordRowList", func( aWord string, s_maxNumRow string, js_function string) {				
				bind_go_passToJs_thisWordRowList( aWord, getInt(s_maxNumRow), js_function) } )
		//---------------------------------------
		ui.Bind("go_passToJs_rowList", func(  s_inpBegRow string,   s_maxNumRow string, js_function string) {	
				bind_go_passToJs_rowList(     getInt(s_inpBegRow), getInt(s_maxNumRow), js_function) } )
		//---------------------------------------
	
		ui.Bind("go_passToJs_rowWordList", func( numIdOut string, s_ixRR string, js_function string ) {
				bind_go_passToJs_rowWordList(numIdOut, getInt(s_ixRR), js_function) } ) 	
		//---------------------------------------	
		ui.Bind("go_write_lang_dictionary", func( langAndVoiceName string) {
			bind_go_write_lang_dictionary( langAndVoiceName ) } ) 
		//---------------------------------------	
		ui.Bind("go_write_word_dictionary", func( listGoWords string) {
			bind_go_write_word_dictionary( listGoWords ) } )  				
		//---------------------------------------			
		ui.Bind("go_write_row_dictionary", func( listGoRows string) {
			bind_go_write_row_dictionary( listGoRows ) } )  
		//---------------------
		ui.Bind("go_passToJs_word_known", func( s_ixWord string,  s_yesNo string,  s_knowCtr string, js_function string) {
			bind_go_passToJs_word_known(       getInt(s_ixWord), getInt(s_yesNo), getInt(s_knowCtr), js_function)  } )  
			
		//----------------------------------------------------------------
		/**
		ui.Bind("go_passToJs_read_wordsToLearn", func( js_function string) {
			bind_go_passToJs_read_wordsToLearn(js_function)  } )  
		**/
		//----------------------------------------------------------------
		ui.Bind("go_passToJs_write_WordsToLearn", func( js_function string) {
			bind_go_passToJs_write_WordsToLearn(js_function)  } )  	
			
		//-----------------------------------------
		/***
		ui.Bind("go_passToJs_updateRowGroup", func( s_index string,  s_inpBegRow string,  s_inpNumRow string,  js_function string) {	
			bind_go_passToJs_updateRowGroup(       getInt(s_index), getInt(s_inpBegRow), getInt(s_inpNumRow),  js_function)   } ) 
		**/
		//-------------------------------------------------------------------
		
		ui.Bind("go_passToJs_getIxRowFromGroup", func( s_rowGrIndex string,  s_html_rowGroup_beginNum string, s_html_rowGroup_numRows string, js_function string) {		
			bind_go_passToJs_getIxRowFromGroup( getInt(s_rowGrIndex),  getInt( s_html_rowGroup_beginNum), getInt( s_html_rowGroup_numRows),  js_function)  } ) 
		
		//----------------------------------------------------------------
}  
//---------------------------------------------
func bind_go_passToJs_html_is_ready( msg1 string,  js_function string) {
	fmt.Println("\n", "go func bind_go_passToJs_html_is_ready () " , "\n\t msg from html: ", msg1 )  
	
	fmt.Println("XXXXXXXXXX   ", green("html has been loaded"), "   XXXXXXXXXXXX")
	
	begin() 
	
	fmt.Println( green( "\n\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n" + 
		"xxxxxxxxxxxxxx you can use the tool xxxxxxxxxxxxxxxxxx\n"  + 
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n\n"  ) ) 
		
	//prova_js_function_treValori();
	
	//go_exec_js_function( js_function, msg1); 	
	
} // end of bind_go_passToJs_html_is_ready
 
//---------------------------------------------
/**
func bind_go_passToJs_html_and_go_ready( msg1 string,  js_function string) {
	fmt.Println("\n", "go func  bind_go_passToJs_html_and_go_ready() " , "\n\t msg from html: ", msg1 )  
	
	fmt.Println("XXXXXXXXXX   PRONTO  XXXXXXXXXXXX")
	
	begin() 
	
} // end of bind_go_passToJs_html_and_go_ready
***/
//-------------------------------------------------------- 

//---------------------------		

func getInt(x string) int {	
	y1, e1 := strconv.Atoi( x ) 
	if e1 == nil { 
		return y1
	} 
	y2, e2 := strconv.Atoi(  "0"+strings.TrimSpace(x)  ) 
	if e2 == nil {
		return y2
	} else {
		fmt.Println("error in getInt(",x,") ", e2) 
	}
	return 0
}

//-------------------------------------
func showErrMsg(errorMSG0 string, msg1 string, func1 string) {
	errorMSG = strings.ReplaceAll( errorMSG0, "\n"," ") 
	fmt.Println(msg1, " func ", func1)
	go_exec_js_function( "js_go_setError", errorMSG ); 	
}
//--------------------------------
func showErrMsg2(errorMSG0 string, msg1 string) {
	errorMSG = strings.ReplaceAll( errorMSG0, "\n"," ") 
	fmt.Println(msg1)
	go_exec_js_function( "js_go_setError", errorMSG ); 	
}
//------------------------------
func showInfoRead( fileName string, startStop string) {
		
	msg1_Js := "file " + fileName + " " + startStop 
 	msg1:=`<br>file &nbsp;&nbsp;<span style="color:blue;" >` + fileName + `</span>`	+ 				
				`<br><span style="font-size:0.7em;color:black;">`	+ startStop	+ `</span>` 		
				
	showErrMsg2(msg1, msg1_Js)	
}
//------------------------------------------------
func isNumber(s string) bool {
    for _, v := range s {
        if v < '0' || v > '9' {
            return false
        }
    }
    return true
}




//-----------------------------------------------------------

//==============================================