package wbfSubPack

	import (
		"fmt"
		"strconv"
	    //"strings"
		"sort"
		"encoding/hex"
		"os"
	)

//------------------------------

func elabWordList() {	


	fmt.Println("\n", green("elabWordList"), "\n")
	
	elabWordAlpha_buildWordFreqList() 		
		
	build_uniqueWord_byFreqAlpha(); 
	
	loadInverseWordSlice()
	
	putWordFrequenceInRowArray()

	//addRowTranslation() 
	
	build_lemma_word_ix()
	
	//antoList_wordSchrift_anto()
	
	fmt.Println("\n", green("end elabWordList"), "\n")
	
} // end of elabWordList()

//-----------------

//---------------------------------

func elabWordAlpha_buildWordFreqList() {
	/*
	put in each row of the inputTextRowSlice the number of its words  
	from wordSliceAlpha list obtain a new list by sorting it by occurrence of the words (totRow) 
	*/
	//preW  := ""; // wordSliceAlpha[0].word;	
	//ix:=0;
	//removeIxWord :=-1 
		/******************** rimosso il 15/11/2023
	
	for nn, wS1 := range wordSliceAlpha {	
		if wS1.wWordSeq == LAST_WORD  { wordSliceAlpha[nn].totRow = LAST_WORD_FREQ }
		//fmt.Println("ANTO elab...FreqList() ", nn, "  alpha=",  wordSliceAlpha[nn]);  	
		//fmt.Println("ANTO elab...FreqList() nn=", nn, ", preW=" + preW + ",  wS1 ", wS1)  
		
		if (wS1.wWordSeq != preW) {		
			removeIxWord = -1
			preW = wS1.wWord2; 
		}
		ix =  wS1.ixRow
		inputTextRowSlice[ix].numWords ++; 		// how many words contains the row (eg. the row "the cat is on the table"  contains 6 words --> .numWords = 6 
		if removeIxWord >= 0 {
			//fmt.Println("\t ANTO 2elab" , "  removeIxWord=" , removeIxWord , " wordSliceAlpha[ removeIxWord ].word=" + wordSliceAlpha[ removeIxWord ].word  )
			
		    if wS1.wWordSeq == wordSliceAlpha[ removeIxWord ].wWordSeq {
				wordSliceAlpha[ nn ].sw_ignore = true 
			}	
		}	
		//fmt.Println("\t ANTO XXXelab"  , "  wS1.sw_ignore = ", wS1.sw_ignore)  
	}
	*****************************/
	
	//----------------
	for nn, wS1 := range wordSliceAlpha {	
		if wS1.wWordSeq == LAST_WORD  { wordSliceAlpha[nn].wTotRow = LAST_WORD_FREQ }	
		inputTextRowSlice[ wS1.wIxRow ].rNumWords ++; 		// number of words in a row (eg. the row "the cat is on the table"  contains 6 words --> .numWords = 6 
	}
	//------------
	// build WordList by frequence in the text
	// the slice is sorted in descending order of frequency   ( ie. firstly the most used)   
	//-----------------------------------
	
	wordSliceFreq  = make([]wordStruct, len(wordSliceAlpha),  len(wordSliceAlpha) ) 	 // la slice destinazione del 'copy' deve avere la stessa lunghezza di quella input  
	
	copy(wordSliceFreq , wordSliceAlpha);
	
	
	// le parole eguali si trovano in righe contigue perchè hanno la stessa frequenza
	
	sort.Slice(wordSliceFreq, func(i, j int) bool {
			if wordSliceFreq[i].wTotRow !=  wordSliceFreq[j].wTotRow {
			   return wordSliceFreq[i].wTotRow > wordSliceFreq[j].wTotRow        // totRow    descending order (how many row contain the word) 
			} else {
				if wordSliceFreq[i].wWordSeq != wordSliceFreq[j].wWordSeq {
					return wordSliceFreq[i].wWordSeq < wordSliceFreq[j].wWordSeq            // word  ascending order (eg.   a before b ) 
				} else {
					return wordSliceFreq[i].wWord2 < wordSliceFreq[j].wWord2  			
				}
			}
		})
		
} // end of elabWordAlpha_buildWordFreqList

//--------------------------------------

