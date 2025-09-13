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

func g34_lookForInverse(invTarg string, inverseSlice []inverseStruct, sw_oneOnly bool, lastIx int) (int, int) {
	
	// find 2 indices of the 2 words nearest to the word to find 
	lenTarg := len(invTarg)		
	low   := 0
	high  := len(inverseSlice) - 1	
	//maxIx := high; 
	
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
	//fmt.Println("g34_lookForInverse ", invTarg,  " sw_oneOnly=",   sw_oneOnly , " low=", low, " high=", high, "  lenTarg=", lenTarg)
	
	fromIx:=-1; toIx:= -2 

	len2 :=0
	//--------------------------------- vai indietro partendo dall'indice  low  
	for z:=low; z >=0; z-- {
		len2 =  len(inverseSlice[z].inInverse) 
		if len2 < lenTarg {continue }
		if sw_oneOnly == false { len2 = lenTarg}
		//fmt.Println("     loop1 confronta =>" + inverseSlice[z].inInverse + "<== len=", len2, " con ==>" + invTarg + "<==")
		if inverseSlice[z].inInverse[0:lenTarg] != invTarg { break  } 
		fromIx = z;		
		//fmt.Println("      1 trovato eguale  fromIx=", fromIx)
	}
	//------------------	
	for z:=fromIx; z < len(inverseSlice); z++ {   //  vai avanti 
		if z < 0 {continue}
		len2 =  len(inverseSlice[z].inInverse) 
		if len2 < lenTarg {continue }
		if sw_oneOnly == false { len2 = lenTarg}
		//fmt.Println("     loop2 confronta =>" + inverseSlice[z].inInverse + "<== len=", len2, " con ==>" + invTarg + "<==")
		if inverseSlice[z].inInverse[0:lenTarg] != invTarg { break  } 
		toIx = z;
		//fmt.Println("      2 trovato eguale  fromIx=", fromIx, " toIx=", toIx)
	}	
	return fromIx, toIx	

} // end of g34_lookForInverse

var inverseLemmaSlice = make([] inverseStruct, 0, 0)
var directLemmaSlice  = make([] inverseStruct, 0, 0)
//-----------------------------------------------
func g34_load_direct_and_inverse_lemma()  {	
	
	ItempLemma := make([] inverseStruct, 0, len( lemmaSlice) )	
	DtempLemma := make([] inverseStruct, 0, len( lemmaSlice) )	
	
	var oneInv inverseStruct
	var oneDir inverseStruct
	/*
	type lemmaStruct struct {
		leLemma    string    
		leNumWords int 
		leFromIxLW  int 
		leToIxLW    int  
		leUnWord_al_IxList []int    // indice delle parole unique che puntano a questo lemma        
		leTran      string 
		//leLevel     string  
		lePara      string  
		leExample   string  
		leNumPara   int	
	} 
	*/
	for lemIndex, dirLem := range lemmaSlice {
		//tutte i lemma anche se non nel testo;  if dirLem.leNumWords == 0 { continue }		// nessuna word del testo ha questo lemma 	
		oneInv.inInverse = reverseString( dirLem.leLemma ) 
		oneInv.inIx      = lemIndex		
		ItempLemma = append( ItempLemma, oneInv )	
		
		oneDir.inInverse = dirLem.leLemma  
		oneDir.inIx      = lemIndex		
		DtempLemma = append( DtempLemma, oneDir )	
	}  
	numInv := len(ItempLemma) 
	
	//fmt.Println( "   g34_load_direct_and_inverse_lemma",   " len(tempLemma)=", len(tempLemma) )	
	
	directLemmaSlice  = make([] inverseStruct, numInv)
	inverseLemmaSlice = make([] inverseStruct, numInv)
	copy(directLemmaSlice , DtempLemma)
	copy(inverseLemmaSlice, ItempLemma)
	
	sort.Slice(inverseLemmaSlice, func(i, j int) bool {
				return inverseLemmaSlice[i].inInverse < inverseLemmaSlice[j].inInverse
			}   )		
	fmt.Println( "func ", green("loadInverseLemmaSlice") , " caricati ", len(inverseLemmaSlice) , " lemma inversi", 
		len(directLemmaSlice), " lemma diretti"	)   
	
	
	//provaInverseLemma("hen", false, 100)
	//provaDirectLemma("welt",true,  100)
	//provaDirectLemma("welt",false, 100)
	
	
}  // end of loadInverseLemmaSlice
//-------------------
func g34_getListInverseLemmaIndex(dirLemmaTarg string, sw_oneLemma bool, maxNum int) []int { 
			
	invLemmaTarg:= reverseString( dirLemmaTarg ) 
	
	fromIx, toIx:= g34_lookForInverse(invLemmaTarg, inverseLemmaSlice, sw_oneLemma,0)
	
	//fmt.Println("g34_getListInverseLemmaIndex  len(inverseLemmaSlice)=", len(inverseLemmaSlice), " fromIx=", fromIx, " toIx=", toIx)
	
	//if toIx < fromIx { fromIx = toIx}
	
	//fmt.Println("get inverse ", invLemmaTarg,  " fromIx=", fromIx, " toIx=", toIx) 	
	
	listInverseLemmaIndex:= make([]int,0,200)    
	
	if toIx < 0 { return listInverseLemmaIndex }
	
	//lenTarg:= len(dirLemmaTarg)
	//lenCk  := 0
	num:=0
	
	if toIx >= len(inverseLemmaSlice) {toIx = len(inverseLemmaSlice)-1 }
	
	for j:=fromIx; j <=toIx; j++ {		
		invLem := inverseLemmaSlice[j]
		/**
		lenCk = len(invLem.inInverse)
		if sw_oneLemma == false {if lenCk > lenTarg { lenCk = lenTarg } }
		if invLem.inInverse[0:lenCk] < invLemmaTarg { continue}
		if invLem.inInverse[0:lenCk] > invLemmaTarg { break }		
		**/
		num++
		if num > maxNum { break }
		listInverseLemmaIndex = append(listInverseLemmaIndex,  invLem.inIx )  //     lem:= sliceLemma[ invLem.ixLemma ]		
	} 
	return listInverseLemmaIndex 
	
} // end of getListInverseLemma

