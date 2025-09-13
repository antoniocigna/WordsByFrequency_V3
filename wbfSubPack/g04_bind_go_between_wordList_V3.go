package wbfSubPack

	import (
		"fmt"
		"strings"		
		"strconv"
		//"slices"
		//"regexp"
		"sort"
	)
//--------------------------------------------------------

func g04_bind_go_passToJs_prefixWordList( numWords int, wordPrefix string, js_function string, js_parm string, js_caller string) {
	
	g04_bind_go_passToJs_betweenWordList_V3( numWords, wordPrefix, js_function, js_parm, js_caller) 
			
} // end of bind_go_passToJs_prefixWordList

//-----------------------------------------

//--------------------------------------------------
func g04_get_word_row_list( maxNumWords int, fromWordPref0 string) (string, []int, []int) {
	
	 
	fromWordPref := stdCode( fromWordPref0)
	
	//fmt.Println( red("\n 0 get_word_row_list ") + fromWordPref0 + " ==>" + fromWordPref ) 
	
	var onlyThisLevel string = "any" ; // "A0"  // questo deve arrivare da parametro  
	var outS1 string; 	
	//----------------------------------------------------------------------
	fromWord   := strings.ToLower(strings.TrimSpace( fromWordPref));  
	lenFrom := len(fromWord) 	
	sw_oneWord := false 
	sw_suffix  := false
	wordPrefixIndexList:= make([]int,0, 2*maxNumWords)
	wordSuffixIndexList:= make([]int,0, 2*maxNumWords)
	
	num1:=0	
	onlyIfExtr := false 
	
	
	if fromWord == "" { 
			//go_exec_js_function( js_function, "");
			return "", wordPrefixIndexList, wordSuffixIndexList
	}
	//----------------------		
	sw_oneOnly:= false;
	if fromWord[0:1] == "-" {
		if  fromWord[lenFrom-1:] == "-" { //     any word which contains the string in fromWord  eg.   -mili-  --> familie 
			// this case cannot be of any use, ignore it    
			return "", wordPrefixIndexList, wordSuffixIndexList
			
		} else {//a suffix request, wanted all words ending with the string in fromWord.    eg.  -en --> gehen, haben, etc.   		
			//fmt.Println("   g04  g34_getListInverseWordIndex  di " + fromWord[1:] + "<==") 
			sw_oneOnly = true;
			wordSuffixIndexList = g34_getListInverseWordIndex( fromWord[1:] , sw_oneOnly, maxNumWords) 
			sw_suffix = true 
		}	
	} else {
		if fromWord[lenFrom-1:] == "-" {// a prefix request,  wanted all words beginning with the string in fromWord.    eg.  geh-    --> geht, gehen, etc.    
			fromWord = fromWord[:lenFrom-1] 
		} else { 
			sw_oneWord = true    // fromWord contains the only word to look for,  it's not a prefix neither a suffix 
		}
	}
	
	//fmt.Println("    wordPrefixIndexList=", wordPrefixIndexList)   
	//fmt.Println("    wordSuffixIndexList=", wordSuffixIndexList)   
	
	lenFrom = len(fromWord)  
	
	fromWordCod:= fromWord		
	from1, _:= lookForWordInUniqueAlpha( fromWordCod,0)		
	fromIx2 :=0	
	fromWordTarg := fromWordCod; //  (strings.Split(fromWordCod,"."))[0]
	lenFrom = len(fromWordTarg)  

	lenCk   :=0
	//--
	if from1 < 0 {from1=0}	
	//---------
	if sw_suffix == false { 
		fromIx2 = from1 
		for k:= from1; k >=0; k-- {
			wAlf   := uniqueWordByAlpha[k]		
			lenCk   = len(wAlf.uWord2)
			if lenCk > lenFrom { lenCk = lenFrom }		
			if wAlf.uWord2[0:lenCk] < fromWordTarg {  break } 		
			fromIx2 = k
		}	
		//---------
		for k:= fromIx2; k < len( uniqueWordByAlpha); k++ {		
			wAlf   := uniqueWordByAlpha[k]
			lenCk   = len(wAlf.uWord2)
			
			if sw_oneWord {
				if wAlf.uWord2[0:lenCk] != fromWordTarg {
					if wAlf.uWord2[0:lenCk] > fromWordTarg { break } 
					continue					
				} 
			} else {			
				if lenCk > lenFrom { lenCk = lenFrom}		
				// compare using the length of the prefix, I shall match just the beginning and nothing else    
				
				if wAlf.uWord2[0:lenCk] < fromWordTarg { //	fmt.Println(" continue "); 
					continue} 		
				if wAlf.uWord2[0:lenCk] > fromWordTarg { //fmt.Println(" break    "); 
					break } 			
			}
			sw, rowW := word_to_row("", onlyIfExtr, onlyThisLevel,  wAlf, -1 )	
			//fmt.Println( green("  2 get_word_row_list "), " k=", k, "  rowW=", rowW, " \nwAlf=", wAlf)  
			
			if sw == false { continue }
			wordPrefixIndexList = append(wordPrefixIndexList, k)
			/**
			if swThereIsSuffixList {
				sufix:= slices.Index(wordSuffixIndexList, k)
				if sufix >=0 {wordSuffixIndexList[sufix] = -1; } // remove index of suffix index list if the word has been already got here			
			}
			**/
			outS1 += rowW 	
			num1++
			if num1 >= maxNumWords { break }
		} // end for k	
	}
	//--------------------	
	if sw_suffix { 
		for z, ixWord:= range wordSuffixIndexList {
			//fmt.Println("wordSuffixIndexList ixWord=",ixWord)
			if (ixWord < 0) {continue} 
			wAlf := uniqueWordByAlpha[ixWord] 
			sw, rowW := word_to_row("", onlyIfExtr, onlyThisLevel,  wAlf,-1 )  	
			//fmt.Println( green(" 3 get_word_row_list" ), " z=" , z, " ixWord=", ixWord ,  " rowW=", rowW)  
			//fmt.Println("    wordSuffixIndexList sw=", sw, " wAlf=", wAlf.uWord2) 
			if sw == false {
				wordSuffixIndexList[z] = -1; 
				continue 
			}
			outS1 += rowW 	
			num1++
			if num1 >= maxNumWords { break }
		}
	}
	//-----------------------------
	if num1 < 1 {	
		//rowW:= notFoundWord_row( fromWordCod, fromWord)		
		//outS1 += rowW  
		outS1 = ""; //   NONE," + fromWord
	}	
	
	outS1 = sortWordToRowByFreq(outS1) 
	
	return outS1, wordPrefixIndexList, wordSuffixIndexList
	
} // end of get_word_row_list
//--------------------------------------
func g04_bind_go_passToJs_betweenWordList_V3( maxNumWords int, fromWordPref string, js_function string, js_parm string, js_caller string) {
	
	outS1, _, _ := g04_get_word_row_list( maxNumWords, fromWordPref) 
	
	go_exec_js_functionPlus( js_function, outS1, js_parm, js_caller); 		
			
} // end of bind_go_passToJs_betweenWordList
//------------------

