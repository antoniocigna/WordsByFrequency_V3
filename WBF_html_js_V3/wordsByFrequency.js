"use strict";
/*  
Words By Frequence: A tool to practice language comprehension
Antonio Cigna 2023
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//-----------------------------------------------

var html_fromIxRow  = 0;
var html_toIxRow  = 0;
//------------------------------------------
var numeroWord_TR=0;
const wSep = "§";
const endOfLine = ";;\n"; 
//var apiceInverso = `40`
let is_selected_row_only = false; 
let index_onlySelRowsWanted = set_ix_selectedRow(); //0

let swBegin=false;
let ele_allowWordTranslation= document.getElementById("id_allowWordTranslation"); 
let ele_wordsToTranslate 	= document.getElementById("id_words_to_translate");
let ele_wordsTranslated  	= document.getElementById("id_words_translated"  );
let ele_wordListDisplay  	= document.getElementById("id_wordListDisplay"   );
let ele_wordList = document.getElementById("id_wordList1");

//----------------------------------------------
var html_rowGroup_index_gr = 0 ;   // from onchange_rowGroupSelectChange
var html_rowGroup_beginNum = 0 ;   // from onchange_rowGroupNumBegChange 
var html_rowGroup_numRows  = 0 ;   // from onchange_rowGroupNumRowsChange 
var html_sel_extrRow       = "";   // from onchange_mostFreqWordList_extrRow
var last_html_rowGroup_index_gr = ""; 
var last_html_rowGroup_beginNum = "";
var last_html_rowGroup_numRows  = "";
var last_sel_extrRow_freqWord_list = ""; 
var sw_somethingChanged	= false;    // resetted  only by onclick_mostFreqWordList_require   

onchange_mostFreqWordList_extrRow(); 
onchange_rowGroupSelectChange(false); 

//var sw_rowGroupSelectChange       = false; 	
//var sw_rowGroupSelectChange_group = false; 	

//let ele_word     = document.getElementById("id_word"      );
//let ele_wRowList = document.getElementById("id_wRowList1");
//let ele_wordLisH = document.getElementById("id_wordListH");

let myPage01 = document.getElementById("id_myPage01");
let myPage02 = document.getElementById("id_myPage02");
let myPage03 = document.getElementById("id_myPage03");
let myPage04 = document.getElementById("id_myPage04");
let myPage05 = document.getElementById("id_myPage05");
let maxNumRow = 99999999999; // 100;
let wordToStudy_list ;
let numberOf_uniW = 0;
let numberOf_totW = 0;
let	numberOf_Row  = 0;
let prev_voice_ix        =  ""; 	
let	prev_voiceLang2      =  ""; 		
let	prev_voiceLangRegion =  "";  
let	prev_voiceName       =  ""; 
let listaInputFile   =""; 
let prevRunListFile  = ""
let lastRunLanguage  = "" 
const wordTTEnd     = ";:"
const wordTTBegin   = ";"
const wordSepEndBeg = wordTTEnd + wordTTBegin
//-------
let rowToStudy_list; 
let ele_toTranslate_textarea 	= document.getElementById("txt_pagOrig");  
let ele_translated_textarea  	= document.getElementById("txt_pagTrad"  );
let sw_newRowTranWritten = false;   // when true new translations  cannot  be asked 
 
//----------------
//let sw_newWordTranWritten = false;   // when true new translations  cannot  be asked 
let sw_ignore_missTranWord  = true;    // ignore missing translation 
let sw_ignore_missTranRow   = true;    // ignore missing translation 

let newTran;  // is an array with the same length as wordToStudy_list, when the element is = true  then a new translated word has been written   
let newRowTran;
let sw_firstDictLine_already_existed = false; 
//console.log("window dimensions = w=" , window.innerWidth, "  h=", window.innerHeight);   
//console.log(" javascript " + screen.width + " x " + screen.height);
//---------------------------------------
let ele_where = document.getElementById("id_where");
let ele_bar   =  document.getElementById("id_progrBar");
let ele_bar2  =  document.getElementById("id_progrBar2");
let ele_bar3  =  document.getElementById("id_progrBar3");
//------------------
function black(   str1 ) { return "\u001b[30m" + str1 }
function red(     str1 ) { return '\u001b[31m' + str1 }
function green(   str1 ) { return '\u001b[32m' + str1 }
function yellow(  str1 ) { return "\u001b[33m" + str1 }
function blue(    str1 ) { return "\u001b[34m" + str1 }
function magenta( str1 ) { return "\u001b[35m" + str1 }
function cyan(    str1 ) { return "\u001b[36m" + str1 }
function white(   str1 ) { return "\u001b[37m" + str1 }

//------------------------

function set_ix_selectedRow() {
	var x2 = document.getElementById("id_sel_2_extrRow");
    for (var i=0; i < x2.children.length; i++) {
		if (x2.children[i].id=="extrRow") {
			return i; 
		}
	}
	return 0; 
} // 

//-------------------------
function OLDextract_level(data0) {
	console.log("extract_level(data0=", data0) 
	/*
	result += "<br>" + fmt.Sprintln( sS.uniqueWords , " words (",  
				sS.uniquePerc, "%), make up ", sS.totPerc,"% of the text (", sS.totWords, " words)") 
	*/
	//var eleSelect = document.getElementById("id_sel_1_levTOLTO")
	/**
		<select id="id_sel_1_lev"> 
			<option value="any">qualsiasi</option>
			<option value="A0">A0</option>
			<option value="A1">A1</option><option value="A2">A2
		</select>
	**/
	
	//var newSelect = '   <option value="any">qualsiasi</option> \n' ; 
	
	var newSelect = '';
	
	var una, j1, oneLevel;
	j1 = (data0+ " ..end" ).indexOf("..end")
	var data = data0.substring(0, j1) 
	var righe = data.split(":");
	
	for (var z1=0; z1 < righe.length; z1++) {
		una = righe[z1].trim();
		if (una.lastIndexOf("-oth-") > 0) { continue; } 
		console.log("extract_level(data0) z1=", z1 , "  ". una) 
		j1=una.lastIndexOf(" ") 
		if (j1 < 0) {continue}
		oneLevel = una.substring(j1).trim()	
		if (oneLevel == "") { continue }		
		newSelect  += '   <option id="' + oneLevel + '">' + oneLevel + '</option> \n' ; 
	}	
	//eleSelect.innerHTML = newSelect; 
	
} // end of OLDextract_level 

//---------------------------


function js_go_updateStatistics( data, js_parm, jsFunc, goFunc) {
	
	function formatRight(num1, numDigit) { 
		var numS = "" + num1; 
		var nLen = numS.length
		if (nLen >= numDigit ) { return numS;} 
		return ("                ".substr(0, numDigit - nLen)) + numS; 
	}	
	
	var statRow = data.split("<br>");   
	//console.log("js_go_updateStatistics()  statRow=", statRow.join("<br>")) 
	
	var st1, field;
	var result="<table> \n";
	
	for (var z1=1; z1 < statRow.length; z1++) {  // ignore the first 
		st1 = statRow[z1];
		field = st1.split(",");
		if (field.length < 4) { continue;}		
		
		var line= "<tr>" +
				'<td style="text-align:right">' + field[0] + '</td><td style="text-align:left">' + "words ("           + "</td>" +
				'<td style="text-align:right">' + field[1] + '</td><td style="text-align:left">' + "%), make up "      + "</td>" +
				'<td style="text-align:right">' + field[2] + '</td><td style="text-align:left">' + "% of the text ("   + "</td>" + 
				'<td style="text-align:right">' + field[3] + '</td><td style="text-align:left">' + " words)" + "</td>" +
				"</tr> \n"	;		
		result += line; 
	}  	
	result += "</table>\n";
	
	document.getElementById("id_frequenze").innerHTML = result;
	
} // end of js_go_updateStatistics

//---------------------------------
//---------------------------------------------------------

function onclick_mostFreqWordList_require() {	
	
	fun_require_mostFreqWordList( false , "HTML page onclick_mostFreqWordList_require");
	
} // end of onclick_mostFreqWordList_require

//-----------------------------------------

function onchange_mostFreqWordList_extrRow() {	
	
	fun_require_mostFreqWordList( true , "HTML page onchange_mostFreqWordList_extrRow");
	
} // end of onchange_mostFreqWordList_require

//---------------------------------

function fun_require_mostFreqWordList( swFromOnChangeExtr , caller) {	
		
	document.getElementById("id_inpBegError").style.display = "none"; 	
	
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	var swChg=false;
	var fromWord = getInt( document.getElementById("id_inpBegFreqWList").value);	
	var numWords = getInt( document.getElementById("id_inpMaxNumWords" ).value);
	if (fromWord < 1) { fromWord=1;     document.getElementById("id_inpBegFreqWList").value = 1; }
	if (numWords < 1) { numWords = 1;   document.getElementById("id_inpMaxNumWords" ).value = 1; }
	
	//---
	var sel_level = "any"; //  x.options[i].id;
	//---
	var x2 = document.getElementById("id_sel_2_extrRow");
    var i = x2.selectedIndex;
	
	html_sel_extrRow = x2.options[i].id; 
	
	if (last_sel_extrRow_freqWord_list == "") {last_sel_extrRow_freqWord_list = html_sel_extrRow; }  
	
	
	is_selected_row_only = ( i == index_onlySelRowsWanted); //1 onclick_require_mostFreqWordLi
	
	fun_selRowsWanted_changed();
	//---	
	var xTbl = document.getElementById("id_sel_tblwords");
    var i2 = xTbl.selectedIndex;
	var sel_toBeLearned = xTbl.options[i2].id;    //( 0 = 'allWords'   1 = 'toBeLearned' )
	//--
	swChg = isExtrRowChanged() 
	if (swChg) {
		if ((last_sel_extrRow_freqWord_list == html_sel_extrRow) && (html_sel_extrRow == "anyRow")) {
			swChg = false; 
			console.log("onchange_mostFreqWordList_extrRow " , " isExtrRowChanged=", true, " but sel_extrRow=anyRow as before,  reset isExtrRowChanged= false");	
		}
	}
	if (swFromOnChangeExtr) {
		//console.log("onchange_mostFreqWordList_extrRow " , " isExtrRowChanged=", swChg , " return");	
		return; 
	}
	
	//console.log("onclick_mostFreqWordList_require " , " isExtrRowChanged=", swChg , " go_passToJs_wordList --> js_go_showWordList_lev2( , button=1)");	
	
	//var caller = "HTML page onclick_require_rowList1" //  (new Error()).stack?.split("\n")[2]?.trim().split(" ")[1] ;
	//if (caller == undefined) { caller = ""; }
	
	//go_passToJs_rowList(""+inpBegRow, ""+numRows, "js_go_rowList," + caller); 
	
	go_passToJs_wordList( swChg, ""+fromWord, ""+numWords, sel_level, html_sel_extrRow, sel_toBeLearned, "js_go_showWordList_lev2(1)," + caller);
	
	
} // end of fun_require_mostFreqWordList

//------------------------------------------
function js_go_console( str1 ) {
	//console.log( str1 )	
} 

//-------------------------------------
function sortAlpha(wordToStudy_listStr) {
	//wordToStudy_listStr
	var ww, col1, key;
	var listKey = [];
	for (var z=0; z < wordToStudy_listStr.length; z++) {
		col1 = (wordToStudy_listStr[z].trim() + ";.;.;.;.;.;.;.;.;.;.").split(";.")
		key = col1[0];  // wordCod		
		listKey.push(key  + ":" + z ); 
	}
	return listKey.sort();
	
} // end of sortAlpha 
//-------------------------------------
function sortFreq(wordToStudy_listStr) {
	//wordToStudy_listStr
	var ww, col1, key, freq1, freq2;
	var listKey = [];
	
	for (var z=0; z < wordToStudy_listStr.length; z++) {		
		col1 = (wordToStudy_listStr[z].trim() + ";.;.;.;.;.;.;.;.;.;.").split(";.")
		freq1 = 1*("0" + col1[3].trim());
		freq2 = 10000000 -  freq1;
		key = freq2 + " " + col1[0] 
		//console.log(wordToStudy_listStr[z], "\n\t col0=", col1[0], " col[3]=", col1[3], " freq1=", freq1, " freq2=", freq2, " key=", key)
		listKey.push(key  + ":" + z ); 
	}

	return listKey.sort();
	
} // end of sortFreq 

//-------------------------------------
function errorNoWord1() {
	var startIx = getInt( document.getElementById("id_inpBegFreqWList").value );
	if (numberOf_uniW < startIx) {
		document.getElementById("id_inpBegErrMsg").innerHTML =  " il numero di partenza "+  startIx + " supera il numero di parole " + numberOf_uniW + " (forzato 1)"; 
		document.getElementById("id_inpBegFreqWList").value = 1
	} else {
		document.getElementById("id_inpBegErrMsg").innerHTML = ""
	}	
	document.getElementById("id_inpBegError").style.display = "inline-block"; 		
} 
//----------------------
function js_go_showWordList_lev2(wordListStr00, numButton, jsFunc,goFunc) {
	//console.log("function js_go_showWordList_lev2() ", "wordListStr.length=",wordListStr00.length  ," numButton=", numButton, " <-- " + goFunc + " <-- " + jsFunc) ;
	// numButton=1 default ==> from onclick most frequent word list  
	// numButton=2         ==> from onclick BetweenWordList or prefix wordlist   
	// numButton=3         ==> from onclick Lemma word list   
	// numButton=0         ==> from word list from word, lemma, ?   
	if (numButton==1 ) {
		sw_somethingChanged = false; 
	} 
	word_to_underline_list = []
	var wordListStr = wordListStr00.trim();
	var len = wordListStr.length	

	if (wordListStr.substring(len-1) == ";") { len = len - 1; wordListStr = wordListStr.substring(0, len) }
	if (wordListStr.substring(len-1) == ";") { len = len - 1; wordListStr = wordListStr.substring(0, len) }
		
    // triggered by go func (  go _ passToJs_wordList )
    if (wordListStr == undefined) {
        //console.log("js_showWordList: parameter is undefined");
		onclick_jumpFromToPage( myPage02,myPage03,  myPage01); 
        return;
    }
    if (wordListStr == "") {
        console.log("js_showWordList: parameter is empty");
		if (numButton==1) { errorNoWord1()}
		onclick_jumpFromToPage( myPage02,myPage03,  myPage01); 
        return;
    }	
	var wLemmaListU, wTranListU, wLevelListU,	wParaListU, wExampleListU, wIxLemmaListU;   
	var wLemmaList,  wTranList,  wLevelList,    wParaList,  wExampleList , wIxLemmaList ;
	var word2, ixUnW2, totRow2, totExtrRow2 
	var knowYesCtr, knowNoCtr ;
	var wordCod, chk_ix, chk_ixLemma;
	
    var wordToStudy_listStr = wordListStr.split( endOfLine );	
		
	/*
wordListStr=
                     0                 1         2     3     4        5              6    7   8   9  10   11 
genannt.genannt      ;.  genannt    ;.505;.    145;. 123;. 123;. nennen  ;.  nome    ;.   ;.  ;.  ;.  0;.  ;					 
\ngen.gen            ;.  gen        ;.7548;.     2;. 123;. 123;.  gen    ;.  gen     ;.   ;.  ;.  ;.  0;.  ;;
\ngenannte.genannte  ;.  genannte   ;.13717;.    1;. 123;. 123;. genannt ;.  chiamato;.   ;.  ;.  ;.  0;.  ;;

	*/	
	wordToStudy_list = []
	var ixNumPlus; 
	var z;
	
	
	var listKey=[]; var keyS, keyIx;
	
	if (numButton == 1) {		
			listKey = sortFreq(wordToStudy_listStr) 
			for (var x=0; x < listKey.length; x++ ) {
				[keyS, keyIx] = listKey[x].split(":") 
				oneElemToStudy(keyIx)
			} 		 
	}
	if (numButton > 1) {
		
		listKey = sortAlpha(wordToStudy_listStr)   // cod 
		
		for (var x=0; x < listKey.length; x++ ) {
				[keyS, keyIx] = listKey[x].split(":") 
				oneElemToStudy(keyIx)
		} 
	}
	//--------------------
	
	function oneElemToStudy(z) {		
		
		var wordLineZ = wordToStudy_listStr[z].trim();  
		if (wordLineZ == "") return; 
		
		
		//[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList] = getFieldsFromWordToStudy( wordToStudy_listStr[z],"1wordToStudy_"+z ); 	
		
		//var col0 = getFieldsFromWordToStudy( wordLineZ,"1wordToStudy_"+z ); 
		
		var ww0 = ((wordLineZ + ";.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.").split(";.") ).slice(0,15);
		
		[wordCod, word2, chk_ix, ixUnW2, totRow2, wLemmaListU, wTranListU,	wLevelListU,	wParaListU, wExampleListU,
					   totExtrRow2, knowYesCtr, knowNoCtr,   chk_ixLemma, wIxLemmaListU] = ww0;  
				   
		wLemmaList   = wLemmaListU.split(   wSep ) 			   
		wTranList    = wTranListU.split(    wSep ) 		
		wLevelList   = wLevelListU.split(   wSep ) 		
		wParaList    = wParaListU.split(    wSep ) 		
		wExampleList = wExampleListU.split( wSep ) 		   
		wIxLemmaList = wIxLemmaListU.split( wSep ) 		   
					   
		//	0       1      2       3       4        5           6            7          8            9         
		//   				10            11         12           13          14
		/**
		if ((chk_ix != "ix") || (chk_ixLemma != "ixLemma"))  {
			console.log( red("error "), " in ", green("js_go_showWordList_lev2"), 
				" in string argument got from GO (field3 not 'ix' or field13 not 'ixLemma' 0\n\t", ww)   			
		}	
		***/		
		/**
		if (z==1) {
			console.log( green("oneElemToStudy "), z,"  ", wordLineZ, "\n\t" ,  chk_ix,ixUnW2, chk_ixLemma, wIxLemmaList)			
		}
		**/
			
		if ((word2 == "") || (word2 == "...") ) return ; 
		if (word2.indexOf("…") >= 0) return;  
		var trattini = "-_*."; 
		if (trattini.indexOf( word2.substring(0,1) ) >=0) return;
		
		wordToStudy_list.push(  [word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, 
									totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList ] ); 
			
	}  // end of oneElemToStudy
	
	newTran = [];
	
	//console.log( green("js_go_showWordList_lev2"), " wordToStudy_list.length=" , wordToStudy_list.length);
	
	for ( z=0; z < wordToStudy_list.length; z++) {
		//console.log("\t ", z, " wordToStudy_list = ", wordToStudy_list[z]); 
		newTran.push( 0 )
	}

	fun_showWordList("3") 
	
} // end of js_go_showWordList_lev2