//-----------------------------------------------------

func g34_getListDirectLemmaIndex(dirLemmaTarg string, sw_oneLemma bool, maxNum int) []int { 
	//fmt.Println("func ", green("g34_getListDirectLemmaIndex "),  dirLemmaTarg, " sw_oneLemma=",  sw_oneLemma )	
	invLemmaTarg:= dirLemmaTarg
	
	fromIx, toIx:= g34_lookForInverse(invLemmaTarg, directLemmaSlice,sw_oneLemma, 0)
	
	//fmt.Println("   esegue g34_lookForInverse(", invLemmaTarg, " in directLemmaSlice ( len=", len(directLemmaSlice), ") ==> ", " fromIx=", fromIx, ", toIx=", toIx)
	
	listDirectLemmaIndex:= make([]int,0,200)     
	if toIx < 0 {  return listDirectLemmaIndex }
	//lenTarg:= len(dirLemmaTarg)
	if toIx >= len(directLemmaSlice) {toIx = len(directLemmaSlice)-1}
	num:=0
	for j:=fromIx; j <= toIx; j++ {		
		invLem := directLemmaSlice[j]		
		/**
		lenCk = len(invLem.inInverse)
		if sw_oneLemma == false {if lenCk > lenTarg { lenCk = lenTarg } }
		if invLem.inInverse[0:lenCk] < invLemmaTarg { continue}
		if invLem.inInverse[0:lenCk] > invLemmaTarg { break }	
		**/		
		num++
		if num > maxNum { break }
		listDirectLemmaIndex = append(listDirectLemmaIndex,  invLem.inIx )  //     lem:= sliceLemma[ invLem.ixLemma ]		
	}  
	//fmt.Println("    listDirectLemmaIndex =", listDirectLemmaIndex ) 
	
	return listDirectLemmaIndex 
	
} // end of getListInverseLemma
//------------------------------------------
func provaInverseLemma( finalLemma string, sw_oneLemma bool, maxNum int) {

	fmt.Println( "\n\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n     prova finalLemma = ", finalLemma ) 	
	
	listInverseLemmaIndex := g34_getListInverseLemmaIndex( finalLemma, sw_oneLemma, maxNum) 
	
	//fmt.Println( "     prova indici = ",listInverseLemmaIndex ) 
	
	for _,ixL:= range listInverseLemmaIndex {  
		fmt.Println( "trovato lemma con finale=", finalLemma , " ==> ", lemmaSlice[ixL].leLemma )
	}
	//---	
}
//------------
func provaDirectLemma( finalLemma string, sw_oneLemma bool, maxNum int) {

	fmt.Println( "\n\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n   ", red("prova direct Lemma"), " = ", finalLemma , "  sw_oneLemma =",  sw_oneLemma ) 
	
	listDirectLemmaIndex := g34_getListDirectLemmaIndex(  finalLemma, sw_oneLemma, maxNum) 	 
	
	for _,ixL:= range listDirectLemmaIndex {  
		fmt.Println( "trovato lemma con prefisso=", finalLemma , " ==> ", lemmaSlice[ixL].leLemma )
	}
} // end of testDirectLemma

