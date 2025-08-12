package wbfSubPack

import (  
	"fmt"
    "strings"
	"strconv"	
)

//------------------------------
func g32_buildStatistics() {		
		//fmt.Println(cyan("buildStatistics"))
		//var rows []string
		var result string = ""
		 g32_build_stat() 
		//if len( only_level_numWords ) < 1 { return }
		
		/***	
		msgLevelStat = "" 		
		if percA0 > 0 {msgLevelStat += ", A0: " + strconv.Itoa(percA0) + "%" }
		if percA1 > 0 {msgLevelStat += ", A1: " + strconv.Itoa(percA1) + "%" }
		if percA2 > 0 {msgLevelStat += ", A2: " + strconv.Itoa(percA2) + "%" }
		if percB1 > 0 {msgLevelStat += ", B1: " + strconv.Itoa(percB1) + "%" }
		if percOth > 0 {msgLevelStat += ", Oth: " + strconv.Itoa(percOth) + "%" }
		if len(msgLevelStat) > 1 {msgLevelStat = msgLevelStat[2:] } 
		**/
		
		msgLevelStat = "" 
		/**
		for f:=1; f < len( only_level_numWords ) ; f++ {
			//if only_level_numWords[f] == 0 { continue }
			if perc_level[f] == 0 { continue }
			msgLevelStat += ", " + list_level[f] + ": " + strconv.Itoa( perc_level[f] ) + "%"  
		}	
		if only_level_numWords[0] > 0 {  
			msgLevelStat += ", " + list_level[0] + ": " + strconv.Itoa( perc_level[0] ) + "%"  
		}
		**/
		if len(msgLevelStat) > 1 {msgLevelStat = msgLevelStat[2:] } 

		result += "livello " + msgLevelStat //  + "..endLevel ";  
		
		//fmt.Println("statistcs len(wordStatistic_tx)=", len(wordStatistic_tx) )
		
		for _, sS:= range wordStatistic_tx {	
			if sS.totWords == 0 { continue; }
			//if sS.uniqueWords < 100 { continue}
			//fmt.Println( sS.uniqueWords , " words (",  sS.uniquePerc, "%), found ", 
			//	sS.totWords,  " times in the text(", sS.totPerc,"%)" ) 
			
			//result += "<br>" + fmt.Sprintln( sS.uniqueWords , " words (",  
			//	sS.uniquePerc, "%), make up ", sS.totPerc,"% of the text (", sS.totWords, " words)") 
			result += "<br>" + fmt.Sprintln( sS.uniqueWords, ",", sS.uniquePerc, ",", sS.totPerc, ",", sS.totWords) 	
		}  		
		result += "<br>" 
		
		//fmt.Println("statistics ", result); 
		
		go_exec_js_function("js_go_updateStatistics", result )		
	
}	

//----------------------------------------
/**
func stat_level( lemmaLevel []string, numWords int) {	
	
	// get the first level of the first lemma 
	
	if len(lemmaLevel) < 1 { return }
	if numWords < 1 { return }
	
	level2 := strings.Split( lemmaLevel[0], "|" ) 
	if len(level2) < 1 { return }
	
	levelToText := level2[0]
	
	sw_oth:=true
	for m:=0; m < len(list_level); m++ {	
		if levelToText == list_level[m] {
			only_level_numWords[m] += numWords 
			sw_oth = false; 
			break
		} 
	} 
	if sw_oth {
		only_level_numWords[0] += numWords 
	}
	
}
***/
//-----------------------------------

func stat_useWord() {
	len1:=  len(uniqueWordByFreq)
	len2:= float64(len1)/100
	
	fmt.Println("len1=", len1, " ", len2) 	

	lisPerc := [29]float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9,1,2,3,4,5,6,7,8,9, 10,20,30,40,50,60,70,80,90,100}
	listIxPerc:= make([]int,0,40)
	for z:=0; z < len(lisPerc); z++ {
		per1 := lisPerc[z]
		per2 := int( float64(per1) * len2)
		listIxPerc = append( listIxPerc, per2 )
		//fmt.Println("stats ", per1, "% = num.Elem.",  per2 )  	
	}   
	
	lastTot:=0;
	ixNow:=0	
	lenU := len(uniqueWordByFreq)
	if lenU < 1 {return } 
	
	for z:=0; z < len(listIxPerc); z++ {
		//from1 = ixNow
		ixNow =  listIxPerc[z]-1
		if ((ixNow< 0) || (ixNow >= lenU)) { continue }
		
		if uniqueWordByFreq[ixNow].fuTotRow == lastTot { continue }
		
		//fmt.Println("stats ", lisPerc[z], "% = num.Elem.",  listIxPerc[z], " toIx=", ixNow,   
		//					" num.Rows per word=",uniqueWordByFreq[ixNow].totRow )
		if listIxPerc[z] >= 1 {	
			fmt.Println( "stats ",  prtFloat( lisPerc[z] , 5,1 ) ,"% = num.Elem.",prtInt( listIxPerc[z] , 5 )," sono usate ", prtInt( uniqueWordByFreq[ixNow].fuTotRow, 5 ), 
					" o più volte  (",  prtFloat( lisPerc[z] , 5,1 ),"% delle parole non sono usate più di ", prtInt(  uniqueWordByFreq[ixNow].fuTotRow, 5 ), " volte)")					
		}
		lastTot = uniqueWordByFreq[ixNow].fuTotRow 	
	} 
	
} // end of stat_useWord

//------------------------------------

