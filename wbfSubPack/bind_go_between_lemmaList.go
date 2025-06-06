package wbfSubPack

	import (
		"fmt"
		"strings"		
		//"strconv"
		//"regexp"
		//"sort"
	)

//-----------------------------------------
//-----------------------------------------

func bind_go_passToJs_betweenLemmaList_V3( maxNumLemmas int, fromLemmaPref string, js_function string) {
	
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
	
} // end of bind_go_passToJs_betweenLemmaList_V3

//------------------------------------------
func bind_go_passToJs_betweenLemmaList_V4( maxNumLemmas int, fromLemmaPref string, js_function string) {
	
	var outS1 string; 		
	fmt.Println( "bind_go_passToJs_betweenLemmaList 1 fromLemmaPref=>" + fromLemmaPref  )    
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
	
	fromIx, toIx := lookForLemma(fromLemma) // find the index of lemma in lemmaSlice 
	fmt.Println( green("bind_go_passToJs_betweenLemmaList"), " fromLemma=", fromLemma, " fromIx=", fromIx, " toIx=", toIx)
	if toIx < 0 {
		bind_go_passToJs_betweenWordList_V3(maxNumLemmas, fromLemmaPref, "js_go_showBetweenWordList")	
		return; 
	}
	
	fmt.Println( "bind_go_passToJs_betweenLemmaList 2  fromLemma=" + fromLemma + ", fromIx=", fromIx,  "  toIx=", toIx) 
	/**
	preLe:=""	
	for k:= 0; k < len( lemmaSlice); k++ {
		if k > (fromIx+10) { break}
		if lemmaSlice[k].leLemma < preLe {  fmt.Println("prova lemma[",k,"]=" , lemmaSlice[k].leLemma, "   fuori sequnza pre=", preLe ); break; }
		preLe = lemmaSlice[k].leLemma 
		if k < 20 {   fmt.Println("prova lemma[",k,"]=" , lemmaSlice[k].leLemma) }	
		if len(preLe) >=4 { if preLe[0:4] == "wass" {   fmt.Println("prova lemma[",k,"]=" , lemmaSlice[k].leLemma) }}	
	}
	**/
	//-------
	
	//-------------
	lenCk   :=0
	num1:=0
	
	fromIx2 := fromIx 
	
	//---------
	for k:= fromIx; k >=0; k-- {
		myLem1 := lemmaSlice[k]
		lenCk = len(myLem1.leLemma)
		if lenCk > lenFrom { lenCk = lenFrom }
		fmt.Println("vai all'indietro k=", k,  " fromLemma=", fromLemma, "  lenCk=", lenCk , " leLemma[0:lenCk]=",myLem1.leLemma[0:lenCk], " lemma=", myLem1.leLemma)
		
		if myLem1.leLemma[0:lenCk] < fromLemma {  break } 
		
		fmt.Println("     caricato ", " k=", k )
		fromIx2 = k
	}	
	//---------
	
	for k:= fromIx2; k < len( lemmaSlice); k++ {
		myLem1 := lemmaSlice[k]	
		fmt.Println("loop k=", k , " lemma=", myLem1.leLemma,   "   fromLemma=", fromLemma)	
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
	if num1 < 1 {
			//rowW:= notFoundLemma_row( fromLemma, fromLemma)		
			//outS1 += rowW 			
			//outS1 += "NONE," + fromLemma		
			bind_go_passToJs_betweenWordList_V3(maxNumLemmas, fromLemmaPref, "js_go_showBetweenWordList")	
			return; 
	}
	//-----------
	
	//fmt.Println( "bind_go_passToJs_betweenLemmaList 5  outS1=", outS1)
	
	go_exec_js_function( js_function, outS1 ); 	
	
} // end of bind_go_passToJs_betweenLemmaList_V4

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
		//fmt.Println("bind_go_passToJs_suffixLemmaList k=", k, " lemma=", myLem1.leLemma)
		rowW := build_one_lemma_row_word( myLem1 ) 
		
		outS1 += rowW 	
		num1++
		if num1 >= maxNumLemmas { break }
	}
	//------------------
	if num1 < 1 {		
			//outS1 += "NONE," + fromLemmaSuff		
			outS1 = ""			
	}

	go_exec_js_function( js_function, outS1 ); 	
		
} // end of  bind_go_passToJs_suffixLemmaList

//-------------------------------------------------