func word_to_row_onlyOneLemma(onlyThisLemma string, onlyIfExtr bool, onlyThisLevel string, xWordAlpha wordUnAlphaStruct, numRowsInW int) (bool, string)  {
	
	
	sw:= true
	var lastLemma lemmaStruct
	
	
	totNumRow:= numRowsInW
	if totNumRow < 1 {
		totNumRow = xWordAlpha.uTotExtrRow
	}
	
	ix2 := -1
	for _, ix1 := range xWordAlpha.uIxLemmaL {		
		lastLemma = lemmaSlice[ ix1 ]
		if lastLemma.leLemma == onlyThisLemma {
			ix2=ix1; 
			break
		}  	
	}	
	var x2Lemma, x2Tran, x2IxLemma string
	var ixL1 int
	
	if ix2 >=0 {
		x2Lemma		= xWordAlpha.uLemmaL[ix2] 
		ixL1        = xWordAlpha.uIxLemmaL[ix2]
		x2Tran      = lemmaSlice[ixL1].leTran
		x2IxLemma 	= strconv.Itoa( ixL1) //   fmt.Sprint( xWordAlpha.uIxLemmaL[ix2] )
	}	
	  
	return sw, xWordAlpha.uWord2 + separ1 + xWordAlpha.uWord0 + separ1 + 
		"ix" + separ1 + 
		strconv.Itoa(xWordAlpha.uIxUnW_fr) + separ1 + strconv.Itoa(xWordAlpha.uTotRow)  + separ1 + 
		x2Lemma       					+ separ1 + 
		x2Tran  						+ separ1 +  
		separ1 							+  
		lastLemma.lePara                + separ1 +  
		lastLemma.leExample             + separ1 +  
		strconv.Itoa(totNumRow) 		+ separ1 +  		
		xWordAlpha.uLearnedYN              + separ1 + 	
		"ixLemma" + separ1 + x2IxLemma 	+ separ1 + 	
		endOfLine 		
	
	
} // end of word_to_row_onlyOneLemma
//------------------------------------------------------
func sortWordToRowByFreq( str1 string) string {
	return str1	
}	
func VEROsortWordToRowByFreq( str1 string) string {	
	righe:= strings.Split(str1, endOfLine); 
	type rowWor struct {
		myKeyF   int
		myKeyW   string	
		myField []string
	}
	var oneW rowWor; 
	rowLis:= make([]rowWor,0, 10+len(righe) ) 
	for j3, oneRow:= range righe {
		if len(oneRow) < 1 {continue} 		
		//if len(oneRow) < 5 {fmt.Println( green("sortWord  oneRow=") + oneRow + "<==") }		
		oneW.myField = strings.Split(oneRow, separ1) 
		oneW.myKeyF, _ = strconv.Atoi( oneW.myField[4] )   // totRow
		oneW.myKeyW = oneW.myField[1]    // word0
		if (j3 < 10) {fmt.Println( green("sortWord  oneW ") + "keyF=", oneW.myKeyF , " keyW=", oneW.myKeyW, " fields=", oneW.myField) }	
		rowLis = append(rowLis, oneW)
	}
	//----------------
	sort.Slice(rowLis, func(i, j int) bool {
		if rowLis[i].myKeyF != rowLis[j].myKeyF {
			return rowLis[i].myKeyF > rowLis[j].myKeyF    // tot.row  in descending order
		} else {
			return rowLis[i].myKeyW < rowLis[j].myKeyW   
		}	
	})
	//-----------		
	outS2:=""; 
	for j3, oneW:= range rowLis {
		outS2 += strings.Join(oneW.myField, separ1) + endOfLine
		if (j3 < 10) { fmt.Println( green("sortato ") + "keyF=", oneW.myKeyF , " keyW=", oneW.myKeyW, " fields=", oneW.myField) }
	}	
	return outS2
	
} // end of sortWordToRowByFreq 

