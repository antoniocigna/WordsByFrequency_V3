package wbfSubPack

import (  
	"fmt"
    //"strings"
	"sort"
)
//--------------------------------------------------------------------------

type inverseStruct  struct {
	inInverse string  
	inIx      int 
}
//-----------------------------------------------------------
func reverseString(s string) string {
    runes := []rune(s)
    length := len(runes)

    for i := 0; i < length/2; i++ {
        // Swap characters from the start and end
        runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
    }

    return string(runes)
}

//-----------------------------------------------------------

func lookForInverse(invTarg string, inverseSlice []inverseStruct) (int, int) {
	
	// find 2 indices of the 2 words nearest to the word to find 

	low   := 0
	high  := len(inverseSlice) - 1	
	maxIx := high; 
	
	//----
	for low <= high{
		median := (low + high) / 2
		if inverseSlice[median].inInverse < invTarg {  
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

} // end of lookForInverse

//=============================================================================

var inverseLemmaSlice = make([] inverseStruct, 0, 0)
//-----------------------------------------------
func loadInverseLemmaSlice()  {
	
	inverseLemmaSlice = make([] inverseStruct, 0, len( lemmaSlice) )
	var oneInv inverseStruct
	
	for ix1, dirLem := range lemmaSlice {
		oneInv.inInverse = reverseString( dirLem.leLemma ) 
		oneInv.inIx      = ix1		
		inverseLemmaSlice = append( inverseLemmaSlice, oneInv )
	}  
		
	sort.Slice(inverseLemmaSlice, func(i, j int) bool {
				return inverseLemmaSlice[i].inInverse < inverseLemmaSlice[j].inInverse
			}   )		
	fmt.Println( green("loadInverseLemmaSlice") , " caricati ", len(inverseLemmaSlice) , " lemma inversi")   
	
	
}  // end of loadInverseLemmaSlice
//-------------------
func getListInverseLemmaIndex(dirLemmaTarg string, maxNum int) []int { 
	
	invLemmaTarg:= reverseString( dirLemmaTarg ) 
	fromIx, toIx:= lookForInverse(invLemmaTarg, inverseLemmaSlice)
	if toIx < fromIx { fromIx = toIx}
	
	//fmt.Println("get inverse ", invLemmaTarg,  " fromIx=", fromIx) 	
	
	listInverseLemmaIndex:= make([]int,0,200)     
	lenTarg:= len(dirLemmaTarg)
	lenCk  := 0
	num:=0
	for j:=fromIx; j < len(inverseLemmaSlice); j++ {		
		invLem := inverseLemmaSlice[j]
		lenCk = len(invLem.inInverse)
		if lenCk > lenTarg { lenCk = lenTarg }
		//fmt.Println("get inverse for j "," invLem = ", invLem.inIx, " => ",  lemmaSlice[ invLem.inIx ].leLemma	 ) 
		//fmt.Println("      invLem.inInverse= ", invLem.inInverse , " invLemmaTarg=",invLemmaTarg) 
		if invLem.inInverse[0:lenCk] < invLemmaTarg { continue}
		if invLem.inInverse[0:lenCk] > invLemmaTarg { break }		
	
		//fmt.Println("      OK append ", invLem.inIx  ) 
		num++
		if num > maxNum { break }
		listInverseLemmaIndex = append(listInverseLemmaIndex,  invLem.inIx )  //     lem:= sliceLemma[ invLem.ixLemma ]		
	}   
	
	return listInverseLemmaIndex 
	
} // end of getListInverseLemma
//------------------------------------------
func provaInverseLemma( finalLemma string, maxNum int) {

	fmt.Println( "\n\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n     prova finalLemma = ", finalLemma ) 	
	
	listInverseLemmaIndex := getListInverseLemmaIndex( finalLemma, maxNum) 
	
	//fmt.Println( "     prova indici = ",listInverseLemmaIndex ) 
	
	for _,ixL:= range listInverseLemmaIndex {  
		fmt.Println( "trovato lemma con finale=", finalLemma , " ==> ", lemmaSlice[ixL].leLemma )
	}
	//----------------------------
	
} // end of testInverseLemma

//------------------------------------------
/**
func testInverseLemma() {
	maxNum:= 20
	provaInverseLemma( "end"   , maxNum ) 
	provaInverseLemma( "weise" , maxNum ) 
	//provaInverseLemma( "en"  , maxNum ) 
}	
**/
//----------------------------------

//============================================================================================

var inverseWordSlice = make([] inverseStruct, 0, 0)
//-----------------------------------------------
func loadInverseWordSlice()  {
	
	inverseWordSlice = make([] inverseStruct, 0, len(uniqueWordByAlpha) )
	var oneInv inverseStruct
	                                                
	for ix1, oneWord := range uniqueWordByAlpha {	
		oneInv.inInverse = reverseString(  oneWord.uWordSeq )
		oneInv.inIx      = ix1		
		inverseWordSlice = append( inverseWordSlice, oneInv )
	}  
		
	sort.Slice(inverseWordSlice, func(i, j int) bool {
				return inverseWordSlice[i].inInverse < inverseWordSlice[j].inInverse
			}   )		
	fmt.Println( green("loadInverseWordSlice") , " caricati ", len(inverseWordSlice) , " word inversi")   
	
	
}  // end of loadInverseWordSlice
//-------------------------------------------

func getListInverseWordIndex(dirWordTarg string, maxNum int) []int { 

													
	targWordCoded := seqCode( dirWordTarg) 	
	
	invWordTarg:= reverseString(targWordCoded) 
		
	fromIx, toIx:= lookForInverse(invWordTarg, inverseWordSlice)
	if toIx < fromIx { fromIx = toIx}
	
	//fmt.Println("get inverse ", invWordTarg,  " fromIx=", fromIx) 	
	
	listInverseWordIndex:= make([]int,0,200)     
	lenTarg:= len(targWordCoded)
	lenCk  := 0
	num:=0
	for j:=fromIx; j < len(inverseWordSlice); j++ {		
		invLem := inverseWordSlice[j]
		lenCk = len(invLem.inInverse)
				
		if lenCk > lenTarg { lenCk = lenTarg }
		if invLem.inInverse[0:lenCk] < invWordTarg { continue}
		if invLem.inInverse[0:lenCk] > invWordTarg { break }		
	
		num++
		if num > maxNum { break }
		listInverseWordIndex = append(listInverseWordIndex,  invLem.inIx )
	}   
	
	return listInverseWordIndex 
	
} // end of getListInverseWord

//------------------------------------------

func provaInverseWord( finalWord string, maxNum int) {

	fmt.Println( "\n\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n     prova finalWord = ", finalWord ) 	
	
	listInverseWordIndex := getListInverseWordIndex( finalWord, maxNum) 
	
	//fmt.Println( "     prova indici = ",listInverseWordIndex ) 
	
	for _,ixL:= range listInverseWordIndex {  
		fmt.Println( "trovato word con finale=", finalWord , " ==> ", uniqueWordByAlpha[ixL].uWordSeq, "  -  ", uniqueWordByAlpha[ixL].uWord2)
	}
	//----------------------------
	
} // end of testInverseLemma
//-------------------------------------------------------


//-----------------------------------------------------------