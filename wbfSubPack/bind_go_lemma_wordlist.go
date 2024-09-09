package wbfSubPack

	import (		
		"fmt"
		"strconv"
		/**
		"strings"		
		
		"regexp"
		"sort"
		**/
	)
//--------------------------------------------------------

func bind_go_passToJs_lemmaWordList(lemmaToFind0 string, inpMaxWordLemma int, js_function string)   {
		/**
			lista tutte le parole col lemma indicato		
		**/
		//lemmaCod:= newCode( lemmaToFind0 )
		//var onlyThisLevel string = "any" ; // "A0"  // questo deve arrivare da parametro  
		//onlyIfExtr := false 
		
		outS1 := "" 
		
		fromIx, toIx:= lookForLemma(lemmaToFind0) // find the index of lemma in lemmaSlice 
		ixLe :=-1
		for k:= fromIx; k <= toIx; k++ {
			if lemmaSlice[k].leLemma == lemmaToFind0 { ixLe = k;  break }
		}
		if ixLe < 0 {
			go_exec_js_function( js_function,   "NONE," + lemmaToFind0 )
			return
		}
		
		//lem1_prefix      := ""
		//lem1_prefix_tran := ""
		myLem1_name      := lemmaToFind0
		myLem2_name      := ""
		myLem2_ix        := -1
		var myLem2 lemmaStruct
		myLem1 := lemmaSlice[ixLe]
		
		/*					
			type lemmaStruct struct {
				leLemma    string        // ls_lemma_einstellen einstellen				
				leNumWords int 
				leFromIxLW  int 
				leToIxLW    int  
				leTran      string 
				leLevel     string  
				lePara      string  
				leExample   string  
				ls_lemma_ix_stellen  int	
				ls_lemma_stellen     string
				ls_pref_ein          string
				ls_pref_tran         string 
				ls_lemma_einStellenList []int 
			}		
		*/
		
		
		//fmt.Println("bind_go_passToJs_lemmaWordList target lemma=", lemmaToFind0)
		//fmt.Println("\t", "               myLem1=", myLem1)   
		
		/**
		bind_go_passToJs_lemmaWordList target lemma= einstellen  trovato myLem= {einstellen 2  5678  5679 impostare    14893 stellen ein dentro, verso l'interno []}
					servono le parole con lemma einstellen, 
						ma anche quelle con lemma stellen:   es. stelle appartiene a stellen e insieme ad ein al lemma einstellen 
						ma poi le presento come se fossero tutte dei einstellen  			
		----
		bind_go_passToJs_lemmaWordList target lemma= stellen     trovato myLem= {stellen   21 20259 20279 fornire         -1                                     [718 1170 1362 4097 7175 7296 17303]}		
		
					servono solo le parole con lemma stellen, es. stellt, gestellt 
		**/
		
		myLem2_ix = myLem1.ls_lemma_ix_stellen
		if myLem2_ix >= 0 {
			myLem2_name = myLem1.ls_lemma_stellen
			//lem1_prefix = myLem1.ls_pref_ein
			//lem1_prefix_tran = myLem1.ls_pref_tran
			myLem2 = lemmaSlice[myLem2_ix]
			//fmt.Println("\t", " myLem2_name=", myLem2_name , "myLem2=", myLem2,    "  lem1_prefix =", lem1_prefix, " tran=", lem1_prefix_tran) 
		}
		//---------------------------------
		//  .leFromIxLW / .leToIxLW    are indicies of element in "lemma_word_ix"
		listaIxLemmaWordIx:= make([]int, 0, 200 ) 
		for k := myLem1.leFromIxLW; k <= myLem1.leToIxLW; k++ {
			listaIxLemmaWordIx = append(listaIxLemmaWordIx, k) 
		} 
		if myLem2_ix >= 0 {
			for k := myLem2.leFromIxLW; k <= myLem2.leToIxLW; k++ {				
				if containIntList(listaIxLemmaWordIx, k) == false {  
					listaIxLemmaWordIx = append(listaIxLemmaWordIx, k) 
				}
			} 
		}
		//fmt.Println(" listaIxLemmaWordIx=", listaIxLemmaWordIx)
		numO:= 0
		//---------------------------
		
		for _,k:= range listaIxLemmaWordIx {
			lw1:= lemma_word_ix[k]
			
			
			if ((lw1.lw_origLemma != myLem1_name) && (lw1.lw_origLemma != myLem2_name)) {
				continue
			} 
			
			ix := lemma_word_ix[k].lw_ixWordUnFr 		
			
				
			rowW := lemma_word_to_row(myLem1_name, myLem2_name, lw1.lw_lemma2, uniqueWordByFreq[ix] )  
				
				numO++
				if numO > inpMaxWordLemma { break } 
				outS1 += rowW 
			  	
			//fmt.Println("    rowW=", rowW)	
		} 
		
		if len(outS1)< 3 {
			outS1 = "NONE," + lemmaToFind0; 			
			//fmt.Println(" bind_go_passToJs_lemmaWordList() ", outS1) 
		} 	
		
		go_exec_js_function( js_function, outS1 ); 		
		
	
				
} // end of bind_go_passToJs_lemmaWordList

