package wbfSubPack

import (  
	"fmt"
	"os"	
    "bufio"	
	"sort"
	"time"
	//"strings"
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

func rewrite_LemmaTranDict_file() {

	//fmt.Println( "GO ", green("rewrite_LemmaTranDict_file" ))
	
	//outFile := FOLDER_IO_lastTRAN +  string(os.PathSeparator) + FILE_ outLemmaTranDict;
	
	outFile := FOLDER_IO_lastTRAN  +  string(os.PathSeparator) + FILE_last_updated_dict_words 

	pkey := ""; key := ""
	
	lines:= make([]string, 0, 10+len(dictLemmaTran) )
	lines = append(lines,  "__" + outFile + "\n" + "_lemma	_traduzione")
	
	for z:=0; z < len(dictLemmaTran); z++ {
	
		pkey=key
		
		
		
		key = dictLemmaTran[z].dL_lemma2 + "|"  + dictLemmaTran[z].dL_tran  		////cigna1_3
		
		//if strings.Index(key,"eindhoven")>=0 { fmt.Println("rewrite_LemmaTranDict_file ",  dictLemmaTran[z],  " key=",key, " pkey=", pkey) }
		
		if pkey == key { continue}
		
		lines = append(lines, key ) 
	}
	writeList( outFile, lines )
	//--------------------

	currentTime := time.Now()		
	outF1 		:= FOLDER_O_arc_TRAN_words +  string(os.PathSeparator) + "dictL"  		
	outFile2 := outF1 + currentTime.Format("20060102150405") + ".txt"
	
	writeList( outFile2, lines )
	
	
} // end of rewrite_LemmaTranDict_file

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