//============================================================================================

var inverseWordSlice = make([] inverseStruct, 0, 0)
//-----------------------------------------------
func loadInverseWordSlice()  {
	
	inverseWordSlice = make([] inverseStruct, 0, len(uniqueWordByAlpha) )
	var oneInv inverseStruct
	                                                
	for ix1, oneWord := range uniqueWordByAlpha {	
		oneInv.inInverse = reverseString(  oneWord.uWord2 )
		oneInv.inIx      = ix1		
		inverseWordSlice = append( inverseWordSlice, oneInv )
	}  
		
	sort.Slice(inverseWordSlice, func(i, j int) bool {
				return inverseWordSlice[i].inInverse < inverseWordSlice[j].inInverse
			}   )		
	fmt.Println( green("loadInverseWordSlice") , " caricati ", len(inverseWordSlice) , " word inversi")   
	
	/**
	for j:=0; j < len(inverseWordSlice); j++ {		
		fmt.Println( green("inverse Word "), inverseWordSlice[j] )		
	} 
	**/
	
}  // end of loadInverseWordSlice
//-------------------------------------------

func g34_getListInverseWordIndex(dirWordTarg string, sw_oneOnly bool, maxNum int) []int { 

	//fmt.Println("g34_getListInverseWordIndex(dirWordTarg=",dirWordTarg	)
	
	targWordCoded := dirWordTarg 	
	
	invWordTarg:= reverseString(targWordCoded) 
		
	fromIx, toIx:= g34_lookForInverse(invWordTarg, inverseWordSlice, sw_oneOnly,0)
		
	//fmt.Println("    g34_getListInverseWordIndex  invWordTarg=", invWordTarg, " len(inverseWordSlice)=", len(inverseWordSlice), " fromIx=", fromIx, " toIx=", toIx)
	
	listInverseWordIndex:= make([]int,0,200)     
	if toIx < 0 { return listInverseWordIndex }
	
	//lenTarg:= len(targWordCoded)
	//lenCk  := 0
	num:=0
	if toIx >=  len(inverseWordSlice) { toIx = len(inverseWordSlice) -1} 
	//---------------------
	for j:=fromIx; j <=toIx; j++ {		
		invLem := inverseWordSlice[j]
		/**
		lenCk = len(invLem.inInverse)
		//fmt.Println("   j=", j, " invLem=", invLem, " lenCk=",  lenCk, " lenTarg=", lenTarg, " invWordTarg=", invWordTarg)
		if lenCk > lenTarg { lenCk = lenTarg }
		if invLem.inInverse[0:lenCk] < invWordTarg { continue}
		if invLem.inInverse[0:lenCk] > invWordTarg { break }
		**/
		num++
		if num > maxNum { break }
		listInverseWordIndex = append(listInverseWordIndex, invLem.inIx)
	}   
	
	return listInverseWordIndex 
	
} // end of getListInverseWord

//------------------------------------------

func provaInverseWord( finalWord string, maxNum int) {

	fmt.Println( "\n\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n     prova finalWord = ", finalWord ) 	
	sw_oneOnly := false;
	listInverseWordIndex := g34_getListInverseWordIndex( finalWord, sw_oneOnly, maxNum) 
	
	//fmt.Println( "     prova indici = ",listInverseWordIndex ) 
	
	for _,ixL:= range listInverseWordIndex {  
		fmt.Println( "trovato word con finale=", finalWord , " ==> ", uniqueWordByAlpha[ixL].uWord2, "  -  ", uniqueWordByAlpha[ixL].uWord0)
	}
	//----------------------------
	
} // end of testInverseLemma
//-------------------------------------------------------


//-----------------------------------------------------------