package wbfSubPack

	import (
		"fmt"
		"strings"		
		//"strconv"
		//"regexp"
		//"sort"
	)

//--------------------------------------
func g03_bind_go_passToJs_betweenLemmaList_V3( maxNumLemmas int, fromLemmaPref string, js_function string) {
	fmt.Println("func ", green("g03_bind_go_passToJs_betweenLemmaList_V3"), " -->   g03_get_lemma_row_list")
	outS1 := g03_get_lemma_row_list( maxNumLemmas, fromLemmaPref) 
	
	go_exec_js_function( js_function, outS1 ); 		
			
} // end of bind_go_passToJs_betweenWordList
//------------------

//--------------------------------------------------
func g03_get_lemma_row_list( maxNumLemmas int, fromLemmaPref string) string {
	
	fmt.Println( red("\n 0 get_lemma_row_list ") + fromLemmaPref )  
	
	
	var outS1 string; 	
	//----------------------------------------------------------------------
	fromLemma   := strings.ToLower(strings.TrimSpace( fromLemmaPref));  
	lenFrom := len(fromLemma) 	
	sw_oneLemma := false 
	
	lemmaPrefixIndexList:= make([]int,0, 2*maxNumLemmas)
	lemmaSuffixIndexList:= make([]int,0, 2*maxNumLemmas)	
	num1:=0	
	
	if fromLemma == "" { 
		//go_exec_js_function( js_function, "");
		return ""
	}
	//----------------------		
	
	if fromLemma[0:1] == "-" {
		if  fromLemma[lenFrom-1:] == "-" { //     any lemma which contains the string in fromLemma  eg.   -mili-  --> familie 
			// this case cannot be of any use, ignore it    
			return ""			
		} else {//a suffix request, wanted all lemmas ending with the string in fromLemma.    eg.  -en --> gehen, haben, etc. 				
			lemmaSuffixIndexList = g34_getListInverseLemmaIndex( fromLemma[1:] , sw_oneLemma, maxNumLemmas) 
		}	
	} else {
		if fromLemma[lenFrom-1:] == "-" {// a prefix request,  wanted all lemmas beginning with the string in fromLemma. eg.  geh-    --> geht, gehen, etc.    
			fromLemma = fromLemma[:lenFrom-1]  
		} else { 
			sw_oneLemma = true    // fromLemma contains the only lemma to look for,  it's not a prefix neither a suffix 
		}
		//lemmaPrefixIndexList = getListLemmaIndex(fromLemma, sw_oneLemma, maxNumLemmas) 
		lemmaPrefixIndexList = g34_getListDirectLemmaIndex(fromLemma, sw_oneLemma, maxNumLemmas) 
	}
	
	fmt.Println("g03_get_lemma_row_list   lemmaPrefixIndexList=", lemmaPrefixIndexList)   
	fmt.Println("g03_get_lemma_row_list   lemmaSuffixIndexList=", lemmaSuffixIndexList)   
	
	listA := make([]int, 0, len(lemmaPrefixIndexList) + len( lemmaSuffixIndexList ) ) 
	listA = append(listA,lemmaPrefixIndexList...) 
	listA = append(listA,lemmaSuffixIndexList...) 	

	totNumLemmas:= len(listA)	
	
	sw1:= ((fromLemma == "familie") || (fromLemma == "gehen") )
	if sw1 { fmt.Println("bind_go_passToJs_betweenLemmaList_V3 ", " num indici=", len(listA), " ", listA) }
	var rowW string
	for j1, ix2:= range listA {
		if sw1 {fmt.Println("\n", green("LEMMA lemmaSlice"),"[", ix2, "]=", lemmaSlice[ix2] ) }
		
		if totNumLemmas > 1 { 
			rowW = g05_build_one_lemma_row_word_ManyLemma(   lemmaSlice[ix2], totNumLemmas)		
		} else {		
			rowW = g05_build_one_lemma_row_word_OnlyOneLemma(lemmaSlice[ix2], totNumLemmas)
		}
		if sw1 {fmt.Println("loop", j1, "\n\t", strings.ReplaceAll(rowW, "\n", "\n\t") ) }
		outS1 += rowW
		num1++
		if num1 >= maxNumLemmas { break }	
	}	
	return outS1
	
} // end of bind_go_passToJs_betweenLemmaList_V3

//------------------------------------------
//-------------------------------------------------------

func notFoundLemma_row( fromWordCod string, fromWord string ) string {
	
		return  "" + ";." + "" + ";." + 
				"ix" + ";." + 
				"-1" + ";." + "-1" + ";." +
				"_lemma_not_found_" + ";." + 
				""    + ";." +  
				""   + ";." +  
				""   + ";." +  
				""   + ";." +  
				"0"  + ";." +  					
				"0"  + ";." + "0" + ";." + 				
				"ixLemma" + ";." + "-1" + ";." + 		
				endOfLine 
					
} // end of notFoundWord_row

//--------------------------------------------

 func TOGLIgetListLemmaIndex(fromLemma string, sw_oneLemma bool, maxNum int) []int { 
	
	indexList:= make([]int,0, 2*maxNum)	
	
	lenFrom := len(fromLemma)
	
	from1, to1 := lookForLemma(fromLemma,0) // find the index of lemma in lemmaSlice 
	
	if to1 < from1 { from1 = to1}	
	if from1 < 0 {from1=0}	
	
	fromIx2:= from1 
	numOut :=0
	lenCk  :=0
	//----------
	// non sono sicuro che from1, to1 contengano gli elementi richiesti e nemmeno se non ce ne sia qualcuno prima o dopo  
	// vado indietro e mi fermo quand trovo elementi diversi	
	
	for k:= from1; k >=0; k-- {
		myLem1 := lemmaSlice[k]			
		lenCk = len(myLem1.leLemma)
		if lenCk > lenFrom { lenCk = lenFrom}	
		if myLem1.leLemma[0:lenCk] < fromLemma { break } 
		fromIx2 = k
	}	
	//---------
	if sw_oneLemma {
		for k:= fromIx2; k < len( lemmaSlice); k++ {
			myLem1 := lemmaSlice[k]		
			if myLem1.leLemma > fromLemma { break } 
			if myLem1.leLemma < fromLemma { continue }	
			if myLem1.leNumWords == 0 { continue }		// nessuna word del testo ha questo lemma 	
			indexList = append(indexList, k)	
			numOut ++
			if numOut >= maxNum {break} 
		}
		return indexList	
	}
	//-----------------------------------
	
	// looking for prefixes 
	
	for k:= fromIx2; k < len( lemmaSlice); k++ {
		myLem1 := lemmaSlice[k]	
		lenCk = len(myLem1.leLemma)
		if lenCk > lenFrom { lenCk = lenFrom}			
		if myLem1.leLemma[0:lenCk] > fromLemma { break } 	
		if myLem1.leLemma[0:lenCk] < fromLemma { continue} 	
		if myLem1.leNumWords == 0 { continue }		// nessuna word del testo ha questo lemma	
		indexList = append(indexList, k)	
		numOut ++
		if numOut >= maxNum {break} 
	} // end no suffix 
	
	return indexList
	
 } // end of getListLemmaIndex	
 
 //------------------------------------------