package wbfSubPack

	import (
		"fmt"
		"strings"		
		"strconv"
		//"slices"
		//"regexp"
		//"sort"
	)
//--------------------------------------------------------

func g04_bind_go_passToJs_prefixWordList( numWords int, wordPrefix string, js_function string) {
	
	g04_bind_go_passToJs_betweenWordList_V3( numWords, wordPrefix, js_function) 
			
} // end of bind_go_passToJs_prefixWordList

//-----------------------------------------

//--------------------------------------------------
func g04_get_word_row_list( maxNumWords int, fromWordPref string) (string, []int, []int) {
	
	//fmt.Println( red("\n 0 get_word_row_list ") + fromWordPref )  
	
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
	
	fromWordCod:= seqCode(fromWord)		
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
			lenCk   = len(wAlf.uWordSeq)
			if lenCk > lenFrom { lenCk = lenFrom }		
			if wAlf.uWordSeq[0:lenCk] < fromWordTarg {  break } 		
			fromIx2 = k
		}	
		//---------
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
			//fmt.Println("    wordSuffixIndexList sw=", sw, " wAlf=", wAlf.uWordSeq) 
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
	
	return outS1, wordPrefixIndexList, wordSuffixIndexList
	
} // end of get_word_row_list
//--------------------------------------
func g04_bind_go_passToJs_betweenWordList_V3( maxNumWords int, fromWordPref string, js_function string) {
	
	outS1, _, _ := g04_get_word_row_list( maxNumWords, fromWordPref) 
	
	go_exec_js_function( js_function, outS1 ); 		
			
} // end of bind_go_passToJs_betweenWordList
//------------------

func word_to_row_onlyOneLemma(onlyThisLemma string, onlyIfExtr bool, onlyThisLevel string, xWordAlpha wordUnAlphaStruct, numRowsInW int) (bool, string)  {
	
	separ1 := ";."
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
	  
	return sw, xWordAlpha.uWordSeq + separ1 + xWordAlpha.uWord2 + separ1 + 
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
		
		row:= xWordAlpha.uWordSeq + separ1 + xWordAlpha.uWord2 + separ1 + 
				"ix" + separ1 + 
				strconv.Itoa(xWordAlpha.uIxUnW_fr) + separ1 + 
				strconv.Itoa(xWordAlpha.uTotRow)   + separ1 +
				lastLemma.leLemma            	+ separ1 + 
				lastLemma.leTran 				+ separ1 +  
				separ1 							+  
				lastLemma.lePara	 			+ separ1 +  
				lastLemma.leExample	 			+ separ1 +  
				strconv.Itoa(totNumRow) 		+ separ1 +  			
				xWordAlpha.uLearnedYN              + separ1 + 			
				"ixLemma" + separ1 + strconv.Itoa(ix2) + separ1 +  	
				endOfLine 	
		outL += row		
		//fmt.Println(" word_to_row ", row)  
	} 
	//--------------------	
	return sw, outL	
	
} // end of word_to_row 

//------------------------------------------------------

func TOGLIword_to_row(onlyThisLemma string, onlyIfExtr bool, onlyThisLevel string, xWordF2 wordUnAlphaStruct, numRowsInW int) (bool, string)  {
	
	//fmt.Println("\nXXXXXXX word_to_row ( onlyThisLevel=" + onlyThisLevel + 
	//  	"<==   onlyThisLevel=" + level_other +  "<==" + "  sw_list_Word_if_in_ExtrRow=" , sw_list_Word_if_in_ExtrRow); 
	separ1 := ";."
	sw:= true
	var lastLemma lemmaStruct
	
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
		for _, ix1 := range xWordF2.uIxLemmaL {		
			lastLemma = lemmaSlice[ ix1 ]
			if lastLemma.leLemma == onlyThisLemma {
				ix2=ix1; 
				break
			}  	
		}		
		if ix2 >=0 {			
			fmt.Println("word to row 1 " , xWordF2.uWord2, " ix2=", ix2, " onlyIfExtr=", onlyIfExtr,  
				"    WordF2.uTotExtrRow =" ,xWordF2.uTotExtrRow, 
			  "	lastLemma.lePara=",  lastLemma.lePara)
			
			ixL1 := xWordF2.uIxLemmaL[ix2]
	
			return sw, xWordF2.uWordSeq + separ1 + xWordF2.uWord2 + separ1 + 
				"ix" + separ1 + 
				strconv.Itoa(xWordF2.uIxUnW_fr) + separ1 + strconv.Itoa(xWordF2.uTotRow)  + separ1 + 
				xWordF2.uLemmaL[ix2]              + separ1 + 
				lemmaSlice[ixL1].leTran           + separ1 +  
				separ1 +  
				lastLemma.lePara                  + separ1 +  
				lastLemma.leExample               + separ1 +  
				strconv.Itoa(totNumRow) + separ1 +  		
				xWordF2.uLearnedYN                + separ1 + 	
				"ixLemma" + separ1 + fmt.Sprint( xWordF2.uIxLemmaL[ix2] ) + separ1 + 	
				endOfLine 		
		}
	}
	
	//---------------------------	
var sw2 = (xWordF2.uWord2 == "personen")
	lisPara:= ""
	lisExa := ""
	for _, ix2 := range xWordF2.uIxLemmaL {		
		lastLemma = lemmaSlice[ ix2 ]
		lisPara = wSep + lastLemma.lePara	
		lisExa  = wSep + lastLemma.leExample	
		if sw2 {  fmt.Println( "1  ", xWordF2.uWord2, " ix2=", ix2, " lisPara=", lisPara," lisExa=", lisExa )  	}	
	}
	if len(lisPara) > 0 {
		lisPara = lisPara[ len(wSep):]
		lisExa  = lisExa[  len(wSep):]
		if sw2 {  fmt.Println( "2  ", xWordF2.uWord2, " lisPara=", lisPara)  	}	
	}		
	if sw2 {  fmt.Println( "3  ", xWordF2.uWord2,  " lisPara=", lisPara)  	}	
	
	return sw, xWordF2.uWordSeq + separ1 + xWordF2.uWord2 + separ1 + 
				"ix" + separ1 + 
				strconv.Itoa(xWordF2.uIxUnW_fr) + separ1 + strconv.Itoa(xWordF2.uTotRow)  + separ1 +
				fmt.Sprint( strings.Join(xWordF2.uLemmaL,  wSep)  ) + separ1 + 
				//fmt.Sprint( strings.Join(xWordF2.uTranL,   wSep)  ) + separ1 +  
				listStringLemmaSlice_Tran(xWordF2) + separ1 +  
				separ1 +  
				lisPara + separ1 +  
				lisExa  + separ1 +  
				strconv.Itoa(totNumRow) + separ1 +  			
				xWordF2.uLearnedYN              + separ1 + 			
				"ixLemma" + separ1 + intSliceToString( xWordF2.uIxLemmaL,wSep )  + separ1 + 		
				endOfLine 
				
				//fmt.Sprint( strings.Join(xWordF2.uPara,    wSep)  ) + separ1 +  
				//fmt.Sprint( strings.Join(xWordF2.uExample, wSep)  ) + separ1 +  
				 
} // end of TOGLIword_to_row 
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
