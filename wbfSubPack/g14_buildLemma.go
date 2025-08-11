package wbfSubPack

import (  
	"fmt"
    "strings"
	"sort"
)
//--------------------------------------------------------------------------

func TOGLIestraeCoppieWordLemmaInTesto( wordLemmaPairTMP []wordLemmaPairStruct) {
	
	//sw_WW00:= false
	//sw_WW11:= false
	//--------------------------------------
	// sort x lemma, L_W, l_Word2               	
	sort.Slice(wordLemmaPairTMP, func(i, j int) bool {
			if (wordLemmaPairTMP[i].lWordSeq != wordLemmaPairTMP[j].lWordSeq) {
				return wordLemmaPairTMP[i].lWordSeq < wordLemmaPairTMP[j].lWordSeq
			} else {				
				return wordLemmaPairTMP[i].lLemma < wordLemmaPairTMP[j].lLemma					
			}
		} )	 
	//---------------	
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


	fmt.Println( red("variato , eliminato L_W   ( = 0,1 , 9  sostituire con ricerca diretta del lemma "))	



	wordLemmaPairEXTR := make([]wordLemmaPairStruct, 0 , len(wordLemmaPairTMP) )
	//num11:=0; //tot11:=0; totMiss:=0
	for _, lemX := range wordLemmaPairTMP {		
		tuttoCodice = fmt.Sprint( lemX.lWordSeq , "-" ,  lemX.lLemma) 
		if tuttoCodice == preTutto { continue }
		preTutto = tuttoCodice 
		
		if lemX.lWordSeq > preWord {
			//sw_WW00 = false
			//sw_WW11 = false
			/***
			if lemX.lL_W == 0 { 
				sw_WW00 = true;  // significa che è presente almeno un 'word' del testo
				preWord = lemX.lWordSeq	
				//tot11 += num11
				//num11=0
				continue 			}
			***/
		}
		preWord = lemX.lWordSeq	
		/**
		if lemX.lL_W == 1 { 
			if sw_WW00 == false { // non è presente nemmeno una 'word' , allora ignora tutte le coppia word-lemma 
				continue
			}
			//sw_WW11 = true   // è presente almeno una coppia word-lemma
			//num11++
		}
		**/
		/***
		if lemX.lL_W == 9 {
			//if sw_WW11 == false { totMiss+= num11 }
			continue 
		}	
		***/		
		nnOut++
		wordLemmaPairEXTR = append( wordLemmaPairEXTR, lemX)		
	}	
	//------------------------------------------
	
	sort.Slice(wordLemmaPairEXTR, func(i, j int) bool {
			if (wordLemmaPairEXTR[i].lLemma != wordLemmaPairEXTR[j].lLemma) {
				return wordLemmaPairEXTR[i].lLemma < wordLemmaPairEXTR[j].lLemma
			} else {
					return wordLemmaPairEXTR[i].lWord2 < wordLemmaPairEXTR[j].lWord2				 
			}
		} )	 	
	//-------------------------------------------
	wordLemmaPairNEW := make([]wordLemmaPairStruct, 0 , len(wordLemmaPairTMP) )
	var NW wordLemmaPairStruct 
	//-----------------------------------
	for zz:=0; zz < len( listAllLemmaFromFile); zz++ {
		nome:= listAllLemmaFromFile[zz] + "<="
		if strings.Index(nome,"gehen<=") >= 0  {
			fmt.Println( "listAllLemmaFromFile[",zz, "]= ", listAllLemmaFromFile[zz]  )		
		}
	} 
	ixAF2 := binarySearch_string(listAllLemmaFromFile, "gehen") 
	fmt.Println(`binarySearch_string(listAllLemmaFromFile, "gehen") --> ixAF2=`, ixAF2)
	//--------------------------
	preLe:=""
	listLem1:= make([]string,0,100)
	sw1:= false
	for _, WL2:= range wordLemmaPairEXTR {
		sw1 = ( WL2.lLemma == "gehen" ) 
		if sw1 { fmt.Println( green("WL2=" + WL2.lLemma) )  }
		if WL2.lLemma != preLe {
			listLem1 = make([]string,0,100)//nuovo lemma
			for _, p:= range separPrefList {			
				newLemma := p.sPrefix + WL2.lLemma
				ixAF := binarySearch_string(listAllLemmaFromFile, newLemma) 
				if sw1 { fmt.Println( "  newLemma=",  newLemma, "    ixAF=", ixAF)   }
				if (ixAF < 0) { continue }
				listLem1 = append(listLem1, newLemma) 				
			}	
			preLe = WL2.lLemma
		}
		if len(listLem1) < 1 {continue}
		NW = WL2
		for _,newL:= range listLem1 {
			NW.lLemma = newL
			wordLemmaPairNEW = append(wordLemmaPairNEW, NW) 
		} 
	}	
	//------
	fmt.Println("listAllLemmaFromFile len=", len(listAllLemmaFromFile) )
	fmt.Println("wordLemmaPairEXTR    len=", len(wordLemmaPairEXTR) )
	fmt.Println("separable prefix     len=", len(separPrefList) )
 	fmt.Println("wordLemmaPairNEW     len=", len(wordLemmaPairNEW) )
	
	return	
	
	/**
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
 	**/
	/**
	loop 
	per ogni lemma di wordLemmaPairEXTR
		premetti ad ognuno tutti i prefissi e controlla che il nuovo lemma risultante esista
		
		var ixAF = binarySearch_string(listAllLemmaFromFile, nuovoLemmaPref+OldLemma) 
		if (ixAf >= 0 ) {
			se trovato aggiungi  tutte le coppie word-oldLemma facendole diventare word-newLemma   
			
		}
	
	//---------------------------
	
	//-------------------
	fmt.Println("estraeCoppieWordLemmaInTesto:  trovate ", len(wordLemmaPairEXTR), " coppie word-lemma, potrebbero esistere parole di testo senza coppia ");  
	***/ 
	
} // end of estraeCoppieWordLemmaInTesto		
//--------------------------------------------------------------------------

