package wbfSubPack

import (  
	//"fmt"
    //"strings"
	//"sort"
)
//--------------------------------------------------------------------------

func g14_appendOneLemma( xLemma string, fromIx int, toIx int, numLemmaOrig int, numLemmaAdded int	) (int, int) {

	var leV lemmaStruct; 
	iixLem:=0	
	leV.leLemma    = xLemma
	leV.leNumWords = 0 
	leV.leFromIxLW = fromIx 
	leV.leToIxLW   = toIx  
	leV.leTran     = ""
	//leV.leLevel    = ""   
	leV.lePara     = ""   
	leV.leExample  = "" 
	//-----------			
	ixTra := lookForAllTran( xLemma,0 ) 
	if ixTra >= 0 { 
		leV.leTran = dictLemmaTran[ixTra].dL_tran 
	} 		
	
	numLemmaOrig++
	lemmaSlice = append(lemmaSlice, leV ) 
	iixLem = len(lemmaSlice) -1 
	for h:=fromIx; h<= toIx; h++ {
		wordLemmaPair[h].lIxLemma = iixLem    
	}	
	return numLemmaOrig, numLemmaAdded			
} // end of g14_appendOneLemma 

//---------------------