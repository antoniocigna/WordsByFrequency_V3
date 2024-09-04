package wbfSubPack

import (  
	"fmt"
    "strings"
    "sort"
)
//------------------------------------------------
//-----------------
const LAST_WORD_FREQ = 999999999 
//---------------------
var righe      =  []string{} 
//------------
//------------
var newWordLemmaPair    [] wordLemmaPairStruct // all word-lemma pair 
var lemma_word_ix []lemmaWordStruct  

var lemmaSlice       [] lemmaStruct         // lemma , translation 

var wordLemmaPair    [] wordLemmaPairStruct // all word-lemma pair  
var numLemmaDict int =0 
var sw_stop bool = false
var errorMSG = ""; 
//----------------------------------------

func read_lemma_file( path1 string, inpLemmaFile_wordLemma, inpLemmaFile_lemmaWord string) {
	
	//showInfoRead( inpLemmaFile, " inizio lettura " )
	
	bytesPerRow:=10
	numLemmaDict=0; 
	
	var wordLemma1 wordLemmaPairStruct 
	//------
	file1_bytes := getFileByteSize(path1, inpLemmaFile_wordLemma)
	file2_bytes := getFileByteSize(path1, inpLemmaFile_lemmaWord)
	fmt.Println("file ", inpLemmaFile_wordLemma, "  ", file1_bytes , " bytes") 
	fmt.Println("file ", inpLemmaFile_lemmaWord, "  ", file2_bytes , " bytes") 
	numEleMax:= int(  (file1_bytes + file2_bytes) / bytesPerRow ); 
	if numEleMax < 10 {numEleMax=10}
	//----------------
    lineS:= rowListFromFile( path1, inpLemmaFile_wordLemma, "1assoc. word-lemma", "read_lemma_file", bytesPerRow)  		
	
	var wordLemmaPairTMP = make( []wordLemmaPairStruct, 0, numEleMax)
	
	if (sw_stop == false) {	
		// read word lemma
		for z:=0; z< len(lineS); z++ { 
			lineZ0 := strings.TrimSpace(lineS[z])   //  format:     word   lemma		
			if lineZ0 == "" {continue}
			lineZ := strings.ReplaceAll( lineZ0, "\t" , " ")			
			
			cols:= strings.Fields( strings.ToLower( lineZ ) )   // Fields   split using whitespace,  treats consecutive whitespace characters as a single separator		
			if len(cols) < 2 { continue } 
			wordLemma1.lWord2   = stdCode( cols[0] ) 		
			wordLemma1.lLemma   = stdCode( cols[1] )	
			
			if len(wordLemma1.lLemma) < 1 { continue;  } 
			if ((wordLemma1.lLemma == "-") || (wordLemma1.lLemma[0:1] < "A")) { continue;  }   // ignore number  
			
			wordLemma1.lWordCod = newCode( wordLemma1.lWord2)
			wordLemma1.lIxLemma = -1	
			
			wordLemmaPairTMP = append(wordLemmaPairTMP, wordLemma1 ) 
			numLemmaDict++		
		}
		fmt.Println(" read ", len(lineS), " input lemma: format word-lemma")
	}
	
	//----------------
    lineS = rowListFromFile( path1, inpLemmaFile_lemmaWord, "2assoc. lemma-word", "read_lemma_file", bytesPerRow)  		
	
	if (sw_stop == false) {	
		// read word lemma
		for z:=0; z< len(lineS); z++ { 
			lineZ0 := strings.TrimSpace(lineS[z])   //  format:     word   lemma		
			if lineZ0 == "" {continue}
			lineZ := strings.ReplaceAll( lineZ0, "\t" , " ")			
			
			cols:= strings.Fields( strings.ToLower( lineZ ) )   // Fields   split using whitespace,  treats consecutive whitespace characters as a single separator		
			if len(cols) < 2 { continue } 	
			wordLemma1.lWord2   = stdCode( cols[0] ) 		
			wordLemma1.lLemma   = stdCode( cols[1] )	
			
			if len(wordLemma1.lLemma) < 1 { continue;  } 
			if ((wordLemma1.lLemma == "-") || (wordLemma1.lLemma[0:1] < "A")) { continue;  } 
			
			wordLemma1.lWordCod = newCode( wordLemma1.lWord2)
			wordLemma1.lIxLemma = -1
				
			wordLemmaPairTMP = append(wordLemmaPairTMP, wordLemma1 ) 
			numLemmaDict++		
		}
		fmt.Println(" read ", len(lineS), " input lemma: format lemma-word")
	}
	
	//-----------------------------------
	lineS = nil
	//---------------------------------------	
	//-----		
	fmt.Println( "lette " , numLemmaDict ,  " coppie word-lemma", "\n")
	//-----	
	// sort x lemma, word
	sort.Slice(wordLemmaPairTMP, func(i, j int) bool {
			if (wordLemmaPairTMP[i].lLemma != wordLemmaPairTMP[j].lLemma) {
				return wordLemmaPairTMP[i].lLemma < wordLemmaPairTMP[j].lLemma
			} else {
					return wordLemmaPairTMP[i].lWord2 < wordLemmaPairTMP[j].lWord2
			}
		} )	 
	/***
	sort.Slice(wordLemmaPairTMP, func(i, j int) bool {
			if (wordLemmaPairTMP[i].lLemma != wordLemmaPairTMP[j].lLemma) {
				return wordLemmaPairTMP[i].lLemma < wordLemmaPairTMP[j].lLemma
			} else {
				if wordLemmaPairTMP[i].lype != wordLemmaPairTMP[j].lype {
					return wordLemmaPairTMP[i].lype < wordLemmaPairTMP[j].lype 
				} else {
					return wordLemmaPairTMP[i].lWord2 < wordLemmaPairTMP[j].lWord2
				}
			}
		} )	 	
		***/
	//-------------------------------------
	
	wordLemmaPair = make( []wordLemmaPairStruct, 0, len(wordLemmaPairTMP)	)
	
	buildListLemmaSlice(wordLemmaPairTMP)
	
	
	//-------------------------------------
	fmt.Println( green("lemmaSlice"), "  composto da ", len(lemmaSlice) , " elementi")    
	
	//-----------------------------
	/**
	seq:=""; swerr:=false
	for  _, lem := range lemmaSlice {
		if lem.leLemma < seq {
			fmt.Println(red("ERRORE lemmaSlice fuori sequenza "), " pre=", seq, "   new=", lem.leLemma )
			//swerr = true
			break;
		}
		seq = lem.leLemma
	}
	
	if swerr == false { fmt.Println( green("lemmaSlice IN SEQUENZA")) 	}
	**/
	//---------------------------------
	
	//--------------------------------
	// sort x word , lemma 
	sort.Slice(wordLemmaPair, func(i, j int) bool {
			if (wordLemmaPair[i].lWordCod != wordLemmaPair[j].lWordCod) {
				return wordLemmaPair[i].lWordCod < wordLemmaPair[j].lWordCod
			} else {
				if (wordLemmaPair[i].lWord2 != wordLemmaPair[j].lWord2) {
					return wordLemmaPair[i].lWord2 < wordLemmaPair[j].lWord2 
				} else {
					return wordLemmaPair[i].lLemma < wordLemmaPair[j].lLemma
				}
			}
		} )	 
	//------------------------	
	
	check_wordLemma_sameCode()
	
	
	
	//-------------------------------
} // end of  read_lemma_file

