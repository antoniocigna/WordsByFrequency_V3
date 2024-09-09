package wbfSubPack

import (  
	"fmt"
    "strings"
	"sort"
)
//-----------------------------------------------

func read_dictLemmaTran_file(path1 string, inpFile string) {
	bytesPerRow:= 10
    lineD := rowListFromFile( path1, inpFile, "traduzione lemma", "read_dictLemmaTran_file", bytesPerRow)  
	if sw_stop { return }
	
	// 	abnutzbarkeit vestibilità   ==>  lemma     \t traduzione                                                                   |    |  |        
	
	lineZ := ""
	
	var ele1 lemmaTranStruct       //  lemmaTranStruct: dL_lemmaSeq string,  dL_lemma2 string, dL_tran string  
	
	//---------------
	cod1:= "" 	
	lastNumDict ++;
	//-----------------
	for z:=0; z< len(lineD); z++ { 
		
		lineZ = strings.TrimSpace(lineD[z]) 
		
		// eg. abnutzbarkeit \t	vestibilità     ==>  lemma \t translation  
		
		if lineZ == "" { continue }
		j1:= strings.Index(lineZ, "|")
		if j1 < 0 { continue }
		cod1 = lineZ[0:j1]
		ele1.dL_lemmaSeq = seqCode(cod1) 
		ele1.dL_lemma2   = strings.TrimSpace( cod1 )
		ele1.dL_numDict  = lastNumDict 
		ele1.dL_tran     = strings.TrimSpace( lineZ[j1+1:] )   ////cigna1_2
		
		
		dictLemmaTran = append( dictLemmaTran, ele1 ) 	
		
	}	
	
	fmt.Println("read_dictLemmaTran_file", " len(lineD)=", len(lineD), " len( dictLemmaTran)=", len(dictLemmaTran) )
	
	sort_lemmaTran2();
	
	fmt.Println( len(dictLemmaTran) , "  lemma - translation elements of  dictLemmaTran" , "( input: ", inpFile, ")"  )   
	
} // end of read_dictLemmaTran_file  


//---------------------------------------

//---------------------------------

func sort_lemmaTran2() {
	
	if len(dictLemmaTran) < 1 { return }
	
	sort.Slice(dictLemmaTran, func(i, j int) bool {
			if (dictLemmaTran[i].dL_lemmaSeq != dictLemmaTran[j].dL_lemmaSeq) { 
				return dictLemmaTran[i].dL_lemmaSeq < dictLemmaTran[j].dL_lemmaSeq 
			} else {
				return dictLemmaTran[i].dL_numDict < dictLemmaTran[j].dL_numDict				
			}
		} )		
	//------------	
	var pre lemmaTranStruct
	pre = dictLemmaTran[0]
	//nelle doppie mette codice X'ff'  valore massio di un byte
	for g2:=1; g2 < len(dictLemmaTran); g2++ {
		if (dictLemmaTran[g2].dL_lemmaSeq == pre.dL_lemmaSeq) {
			dictLemmaTran[g2 -1].dL_lemmaSeq = "";  //LAST_WORD; 
		} 
		pre = dictLemmaTran[g2] 
	}
	//--------------------------
	// sort in modo da sbattere i codice XX'ff ( cioè le doppie) alla fine 
	sort.Slice(dictLemmaTran, func(i, j int) bool {
			if (dictLemmaTran[i].dL_lemmaSeq != dictLemmaTran[j].dL_lemmaSeq) { 
				return dictLemmaTran[i].dL_lemmaSeq < dictLemmaTran[j].dL_lemmaSeq 
			} else {
				return dictLemmaTran[i].dL_numDict < dictLemmaTran[j].dL_numDict				
			}
		} )		
	//--------------------------------	
	var numLin = 0
	// cerca dove si trova il primo codice X'ff' per trovare la lunghezza effettiva dell'array
	
	firstIx:=-1
	
	for g2:=0; g2 < len(dictLemmaTran); g2++ {
		
		//if strings.Index(dictLemmaTran[g2].dL_lemmaSeq, "eindhoven") >= 0 { fmt.Println( "in sort_lemaa_tran2 ", " g2=", g2, " ",  dictLemmaTran[g2] ) }
		
		if (dictLemmaTran[g2].dL_lemmaSeq == "") {
			firstIx = g2 	
		} else {
			numLin++
		}
	}
	//fmt.Println("sort_lemmaTran2 ", " last blank=", firstIx, " len=",  len(dictLemmaTran)); 
	firstIx++
	if numLin > 0 {
		dictLemmaTran = dictLemmaTran[firstIx:]
	}
	
	
} // end of sort_lemmaTran2() 

//-------------------------