func putWordFrequenceInRowArray() {

	ix:=0;

	//-------------------------	
	/***
	//  in each element of inputTextRowSlice define an empty slice to contain the frequence of each of its words
	for k, _ := range inputTextRowSlice {	
		tot1 := inputTextRowSlice[k].rNumWords;
		inputTextRowSlice[k].rListIxUnF =  make( []int, tot1, tot1 )		
		inputTextRowSlice[k].rListFreq  =  make( []int, tot1, tot1 )			
	}
	***/
	//---------------------------------
	//  fill each row with the frequence of its words
	//-------------------	
	
	for _, wS1 := range wordSliceFreq {	
		//fmt.Println("ANTO putWordFrequenceInRowArray() wS1 ", wS1)  
		
		ix = wS1.wIxRow; 
		ixPos := wS1.wIxPosRow; 
		/****
		num2 := len(inputTextRowSlice[ix].rListFreq) 
		if (num2 <= ixPos) {		
			fmt.Println("error " , wS1.wWord2, " ix=" , ix  ," ixPos=",   ixPos, " row ", " num2=", num2, " tot1=",  
				inputTextRowSlice[ix].rNumWords, " list=" , inputTextRowSlice[ix].rListFreq , " " , inputTextRowSlice[ix].rRow1); 			
		}
		if ((ixPos<1) || (ix<1)) {
			fmt.Println( "ERRORe  putWordFrequenceInRowArray() nx=", nx, " WS1=", wS1, "\n\t", "ix=", ix, ", ixPos=", ixPos)  
		}
		***/
		if ixPos < inputTextRowSlice[ix].rNumWords { 
			inputTextRowSlice[ix].rListIxUnF[ixPos] = wS1.wIxUniq // index of the word in the uniqueWordByFreq  	
			inputTextRowSlice[ix].rListFreq[ ixPos] = wS1.wTotRow // for each word in the row  set its frequence of use (how many times the word is used in the whole text)  
		} else {
			fmt.Println("errore in func ", red( "putWordFrequenceInRowArray"), " row n.", ix, " word pos=", ixPos, " num words in row=",  
				inputTextRowSlice[ix].rNumWords, " word=", wS1.wWord2, " row=", inputTextRowSlice[ix].rRow1)
		}	
	}
	
	//---------------------------
} // end of putWordFrequenceInRowArray

//------------------------------------------------

func put_a_priority_to_the_row_of_each_word() {
	
	//fmt.Println("put_a_priority_to_the_row_of_each_word()")
	
	//the row importance is assigned by the number of its unknown words	
		
	for k, wS1 := range wordSliceFreq {	
		
		ix := wS1.wIxRow; 
		numMinor :=0 	
		wordFreq := wS1.wTotRow; 		
		for _, frw:= range   inputTextRowSlice[ix].rListFreq {
			if frw < wordFreq { numMinor++; }
		} 
		wordSliceFreq[k].wTotMinRow = numMinor;	             // number of words with frequency < this word   	
		wordSliceFreq[k].wTotWrdRow = inputTextRowSlice[ix].rNumWords   // number of words in this row    
	}		
	
	sortWordListByFreq_and_row_priority() 
		
} // end of put_a_priority

