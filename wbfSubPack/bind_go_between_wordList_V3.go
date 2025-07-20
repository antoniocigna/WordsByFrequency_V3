package wbfSubPack

	import (
		"fmt"
		"strings"		
		"strconv"
		"slices"
		//"regexp"
		//"sort"
	)
//--------------------------------------------------------

func bind_go_passToJs_prefixWordList( numWords int, wordPrefix string, js_function string) {
	
	bind_go_passToJs_betweenWordList_V3( numWords, wordPrefix, js_function) 
			
} // end of bind_go_passToJs_prefixWordList

//-----------------------------------------

//--------------------------------------------------
func get_word_row_list( maxNumWords int, fromWordPref string) (string, []int, []int) {
	
	//fmt.Println( red("\n 0 get_word_row_list" + fromWordPref) )  
	
	var onlyThisLevel string = "any" ; // "A0"  // questo deve arrivare da parametro  
	var outS1 string; 	
	//----------------------------------------------------------------------
	fromWord   := strings.ToLower(strings.TrimSpace( fromWordPref));  
	lenFrom := len(fromWord) 	
	sw_oneWord := false 
	wordPrefixIndexList:= make([]int,0, 2*maxNumWords)
	wordSuffixIndexList:= make([]int,0, 2*maxNumWords)
	swThereIsSuffixList := false
	
	if fromWord == "" { 
			//go_exec_js_function( js_function, "");
			return "", wordPrefixIndexList, wordSuffixIndexList
	}	
	if fromWord[0:1] == "-" {
		// this is a suffix request 
		if fromWord[lenFrom-1:] == "-" {  // the word might be a prefix or a suffix
			wordSuffixIndexList = getListInverseWordIndex( fromWordPref[1:lenFrom-1], maxNumWords) 
		} else {	
			wordSuffixIndexList = getListInverseWordIndex( fromWordPref[1:]         , maxNumWords) 
		}
		swThereIsSuffixList = (len(wordSuffixIndexList) > 0)
	}
	if fromWord[lenFrom-1:] == "-" {
		// prefisso 
		fromWord = fromWord[:lenFrom-1] 
		if fromWord[0:1] == "-" {fromWord = fromWord[1:] }   // word is a prefix and a suffix  
	} else {
		sw_oneWord = true    // fromWord contains the only word to look for,  it's not a prefix neither a suffix 
	}
	
	lenFrom = len(fromWord)  
	
	fromWordCod:= seqCode(fromWord)		
	from1, _:= lookForWordInUniqueAlpha( fromWordCod)		
	fromIx2 :=0	
	fromWordTarg := (strings.Split(fromWordCod,"."))[0]
	lenFrom = len(fromWordTarg)  

	lenCk   :=0
	//--
	if from1 < 0 {from1=0}	
	//---------
	fromIx2 = from1 
	for k:= from1; k >=0; k-- {
		wAlf   := uniqueWordByAlpha[k]		
		lenCk   = len(wAlf.uWordSeq)
		if lenCk > lenFrom { lenCk = lenFrom }		
		if wAlf.uWordSeq[0:lenCk] < fromWordTarg {  break } 		
		fromIx2 = k
	}	
	//---------
	num1:=0	
	onlyIfExtr := false 
	//----
	for k:= fromIx2; k < len( uniqueWordByAlpha); k++ {		
		wAlf   := uniqueWordByAlpha[k]
		lenCk   = len(wAlf.uWordSeq)
		
		if sw_oneWord {
			if wAlf.uWordSeq[0:lenCk] != fromWordTarg {
				if wAlf.uWordSeq[0:lenCk] > fromWordTarg { break } 
				continue					
			} 
		} else {			
			if lenCk > lenFrom { lenCk = lenFrom}		
			// compare using the length of the prefix, I shall match just the beginning and nothing else    
			
			if wAlf.uWordSeq[0:lenCk] < fromWordTarg { //	fmt.Println(" continue "); 
				continue} 		
			if wAlf.uWordSeq[0:lenCk] > fromWordTarg { //fmt.Println(" break    "); 
				break } 			
		}
		sw, rowW := word_to_row("", onlyIfExtr, onlyThisLevel,  wAlf, -1 )	
		//fmt.Println( green("  2 get_word_row_list "), " k=", k, "  rowW=", rowW, " \nwAlf=", wAlf)  
		
		if sw == false { continue }
		wordPrefixIndexList = append(wordPrefixIndexList, k)
		if swThereIsSuffixList {
			sufix:= slices.Index(wordSuffixIndexList, k)
			if sufix >=0 {wordSuffixIndexList[sufix] = -1; } // remove index of suffix index list if the word has been already got here			
		}
		outS1 += rowW 	
		num1++
		if num1 >= maxNumWords { break }
	} // end for k	
	
	//--------------------	
	
	for z, ixWord:= range wordSuffixIndexList {
		//fmt.Println("wordSuffixIndexList ixWord=",ixWord)
		if (ixWord < 0) {continue} 
		wAlf := uniqueWordByAlpha[ixWord] 
		sw, rowW := word_to_row("", onlyIfExtr, onlyThisLevel,  wAlf,-1 )  	
		//fmt.Println( green(" 3 get_word_row_list" ), " z=" , z, " ixWord=", ixWord ,  " rowW=", rowW)  
		//fmt.Println("    wordSuffixIndexList sw=", sw, " wAlf=", wAlf.uWordSeq) 
		if sw == false {
			wordSuffixIndexList[z] = -1; 
			continue 
		}
		outS1 += rowW 	
		num1++
		if num1 >= maxNumWords { break }
	}
	//-----------------------------
	if num1 < 1 {	
		//rowW:= notFoundWord_row( fromWordCod, fromWord)		
		//outS1 += rowW  
		outS1 = ""; //   NONE," + fromWord
	}	
	
	return outS1, wordPrefixIndexList, wordSuffixIndexList
	
} // end of get_word_row_list
//--------------------------------------
func bind_go_passToJs_betweenWordList_V3( maxNumWords int, fromWordPref string, js_function string) {
	
	outS1, _, _ := get_word_row_list( maxNumWords, fromWordPref) 
	
	go_exec_js_function( js_function, outS1 ); 		
			
} // end of bind_go_passToJs_betweenWordList
//------------------



