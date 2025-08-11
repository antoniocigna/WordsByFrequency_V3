package wbfSubPack

	import (
		//"fmt"
		//"strings"
		//"strconv"
		//"sort"
		//"regexp"
		//"encoding/hex"
	)

//------------------------------


//---------------------------------
	
func g13_addLemmaTranParadigmaToUniqueWord( wF wordUnAlphaStruct, lis_ixLemma []int)  {
	
		var wP lemmaTranStruct  
			
		//lemma_word_ix = make([]lemmaWordStruct, 0,  len(uniqueWordByAlpha)  )  
		
		//var LW lemmaWordStruct		
		
		nele := len( lis_ixLemma )
			
		lis_lemmaName := make( [] string, 0, nele )		
		lis_tran  := make( [] string, 0, nele )
		//lis_level := make( [] string, 0, nele )
		//lis_para  := make( [] string, 0, nele )
		//lis_exam  := make( [] string, 0, nele )
	
		
		//-----------------------
		for _, ixLemma2:= range lis_ixLemma { 
			lem1Slice := lemmaSlice[ixLemma2]
			//lemStru := lemmaSlice[ixLemma2] 
			lem     := lem1Slice.leLemma 
			lis_lemmaName   = append( lis_lemmaName,  lem )  				
			lemmaSlice[ixLemma2].leNumWords++			
			lemmaSlice[ixLemma2].leUnWord_al_IxList = append( lemmaSlice[ixLemma2].leUnWord_al_IxList, wF.uIxUnW_al)
			//-------------
			for g:=lem1Slice.leFromIxLW; g<= lem1Slice.leToIxLW;g++ {					
				if wordLemmaPair_lemmaWordSeq[g].lWordSeq ==  wF.uWordSeq {
					wordLemmaPair_lemmaWordSeq[g].lIxUnWord_al = wF.uIxUnW_al
					break
				}
			}	
			//-----------			
			ixTra := lookForAllTran( lem,0 ) 
			if ixTra >= 0 { 
				wP = dictLemmaTran[ixTra] 
				lis_tran = append( lis_tran, wP.dL_tran ) 		////cigna1	
				if ixLemma2 >=0 {	lemmaSlice[ixLemma2].leTran = wP.dL_tran } 
			} else {
				lis_tran = append( lis_tran, ""         ) 	
				if ixLemma2 >=0 {	lemmaSlice[ixLemma2].leTran = "" }  		
			}				
		} // end of for , lem 
		//-----------
			
		wF.uIxLemmaL= make( []int,     nele, nele )    
		wF.uLemmaL  = make( []string,  nele, nele )    
		//wF.uTranL   = make( []string,  nele, nele )        
		//wF.uLevel   = make( []string,  nele, nele )   
		//wF.uPara    = make( []string,  nele, nele )   
		//wF.uExample = make( []string,  nele, nele )   
		
		
		copy( wF.uIxLemmaL, lis_ixLemma )		
		copy( wF.uLemmaL  , lis_lemmaName ) 
		//copy( wF.uTranL   , lis_tran  )  
		//copy( wF.uLevel   , lis_level ) 
		//copy( wF.uPara    , lis_para  ) 
		//copy( wF.uExample , lis_exam  ) 
		
		uniqueWordByAlpha = append( uniqueWordByAlpha, wF)
		
	
} // end of addLemmaTranParadigmaToUniqueWord

//---------------------------------------------
func elabWordAlpha_buildWordFreqListDUE() {
}
//-----------------------------------