package wbfSubPack

	import (
		"fmt"
		"os"
		"strings"		
		"strconv"
		"time"
	)
//--------------------------------------------------------
func bind_go_write_lang_dictionary( langAndVoiceName  string) { 
		
		outFileName		:= FOLDER_INPUT_OUTPUT + string(os.PathSeparator) + FILE_language;  		
		
		//fmt.Println("bind_go_write_lang_dictionary file=" + outFileName +"\n\t",  langAndVoiceName );  	
		
		if langAndVoiceName[0:9] != "language=" {
			fmt.Println("bind_go_write_lang_dictionary() ERROR =>" +  langAndVoiceName[0:9] + "<==");  
			return 
		}	
		
		f, err := os.Create( outFileName )
		check(err)
		defer f.Close();

		_, err = f.WriteString( langAndVoiceName )
		check(err)

		f.Sync()
		
} // end of bind_go_lang_word_dictionary 	

//-----------------------------------------------------------
func bind_go_write_word_dictionary( listGoWords string) { 
		
		// listGoWords = list of NEW translated words
		
		//fmt.Println( "GO ",red("bind_go_write_word_dictionary"),  listGoWords  );  	
		
		//js_console_log("GO esegue bind_go_write_word_dictionary ")
	
		
		if len(listGoWords) < 1 { return }
		if len(listGoWords) > 9 {
			if listGoWords[0:9] == "language=" { return }
		}
			
		lemmaTranStr := ""; //  "__" + outFileName + "\n" + "_lemma	_traduzione"
		lastNumDict++; 
		lemmaTranStr += split_ALL_word_dict_row( listGoWords )	
		
		//js_console_log("GO bind_go_write_word_dictionary " + "listGoWords=" + listGoWords   + "lemmaTranStr=" + lemmaTranStr )
		
		sort_lemmaTran2();  // sort_lemmaX2 in write_word_dict... utilizzabili già in questo run 	
		
		/**------------------
		f, err := os.Create( outFileName )
		check(err)
		defer f.Close();

		_, err = f.WriteString( lemmaTranStr );  
		check(err) 

		f.Sync()
		**/
		
		//----------------------
		
		//sort_lemmaTran();  // sort_lemma2 in write_word_dict... utilizzabili già in questo run 
		
		rewrite_LemmaTranDict_file() 
		
		//js_console_log("GO fine esecuzione bind_go_write_word_dictionary ")
		
} // end of bind_go_write_word_dictionary 	

//-----------------------------------------------------------

func split_one_word_dict_row( row1 string ) (string, int, []string, []string) {

	// eg. einem;14; ein§einem§einer	
	
	lemmaLis := make( []string,0,0 )
	tranLis  := make( []string,0,0 )
	
	if row1 == "" { return "",	-1,	lemmaLis, tranLis }                           
	
	//row1:= strings.ReplaceAll( strings.ReplaceAll( row0, ....  parentesi quadre 
	
	var field = strings.Split( row1, ";");     // eg. einem;  14; ein§einem§einer; uno§uno§uno
											   //	 field[0]; 1;               2; 3 
	
	lemmaLis = strings.Split( strings.TrimRight(field[2],wSep), wSep)  // eg. ein , einem,   einer       
	tranLis  = strings.Split( strings.TrimRight(field[3],wSep), wSep)  //     uno , uno,     uno  
	
	//fmt.Println( green("split_one_word_dict_row"), "\n\t lemma=",len(lemmaLis), " ", strings.Join(lemmaLis, " - ")  ,    
	//			 "\n\t tranl=",len(tranLis), " ", strings.Join(tranLis, " - ") )	
	
	
	ix1, err := strconv.Atoi( field[1] )
	if err != nil { return "",	-1,	lemmaLis, tranLis }                           //error
	 
	return field[0],ix1, lemmaLis, tranLis     // eg. return einem; 14, [ein , einem,   einer], [uno , uno,     uno] 
		
} // end of	split_one_Word_Dict_row(	

