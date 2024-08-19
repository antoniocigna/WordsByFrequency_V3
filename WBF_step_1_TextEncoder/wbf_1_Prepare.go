
package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"		
    "io"	
)
//-------------------------------------------
var separRow   = "[\r\n.;:?!]";
//var separRowFalse = "[\r\n]";
var separWord = "["     +
			"\r\n.;?!" + 
			"\t"       +
			" "        + 
			",:"       +
            "|"        +
			"¿"        +	
			"°"        + 
			"¡"        + 
			"\\"       + 
			"_"        +
			"\\+"      +
			"\\*"      +
			"()<>"     + 
			"\\]\\["   +
			"`{}«»‘’‚‛“””„‟" +		
			"\""       + 
			"'"        + 	
			"\\/"      +  
			wSep       + 
			"]" ; 	
var separRowList = make([]string,0,0)		

//=============================================
//----------
var pref          = ""
var fromNum int   = 0
var outFileName   = ""
var inputFileName = ""
var sw_ignore_newLine = false 
var sw_nl_only        = false  

//--------------

var sw_stop       = false
//---------
const wSep = "§";                         // used to separe word in a list 
const endOfLine = ";;\n"
var newRows = make([]string, 0, 0)
//-----------------------

//============================================
func main() {
	
	getPgmArgs()	
	
	//fmt.Println( "p=",pref, " from", fromNum, " oFn=", outFileName, "  Ifn=",inputFileName)
	
	read_main_text_file()
	
	
} // end of main
//---------------------------------------------------

//-------------------------------
func getPgmArgs() {  
	
	args    :=  os.Args
	
	
	fromNumS := "0"
	sw_nl_ignoreS := "" 
	sw_nl_onlyS   := "" 
	
	for a:=1; a < len(args)-1; a++  {
		
		//fmt.Println("a=",a, "==>" + args[a] + "<==")
		 
		switch args[a] {
			case "-pref"        : pref          = args[a+1] 
			case "-fromNum"     : fromNumS      = args[a+1]
			case "-outputFile"  : outFileName   = args[a+1]
			case "-inputFile"   : inputFileName = args[a+1]
			case "-sw_nl_ignore": sw_nl_ignoreS = args[a+1] 
			case "-sw_nl_only"  : sw_nl_onlyS   = args[a+1] 
		}
	} 
	fromNum, _ = strconv.Atoi( strings.TrimSpace(fromNumS) )
	if ((sw_nl_ignoreS == "not") || (sw_nl_ignoreS == "no") || (sw_nl_ignoreS == "false") ) {
		sw_ignore_newLine = false 
	} else {
		sw_ignore_newLine = true 
	}
	if ((sw_nl_onlyS == "not") || (sw_nl_onlyS == "no") || (sw_nl_onlyS == "false") ) {
		sw_nl_only = false 
	} else {
		sw_nl_only = true 
	}	
	fmt.Println("\n-----------------------------------") 
	fmt.Println("inputFile     = " , inputFileName)
	fmt.Println("") 
	fmt.Println("outputFile    = " , outFileName )
	fmt.Println("pref          = " , pref    ) 
	fmt.Println("fromNum       = " , fromNum )
	
	fmt.Println("sw_nl_ignore  = " , sw_ignore_newLine , "\n\t\t( true => le righe di input vengono accodate in una sola riga " +
			 "\n\t\te spezzate al punto (punto fermo, interrogativo, esclamativo) ") 
	fmt.Println("sw_nl_only    = " , sw_nl_only        , " \t( true => le righe di input non vengono spezzate)")
	fmt.Println("-----------------------------------\n") 
	
	
} // end of getPgmArgs

//------------------
func myOpenRead( path1 string,   fileName string,   descr string,  func1    string) (*os.File, int) {
	path2:="";
	path10:=""
	if path1 != "" {
		path10 = " in " + path1
		path2 = path1 + string(os.PathSeparator) 
	} 	
	fileN := path2 + fileName 
	sizeByte:=0
	fmt.Println("\n-------------------------------", "\nopen file ",  fileName, path10 )
	
	readFile, err := os.Open( fileN )  
    if err == nil {			
		fileInfo, _ := os.Stat( fileN )  
		sizeByte = int( fileInfo.Size() )
		fmt.Println( "\t", "size: ", fileInfo.Size(), " bytes" )	
		return readFile, sizeByte
	}
	msg1:= `il file "` + fileN + `" (` + descr + " " + func1 + ")" + " non esiste"
	
	fmt.Println(msg1, " func ", func1)	
	
	return readFile, 0		
	
} // end of myOpenRead