//--------------------------------------
func build_uniqueWord_byFreqAlpha() {
	
		
	put_a_priority_to_the_row_of_each_word() 	
	
	
	
	preW := ""
	numWordUn := 0
	numWordRi := 0	
	num_word:=0
	
	numWordUn_0 := 0
	numWordRi_0 := 0	
	num_word_0 :=0 
	//--------------------
		
	for _, wS1 := range wordSliceFreq {	
		num_word++
		//if wS1.sw_ignore == false { 
		num_word_0++
		//}
		
		if wS1.wWordSeq != preW {
			preW = wS1.wWordSeq;
			numWordUn += 1 
			numWordRi += wS1.wTotRow 
			//if wS1.sw_ignore == false { 
			numWordUn_0 += 1 
			numWordRi_0 += wS1.wTotRow 
			//}
		}  
	}
		//------------
	if num_word_0 != num_word {
		fmt.Println( "PAROLE SINGOLE File0= ", numWordUn_0, ", PAROLE Totale=", numWordRi_0,  "  numberOfWords=" , num_word_0 , " "  );
	}
	fmt.Println( "PAROLE SINGOLE tutti= ", numWordUn, ", PAROLE Totale=", numWordRi,  "  numberOfWords=" , num_word , "\n");
		//--
	//numberOfUniqueWords = numWordUn;
	numberOfUniqueWords = numWordUn_0;
	preW      = ""
	numWordUn = 0
	numWordRi = 0	
	
	percIx := 0; 	

	//result_word2 ="";
	
	var xWordF wordIxStruct;   
	var sS  statStruct;
	
	//numWordUn = -1
	numWordUn = 0
	
	//-----------------------------
	/**** removed 2023_11_15
	// remove elements to ignore  ( those of the files after the first )
	wrk := wordSliceFreq[:0]
	for _, wS1 := range wordSliceFreq {				
			if wS1.sw_ignore { continue }
			wrk = append( wrk, wS1) 
	}
	wordSliceFreq = wrk		
	**/
	//--------------------
	/**
	for n0, wS1 := range wordSliceFreq {	
		if n0 > 20 {fmt.Println("ANTONIO1 build_uniqueWord_byFreqAlpha() ", "... continua");  break;  }
		fmt.Println("ANTONIO1 build_uniqueWord_byFreqAlpha() ", wS1) 
	}
	***/
	//---------
	
	//fmt.Println( "build_uniqueWord_byFreqAlpha() loop 1 ");
	
	for n1, wS1 := range wordSliceFreq {	
		if wS1.wWordSeq != preW {
			preW = wS1.wWordSeq;
			
			if wS1.wTotRow >= LAST_WORD_FREQ {
				wS1.wTotRow = 0
			}
			xWordF.uWordSeq  = wS1.wWordSeq;
			xWordF.uWord2    = wS1.wWord2;
			xWordF.uTotRow   = wS1.wTotRow
			xWordF.uTotExtrRow = wS1.wTotExtrRow  
			xWordF.uSwSelRowR = wS1.wSwSelRowR 
			xWordF.uSwSelRowG = wS1.wSwSelRowG
			xWordF.uIxWordFreq = n1
			
			//if wS1.wWord2=="von" { if  wS1.wTotExtrRow>0 { fmt.Println("build_uniqueWord_byFreqAlpha() extr ", wS1.wWord2, " xWordF.uTotExtrRow =", xWordF.uTotExtrRow) }  }
			
			
			//xWordF.wTran = "" 
			xWordF.uIxUnW     = len(uniqueWordByFreq)  
			uniqueWordByFreq = append( uniqueWordByFreq, xWordF);  
			numWordUn += 1 
						
			numWordRi += wS1.wTotRow 
			percIx = int(numWordUn * 100 / numberOfUniqueWords); 
			
			//fmt.Println( "   n1=", n1, " \t ", xWordF.uWord2, " \t xWordF.uSwSelRow =" , xWordF.uSwSelRow)
			
			sS.uniqueWords = numWordUn 
			sS.totWords    = numWordRi
			sS.uniquePerc  = percIx 
			sS.totPerc     = int(numWordRi * 100 / numberOfWords);
			
			//if sS.totPerc > 100 {  fmt.Println("AN TONIO4 n1=", n1, " len(wordSliceFreq)=",  len(wordSliceFreq) , " wS1.wWord2=" + wS1.wWord2 + " wS1.totRow=", wS1.totRow, " numWordRi=", numWordRi ) }
			
			/**
			if strconv.Itoa(1000 + percIx)[3:] == "0" {  
				wordStatistic_un[percIx] = sS; 	
			}
			***/				
			if sS.totPerc <= 200 {   // esistono perc > 100%,  probabilmente c'è un errore di logica 
				if strconv.Itoa(1000 + sS.totPerc)[3:] == "0" {				
					wordStatistic_tx[sS.totPerc] = sS; 
				}
			}
			//fmt.Println("STAT. ", n1, " ", xWordF.word, " numWordUn=", numWordUn,  " numWordRi=", numWordRi, " percIx=", percIx, " ", sS.uniquePerc,  " sS.totPerc=" ,  sS.totPerc); 
		} 			
	}
	//---------
	
	highestValueByte, err := hex.DecodeString("ffff")   
	if err != nil { panic(err) }
	var highestValue = string( highestValueByte ) + "end_of_list"	
	xWordF.uWordSeq = highestValue 
	xWordF.uWord2   = highestValue 	
	
	xWordF.uTotRow = 1 ; // the lowest frequency
	xWordF.uTotExtrRow = 0               
	xWordF.uIxWordFreq = len(uniqueWordByFreq)   
	xWordF.uIxUnW      = len(uniqueWordByFreq)  	
	xWordF.uKnow_yes_ctr = 0 
	xWordF.uKnow_no_ctr  = 0 	
	//xWordF.uTranL      = []string{ xWordF.uWord2 }        // ??anto8 .uTranL
	uniqueWordByFreq   = append( uniqueWordByFreq, xWordF);  
	
	//--------------------------
	
	addWordLemmaTranLevelParadigma()   
		 
	add_ixWord_to_WordSliceFreq()
	
	
	//addWordTranslation()		
	
	//---------------------
	uniqueWordByAlpha = make([]wordIxStruct, len(uniqueWordByFreq),  len(uniqueWordByFreq))	 // la slice destinazione del 'copy' deve avere la stessa lunghezza di quella input  
	
	//stat_useWord();
	
	copy( uniqueWordByAlpha, uniqueWordByFreq); 
	
	
	//fmt.Println("build_uniqueWord_byFreqAlpha() PROVA len unique Freq() = ", len(uniqueWordByFreq)   )
	
	//------------
	sort.Slice(uniqueWordByAlpha, func(i, j int) bool {
		if uniqueWordByAlpha[i].uWordSeq != uniqueWordByAlpha[j].uWordSeq {
			return uniqueWordByAlpha[i].uWordSeq < uniqueWordByAlpha[j].uWordSeq            // word  ascending order (eg.   a before b ) 
		} else {
			return uniqueWordByAlpha[i].uWord2 < uniqueWordByAlpha[j].uWord2  			
		}
	})
	//---------

	//console( "\nlista uniqueWordByAlpha")	
	// update alpha index  // ixUnW,  ixUnW_al	
	
	//fmt.Println("\n build_uniqueWord_byFreqAlpha() PROVA len unique alpha() = ", len(uniqueWordByAlpha)   )
	
	for u:=0; u < len(uniqueWordByAlpha); u++ {
		f:= uniqueWordByAlpha[u].uIxUnW
		uniqueWordByFreq[f].uIxUnW_al  = u; 		
		uniqueWordByAlpha[u].uIxUnW_al = u
				
	}
 	//console( "------------------\n")
	
	//fmt.Println("\n---------------- build_uniqueWord_byFreqAlpha() end  \n")
	

} // end of build_uniqueWord_byFreqAlpha

