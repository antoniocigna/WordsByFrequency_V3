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


//-------------------------------

func split_ALL_word_dict_row(  strRows string) string {
	//fmt.Println( "ANTONIO xxxxxxxxxxxxxxxxxxxxxxxxxxxx  split_ALL_word_dict_row( strRows=", strRows); 
	// eg. einem;14 ; ein einem einer ;  a uno uno;	  ==> word ; ix : list of lemmas ; list of translations	
	/**	
	ANTONIO xxxxxxxxxxxxxxxxxxxxxxxxxxxx  split_ALL_word_dict_row( 
	strRows= 
dem;31;dem;dem§
den;7;den;tana§
der;0;der;il§
die;8;die;il§	
	***/
	lemmaTranStr := ""
	
	lines := strings.Split( strRows, "\n");	
	
	var ele1 lemmaTranStruct     
	/**
	//--------------------------
		type lemmaTranStruct struct {
			dL_lemmaCod  string 
			dL_lemma2    string  
			dL_numDict   int	  
			dL_tran      string
		} 
	**/
	
	lenIx1:= len(uniqueWordByFreq)
	
	
	//------------  solo per controllare----	
	/**
	for ix1:=0;  ix1 < 28; ix1++ {   
		fmt.Println( "ANTONIO error7 check inizio ix1=", ix1, " uniqueWordByFreq[ix1]=" , uniqueWordByFreq[ix1] )
	}
	**/	
	/***
	fmt.Println( "ANTONIO xxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	for ix1:=(lenIx1-5);  ix1 < lenIx1; ix1++ {   
		fmt.Println( "ANTONIO error7 check fine   ix1=", ix1, " uniqueWordByFreq[ix1].ixUnW ="  , uniqueWordByFreq[ix1].ixUnW )
	}
	***/
	
	//-------------- fine solo per controllare-----
	
	for z:=0;  z < len(lines); z++ {   
		
		//fmt.Println("\nsplit_ALL_word_dict_row() lines[",z,"]=", lines[z] )	
		
		_, ix1, lemmaLis,tranLis := split_one_word_dict_row( lines[z] )
		if ix1 < 0 { continue }		
		
		if ix1 >= lenIx1 { 
			fmt.Println("error7 1 len(uniqueWordByFreq)=", len(uniqueWordByFreq), " ix1=", ix1 ,  " lines[z=", z, "]=", lines[z] )
			continue 
		}
		//---------------
		//fmt.Println("   ix1=", ix1, " \tlemmaLis=", lemmaLis, "\ttranLis=", tranLis )  
		
		/**
		  ix1= 37      lemmaLis= [ihr ihre ihrer sein]         tranLis= [tu loro loro essere ]
		**/
		
		//---------------
		if uniqueWordByFreq[ix1].uIxUnW != ix1 {	
			fmt.Println("error7 2 len(uniqueWordByFreq)=", len(uniqueWordByFreq), " ix1=", ix1 ,  " lines[z=", z, "]=", lines[z] )
			continue 
		}  // error	
		
		//swUpdList[ix1] = true 
		//numUpd++  		
		//ixAlfa := uniqueWordByFreq[ix1].uIxUnW_al  	
		
		//fmt.Println("  uniqueWordByFreq[ix1=", ix1, "] = ",uniqueWordByFreq[ix1], "\n\tuLemmaL=", uniqueWordByFreq[ix1].uLemmaL) 
		
		/***
		 {ihren.ihren ihren 37 29 1 1 63 1 1 0 0 [25114 25115 25116 45916] [ihr ihre ihrer sein] [   ] [   ] [   ]}
		 
			uLemmaL= [ihr ihre ihrer sein]	
		
		***/
		len1:= len(uniqueWordByFreq[ix1].uLemmaL)
		
		//fmt.Println("\tsplit_ALL_word_dict_row() 2 ", " ix1=", ix1, " len1=", len1, ", len(lemmaLis)=", len(lemmaLis), ",     len(tranLis)=",  len(tranLis) )
		if len1 != len(lemmaLis) { 
			fmt.Println(red("split_ALL_word_dict_row"),"() 1 ", lines[z] ," number of lemma not equal ")
			for mio1:=0; mio1 < len1; mio1++ {
				fmt.Println("\tsplit_ALL_word_dict_row() 2.1 ", "  uniqueWordByFreq[ix1].uLemmaL[", mio1,"] =>" + uniqueWordByFreq[ix1].uLemmaL[mio1] + "<==")
			}   
		}   
		if len1 != len(lemmaLis) { fmt.Println( red("error1"), " split_ALL_word_dict_row!") ;  continue }               // error 
		
		if len1 != len(tranLis)  { fmt.Println( red("error2"), " split_ALL_word_dict_row!") ;continue }               // error 
		
		//---------------------
		oneW := uniqueWordByFreq[ix1]
		
		for m:=0; m < len1; m++ {
			mLemm := strings.TrimSpace( lemmaLis[m] )
			mTran := strings.TrimSpace( tranLis[m] )
			if mLemm == oneW.uLemmaL[m] {  
				// fmt.Println("lemma trovato con m=",m," ", mLemm)
				ixLe:= oneW.uIxLemmaL[m]
				lemmaSlice[ixLe].leTran             = mTran
				//uniqueWordByFreq[ix1].uTranL[m]     = mTran 	
				//uniqueWordByAlpha[ixAlfa].uTranL[m] = mTran 	
			} else { 
				//fmt.Println("lemma NON trovato, lo cerco ", mLemm)
				for m2:=0; m2 < len(oneW.uLemmaL); m2++ {
					if mLemm == oneW.uLemmaL[m2] {
						// fmt.Println("lemma trovato con m2=",m2," ", mLemm)
						ixLe:= oneW.uIxLemmaL[m2]
						lemmaSlice[ixLe].leTran              = mTran
						//uniqueWordByFreq[ix1].uTranL[m2]     = mTran 	
						//uniqueWordByAlpha[ixAlfa].uTranL[m2] = mTran 	
						break;
					}
				} 
			}
		} 
		//---------------------------
		
		for m:=0; m < len1; m++ {		
			mLemm := strings.TrimSpace( lemmaLis[m] )
			mTran := strings.TrimSpace( tranLis[m] 	)	
				
			lemmaTranStr += "\n" + mLemm + "|" + mTran 	
			
			ele1.dL_lemmaCod = newCode(mLemm) 
			ele1.dL_lemma2   = mLemm 
			ele1.dL_numDict  = lastNumDict
			ele1.dL_tran     = mTran          
			if mTran != "" { 		
				dictLemmaTran = append( dictLemmaTran, ele1 ) 
			}
		}
		
	} // end of z 
	
	return lemmaTranStr
	
} // end of split_ALL_word_dict_row(

