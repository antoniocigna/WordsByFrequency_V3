package wbfSubPack

import (  
	"fmt"
    "strings"
	"strconv"	
)

//------------------------------
func buildStatistics() {		
		//var rows []string
		var result string = ""
		
		if len( only_level_numWords ) < 1 { return }
		
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
		for f:=1; f < len( only_level_numWords ) ; f++ {
			//if only_level_numWords[f] == 0 { continue }
			if perc_level[f] == 0 { continue }
			msgLevelStat += ", " + list_level[f] + ": " + strconv.Itoa( perc_level[f] ) + "%"  
		}	
		if only_level_numWords[0] > 0 {  
			msgLevelStat += ", " + list_level[0] + ": " + strconv.Itoa( perc_level[0] ) + "%"  
		}
		if len(msgLevelStat) > 1 {msgLevelStat = msgLevelStat[2:] } 

		result += "livello " + msgLevelStat //  + "..endLevel ";  
		
		for _, sS:= range wordStatistic_tx {	
			if sS.totWords == 0 { continue; }
			if sS.uniqueWords < 100 { continue}
			//fmt.Println( sS.uniqueWords , " words (",  sS.uniquePerc, "%), found ", 
			//	sS.totWords,  " times in the text(", sS.totPerc,"%)" ) 
			
			//result += "<br>" + fmt.Sprintln( sS.uniqueWords , " words (",  
			//	sS.uniquePerc, "%), make up ", sS.totPerc,"% of the text (", sS.totWords, " words)") 
			result += "<br>" + fmt.Sprintln( sS.uniqueWords, ",", sS.uniquePerc, ",", sS.totPerc, ",", sS.totWords) 	
		}  		
		result += "<br>" 
		go_exec_js_function("js_go_updateStatistics", result )		
	
}	

//----------------------------------------

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
	for z:=0; z < len(listIxPerc); z++ {
		//from1 = ixNow
		ixNow =  listIxPerc[z]-1
		if ixNow< 0 { ixNow=0;}
		
		if uniqueWordByFreq[ixNow].uTotRow == lastTot { continue }
		
		//fmt.Println("stats ", lisPerc[z], "% = num.Elem.",  listIxPerc[z], " toIx=", ixNow,   
		//					" num.Rows per word=",uniqueWordByFreq[ixNow].totRow )
		if listIxPerc[z] >= 1 {	
			fmt.Println( "stats ",  prtFloat( lisPerc[z] , 5,1 ) ,"% = num.Elem.",prtInt( listIxPerc[z] , 5 )," sono usate ", prtInt( uniqueWordByFreq[ixNow].uTotRow, 5 ), 
					" o più volte  (",  prtFloat( lisPerc[z] , 5,1 ),"% delle parole non sono usate più di ", prtInt(  uniqueWordByFreq[ixNow].uTotRow, 5 ), " volte)")					
		}
		lastTot = uniqueWordByFreq[ixNow].uTotRow 	
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