//-----------------------

//-----------------------

func build_lemma_word_ix() {
	
	//	build a slice with all lemma with all words 
	//  lemma_word_ix  loaded in addWordLemmaTranLevelParadigma

	var LW lemmaWordStruct
	/***						
			type lemmaStruct struct {
				leLemma    string
				leNumWords int 
				leFromIxLW  int 
				leToIxLW    int  
				leTran     string  
			}
			//---------------
			type lemmaWordStruct struct {
				lw_lemmaSeq string 	
				lw_lemma2   string 	
				lw_word     string 
				lw_ixLemma    int
				lw_ixWordUnFr int
			}
			//------------
	***/	
	//--------------------
	
	// order by lemma and word 	
	sort.Slice(lemma_word_ix, func(i, j int) bool {
			if  lemma_word_ix[i].lw_lemmaSeq != lemma_word_ix[j].lw_lemmaSeq {
				return lemma_word_ix[i].lw_lemmaSeq < lemma_word_ix[j].lw_lemmaSeq 	
			} else {
				if  lemma_word_ix[i].lw_lemma2 != lemma_word_ix[j].lw_lemma2 {
					return lemma_word_ix[i].lw_lemma2 < lemma_word_ix[j].lw_lemma2	
				} else {
					return lemma_word_ix[i].lw_word < lemma_word_ix[j].lw_word 
				}				
			}
		} )		
	fromIxWL:=-1
	toIxWL  :=-1	
	preIxLem:=-1 
	
	
	for z2:=0; z2 < len( lemma_word_ix); z2++ { 
		LW = lemma_word_ix[z2];
		if LW.lw_ixLemma != preIxLem  {	
			if preIxLem >=0 {
				lemmaSlice[preIxLem].leFromIxLW = fromIxWL
				lemmaSlice[preIxLem].leToIxLW   = toIxWL
			}
			preIxLem = LW.lw_ixLemma
			fromIxWL = z2
		} 
		toIxWL = z2	
	}
	if preIxLem >=0 {
				lemmaSlice[preIxLem].leFromIxLW = fromIxWL
				lemmaSlice[preIxLem].leToIxLW   = toIxWL
			}
	//---------------
	/**
	for z2:=0; z2 < len( lemma_word_ix); z2++ { 
		if z2 < 20 { fmt.Println(   "lemma_word_ix[", z2, "] = ", lemma_word_ix[z2]  ) } else { break }
		
	}
	**/
	//---------------
	/**
	for z2:=0; z2 < len( lemmaSlice); z2++ { 
		LE := lemmaSlice[z2];
		if LE.leLemma == "werkzeug" {
			fmt.Println("lemma = ", LE )
			for z3:= LE.leFromIxLW; z3 <= LE.leToIxLW; z3++ {
				LW = lemma_word_ix[z3]
				fmt.Println(" \t word = ", LW) 
			} 
			break;
		}
	}
	***/
	
} // end of build_lemma_word_ix  
 
//------------------------------------------


//--------------------------

