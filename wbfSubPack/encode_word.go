package wbfSubPack

import (  
	//"fmt"
    "strings"
)

//=================================================
//    metti queste 2 righe su file, nello stesso file metti prefissi separabili e non separabili con identificatore di separab/non-separab 

var translate_chars_std_inp = "" ; // "ä,ö,ü,ß";     
var translate_chars_std_out = "" ; // "ae,oe,ue,ss"
//--   
var translate_chars_SEQ_inp = "ä,ö,ü,ß";      
var translate_chars_SEQ_out = "aä,oö,uü,ssß"

var translate_chars_std_inpList =  strings.Fields( strings.ReplaceAll(translate_chars_std_inp, ","," ") )	     
var translate_chars_std_outList =  strings.Fields( strings.ReplaceAll(translate_chars_std_out, ","," ") )	 
//--   
var translate_chars_SEQ_inpList =  strings.Fields( strings.ReplaceAll(translate_chars_SEQ_inp, ","," ") )	       
var translate_chars_SEQ_outList =  strings.Fields( strings.ReplaceAll(translate_chars_SEQ_out, ","," ") )	 

//=================================================

func stdCode(inpCode string ) string {		
	
	if len(translate_chars_std_inpList) == 0 {return inpCode}
	
	outCode:= inpCode 
	for nn,ch1:= range translate_chars_std_inpList {
		outCode = strings.ReplaceAll(outCode, ch1, translate_chars_std_outList[nn])  
	}		
	
	return outCode  

} // end of stdCode

//----------------------------
func seqCode( inpCode string ) string {	
	/*
	serve soprattutto per mettere le parole in sequenza alfabetica più naturale di quella dettata dal codifica asci o utf8 
	es. per il tedesco es.  ä, ö, ü, ß vicini rispettivamente ad a, o, u, ss)   	
	*/	
	
	if len(translate_chars_SEQ_inpList) == 0 {return inpCode}
	
	outCode:= inpCode 
	for nn,ch1:= range translate_chars_SEQ_inpList {
		outCode = strings.ReplaceAll(outCode, ch1, translate_chars_SEQ_outList[nn])  
	}		
	return outCode  
	
}// end of seqCode					

//------------------------------------------------
