package wbfSubPack

import (  
	"fmt"
	"os"	
    "bufio"	
	"sort"
	"time"
	"strconv"
	"strings"
)
//--------------------------------

func rewrite_word_lemma_dictionary() {	
	
	//----------------------------------------------------
	sort.Slice(newWordLemmaPair, func(i, j int) bool {
			if (newWordLemmaPair[i].lWordSeq != newWordLemmaPair[j].lWordSeq) {
				return newWordLemmaPair[i].lWordSeq < newWordLemmaPair[j].lWordSeq
			} else {
				if (newWordLemmaPair[i].lWord2 != newWordLemmaPair[j].lWord2) {
					return newWordLemmaPair[i].lWord2 < newWordLemmaPair[j].lWord2 
				} else {
					return newWordLemmaPair[i].lLemma < newWordLemmaPair[j].lLemma
				}
			}
		} )		 	
	//------------			
	outFile := FOLDER_OUTPUT +  string(os.PathSeparator) + FILE_outWordLemmaDict ;		
	
	lines:= make([]string, 0, 10+len(newWordLemmaPair) )

	lines = append(lines,  "__" + outFile + "\n" + "_word _lemma ")
	
 	for z:=0; z < len( newWordLemmaPair); z++ {
		//lines = append(lines,  newWordLemmaPair[z].lWord2 + "|" + newWordLemmaPair[z].lLemma) 
		lines = append(lines,  newWordLemmaPair[z].lWord2 + " " + newWordLemmaPair[z].lLemma) 
	}  	
	
    writeList( outFile, lines )
	
	
} // end of rewrite_word_lemma_dictionary

//--------------------------------
func g35_rewriteUP_LemmaTranDict_file() {
	
	//fmt.Println("rewriteUP_LemmaTranDict_file  len(dictLemmaTranUP) = ", len(dictLemmaTranUP) )
	
	outFile := FOLDER_IO_lastTRAN  +  string(os.PathSeparator) + FILE_last_updated_dict_words 

	pkey := ""; key := ""	
	pkeyLemma:=""
	keyLemma:=""
	
	lines:= make([]string, 0, 10+len(dictLemmaTranUP) )
	lines = append(lines,  "__" + outFile + "\n" + "_lemma	_traduzione")
	
	sort_lemmaTranUP()   // necessario di nuovo il sort perchè dopo il caricamento inziale sono stati accodati degli elementi 
	//------------------
	for z:=0; z < len(dictLemmaTranUP); z++ {
		//fmt.Println( "dictLemmaTranUP[",z,"]=", dictLemmaTranUP[z])
		
		keyLemma = strings.TrimSpace( dictLemmaTranUP[z].dL_lemma2 )
		
		if len(keyLemma) < 1 {continue}
		if keyLemma[:1] < "a" { continue} 	
		
		key = keyLemma + "|" + dictLemmaTranUP[z].dL_tran  + "|" + strconv.Itoa(dictLemmaTranUP[z].dL_numDict)  //  dictLemmaTranUP[z].dL_lemma2 + "|"  + dictLemmaTranUP[z].dL_tran  			
		if pkeyLemma != keyLemma { 
		   if pkey != "" { 
				lines = append(lines, pkey ) 
		   }
		   pkeyLemma = keyLemma	
		}
		pkey=key
	}
	//---
	if pkey != "" { 
		lines = append(lines, pkey ) 	
	}
	fmt.Println("lemma translation update ", " letti ", len(dictLemmaTranUP), ", scritti ", len(lines) , " sul file ", outFile  )
		   
	writeList( outFile, lines )
	//--------------------

	currentTime := time.Now()		
	outF1 		:= FOLDER_O_arc_TRAN_words +  string(os.PathSeparator) + "dictL"  		
	outFile2 := outF1 + currentTime.Format("20060102150405") + ".txt"
	
	writeList( outFile2, lines )
	
	
} // end of rewrite_LemmaTranDict_file
//----------------------------------------------

func g35_bind_go_rewrite_allTran() {
	
	
	outFile := "__all_updated_lemma_tran.csv.txt" 

	pkey := ""; key := "";	pkey00:="";	key00:=""; pkeyLemma:=""; keyLemma:=""
		
	lines:= make([]string, 0, 10+len(dictLemmaTran) )
	lines = append(lines,  "__" + outFile + "\n" + "_lemma	_traduzione")
		
	//------------------
	for z:=0; z < len(dictLemmaTran); z++ {
		keyLemma = strings.TrimSpace(dictLemmaTran[z].dL_lemma2) 
		if len(keyLemma) < 1 {continue}
		if keyLemma[:1] < "a" { continue} 	
			
		key00 = keyLemma + "|" + dictLemmaTran[z].dL_tran 
		
		key = key00 + "|" + strconv.Itoa(dictLemmaTran[z].dL_numDict)  //  dictLemmaTranUP[z].dL_lemma2 + "|"  + dictLemmaTranUP[z].dL_tran  			
		if pkeyLemma != keyLemma { 
		   if pkey != "" { 
				lines = append(lines, pkey00) 
		   }
		}
		pkey      = key	
		pkey00    = key00	
		pkeyLemma = keyLemma	
	}
	//---
	if pkey != "" { 
		lines = append(lines, pkey00 ) 		
	}
	fmt.Println("lemma translation ", " letti ", len(dictLemmaTran), ", scritti ", len(lines) , " sul file ", outFile )
	fmt.Println(red("sostituisci con questo file (dopo averlo rinominato) "), " il file ", 
		`"WBF_INPUT_DE/inputTranslation/input_dict_tran_words.csv" `)
	fmt.Println("\t e poi azzera il file ", `"INPUT_OUTPUT\lastTRAN\lastUpdated_dict_tran_words.csv"` )
	
		   
	writeList( outFile, lines )
	
	
} // end of g35_bind_go_rewrite_allTran
//----------------------
func writeList( fileName string, lines []string)  {
	// create file
    f, err := os.Create( fileName )
    if err != nil {
        fmt.Println( red("error")," in writeList file=", fileName,"\n\t" , err ) //  log.Fatal(err)
    }
    // remember to close the file
    defer f.Close()

    // create new buffer
    buffer := bufio.NewWriter(f)

    for _, line := range lines {
        _, err := buffer.WriteString(line + "\n")
        if err != nil {
           fmt.Println( red("error"), " in buffer.WriteString file=", fileName,"\n\t" , err ) //log.Fatal(err)
        }
    }
    // flush buffered data to the file
    if err := buffer.Flush(); err != nil {
        fmt.Println( red("error"), " in buffer.Flush()cls file=", fileName,"\n\t" , err ) //  log.Fatal(err)
    }
} 
//----------------------------------------