func prtFloat( input float64, maxL int, dec1 int ) string {
	//  space character after % is the padding character which will be repeated by the value replacing the first *   
    //  first  * character is replaced by the difference between the maximum length and the actual length of the number converted to string 
    //  second * character is replaced by dec1 value ( how many decimal)    	
	return fmt.Sprintf("% *s%.*f", maxL-len( strconv.FormatFloat(input, 'f', 2, 64)), "", dec1,  input )
} 
//--------------------------
func prtFloat1( input float64, maxL int ) string {
  return fmt.Sprintf("% *s%.1f", maxL-len( strconv.FormatFloat(input, 'f', 2, 64)), "",  input )
} 
//---
func prtInt( input int, maxL int ) string {
  return fmt.Sprintf("% *s%d", maxL-len( strconv.Itoa(input)), "",  input )
}  
//-----------
/***
var wordSliceAlpha = make([]wordStruct, 0, 0)  
var uniqueWordByFreq  []wordUnFreqStruct;   // elenco delle parole in ordine di frequenza
var uniqueWordByAlpha []wordUnAlphaStruct;    // elenco delle parole in ordine alphabetico

//--
type wordStruct struct {       // a word is repeated several time one for each row containing it  
	wWordSeq  string
    wWord2    string
	wIxThisWord_al int 
	wIxUniq_al   int               // index of uniqueWordByFreq 	
	wIxUniq_fr   int               // index of uniqueWordByFreq 	
	wNfile    int 
	wSwSelRowG int
	wSwSelRowR int
    wIxRow    int       
	wIxPosRow int 
	wListPref string 
	wTotRow   int              // number of rows 
	wTotExtrRow int            // number of extracted rows 
	wTotMinRow int
	wTotWrdRow int 	
	wIxLemmaList []int
}
//--

//--
type wordUnAlphaStruct struct {    // uniqueWordByAlpha
	uWordSeq    string	
    uWord2      string		
	uIxUnW_al   int            // index of this word in the uniqueWordByAlpha 	
	uIxUnW_fr   int            // index of this word in the uniqueWordByFreq	
	uTotRow     int 
	uTotExtrRow int
	uIxFromWord_al int          // index of this word in the wordSliceAlpha (first occurrence, last = uIxFromWordAl + uTotRow-1	
    //uIxWordFreq int            // index of this word in the wordSliceFreq	
	uSwSelRowG  int
	uSwSelRowR   int  
	uLearnedYN   string         // y n ( ie.yes,I learned / not yet  
	//uKnow_yes_ctr int 
	//uKnow_no_ctr  int         // a value > 0  means that this is a word that I don't know, ie. it's to be learned   
	uIxLemmaL  []int  
	uLemmaL    []string       // list of lemma 	
	//uPara      []string  
	//uExample   []string  
}	

//--
type wordUnFreqStruct struct {    // uniqueWordByFreq
	fuWordSeq    string	
    fuWord2      string		
	fuIxUnW_al   int            // index of this word in the uniqueWordByAlpha 	
	fuIxUnW_fr   int            // index of this word in the uniqueWordByFreq	
	fuTotRow     int 
}	



***/
//--------------------------------------
func g32_build_stat() {
	fmt.Println("\n", green("STATISTICHE build_uniqueWord_byFreqAlpha"),  " crea uniqueWordByFreq" )
		
	//put_a_priority_to_the_row_of_each_word() 	
	
	fmt.Println(" len(wordSliceAlpha)   = ", len(wordSliceAlpha)    )
	fmt.Println(" len(uniqueWordByFreq) = ", len(uniqueWordByFreq)  )
	fmt.Println(" len(uniqueWordByAlpha)= ", len(uniqueWordByAlpha) )
	
	preW := ""
	numWordUn := 0
	numWordRi := 0	
	num_word:=0
	
	numWordUn_0 := 0
	numWordRi_0 := 0	
	num_word_0 :=0 
	//--------------------		
	for _, wS1 := range wordSliceAlpha {	
		if strings.Index(wS1.wWordSeq, "...") > 0 { continue }
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
	
	
	var sS  statStruct;
	
	//numWordUn = -1
	numWordUn = 0
	
	//---------
	const LAST_WORD_FREQ = 999999999 
	//---------------------------
	
	numWordUn = 0 
	numWordRi = 0
	//-------------------
	fmt.Println("numberOfWords=",numberOfWords, "  numberOfUniqueWords=",  	numberOfUniqueWords)
	
	for _, frW:= range uniqueWordByFreq{
		if strings.Index(frW.fuWord2,"...") > 0  { continue }  
		/**
		fuWordSeq    string	
		fuWord2      string		
		fuIxUnW_al   int            // index of this word in the uniqueWordByAlpha 	
		fuIxUnW_fr   int            // index of this word in the uniqueWordByFreq	
		fuTotRow     int 
		**/
		ixA:= frW.fuIxUnW_al
		alW:= uniqueWordByAlpha[ixA]
		numWordUn += 1
		numWordRi += alW.uTotRow
		percIx = int(numWordUn * 100 / numWordUn_0); 
		if percIx < 1 {continue}	
		sS.uniqueWords = numWordUn 		
		sS.totWords    = numWordRi
		sS.uniquePerc  = percIx 
		sS.totPerc     = int(numWordRi * 100 / numWordRi_0);
		fmt.Println("statistic ", "sS.totWords=", sS.totWords, " numWordRi=",numWordRi, "  sS.totPerc=",  sS.totPerc )
			
		if sS.totPerc <= 200 {   // esistono perc > 100%,  probabilmente c'è un errore di logica 
			/**
			if sS.totPerc > 90 {
				wordStatistic_tx[sS.totPerc] = sS; 
				continue
			}
			**/
			if strconv.Itoa(1000 + sS.totPerc)[3:] == "0" {				
				wordStatistic_tx[sS.totPerc] = sS; 
			}
		}
	}


	
	

} // end of build_stat 	
//----------------------------
