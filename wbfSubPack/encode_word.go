package wbfSubPack

import (  
	//"fmt"
    "strings"
)

//=================================================
/**
var translate_chars_std_inp = "" ; // "ä,ö,ü,ß";     
var translate_chars_std_out = "" ; // "ae,oe,ue,ss"
//--   
var translate_chars_SEQ_inp = "ä,ö,ü,ß";      
var translate_chars_SEQ_out = "aä,oö,uü,ssß"
**/
// 	 can be used in case a character might be written in more than one way:   eg. ß = ss        
var translate_chars_std_inpList []string	 // filled by inputLanguage.txt file in func read_languageFile     
var translate_chars_std_outList []string	 // filled by inputLanguage.txt file in func read_languageFile       
//--   
//   to modify the the place of a character in the alphabetic sequence order: eg.  'aä' replaces 'ä' so that the 'ä' is immediately after the normal 'a' 
//					                eg. german ==>  inp = "ä,ö,ü,ß"  / out "aä,oö,uü,ssß"

var translate_chars_SEQ_inpList []string     // filled by inputLanguage.txt file in func read_languageFile           
var translate_chars_SEQ_outList []string	 // filled by inputLanguage.txt file in func read_languageFile     

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
