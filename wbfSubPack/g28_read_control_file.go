package wbfSubPack

import (  
	"fmt"
    "strings"
	"strconv"
)
//------------------------------------------------

func g28_read_control_file() {
	
	fmt.Println("file control: ", file_inputControl)
	
	bytesPerRow:= 40
    lineD := rowListFromFile( "", file_inputControl, "input control", "read_control_file", bytesPerRow)  
	if sw_stop { return }
	
	sw_nl_only  = false	
	
	
	trim1 := string(`"`)
	trim2 := string(`'`)
	_, err := strconv.Atoi("0")
	//--------	
	type ctlSk struct {
		var1 string
		val1 string
	}
	var c1 ctlSk
	ctlLine:= make([]ctlSk,0,20)
	var commonInputFolder string
	//---------------------------
	for _, fline00:= range( lineD ) {
	
		if (fline00 == "") {continue}    // ignore zero length line 
		fline:= strings.Split(fline00, "//")[0]           // ignore all after // 
		fline = strings.Split(fline,   "/*")[0]           // ignore all after /* 
		fli  := strings.Split(fline,   "=")               //  dictionary_folder=folder of the dictionary files,   or file = filename  
		if len(fli) < 2 { continue} 
		varia1 := strings.ToLower( strings.TrimSpace(fli[0]) ) 
		
		value1 := strings.TrimSpace(fli[1]) 
		value1 = strings.Trim(value1, trim1) 
		value1 = strings.Trim(value1, trim2)
		if varia1 == "commoninputfolder" {
			commonInputFolder = value1 
			continue
		}			
		c1.var1 = varia1
		c1.val1 = value1
		ctlLine = append(ctlLine, c1) 
	}
	//---------------	
	if commonInputFolder != "" {
		commonInputFolder = strings.ReplaceAll(commonInputFolder, "\\","/")
		if commonInputFolder[ len(commonInputFolder)-1:] != "/" {
			commonInputFolder += "/"
		}
	}
	fmt.Println("commonInputFolder=", commonInputFolder)
	//----------------------------
	for _,line1:= range ctlLine {
		varia1:= line1.var1
		value1:= line1.val1
		
		rowArrayCap   := 0	
		wordSliceCap  := 0
		uniqueWordsCap:= 0
		//-------------------
		switch varia1 {
			case "word_lemma_file"     	 : FILE_inputWordLemma      = commonInputFolder + value1
			case "word_lemmaplus_file" 	 : FILE_inputWordLemmaPlus  = commonInputFolder + value1	
			case "paradigma_file"      	 : FILE_inpParadigma        = commonInputFolder + value1	
			case "inputlanguage_file"  	 : FILE_inputLanguage       = commonInputFolder + value1
			case "inputtranslation_file" : FILE_inputTranslation  	= commonInputFolder + value1
				
			case "write_numbered_text" :     //       = true 
				sw_write_numbered_text = (value1 == "true") 				
			case "read_numbered_text"  :     //         = false 
				sw_read_numbered_text = (value1 == "true") 
			case "out_numbered_text_fn_prefix"   :     //          = outNumText.txt
				fPrefix_out_numbered_text = strings.ReplaceAll(value1,"\\","/") 
			case "inp_numbered_text_fn_prefix"   :     //          = inpNumText.txt
				fPrefix_inp_numbered_text = strings.ReplaceAll(value1,"\\","/") 
			case "out_number_begin" :
				outNumberBegin, err = strconv.Atoi(value1)
				if err != nil { outNumberBegin = 0;}  	
				
				
			//------------
			case "max_num_lines" :
				rowArrayCap, err = strconv.Atoi(value1)
				if err != nil { rowArrayCap = 0;}  
				if rowArrayCap > 0 {
					inputTextRowSlice    = make( []rowStruct,   0, rowArrayCap) 					
					isUsedArray          = make( []bool       , 0, rowArrayCap)  
					//dictionaryRow        = make( []rDictStruct, 0, rowArrayCap)    
					fmt.Println("max_num_lines     :", rowArrayCap, " (inputTextRowSlice capacity)")  
				}
				
			case "max_num_words" : 
				wordSliceCap, err = strconv.Atoi(value1)
				if err != nil { wordSliceCap = 0;}  
				if wordSliceCap > 0 {
					wordSliceAlpha = make([]wordStruct, 0, wordSliceCap)   
					fmt.Println("max_num_words     :", wordSliceCap, " (wordSliceFreq capacity)")  
				}
				
			case "max_num_unique":
				uniqueWordsCap, err = strconv.Atoi(value1)
				if err != nil { uniqueWordsCap = 0;}  
				if uniqueWordsCap > 0 {
					uniqueWordByFreq    = make([]wordUnFreqStruct, 0, uniqueWordsCap)  					
					dictionaryWord      = make([]wDictStruct,  0, uniqueWordsCap)  
					fmt.Println("max_num_uniques   :", uniqueWordsCap, " (uniqueWordsByFreq capacity)")  
				}	
			/**	
			case "text_split_ignore_newline" :           // if true, newLine Character (\n) are ignored and the text is split only by full stop or any of other character as .;:!?    
				value1 = strings.ToLower(value1)					
				fmt.Println("text_split_ignore_newline :", value1)  
				if value1 == "true" {   
					sw_ignore_newLine = true 
				}
				
			case "text_split_by_newline_only" :   		
				value1 = strings.ToLower(value1)					
				fmt.Println("text_split_by_newline_only :", value1)  
				if value1 == "true" {  
					sw_nl_only = true
				}  				
							
			case "main_text_file"  :
				main_input_text_file = value1 				
			**/
			
			case "rewrite_word_lemma_dictionary" :
				sw_rewrite_wordLemma_dict = (value1 == "true") 				
			
		}
		
		
    } // end for fline00 range
	//-----------------------------------
	
	fmt.Println("\t", "word_lemma_file       = " , strings.ReplaceAll( FILE_inputWordLemma ,  commonInputFolder,""), 
			"\n\t",   "word_lemmaplus_file   = " , strings.ReplaceAll( FILE_inputWordLemmaPlus,  commonInputFolder,""), 
			"\n\t",   "paradigma_file        = " , strings.ReplaceAll( FILE_inpParadigma,     commonInputFolder,""), 
			"\n\t",   "inputLanguage_file    = " , strings.ReplaceAll( FILE_inputLanguage   , commonInputFolder,""), 
			"\n\t",   "inputTranslation_file = " , strings.ReplaceAll( FILE_inputTranslation, commonInputFolder,""), 
			"")
				
	fmt.Println("sw_write_numbered_text  = ", sw_write_numbered_text, 
				", sw_read_numbered_text   = ", sw_read_numbered_text )
				
	if sw_read_numbered_text {			
	} else {	
		if main_input_text_file == "" {
			sw_stop = true 	
			msg1_Js:= `in InputControl il parametro main_input_text_file è vuoto, ma sw_read_numbered_text = false` 			
			errorMSG = `<br><br><span style="font-size:0.7em;color:black;">` + msg1_Js + `</span>` 				
			showErrMsg2(errorMSG, msg1_Js)	
			return
		}
		// when main text file is read, the numbered file must be written  in any case	
		sw_write_numbered_text = true; 
	}
	if sw_write_numbered_text {			
		if fPrefix_out_numbered_text != "" {	
			fname_out_numbered_text = FOLDER_outNumText + "/" + fPrefix_out_numbered_text	
			fmt.Println("fname_out_numbered_text  =" + fname_out_numbered_text  + " first number to use=", outNumberBegin); 
		} else {
			//fmt.Println("required out numbered text file, but \"fname_out_numbered_text\" not specified in inputControl")  
		}
	}
	
	//---------------------
   
	
	if sw_ignore_newLine && sw_nl_only {	
		fmt.Println("text_split_ignore_newline = true and text_split_by_newline_only = true,  this is incompatible, both are ignored"  )    
	}
	
	if (sw_read_numbered_text) {
		fmt.Println("since " + "sw_read_numbered_text  = ", sw_read_numbered_text, ", then main text file is ignored")			
	} 
	
    
}  // end of read_control_file()
//----------------------------------------
