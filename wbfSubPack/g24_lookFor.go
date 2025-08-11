package wbfSubPack

import (  
	"fmt"
    "strings"
)
//-----------------------------
func lookForLemma(lemmaTarg string, _ int) (int, int) {
	  
	// find 2 indices of the 2 words nearest to the word to find 
	/**
				
			//---
			type lemmaStruct struct {
				leLemma     string
				leNumWords  int 
				leFromIxLW  int 
				leToIxLW    int  
				leTran      string 
				leLevel     string  
				lePara      string  
				leExample   string  	
			} 	
			var lemmaSlice       [] lemmaStruct         // lemma , translation 
	***/
	
	low   := 0
	high  := len(lemmaSlice) - 1	
	maxIx := high; 
	/**
	if lookFromLemma_LastIx >=0 {
		if lookFromLemma_LastIx >= len(lemmaSlice) {
			lookFromLemma_LastIx=0
		} else { 
			if lemmaSlice[lookFromLemma_LastIx].leLemma < lemmaTarg {  low = lookFromLemma_LastIx }
		}
		fmt.Println(red("lookForLemma "),  lemmaTarg,  " lookFromLemma_LastIx =", lookFromLemma_LastIx )
	}
	numLoop:=0
	**/
	//----
	for low <= high{
		//numLoop++
		median := (low + high) / 2
		if lemmaSlice[median].leLemma < lemmaTarg {  
			low = median + 1
		}else{
			high = median - 1
		}
	} 
	//---
	fromIx:= low; toIx := high; 
	if fromIx > toIx { fromIx = high; toIx = low;}
	if fromIx < 0 { fromIx=0} 
	if toIx  > maxIx { toIx = maxIx}
	
	//lookFromLemma_LastIx = fromIx
	
	//fmt.Println(red("     lookForLemma "),  lemmaTarg,   " numero cicli per trovarlo =", numLoop)
	
	return fromIx, toIx	

} // end of LookForLemma

//------------------------------------------

func lookForAllLemmas(  wordToFind string, lastIx int) []int {

	wordToFindCod:= seqCode( wordToFind )	
	ixFoundList := lookForAllLemmas2(  wordToFindCod, lastIx ) 
	if len(ixFoundList) > 0 {
		return ixFoundList
	}
	
	inp1:= strings.ReplaceAll( 
				strings.ReplaceAll( 
						strings.ReplaceAll( 
							strings.ReplaceAll( wordToFind, "ae","ä"),  
							"oe","ö"), 
						"ue","ü"), 		
				"ss","ß") 	
				
	wordToFindCod2 := seqCode( inp1 )
	ixFoundList2 := lookForAllLemmas2(  wordToFindCod2, lastIx ) 
	if len(ixFoundList2) == 0 {
		ixFoundList2 = append( ixFoundList,  -1)	// if lemma is missing use the original word to find 
		lemmaNotFoundList = append( lemmaNotFoundList, wordToFind ) 
	} 	
	return ixFoundList2 
	
} // end of lookForAllLemmas
//------------------------------------------
func lookForAllLemmas2(  wordToFindCod string, lastIx int) []int {

	//wordToFindCod:= seqCode( wordToFind )

	// get the index of a word in word-lemma dictionary (-1 if not found)  
	var ixFoundList = make( []int, 0,0) 
	
	if len(wordLemmaPair) == 0 { return ixFoundList}
	
	fromIxX, toIx := lookForWordLemmaPair(wordToFindCod, lastIx)
	
	if toIx < 0 { return ixFoundList }
	
	fromIx:= fromIxX
	
	for k:= fromIxX; k >= 0; k-- {
		if wordLemmaPair[k].lWordSeq < wordToFindCod { break }
		fromIx = k
	}
	for k:= fromIx; k < len(wordLemmaPair); k++ {
		if wordLemmaPair[k].lWordSeq == wordToFindCod {
			ixFoundList = append( ixFoundList, k) ; //    wordLemmaPair[k].lLemma )	
		} else {
			if wordLemmaPair[k].lWordSeq > wordToFindCod { break }
		}
	} 
			

	//fmt.Println("lookForLemma( ==>" + wordToFind + "<== lemmaList=" , lemmaList, "   numLemmaDict=" , numLemmaDict)
	
	return ixFoundList 	
	
} // end of lookForAllLemmas2

//-----------------------------



