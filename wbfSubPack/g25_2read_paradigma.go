package wbfSubPack

import (  
	"fmt"
    "strings"
	"sort"
)
//-----------------------------------------------
func findIndex(slice []int, target int) int {
    for i, v := range slice {
        if v == target {
            return i
        }
    }
    return -1
}

//-----------------------------------
func g25_2read_ParadigmaFile( path1 string, inpFile string) {
	bytesPerRow:= 40
    righe := rowListFromFile( path1, inpFile, "paradigma", "read_ParadigmaFile", bytesPerRow)  
	if sw_stop { return }
	
	/*
		  0  	| 1 |                 2              |                                  3                   
		ab  	| A2| abholen, holt ab, hat abgeholt |Wann kann ich die Sachen bei dir abholen<br>Wir müssen noch meinen Bruder abholen  | 
	    
		aber	| A1|aber	 |Der Film ist traurig, aber sehr schön.
		aber	| A2|aber    |Heute kann ich nicht kommen, aber morgen habe ich Zeit.<br>Wir haben nur eine kleine Wohnung, sind aber damit zufrieden.<br>Es war sehr schön.<br>Jetzt muss ich aber gehen.<br>Das ist aber nett von dir.
		aber	| B1|aber    |1. Heute kann ich nicht, aber morgen ganz bestimmt.<br>2. Es lag sehr viel Schnee, aber Enzo ist trotzdem mit dem Motorrad gefahren.<br>3. Wir haben nur eine kleine Wohnung, sind aber damit zufrieden.<br>4. Es war sehr schön.<br>Jetzt muss ich aber gehen.<br>5. Ich würde gerne kommen, aber es geht leider nicht.<br>6. Darf ich dich zu einem Kaffee einladen?<br>– Aber ja, sehr gern.<br>7. Du spielst aber gut Klavier. |                                                                  |    |  |        
	
		abfahren| A1|abfahren|Der Zug fährt gleich ab.
		abfahren| B1|abfahren, fährt ab, fuhr ab, ist abgefahren |Unser Zug ist pünktlich abgefahren.
		gehen	| A1|gehen, geht, ging, ist gegangen.|Das geht nicht!
		gehen	| A1|gehen, geht, ging, ist gegangen.|Ich muss zum Arzt gehen.
		gehen	| A1|gehen, geht, ging, ist gegangen.|Ich weiß nicht, wie das geht.
		gehen	| A1|gehen, geht, ging, ist gegangen.|Jetzt muss ich (aber) leider gehen.
	*/ 
	
	fmt.Println("\nletti paradigma file ", inpFile  + " " , len(righe) , " righe")   
	
	lemma_para_list = make([]paraStruct, 0, len(righe)+4 )   
	var wP paraStruct
	//var pP paraStruct
	var pkeyL, keyL string
	var pkeyL2, keyL2 string
	var sumExample string
	var xLem, xPara, xExa string
	
	sk:=0
	sw1:=false
	//--------------
	for z1:=0; z1 < len(righe); z1++ {		
		col := strings.Split((righe[z1]+"||||"), "|") 
		xLem   = strings.TrimSpace( col[0] )
		if xLem == "" {continue}
		xPara    = strings.TrimSpace( col[1] ) 	
		xExa     = strings.TrimSpace( col[2] ) 
		keyL2 = xLem + "." +  xPara
		keyL  = xLem + "." +  xPara + "." + xExa
		
		if keyL == pkeyL { sk++; continue }
		if keyL2 != pkeyL2 {
			if pkeyL2 != "" {
				wP.p_example = sumExample
				lemma_para_list = append(lemma_para_list , wP ) 
				if ((wP.p_lemma == "gehen") || (wP.p_lemma == "familie"))  {fmt.Println("carica paradigma ", wP)  } 
			}	
			pkeyL2 = keyL2
			sumExample   = ""
			wP.p_lemma   = xLem
			wP.p_para    = xPara 	
			wP.p_example = "" 
		}
		sw1 =((xLem == "gehen") || (xLem == "familie")) 
		if sw1 { fmt.Println("legge  paradigma ", keyL)  } 
		pkeyL = keyL
		if len(sumExample) > 0 {  
			if len(xExa) > 0 {  
				if strings.Index(".!?", sumExample[ len(sumExample)-1:] ) < 0 {xExa = ". " + xExa } else {xExa = " " + xExa}	
			}
		}
		sumExample += xExa 
	} // end of for_z1	
	//-------
	wP.p_example = sumExample 
	lemma_para_list = append(lemma_para_list , wP ) 	
	//------------
	fmt.Println("    scartate ", sk, " righe doppie, ", len( lemma_para_list ), " righe caricate in lemma_para_list")
	
	sort_lemmaPara() 
	
	numFound:=0 
	notFound:=0
	swFound:=false	
	
	/**
	//-----------------------
	for f2, lex := range lemmaSlice {
		if lex.leLemma == "person" {
			fmt.Println("g25_2read_para... 1 lemma ", f2, " ==> ", lex.leLemma, " para=", lex.lePara, " ", lex)
		}	
	}
	**/
	//------------------------------------------------ 		
	for z2, wP2 := range lemma_para_list {	
		
		code1:= wP2.p_lemma
		ix1, ix2 := lookForLemma(code1,0)		
		swFound =false
		for z3:=ix1; z3 <=ix2; z3++ {
			LEM:= lemmaSlice[z3]					
			if LEM.leLemma != code1 { continue }
			swFound = true
		}
		if swFound == false {				
			lemmaCod :=  stdCode(code1) 
			if lemmaCod == code1 {lemmaCod =  std2Code(code1)  } 
			if lemmaCod != code1 {
				//fmt.Println(" LEMMA ", red(code1), " lemmaCod=", lemmaCod)
				wP2.p_lemma = lemmaCod	
				ix1, ix2 = lookForLemma(wP2.p_lemma,0)			
			}				
		}	
		for z3:=ix1; z3 <=ix2; z3++ {
			LEM:= lemmaSlice[z3]					
			if LEM.leLemma != wP2.p_lemma { continue }
			swFound = true
			lemma_para_list[z2].p_ixLemma = z3   // indice di lemmaSlice  
			LEM.leNumPara ++
			if LEM.leNumPara > 1 {   
				LEM.lePara    += "|" + wP2.p_para       
				LEM.leExample += "|" + wP2.p_example   
			} else {	      
				LEM.lePara    = wP2.p_para       
				LEM.leExample = wP2.p_example 
			}				
			lemmaSlice[z3] = LEM	
		}
		if swFound {
			numFound++	
		} else {
			notFound++				
			if notFound < 50 { 
				fmt.Println(code1 + "|" + code1 + "| |" + "lemma_para_list ", red("NOT FOUND in lemmaSlice" ), wP2.p_para)
			}
		}	
	}
	//-----------------------
	/**
	for f2, lex := range lemmaSlice {
		if lex.leLemma == "person" {
			fmt.Println("g25_2read_para... 2 lemma ", f2, " ==> ", lex.leLemma, " para=", lex.lePara, " ", lex)
		}	
	}
	**/
	//--------------
	fmt.Println("\t\t", numFound , " lemma_para_list FOUND in lemmaSlice, ",  notFound, " not found")       
	
} // end of read_ParadigmaFile

