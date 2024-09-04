package wbfSubPack

	import (
		//"fmt"
		"strings"		
		//"strconv"
		//"regexp"
		//"sort"
	)

//-----------------------------------------

func bind_go_passToJs_betweenLemmaList( maxNumLemmas int, fromLemmaPref string, js_function string) {
	
	var outS1 string; 		
	//fmt.Println( "bind_go_passToJs_betweenLemmaList 1 fromLemmaPref=>" + fromLemmaPref  + " toLemmaPref=" + toLemmaPref )    
	//-----------------------------
	fromLemma   := strings.ToLower(strings.TrimSpace( fromLemmaPref )); 	
	lenFrom := len(fromLemma)  
	
	sw_oneWord := false 
	
	if fromLemma == "" {
		go_exec_js_function( js_function, "")
		return	
	}
	if fromLemma[0:1] == "-" {
		// this is a suffix request 
		bind_go_passToJs_suffixLemmaList( maxNumLemmas, fromLemmaPref[1:], js_function)
		return
	}
	if fromLemma[lenFrom-1:] == "-" {
		// prefisso 
		fromLemma = fromLemma[:lenFrom-1] 
	} else {
		sw_oneWord = true    // fromLemma contains the only word to look for,  it's not a prefix neither a suffix 
	}
	
	lenFrom = len(fromLemma)  
	
	fromIx, _ := lookForLemma(fromLemma) // find the index of lemma in lemmaSlice 
	
	lenCk   :=0
	//-------
	//fmt.Println( "bind_go_passToJs_betweenLemmaList 2   fromIx=", fromIx,  "  toIx=", toIx) 
	//-------------
	
	num1:=0
	
	fromIx2 := fromIx 
	
	//---------
	for k:= fromIx; k >=0; k-- {
		myLem1 := lemmaSlice[k]
		lenCk = len(myLem1.leLemma)
		if lenCk > lenFrom { lenCk = lenFrom }
		//fmt.Println("vai all'indietro k=", k,  " fromLemma=", fromLemma, "  lenCk=", lenCk , " leLemma[0:lenCk]=",myLem1.leLemma[0:lenCk], " lemma=", myLem1.leLemma)
		
		if myLem1.leLemma[0:lenCk] < fromLemma {  break } 
		
		//fmt.Println("     caricato ", " k=", k )
		fromIx2 = k
	}	
	//---------
	
	for k:= fromIx2; k < len( lemmaSlice); k++ {
		myLem1 := lemmaSlice[k]		
		if sw_oneWord {
			if myLem1.leLemma != fromLemma {
				if myLem1.leLemma > fromLemma { break } 
				continue					
			} 
		} else {
			lenCk = len(myLem1.leLemma)
			if lenCk > lenFrom { lenCk = lenFrom }
			if myLem1.leLemma[0:lenCk] < fromLemma { continue} 		
			if myLem1.leLemma[0:lenCk] > fromLemma { break } 		
		}
		rowW := build_one_lemma_row_word( myLem1 ) 
		
		outS1 += rowW 	
		num1++
		if num1 >= maxNumLemmas { break }
	}
	//------------------
	if num1 < 0 {
			//rowW:= notFoundLemma_row( fromLemma, fromLemma)		
			//outS1 += rowW 			
			outS1 += "NONE," + fromLemma			
	}
	//-----------
	
	//fmt.Println( "bind_go_passToJs_betweenLemmaList 5  outS1=", outS1)
	
	go_exec_js_function( js_function, outS1 ); 	
	
} // end of bind_go_passToJs_betweenLemmaList

//-------------------------------------------------------
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


func bind_go_passToJs_suffixLemmaList( maxNumLemmas int, fromLemmaSuff string, js_function string) {
	
	var outS1 string;
	
	listInverseLemmaIndex := getListInverseLemmaIndex( fromLemmaSuff, maxNumLemmas) 
	
	num1:=0	
	for _,k:= range listInverseLemmaIndex { 
		myLem1 := lemmaSlice[k]

		rowW := build_one_lemma_row_word( myLem1 ) 
		
		outS1 += rowW 	
		num1++
		if num1 >= maxNumLemmas { break }
	}
	//------------------
	if num1 < 0 {		
			outS1 += "NONE," + fromLemmaSuff			
	}

	go_exec_js_function( js_function, outS1 ); 	
		
} // end of  bind_go_passToJs_suffixLemmaList

//-------------------------------------------------