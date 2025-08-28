package wbfSubPack

	import (
		"fmt"	
		//"sort"
	)

//------------------------------

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