//-----------------------------------
func TOGLI2g25_2read_ParadigmaFile( path1 string, inpFile string) {
	bytesPerRow:= 40
    righe := rowListFromFile( path1, inpFile, "paradigma", "read_ParadigmaFile", bytesPerRow)  
	if sw_stop { return }
	
	/*
		  0  	| 1 |                 2              |                                  3                   
		ab  	| A2| abholen, holt ab, hat abgeholt |Wann kann ich die Sachen bei dir abholen<br>Wir müssen noch meinen Bruder abholen  | 
	    
		aber	| A1|aber	 |Der Film ist traurig, aber sehr schön.
		aber	| A2|aber    |Heute kann ich nicht kommen, aber morgen habe ich Zeit.<br>Wir haben nur eine kleine Wohnung, sind aber damit zufrieden.<br>Es war sehr schön.<br>Jetzt muss ich aber gehen.<br>Das ist aber nett von dir.
		aber	| B1|aber    |1. Heute kann ich nicht, aber morgen ganz bestimmt.<br>2. Es lag sehr viel Schnee, aber Enzo ist trotzdem mit dem Motorrad gefahren.<br>3. Wir haben nur eine kleine Wohnung, sind aber damit zufrieden.<br>4. Es war sehr schön.<br>Jetzt muss ich aber gehen.<br>5. Ich würde gerne kommen, aber es geht leider nicht.<br>6. Darf ich dich zu einem Kaffee einladen?<br>– Aber ja, sehr gern.<br>7. Du spielst aber gut Klavier. |                                                                  |    |  |        
	
		abfahren| A1|abfahren|Der Zug fährt gleich ab.
		abfahren| B1|abfahren, fährt ab, fuhr ab, ist abgefahren |Unser Zug ist pünktlich abgefahren.
	
	*/ 
	
	fmt.Println("\nletti paradigma file ", inpFile  + " " , len(righe) , " righe")   
	
	lemma_para_list = make([]paraStruct, 0, len(righe)+4 )   
	var wP paraStruct
	//var pP paraStruct
	var pkeyL, keyL string
	sk:=0
	//--------------
	for z1:=0; z1 < len(righe); z1++ {		
		col := strings.Split((righe[z1]+"||||"), "|") 
		wP.p_lemma   = strings.TrimSpace( col[0] )
		if wP.p_lemma == "" {continue}
		//wP.p_level   = strings.TrimSpace( col[1] ) 
		wP.p_para    = strings.TrimSpace( col[2] ) 	
		wP.p_example = strings.TrimSpace( col[3] ) 
		
		//keyL = wP.p_lemma + "." + wP.p_level + "." + wP.p_para + "." + wP.p_example
		keyL = wP.p_lemma + "." +  wP.p_para + "." + wP.p_example
		if ((wP.p_lemma == "gehen") || (wP.p_lemma == "familie")) { fmt.Println("carica paradigma ", keyL)  } 
		if keyL == pkeyL { sk++; continue }
		pkeyL = keyL

		
		//pP = wP ;  
		lemma_para_list = append(lemma_para_list , wP ) 	
		
	} // end of for_z1	
	//------------
	fmt.Println("    scartate ", sk, " righe doppie, ", len( lemma_para_list ), " righe caricate in lemma_para_list")
	
	sort_lemmaPara() 
	
	numFound:=0; 
	list_z3 := make([]int,0,20)
	for z2, wP2 := range lemma_para_list {			
		ix1, ix2 := lookForLemma(wP2.p_lemma,0)
		for z3:=ix1; z3 <=ix2; z3++ {
			LEM:= lemmaSlice[z3]
			if findIndex(list_z3, z3) < 0 {	list_z3 = append(list_z3,z3) }
			//if z2 < 10 {fmt.Println("     z3=", z3, " LEM.leLemma=", LEM.leLemma ) }
			
			if LEM.leLemma == wP2.p_lemma {
				lemma_para_list[z2].p_ixLemma = z3   // indice di lemmaSlice  
				LEM.leNumPara ++
				if LEM.leNumPara > 1 {
					//LEM.leLevel   += "|" + wP2.p_level      
					LEM.lePara    += "|" + wP2.p_para       
					LEM.leExample += "|" + wP2.p_example   
				} else {	
					//LEM.leLevel   += wP2.p_level       
					LEM.lePara    += wP2.p_para       
					LEM.leExample += wP2.p_example 
				}				
				lemmaSlice[z3] = LEM	
				numFound++	
				//if numFound < 10 { fmt.Println(" lemma_para_list ", wP2 , " \t XXX lemmaSlice[",z3,"] = ", lemmaSlice[z3] )   	}
			}
		} 
	}
	//--------------
	
	
	fmt.Println("\t\t", numFound , " lemma_para_list FOUND in lemmaSlice     ",  (len( lemma_para_list ) - numFound), " not found")       
	
} // end of TOGLI2read_ParadigmaFile

