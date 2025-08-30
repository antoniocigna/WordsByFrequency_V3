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
const YES1 = "Yes";
const NOT_YET1 = "Not Yet";
const MAX_NUM_WORD_LEARN = 10;
let numWordsKnownChanged = 0;
const NUM_CELL_LEMMA = 6;  
//----------------------------------
var html_fromIxRow  = 0;
var html_toIxRow  = 0;
var listaGruppiTesto=""; 
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
onchange_rowGroupSelectChange(false,10); 

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
	//console.log("result=", result)
	document.getElementById("id_frequenze").innerHTML = result;
	
} // end of js_go_updateStatistics

//---------------------------------
//---------------------------------------------------------

function onclick_mostFreqWordList_require(anyRow,toLearn) {	
	
	fun_require_mostFreqWordList( false , "HTML page onclick_mostFreqWordList_require", anyRow, toLearn);
	
} // end of onclick_mostFreqWordList_require

//-----------------------------------------

function onchange_mostFreqWordList_extrRow() {	
	
	fun_require_mostFreqWordList( true , "HTML page onchange_mostFreqWordList_extrRow",'','');
	
} // end of onchange_mostFreqWordList_require

//---------------------------------

function fun_require_mostFreqWordList( swFromOnChangeExtr , caller, anyRow="", toLearn="") {	
	
	console.log("%c	require_mostFreqWordList", "color:blue;")
	
	
	document.getElementById("id_inpBegError").style.display = "none"; 	
	
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	var swChg=false;
	var fromWord = getInt( document.getElementById("id_inpBegFreqWList").value);	
	var numWords = getInt( document.getElementById("id_inpMaxNumWords" ).value);
	if (fromWord < 1) { fromWord=1;     document.getElementById("id_inpBegFreqWList").value = 1; }
	if (numWords < 1) { numWords = 1;   document.getElementById("id_inpMaxNumWords" ).value = 1; }
	console.log("anyRow=", anyRow, " fromWord=",fromWord, " numWords=", numWords, " toLearn=", toLearn); 
	//---
	var sel_level = "any"; //  x.options[i].id;
	//---
	var x2 = document.getElementById("id_sel_2_extrRow");
    var i = x2.selectedIndex;
	
	html_sel_extrRow = x2.options[i].id; 
	if (anyRow != "") { html_sel_extrRow = anyRow; }
	console.log("html_sel_extrRow=", html_sel_extrRow)
	
	//console.log("  require_mostFreqWordList", " html_sel_extrRow =",html_sel_extrRow )
	
	if (last_sel_extrRow_freqWord_list == "") {last_sel_extrRow_freqWord_list = html_sel_extrRow; }  
	
	//console.log("  require_mostFreqWordList", " last_sel_extrRow_freqWord_list=", last_sel_extrRow_freqWord_list);
	
	is_selected_row_only = ( i == index_onlySelRowsWanted); //1 onclick_require_mostFreqWordLi
	
	console.log("  require_mostFreqWordList", " is_selected_row_only =", is_selected_row_only )
	
	fun_selRowsWanted_changed();
	//---	
	var xTbl = document.getElementById("id_sel_tblwords");
    var i2 = xTbl.selectedIndex;
	var sel_toBeLearned = xTbl.options[i2].id;    //( 0 = 'allWords'   1 = 'toBeLearned' )
	if (toLearn != "") {
		sel_toBeLearned = toLearn; 
	}
	console.log("sel_toBeLearned=" ,sel_toBeLearned); 
	
	//--
	swChg = isExtrRowChanged() 
	if (swChg) {
		if ((last_sel_extrRow_freqWord_list == html_sel_extrRow) && (html_sel_extrRow == "anyRow")) {
			swChg = false; 
			console.log("onchange_mostFreqWordList_extrRow " , " isExtrRowChanged=", true, " but sel_extrRow=anyRow as before,  reset isExtrRowChanged= false");	
		}
	}
	if (swFromOnChangeExtr) {
		console.log("onchange_mostFreqWordList_extrRow " , " isExtrRowChanged=", swChg , " return");	
		return; 
	}
	
	console.log("onclick_mostFreqWordList_require " , " isExtrRowChanged=", swChg , " go_passToJs_wordList --> js_go_showWordList_lev2  ( , button=1)");	
	
	//var caller = "HTML page onclick_require_rowList1" //  (new Error()).stack?.split("\n")[2]?.trim().split(" ")[1] ;
	//if (caller == undefined) { caller = ""; }
	
	console.log("%cfun_require_mostFreqWordList fun_require_mostFreqWordList", "color:green;"); 
	console.log("go_passToJs_wordList ", "swChg=", swChg, " fromWord=", fromWord, " numWords=", numWords, " sel_level=", sel_level, 
			" html_sel_extrRow=", html_sel_extrRow, " sel_toBeLearned=", sel_toBeLearned ); 
	
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
		console.log(wordToStudy_listStr[z], "\n\t col0=", col1[0], " col[3]=", col1[3], " freq1=", freq1, " freq2=", freq2, " key=", key)
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

//--------------------------------------------------------------

function sortWordFreqFirst( wordToStudy_listStr ) { 
	
	var col1, key1 , key2; 
	
	var MAXKEY = 1000000;  
	var listKey=[]; 
	
	//--------------------
	/*
	 0 xWordF2.uWordSeq + ";." + 
	 1 xWordF2.uWord2 + ";." + 
	 2 "ix" + ";." +
	 3 strconv.Itoa(xWordF2.uIxUnW) + ";." + 
	 4 strconv.Itoa(xWordF2.uTotRow)  + ";." +
	 5 fmt.Sprint( strings.Join(xWordF2.uLemmaL,  wSep)  ) + ";." + /fmt.Sprint( strings.Join(xWordF2.uTranL,   wSep)  ) + ";." +  
	 6 listStringLemmaSlice_Tran(xWordF2) + ";." +  
	 7fmt.Sprint( strings.Join(xWordF2.uLevel,   wSep)  ) + ";." +  
	 8 fmt.Sprint( strings.Join(xWordF2.uPara,    wSep)  ) + ";." +  
	 9		fmt.Sprint( strings.Join(xWordF2.uExample, wSep)  ) + ";." +  
	 10 strconv.Itoa(xWordF2.uTotExtrRow) + ";." +  					
					strconv.Itoa(xWordF2.uKnow_yes_ctr)  + ";." + strconv.Itoa(xWordF2.uKnow_no_ctr)  + ";." + 				
					"ixLemma" + ";." + intSliceToString( xWordF2.uIxLemmaL,wSep )  + ";." + 		
					endOfLine 
	*/
	const ix_uTotRow  = 4;
	const ix_uWordSeq = 0; 
	const ix_lemma    = 5; 
	const ix_uTotExtrRow = 6;
	const isNumber = true; 
	const isAlpha  = false; 
	const ascending = "a"; 
	const descending = "d";
	//-----------	
	
	for (var z=0; z < wordToStudy_listStr.length; z++) {		
		col1 = (wordToStudy_listStr[z].trim() + ";.;.;.;.;.;.;.;.;.;.").split(";.");		
		key1 = setKey0(isNumber, col1[ix_uTotRow ], descending, MAXKEY);
		key2 = setKey0(isAlpha,  col1[ix_uWordSeq], ascending,  MAXKEY);			
		listKey.push( key1 + ";;" + key2 + ";;" + (MAXKEY + z) + ":" + z  ); 		
		//if (z < 10) { console.log("lista ", " z=", z, " \t", col1[1] + "\t ix=", col1[3],  "\t", "  ix_uTotRow=", col1[ix_uTotRow ] ) }	
	}
	return listKey.sort();
	
} // end of sortWordFreqFirst	

//------------------------------

function sentenceOneRow( str1 ) {
	if (str1 == undefined)  { return str2; }
	var str2 = (""+str1).replaceAll("\n"," ")
	str2 = str2.replaceAll(". ",".<br>").replaceAll("? ","?<br>").replaceAll("! ","!<br>"); 
	/**
	str2 = str2.replaceAll(" 1.","<br>1.").
			replaceAll(" 2.","<br>1."). 	
			replaceAll(" 3.","<br>2."). 	
			replaceAll(" 4.","<br>3."). 	
			replaceAll(" 5.","<br>4."). 	
			replaceAll(" 6.","<br>5."). 	
			replaceAll(" 7.","<br>6."). 	
			replaceAll(" 8.","<br>7."). 	
			replaceAll(" 9.","<br>8.")  ;
	**/		
	str2 = str2.
			replaceAll("1.<br>", "1. "). 	
			replaceAll("2.<br>", "2. "). 	
			replaceAll("3.<br>", "3. "). 	
			replaceAll("4.<br>", "4. "). 	
			replaceAll("5.<br>", "5. "). 	
			replaceAll("6.<br>", "6. "). 	
			replaceAll("7.<br>", "7. "). 	
			replaceAll("8.<br>", "8. ").
			replaceAll("9.<br>", "9. ").
			replaceAll("0.<br>", "0. ") ;		
	
	/**
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
	**/
	//if (str1.indexOf("Hund") >= 0) { console.log("ANTONIO str1=" + str1 + "\nstr2=" + str2) }
	
	return str2	
}

//-------------------------------------------

function fun_selRowsWanted_changed() {
	var ele_sel = document.getElementById("id_sel_2_extrRow");
	//var ele_listBut = document.getElementById("id_list_Righe_TD_But") 
	//var ele_listNum = document.getElementById("id_list_Righe_TD_Num") 

	if (is_selected_row_only) {
		ele_sel.style.border = "4px solid blue"; 
		ele_sel.style.color  = "blue"; 
		//ele_listBut.style.border = "4px solid blue";
		//ele_listNum.style.border = "4px solid blue"; 
	} else {
		ele_sel.style.border     = null; 
		ele_sel.style.color      = null; 
		//ele_listBut.style.border = null;
		//ele_listNum.style.border = null; 
	}
	
} // end of fun_selRowsWanted_changed

//---------------------

function writeLanguageChoise() {
	
	return; 
	
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
	let word1, ix1, nrow, totExtrRow2,  wLemma1, wordTran;  
	var uLearnedYN;
	var wLemmaList, wTranList, wLevelList, wParaList, wExampleList, wIxLemmaList;
	var newTranWord=0;
	var listNewTranWords = "";
	
    for (var ixW2StudyLs = 0; ixW2StudyLs < wordToStudy_list.length; ixW2StudyLs++) {
		[word1, ix1,        nrow, wLemmaList, wTranList, wLevelList, wParaList, 
				wExampleList, totExtrRow2, uLearnedYN, wIxLemmaList ] = wordToStudy_list[ixW2StudyLs]; 
		word1 = word1.trim(); 
		if (ix1 == undefined) {
            continue;
        }  		
		if (ix1 < 0) continue; 
				
		if (newTran[ixW2StudyLs]==1) {
			if (ix1 == -1) { continue; }	
			newTranWord++; 	
			//listNewTranWords += "\n" + word1 + ";" + ix1 + ";" + wLemmaList.join( wSep ) + ";" + wTranList.join( wSep )  ;   // new line for dictionary 	
			listNewTranWords += "\n" + word1 + ";" + ix1 + ";" + wLemmaList + ";" + wTranList ;   // new line for dictionary 	
		
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

//-----------------------------------------

function onclick_reWriteAllTran() {
	
	go_rewrite_allTran(); 
	
} // end of onclick_reWriteAllTran

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

//------------------------------------------------------

function onclick_require_betweenWordList() {
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	
	var eleFromW = document.getElementById("id_inpFromABC");
	
	var fromWordPref = eleFromW.value.trim(); 		
	
	var maxNumWords = getInt(  document.getElementById("id_inpMaxNumABC" ).value );	
	if (maxNumWords < 1) {  maxNumWords = 1; document.getElementById("id_inpMaxNumABC" ).value = 1; }	
	
	if (fromWordPref == "") { return; }
	
	eleFromW.style.color = null;
	eleFromW.parentElement.style.backgroundColor = null;
	
	go_passToJs_betweenWordList(""+maxNumWords, fromWordPref, "js_go_showBetweenWordList"); // ask 'go' to give wordlist by js_... function  
	
} // end of onclick_require_prefixWordList


//------------------------------------
function js_go_showBetweenWordList(wordListStr, js_parm, jsFunc,goFunc) {
	//console.log("function js_go_showBetweenWordList () ", " js_parm=", js_parm, " <-- " + goFunc + " <-- " + jsFunc) ;
	//console.log("wordListStr=\n"+wordListStr +"\n--------------------------\n")
	
	if (wordListStr == "") {
		document.getElementById("id_bW_err").style.display ="block";   // no entry found
    } else {
		document.getElementById("id_bW_err").style.display ="none"; 
	}
	
	onclick_jumpFromToPage( myPage01,0, myPage02);  //myPage01
	
	js_go_showWordList_lev2(wordListStr, 2,jsFunc, goFunc) //2 js_go_showBetweenWordList  onclick_require_prefixWordList

} // end of js_go_showBetweenWordList


//------------------------------------------------------

function onclick_require_betweenLemmaList() {
	
	//console.log("%conclick_require_betweenLemmaList", "color:green;")
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	
	var eleFromW = document.getElementById("id_inpLemmaFromABC");
	//var eleToW   = document.getElementById("id_inpLemmaToABC"  ); 
	var fromWordPref = eleFromW.value.trim(); 		
	//var toWordPref   = eleToW.value.trim(); 		
    
	
	var maxNumLemma = getInt(  document.getElementById("id_inpLemmaMaxNumABC" ).value );	
	if (maxNumLemma < 1) {  maxNumLemma = 1; document.getElementById("id_inpLemmaMaxNumABC" ).value = 1; }	
	
	if (fromWordPref == "") { return; }
	
	/**
	if (toWordPref == "") { 
		if (fromWordPref == "") { return; }
		toWordPref = fromWordPref;
		//document.getElementById("id_inpToABC"  ).value = toWordPref; 	
	} else {
		if (fromWordPref == "") { 
			fromWordPref = toWordPref;
			document.getElementById("id_inpLemmaFromABC").value = fromWordPref; 
		}
	} 
	if (toWordPref < fromWordPref) {
		//document.getElementById("id_inpToABC"  ).value = fromWordPref; 	
		document.getElementById("id_inpFromABC").value = toWordPref;
		fromWordPref = document.getElementById("id_inpLemmaFromABC").value.trim(); 		
	    //toWordPref   = document.getElementById("id_inpLemmaToABC"  ).value.trim(); 		
	}
	**/
	eleFromW.style.color = null;
	eleFromW.parentElement.style.backgroundColor = null;
	var eleLemma = document.getElementById("id_inpLemmaFromABC");  
	eleLemma.style.color = null;
	eleLemma.parentElement.style.backgroundColor = null;
	//var eleLemma2 = document.getElementById("id_inpLemmaToABC");  
	//eleLemma2.style.color = null;
	//eleLemma2.parentElement.style.backgroundColor = null;
	
	//console.log("go_passToJs_betweenLemmaList"   , " fromWordPref=", fromWordPref)
	
	go_passToJs_betweenLemmaList(""+maxNumLemma, fromWordPref, "js_go_showBetweenLemmaList"); // ask 'go' to give wordlist by js_... function  
	
} // end of onclick_require_betweenLemmaList


//------------------------------------
function js_go_showBetweenLemmaList(lemmaListStr, js_parm, jsFunc,goFunc) {
	//console.log("js_go_showBetweenLemmaList")
	if (lemmaListStr == "") {
		document.getElementById("id_bW0_err").style.display ="block";   // no entry found
    } else {
		document.getElementById("id_bW0_err").style.display ="none"; 
	}
	
	onclick_jumpFromToPage( myPage01,0, myPage02);  //myPage01
	
	js_go_showWordList_lev2(lemmaListStr,5, jsFunc, goFunc)

} // end of js_go_showBetweenLemmaList

//-------------------------------------------------------
function onclick_require_lemmaWordList2(aLemma) {
	if (aLemma=="") return; 
	word_to_underline_list = []
	
	//console.log("%conclick_require_lemmaWordList2 ", "color:green;font-weight:bold;"); console.log("lista le parole con questo lemma")
	
	var eleMax = document.getElementById("idTabWRLL3")
	
	var inpMaxWordLemma = eleMax.value; 
    aLemma = aLemma.trim();    
	if (aLemma == "") {
		//ele_wordList.innerHTML ='<span style="color:red;">manca il lemma</span>';
		return;
	}		
	/***
	var eleFromW = document.getElementById("id_inpFromABC");
	eleFromW.style.color = null;
	eleFromW.parentElement.style.backgroundColor = null;
	var eleLemma = document.getElementById("id_inpLemma");  
	eleLemma.style.color = null;
	eleLemma.parentElement.style.backgroundColor = null;
	***/
	//myPage01.style.display = "none"; 
	var caller = "HTML page onclick_require_lemmaWordList2"
	//console.log("go_passToJs_lemmaWordList ( ", aLemma );
	go_passToJs_lemmaWordList(aLemma, inpMaxWordLemma, "js_go_showLemmaWordList," + caller);  	

	
} // end of onclick_require_lemmaWordList2
//------------------------------
/***
function onclick_require_lemmaWordList() {
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	
	var inpMaxWordLemma = document.getElementById("id_inpMaxWordLemma").value; 
   // var aLemma = document.getElementById("id_inpLemma").value.trim();    
	**
	if (aLemma == "") {
		ele_wordList.innerHTML ='<span style="color:red;">manca il lemma</span>';
		return;
	}	
	**
	var eleFromW = document.getElementById("id_inpFromABC");
	eleFromW.style.color = null;
	eleFromW.parentElement.style.backgroundColor = null;
	//var eleLemma = document.getElementById("id_inpLemma");  
	//eleLemma.style.color = null;
	//eleLemma.parentElement.style.backgroundColor = null;
	
	//myPage01.style.display = "none"; 
	var caller = "HTML page onclick_require_lemmaWordList"
	
	go_passToJs_lemmaWordList(aLemma, inpMaxWordLemma, "js_go_showLemmaWordList," + caller);  	

} // end of onclick_require_lemmaWordList
***/
//------------------------------------
function js_go_showLemmaWordList(wordListStr,  js_parm, jsFunc,goFunc) {
	
	//console.log(" js_go_showLemmaWordList () ", "wordListStr=\n" + wordListStr + "\n-------------------\n")
	
	
	if (wordListStr.substring(0,5) == "NONE,") {
		//document.getElementById("id_inpLemma_word").innerHTML = wordListStr.substring(5) 
		//document.getElementById("id_inpLemma_msg").style.display = "block"			
		myPage01.style.display = "flex"; 
		
		//onclick_jumpFromToPage( myPage02,myPage03, myPage04); 
		return
	}
	//document.getElementById("id_inpLemma_msg").style.display = "none"
	onclick_jumpFromToPage( myPage01,0, myPage02);  
	myPage01.style.display = "none"; 
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
	
	//get_first_tr_visible();  // memorizza la prima TR visibile delle frasi in cui si trova questa funzione 
	
	var eleFromW = document.getElementById("id_inpFromABC");
	//var eleToW   = document.getElementById("id_inpToABC"  ); 
	//eleToW.value = ""; 
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
	//var eleToW   = document.getElementById("id_inpToABC"  ); 
	//eleToW.value = ""; 
	eleFromW.value = word1;  
    eleFromW.style.color = "blue";
	eleFromW.parentElement.style.backgroundColor = "yellow";
	//var eleLem  = document.getElementById("id_inpLemma")
	//var eleLemTD = eleLem.parentElement
	//var eleLemTR = eleLemTD.parentElement
	//eleLem.value = lemma1; 
	//eleLem.style.color = "blue"; 
	//eleLemTD.style.backgroundColor = "yellow";
   
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
	
	//get_first_tr_visible();  // memorizza la prima TR visibile delle frasi in cui si trova questa funzione 
	
	//myPage01.style.display = "none"; 
	//go_passToJs_thisWordRowList(      aWord, ""+maxNumRow5, "js_go_showWrdRowList"); 
	go_passToJs_someWordsRowList( "", aWord, ""+maxNumRow5, "js_go_showWrdRowList");   
	
} // end of onclick_require_rowListWithThisWord2


//--------------------------

function onclick_require_rowList1(selFrasiParole12) {	
	
	//console.log("onclick_require_rowList1 selFrasiParole12=", selFrasiParole12 )
	
	document.getElementById("id_inpRowEmpty").style.display = "none";
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	//ele_wRowList.innerHTML = "";
    //ele_word.innerHTML     = ""; 	 	
	
	var inpBegRow = getInt( document.getElementById("id_fromIx_row").innerHTML );
	var numRows   = getInt( document.getElementById("id_inpNumRow" ).innerHTML );  
	var inpEndRow = inpBegRow+numRows-1;
		
	//console.log("onclick_require_rowList1 ", " XXX numRows=", numRows, " inpBegRow=" , inpBegRow, " inpEndRow=",  inpEndRow  ) 
	
	//myPage01.style.display = "none"; 
	document.getElementById("id_headWord").innerHTML = ""; //head1; 
		
	var caller = "HTML page onclick_require_rowList1" //  (new Error()).stack?.split("\n")[2]?.trim().split(" ")[1] ;
	if (caller == undefined) { caller = ""; }
	
	//console.log("2 onclick_require_rowList1 esegue go_passToJs_rowList(" + inpBegRow + "," +numRows + "," + selFrasiParole12 + ","+ "js_go_rowList" + ", " + "js_go_showWordList_lev2(1)" )
	
	go_passToJs_rowList(""+inpBegRow, ""+numRows, ""+selFrasiParole12, "js_go_rowList" , "js_go_showWordList_lev2(1)", caller); 
		
} // end of onclick_require_rowList1

//--------------
function js_go_build_rowGruppi( gruppi_option) {
	// run just after the reading of the input text 
	/*
	<option>13 file: soloParoleGoetheLista_A1.csv   (783 righe)</option>
	<option>14 Esempi da dizionario</option>
	<option>15 140 verbi irregolari</option>
	*/
	var optLine = gruppi_option.split("<option>");
	var opt1, optVV;
	listaGruppiTesto = ",";
	for(var v=0; v < optLine.length; v++) {
		opt1 = (""+optLine[v]).trim();
		if (opt1 == "") continue
		optVV = opt1.split(" ");
		if (optVV.length < 1 ) {continue;}
		listaGruppiTesto = listaGruppiTesto + trimLeftZero( optVV[0] ) +","
	}	
	//console.log("js_go_build_rowGruppi ", gruppi_option, "\nlistGruppiTesto=" + listaGruppiTesto);
	
	document.getElementById("id_gruppi_sel").innerHTML = gruppi_option;  	
	html_rowGroup_index_gr = 0;  // will be updated from last file values 
	document.getElementById("id_gruppi_sel").selectedIndex = html_rowGroup_index_gr;
	 
	setLastValuesOfExtrRowChanged("js_go_build_rowGruppi");
	
} 
//--------------------------
function getInt( sInt ) {	
	if (sInt == undefined) return -1
	try {
        return parseInt( sInt );
    } catch (err) {
		return -1
	}   
} 
//-------------------------------

function onclick_rowsByIxWord(sIxWord) {
	if (sIxWord == "") return; 
	if (getInt(sIxWord) < 0) return;
	//console.log("%conclick_rowsByIxWord", "color:green; font-weight:bold;");
	//console.log("lista le righe con questa parola", " sIxWord=", sIxWord, " getInt(sIxWord)=", getInt(sIxWord) )
	var max_num_row4word  = document.getElementById("idTabWRoW1").value  
	
    go_passToJs_getRowsByIxWord(""+sIxWord, ""+max_num_row4word, "js_go_showWrdRowList"); // ask 'go' to give the rows of the word  by the go function js_go...  

} // end of onclick_rowsByIxWord

//-------------------------------
function onclick_rowsByIxLemma(sIxLemma) {
	if (sIxLemma == "") return; 
	if (getInt(sIxLemma) < 0) return;
	var max_num_row4lemma = document.getElementById("idTabWRoL2").value 

	//console.log( "%conclick_rowsByIxLemma","color:green;font-weight:bold;" ); 
	//console.log( "LISTA le RIGHE con questo lemma ", "  sixLemma=", sIxLemma, " max_num_row4lemma=", max_num_row4lemma)

    go_passToJs_getRowsByIxLemma(""+sIxLemma, ""+max_num_row4lemma, "js_go_showLemmaRowList4"); // ask 'go' to give the rows of the word  by the go function js_go...  

} // end of onclick_rowsByIxLemma

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
//-------------------------------
function headerSomeWords(str1) {
	/*
	<div id="id_model_tSHeadW" style="display:none;"> 
		<div class="centerXY" style="width:100%;text-align:left;font-size:0.8em;">						
			<table id="tsHead_2" style="display:none;"> ...</table>
			<table id="tsHead_3" style="display:none;"> ...</table>
			...
			<table id="tsHead_4" style="display:none;">  
				<tbody id="id_model_tSHead_bdy_4">
					<tr>
						<td class="c_word" >§4word§</td>																
						<td class="c_lemma">§4lemma§</td>
						<td class="c_tranW">§4tran§</td> 
					</tr>														
			
		</tbody>
	</table>	
	*/
	/*
	var ele_model_tSHeadW_bdy;
	var model_tSHeadW_div;	
	document.getElementById("tsHead_4").style.display=" ERRORE non usare block per le table metti table-block"; 		
	ele_model_tSHeadW_bdy = document.getElementById("id_model_tSHead_bdy_4"); 	
	str2 += ele_model_tSHeadW_bdy.replace("§1lemma§",nuovoLemma).
											replace("§1tran§",     nuovoTran) + 
											"\n\n";  		
	*/
	
	document.getElementById("tsHead_4").style.display="block"; 		
	var ele_model_tSHeadW_bdy_inner = document.getElementById("id_model_tSHead_bdy_4").innerHTML ; 	
	var head1 = str1.split("\n"); 
	var str2="";
	for(var v=0; v < head1.length; v++) {
		var aLine = head1[v] + "|||";
		var col1 = aLine.split("|") 
		str2 += ele_model_tSHeadW_bdy_inner.replace("§4word§",col1[0]).replace("§4lemma§",col1[1]).replace("§4tran§", col1[2]) + "\n";  		
	} 	
	var ele_model_tSHeadW_DIV = document.getElementById("id_model_tSHeadW"); 
	var model_tSHeadW_div = ele_model_tSHeadW_DIV.innerHTML; 	
	var jBody  = model_tSHeadW_div.indexOf("<tbody");
	var jBody2 = model_tSHeadW_div.indexOf("<tr", jBody);		
	var newDiv = model_tSHeadW_div.substr(0, jBody2) +"\n" + str2.trim()  + "\n</tbody></table></div>\n";  
	
	return newDiv.replaceAll('display:none', 'display:block');  
	
} // end of headerSomeWords
//---------------------------------------
function buildHeaderTable( str1 ) {	

	//console.log( green("   buildHeaderTable") );  console.log(str1, "\n");
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
	if (str1.indexOf("some:") >= 0) {
		return headerSomeWords(str1.substring( str1.indexOf("some:") + 5) ); 
	}
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
	
	//console.log( "   2 buildHeaderTable", " oneWordOnly  =",oneWordOnly, "  oneLemmaOnly =", oneLemmaOnly ) ;  

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
		//console.log( green("1 CASO 4 "), "4XXXXXXXX many WORD  and many LEMMA XXXXXXXXXX");
		
		document.getElementById("tsHead_4").style.display="block"; 
		
		ele_model_tSHeadW_bdy = document.getElementById("id_model_tSHead_bdy_4"); 	
		
		//console.log( green("2 CASO 4 "), "4XXXXXXXX many WORD  and many LEMMA XXXXXXXXXX")
	}
	//----------------------------------
	var model_tSHeadW_lemma  = ele_model_tSHeadW_bdy.innerHTML; 
	//console.log("aaa ", model_tSHeadW_lemma)
	var model_tSHeadW13_row2 = document.getElementById("id_model_tSHead_bdy_13_row2").innerHTML; 	
	
	
	
	//-----
	switch( type ) {			
		 case 1: 
			//console.log("caso 1  una parola e un lemma")	
			case_type1_3(); break;
		 case 2: 
			//console.log("caso 2  un lemma diverse parole")			
			case_type2();
			break;
		 case 3: 
			//console.log("caso 3  una parola e diversi lemma")	
			case_type1_3(); 		
			break;
		 case 4: 
			//console.log("caso 4  diverse parole con diversi lemma")
			case_type4(); 
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
						str2 += model_tSHeadW_lemma.replace("§1lemma§", nuovoLemma).
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
			// console.log("caso 2  un lemma diverse parole")			
			/**
													<!-- case 2:  one lemma many words  --> 											
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
		var listParole = "", lis1="";
		var nuovoLemma = "",  nuovoTran = "";
		var len2=0;
		var lenT=0;
		var LENMAX = 70
		
		for(var z1=0; z1 < len1; z1++) {	
			var oneTr1 = lineTr[z1]	
			var jT = oneTr1.indexOf(":tran=");
			var jW = oneTr1.indexOf(":wordsInLemma=");
			if (z1==0) {
				nuovoLemma   = oneTr1.substring(0,jT    ).trim();
				nuovoTran    = oneTr1.substring(jT+6,jW ).trim();
			}
			var nuovoLisWord = oneTr1.substring(jW+14   ).trim();	
			
			len2 = 4 + (""+nuovoLisWord).length
			lenT = lenT + len2
			if (lenT > LENMAX) {
				listParole += lis1 + "<br>"
				lis1=""
				lenT= len2
			}
			
			//console.log("???antoA z1=", z1, " nuovoLisWord =", nuovoLisWord , " len1=", len1, " lenT=", lenT) 
			lis1 += '<span style="margin-right:2em;">' + nuovoLisWord + "</span>"
			//listParole += '<span style="margin-right:2em;">' + nuovoLisWord + "</span>"
		}
		listParole += lis1
		str2 += model_tSHeadW_lemma.replace("§1lemma§", nuovoLemma).
					replace("§1tran§",     nuovoTran). 
					replace("§1wordXlem§", listParole) + "\n\n";  	
					
	} // end of case_type2
	//-----------------------------------
	
	function case_type4() {
				console.log("caso 4  diverse parole e diversi lemma")			
					
				for(var z1=0; z1 < len1; z1++) {
					var oneTr1 = lineTr[z1]	
					var jT = oneTr1.indexOf(":tran=");
					var jW = oneTr1.indexOf(":wordsInLemma=");
					var nuovoLemma   = oneTr1.substring(0,jT    ).trim();
					var nuovoTran    = oneTr1.substring(jT+6,jW ).trim();
					var nuovoLisWord = oneTr1.substring(jW+14   ).trim();						
					str2 += ele_model_tSHeadW_bdy.replace("§1lemma§",nuovoLemma).
											replace("§1tran§",     nuovoTran) + 
											"\n\n";  									
				} // end for z1
	} // end of case_type4
	//------------------------
	
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
	
		<table id="tsHead_1" style="display:none;">
			<tbody id="id_model_tSHeadW_bdy1">
				<tr><td colspan="2" class="c_lemma">§1lemma§</td> </tr>															
				<tr><td style="width:1em;">&nbsp;</td><td class="c_word">§1wordXlem§</td></tr> 
				<tr><td style="width:1em;">&nbsp;</td><td class="c_tranW">§1tran§</td></tr> 
			</tbody>
		</table>
		
		<table id="tsHead_2" style="display:none;">
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
		document.getElementById("tsHead_2").style.display="block"; 
		ele_model_tSHeadW_bdy = document.getElementById("id_model_tSHeadW2_bdy"); 				
	} else {
		document.getElementById("tsHead_1").style.display="block"; 
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
	
	//console.log("%cjs_go_showWrdRowList ","color:red;"); console.log("  inpstr=" + inpstr);  
	
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
	myPage01.style.display = "none"; 
	
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
	
	//console.log("2 js_go_showWrdRowList ");
	
	myPage05.style.display = "none";
	
		
	var h_wordListStr = "", h_wordTab = "";
	var ks = inpstr.indexOf("</HEADER>"); 
	if (ks < 0) { return } 	
	
	//console.log("3 js_go_showWrdRowList ");
	
	var inpHeader = inpstr.substring( 0, ks + 9);
	
	//console.log("js_go_showWrdRowList inpHeader + \n" +  inpstr.substring( 0, ks + 20)   +"\n-------------------\n") 
	
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
	
	//console.log("4 js_go_showWrdRowList buildHeaderTable");
	
	h_wordTab     = buildHeaderTable( col1[1] )
	
	//console.log("5 js_go_showWrdRowList");
	//-------------------
	
	word_to_underline_list = h_wordListStr.trim().split(" ")                        
	
	//console.log("1word_to_underline_list=", word_to_underline_list.length, " parole") 
	
	var word3, ixUnW3, totRow3, wLemma3, wTran3;

	
	
	document.getElementById("id_headWord").innerHTML = h_wordTab.replaceAll('display:none','display:block').replaceAll("tsHead_1","tsHead_00") ;
	
	//console.log("6 js_go_showWrdRowList");
	
	onclick_jumpFromToPage( myPage02,myPage03, myPage04);  //   
	
	js_go_rowList( thisListRow );  
	
	
} // end of js_go_showWrdRowList  NEW
//--------------------------------------------

function js_go_rowList2( inpstr )  {

} // end of  js_go_rowList2

//--------------------------------------
function js_go_rowList( inpstr, js_parm, jsFunc,goFunc) {
	
	// triggered by go ( go_passToJs_rowList and js_go_showWrdRowList)
	
	//console.log("function js_go_rowList() js_parm=" + js_parm + "\n\t jsFunc=" + jsFunc , "\n\t goFunc=" + goFunc ) 
	//console.log("	inpstr=" +inpstr ) 
	
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
	myPage01.style.display = "none"; 
	
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
	str1 = parseInt(numUniW).toLocaleString() + " parole diverse, " + 
			" in totale " + parseInt(numTotW).toLocaleString() + " parole " + 
			"su un testo di " + parseInt(numRow).toLocaleString() + " righe"  ;
			
	document.getElementById("id_inpFileList").innerHTML	= "(" + str1 + ")" ; 	
	/**
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
	***/
	
}
//---------------
/**
function set_dragDiv() {
		dragElement(document.getElementById("dragButt1"  ), document.getElementById("dragButt1_header"  ));
		dragElement(document.getElementById("dragButt2"  ), document.getElementById("dragButt2_header"  ));
		dragElement(document.getElementById("dragButt3"  ), document.getElementById("dragButt3_header"  ));
		dragElement(document.getElementById("dragButt4"  ), document.getElementById("dragButt4_header"  ));
		dragElement(document.getElementById("dragWLsTip5"), document.getElementById("dragWLsTip5_Header")); 
}
**/
//------------------------------------
function js_go_ready( prevRun00) {
	console.log("************************* js_go_ready(" + prevRun00.trim() + ")" ); 
	
	/**
	5,de,de-DE,Microsoft Stefan - German (Germany):mainpage_value=5491,10,1,100,any,,
	**/
	
	//set_dragDiv();
	
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
		console.log("language file has been read ==>" +  lastRunLanguage) 
	} 	
  
    document.getElementById("id_start001").style.display = "none";
    myPage01.style.display = "flex";
	
    document.getElementById("id_showButt").style.display = "block";
    document.getElementById("id_start_tab").style.display = "none";
	
	//scroll_1_init() 
	
	if (prevRunLanguage != "") { 
		lastRunLanguage = prevRunLanguage
		loadPrevLang( prevRunLanguage ) 
		console.log("js_go_ready()  prevRunLanguage=", prevRunLanguage, "  js_go_ready() NON chiama  fcommon_load_all_voices()");  
	}	else {
		console.log("js_go_ready() call fcommon_load_all_voices()");
		
		fcommon_load_all_voices(); // at end calls tts_1_toBeRunAfterGotVoices()		
		// WARNING: the above function contains asynchronous code.  
		// 			Any statement after this line is executed immediately without waiting its end			
	}
	
	
	//onclick_getRowGroup(  document.getElementById("id_gruppi_sel") )
	
	let cellWord_TrTD = document.getElementById("idTableWordList_tbody").children[0].children[5]; 		
	let cellWord_TrTH = document.getElementById("idTableWordList_thead").children[0].children[5]; 		
		console.log("%ccellWord_TrTH=",cellWord_TrTH.innerHTML) 
		console.log("%ccellWord_TrTD=",cellWord_TrTD.innerHTML) 
		
		console.log("%ccellWord_TrTD  xxx =",document.getElementById("idTableWordList_tbody").rows[0]) 

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

function evidenzia( unaparola, class_targ, txtinp1) {
		
	
	var newRow="";
	unaparola = unaparola.toLowerCase().trim();
	
	var lenParola = unaparola.length; 
	
	var lowinp0 = txtinp1.toLowerCase().replaceAll("§"," ") + "§"; 
	var lowinp1 = lowinp0.replace(/[\s;,:"'\.<>»«()\[\]\!\?„“]/g,"§")  
	
	var j0=-1, j1=0;
	var jNew=0
	var swBold=false;
	var parola1, parolaT;
	j1=-1; 
	for(var i=0; i < lowinp1.length; i++) {	    
		j0=j1+1
		j1 = lowinp1.indexOf("§", j0);  
		if (j1 < 0) break;
		parola1 = lowinp1.substring(j0,j1) 
		parolaT = stdCode(parola1)
		newRow += txtinp1.substring(jNew, j0) 	
		swBold = (parolaT == unaparola) 
		if (swBold) {
			newRow += '<span class="' + class_targ + '">'
		}	
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
} // end of onclick_jumpFromTo1_2Page

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
	
	
	// if all words are translated continue 
	// is some translations are missing, ask if they must be ignored, otherwise loop till no translations are missing     
	if (type == 1) {
		sw_ignore_missTranWord = true;  // Ignore missing translations and continue
	} else {
		sw_ignore_missTranWord = false; // Translations added, check again</button	
	}  	
	
	var wordTranList = ele_wordsTranslated.value.trim().split( wordTTEnd );    // 
	
	console.log("X3 onclick_showWordsButton() ", " type=", type, " wordTranList=\n", wordTranList ) 	
	
	var lenTran =  wordTranList.length;	
	var wordTran
		
	var ixUnTrad, ixUnWtS;
	var lenW = wordToStudy_list.length;
	var word1, nrow, totExtrRow2, totExtrRow2, wLemma1, wtran, uLearnedYN  ; 
	var newTran_f = "";
	//------------------------------------------
	var wX, wT, ixz1,  ixz2, ixzNum   
	
	//-----------
	var wList;
	var sw_someTranMissingW = false;  // sometimes the automatic translator does some mistakes 
	var wIxLemmaList, wLemmaList, wTranList, newTranList , wLevelList, wParaList, wExampleList; 
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
		[word1, ixUnWtS, nrow, wLemmaList,   wTranList, wLevelList, 
				wParaList, wExampleList, totExtrRow2, uLearnedYN, wIxLemmaList  ] = wordToStudy_list[ixzNum]; 
		if (word1== "Xabfahren") { console.log("word1=", word1, " wParaList=", wParaList, " wExampleList=", wExampleList)}
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
			console.log("%conclick_showWordsButton  1 return ", "color:red;")
			return;
		}		
	}
	//console.log("onclick_showWordsButton() wordToStudy_list.length=" , wordToStudy_list.length)
	for (var i = 0; i < wordToStudy_list.length; i++) {
		[word1, ixUnWtS, nrow,   wLemmaList, wTranList, 
					wLevelList, wParaList, wExampleList, totExtrRow2, uLearnedYN , wIxLemmaList  ] = wordToStudy_list[i];  
		if (word1== "Xabfahren") { console.log("word1=", word1, " wParaList=", wParaList, " wExampleList=", wExampleList)}			
		//console.log("onclick_showWordsButton() 3 i=",i, " wordToStudy_list[i]=" , wordToStudy_list[i])
		if (word1 == "") { continue; }
				
		for(var ixLemma = 0 ; ixLemma < wLemmaList.length; ixLemma++) {	
			if (wLemmaList[ixLemma] == "") { continue }	  // word senza lemma 
			if (wTranList[ixLemma] == "") {	
				sw_someTranMissingW = true;  
				if (sw_ignore_missTranWord == false) {
					fun_showWordList("2", i)
					//console.log("fun_showWordList(2,", i)
					console.log("%conclick_showWordsButton  2 return ", "color:red;")
					console.log("errore in wordByFrequency.js onclick_showWordsButton  riga 2704 "); 
					return;	
				}				
			}
		}
	}	
	
	console.log(" 1 onclick_showWordsButton() call write word dictionary ") 
	
	write_word_dictionary(); 
	
	//console.log("onclick_showWordsButton() ",red("call write word dictionary ")) 
	
	if (sw_Minus1) {
		for (var i = 0; i < wordToStudy_list.length; i++) {  
			[word1, ixUnWtS, nrow,   wLemmaList, wTranList, wLevelList, 
						wParaList, wExampleList, totExtrRow2, uLearnedYN, wIxLemmaList ] = wordToStudy_list[i]; 
			if (word1== "Xabfahren") { console.log("word1=", word1, " wParaList=", wParaList, " wExampleList=", wExampleList)}				
			if (ixUnWtS < 0) {
				wordToStudy_list[i] = [word1, 0, nrow,   wLemmaList, wTranList, wLevelList, 
							wParaList, wExampleList, totExtrRow2, uLearnedYN, wIxLemmaList  ] 
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
	
	prev_voiceLang2 = prevLanguage
	
	// language=5,de,de-DE,Microsoft Stefan - German (Germany)   ( it's in the first line of dictionary) 
	//          0,1 ,2    ,3  
	
	/***
	var pvLang = prevLanguage.trim).split(",") 	
	prev_voice_ix        =  pvLang[0] ; 	
	prev_voiceLang2      =  pvLang[1] ; 		
	prev_voiceLangRegion =  pvLang[2] ;  
	prev_voiceName       =  pvLang[3] ; 
	**/
	 
	console.log(" wordByFrequence.js loadPrevLang()" + " prev_voiceLang2 = "+  prev_voiceLang2 );  
	
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
	if ((ele_wordset == null) || (ele_wordset == false) ) {
			console.log("onclick_tts_seeWordsGO1(this1," , numId00, ") ==> ", eleTR.outerHTML, "\n\tid_analWords=" + id_analWords , " ERROR ele_wordset == false" )
		return
	}
	
	get_first_tr_visible();  // memorizza la prima TR visibile delle frasi in cui si trova questa funzione 
	
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
	divWord += `<div style="font-size:0.6em;color: black;text-align:center;margin-top:0.5em;">					
				clicca su una parola per ottenere paradigma, traduzione, esempi, toccala senza cliccare per avere la traduzione   							
			</div> \n`; 
	divWord += table_txt;
	
	divWord += anyOtherWord();  // $$anto see appl: lineByLine_v3,  file: lbl2_lineByLine_script3_GOHTML.js  function: onclickSelectWord
	
	//if (table_txt.indexOf("creazione")>= 0 ) { console.log("table_txt=" , table_txt); }
	ele_wordset.innerHTML = divWord; 
	
	//console.log("\nXXXXXXXXXXXXX\nXXXXXXXXXXXX\n ele_wordset.innerHTML = " + ele_wordset.innerHTML ) 
	//console.log("ixLastEle=", ixLastEle)
	if (ixLastEle > 0) {
		//console.log("ULTIMO")
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
 // XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

    function tts_3_spezzaRiga3(rowWordList) {

        //console.log("%cspezzaRiga", "color:red;");

        let prototype_one_lemma = `
							<div style="text-align:left;font-size:0.8em;padding-left:1em;">
								<span class="c_paradigma2" style="font-weight:bold;">§7alemma§</span>	
								&nbsp;&nbsp;&nbsp;<span class="c_wordTran2" >§7atran§</span>	
								<span class="c_example2"  >§7aesempi§</span>
							</div> `; // end of prototype_one_lemma 
        //--
        let prototype_all_lemma_header = `<div style="display:none" onclick="onclick_showHideLemma(this)"> `;

        let all_lemma_for_thisWord = "";


        let endix2 = -1;
        let maxNumRow5 = 100;
        let listaParole = [];
        let listaParo_tts = [];
        let listaParo_lemma = [];
        let listaParo_tran = [];
        let listaParo_nFrasi = [];

        let row1, k1, k2, col1;
        let part1;
        let Xword1;
        let listLemS, listTranS;
        let listLem, listTran;
        let xWord_numFrasi;

        let pezRiga, pezzo0, pezzoN, pez2;
        let lemma, paradigma, traduz, esempi, numWords;
        let parola1, paro_nFrasi, paro_tts;
        let trad1;
        let frase_showTxt = prototype_word_table_header; //  '<div><table>
        //--------------



        //-------------
		 let last1 = 0;
        //--------------------------------------------------
        for (let k = 0; k < rowWordList.length; k++) {

            let unaRiga = rowWordList[k];
            //console.log("k=", k, "   ", unaRiga);

            pezRiga = unaRiga.split("[");
            if (pezRiga.length < 1) continue;
            pezzo0 = pezRiga[0];
            //-----------

            part1 = (pezzo0 + ";;;;").split(";");
            col1 = part1[0].split(",");
            parola1 = col1[0].trim();
			if (parola1 == "") continue;
            paro_nFrasi = col1[2];
            paro_tts = parola1;
            //-------------------
            //console.log("pezzo0=", pezzo0, "    pezRiga.length=", pezRiga.length);

            all_lemma_for_thisWord = prototype_all_lemma_header;
            trad1 = "";
			let traduz1;
            for (let p = 1; p < pezRiga.length; p++) {
                pezzoN = pezRiga[p].replace("]", "");
                [lemma, paradigma, traduz, esempi, numWords] = ((pezzoN + ";;;;;").split(";")).slice(0, 5);
				traduz1 = " " + traduz+" ";
				if (trad1.indexOf(traduz1) < 0) { trad1 += traduz1;	}				
                if (paradigma != "") lemma = paradigma;
				if (esempi !="") esempi = "<br>" + esempi;
                all_lemma_for_thisWord += prototype_one_lemma.replaceAll("§7alemma§", lemma).replaceAll("§7atran§", traduz).replaceAll("§7aesempi§", esempi);
            }
            all_lemma_for_thisWord += "\n						</div>";
			trad1 = trad1.trim().replaceAll("  ",", "); 
            frase_showTxt += getWord_tr2(k, parola1, paro_tts, trad1, all_lemma_for_thisWord, paro_nFrasi, maxNumRow5) + "\n";
			last1 = k;
        } // end for k 	

       

        frase_showTxt += prototype_word_table_end; // </table></div>

        //console.log("%cfine frase showword", "color:red;");
        //console.log("last1=", last1);

        return [last1, frase_showTxt];

    } //  end of  spezzaRiga3()

    //-------------------------------------
  

// ===================================================================
function OLDtts_3_spezzaRiga3( rowWordList ) {
	
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

} //  end of  OLDspezzaRiga3()

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
		
		//console.log("  xxx  NUOVA rowToStudy_list[ouIxRow=",ouIxRow , "] = " + rowToStudy_list[ouIxRow]	)
		
		newAddedTran++; 
					
	
	}  // end of for(var z ...
	//--------------------------	
	
	[numeroTS_Row, numeroTS_OkTran, numeroTS_NoTran ] = build_Page1_rowsToTranslate("2onclick_showRowsButton(type=", type);
	
	document.getElementById("id_notTranNumRow").innerHTML = numeroTS_NoTran; 
	
	write_row_dictionary(1); 
	
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
	
	//console.log("showRowsAndTranButton  (wh=",wh)
	
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
		
		//console.log(  "showRowsAndTranButton () XXX i1=", i1, " rowToStudy_list[i1]=", rowToStudy_list[i1] );  	
		
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
	
	//console.log("showRowsAndTranButton 2")
	//---------------------------
	//var endLine1 = "_endLine1_" 
	//newRowList1.push(endLine1);
	var newRowList = newRowList1.concat(newRowList2);
	//---------------------------
	var iNumTr =0;
	var idRow1, idRow2; 
	
	var PREF_MARKER = ":PREF:"
	
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
		
		var unaparola; 
		var class_targList = [ "c_wordTarg", "c_wordTarg2" ]
		var type=0;
		
		for(var hx=0; hx < word_to_underline_list.length; hx++) {  
			unaparola = word_to_underline_list[hx].trim()
			if (unaparola == "") { continue }
			if (unaparola == PREF_MARKER) {type=1;  continue; } 
			
			txt1p = evidenzia(unaparola, class_targList[type], txt1p); 			
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
		
		
		//let txt1p_n   =   txt1p.replaceAll("%20"," ").replaceAll("/","/ ").replaceAll("</ ","</");
		//let tranRow_n = tranRow.replaceAll("%20"," ").replaceAll("/","/ ").replaceAll("</ ","</");	
		let txt1p_n   =   txt1p
		let tranRow_n = tranRow
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
		
		if (first < 0) first = iNumTr;
		last = iNumTr;
		
		showList    += riga + "\n";	
				
		
   } // end for i
	//---------------------------------------------------
	eleTabSub_tbody.innerHTML = showList;
	
	//console.log("showRowsAndTranButton 3", " first=", first,  " last=", last)
	
	//scroll_1_init()
	
	if ( (last - first) > 0) {
		let eleF = document.getElementById("b1_" + first);
		let eleT = document.getElementById("b2_" + last);
		onclick_tts_arrowFromIx(eleF, first, 5);
		onclick_tts_arrowToIx(  eleT, last , "3showRowsAndTranButton" );		
	}
	//console.log("showRowsAndTranButton 4")
	
	onclick_jumpFromToPage( myPage04,0, myPage05);  
	
	//eleTabSub_tbody.scrollIntoView(true);  // non sempre l'elemento risulta preciso al top , forzo invece zero  direttamente sull'elemento scrollabile  
	
	eleTabSub_tbody.parentElement.parentElement.scrollTop = 0;
	
} // end of showRowsAndTranButton

//-------------------------------------------------

function write_row_dictionary(wh) {
	var nfileW,ixRowW, rowW, tranW; 
	var word1, ix1, nrow, wLemma1, wordTran, col1;
	
	var newTranRow=0;
	var listNewTranRows = "";
	
	//console.log("\n----------------------\nwrite_row_dictionary() rowToStudy_list=" ,  rowToStudy_list ,"\n----------")
	var idRow1, ixRow1; 
	var id_ix_Row;
	
	for (var i = 0; i < rowToStudy_list.length; i++) {
		if ( rowToStudy_list[i] == "") continue; 
		//console.log("2977write " ,  rowToStudy_list[i] )
		
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
	console.log("write_row_dictionary (",wh,") ",  newTranRow, " righe tradotte"); 	
	
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

function mouseOverWord(this1,nch) {
	if (this1.parentElement.children.length > nch) { 
		this1.parentElement.children[nch].style.visibility = "visible";
	}
}
//---------------------
function mouseOutWord(this1,nch) {
	if (this1.parentElement.children.length <= nch) { 
		return;
	}
	var ele_tran = this1.parentElement.children[nch];	
	if (ele_tran == undefined) return 
	//if (ele_tran.contentEditable == true) { return; }		
	try {
		ele_tran.style.visibility = "hidden";
	} catch(e) {
		console.log("error ", e , "\n\t this1.parentElement.parentElement=", this1.parentElement.parentElement.innerHTML)
	}		
} // end of mouseOutWord
//-------------------------------
function TOGLImouseOverWord(this1,nch) {
	if (this1.parentElement.children.length > nch) { 
		this1.parentElement.children[nch].style.display = "inline-block";
	}
}
//---------------------
function TOGLImouseOutWord(this1,nch) {
	if (this1.parentElement.children.length <= nch) { 
		return;
	}
	var ele_tran = this1.parentElement.children[nch];	
	if (ele_tran == undefined) return 
	//if (ele_tran.contentEditable == true) { return; }		
	try {
		ele_tran.style.display = "none";
	} catch(e) {
		console.log("error ", e , "\n\t this1.parentElement.parentElement=", this1.parentElement.parentElement.innerHTML)
	}		
} // end of mouseOutWord
//-------------------------------

function OLDmouseOverWord2(this1) {
	if (this1.children.length > 1) { 
		this1.children[2].style.display = "inline-block";
	}
}
//---------------------
function OLDmouseOutWord2(this1) {
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
	console.log("js_go_word=", str1);   
}// end of js_go_word_known 	
	
//-------------------------------------------------	
		
function onclick_sortWordBy_ixField(nField1,nChild1,isNumber1,ascending1, 
									nField2,nChild2,isNumber2,ascending2,swTran = false) {
	
	if (arguments.length < 8) {
		console.log("error: wrong number of arguments in onclick_sortWordBy_ixField( ", nField1,nChild1,isNumber1,ascending1, 
									nField2,nChild2,isNumber2,ascending2, 
					"\n check  'prototype_tableWordList_Header'  in 'wordsByFrequency_prototype_script.html_js' file " ); 		
	}
	/**
	console.log("onclick_sortWordBy_ixField(", nField1,",",nChild1,",", isNumber1,",",ascending1, ",",
									nField2,",",nChild2,",",isNumber2,",",ascending2,",",swTran)
	**/									
									
	document.getElementById("id_wordList1").scrollTop = 0;
		
	var ele_tbody = document.getElementById("idTableWordList_tbody"); 
	var num_tr = ele_tbody.children.length; 
	var ele_tr, ele_td, ele_butt,  ele_div0, ele_tran;  
	var num_td =0; 
	var trad="";
	
	const EMPTY  = "_none_"; 
	
	var listKey=[ EMPTY ]; // lascio l'entrata 0 occupata 
	var key1 , key2; 
	var ke2, ix1, ix2; 
	var MAXKEY = 1000000;  
	var ele_td5 ;
	//---------------------------------
	for(var g=0; g < num_tr; g++) {
		//if (g > 20) { break; }
		
		ele_tr = ele_tbody.children[g]; 
		
		//if (g==2) {console.log("onclick_sortWordBy_ixField()1.1 ", " num_tr=", g, " ==> TR=", ele_tr.innerHTML);}
		
		num_td = ele_tr.children.length; 
		//--
		ele_td = ele_tr.children[nField1]; 	
		ele_butt = ele_td 		
		for(var f2=0; f2 < 10; f2++) {
			if (ele_butt.children.length == 0) break
			ele_butt = ele_butt.children[0]; 
		} 		
		key1 = setKey0(isNumber1, ele_butt.innerHTML, ascending1, MAXKEY);			
		if (swTran) {
			trad="2";		
			ele_div0 = ele_td.children[0];
			ele_tran = ele_div0.children[1].children[3];
			if (ele_tran) {
				if (ele_tran.innerHTML.trim() != "" )  trad="1"; 
			}
			key1 = trad + "_" + key1;	
			ele_td5 = ele_tr.children[5];	
			
			if (trad == "1") {
				ele_td5.style.backgroundColor = "lightgrey";				
			} else {
				ele_td5.style.backgroundColor = "white";		
			}	
		}	
		
		//--
		ele_td = ele_tr.children[nField2];	
		ele_butt = ele_td 		
		for(var f2=0; f2 < 10; f2++) {
			if (ele_butt.children.length == 0) break
			ele_butt = ele_butt.children[0]; 
		} 	
		key2 = setKey0(isNumber2, ele_butt.innerHTML, ascending2, MAXKEY);		
		
		listKey.push( key1 + "::" + key2 + "::" + (MAXKEY + g)  ); 
		
		//console.log("SORT key1=" + key1 + " KEY2 ="+key2); 
	} 
	//----------------------------------------
	
	//console.log("\nprima di sort ", listKey.join("\n") )
	
	listKey.sort(); 
	
	//console.log("\ndopo di sort ", listKey.join("\n") )
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
			if (ele_tr.children.length > 0) { 
				var child0 = ele_tr.children[0]; 
				if (child0.children) {
					if (child0.children.length > 0) {						
						gg++
						child0.children[0].innerHTML = gg;						
						//newTd = '<td style="text-align:center;">' + gg + '</td>' ; 
						//ele_tr.children[0].outerHTML = newTd; 	
					}					
				}
			}				
			newBodyInner += ele_tr.outerHTML + "\n"; 			
		} 
		
	} 
	//-----------------------
	
	ele_tbody.innerHTML = newBodyInner; 	
	
	//-------------------------------


} // end of onclick_sortWordBy_ixField

//------------------------------------------
/**
	function setKeyTD(isNumber, tdElem, ascending, MAXKEY) {
		try 
		
		if (!isNumber) {			
			return keyAlphaCod(sNum);		
		}		
		var num1 =0;	
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
**/
	//---------------

//------------------------------------------
	function setKey0(isNumber, sNum, ascending, MAXKEY) {
		if (!isNumber) {			
			return (keyAlphaCod(""+sNum)).trim();		
		}		
		var num1 =0;	
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
	
	
	var newBodyInner = "";
	
	
	const YESNO_NO_field = 3; 
	
	
	for(var g=0; g < num_tr; g++) {
		
		ele_tr = ele_tbody.children[g]; 		
		ele_td = ele_tr.children[YESNO_NO_field];  
		if (ele_td.innerHTML == YES1) {
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


function onchange_rowGroupSelectChange(sw_newGr,where) {
	
	// <select ele_gruppi ></select>	
	var ele_gruppi = document.getElementById("id_gruppi_sel"          )
	var ele_begNum = document.getElementById("id_gruppi_iBegNum"  )
	var ele_numRow = document.getElementById("id_gruppi_iNumRows" )

	if (ele_gruppi.selectedIndex < 0) { return; }
	
	//console.log("%conchange_rowGroupSelectChange", "color:blue;"); console.log("onchange_rowGroupSelectChange ", "sw_newGR=", sw_newGr, " where=", where, "  ele_gruppi.selectedIndex=", ele_gruppi.selectedIndex, " ele_begNum",ele_begNum.id, " ==>", ele_begNum.value  ) ;
	
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
					<div style="display:none;">§ixRow2StudyLs§</div>
				</div>			
				-------------------- to be added BY DOUBLE CLICK ---- 
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
function onclickModifyRowTran(this1) {
	var eleDiv = this1.parentElement;
	var eleTD  = eleDiv.parentElement; 
	var newDiv;
	if (eleTD.children.length >= 2) { 		// elimina div che permette variazione/immissione traduzione della riga 
		newDiv = eleTD.children[1]; 	
		newDiv.remove();
		return; 
	} 	
	var ele_tran = eleDiv.children[1];		// aggiunge div che permette variazione/immissione traduzione della riga 
	newDiv = document.createElement("div");
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
	/**
		<td style="text-align:center;" class="fixWordTdWidthN borderVert_L">
			<div class="c_size_1_line">
				<div class="hpad top left1">
					<span ondblclick="onclickDoubleWordTran(this)"> 
						<span class="c_wordOrig"><b>zwei</b></span>						
					</span> 					
					<span class="c_wordTran" style="display:none;"></span>		
				</div>					
				
				insert here
				
			</div>	
		</td>	
	**/
	
	
	
	var eleDiv0  = this1.parentElement;
	var eleDiv1  = eleDiv0.parentElement; 
	var eleTD    = eleDiv1.parentElement; 
	
	var eleLemmaTD   = eleTD.nextElementSibling
	var eleLemDiv1   = eleLemmaTD.children[0]
	var eleLemmaSpan = eleLemDiv1.children[0]
	if (eleLemmaSpan.innerHTML == "") return; 
	
	
	//if (eleTD.children.length < 1) { return; } 
	
	//console.log("%conclickDoubleWordTran", "color:red;"); console.log(" 1 eleTD=", eleTD.outerHTML)
	
	//if (eleTD.children.length >= 2) { return; } 
	if (eleDiv1.children.length >= 2) { return; } 
	
	var ele_tran = eleDiv0.children[1];
	
	const newDiv = document.createElement("div");
	newDiv.style.textAlign = "left";
	//eleTD.appendChild(newDiv);
	eleDiv1.appendChild(newDiv);
	//console.log(red("medio eleTD="), eleTD.outerHTML) ; 	
	var newInn=""
	newInn += '<div  style="font-size:0.6em;width:100%;">add/modify translation</span></div>' + '\n';
	newInn += '<div  class="c_wordTran" ' +
		'style="background-color:lightgrey; color:black; font-weight:bold;border:2px solid black;min-width:100%;text-align:left;" ' +
		'contentEditable=true>' + 
		ele_tran.innerHTML + '</div>' + '\n'		
	newInn += '<div  style="width:100%;"><button onclick="onclick_saveNewWordTran(this)">Salva tutte le nuove traduzioni</button>' +
				'</div>\n'; 
	//eleTD.children[1].innerHTML = newInn;    	
	eleDiv1.children[1].innerHTML = newInn;    
	
	//onclick_vertResizeLemma(this);
	
	
	//console.log( "    new eleTD=", eleTD.outerHTML) ; 	
	/**
	new eleTD= 
	<td style="text-align:center;" class="borderVert_L">					
		<!--
		<div class="hpad top left1">
				<span class="c_wordOrig">   ondblclick="onclickDoubleWordTran(this)"<b>leben</b></span>		
		</div>
		-->
		<div class="hpad top left1">
			<span ondblclick="onclickDoubleWordTran(this)"> 
				<span class="c_wordOrig"><b>leben</b></span>						
			</span> 					
			<span class="c_wordTran" style="display:none;">vivere</span>		
		</div>
		
		<div style="text-align: left;">
			<div style="font-size:0.6em;width:100%;">add/modify translation</div>
			<div class="c_wordTran" style="background-color:lightgrey; color:black; font-weight:bold;border:2px solid black;min-width:100%;text-align:left;" 
					contenteditable="true">
					vivere
			</div>
			<div style="width:100%;">
				<button onclick="onclick_saveNewWordTran(this)">Salva tutte le nuove traduzioni</button>
			</div>
		</div>
	</td>
	
	**/
} // end of onclickDoubleWordTran


//------------------------
function vertResizeWord(this1) {	
	var class01 = "c_size_1_line";
	var classNN = "c_size_nn_line";
	var divToResize = this1.parentElement.parentElement.parentElement
	var swClass = ( divToResize.classList.contains( class01 )	)
	//----------------------
	if (swClass) {	
		divToResize.classList.remove( class01 );			
		divToResize.classList.add(    classNN );		
	} else {	
		divToResize.classList.remove( classNN );			
		divToResize.classList.add(    class01 );	
	} 	
} // end of onclick_vertResizeWord 


//----------------------------------------
function TOGLIonclick_saveNewWordTran(this1) {
	//===
	let wordx, ix12, nrow, totExtrRow2,  wLemma1, wordTran, uLearnedYN ;
	var wLemmaList, wTranList, wLevelList, wParaList, wExampleList,  wIxLemmaList;	
	
	// =================
	//console.log("1 onclick_saveNewWordTran"); 
	if (this1 == null) return; 
	
	vertResizeWord(this1);
	
	var eleTR = this1.parentElement; 
	var eleTD;
	var numCellWrd = -1; 
	for(var z1=0; z1 < 10; z1++) {
		if (eleTR == null) break; 
		if (eleTR.tagName == "TR") { break; } 
		if (eleTR.tagName == "TD") { 
			eleTD = eleTR; 
			//console.log("eleTD=", eleTD.outerHTML); 
			numCellWrd = eleTD.cellIndex;    // indice della td dove si trova onclick_saveNewWordTran
		} 
		eleTR = eleTR.parentElement; 
	} 	
	
	if (eleTR == null) return; 
	if (eleTR.tagName != "TR") { return; }
	if (numCellWrd < 0) {
		console.log("%error onclick_saveNewWordTran cellIndex < 0","color:red;");	
		console.log("    eleTr=", eleTr.outerHTML); 
		return;	
	}	
	
	//console.log("2 onclick_saveNewWordTran", " cellIndex=", numCellWrd, " TR outerHTML=", eleTR.outerHTML ); 
	
	var elePareTr = eleTR.parentElement; // tbody
	if (elePareTr == null) return; 
	
	var eleTr2;

	let word1, ixW2StudyLs, ix1, ixLemma; 
	
	var oldDivTran, eleOldTran, oldTranslation; 
	var newDivTran, eleNewTran, newTranslation; 
	var swChg=false;
	//var listNewTranIx = [];
	//var listNewTransla= [];
	var eleTD_0, eleTD_5; var eleTD0_val;
	var eleTD_6, ele_details, ele_details, ele_summ, ele_tranLemma; 
	
	var newUp = 0; 
	
	/*
	il primo TD contiene il numero d'ordine visibile seguito da parola, 2 indici, indice lemma (visbile per mouse over)  es:  6  die 3 1 1 4472   
	
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
	//console.log("3 onclick_saveNewWordTran"  , "  elePareTr.children.length=", elePareTr.children.length ); 
	
	for(var t1=0; t1 < elePareTr.children.length; t1++) {
		eleTr2 = elePareTr.children[t1];	// scan su tutti  i TR ( tutte le righe )
		if (eleTr2 == null) { continue; }
		eleTD_0 = eleTr2.children[0]; 
		if (eleTD_0 == null) {continue; }
		eleTD0_val = eleTD_0.children[1]; 
		if (eleTD0_val == null) {continue; }
		word1       = eleTD0_val.children[0].innerHTML; 
		ixW2StudyLs = eleTD0_val.children[1].innerHTML; // indice wordToStudy_List
		ix1         = eleTD0_val.children[2].innerHTML; // indice word (in go) 
		ixLemma     = eleTD0_val.children[3].innerHTML; // indice lemma x lemma e tran list 
		oldTranslation ="";
		newTranslation = ""; 
		swChg=false
		
		eleTD_5 = eleTr2.children[ numCellWrd ];          // colonna WORD   dobe si trova onclick_saveNewWordTran 
		if (eleTD_5 == null) {console.log("4 cont"); continue; }
		
		
		//console.log("numCellWrd=",numCellWrd ,  "  eleTD_5=", eleTD_5.outerHTML)
			/*
				<td style="text-align:center;" class="borderVert_L">				
					<div class="hpad top left1"  >
						<span ondblclick="onclickDoubleWordTran(this)"> 
							<span class="c_wordOrig"><b>§one-word1§</b></span>						
						</span> 					
						<span  class="c_wordTran" style="display:none;">§one-f_tran§</span>		
					</div>
					<div style="text-align: left;">
						<div style="font-size:0.6em;width:100%;">add/modify translation</div>
						<div class="c_wordTran" style="background-color:lightgrey; color:black; font-weight:bold;border:2px solid black;min-width:100%;text-align:left;" 
								contenteditable="true">
								vivere
						</div>
						<div style="width:100%;">
							<button onclick="onclick_saveNewWordTran(this)">Salva tutte le nuove traduzioni</button>
						</div>
					</div>
				</td>	
			*/
	
		var eleTD_5div = eleTD_5.children[0];
		oldDivTran = eleTD_5div.children[0];
		if (oldDivTran) {
			eleOldTran = oldDivTran.children[1];  
			if (eleOldTran) {
				oldTranslation = eleOldTran.innerHTML ;
			}	
			//console.log("oldTranslation =",  oldTranslation )
		}
		newDivTran = eleTD_5div.children[1];
		if (newDivTran) {
			//console.log(" newDivTran=", newDivTran.outerHTML)
			eleNewTran = newDivTran.children[1]; 
			if (eleNewTran) {
				newTranslation = eleNewTran.innerHTML; 
				//console.log(" newTranslation=", newTranslation)
				if (newTranslation != oldTranslation) {
					swChg=true; 
				}		
			} 
		} 
		/***				
			<td style="text-align:center;" class="borderVert_R">					
				<span style="display:none;">§one-f_lemma§</span>
				<div class="hpad top left1">
					<details>
						<summary §summarystyle§> 				
								 <span onmouseover='mouseOverWord(this)' onmouseout='mouseOutWord(this)' "> 								 
										<b>§one-f_lemma§</b>
								 </span>
								 <span  class="c_wordTran" style="display:none;">§one-f_tran§</span>		
						</summary>
						<div> 					
		***/
		
		if (swChg) {  
			//console.log( "ix=", ix1, " \t ", word1 , " \t oldTran=", oldTranslation, "\t newTran=", newTranslation, " ixW2StudyLs=" + ixW2StudyLs + "<==",
			//	" wordToStudy_list.length=", wordToStudy_list.length); 
			 
			[wordx, ix12,  nrow, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2,uLearnedYN, wIxLemmaList  ] = wordToStudy_list[ixW2StudyLs]; 
			/**
			console.log( "wordx=", wordx, " ixLemma=", ixLemma, " wLemmaList type=",typeof wLemmaList, " ",  wLemmaList, 
				" wTranList=", wTranList) 
			**/
			/**
			if (wTranList[ixLemma] == oldTranslation) {
				wTranList[ixLemma] = newTranslation;
			**/	
			if (wTranList == oldTranslation) {
				wTranList = newTranslation;
				/**
				if (wordx== "Xabfahren") { 
					console.log("wordx=", wordx, " wParaList=", wParaList, " wExampleList=", wExampleList)}	
				**/	
				wordToStudy_list[ixW2StudyLs] = [wordx, ix12, nrow, 
						wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, uLearnedYN, wIxLemmaList   ] ;
				newTran[ixW2StudyLs]=1; 
				newUp++;
				eleOldTran.innerHTML = newTranslation;  
		
				eleTD_6 =  eleTr2.children[(numCellWrd+1)];          // colonna Lemma 
				
				//console.log("numCellWrd+1=", (numCellWrd+1), " eleTD_6=", eleTD_6.outerHTML)
				
				var eleTD_6div =  eleTD_6.children[0];  
				var ele_divSu = eleTD_6div.children[1]
				//console.log("  ele_divSu=", ele_divSu.outerHTML)
				if (ele_divSu) {
					var ele_tranLemma = ele_divSu.children[3]
					//console.log("  spanTran =", ele_tranLemma.outerHTML)
					if (ele_tranLemma) {
							ele_tranLemma.innerHTML = newTranslation; 						
					}						
				}
			} 		
		} 
		if (newDivTran) {
			//console.log("CXXXXXX  to REMOVE newDivTran=", newDivTran.outerHTML)
			//eleTD_5.children[1].remove(); 	
			newDivTran.remove(); 
		}	
	}	
	//---------------------
	//console.log("9 onclick_saveNewWordTran" , "   newUp=", newUp ); 
	
	if (newUp > 0) {
		//console.log( red(" 2 onclick_saveNewWordTran	call  write_word_dictionary"))		
		write_word_dictionary()
	}
	
} // end of TOGLIonclick_saveNewWordTran	

//------------------------------------
/***
<td class="borderVert"> 
				<div class="divRowText" >
					<div class="suboLine" style="display:none;" id="idc_§1§"  ondblclick="onclickDoubleRowTran(this)">§4txt§</div>
					<div class="tranLine" style="display:none;" id="idt_§1§">§5txt§<br></div>	
					<div id="idw_§1§" class="center" style="width:100%;border:0px solid red;"></div>				
					<div style="display:none;" id="idtts§1§">§ttstxt§</div>
					<div style="display:none;">§ixRow2StudyLs§</div>
				</div>			
				-------------------- to be added BY DOUBLE CLICK ---- 
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
//----------------------------------------
function onclick_saveNewRowTran(this1) {	
	// questo onclick appare nella riga creata da onclickDoubleRowTran(this)
	/*
	questa funzione salva tutte le nuove traduzioni inserite trovate in tutte le TR del TBODY
	*/
	var eleDiv1 = this1.parentElement;  
	var eleDiv2 = eleDiv1.parentElement; 	
	var eleTDs  = eleDiv2.parentElement;
	var eleTR   = eleTDs.parentElement; 
	/*
	var eleTR = this1.parentElement; 
	for(var z1=0; z1 < 10; z1++) {
		if (eleTR.tagName == "TR") { break; } 
		eleTR = eleTR.parentElement; 
	} 
	**/
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
		if (eleTr2.id.indexOf("_m") >0) {continue}
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
		write_row_dictionary(2)
	}
	
} // end of onclick_saveNewWordTran	


//-------------------------
function anyOtherWord() {
	// vedi lineByLine onclickSelectWord
	//--------------------
	let str1 = `
		<div>
				<div style="font-size:0.5em;color: black;text-align:left;margin-top:1.5em;">	
					è possibile ottenere la lista delle frasi che contengono una qualunque parola o combinazione di parole, 
					inserendo le parole richieste nelle 2 liste e poi premendo il pulsante di ricerca.
					<br>Per soddisfare il criterio di scelta una parola delle lista1 deve essere presente insieme ad una parola della lista2  (una delle due lista può essere vuota).
					<br>(es. la prima lista potrebbe contenere il prefisso di un verbi composto tedesco, la seconda le voci di paradigma dello stesso verbo)   
					<table>								
						<tbody><tr><td>1) lista di parole</td><td><textarea id="idwS1_§x1§" cols="50" rows="1" placeholder="ab"></textarea></td></tr>
						<tr><td>2) lista di parole</td><td><textarea id="idwS2_§x1§" cols="50" rows="2" placeholder="gebe gebt gab gegeben"></textarea></td></tr>
					</tbody></table>
				</div>	
				<div style="font-size:0.6em;text-align:center;margin-bottom:1.5em;padding-bottom:2em; border-bottom:1px solid black;">
					<button class="but_word" style="font-size:1.0em;" onclick="onclickSelectWord2('§x1§')">cerca tutte le frasi che contengono le parole delle liste1 e 2</button>	
					<span style="font-size:0.8em;">&nbsp;&nbsp; (massimo </span>
					<input id="idwS3_§x1§"  type="number" min="1" max="400" value="100" style="text-align:right; width:4em; background-color:white;" >
					<span style="font-size:0.8em;"> righe)</span>
					<br><span id="err_§x1§" +="" style="color:red; font-weight:bold;display:none;"></span>
				</div>
		</div> \n`;
	return "\n" + str1.trim();

} // end of anyOtherWord()
	
//-----------------------------------------	
function onclickSelectWord2(id1) {	
	//-------------------------
	//                                          vedi onclick_require_rowListWithThisWord2
	//                                                    go_passToJs_thisWordRowList(aWord, ""+maxNumRow5, "js_go_showWrdRowList");  
	//
	//                   go  bind_go_passToJs_someWordsRowList( aWordList1 string, aWordList2 string,  maxNumRow int, js_function string) 
	//                        nel file bind_go_word_to_row.go
	/**
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
	
	**/
	
	var maxNumRow5 = 100; 
	
	var wordLista1 = "", wordLista2=""; 
	var ele_lista1 , ele_lista2;  	
	
	
	var id_wordLista1 = "idwS1_" + id1; 
	var id_wordLista2 = "idwS2_" + id1; 
	var id_maxNum     = "idwS3_" + id1; 
	if (document.getElementById(id_maxNum)) {
		var numVal=	getInt( document.getElementById(id_maxNum).value);
		if (numVal > 0) maxNumRow5 = numVal; 	
	}
	
	ele_lista1 = document.getElementById(id_wordLista1);  
	ele_lista2 = document.getElementById(id_wordLista2);  
	if (ele_lista1) wordLista1 = ele_lista1.value;
	if (ele_lista2) wordLista2 = ele_lista2.value;	
	
	//console.log("%conclickSelectWord2 ", "color:red;"); console.log("maxNumRow5=", maxNumRow5, " wordLista1=", wordLista1, " wordLista2=", wordLista2); 
	//console.log("%cCERCA  PAROLA "+ wordLista1 + " "+ wordLista2 , "color: green;") 
		
	//get_first_tr_visible();  // memorizza la prima TR visibile delle frasi in cui si trova questa funzione 
	
	//myPage01.style.display = "none"; 
	go_passToJs_someWordsRowList(wordLista1, wordLista2, ""+maxNumRow5, "js_go_showWrdRowList"); 	
	
	
	return 
	var idMsg = "err_" + id1;
	var eleMsg = document.getElementById(idMsg);  
	if (eleMsg)  eleMsg.style.display ="none";	
	
	//myPage01.style.display = "none"; 
	go_passToJs_someWordsRowList(aWord, ""+maxNumRow5, "js_go_showWrdRowList"); 	
	
} // end of onclickSelectWord2

//=========================================================

function get_first_tr_visible() {	
	var divConTab1 = document.getElementById("id_div_tabSub");  // elemento scrollabile che contiene la tabella
	var eleTabBody = document.getElementById("id_tabSub_tbody");
	var eleTrList  = eleTabBody.children;          //righe tabella che possono scomparire/apparire qusndo  il cursore sposta la vista 
	var container  =  document.getElementById("id_div_tabSub");  // elemento scrollabile che contiene la tabella
	
	const containerTop = container.scrollTop;
	const containerBottom = containerTop + container.clientHeight;
	var first_visible_tr_id = -1;
	for (let i = 0; i < eleTrList.length; i++) {
		var ele = eleTrList[i]
		const eleTop = ele.offsetTop;
		const eleBottom = eleTop + ele.clientHeight;
		if (eleTop >= containerTop && eleBottom <= containerBottom) { // The element is fully visible in the container
			 first_visible_tr_id = i;
			 break; 
		}
	} // end for i 
	if (first_visible_tr_id < 0) {
		//console.log("%cnessuna tr è visibile", "color:red;")
	} else {
		//console.log("%cLa prima tr visibile è " + first_visible_tr_id, "color:blue;")
		var eleFromNum=document.getElementById("id_gruppi_iBegNum");  
		eleFromNum.value = parseInt(eleFromNum.value) + first_visible_tr_id;  	               // set on first page
		onchange_rowGroupSelectChange(false,11);
	}
	
	
} // end of get_first_tr_visible	
//-----------------------------------
function back_from_listaRighe(  myPage05, zero,myPage03, myPage01) {
	 
	get_first_tr_visible() 
	 
	onclick_jumpFromTo1_2Page( myPage05, 0,myPage03, myPage01)
	
} // end of back_fromn_listaRighe

//-------------------------
window.onbeforeunload = function(){
	get_first_tr_visible();
	//return 'Are you sure you want to leave?';
};


//----------------------------------------
function onclick_setPrevRowsLearned(this1) {
	get_first_word_tr_visible();
} // end of onclick_setPrevRowsLearned 
//-----------------------------------------
function get_first_word_tr_visible() {	
	var eleTabBody = document.getElementById("idTableWordList_tbody");
	var eleTrList  = eleTabBody.children;          //righe tabella che possono scomparire/apparire qusndo  il cursore sposta la vista 
	var container  =  document.getElementById("id_wordList1");  // elemento scrollabile
	
	const containerTop = container.scrollTop;
	const containerBottom = containerTop + container.clientHeight;
	var first_visible_tr_id = -1;
	for (let i = 0; i < eleTrList.length; i++) {
		var ele = eleTrList[i]
		const eleTop = ele.offsetTop;
		const eleBottom = eleTop + ele.clientHeight;
		if (eleTop >= containerTop && eleBottom <= containerBottom) { // The element is fully visible in the container
			 first_visible_tr_id = i;
			 break; 
		}
	} // end for i 
	if (first_visible_tr_id < 0) {
		//console.log("%cnessuna tr è visibile", "color:red;")
	} else {
		var ele = eleTrList[first_visible_tr_id]
		var eleTD0 = ele.children[0];
		console.log("%cLa prima tr visibile è " + first_visible_tr_id, "color:blue;")
		console.log(eleTD0.children[0].innerHTML, " => ", eleTD0.children[1].innerHTML );  
		/**
		var eleFromNum=document.getElementById("id_gruppi_iBegNum");  
		eleFromNum.value = parseInt(eleFromNum.value) + first_visible_tr_id;  	               // set on first page
		onchange_rowGroupSelectChange(false,11);
		**/
	}
	
	
} // end of get_first_word_tr_visible	
//-----------------------------------	
//--------------------------------------------
function onclick_word_known2(sIxWord, this1) {
	var ixWord = 0, yes_not_len1;
    try {
        ixWord = parseInt(sIxWord);
    } catch (err) {}	
	
	var sw_yesNo = (this1.innerHTML != YES1)
  
	if (sw_yesNo) {
		this1.innerHTML = YES1;		
		yes_not_len1= "y";
		this1.className = "learnedButt";			
	} else {
		this1.innerHTML = NOT_YET1;		
		yes_not_len1= "n";
		this1.className = "notYetLearnedButt";			
	}	
	numWordsKnownChanged++; 
	document.getElementById("id_buttLearnNumW").innerHTML = numWordsKnownChanged;
	
	go_passToJs_word_known2(""+ixWord, yes_not_len1, "js_go_word_known2"); // ask 'go' to update yes/no word known ctr  
	
} // end of onclick_word_know_yes		

//----------------------------------------
function js_go_word_known2(str1) {
	//console.log("js_go_word=", str1);   
	
	if (numWordsKnownChanged == 1) {		
		document.getElementById("id_buttLearnDiv").style.display = "block";
	} 	
	if (numWordsKnownChanged > MAX_NUM_WORD_LEARN) {		
		onclick_write_words_to_learn();
	} 
	
}// end of js_go_word_known 	

//-------------------------------------------------------------
function js_go_file_words_to_learn_written( str1 ) {
	//document.getElementById("id_w_to_learn_written").innerHTML = str1 ;	
	numWordsKnownChanged = 0; 
	document.getElementById("id_buttLearnNumW").innerHTML = numWordsKnownChanged;
	document.getElementById("id_buttLearnDiv").style.display = "none";
}
//-----------------------------------------------------------------

function onclick_hideShowWordTran(this1) {	

	var swEle = this1.previousElementSibling;
	
	//console.log("%conclick_hideShowWordTran", "color:red;"); console.log(" swEle=", swEle)

	// nasconde o mostra la traduzione di tutte le parole 
	var eleBody = document.getElementById("idTableWordList_tbody")
	var rows1 = eleBody.rows; 	
	//--------
	var lemmaCell, eleTran, visib;		
	//-------------
	if (swEle.innerHTML == "y") {
		swEle.innerHTML = "n"; 
		//visib = "visible";		
		visib = "inline-block";
	} else {
		swEle.innerHTML = "y";
		//visib = "hidden";
		visib = "none"; 
	}	
	//console.log( "            swEle.innerHTML=", swEle.innerHTML,   " visib=", visib, " rows1.length=", rows1.length)
	//----------------------
	for (var g=0; g < rows1.length; g++ ) {	
		try {
			lemmaCell = rows1[g].cells[NUM_CELL_LEMMA]
			eleTran = lemmaCell.children[0].children[1].children[1]; 
			if (eleTran) eleTran.style.display = visib;
			//console.log(" lemmaCell=", lemmaCell, "  VISIBIL=",eleTran.style.visibility, " eleTran=", eleTran.outerHTML )
		} catch(e1) {
			continue
		}
	}
	
} // end of onclick_hideShowWordTran 
//-------------------------------------------------------

var hideShowTranExample_H1 = "mostra traduzione ed esempi"; 
var hideShowTranExample_H2 = "nascondi traduzione ed esempi"; 
//----------------
function onclick_hideShowTranExample(this1) {
	//console.log("onclick_hideShowTranExample")
	var display1;
	var lemmaCell, eleTran;
	var buttHeader = this1.innerHTML; 
	if (buttHeader == hideShowTranExample_H1) {
		display1="block";
		this1.innerHTML = hideShowTranExample_H2
	} else {
		display1="none";
		this1.innerHTML = hideShowTranExample_H1
	}	
	var eleBody = document.getElementById("idTableWordList_tbody")
	var rows1 = eleBody.rows; 
	//--------------------------
	for (var g=0; g < rows1.length; g++ ) {	
		try {
			lemmaCell = rows1[g].cells[NUM_CELL_LEMMA]
			eleTran = lemmaCell.children[0].children[2]; 
			if (eleTran) eleTran.style.display = display1;
		} catch(e1) {
			console.log("%cerrore " + e1, "color:red;")
			continue
		}
	}
	
} // end of onclick_hideShowTranExample
//------------------------
function TOGLIonclick_vertResizeLemma(this1) {	
	var eleBody = document.getElementById("idTableWordList_tbody")
	var rows1 = eleBody.rows; 		
	var class01 = "c_size_1_line";
	var classNN = "c_size_nn_line";
	if (rows1.length < 1) return
	var lemmaCell = rows1[0].cells[NUM_CELL_LEMMA].children[0]
	var swClass = ( lemmaCell.classList.contains( class01 )	)
	//----------------------
	if (swClass) {		
		this1.innerHTML = "mostra solo 2 righe"
		for (var g=0; g < rows1.length; g++ ) {	
			lemmaCell = rows1[g].cells[NUM_CELL_LEMMA].children[0]
			lemmaCell.classList.remove( class01 );
			lemmaCell.classList.add(    classNN );		
		} // end for g
	} else {		
		this1.innerHTML = "mostra tutte le righe"
		for (var g=0; g < rows1.length; g++ ) {	
			lemmaCell = rows1[g].cells[NUM_CELL_LEMMA].children[0]
			lemmaCell.classList.remove( classNN );
			lemmaCell.classList.add(    class01 );				
		}
	} 	
} // end of TOGLIonclick_vertResizeLemma 

//-----------------------------------------------------------------

function onclick_listNoParadigmaLemma(this1) {	

	var downfilename= "parole_senza_paradigma.txt" 
	var outText = "Lista Parole senza paradigma " + "\n\n" ;
	
	
	var eleBody = document.getElementById("idTableWordList_tbody")
	var rows1 = eleBody.rows; 	
	//--------
	var lemmaCell, eleTran, elePara, eleDivSup;	
	//----------------------
	for (var g=0; g < rows1.length; g++ ) {	
		try {
			lemmaCell = rows1[g].cells[NUM_CELL_LEMMA]
			eleDivSup = lemmaCell.children[0].children[1];
			elePara = eleDivSup.children[1]; 
			
			if (elePara.innerHTML == "") {
				outText += eleDivSup.children[0].innerHTML.replaceAll("<b>","").replaceAll("</b>","") + "  | \n";;				
			}
		} catch(e1) {
			console.log(e1)
			continue
		}
	} // end for g
	
	download(downfilename, outText );
		
} // end of onclick_listNoParadigmaLemma 

//-----------------------------------------------------------------

function onclick_listNoTranLemma(this1,swNoTran) {	
	
	var downfilename= "Lemma_senza_traduzione.txt" 
	var outText = "Lista Lemma senza Traduzione"+ "\n" ;
	if (swNoTran == false) {
		downfilename= "lista_tutti_lemma.txt" 
	    outText = "Lista di tutti i lemma"+ "\n" ;
	} 
	
	var eleBody = document.getElementById("idTableWordList_tbody")
	var rows1 = eleBody.rows; 	
	//--------
	var lemmaCell, eleTran, elePara, eleDivSup;	
	//----------------------
	var noTranL = []
	var tranS;
	for (var g=0; g < rows1.length; g++ ) {	
		try {
			lemmaCell = rows1[g].cells[NUM_CELL_LEMMA]
			eleDivSup = lemmaCell.children[0].children[1];
			eleTran = eleDivSup.children[3]; 
			if (eleTran) tranS = eleTran.innerHTML;
			else tranS = ""
			if (swNoTran) {
				if (tranS != "") continue;
			}
			noTranL.push(  eleDivSup.children[0].innerHTML.replaceAll("<b>","").replaceAll("</b>","") + "|" + tranS );					
			
		} catch(e1) {
			console.log(e1)
			continue
		}
	} // end for g
	//-----------------
	noTranL.sort();
	//------------------	
	var preW="", wo="";
	var nn=0;
	for (g=0; g < noTranL.length; g++ ) {
		wo = noTranL[g]
		if (wo == preW) continue;
		preW = wo; 
		nn++;
		outText += "\n|" + nn + "|" + wo;		
	}	
	
	download(downfilename, outText + "\n" );
	
} // end of onclick_listNoTranLemma 
//----------------------------------------------


function onclick_listWords(this1) {	
	
	var downfilename= "lista_parole.txt" 
	var outText = "Lista Parole"+ "\n" ;	
	var eleBody = document.getElementById("idTableWordList_tbody")
	var rows1 = eleBody.rows; 	
	//--------
	var wordCell, lemmaCell, eleTran, elePara, eleDivSup, eleWrd;	
	var lemmaT, wordT;
	//----------------------
	var noTranL = []
	var tranS;
	for (var g=0; g < rows1.length; g++ ) {	
		try {
			wordCell = rows1[g].cells[4];
			eleWrd = wordCell.children[0].children[0].children[0].children[0]; 
			if (eleWrd == undefined) continue; 
			
			lemmaCell = rows1[g].cells[NUM_CELL_LEMMA]
			eleDivSup = lemmaCell.children[0].children[1];
			eleTran = eleDivSup.children[3]; 
			if (eleTran) tranS = eleTran.innerHTML;
			else tranS = ""	;	
			wordT = eleWrd.innerHTML.replaceAll("<b>","").replaceAll("</b>","") ;
			lemmaT = 	eleDivSup.children[0].innerHTML.replaceAll("<b>","").replaceAll("</b>","");
			noTranL.push( lemmaT+ " §§" +wordT + "|" + lemmaT + "|" + tranS );				
		} catch(e1) {
			console.log(e1)
			continue
		}
	} // end for g
	//-----------------
	noTranL.sort();
	//------------------	
	var preW="", wo="";
	var nn=0;
	for (g=0; g < noTranL.length; g++ ) {
		wo = noTranL[g].split("§§")[1];
		if (wo == preW) continue;
		preW = wo; 
		nn++;
		outText += "\n|" + nn + "|" + wo;		
	}	
	
	download(downfilename, outText + "\n" );
	
} // end of onclick_listWords 
//------------------------------------------	
function download(filename, text) {

    var element = document.createElement('a');

    element.setAttribute('href', 'data:text/plain;charset=utf-8,' +
        encodeURIComponent(text));
   
    element.setAttribute('download', filename);

    element.style.display = "none";
    document.body.appendChild(element);
    element.click();
    document.body.removeChild(element);
}

//--------------------------------------------------
function onclick_split_newText() {
		var eleTxt = document.getElementById("id_newText");
		var str1 = (""+eleTxt.value).trim();
		if (str1.length == 0) return;
		/**
		console.log("punto ==> ",  str1.replaceAll(".",".\n") )
		console.log("punto? ==> ",  str1.replaceAll("?","?\n") )
			console.log("punto! ==> ",  str1.replaceAll("!","!\n") )		
		**/
		str1 = str1.replaceAll(".",".\n").replaceAll("?","?\n").replaceAll("!","!\n").replaceAll(";",";\n").replaceAll("\n\n","\n");
		
		//console.log("tutti ==> ",  str1)
			
		eleTxt.value = str1;			
} // end of onclick_split_newText		  
/**
1_0|O|file: prova.csv
1_0|T|file: prova.csv
1_1|O|Zu meiner Familie gehören vier Personen.
1_1|T|Ci sono quattro persone nella mia famiglia.
**/
//-------------------
function trimLeftZero( num1 ) {
	num1 = ("" + num1).trim();
	for(var g=0; g < 5; g++) {   
		if (num1.length < 2) { break;}
		if (num1.substr(0,1) == "0") num1 = num1.substring(1); 
	}
	return num1;
}	
//--------------
function onclick_get_newText() { 			
	var eleTxt = document.getElementById("id_newText");
	var str1 = (""+eleTxt.value).trim();
	if (str1.length == 0) return;			
	var ele_id1 = document.getElementById("id_newTxt_id")
	var ele_title1 = document.getElementById("id_newTxt_title"); 
	var id1 = ele_id1.value.trim();
	var title1 = ele_title1.value;	
	
	var msg1="";
	if (id1=="") {msg1=" manca identificativo"; }
	if (title1=="") { msg1+=" manca titolo del testo"; }
	if (id1 != "") {
		id1 = trimLeftZero( id1 );
		if (listaGruppiTesto.indexOf( "," + id1 + ",") >=0 ) {
			msg1 += "gruppo " + id1 + " non può essere utilizzato, è già esistente";   
		}	
	} 	
	if (msg1 != "") msg1="<br>errore: " + msg1;	
		
	document.getElementById("id_newTxtMsg1").innerHTML = msg1;
	if (msg1 != "") return;
	
	var righe=str1.split("\n");
	var riga;
	var outText = id1+"_0|O|file: " + title1 ;
	outText += "\n" + id1+"_0|T|file: " + title1 ;
	for (var f=0; f < righe.length; f++) {
		riga = "\n" + id1+ "_" + (f+1) + "|O|" + righe[f].trim();
		outText += riga;
	}	
	eleTxt.value = outText;
	console.log(" onclick_get_newText ", " --> go_write_new_row_dictionary ", outText)
	go_write_new_row_dictionary( outText, "js_go_new_row_written")
	
	
}// end of onclick_get_newText		
//---
/**
function go_write_new_row_dictionary( str1 ) {
	console.log("passato a go ", str1 )
	js_go_new_row_written("fine prova" )
}  
***/
//-------------------------------------
function js_go_new_row_written(str1 ) {
	var eleTxt = document.getElementById("id_newText").value = ""
	var ele_id1 = document.getElementById("id_newTxt_id")
	var ele_title1 = document.getElementById("id_newTxt_title"); 
	var str1 = "il nuovo testo con identificativo " + ele_id1.value + " è stato accettato";
	//var msg1 = "chiudi e riesegui l'applicazione"
	document.getElementById("id_newTxtMsg0").innerHTML = "<br>" + str1;
	document.getElementById("id_newTxtMsg1").innerHTML = "";
	ele_id1.value = "";  
	ele_title1.value= ""; 
	
}  // end of js_go_new_row_written
//-----------------------------------------------------
		 // display traduzione righe toccando col mouse la riga originale 
		 /*
				<div class="suboLine" style="display: none;" id="idc_1" ondblclick="onclickDoubleRowTran(this)"					
						 onmouseover='onmouseOverRow(this)' onmouseout='onmouseOutRow(this)'			
				>Zu meiner Familie gehören vier Personen.</div>
				<div class="tranLine" style="display:none;" id="idt_1">Ci sono quattro persone nella mia famiglia.<br></div>	
		 */
		
//------------------------------------	 
let sw_mouseoverActiv = false;
//----------------------------------   
function onclick_moveOverActivate(this1) {
	if (sw_mouseoverActiv) {
		sw_mouseoverActiv = false;
		this1.innerHTML = "attivare la traduzione al passaggio del mouse"
	} else {
		sw_mouseoverActiv = true;
		this1.innerHTML = "disattivare la traduzione al passaggio del mouse"
	}
	
} // end of onclick_moveOverActivate
//---------------------------------------------
function onmouseOverRow(this1) {  // toccando la riga originale, rende visibile la traduzione  
	if (sw_mouseoverActiv == false) return;
	var id1 = this1.id;
	if (id1 == undefined) return 
	var idNum = id1.replaceAll("idc_","idt_"); 
	var eleTran = document.getElementById(idNum); 
	if (eleTran == undefined) return;
	eleTran.style.display = "block";			 
} // end of onmouseOverRow
//-----------------------------------------------------------	
function onmouseOutRow(this1) {   // allontanando il mouse dalla riga originale, nasconde la traduzione (a meno che non sia attivo il tasto show T.  
	if (sw_mouseoverActiv == false) return;	
	var id1 = this1.id
	if (id1 == undefined) return 
	var idNum = id1.replaceAll("idc_","idt_"); 
	var eleTran = document.getElementById(idNum); 
	if (eleTran == undefined) return;
	// sarebbe naturale rimettere display none, ma fintanto che esistono i pulsanti di show/hide transl. forzo lo stato dettato da questi 
	var idNumButtT = id1.replaceAll("idc_","idbT_");
	if (idNumButtT == undefined) return;
	var eleTranButt = document.getElementById(idNumButtT); 
	var eleTbutCh   = eleTranButt.children[0]; 
	if (eleTbutCh == undefined) return;
	eleTran.style.display = eleTbutCh.style.display;			 
} // end of onmouseOutRow
//---------------------------------------------------------------------		

function onmouseOverRow2( num1 ) {  // toccando la riga originale, rende visibile la traduzione  
	
	var idNum = "idt_" + num1; 
	var eleTran = document.getElementById(idNum); 
	if (eleTran == undefined) return;
	eleTran.style.display = "block";	
	
} // end of onmouseOverRow
//-----------------------------------------------------------	
function onmouseOutRow2( num1 ) {   // allontanando il mouse dalla riga originale, nasconde la traduzione (a meno che non sia attivo il tasto show T.  
	var idNum = "idt_" + num1; 
	var eleTran = document.getElementById(idNum); 
	if (eleTran == undefined) return;
	eleTran.style.display = "none";	 
	
} // end of onmouseOutRow
//---------------------------------------------------------------------		

//----------------------------------------------------------
function onclick_changeTraduzione(this1 ) {	
	
	var eleTD0  = this1.parentElement;
	
	var eleLemmaTD = eleTD0.nextElementSibling;
	
	/***
		// td lemma     tr.cells[NUM_CELL_LEMMA] 	
		<td style="text-align:center;" class="fixWordTdWidthL borderVert_L borderVert_R">
			<div class="c_size_1_line">					
				<div style="display:none;">zwei</div>
				<div class="hpad top left1 " style="width:100%;">										
					<span style="display:none;" class="c_lemma" onmouseover="mouseOverWord(this,3)" onmouseout="mouseOutWord(this,3)"><b>zwei</b></span>
					<span class="c_paradigma" onmouseover="mouseOverWord(this,3)" onmouseout="mouseOutWord(this,3)">zwei</span>						
					<br class="c_paradigma">	
					<div class="c_wordTran" style="visibility: hidden;">due</div>						
					<div class="c_example"></div>
				</div> 				
			</div>
			----
			qui la parte per variare la traduzione 	
			----
		</td>
	
	***/
		
	var eleWordTD  = eleTD0.previousElementSibling;
		
	var eleTD    =  eleWordTD; 
		
	var eleLemDiv1   = eleLemmaTD.children[0]
	var eleLemmaSpan = eleLemDiv1.children[0]
	if (eleLemmaSpan.innerHTML == "") return; 	
		
	if (eleLemmaTD.children.length >= 2) { return; } 
	
	var ele_oldTran = eleLemmaTD.children[0].children[1].children[1];
	const newDiv = document.createElement("div");
	newDiv.style.textAlign = "left";
	
	eleLemmaTD.appendChild(newDiv);
	var newInn=""
	newInn += '<div  style="font-size:0.6em;width:100%;">add/modify translation</span></div>' + '\n';
	newInn += '<div  class="c_wordTran" ' +
		'style="background-color:lightgrey; color:black; font-weight:bold;border:2px solid black;min-width:100%;text-align:left;" ' +
		'contentEditable=true>' + 
		ele_oldTran.innerHTML + '</div>' + '\n'		
	newInn += '<div  style="width:100%;"><button onclick="onclick_saveNewWordTran(this)">Salva tutte le nuove traduzioni</button>' +
				'</div>\n'; 
	
	eleLemmaTD.children[1].innerHTML = newInn;    
	
} // end of onclick_changeTraduzione

//----------------------------------------
function onclick_saveNewWordTran(this1) {
	
	/******
	
	<td style="text-align:center;" class="fixWordTdWidthL borderVert_L borderVert_R">
		<div class="c_size_1_line">
			<div style="display:none;">zwei</div>
			<div class="hpad top left1 " style="width:100%;">										
				<span style="display:none;" class="c_lemma" onmouseover="mouseOverWord(this,3)" onmouseout="mouseOutWord(this,3)"><b>zwei</b></span>
				<span class="c_paradigma" onmouseover="mouseOverWord(this,3)" onmouseout="mouseOutWord(this,3)">zwei</span>						
				<br class="c_paradigma">	
				<div class="c_wordTran" style="visibility: hidden;">due</div>						
				<div class="c_example"></div>
			</div> 
		</div>
					
		<div style="text-align: left;">
			<div style="font-size:0.6em;width:100%;">add/modify translation</div>
			<div class="c_wordTran" style="background-color:lightgrey; color:black; 
					font-weight:bold;border:2px solid black;min-width:100%;text-align:left;" 
					contenteditable="true">
					due
			</div>
			<div style="width:100%;">
					<button onclick="onclick_saveNewWordTran(this)">Salva tutte le nuove traduzioni</button>
			</div>
		</div>
	</td>	
	
	*********/
	
	
	var wLemmaList, wTranList, wLevelList, wParaList, wExampleList,  wIxLemmaList;	
	let wordx, ix12, nrow, totExtrRow2,  wLemma1, wordTran, uLearnedYN ;
	let word1, ixW2StudyLs, ix1, ixLemma; 	
	var oldDivTran, eleOldTran, oldTranslation; 
	var newDivTran, eleNewTran, newTranslation; 
	var swChg=false;
	var eleTD_0, eleTD_5; var eleTD0_val;
	var eleTD_6, ele_details, ele_details, ele_summ, ele_tranLemma; 
	
	var newUp = 0; 
	
	var div_onsave00 = this1.parentElement; 	
	var divAdded00   = div_onsave00.parentElement;  
	var eleTdLemma00 = divAdded00.parentElement; 
	if (eleTdLemma00.tagName != "TD") { console.log("%cERRORE eleTdLemma.tagName not equal TD =" + +eleTdLemma00.tagName );return; } 	
	var eleTR = eleTdLemma00.parentElement; 	
	if (eleTR.tagName != "TR") { return; }
	var elePareTr = eleTR.parentElement; // tbody
	if (elePareTr == null) return; 
	
	//----------------------
	// cerca tutti i lemma con traduzione pendente  
	
	var tranToUpdList = [];
	var bodyChild = elePareTr.children
	for (var p1=0; p1 < bodyChild.length; p1++) {
		var eleTr0 = bodyChild[p1];
		var oneTdLemma = eleTr0.children[NUM_CELL_LEMMA]; 		
		if (oneTdLemma.children.length < 2) { continue; }	
		var oneNewLemTran = oneTdLemma.children[1]; 
		var oneNewLemTranInner = oneNewLemTran.children[1].innerHTML;
		if (oneNewLemTranInner == "") { continue; }	
		tranToUpdList.push( [ oneTdLemma.children[0].children[0].innerHTML, oneNewLemTranInner, oneNewLemTran ] ); 		
	} // end for p1	
	//-------------------
	for(var t2=0; t2 < tranToUpdList.length; t2++) {
		var toUpdLemma   = tranToUpdList[t2][0];	
		var toUpdTran    = tranToUpdList[t2][1];	
		var addedLivTran = tranToUpdList[t2][2];
		for (var p2=0; p2 < bodyChild.length; p2++) {
			updateOneTran(  bodyChild[p2], toUpdLemma, toUpdTran); 
		}	
		addedLivTran.remove(); 			
	}  // end for t2 
	//---------------------
	if (newUp > 0) {
		//console.log( " onclick_saveNewWordTran	call  write_word_dictionary",    "  ", newUp, " variazioni")		
		write_word_dictionary();
	}
	return
	
	//------------------------------
	function updateOneTran( eleTr2, targetLemmaName, newLemTran ) {
		var eleTdLemma = eleTr2.children[ NUM_CELL_LEMMA] ;
		var ele_target_lemmaName = eleTdLemma.children[0].children[0].innerHTML; 			
		if (ele_target_lemmaName != targetLemmaName) { 
			return;
		}	
		//-----
		eleTD_0 = eleTr2.children[0];
		eleTD0_val = eleTD_0.children[1]; 
		word1       = eleTD0_val.children[0].innerHTML; 
		ixW2StudyLs = eleTD0_val.children[1].innerHTML; // indice wordToStudy_List
		ix1         = eleTD0_val.children[2].innerHTML; // indice word (in go) 
		ixLemma     = eleTD0_val.children[3].innerHTML; // indice lemma x lemma e tran list 

		var ele_target_lemTran = eleTdLemma.children[0].children[1].children[1]; 		
		var ele_target_tran2   = eleTdLemma.children[0].children[2].children[0].children[0]; 
		//console.log("ele_target_lemTran=", ele_target_lemTran.innerHTML)
		//console.log("ele_target_tran2  =", ele_target_tran2.innerHTML)
		
		if (ele_target_lemTran.innerHTML ==  newLemTran) { 
			console.log("    traduzione eguale a prima, ignoro")
			return;
		}
		ele_target_lemTran.innerHTML = newLemTran;
		ele_target_tran2.innerHTML = newLemTran;
		//------		
		[wordx, ix12,  nrow, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2,uLearnedYN, wIxLemmaList  ] = wordToStudy_list[ixW2StudyLs]; 
		wTranList = newLemTran;
		wordToStudy_list[ixW2StudyLs] = [wordx, ix12, nrow, 
						wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, uLearnedYN, wIxLemmaList   ] ;
		newTran[ixW2StudyLs]=1; 	 
		newUp++;
		
	} //end of updateOneTran 
	//-------------------------------------
	
} // end of onclick_saveNewWordTran	

//------------------------------------

function vertResizeLemma( eleTdLemma ) {	
	var class01 = "c_size_1_line";
	var classNN = "c_size_nn_line";
	var divToResize = eleLemma.children[0] ; 
	var swClass = ( divToResize.classList.contains( class01 )	)
	//----------------------
	if (swClass) {	
		divToResize.classList.remove( class01 );			
		divToResize.classList.add(    classNN );		
	} else {	
		divToResize.classList.remove( classNN );			
		divToResize.classList.add(    class01 );	
	} 	
} // end of onclick_vertResizeLemma 

//-------------------------------------

function onclick_showHideLemma(this1) {
	let ele2 = this1.nextElementSibling;
	if (ele2.style.display == "none") ele2.style.display = "block";
	else ele2.style.display = "none";
} // end of onclick_showHideLemma 

//------------------------------------

function getWord_tr2(z3, parola1, paro_tts, trad1, all_lemma_for_thisWord, paro_nFrasi, maxNumRow) {

	var wordTR = prototype_word_tr_tts.replaceAll("§1§", z3).replaceAll("§4txt§", parola1).replaceAll("§4maxNumRow§", "" + maxNumRow).
	replaceAll("§ttsWtxt§", paro_tts).replaceAll("§8numfrasi§", paro_nFrasi).replaceAll("§6tran§", trad1).
	replaceAll("§6alllemmaWord§", all_lemma_for_thisWord);
	return wordTR;

} // end of getWord_tr 

//-----------------------------------

function mouseOverWord3(this1) {
    if (this1.children.length > 1) {
        this1.children[1].style.display = "inline-block";
    }
}
//---------------------
function mouseOutWord3(this1) {
    if (this1.children.length > 1) {
        this1.children[1].style.display = "none";
    }
}
//-------------------------------