//----------------------------------

func rowListFromFile( path1 string, fileName string, descr string, func1 string, bytesPerRow int ) []string { 
	
	file, sizeByte := myOpenRead( path1, fileName, descr, func1 )  
	if file == nil {			
		sw_stop = true 
		return nil
	} 
	numEleMax:= int( sizeByte / bytesPerRow ); 
	if numEleMax < 10 {numEleMax=10}
	
	//fmt.Println("    allocate for a maximum of ", numEleMax, " rows (assumed ", bytesPerRow, " bytes per row as average)" )
	
	retRowL := make( [] string, 0, numEleMax)	
	
	r := bufio.NewReader(file)
	for {
	  line, _, err := r.ReadLine()
	  if err != nil {
		if err == io.EOF {
			break
		}
		break
	  }	 
	  retRowL = append( retRowL, string(line) ) 	
	}
	defer file.Close()
	
	//fmt.Println("letto file " , fileName, "  num lines=", len(retRowL) )
	fmt.Println("letti ", len(retRowL), " righe")
	
	return retRowL 
	
} // end of rowListFromFile	

//-------------------------------------------------

//----------------------------------------------
func read_main_text_file()  { 
	
	fn:=  strings.ReplaceAll(inputFileName,"\\","/")
	if fn == "" {
		fmt.Println("il file input non è stato specificato  (usa parametro -inputFile )" )
		return 
	}
	 
	v2:=0
	for v0:=len(fn)-1; v0 > 0; v0-- {
		if fn[v0] == '/' {
			v2=v0; 
			break;
		}
	}  
	path1:=    fn[0:v2];
	fileName:= fn[v2+1:]	
	
	//fmt.Println("read_main_text_file  main_input_text_file=" + inputFileName, "\nfn=" + fn , " v2=", v2, " path1=", path1, " fileName=", fileName) 	
	
    //--------------------------
    // leggi file di testo 
    //-------------------------	
    
	
	bytesPerRow:= 10
    righeInp := rowListFromFile( path1, fileName, "main text file", "read_main_text_file", bytesPerRow)  
	if sw_stop { return }
	 
	numRigheInp:= len(righeInp)

	newRows = make([]string, 0, 2*numRigheInp)

	//---------------------
	
	line:=""
	//---------------------------------------------
	
	/**
	for _, line0:= range( righeInp ) {
		fmt.Println("input ", line0)
	}
	fmt.Println("--------------------------------\n")
	**/
	
	for _, line0:= range( righeInp ) {
	
		line = strings.ReplaceAll( string(line0), "<br>", "\n");
		line = strings.ReplaceAll( line, "\r\n", "\n") 
		line = strings.ReplaceAll( line, "\r"  , "\n") 
		
		//--------------------------------------
		//  sw_ignore_newLine = true  ,   sw_nl_only = true      ==>  incompatibili
		//  sw_ignore_newLine = true  ,   sw_nl_only = false     ==>  split .?!;  
		//  sw_ignore_newLine = false ,   sw_nl_only = false     ==>  split .?!;  and  new line  
		//  sw_ignore_newLine = false ,   sw_nl_only = true      ==>  only new line 
		//------------------------------------------------------------------------------		
		if sw_nl_only {                      // append read line 
			newRows = append(newRows, line )			
			continue
		}
		//--------------------------------------
		
		split1:= splitHoldSep( string(line) )  //split the line according to several separators (eg.  .;?! )  
		//--------------------------------------
		if sw_ignore_newLine {                 //  sw_ignore_newLine = true  ,   sw_nl_only = false     ==>  split .?!;   											
			for c1, row1 := range split1 {				
				if c1 == 0  {				
					newRows[ len(newRows)-1] += " " + strings.TrimSpace(row1)  // append text to the previous entry 
				} else {
					newRows = append(newRows, row1 )			
				}	
			}	
			continue
		} 
		//---------------------------------------
		//  sw_ignore_newLine = false ,   sw_nl_only = false     ==>  split .?!;  and  new line   											
		for _, row1 := range split1 {
			//fmt.Println("for  n1 ", n1, " " , row1)
			newRows = append(newRows,  row1) 	
		}	
		
    } // end of for line0 range
	
	//--------------------------
	
	
	removeWordContinuation()
	
	numberOfRows := len(newRows)
	
	fmt.Println( "\n", numRigheInp, " rows read from text file ", fileName , "\n\t", numberOfRows, " rows" )
  
    if numberOfRows < 1 {
		msg1:= "il file " + fileName  + " not contiene nessuna riga valida" 	
		func1:= "read_main_text_file"			
		fmt.Println(msg1, " func ", func1)			
		sw_stop = true 
		return	 
	}
	//------------------
	if sw_stop { return;}
	
	
	writeTextRowSlice()  
	
	return;  
	
} // end of read_main_text_file
 