func sortWordListByFreq_and_row_priority() {
	
	//fmt.Println(" sortWordListByFreq_and_row_priority()")
	
	sort.Slice(wordSliceFreq, func(i, j int) bool {
	
		if wordSliceFreq[i].wTotExtrRow !=  wordSliceFreq[j].wTotExtrRow {
		   return wordSliceFreq[i].wTotExtrRow > wordSliceFreq[j].wTotExtrRow  // wTotExtrRow  descending order ( number of extracted rows  if id_sel_2_extrRow option = id="extrRow", else =0) 	
		}
		
		if wordSliceFreq[i].wTotRow !=  wordSliceFreq[j].wTotRow {
		   return wordSliceFreq[i].wTotRow > wordSliceFreq[j].wTotRow         // totRow    descending order (how many rows contain the word)  	  		   
		}	
		
		if wordSliceFreq[i].wWordSeq !=  wordSliceFreq[j].wWordSeq {
		   return wordSliceFreq[i].wWordSeq < wordSliceFreq[j].wWordSeq            // word      ascending order	  		   
		}	
		if wordSliceFreq[i].wWord2 !=  wordSliceFreq[j].wWord2 {
		   return wordSliceFreq[i].wWord2 < wordSliceFreq[j].wWord2                // word      ascending order	  		   
		}				
		if wordSliceFreq[i].wTotMinRow !=  wordSliceFreq[j].wTotMinRow {
		   return wordSliceFreq[i].wTotMinRow < wordSliceFreq[j].wTotMinRow  // totMinRow ascending order	(how many words in the row are not yet learned) 
		}
		if wordSliceFreq[i].wSwSelRowR !=  wordSliceFreq[j].wSwSelRowR {
		   return wordSliceFreq[i].wSwSelRowR < wordSliceFreq[j].wSwSelRowR    // first the extracted row  
		}
		return wordSliceFreq[i].wIxRow < wordSliceFreq[j].wIxRow             // ixRow     ascending order	( first the rows which were first in the text)  		   
			
	})	
	
	/**
		fmt.Println( "\nANTONIO5 wordSliceFreq -------------------------")
	for g:=0; g < len( wordSliceFreq ); g++ {
		fmt.Println( "ANTONIO5 alpha ", wordSliceFreq[g] )
	}	
	**/
	
} // end of sortWordListByFreq_and_row_priority

//-----------------------------

