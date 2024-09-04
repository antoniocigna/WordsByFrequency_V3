package wbfSubPack

	import (
		"fmt"
		"strings"
		"strconv"
		"sort"
		"regexp"
	)

//------------------------------

func buildWordList() {
    /*
	write a line in wordSliceFreq and wordSliceAlpha  for each word in the row 
	*/
	
	if  len(inputTextRowSlice) < 1 { return }  
	
	var wS1 wordStruct;
	//numMio:=0
	
	
	numberOfWords=0; 
	nn:=0
		
	lastPerc = 10;
	
	//----
	delta1 := (37.0- float64(lastPerc) ) / float64( len( inputTextRowSlice) )  
	percX1 := float64( lastPerc )  
	//------------------------------------------
	
	//sw_list_Word_if_in_ExtrRow = (last_mainpage_val_sel_extRow == "extrRow")   
	sw_list_Word_if_in_ExtrRow = (last_sel_extrRow == "extrRow")  
	
	//fromN := last_mainpage_val_inpBegRow
	//toN   := last_mainpage_val_inpBegRow + last_mainpage_val_maxNumRow -1 
	
	fromN := last_ixRowBeg 
	toN   := last_ixRowEnd      
	
	thisR_SelRow := SEL_NO_EXTR_ROW
	
	fmt.Println("buildWordList() fromN=", fromN,  "  toN=", toN , " len=" ,len(inputTextRowSlice), " sw_list_Word_if_in_ExtrRow =", sw_list_Word_if_in_ExtrRow )
	
	//antoCtr_rowSchrift :=0 ;
	//antoCtr_wordSchrift:=0; 
	//---------------------
	for ixR, rS2 := range inputTextRowSlice {	//  for each text row 
		
		//fmt.Println( " cerca parole loop ixR=", ixR, "  ", rS2);
		
		row2   := rS2.rRow1;	
		if row2 == LAST_WORD { continue }
		if sw_HTML_ready {
			percX1 += delta1 
			if ixR == (1000 * int(ixR/1000)) {
				//fmt.Println("ixR=", ixR, " percX1=", int( percX1 ) )
				go_exec_js_function( "showProgress", strconv.Itoa( int( percX1 ) ) ) 	
			}
		}		
		wordA  := regexp.MustCompile(separWord).Split(row2, -1);  // split row into words 
		
		tot1:= len(wordA) 
		
		rS2.rNumWords  = tot1      // number of words in the row 
		rS2.rListIxUnF = make( []int, tot1, tot1 )	
		
		rS2.rListFreq  = make( []int, tot1, tot1 )	
		
		inputTextRowSlice[ixR] = rS2
	
        z:= -1;
		thisR_SelRow = SEL_NO_EXTR_ROW
		if (sw_list_Word_if_in_ExtrRow) {
				if ((ixR >= fromN) && (ixR <= toN)) {
					thisR_SelRow = SEL_EXTR_ROW
				} 
			}	
			
		//swTEST := ((ixR == 4130) || (ixR == 4192) || (ixR == 4196))  // 
		
		//if ixR==211 { fmt.Println( " cerca parole loop ixR=", ixR, "  ", row2, "\n\t", wordA,  "\n\t", inputTextRowSlice[ixR] )}
		
		//if swTEST {  fmt.Println( " row2=", row2); antoCtr_rowSchrift++; } 
		
		for _, wor1 := range wordA {
			wS1.wWord2 = checkTheWord( wor1 ) ;
			if wS1.wWord2 == "" { continue }					
			z++;
			nn++				
			wS1.wNfile    = rS2.rNfile1 
			wS1.wSwSelRowR= thisR_SelRow			
			wS1.wIxRow    = ixR   // index of row containing the word 
			wS1.wIxPosRow = z;    // position of the word in the row 
			wordSliceAlpha = append(wordSliceAlpha, wS1);	
			/**
			if swTEST {  
				fmt.Println( "\t ANTO ", "word2=", wS1.wWord2, " wIxRow =ixR=", strconv.Itoa( ixR), " wordSliceAlpha=", wS1) 
				antoCtr_wordSchrift++
			} else {
				if wS1.wWord2 == "schrift" {
					 fmt.Println( "ANTO2 row2=", row2, "\n\t ANTO2 ", "word2=", wS1.wWord2, " wIxRow =ixR=", strconv.Itoa( ixR), " wordSliceAlpha=", wS1)
					 antoCtr_rowSchrift++
					 antoCtr_wordSchrift++
				} 
			}
			***/
		}
	} // end of for_ixR 
	
	//fmt.Println("anto3 antoCtr_rowSchrift=", antoCtr_rowSchrift, " antoCtr_wordSchrift=", antoCtr_wordSchrift)
	
	//---------------------------------------
	numberOfWords = len(wordSliceAlpha); 

	fmt.Println("numberOfWords=", numberOfWords)
			
	fmt.Println("number of words in text lines ", numberOfWords);
	
	//----	
	sort.Slice(wordSliceAlpha, func(i, j int) bool {
		return wordSliceAlpha[i].wWord2 < wordSliceAlpha[j].wWord2            // word  ascending order (eg.   a before b ) 		
	})
	//------------------------------
	addCodedWordToWordSlice()
	//---------------------------------
	
	// now wordSliceAlpha is in order by coded word and actual word ( eg. both actual word "über"   and "ueber" have "uber" as coded word) 
		
	/**
	for g:=0; g < len( wordSliceAlpha ); g++ {
		fmt.Println( "ANTONIO2 alpha ", wordSliceAlpha[g] )
	}	
	fmt.Println( "ANTONIO2 alpha \n")
	**/
	
	addTotRowToWord()
	
	
	
	
} // end of buildWordList
//-----------------