//-------------------------------

//-------------------------------------
func bind_go_write_row_dictionary( listGoRows string) {
		/**
		2;Primo capitolo
		4;Gustav Aschenbach o von Aschenbach, come ha fatto sin dai suoi cinquant'anni
		5;compleanno, era il suo nome ufficiale, era l'uno
		**/
		
		//fmt.Println("bind_go_write_row_dictionary ()  esegue  update_rowTransLATION () ")
		
		update_rowTranslation(   strings.Split(listGoRows,"\n") ) ;
		
		//fmt.Println("bind_go_write_row_dictionary ()  esegue  writeTextRowSlice () ")
		
		writeTextRowSlice()  // after some new translated rows  the dictR file is rewritten  

} // end of  bind_go_write_row_dictionary 	

//--------------------------------------------
func writeTextRowSlice() {		

		nout:=0  
		lines:= make([]string, 0, 10+len( inputTextRowSlice) )
		
		for z:=0; z < len( inputTextRowSlice); z++ {
			rS1 := inputTextRowSlice[z]
			
			lines = append(lines, rS1.rIdRow + "|O|" + rS1.rRow1 )  
			
			if rS1.rTran1 != "" {
				lines = append(lines, rS1.rIdRow + "|T|" + rS1.rTran1)  
			}
			
			//fmt.Println("writeTextRowSlice () z=", z, "  lines[]= ", lines[ len(lines)-1 ] )
			nout++
		}  
		
		// riscrivi tutte le righe con text origine e traduzione  (ogni file ha data e ora nel nome, bisogna leggere l'ultimo, e ogni tanto cancellare i file vecchi)	
		
		currentTime := time.Now()	

		last_written_dict_rowFile = "dictR" + currentTime.Format("20060102150405") + ".txt"			
		//-----------------------
		outF1 		:= FOLDER_O_arc_TRAN_rows + string(os.PathSeparator)    		
		outFileName := outF1 + last_written_dict_rowFile				
		fmt.Println("wrote ", nout , " lines on ", outFileName ) 	
		writeList( outFileName, lines )	
		//-----------------
		outF12 		:= FOLDER_IO_lastTRAN + string(os.PathSeparator)    		
		outFileName2 := outF12 + FILE_last_updated_dict_rows				
		fmt.Println("wrote ", nout , " lines on ", outFileName2 ) 	
		writeList( outFileName2, lines )	
		//-----------------
		
		
} // end of bind_go_write_row_dictionary 	


//-----------------------------------------------

func bind_go_passToJs_write_WordsToLearn(js_function string)  {
						
	outFile := FOLDER_INPUT_OUTPUT  + string(os.PathSeparator) + FILE_words_to_learn ;

	lines:= make([]string, 0, len(uniqueWordByFreq) )
	
	numOut:=0
	for i:= 0; i < len(uniqueWordByFreq); i++ { 
		var xWordF     = uniqueWordByFreq[i]
		if ( (xWordF.uKnow_yes_ctr < 1 ) && (xWordF.uKnow_no_ctr < 1 ))  { continue }
		lines = append( lines,   fmt.Sprintf( "%s|%d|%d",  xWordF.uWord2, xWordF.uKnow_yes_ctr,  xWordF.uKnow_no_ctr) ) 
		numOut++	
	}		
	writeList( outFile, lines )
	//--------
	currentTime := time.Now()	
	outF1 		:= FOLDER_O_arc_TO_learn + string(os.PathSeparator)  
	outFile2    := outF1 + "wordsToLearn_"  + currentTime.Format("20060102150405") + ".txt"	
	writeList( outFile2, lines )	
	//-----------------	
	
	if (js_function == "") { return }
	outS1:= fmt.Sprintf("scritti %d words to learn nel file %s", numOut,  FILE_words_to_learn)
	
	go_exec_js_function( js_function, outS1 ); 	
	
} // end of  bind_go_passToJs_write_WordsToLearn

//----------------------------------------