//----------------
function OLDgetFieldsFromWordToStudy(  ww , where="") {
		
		//const citrus = fruits.slice(1, 3);     


		//  xWordF.word + "," + strconv.Itoa(xWordF.ixUnW) + "," + strconv.Itoa(xWordF.totRow)  + ";[" +xWordF.wLemmaL + "];[" + xWordF.wTranL) + "];"
		
		if (ww=="") { return ["", "",null,null,null,null,null,null,null,null,null,null]}
		//if (ww.indexOf("]") >0) { ww = (ww+";").replaceAll(";[",";").replaceAll("];",";")   }
		
		//console.log("antonio get FieldsFromWordToStudy( ww=" + ww + "<===") 
		
		var ww0 = ((ww + ";.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.").split(";.") ).slice(0,15)
		/****
		[wordCod, word2, chk_ix, ixUnW2, totRow2,wLemmaList, wTranList,	wLevelList,	wParaList, wExampleList,
					   totExtrRow2, knowYesCtr, knowNoCtr,   chk_ixLemma, wIxLemmaList] = ww0;  
		//	0       1      2       3       4        5           6            7          8            9         
		//   				10            11         12           13          14
		if ((chk_ix != "ix") || (chk_ixLemma != "ixLemma"))  {
			console.log( red("error "), " in ", green("js_go_showWordList_lev2"), 
				" in string argument got from GO (field3 not 'ix' or field13 not 'ixLemma' 0\n\t", ww)   			
		}
		***/
		
		var varcode = ww0[0]
		var piece = ww0.slice(1);
	/*
			 0 bin.bin;.
			 1 bin;.
			 2 ix;.
			 3 3021;.
			 4 6;.
			 5 sein;.
			 6 essere;.
			 7 ;.
			 8 ;.
			 9 ;.
			10 ;.
			11 ;.
			12 ;.
			13 ixLemma;.
			14 13927;.
			;;
		**/	      	
		
		
		// [word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList]	
		//    0       1       2          3          4           5          6           7     
			
		//console.log("\t\tantonio  piece7 =" + piece[7] ) 	
		//console.log("\t\tantonio  split  piece7 =" + ( piece[7].split( wSep) ).length , " uno=" + piece[7].split( wSep)[1] ) 	
			
		return [ varcode, piece[0], piece[1], piece[2],  
			piece[3].split( wSep ) , 
			piece[4].split( wSep ) , 
			piece[5].split( wSep ) , 
			piece[6].split( wSep ) ,  
			piece[7].split( wSep ) , 
			piece[8].split( wSep ) ,
			piece[9].split( wSep ) , 
			piece[10].split( wSep ),
			piece[11].split( wSep )  	
			] ;
	
} // end of OLDgetFieldsFromWordToStudy  
//------------------------------------
function fun_showWordList(wh, ix1=-1) {	
	
	var numNoTran = 0; // -1 
	var word2, ixUnW2, totRow2, totExtrRow2, wLemmaList, wTranList , wLevelList, wParaList, wExampleList, knowYesCtr, knowNoCtr , wIxLemmaList  ; 
	var words_to_translate_str = wordTTBegin   //  ; 
	
	for (var z=0; z < wordToStudy_list.length; z++) {			
		//[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList ] ???antoX
		[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList ] = wordToStudy_list[z] ; 	
		
		//console.log("fun_showWordList(", wh, " ix1=", ix1, ") wordToStudy_list: ", "word2=", word2, "  ixUnW2=",  ixUnW2, " wLemmaList=", wLemmaList, " wTranList=", wTranList)
		
		for(var f = 0; f < wLemmaList.length; f++) {	
			
			if (wTranList[f] == wLemmaList[f]) {
				wTranList[f] += " _" ;
				//wTranList[f] = "" ; // per evitare parole non tradotte che sembrano tradotte,  se è davvero la traduzione aggiungi un carattere per diversificarla per esempio  " _"  
			}  	
			if (wTranList[f] == "") {			
				numNoTran++			
				words_to_translate_str += z + ";" + ixUnW2 + ";" + f + "; " +wLemmaList[f] + wordSepEndBeg;  // ;,;    // blank dopo ; prima di wLemma (serve per evitare problemi google)
			} else {				
				if (wTranList[f] == "_word_not_found_") {
					numNoTran++			
					words_to_translate_str += z + ";" + "-1" + ";" + f + "; " +wLemmaList[f] + wordSepEndBeg;  	
				}
			}
		}
	}
	
	if (numNoTran < 1) {		
		showWordsAndTranButton("2")
		return
	}	
	//console.log("X fun_showWordList(", wh, ")  words_to_translate_str =\n", words_to_translate_str)	
		
	ele_wordsToTranslate.value = words_to_translate_str; //  .replaceAll("\n"," "); 
	ele_wordsTranslated.value = ""; 
	document.getElementById("id_notTranNum").innerHTML = numNoTran; 
	
	onclick_jumpFromToPage( myPage01,0,myPage02);  
	

} // end of fun_showWordList

//------------------------------

function sentenceOneRow( str1 ) {
	
	var str2 = str1.replaceAll("\n"," ")
	str2 = str2.replaceAll(" 1.","<br>1.").
			replaceAll(" 2.","<br>1."). 	
			replaceAll(" 3.","<br>2."). 	
			replaceAll(" 4.","<br>3."). 	
			replaceAll(" 5.","<br>4."). 	
			replaceAll(" 6.","<br>5."). 	
			replaceAll(" 7.","<br>6."). 	
			replaceAll(" 8.","<br>7."). 	
			replaceAll(" 9.","<br>8.")  ;
	str2 = str2.replaceAll(". ",".<br>").replaceAll("? ","?<br>").replaceAll("! ","!<br>"); 
	
	//if (str1.indexOf("Hund") >= 0) { console.log("ANTONIO str1=" + str1 + "\nstr2=" + str2) }
	
	return str2	
}

//-------------------------------------------

function oneTR_lemma( ixW2StudyLs, ixLemma, clas1, word1, ix1, nrow, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, 
					totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList) {	
			
			var wLemma1;
			var riga;
			var wordOrig2, wordTran2;
			
			
			var wordTran = ""; 
			var wLemma3, wTran3;
			var nSpanV, spanV;
			
			var showList = ""; 
		
			var f_lemma   = wLemmaList[  ixLemma ];
			var f_ixLemma = wIxLemmaList[ixLemma ];
			var f_tran    = wTranList[   ixLemma ]; 
			
			var f_level   = wLevelList[  ixLemma ]; 
			var f_para    = wParaList[   ixLemma ]; 
			var f_example = wExampleList[ixLemma ]; 
			
			var nn_level   = f_level.split("|")	
			var nn_para    = f_para.split("|")	
			var nn_example = f_example.split("|")	
				
			var x_level, x_para, x_example; 
			var x_para1, x_para2, x_example1, x_example2; 	
			var jbr; 			
			
			var riga00 = ""	
			
			var hig = 1.2 
						
			var XixW2StudyLs = ""+ixW2StudyLs;  
			var Xix1   = ""+ix1;  
			var Xnrow  = ""+nrow; 	
			var XnExtrRow = "";
			if (totExtrRow2) { if (totExtrRow2 > 0) { XnExtrRow = ""+totExtrRow2}};  	
				
			var Xword1 = word1;
		
			var Xf_lemma = f_lemma;
			var Xf_para  = f_para; 
			var Xf_tran  = f_tran.replaceAll("|", "<br>") ; 
			
			//-----------------------	
			var key="",	pKey=""
			var numLev = nn_level.length;
			
			var num1=0
			
			for (var m=0; m < numLev; m++) {	
				pKey = key
				x_level = nn_level[m]; 
				x_para  = nn_para[m];
				x_example = nn_example[m]; 
				key = x_level + " " + x_para + " " + x_example 
				if (key == pKey) {continue; }	
				num1++
			}

			key="";

			// num > 1, significa che ci sono diverse righe
			
			for (var m=0; m < numLev; m++) {	
				pKey = key
				x_level = nn_level[m]; 
				x_para    = sentenceOneRow( nn_para[m] );
				x_example = sentenceOneRow( nn_example[m] ); 
				
				//key = x_level + " " + x_para + " " + x_example 
				//if (key == pKey) {continue; }				
				
			
				jbr = x_para.indexOf("<br>") 			
				if (jbr > 0) { x_para1 = x_para.substring(0, jbr); x_para2 = x_para.substring(jbr+4).trim() }	
				else { x_para1 = x_para; x_para2 = ""		 }
							
				jbr = x_example.indexOf("<br>") 			
				if (jbr > 0) { x_example1 = x_example.substring(0, jbr); x_example2 = x_example.substring(jbr+4).trim() }	
				else { x_example1 = x_example; x_example2 = ""		 }		
				
				if ( m !=0) {
					//Xf_lemma = "";
					Xf_para  = ""; 
					Xf_tran  = "";					
				}
				
				
				var showAltre = "";
				
				if (m == 0) {
					if (x_example2 != "") { riga += '	<span style="display:none;">' + x_example2 + '</span>'; }
					riga += '</td> \n'	
					
					if ((num1 == 1) && (x_example2 == "")) {  // se solo un riga da mostrare  non mostrare i tre puntini  di mostra altro 
						showAltre = "";
						riga += '<td></td> \n';	
							}  else {
						riga += '<td>' + 
							'<span onclick="show_altreRighe(this,' + (numLev - 1) + ')">&hellip;</span> ' +
							'<span style="display:none;">none</span>' +  
							'</td>	\n' ;
						showAltre = '<span onclick="show_altreRighe(this,' + (numLev - 1) + ')">&hellip;</span><span style="display:none;">none</span> ' ;				
					}
				} else {
					if (x_example2 != "") { riga += '	<span style="display:block;">' + x_example2 + '</span>'; }						
					riga += '</td> \n'						
					//  non mostrare i tre puntini  di mostra altro 
					riga += '<td></td> \n';	
				}	
						
						
				riga  +='</tr> \n' ;   		
				
				
				if ((num1 == 1) && (x_example2 == "")) {  // se solo un riga da mostrare  non mostrare i tre puntini  di mostra altro 
						showAltre="";
					}  else {						
						showAltre = '<span onclick="show_altreRighe(this,' + (numLev - 1) + ')">&hellip;</span><span style="display:none;">none</span> ' ;							
					}
				
				numeroWord_TR++;				
				var newTr = newTr_from_prototype( numeroWord_TR, clas1, XixW2StudyLs, ""+f_ixLemma, ixLemma, Xnrow, 
					Xix1, Xword1, Xf_lemma, x_level, Xf_para, Xf_tran, x_para1, x_para2, x_example1, x_example2,  showAltre, m, 
					XnExtrRow, knowYesCtr, knowNoCtr ); 
				//console.log("oneTR_lemma ", newTr); 
				showList    += newTr ;	
			}
			
		return showList
		
}  // end of oneTR_lemma2
//----------------------------------------
function fun_selRowsWanted_changed() {
	var ele_sel = document.getElementById("id_sel_2_extrRow");
	//var ele_listBut = document.getElementById("id_list_Righe_TD_But") 
	var ele_listNum = document.getElementById("id_list_Righe_TD_Num") 

	if (is_selected_row_only) {
		ele_sel.style.border = "4px solid blue"; 
		ele_sel.style.color  = "blue"; 
		//ele_listBut.style.border = "4px solid blue";
		ele_listNum.style.border = "4px solid blue"; 
	} else {
		ele_sel.style.border     = null; 
		ele_sel.style.color      = null; 
		//ele_listBut.style.border = null;
		ele_listNum.style.border = null; 
	}
	
} // end of fun_selRowsWanted_changed
//-------------------------------------------
function showWordsAndTranButton(wh) {
		
	
	var showList = prototype_tableWordList_Header;  
		
	var x2 = document.getElementById("id_sel_2_extrRow");
    var i = x2.selectedIndex;
	var sel_extrRow = x2.options[i].id;
	is_selected_row_only = ( i == index_onlySelRowsWanted);  // 2showWordsAndTranButton(wh) 
	fun_selRowsWanted_changed();
	//console.log("require most freq.:  is_selected_row_only =",  is_selected_row_only); 
	if (is_selected_row_only) {
		showList = showList.replace("§sel2collapse§", "visible");			
	} else {
		showList = showList.replace("§sel2collapse§", "collapse"); 
	}
	//console.log("showWordsAdTran:  is_selected_row_only=", is_selected_row_only, " showList=", showList.substr(0,300));  
	
    var word2, ix1,ixUnW2, totRow2, nrow, totExtrRow2, wLemma1;
	var riga;
	
	var wordOrig2, wordTran2;
	var wIxLemmaList, wLemmaList, wTranList , wLevelList, wParaList, wExampleList;
	var knowYesCtr, knowNoCtr;
	
	var wordTran = ""; 
	var wLemma3, wTran3;
	var nSpanV, spanV;
	
	var clas1;
	var hig=1
	var swP; 
	var col1; 
	numeroWord_TR = 0;
	//-------------------
	
    for (var ixW2StudyLs = 0; ixW2StudyLs < wordToStudy_list.length; ixW2StudyLs++) {

		[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, 
				totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList ] = wordToStudy_list[ixW2StudyLs]; 
		
		if (ixUnW2 < 0) {continue}
		
		nSpanV =  wLemmaList.length
		//------------------------	
		for(var ixixLemma = 0 ; ixixLemma < wLemmaList.length; ixixLemma++) {	
			showList += oneTR_lemma(ixW2StudyLs, ixixLemma, "", word2, ixUnW2, totRow2, 
					wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList ); 
		}
		//-------------------
    }
	//---------------------------------------------------
	
	showList += '   </tbody>  \n' +
		'</table> \n'; 	
    ele_wordList.innerHTML = showList;
   	
	onclick_jumpFromToPage( myPage01,myPage02, myPage03);  
	
} // end of showWordsAndTranButton

//---------------------

function writeLanguageChoise() {
	//---------------------------------------------------
	// triggered by onchange_tts_get_oneLangVoice(this1) 
	
	//console.log("writeLanguageChoise()"); 
	
	var langRow = ""; 
	if ( isVoiceSelected ) {		
		langRow += selected_voice_ix + "," + selected_voiceLang2 + "," + selected_voiceLangRegion + "," +  selected_voiceName;
	}
	// langRow += "<file>" + prevRunListFile
	
	if (langRow == "") return
	if (langRow == lastRunLanguage) return
	
	//console.log(" write file language " + langRow); 
	
	js_go_ready( langRow);  
	
	go_write_lang_dictionary(   "language="  + langRow );  
	
} // end of writeLanguageChoise
//--------------------------------------------

//--------------------------------------------------
function write_word_dictionary() {
	//console.log("write word dictionary()")
	let word1, ix1, nrow, totExtrRow2,  wLemma1, wordTran, knowYesCtr, knowNoCtr   ;
	var wLemmaList, wTranList, wLevelList, wParaList, wExampleList, wIxLemmaList;
	var newTranWord=0;
	var listNewTranWords = "";
	
    for (var ixW2StudyLs = 0; ixW2StudyLs < wordToStudy_list.length; ixW2StudyLs++) {
		//[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList ] ???antoX
		[word1, ix1,        nrow, wLemmaList, wTranList, wLevelList, wParaList, 
				wExampleList, totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList ] = wordToStudy_list[ixW2StudyLs]; 
		word1 = word1.trim(); 
		if (ix1 == undefined) {
            continue;
        }  		
		if (ix1 < 0) continue; 
		
		if (newTran[ixW2StudyLs]==1) {
			if (ix1 == -1) { continue; }	
			newTranWord++; 	
			listNewTranWords += "\n" + word1 + ";" + ix1 + ";" + wLemmaList.join( wSep ) + ";" + wTranList.join( wSep )  ;   // new line for dictionary 	
			//console.log("\t NEWLINE DICT WORD=",  	 word1 + ";" + ix1 + ";" + wLemmaList.join( wSep ) + ";" + wTranList.join( wSep ) );
		}
    }
	
	if (newTranWord < 1) {
		//console.log("write_word_dictionary",  " nessuna nuova traduzione"); 	
		return;
	}
	//console.log("write_word_dictionary ",  newTranWord, " parole tradotte"); 	
	//console.log(red("\tJS  run go_write_word_dictionary("),  listNewTranWords.substring(1)   ); 	
	
	go_write_word_dictionary(  listNewTranWords.substring(1)  ); 
	
	//console.log(red("\tJS  dopo l'istruzione di run go_write_word_dictionary")); 	
	
} // end of write_word_dictionary

//---------------------------------------------------------

function onclick_require_prefixWordList() {
	 word_to_underline_list = []
	ele_wordList.innerHTML ="";
	//ele_wRowList.innerHTML = "";
    //ele_word.innerHTML     = ""; 	 
	//ele_wordLisH.style.display = "none";
    var sNumWords = document.getElementById("id_inpMaxNumWords").value;
    var wordPrefix = document.getElementById("id_inpPref").value.trim();   
	

	var numWords=0; 
    try {
        numWords = parseInt(sNumWords);
    } catch (err) {
		ele_wordList.innerHTML ='<span style="color:red;">errore numWords non numerico =>' + numWords +    '<== sNumwords=' + sNumWords + '<==  </span>';
		return;
	}
	if (numWords < 1) {
		ele_wordList.innerHTML ='<span style="color:red;">massimo numero di parole minore di 1</span>';
		return; 
	}	
	if (wordPrefix == "") {
		ele_wordList.innerHTML ='<span style="color:red;">manca il prefisso</span>';
		return;
	}	
	
    go_passToJs_prefixWordList(""+numWords, wordPrefix, "js_go_showPrefixWordList"); // ask 'go' to give wordlist by js_... function  
	
} // end of onclick_require_prefixWordList

//-------------
//---------------------------------------------------------

//------------------------------------
function js_go_showPrefixWordList(wordListStr, js_parm, jsFunc,goFunc) { 

	
	// triggered by go ( go_passToJs_rowList and js_go_showWrdRowList)
	
	console.log("function js_go_showPrefixWordList() ", " js_parm=", js_parm, " <-- " + goFunc + " <-- " + jsFunc) 
	//console.log("inpstr=" +inpstr ) 

    // triggered by go func  (go_passToJs_prefixWordList)
	console.log( "js_go_showPrefixWordList(wordListStr = " + wordListStr)
	
	if (wordListStr.substring(0,5) == "NONE,") {
			document.getElementById("id_inpPref_msgWord").innerHTML = wordListStr.substring(5) 
			document.getElementById("id_inpPref_msg").style.display = "block"
			myPage01.style.display = "flex"; 
			//onclick_jumpFromToPage( myPage02,myPage03, myPage04); 
		return
	}
	document.getElementById("id_inpPref_msg").style.display = "none"
	onclick_jumpFromToPage( myPage01,0, myPage02);  
	
	js_go_showWordList_lev2(wordListStr, 2);

} // end of js_go_showPrefixWordList

//------------------------------------------------------

function onclick_require_betweenWordList() {
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	
	var eleFromW = document.getElementById("id_inpFromABC");
	var eleToW   = document.getElementById("id_inpToABC"  ); 
	var fromWordPref = eleFromW.value.trim(); 		
	var toWordPref   = eleToW.value.trim(); 		
    /***
	var sMaxWords = document.getElementById("id_inpMaxNumABC" ).value;							
    var maxNumWords = 0;
    try {
        maxNumWords = parseInt(sMaxWords);
    } catch (err) {}    
    if (maxNumWords < 1) {
        maxNumWords = 1;
    }	
	***/
	var maxNumWords = getInt(  document.getElementById("id_inpMaxNumABC" ).value );	
	if (maxNumWords < 1) {  maxNumWords = 1; document.getElementById("id_inpMaxNumABC" ).value = 1; }	
	
	if (toWordPref == "") { 
		if (fromWordPref == "") { return; }
		toWordPref = fromWordPref;
		//document.getElementById("id_inpToABC"  ).value = toWordPref; 	
	} else {
		if (fromWordPref == "") { 
			fromWordPref = toWordPref;
			document.getElementById("id_inpFromABC").value = fromWordPref; 
		}
	} 
	if (toWordPref < fromWordPref) {
		document.getElementById("id_inpToABC"  ).value = fromWordPref; 	
		document.getElementById("id_inpFromABC").value = toWordPref;
		fromWordPref = document.getElementById("id_inpFromABC").value.trim(); 		
	    toWordPref   = document.getElementById("id_inpToABC"  ).value.trim(); 		
	}
	eleFromW.style.color = null;
	eleFromW.parentElement.style.backgroundColor = null;
	var eleLemma = document.getElementById("id_inpLemma");  
	eleLemma.style.color = null;
	eleLemma.parentElement.style.backgroundColor = null;

	go_passToJs_betweenWordList(""+maxNumWords, fromWordPref, toWordPref, "js_go_showBetweenWordList"); // ask 'go' to give wordlist by js_... function  
	
} // end of onclick_require_prefixWordList


//------------------------------------
function js_go_showBetweenWordList(wordListStr, js_parm, jsFunc,goFunc) {
	//console.log("function js_go_showBetweenWordList() ", " js_parm=", js_parm, " <-- " + goFunc + " <-- " + jsFunc) ;
	//console.log("wordListStr=\n"+wordListStr +"\n--------------------------\n")
	
	if (wordListStr == "") {
		document.getElementById("id_bW_err").style.display ="block";   // no entry found
    } else {
		document.getElementById("id_bW_err").style.display ="none"; 
	}
	
	onclick_jumpFromToPage( myPage01,0, myPage02);  //myPage01
	
	js_go_showWordList_lev2(wordListStr,2, jsFunc, goFunc)

} // end of js_go_showBetweenWordList


//------------------------------
function onclick_require_lemmaWordList() {
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	
	var inpMaxWordLemma = document.getElementById("id_inpMaxWordLemma").value; 
    var aLemma = document.getElementById("id_inpLemma").value.trim();    
	if (aLemma == "") {
		ele_wordList.innerHTML ='<span style="color:red;">manca il lemma</span>';
		return;
	}	
	
	
	var eleFromW = document.getElementById("id_inpFromABC");
	eleFromW.style.color = null;
	eleFromW.parentElement.style.backgroundColor = null;
	var eleLemma = document.getElementById("id_inpLemma");  
	eleLemma.style.color = null;
	eleLemma.parentElement.style.backgroundColor = null;
	
	myPage01.style.display = "none"; 
	var caller = "HTML page onclick_require_lemmaWordList"
	
	go_passToJs_lemmaWordList(aLemma, inpMaxWordLemma, "js_go_showLemmaWordList," + caller);  	

} // end of onclick_require_lemmaWordList
//------------------------------------
function js_go_showLemmaWordList(wordListStr,  js_parm, jsFunc,goFunc) {
	//console.log(" js_go_showLemmaWordList () ", "wordListStr=\n" + wordListStr + "\n-------------------\n")
	
	
	if (wordListStr.substring(0,5) == "NONE,") {
		//document.getElementById("id_inpLemma_word").innerHTML = wordListStr.substring(5) 
		document.getElementById("id_inpLemma_msg").style.display = "block"			
		myPage01.style.display = "flex"; 
		//onclick_jumpFromToPage( myPage02,myPage03, myPage04); 
		return
	}
	document.getElementById("id_inpLemma_msg").style.display = "none"
	onclick_jumpFromToPage( myPage01,0, myPage02);  
	
	//console.log("js_go_showLemmaWordList ()  chiama js_go_showWordList_lev2")
	
	js_go_showWordList_lev2(wordListStr, 3, jsFunc,goFunc  );

} // end of js_go_showLemmaWordList

