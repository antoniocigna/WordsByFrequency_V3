package wbfSubPack

import (  
	"fmt"
    "strings"
)
//--------------------------------------------------------------------------

func buildListLemmaSlice( wordLemmaPairTMP []wordLemmaPairStruct) {
	
	preLemS12:=  ""
	lemS12:= ""		
	preLemma:=  ""
	numLemmaAdded:=0
	numLemmaOrig:=0 
	doppi:=0
	z:=-1 // minus 1
	fromIx:=0
	toIx:=0
	numW:=0
	//-----------------	
	for _, lemX := range wordLemmaPairTMP {
		lemS12 = lemX.lLemma + " " + lemX.lWord2 
		if preLemS12 == lemS12 {
			doppi++
			continue
		}	
		preLemS12 = lemS12
		wordLemmaPair = append( wordLemmaPair, lemX)
		z++
		if preLemma != lemX.lLemma { 	
			if numW > 0 {
				//scrive lemma precedente 	
				numLemmaOrig, numLemmaAdded	= appendOneLemma( preLemma, fromIx, toIx, numLemmaOrig, numLemmaAdded)
			}			
			numW=0
			fromIx=z;  
		} 
		numW++
		toIx=z	
		preLemma = lemX.lLemma		

	} // end for wordLemmaPairTMP
	//------------------------------
	
	if numW > 0 {
		//scrive lemma precedente 				
		numLemmaOrig, numLemmaAdded	= appendOneLemma( preLemma, fromIx, toIx, numLemmaOrig, numLemmaAdded)	
	}
	
	//------------------------------------
	if doppi > 0 {
		fmt.Println(" scartate ", doppi, " entrate doppie in lemma - word ") 
	}
	numLemmaDict = len(wordLemmaPair)
	
	fmt.Println( "caricate " , numLemmaDict ,  " coppie word-lemma", "\n")
	
	//----------------------------------
	
	
	//----------------------------
	/**
	fmt.Println("\n---------------------------------")
	
	for nn, lem1 := range lemmaSlice {
		if nn > 40 { break }
		fmt.Println(" lista lemmaSlice = ", lem1 ) 
	} 
	**/
	
	fmt.Println("-----------")
	fmt.Println("num lemma Orig=", numLemmaOrig, " num Lemma added=", numLemmaAdded, " num Tot=", len( lemmaSlice) )
	fmt.Println("---------------------------------\n")
	
	//----------------
	lemmaSliceUpdateSubLemma()
	
	/**
	fmt.Println("\n------  ix update ---------------------------")
	for nn, lem1 := range lemmaSlice {
		if nn > 40 { break }
		fmt.Println(" lista lemmaSlice = ", lem1 ) 
	} 
	***/
	
	/***
				DATI VERI:  num lemma Orig= 19551  num Lemma added= 1593  num Tot= 21144
	***/
	//-----------------------------	
	
	//------------------------
	// update  einStellenList
	for z2, wD:= range wordLemmaPair {

		//if z2 < 0 {	fmt.Println( "wordLemmaPair ", wD) }
		
		ix2 :=wD.lIxLemma
		if ix2 < 0 { continue}
		wD.lIx_einStellenList = lemmaSlice[ ix2 ].ls_lemma_einStellenList 
		wordLemmaPair[z2] = wD 
		//if strings.Index(wD.lLemma,"stellen") >=0  { fmt.Println( "??anto2 buildLemma ",  wD) }
	}  
	//-----------
	
	//-----------------------------	
	fmt.Println("") 
	/**
	for _, wD:= range wordLemmaPair {
		fmt.Println( "wordLemmaPair ", wD) 
	}  
	**/
	
	//-----------------------------
	
	loadInverseLemmaSlice()	
	
	//------------------------
	
} // end of buildListLemmaSlice	

//--------------------------------

func appendOneLemma( xLemma string, fromIx int, toIx int, numLemmaOrig int, numLemmaAdded int	) (int, int) {

	var leV lemmaStruct; 
	iixLem:=0	
	leV.leLemma    = xLemma
	leV.leNumWords = 0 
	leV.leTran     = ""
	
	leV.ls_lemma_ix_stellen = -1
	leV.ls_lemma_stellen    = ""		
	leV.ls_pref_ein         = "" 
	leV.ls_pref_tran        = ""	
	leV.ls_lemma_einStellenList = nil; 
	
	// eg. einstellen  = ein + stellen 	
	
	/**
	works even with multibyte characters: eg. if prefix="日本"  then   lemma="日本語語語"  is broken down into:  "日本" + "語語語"   
	**/
	for _, p:= range separPrefList {
		//fmt.Println("separable prefix ", p.sPrefix , " \t ", p.sPrefTran)
	
		mightLemma, swYes:= strings.CutPrefix(xLemma, p.sPrefix) 
		if len(xLemma) < (len(p.sPrefix)+2) {swYes = false}  
		if swYes  {			
			leV.ls_lemma_stellen    = mightLemma
			leV.ls_pref_ein         = p.sPrefix 
			leV.ls_pref_tran        = p.sPrefTran 
			numLemmaAdded++ 
			break
		}
	}
	
	numLemmaOrig++
	lemmaSlice = append(lemmaSlice, leV ) 
	iixLem = len(lemmaSlice) -1 
	for h:=fromIx; h<= toIx; h++ {
		wordLemmaPair[h].lIxLemma = iixLem    
		wordLemmaPair[h].lIx_einStellenList = nil 
	}	
	return numLemmaOrig, numLemmaAdded			
} // end of appendOneLemma 

//---------------------
/** 
 lista lemmaSlice =  {herstellen -1 stellen her avanti 0 0 0    }
 lista lemmaSlice =  {meile -1    0 0 0    }		
**/
func lemmaSliceUpdateSubLemma() {
	
	for z1, lem1:= range lemmaSlice {	
		if lem1.ls_lemma_stellen == "" { continue }
		ix1, ix2 := lookForLemma( lem1.ls_lemma_stellen )
		ixFound:=-1
		for z3:=ix1; z3 <=ix2; z3++ {
			if lemmaSlice[z3].leLemma == lem1.ls_lemma_stellen {
				ixFound = z3
				break
			}
		}	
		if ixFound < 0 {
			lem1.ls_lemma_stellen = ""
			lem1.ls_pref_ein = ""
			lem1.ls_pref_tran= ""
		} else {
			lem1.ls_lemma_ix_stellen = ixFound
			lemmaSlice[ixFound].ls_lemma_einStellenList = append(lemmaSlice[ixFound].ls_lemma_einStellenList, z1)		
		}
		lemmaSlice[z1] = lem1 
	}	

} // end of  lemmaSliceUpdateSubLemma

//-------------------------------------------------------------