//-------------------
func TOGLIread_ParadigmaFile( path1 string, inpFile string) {
	bytesPerRow:= 40
    righe := rowListFromFile( path1, inpFile, "paradigma", "read_ParadigmaFile", bytesPerRow)  
	if sw_stop { return }
	
	/*
		  0  	| 1 |                 2              |                                  3                   
		ab  	| A2| abholen, holt ab, hat abgeholt |Wann kann ich die Sachen bei dir abholen<br>Wir müssen noch meinen Bruder abholen  | 
	    
		aber	| A1|aber	 |Der Film ist traurig, aber sehr schön.
		aber	| A2|aber    |Heute kann ich nicht kommen, aber morgen habe ich Zeit.<br>Wir haben nur eine kleine Wohnung, sind aber damit zufrieden.<br>Es war sehr schön.<br>Jetzt muss ich aber gehen.<br>Das ist aber nett von dir.
		aber	| B1|aber    |1. Heute kann ich nicht, aber morgen ganz bestimmt.<br>2. Es lag sehr viel Schnee, aber Enzo ist trotzdem mit dem Motorrad gefahren.<br>3. Wir haben nur eine kleine Wohnung, sind aber damit zufrieden.<br>4. Es war sehr schön.<br>Jetzt muss ich aber gehen.<br>5. Ich würde gerne kommen, aber es geht leider nicht.<br>6. Darf ich dich zu einem Kaffee einladen?<br>– Aber ja, sehr gern.<br>7. Du spielst aber gut Klavier. |                                                                  |    |  |        
	
		abfahren| A1|abfahren|Der Zug fährt gleich ab.
		abfahren| B1|abfahren, fährt ab, fuhr ab, ist abgefahren |Unser Zug ist pünktlich abgefahren.
	
	*/ 
	
	fmt.Println("\nletti paradigma file ", inpFile  + " " , len(righe) , " righe")   
	
	lemma_para_list = make([]paraStruct, 0, len(righe)+4 )   
	var wP paraStruct
	//var pP paraStruct
	var pkeyL, keyL string
	sk:=0
	//--------------
	for z1:=0; z1 < len(righe); z1++ {		
		col := strings.Split((righe[z1]+"||||"), "|") 
		wP.p_lemma   = strings.TrimSpace( col[0] )
		if wP.p_lemma == "" {continue}
		//wP.p_level   = strings.TrimSpace( col[1] ) 
		wP.p_para    = strings.TrimSpace( col[2] ) 	
		wP.p_example = strings.TrimSpace( col[3] ) 
		
		//keyL = wP.p_lemma + "." + wP.p_level + "." + wP.p_para + "." + wP.p_example
		keyL = wP.p_lemma + "." +  wP.p_para + "." + wP.p_example
		if keyL == pkeyL { sk++; continue }
		pkeyL = keyL

		//level1 := " " +  wP.p_level + " "
		//if strings.Index(list_level_str, level1) < 0 { list_level_str += level1 } 	
		/**
		if wP.p_para != "" {
			//    0 = x48, A = x65, Z =x90,  a = x97
			ch1 := wP.p_para[0:1]
			if ch1 < "a" { 
				if ch1 < "A" || ch1 > "Z" {
					wP.p_para = fseq + wP.p_para  // per la chiave di sort,  serve per spostare la riga alla fine  se il paradigma inizia con ( [ o altro 
				}	 
			}
		} 
		**/
		
		
		//pP = wP ;  
		lemma_para_list = append(lemma_para_list , wP ) 	
		
	} // end of for_z1	
	//------------
	fmt.Println("    scartate ", sk, " righe doppie, ", len( lemma_para_list ), " righe caricate in lemma_para_list")
	
	/**
	list_level = strings.Split(strings.TrimSpace(list_level_str), " ") 
	
	//fmt.Println("XXX livelli: string=>" + list_level_str + "<== \nlivelli=", list_level) 
	
	only_level_numWords = make([]int, len(list_level), len(list_level) )
	perc_level          = make([]int, len(list_level), len(list_level) )
	**/
	sort_lemmaPara() 
	
	numFound:=0; 
	list_z3 := make([]int,0,20)
	for z2, wP2 := range lemma_para_list {		
		//if z2 < 10 {fmt.Println("XXX PARADIGMA: ", z2, " ",  wP2, "   wP2.p_lemma=" +wP2.p_lemma  )   }
		
		ix1, ix2 := lookForLemma(wP2.p_lemma,0)
		
		//if z2 < 10 {fmt.Println("     ix1=", ix1,  " ix2=", ix2) }
		
		for z3:=ix1; z3 <=ix2; z3++ {
			LEM:= lemmaSlice[z3]
			if findIndex(list_z3, z3) < 0 {	list_z3 = append(list_z3,z3) }
			//if z2 < 10 {fmt.Println("     z3=", z3, " LEM.leLemma=", LEM.leLemma ) }
			
			if LEM.leLemma == wP2.p_lemma {
				lemma_para_list[z2].p_ixLemma = z3   // indice di lemmaSlice  
				LEM.leNumPara ++
				if LEM.leNumPara > 1 {
					//LEM.leLevel   += "|" + wP2.p_level      
					LEM.lePara    += "|" + wP2.p_para       
					LEM.leExample += "|" + wP2.p_example   
				} else {	
					//LEM.leLevel   += wP2.p_level       
					LEM.lePara    += wP2.p_para       
					LEM.leExample += wP2.p_example 
				}				
				lemmaSlice[z3] = LEM	
				numFound++	
				//if numFound < 10 { fmt.Println(" lemma_para_list ", wP2 , " \t XXX lemmaSlice[",z3,"] = ", lemmaSlice[z3] )   	}
			}
		} 
	}
	//--------------
	
	
	fmt.Println("\t\t", numFound , " lemma_para_list FOUND in lemmaSlice     ",  (len( lemma_para_list ) - numFound), " not found")       
	
} // end of TOGLIread_ParadigmaFile

//-----------------------------------------------

func sort_lemmaPara() {

	sort.Slice(lemma_para_list, func(i, j int) bool {
		if lemma_para_list[i].p_lemma != lemma_para_list[j].p_lemma {
			return lemma_para_list[i].p_lemma < lemma_para_list[j].p_lemma         // lemma ascending order (eg.   a before b ) 
		} else {
			return lemma_para_list[i].p_para < lemma_para_list[j].p_para       // level ascending order (eg.   a  before b ) 			
		}	
	} )	
	/**
	sort.Slice(lemma_para_list, func(i, j int) bool {
		if lemma_para_list[i].p_lemma != lemma_para_list[j].p_lemma {
			return lemma_para_list[i].p_lemma < lemma_para_list[j].p_lemma         // lemma ascending order (eg.   a before b ) 
		} else {
			if lemma_para_list[i].p_level != lemma_para_list[j].p_level {
				return lemma_para_list[i].p_level < lemma_para_list[j].p_level     // level ascending order (eg.   A1 before A2 ) 
			} else {
				return lemma_para_list[i].p_para < lemma_para_list[j].p_para       // level ascending order (eg.   a  before b ) 
			}
		}	
	} )	
	**/
} // end of sort_lemmaPara() 

//-----------------------------------------------