//===========================================================================
func addWordLemmaTranLevelParadigma() {
	
	//fmt.Println("addWordLemmaTranLevelParadigma()" )
	
	newWordLemmaPair = make( []wordLemmaPairStruct, 0, len(uniqueWordByFreq)  )		
	var newWL wordLemmaPairStruct   // 	lWordSeq string, lWord2 string , lLemma string
	var wP lemmaTranStruct  
		
	lemma_word_ix = make([]lemmaWordStruct, 0,  len(uniqueWordByFreq)  )  
	var LW lemmaWordStruct
	
	list1Level:= ""
	list1Para := ""
	list1Exam := ""	
	
	//var swMio bool = false
	
	//--------------------------------------
	for zz:=0; zz < len(uniqueWordByFreq); zz++ {
		wF:= uniqueWordByFreq[zz]
		
		
		//swprova:= ((wF.uWord2 == "cäsar") || (wF.uWord2 == "caesar") || (wF.uWord2 == "casar")) 
		
		
		
		ixLemmaPairFoundList := lookForAllLemmas( wF.uWord2 ) // 
		
		//if swprova { fmt.Println("1 loop unique x lemma ", wF.uWord2,  " ixLemmaPairFoundList=",  ixLemmaPairFoundList ) }

				/**
						//
						var lemmaSlice       [] lemmaStruct         // lemma , translation 
						var wordLemmaPair    [] wordLemmaPairStruct // all word-lemma pair  
						//---
						
						type lemmaStruct struct {
							leLemma    string        // ls_lemma_einstellen einstellen
							
							leNumWords int 
							leFromIxLW  int 
							leToIxLW    int  
							leTran      string 
							leLevel     string  
							lePara      string  
							leExample   string  
							ls_lemma_ix_stellen  int	
							ls_lemma_stellen     string
							ls_pref_ein          string
							ls_pref_tran         string 
							ls_lemma_einStellenList []int 
						} 
						//---------
						type lemmaWordStruct struct {
							lw_lemmaSeq string 	
							lw_lemma2   string 	
							lw_word     string 
							lw_ixLemma    int
							lw_ixWordUnFr int
						}
						//-------------------------------
						type wordLemmaPairStruct struct {
							lWordSeq string 
							lWord2   string 
							lLemma   string
							lIxLemma int
						} 
						//---
				**/ 


		lis_ixLemma, lis_origLemma:= build_listOfLemmaForAWord(wF.uWord2, ixLemmaPairFoundList)
 		
		nele := len( lis_ixLemma )
			
		lis_lemmaName := make( [] string, 0, nele )		
		lis_tran  := make( [] string, 0, nele )
		lis_level := make( [] string, 0, nele )
		lis_para  := make( [] string, 0, nele )
		lis_exam  := make( [] string, 0, nele )
		

		//ixLemma:=-1
		
		
		newWL.lWord2   = wF.uWord2
		newWL.lWordSeq = seqCode( wF.uWord2 )
		newWL.lLemma = ""
		newWL.lIxLemma = -1
		
		//sw1:= (strings.Index(wF.uWord2, "stellen")>=0 )

		//if sw1 {  fmt.Println("") }

		// each word can have many lemmas 
		//                       each lemma can have many levels, paradigmas, translations ( they are separated by "|", eg "A1|A2|B1" ) 
		//---------------
		for n1, ixLemma2:= range lis_ixLemma { 
			lem := lemmaSlice[ixLemma2].leLemma 
			lis_lemmaName   = append( lis_lemmaName,  lem )  			
			
			//if sw1 {  fmt.Println(" xxx word=",wF.uWord2,  sPrintOneLemma( ixLemma2,   lemmaSlice[ixLemma2] )) }
			
			lemmaSlice[ixLemma2].leNumWords++
			
			list1Level, list1Para, list1Exam = fromLemmaTo3List( lem )  // ogni list può contenere più elementi separati da |   ( lo stesso num.di elementi per tutte le liste: A1|A2,par1|par2, ex1|ex2 )  
			
			list1Exam = compound_lemma(lemmaSlice[ixLemma2] )
	
			lis_level = append( lis_level, list1Level )
			lis_para  = append( lis_para , list1Para  )
			lis_exam  = append( lis_exam , list1Exam  )
			ixTra := lookForAllTran( lem ) 
			if ixTra >= 0 { 
				wP = dictLemmaTran[ixTra] 
				lis_tran = append( lis_tran, wP.dL_tran ) 		////cigna1	
				if ixLemma2 >=0 {	lemmaSlice[ixLemma2].leTran = wP.dL_tran } 
												// lab_word_list.go   wP.dL_tran= eindhoven   lemmaSlice[ixLemma2].leTran = eindhoven
				//if lem == "eindhoven" { fmt.Println("elab_word_list.go ", " wP.dL_tran=",wP.dL_tran,"  lemmaSlice[ixLemma2].leTran =",  lemmaSlice[ixLemma2].leTran ) }
			} else {
				lis_tran = append( lis_tran, ""         ) 	
				if ixLemma2 >=0 {	lemmaSlice[ixLemma2].leTran = "" }  
				//if lem == "eindhoven" { fmt.Println("elab_word_list.go ", " NO NO   lemmaSlice[ixLemma2].leTran =",  lemmaSlice[ixLemma2].leTran ) }
		
			}	
			if sw_rewrite_wordLemma_dict { 
				newWL.lLemma = lem 
				newWordLemmaPair = append( newWordLemmaPair, newWL ) 				
			}			
			//--			
			LW.lw_lemma2     = lem                                 
			LW.lw_lemmaSeq   = seqCode( lem )			
			LW.lw_word       = wF.uWord2 
			LW.lw_ixLemma    = ixLemma2
			LW.lw_origLemma  = lis_origLemma[n1] 
			if LW.lw_origLemma == LW.lw_lemma2 { LW.lw_origLemma = "" }
			LW.lw_ixWordUnFr = wF.uIxUnW
			lemma_word_ix  = append( lemma_word_ix, LW )		
			
		} // end of for , lem 
		//-----------
			
		wF.uIxLemmaL= make( []int,     nele, nele )    
		wF.uLemmaL  = make( []string,  nele, nele )    
		//wF.uTranL   = make( []string,  nele, nele )        
		wF.uLevel   = make( []string,  nele, nele )   
		wF.uPara    = make( []string,  nele, nele )   
		wF.uExample = make( []string,  nele, nele )   
		
		
		copy( wF.uIxLemmaL, lis_ixLemma )		
		copy( wF.uLemmaL  , lis_lemmaName ) 
		//copy( wF.uTranL   , lis_tran  )  
		copy( wF.uLevel   , lis_level ) 
		copy( wF.uPara    , lis_para  ) 
		copy( wF.uExample , lis_exam  ) 
		uniqueWordByFreq[zz] = wF
		
		//if swMio {  fmt.Println("ANTONIO unique word=", newWL.lWord2, " wFixLemma=",wF.uIxLemmaL,   " wF uLemmaL=" ,wF.uLemmaL ) }  // , ",    tran=",  wF.uTranL ) }
	
		/**
		if ((wF.uWord2== "tun") || (wF.uWord2 == "umwelt") ) {
				fmt.Println("lookFromLemma( uniqueWordByFreq[",zz,"] = ", uniqueWordByFreq[zz]) 
		} 
		**/
		
		//if sw1 {  fmt.Println( "word=", cyan(wF.uWord2), " lemma=" , wF.uLemmaL) }
		
	}  // end of for zz 
	//-----------
	
	if len(lemmaNotFoundList) > 0 {
		outFile := FOLDER_OUTPUT +  string(os.PathSeparator) + FILE_outLemmaNotFound;
		writeList( outFile, lemmaNotFoundList )		
	}
	
	//countWordLemmaUse() 
	
} // end of addWordLemmaTranLevelParadigma

//---------------------------------------------

