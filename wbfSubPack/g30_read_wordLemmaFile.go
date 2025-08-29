package wbfSubPack

import (  
	"fmt"
    "strings"
    "sort"
	//"slices"
	"regexp"
	"time"
)
//------------------------------------------------
//-----------------
const LAST_WORD_FREQ = 999999999 
//---------------------
var righe      =  []string{} 
//------------

var numLemmaDict int =0 
//----------------------------------------
func getLemmaPathAndFile(inpLemmaFile_wordLemma string) (string, string) {
	path1:=""
	bar1:= strings.Index(inpLemmaFile_wordLemma,"/") 
	if bar1 < 0 {bar1 = strings.Index(inpLemmaFile_wordLemma,"\\") }	
	if bar1 > 0 {
		barZ:= inpLemmaFile_wordLemma[bar1:bar1+1]
		for z:= len(inpLemmaFile_wordLemma)-1; z > 0; z-- {
			if inpLemmaFile_wordLemma[z:z+1] == barZ {
				path1 = inpLemmaFile_wordLemma[0:z]
				inpLemmaFile_wordLemma = inpLemmaFile_wordLemma[z+1:]
				break
			}
		} 	
	}  	
	return path1, inpLemmaFile_wordLemma 
	
}

//----------------
const NO_LEMMA_WORD  = "aaalemmanotfound"
const NO_LEMMA_LEMMA = "_lemma_not_found"
var NO_LEMMA_INDEX = 0
//----------------------------------
func g30_read_wordLemma_file( path0 string, inpLemmaFile_wordLemma0 string,  inpLemmaFile_wordLemmaPlus0 string) {
	startTime := time.Now()
	bytesPerRow:=20
	numLemmaDict=0; 
	
	var wordLemma1 wordLemmaPairStruct
	
	path1, inpLemmaFile_wordLemma := getLemmaPathAndFile(inpLemmaFile_wordLemma0)
	if path1 == "" { path1 = path0}
	
	path2, inpLemmaFile_wordLemmaPlus := getLemmaPathAndFile(inpLemmaFile_wordLemmaPlus0)
	if path2 == "" { path2 = path0}
	
	fmt.Println(green("read_wordLemma_file 1 "), 
		" path = ", path1 ,
		"\n		inputLemmaFile     =", inpLemmaFile_wordLemma, 
		"\n		inputLemmaFilePlus =", inpLemmaFile_wordLemmaPlus, 
		"\n-------------------------------------------------" ) 
	
	
		
	fmt.Println(green("\n  read_wordLemma_file 2 read "), " file ", inpLemmaFile_wordLemma, " in folder ", path1)
	//------
	file1_bytes := getFileByteSize(path1, inpLemmaFile_wordLemma)
	fmt.Println("file ", inpLemmaFile_wordLemma, "  ", file1_bytes , " bytes") 
	numEleMax:= int(  file1_bytes / bytesPerRow ); 
	if numEleMax < 10 {numEleMax=10}
	//----------------
    lineS:= rowListFromFile( path1, inpLemmaFile_wordLemma, "1assoc. word-lemma", "read_wordLemma_file", bytesPerRow)  	
	if len(lineS) == 0 { sw_stop = false }	
	if sw_stop { return }
	fmt.Println("lette ", len(lineS), " coppie word-lemma")  
	lineS = append(lineS, NO_LEMMA_WORD + "|" + NO_LEMMA_LEMMA + "|" )
	/**
	for z, line := range lineS {	// eseguito in ordine di lemma e poi word 	
		if z < 10 { fmt.Println("lines ", z , " " , line) }
	}
	**/
	//-----------------
	if inpLemmaFile_wordLemmaPlus != "" {
		fmt.Println(green("\n  read_wordLemma_file 3 read"), " file ", inpLemmaFile_wordLemmaPlus, " in folder ", path2)
		//------
		file2_bytes := getFileByteSize(path1, inpLemmaFile_wordLemmaPlus)
		fmt.Println("file ", inpLemmaFile_wordLemmaPlus, "  ", file2_bytes , " bytes") 
		numEleMax2:= int(  file2_bytes / bytesPerRow ); 
		if numEleMax2 < 10 {numEleMax2=10}
		//----------------
		lineS2:= rowListFromFile( path2, inpLemmaFile_wordLemmaPlus, "1assoc. word-lemma", "read_wordLemma_file", bytesPerRow)  	
		
		if len(lineS2) == 0 { sw_stop = false }	
		fmt.Println("lette ", len(lineS2), " coppie word-lemmaPlus")  	
		if len(lineS2) > 0 {
			lineS = append(lineS, lineS2...)
			numEleMax += numEleMax2
		}	
		if sw_stop {fmt.Println( red("sw_stop in g30_read_wordLemma_file 1"))	}
	} 
	//-----------------------------------
	if sw_stop {fmt.Println( red("sw_stop in g30_read_wordLemma_file 2")) }
	
	listAllLemmaFromFile = make([]string, 0, numEleMax )
	
	if (sw_stop) {	return }
		
	// read word lemma
	for z:=0; z< len(lineS); z++ { 
		lineZ0 := strings.ToLower( strings.TrimSpace(lineS[z]) )   //  format:     word   lemma		
		if lineZ0 == "" {continue}
		
		cols := strings.Fields( regexp.MustCompile(separWord).ReplaceAllString(lineZ0," ") ) 
		//cols:= strings.Split( strings.ToLower( lineZ0 ), "|" )   // Fields   split using whitespace,  treats consecutive whitespace characters as a single separator		
		if len(cols) < 2 { continue } 
		wordLemma1.lWord2   = stdCode( strings.TrimSpace( cols[0] ) ) 		
		wordLemma1.lLemma   = stdCode( strings.TrimSpace( cols[1] )	)
		
		if len(wordLemma1.lLemma) < 1 { continue;  } 
		if ((wordLemma1.lLemma == "-") || (wordLemma1.lLemma[0:1] < "A")) { continue;  }   // ignore number  
		
		//wordLemma1.lWordSeq     = seqCode( wordLemma1.lWord2)
		wordLemma1.lIxLemma     = -1
		wordLemma1.lIxUnWord_al = -1
		
		wordLemmaPair = append(wordLemmaPair, wordLemma1 ) 
		numLemmaDict++				
	}
	fmt.Println(" read ", len(lineS), " input lemma: format word-lemma")
	
	printTimeDiff(startTime, "tempo di esecuzione di ...wordLemma_file - read file")
	
	//--------------------------------------------------------------
 	// sort x lemma, l_WordSeq               	
 	sort.Slice(wordLemmaPair, func(i, j int) bool {
 			if (wordLemmaPair[i].lLemma != wordLemmaPair[j].lLemma) {
 				return wordLemmaPair[i].lLemma < wordLemmaPair[j].lLemma
 			} else {
 				return wordLemmaPair[i].lWord2 < wordLemmaPair[j].lWord2 				 
 			}
 		} )	 
 	//------------------------------
		
	numLemma:=0
	preLemma:=  ""
	numLemmaAdded:=0
	numLemmaOrig:=0 
	z:=-1 // minus 1
	fromIx:=0
	toIx:=0
	numW:=0
	//------------
	
	
	startTime = time.Now()
	/*  ----------------------------
		crea lemmaSlice ogni elemento contiene i range di indice fromIx toIx che puntano alle entrate word-lemma corrispondente
		( word_lemma sono in ordine prima di Lemma e poi di Word) 
		e le entrate word_lemma contengono l'indice del lemma a cui si riferiscono 
	---------------------- 
	*/
	for _, wL := range wordLemmaPair {	// eseguito in ordine di lemma e poi word 			
		z++
		//--
		if wL.lWord2 == NO_LEMMA_WORD {
			if wL.lLemma == NO_LEMMA_LEMMA {
				NO_LEMMA_INDEX = z 		
				fmt.Println( red("no lemma entry index="), NO_LEMMA_INDEX, " entry: ", wL )
			}
		}
		//---
		if preLemma != wL.lLemma { 	
			if numW > 0 {
				//scrive lemma precedente 	
				numLemmaOrig, numLemmaAdded	= g14_appendOneLemma( preLemma, fromIx, toIx, numLemmaOrig, numLemmaAdded)
				listAllLemmaFromFile = append(listAllLemmaFromFile, preLemma)  // non mi serve lemmaSlice è nello stesso ordine
			}			
			numW=0
			fromIx=z;  
		} 
		numW++
		toIx=z	
		preLemma = wL.lLemma	
	} 
	///-------------
	if numW > 0 {
		//scrive lemma precedente 				
		numLemmaOrig, numLemmaAdded	= g14_appendOneLemma( preLemma, fromIx, toIx, numLemmaOrig, numLemmaAdded)	
		listAllLemmaFromFile = append(listAllLemmaFromFile, preLemma)
	}
	
	g14_updateLemmaWithTranPara()
	
	printTimeDiff(startTime, "tempo di esecuzione di ...wordLemma_file - upd tran")
	
	//------------------------------	
	
	wordLemmaPair_lemmaWordSeq = make( []wordLemmaPairStruct, len(wordLemmaPair) )
	copy(wordLemmaPair_lemmaWordSeq,wordLemmaPair) 

	
	//--------------------------------------
 	// sort x lemma, L_W, l_Word2               	
 	sort.Slice(wordLemmaPair, func(i, j int) bool {
 			if (wordLemmaPair[i].lWord2 != wordLemmaPair[j].lWord2) {
 				return wordLemmaPair[i].lWord2 < wordLemmaPair[j].lWord2
 			} else {
 				return wordLemmaPair[i].lLemma < wordLemmaPair[j].lLemma 					
 			}
 		} )	 
	//---------------------
	/**	
	for z, wL := range wordLemmaPair {	// eseguito in ordine di lemma e poi word 	
		if wL.lWord2 == NO_LEMMA_WORD {
			if wL.lLemma == NO_LEMMA_LEMMA {
				NO_LEMMA_INDEX = z 		
				fmt.Println("no lemma entry index=", NO_LEMMA_INDEX, " entry: ", wL )
				break
			}
		}
	}	
	**/
	//----------------------------------			
	fmt.Println( "lette " , numLemmaDict ,  " coppie word-lemma", "    ", numLemma, " lemma")
	fmt.Println( "wordLemmaPair è in ordine di wordSeq, lemma")
	//----------------------------------
	
	//g34_loadInverseLemmaSlice()	
	
	//--------------------------
	
}  // end of read_lemma_file	

//-----------------------------------------