//-----------------------------------------------------------
func containIntList( list1 []int, s int) bool{
	for _,tar:= range list1 {
		if tar == s { return true}
	} 
	return false 
}

//------------------------------------------------------

func lemma_word_to_row( myLem1_name string, myLem2_name string, lw_lemma2 string, xWordF2 wordIxStruct) string  {
	
	
		ix2 := -1
		for x1, oneLemma := range xWordF2.uLemmaL {
			if oneLemma == lw_lemma2 {
				ix2=x1; 
				break
			}  	
		}
			
		return xWordF2.uWordSeq + ";." + xWordF2.uWord2 + ";." + 
			"ix" + ";." + 
			strconv.Itoa(xWordF2.uIxUnW) + ";." + strconv.Itoa(xWordF2.uTotRow)  + ";." + 
			xWordF2.uLemmaL[ix2]              + ";." + 
			tranFromIxLemma( xWordF2, ix2)    + ";." +  
			xWordF2.uLevel[ix2]               + ";." +  
			xWordF2.uPara[ix2]                + ";." +  
			xWordF2.uExample[ix2]             + ";." +  
			strconv.Itoa(xWordF2.uTotExtrRow) + ";." +  					
			strconv.Itoa(xWordF2.uKnow_yes_ctr)  + ";." + strconv.Itoa(xWordF2.uKnow_no_ctr)  + ";." + 				
			"ixLemma" + ";." + fmt.Sprint( xWordF2.uIxLemmaL[ix2] ) + ";." + 	
			endOfLine 				
			 
} // end of lemma_word_to_row 

//-----------------------------------------------------------

func build_one_lemma_row_word( myLem1 lemmaStruct) string {	

		//lem1_prefix      := ""
		//lem1_prefix_tran := ""

		myLem1_name      := myLem1.leLemma

		myLem2_name      := ""
		myLem2_ix        := -1
		
		//var myLem2 lemmaStruct	
		
		//fmt.Println("build_one_lemma_row_word   target lemma=", myLem1_name)
		
		
		myLem2_ix = myLem1.ls_lemma_ix_stellen
		if myLem2_ix >= 0 {
			myLem2_name = myLem1.ls_lemma_stellen
		}
		//---------------------------------
		//  .leFromIxLW / .leToIxLW    are indicies of element in "lemma_word_ix"
		
		listaIxLemmaWordIx:= make([]int, 0, 200 ) 
		
		for k := myLem1.leFromIxLW; k <= myLem1.leToIxLW; k++ {
			listaIxLemmaWordIx = append(listaIxLemmaWordIx, k) 
			break  //  only once 
		} 
		
		//fmt.Println(" listaIxLemmaWordIx=", listaIxLemmaWordIx)
		//numO:= 0
		//---------------------------
		
		for _,k:= range listaIxLemmaWordIx {
			lw1:= lemma_word_ix[k]
		
			
			ix := lemma_word_ix[k].lw_ixWordUnFr 		
			
			//fmt.Println("           word=", uniqueWordByFreq[ix].uWord2 )  
				
			rowW := lemma_word_to_row(myLem1_name, myLem2_name, lw1.lw_lemma2, uniqueWordByFreq[ix] )  
			
			//fmt.Println( "return rowW=", rowW)
			
			return  rowW // only once 	
		} 
	
	//fmt.Println( "return spazio")
	
	return ""
	
} // end of build_one_lemma_row_word		

//-------------------------------------------