//----------------------------------------

func checkTheWord( word0 string ) string {
	// check word 
	//  return space if not valid
	//  return the same word if OK, sometime with the first character removed 
	//---
	var wor = strings.ToLower(  strings.TrimSpace( word0 ) )
	if ((wor == "") || ( wor == "&amp" )) { 
		return ""
	}
	if wor[0:1] < " "  {
		return ""
	}
	//--------------
	var j1 = strings.IndexAny(wor, "0123456789%|\\_*•-=^&~.,;?!\"'")
	if j1 >=0 {return ""}
	//------------
	var toRemove = "°¿¡€$£"
	var j2 = strings.IndexAny( wor, toRemove)
	if j2 >= 0 {
		// la parola "wor" contiene un carattere da rimuovere
		// se il primo carattere è tra quelli da rimuovere elimino il primo carattere e prendo il resto
		// se non è il primo, allora elimino tutta la parola
		//
		// voglio controllare soltanto il primo carattere,  
		// 		ma un carattere potrebbe essere più lungo di un byte, non posso usare substring  
		
		var wor2 = ""
		for i, letterR := range wor {
			var letter = string(letterR)
			//fmt.Println( "loop rune ", wor,  "  i=", i, " letter=", letter) 
			if (i == 0) { // test primo carattere 
				if strings.IndexAny( string(letter), toRemove) < 0 {
					return ""   // il primo è ok, ma gli altri no 
				} 
				continue
			}
			wor2 += letter 
			//fmt.Println( "loop rune ", wor,  "  i=", i, " letter=", letter,  " wor2=", wor2) 
		}
		wor = wor2
		if strings.IndexAny( wor, toRemove) >= 0 {
			// il carattere strano continua ad esserci, quindi elimino la parola 
			//fmt.Println(" ex loop ",  wor, "    ", "il carattere strano continua ad esserci, quindi elimino la parola " )	
			return ""    
		} 
	} 
	return stdCode(wor)
}
//---------------------------------------

