package wbfSubPack

import (  
	"fmt"
    "strings"
)
//-----------------------------

func lookForLemma(lemmaTarg string) (int, int) {
	
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
	
	//----
	for low <= high{
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
	return fromIx, toIx	

} // end of LookForLemma

//------------------------------------------

func lookForAllLemmas(  wordToFind string) []int {

	wordToFindCod:= newCode( wordToFind )	
	ixFoundList := lookForAllLemmas2(  wordToFindCod ) 
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
				
	wordToFindCod2 := newCode( inp1 )
	ixFoundList2 := lookForAllLemmas2(  wordToFindCod2 ) 
	if len(ixFoundList2) == 0 {
		ixFoundList2 = append( ixFoundList,  -1)	// if lemma is missing use the original word to find 
		lemmaNotFoundList = append( lemmaNotFoundList, wordToFind ) 
	} 	
	return ixFoundList2 
	
} // end of lookForAllLemmas
//------------------------------------------
func lookForAllLemmas2(  wordToFindCod string) []int {

	//wordToFindCod:= newCode( wordToFind )

	// get the index of a word in word-lemma dictionary (-1 if not found)  
	var ixFoundList = make( []int, 0,0) 
	
	if len(wordLemmaPair) == 0 { return ixFoundList}
	
	fromIxX, toIx := lookForWordLemmaPair(wordToFindCod)
	
	if toIx < 0 { return ixFoundList }
	
	fromIx:= fromIxX
	
	for k:= fromIxX; k >= 0; k-- {
		if wordLemmaPair[k].lWordCod < wordToFindCod { break }
		fromIx = k
	}
	for k:= fromIx; k < len(wordLemmaPair); k++ {
		if wordLemmaPair[k].lWordCod == wordToFindCod {
			ixFoundList = append( ixFoundList, k) ; //    wordLemmaPair[k].lLemma )	
		} else {
			if wordLemmaPair[k].lWordCod > wordToFindCod { break }
		}
	} 
			

	//fmt.Println("lookForLemma( ==>" + wordToFind + "<== lemmaList=" , lemmaList, "   numLemmaDict=" , numLemmaDict)
	
	return ixFoundList 	
	
} // end of lookForAllLemmas2

//-----------------------------



func lookForAllParadigma( lemma3 string ) (int, int) {
		
	fromIxX, toIxX := lookForParadigma(lemma3 )
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
func lookForParadigma(lemmaToFind string) (int, int) {

	// find 2 indices of the 2 words nearest to the word to find 
	
	low   := 0
	high  := len(lemma_para_list) - 1	
	maxIx := high; 
	
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

func lookForAllTran ( lemma30 string ) int {
	lemma3Cod:= newCode(lemma30)	
	fromIxX, toIxX := lookForTranslation( lemma3Cod )
	if toIxX < 0 { return -1 }
	
	
	z:=-1
	fromIx:= fromIxX
	for k:= fromIxX; k >= 0; k-- {
		if dictLemmaTran[k].dL_lemmaCod == lemma3Cod {
			z=k
			break	
		}
		if dictLemmaTran[k].dL_lemmaCod < lemma3Cod { break }
		fromIx = k
	}
	if z < 0 {
		for k:= fromIx; k < len( dictLemmaTran); k++ {
			if dictLemmaTran[k].dL_lemmaCod == lemma3Cod {
				z=k
				break
			}
		} 
	}
	
	if z < 0 { return -1 }
	
	return z 
	
} // end of lookForAllTran

//-----------------------------

func lookForTranslation(lemmaToFindCod string) (int, int) {

	// find 2 indices of the 2 words nearest to the word to find 
	
	low   := 0
	high  := len(dictLemmaTran) - 1	
	maxIx := high; 
	
	//----
	for low <= high{
		median := (low + high) / 2
		if dictLemmaTran[median].dL_lemmaCod < lemmaToFindCod {  
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

func lookForWordLemmaPair(wordToFindCod string) (int, int) {
	
	// find 2 indices of the 2 words nearest to the word to find 
	
	low   := 0
	high  := numLemmaDict - 1	
	maxIx := high; 
	
	if high < 1 { return -1, -1 } 
	
	//----
	for low <= high{
		median := (low + high) / 2
		if wordLemmaPair[median].lWordCod < wordToFindCod {  
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
func lookForWordInUniqueAlpha(wordCoded string) (int, int) {
	
	// find 2 indices of the 2 words nearest to the word to find 	
	
	low   := 0
	high  := numberOfUniqueWords - 1	
	maxIx := high; 
	//----
	for low <= high{
		median := (low + high) / 2
		if median >= len(uniqueWordByAlpha) {
			fmt.Println("errore in lookForWordInUniqueAlpha: median=", median , "     len(uniqueWordByAlpha)=" ,  len(uniqueWordByAlpha) )
		}
		if uniqueWordByAlpha[median].uWordCod < wordCoded {
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

func searchAllWordWithPrefixInAlphaList(  wordPref string) (int, int) {
	
	// get the indicies of the first and the last word beginning with the required prefix (-1,-1 if not found)  
	
	wordPref = strings.ToLower(strings.TrimSpace( wordPref));  
	wordCodPref:= newCode(wordPref)
	
	lenPref:= len(wordPref); 
	ixTo := -1; ixFrom:= -1;	
	
	if lenPref == 0 { return ixFrom, ixTo }
	
	ix1, ix2:= lookForWordInUniqueAlpha(wordCodPref)	
	
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
		wA =  uniqueWordByAlpha[k].uWordCod + spaceFill
		if wA[0:lenPref] < wordCodPref { break; }
		ixFrom = k; 
	}  
	
	if (ixFrom >=0) { ixTo = ixFrom; }  //  se ixFrom è valido, deve essere valido anche ixTo   
	
	for k:= ix2; k < numberOfUniqueWords; k++ {
		wA =  uniqueWordByAlpha[k].uWordCod + spaceFill  
		if wA[0:lenPref] > wordCodPref { break; }
		ixTo = k; 
		if (ixFrom < 0) {ixFrom = ixTo;}  //  se ixTo è valido, deve essere valido anche ixFrom   
	}  
	return ixFrom, ixTo 
	
} // end of searchAllWordWithPrefixInAlphaList

//------------------------------------------------------
func testGenericWord(pref string ) {
	//fmt.Println( "cerca tutte le parole che iniziano con " + pref);	
	var xWordF wordIxStruct;  
	
	from1, to1 := searchAllWordWithPrefixInAlphaList( pref )
	if (to1 < 0) {
		fmt.Println( "nessuna parola che inizia con " , pref); 
	} else {	
		for i:=from1; i <=to1; i++ {	
			xWordF =  uniqueWordByAlpha[i] 
			fmt.Println( "trovato ", xWordF.uWord2, " ix=", i,   " totRow=", xWordF.uTotRow, " ixwordFre=", xWordF.uIxWordFreq); 
		}
	}
	fmt.Println( "-------------------------------------------------\n" ); 
	return 
}
//---------------------


//------------------------------------------------

func lookForLemmaWord(lemmaCode string) (int, int) {
	
	// find 2 indices of the 2 words nearest to the word to find 
	
	low   := 0
	high  := len(lemma_word_ix) - 1	
	maxIx := high; 
	
	//----
	for low <= high{
		median := (low + high) / 2
		if lemma_word_ix[median].lw_lemmaCod < lemmaCode {  
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