//-----------------------------------------


func word_to_row(onlyThisLemma string, onlyIfExtr bool, onlyThisLevel string, xWordAlpha wordUnAlphaStruct, numRowsInW int) (bool, string)  { 	
	
	//fmt.Println("\nXXXXXXX word_to_row ( onlyThisLevel=" + onlyThisLevel + 
	//  	"<==   onlyThisLevel=" + level_other +  "<==" + "  sw_list_Word_if_in_ExtrRow=" , sw_list_Word_if_in_ExtrRow); 
	separ1 := ";."
	sw:= true
	var lastLemma lemmaStruct
	
	if onlyIfExtr {
		if (sw_list_Word_if_in_ExtrRow) {
			if (xWordAlpha.uSwSelRowG == SEL_NO_EXTR_ROW) {   //    1 or 2: 1 SEL_EXTR_ROW, 2 SEL_NO_EXTR_ROW  
				//             xWordAlpha.uSwSelRowG è stato impostato nell'ultima esecuzione di "Lista le Righe" dove è stato scelto il gruppo di righe ed il range da-a
				sw = false
				//fmt.Println("word to row " , xWordAlpha.uWord2, " da ignorare")
				return sw, ""
			} 
		}
	}
	//fmt.Println("word to row " , xWordAlpha.uWord2, " \t\t\t XXXXXXXXXXXX (xWordAlpha.uTotExtrRow =" ,xWordAlpha.uTotExtrRow , " xxxxxxxxxxxxxxxxxxx   accettato")
	//-----------
	totNumRow:= numRowsInW
	if totNumRow < 1 {
		totNumRow = xWordAlpha.uTotExtrRow
	}	
	if onlyThisLemma != "" {
		return word_to_row_onlyOneLemma( onlyThisLemma, onlyIfExtr, onlyThisLevel, xWordAlpha, numRowsInW )  
	}	
	//---------------------------	
	outL:= ""
	for k1, ix2 := range xWordAlpha.uIxLemmaL {		
		lastLemma = lemmaSlice[ ix2 ]
		if lastLemma.leLemma != xWordAlpha.uLemmaL[k1] {
			fmt.Println("g04_bind... word_to_row ", "k1=", k1, " ix2=", ix2, " xWordAlpha.uLemmaL[ix2] =", xWordAlpha.uLemmaL[ix2], 
					" lemmaSlice[ix2].leLemma=", lastLemma.leLemma)
			return false, ""
		}  
		
		row:= xWordAlpha.uWord2 + separ1 + xWordAlpha.uWord0 + separ1 +
				"ix" + separ1 + 
				strconv.Itoa(xWordAlpha.uIxUnW_fr) + separ1 + 
				strconv.Itoa(xWordAlpha.uTotRow)   + separ1 +
				lastLemma.leLemmaOr           	+ separ1 + 
				lastLemma.leTran 				+ separ1 +  
				separ1 							+  
				lastLemma.lePara	 			+ separ1 +  
				lastLemma.leExample	 			+ separ1 +  
				strconv.Itoa(totNumRow) 		+ separ1 +  			
				xWordAlpha.uLearnedYN              + separ1 + 			
				"ixLemma" + separ1 + strconv.Itoa(ix2) + separ1 +  	
				endOfLine 	
		//fmt.Println(green("word_to_row "), row )	
		outL += row		
		//fmt.Println(" word_to_row ", row)  
	} 
	//--------------------	
	return sw, outL	
	
} // end of word_to_row 

//------------------------------------------------------

func intSliceToString(mySlice []int, wSep string) string {
	output := ""
	for _, v := range mySlice {
		output += (strconv.Itoa(v) + wSep)
	}
	return output
}
//------------------------------------------------------

func notFoundWord_row( fromWordCod string, fromWord string ) string {
	separ1:= ";."
	return fromWordCod + separ1 + fromWord + separ1 + 
			"ix" + separ1 + 
			strconv.Itoa(0) + separ1 + strconv.Itoa(0)  + separ1 + 
			fmt.Sprint( fromWord ) + separ1 + 
			fmt.Sprint( "_word_not_found_" ) + separ1 +  
			fmt.Sprint( "" ) + separ1 +  
			fmt.Sprint( "" ) + separ1 +  
			fmt.Sprint( "" ) + separ1 +  
			strconv.Itoa(0) + separ1 +  					
			strconv.Itoa(0) + separ1 + strconv.Itoa(0)  + separ1 + 
			"ixLemma" + separ1 + "" + separ1 + 	
			endOfLine 
					
} // end of notFoundWord_row

//---------------------------------------------------