func addCodedWordToWordSlice() {
	/*
	add a sortKeyWord to each word element
	*/
	preW     := "" 	
	preCoded := ""
	//----------------
	for i, wS1 := range wordSliceAlpha {
		if (wS1.wWord2 != preW) { 
			preW = wS1.wWord2	
			preCoded = newCode(preW)
		}
		wordSliceAlpha[i].wWordCod = preCoded;
	}	
	//----	
	sort.Slice(wordSliceAlpha, func(i, j int) bool {
		if wordSliceAlpha[i].wWordCod != wordSliceAlpha[j].wWordCod {
			return wordSliceAlpha[i].wWordCod < wordSliceAlpha[j].wWordCod            // word  ascending order (eg.   a before b ) 
		} else {
			if wordSliceAlpha[i].wWord2 != wordSliceAlpha[j].wWord2 {
				return wordSliceAlpha[i].wWord2 < wordSliceAlpha[j].wWord2  
			} else {
				return wordSliceAlpha[i].wNfile < wordSliceAlpha[j].wNfile          // nFile ascending order (eg.   0 before 1 ) 
			}
		}
	})
	//------------------------------
	
	
} // end of addCodedWordToWordSlice

//------------------------------------------------

func addTotRowToWord() {
	/*
	each element of wordSliceAlpha contains a word (the same word may be in several rows) 
	the number of repetition of a word (totRow) is put in its element  ( later will be put in each row that contain it) 
		eg.  one 3, one 3, one 3, two 4, two 4, two 4, two 4	
	*/
	preW  := wordSliceAlpha[0].wWordCod;	
	totR  := 0	
	ix1   := 0
	ix2   :=-1
	pre_wSwSelRow := SEL_NO_EXTR_ROW 
	//----------------
	tot_extrRow:=0
	for i, wS1 := range wordSliceAlpha {
		
		//fmt.Println("ANTO addTot ... i=" , i , " => ", wS1 )
		
		if (wS1.wWordCod != preW) {
			ix2 = i; 
			for i2 := ix1; i2 < ix2;i2++ {
				 wordSliceAlpha[i2].wSwSelRowG = pre_wSwSelRow; 	// se esiste almeno un richiamo a una riga estratta ( wSwSelRowR)allora questo segnale è ripetuto come wSwSelRowG
				 wordSliceAlpha[i2].wTotExtrRow = tot_extrRow   
				 wordSliceAlpha[i2].wTotRow    = totR;   // se una parola è ripetuta 3 volte, ad ogni parola è associato 3  
				 //if wordSliceAlpha[i2].wWord2=="von" { if  tot_extrRow>0 { fmt.Println("addTotRowToWord () 1 extr ", wordSliceAlpha[i2].wWord2, " .uTotExtrRow =", tot_extrRow) }  }
				
				//if wordSliceAlpha[i2].wWord2 == "schrift" { fmt.Println("ANTO2 addTotRowToWord  ", wordSliceAlpha[i2], 
				//		" .wSwSelRowR=", wordSliceAlpha[i2].wSwSelRowR, "  tot_extrRow=", wordSliceAlpha[i2].wTotExtrRow)  }  
	
			}			
			pre_wSwSelRow = SEL_NO_EXTR_ROW 
			totR = 0
			tot_extrRow = 0
			ix1  = i; 
			preW = wS1.wWordCod; 
		} 
		
		if wS1.wSwSelRowR == SEL_EXTR_ROW {   // se almeno uno è "estratto", tutti lo sono 
			pre_wSwSelRow = SEL_EXTR_ROW 
			tot_extrRow++	
			//if (wS1.wWord2 == "schrift") { fmt.Println("ANTO addTotRowToWord  ",  wS1 , " SEL_EXTR_ROW=", SEL_EXTR_ROW, "  tot_extrRow=", tot_extrRow) }  //
		} 
		totR++;     	
	}
	ix2++; 
	for i2 := ix1; i2 < len(wordSliceAlpha);i2++ {
		wordSliceAlpha[i2].wTotRow   = totR; 
		wordSliceAlpha[i2].wTotExtrRow = tot_extrRow   
		wordSliceAlpha[i2].wSwSelRowG = pre_wSwSelRow;
		// if wordSliceAlpha[i2].wWord2=="von" { if  tot_extrRow>0 { fmt.Println("addTotRowToWord () 2 extr ", wordSliceAlpha[i2].wWord2, " .uTotExtrRow =", tot_extrRow) }  }
	
	}
	//------		
	/**
	for g:=0; g < len( wordSliceAlpha ); g++ {
		fmt.Println( "ANTONIO3 alpha ", wordSliceAlpha[g] )
	}	
	**/
	
} // end of addTotRowToWord 

//---------------------------------
