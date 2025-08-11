package wbfSubPack

import (  
	"fmt"
    "strings"
    "sort"
	//"slices"
)
//----------------------------------------------------
//------------
var newWordLemmaPair    [] wordLemmaPairStruct // all word-lemma pair 
var lemma_word_ix []lemmaWordStruct  

var lemmaSlice       [] lemmaStruct         // lemma , translation 

//var wordLemmaPair    [] wordLemmaPairStruct // all word-lemma pair  
//----------------------------------------

func TOGLIg10_build_extrLemmaPair() {
	fmt.Println( red("variato L_W  sostituito con riecrca diretta del word o lemma"))
	var wordLemma1 wordLemmaPairStruct 
	fmt.Println(" 2 wordLemmaPair  len=", len(wordLemmaPair) )
	prePa :=""
	for _, unaParola := range all_words {	// tutte le parole delle righe di testo
		if unaParola == prePa { continue }
		//fmt.Println("parola ", unaParola); 
		prePa = unaParola	
		parolaZ:= strings.ToLower( strings.TrimSpace( strings.ReplaceAll( unaParola, "\t" , " ") )  ) 			
		wordLemma1.lWord2   = stdCode( parolaZ ) 		
		//wordLemma1.lLemma   = wordLemma1.lWord2			
		wordLemma1.lLemma   = LEMMA_MISSING            // segnala che il lemma è mancante  (questo wordLemma1 struct sarà ignorato se esiste un'entrata valida   		
		wordLemma1.lWordSeq = seqCode( wordLemma1.lWord2)
		wordLemma1.lIxLemma = -1
		//wordLemma1.lL_W     = 9           // indica che l'origine della coppia è il file di testo 				
		wordLemmaPair = append(wordLemmaPair, wordLemma1 ) 
		//wordLemma1.lL_W     = 0           // indica che l'origine della coppia è il file di testo 	
		wordLemmaPair = append(wordLemmaPair, wordLemma1 ) 
	}
	fmt.Println(" 3 wordLemmaPair  len=", len(wordLemmaPair) )
	//---------------
	fmt.Println("aggiunti alle coppie word-lemma ", len(all_words), " coppie ottenute da tutte le parole del testo (nel caso in cui i lemma mancano)")  
	//-----------------------------------
	

	// sort x lemma, L_W, l_Word2  
	/**	
	sort.Slice(wordLemmaPair, func(i, j int) bool {
			if (wordLemmaPair[i].lLemma != wordLemmaPair[j].lLemma) {
				return wordLemmaPair[i].lLemma < wordLemmaPair[j].lLemma
			} else {
				if (wordLemmaPair[i].lWord2 != wordLemmaPair[j].lL_W) {
					return wordLemmaPair[i].lWord2 < wordLemmaPair[j].lL_W
				} else {
					return wordLemmaPair[i].lL_W < wordLemmaPair[j].lWord2
				}	
			}
		} )	 	
	**/
	wordLemmaPair = make( []wordLemmaPairStruct, 0, len(wordLemmaPair)	)
	
	fmt.Println(cyan(" 4 wordLemmaPair"), "  len=", len(wordLemmaPair), " coppie" )
	fmt.Println("     esistono ", len(separPrefList), " prefissi separabili")
	
	TOGLIestraeCoppieWordLemmaInTesto(wordLemmaPair) 
	
	TOGLIg14_buildListLemmaSlice(wordLemmaPair)
	/**
	for _, wD:= range wordLemmaPair {
		if wD.lWord2 == "am" {  fmt.Println( "dopo g14_buildListLemmaSlice  wordLemmaPair=", wD) }
	}
	**/
	//-------------------------------------
	fmt.Println( green("lemmaSlice"), "  composto da ", len(lemmaSlice) , " elementi")    
	
	//-----------------------------
	/***
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
	//--------------------------------
	// sort x word , lemma 
	sort.Slice(wordLemmaPair, func(i, j int) bool {
			if (wordLemmaPair[i].lWordSeq != wordLemmaPair[j].lWordSeq) {
				return wordLemmaPair[i].lWordSeq < wordLemmaPair[j].lWordSeq
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


//var ixAF = binarySearch_string(listAllLemmaFromFile, myLemma) 
//-----------------------------
func addToCurrentLemmaPair(wordLemmaPair []wordLemmaPairStruct ) {
	sort.Strings(lemmaNotFoundList)
	preL := ""
	for _, oneLemma:= range lemmaNotFoundList { 
		if oneLemma == preL { continue } 
		preL = oneLemma  
		var wordLemma1 wordLemmaPairStruct 
		wordLemma1.lWord2   = stdCode( oneLemma ) 		
		wordLemma1.lLemma   = wordLemma1.lWord2 
		if len(wordLemma1.lLemma) < 1 { continue;  } 
		if ((wordLemma1.lLemma == "-") || (wordLemma1.lLemma[0:1] < "A")) { continue;  }   // ignore number  
		wordLemma1.lWordSeq = seqCode( wordLemma1.lWord2)
		wordLemma1.lIxLemma = -1	
		wordLemmaPair = append(wordLemmaPair, wordLemma1 ) 
	}
}
//-----------------------------
func check_wordLemma_sameCode() {
	fmt.Println( green("check_wordLemma_sameCode") , "()"  )
	// check same words  written in diffent way (eg. caesar   and  "cäsar")
	pre_wordCod := ""
	pre_word2   := ""	
	//pre_lemma   := ""
	pre_z := -1
	
	for z, wordPair := range wordLemmaPair {	
		if ((  wordPair.lWord2 == "abgehauen") || (wordPair.lLemma == "abhauen") ) { fmt.Println(green("check_wordLemma_sameCode abhauen "), "z=", z,  " wordPair=" , wordPair) }
	
		if (wordPair.lWordSeq != pre_wordCod) {
			pre_wordCod = wordPair.lWordSeq 
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