//---------------------------------------------------------

function onclick_require_PrefWordFromRowList(word1) {
	/*
	nella lista di frasi 
	 per ogni frase c'è la possibilità di sploderla in tutte le sue parole  ( tasto lente ingrandimento)
	 per ogni parola c'è un pulsante che richiama questa funzione ("pref") 
	Questa funzione lista tutte le parole che iniziano con questa parola 	
	*/
	word1 = word1.trim()
	if (word1 == "") {
		return;
	}	
	var eleFromW = document.getElementById("id_inpFromABC");
	var eleToW   = document.getElementById("id_inpToABC"  ); 
	eleToW.value = ""; 
	eleFromW.value = word1;  
    eleFromW.style.color = "blue";
	eleFromW.parentElement.style.backgroundColor = "yellow";
	
   
	myPage01.style.display = "flex";  
	myPage03.style.display = "none"; 
	myPage05.style.display = "none"; 
} // end of onclick_require_PrefWord
//---------------------------------------------------------

//---------------------------------------------------------

function onclick_require_PrefWord_lemma(word1, lemma1) {
			
	//var max_num_word4LeWor = document.getElementById("idTabWoWL3").value 

	/*
	nella lista di frasi 
	 per ogni frase c'è la possibilità di sploderla in tutte le sue parole  ( tasto lente ingrandimento)
	 per ogni parola c'è un pulsante che richiama questa funzione ("pref") 
	Questa funzione lista tutte le parole che iniziano con questa parola 	
	*/
	word1 = word1.trim()
	if (word1 == "") {
		return;
	}	
	var eleFromW = document.getElementById("id_inpFromABC");
	var eleToW   = document.getElementById("id_inpToABC"  ); 
	eleToW.value = ""; 
	eleFromW.value = word1;  
    eleFromW.style.color = "blue";
	eleFromW.parentElement.style.backgroundColor = "yellow";
	var eleLem  = document.getElementById("id_inpLemma")
	var eleLemTD = eleLem.parentElement
	var eleLemTR = eleLemTD.parentElement
	eleLem.value = lemma1; 
	eleLem.style.color = "blue"; 
	eleLemTD.style.backgroundColor = "yellow";
   
	myPage01.style.display = "flex";  
	myPage03.style.display = "none"; 
	myPage05.style.display = "none"; 
} // end of onclick_require_PrefWord
//--------------------------------------------------------

//------------------------------------------------------------------

function onclick_require_rowListWithThisWord2(type,word1, maxNumRow5) {
	//console.log("\n\nonclick_require_rowListWithThisWord2(" , "type=", type, ", word1=" + word1+")" );   
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	//ele_wRowList.innerHTML = "";
    //ele_word.innerHTML     = ""; 	 	
	//ele_wordLisH.style.display = "none";
	//console.log('document.getElementById("id_inpWordFra") =' , document.getElementById("id_inpWordFra").outerHTML) 
    var aWord ="";  
	if (type==2) {
		aWord = word1; 	
	} else {		
		aWord     = document.getElementById("id_inpWordFra").value.trim();  
	}	
	if (aWord == "") {
		ele_wordList.innerHTML ='<span style="color:red;">manca la parola da cercare</span>';
		return;
	}	
	myPage01.style.display = "none"; 
	go_passToJs_thisWordRowList(aWord, ""+maxNumRow5, "js_go_showWrdRowList"); 
		
} // end of onclick_require_rowListWithThisWord2


//--------------------------

function onclick_require_rowList1() {	
	
	//console.log("onclick_require_rowList1 ")
	
	document.getElementById("id_inpRowEmpty").style.display = "none";
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	//ele_wRowList.innerHTML = "";
    //ele_word.innerHTML     = ""; 	 	
	
	var inpBegRow = getInt( document.getElementById("id_fromIx_row").innerHTML );
	var numRows   = getInt( document.getElementById("id_inpNumRow" ).innerHTML );  
	var inpEndRow = inpBegRow+numRows-1;
		
	//console.log("onclick_require_rowList1 ", " XXX numRows=", numRows, " inpBegRow=" , inpBegRow, " inpEndRow=",  inpEndRow  ) 
	
	myPage01.style.display = "none"; 
	document.getElementById("id_headWord").innerHTML = ""; //head1; 
		
	var caller = "HTML page onclick_require_rowList1" //  (new Error()).stack?.split("\n")[2]?.trim().split(" ")[1] ;
	if (caller == undefined) { caller = ""; }
	go_passToJs_rowList(""+inpBegRow, ""+numRows, "js_go_rowList," + caller); 
		
} // end of onclick_require_rowList1

//--------------
function js_go_build_rowGruppi( gruppi_option) {
	// run just after the reading of the input text 
	document.getElementById("id_gruppi_sel").innerHTML = gruppi_option;  	
	html_rowGroup_index_gr = 0;  // will be updated from last file values 
	document.getElementById("id_gruppi_sel").selectedIndex = html_rowGroup_index_gr;
	 
	setLastValuesOfExtrRowChanged("js_go_build_rowGruppi");
	
} 
//--------------------------

function getInt( sInt ) {
	 
	 try {
        return parseInt( "0" + sInt.trim() );
    } catch (err) {
		return 0
	}   
} 

//-------------------------------

function onclick_wordByIndex(sIxWord) {
    // triggered by clicking on a word  
 	
    var ixWord = 0;
    try {
        ixWord = parseInt(sIxWord);
    } catch (err) {}
	
    go_passToJs_getWordByIndex(""+ixWord, "999999","js_go_showWrdRowList"); // ask 'go' to give the rows of the word  by the go function js_go...  

} // end of onclick_wordByIndex
//----------------------

//-------------------------------
function onclick_rowsByIxWord(sIxWord) {
	
	var max_num_row4word  = document.getElementById("idTabWRoW1").value  
	
    go_passToJs_getRowsByIxWord(""+sIxWord, ""+max_num_row4word, "js_go_showWrdRowList"); // ask 'go' to give the rows of the word  by the go function js_go...  

} // end of onclick_rowsByIxWord

//-------------------------------
function onclick_rowsByIxLemma(sIxLemma) {
	
	var max_num_row4lemma = document.getElementById("idTabWRoL2").value 

	console.log( green("onclick_rowsByIxLemma"), " sixLemma=", sIxLemma, " max_num_row4lemma=", max_num_row4lemma)

    go_passToJs_getRowsByIxLemma(""+sIxLemma, ""+max_num_row4lemma, "js_go_showLemmaRowList4"); // ask 'go' to give the rows of the word  by the go function js_go...  

} // end of onclick_rowsByIxLemma

//-------------------------
//-------------------------------
function OLDonclick_wordByIndex2(sIxWord, sIxLemma, swOnlyThisWordRows) {
    // triggered by clicking on a word  
	
	var max_num_row4word  = document.getElementById("idTabWRoW1").value 
	var max_num_row4lemma = document.getElementById("idTabWRoL2").value 

    var ixWord = 0;
    try {
        ixWord = parseInt(sIxWord);
    } catch (err) {}	
	
	//console.log("onclick_wordByIndex2() -->  go_passToJs_getWordByIndex2(" , "ixWord=", ixWord, "  swOnlyThisWordRows ",swOnlyThisWordRows, " maxNumRow=", maxNumRow)
	
    go_passToJs_getWordByIndex2(""+ixWord, ""+sIxLemma, swOnlyThisWordRows, ""+max_num_row4word, ""+max_num_row4lemma,"js_go_showWrdRowList"); // ask 'go' to give the rows of the word  by the go function js_go...  

} // end of OLDonclick_wordByIndex2
//-------------------------
//-------------------------
function firstUpper(str1) {	
	return str1.substring(0,1).toUpperCase() + str1.substring(1).toLowerCase(); 
}
//-------------------
function checkUpper( thisWord, thisListRow) {
	/* if the word with the first letter capitalized 
	 is found in the rows but not at the beginning of the row 
	 then probably normally is the correct way to write it (eg. a Person Name)
	*/ 
	var upperThisW = " " + firstUpper(thisWord);  
	var righe = thisListRow.split(";;"); 
	for(var z1=0; z1 < righe.length; z1++)  {
		var riga = righe[z1].trim()
		if (riga.indexOf( upperThisW ) > 0 ) return true; 
	}  	
	return false;  
}
//------------------------

function splitHeader( inpHeader ) {
	//console.log("splitHeader(", 	inpHeader); 
	/**	
		<HEADER>
			<WORD>ihren,L:ihren</WORD> 
			:lemma=ihr 		:tran=tu 		:wordsInLemma=ihren 
			:lemma=ihre 	:tran=loro		:wordsInLemma=ihren 
			:lemma=ihrer 	:tran=loro 		:wordsInLemma=ihren 
			:lemma=sein 	:tran=essere	:wordsInLemma=ihren 
		</HEADER>
	**/
	
	var wordList="", wordTab=""
	var h1 = inpHeader.indexOf("<WORD>")
	var h2 = inpHeader.indexOf("</WORD>")
	//var h3 = inpHeader.indexOf("<TABLE")
	//var h4 = inpHeader.indexOf("</TABLE>")
	var h4 = inpHeader.indexOf("</HEADER>")
	if (h2 < (h1+6)) { 
		//console.log("errore  splitHeader( inpHeader ) h1=",h1, " h2=", h2, " h3=", h3, " h4=", h4 )  
		return []
	} 
	
	if (h4 < 0) {  h4 = inpHeader.length}
	wordList = inpHeader.substring(h1+6, h2)
	wordTab = inpHeader.substring(h2+7, h4).trim(); 
	
	//wordTab =  inpHeader.substring(h3, h4+8)	
	//console.log("splitHEADER h3=", h3, " h4=", h4, ",   wordTab=", wordTab) 
	return [wordList, wordTab]
}

//---------------------------------------
function buildHeaderTable( str1 ) {	

	//console.log( green("buildHeaderTable"), " str1=\n", str1);
	/***
	buildHeaderTable () str1= 
		one Lemma, many words 
		:lemma=sein :tran=essere :wordsInLemma=bin  
		:lemma=sein :tran=essere :wordsInLemma=bist  
		:lemma=sein :tran=essere :wordsInLemma=gewesen  
		:lemma=sein :tran=essere :wordsInLemma=ihren
	
		many Lemma - One Word only
		<HEADER>
			<WORD>ihren,L:ihren</WORD> 
			:lemma=ihr 		:tran=tu 		:wordsInLemma=ihren 
			:lemma=ihre 	:tran=loro		:wordsInLemma=ihren 
			:lemma=ihrer 	:tran=loro 		:wordsInLemma=ihren 
			:lemma=sein 	:tran=essere	:wordsInLemma=ihren 
		</HEADER>
	**/
	
	
	var lineTr = str1.replaceAll("|", "<br>").replaceAll("\n", " ").split(":lemma="); 
	if (lineTr[0] == "") {  lineTr = lineTr.slice(1);}
	
	var str2 = '' ; 
	var len1 = lineTr.length
	//-------------------------------
	// caso 1: un solo lemma e una voce 
	// caso 2: un solo lemma e diverse voci 
	// caso 3: diversi lemma e una sola voce
	// caso 4: diversi lemma e diverse voci    ( non previsto )   	
	//---------------------------------------
	var oneWordOnly, oneLemmaOnly; 
	
	//-----------
	var numVoci=0
	var numLemma=0
	var preLem = "", lem1="", tran1=""
	var preVoce=""
	for(var z1=0; z1 < len1; z1++) {
		var oneTr1 = lineTr[z1].split(":")	
		lem1 = oneTr1[0].trim()
		if (lem1 != preLem) {
			numLemma++
			preLem = lem1
			tran1 = oneTr1[1].trim() 
		}
		var voce =  oneTr1[2].trim() 
		if (preVoce == "") {
			preVoce = voce; 
			numVoci++
		} else {
			if (voce != preVoce ) {
				numVoci++
				preVoce =voce; 
			}
		} 	
	} 
	oneWordOnly  = (numVoci == 1)
	oneLemmaOnly = (numLemma == 1)
	//------------------------	
	var ele_model_tSHeadW_bdy;
	var model_tSHeadW_div;	
	
	

	//--------------------------------
	var type = 0;

	if (oneWordOnly) {	
		document.getElementById("tsHead_1_3").style.display="block"; 
		ele_model_tSHeadW_bdy = document.getElementById("id_model_tSHead_bdy_13"); 	
		if (oneLemmaOnly) {
			type=1;
			//console.log( green(" CASO 1 "), "1XXXXXXXX ONE WORD Only and ONE LEMMA only XXXXXXXXXX")
		} else {
			type=3;
			//console.log( green(" CASO 3 "), "1XXXXXXXX ONE WORD and many LEMMA XXXXXXXXXX")			
		}	
	} 
	if (oneLemmaOnly) {
		if (oneWordOnly) {
			type=1;
			//console.log( green(" CASO 1 "), "2XXXXXXXX ONE WORD Only and ONE LEMMA XXXXXXXXXX")		
		} else {
			type=2;
			//console.log( green(" CASO 2 "), "1XXXXXXXX many WORD  and ONE LEMMA XXXXXXXXXX")
		}	
		document.getElementById("tsHead_2").style.display="block"; 
		ele_model_tSHeadW_bdy = document.getElementById("id_model_tSHead_bdy_2"); 	
	}
	if ((oneWordOnly == false) && (oneLemmaOnly == false)) {
		type=4;
		//console.log( green(" CASO 4 "), "3XXXXXXXX many WORD  and many LEMMA XXXXXXXXXX")
	}
	//----------------------------------
	var model_tSHeadW_lemma  = ele_model_tSHeadW_bdy.innerHTML; 
	//console.log("aaa ", model_tSHeadW_lemma)
	var model_tSHeadW13_row2 = document.getElementById("id_model_tSHead_bdy_13_row2").innerHTML; 	
	
	
	
	//-----
	switch( type ) {			
		 case 1: //console.log("caso 1  una parola e un lemma")	
			case_type1_3(); break;
		 case 2: //console.log("caso 2  un lemma diverse parole")			
			case_type2();
			break;
		 case 3: 
			//console.log("caso 3  una parola e diversi lemma")	
			case_type1_3(); 		
			break;
		 case 4: 
			//console.log("caso 4  diverse parole con diversi lemma")
			break;
		 default:
			break;
	}
	//----------------
	function case_type1_3() {
					//console.log("caso 1  una parola e un lemma")
					//console.log("caso 3  una parola e diversi lemma")				/*	
				for(var z1=0; z1 < len1; z1++) {
					var oneTr1 = lineTr[z1]	
					var jT = oneTr1.indexOf(":tran=");
					var jW = oneTr1.indexOf(":wordsInLemma=");
					var nuovoLemma   = oneTr1.substring(0,jT    ).trim();
					var nuovoTran    = oneTr1.substring(jT+6,jW ).trim();
					var nuovoLisWord = oneTr1.substring(jW+14   ).trim();						
					if (z1 == 0) {
						str2 += model_tSHeadW_lemma.replace("§1lemma§",nuovoLemma).
											replace("§1tran§",     nuovoTran).
											replace("§1wordXlem§", nuovoLisWord). 
											replace("§1tran§",     nuovoTran) + 
											"\n\n";  				
					} else {
						str2 += model_tSHeadW13_row2.replace("§1lemma§",nuovoLemma).
											replace("§1tran§",     nuovoTran) + 
											"\n\n";  		
					}			
				} // end for z1
	} // end of case_type1_3	
	//------------------
		
	function case_type2() {
			// console.log("caso 2  un lemma diverse parole")			/**
													<!-- case 2:  one lemma many words  --> 
			/*											
													<table id="tsHead_2" style="display:none;">
														<tbody id="id_model_tSHead_bdy_2">
															<tr>
																<td colspan="2">
																	<span class="c_lemma">§1lemma§</span>
																	<span class="c_tranW" style="padding-left:5em;">§1tran§</span> 
																</td> 
															</tr>															
															<tr>
																<td style="width:1em;">&nbsp;</td>
																<td class="c_word">§1wordXlem§</td>
															</tr> 
														</tbody>
													</table>
			**/
		var listParole = ""
		var nuovoLemma = "",  nuovoTran = ""
		for(var z1=0; z1 < len1; z1++) {	
			var oneTr1 = lineTr[z1]	
			var jT = oneTr1.indexOf(":tran=");
			var jW = oneTr1.indexOf(":wordsInLemma=");
			if (z1==0) {
				nuovoLemma   = oneTr1.substring(0,jT    ).trim();
				nuovoTran    = oneTr1.substring(jT+6,jW ).trim();
			}
			var nuovoLisWord = oneTr1.substring(jW+14   ).trim();		
			listParole += '<span style="margin-right:2em;">' + nuovoLisWord + "</span>"
		}
		str2 += model_tSHeadW_lemma.replace("§1lemma§", nuovoLemma).
					replace("§1tran§",     nuovoTran). 
					replace("§1wordXlem§", listParole) + "\n\n";  		
		
	} // end of case_type2
	//-----------------------------------
	
	var ele_model_tSHeadW_DIV = document.getElementById("id_model_tSHeadW"); 
	var model_tSHeadW_div = ele_model_tSHeadW_DIV.innerHTML; 	
	var jBody  = model_tSHeadW_div.indexOf("<tbody");
	var jBody2 = model_tSHeadW_div.indexOf("<tr", jBody);	
	
	var newDiv = model_tSHeadW_div.substr(0, jBody2) +"\n" + str2.trim()  + "\n</tbody></table></div>\n";  
		
	//console.log("6 buildHeaderTable () newDiv = ", newDiv)
	return newDiv;  
	
} // end of buildHeaderTable
//======================================================
//---------------------------------------
function OLD2buildHeaderTable( str1 ) {
	/***
	1 buildHeaderTable () str1= 
		:lemma=sein :tran=essere :wordsInLemma=bin  
		:lemma=sein :tran=essere :wordsInLemma=bist  
		:lemma=sein :tran=essere :wordsInLemma=gewesen  
		:lemma=sein :tran=essere :wordsInLemma=ihren
	***/
		
	/**	
		many Lemma - One Word only
		<HEADER>
			<WORD>ihren,L:ihren</WORD> 
			:lemma=ihr 		:tran=tu 		:wordsInLemma=ihren 
			:lemma=ihre 	:tran=loro		:wordsInLemma=ihren 
			:lemma=ihrer 	:tran=loro 		:wordsInLemma=ihren 
			:lemma=sein 	:tran=essere	:wordsInLemma=ihren 
		</HEADER>
		---------
		one lemma - many words		
		<HEADER>
			<WORD>gewesen,L:bin bist gewesen ihren ihrer ist sei seien sein seine seinem seinen seiner seines sind war waren warst wäre wären</WORD> 
			:lemma=sein :tran=essere :wordsInLemma=bin bist gewesen ihren ihrer ist sei seien sein seine seinem seinen seiner seines sind war waren warst wäre wären 
		</HEADER>
	**/
	
	var lineTr = str1.replaceAll("|", "<br>").replaceAll("\n", " ").split(":lemma="); 
	if (lineTr[0] == "") {  lineTr = lineTr.slice(1);}
	
	console.log("1 buildHeaderTable () str1=" , str1); 
 	
		
	//console.log("2 buildHeaderTable () model_tSHeadW_div=> " + model_tSHeadW_div  + "<==="); 
	
	/**********
	
		<table id="tsHead1" style="display:none;">
			<tbody id="id_model_tSHeadW_bdy1">
				<tr><td colspan="2" class="c_lemma">§1lemma§</td> </tr>															
				<tr><td style="width:1em;">&nbsp;</td><td class="c_word">§1wordXlem§</td></tr> 
				<tr><td style="width:1em;">&nbsp;</td><td class="c_tranW">§1tran§</td></tr> 
			</tbody>
		</table>
		
		<table id="tsHead2" style="display:none;">
			<tbody id="id_model_tSHeadW2_bdy">
				<tr><td colspan="2" class="c_word">§1wordXlem§</td> </tr>															
				<tr><td style="width:1em;">&nbsp;</td><td class="c_lemma">§1lemma§</td>
								<td class="c_tranW">§1tran§</td></tr> 
			</tbody>
		</table>
	
	
	
	********/
	
	
	/**
	var ele_model_tSHeadW_bdy = document.getElementById("id_model_tSHeadW_bdy"); 
	var model_tSHeadW_lemma = ele_model_tSHeadW_bdy.innerHTML; 
	
	
	
	var jBody  = model_tSHeadW_div.indexOf("<tbody "); 
	 
	var jBodyEnd = model_tSHeadW_div.indexOf("</tbody>", jBody); 
	**/

	
	var str2 = '' ; 
	var len1 = lineTr.length
	//-------------------------------
	// caso 1: un solo lemma e una voce 
	// caso 2: un solo lemma e diverse voci 
	// caso 3: diversi lemma e una sola voce
	// caso 4: diversi lemma e diverse voci    ( non previsto )   	
	//---------------------------------------
	var oneWordOnly = true; 
	var ele_model_tSHeadW_bdy;
	//-----------
	for(var z1=0; z1 < len1; z1++) {
		var oneTr1 = lineTr[z1]	
		var jW = oneTr1.indexOf(":wordsInLemma=");
		var nuovoLisWord = oneTr1.substring(jW+14   ).trim();				
		var wor1arr = nuovoLisWord.split("<br>");
		if (wor1arr.length > 1) { 
			oneWordOnly = false;			
		}	
	} // end for z1
	//------------------------	

	var	ele_model_tSHeadW_bdy 
	var model_tSHeadW_div
	if (oneWordOnly) {	
		console.log("XXXXXXXX oneWordOnly = true  XXXXXXXXXX")
		document.getElementById("tsHead2").style.display="block"; 
		ele_model_tSHeadW_bdy = document.getElementById("id_model_tSHeadW2_bdy"); 				
	} else {
		document.getElementById("tsHead1").style.display="block"; 
		ele_model_tSHeadW_bdy = document.getElementById("id_model_tSHeadW1_bdy"); 	
	}
	

	
	//var model_tSHeadW_div = ele_model_tSHeadW.innerHTML; 

	
	var model_tSHeadW_lemma  = ele_model_tSHeadW_bdy.innerHTML; 
	var model_tSHeadW2_row2 = document.getElementById("id_model_tSHeadW2_row2_bdy").innerHTML; 	
	
	//console.log("XXXXXXXX  model_tSHeadW_lemma = ",   model_tSHeadW_lemma )
	
	//var ele_model_tSHeadW_bdyInner = ele_model_tSHeadW_bdy.innerHTML
	
	//var	jBody    = ele_model_tSHeadW_bdy.innerHTMLele_model_tSHeadW_bdymodel_tSHeadW_div.indexOf("<tbody ");	 
	//var jBodyEnd = model_tSHeadW_div.indexOf("</tbody>", jBody);  
	//----------------------------
	for(var z1=0; z1 < len1; z1++) {
		var oneTr1 = lineTr[z1]	
		var jT = oneTr1.indexOf(":tran=");
		var jW = oneTr1.indexOf(":wordsInLemma=");
		var nuovoLemma   = oneTr1.substring(0,jT    ).trim();
		var nuovoTran    = oneTr1.substring(jT+6,jW ).trim();
		var nuovoLisWord = oneTr1.substring(jW+14   ).trim();		
			
		var wor1arr = nuovoLisWord.split("<br>");
		var newLe2=""; 
		if (wor1arr.length == 1) { 
			var wor11 = wor1arr[0].trim(); 
			var lem1arr = nuovoLemma.split("<br>")    ;
			for(var h1=0; h1 < lem1arr.length; h1++) {
				if (lem1arr[h1].trim() == wor11) {continue;}
				newLe2 += "<br>" + lem1arr[h1].trim() ; 
			}
			if (newLe2 != "") {
				newLe2 = newLe2.substr(4); 
			}
			nuovoLemma = newLe2; 
		} 	
		if (oneWordOnly) {
			if (z1 == 0) {
				str2 += model_tSHeadW_lemma.replace("§1lemma§",nuovoLemma).
									replace("§1tran§",     nuovoTran).
									replace("§1wordXlem§", nuovoLisWord) + 
									"\n\n";  				
			} else {
				str2 += model_tSHeadW2_row2.replace("§1lemma§",nuovoLemma).
									replace("§1tran§",     nuovoTran) + 
									"\n\n";  		
			}
			continue		
		}

		if ((nuovoLemma == nuovoLisWord) && (len1==1)) {
			str2 += model_tSHeadW_lemma.replace("§1lemma§","").
									replace("§1tran§",     nuovoTran).
									replace("§1wordXlem§", nuovoLisWord) + 
									"\n\n";  
		} else {
			str2 += model_tSHeadW_lemma.replace("§1lemma§",nuovoLemma).
									replace("§1tran§",     nuovoTran).
									replace("§1wordXlem§", nuovoLisWord) + 
									"\n\n";  
		}							
	} // end for z1
	//------------
	var ele_model_tSHeadW_DIV = document.getElementById("id_model_tSHeadW"); 
	var model_tSHeadW_div = ele_model_tSHeadW_DIV.innerHTML; 	
	var jBody  = model_tSHeadW_div.indexOf("<tbody");
	var jBody2 = model_tSHeadW_div.indexOf("<tr", jBody);	
	
	var newDiv = model_tSHeadW_div.substr(0, jBody2) +"\n" + str2.trim()  + "\n</tbody></table></div>\n";  
		
	console.log("6 buildHeaderTable () newDiv = ", newDiv)
	
	return newDiv;  
	
} // end of OLD2buildHeaderTable

