package wbfSubPack

import (  
	"fmt"
    "strings"
	"strconv"
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
		
		/***
		if len(fields) == 4 {
			field01 := strings.TrimSpace(fields[0])[0:1]
			if z < 10 { fmt.Println( "1 toLearn ", fields) }
			if ((field01 >="0") && (field01 <="9")) { // si tratta della vecchia versione con indice nella prima posizione  
				fields = fields[1:]
				swWrite = true
			}	
			if z < 10 { fmt.Println( "2 toLearn ", fields) 	}
		}
		
		if swWrite {
			outLearn = append( outLearn, ( fields[0] + "|" +  fields[1] + "|" + fields[2] ) )   
		}
		****/
		
		if len(fields) < 3 { continue}
		 
		r_word2          := strings.TrimSpace( fields[0] ) 	
		known_yes_ctr, _ := strconv.Atoi( strings.TrimSpace( fields[1] ) ) 
		known_no_ctr , _ := strconv.Atoi( strings.TrimSpace( fields[2] ) ) 
		
		wordCod:= newCode( r_word2)		
	
		/**
			uWordCod    string	
			uWord2      string	
			uIxUnW      int            // index of this word in the uniqueWordByFreq	
			uIxUnW_al   int            // index of this word in the uniqueWordByAlpha 	
			uTotRow     int 
			uTotExtrRow int
			uIxWordFreq int            // index of this word in the wordSliceFreq	
			uSwSelRowG  int
			uSwSelRowR   int  
			uKnow_yes_ctr int 
			uKnow_no_ctr  int 
		**/
		
		// ignoro l'indice dell'input (potrebbero esserci state delle variazioni nella freq. delle parole) e lo ricalcolo ( ottengo in realtà un range di indici che dovrebbero coincidere)
		ixF, ixT:= lookForWordInUniqueAlpha( wordCod)		
		if (ixT < 0) {
				fmt.Println(red("error in" + "wordToLearn "), r_word2 , " not found in wordUniqueAlpha" ); 	
				continue
		}		
		for ixA:= ixF; ixA <= ixT; ixA++ {
			xWordA :=  uniqueWordByAlpha[ixA]
			if xWordA.uWordCod != wordCod {
				continue
			}			
			uniqueWordByAlpha[ixA].uKnow_yes_ctr = known_yes_ctr
			uniqueWordByAlpha[ixA].uKnow_no_ctr  = known_no_ctr
			if (ixA != xWordA.uIxUnW_al) {  
				fmt.Println( red("error in" + "wordToLearn "), r_word2 , " z=",z," fields=", fields, " ixA=", ixA, " xWordA.uIxUnW_al=",xWordA.uIxUnW_al);  
				continue
			}
			ix1 := xWordA.uIxUnW
			uniqueWordByFreq[ix1].uKnow_yes_ctr = known_yes_ctr
			uniqueWordByFreq[ix1].uKnow_no_ctr  = known_no_ctr 		

			//if nread < 10 { fmt.Println( "word to learn ", uniqueWordByFreq[ix1] , " no_ctr=", uniqueWordByFreq[ix1].uKnow_no_ctr) }
			
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