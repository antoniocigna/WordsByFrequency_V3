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
		
		fmt.Println( "GO ",red("bind_go_write_word_dictionary"),  listGoWords  );  	
		
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
	
	//lemmaLis = strings.Fields( strings.ReplaceAll(field[2],wSep," ") )       
	//tranLis  = strings.Fields( strings.ReplaceAll(field[3],wSep," ") )  
	lemmaLis = strings.Split( strings.Trim(field[2],wSep), wSep )  
	tranLis  = strings.Split( strings.Trim(field[3],wSep), wSep )  
	
	//fmt.Println("  xx split_one_word_dict_row(", row1, ") file=", field, " lemmaLis=", lemmaLis, "tranLis=", tranLis )  
	
	if len(lemmaLis) != len(tranLis)  {
		return "",	-1,	lemmaLis, tranLis    
	} 
		
	ix1, err := strconv.Atoi( field[1] )
	if err != nil { return "",	-1,	lemmaLis, tranLis }                           //error
	 
	return field[0],ix1, lemmaLis, tranLis     // eg. return einem; 14, [ein , einem,   einer], [uno , uno,     uno] 
		
} // end of	split_one_Word_Dict_row(	

//-------------------------------

//--------------------

func update_rowTranslation( rowDictRow [] string)  {
	/*
	rowDictRow is filled in javascript, each line: rowIndex;row Translation   
	*/
	var len1 = len(inputTextRowSlice)
	
	//fmt.Println("update_rowTraslation() len(rowDictRow) =",len(rowDictRow), " = ", rowDictRow)
	
	//ion() len(rowDictRow) =  [ 2 206(2_206 211|211|1. Introduzione Preistoria 2 207(2_207 212|212|La prei
	
	for z:=0; z < len(rowDictRow); z++ {  		
		/**
		2;Primo capitolo
		**/
		row1dict := rowDictRow[z] 	
		if row1dict == "" { continue }	
		rowCols	 := strings.Split(row1dict, "|")
		
		//fmt.Println("\trowCols= len=", len(rowCols), " rowcols=",    rowCols) 
		
		if len(rowCols) < 3 { 
			fmt.Println("update_rowTraslation() len(rowDictRow) z=", z, " row1dict=", row1dict, " len(rowCols)=", len(rowCols))
			continue 
		}
		
		idS   := rowCols[0]   
		ixS   := rowCols[1] 
			
		tranS := rowCols[2]
		
		//fmt.Println("\tidS=", idS, "  ixS=", ixS,  " tranS=", tranS ) 
		
	
		ixRow, err := strconv.Atoi(  strings.TrimSpace( ixS) )		
		
		//fmt.Println("\tixRow=", ixRow, " err=", err)
		
		if err != nil {
			return 
		}
		if ixRow >= len1 { return }  // error 	

		//fmt.Println("\tinputTextRowSlice[ixRow].rIdRow = ", inputTextRowSlice[ixRow].rIdRow )
		
		if strings.Index( idS, inputTextRowSlice[ixRow].rIdRow) < 0  {
			// error
			fmt.Println("error func update_rowTranslation () inputTextRowSlice[ixRow=", ixRow, "].rIdRow = ", inputTextRowSlice[ixRow].rIdRow,   " idS=",  idS )
			continue
		}  
		inputTextRowSlice[ixRow].rTran1 = strings.TrimSpace(tranS ) 
		
		//fmt.Println("update_rowTraslation() row1dict=", row1dict, " \n\t ixRow=", ixRow, " inputTextRowSlice[ixRow]=", inputTextRowSlice[ixRow] )
		
	} // end for z
	
} // end of update_rowTranslation

//---------------------------------------------------------
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

