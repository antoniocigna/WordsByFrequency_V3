package wbfSubPack

import (  
	"fmt"
    "strings"
	//"strconv"
)

//---------------------------
func read_wordsToLearn() {

	//swWrite:=false; outLearn:= make([]string,0,3000) 
	
	bytesPerRow:= 10
    lineD := rowListFromFile( FOLDER_INPUT_OUTPUT, FILE_words_to_learn, "words to learn", " bind_go_passToJs_read_wordsToLearn", bytesPerRow)  
	if sw_stop {  // this file might be missing
		sw_stop=false		
		return
	}
	
	nread:=0
	for z:=0; z< len(lineD); z++ { 
		fields:= strings.Split( lineD[z] ,"|") 
		
		
		if len(fields) < 2 { continue}
		 
		r_word2          := strings.TrimSpace( fields[0] ) 	
		yesNo:= strings.ToLower( strings.TrimSpace( fields[1] ) )	 
		if len(yesNo) > 1  { yesNo = yesNo[0:1]}
		if yesNo != LEARNED_YES { yesNo = LEARNED_NOT }	
		
		wordCod:= seqCode( r_word2)		
	
		/**
			uWordSeq    string	
			uWord2      string	
			uIxUnW      int            // index of this word in the uniqueWordByFreq	
			uIxUnW_al   int            // index of this word in the uniqueWordByAlpha 	
			uTotRow     int 
			uTotExtrRow int
			uIxWordFreq int            // index of this word in the wordSliceFreq	
			uSwSelRowG  int
			uSwSelRowR   int  
			uLearnedYN   string  
		**/
		
		// ignoro l'indice dell'input (potrebbero esserci state delle variazioni nella freq. delle parole) e lo ricalcolo ( ottengo in realtà un range di indici che dovrebbero coincidere)
		ixF, ixT:= lookForWordInUniqueAlpha( wordCod,0)		
		if (ixT < 0) {
				fmt.Println(red("error in" + "wordToLearn "), r_word2 , " not found in wordUniqueAlpha" ); 	
				continue
		}		
		for ixA:= ixF; ixA <= ixT; ixA++ {
			xWordA :=  uniqueWordByAlpha[ixA]
			if xWordA.uWordSeq != wordCod {
				continue
			}
			uniqueWordByAlpha[ixA].uLearnedYN = yesNo;
			
			if (ixA != xWordA.uIxUnW_al) {  
				fmt.Println( red("error in" + "wordToLearn "), r_word2 , " z=",z," fields=", fields, " ixA=", ixA, " xWordA.uIxUnW_al=",xWordA.uIxUnW_al);  
				continue
			}				
			nread++	
		}
	}
	
	/**
	if swWrite {
			writeList("NUOVO_wordToLearn.txt", outLearn) 	
	}
	**/
	
	fmt.Printf("letti %d parole da imparare  dal file %s\n", nread,  FILE_words_to_learn)
	
	//go_exec_js_function( js_function, outS1 ); 	
	
 } // end of read_words_to_learn_file 

//---------------------------------------