//------------------------------------------

//--------------------

func update_rowTranslation( rowDictRow [] string)  {
	/*
	rowDictRow is filled in javascript, each line: rowIndex;row Translation   
	*/
	var len1 = len(inputTextRowSlice)
	
	//fmt.Println("update_rowTraslation() len(rowDictRow) =",len(rowDictRow), " = ", rowDictRow)
	
	//ion() len(rowDictRow) =  [ 2 206(2_206 211|211|1. Introduzione Preistoria 2 207(2_207 212|212|La prei
	
	for z:=0; z < len(rowDictRow); z++ {  		
		/**
		2;Primo capitolo
		**/
		row1dict := rowDictRow[z] 	
		if row1dict == "" { continue }	
		rowCols	 := strings.Split(row1dict, "|")
		
		//fmt.Println("\trowCols= len=", len(rowCols), " rowcols=",    rowCols) 
		
		if len(rowCols) < 3 { 
			fmt.Println("update_rowTraslation() len(rowDictRow) z=", z, " row1dict=", row1dict, " len(rowCols)=", len(rowCols))
			continue 
		}
		
		idS   := rowCols[0]   
		ixS   := rowCols[1] 
			
		tranS := rowCols[2]
		
		//fmt.Println("\tidS=", idS, "  ixS=", ixS,  " tranS=", tranS ) 
		
	
		ixRow, err := strconv.Atoi(  strings.TrimSpace( ixS) )		
		
		//fmt.Println("\tixRow=", ixRow, " err=", err)
		
		if err != nil {
			return 
		}
		if ixRow >= len1 { return }  // error 	

		//fmt.Println("\tinputTextRowSlice[ixRow].rIdRow = ", inputTextRowSlice[ixRow].rIdRow )
		
		if strings.Index( idS, inputTextRowSlice[ixRow].rIdRow) < 0  {
			// error
			fmt.Println("error func update_rowTranslation () inputTextRowSlice[ixRow=", ixRow, "].rIdRow = ", inputTextRowSlice[ixRow].rIdRow,   " idS=",  idS )
			continue
		}  
		inputTextRowSlice[ixRow].rTran1 = strings.TrimSpace(tranS ) 
		
		//fmt.Println("update_rowTraslation() row1dict=", row1dict, " \n\t ixRow=", ixRow, " inputTextRowSlice[ixRow]=", inputTextRowSlice[ixRow] )
		
	} // end for z
	
} // end of update_rowTranslation

//---------------------------------------------------------
//==============================================