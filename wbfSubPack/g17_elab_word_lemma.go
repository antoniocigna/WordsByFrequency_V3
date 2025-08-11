package wbfSubPack

	import (
		"fmt"	
		"sort"
	)

//------------------------------

func TOGLIg17_build_lemma_word_ix() {
	fmt.Println("func ", green("build_lemma_word_ix "),   )
	//	build a slice with all lemma with all words 
	//  lemma_word_ix  loaded in addWordLemmaTranLevelParadigma

	var LW lemmaWordStruct
	/***						
			type lemmaStruct struct {
				leLemma    string
				leNumWords int 
				leFromIxLW  int 
				leToIxLW    int  
				leTran     string  
			}
			//---------------
			type lemmaWordStruct struct {
				lw_lemmaSeq string 	
				lw_lemma2   string 	
				lw_word     string 
				lw_ixLemma    int
				lw_ixWordUnFr int
			}
			//------------
	***/	
	//--------------------
	
	// order by lemma and word 	
	sort.Slice(lemma_word_ix, func(i, j int) bool {
			if  lemma_word_ix[i].lw_lemmaSeq != lemma_word_ix[j].lw_lemmaSeq {
				return lemma_word_ix[i].lw_lemmaSeq < lemma_word_ix[j].lw_lemmaSeq 	
			} else {
				if  lemma_word_ix[i].lw_lemma2 != lemma_word_ix[j].lw_lemma2 {
					return lemma_word_ix[i].lw_lemma2 < lemma_word_ix[j].lw_lemma2	
				} else {
					return lemma_word_ix[i].lw_word < lemma_word_ix[j].lw_word 
				}				
			}
		} )		
	fromIxWL:=-1
	toIxWL  :=-1	
	preIxLem:=-1 
	
	fmt.Println("func ", green("build_lemma_word_ix "),  " len( lemma_word_ix) =",  len( lemma_word_ix)    )
	
	for z2:=0; z2 < len( lemma_word_ix); z2++ { 
		LW = lemma_word_ix[z2];
		if LW.lw_ixLemma != preIxLem  {	
			if preIxLem >=0 {
				lemmaSlice[preIxLem].leFromIxLW = fromIxWL
				lemmaSlice[preIxLem].leToIxLW   = toIxWL
			}
			preIxLem = LW.lw_ixLemma
			fromIxWL = z2
		} 
		toIxWL = z2	
	}
	if preIxLem >=0 {
				lemmaSlice[preIxLem].leFromIxLW = fromIxWL
				lemmaSlice[preIxLem].leToIxLW   = toIxWL
			}
	//---------------
	/**
	for z2:=0; z2 < len( lemma_word_ix); z2++ { 
		if z2 < 20 { fmt.Println(   "lemma_word_ix[", z2, "] = ", lemma_word_ix[z2]  ) } else { break }
		
	}
	**/
	//---------------
	/**
	for z2:=0; z2 < len( lemmaSlice); z2++ { 
		LE := lemmaSlice[z2];
		if LE.leLemma == "werkzeug" {
			fmt.Println("lemma = ", LE )
			for z3:= LE.leFromIxLW; z3 <= LE.leToIxLW; z3++ {
				LW = lemma_word_ix[z3]
				fmt.Println(" \t word = ", LW) 
			} 
			break;
		}
	}
	***/
	
} // end of build_lemma_word_ix  
 