//----------------------------------------------


//----------------------------------------------

func buildSeparRowlist() {
	separRow1 := separRow
	if separRow1[:1] == "[" { separRow1 = separRow1[1:] }
	if separRow1[ len(separRow1)-1:] == "]" { separRow1 = separRow1[0: len(separRow1)-1] }
	
	//fmt.Println("separ =", separRow )
	//fmt.Println("separ1= ", separRow1)
	
	var sepStr=""
	for _, chr1 := range separRow1 {
		sepStr = sepStr+ string(chr1)
		separRowList = append(separRowList, string(chr1) ) 
	}	
	fmt.Println("\n separRowList="  + sepStr)
	
}	
//---------------------------	
/**
eliminare . . . . .    oppure ...  prima di spezzare le righe  ==> replace ". ."  con ""
eliminare ... e .. 
eliminare -- e __
eliminare ==

eliminare " . " 
------------------
trasforma numero seguito da punto e spazio, in numero°  es.    "2. " --> "2° " ,  "12. " in "12° "
	o meglio  numero seguito da punto e spazio non provoca un a capo 
==>  numero seguito da punto non provoca un newLine 
idem per numeri romani:   ultima lettera prima del punto è maiscola ed è una di queste : I V X L C D M  

------------------------------------
- conserva il newLine esistente     
- le parole cone terminano la riga con "-", continuano con la riga successiva. es.  fine riga "... mög-", inizio nuova riga "lich altro ..." --> diventa  fine riga "... möglich" inizio successiva "altro ..."  
--------------
**/
//--------------------------------
func splitHoldSep( str1 string) []string {
	/*
	split string according to several separator characters
	keeping them in the text  	
	*/
	//fmt.Println( "\nsplitHoldSep( str1=" + str1) 
	if len(separRowList)== 0 { buildSeparRowlist() }
	
	newY  := str1  + " "
	
	//fmt.Println("splitHoldSep() str1=" + str1 + "<==")
	
	//-------------------------------------
	// manage fullstop exception ( mask fullstop if it's after a number)  
	fullStopReplace := "°";
	swChg := false 
	
	if (strings.Index( newY, ".") > 0) {
	
		if (strings.Index( newY, "°") > 0) { newY = strings.ReplaceAll(newY,fullStopReplace," ") }
		
		newY = check_BC(newY)  // check Before Christ 
		newY = check_manyDots(newY)  
		if (strings.Index( newY, "//") > 0) {  // eg https://abc.org  // http://abc.org,  i's better to leave the line untouched  (there might be ? ! etc.)  
			//newY = strings.ReplaceAll(newY,".", fullStopReplace) 
			split1:=  []string{ newY } // it must return an array     
			return split1
		}	

		romanNumber := "IVXLCDM"   // must be uppercase 
		for a:=0; a < len(str1); a++ {
			ix1:= strings.Index( newY, ".")
			if (ix1 < 0) { break }
			if (ix1 < 1) { 
				newY = strings.ReplaceAll(newY,".",fullStopReplace); 
				continue 
			}
			pre:= newY[ix1-1:ix1];
			if ((pre >= "0") && (pre <= "9")) {  
				// if before fullstop there is a number, the line might be the member of a list or the number must be read as an ordinal number (eg. 1. == the first, 2. the second)  
				//  number.  eg. 2.  replaced by 2°    ( to avoid to force newline after fullstop)
  				newY = strings.ReplaceAll(newY,".",fullStopReplace);
			} 		
			if ( strings.Index(romanNumber, pre) >=0) {
				if (newY[ix1+1:ix1+2] == " ") {
					// I suppose it's a roman numbered list member    
					newY = strings.ReplaceAll(newY,".",fullStopReplace);
				}
			}	
		}
	}  
	swChg = (strings.Index( newY, fullStopReplace) >=0 )
	//--------------------------------------------------
	// split according to fullstops and other separators 
	split1:= []string{}
	
	//fmt.Println("  split .. newY=" + newY + "<==")
	
	for  _, sep1:= range separRowList {    
		split1 = strings.SplitAfter(newY, sep1)
		newY   = strings.Join(split1[:], "§")
		//fmt.Println("  split .. newY=" + newY + "<==")
	}
	//------
	if (swChg)  {
		newY = strings.ReplaceAll(newY,fullStopReplace, "."); // replace masked fullstop 
	}
	
	newY = strings.TrimRight( strings.TrimSpace(newY) , "§") 
	
	//fmt.Println("  split .. prima di split newY=" + newY + "<==")
	
	return strings.Split(   strings.TrimRight(newY,"§") , "§")	
	
} // end of  splitHoldSep	
//---------------
func check_BC(str1 string) string {
	//fmt.Println("  check_BC  input ==>" + str1 + "<==") 
	
	if (strings.Index(str1, " v.Chr." )>=0) { return strings.ReplaceAll(str1, " v.Chr."   , " v°Chr°" ) }   // German
	if (strings.Index(str1, " v. Chr.")>=0) { return strings.ReplaceAll(str1, " v. Chr."  , " v° Chr°") }
	if (strings.Index(str1, " a.C."   )>=0) { return strings.ReplaceAll(str1, " a.C."  , " a°C°"  ) }         // Italian   
	if (strings.Index(str1, " a. C."  )>=0) { return strings.ReplaceAll(str1, " a. C." , " a° C°" ) }
	if (strings.Index(str1, " A.C."   )>=0) { return strings.ReplaceAll(str1, " A.C."  , " A°C°"  ) }
	if (strings.Index(str1, " A. C."  )>=0) { return strings.ReplaceAll(str1, " A. C." , " A° C°" ) }
	if (strings.Index(str1, " B.C."   )>=0) { return strings.ReplaceAll(str1, " B.C."  , " B°C°"  ) }         // English 
	if (strings.Index(str1, " B. C."  )>=0) { return strings.ReplaceAll(str1, " B. C." , " B° C°" ) }
	
	if (strings.Index(str1, " v.Chr" )>=0) { return strings.ReplaceAll(str1, " v.Chr"   , " v°Chr" ) }   // German
	if (strings.Index(str1, " v. Chr")>=0) { return strings.ReplaceAll(str1, " v. Chr"  , " v° Chr") }
	if (strings.Index(str1, " a.C"   )>=0) { return strings.ReplaceAll(str1, " a.C"  , " a°C"  ) }         // Italian   
	if (strings.Index(str1, " a. C"  )>=0) { return strings.ReplaceAll(str1, " a. C" , " a° C" ) }
	if (strings.Index(str1, " A.C"   )>=0) { return strings.ReplaceAll(str1, " A.C"  , " A°C"  ) }
	if (strings.Index(str1, " A. C"  )>=0) { return strings.ReplaceAll(str1, " A. C" , " A° C" ) }
	if (strings.Index(str1, " B.C"   )>=0) { return strings.ReplaceAll(str1, " B.C"  , " B°C"  ) }         // English 
	if (strings.Index(str1, " B. C"  )>=0) { return strings.ReplaceAll(str1, " B. C" , " B° C" ) }	
	
	return str1
}	
//--------------------------
func check_manyDots(newY string) string {  
		/*
		remove ......  and . . . . .
		*/
		dots2 := ". . ."
		dots1 := "..."	
		j2:= strings.Index(newY, dots2);   	// ". . . "	
		j1:= strings.Index(newY, dots1);    // "....."
		jIncr:=1; 
		dotsX := dots2;
		if (j1 < 0) {		
			if (j2 < 0) {
				return newY 
			} else {
				j1 = j2
				jIncr = 2; 
				dotsX = dots2[0:2]
			}		
		} else {   // js1 > 0  ==> ....
			if (j2 >= 0) {  // js2 > 0 ==> . . . . 
				if (j2 < j1) { 
					j1=j2; jIncr=2;dotsX = dots2[0:2]   // dotsX = ". "
				} else { 
					jIncr=1;dotsX = dots1[0:1]     // dotsX = "."
				}
			} else {
				jIncr=1;dotsX = dots1[0:1]     // dotsX = "."
			}
		}
		//fmt.Println("check_manyDots() jIncr=", jIncr, " dotsX=" + dotsX + "<== input str=" + newY + "<=="); 
		newY += " ";
		
		ll1:= len(newY) 
		ll2:= ll1 - jIncr +1; 
		jEnd:=-1; 
		newY += "  "; 
		if (j1 > 0) {
			jEnd = -1; 
			for a:=j1; a < ll2; a+= jIncr {
				xy1:= newY[a:a+jIncr]				
				//fmt.Println("check_manyDots() j1=", j1, " a=", a, " xy1=" + xy1 + "<=="); 
				if (xy1 != dotsX) {
					jEnd=a;
					break; 
				}
			}
			if (jEnd < 0) { jEnd = ll1; } 
			if (newY[jEnd:jEnd+1] == ".") { jEnd++; }; 
			//fmt.Println("check_manyDots() xx j1=", j1 , " jEnd=", jEnd, "  ll1=", ll1);  
			newY = newY[0:j1] + " " + newY[jEnd:]; 
		}
		
		
		newY = strings.ReplaceAll(newY," . "," ");
		//fmt.Println("check_manyDots() j1=", j1, " jEnd=", jEnd, " out=" + newY); 
		
		return strings.TrimSpace( newY )
}		
		
