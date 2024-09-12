package wbfSubPack

import (
    "fmt"
	"strings"
	"sort"
)
//-----------------------------------------------

//--------------------------------------------
func read_languageFile( path1 string, inpFile string) {
	
	fmt.Println( green( "read_languageFile" ) )
	
	bytesPerRow:= 40
    righe := rowListFromFile( path1, inpFile, "paradigma", "read_ParadigmaFile", bytesPerRow)  
	if sw_stop { return }
	fmt.Println("\nletto language file ", inpFile  + " " , len(righe) , " righe") 
	
	var translate_chars_std_inp = "" 
	var translate_chars_std_out = ""
	var translate_chars_SEQ_inp = ""     
	var translate_chars_SEQ_out = "" 
	//-----------------------------------------------  

	var sepPref =""	
	
	for _, line0:= range righe {
		line := strings.TrimSpace( line0 )
		
		j1:= strings.Index( line, "//" )
		if j1 >0 { line = line[0:j1] }
		
		if len(line) < 1 { continue } 
		
		j1 = strings.Index( line, "=" ) 
		if j1 < 0 { continue}
		
		value1:= strings.TrimSpace(line[j1+1:])
		var1 := strings.TrimSpace(line[0:j1])	
		
		if line[0:9] == "sep_pref " {
			sepPref += strings.TrimSpace(line[9:]) +";"  
			continue
		}	
		
		switch var1 {
			case "chars_std_inp":  translate_chars_std_inp = value1
			case "chars_std_out":  translate_chars_std_out = value1
			case "chars_SEQ_inp":  translate_chars_SEQ_inp = value1
			case "chars_SEQ_out":  translate_chars_SEQ_out = value1
			case "sep_pref"     :  sepPref += strings.TrimSpace(line[9:]) +";"  
		}
	} 	
	
	translate_chars_std_inpList =  strings.Fields( strings.ReplaceAll(translate_chars_std_inp, ","," ") )	     
	translate_chars_std_outList =  strings.Fields( strings.ReplaceAll(translate_chars_std_out, ","," ") )	 
   
	translate_chars_SEQ_inpList =  strings.Fields( strings.ReplaceAll(translate_chars_SEQ_inp, ","," ") )	       
	translate_chars_SEQ_outList =  strings.Fields( strings.ReplaceAll(translate_chars_SEQ_out, ","," ") )	 
	
	//------------------------------
	fmt.Println("translate_chars_std_inpList = " , translate_chars_std_inpList, "\n"+"translate_chars_std_outList = ", translate_chars_std_outList ) 
	fmt.Println("translate_chars_SEQ_inpList = " , translate_chars_SEQ_inpList, "\n"+"translate_chars_SEQ_outList = ", translate_chars_SEQ_outList ) 
	
	get_separablePrefix( sepPref ) 
	
} // end of read_languageFile	

//-----------------------------------------------
func get_separablePrefix( stringPrefissiSeparabili string ) {
	//----------------------------------------
	/***
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
	****/	
		
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
	
	fmt.Println("\n\nXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	for _, p:= range separPrefList {
		fmt.Println("separable prefix ", p.sPrefix , " \t ", p.sPrefTran)
	}
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\n")
	
}
//---------------------------------------------------------------------------------