//-------------------------------
//  called by bind_go_write 
//----------------
func split_ALL_word_dict_row(  strRows string) string {
	//fmt.Println( "ANTONIO xxxxxxxxxxxxxxxxxxxxxxxxxxxx  split_ALL_word_dict_row( strRows=", strRows); 
	// eg. einem;14 ; ein§einem§einer ;  a§uno§uno;	  ==> word ; ix : list of lemmas ; list of translations	
	
	lemmaTranStr := ""
	
	lines := strings.Split( strRows, "\n");	
	
	var ele1 lemmaTranStruct     
	/**
	//--------------------------
		type lemmaTranStruct struct {
			dL_lemmaSeq  string 
			dL_lemma2    string  
			dL_numDict   int	  
			dL_tran      string
		} 
	**/
	
	lenIx1:= len(uniqueWordByFreq)

	
	for z:=0;  z < len(lines); z++ {   
		
		//fmt.Println("\nsplit_ALL_word_dict_row 1 lines[",z,"]=", lines[z] )	
		
		_, ix1, lemmaLis,tranLis := split_one_word_dict_row( lines[z] )
		
		//fmt.Println("\nsplit_ALL_word_dict_row 1 lines[",z,"]=", lines[z] , " ix1=", ix1, " lemmaLis=", lemmaLis, " tranLis=", tranLis) 	
		
		if ix1 < 0 { continue }		
		
		//fmt.Println("       split_ALL_word_dict_row 2 ", " ix1=", ix1)
		
		if ix1 >= lenIx1 { 
			fmt.Println("error7 1 len(uniqueWordByFreq)=", len(uniqueWordByFreq), " ix1=", ix1 ,  " lines[z=", z, "]=", lines[z] )
			continue 
		}
	
		//fmt.Println("       split_ALL_word_dict_row 3 ")
		
		//---------------
		if uniqueWordByFreq[ix1].uIxUnW != ix1 {	
			fmt.Println("error7 2 len(uniqueWordByFreq)=", len(uniqueWordByFreq), " ix1=", ix1 ,  " lines[z=", z, "]=", lines[z] )
			continue 
		}  // error	
		
		//fmt.Println("       split_ALL_word_dict_row 4 ")	
		
		/***
		 {ihren.ihren ihren 37 29 1 1 63 1 1 0 0 [25114 25115 25116 45916] [ihr ihre ihrer sein] [   ] [   ] [   ]}
		 
			uLemmaL= [ihr ihre ihrer sein]	
		
		***/
		
		//len1:= len(uniqueWordByFreq[ix1].uLemmaL)		
		
		//---------------------
		oneW := uniqueWordByFreq[ix1]
		
		//fmt.Println("       split_ALL_word_dict_row 5 ", " oneW=", oneW)
		
		for m, lemmaFreq:= range oneW.uLemmaL {
			
			newL:=-1
			for l1:=0; l1 < len(lemmaLis); l1++ {         // translation from dict tran file 
				if lemmaFreq == lemmaLis[l1] {
					newL = l1
					break	
				}			
			}	
			
			//fmt.Println("       split_ALL_word_dict_row 6 for 1 ")
			
			if newL < 0 { continue }	
			
			//fmt.Println("       split_ALL_word_dict_row 6 for 2 ")
			
			if lemmaLis[newL] != oneW.uLemmaL[m] { continue }     // error?   
				
			//fmt.Println("       split_ALL_word_dict_row 6 for 3 ")
			
			ixLe:= oneW.uIxLemmaL[m]
			lemmaSlice[ixLe].leTran    = tranLis[newL]            // update lemmaSlice lemmaStruct .le...
			//uniqueWordByFreq[ix1].uTranL[m]     = mTran 	
			//uniqueWordByAlpha[ixAlfa].uTranL[m] = mTran 
			
			lemmaTranStr += "\n" + lemmaLis[newL] + "|" + tranLis[newL]  	
			ele1.dL_lemmaSeq = seqCode(lemmaLis[newL] )                  //  lemmaTranStruct struct   .dL...
			ele1.dL_lemma2   = lemmaLis[newL] 
			ele1.dL_numDict  = lastNumDict
			ele1.dL_tran     = tranLis[newL]    
			
			//fmt.Println("       split_ALL_word_dict_row 6 for 4 ")
			
			if  tranLis[newL] != "" { 		
				dictLemmaTran = append( dictLemmaTran, ele1 ) 
				
				//fmt.Println("         split_ALL_word_dict_row 6.1 =>  dictLemmaTran[]=", dictLemmaTran[ len(dictLemmaTran)-1 ] ) 
			
			}
		} 
		//---------------------------
		/****
		//  build new dictLemmaTran  from uniqueWordByFreq
		for _, oneW:= range uniqueWordByFreq {	
			for m, lemmaFreq:= range oneW.uLemmaL {
				mLemm := strings.TrimSpace( lemmaLis[m] )
				mTran := strings.TrimSpace( tranLis[m] 	)	
					
				lemmaTranStr += "\n" + mLemm + "|" + mTran 	
				
				ele1.dL_lemmaSeq = seqCode(mLemm)                         //  lemmaTranStruct struct   .dL...
				ele1.dL_lemma2   = mLemm 
				ele1.dL_numDict  = lastNumDict
				ele1.dL_tran     = mTran          
				if mTran != "" { 		
					dictLemmaTran = append( dictLemmaTran, ele1 ) 
				}
			}	
		}
		***/
		
	} // end of z 
	
	return lemmaTranStr
	
} // end of split_ALL_word_dict_row(

//------------------------------------------


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
