package wbfSubPack

import (  
	"fmt"
    "strings"
	"sort"
)
//-----------------------------------------------

func g25_3read_dictLemmaTran_file(path1 string, inpFile string, swUp bool) {
	bytesPerRow:= 10
    lineD := rowListFromFile( path1, inpFile, "traduzione lemma", "read_dictLemmaTran_file", bytesPerRow)  
	if len(lineD) == 0 {sw_stop = false }
	if sw_stop { return }
	
	var ele1 lemmaTranStruct      
	
	//---------------
	cod1:= "" 	
	if swUp { lastNumDict = 100000}
	//-----------------
	for z:=0; z< len(lineD); z++ { 
				
		lineCol := strings.Split( (lineD[z]+"|||"), "|" ) 
		
		lastNumDict ++;
		cod1 =  stdCode( strings.TrimSpace( lineCol[0] ) )
		ele1.dL_lemma    = cod1 
		ele1.dL_numDict  = lastNumDict 
		ele1.dL_tran     = strings.TrimSpace( lineCol[1] ) 
		
		if swUp {
			dictLemmaTranUP = append( dictLemmaTranUP, ele1 ) 	
		} 
		dictLemmaTran = append( dictLemmaTran, ele1 ) 	
				
	}	
	//-----------------
	sort_lemmaTranUP()
	//---------------
	
	fmt.Println("read_dictLemmaTran_file", " len(lineD)=", len(lineD), " len( dictLemmaTran)=", len(dictLemmaTran) )
	
	sort_lemmaTran2();
	
	fmt.Println( len(dictLemmaTran) , "  lemma - translation elements of  dictLemmaTran" , "( input: ", inpFile, ")"  )   
	
} // end of read_dictLemmaTran_file  


//---------------------------------------

func sort_lemmaTranUP() {
	
	if len(dictLemmaTranUP) < 1 { return }	
	sort.Slice(dictLemmaTranUP, func(i, j int) bool {
			if (dictLemmaTranUP[i].dL_lemma != dictLemmaTranUP[j].dL_lemma) { 
				return dictLemmaTranUP[i].dL_lemma < dictLemmaTranUP[j].dL_lemma 
			} else {
				return dictLemmaTranUP[i].dL_numDict < dictLemmaTranUP[j].dL_numDict				
			}
		} )		
		
}  // end of sort_lemmaTranUP
//---------------------------------

func sort_lemmaTran2() {
	
	if len(dictLemmaTran) < 1 { return }
	
	sort.Slice(dictLemmaTran, func(i, j int) bool {
			if (dictLemmaTran[i].dL_lemma != dictLemmaTran[j].dL_lemma) { 
				return dictLemmaTran[i].dL_lemma < dictLemmaTran[j].dL_lemma 
			} else {
				return dictLemmaTran[i].dL_numDict < dictLemmaTran[j].dL_numDict				
			}
		} )		
	//------------	
	
	
} // end of sort_lemmaTran2() 

//-------------------------