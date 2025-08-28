package wbfSubPack

import (  
	//"fmt"
    //"strings"
    "sort"
	//"slices"
)
//----------------------------------------------------
//------------
var newWordLemmaPair    [] wordLemmaPairStruct // all word-lemma pair 
 

var lemmaSlice       [] lemmaStruct         // lemma , translation 

//var wordLemmaPair    [] wordLemmaPairStruct // all word-lemma pair  
//----------------------------------------

func addToCurrentLemmaPair(wordLemmaPair []wordLemmaPairStruct ) {
	sort.Strings(lemmaNotFoundList)
	preL := ""
	for _, oneLemma:= range lemmaNotFoundList { 
		if oneLemma == preL { continue } 
		preL = oneLemma  
		var wordLemma1 wordLemmaPairStruct 
		wordLemma1.lWord2   = stdCode( oneLemma ) 		
		wordLemma1.lLemma   = wordLemma1.lWord2 
		if len(wordLemma1.lLemma) < 1 { continue;  } 
		if ((wordLemma1.lLemma == "-") || (wordLemma1.lLemma[0:1] < "A")) { continue;  }   // ignore number  
		//wordLemma1.lWordSeq = seqCode( wordLemma1.lWord2)
		wordLemma1.lIxLemma = -1	
		wordLemmaPair = append(wordLemmaPair, wordLemma1 ) 
	}
}
//-----------------------------
/**
func check_wordLemma_sameCode() {
	fmt.Println( green("check_wordLemma_sameCode") , "()"  )
	// check same words  written in diffent way (eg. caesar   and  "cäsar")
	pre_wordCod := ""
	pre_word2   := ""	
	//pre_lemma   := ""
	pre_z := -1
	
	for z, wordPair := range wordLemmaPair {	
		//if ((  wordPair.lWord2 == "abgehauen") || (wordPair.lLemma == "abhauen") ) { fmt.Println(green("check_wordLemma_sameCode abhauen "), "z=", z,  " wordPair=" , wordPair) }
	
		if (wordPair.lWordSeq != pre_wordCod) {
			//pre_wordCod = wordPair.lWordSeq 
			pre_word2   = wordPair.lWord2 
			//pre_lemma   = wordPair.lLemma 
			pre_z = z
			continue
		}
		if (wordLemmaPair[z].lWord2 == pre_word2) {
			continue
		}
		//--------
		fmt.Println( green("check_wordLemma_sameCode") )
		for x:= pre_z; x<= z; x++ {
			fmt.Println("\t", " wordLemmaPair[",x,"] = ", wordLemmaPair[x] )   
		} 
		
	}	
	
} // end of check_wordLemma_sameCode
***/
//--------------------------------------