//-------------------------------------------

func check_wordLemma_sameCode() {
	fmt.Println( green("check_wordLemma_sameCode") , "()"  )
	// check same words  written in diffent way (eg. caesar   and  "cäsar")
	pre_wordCod := ""
	pre_word2   := ""	
	//pre_lemma   := ""
	pre_z := -1
	
	for z, wordPair := range wordLemmaPair {	
			//if ((  wordPair.lWord2 == "cäsar") || (wordPair.lWord2 == "caesar") || (wordPair.lWord2 == "casar") ) { fmt.Println(" check 222 Lemma ", z,  " wordPair=" , wordPair) }
	
		if (wordPair.lWordCod != pre_wordCod) {
			pre_wordCod = wordPair.lWordCod 
			pre_word2   = wordPair.lWord2 
			//pre_lemma   = wordPair.lLemma 
			pre_z = z
			continue
		}
		if (wordLemmaPair[z].lWord2 == pre_word2) {
			continue
		}
		//--------
		fmt.Println( green("check_wordLemma_sameCode") )
		for x:= pre_z; x<= z; x++ {
			fmt.Println("\t", " wordLemmaPair[",x,"] = ", wordLemmaPair[x] )   
		} 
		
	}	
	
} // end of check_wordLemma_sameCode

//--------------------------------------


//----------------------------------------
func stdCode(inpCode string ) string {	
	CoerInp:= strings.ReplaceAll( 
					strings.ReplaceAll( 
						strings.ReplaceAll( 
							strings.ReplaceAll(inpCode, "ae","ä"),  
							"oe","ö"), 
						"ue","ü"),
					"ß","ss")  	           // non tutte le ss sono 	ß, ma è vero il contrario				
					
	return CoerInp  
}
//----------------------------
func newCode( inpCode string ) string {	

	//  pronto soltanto per il tedesco 

	/*
	1) serve soprattutto per mettere le parole in sequenza alfabetica coerente 
		( es. per il tedesco es.  ä, ö, ü, ß vicini rispettivamente ad a, o, u, ss)   
	2) a volte ä, ö, ü, ß sono scritti come ae, oe, ue, ss, in questo caso li sostituiamo con a, o, u, ss
    3) a volte eu dovrebbe rimanere tale (es. Treue), non ho modo di distinguere per cui la sequenza è falsata ue è prima di ua o di uz 
		l'alternativa potrebbe essere tradurre ü con ue, ma questo porterebbe alla sequenza errata  ue nel posto ua, ue, uz invece di u ua uz
	----------
	questo codice di sequenza   è usato per word e lemma
	la scrittura alternativa a quella ufficiale (es. ue invece di ü) può essere trovata in un testo.
	Il lemma si trova in un file "ufficiale"  quindi improbabile che venga usata la scrittura alternativa.

	a) nel lemma il newCode mi serve soltanto per correggere la sequenza
	b) nel word (che si trova nel testo analizzato) il new code mi serve per confrontare 
		però è probabile che un testo sia scritto o in codice alternativo o in modo normale, improbabile in entrambi modi. 
	c) si potrebbe pensare ad un switch da impostare 
	d) cosa devo confrontare?
		lemma ( non mi serve tradurre eventuale codice alternativo, però newCode mi serve per la sequenza )
			confronto per collegarlo a word e per trovare la traduzione
		word  ( confronto per assegnare il lemma )       	
	
	*/
	SQinp1:= strings.ReplaceAll( 
						strings.ReplaceAll( 
							strings.ReplaceAll(inpCode, "ae","a"),  
							"oe","o"), 
						"ue","u") 		
	SQinp2:= strings.ReplaceAll( 
					strings.ReplaceAll( 
						strings.ReplaceAll( 
							strings.ReplaceAll(SQinp1, "ä","a"),  
							"ö","o"), 
						"ü","u"), 
					"ß","ss")   					
	
	return SQinp2 + "." +  stdCode(inpCode)  
	
}// end of newCode					

//------------------------------------------------
