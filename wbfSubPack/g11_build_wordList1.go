package wbfSubPack

	import (
		"fmt"
		"strings"
		"strconv"
		"sort"
		"regexp"
		//"encoding/hex"
	)

//------------------------------
func g11_buildWordList() {
    /*
	write a line in wordSliceFreq and wordSliceAlpha  for each word in the row 
	*/
	fmt.Println("func ", green("buildWordList") )
	
	
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
	
	//fmt.Println("buildWordList() fromN=", fromN,  "  toN=", toN , " len=" ,len(inputTextRowSlice), " sw_list_Word_if_in_ExtrRow =", sw_list_Word_if_in_ExtrRow )
	/***
	???
	separPrefList = make([]separPrefStruct, 0, 200) 
	var sP separPrefStruct 
	sP.sLenPref  int 
	.sPrefix   string 
	.sPrefTran string 	

	
	
	***/
	
	//-------
	
	fmt.Println(" stdCode: from ",  translate_chars_std_inpList, " to ", translate_chars_std_outList)
	fmt.Println(" seqCode: from ",  translate_chars_SEQ_inpList, " to ", translate_chars_SEQ_outList)
	//------------
	//antoCtr_rowSchrift :=0 ;
	//antoCtr_wordSchrift:=0; 
	//fmt.Println("separPrefList=" , separPrefList)
	allPrefStringList := " " 
	for _,sP := range separPrefList {
		allPrefStringList += sP.sPrefix + " "
	}
	
	fmt.Println("allPrefStringList=" + allPrefStringList)
	
	all_words = make([]string,0, 5*len(inputTextRowSlice) ) 
	
	//---------------------
	for ixR, rS2 := range inputTextRowSlice {	//  for each text row 
		//if (ixR < 25) { fmt.Println( green("inputTextRowSlice "), rS2.rRow1) }
		//fmt.Println( " cerca parole loop ixR=", ixR, "  ", rS2);
		
		row2   := rS2.rRow1;	
		if sw_HTML_ready {
			percX1 += delta1 
			if ixR == (1000 * int(ixR/1000)) {
				//fmt.Println("ixR=", ixR, " percX1=", int( percX1 ) )
				go_exec_js_function( "showProgress", strconv.Itoa( int( percX1 ) ) ) 	
			}
		}		
		//wordA  := regexp.MustCompile(separWord).Split(row2, -1);  // split row into words 
		wordA  := strings.Fields( regexp.MustCompile(separWord).ReplaceAllString(row2," ") ) 
		
		//if (ixR < 5) { fmt.Println( "  wordA=", wordA) }
		tot1:= len(wordA) 
		
		all_words = append(all_words, wordA...)
				
        z:= -1;
		thisR_SelRow = SEL_NO_EXTR_ROW
		if (sw_list_Word_if_in_ExtrRow) {
				if ((ixR >= fromN) && (ixR <= toN)) {
					thisR_SelRow = SEL_EXTR_ROW
				} 
		}							
		//------------------
		
		pref_inThisLine := "  " 
		
		//-------------------
		for _, wor1 := range wordA {			
			oneW0 := checkTheWord( wor1 ) ;
			if oneW0 == "" { continue}
			oneW1:= " "+oneW0+" "	
			if strings.Index(allPrefStringList, oneW1) < 0 {continue}
			if strings.Index(pref_inThisLine, oneW1) < 0 {
				pref_inThisLine += oneW1
			}
		}	
		//pref_inThisLineList := strings.Fields(pref_inThisLine)
		//if (ixR < 25) { fmt.Println( " pref_inThisLine=", pref_inThisLine) }
		//-------------------
		for _, wor1 := range wordA {
			//if nn < 20 { fmt.Println( "buildWordList ", nn, "  ", wor1)}
			
			wS1.wWord2 = checkTheWord( wor1 ) 
			if wS1.wWord2 == "" { continue }	
			wS1.wWordSeq = seqCode(wS1.wWord2)
			//------------------
			if len(pref_inThisLine) == 0 {
				wS1.wListPref = ""
			} else {				
				if strings.Index(pref_inThisLine, " " + wS1.wWord2 + " ") < 0 {
					wS1.wListPref = strings.TrimSpace(pref_inThisLine)
				} else {
					wS1.wListPref = ""
				}	
			}
			//if (ixR < 25) { fmt.Println( " wor1=", wor1 , ", wS1.wListPref=" , wS1.wListPref) }
			//----------------	
			z++;
			nn++				
			wS1.wNfile    = rS2.rNfile1 
			wS1.wSwSelRowR= thisR_SelRow			
			wS1.wIxRow    = ixR   // index of row containing the word 
			wS1.wIxPosRow = z;    // position of the word in the row 
			wordSliceAlpha = append(wordSliceAlpha, wS1);	
			
		}
		tot1 = 1+z 
		rS2.rNumWords  = tot1      // number of words in the row 
		rS2.rListIxUnF = make( []int, tot1, tot1 )			
		rS2.rListFreq  = make( []int, tot1, tot1 )	
		
		inputTextRowSlice[ixR] = rS2
		
	} // end of for_ixR 
	
	//---------------------------------------
	numberOfWords = len(wordSliceAlpha); 

	fmt.Println("numberOfWords=", numberOfWords)
			
	fmt.Println("number of words in text lines ", numberOfWords);
	//-------------------------	
	var wS0 wordStruct;
	wS0.wWord2   = LAST_WORD
	wS0.wWordSeq = LAST_WORD  
	wordSliceAlpha = append(wordSliceAlpha, wS0);
	all_words      = append(all_words,wS0.wWord2)
	//-----------------
	sort.Strings(all_words)
	
	//----	
	fmt.Println("sort wordSliceAlpha  in ordine .wWordSeq, .wWord2, .wNfile") 
	
	/***
	sort.Slice(wordSliceAlpha, func(i, j int) bool {
		return wordSliceAlpha[i].wWord2 < wordSliceAlpha[j].wWord2            // word  ascending order (eg.   a before b ) 		
	})
	//----------------------
	***/
	//----	
	sort.Slice(wordSliceAlpha, func(i, j int) bool {
		if wordSliceAlpha[i].wWordSeq != wordSliceAlpha[j].wWordSeq {
			return wordSliceAlpha[i].wWordSeq < wordSliceAlpha[j].wWordSeq            // word  ascending order (eg.   a before b ) 
		} else {
			if wordSliceAlpha[i].wWord2 != wordSliceAlpha[j].wWord2 {
				return wordSliceAlpha[i].wWord2 < wordSliceAlpha[j].wWord2  
			} else {
				return wordSliceAlpha[i].wNfile < wordSliceAlpha[j].wNfile          // nFile ascending order (eg.   0 before 1 ) 
			}
		}
	})
	//------------------------------		
	
	g12_add_totRow_and_indexLemmaPair()
	
	
	// UNIQUE FREQUENZA  -------------
	
	uniqueWordByFreq  = make([]wordUnFreqStruct, 0, len(uniqueWordByAlpha) ) 	 // la slice destinazione del 'copy' deve avere la stessa lunghezza di quella input  
	
	var ffww wordUnFreqStruct
	
	for _, aW:= range uniqueWordByAlpha {
		ffww.fuWordSeq  = aW.uWordSeq
		ffww.fuWord2    = aW.uWord2
		ffww.fuIxUnW_al = aW.uIxUnW_al
		ffww.fuIxUnW_fr = aW.uIxUnW_fr
		ffww.fuTotRow   = aW.uTotRow
		uniqueWordByFreq = append(uniqueWordByFreq , ffww)
	} 	
	// le parole eguali si trovano in righe contigue perchè hanno la stessa frequenza	
	sort.Slice(uniqueWordByFreq, func(i, j int) bool {
			if uniqueWordByFreq[i].fuTotRow !=  uniqueWordByFreq[j].fuTotRow {
			   return uniqueWordByFreq[i].fuTotRow > uniqueWordByFreq[j].fuTotRow        // totRow    descending order (how many row contain the word) 
			} else {
				if uniqueWordByFreq[i].fuWordSeq != uniqueWordByFreq[j].fuWordSeq {
					return uniqueWordByFreq[i].fuWordSeq < uniqueWordByFreq[j].fuWordSeq            // word  ascending order (eg.   a before b ) 
				} else {
					return uniqueWordByFreq[i].fuWord2 < uniqueWordByFreq[j].fuWord2  			
				}
			}
		})
	//---------------------------------------------------	
	for n2, aa:= range uniqueWordByFreq {
		ixUA := aa.fuIxUnW_al
		uniqueWordByFreq[ n2  ].fuIxUnW_fr = n2 
		uniqueWordByAlpha[ixUA].uIxUnW_fr  = n2 
		ixFromW := uniqueWordByAlpha[ixUA].uIxFromWord_al
		ixToW   := uniqueWordByAlpha[ixUA].uTotRow + ixFromW 
		for x1:= ixFromW; x1 < ixToW; x1++ {
			wordSliceAlpha[x1].wIxUniq_fr = n2
		}	  
	}	
	numberOfUniqueWords = len(uniqueWordByFreq)
	//------------------------------------
	
	loadInverseWordSlice()
	
	putWordFrequenceInRowArray1()
	//putWordFrequenceInRowArray2()
	
	read_wordsToLearn()
	
	
} // end of buildWordList