func compound_lemma( leS lemmaStruct)  string { 

	  str1:= ""
	  
	  if leS.ls_lemma_ix_stellen >=0 {
			str1+= "<br>" + leS.ls_pref_ein +  "(" +   leS.ls_pref_tran+")" + " + " +  leS.ls_lemma_stellen 
	  }
	  
	  if len(leS.ls_lemma_einStellenList) > 0 {
			for _, ixle2:= range leS.ls_lemma_einStellenList {
				lelem:= lemmaSlice[ixle2]
				str1 += "<br>" + lelem.leLemma  + " (" + lelem.leTran + ")"  + 
					" = " + lelem.ls_pref_ein + "(" + lelem.ls_pref_tran +")" + " + " + lelem.ls_lemma_stellen     
			}
	  } 	
	  if  len(str1) < 4 {return str1} 	  
	  return str1[4:]
}						

//----------------------------
func sPrintOneLemma( ix int, leS lemmaStruct)  string { 

	  str1:= fmt.Sprint(" ixLemma=", ix, " lemma=", leS.leLemma ,  leS.leTran) 
	  
	  if leS.ls_lemma_ix_stellen >=0 {
		str1 += fmt.Sprint(" composto da ",  leS.ls_lemma_stellen, "(", leS.ls_lemma_ix_stellen,") + ", leS.ls_pref_ein , "(" +   leS.ls_pref_tran+")" )
	  }
	  
	  if len(leS.ls_lemma_einStellenList) > 0 {
			str1 += fmt.Sprint(" einStelleList=", leS.ls_lemma_einStellenList )
			for _, ixle2:= range leS.ls_lemma_einStellenList {
				str1 += "\n\t\t" + lemmaSlice[ixle2].leLemma 
			}
	  } 
	  
	  return str1	
		/***
		return fmt.Sprint(" ixLemma=", ix, " lemma=", leS.leLemma , 
		" leTran=", leS.leTran, 
		" ixStell=", leS.ls_lemma_ix_stellen,
		" lemmaStell=", leS.ls_lemma_stellen,
		" pref=", leS.ls_pref_ein,
		" prefTran=", leS.ls_pref_tran,
		" ixLemmaList=", leS.ls_lemma_einStellenList )
		***/
}						

//-----------------------------------------------
func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
//------------------
func build_listOfLemmaForAWord(uWord string, ixLemmaPairFoundList []int) ([]int, []string) {
	
	//sw1:= (strings.Index(uWord, "stellen") >= 0)
	
		nele := len(ixLemmaPairFoundList)
		nele += 20
		
		lis_ixLemma1:=make( [] int,    0, nele )				
		lis_ixLemma2:=make( [] int,    0, nele )		
		lis_ixLemma3:=make( [] int,    0, nele )
		lis_origLemma1:= make( []string,    0, nele )	
		lis_origLemma2:= make( []string,    0, nele )
		lis_origLemma3:= make( []string,    0, nele )	
		
		ixLemma:=-1
		numLerr:=0; maxNumLerr:=100; 
		
		for  _, ixLp := range ixLemmaPairFoundList { 
			
			if numLerr > maxNumLerr { break}
			if ixLp < 0 {
				lemma3 := "?" + uWord	
				ixLemma = addUnknowToLemma(lemma3) 
			} else {
				newWL := wordLemmaPair[ixLp]
				if newWL.lWord2 != uWord { // error 
					continue; 
				}
				ixLemma = newWL.lIxLemma 
			}
			if ixLemma < 0 { continue } // error
			
			leSL:= lemmaSlice[ixLemma]
			
			lis_ixLemma1   = append(  lis_ixLemma1, ixLemma      )  
			lis_origLemma1 = append(lis_origLemma1, leSL.leLemma )  
			
			//if sw1 {  fmt.Println(" lemma1 ", lemmaSlice[ixLemma], " ixLemma=",ixLemma, " orig=", leSL.leLemma) }
			
			// add lemma in case of lemma made of prefix plus other lemma 
			 

			if leSL.leLemma == uWord {  continue } 
			
			if leSL.ls_lemma_ix_stellen >= 0 {
				newIxLemma00 := leSL.ls_lemma_ix_stellen 
				lis_ixLemma2   = append(  lis_ixLemma2, newIxLemma00 ) 	
				lis_origLemma2 = append(lis_origLemma2, leSL.leLemma )  	
			}
			for _, newIxLemma := range  leSL.ls_lemma_einStellenList {
				if newIxLemma >= 0 {
					lis_ixLemma3   = append(  lis_ixLemma3, newIxLemma ) 	
					lis_origLemma3 = append(lis_origLemma3, leSL.leLemma )  	
				}
			} 
				
		}	
		for n2, ix2:= range lis_ixLemma2{
			if contains(lis_ixLemma1, ix2) { continue } 
			lis_origLemma1 = append(lis_origLemma1, lis_origLemma2[n2]  )  
			//if sw1 {  fmt.Println(" lemma2 ", lemmaSlice[ix2], " ixLemma=",ix2, " orig=",  lis_origLemma2[n2] )  }
		}
		for n3, ix3:= range lis_ixLemma3{
			if contains(lis_ixLemma1, ix3) { continue } 
			lis_ixLemma1 = append( lis_ixLemma1, ix3) 
			lis_origLemma1 = append(lis_origLemma1, lis_origLemma3[n3]  )   
			//if sw1 {  fmt.Println(" lemma3 ", lemmaSlice[ix3], " ixLemma=",ix3, " orig=",  lis_origLemma3[n3] )  }	
		}
		
		return lis_ixLemma1,lis_origLemma1 
		
} // end of build_listOfLemmaForAWord