func TOGLIg14_buildListLemmaSlice( wordLemmaPairTMP []wordLemmaPairStruct) {
	
	//preLemS12:=  ""
	//lemS12:= ""		
	
	//sw_WW00:= false
	//sw_WW11:= false
	//--------------------------------------
	// sort x lemma, L_W, l_Word2               	
	sort.Slice(wordLemmaPairTMP, func(i, j int) bool {
			if (wordLemmaPairTMP[i].lWordSeq != wordLemmaPairTMP[j].lWordSeq) {
				return wordLemmaPairTMP[i].lWordSeq < wordLemmaPairTMP[j].lWordSeq
			} else {
					return wordLemmaPairTMP[i].lLemma < wordLemmaPairTMP[j].lLemma
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
		//if strings.ToLower(lemX.lWord2) == "abend"    {fmt.Println( " leggo wordLemmaPairTMP ", lemX) }
		//if strings.ToLower(lemX.lWord2) == "antonio"  {fmt.Println( " leggo wordLemmaPairTMP ", lemX) }
		tuttoCodice = fmt.Sprint( lemX.lWordSeq , "-" ,  lemX.lLemma) 
		if tuttoCodice == preTutto { continue }
		preTutto = tuttoCodice 
		//if strings.ToLower(lemX.lWord2) == "abend"    {fmt.Println( " 2leggo wordLemmaPairTMP ", lemX) }
		//if strings.ToLower(lemX.lWord2) == "antonio"  {fmt.Println( " 2leggo wordLemmaPairTMP ", lemX) }
	  fmt.Println(red("VARIAto,  eliminato L_W, sostituito con ricerca diretta del lemma o della parola") )	
		if lemX.lWordSeq > preWord {
			//sw_WW00 = false
			//sw_WW11 = false
			/**
			if lemX.lL_W == 0 { 
				sw_WW00 = true;  // significa che è presente almeno un 'word' del testo
				preWord = lemX.lWordSeq	
				continue 
			}
			**/
		}
		preWord = lemX.lWordSeq	
		/**
		if lemX.lL_W == 1 { 
			if sw_WW00 == false { // non è presente nemmeno una 'word' , allora ignora tutte le coppia word-lemma 
				//fmt.Println("  sw_WW00=", sw_WW00, "   ignoro appena letto") 
				continue
			}
			sw_WW11 = true   // è presente almeno una coppia word-lemma
		}
		**/
		/**
		if lemX.lL_W == 9 {
			if sw_WW11 == true { // esiste alemeno una coppia word-lemma, non serve aggiungere un lemma false '_lemma_i_missing_'			
				//fmt.Println("  sw_WW11=", sw_WW11, "   ignoro appena letto") 
				continue 
			}
		}	
		**/		
		//fmt.Println("SCRIVO " , lemX, "     xxx   sw_WW00=", sw_WW00, " sw_WW11=", sw_WW11)
		nnOut++
		wordLemmaPair = append( wordLemmaPair, lemX)
		
		//fmt.Println("caricate wordLemmaPair ", lemX)
		
	}	
	//-------------------------------------
	/**
	for _, lemX := range wordLemmaPair {
		if strings.ToLower(lemX.lWord2) == "abend"    {fmt.Println( " prima del sort LEGGO wordLemmaPair ", lemX) }
		if strings.ToLower(lemX.lWord2) == "antonio"  {fmt.Println( " prima del sort LEGGO wordLemmaPair ", lemX) }
	}
	**/	
	//------------------
	fmt.Println("g14_buildListLemmaSlice ", "nnOut=", nnOut,  " len( wordLemmaPair)=", len( wordLemmaPair))
	//--------------------------------------------------------------
	// sort x lemma, L_W, l_Word2               	
	sort.Slice(wordLemmaPair, func(i, j int) bool {
			if (wordLemmaPair[i].lLemma != wordLemmaPair[j].lLemma) {
				return wordLemmaPair[i].lLemma < wordLemmaPair[j].lLemma
			} else {
				return wordLemmaPair[i].lWord2 < wordLemmaPair[j].lWord2
			}
		} )	 
	//------------------------------	
	
	//---------------------------------------
	preLemma:=  ""
	numLemmaAdded:=0
	numLemmaOrig:=0 
	//doppi:=0
	z:=-1 // minus 1
	fromIx:=0
	toIx:=0
	numW:=0
	//L_W_doppi:=0
	//L_W_aggiunti:=0
	//------------------------------
	fmt.Println("  2    len( wordLemmaPair)=", len( wordLemmaPair))
	//-----------------------------------
	//zz:=0
	
	for _, lemX := range wordLemmaPair {
		//if strings.ToLower(lemX.lWord2) == "abend"    {fmt.Println( " LEGGO wordLemmaPair ", lemX) }
		//if strings.ToLower(lemX.lWord2) == "antonio"  {fmt.Println( " LEGGO wordLemmaPair ", lemX) }
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
				numLemmaOrig, numLemmaAdded	= g14_appendOneLemma( preLemma, fromIx, toIx, numLemmaOrig, numLemmaAdded)
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
		numLemmaOrig, numLemmaAdded	= g14_appendOneLemma( preLemma, fromIx, toIx, numLemmaOrig, numLemmaAdded)	
	}
	
	//------------------------------------
	/**
	if doppi > 0 {
		fmt.Println(" scartate ", (doppi - L_W_doppi), " entrate doppie in lemma - word ") 
		fmt.Println("      scartate ", L_W_doppi, " coppie doppie perché L_W (origine parole da righe di testo) ")
	}
	**/
	//fmt.Println(" aggiunte ", L_W_aggiunti , " coppie ottenute dalle parole da righe di testo)")
	
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
	//lemmaSliceUpdateSubLemma()
	
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
	
	/**
	// update  einStellenList
	for z2, wD:= range wordLemmaPair {

		//if z2 < 0 {	fmt.Println( "wordLemmaPair ", wD) }
		
		ix2 :=wD.lIxLemma
		if ix2 < 0 { continue}
		wD.lIx_einStellenList = lemmaSlice[ ix2 ].ls_lemma_einStellenList 
		wordLemmaPair[z2] = wD 
		//if strings.Index(wD.lLemma,"stellen") >=0  { fmt.Println( "??anto2 buildLemma ",  wD) }
	}  
	**/
	
	//-----------------------------	
	fmt.Println("") 
	/**
	for _, wD:= range wordLemmaPair {
		fmt.Println( "wordLemmaPair ", wD) 
	}  
	**/
	
	//-----------------------------
	
	//g34_loadInverseLemmaSlice()	spostato a dopo il caricamento del testo 
	
	//------------------------
	
} // end of TOGLIg14_buildListLemmaSlice	

//--------------------------------

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
	/**	
	leV.ls_lemma_ix_stellen = -1
	leV.ls_lemma_stellen    = ""		
	leV.ls_pref_ein         = "" 
	leV.ls_pref_tran        = ""	
	leV.ls_lemma_einStellenList = nil; 
	**/

	// eg. einstellen  = ein + stellen 	
	
	/**
	works even with multibyte characters: eg. if prefix="日本"  then   lemma="日本語語語"  is broken down into:  "日本" + "語語語"   
	**/
	/**
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
	**/
	
	numLemmaOrig++
	lemmaSlice = append(lemmaSlice, leV ) 
	iixLem = len(lemmaSlice) -1 
	for h:=fromIx; h<= toIx; h++ {
		wordLemmaPair[h].lIxLemma = iixLem    
		//wordLemmaPair[h].lIx_einStellenList = nil 
		//if ((leV.leLemma == "am") || (leV.leLemma == "wochenende"))  { fmt.Println("g14_appendOneLemma lemmaSlice=",  leV ,  " wordLemmaPair=", wordLemmaPair[h] ) }
	}	
	return numLemmaOrig, numLemmaAdded			
} // end of g14_appendOneLemma 

//---------------------