//---------------------------------------
function OLDbuildHeaderTable( str1 ) {
	/**	
		<HEADER>
			<WORD>ihren,L:ihren</WORD> 
			:lemma=ihr 		:tran=tu 		:wordsInLemma=ihren 
			:lemma=ihre 	:tran=loro		:wordsInLemma=ihren 
			:lemma=ihrer 	:tran=loro 		:wordsInLemma=ihren 
			:lemma=sein 	:tran=essere	:wordsInLemma=ihren 
		</HEADER>
	**/
	
	var lineTr = str1.replaceAll("|", "<br>").replaceAll("\n", " ").split(":lemma="); 
	if (lineTr[0] == "") {  lineTr = lineTr.slice(1);}
	
	console.log("1 buildHeaderTable () str1=" , str1); 
 	
	var ele_model_tSHeadW = document.getElementById("id_model_tSHeadW"); 
	var model_tSHeadW_div = ele_model_tSHeadW.innerHTML; 
	
	//console.log("2 buildHeaderTable () model_tSHeadW_div=> " + model_tSHeadW_div  + "<==="); 
	
	var ele_model_tSHeadW_bdy = document.getElementById("id_model_tSHeadW_bdy"); 
	var model_tSHeadW_lemma = ele_model_tSHeadW_bdy.innerHTML; 
	
	
	
	var jBody  = model_tSHeadW_div.indexOf("<tbody "); 
	 
	var jBodyEnd = model_tSHeadW_div.indexOf("</tbody>", jBody); 
	

	
	var str2 = '' ; 
	var len1 = lineTr.length
	//-------------------------------
	// caso 1: un solo lemma e una voce 
	// caso 2: un solo lemma e diverse voci 
	// caso 3: diversi lemma e una sola voce
	// caso 4: diversi lemma e diverse voci    ( non previsto )   	
	//---------------------------------------
	for(var z1=0; z1 < len1; z1++) {
		var oneTr1 = lineTr[z1]
	
		var jT = oneTr1.indexOf(":tran=");
		var jW = oneTr1.indexOf(":wordsInLemma=");
		var nuovoLemma   = oneTr1.substring(0,jT    ).trim();
		var nuovoTran    = oneTr1.substring(jT+6,jW ).trim();
		var nuovoLisWord = oneTr1.substring(jW+14   ).trim();		
			
		var wor1arr = nuovoLisWord.split("<br>") ;
		var newLe2=""; 
		if (wor1arr.length == 1) {
			var wor11 = wor1arr[0].trim(); 
			var lem1arr = nuovoLemma.split("<br>")    ;
			for(var h1=0; h1 < lem1arr.length; h1++) {
				if (lem1arr[h1].trim() == wor11) {continue;}
				newLe2 += "<br>" + lem1arr[h1].trim() ; 
			}
			if (newLe2 != "") {
				newLe2 = newLe2.substr(4); 
			}
			nuovoLemma = newLe2; 
		}
			
		if ((nuovoLemma == nuovoLisWord) && (len1==1)) {
			str2 += model_tSHeadW_lemma.replace("§1lemma§","").
									replace("§1tran§",     nuovoTran).
									replace("§1wordXlem§", nuovoLisWord) + 
									"\n\n";  
		} else {
			str2 += model_tSHeadW_lemma.replace("§1lemma§",nuovoLemma).
									replace("§1tran§",     nuovoTran).
									replace("§1wordXlem§", nuovoLisWord) + 
									"\n\n";  
		}							
	}
	str2 += '';
	
	//console.log("6 buildHeaderTable () str2 = " + str2 + "<==" ) ; 

	
	var newDiv = model_tSHeadW_div.substr(0, jBody) + str2 + model_tSHeadW_div.substr(jBodyEnd+8);  
	
	return newDiv;  
	
} // end of OLDbuildHeaderTable

//--------------------------------------------------
function js_go_showWordRowList3(inpstr) {
	// vedi js_go_showWrdRowList(inpstr)
	console.log(" js_go_showWordRowList2 (inpstr=" + inpstr);  
}
//--------------------------------------------------
function js_go_showLemmaRowList4(inpstr) {
	// vedi js_go_showWrdRowList(inpstr)	
	//console.log(" js_go_showLemmaRowList2 (inpstr=" + inpstr);  
	js_go_showWrdRowList(inpstr)
}
//------------------------
function js_go_showWrdRowList(inpstr) {

	// triggered by go ( bild go_passToJs_getWordByIndex )
	
	//console.log(" js_go_showWrdRowList (inpstr=" + inpstr);  
	
    if (inpstr == undefined) {
		//console.log(" js_go_showWrdRowList () 1 return inpstr undefined ");  
		onclick_jumpFromToPage( myPage02,myPage03, myPage01);  //   
        return;
    }
    if (inpstr == "") {		
		//console.log(" js_go_showWrdRowList () 2 return inpstr vuoto");  
		onclick_jumpFromToPage( myPage02,myPage03, myPage01);    
        return;
    }
	
	if (inpstr.substring(0,5) == "NONE,") {
		console.log(" js_go_showWrdRowList () 3 return inpstr = " +inpstr);  
		document.getElementById("id_inpWordFra_msgWord").innerHTML = inpstr.substring(5) ;
		//document.getElementById("id_inpWordFra_msg").style.display = "block";
		myPage01.style.display = "flex"; 
		myPage05.style.display = "none";
		//onclick_jumpFromToPage( myPage02,myPage03, myPage04);  //   
		return
	}	
	//document.getElementById("id_inpWordFra_msg").style.display = "none";
	
	myPage05.style.display = "none";
	
	
	
	
	var h_wordListStr = "", h_wordTab = "";
	var ks = inpstr.indexOf("</HEADER>"); 
	if (ks < 0) { return } 	
	
	var inpHeader = inpstr.substring( 0, ks + 9);
	
	//console.log("js_go_showWrdRowList inpHeader=\n" + inpHeader +"\n-----------------\n") 
	/**	
		<HEADER>
			<WORD>ihren,L:ihren</WORD> 
			:lemma=ihr 		:tran=tu 		:wordsInLemma=ihren 
			:lemma=ihre 	:tran=loro		:wordsInLemma=ihren 
			:lemma=ihrer 	:tran=loro 		:wordsInLemma=ihren 
			:lemma=sein 	:tran=essere	:wordsInLemma=ihren 
		</HEADER>
	**/
	
	
	var thisListRow = inpstr.substring(ks+9) ; 
	//console.log("ANTONIO resto \n", inpstr	, "\n------------------------------------ fine ----")
	
	var col1 = splitHeader( inpHeader );
	var h_wordListStr00 = col1[0]
	
	var inpReqWord = "";
	var jh = h_wordListStr00.indexOf(",L:");	
	if (jh < 1) {
		h_wordListStr = h_wordListStr00; 
	} else {
		inpReqWord    = h_wordListStr00.substring(0,jh).trim(); 
		h_wordListStr = h_wordListStr00.substring(jh+3).trim();		
	}	
		
	//console.log("ANTONIO _showWordRowList ", "inpReqWord=" + inpReqWord + ",h_wordListStr=" + h_wordListStr + "\n-------------------------------------\n") ;
	
	h_wordTab     = buildHeaderTable( col1[1] )
	
	
	//-------------------
	
	word_to_underline_list = h_wordListStr.split(" ")                        
	
	//console.log("1word_to_underline_list=", word_to_underline_list) 
	
	var word3, ixUnW3, totRow3, wLemma3, wTran3;

	
	
	document.getElementById("id_headWord").innerHTML = h_wordTab.replaceAll('display:none','display:block').replaceAll("tsHead1","tsHead00") ;
	
	onclick_jumpFromToPage( myPage02,myPage03, myPage04);  //   
	
	js_go_rowList( thisListRow );  
	
	
} // end of js_go_showWrdRowList  NEW
//--------------------------------------------

function js_go_rowList2( inpstr )  {

} // end of  js_go_rowList2

//--------------------------------------
function js_go_rowList( inpstr, js_parm, jsFunc,goFunc) {
	
	// triggered by go ( go_passToJs_rowList and js_go_showWrdRowList)
	
	//console.log("function js_go_rowList() <-- " + goFunc + " <-- " + jsFunc ) 
	//console.log("inpstr=" +inpstr ) 
	
	rowToStudy_list = [];
	newRowTran = [];
	var numeroTS_Row=0, numeroTS_OkTran=0, numeroTS_NoTran=0;
    if (inpstr == undefined) {
		console.log("js_go_rowList() 1 return inpstr undefined "); 	
		onclick_jumpFromToPage( myPage02,myPage03, myPage01);  //   
        return;
    }
	
    if (inpstr == "") {		
		console.log("js_go_rowList() 2 return inpstr vuoto");  
		document.getElementById("id_inpRowEmpty").style.display = "inline-block";
		onclick_jumpFromToPage( myPage02,myPage03, myPage01);  //   
	    return;
    }
	
	rowToStudy_list =  inpstr.split("<br>");	
	
	//console.log("js_go_rowList() 1 rowToStudy_list.length=", rowToStudy_list.length , "  maxNumRow=",  maxNumRow , " type=", typeof maxNumRow);
	
	if (rowToStudy_list.length >= maxNumRow) {
		rowToStudy_list = rowToStudy_list.slice(0, maxNumRow+1) ; 
	}  
	
	//console.log("js_go_rowList() 2 rowToStudy_list.length=", rowToStudy_list.length );
	
	for (var z=0; z < rowToStudy_list.length; z++) {	
		newRowTran.push( 0 )
	}
		
	[numeroTS_Row, numeroTS_OkTran, numeroTS_NoTran] = build_Page1_rowsToTranslate("3js_go_rowList"); 
	
	//console.log("js_go_rowList() numeroTS=", numeroTS_Row, " numeroTS_OkTran=", numeroTS_OkTran, " numeroTS_NoTran=", numeroTS_NoTran ) ;
	
	
} // end of js_go_rowList
//-------------------

function build_Page1_rowsToTranslate( wh ) {
	
	
	var numeroTS_Row=0, numeroTS_OkTran=0, numeroTS_NoTran=0;
	var numRowNoTranR1 = 0 
	var rows_to_translate_str = "";  
	
	var nfile,idRow, ixRow,p3,rowS, tranS, nfileS, idRowS, ixRowS, oT; 
	
	//console.log("\nXXXXXXXXXXXXXXXXXXXXXXXX   build_Page1_rowsToTranslate (", wh,")",  "  1  length=", rowToStudy_list.length )
	
	for (var z=0; z < rowToStudy_list.length; z++) {		
		/*
		;;0;;2;; Erstes Kapitel;; Primo Capitolo 
		;;0;;3;; 
		;;0;;4;; Gustav Aschenbach oder von Aschenbach, wie seit seinem fuenfzigsten;; Gustav Aschenbach o von Aschenbach, come ha fatto fin dai cinquant'anni
		;;0;;5;; Geburtstag amtlich sein Name lautete, hatte an einem;;Il suo compleanno ufficiale cadeva l'una		
		*/
		oT = rowToStudy_list[z]
		
		//console.log( "build_Page1_rowsToTranslate (", wh, ") z=", z, " rowToStudy_list[z] m=" + oT + "<==") 
		
		oT = oT.trim() + "|||||";	
		
		[nfileS,idRowS,ixRowS, rowS, tranS] =  oT.split("|") ;  // eg. ;;0;2;; Primo capitolo	
		rowS  = rowS.trim()
		tranS = tranS.trim()
		
		if (rowS == "") {  continue;}
				
		//console.log("build_Page1_rowsToTranslate ( z=", z, "  rowToStudy_list[]: " , " nfileS=",nfileS, " idRows=", idRowS, " ixRows=", ixRowS, " rows=",rowS, " tranS=", tranS)   
		
		numeroTS_Row++; 
		if (tranS == "") {
			numeroTS_NoTran++; 
			numRowNoTranR1++
			rows_to_translate_str += z + ";;" + ixRowS + ";; " + rowS + "\n"; // punto e virgola come separatore nella speranza che il traduttore google non colleghi campi diversi; metto spazio prima di rowS per cercare di evitare problemi nelle traduzione con google 
			//console.log("\t build_Page1_rowsToTranslate () 2.2 rows_to_translate_str = ",  z + ";" + ixRowS + ";" + rowS + "<== MANCA TRANS XXXXXXXX\n" )
		} else {
			numeroTS_OkTran++; 
			//console.log("\t build_Page1_rowsToTranslate () 2.3 tranS=" + tranS + "<==  TROVATO ") 			
		}
	}
	//------------------	
	if (numRowNoTranR1 < 1) {		
		showRowsAndTranButton("2")
		return [numeroTS_Row, numeroTS_OkTran, numeroTS_NoTran ];
	}	
	document.getElementById("id_notTranNumRow").innerHTML = numRowNoTranR1; 
	
	ele_toTranslate_textarea.value = rows_to_translate_str;
	ele_translated_textarea.value = ""; 
	
	//console.log("build... ele_toTranslate_textarea = ", rows_to_translate_str); 
	
	//onclick_jumpFromToPage( myPage01,0,myPage04);   
	onclick_jumpFromToPage( myPage01,0,myPage04);   
	
	return [numeroTS_Row, numeroTS_OkTran, numeroTS_NoTran ];
	
} // end of build_Page1_rowsToTranslate	

//-------------------------------------------------