//--------------------------------------
func removeWordContinuation() {

	// manage continuation: eg. 2 sentences, the first ending with "Ita-"  and the next beginning with "lia"  
	//										will be replaced by their concatenation without '-'  between them.  
	
	ll2 := len(newRows) -1; 
	
	for ix2, rline:= range newRows {		
		if (ix2 >= ll2            ) { continue } 	
		line := strings.TrimSpace(rline) 			
		ll:= len(line)
		line = strings.TrimRight(line, "\n") 
		line = strings.TrimRight(line, "\r")  
		if (ll < 2) { continue }
		if (line[ll-1:ll] != "-"  ) { continue }
		if ((line[ll-2:ll-1] == "-") || (line[ll-2:ll-1] == " ") ) { continue }
		line2 := strings.TrimSpace(newRows[ ix2+1 ])
		line2 = strings.TrimRight(line2, "\n") 
		line2 = strings.TrimRight(line2, "\r")  
		line2 = line2 + " "
		ix3 := strings.Index(line2, " ")	
		//fmt.Println("removeWordContinuation() ix2=",  ix2, " line2=" + line2 + "<==" , " ix3=", ix3)

		newRows[ ix2  ] = line[0:ll-1] + line2[0:ix3]      // append the first word of the next row
		newRows[ ix2+1] = strings.TrimSpace( line2[ix3:])  // cut the first word of the next row 
    }
} 

