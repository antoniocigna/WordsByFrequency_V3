package wbfSubPack

import (  
	"fmt"
    "strings"
	"sort"
)
//--------------------------------------------------------------------------

func buildListLemmaSlice( wordLemmaPairTMP []wordLemmaPairStruct) {
	
	//preLemS12:=  ""
	//lemS12:= ""		
	
	sw_WW00:= false
	sw_WW11:= false
	//--------------------------------------
	// sort x lemma, L_W, l_Word2               	
	sort.Slice(wordLemmaPairTMP, func(i, j int) bool {
			if (wordLemmaPairTMP[i].lWordSeq != wordLemmaPairTMP[j].lWordSeq) {
				return wordLemmaPairTMP[i].lWordSeq < wordLemmaPairTMP[j].lWordSeq
			} else {
				if (wordLemmaPairTMP[i].lL_W != wordLemmaPairTMP[j].lL_W) {
					return wordLemmaPairTMP[i].lL_W < wordLemmaPairTMP[j].lL_W
				} else {
					return wordLemmaPairTMP[i].lLemma < wordLemmaPairTMP[j].lLemma
				}	
			}
		} )	 
	//---------------
	/**
	for z1, lemX := range wordLemmaPairTMP {
		fmt.Println("XXXXXXXX   lista wordLemmaPair TMP z1=", z1, " => ", lemX)
	}	
	**/
	/***
	esempio 		
		antoniocigna antoniocigna _lemma_is_missing 0 -1 []}
		antoniocigna antoniocigna _lemma_is_missing 9 -1 []}
		
		mut          mut          mut               1 -1 []}
		 
		personen     personen     _lemma_is_missing 0 -1 []}
		personen     personen     person            1 -1 []}
		personen     personen     person            1 -1 []}
		personen     personen     _lemma_is_missing 9 -1 []}
	**/
	//------------------
	preWord:=""	
	//-----------------
	nnOut:=0
	//-------------------------------------
	preTutto:=""
	tuttoCodice := ""
	for _, lemX := range wordLemmaPairTMP {
		//fmt.Println("leggo " , lemX)
		//-------------------------------------
		if strings.ToLower(lemX.lWord2) == "abend"    {fmt.Println( " leggo wordLemmaPairTMP ", lemX) }
		if strings.ToLower(lemX.lWord2) == "antonio"  {fmt.Println( " leggo wordLemmaPairTMP ", lemX) }
		tuttoCodice = fmt.Sprint( lemX.lWordSeq , "-" , lemX.lL_W , "-" , lemX.lLemma) 
		if tuttoCodice == preTutto { continue }
		preTutto = tuttoCodice 
		if strings.ToLower(lemX.lWord2) == "abend"    {fmt.Println( " 2leggo wordLemmaPairTMP ", lemX) }
		if strings.ToLower(lemX.lWord2) == "antonio"  {fmt.Println( " 2leggo wordLemmaPairTMP ", lemX) }
		
		if lemX.lWordSeq > preWord {
			sw_WW00 = false
			sw_WW11 = false
			if lemX.lL_W == 0 { 
				sw_WW00 = true;  // significa che è presente almeno un 'word' del testo
				preWord = lemX.lWordSeq	
				continue 
			}
		}
		preWord = lemX.lWordSeq	
		if lemX.lL_W == 1 { 
			if sw_WW00 == false { // non è presente nemmeno una 'word' , allora ignora tutte le coppia word-lemma 
				//fmt.Println("  sw_WW00=", sw_WW00, "   ignoro appena letto") 
				continue
			}
			sw_WW11 = true   // è presente almeno una coppia word-lemma
		}
		if lemX.lL_W == 9 {
			if sw_WW11 == true { // esiste alemeno una coppia word-lemma, non serve aggiungere un lemma false '_lemma_i_missing_'			
				//fmt.Println("  sw_WW11=", sw_WW11, "   ignoro appena letto") 
				continue 
			}
		}		
		//fmt.Println("SCRIVO " , lemX, "     xxx   sw_WW00=", sw_WW00, " sw_WW11=", sw_WW11)
		nnOut++
		wordLemmaPair = append( wordLemmaPair, lemX)
		
		//fmt.Println("caricate wordLemmaPair ", lemX)
		
	}	
	//-------------------------------------
	for _, lemX := range wordLemmaPair {
		if strings.ToLower(lemX.lWord2) == "abend"    {fmt.Println( " prima del sort LEGGO wordLemmaPair ", lemX) }
		if strings.ToLower(lemX.lWord2) == "antonio"  {fmt.Println( " prima del sort LEGGO wordLemmaPair ", lemX) }
	}	
	//------------------
	fmt.Println("nnOut=", nnOut,  " len( wordLemmaPair)=", len( wordLemmaPair))
	//--------------------------------------------------------------
	// sort x lemma, L_W, l_Word2               	
	sort.Slice(wordLemmaPair, func(i, j int) bool {
			if (wordLemmaPair[i].lLemma != wordLemmaPair[j].lLemma) {
				return wordLemmaPair[i].lLemma < wordLemmaPair[j].lLemma
			} else {
				if (wordLemmaPair[i].lWord2 != wordLemmaPair[j].lWord2) {
					return wordLemmaPair[i].lWord2 < wordLemmaPair[j].lWord2
				} else {
					return wordLemmaPair[i].lL_W < wordLemmaPair[j].lL_W
				}	
			}
		} )	 
	//------------------------------	
	
	preLemma:=  ""
	numLemmaAdded:=0
	numLemmaOrig:=0 
	doppi:=0
	z:=-1 // minus 1
	fromIx:=0
	toIx:=0
	numW:=0
	L_W_doppi:=0
	L_W_aggiunti:=0
	//------------------------------
	fmt.Println("  2    len( wordLemmaPair)=", len( wordLemmaPair))
	//-----------------------------------
	//zz:=0
	
	for _, lemX := range wordLemmaPair {
		if strings.ToLower(lemX.lWord2) == "abend"    {fmt.Println( " LEGGO wordLemmaPair ", lemX) }
		if strings.ToLower(lemX.lWord2) == "antonio"  {fmt.Println( " LEGGO wordLemmaPair ", lemX) }
		//fmt.Println(zz, " LEGGO wordLemmaPair ", lemX)
		/**
		lemS12 = lemX.lLemma + " " + lemX.lWord2 
		if preLemS12 == lemS12 {   // se c'è il caso lL_W = 9, è questo ad essere scartato  
			if lemX.lL_W == 9 { L_W_doppi++ }
			doppi++
			continue
			
		}	
		preLemS12 = lemS12
				
		if preLemma != lemX.lLemma { 	
			if lemX.lL_W == 9 {  L_W_aggiunti++ }
		}
		**/
		
		//wordLemmaPair = append( wordLemmaPair, lemX)
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

	} // end for wordLemmaPair
	//------------------------------
	
	if numW > 0 {
		//scrive lemma precedente 				
		numLemmaOrig, numLemmaAdded	= appendOneLemma( preLemma, fromIx, toIx, numLemmaOrig, numLemmaAdded)	
	}
	
	//------------------------------------
	if doppi > 0 {
		fmt.Println(" scartate ", (doppi - L_W_doppi), " entrate doppie in lemma - word ") 
		fmt.Println("      scartate ", L_W_doppi, " coppie doppie perché L_W (origine parole da righe di testo) ")
	}
	fmt.Println(" aggiunte ", L_W_aggiunti , " coppie ottenute dalle parole da righe di testo)")
	
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
	leV.leFromIxLW = -1 
	leV.leToIxLW   = -2  
	leV.leTran     = ""
	leV.leLevel    = ""   
	leV.lePara     = ""   
	leV.leExample  = ""   
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