//----------------------------------------


func putWordFrequenceInRowArray1() {

	ix:=0;
	
	//---------------------------------
	//  fill each row with the frequence of its words
	//-------------------	
	
	for _, wS1 := range wordSliceAlpha {			
		ix = wS1.wIxRow;            // numero riga
		ixPos := wS1.wIxPosRow; 	// numero di parola nella riga	
		
		if ( strings.Index( wS1.wWord2, "...") >= 0 ) {		
			inputTextRowSlice[ix].rListIxUnF = append( inputTextRowSlice[ix].rListIxUnF, wS1.wIxUniq_fr )	
			inputTextRowSlice[ix].rListFreq  = append( inputTextRowSlice[ix].rListFreq,  wS1.wTotRow    )	
			inputTextRowSlice[ix].rNumWords  = len( inputTextRowSlice[ix].rListIxUnF )
			continue		
		}
		if ixPos < inputTextRowSlice[ix].rNumWords { 
			inputTextRowSlice[ix].rListIxUnF[ixPos] = wS1.wIxUniq_fr // index of the word in the uniqueWordByFreq  	
			inputTextRowSlice[ix].rListFreq[ ixPos] = wS1.wTotRow    // for each word in the row  set its frequence of use (how many times the word is used in the whole text)  
		} else {
			fmt.Println("errore in func ", red( "putWordFrequenceInRowArray"), " row n.", ix, " word pos=", ixPos, " num words in row=",  
				inputTextRowSlice[ix].rNumWords, " word=", wS1.wWord2, " row=", inputTextRowSlice[ix].rRow1)
		}	
	}
	
	//---------------------------
} // end of putWordFrequenceInRowArray1
//--------------------------------------------------