func lookForAllParadigma( lemma3 string, lastIx int ) (int, int) {
		
	fromIxX, toIxX := lookForParadigma(lemma3, lastIx )
	if toIxX < 0 { return -1, -99999 }
	
	
	minIx:= -1; maxIx := -1 
	
	fromIx:= fromIxX
	
	// get the smaller index  with the right lemma  
	
	for k:= fromIx; k >= 0; k-- {
		if lemma_para_list[k].p_lemma == lemma3 {
			minIx=k
		} else { 
			if lemma_para_list[k].p_lemma < lemma3 { break }
		}
	}
	
	// get the maximum index with the right lemma
	maxIx = minIx
	for k:= fromIx+1; k < len( lemma_para_list); k++ {
		if lemma_para_list[k].p_lemma == lemma3 {
			maxIx = k; 
			if minIx < 0 { minIx = k}
		}  else { 
			if lemma_para_list[k].p_lemma > lemma3 { break }
		}
	}	
	if maxIx < 0 { maxIx = -9999; minIx = -1 } 
	return minIx, maxIx
	
} // end of lookForAllParadigma

//-------------------------------------
func lookForParadigma(lemmaToFind string, lastIx int) (int, int) {

	// find 2 indices of the 2 words nearest to the word to find 
	
	low   := 0
	high  := len(lemma_para_list) - 1	
	maxIx := high; 
	if lastIx >=0 {
		if lemma_para_list[lastIx].p_lemma < lemmaToFind {  low = lastIx }
	}
	//----
	for low <= high{
		median := (low + high) / 2
		if lemma_para_list[median].p_lemma < lemmaToFind {  
			low = median + 1
		}else{
			high = median - 1
		}
	} 
	//---
	fromIx:= low; toIx := high; 
	if fromIx > toIx { fromIx = high; toIx = low;}
	if fromIx < 0 { fromIx=0} 
	if toIx  > maxIx { toIx = maxIx}
	
	return fromIx, toIx	
	
} // end of lookForParadigma

//-----------------------------------------

func lookForAllTran ( lemma30 string , lastIx int) int {
	lemma3Cod:= seqCode(lemma30)	
	fromIxX, toIxX := lookForTranslation( lemma3Cod , lastIx)
	if toIxX < 0 { return -1 }
	
	
	z:=-1
	fromIx:= fromIxX
	for k:= fromIxX; k >= 0; k-- {
		if dictLemmaTran[k].dL_lemmaSeq == lemma3Cod {
			z=k
			break	
		}
		if dictLemmaTran[k].dL_lemmaSeq < lemma3Cod { break }
		fromIx = k
	}
	if z < 0 {
		for k:= fromIx; k < len( dictLemmaTran); k++ {
			if dictLemmaTran[k].dL_lemmaSeq == lemma3Cod {
				z=k
				break
			}
		} 
	}
	
	if z < 0 { return -1 }
	
	return z 
	
} // end of lookForAllTran

//-----------------------------

func lookForTranslation(lemmaToFindCod string, lastIx int) (int, int) {
	if len(dictLemmaTran) == 0 { return -1, -999 }
	// find 2 indices of the 2 words nearest to the word to find 
	
	low   := 0
	high  := len(dictLemmaTran) - 1	
	maxIx := high; 
	if lastIx >=0 {
		if dictLemmaTran[lastIx].dL_lemmaSeq < lemmaToFindCod {  low = lastIx }
	}
	//----
	for low <= high{
		median := (low + high) / 2
		if dictLemmaTran[median].dL_lemmaSeq < lemmaToFindCod {  
			low = median + 1
		}else{
			high = median - 1
		}
	} 
	//---
	fromIx:= low; toIx := high; 
	if fromIx > toIx { fromIx = high; toIx = low;}
	if fromIx < 0 { fromIx=0} 
	if toIx  > maxIx { toIx = maxIx}
	
	return fromIx, toIx	

} // end of lookForTranslation

//-----------------------------

func lookForWordLemmaPair(wordToFindCod string, lastIx int) (int, int) {
	
	// find 2 indices of the 2 words nearest to the word to find 
	
	low   := 0
	high  := numLemmaDict - 1	
	maxIx := high; 
	if high < 1 { return -1, -1 } 
	
	if lastIx >=0 {
		if wordLemmaPair[lastIx].lWordSeq < wordToFindCod {  low = lastIx }
	}
	//----
	for low <= high{
		median := (low + high) / 2
		if wordLemmaPair[median].lWordSeq < wordToFindCod {  
			low = median + 1
		}else{
			high = median - 1
		}
	} 
	//---
	fromIx:= low; toIx := high; 
	if fromIx > toIx { fromIx = high; toIx = low;}
	
	if fromIx < 0 { fromIx=0} 
	
	if toIx  > maxIx { toIx = maxIx}
	return fromIx, toIx	

} // end of lookForWordLemmaPair

//---------------------------------------------------

