package wbfSubPack

import (
    //"fmt"
	"strings"
	"sort"
)

//-----------------------------------------------
func get_separablePrefix() {
	//----------------------------------------
	//  "ab, an, auf, aus, bei, ein, fern,   her, herein, hin, hinaus, los, mit, nach, vor, vorbei, weg, weiter, zu, zurück, zusammen, zuruck"  
	var stringPrefissiSeparabili = ""+ 
		"ab - da;" + 
		"an  - a, addosso       ;"	+ 	
		"auf - su, verso l'alto ;"  +
		"aus - fuori, all'infuori;" + 
		"bei   - a;" +
		"ein - dentro, verso l'interno;" + 
		"fern - lontano;" +
		"her - avanti;"+
		"herein - dentro;" +
		"hin - a;" +
		"hinaus - fuori;" +
		"los - andare;" + 
		"mit - con;" +
		"nach - a;" + 
		"vor  - davanti, prima;" +
		"vorbei- passato;" + 
		"weg - lontano;" +
		"weiter - più lontano, continuare;"   +
		"zu     - chiuso, verso il basso;" +
		"zurück - indietro, ritorno;"      + 		 
		"zusammen - insieme;" + 
		"zuruck - indietro;" + 
		"";
		
	//----------------------------------------
	
	separPrefList = make([]separPrefStruct, 0, 200) 
	var sP separPrefStruct 
	listPr:= strings.Split(  strings.ToLower(stringPrefissiSeparabili),  ";" ) 	
	for _,lP:= range listPr {
		if lP == "" {continue}
		lP+= " - - - "
		pp:= strings.Split(lP,"-")
		
	    sP.sPrefix  = strings.TrimSpace( pp[0] ) 
		sP.sPrefTran= strings.TrimSpace( pp[1] ) 
		sP.sLenPref = len(sP.sPrefix) 
		separPrefList = append(separPrefList, sP) 
	}	
	//----- sort firstly the longest prefix, then ascending prefix order   
	sort.Slice(separPrefList, func(i, j int) bool { 
			if separPrefList[i].sLenPref != separPrefList[j].sLenPref {
				return separPrefList[i].sLenPref > separPrefList[j].sLenPref
			} else {			
				return separPrefList[i].sPrefix < separPrefList[j].sPrefix 		 	
			} 
		})	 	
	/**
	fmt.Println("\n\nXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	for _, p:= range separPrefList {
		fmt.Println("separable prefix ", p.sPrefix , " \t ", p.sPrefTran)
	}
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\n")
	**/
}
//---------------------------------------------------------------------------------