function js_go_showReadFile( str1 ) {
	//console.log("ANTONIO js_go_showReadFile( str1=" , str1 )
	var j1 = str1.indexOf( "))" ); 
	var mainNum     = str1.substring(0,j1);
	var fileListStr = str1.substring(j1+2); 
	
	//  "level " + msgLevelStat + "))" 	
	var numUniW, numTotW, numRow, levelStats = "";
	[numUniW, numTotW, numRow, levelStats] = mainNum.split(";")
	numberOf_uniW = numUniW;
	numberOf_totW = numTotW;
	numberOf_Row  = numRow;
	str1 = parseInt(numUniW).toLocaleString() + " parole diverse" + 
			"<br>in totale " + parseInt(numTotW).toLocaleString() + " parole " + 
			"su un testo di " + parseInt(numRow).toLocaleString() + " righe"  ;

	//console.log("js_go_showReadFile  levelStats=", levelStats)
	
	if (levelStats.indexOf("-oth-: 99%") < 0)  {		
		str1 += "<br>" + levelStats; 			
	}
	
    var rows = fileListStr.split(";");
    var td1, td2;
    //var modelTR = document.getElementById("id_start_trModel").outerHTML;
	str1 += "<hr>"
	var str2 = '<table style="border:0px solid black">\n';
	//str2 += '<tr><th colspan="2">input</th></tr> \n' ;   	
    for (var i = 0; i < rows.length; i++) {
        if (rows[i] == "") continue;
        [td1, td2] = rows[i].split("<file>");			
		var k1= td2.lastIndexOf("\\")
		var k2= td2.lastIndexOf("/")	
		var k3 = Math.max(k1,k2) ;
		listaInputFile = td2.substring(k3+1).trim() + "," 
		str2 += '<tr><td>' + td2.substring(k3+1) + '</td>' +
			'<td>(' + parseInt(td1).toLocaleString() + " righe" + ')' + '</td></tr> \n' ;  
    }
	str2 += '</table>';
	
	var str0 = '<div style="border:1px solid black; padding:1em;">';
	
	document.getElementById("id_inpFileList").innerHTML = str0 + str1 + ""+ str2 + "</div>";
	
}
//---------------
function set_dragDiv() {
		dragElement(document.getElementById("dragButt1"  ), document.getElementById("dragButt1_header"  ));
		dragElement(document.getElementById("dragButt2"  ), document.getElementById("dragButt2_header"  ));
		dragElement(document.getElementById("dragButt3"  ), document.getElementById("dragButt3_header"  ));
		dragElement(document.getElementById("dragButt4"  ), document.getElementById("dragButt4_header"  ));
		dragElement(document.getElementById("dragWLsTip5"), document.getElementById("dragWLsTip5_Header")); 
}
//------------------------------------
function js_go_ready( prevRun00) {
	//console.log("************************* js_go_ready(" + prevRun00.trim() + ")" ); 
	
	/**
	5,de,de-DE,Microsoft Stefan - German (Germany):mainpage_value=5491,10,1,100,any,,
	**/
	
	set_dragDiv();
	
	prevRun00+="                                   ";
	
	var jj = prevRun00.indexOf(":mainpage_value=");
	var lastMainPageValueS;
	var prevRun = "";
	if (jj < 0) {
		lastMainPageValueS = "";	
		prevRun = prevRun00;
	} else {		
		lastMainPageValueS = prevRun00.substr(jj+16);  	
		prevRun = prevRun00.substr(0,jj)
	}
	var sel_lev  = ""; var sel_ix_lev=0;  var sel_id_lev="";
	var sel_extr = ""; var sel_ix_extr=0; var sel_id_extr="";
	var isAlphaStr=""; 
	
	var ele_sel_1_lev    = document.getElementById("id_sel_1_levTOLTO");
	var ele_sel_2extRow  = document.getElementById("id_sel_2_extrRow");
	 	//-----------	
	
	
	
	if (sel_lev == "" ) {
		//sel_id_lev = ele_sel_1_lev.options[0].id;
		//sel_ix_lev=0; 
	} else {
		//sel_id_lev = ele_sel_1_lev.options.namedItem( sel_lev ).id;
		//sel_ix_lev = ele_sel_1_lev.options.namedItem( sel_lev ).index;
	}
	//ele_sel_1_lev.selectedIndex = sel_ix_lev; 
	
	//console.log("sel_ix_lev=" , sel_ix_lev)
	
	//------------	
	if (sel_extr == "" ) {
		sel_id_extr = ele_sel_2extRow.options[0].id;
		sel_ix_extr=0; 
	} else {
		sel_id_extr = ele_sel_2extRow.options.namedItem( sel_extr ).id;
		sel_ix_extr = ele_sel_2extRow.options.namedItem( sel_extr ).index;
	}
	ele_sel_2extRow.selectedIndex = sel_ix_extr; 
	
	is_selected_row_only = (sel_ix_extr == index_onlySelRowsWanted); //3js_go_ready
	fun_selRowsWanted_changed();
	//console.log("go_ready:  is_selected_row_only = ",is_selected_row_only ); 
	
	//-------
	/***
	if (isAlphaStr == "alpha") {
		ele_alpha.checked = true;
	} else {
		ele_freq.checked  = true;
	}
	var eleSele = document.getElementById("id_orderWord1TOLTO");	
	var swIsAl = ( eleSele.selectedIndex == 0) 
	***/
	/***
	var eleSele = document.getElementById("id_orderWord1TOLTO");	
	if (isAlphaStr == "alpha") {
		eleSele.selectedIndex = 0;  // alphabetic 
	} else {
		eleSele.selectedIndex = 1;  // by frequence 
	}
	***/
	//-------------------
	
	var prevRunLanguage = prevRun.trim(); 
	if (prevRunLanguage != "") {
		sw_firstDictLine_already_existed = true; 
		lastRunLanguage = prevRunLanguage
		//console.log("language file has been read ==>" +  lastRunLanguage) 
	} 	
  
    document.getElementById("id_start001").style.display = "none";
    myPage01.style.display = "flex";
	
    document.getElementById("id_showButt").style.display = "block";
    document.getElementById("id_start_tab").style.display = "none";
	
	if (prevRunLanguage != "") { 
		lastRunLanguage = prevRunLanguage
		loadPrevLang( prevRunLanguage ) 
		//console.log("js_go_ready()  prevRunLanguage=", prevRunLanguage, "  js_go_ready() NON chiama  fcommon_load_all_voices()");  
	}	else {
		//console.log("js_go_ready() call fcommon_load_all_voices()");
		
		fcommon_load_all_voices(); // at end calls tts_1_toBeRunAfterGotVoices()		
		// WARNING: the above function contains asynchronous code.  
		// 			Any statement after this line is executed immediately without waiting its end			
	}
	
	
	//onclick_getRowGroup(  document.getElementById("id_gruppi_sel") )
	
	
} // end of js_go_ready
//-------------------------------------

function onclick_show_or_hide_statistics() {
    var eleFreq   = document.getElementById("id_frequenze");
    var eleStFile = document.getElementById("id_start_tab");


    if (eleFreq.style.display == "block") {
        eleFreq.style.display   = "none";
        eleStFile.style.display = "none";
    } else {
        eleFreq.style.display   = "block";
        eleStFile.style.display = "block";
    }
} // end of onclick_require_statistics

//----------------------------------------
function stdCode(inpCode) {	

	var CoerInp = inpCode.replaceAll( "ae","ä").replaceAll("oe","ö").replaceAll("ue","ü").replaceAll("ß","ss") 
					
	return CoerInp  
}
//----------------------------

