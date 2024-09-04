package wbfSubPack

	import (
		"fmt"
		"strings"		
		"strconv"
		//"regexp"
		//"sort"
	)
//--------------------------------------------------------

func bind_go_passToJs_prefixWordList( numWords int, wordPrefix string, js_function string) {
	
	bind_go_passToJs_betweenWordList( numWords, wordPrefix, js_function) 
			
} // end of bind_go_passToJs_prefixWordList

//-----------------------------------------

func bind_go_passToJs_betweenWordList( maxNumWords int, fromWordPref string, js_function string) {
	
	var onlyThisLevel string = "any" ; // "A0"  // questo deve arrivare da parametro  
	var outS1 string; 	
	//----------------------------------------------------------------------
	fromWord   := strings.ToLower(strings.TrimSpace( fromWordPref));  
	lenFrom := len(fromWord) 	
	sw_oneWord := false 
	
	if fromWord == "" { 
			go_exec_js_function( js_function, "");
			return 
	}	
	if fromWord[0:1] == "-" {
		// this is a suffix request 
		bind_go_passToJs_suffixWordList( maxNumWords, fromWordPref[1:], js_function)
		return
	}
	if fromWord[lenFrom-1:] == "-" {
		// prefisso 
		fromWord = fromWord[:lenFrom-1] 
	} else {
		sw_oneWord = true    // fromWord contains the only word to look for,  it's not a prefix neither a suffix 
	}
	
	lenFrom = len(fromWord)  
	
	fromWordCod:= newCode(fromWord)		
	from1, _:= lookForWordInUniqueAlpha( fromWordCod)		
	fromIx2 :=0	
	fromWordTarg := (strings.Split(fromWordCod,"."))[0]
	lenFrom = len(fromWordTarg)  
	lenCk   :=0
	//--
	if from1 < 0 {from1=0}
	
	//fmt.Println( green("bind_go_passToJs_betweenWordList"), " from1=", from1, "   fromWordPref=", fromWordPref, "  fromWordCod=", fromWordCod, " lenFrom=", lenFrom) 
	
	//---------
	fromIx2 = from1 
	for k:= from1; k >=0; k-- {
		wAlf   := uniqueWordByAlpha[k]
		j:= strings.Index( wAlf.uWordCod, ".") 
		if (j<0) {j = len(wAlf.uWordCod) }
		lenCk   = j; //  len(wAlf.uWordCod[0:j])
		if lenCk > lenFrom { lenCk = lenFrom }
		
		//fmt.Println(" k=", k, " wAlf.uWordCod=", wAlf.uWordCod, " lenCk=", lenCk ,"  wAlf.uWordCod[0:lenCk]=",  wAlf.uWordCod[0:lenCk],   " fromWordTarg=", fromWordTarg )
		
		if wAlf.uWordCod[0:lenCk] < fromWordTarg {  break } 
		fmt.Println("   caricto k=", k)  
		fromIx2 = k
	}	
	//---------
	
	num1:=0	
	onlyIfExtr := false 
	
	//----
	for k:= fromIx2; k < len( uniqueWordByAlpha); k++ {		
		wAlf   := uniqueWordByAlpha[k]
		j:= strings.Index( wAlf.uWordCod, ".")      // the format of wordCod is: coded_word.the_actual_word  
		if (j<0) {j = len(wAlf.uWordCod) }
		lenCk   = j; //  len(wAlf.uWordCod[0:j])
		
		if sw_oneWord {
			if wAlf.uWordCod[0:lenCk] != fromWordTarg {
				if wAlf.uWordCod[0:lenCk] > fromWordTarg { break } 
				continue					
			} 
		} else {			
			if lenCk > lenFrom { lenCk = lenFrom}		
			// compare using the length of the prefix, I shall match just the beginning and nothing else    
			//fmt.Println(" k=", k, " wAlf.uWordCod=", wAlf.uWordCod, " lenCk=", lenCk ,"  wAlf.uWordCod[0:lenCk]=",  wAlf.uWordCod[0:lenCk])
		
			if wAlf.uWordCod[0:lenCk] < fromWordTarg { fmt.Println(" continue "); continue} 		
			if wAlf.uWordCod[0:lenCk] > fromWordTarg { fmt.Println(" break    "); break } 			
		}
		sw, rowW := word_to_row("", onlyIfExtr, onlyThisLevel,  wAlf )  	
		
		if sw == false { continue }
		outS1 += rowW 	
		num1++
		if num1 >= maxNumWords { break }
	}
	if num1 < 1 {	
		rowW:= notFoundWord_row( fromWordCod, fromWord)		
		outS1 += rowW 	
	}	
	//-----------
	go_exec_js_function( js_function, outS1 ); 		
	
			
} // end of bind_go_passToJs_betweenWordList
//------------------



//------------------------------------------------------