//-----------------------------------
func TOGLIbuild_listOfLemmaForAWord(uWord string, ixLemmaPairFoundList []int) ([]int, []string) {
	
	//sw1:= (strings.Index(uWord, "stellen") >= 0)
		
		nele := len(ixLemmaPairFoundList)
		nele += 20
		
		lis_ixLemma1:=make( [] int,    0, nele )				
		lis_ixLemma2:=make( [] int,    0, nele )		
		lis_ixLemma3:=make( [] int,    0, nele )
		lis_origLemma1:= make( []string,    0, nele )	
		lis_origLemma2:= make( []string,    0, nele )
		lis_origLemma3:= make( []string,    0, nele )	
		
		ixLemma:=-1
		numLerr:=0; maxNumLerr:=100; 
		//------------------
		for  _, ixLp := range ixLemmaPairFoundList { 
			
			if numLerr > maxNumLerr { break}
			if ixLp < 0 {
				//lemma3 := "?" + uWord	
				ixLemma = -1; // addUnknowToLemma(lemma3) 
			} else {
				newWL := wordLemmaPair[ixLp]
				if newWL.lWord2 != uWord { // error 
					continue; 
				}
				if newWL.lLemma == LEMMA_MISSING { lemmaNotFoundList = append( lemmaNotFoundList, uWord ) } 
				ixLemma = newWL.lIxLemma 
			}
			if ixLemma < 0 { continue } // error
			
			leSL:= lemmaSlice[ixLemma]
			
			lis_ixLemma1   = append(  lis_ixLemma1, ixLemma      )  
			lis_origLemma1 = append(lis_origLemma1, leSL.leLemma )  			
			
		}	
		for n2, ix2:= range lis_ixLemma2{
			if contains(lis_ixLemma1, ix2) { continue } 
			lis_origLemma1 = append(lis_origLemma1, lis_origLemma2[n2]  )  
			//if sw1 {  fmt.Println(" lemma2 ", lemmaSlice[ix2], " ixLemma=",ix2, " orig=",  lis_origLemma2[n2] )  }
		}
		//------------------------------------
		for n3, ix3:= range lis_ixLemma3{
			if contains(lis_ixLemma1, ix3) { continue } 
			lis_ixLemma1 = append( lis_ixLemma1, ix3) 
			lis_origLemma1 = append(lis_origLemma1, lis_origLemma3[n3]  )   
			//if sw1 {  fmt.Println(" lemma3 ", lemmaSlice[ix3], " ixLemma=",ix3, " orig=",  lis_origLemma3[n3] )  }	
		}
		
		return lis_ixLemma1,lis_origLemma1 
		
} // end of build_listOfLemmaForAWord

//-----------------------------------------
func add_ixWord_to_WordSliceAlpha() {
} // end of add_ixWord_to_WordSliceAlpha	
//-------------------------------------------
func addUnknowToLemma( lemma1 string) int {
	var leV lemmaStruct 
	leV.leLemma    = lemma1
	leV.leNumWords = 0; 
	leV.leTran     = ""
	//leV.ls_lemma_ix_stellen = -1 	
	//leV.ls_lemma_einStellenList = nil

	lemmaSlice = append(lemmaSlice, leV ); 

	return len(lemmaSlice) -1 

} // end of addUnknowToLemma 

//-----------------------------------------

/****
func compound_lemma( leS lemmaStruct)  string { 

	  str1:= ""
	  
	  if leS.ls_lemma_ix_stellen >=0 {
			str1+= "<br>" + leS.ls_pref_ein +  "(" +   leS.ls_pref_tran+")" + " + " +  leS.ls_lemma_stellen 
	  }
	  
	  if len(leS.ls_lemma_einStellenList) > 0 {
			for _, ixle2:= range leS.ls_lemma_einStellenList {
				lelem:= lemmaSlice[ixle2]
				str1 += "<br>" + lelem.leLemma  + " (" + lelem.leTran + ")"  + 
					" = " + lelem.ls_pref_ein + "(" + lelem.ls_pref_tran +")" + " + " + lelem.ls_lemma_stellen     
			}
	  } 	
	  if  len(str1) < 4 {return str1} 	  
	  return str1[4:]
}						
***/
//----------------------------
func sPrintOneLemma( ix int, leS lemmaStruct)  string { 

	  str1:= fmt.Sprint(" ixLemma=", ix, " lemma=", leS.leLemma ,  leS.leTran) 
	  /**
	  if leS.ls_lemma_ix_stellen >=0 {
		str1 += fmt.Sprint(" composto da ",  leS.ls_lemma_stellen, "(", leS.ls_lemma_ix_stellen,") + ", leS.ls_pref_ein , "(" +   leS.ls_pref_tran+")" )
	  }
	  
	  if len(leS.ls_lemma_einStellenList) > 0 {
			str1 += fmt.Sprint(" einStelleList=", leS.ls_lemma_einStellenList )
			for _, ixle2:= range leS.ls_lemma_einStellenList {
				str1 += "\n\t\t" + lemmaSlice[ixle2].leLemma 
			}
	  } 
	  **/
	  return str1	
		/***
		return fmt.Sprint(" ixLemma=", ix, " lemma=", leS.leLemma , 
		" leTran=", leS.leTran, 
		" ixStell=", leS.ls_lemma_ix_stellen,
		" lemmaStell=", leS.ls_lemma_stellen,
		" pref=", leS.ls_pref_ein,
		" prefTran=", leS.ls_pref_tran,
		" ixLemmaList=", leS.ls_lemma_einStellenList )
		***/
}						

//-----------------------------------------------
func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
//------------------