//------------------------------------------------------

func word_to_row(onlyThisLemma string, onlyIfExtr bool, onlyThisLevel string, xWordF2 wordIxStruct, numRowsInW int) (bool, string)  {
	
	//fmt.Println("\nXXXXXXX word_to_row ( onlyThisLevel=" + onlyThisLevel + 
	//  	"<==   onlyThisLevel=" + level_other +  "<==" + "  sw_list_Word_if_in_ExtrRow=" , sw_list_Word_if_in_ExtrRow); 
	
	sw:= true
	
	if onlyIfExtr {
		if (sw_list_Word_if_in_ExtrRow) {
			if (xWordF2.uSwSelRowG == SEL_NO_EXTR_ROW) {   //    1 or 2: 1 SEL_EXTR_ROW, 2 SEL_NO_EXTR_ROW  
				//             xWordF2.uSwSelRowG è stato impostato nell'ultima esecuzione di "Lista le Righe" dove è stato scelto il gruppo di righe ed il range da-a
				sw = false
				//fmt.Println("word to row " , xWordF2.uWord2, " da ignorare")
				return sw, ""
			} 
		}
	}
	//fmt.Println("word to row " , xWordF2.uWord2, " \t\t\t XXXXXXXXXXXX (xWordF2.uTotExtrRow =" ,xWordF2.uTotExtrRow , " xxxxxxxxxxxxxxxxxxx   accettato")
	//-----------
	totNumRow:= numRowsInW
	if totNumRow < 1 {
		totNumRow = xWordF2.uTotExtrRow
	}
	
	if onlyThisLemma != "" {
		ix2 := -1
		for x1, oneLemma := range xWordF2.uLemmaL {
			if oneLemma == onlyThisLemma {
				ix2=x1; 
				break
			}  	
		}
		if ix2 >=0 {
			//fmt.Println("word to row 1 " , xWordF2.uWord2, " ix2=", ix2, " onlyIfExtr=", onlyIfExtr,  "    WordF2.uTotExtrRow =" ,xWordF2.uTotExtrRow	)
			return sw, xWordF2.uWordSeq + ";." + xWordF2.uWord2 + ";." + 
				"ix" + ";." + 
				strconv.Itoa(xWordF2.uIxUnW) + ";." + strconv.Itoa(xWordF2.uTotRow)  + ";." + 
				xWordF2.uLemmaL[ix2]              + ";." + 
				tranFromIxLemma( xWordF2, ix2)    + ";." +  
				xWordF2.uLevel[ix2]               + ";." +  
				xWordF2.uPara[ix2]                + ";." +  
				xWordF2.uExample[ix2]             + ";." +  
				strconv.Itoa(totNumRow) + ";." +  		
				xWordF2.uLearnedYN                + ";." + 	
				"ixLemma" + ";." + fmt.Sprint( xWordF2.uIxLemmaL[ix2] ) + ";." + 	
				endOfLine 		
		}
	}
	
	//---------------------------	
	//fmt.Println("word to row 2 " , xWordF2.uWord2, "  onlyIfExtr=", onlyIfExtr,  "    WordF2.uTotExtrRow =" ,xWordF2.uTotExtrRow	)
	return sw, xWordF2.uWordSeq + ";." + xWordF2.uWord2 + ";." + 
				"ix" + ";." + 
				strconv.Itoa(xWordF2.uIxUnW) + ";." + strconv.Itoa(xWordF2.uTotRow)  + ";." +
				fmt.Sprint( strings.Join(xWordF2.uLemmaL,  wSep)  ) + ";." + 
				//fmt.Sprint( strings.Join(xWordF2.uTranL,   wSep)  ) + ";." +  
				listStringLemmaSlice_Tran(xWordF2) + ";." +  
				fmt.Sprint( strings.Join(xWordF2.uLevel,   wSep)  ) + ";." +  
				fmt.Sprint( strings.Join(xWordF2.uPara,    wSep)  ) + ";." +  
				fmt.Sprint( strings.Join(xWordF2.uExample, wSep)  ) + ";." +  
				strconv.Itoa(totNumRow) + ";." +  			
				xWordF2.uLearnedYN              + ";." + 			
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

		sw, rowW := word_to_row("", false, "any",  myAlf , -1)  	
		
		if sw == false { continue }
		outS1 += rowW 	
		num1++
		if num1 >= maxNumWords { break }
	}
	if num1 < 1 {	
		//rowW:= notFoundWord_row( fromWordSuff,  fromWordSuff)		
		//outS1 += rowW 	
		//outS1 += "NONE," + fromWordSuff
		outS1 = ""; 
	}			
	//------------------
	
	go_exec_js_function( js_function, outS1 ); 	
		
} // end of  bind_go_passToJs_suffixWordList

//-------------------------------------------------