func word_to_row(onlyThisLemma string, onlyIfExtr bool, onlyThisLevel string, xWordF2 wordIxStruct) (bool, string)  {
	
	//fmt.Println("\nXXXXXXX word_to_row ( onlyThisLevel=" + onlyThisLevel + 
	//  	"<==   onlyThisLevel=" + level_other +  "<==" + "  sw_list_Word_if_in_ExtrRow=" , sw_list_Word_if_in_ExtrRow); 
	
	sw:= true
	
	if onlyIfExtr {
		if (sw_list_Word_if_in_ExtrRow) {
			if (xWordF2.uSwSelRowG == SEL_NO_EXTR_ROW) {
				sw = false
				//fmt.Println("word to row " , xWordF2.uWord2, " da ignorare")
				return sw, ""
			} 
		}
	}
	//fmt.Println("word to row " , xWordF2.uWord2, " \t\t\t XXXXXXXXXXXX (xWordF2.uTotExtrRow =" ,xWordF2.uTotExtrRow , " xxxxxxxxxxxxxxxxxxx   accettato")
	//-----------
	
	if onlyThisLemma != "" {
		ix2 := -1
		for x1, oneLemma := range xWordF2.uLemmaL {
			if oneLemma == onlyThisLemma {
				ix2=x1; 
				break
			}  	
		}
		if ix2 >=0 {
			return sw, xWordF2.uWordCod + ";." + xWordF2.uWord2 + ";." + 
				"ix" + ";." + 
				strconv.Itoa(xWordF2.uIxUnW) + ";." + strconv.Itoa(xWordF2.uTotRow)  + ";." + 
				xWordF2.uLemmaL[ix2]              + ";." + 
				tranFromIxLemma( xWordF2, ix2)    + ";." +  
				xWordF2.uLevel[ix2]               + ";." +  
				xWordF2.uPara[ix2]                + ";." +  
				xWordF2.uExample[ix2]             + ";." +  
				strconv.Itoa(xWordF2.uTotExtrRow) + ";." +  					
				strconv.Itoa(xWordF2.uKnow_yes_ctr)  + ";." + strconv.Itoa(xWordF2.uKnow_no_ctr)  + ";." + 				
				"ixLemma" + ";." + fmt.Sprint( xWordF2.uIxLemmaL[ix2] ) + ";." + 	
				endOfLine 		
		}
	}
	
	//---------------------------	
	
	return sw, xWordF2.uWordCod + ";." + xWordF2.uWord2 + ";." + 
				"ix" + ";." + 
				strconv.Itoa(xWordF2.uIxUnW) + ";." + strconv.Itoa(xWordF2.uTotRow)  + ";." +
				fmt.Sprint( strings.Join(xWordF2.uLemmaL,  wSep)  ) + ";." + 
				//fmt.Sprint( strings.Join(xWordF2.uTranL,   wSep)  ) + ";." +  
				listStringLemmaSlice_Tran(xWordF2) + ";." +  
				fmt.Sprint( strings.Join(xWordF2.uLevel,   wSep)  ) + ";." +  
				fmt.Sprint( strings.Join(xWordF2.uPara,    wSep)  ) + ";." +  
				fmt.Sprint( strings.Join(xWordF2.uExample, wSep)  ) + ";." +  
				strconv.Itoa(xWordF2.uTotExtrRow) + ";." +  					
				strconv.Itoa(xWordF2.uKnow_yes_ctr)  + ";." + strconv.Itoa(xWordF2.uKnow_no_ctr)  + ";." + 				
				"ixLemma" + ";." + intSliceToString( xWordF2.uIxLemmaL,wSep )  + ";." + 		
				endOfLine 
				 
} // end of word_to_row 
//-----------------------------------
func intSliceToString(mySlice []int, wSep string) string {
	output := ""
	for _, v := range mySlice {
		output += (strconv.Itoa(v) + wSep)
	}
	return output
}
//------------------------------------------------------

func notFoundWord_row( fromWordCod string, fromWord string ) string {
	
	return fromWordCod + ";." + fromWord + ";." + 
			"ix" + ";." + 
			strconv.Itoa(0) + ";." + strconv.Itoa(0)  + ";." + 
			fmt.Sprint( fromWord ) + ";." + 
			fmt.Sprint( "_word_not_found_" ) + ";." +  
			fmt.Sprint( "" ) + ";." +  
			fmt.Sprint( "" ) + ";." +  
			fmt.Sprint( "" ) + ";." +  
			strconv.Itoa(0) + ";." +  					
			strconv.Itoa(0) + ";." + strconv.Itoa(0)  + ";." + 
			"ixLemma" + ";." + "" + ";." + 	
			endOfLine 
					
} // end of notFoundWord_row

//---------------------------------------------------


//--------------------------------------------


func bind_go_passToJs_suffixWordList( maxNumWords int, fromWordSuff string, js_function string) {
	
	var outS1 string;
	
	listInverseWordIndex := getListInverseWordIndex( fromWordSuff, maxNumWords) 
	
	num1:=0	
	for _,k:= range listInverseWordIndex { 
		myAlf := uniqueWordByAlpha[k]

		sw, rowW := word_to_row("", false, "any",  myAlf )  	
		
		if sw == false { continue }
		outS1 += rowW 	
		num1++
		if num1 >= maxNumWords { break }
	}
	if num1 < 1 {	
		rowW:= notFoundWord_row( fromWordSuff,  fromWordSuff)		
		outS1 += rowW 	
	}			
	//------------------
	
	go_exec_js_function( js_function, outS1 ); 	
		
} // end of  bind_go_passToJs_suffixWordList

//-------------------------------------------------