//--------------------------------------------
func writeTextRowSlice() {

		lines:= make([]string, 0, 10+len( newRows) )
		
		
		chgFileName := strings.Replace(outFileName, ".txt", "", -1)
		chgFileName = chgFileName + "_" + pref + ".txt"
		
		var numOut = fromNum
		
		// set an id for each line 
		
		for g1:=0; g1 < len( newRows); g1++ {
			lines = append(lines,  pref + "_" + strconv.Itoa( numOut ) + "|O|" + strings.TrimSpace( newRows[g1] )  ) 
			numOut++	
		}  
		writeList( chgFileName, lines )	
		
		fmt.Println("write ", len(lines), " lines on ", chgFileName, "\n" )
		
} // end of bind_go_write_row_dictionary 	


//----------------------
func writeList( fileName string, lines []string)  {
	// create file
    f, err := os.Create( fileName )
    if err != nil {
		sw_stop = true
		fmt.Println("error in create file ", fileName, "\n", err)
        return
    }
    // remember to close the file
    defer f.Close()

    // create new buffer
    buffer := bufio.NewWriter(f)

    for _, line := range lines {
        _, err := buffer.WriteString(line + "\n")
        if err != nil {
			sw_stop = true
			fmt.Println("error in writing ", fileName, "\n", err)
			return
        }
    }
    // flush buffered data to the file
    if err := buffer.Flush(); err != nil {
		sw_stop = true
		fmt.Println("error in flush out file ", fileName, "\n", err)
        return
    }
} 
//----------------------------------------


// =================================