//--------------
func lookForWordInUniqueAlpha(wordCoded string, lastIx int) (int, int) {
	
	// find 2 indices of the 2 words nearest to the word to find 	
	
	low   := 0
	high  := numberOfUniqueWords - 1	
	maxIx := high; 
	
	if lastIx >=0 {
		if uniqueWordByAlpha[lastIx].uWordSeq < wordCoded {  low = lastIx }
	}
	
	//----
	for low <= high{
		median := (low + high) / 2
		if median >= len(uniqueWordByAlpha) {
			fmt.Println("errore in lookForWordInUniqueAlpha: median=", median , "     len(uniqueWordByAlpha)=" ,  len(uniqueWordByAlpha) )
		}
		if uniqueWordByAlpha[median].uWordSeq < wordCoded {
			low = median + 1
		}else{
			high = median - 1
		}
	} 
	//---
	fromIx:= low; toIx := high; 
	if fromIx > toIx { fromIx = high; toIx = low;}
	if fromIx < 0 { fromIx=0} 
	if toIx  > maxIx { toIx = maxIx}
	return fromIx, toIx	

} // end of lookForWordInUniqueAlpha

//-----------------------------

func searchAllWordWithPrefixInAlphaList(  wordPref string, lastIx int) (int, int) {
	
	// get the indicies of the first and the last word beginning with the required prefix (-1,-1 if not found)  
	
	wordPref = strings.ToLower(strings.TrimSpace( wordPref));  
	wordCodPref:= seqCode(wordPref)
	
	lenPref:= len(wordPref); 
	ixTo := -1; ixFrom:= -1;	
	
	if lenPref == 0 { return ixFrom, ixTo }
	
	ix1, ix2:= lookForWordInUniqueAlpha(wordCodPref, lastIx)	
	
	/***
	fmt.Println("ANTONIO SEARCH ALPHA wordPref=" + wordPref + " wordCodPref=" +  wordCodPref + " ix1=", ix1, " ix2=", ix2) 
	for k:= ix1; k <= ix2; k++ {
		if k < 0 { continue}
		fmt.Println("ANTONIO SEARCH ALPHA k=", k , " ==>" , uniqueWordByAlpha[k])
	}	
	***/
	
	wA :=""
	spaceFill := "                                                          ";  
	//-----------
	for k:= ix1; k >= 0; k-- {
		wA =  uniqueWordByAlpha[k].uWordSeq + spaceFill
		if wA[0:lenPref] < wordCodPref { break; }
		ixFrom = k; 
	}  
	
	if (ixFrom >=0) { ixTo = ixFrom; }  //  se ixFrom è valido, deve essere valido anche ixTo   
	
	for k:= ix2; k < numberOfUniqueWords; k++ {
		wA =  uniqueWordByAlpha[k].uWordSeq + spaceFill  
		if wA[0:lenPref] > wordCodPref { break; }
		ixTo = k; 
		if (ixFrom < 0) {ixFrom = ixTo;}  //  se ixTo è valido, deve essere valido anche ixFrom   
	}  
	return ixFrom, ixTo 
	
} // end of searchAllWordWithPrefixInAlphaList

//------------------------------------------------------
func testGenericWord(pref string ) {
	//fmt.Println( "cerca tutte le parole che iniziano con " + pref);	
	var xWordF wordUnAlphaStruct;  
	
	from1, to1 := searchAllWordWithPrefixInAlphaList( pref,0 )
	if (to1 < 0) {
		fmt.Println( "nessuna parola che inizia con " , pref); 
	} else {	
		for i:=from1; i <=to1; i++ {	
			xWordF =  uniqueWordByAlpha[i] 
			fmt.Println( "trovato ", xWordF.uWord2, " ix=", i,   " totRow=", xWordF.uTotRow, " uIxFromWord_al=", xWordF.uIxFromWord_al); 
		}
	}
	fmt.Println( "-------------------------------------------------\n" ); 
	return 
}
//---------------------


//------------------------------------------------

func lookForLemmaWord(lemmaCode string, lastIx int) (int, int) {
	
	// find 2 indices of the 2 words nearest to the word to find 
	
	low   := 0
	high  := len(lemma_word_ix) - 1	
	maxIx := high; 
	
	if lastIx >=0 {
		if lemma_word_ix[lastIx].lw_lemmaSeq < lemmaCode {  low = lastIx }
	}
	
	//----
	for low <= high{
		median := (low + high) / 2
		if lemma_word_ix[median].lw_lemmaSeq < lemmaCode {  
			low = median + 1
		}else{
			high = median - 1
		}
	} 
	//---
	fromIx:= low; toIx := high; 
	if fromIx > toIx { fromIx = high; toIx = low;}
	if fromIx < 0 { fromIx=0} 
	if toIx  > maxIx { toIx = maxIx}
	return fromIx, toIx	

} // end of LookForLemmaWord

//-----------------------------------------

func binarySearch_string(myStrList []string, value string) int {
	// SEARCH STRING ARRAY RETURN INDEX OF -1  
    low := 0
    high := len(myStrList)-1
	mid:=0
    for low <= high { 
        mid = (low+high)/2
        if myStrList[mid] > value {
			high = mid-1
		}  else{
			if myStrList[mid] < value { 
				low = mid+1 
			} else { 
				return mid 
			}
		}
	}	
    return -1
} // end of binarySearch	
//-------------------------