function evidenzia( unaparola, txtinp1) {
	var newRow="";
	unaparola = unaparola.toLowerCase().trim();
	
	var lenParola = unaparola.length; 
	
	var lowinp0 = txtinp1.toLowerCase().replaceAll("§"," ") + "§"; 
	var lowinp1 = lowinp0.replace(/[\s;,:"'\.<>»«()\[\]\!\?„“]/g,"§")  
	
	var j0=-1, j1=0;
	var jNew=0
	var swBold=false;
	var parola1, parolaT;
	
	for(var i=0; i < lowinp1.length; i++) {	    
		j0=j1+1
		j1 = lowinp1.indexOf("§", j0);  
		if (j1 < 0) break;
		parola1 = lowinp1.substring(j0,j1) 
		parolaT = stdCode(parola1)
		newRow += txtinp1.substring(jNew, j0) 	
		swBold = (parolaT == unaparola) 
		if (swBold) newRow += '<span class="c_wordTarg">'
		newRow +=  txtinp1.substring(j0, j1)
		jNew = j1;	
		if (swBold) newRow += '</span>' 
	}	

	return newRow; 
	
} // end of evidenzia 
//------------------------------
//----------------------------------------------------


//----------------------------------------------------
function onclick_jumpFromToPage( fromPage1, fromPage2, toPage, blockFlex = "flex") {
	//console.log("onclick_jumpFromToPage()" + "fromPage=", fromPage1,  " toPage=" , toPage, " blockFlex=", blockFlex);  	
	fromPage1.style.display = "none"; 
	if (fromPage2 != 0) {	fromPage2.style.display = "none"; }
	try {
		toPage.style.display = blockFlex;
	} catch(e1) {
		console.log("onclick_jumpFromToPage()" + " toPage=" , toPage);  
			console.log(e1);
	}
}

//----------------------------------------------------
function onclick_jumpFromTo1_2Page( fromPage1, fromPage2, toPage, toPage2, blockFlex = "flex") {
	
	fromPage1.style.display = "none"; 
	if (fromPage2 != 0) {	fromPage2.style.display = "none"; }
	
	if(ele_wordList.innerHTML == "") {
		try {
			toPage2.style.display = blockFlex;
		} catch(e1) {
			console.log("onclick_jumpFromToPage()" + " toPage=" , toPage2);  
			console.log(e1);
		}
		
	} else {	
		try {
			toPage.style.display = blockFlex;
		} catch(e1) {
			console.log("onclick_jumpFromToPage()" + " toPage=" , toPage);  
			console.log(e1);
		}
	}
}
//--------------------------------------------------

function onclick_copyTextAreaValue_to_clipboard( this1 ) {
	// copy textarea value to clipboard (from where you can paste)  
	this1.select();
	this1.setSelectionRange(0, 99999);
	navigator.clipboard.writeText(this1.value);  
}
//--------------------------------------------------------
function copyInnerHTML_to_clipboard( this1 , ele_textarea) {
	// copy innerHTML to textarea value and then ask to copy from that to clipboard (from where you can paste) 
	// textarea might be with display none  if you need 
	ele_textarea.value = this1.innerHTML.replaceAll("<br>","\n") ;
	onclick_copyTextAreaValue_to_clipboard( document.getElementById('myInput') );
}
//--------------------------------------

function onclick_showWordsButton(type) {
	
	console.log("antonio prova", "lorca", "con tre parole") 	

	// if all words are translated continue 
	// is some translations are missing, ask if they must be ignored, otherwise loop till no translations are missing     
	if (type == 1) {
		sw_ignore_missTranWord = true;  // Ignore missing translations and continue
	} else {
		sw_ignore_missTranWord = false; // Translations added, check again</button	
	}  	
	
	var wordTranList = ele_wordsTranslated.value.trim().split( wordTTEnd );    // 
	
	//console.log("X3 onclick_showWordsButton()  wordTranList=\n", wordTranList ) 	
	
	var lenTran =  wordTranList.length;	
	var wordTran
		
	var ixUnTrad, ixUnWtS;
	var lenW = wordToStudy_list.length;
	var word1, nrow, totExtrRow2, totExtrRow2, wLemma1, wtran, knowYesCtr, knowNoCtr   ; 
	var newTran_f = "";
	//------------------------------------------
	var wX, wT, ixz1,  ixz2, ixzNum   
	
	//-----------
	var wList;
	var sw_someTranMissingW = false;  // sometimes the automatic translator does some mistakes 
	var wIxLemmaList, wLemmaList, wTranList, newTranList , wLevelList, wParaList, wExampleList, knowYesCtr, knowNoCtr; 
	var numf=0
	var sw_Minus1 = false
	//----------------------------
	for(var z=0; z < lenTran; z++) {
	
		wT = wordTranList[z].toLowerCase(); 	
		//                             z + ";" + ixUnW2 + ";" + ixLemma + "; " +wTran    ( soltanto UNA traduzione (quella del lemma num.f 
		if (wT == "") {continue;}
		if (wT.substring(0,1) != wordTTBegin) {continue}

		
		
		wList =  wT.substring(1).split(";") ;
		
		//console.log("X3.0 onclick_showWordsButton() z=",z, " wT=", wT , " XXX  wList=", wList) 
		
		if (wList.length < 4) { 
				//console.log("onclick_showWordsButton() errorT  entry z=" + z + " = " + wT ); 
				continue;
		}
		[ixzNum, ixUnTrad, numf, newTran_f] = wList ;              //   ix: translation word	
		
		//console.log("traduz=", wT, "\n\t ixzNum=", ixzNum, "  ixUnTrad=", ixUnTrad, " numf=", numf, " newTran_f=", newTran_f ) // 
		
		//ixzNum is the index of the word in wordToStudy_list		
		try{
			var ix3 = parseInt( ixzNum );
		} catch(e1) {
			continue;			
		}
		if (wordToStudy_list) {
			if (ix3 >= wordToStudy_list.length) {continue;}
		} else {
			continue;
		}
		//	[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList ] ???antoX
		[word1, ixUnWtS, nrow, wLemmaList,   wTranList, wLevelList, 
				wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList  ] = wordToStudy_list[ixzNum]; 
		
		//console.log("\n------------------------\nwordToStudy_list[ixzNum=", ixzNum,"]= ",  wordToStudy_list[ixzNum] , " ==> ", [word1, ixUnWtS, nrow, wLemmaList, wTranList] ) 
		
		//console.log(" wordTranList[z]=", wT, " ===> wList= [ixzNum, ixUnTrad, numf, newTran_f]=", [ixzNum, ixUnTrad, numf, newTran_f] ); 
		
		var numIxTrad= parseInt(ixUnTrad)
		if (numIxTrad == -1) { sw_Minus1 = true; numIxTrad = 0; }
		if ( parseInt(ixUnWtS) != numIxTrad) {
			sw_someTranMissingW= true;	
			if (sw_ignore_missTranWord == false) {
				console.log("onclick_showWordsButton() ", red("error4w"), " entry z=" + z + " = " + wT  + "\n\t wordToStudy_list[ixzNum="+ ixzNum +"]=" +  wordToStudy_list[ixzNum] +  
					"\n\tixUnWtS=" + ixUnWtS + " ixUnTrad=" + ixUnTrad);			
				continue; 
			}	
		} 
		var numLemma = parseInt( numf ); // index of the lemma and its translation in the [lemmalist][tran list] in wordToStudy_list 
		newTran[ixzNum] = 1; 
		
		//console.log("anto newTran[ixzNum=" + ixzNum + "] = 1" ); 
		
		for(var h = 0 ; h < wLemmaList.length; h++) {	
			if (h == numLemma) {
				wTranList[h] = newTran_f.trim();
			} 
		}
		
		//wordToStudy_list[ixzNum] = [word1 , ixUnWtS, nrow, wLemmaList, newTranList.substring(1) ;   // update  element  
		if (ixUnTrad == "-1") { 
			wordToStudy_list[ixzNum][1] = ixUnTrad
		}
		wordToStudy_list[ixzNum][4] = wTranList;    
		
		//console.log(" new wordToStudy_list[ixzNum]=", 	wordToStudy_list[ixzNum] ); 
		
	}  // end of for(var z ...
	//--------------------------	
	if (sw_ignore_missTranWord == false) {
		if (sw_someTranMissingW) {		
			//console.log("onclick_showWordsButton() return 1 sw_someTranMissingW=true") 
			fun_showWordList("1")
			return;
		}		
	}
	//console.log("onclick_showWordsButton() wordToStudy_list.length=" , wordToStudy_list.length)
	for (var i = 0; i < wordToStudy_list.length; i++) {
		//	[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList ] ???antoX
		[word1, ixUnWtS, nrow,   wLemmaList, wTranList, 
					wLevelList, wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr , wIxLemmaList  ] = wordToStudy_list[i];  
		
		if (word1 == "") { continue; }
				
		for(var ixLemma = 0 ; ixLemma < wLemmaList.length; ixLemma++) {	
			if (wTranList[ixLemma] == "") {	
				sw_someTranMissingW = true;  
				if (sw_ignore_missTranWord == false) {
					fun_showWordList("2", i)
					return;	
				}				
			}
		}
	}	
	
	//console.log("onclick_showWordsButton() call write word dictionary ") 
	
	write_word_dictionary(); 
	
	//console.log("onclick_showWordsButton() ",red("call write word dictionary ")) 
	
	if (sw_Minus1) {
		for (var i = 0; i < wordToStudy_list.length; i++) {  
			// [word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList ] ???antoX
			[word1, ixUnWtS, nrow,   wLemmaList, wTranList, wLevelList, 
						wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr , wIxLemmaList ] = wordToStudy_list[i];  
			if (ixUnWtS < 0) {
				//[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList ] ???antoX
				wordToStudy_list[i] = [word1, 0, nrow,   wLemmaList, wTranList, wLevelList, 
							wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList  ] 
			}
		}
	}
	
	showWordsAndTranButton("3");
	
} // end of onclick_showWordsButton()
//---------------------------


//------------------------
function extract_ix_word( i,wordOrig2 ) {
		// var wordOrig2 = wordOrigList[i].trim(); 		
		var pp= wordOrig2.indexOf(";"); 
		
		//console.log("antonio extract_ix_word( i=" + i + "  wordOrig2=" + wordOrig2 + " ==> pp=" + pp)  
		
		if (pp< 0) { return "";}
	
		var ix2 = wordOrig2.substring(0, pp); 
		var	ww2 = wordOrig2.substring(pp+1).trim(); 		
		
		//console.log("\t\t i=" + i + "   ix2=" + ix2 + "   ww2=" + ww2)  
		
		try{
			var ix3 = parseInt( ix2 );
			if (ix3 == i) return ww2;
			return ""; 	
		} catch(e1) {
			return "";			
		}
} // end of extract

//----------------------------------------------------
// set previous run voice
 
function loadPrevLang(prevLanguage) {
	
	// language=5,de,de-DE,Microsoft Stefan - German (Germany)   ( it's in the first line of dictionary) 
	//          0,1 ,2    ,3  
	var pvLang = prevLanguage.split(",") 
		
		
	prev_voice_ix        =  pvLang[0] ; 	
	prev_voiceLang2      =  pvLang[1] ; 		
	prev_voiceLangRegion =  pvLang[2] ;  
	prev_voiceName       =  pvLang[3] ; 
	 
	//console.log(" wordByFrequence.js loadPrevLang()" + " prev_voiceName = "+  prev_voiceName  );  
	
	fcommon_load_all_voices(); // at end calls tts_1_toBeRunAfterGotVoices()
	
	// WARNING: the above function contains asynchronous code.  
	// 			Any statement after this line is executed immediately without waiting its end

} // 
//---------------------------------
function js_go_setError(msg) {
	document.getElementById("id_startwait").innerHTML = msg; 	
}
//------------------------------
var numP=0
function whereIs(here) {
	//numP++
	
	//if (numP < 10) { console.log( "\whereIs(" + here) }
	
	if (here.substring(0,2) != "::") {	
		//ele_where.innerHTML = here ;
		return ;
	}
	
	var cols = here.split("::")
	//ele_where.innerHTML = cols[2];
	var perc=0;
	try{
		perc = parseInt( cols[1] )
		ele_bar.style.width = perc + "%";  
		ele_bar2.style.width = (100-perc) + "%";  
		ele_bar3.innerHTML = perc + "%";  
	} catch(e1) {
	}	
 	
	
} 
//-------------------------
function showProgress(perc0) {
	var perc=0;
	try{
		perc = parseInt( perc0 )
		ele_bar.style.width = perc + "%";  
		ele_bar2.style.width = (100-perc) + "%";  
		ele_bar3.innerHTML = perc + "%";  
	} catch(e1) {
	}		
}
//----------------------------------

function onclick_tts_seeWordsGO1(numTr, ixRow) {
	// numTr è il numero progressivo della riga nella pagina, ixRow è il numero della riga nella lista di tutte le righe del testo (in GO) 

	
	word_to_underline_list = []
	
	var id_analWords = "idw_" + numTr;
	var ele_wordset = document.getElementById(id_analWords);   
	if (ele_wordset == false) {
			console.log("onclick_tts_seeWordsGO1(this1," , numId00, ") ==> ", eleTR.outerHTML, "\n\tid_analWords=" + id_analWords , " ERROR ele_wordset == false" )
		return
	}
	
	ele_wordset.innerHTML = ""; 
		
	go_passToJs_rowWordList(""+numTr,""+ixRow, "js_go_rowWordList"); // ask 'go' to give wordlist by js_... function  
		
} // end of onclick_tts_seeWordsGO1



//------------------------------------
function js_go_rowWordList(wordListStr) {
	
	//console.log("js_go_rowWordList(wordListStr=", wordListStr)
	
    // triggered by go func (  go _ passToJs_wordList )
    if (wordListStr == undefined) {
        console.log("js_showWordList: parameter is undefined");
		onclick_jumpFromToPage( myPage02,myPage03,  myPage01);  //   
        return;
    }
    if (wordListStr == "") {
        console.log("js_showWordList: parameter is empty");
		onclick_jumpFromToPage( myPage02,myPage03,  myPage01); 
        return;
    }	

    var rowWordList = wordListStr.split(endOfLine);	
	
	//console.log("rowWordList=", rowWordList)
	
	var numId_0 = rowWordList[0].split(",");
	var numId = (""+numId_0[0]).trim();  	
	

	var anal_txt = ""; var anal_tts_txt=""; 
	var idc1 = "idc_"  + numId;
	var idtts= "idtts" + numId; 

	var anal_ele_idc   = document.getElementById(idc1 );		
	var anal_ele_idtts = document.getElementById(idtts);	
		
	if (anal_ele_idc) {
		anal_txt = anal_ele_idc.innerHTML;
		anal_tts_txt= anal_ele_idtts.innerHTML; 
	} else {
		console.log("js_showWordList: idc1=", idc1 , "(anal_ele_idc==false)" )  
		return;
	}
	
	var eleTR = anal_ele_idtts.parentElement.parentElement.parentElement; 
	var trHeight = eleTR.offsetHeight; 	
	
	var id_analWords = "idw_" + numId;
	var ele_wordset = document.getElementById(id_analWords);   
	
	
	if (last_ele_analWords_id != "") {	
		// remove the previous 
		var last_ele_analWords = document.getElementById(last_ele_analWords_id);
		if (last_ele_analWords)  {
			
					
			last_ele_analWords.style.height = null;   				
			last_ele_analWords.innerHTML = "";  
			last_ele_analWords_tr.style.height = last_ele_analWords_height;	
			
				
			last_ele_analWords = null; 					
		}	
		if (id_analWords == last_ele_analWords_id) {  // if it's the same as the previous then  remove it  		
			last_ele_analWords = null; 	
			last_ele_analWords_id = "";
			return; 
		}	
		last_ele_analWords_id = "";
	}	
	
	var ixLastEle, table_txt
	[ ixLastEle, table_txt] = tts_3_spezzaRiga3( rowWordList.slice(1) )	 ;   
 	
	var divWord = "";
	
	divWord += table_txt;
	//if (table_txt.indexOf("creazione")>= 0 ) { console.log("table_txt=" , table_txt); }
	ele_wordset.innerHTML = divWord; 
	
	//console.log("\nXXXXXXXXXXXXX\nXXXXXXXXXXXX\n ele_wordset.innerHTML = " + ele_wordset.innerHTML ) 
	
	if (ixLastEle > 0) {
		var eleF = document.getElementById("wb1_0");
		var eleT = document.getElementById("wb2_" + ixLastEle );
		onclick_tts_word_arrowFromIx(eleF, 0,         true, false)
		onclick_tts_word_arrowToIx(  eleT, ixLastEle, true, false)
	}	
	
	var prevTR  = document.getElementById( "idtr_" + (numId-4) ); 
	if (prevTR) tts_5_fun_scroll_tr_toTop( prevTR ); 	// scroll 
		
	last_ele_analWords_id = id_analWords;
	last_ele_analWords_tr = eleTR; 
	last_ele_analWords_height = trHeight;  
	
} // end of js_go_rowWordList

//------------------------------------
// ===================================================================
function tts_3_spezzaRiga3( rowWordList ) {
	
    var endix2 = -1;

    var listaParole = [];
    var listaParo_tts = [];	
    var listaParo_lemma = [];
    var listaParo_tran  = [];
	var listaParo_nFrasi = [];  

    //console.log( list1.length + " " + list2.length); 


   // if (list2.length != list1.length) list2 = orig_riga2.split("§");
	var row1, k1,k2, col1; var part1;
	var Xword1;
	var listLemS, listTranS; 
	var listLem, listTran	
	var xWord_numFrasi;
	
	//console.log("tts_3_spezzaRiga3( rowWordList = " , rowWordList , "<=="  ) 	
	
    for (var k = 0; k < rowWordList.length; k++) {
		row1 = (rowWordList[k]+";").replaceAll(";[",";").replaceAll("];",";")
		if (row1.trim()=="") continue; 
		
		//console.log("\tspezzaRiga3 k=" , k, " row=>" , row1); // spezzaRiga3 k= 6  row=> einem,14,2;[ein einem einer];[a uno uno]
		
		part1 = row1.split(";") 
		if (part1.length < 3) continue
		col1 = part1[0].split(",")
		Xword1 = col1[0];
		xWord_numFrasi = col1[2]; 
		listLemS = part1[1]; 
		listTranS = part1[2]
		listLem = listLemS.split(   wSep ) ;   //  cambia separatore spazio con virgola 
		listTran = listTranS.split( wSep ) ;   //  
		
		var ixLemma;  var strTran=""; var xtt = "" ;
		for( ixLemma=0; ixLemma < listTran.length; ixLemma++) {
			xtt = " " + listTran[ixLemma] + " "
			if (strTran.indexOf(xtt) < 0){ strTran += xtt; }		
		}	
		var strLem=""; var xll = "";  
		for( ixLemma=0; ixLemma < listLem.length; ixLemma++) {
			if (Xword1 == listLem[ixLemma] ) { continue;}
			xll = " " + listLem[ixLemma] + " "
			if (strLem.indexOf(xll) < 0){ strLem += xll; }		
		}	
		strLem =  strLem.trim(); 
		if (strLem != "") { strLem = "(" + strLem + ")"; }
		listaParole.push(   Xword1 ) ;
		listaParo_nFrasi.push( xWord_numFrasi );   
        listaParo_tts.push( Xword1 ) ;
		listaParo_lemma.push( strLem);	
		listaParo_tran.push(  strTran.trim()); 
    }
    var parola1, paro_tts, paro_lemma, paro_tran, paro_nFrasi  ;

    var frase_showTxt = prototype_word_table_header ; //  '<div><table>
	
	//console.log("\t----------------- leng listaParole=", listaParole.length, " ==>", listaParole); // spezza
	
	var maxNumRow5 = 100; 
    for (let z3 = 0; z3 < listaParole.length; z3++) {
        parola1    = listaParole[z3].trim();
		paro_nFrasi= listaParo_nFrasi[z3];
        paro_tts   = listaParo_tts[z3];
		paro_lemma = listaParo_lemma[z3];		
		paro_tran  = listaParo_tran[z3];	
		
		frase_showTxt  += getWord_tr( z3, parola1, paro_tts, paro_lemma, paro_tran, paro_nFrasi,maxNumRow5) + "\n";
      
    } // end of for z3
	var last1 = listaParole.length - 1 
	frase_showTxt +=  prototype_word_table_end ; // </table></div>
	
    return [ last1, frase_showTxt ];

} //  end of  spezzaRiga3()

//====================


function onclick_FR_tts_getInput(elepage) {

  if (fr_tts_1_join_orig_trad() < 0) {
	  return;
  }
  if (elepage == myPage04) {
	  myPage05.style.display = "block";
  } else {
	  myPage04.style.display = "block";
  }
  elepage.style.display = "none";
}
//------------------------------------  

function fr_tts_1_join_orig_trad() {
	document.getElementById("id_msg16").innerHTML = "";

	sw_translation_not_wanted = false;

	var msgerr0 = "";
	var msgerr1 = fr_tts_1_get_orig_subtitle2(); // get orig. text /srt
	var msgerr2 = fr_tts_1_get_tran_subtitle2(msgerr1); // get tran. text/srt	

	msgerr0 += msgerr1 + msgerr2;
	var msgerr3 = "";

	msgerr0 += msgerr3;

	if (msgerr0 != "") {
		tts_1_putMsgerr(msgerr0);
		document.getElementById("id_msg16").style.color = "red";
		return -1;
	}
	document.getElementById("id_msg16").style.color = null;

	tts_9_toBeRunAfterGotVoicesPLAYER(); 

	return 0;

} // end of tts_1_join_orig_trad()

//----------------------

//--------------------------------------------------
 function fr_tts_1_get_orig_subtitle2() {
     
      var msgerr1 = "";
    
      //builder_orig_subtitles_string = document.getElementById("txt _ pagOrig").value.trim();
	  
	  builder_orig_subtitles_string = ele_toTranslate_textarea.value.trim();
      
	  if (builder_orig_subtitles_string != "") {
          sw_inp_sub_orig_builder = true;
      } else {
          sw_inp_sub_orig_builder = false;
          msgerr1 += "<br>" + tts_1_getMsgId("m132"); //  ma22 the source language subtitle file  has not been read or is empty" ;         
      }
	  inp_row_orig = builder_orig_subtitles_string.split("\n") ;
	  inp_row_orig.push("");    
	  numOrig = inp_row_orig.length;
	
      return msgerr1;

  } // end of get_orig_subtitle2()
  //--------------------------------------------------

  function fr_tts_1_get_tran_subtitle2(msgerrOrig) {
     
      var msgerr1 = "";

      //builder_tran_subtitles_string = document.getElementById("txt _ pagTrad").value.trim();
	  
	  builder_tran_subtitles_string = ele_translated_textarea.value.trim();
	 
      sw_inp_sub_tran_builder = false;
      if (builder_tran_subtitles_string != "") {
          sw_inp_sub_tran_builder = true;
      } else {
          sw_inp_sub_tran_builder = false;
          if (sw_translation_not_wanted == false) {
              msgerr1 += "<br>" + tts_1_getMsgId("m133"); //    translated subfile missing     					
              if (msgerrOrig == "") { // only if original srt is Ok  
                  msgerr1 += "<br>" + tts_1_getMsgId("m134").replace("§...§", TRANSLATION_NOT_WANTED); // if there is no translation
			  }
          }
      }
	  inp_row_tran = builder_tran_subtitles_string.split("\n");
	  inp_row_tran.push(""); 
	  numTran = inp_row_tran.length;  
	  
      return msgerr1;

  } // end of tts_1_get_tran_subtitle2()

//--------------------------
function replaceSepar( wT1 ) {	
	/**
	1;;1374;;sie zum zweiten Male jaeh zu verlassen gezwungen war, so hatte er sie
	1, 1374, e fu costretto a lasciarla per la seconda volta, l'ebbe
	**/
	var k1,k2,k3,kx1,kx2, kx10,kx20, num1,num2,new_wT;
	
	k1 = wT1.indexOf(";;"); if (k1 < 0) k1=99999
	k2 = wT1.indexOf(";" ); if (k2 < 0) k2=99999 
	k3 = wT1.indexOf("," ); if (k3 < 0) k3=99999
	kx1 = Math.min(k1,k2,k3)
	if (kx1 > 14) return wT1
	num1 = wT1.substring(0,kx1).trim()
	kx10 = kx1+1
	if (wT1.substr(k1,2) == ";;") kx10 = kx1+2  
		
	k1 = wT1.indexOf(";;", kx10); if (k1 < 0) k1=99999
	k2 = wT1.indexOf(";",  kx10); if (k2 < 0) k2=99999 
	k3 = wT1.indexOf(",",  kx10); if (k3 < 0) k3=99999
	kx2 = Math.min(k1,k2,k3)	
	if (kx1 > 8) return wT1;
	
	num2 = wT1.substring(kx10, kx2).trim() 
	kx20 = kx2+1
	if (wT1.substr(k2,2) == ";;") kx20 = kx2+2  
	
	
	var rest = wT1.substring(kx20).trim(); 
	new_wT = num1 + ";;" + num2 + ";;" + rest;
	try {
		var num1Nu = parseInt(num1);
		var num2Nu = parseInt(num2);
	} catch(e) {
		return wT1; 		
	} 
	
	return new_wT; 
	
} 
//--------------------------------------

function onclick_showRowsButton(type) {
	//    chiamata dopo che è stata copiata la traduzione delle righe 
	/***	
	objOrig.value= 
	;;0;2;; Erstes Kapitel
	;;0;3;; 
	;;0;4;; Gustav Aschenbach oder von Aschenbach, wie seit seinem fuenfzigsten
	;;0;5;; Geburtstag amtlich sein Name lautete, hatte an einem

	objTran.value=
	;;0;2;; Primo capitolo
	;;0;3;;
	;;0;4;; Gustav Aschenbach o von Aschenbach, come ha fatto fin dai cinquant'anni
	;;0;5;; Il suo compleanno ufficiale cadeva l'una	
	***/
	//-----------------------------
	// if all words are translated continue 
	// is some translations are missing, ask if they must be ignored, otherwise loop till no translations are missing     
	if (type == 1) {
		sw_ignore_missTranRow = true;  // Ignore missing translations and continue
	} else {
		sw_ignore_missTranRow = false; // Translations added, check again</button	
	}  	
	
	//console.log("onclick_showRowsButton(type=", type, ")  sw_ignore_missTranRow=", sw_ignore_missTranRow); 
	
	//---------------------------------------------------
		
	//	console.log("onclick_showRowsButton ele_translated_textarea = ", ele_translated_textarea.value); 	 
		
	var rowTranList = ele_translated_textarea.value.trim().split( "\n" );  
	
	var lenTran = rowTranList.length;	
	var rowTran
	var ixRowS, ixRow, ouIxRowS, ouIxRow, rowNewTran; 
	//------------------------------------------

	var nfileW, idRowW, ixRowW, rowW, tranW;
	//-----------
	var wList;
	var sw_someTranMissingR = false;  // sometimes the automatic translator does some mistakes 
	var wT;
	var numNoTranRow = 0;
	var numNoTranRow2 = 0;
	var numeroTS_Row=0, numeroTS_OkTran=0, numeroTS_NoTran=0;
	var newAddedTran = 0; 
	//---------------------------------
	
	var lenW    = rowToStudy_list.length;
	
	
	//----------------------------
	//  scandisce le righe di traduzione copiate dal traduttore google (potrebbero essere meno di rowToStudy_list ) 
	
	for(var z=0; z < lenTran; z++) {
	
		wT = rowTranList[z].trim() + ";;;;;;;;;;;"	; 
		
		//console.log("translation added  rowTranList[", z, "] = " ,  rowTranList[z]  ); 
		
		/**
		1;;2;;Erstes Kapitel
		3;;4;;Gustav Aschenbach oder von Aschenbach, wie seit seinem fuenfzigsten
		4;;5;;Geburtstag amtlich sein Name lautete, hatte an einem
		**/
		if (wT == "") {continue;}
			
		wList = wT.split(";;") ;
		/***
		if (wList.length != 3) { 
			wT = replaceSepar(wT);
			wList = wT.split(";;") ;
		}
		if (wList.length != 3) { 
			continue;
		}
		***/
		[ouIxRowS, ixRowS, rowNewTran] = wList.slice(0,3) ;   			
		
		rowNewTran = rowNewTran.trim();
		
		//console.log("     1  ouIxRowS=", ouIxRowS, " ixRows=", ixRowS, "rowNewTran=", rowNewTran)
		
		if (rowNewTran == "") { continue; }
		
		try {
			ouIxRow = parseInt(ouIxRowS)    // row index  in the row html page   
			ixRow   = parseInt(ixRowS)      // row index  in the row DB    ( that is:  inputTextRowSlice[ixRow] in GO )
		} catch(e) { 
			continue; 
		}	
		if (ouIxRow >= lenW) {
			//console.log("showWordsButton() error2  entry z=" + z + " = " + wt  + " ixRow=" + ixRow , " >= lenW=" , lenW ); 
			continue;
		} 
		
		//console.log("     2  ouIxRow=", ouIxRow, " ixRow=", ixRow)
		
		
		var oldRow = rowToStudy_list[ouIxRow];
		//console.log("     rowToStudy_list[ouIxRow] = "  +  oldRow )
		// accoppia la riga di traduzione con quella originale  
		try {
			[nfileW,idRowW, ixRowW, rowW, tranW] = oldRow.split("|");
		} catch(e1) {
			/**
			console.log(e1); 
			console.log("wT=" , wT)
			console.log("wList=", wList)
			console.log("ouIxRow=" + ouIxRow )  
			**/
			continue;
		}
		
		if ( parseInt(ixRowW) != ixRow) {
			//if (sw_ignore_missTranRow == false) {
			//	sw_someTranMissingR = true;	
			//console.log("showRowsButton() error4 entry z=" + z + " = " + wT  + "\n\t rowToStudy_list[ouIxRow="+ ouIxRow +"]=" +  rowToStudy_list[ouIxRow] +  
			//		"\n\tixRowW=" + ixRowW + " ixRow=" + ixRow);	
			//}						
			continue; 
		}  
		if (tranW == rowNewTran) { continue; }
		
		newRowTran[ouIxRow] = 1; 
		
		rowToStudy_list[ouIxRow] = nfileW + "|" + idRowW + "|" + ixRowW + "|" + rowW + "|" + rowNewTran ; // update with translation 	
		
		console.log("  xxx  NUOVA rowToStudy_list[ouIxRow=",ouIxRow , "] = " + rowToStudy_list[ouIxRow]	)
		
		newAddedTran++; 
					
	
	}  // end of for(var z ...
	//--------------------------	
	
	[numeroTS_Row, numeroTS_OkTran, numeroTS_NoTran ] = build_Page1_rowsToTranslate("2onclick_showRowsButton(type=", type);
	
	document.getElementById("id_notTranNumRow").innerHTML = numeroTS_NoTran; 
	
	write_row_dictionary(); 
	
	if (numeroTS_NoTran > 0) {
		if (sw_ignore_missTranRow == false) {
			return;  
		}
	} 
	
	showRowsAndTranButton("3");
	
} // end of onclick_showRowsButton()
 

//----------------------


//---------------------------

function showRowsAndTranButton(wh) {	
	
	var showList = ''    ;
    var word1, ix1, nrow, wLemma1;
	var riga;
	
	var wordOrig2, wordTran2;
	
	
	var wordTran = ""
	var row0; 
	var nfile, idRow, ixRow, origRow, origTran; 
	
	//string_tr_xx = "\n" + prototype_tr_m2_tts + "\n" + prototype_tr_m1_tts + "\n" + prototype_tr_tts; 
	string_tr_xx = "\n" + prototype_tr_tts; 
	
	//word_tr_allclip =  "\n" + prototype_word_tr_m2_tts + "\n" + prototype_word_tr_m1_tts + "\n" + prototype_word_tr_tts; 
		
	//--------------	
	var txt1p, text_tts, tranRow; 
	var nFileR, nfile_zero;
	var first= -1, last=-1;
	var visib;
	var ixRow2StudyLs;
	var inpBegRow = getInt( document.getElementById("id_fromIx_row").innerHTML );
	var numRows   = getInt( document.getElementById("id_inpNumRow" ).innerHTML );  
	var inpEndRow = inpBegRow+numRows-1;
	//------------------------------------
	var newRowList1 =[];
	var newRowList2 =[];
	//--- 
	for (var i1 = 0; i1 < rowToStudy_list.length; i1++) {	
		
		row0 = rowToStudy_list[i1].trim();	
		
		//console.log("showRowsAndTranButton () XXX i1=", i1, " rowToStudy_list[i1]=", rowToStudy_list[i1] );  	
		
		if (row0=="") {continue;}	
		if (row0 == "\n") {continue}
		
		row0 = i1 + "|" + row0 + "|||||"; 
			
		var cols = row0.split("|")
				
 		try{ 
			[ixRow2StudyLs, nfile, idRow, ixRow, origRow, tranRow] = cols.slice(0,6); 
		} catch(e1) {
			console.log("showRowsAndTranButton () o=", i, " rowToStudy_list[i]=", rowToStudy_list[i], " row0=", row0, " cols=", cols, "\n\t XXX  error ", e1) 			
		}
		
		if ((ixRow >= inpBegRow) && (ixRow <= inpEndRow)) {		
			newRowList1.push(row0)	
		} else {
			newRowList2.push(row0)
		}	
    }  // end for i1
	//---------------------------
	//var endLine1 = "_endLine1_" 
	//newRowList1.push(endLine1);
	var newRowList = newRowList1.concat(newRowList2);
	//---------------------------
	var iNumTr =0;
	var idRow1, idRow2; 
	for (var i = 0; i < newRowList.length; i++) {
		iNumTr = i+1;
		row0 = newRowList[i].trim() + "|||||"; 		
		var cols = row0.split("|")
 				
		try{ 
			[ixRow2StudyLs, nfile, idRow, ixRow, origRow, tranRow] = cols.slice(0,6); 
		} catch(e1) {
			console.log("showRowsAndTranButton () o=", i, " rowToStudy_list[i]=", rowToStudy_list[i], " row0=", row0, " cols=", cols, "\n\t XXX  error ", e1) 			
		}
		//if (i < 5) {  console.log("showRowsAndTranButton () i=", i, " rowToStudy_list[i]=", rowToStudy_list[i]  , " idRow="+ idRow) }
		
		/***
		if (nfile == endLine1) {
			riga = string_tr_xx.replaceAll("§1§", i).
				replaceAll("§4txt§"  , "").
				replaceAll("§5txt§"  , "").
				replaceAll("§ttstxt§", "").
				replaceAll("§6id§"   , "").
				replaceAll("§6ix§"   , "").
				replaceAll("§nfile§" , ""). 
				replaceAll("§visib§" , "");
			
			showList    += riga + "\n";	
			continue; 
		}  
		***/
		
		if (origRow == "") { continue; } 
		if (origRow == undefined) { continue; } 	
		
		origRow = origRow.trim(); 
		txt1p = origRow;
			
		for(var hx=0; hx < word_to_underline_list.length; hx++) {  
			txt1p = evidenzia(word_to_underline_list[hx], txt1p); 
		}
		
		txt1p = txt1p.replaceAll("§§", "");  // have beewn addded in function evidenzia
		
		text_tts = "";
		visib=""; 
		if (origRow == "") {
			visib = "visibility: hidden;"; 			
		} 	
		
		if ((ixRow >= inpBegRow) && (ixRow <= inpEndRow)) {
			nfile = 1
		} else {
			nfile = 2	
		} 
		var idro1 =idRow.split("(");
		if (idro1.length<2) {
			idRow1 = idRow; idRow2 = "";
		} else {
			idRow1 = idro1[0];  idRow2 = idro1[1]; 
		}
		//if (i < 5) { console.log("\t idrow=", idRow, " idro1=", idro1.join(",") ,"   idRow1=", idRow1, " idRow2=", idRow2) }
		
		
		let txt1p_n   =   txt1p.replaceAll("%20"," ").replaceAll("/","/ ").replaceAll("</ ","</");
		let tranRow_n = tranRow.replaceAll("%20"," ").replaceAll("/","/ ").replaceAll("</ ","</");	
        riga = string_tr_xx.replaceAll("§1§", iNumTr).
			replaceAll("§ixRow2StudyLs§"  , ""+ixRow2StudyLs).
			replaceAll("§4txt§"  , txt1p_n).
			replaceAll("§5txt§"  , tranRow_n).
			replaceAll("§ttstxt§", text_tts).
			replaceAll("§6id§"   , idRow1.replace(" "," - ") ).
			replaceAll("§6id2§"   ,idRow2).
			replaceAll("§6ix§"   , ixRow).
			replaceAll("§nfile§" , nfile). 
			replaceAll("§visib§" , visib);
			
		//if (i < 5) { console.log("\t riga=", riga) }	
			
		/**
		if (first < 0) first = i;
		last = i; 	
		**/
		if (first < 0) first = iNumTr;
		last = iNumTr;
		
		showList    += riga + "\n";	
				
		
   } // end for i
	//---------------------------------------------------
	eleTabSub_tbody.innerHTML = showList;
	
	if ( (last - first) > 0) {
		let eleF = document.getElementById("b1_" + first);
		let eleT = document.getElementById("b2_" + last);
		onclick_tts_arrowFromIx(eleF, first, 5);
		onclick_tts_arrowToIx(  eleT, last , "3showRowsAndTranButton" );		
	}
		
	onclick_jumpFromToPage( myPage04,0, myPage05);  
	

} // end of showRowsAndTranButton

//-------------------------------------------------

function write_row_dictionary() {
	var nfileW,ixRowW, rowW, tranW; 
	var word1, ix1, nrow, wLemma1, wordTran, col1;
	
	var newTranRow=0;
	var listNewTranRows = "";
	
	//console.log("\n----------------------\nwrite_row_dictionary() rowToStudy_list=" ,  rowToStudy_list ,"\n----------")
	var idRow1, ixRow1; 
	var id_ix_Row;
	
	for (var i = 0; i < rowToStudy_list.length; i++) {
		if ( rowToStudy_list[i] == "") continue; 
		//console.log("write " ,  rowToStudy_list[i] )
		
		col1 = (rowToStudy_list[i]+ "||||||").split("|");		
		nfileW = col1[0];  
		idRow1 = col1[1]; 
		ixRow1 = col1[2];
		rowW   = col1[3]; 
		tranW  = col1[4].replaceAll("|"," "); 		
		
		//console.log( "\t col1 = ", col1, "    tranW=", tranW)

		
		if (tranW == "") { continue; }
		
		
		//console.log("\t" , [nfileW,ixRowW, tranW] )
		if (newRowTran[i]==1) {                                                                 
			newTranRow++;                                     
			listNewTranRows  += "\n" + idRow1+"|" + ixRow1+ "|" + tranW ;   // new line for dictionary 
			//console.log("write row direct ", idRow1+"|" + ixRow1+ "|" + tranW) ;  
		}
	}
	//------------
		
	if (newTranRow < 1) {
		//console.log("write_row_dictionary",  " nessuna nuova traduzione"); 	
		return;
	}
	console.log("write_row_dictionary ",  newTranRow, " righe tradotte"); 	
	
	go_write_row_dictionary(  listNewTranRows );  		
	
} // end of write_row_dictionary

//------------------------------------------

function show_altreRighe(this1,numRows) {
	
	var eleTD = this1.parentElement
	
	var eleOnOff = eleTD.children[1]		
	
	//console.log("================\nantonio  show_altreRighe() 1 eleOnOff=", eleOnOff.innerHTML  , " display=", eleOnOff.style.display, " eleTD outer=", eleTD.outerHTML) 
	
	var showSPAN, showTR;
	if (eleOnOff.innerHTML == "block") {
		showSPAN = "none"; 
		showTR   = "none"
		eleOnOff.innerHTML = showSPAN; 
	}	else{
		showSPAN = "block"; 
		showTR   = "table-row"
		eleOnOff.innerHTML = showSPAN;
	}
	
	//eleTD.style.display = showSPAN;  
	
	//console.log("antonio  show_altreRighe() 2 eleOnOff=", eleOnOff.innerHTML , " display=", eleOnOff.style.display)
	
	var eleTR = eleTD.parentElement	
	

	var nextTR = eleTR;
	var nextTR2 ;
	
	for (var rr=0; rr < numRows; rr++ ) {			
		nextTR2 = nextTR.nextElementSibling; 
		if (nextTR2 == undefined) { break } 
		nextTR = nextTR2; 
		nextTR.style.display = showTR; 
	} 
	
	//console.log("antonio  show_altreRighe() 9 eleOnOff=", eleOnOff.innerHTML  , " display=", eleOnOff.style.display, " eleTD outer=", eleTD.outerHTML) 
		
} // end of show_altreRighe
//---------------------------------	

function mouseOverWord(this1) {
	if (this1.parentElement.children.length > 1) { 
		this1.parentElement.children[1].style.display = "inline-block";
	}
}
//---------------------
function mouseOutWord(this1) {
	if (this1.parentElement.children.length < 1) { 
		return;
	}
	var ele_tran = this1.parentElement.children[1];	
	//if (ele_tran.contentEditable == true) { return; }		
	
	ele_tran.style.display = "none";
		
} // end of mouseOutWord
//-------------------------------

function mouseOverWord2(this1) {
	if (this1.children.length > 1) { 
		this1.children[2].style.display = "inline-block";
	}
}
//---------------------
function mouseOutWord2(this1) {
	if (this1.children.length > 1) { 
		this1.children[2].style.display = "none";
	}
}
//-------------------------------


//-----------------------------------------
function onclick_word_known(sIxWord, yesNo, this1) {
	var ixWord = 0;
    try {
        ixWord = parseInt(sIxWord);
    } catch (err) {}	
	
	var num1 = 1*this1.innerHTML; 
	var ele_td= this1.parentElement;
	if (yesNo==0) {		
		var ele_nextTd = ele_td.nextElementSibling; 
		
		var next_this1 = ele_nextTd.children[0]
		next_this1.innerHTML = 0;
		next_this1.style.border = null; 
		go_passToJs_word_known(""+ixWord, "1", "0", "js_go_word_known"); // ask 'go' to update yes/no word known ctr  
		return 
	} 
	num1++;
	this1.innerHTML = num1;
	if (num1 > 0) {
		this1.style.border = "5px solid black"; 
	} else {
		this1.style.border = null; 
	}
	
    go_passToJs_word_known(""+ixWord, ""+yesNo, ""+num1, "js_go_word_known"); // ask 'go' to update yes/no word known ctr  
	
} // end of onclick_word_know_yes		

//-----------------------------------------
function js_go_word_known(str1) {
	//console.log("js_go_word=", str1);   
}// end of js_go_word_known 	
	
//-------------------------------------------------	
		
function onclick_sortWordBy_ixField(nField1,nChild1,isNumber1,ascending1, 
									nField2,nChild2,isNumber2,ascending2) {
	
	if (arguments.length != 8) {
		console.log("error: wrong number of arguments in onclick_sortWordBy_ixField( ", nField1,nChild1,isNumber1,ascending1, 
									nField2,nChild2,isNumber2,ascending2, 
					"\n check  'prototype_tableWordList_Header'  in 'wordsByFrequency_prototype_script.html_js' file " ); 		
	}
	
	document.getElementById("id_wordList1").scrollTop = 0;
	
	
	var ele_tbody = document.getElementById("idTableWordList_tbody"); 
	var num_tr = ele_tbody.children.length; 
	var ele_tr, ele_td, ele_butt ; 
	var num_td =0; 
	
	var colgr = document.getElementById("idwcol2"); 
	var colgrX
	for(var f=0; f < colgr.children.length; f++) {
		colgrX = colgr.children[f]
		if (colgrX) {
			if (f == nField1) {
				colgrX.style.borderLeft  = "3px solid green"; 
				colgrX.style.borderRight = "3px solid green"; 
			} else {
				colgrX.style.borderLeft  = null;
				colgrX.style.borderRight = null;
			}
		}
	}  
	
		
	//document.getElementById("idwcol2_" + nField).style.backgroundColor =  "gray"; // "#c4e8db";  //"#b2bbcb";  // "#33ffcc" ; //"#33ff99" ; //"#ccfff2";   
	
	//console.log("onclick_sortWordBy_ixField()1 nField=", nField1,  " isNumber=", isNumber1, "ascending1=", ascending1,
	//			" nField2=", nField2,   " isNumber2=", isNumber2, "ascending2=", ascending2);
	
	
	//console.log("onclick_sortWordBy_ixField() => ", document.getElementById("idTableWordList").innerHTML.substr(0,500));  

	
	const EMPTY  = "_none_"; 
	
	var listKey=[ EMPTY ]; // lascio l'entrata 0 occupata 
	var key1 , key2; 
	var ke2, ix1, ix2; 
	var MAXKEY = 1000000;  
	
	for(var g=0; g < num_tr; g++) {
		//if (g > 20) { break; }
		
		ele_tr = ele_tbody.children[g]; 
		
		//if (g==2) {console.log("onclick_sortWordBy_ixField()1.1 ", " num_tr=", g, " ==> TR=", ele_tr.innerHTML);}
		
		num_td = ele_tr.children.length; 
		//--
		ele_td = ele_tr.children[nField1]; 	
		if (nChild1==0) {
			key1 = setKey0(isNumber1, ele_td.innerHTML, ascending1); 
		} else {
			ele_butt = ele_td.children[0]; 	// might be button or  span						
			key1 = setKey0(isNumber1, ele_butt.innerHTML, ascending1);
		} 		
		
		//--
		ele_td = ele_tr.children[nField2];	
		if (nChild2==0) {
				key2 = setKey0(isNumber2, ele_td.innerHTML, ascending2); 
		} else {	
			ele_butt = ele_td.children[0]; 	// might be button or  span							
			key2 = setKey0(isNumber2, ele_butt.innerHTML, ascending2);
		}		
		
		listKey.push( key1 + "::" + key2 + "::" + (MAXKEY + g)  ); 
		
		//console.log("g=", g, " 2 listKey[", listKey.length-1, "] = ", listKey[listKey.length-1] ); 
	} 
	//----------------------------------------
	
	
	listKey.sort(); 
	
	//console.log("SORT listKey ==> ", listKey.length)    
	
	var newBodyInner = "";
	var newTd;
	//--------------------
	var gg=0
	for(var g=0; g < listKey.length; g++) {
		if (listKey[g] == EMPTY) { 
			continue; 
		}
		ke2 = listKey[g].split("::"); 
		try {
			ix2 = parseInt( ke2[2] ) ;
			ix1 = ix2 - MAXKEY; 
		} catch (err) {
			ix2=0;
			ix1 = 0;
		}
		ele_tr = ele_tbody.children[ix1]; 		
		if (ele_tr) {
			gg++
			newTd = '<td style="text-align:center;">' + gg + '</td>' ; 
			ele_tr.children[0].outerHTML = newTd; 			
			
			newBodyInner += ele_tr.outerHTML + "\n"; 
			
		} 
		
	} 
	//-----------------------
	
	ele_tbody.innerHTML = newBodyInner; 	
	
	//-------------------------------

	function setKey0(isNumber, sNum, ascending) {
		if (!isNumber) {			
			return keyAlphaCod(sNum);		
		}
		var newKey="";
		var num1 =0, num2;	
		try {
			num1 = parseInt(""+sNum) ;
		} catch (err) {
			num1 = 0;
		}
		if (ascending == "a") {
			return MAXKEY + num1;
		} else { 
			return 2*MAXKEY - num1;
		}
    } // end of setKey0

	//---------------

} // end of onclick_sortWordBy_ixField



//-----------------------------------------------------------------
function onclick_write_words_to_learn() {
	
	go_passToJs_write_WordsToLearn("js_go_file_words_to_learn_written"); 
	
}
//-----------------------------------------
function js_go_file_words_to_learn_written( str1 ) {
	document.getElementById("id_w_to_learn_written").innerHTML = str1 ;	
}
//-----------------------------------------------------------------

//-----------------------------------------------------------------
/**
function onclick_read_words_to_learn() {
	
	go_passToJs_read_WordsToLearn("js_go_file_words_to_learn_read"); 
	
}
***/
//-----------------------------------------
function js_go_file_words_to_learn_read( str1 ) {
	//console.log( str1 );	
}
//-----------------------------------------------------------------
function onclick_remove_words_already_known() {	
	
	var ele_tbody = document.getElementById("idTableWordList_tbody"); 
	var num_tr = ele_tbody.children.length; 
	var ele_tr, ele_td, ele_butt ; 
	var num_td =0; 
	
	var swWrite=false; 
	var newBodyInner = "";
	var numYesCtr=0; 
	
	const YESNO_NO_field = 4; 
	
	
	for(var g=0; g < num_tr; g++) {
		
		ele_tr = ele_tbody.children[g]; 
			
		num_td = ele_tr.children.length; 
			
		//if (num_td < 4) { continue } 
		
		ele_td = ele_tr.children[YESNO_NO_field];  // knowYesCtr 
		
		//console.log("remove words ",  ele_tr.children[5].innerHTML, " yesNo=", ele_tr.children[4].innerHTML ); 

		swWrite=false; 
		numYesCtr = 0;
		
		if (ele_td.children.length > 0) {
			ele_butt = ele_td.children[0]; 				
			try {
				numYesCtr = parseInt(ele_butt.innerHTML) ;
			} catch (err) {
				numYesCtr = 0;
			}
		}			
		if (numYesCtr > 0) {		
			swWrite=true; 
			newBodyInner += ele_tr.outerHTML + "\n"; 
		}		
	} 
	//----------------------------------------
	
	ele_tbody.innerHTML = newBodyInner; 
	
	
} // end of onclick_remove_words_already_known 
//-----------------------------------------------------------------

function keyAlphaCod(inp1) {
	// serve per mettere vicini le vocali a prescindere dall'accento ( serve x tedesco e italiano) 
	if (!inp1   ) { return ""; }
	if (inp1=="") { return ""; }
	
	var inp2 = inp1.trim().replaceAll( "ae","a" );  
	
	inp2 = inp2.replaceAll( "oe","o" );
	inp2 = inp2.replaceAll( "ue","u" );  		

	inp2 = inp2.replaceAll( "ä","a") ;  
	inp2 = inp2.replaceAll( "ö","o") ; 
	inp2 = inp2.replaceAll( "ü","u") ; 
	inp2 = inp2.replaceAll( "ß","ss");  
	
	inp2 = inp2.replaceAll( "à","a");
	inp2 = inp2.replaceAll( "é","e");  
	inp2 = inp2.replaceAll( "è","e"); 
	inp2 = inp2.replaceAll( "ì","i");  
	inp2 = inp2.replaceAll( "ò","o");  
	inp2 = inp2.replaceAll( "ù","u");  
	
	return inp2;

} // end of keyAlphaCod

//--------------------------	


function onchange_rowGroupSelectChange(sw_newGr) {
	// <select ele_gruppi ></select>	
	var ele_gruppi = document.getElementById("id_gruppi_sel"          )
	var ele_begNum = document.getElementById("id_gruppi_iBegNum"  )
	var ele_numRow = document.getElementById("id_gruppi_iNumRows" )

	if (ele_gruppi.selectedIndex < 0) { return; }
	
	console.log("onchange_rowGroupSelectChange ele_gruppi.selectedIndex=", ele_gruppi.selectedIndex) ;
	html_rowGroup_index_gr = ele_gruppi.selectedIndex ;  // indice gruppo 	
    html_rowGroup_beginNum = getInt( ele_begNum.value);  // il gruppo inizia dalla riga in id_gruppi_iBegNum 	
	html_rowGroup_numRows  = getInt( ele_numRow.value);  // numero di righe richieste in id_gruppi_iNumRows
	if (html_rowGroup_beginNum < 1) {html_rowGroup_beginNum = 1; ele_begNum.value = 1; }
	if (html_rowGroup_numRows  < 1) {html_rowGroup_numRows  = 1; ele_numRow.value = 1; }
	
	//----------
	// if the group  changes, reset beginning and number of rows 
	if (sw_newGr) {
		html_rowGroup_beginNum = 1;      // relativo all'inizio del gruppo 
		html_rowGroup_numRows  = 999999;
		ele_begNum.value = html_rowGroup_beginNum
		ele_numRow.value = html_rowGroup_numRows
	}
	if (last_html_rowGroup_index_gr == "") {last_html_rowGroup_index_gr = html_rowGroup_index_gr; } 
	if (last_html_rowGroup_beginNum == "") {last_html_rowGroup_beginNum = html_rowGroup_beginNum; } 
	if (last_html_rowGroup_numRows  == "") {last_html_rowGroup_numRows  = html_rowGroup_numRows; }  
		
	go_passToJs_getIxRowFromGroup( ""+html_rowGroup_index_gr,  ""+html_rowGroup_beginNum, ""+html_rowGroup_numRows, "js_go_gotIxRowFromGroup");
	
} // end of onchange_rowGroupSelectChange 

//----------------------------------------------
function isExtrRowChanged() {  // called by fun_require_mostFreqWordList,    sw_something_changed resetted by js_go_showWordList_lev2
	if (sw_somethingChanged) {extrRowBecause("0PrevChange");  return true; }
	if (last_html_rowGroup_index_gr != html_rowGroup_index_gr) {
		console.log("isExtrRowChanged ", "last_html_rowGroup_index_gr =" + last_html_rowGroup_index_gr + ", html_rowGroup_index_gr =" + html_rowGroup_index_gr + "<==")
		extrRowBecause("1groupIndex"); sw_somethingChanged=true; return true; 
	} 
	if (last_html_rowGroup_beginNum != html_rowGroup_beginNum) {extrRowBecause("2beginNum");   sw_somethingChanged=true; return true; } 
	if (last_html_rowGroup_numRows  != html_rowGroup_numRows ) {extrRowBecause("3numRows");    sw_somethingChanged=true; return true; }   
	
	if (last_sel_extrRow_freqWord_list  != html_sel_extrRow  ) {extrRowBecause("4selExtrRow"); sw_somethingChanged=true; return true; }    
	
	
	
	//console.log("sw_somethingChanged=",  false)  
	
	return false; 
	
	function extrRowBecause(wh) {
		console.log("sw_somethingChanged true, reason=", wh) 
		console.log("1 last_html_rowGroup_index_gr    = ", last_html_rowGroup_index_gr,    "  html_rowGroup_index_gr = ", html_rowGroup_index_gr )		
		console.log("2 last_html_beginNum             = ", last_html_rowGroup_beginNum,    "  html_rowGroup_beginNum = ", html_rowGroup_beginNum)
		console.log("3 last_html_numRows              = ", last_html_rowGroup_numRows,     "  html_rowGroup_numRows  = ", html_rowGroup_numRows )
		console.log("4 last_sel_extrRow_freqWord_list = ", last_sel_extrRow_freqWord_list, "  html_sel_extrRow       = ", html_sel_extrRow      )
	}		
} // end of isExtrRowChanged() 

//----------------------------------------------

function setLastValuesOfExtrRowChanged( agent ) {
	
	//console.log("setLastValuesOfExtrRowChanged run by ", agent); 
  	
	var ele_gruppi = document.getElementById("id_gruppi_sel"          )
	var ele_begNum = document.getElementById("id_gruppi_iBegNum"  )
	var ele_numRow = document.getElementById("id_gruppi_iNumRows" )
	//console.log("onchange_rowGroupSelectChange ele_gruppi.selectedIndex=", ele_gruppi.selectedIndex) ;
	html_rowGroup_index_gr = ele_gruppi.selectedIndex ;  // indice gruppo 	
    html_rowGroup_beginNum = getInt( ele_begNum.value);  // il gruppo inizia dalla riga in id_gruppi_iBegNum 	
	html_rowGroup_numRows  = getInt( ele_numRow.value);  // numero di righe richieste in id_gruppi_iNumRows
	
	var x2 = document.getElementById("id_sel_2_extrRow");
    var i = x2.selectedIndex;	
	html_sel_extrRow = x2.options[i].id; 
	
	last_html_rowGroup_index_gr = html_rowGroup_index_gr; 
	last_html_rowGroup_beginNum = html_rowGroup_beginNum; 
	last_html_rowGroup_numRows  = html_rowGroup_numRows ;  

	last_sel_extrRow_freqWord_list  = html_sel_extrRow  ;
	
	/**
	console.log("	1 last_html_rowGroup_index_gr    = ", last_html_rowGroup_index_gr )		
	console.log("	2 last_html_beginNum             = ", last_html_rowGroup_beginNum )
	console.log("	3 last_html_numRows              = ", last_html_rowGroup_numRows  )
	console.log("	4 last_sel_extrRow_freqWord_list = ", last_sel_extrRow_freqWord_list )
	**/
	return false; 
} // end of setExtrRowChanged 


//-------------------------------------------------

function js_go_gotIxRowFromGroup( gostr1 ) {
	
	/*
	js ==> go : go_passToJs_getIxRowFromGroup( ""+html_rowGroup_index_gr,  ""+html_rowGroup_beginNum, ""+html_rowGroup_numRows, "js_go_gotIxRowFromGroup")
	go ==> js :  
	outS1:= fmt.Sprintf( "inp,%d,%d,%d,rG_,%d,%s,%d,%d,ixR,%d,%d, %s",
				rowGrIndex, html_rowGroup_beginNum, html_rowGroup_numRows,
				rG.rG_ixSelGrOption, 
				rG.rG_group, 
				rG.rG_firstIxRowOfGr, 
				rG.rG_lastIxRowOfGr,
				ixRowBeg, ixRowEnd, 	
				inputTextRowSlice[  rG.rG_firstIxRowOfGr ].rRow1   )
	*/
	
	console.log("js_go_gotIxRowFromGroup gostr1=", gostr1) 

	var col1 = gostr1.split(",")
	if ( (col1.length < 12) || ( (col1[0] != "inp") || (col1[4] != "gr") || (col1[9] != "ixr") )  ) {
		console.log("Errore1 in js_go_gotIxRowFromGroup (gostr1=", gostr1 , "\n\t non inizia con inp il formato deve essere ",  	
				`\n "inp,%d,%d,%d,rG_,%d,%s,%d,%d,ixR,%d,%d, %s",
				rowGrIndex, html_rowGroup_beginNum, html_rowGroup_numRows,
				rG.rG_ixSelGrOption, 
				rG.rG_group, 
				rG.rG_firstIxRowOfGr, 
				rG.rG_lastIxRowOfGr,
				ixRowBeg, ixRowEnd, 	
				inputTextRowSlice[  rG.rG_firstIxRowOfGr ].rRow1   ` ) 
				return;  
	}
	//inp,0,1,6,gr,0,1,0,5,ixr, 0,  5, Die Elemente
	//    1 2 3  4 5 6 7 8 9   10, 11, 12     
	let x_rowGrIndex             = col1[1]
	let x_html_rowGroup_beginNum = col1[2]
	
	let html_rowGroup_numRows    = getInt(  col1[3] )
	
	let x_rG_ixSelGrOption       = col1[5] 
	let x_rG_group               = col1[6]  
	let x_rG_firstIxRowOfGr      = col1[7] 
	let x_rG_lastIxRowOfGr       = col1[8]
	
	let html_fromIxRow          =  getInt(  col1[10] )
	let html_toIxRow            =  getInt(  col1[11] )
	
	let firstRowOfGroup  = col1.slice(12).join(",")
	
	if ((x_rowGrIndex != html_rowGroup_index_gr) || (x_html_rowGroup_beginNum != html_rowGroup_beginNum ) || (x_rowGrIndex != x_rG_ixSelGrOption)) {
		console.log("Errore2 in js_go_gotIxRowFromGroup (gostr1=", gostr1 , 
			"html_rowGroup_index_gr oppure html_rowGroup_beginNum sono cambiati" ,
			" html_rowGroup_index_gr =", html_rowGroup_index_gr , " nuovo=", x_rowGrIndex, 
			" html_rowGroup_beginNum=", html_rowGroup_beginNum , " nuovo=", x_html_rowGroup_beginNum,
			" x_rowGrIndex = ", x_rowGrIndex  , "  x_rG_ixSelGrOption=", x_rG_ixSelGrOption 			
			) 
		return; 		
	}
	
	document.getElementById("id_gruppi_sel"        ).selectedIndex = html_rowGroup_index_gr;
	
	document.getElementById("id_gruppi_oSelIx" ).innerHTML = x_rG_ixSelGrOption;
	document.getElementById("id_gruppo_val"    ).innerHTML = x_rG_group;   
	document.getElementById("id_inizioGruppo"  ).innerHTML = x_rG_firstIxRowOfGr;
	document.getElementById("id_fineGruppo"    ).innerHTML = x_rG_lastIxRowOfGr;

	document.getElementById("id_gruppo_numTotRow1" ).innerHTML = getInt(x_rG_lastIxRowOfGr) -  getInt(x_rG_firstIxRowOfGr) + 1; 
	document.getElementById("id_gruppo_numTotRow2" ).innerHTML = getInt(x_rG_lastIxRowOfGr) -  getInt(x_rG_firstIxRowOfGr) + 1; 

	document.getElementById("id_ixTextAskedRow").innerHTML  = html_fromIxRow;
	document.getElementById("id_gruppi_firstRow").innerHTML = firstRowOfGroup;
	
	document.getElementById("id_gruppi_oBegNum").innerHTML  = html_rowGroup_beginNum
		
	document.getElementById("id_gruppi_iNumRows").value     = html_rowGroup_numRows   
	
	document.getElementById("id_fromIx_row"     ).innerHTML = html_fromIxRow  ;   
	document.getElementById("id_toIx_row"       ).innerHTML = html_toIxRow    ; 
	
	document.getElementById("id_inpNumRow").innerHTML = html_rowGroup_numRows ;
	
} // end of js_go_gotIxRowFromGroup 

//--------------------------------------------------

//---------------------------------------------------
function js_go_valueFromLastRun( gostr1 ) {
	
	
	var col1 = gostr1.split(",")
	if (col1.length < 16) {
		console.log("errore1 in js_go_valueFromLastRun il numero di valori tra virgola (", col1.length,") < 16" ,   " gostr1=" + gostr1 );
		return	;	
	}
	//  1,2,0,0,html,1,2,28,ix,0,0,w,1,774,extrRow, :row=,Die Elemente
	var rS_ixSelGrOption = col1[ 0 ] 
	var rS_group         = col1[ 1 ]   	
	var last_rG_firstIxRowOfGr = col1[ 2 ]      
	var last_rG_lastIxRowOfGr  = col1[ 3 ]    
				
	html_rowGroup_index_gr = col1[ 5 ] 
	
	if (html_rowGroup_index_gr < 0) {  console.log("js_go_valueFromLastRun html_rowGroup_index_gr=", html_rowGroup_index_gr) }
	
	html_rowGroup_beginNum = col1[ 6 ] 
	html_rowGroup_numRows  = col1[ 7 ]  			
	
	html_fromIxRow   = col1[ 9 ]     
	html_toIxRow     = col1[ 10 ]       
		
	var word_fromWord    = col1[ 12 ] 	
	var word_numWords    = col1[ 13 ]			
	var sel_extrRow      = col1[ 14 ] 
	
	var row                   = col1.slice(16).join(",")
	
	
	document.getElementById("id_gruppi_iNumRows").value     = html_rowGroup_numRows
	
	document.getElementById("id_fromIx_row").innerHTML = html_fromIxRow  ;    // relativo all'inizio della lista righe
	document.getElementById("id_toIx_row").innerHTML   = html_toIxRow    ;    // relativo all'inizio della lista righe
	
	
	document.getElementById("id_inpNumRow").innerHTML = html_rowGroup_numRows ;
	
	document.getElementById("id_gruppi_firstRow").innerHTML = row; 
	
	//---
	document.getElementById("id_gruppi_sel").selectedIndex    = html_rowGroup_index_gr;
	last_html_rowGroup_index_gr = html_rowGroup_index_gr;
	
	
	document.getElementById("id_gruppi_oSelIx").innerHTML = html_rowGroup_index_gr;
	
	document.getElementById("id_gruppi_iBegNum" ).value     = html_rowGroup_beginNum ;    // relativo all'inizio del gruppo 
	document.getElementById("id_gruppi_oBegNum" ).innerHTML = html_rowGroup_beginNum ;  // relativo all'inizio del gruppo 
	
	//console.log("html_fromIxRow = ", html_fromIxRow ,"  html_toIxRow=", html_toIxRow) 
	
	
	document.getElementById("id_gruppo_numTotRow1" ).innerHTML = getInt(last_rG_lastIxRowOfGr) -  getInt(last_rG_firstIxRowOfGr) + 1; 
	document.getElementById("id_gruppo_numTotRow2" ).innerHTML = getInt(last_rG_lastIxRowOfGr) -  getInt(last_rG_firstIxRowOfGr) + 1; 
	
	document.getElementById("id_gruppo_val"      ).innerHTML = rS_group;   
	document.getElementById("id_inizioGruppo"    ).innerHTML = last_rG_firstIxRowOfGr; 
	document.getElementById("id_fineGruppo"      ).innerHTML = last_rG_lastIxRowOfGr; 
	document.getElementById("id_ixTextAskedRow"  ).innerHTML = html_fromIxRow ; 
	
	document.getElementById( sel_extrRow        ).selected = "true"; 
	document.getElementById("id_inpMaxNumWords" ).value = word_numWords ; 
    document.getElementById("id_inpBegFreqWList").value = word_fromWord ;

	setLastValuesOfExtrRowChanged("js_go_valueFromLastValue");
		
} // end of js_go_valueFromLastValue()

//-------------------------
/***
<td class="borderVert"> 
				<div class="divRowText" >
					<div class="suboLine" style="display:none;" id="idc_§1§"  ondblclick="onclickDoubleRowTran(this)">§4txt§</div>
					<div class="tranLine" style="display:none;" id="idt_§1§">§5txt§<br></div>	
					<div id="idw_§1§" class="center" style="width:100%;border:0px solid red;"></div>				
					<div style="display:none;" id="idtts§1§">§ttstxt§</div>
				</div>
				-------------------- to be added ---- 
				<div>
					<div>add/modify translation</div>' 
					<div  style="background-color:lightgrey; color:black; font-weight:bold;border:2px solid black;min-width:100%;text-align:left;"
						contentEditable=true>
						newTranslation 
					</div>
					<div><button onclick="onclick_saveNewWordTran(this)">Salva tutte le nuove traduzioni</button> </div> 
				</div>
			</td>	

***/ 
//---
function onclickDoubleRowTran(this1) {
	var eleDiv = this1.parentElement;
	var eleTD  = eleDiv.parentElement; 
	if (eleTD.children.length >= 2) { return; } 	
	var ele_tran = eleDiv.children[1];	
	const newDiv = document.createElement("div");
	newDiv.style.textAlign = "left";
	eleTD.appendChild(newDiv);
	
	var newInn=""
	newInn += '<div  style="font-size:0.6em;width:100%;">add/modify translation</span></div>' + '\n';
	newInn += '<div style="background-color:lightgrey; color:black; font-weight:bold;border:2px solid black;min-width:100%;text-align:left;" ' +
		'contentEditable=true>' + 
		ele_tran.innerHTML + '</div>' + '\n'		
	newInn += '<div  style="width:100%;"><button onclick="onclick_saveNewRowTran(this)">Salva tutte le nuove righe tradotte</button></div>\n'; 
	eleTD.children[1].innerHTML = newInn;    
	eleDiv.children[1].style.display = "block"; 
} 
//-------------------------------------------------


/***			
		<tr> 
			<td style="text-align:center;font-size:0.8em;font-weight:100;">
					<span>§one-numTR§       4</span>
					<span style="display:none;">
						<span>§one-word1§   die</span>    
						<span>§ixW2StudyLs§ 3</span>
						<span>§one-ix1§     0</span>
						<span>§one-ixLemma§ 0</span>	
					</span>	
			</td>  
			...
			<td style="text-align:center;" class="borderVert_L">		                               eleTD         			
				<div class="hpad top left1" style="border:2px solid green;">                           eleDiv   
					<span onmouseover="mouseOverWord(this)" onmouseout="mouseOutWord(this)" 
						ondblclick="onclickDoubleWordTran(this)"> 
							<span class="c_wordOrig"><b>die</b></span>						
					</span> 
					<span class="c_wordTran" style="display:none;">il</span>                           ele_tran  			
				</div>
				---------------  to be added ----------------
				<div>                                                                                  newDiv 
					<span>add/modify translation</span>' 
					<div  class="c_wordTran" style="background-color:lightgrey; color:black; font-weight:bold;border:2px solid black;min-width:100%;text-align:left;"
						contentEditable=true>
						newTranslation 
					</div>
					<button onclick="onclick_saveNewWordTran(this)">Salva tutte le nuove traduzioni</button>  
				</div>
			</td>	
			...
		</tr>
***/
function onclickDoubleWordTran(this1) {
	var eleDiv = this1.parentElement;
	var eleTD  = eleDiv.parentElement; 
	
	//if (eleTD.children.length < 1) { return; } 
	
	if (eleTD.children.length >= 2) { return; } 
	
	var ele_tran = eleDiv.children[1];
	
	const newDiv = document.createElement("div");
	newDiv.style.textAlign = "left";
	eleTD.appendChild(newDiv);
	
	var newInn=""
	newInn += '<div  style="font-size:0.6em;width:100%;">add/modify translation</span></div>' + '\n';
	newInn += '<div  class="c_wordTran" ' +
		'style="background-color:lightgrey; color:black; font-weight:bold;border:2px solid black;min-width:100%;text-align:left;" ' +
		'contentEditable=true>' + 
		ele_tran.innerHTML + '</div>' + '\n'		
	newInn += '<div  style="width:100%;"><button onclick="onclick_saveNewWordTran(this)">Salva tutte le nuove traduzioni</button></div>\n'; 
	eleTD.children[1].innerHTML = newInn;    	
	
	//console.log("new eleTD=", eleTD.tagName, " ==>" , eleTD.innerHTML) ; 	
	
} // end of onclickDoubleWordTran

//----------------------------------------
function onclick_saveNewWordTran(this1) {
	//===
	let wordx, ix12, nrow, totExtrRow2,  wLemma1, wordTran, knowYesCtr, knowNoCtr   ;
	var wLemmaList, wTranList, wLevelList, wParaList, wExampleList;	
	/**
	for (var z2=0; z2 < wordToStudy_list.length; z2++) {
		[wordx, ix12,  nrow, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr  ] = wordToStudy_list[z2]; 
		console.log("wordToStudy_list",z2,"] : ",  "word=", wordx, ",ix12=", ix12, " wLemmaList=", typeof wLemmaList, " ", wLemmaList,   " wTranlist=", typeof wTranList, " ", wTranList);  
	}
	**/
	// =================
	var eleTR = this1.parentElement; 
	for(var z1=0; z1 < 10; z1++) {
		if (eleTR.tagName == "TR") { break; } 
		eleTR = eleTR.parentElement; 
	} 
	if (eleTR.tagName != "TR") { return; } 	
	var elePareTr = eleTR.parentElement; // tbody
	var eleTr2;

	let word1, ixW2StudyLs, ix1, ixLemma; 
	
	var oldDivTran, eleOldTran, oldTranslation; 
	var newDivTran, eleNewTran, newTranslation; 
	var swChg=false;
	//var listNewTranIx = [];
	//var listNewTransla= [];
	var eleTD_0, eleTD_5; var eleTD0_val;
	var newUp = 0; 
	
	/*
	<tr> 
		<td style="text-align:center;font-size:0.8em;font-weight:100;">   eleTD_0
			<span>§one-numTR§</span>
			<span style="display:none;">                                  eleTD0_val
				<span>§one-word1§</span>    
				<span>§ixW2StudyLs§</span>
				<span>§one-ix1§</span>
				<span>§one-ixLemma§</span>	
			</span>	
		</td>  
		...
		<td                                                               eleTD_5
				style="text-align:center;" class="borderVert_L" style="border:2px solid red;">					
			<div class="hpad top left1"  style="border:2px solid green;" >
				<span onmouseover="mouseOverWord(this)" onmouseout="mouseOutWord(this)" ondblclick="onclickDoubleWordTran(this)"> 
					<span class="c_wordOrig"><b>§one-word1§</b></span>						
				</span> 
				<span  class="c_wordTran" style="display:none;">§one-f_tran§</span>			
			</div>
			---------------  might have been added ----------------
			<div>                                                                                  newDiv 
				<span>add/modify translation</span>' 
				<div  class="c_wordTran" style="background-color:lightgrey; color:black; font-weight:bold;border:2px solid black;min-width:100%;text-align:left;"
					contentEditable=true>
					newTranslation 
				</div>
				<button onclick="onclick_saveNewWordTran(this)">Salva tutte le nuove traduzioni</button>  
			</div>
		</td>		
	*/
	for(var t1=0; t1 < elePareTr.children.length; t1++) {
		eleTr2 = elePareTr.children[t1];	
		eleTD_0 = eleTr2.children[0]; 
		eleTD0_val = eleTD_0.children[1]; 
		word1       = eleTD0_val.children[0].innerHTML; 
		ixW2StudyLs = eleTD0_val.children[1].innerHTML; // indice wordToStudy_List
		ix1         = eleTD0_val.children[2].innerHTML; // indice word (in go) 
		ixLemma     = eleTD0_val.children[3].innerHTML; // indice lemma x lemma e tran list 
		oldTranslation ="";
		newTranslation = ""; 
		swChg=false
		
		eleTD_5 = eleTr2.children[5]; 
		oldDivTran = eleTD_5.children[0]
		if (oldDivTran) {
			eleOldTran = oldDivTran.children[1];  
			if (eleOldTran) {
				oldTranslation = eleOldTran.innerHTML ;
			}	
		}
		newDivTran = eleTD_5.children[1]
		if (newDivTran) {
			eleNewTran = newDivTran.children[1]; 
			if (eleNewTran) {
				newTranslation = eleNewTran.innerHTML; 
				if (newTranslation != oldTranslation) {
					swChg=true; 
				}		
			} 
		} 
		if (swChg) {  
			//console.log( "ix=", ix1, " \t ", word1 , " \t oldTran=", oldTranslation, "\t newTran=", newTranslation, " ixW2StudyLs=" + ixW2StudyLs + "<==",
			//	" wordToStudy_list.length=", wordToStudy_list.length); 
			 
			[wordx, ix12,  nrow, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr  ] = wordToStudy_list[ixW2StudyLs]; 
			if (wTranList[ixLemma] == oldTranslation) {
				wTranList[ixLemma] = newTranslation;
				
				//[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList ] ???antoX
				
				wordToStudy_list[ixW2StudyLs] = [wordx, ix12, nrow, 
						wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, knowYesCtr, knowNoCtr, wIxLemmaList   ] ;
				newTran[ixW2StudyLs]=1; 
				newUp++;
				eleOldTran.innerHTML = newTranslation; 
			} 		
		} 
		if (newDivTran) {
			eleTD_5.children[1].remove(); 	
		}	
	}	
	//---------------------

	
	if (newUp > 0) {
		console.log( red(" onclick_saveNewWordTran	call  write_word_dictionary"))
		
		write_word_dictionary()
	}
	
} // end of onclick_saveNewWordTran	

//------------------------------------
//----------------------------------------
function onclick_saveNewRowTran(this1) {	
	
	var eleTR = this1.parentElement; 
	for(var z1=0; z1 < 10; z1++) {
		if (eleTR.tagName == "TR") { break; } 
		eleTR = eleTR.parentElement; 
	} 
	if (eleTR.tagName != "TR") { return; } 	
	var elePareTr = eleTR.parentElement; // tbody
	var eleTr2;


	let word1, ixW2StudyLs, ix1, ixLemma; 
	
	var oldDivTran, eleOldTran, oldTranslation; 
	var newDivTran, eleNewTran, newTranslation; 
	var swChg=false;
	//var listNewTranIx = [];
	//var listNewTransla= [];
	var eleTD_0, eleTD_5; var eleTD0_val;
	var newUp = 0; 
	let ixRow2StudyLs, nfile, idRow, ixRow, origRow, tranRow; 
	var row0, cols;
	/**
		<td class="borderVert">                 xx   eleTD_5
			<div class="divRowText">            xx   oldDivTran
				<div class="suboLine" style="display: block;" id="idc_4" ondblclick="onclickDoubleRowTran(this)">und wie sich ihr Leben allmählich veränderte,</div>
				<div class="tranLine" style="display: block;" id="idt_4">e come la sua vita è gradualmente cambiata,<br></div>	
				<div id="idw_4" class="center" style="width:100%;border:0px solid red;"></div>				
				<div style="display:none;" id="idtts4"></div>
				<div style="display:none;">4</div>
			</div>
			<div style="text-align: left;">    xx   newDivTran
				<div style="font-size:0.6em;width:100%;">add/modify translation</div>
				<div style="background-color:lightgrey; color:black; font-weight:bold;border:2px solid black;min-width:100%;text-align:left;" 
					contenteditable="true">e come la sua vita gradualmente cambiò,<br></div>
				<div style="width:100%;"><button onclick="onclick_saveNewRowTran(this)">Salva tutte le nuove righe tradotte</button></div>
			</div>
		</td>
	**/
	for(var t1=0; t1 < elePareTr.children.length; t1++) {
		eleTr2 = elePareTr.children[t1];	
		
		ixRow2StudyLs = -1;
		oldTranslation ="";
		newTranslation = ""; 
		swChg=false
		
		eleTD_5 = eleTr2.children[5]; 
		
		oldDivTran = eleTD_5.children[0]
		if (oldDivTran) {
			eleOldTran = oldDivTran.children[1];  
			if (eleOldTran) {
				oldTranslation = eleOldTran.innerHTML; 
			}
			ixRow2StudyLs  = oldDivTran.children[4].innerHTML	
		}
		newDivTran = eleTD_5.children[1];
		if (newDivTran) {			
			eleNewTran = newDivTran.children[1]; 
			if (eleNewTran) {
				newTranslation = eleNewTran.innerHTML; 
				if (newTranslation != oldTranslation) {
					swChg=true; 
				}		
			} 
		} 
		if (swChg == false) { 
			if (newDivTran) {
				eleTD_5.children[1].remove();	
			}	
			continue; 
		}  
		//console.log( "FRASI ", eleTr2.id, " ixRow2StudyLs=", ixRow2StudyLs , " orig=", oldDivTran.children[0].innerHTML, "\n\tTRAN OLD=", oldTranslation, "\n\tTRAN NEW=", newTranslation)
		if (ixRow2StudyLs < 0) { continue; } 	
		
		row0 = rowToStudy_list[ixRow2StudyLs].trim() + "|||||"; 	
		cols = row0.split("|")
		try{ 
			[nfile, idRow, ixRow, origRow, tranRow] = cols.slice(0,5); 	
		} catch(e1) {	
			continue
		}
		if (oldTranslation.indexOf(tranRow)>=0) {  // non sono esattamente eguali, oldTranslation termina con <br>
			tranRow = newTranslation;
			rowToStudy_list[ixRow2StudyLs] = nfile + "|" + idRow + "|" + ixRow + "|" + origRow + "|" + tranRow;  
			newRowTran[ixRow2StudyLs]=1; 
			newUp++;
			eleOldTran.innerHTML = newTranslation; 	
		} 								
		eleTD_5.children[1].remove();	
		
	}	// end for t1
	//---------------------

	
	if (newUp > 0) {
		write_row_dictionary()
	}
	
} // end of onclick_saveNewWordTran	


//-------------------------