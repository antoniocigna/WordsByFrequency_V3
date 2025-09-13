package wbfSubPack

import (  
	//"fmt"
    //"strings"
    "sort"
	//"slices"
)
//----------------------------------------------------
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
		wordLemma1.lWord0   = oneLemma  		
		wordLemma1.lLemmaOr = oneLemma
		
		if len(wordLemma1.lLemma) < 1 { continue;  } 
		if ((wordLemma1.lLemma == "-") || (wordLemma1.lLemma[0:1] < "A")) { continue;  }   // ignore number  
		wordLemma1.lIxLemma = -1	
		wordLemmaPair = append(wordLemmaPair, wordLemma1 ) 
	}
}
//-----------------------------