//-----------------------------------------
func add_ixWord_to_WordSliceFreq() {
	var ixFromList, ixToList int
	tot:=0
	for ixWord:=0; ixWord < len(uniqueWordByFreq); ixWord++ {		
		xWordF := uniqueWordByFreq[ixWord] 			
		stat_level( xWordF.uLevel, xWordF.uTotRow)		
		tot+=  xWordF.uTotRow
		ixFromList = xWordF.uIxWordFreq 
		ixToList   = ixFromList + xWordF.uTotRow;
		if ixToList > numberOfWords { ixToList = numberOfWords; }		
		for n1 := ixFromList; n1 < ixToList; n1++  {
			wordSliceFreq[n1].wIxUniq = ixWord 			
		} 	
	}  
	
	fmt.Println(" num. words tutte = ", tot) 
	
	/**
	percA0  =  only_A0 * 100 / tot ;     
	percA1  =  only_A1 * 100 / tot ;     
	percA2  =  only_A2 * 100 / tot ;     
	percB1  =  only_B1 * 100 / tot ;     
	percOth =  only_Ot* 100 / tot ;   
	
	fmt.Println(
		  " num. words A0 = ", only_A0,    " \t", percA0 , "%" , 
		"\n num. words A1 = ", only_A1,    " \t", percA1 , "%" ,       
		"\n num. words A2 = ", only_A2,    " \t", percA2 , "%" ,       
		"\n num. words B1 = ", only_B1,    " \t", percB1 , "%" ,        
		"\n num. words altro= ", only_Ot, " \t", percOth, "%"  ) 
	****/	
	
	for f:=1; f < len( only_level_numWords ) ; f++ {
		if only_level_numWords[f] == 0 { continue }
		perc_level[f] = only_level_numWords[f] * 100 / tot ;     
		//fmt.Println(f, " num. words ","list_level[f]", list_level[f], " = " , "only_level_numWords[f]=", only_level_numWords[f],    " \t", perc_level[f] , "%" ) 
	}	

	if only_level_numWords[0] > 0 {  
		perc_level[0] = only_level_numWords[0] * 100 / tot ;     
		//fmt.Println(" num. words ", list_level[0], " = ", only_level_numWords[0],    " \t", perc_level[0] , "%" ) 	
	}

	
} // end of add_index_toWordSliceFreq

//---------------

//---------------------------------

func addUnknowToLemma( lemma1 string) int {
	var leV lemmaStruct 
	leV.leLemma    = lemma1
	leV.leNumWords = 0; 
	leV.leTran     = ""
	leV.ls_lemma_ix_stellen = -1 	
	leV.ls_lemma_einStellenList = nil

	lemmaSlice = append(lemmaSlice, leV ); 

	return len(lemmaSlice) -1 

} // end of addUnknowToLemma 

//-----------------------------------------

func fromLemmaTo3List( lemma string) (string, string, string) { 

		fromIx, toIx  := lookForAllParadigma( lemma ) 
		
		//fmt.Println( "         fromIx=", fromIx, "  toIx=", toIx )
		
		listLev := ""
		listPara:= ""
		listExam:= ""		
		for ix:= fromIx; ix <= toIx; ix++ {
			listLev  += "|" + lemma_para_list[ix].p_level   
			parad    := lemma_para_list[ix].p_para			
			if len(parad) > 2 { 				
				if parad[0:lenFseq] == fseq {
					parad = parad[lenFseq:] 
				}
			}
			listPara += "|" + parad  		 	
			listExam += "|" + lemma_para_list[ix].p_example 	
	 		//fmt.Println("\t ", ix, " listLev =", listLev)    
		}	
		if toIx >= fromIx {
			listLev = listLev[1:]
			listPara= listPara[1:]	
			listExam = listExam[1:]	
		}
		return listLev, listPara, listExam 
		
} // end of fromLemmaTo3List 		

//-----------------------------------------