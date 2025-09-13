"use strict";
/*  
Words By Frequence: A tool to practice language comprehension
Antonio Cigna 2023
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint esversion:  8 */
/* jshint strict: true  */
/* jshint undef: true, unused: true */
/* jshint varstmt: true */
/* jshint asi: true     */
//-----------------------------------------------
const YES1 = "Yes";
const NOT_YET1 = "Not Yet";
const MAX_NUM_WORD_LEARN = 10;
let numWordsKnownChanged = 0;
const NUM_CELL_LEMMA = 6;  
let sw_rowListFrom_onclick = false
let listaGruppiTesto=""; 
//------------------------------------------
let numeroWord_TR=0;
const wSep = "§";
const endOfLine = ";.\n"; 
//let apiceInverso = `40`
let is_selected_row_only = false; 
let index_onlySelRowsWanted = set_ix_selectedRow(); //0

let swBegin=false;
let ele_allowWordTranslation= getById("id_allowWordTranslation"); 
let ele_wordsToTranslate 	= getById("id_words_to_translate");
let ele_wordsTranslated  	= getById("id_words_translated"  );

let ele_wordList = getById("id_wordList1");
//----------------------------------------
//  dati che rappresentano il gruppo di testo nel DB  
let x_rG_group           // numero identificativo del gruppo 
let x_rG_firstIxRowOfGr  // indice dell'inizio del gruppo    (se il gruppo precedente termina all'indice 250, questa variabile contiene 251)   
let x_rG_lastIxRowOfGr   // indice della fine del gruppo 
let x_rowGrIndex
let x_id_gruppi_iBegNum_html_rowGroup_beginNum
let x_rG_ixSelGrOption

//----------------------------------------------
let SAVE_fromIx_row  ;   
let	SAVE_toIx_row    ; 	
//---------------
let id_gruppi_selected_groupIndex = 0 ; // from onchange_rowGroupSelectChange ( select nella lista dei gruppi, attenzione il gruppo "1" ha indice 0 )
let id_gruppi_SelectedBegin       = 0 ; // from onchange_rowGroupNumBegChange ( scelta della riga di inizio del gruppo ( es. gruppo 3, inizio 0, num.righe 150 )     
let id_gruppi_SelectedNumRow      = 0 ; // from onchange_rowGroupNumBegChange ( scelta del numero righe - default = 1+indice finale - indice iniziale o inizio gruppo se inizio > 0 ) 
//-----------------------------------------------------------
let html_sel_extrRow       = "";   // from onchange_mostFreqWordList_extrRow
let last_id_gruppi_sel_html_rowGroup_index_gr 		= ""; 
let last_id_gruppi_iBegNum_html_rowGroup_beginNum	= "";
let last_html_rowGroup_numRows 		= "";
let last_sel_extrRow_freqWord_list	= ""; 

let last_ele_analWords_id 		= "";
let	last_ele_analWords_tr 		; 
let	last_ele_analWords_height 	;  


let sw_somethingChanged	= false;    // resetted  only by onclick_mostFreqWordList_require   

onchange_mostFreqWordList_extrRow(); 
//onchange_rowGroupSelectChange(false,10); 

//let sw_rowGroupSelectChange       = false; 	
//let sw_rowGroupSelectChange_group = false; 	

//let ele_word     = getById("id_word"      );
//let ele_wRowList = getById("id_wRowList1");
//let ele_wordLisH = getById("id_wordListH");

let myPage01 = getById("id_myPage01");
let myPage02 = getById("id_myPage02");
let myPage03 = getById("id_myPage03");
let myPage04 = getById("id_myPage04");
let myPage05 = getById("id_myPage05");
let maxNumRow = 99999999999; // 100;
let wordToStudy_list = [];
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
let ele_toTranslate_textarea 	= getById("txt_pagOrig");  
let ele_translated_textarea  	= getById("txt_pagTrad"  );
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
//let ele_where = getById("id_where");
let ele_bar   = getById("id_progrBar");
let ele_bar2  = getById("id_progrBar2");
let ele_bar3  = getById("id_progrBar3");
//---------------------------------
const numButton1listW = 1; // onclick most frequent word list  
const numButton2prefW = 2; // onclick BetweenWordList or prefix wordlist   
const numButton3prefL = 3; // onclick Lemma word list   
const numButton4rowL  = 5; // onclick Lemma list   
const numButton5      = 0; // ?   
//----------------------------------
//  bottoni di sinistra
const xOnclick1lisR = "1L" // Lista le Righe scelte                                      onclick_require_rowList1(1)
const xOnclick2lisW = "2L" // Lista Le Parole delle righe scelte                         onclick_require_rowList1(2)
const xOnclick3lisL = "3L" // Lista Soltanto le parole delle righe scelte da imparare    onclick_require_rowList1(3)	
const xOnclick4lisP = "4L" // Lista le Righe scelte per priorità                         onclick_require_rowList1(4)
const tagLeftOnclickArr = ["",xOnclick1lisR,xOnclick2lisW, xOnclick3lisL, xOnclick4lisP ]
//-------------------------------------------------
// bottoni di destra
const xOnclick1lisAW = "1R" // Lista le Parole di tutti i testi             onclick_mostFreqWordList_require('anyRow','allWords')" 
const xOnclick2lisALr= "2R" // Lista le Parole da imparare                  onclick_mostFreqWordList_require('anyRow','toBeLearned')
const xOnclick3lisSW = "3R" // Lista le Parole richieste (pref suffix)      onclick_require_betweenWordList()
const xOnclick4lisSLm= "4R" // Lista i lemma richiesti   (pref suffix)      onclick_require_betweenLemmaList() 
const tagRightOnclickArr = ["",xOnclick1lisAW,xOnclick2lisALr, xOnclick3lisSW, xOnclick4lisSLm ]
//--------------------------------------
// bottoni sulla parte destra della pagina delle parole
const xOnclick1ixWord = "RW1" //  onclick_rowsByIxWord(           \'§one-ix1§\'     , 'RW1')
const xOnclick2ixLemma= "RW2" //  onclick_rowsByIxLemma(          \'§one-ixLemma§\' , 'RW2')
const xOnclick3LemWord= "RW3" //  onclick_require_lemmaWordList2  \'§one-f_lemma§\' , 'RW3')


//------------------
function black(   str1 ) { return "\u001b[30m" + str1 }
function red(     str1 ) { return '\u001b[31m' + str1 }
function green(   str1 ) { return '\u001b[32m' + str1 }
function yellow(  str1 ) { return "\u001b[33m" + str1 }
function blue(    str1 ) { return "\u001b[34m" + str1 }
function magenta( str1 ) { return "\u001b[35m" + str1 }
function cyan(    str1 ) { return "\u001b[36m" + str1 }
function white(   str1 ) { return "\u001b[37m" + str1 }
let js_parm = "";
let js_caller = "";
let js_func = ""; 
//------------------------------------------------
function js_call_go() {  // called by  html page body onload
	let msg1 = "html loaded"; 
	console.log("html function js_call_go: ", msg1, "\n") 
	
	go_passToJs_html_is_ready(msg1,  "", js_parm, js_caller) ;  //  "js_go_go_is_ready"); 
}
//-----------------------------------------
function getCaller(num) {	
	let stack1="";
	try {
		stack1 = (new Error()).stack?.split("\n")[num].toString() ;
		let s1 = stack1.lastIndexOf("/") 
		if (s1 > 0) { stack1 = stack1.substring(s1+1) }
	} catch(e1) {
		stack1 = ( new Error().stack.toString() )		
	}	
	return stack1	
}
//------------------------
function getById_inner( uno ,swTest=false) {
	let ele1 = getById( uno ,swTest) 
	if (ele1) {
		try{ 
			if (ele1.innerHTML) {
				return [ele1,true,true]
			} else {
				return [ele1,true,false]
			}	
		} catch(e1) {
			return [ele1,false,false]
		}		
	} else {
		return [ele1,false,false]
	}	
}
//------------------------
function getById_children( uno ,swTest=false) {
	let ele1 = getById( uno ,swTest) 
	if (ele1) {
		try{ 
			if (ele1.children) {
				return [ele1,true,true, ele1.children.length]
			} else {
				return [ele1,true,false,0]
			}	
		} catch(e1) {
			return [ele1,false,false,0]
		}		
	} else {
		return [ele1,false,false,0]
	}	
}
//----------------
function getById( uno ,swTest=false) {
	try {
		let ele1 = document.getElementById( uno );
		if (ele1) {
			return ele1
		} else {
			if (swTest) return null;
			document.getElementById("id_error00").innerHTML = 'errore non trovato id="' + uno + '" in ' + getCaller(3).replace(")"," ") ;			
			logColor("%%red", 'ERRORE id="', "%%blue;font-weight:bold;",  uno, "%%black", '" in ', "%%blue;font-weight:bold;",  getCaller(3).replace(")"," ") , "%%red", ' non esiste'  );			
		}	
	} catch(e1) {
		logColor("%%red", e1);
	}	
}
/**
let uno="uno"; let due="due"
logColor("%%red", 'ERRORE id="', "%%blue",  uno, due, "%%black", '" in ', "%%violet",  "tre" , "%%red", ' non esiste'  );
logColor("%%red", 'ERRORE id="', "%%blue;font-size:2.0em;",  uno, due, "%%black", '" in ', "%%violet",  "tre" , "%%red", ' non esiste'  );
**/
//------------------------------------
function logColor() {    //  logColor( "%%red", uno, due, tre, "%%green",quattro)

	try {	
		let numAr = arguments.length ; 
		let colori = [""];
		let strList = [];
		let arg1, argS;
		let jL = -1;
		let where = ""
		let i0=0; 
		//-----------
		if (arguments.length > 1) {
			arg1 = ""+ arguments[0];
			argS = arg1.trim();
			if (argS.substr(0,2) != "%%") { 
				where = argS
				i0=1;
			}
		}	
		//-----------------		
		for (let i = i0; i < arguments.length; i++) {
			arg1 = ""+ arguments[i];
			argS = arg1.trim();
			if (argS.substr(0,2) == "%%") {
				colori.push( "color:" +argS.substring(2).trim() +";" );
				strList.push("")
				jL++
			}	
			else {
				strList[jL] += " " + arg1; 
			}
		}	
		//-------------
		if (where != "") {
			let lineStm = getCaller(3);
			colori.push( "color:" +"black;font-size:0.8em;font-style:oblique;" );
			strList.push("")
			jL++
			strList[jL] += "     (in function " + where + " file " + lineStm ;
		}
		//-------------------
		let uno=""
		for(let i2=0; i2 < strList.length; i2++) {
			uno += "%c" + strList[i2].substring(1);		
		}  		
		colori[0] = uno;
		console.log(...colori) ;	
	} catch(e1) {
	  console.log("%cerrore in logColor(","color:red;",  arguments , ")" );
	  console.log("%c" + e1, "color:red;")	  
	}	

} // end of logColor
//------------------------

function set_ix_selectedRow() {
	let x2 = getById("id_sel_2_extrRow");
    for (let i=0; i < x2.children.length; i++) {
		if (x2.children[i].id=="extrRow") {
			return i; 
		}
	}
	return 0; 
} // 

//---------------------------

function js_go_updateStatistics( data, js_parm, jsFunc, goFunc) {
	
	function formatRight(num1, numDigit) { 
		let numS = "" + num1; 
		let nLen = numS.length
		if (nLen >= numDigit ) { return numS;} 
		return ("                ".substr(0, numDigit - nLen)) + numS; 
	}	
	
	let statRow = data.split("<br>");   
	//console.log("js_go_updateStatistics()  statRow=", statRow.join("<br>")) 
	
	let st1, field;
	let result="<table> \n";
	
	for (let z1=1; z1 < statRow.length; z1++) {  // ignore the first 
		st1 = statRow[z1];
		field = st1.split(",");
		if (field.length < 4) { continue;}		
		
		let line= "<tr>" +
				'<td style="text-align:right">' + field[0] + '</td><td style="text-align:left">' + "words ("           + "</td>" +
				'<td style="text-align:right">' + field[1] + '</td><td style="text-align:left">' + "%), make up "      + "</td>" +
				'<td style="text-align:right">' + field[2] + '</td><td style="text-align:left">' + "% of the text ("   + "</td>" + 
				'<td style="text-align:right">' + field[3] + '</td><td style="text-align:left">' + " words)" + "</td>" +
				"</tr> \n"	;		
		result += line; 
	}  	
	result += "</table>\n";
	//console.log("result=", result)
	getById("id_frequenze").innerHTML = result;
	
} // end of js_go_updateStatistics

//---------------------------------
//---------------------------------------------------------

function onclick_mostFreqWordList_require(anyRow,toLearn, tagSel) {	
	
	fun_require_mostFreqWordList( false , "HTML page onclick_mostFreqWordList_require", anyRow, toLearn, tagSel);
	
} // end of onclick_mostFreqWordList_require

//-----------------------------------------

function onchange_mostFreqWordList_extrRow() {	
	
	fun_require_mostFreqWordList( true , "HTML page onchange_mostFreqWordList_extrRow",'','','');
	
} // end of onchange_mostFreqWordList_require

//---------------------------------

function fun_require_mostFreqWordList( swFromOnChangeExtr , caller, anyRow="", toLearn="", tagSel="") {	
	/**
	//-------------------------------------------------
// bottoni di destra
const xOnclick1lisAW = "1R" // Lista le Parole di tutti i testi             onclick_mostFreqWordList_require('anyRow','allWords')" 
const xOnclick2lisALr= "2R" // Lista le Parole da imparare                  onclick_mostFreqWordList_require('anyRow','toBeLearned')
const xOnclick3lisSW = "3R" // Lista le Parole richieste (pref suffix)      onclick_require_betweenWordList()
const xOnclick4lisSLm= "4R" // Lista i lemma richiesti   (pref suffix)      onclick_require_betweenLemmaList() 
const tagRightOnclickArr = ["",xOnclick1lisAW,xOnclick2lisALr, xOnclick3lisSW, xOnclick4lisSLm ]
	let sel_rightButton = 
	
	**/
	//console.log("%c	require_mostFreqWordList", "color:blue;")
	
	
	getById("id_inpBegError").style.display = "none"; 	
	
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	let swChg=false;
	let fromWord = Number( getById("id_inpBegFreqWList").value);	
	let numWords = Number( getById("id_inpMaxNumWords" ).value);
	if (fromWord < 1) { fromWord=1;     getById("id_inpBegFreqWList").value = 1; }
	if (numWords < 1) { numWords = 1;   getById("id_inpMaxNumWords" ).value = 1; }
	
	//console.log("anyRow=", anyRow, " fromWord=",fromWord, " numWords=", numWords, " toLearn=", toLearn); 
	
	//---
	let sel_level = "any"; //  x.options[i].id;
	//---
	let x2 = getById("id_sel_2_extrRow");
    let i = x2.selectedIndex;
	
	html_sel_extrRow = x2.options[i].id; 
	if (anyRow != "") { html_sel_extrRow = anyRow; }
	//console.log("html_sel_extrRow=", html_sel_extrRow)
	
	//console.log("  require_mostFreqWordList", " html_sel_extrRow =",html_sel_extrRow )
	
	if (last_sel_extrRow_freqWord_list == "") {last_sel_extrRow_freqWord_list = html_sel_extrRow; }  
	
	//console.log("  require_mostFreqWordList", " last_sel_extrRow_freqWord_list=", last_sel_extrRow_freqWord_list);
	
	is_selected_row_only = ( i == index_onlySelRowsWanted); //1 onclick_require_mostFreqWordLi
	
	//console.log("  require_mostFreqWordList", " is_selected_row_only =", is_selected_row_only )
	
	fun_selRowsWanted_changed();
	//---	
	let xTbl = getById("id_sel_tblwords");
    let i2 = xTbl.selectedIndex;
	let sel_toBeLearned = xTbl.options[i2].id;    //( 0 = 'allWords'   1 = 'toBeLearned' )
	if (toLearn != "") {
		sel_toBeLearned = toLearn; 
	}
	//console.log("sel_toBeLearned=" ,sel_toBeLearned); 
	
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
	
	let js_parm = JSON.stringify(  [1, anyRow, toLearn, "tag",tagSel] );
	
	go_passToJs_wordList( swChg, ""+fromWord, ""+numWords, sel_level, html_sel_extrRow, sel_toBeLearned, "js_go_showWordList_lev2", js_parm, caller);
	
	
} // end of fun_require_mostFreqWordList

//------------------------------------------
function js_go_console( str1 ) {
	//console.log( str1 )	
} 

//-------------------------------------
function sortAlpha(wordToStudy_listStr) {
	//wordToStudy_listStr
	let ww, col1, key;
	let listKey = [];
	for (let z=0; z < wordToStudy_listStr.length; z++) {
		col1 = (wordToStudy_listStr[z].trim() + ";.;.;.;.;.;.;.;.;.;.").split(";.")
		key = col1[0];  // wordCod		
		listKey.push(key  + ":" + z ); 
	}
	return listKey.sort();
	
} // end of sortAlpha 
//-------------------------------------

function sortFreq(wordToStudy_listStr) {
	//wordToStudy_listStr
	let ww, col1, key, freq1, freq2;
	let listKey = [];
	
	for (let z=0; z < wordToStudy_listStr.length; z++) {		
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
	let startIx = Number( getById("id_inpBegFreqWList").value );
	if (numberOf_uniW < startIx) {
		getById("id_inpBegErrMsg").innerHTML =  " il numero di partenza "+  startIx + " supera il numero di parole " + numberOf_uniW + " (forzato 1)"; 
		getById("id_inpBegFreqWList").value = 1
	} else {
		getById("id_inpBegErrMsg").innerHTML = ""
	}	
	getById("id_inpBegError").style.display = "inline-block"; 		
} 

//--------------------------------------------------------------
/*
				row:=   xWordAlpha.uWord2 + separ1 + 
						xWordAlpha.uWord0 + separ1 +
		 				"ix" + separ1 + 
		 				strconv.Itoa(xWordAlpha.uIxUnW_fr) + separ1 + 
		 				strconv.Itoa(xWordAlpha.uTotRow)   + separ1 +
		 				lastLemma.leLemma            	+ separ1 + 
		 				lastLemma.leTran 				+ separ1 +  
		 				separ1 							+  
		 				lastLemma.lePara	 			+ separ1 +  
		 				lastLemma.leExample	 			+ separ1 +  
		 				strconv.Itoa(totNumRow) 		+ separ1 +  			
		 				xWordAlpha.uLearnedYN              + separ1 + 			
		 				"ixLemma" + separ1 + strconv.Itoa(ix2) + separ1 +  	
		 				endOfLine 	
		 		//fmt.Println(green("word_to_row "), row )		
	*/	
//--------------------	
function sortWordFreqFirst( wordToStudy_listStr ) { 
	
	let col1, key1 , key2; 
	
	let MAXKEY = 1000000;  
	let listKey=[]; 
	
	//--------------------
	/*
	 0 xWordF2.uWord2+ ";." + 
	 1 xWordF2.uWord0 + ";." + 
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
	const ix_uWord2   = 0; 
	const ix_lemma    = 5; 
	const ix_uTotExtrRow = 6;
	const isNumber = true; 
	const isAlpha  = false; 
	const ascending = "a"; 
	const descending = "d";
	//-----------	
	
	for (let z=0; z < wordToStudy_listStr.length; z++) {		
		col1 = (wordToStudy_listStr[z].trim() + ";.;.;.;.;.;.;.;.;.;.").split(";.");		
		key1 = setKey0(isNumber, col1[ix_uTotRow ], descending, MAXKEY);
		key2 = setKey0(isAlpha,  col1[ix_uWord2], ascending,  MAXKEY);			
		listKey.push( key1 + ";;" + key2 + ";;" + (MAXKEY + z) + ":" + z  ); 		
		//if (z < 10) { console.log("lista ", " z=", z, " \t", col1[1] + "\t ix=", col1[3],  "\t", "  ix_uTotRow=", col1[ix_uTotRow ] ) }	
	}
	return listKey.sort();
	
} // end of sortWordFreqFirst	

//------------------------------

function fun_selRowsWanted_changed() {
	let ele_sel = getById("id_sel_2_extrRow");
	//let ele_listBut = getById("id_list_Righe_TD_But") 
	//let ele_listNum = getById("id_list_Righe_TD_Num") 

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
	
} // end of writeLanguageChoise
//--------------------------------------------

//--------------------------------------------------
function write_word_dictionary() {
	//console.log("write word dictionary()")
	let word1, ix1, nrow, totExtrRow2,  wLemma1, wordTran;  
	let uLearnedYN;
	let wLemmaList, wTranList, wLevelList, wParaList, wExampleList, wIxLemmaList;
	let newTranWord=0;
	let listNewTranWords = "";
	
    for (let ixW2StudyLs = 0; ixW2StudyLs < wordToStudy_list.length; ixW2StudyLs++) {
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
    let sNumWords = getById("id_inpMaxNumWords").value;
    let wordPrefix = getById("id_inpPref").value.trim();   
	

	let numWords=0; 
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
	
    go_passToJs_prefixWordList(""+numWords, wordPrefix, "js_go_showPrefixWordList", js_parm, js_caller); // ask 'go' to give wordlist by js_... function  
	
} // end of onclick_require_prefixWordList

//------------------------------------------------------

function onclick_require_betweenWordList(tagSel) {
	
	//logColor("onclick_require_betweenWordList", "%%blue","onclick_require_betweenWordList ", "%%black", tagSel)
	
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	
	let eleFromW = getById("id_inpFromABC");
	
	let fromWordPref = eleFromW.value.trim(); 		
	
	let maxNumWords = Number(  getById("id_inpMaxNumABC" ).value );	
	if (maxNumWords < 1) {  maxNumWords = 1; getById("id_inpMaxNumABC" ).value = 1; }	
	
	if (fromWordPref == "") { return; }
	
	eleFromW.style.color = null;
	eleFromW.parentElement.style.backgroundColor = null;
	/**
	// numButton=1 default ==> from onclick most frequent word list  
	// numButton=2         ==> from onclick BetweenWordList or prefix wordlist   
	// numButton=3         ==> from onclick Lemma word list   
	// numButton=5         ==> from onclick Lemma list   
	// numButton=0         ==> from word list from word, lemma, ?   
	**/
	
	let js_parm   = JSON.stringify(  [2, "tag",tagSel] );	
	let js_caller = "HTML page onclick_require_betweenWordList"
	
	//console.log("go_passToJs_betweenWordList ( fromWordPref="+fromWordPref )
	
	go_passToJs_betweenWordList(""+maxNumWords, fromWordPref, "js_go_showBetweenWordList", js_parm, js_caller); // ask 'go' to give wordlist by js_... function  
	
} // end of onclick_require_prefixWordList
//--------------------------------------

function js_go_showBetweenWordList(wordListStr, js_parm, js_caller,goFunc) {
	//logColor("%%red", "js_go_showBetweenWordList ", "%%blue", "js_parm=", js_parm, " js_caller=", js_caller )
	//logColor("%%blue", "wordListStr=", "%%black", wordListStr); 
	if (wordListStr == "") {
		getById("id_bW_err").style.display ="block";   // no entry found
    } else {
		getById("id_bW_err").style.display ="none"; 
	}
	js_go_showWordList_lev2(wordListStr, js_parm, js_caller, goFunc); 
	
} // end of js_go_showBetweenWordList	
//------------------------------------
function TOGLIjs_go_showBetweenWordList(wordListStr, js_parm, js_caller,goFunc) {
	console.log("function js_go_showBetweenWordList () ", " js_parm=", js_parm, " <-- " + goFunc + " <-- " + js_caller) ;
	//console.log("wordListStr=\n"+wordListStr +"\n--------------------------\n")
	
	if (wordListStr == "") {
		getById("id_bW_err").style.display ="block";   // no entry found
    } else {
		getById("id_bW_err").style.display ="none"; 
	}
	
	onclick_jumpFromToPage( myPage01,0, myPage02);  //myPage01
	
	//let js_parm = JSON.stringify(  [2, "", ""] );	
	js_go_showWordList_lev2(wordListStr, js_parm, js_caller, goFunc) // numButton=2         ==> from onclick BetweenWordList or prefix wordlist 
 
} // end of js_go_showBetweenWordList


//------------------------------------------------------

function onclick_require_betweenLemmaList(tagSel) {
	//logColor("onclick_require_betweenLemmaList", "%%blue","onclick_require_betweenLemmaList ", "%%black", tagSel)
	
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	
	let eleFromW = getById("id_inpLemmaFromABC");
	//let eleToW   = getById("id_inpLemmaToABC"  ); 
	let fromWordPref = eleFromW.value.trim(); 		
	//let toWordPref   = eleToW.value.trim(); 		
 
	
	let maxNumLemma = Number(  getById("id_inpLemmaMaxNumABC" ).value );	
	if (maxNumLemma < 1) {  maxNumLemma = 1; getById("id_inpLemmaMaxNumABC" ).value = 1; }	
	
	if (fromWordPref == "") { return; }
	
	/**
	if (toWordPref == "") { 
		if (fromWordPref == "") { return; }
		toWordPref = fromWordPref;
		//getById("id_inpToABC"  ).value = toWordPref; 	
	} else {
		if (fromWordPref == "") { 
			fromWordPref = toWordPref;
			getById("id_inpLemmaFromABC").value = fromWordPref; 
		}
	} 
	if (toWordPref < fromWordPref) {
		//getById("id_inpToABC"  ).value = fromWordPref; 	
		getById("id_inpFromABC").value = toWordPref;
		fromWordPref = getById("id_inpLemmaFromABC").value.trim(); 		
	    //toWordPref   = getById("id_inpLemmaToABC"  ).value.trim(); 		
	}
	**/
	eleFromW.style.color = null;
	eleFromW.parentElement.style.backgroundColor = null;
	let eleLemma = getById("id_inpLemmaFromABC");  
	eleLemma.style.color = null;
	eleLemma.parentElement.style.backgroundColor = null;
	//let eleLemma2 = getById("id_inpLemmaToABC");  
	//eleLemma2.style.color = null;
	//eleLemma2.parentElement.style.backgroundColor = null;
	
	//console.log("go_passToJs_betweenLemmaList"   , " fromWordPref=", fromWordPref)
	
	/**
	// numButton=1 default ==> from onclick most frequent word list  
	// numButton=2         ==> from onclick BetweenWordList or prefix wordlist   
	// numButton=3         ==> from onclick Lemma word list   
	// numButton=5         ==> from onclick Lemma list   
	// numButton=0         ==> from word list from word, lemma, ?   
	**/
	
	let js_parm   = JSON.stringify(  [3, "tag", tagSel] );	
	let js_caller = "HTML page onclick_require_betweenLemmaList"
	
	//console.log("  go_passToJs_betweenLemmaList  ( fromWordPref= " , fromWordPref ) 
	
	go_passToJs_betweenLemmaList(""+maxNumLemma, fromWordPref, "js_go_showBetweenLemmaList", js_parm, js_caller); // ask 'go' to give wordlist by js_... function  
	
} // end of onclick_require_betweenLemmaList


//------------------------------------
function js_go_showBetweenLemmaList(lemmaListStr, js_parm, js_caller,goFunc) {
	//logColor("%%green", "js_go_showBetweenLemmaList","%%black"," lemmaListStr.length=",  lemmaListStr.length, " js_parm=", js_parm )
	//console.log("	lemmaListStr=", lemmaListStr.substring(0,100), " ...");
	
	if (lemmaListStr == "") {
		getById("id_bW0_err").style.display ="block";   // no entry found
    } else {
		getById("id_bW0_err").style.display ="none"; 
	}
	
	//onclick_jumpFromToPage( myPage01,0, myPage02);  //myPage01
	
	//console.log("	esegue js_go_showWordList_lev2  ");
	//console.log("AAA");
	
	js_go_showWordList_lev2(lemmaListStr, js_parm, js_caller, goFunc);
	
	//console.log("	FINITO js_go_showWordList_lev2  ");
	  
} // end of js_go_showBetweenLemmaList

//-------------------------------------------------------
function onclick_require_lemmaWordList2(aLemma, tagSel) {
	// questa function è richiamata sull'ultima colonna di destra nella pagina delle lista parole 
	
	//logColor( "%%blue;font-weight:bold;","\nonclick_require_lemmaWordList2 ", "%%black","lista le parole con lemma ", aLemma, " tag=", tagSel)
	
	if (aLemma=="") return; 
	word_to_underline_list = []
	
	
	
	let eleMax = getById("idTabWRLL3")
	
	let inpMaxWordLemma = eleMax.value; 
    aLemma = aLemma.trim();    
	if (aLemma == "") {
		//ele_wordList.innerHTML ='<span style="color:red;">manca il lemma</span>';
		return;
	}		
		
	let js_parm   = JSON.stringify(  ["tag", tagSel] );	
	
	let js_caller = "HTML page onclick_require_lemmaWordList2"
	
	go_passToJs_lemmaWordList(aLemma, inpMaxWordLemma, "js_go_showLemmaWordList", js_parm, js_caller);  	

	
} // end of onclick_require_lemmaWordList2
//------------------------------
/***
function onclick_require_lemmaWordList() {
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	
	let inpMaxWordLemma = getById("id_inpMaxWordLemma").value; 
   // let aLemma = getById("id_inpLemma").value.trim();    
	**
	if (aLemma == "") {
		ele_wordList.innerHTML ='<span style="color:red;">manca il lemma</span>';
		return;
	}	
	**
	let eleFromW = getById("id_inpFromABC");
	eleFromW.style.color = null;
	eleFromW.parentElement.style.backgroundColor = null;
	//let eleLemma = getById("id_inpLemma");  
	//eleLemma.style.color = null;
	//eleLemma.parentElement.style.backgroundColor = null;
	
	//myPage01.style.display = "none"; 
	let caller = "HTML page onclick_require_lemmaWordList"
	
	go_passToJs_lemmaWordList(aLemma, inpMaxWordLemma, "js_go_showLemmaWordList," + caller);  	

} // end of onclick_require_lemmaWordList
***/
//------------------------------------
function js_go_showLemmaWordList(wordListStr,  js_parm, js_caller,goFunc) {
	
	//console.log(" js_go_showLemmaWordList () ", "wordListStr=\n" + wordListStr + "\n-------------------\n")
	
	
	if (wordListStr.substring(0,5) == "NONE,") {
		//getById("id_inpLemma_word").innerHTML = wordListStr.substring(5) 
		//getById("id_inpLemma_msg").style.display = "block"			
		myPage01.style.display = "flex"; 
		
		//onclick_jumpFromToPage( myPage02,myPage03, myPage04); 
		return
	}
	//getById("id_inpLemma_msg").style.display = "none"
	onclick_jumpFromToPage( myPage01,0, myPage02);  
	myPage01.style.display = "none"; 
	//console.log("js_go_showLemmaWordList ()  chiama js_go_showWordList_lev2")
		
	//let js_parm = JSON.stringify(  [3, "", ""] );	// numButton=3         ==> from onclick Lemma word list   
	js_go_showWordList_lev2(wordListStr, js_parm, js_caller, goFunc) 	
	
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
	
	//get_first_row_tr_visible();  // memorizza la prima TR visibile delle frasi in cui si trova questa funzione 
	
	let eleFromW = getById("id_inpFromABC");
	//let eleToW   = getById("id_inpToABC"  ); 
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
			
	//let max_num_word4LeWor = getById("idTabWoWL3").value 

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
	let eleFromW = getById("id_inpFromABC");
	//let eleToW   = getById("id_inpToABC"  ); 
	//eleToW.value = ""; 
	eleFromW.value = word1;  
    eleFromW.style.color = "blue";
	eleFromW.parentElement.style.backgroundColor = "yellow";
	//let eleLem  = getById("id_inpLemma")
	//let eleLemTD = eleLem.parentElement
	//let eleLemTR = eleLemTD.parentElement
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
	//console.log('getById("id_inpWordFra") =' , getById("id_inpWordFra").outerHTML) 
    let aWord ="";  
	if (type==2) {
		aWord = word1; 	
	} else {		
		aWord     = getById("id_inpWordFra").value.trim();  
	}	
	if (aWord == "") {
		ele_wordList.innerHTML ='<span style="color:red;">manca la parola da cercare</span>';
		return;
	}	
	
	//get_first_row_tr_visible();  // memorizza la prima TR visibile delle frasi in cui si trova questa funzione 
	
	logColor("%%blue","onclick_require_rowListWithThisWord2 --> CERCA UNA PAROLA go_passToJs_someWordsRowList ", "%%black", aWord, " maxNumRow5=",maxNumRow5);
	
	//myPage01.style.display = "none"; 
	//go_passToJs_thisWordRowList(      aWord, ""+maxNumRow5, "js_go_showWrdRowList"); 
	go_passToJs_someWordsRowList( "", aWord, ""+maxNumRow5, "js_go_showWrdRowList", js_parm, js_caller);   
	
} // end of onclick_require_rowListWithThisWord2


//--------------------------
function onclick_require_rowList1(selFrasiParole12, tagSel) {
	
	//console.log("onclick_require_rowList1 selFrasiParole12=", selFrasiParole12)
	
	getById("id_inpRowEmpty").style.display = "none";
	word_to_underline_list = []
	ele_wordList.innerHTML ="";
	
	let indexGroup= Number( getById("id_gruppi_sel").selectedIndex );
	let inpBegRow = Number(getById("id_gruppi_iBegNum"  ).value );
	let numRows   = Number(getById("id_gruppi_iNumRows" ).value );
	
	//let tagSel = tagLeftOnclickArr[selFrasiParole12]
	
	let inpEndRow = inpBegRow+numRows-1;
		
	
	getById("id_headWord").innerHTML = ""; //head1; 
		
	let caller = "HTML page onclick_require_rowList1(" + selFrasiParole12+")" //  (new Error()).stack?.split("\n")[2]?.trim().split(" ")[1] ;
	if (caller == undefined) { caller = ""; }
	
	let js_parm = JSON.stringify(  [selFrasiParole12, "tag",tagSel] );
	
	go_passToJs_rowList("" + indexGroup, ""+inpBegRow, ""+numRows, ""+selFrasiParole12, "js_go_rowList" , "js_go_showWordList_lev2", js_parm, caller); 
		
} // end of onclick_require_rowList1


//-------------------------------

function onclick_rowsByIxWord(sIxWord, tagSel) {
	if (sIxWord == "") return; 
	if (Number(sIxWord) < 0) return;
	
	//console.log("%conclick_rowsByIxWord", "color:green; font-weight:bold;");
	//console.log("lista le righe con questa parola", " sIxWord=", sIxWord, " Number(sIxWord)=", Number(sIxWord) )
	
	let max_num_row4word  = getById("idTabWRoW1").value  
	
	let caller = "HTML page onclick_rowsByIxWord(" + sIxWord +")" 
	if (caller == undefined) { caller = ""; }

	let js_parm = JSON.stringify(  ["tag",tagSel] );
	
    go_passToJs_getRowsByIxWord(""+sIxWord, ""+max_num_row4word, "js_go_showWrdRowList", js_parm, caller);  

} // end of onclick_rowsByIxWord

//-------------------------------
function onclick_rowsByIxLemma(sIxLemma, tagSel) {
	if (sIxLemma == "") return; 
	if (Number(sIxLemma) < 0) return;
	let max_num_row4lemma = getById("idTabWRoL2").value 
	
	let caller = "HTML page onclick_rowsByIxLemma(" + sIxLemma +")" 
	if (caller == undefined) { caller = ""; }

	//console.log( "%conclick_rowsByIxLemma","color:green;font-weight:bold;" ); 
	//console.log( "LISTA le RIGHE con questo lemma ", "  sixLemma=", sIxLemma, " max_num_row4lemma=", max_num_row4lemma)

	let js_parm = JSON.stringify(  ["tag",tagSel] );
	
    go_passToJs_getRowsByIxLemma(""+sIxLemma, ""+max_num_row4lemma, "js_go_showLemmaRowList4", js_parm, caller); // ask 'go' to give the rows of the word  by the go function js_go...  

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
	let upperThisW = " " + firstUpper(thisWord);  
	let righe = thisListRow.split(";;"); 
	for(let z1=0; z1 < righe.length; z1++)  {
		let riga = righe[z1].trim()
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
	
	let wordList="", wordTab=""
	let h1 = inpHeader.indexOf("<WORD>")
	let h2 = inpHeader.indexOf("</WORD>")
	//let h3 = inpHeader.indexOf("<TABLE")
	//let h4 = inpHeader.indexOf("</TABLE>")
	let h4 = inpHeader.indexOf("</HEADER>")
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
	let ele_model_tSHeadW_bdy;
	let model_tSHeadW_div;	
	getById("tsHead_4").style.display=" ERRORE non usare block per le table metti table-block"; 		
	ele_model_tSHeadW_bdy = getById("id_model_tSHead_bdy_4"); 	
	str2 += ele_model_tSHeadW_bdy.replace("§1lemma§",nuovoLemma).
											replace("§1tran§",     nuovoTran) + 
											"\n\n";  		
	*/
	
	getById("tsHead_4").style.display="block"; 		
	let ele_model_tSHeadW_bdy_inner = getById("id_model_tSHead_bdy_4").innerHTML ; 	
	let head1 = str1.split("\n"); 
	let str2="";
	for(let v=0; v < head1.length; v++) {
		let aLine = head1[v] + "|||";
		let col1 = aLine.split("|") 
		str2 += ele_model_tSHeadW_bdy_inner.replace("§4word§",col1[0]).replace("§4lemma§",col1[1]).replace("§4tran§", col1[2]) + "\n";  		
	} 	
	let ele_model_tSHeadW_DIV = getById("id_model_tSHeadW"); 
	let model_tSHeadW_div = ele_model_tSHeadW_DIV.innerHTML; 	
	let jBody  = model_tSHeadW_div.indexOf("<tbody");
	let jBody2 = model_tSHeadW_div.indexOf("<tr", jBody);		
	let newDiv = model_tSHeadW_div.substr(0, jBody2) +"\n" + str2.trim()  + "\n</tbody></table></div>\n";  
	
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
	let lineTr = str1.replaceAll("|", "<br>").replaceAll("\n", " ").split(":lemma="); 
	if (lineTr[0] == "") {  lineTr = lineTr.slice(1);}
	
	let str2 = '' ; 
	let len1 = lineTr.length
	//-------------------------------
	// caso 1: un solo lemma e una voce 
	// caso 2: un solo lemma e diverse voci 
	// caso 3: diversi lemma e una sola voce
	// caso 4: diversi lemma e diverse voci    ( non previsto )   	
	//---------------------------------------
	let oneWordOnly, oneLemmaOnly; 
	
	//-----------
	let numVoci=0
	let numLemma=0
	let preLem = "", lem1="", tran1=""
	let preVoce=""
	for(let z1=0; z1 < len1; z1++) {
		let oneTr1 = lineTr[z1].split(":")	
		lem1 = oneTr1[0].trim()
		if (lem1 != preLem) {
			numLemma++
			preLem = lem1
			tran1 = oneTr1[1].trim() 
		}
		let voce =  oneTr1[2].trim() 
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
	let ele_model_tSHeadW_bdy;
	let model_tSHeadW_div;	
	
	//console.log( "   2 buildHeaderTable", " oneWordOnly  =",oneWordOnly, "  oneLemmaOnly =", oneLemmaOnly ) ;  

	//--------------------------------
	let type = 0;

	if (oneWordOnly) {	
		getById("tsHead_1_3").style.display="block"; 
		ele_model_tSHeadW_bdy = getById("id_model_tSHead_bdy_13"); 	
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
		getById("tsHead_2").style.display="block"; 
		ele_model_tSHeadW_bdy = getById("id_model_tSHead_bdy_2"); 	
	}
	if ((oneWordOnly == false) && (oneLemmaOnly == false)) {
		type=4;
		//console.log( green("1 CASO 4 "), "4XXXXXXXX many WORD  and many LEMMA XXXXXXXXXX");
		
		getById("tsHead_4").style.display="block"; 
		
		ele_model_tSHeadW_bdy = getById("id_model_tSHead_bdy_4"); 	
		
		//console.log( green("2 CASO 4 "), "4XXXXXXXX many WORD  and many LEMMA XXXXXXXXXX")
	}
	//----------------------------------
	let model_tSHeadW_lemma  = ele_model_tSHeadW_bdy.innerHTML; 
	//console.log("aaa ", model_tSHeadW_lemma)
	let model_tSHeadW13_row2 = getById("id_model_tSHead_bdy_13_row2").innerHTML; 	
	
	
	
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
				for(let z1=0; z1 < len1; z1++) {
					let oneTr1 = lineTr[z1]	
					let jT = oneTr1.indexOf(":tran=");
					let jW = oneTr1.indexOf(":wordsInLemma=");
					let nuovoLemma   = oneTr1.substring(0,jT    ).trim();
					let nuovoTran    = oneTr1.substring(jT+6,jW ).trim();
					let nuovoLisWord = oneTr1.substring(jW+14   ).trim();						
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
		let listParole = "", lis1="";
		let nuovoLemma = "",  nuovoTran = "";
		let len2=0;
		let lenT=0;
		let LENMAX = 70
		
		for(let z1=0; z1 < len1; z1++) {	
			let oneTr1 = lineTr[z1]	
			let jT = oneTr1.indexOf(":tran=");
			let jW = oneTr1.indexOf(":wordsInLemma=");
			if (z1==0) {
				nuovoLemma   = oneTr1.substring(0,jT    ).trim();
				nuovoTran    = oneTr1.substring(jT+6,jW ).trim();
			}
			let nuovoLisWord = oneTr1.substring(jW+14   ).trim();	
			
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
					
				for(let z1=0; z1 < len1; z1++) {
					let oneTr1 = lineTr[z1]	
					let jT = oneTr1.indexOf(":tran=");
					let jW = oneTr1.indexOf(":wordsInLemma=");
					let nuovoLemma   = oneTr1.substring(0,jT    ).trim();
					let nuovoTran    = oneTr1.substring(jT+6,jW ).trim();
					let nuovoLisWord = oneTr1.substring(jW+14   ).trim();						
					str2 += ele_model_tSHeadW_bdy.replace("§1lemma§",nuovoLemma).
											replace("§1tran§",     nuovoTran) + 
											"\n\n";  									
				} // end for z1
	} // end of case_type4
	//------------------------
	
	let ele_model_tSHeadW_DIV = getById("id_model_tSHeadW"); 
	model_tSHeadW_div = ele_model_tSHeadW_DIV.innerHTML; 	
	let jBody  = model_tSHeadW_div.indexOf("<tbody");
	let jBody2 = model_tSHeadW_div.indexOf("<tr", jBody);	
	
	let newDiv = model_tSHeadW_div.substr(0, jBody2) +"\n" + str2.trim()  + "\n</tbody></table></div>\n";  
		
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
	
	let lineTr = str1.replaceAll("|", "<br>").replaceAll("\n", " ").split(":lemma="); 
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
	let ele_model_tSHeadW_bdy = getById("id_model_tSHeadW_bdy"); 
	let model_tSHeadW_lemma = ele_model_tSHeadW_bdy.innerHTML; 
	
	
	
	let jBody  = model_tSHeadW_div.indexOf("<tbody "); 
	 
	let jBodyEnd = model_tSHeadW_div.indexOf("</tbody>", jBody); 
	**/

	
	let str2 = '' ; 
	let len1 = lineTr.length
	//-------------------------------
	// caso 1: un solo lemma e una voce 
	// caso 2: un solo lemma e diverse voci 
	// caso 3: diversi lemma e una sola voce
	// caso 4: diversi lemma e diverse voci    ( non previsto )   	
	//---------------------------------------
	let oneWordOnly = true; 
	let ele_model_tSHeadW_bdy;
	//-----------
	for(let z1=0; z1 < len1; z1++) {
		let oneTr1 = lineTr[z1]	
		let jW = oneTr1.indexOf(":wordsInLemma=");
		let nuovoLisWord = oneTr1.substring(jW+14   ).trim();				
		let wor1arr = nuovoLisWord.split("<br>");
		if (wor1arr.length > 1) { 
			oneWordOnly = false;			
		}	
	} // end for z1
	//------------------------	

	
	let model_tSHeadW_div
	if (oneWordOnly) {	
		console.log("XXXXXXXX oneWordOnly = true  XXXXXXXXXX")
		getById("tsHead_2").style.display="block"; 
		ele_model_tSHeadW_bdy = getById("id_model_tSHeadW2_bdy"); 				
	} else {
		getById("tsHead_1").style.display="block"; 
		ele_model_tSHeadW_bdy = getById("id_model_tSHeadW1_bdy"); 	
	}
	

	
	//let model_tSHeadW_div = ele_model_tSHeadW.innerHTML; 

	
	let model_tSHeadW_lemma  = ele_model_tSHeadW_bdy.innerHTML; 
	let model_tSHeadW2_row2 = getById("id_model_tSHeadW2_row2_bdy").innerHTML; 	
	
	//console.log("XXXXXXXX  model_tSHeadW_lemma = ",   model_tSHeadW_lemma )
	
	//let ele_model_tSHeadW_bdyInner = ele_model_tSHeadW_bdy.innerHTML
	
	//let	jBody    = ele_model_tSHeadW_bdy.innerHTMLele_model_tSHeadW_bdymodel_tSHeadW_div.indexOf("<tbody ");	 
	//let jBodyEnd = model_tSHeadW_div.indexOf("</tbody>", jBody);  
	//----------------------------
	for(let z1=0; z1 < len1; z1++) {
		let oneTr1 = lineTr[z1]	
		let jT = oneTr1.indexOf(":tran=");
		let jW = oneTr1.indexOf(":wordsInLemma=");
		let nuovoLemma   = oneTr1.substring(0,jT    ).trim();
		let nuovoTran    = oneTr1.substring(jT+6,jW ).trim();
		let nuovoLisWord = oneTr1.substring(jW+14   ).trim();		
			
		let wor1arr = nuovoLisWord.split("<br>");
		let newLe2=""; 
		if (wor1arr.length == 1) { 
			let wor11 = wor1arr[0].trim(); 
			let lem1arr = nuovoLemma.split("<br>")    ;
			for(let h1=0; h1 < lem1arr.length; h1++) {
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
	let ele_model_tSHeadW_DIV = getById("id_model_tSHeadW"); 
	model_tSHeadW_div = ele_model_tSHeadW_DIV.innerHTML; 	
	let jBody  = model_tSHeadW_div.indexOf("<tbody");
	let jBody2 = model_tSHeadW_div.indexOf("<tr", jBody);	
	
	let newDiv = model_tSHeadW_div.substr(0, jBody2) +"\n" + str2.trim()  + "\n</tbody></table></div>\n";  
		
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
	
	let lineTr = str1.replaceAll("|", "<br>").replaceAll("\n", " ").split(":lemma="); 
	if (lineTr[0] == "") {  lineTr = lineTr.slice(1);}
	
	console.log("1 buildHeaderTable () str1=" , str1); 
 	
	let ele_model_tSHeadW = getById("id_model_tSHeadW"); 
	let model_tSHeadW_div = ele_model_tSHeadW.innerHTML; 
	
	//console.log("2 buildHeaderTable () model_tSHeadW_div=> " + model_tSHeadW_div  + "<==="); 
	
	let ele_model_tSHeadW_bdy = getById("id_model_tSHeadW_bdy"); 
	let model_tSHeadW_lemma = ele_model_tSHeadW_bdy.innerHTML; 
	
	
	
	let jBody  = model_tSHeadW_div.indexOf("<tbody "); 
	 
	let jBodyEnd = model_tSHeadW_div.indexOf("</tbody>", jBody); 
	

	
	let str2 = '' ; 
	let len1 = lineTr.length
	//-------------------------------
	// caso 1: un solo lemma e una voce 
	// caso 2: un solo lemma e diverse voci 
	// caso 3: diversi lemma e una sola voce
	// caso 4: diversi lemma e diverse voci    ( non previsto )   	
	//---------------------------------------
	for(let z1=0; z1 < len1; z1++) {
		let oneTr1 = lineTr[z1]
	
		let jT = oneTr1.indexOf(":tran=");
		let jW = oneTr1.indexOf(":wordsInLemma=");
		let nuovoLemma   = oneTr1.substring(0,jT    ).trim();
		let nuovoTran    = oneTr1.substring(jT+6,jW ).trim();
		let nuovoLisWord = oneTr1.substring(jW+14   ).trim();		
			
		let wor1arr = nuovoLisWord.split("<br>") ;
		let newLe2=""; 
		if (wor1arr.length == 1) {
			let wor11 = wor1arr[0].trim(); 
			let lem1arr = nuovoLemma.split("<br>")    ;
			for(let h1=0; h1 < lem1arr.length; h1++) {
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

	
	let newDiv = model_tSHeadW_div.substr(0, jBody) + str2 + model_tSHeadW_div.substr(jBodyEnd+8);  
	
	return newDiv;  
	
} // end of OLDbuildHeaderTable

//--------------------------------------------------
/**
function js_go_showWordRowList3(inpstr) {
	// vedi js_go_showWrdRowList(inpstr)
	console.log(" js_go_showWordRowList2 (inpstr=" + inpstr);  
}
**/
//--------------------------------------------------
function js_go_showLemmaRowList4(inpstr, js_parm, jsFunc,goFunc) {
	// vedi js_go_showWrdRowList(inpstr)	
	//console.log(" js_go_showLemmaRowList2 (inpstr=" + inpstr);  
	js_go_showWrdRowList(inpstr, js_parm, jsFunc,goFunc)
}
//------------------------
function js_go_showWrdRowList(inpstr, js_parm, jsFunc,goFunc) {
	
	// triggered by go ( go_passToJs_rowList and js_go_showWrdRowList)
	
	//logColor("%%blue", "js_go_showWrdRowList () ", "%%black","  js_parm=" + js_parm + "\n\t jsFunc=" + jsFunc , "\n\t goFunc=" + goFunc ) 
	
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
		//console.log(" js_go_showWrdRowList () 3 return inpstr = " +inpstr);  
		getById("id_inpWordFra_msgWord").innerHTML = inpstr.substring(5) ;
		//getById("id_inpWordFra_msg").style.display = "block";
		myPage01.style.display = "flex"; 
		myPage05.style.display = "none";
		//onclick_jumpFromToPage( myPage02,myPage03, myPage04);  //   
		return
	}	
	//getById("id_inpWordFra_msg").style.display = "none";
	
	//console.log("2 js_go_showWrdRowList ");
	
	myPage05.style.display = "none";
	
		
	let h_wordListStr = "", h_wordTab = "";
	let ks = inpstr.indexOf("</HEADER>"); 
	if (ks < 0) { return } 	
	
	//console.log("3 js_go_showWrdRowList ");
	
	let inpHeader = inpstr.substring( 0, ks + 9);
	
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
	
	
	let thisListRow = inpstr.substring(ks+9) ; 
	//console.log("ANTONIO resto \n", inpstr	, "\n------------------------------------ fine ----")
	
	let col1 = splitHeader( inpHeader );
	let h_wordListStr00 = col1[0]
	
	let inpReqWord = "";
	let jh = h_wordListStr00.indexOf(",L:");	
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
	
	let word3, ixUnW3, totRow3, wLemma3, wTran3;

	
	
	getById("id_headWord").innerHTML = h_wordTab.replaceAll('display:none','display:block').replaceAll("tsHead_1","tsHead_00") ;
	
	//console.log("6 js_go_showWrdRowList");
	
	onclick_jumpFromToPage( myPage02,myPage03, myPage04);  //   
	
	//js_go_rowList ( thisListRow );  
	js_go_rowList( thisListRow, js_parm, jsFunc,goFunc)
	
	
} // end of js_go_showWrdRowList  NEW

//--------------------------------------
function js_go_rowList( inpstr, json_parmStr, jsFunc,goFunc) {
	//console.log("1 js_go_rowList " ,  "myPage01=", myPage01.style.display,   ",  myPage04=", myPage04.style.display); 
	// triggered by go ( go_passToJs_rowList and js_go_showWrdRowList)
	
	//console.log("function js_go_rowList () js_parm=" + js_parm + "\n\t jsFunc=" + jsFunc , "\n\t goFunc=" + goFunc ) 
	//console.log("	inpstr=" +inpstr ) 
	let js_parm_array = fromJsonStringToArray(json_parmStr)
	
	//sw_rowListFrom_onclick = (jsFunc.indexOf("HTML page onclick_require_rowList1(1)" ) >= 0) ;
	
	sw_rowListFrom_onclick = (js_parm_array[0] == "1") ;
	/**
	logColor("js_go_rowList", "%%red", "js_go_rowList ", "%%blue", " sw_rowListFrom_onclick=", "%%black", sw_rowListFrom_onclick, 
		 "\n\t js_parm_array=" + js_parm_array + "\n\t jsFunc=" + jsFunc , "\n\t goFunc=" + goFunc  )
	**/
	
	getById("id_caller").innerHTML = jsFunc;
	
	
	rowToStudy_list = [];
	newRowTran = [];
	let numeroTS_Row=0, numeroTS_OkTran=0, numeroTS_NoTran=0;
    if (inpstr == undefined) {
		console.log("js_go_rowList () 1 return inpstr undefined "); 	
		onclick_jumpFromToPage( myPage02,myPage03, myPage01);  //   
        return;
    }
	
    if (inpstr == "") {		
		console.log("js_go_rowList () 2 return inpstr vuoto");  
		getById("id_inpRowEmpty").style.display = "inline-block";
		onclick_jumpFromToPage( myPage02,myPage03, myPage01);  //   
	    return;
    }
	//myPage01.style.display = "none"; 
	
	rowToStudy_list =  inpstr.split("<br>");	
	
	//console.log("js_go_rowList () 1 rowToStudy_list.length=", rowToStudy_list.length , "  maxNumRow=",  maxNumRow , " type=", typeof maxNumRow);
	
	if (rowToStudy_list.length >= maxNumRow) {
		rowToStudy_list = rowToStudy_list.slice(0, maxNumRow+1) ; 
	}  
	
	//console.log("js_go_rowList () 2 rowToStudy_list.length=", rowToStudy_list.length );
	
	for (let z=0; z < rowToStudy_list.length; z++) {	
		//if (z < 10) { console.log("js_go_rowList ", z, "  ", rowToStudy_list[z]   )  }
		newRowTran.push( 0 )
	}
	
	//console.log("2 js_go_rowList build_Page1_rowsToTranslate " ,  "myPage01=", myPage01.style.display,   ",  myPage04=", myPage04.style.display);   
	
	[numeroTS_Row, numeroTS_OkTran, numeroTS_NoTran] = build_Page1_rowsToTranslate("3js_go_rowList"); 
	
	//console.log("js_go_rowList () numeroTS=", numeroTS_Row, " numeroTS_OkTran=", numeroTS_OkTran, " numeroTS_NoTran=", numeroTS_NoTran ) ;
	
	myPage01.style.display = "none"; 
	
	
} // end of js_go_rowList
//-------------------

function build_Page1_rowsToTranslate( wh ) {
	//console.log("1 build_Page1_rowsToTranslate " ,  "myPage01=", myPage01.style.display,   ",  myPage04=", myPage04.style.display);   
	
	
	let numeroTS_Row=0, numeroTS_OkTran=0, numeroTS_NoTran=0;
	let numRowNoTranR1 = 0 
	let rows_to_translate_str = "";  
	
	let nfile,idRow, ixRow,p3,rowS, tranS, nfileS, idRowS, ixRowS, oT; 
	
	//console.log("\nXXXXXXXXXXXXXXXXXXXXXXXX   build_Page1_rowsToTranslate (", wh,")",  "  1  length=", rowToStudy_list.length )
	
	for (let z=0; z < rowToStudy_list.length; z++) {		
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
	getById("id_notTranNumRow").innerHTML = numRowNoTranR1; 
	
	ele_toTranslate_textarea.value = rows_to_translate_str;
	ele_translated_textarea.value = ""; 
	
	//console.log("build... ele_toTranslate_textarea = ", rows_to_translate_str); 
	 
	//onclick_jumpFromToPage( myPage01,0,myPage04);   
	onclick_jumpFromToPage( myPage01,0,myPage04);   
	
	return [numeroTS_Row, numeroTS_OkTran, numeroTS_NoTran ];
	
} // end of build_Page1_rowsToTranslate	

//-------------------------------------------------

function go_run_js_showReadFile( str1 ) {
	//console.log("ANTONIO go_run_js_showReadFile( str1=" , str1 )
	let j1 = str1.indexOf( "))" ); 
	let mainNum     = str1.substring(0,j1);
	let fileListStr = str1.substring(j1+2); 
	
	//  "level " + msgLevelStat + "))" 	
	let numUniW, numTotW, numRow, levelStats = "";
	[numUniW, numTotW, numRow, levelStats] = mainNum.split(";")
	numberOf_uniW = numUniW;
	numberOf_totW = numTotW;
	numberOf_Row  = numRow;
	str1 = parseInt(numUniW).toLocaleString() + " parole diverse, " + 
			" in totale " + parseInt(numTotW).toLocaleString() + " parole " + 
			"su un testo di " + parseInt(numRow).toLocaleString() + " righe"  ;
			
	getById("id_inpFileList").innerHTML	= "(" + str1 + ")" ; 	
	/**
	//console.log("go_run_js_showReadFile  levelStats=", levelStats)
	
	if (levelStats.indexOf("-oth-: 99%") < 0)  {		
		str1 += "<br>" + levelStats; 			
	}
	
    let rows = fileListStr.split(";");
    let td1, td2;
    //let modelTR = getById("id_start_trModel").outerHTML;
	str1 += "<hr>"
	let str2 = '<table style="border:0px solid black">\n';
	//str2 += '<tr><th colspan="2">input</th></tr> \n' ;   	
    for (let i = 0; i < rows.length; i++) {
        if (rows[i] == "") continue;
        [td1, td2] = rows[i].split("<file>");			
		let k1= td2.lastIndexOf("\\")
		let k2= td2.lastIndexOf("/")	
		let k3 = Math.max(k1,k2) ;
		listaInputFile = td2.substring(k3+1).trim() + "," 
		str2 += '<tr><td>' + td2.substring(k3+1) + '</td>' +
			'<td>(' + parseInt(td1).toLocaleString() + " righe" + ')' + '</td></tr> \n' ;  
    }
	str2 += '</table>';
	
	let str0 = '<div style="border:1px solid black; padding:1em;">';
	
	getById("id_inpFileList").innerHTML = str0 + str1 + ""+ str2 + "</div>";
	***/
	
}
//---------------
/**
function set_dragDiv() {
		dragElement(getById("dragButt1"  ), getById("dragButt1_header"  ));
		dragElement(getById("dragButt2"  ), getById("dragButt2_header"  ));
		dragElement(getById("dragButt3"  ), getById("dragButt3_header"  ));
		dragElement(getById("dragButt4"  ), getById("dragButt4_header"  ));
		dragElement(getById("dragWLsTip5"), getById("dragWLsTip5_Header")); 
}
**/
//------------------------------------
function go_run_js_ready( prevRun00) {
	console.log("************************* go_run_js_ready(" + prevRun00.trim() + ")" ); 
	
	//try{
	/**
	5,de,de-DE,Microsoft Stefan - German (Germany):mainpage_value=5491,10,1,100,any,,
	**/
	
	//set_dragDiv();
	
	prevRun00+="                                   ";
	
	let jj = prevRun00.indexOf(":mainpage_value=");
	let lastMainPageValueS;
	let prevRun = "";
	if (jj < 0) {
		lastMainPageValueS = "";	
		prevRun = prevRun00;
	} else {		
		lastMainPageValueS = prevRun00.substr(jj+16);  	
		prevRun = prevRun00.substr(0,jj)
	}
	let sel_lev  = ""; let sel_ix_lev=0;  let sel_id_lev="";
	let sel_extr = ""; let sel_ix_extr=0; let sel_id_extr="";
	let isAlphaStr=""; 
	
	//let ele_sel_1_lev    = getById("id_sel_1_levTOLTO");
	let ele_sel_2extRow  = getById("id_sel_2_extrRow");
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
	
	is_selected_row_only = (sel_ix_extr == index_onlySelRowsWanted); //3go_run_js_ready
	fun_selRowsWanted_changed();
	//console.log("go_ready:  is_selected_row_only = ",is_selected_row_only ); 
	
	//-------
	/***
	if (isAlphaStr == "alpha") {
		ele_alpha.checked = true;
	} else {
		ele_freq.checked  = true;
	}
	let eleSele = getById("id_orderWord1TOLTO");	
	let swIsAl = ( eleSele.selectedIndex == 0) 
	***/
	/***
	let eleSele = getById("id_orderWord1TOLTO");	
	if (isAlphaStr == "alpha") {
		eleSele.selectedIndex = 0;  // alphabetic 
	} else {
		eleSele.selectedIndex = 1;  // by frequence 
	}
	***/
	//-------------------
	
	let prevRunLanguage = prevRun.trim(); 
	if (prevRunLanguage != "") {
		sw_firstDictLine_already_existed = true; 
		lastRunLanguage = prevRunLanguage
		console.log("language file has been read ==>" +  lastRunLanguage) 
	} 	
  
    getById("id_start001").style.display = "none";
    myPage01.style.display = "flex";
	
    getById("id_showButt").style.display = "block";
    getById("id_start_tab").style.display = "none";
	
	//scroll_1_init() 
	
	if (prevRunLanguage != "") { 
		lastRunLanguage = prevRunLanguage
		loadPrevLang( prevRunLanguage ) 
		console.log("go_run_js_ready()  prevRunLanguage=", prevRunLanguage, "  go_run_js_ready() NON chiama  fcommon_load_all_voices()");  
	}	else {
		console.log("go_run_js_ready() call fcommon_load_all_voices()");
		
		fcommon_load_all_voices(); // at end calls tts_1_toBeRunAfterGotVoices()		
		// WARNING: the above function contains asynchronous code.  
		// 			Any statement after this line is executed immediately without waiting its end			
	}
	
	
	//onclick_getRowGroup(  getById("id_gruppi_sel") )
	
	let cellWord_TrTD ; // = getById("idTableWordList_tbody").children[0].children[5]; 		
	let cellWord_TrTH ; // = getById("idTableWordList_thead").children[0].children[5]; 
		
	
	//} catch(e1) { console.log("%cerrore in go_run_js_ready () " +  e1, "color:red;") }
	
} // end of go_run_js_ready
//-------------------------------------

function onclick_show_or_hide_statistics() {
    let eleFreq   = getById("id_frequenze");
    let eleStFile = getById("id_start_tab");


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

	let CoerInp = inpCode.replaceAll( "ae","ä").replaceAll("oe","ö").replaceAll("ue","ü").replaceAll("ß","ss") 
					
	return CoerInp  
}
//----------------------------

function evidenzia( unaparola, class_targ, txtinp1) {
		
	
	let newRow="";
	unaparola = unaparola.toLowerCase().trim();
	
	let lenParola = unaparola.length; 
	
	let lowinp0 = txtinp1.toLowerCase().replaceAll("§"," ") + "§"; 
	let lowinp1 = lowinp0.replace(/[\s;,:"'\.<>»«()\[\]\!\?„“]/g,"§")  
	
	let j0=-1, j1=0;
	let jNew=0
	let swBold=false;
	let parola1, parolaT;
	j1=-1; 
	for(let i=0; i < lowinp1.length; i++) {	    
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
	onclick_copyTextAreaValue_to_clipboard( getById('myInput') );
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
	
	let wordTranList = ele_wordsTranslated.value.trim().split( wordTTEnd );    // 
	
	console.log("X3 onclick_showWordsButton() ", " type=", type, " wordTranList=\n", wordTranList ) 	
	
	let lenTran =  wordTranList.length;	
	let wordTran
		
	let ixUnTrad, ixUnWtS;
	let lenW = wordToStudy_list.length;
	let word1, nrow, totExtrRow2, wLemma1, wtran, uLearnedYN  ; 
	let newTran_f = "";
	//------------------------------------------
	let wX, wT, ixz1,  ixz2, ixzNum   
	
	//-----------
	let wList;
	let sw_someTranMissingW = false;  // sometimes the automatic translator does some mistakes 
	let wIxLemmaList, wLemmaList, wTranList, newTranList , wLevelList, wParaList, wExampleList; 
	let numf=0
	let sw_Minus1 = false
	//----------------------------
	for(let z=0; z < lenTran; z++) {
	
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
		let ix3;	
		try{
			ix3 = parseInt( ixzNum );
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
		
		let numIxTrad= parseInt(ixUnTrad)
		if (numIxTrad == -1) { sw_Minus1 = true; numIxTrad = 0; }
		if ( parseInt(ixUnWtS) != numIxTrad) {
			sw_someTranMissingW= true;	
			if (sw_ignore_missTranWord == false) {
				console.log("onclick_showWordsButton() ", red("error4w"), " entry z=" + z + " = " + wT  + "\n\t wordToStudy_list[ixzNum="+ ixzNum +"]=" +  wordToStudy_list[ixzNum] +  
					"\n\tixUnWtS=" + ixUnWtS + " ixUnTrad=" + ixUnTrad);			
				continue; 
			}	
		} 
		let numLemma = parseInt( numf ); // index of the lemma and its translation in the [lemmalist][tran list] in wordToStudy_list 
		newTran[ixzNum] = 1; 
		
		//console.log("anto newTran[ixzNum=" + ixzNum + "] = 1" ); 
		
		for(let h = 0 ; h < wLemmaList.length; h++) {	
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
		
	}  // end of for(let z ...
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
	for (let i = 0; i < wordToStudy_list.length; i++) {
		[word1, ixUnWtS, nrow,   wLemmaList, wTranList, 
					wLevelList, wParaList, wExampleList, totExtrRow2, uLearnedYN , wIxLemmaList  ] = wordToStudy_list[i];  
		if (word1== "Xabfahren") { console.log("word1=", word1, " wParaList=", wParaList, " wExampleList=", wExampleList)}			
		//console.log("onclick_showWordsButton() 3 i=",i, " wordToStudy_list[i]=" , wordToStudy_list[i])
		if (word1 == "") { continue; }
				
		for(let ixLemma = 0 ; ixLemma < wLemmaList.length; ixLemma++) {	
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
		for (let i = 0; i < wordToStudy_list.length; i++) {  
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
		// let wordOrig2 = wordOrigList[i].trim(); 		
		let pp= wordOrig2.indexOf(";"); 
		
		//console.log("antonio extract_ix_word( i=" + i + "  wordOrig2=" + wordOrig2 + " ==> pp=" + pp)  
		
		if (pp< 0) { return "";}
	
		let ix2 = wordOrig2.substring(0, pp); 
		let	ww2 = wordOrig2.substring(pp+1).trim(); 		
		
		//console.log("\t\t i=" + i + "   ix2=" + ix2 + "   ww2=" + ww2)  
		
		try{
			let ix3 = parseInt( ix2 );
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
	let pvLang = prevLanguage.trim).split(",") 	
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
function go_run_js_setError(msg) {
	getById("id_startwait").innerHTML = msg; 	
}
//------------------------------
let numP=0
function whereIs(here) {
	//numP++
	
	//if (numP < 10) { console.log( "\whereIs(" + here) }
	
	if (here.substring(0,2) != "::") {	
		//ele_where.innerHTML = here ;
		return ;
	}
	
	let cols = here.split("::")
	//ele_where.innerHTML = cols[2];
	let perc=0;
	try{
		perc = parseInt( cols[1] )
		ele_bar.style.width = perc + "%";  
		ele_bar2.style.width = (100-perc) + "%";  
		ele_bar3.innerHTML = perc + "%";  
	} catch(e1) {
	}	
 	
	
} 
//-------------------------
function go_run_js_showProgress(perc0) {
	let perc=0;
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
	
	let id_analWords = "idw_" + numTr;
	let ele_wordset = getById(id_analWords);   
	if ((ele_wordset == null) || (ele_wordset == false) ) {
			console.log("onclick_tts_seeWordsGO1(numTr=", numTr, " ixRow=", ixRow, ") ==> ", ele_wordset.outerHTML, "\n\tid_analWords=" + id_analWords , " ERROR ele_wordset == false" )
		return
	}
	
	get_first_row_tr_visible();  // memorizza la prima TR visibile delle frasi in cui si trova questa funzione 
	
	ele_wordset.innerHTML = ""; 
		
	go_passToJs_rowWordList(""+numTr,""+ixRow, "js_go_rowWordList", js_parm, js_caller); // ask 'go' to give wordlist by js_... function  
		
} // end of onclick_tts_seeWordsGO1



//------------------------------------
function js_go_rowWordList(wordListStr) {
	
	logColor("%%blue","js_go_rowWordList(wordListStr=", "%%black", wordListStr)
	// riga0: numIdOut,ixRR,rNumWords, len(rowX.rListIxUnF))
	// riga n:  xWordAlpha.uWord2,uIxUnW_fr,[Lemma;paradigma;traduzione;example;numWords]
	/*
	esempio wordListStr=
	2,2138,2,2;;
	der,         6,1995;[der;       der;                il;Hier ist der Brief, den du suchst.;11] ;;
	skiurlaub,7856,   1;[skiurlaub;    ;vacanza sulla neve;                                  ; 1] ;;
	*/
	
    // triggered by go func (  go _ passToJs_wordList )
    if (wordListStr == undefined) wordListStr = ""
    if (wordListStr == "") {
        console.log("js_showWordList: parameter is empty");
		onclick_jumpFromToPage( myPage02,myPage03,  myPage01); 
        return;
    }	

    let rowWordList = wordListStr.split(endOfLine);	
	
	//console.log("rowWordList=", rowWordList)
	/**
	wordListStr=
	2,2195,7,7;;
	ja,18,783;[ja;ja;sì;Sind Sie Herr Watanabe? / Ja.;1] ;;
	ich,0,5024;[ich;ich;Io;Ich heiße Veronika.;1] ;;wohne,424,42;[wohne;;vivi;;1] [wohnen;wohnen;vivere;Ich wohne in München.;4] ;;erst,237,75;[erst;;prima;;1] ;;seit,193,93;[seit;seit;dal;Ich wohne seit drei Jahren in Köln.;1] ;;gestern,218,84;[gestern;gestern;ieri;Gestern war ich krank.;1] ;;
	hier,37,484;[hier;hier;qui;Hier ist 06131-553221, Pamela Linke.;1] ;;
	// riga n:  xWordAlpha.uWord2,uIxUnW_fr,[Lemma;paradigma;traduzione;example;numWords]
			   blonde,3613,2;[blond;;biondo;;3] [blonde;;biondo;;1] ;.
	
	
	**/
	let numId_0 = rowWordList[0].split(",");
	let numId = (""+numId_0[0]).trim();  	
	
	console.log("numId=", numId) ;
		
	let prevTR  = getById( "idtr_" + numId); 
	if ( prevTR  == undefined) {
		return;
	}
	console.log("2 js_go_rowWordList ") 
	
	let anal_txt = ""; 
	let anal_tts_txt=""; 
	
	let idc1 			= "idc_"  + numId;
	let idtts			= "idtts" + numId; 
	let id_analWords 	= "idw_"  + numId;
	
	let anal_ele_idc   	= getById(idc1 );		
	let anal_ele_idtts 	= getById(idtts);
	let ele_wordset 	= getById(id_analWords);   
	
	console.log("3 js_go_rowWordList ") 
	
	if (anal_ele_idc) {
		anal_txt = anal_ele_idc.innerHTML;
		anal_tts_txt= anal_ele_idtts.innerHTML; 
		console.log("4 js_go_rowWordList ") 
	} else {
		console.log("js_showWordList: idc1=", idc1 , "(anal_ele_idc==false)" )  
		return;
	}
	console.log("5 js_go_rowWordList ") 
	
	let eleTR = anal_ele_idtts.parentElement.parentElement.parentElement; 
	let trHeight = eleTR.offsetHeight; 	
	
	console.log("last_ele_analWords_id=", last_ele_analWords_id)
	let last_ele_analWords
	if (last_ele_analWords_id != "") {	
		last_ele_analWords = getById(last_ele_analWords_id, true);
		if (last_ele_analWords == undefined)  last_ele_analWords=""
	}
	if (last_ele_analWords_id != "") {		
		// remove the previous 
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
	
	let ixLastEle, table_txt
	[ ixLastEle, table_txt] = tts_3_spezzaRiga3( rowWordList.slice(1) )	 ;   
 	
	let divWord = "";
	divWord += `<div style="font-size:0.6em;color: black;text-align:center;margin-top:0.5em;">					
				clicca su una parola per ottenere paradigma, traduzione, esempi, toccala senza cliccare per avere la traduzione   							
			</div> \n`; 
	divWord += table_txt;
	
	divWord += anyOtherWord();  
	
	//if (table_txt.indexOf("creazione")>= 0 ) { console.log("table_txt=" , table_txt); }
	ele_wordset.innerHTML = divWord; 
	
	//console.log("\nXXXXXXXXXXXXX\nXXXXXXXXXXXX\n ele_wordset.innerHTML = " + ele_wordset.innerHTML ) 
	//console.log("ixLastEle=", ixLastEle)
	if (ixLastEle > 0) {
		//console.log("ULTIMO")
		let eleF = getById("wb1_0");
		let eleT = getById("wb2_" + ixLastEle );
		onclick_tts_word_arrowFromIx(eleF, 0,         true, false)
		onclick_tts_word_arrowToIx(  eleT, ixLastEle, true, false)
	}	
	
	/**
	let prevTR; 
	for (var k=numId-4; k <=numId; k++) {       // try to start from some rows before the required  
		prevTR  = getById( "idtr_" + k, true); 
		if (prevTR) break; 
	}
	**/
	
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
	getById("id_msg16").innerHTML = "";

	sw_translation_not_wanted = false;

	let msgerr0 = "";
	let msgerr1 = fr_tts_1_get_orig_subtitle2(); // get orig. text /srt
	let msgerr2 = fr_tts_1_get_tran_subtitle2(msgerr1); // get tran. text/srt	

	msgerr0 += msgerr1 + msgerr2;
	let msgerr3 = "";

	msgerr0 += msgerr3;

	if (msgerr0 != "") {
		tts_1_putMsgerr(msgerr0);
		getById("id_msg16").style.color = "red";
		return -1;
	}
	getById("id_msg16").style.color = null;

	tts_9_toBeRunAfterGotVoicesPLAYER(); 

	return 0;

} // end of tts_1_join_orig_trad()

//----------------------

//--------------------------------------------------
 function fr_tts_1_get_orig_subtitle2() {
     
      let msgerr1 = "";
    
      //builder_orig_subtitles_string = getById("txt _ pagOrig").value.trim();
	  
	  builder_orig_subtitles_string = ele_toTranslate_textarea.value.trim();
      
	  if (builder_orig_subtitles_string != "") {
          sw_inp_sub_orig_builder = true;
      } else {
          sw_inp_sub_orig_builder = false;
          msgerr1 += "<br>" + tts_1_getMsgId("m132"); //  ma22 the source language subtitle file  has not been read or is empty" ;         
      }
	  inp_row_orig = builder_orig_subtitles_string.split("\n") ;
	  inp_row_orig.push("");    
	  //numOrig = inp_row_orig.length;
	
      return msgerr1;

  } // end of get_orig_subtitle2()
  //--------------------------------------------------

  function fr_tts_1_get_tran_subtitle2(msgerrOrig) {
     
      let msgerr1 = "";

      //builder_tran_subtitles_string = getById("txt _ pagTrad").value.trim();
	  
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
	  //numTran = inp_row_tran.length;  
	  
      return msgerr1;

  } // end of tts_1_get_tran_subtitle2()

//--------------------------
function replaceSepar( wT1 ) {	
	/**
	1;;1374;;sie zum zweiten Male jaeh zu verlassen gezwungen war, so hatte er sie
	1, 1374, e fu costretto a lasciarla per la seconda volta, l'ebbe
	**/
	let k1,k2,k3,kx1,kx2, kx10,kx20, num1,num2,new_wT;
	
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
	
	
	let rest = wT1.substring(kx20).trim(); 
	new_wT = num1 + ";;" + num2 + ";;" + rest;
	try {
		let num1Nu = parseInt(num1);
		let num2Nu = parseInt(num2);
	} catch(e) {
		return wT1; 		
	} 
	
	return new_wT; 
	
} 
//--------------------------------------

function onclick_showRowsButton(type) {
	
	//console.log("1 onclick_showRowsButton " ,  "myPage01=", myPage01.style.display,   ",  myPage04=", myPage04.style.display);   
	
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
	
	//console.log("onclick_showRowsButton (type=", type, ")  sw_ignore_missTranRow=", sw_ignore_missTranRow); 
	
	//---------------------------------------------------
		
	//	console.log("onclick_showRowsButton ele_translated_textarea = ", ele_translated_textarea.value); 	 
		
	let rowTranList = ele_translated_textarea.value.trim().split( "\n" );  
	
	let lenTran = rowTranList.length;	
	let rowTran
	let ixRowS, ixRow, ouIxRowS, ouIxRow, rowNewTran; 
	//------------------------------------------

	let nfileW, idRowW, ixRowW, rowW, tranW, gruppoW;
	//-----------
	let wList;
	let sw_someTranMissingR = false;  // sometimes the automatic translator does some mistakes 
	let wT;
	let numNoTranRow = 0;
	let numNoTranRow2 = 0;
	let numeroTS_Row=0, numeroTS_OkTran=0, numeroTS_NoTran=0;
	let newAddedTran = 0; 
	//---------------------------------
	
	let lenW    = rowToStudy_list.length;
	
	
	//----------------------------
	//  scandisce le righe di traduzione copiate dal traduttore google (potrebbero essere meno di rowToStudy_list ) 
	
	for(let z=0; z < lenTran; z++) {
	
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
		
		
		let oldRow = rowToStudy_list[ouIxRow];
		//console.log("     rowToStudy_list[ouIxRow] = "  +  oldRow )
		// accoppia la riga di traduzione con quella originale  
		try {
			[nfileW,idRowW, ixRowW, rowW, tranW, gruppoW] = oldRow.split("|");
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
		
		rowToStudy_list[ouIxRow] = nfileW + "|" + idRowW + "|" + ixRowW + "|" + rowW + "|" + rowNewTran + "|" + gruppoW ; // update with translation 	
		
		//console.log("  xxx  NUOVA rowToStudy_list[ouIxRow=",ouIxRow , "] = " + rowToStudy_list[ouIxRow]	)
		
		newAddedTran++; 
					
	
	}  // end of for(let z ...
	//--------------------------	
	//console.log("2 onclick_showRowsButton build_Page1_rowsToTranslate " ,  "myPage01=", myPage01.style.display,   ",  myPage04=", myPage04.style.display);   
	
	[numeroTS_Row, numeroTS_OkTran, numeroTS_NoTran ] = build_Page1_rowsToTranslate("2onclick_showRowsButton (type=", type);
	
	getById("id_notTranNumRow").innerHTML = numeroTS_NoTran; 
	
	write_row_dictionary(1); 
	
	if (numeroTS_NoTran > 0) {
		if (sw_ignore_missTranRow == false) {
			return;  
		}
	} 
	
	showRowsAndTranButton("3");
	
} // end of onclick_showRowsButton
//-------------------------------------------

function write_row_dictionary(wh) {
	let nfileW,ixRowW, rowW, tranW, gruppoW; 
	let word1, ix1, nrow, wLemma1, wordTran, col1;
	
	let newTranRow=0;
	let listNewTranRows = "";
	
	//console.log("\n----------------------\nwrite_row_dictionary() rowToStudy_list=" ,  rowToStudy_list ,"\n----------")
	let idRow1, ixRow1; 
	let id_ix_Row;
	
	for (let i = 0; i < rowToStudy_list.length; i++) {
		if ( rowToStudy_list[i] == "") continue; 
		//console.log("2977write " ,  rowToStudy_list[i] )
		
		col1 = (rowToStudy_list[i]+ "||||||").split("|");		
		nfileW = col1[0];  
		idRow1 = col1[1]; 
		ixRow1 = col1[2];
		rowW   = col1[3]; 
		tranW  = col1[4];
		gruppoW= col1[5].replaceAll("|"," "); 		
		
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
	
	go_write_row_dictionary(  listNewTranRows ,"", "", "");  		
	
} // end of write_row_dictionary

//------------------------------------------

function show_altreRighe(this1,numRows) {
	
	let eleTD = this1.parentElement
	
	let eleOnOff = eleTD.children[1]		
	
	//console.log("================\nantonio  show_altreRighe() 1 eleOnOff=", eleOnOff.innerHTML  , " display=", eleOnOff.style.display, " eleTD outer=", eleTD.outerHTML) 
	
	let showSPAN, showTR;
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
	
	let eleTR = eleTD.parentElement	
	

	let nextTR = eleTR;
	let nextTR2 ;
	
	for (let rr=0; rr < numRows; rr++ ) {			
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
	let ele_tran = this1.parentElement.children[nch];	
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
	let ele_tran = this1.parentElement.children[nch];	
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
	let ixWord = 0;
    try {
        ixWord = parseInt(sIxWord);
    } catch (err) {}	
	
	let num1 = 1*this1.innerHTML; 
	let ele_td= this1.parentElement;
	if (yesNo==0) {		
		let ele_nextTd = ele_td.nextElementSibling; 
		
		let next_this1 = ele_nextTd.children[0]
		next_this1.innerHTML = 0;
		next_this1.style.border = null; 
		go_passToJs_word_known(""+ixWord, "1", "0", "js_go_word_known", js_parm, js_caller); // ask 'go' to update yes/no word known ctr  
		return 
	} 
	num1++;
	this1.innerHTML = num1;
	if (num1 > 0) {
		this1.style.border = "5px solid black"; 
	} else {
		this1.style.border = null; 
	}
	
    go_passToJs_word_known(""+ixWord, ""+yesNo, ""+num1, "js_go_word_known", js_parm, js_caller); // ask 'go' to update yes/no word known ctr  
	
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
	getById("id_lastWSort").innerHTML = nField1;							
	getById("id_wordList1").scrollTop = 0;
		
	let ele_tbody = getById("idTableWordList_tbody"); 
	let num_tr = ele_tbody.children.length; 
	let ele_tr, ele_td, ele_butt,  ele_div0, ele_tran;  
	let num_td =0; 
	let trad="";
	
	const EMPTY  = "_none_"; 
	
	let listKey=[ EMPTY ]; // lascio l'entrata 0 occupata 
	let key1 , key2; 
	let ke2, ix1, ix2; 
	let MAXKEY = 1000000;  
	let ele_td5 ;
	//---------------------------------
	for(let g=0; g < num_tr; g++) {
		//if (g > 20) { break; }
		
		ele_tr = ele_tbody.children[g]; 
		
		//if (g==2) {console.log("onclick_sortWordBy_ixField()1.1 ", " num_tr=", g, " ==> TR=", ele_tr.innerHTML);}
		
		num_td = ele_tr.children.length; 
		//--
		ele_td = ele_tr.children[nField1]; 	
		ele_butt = ele_td 		
		for(let f2=0; f2 < 10; f2++) {
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
		for(let f2=0; f2 < 10; f2++) {
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
	
	let newBodyInner = "";
	let newTd;
	//--------------------
	let gg=0
	for(let g=0; g < listKey.length; g++) {
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
				let child0 = ele_tr.children[0]; 
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
		let num1 =0;	
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
		let num1 =0;	
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
	
	go_passToJs_write_WordsToLearn("js_go_file_words_to_learn_written", js_parm, js_caller); 
	
}
//-----------------------------------------
function js_go_file_words_to_learn_written( str1 ) {
	getById("id_w_to_learn_written").innerHTML = str1 ;	
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
	
	let ele_tbody = getById("idTableWordList_tbody"); 
	let num_tr = ele_tbody.children.length; 
	let ele_tr, ele_td, ele_butt ; 
	let num_td =0; 
	
	
	let newBodyInner = "";
	
	
	const YESNO_NO_field = 3; 
	
	
	for(let g=0; g < num_tr; g++) {
		
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
	
	let inp2 = inp1.trim().replaceAll( "ae","a" );  
	
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


//----------------------------------------------
function isExtrRowChanged() {  // called by fun_require_mostFreqWordList,    sw_something_changed resetted by js_go_showWordList_lev2
	
	if (sw_somethingChanged) {
		extrRowBecause("0PrevChange");  
		return true; 
	}
	if (last_id_gruppi_sel_html_rowGroup_index_gr != id_gruppi_selected_groupIndex) {
		console.log("isExtrRowChanged ", "last_id_gruppi_sel_html_rowGroup_index_gr =" + last_id_gruppi_sel_html_rowGroup_index_gr + ", id_gruppi_selected_groupIndex =" + id_gruppi_selected_groupIndex + "<==")
		extrRowBecause("1groupIndex"); 
		sw_somethingChanged=true; 
		return true; 
	} 
	if (last_id_gruppi_iBegNum_html_rowGroup_beginNum != id_gruppi_SelectedBegin) {
		extrRowBecause("2beginNum");   
		sw_somethingChanged=true; 
		return true; 
	} 
	if (last_html_rowGroup_numRows  != id_gruppi_SelectedNumRow ) {
		extrRowBecause("3numRows");    
		sw_somethingChanged=true; 
		return true; 
	}   	
	if (last_sel_extrRow_freqWord_list  != html_sel_extrRow  ) {
		extrRowBecause("4selExtrRow"); 
		sw_somethingChanged=true; 
		return true; 
	}    

	return false; 
	
	function extrRowBecause(wh) {
		return;
		/**
		console.log("sw_somethingChanged true, reason=", wh) 
		console.log("1 last_id_gruppi_sel_html_rowGroup_index_gr    = ", last_id_gruppi_sel_html_rowGroup_index_gr,    "  id_gruppi_selected_groupIndex = ", id_gruppi_selected_groupIndex )		
		console.log("2 last_html_beginNum             = ", last_id_gruppi_iBegNum_html_rowGroup_beginNum,    "  id_gruppi_SelectedBegin = ", id_gruppi_SelectedBegin)
		console.log("3 last_html_numRows              = ", last_html_rowGroup_numRows,     "  id_gruppi_SelectedNumRow  = ", id_gruppi_SelectedNumRow )
		console.log("4 last_sel_extrRow_freqWord_list = ", last_sel_extrRow_freqWord_list, "  html_sel_extrRow       = ", html_sel_extrRow      )
		**/
	}	
	
} // end of isExtrRowChanged() 

//----------------------------------------------

function js_go_gotIxRowFromGroup( gostr1 ) {
	
	//console.log("js_go_gotIxRowFromGroup") 
	//  onchange_rowGroupSelectChange --> go go_passToJs_getIxRowFromGroup -->  js_go_gotIxRowFromGroup
	/*
	js ==> go : go_passToJs_getIxRowFromGroup( ""+id_gruppi_selected_groupIndex,  ""+id_gruppi_SelectedBegin, ""+id_gruppi_SelectedNumRow, "js_go_gotIxRowFromGroup")
	go ==> js :  
	outS1:= fmt.Sprintf( "inp,%d,%d,%d,rG_,%d,%s,%d,%d,ixR,%d,%d, %s",
				rowGrIndex, id_gruppi_SelectedBegin, id_gruppi_SelectedNumRow,
				rG.rG_ixSelGrOption, 
				rG.rG_group, 
				rG.rG_firstIxRowOfGr, 
				rG.rG_lastIxRowOfGr,
				ixRowBeg, ixRowEnd, 	
				inputTextRowSlice[  rG.rG_firstIxRowOfGr ].rRow1   )
	*/
	
	let col1 = gostr1.split(",")
		if ( (col1.length < 12) || ( (col1[0] != "inp") || (col1[4] != "gr") || (col1[9] != "ixr") )  ) {
		console.log("Errore1 in js_go_gotIxRowFromGroup (gostr1=", gostr1 , "\n\t non inizia con inp il formato deve essere ",  	
				`\n "inp,%d,%d,%d,rG_,%d,%s,%d,%d,ixR,%d,%d, %s",
				rowGrIndex, id_gruppi_SelectedBegin, id_gruppi_SelectedNumRow,
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
	x_rowGrIndex             = Number( col1[1] )
	x_id_gruppi_iBegNum_html_rowGroup_beginNum = Number( col1[2] )
	
	let id_gruppi_SelectedNumRow = Number(  col1[3] ); // Number(   col1[3] )
	
	x_rG_ixSelGrOption       = Number( col1[5] )
	x_rG_group               = Number( col1[6] ) 
	x_rG_firstIxRowOfGr      = Number( col1[7] ) 
	x_rG_lastIxRowOfGr       = Number( col1[8] )
	
	SAVE_fromIx_row          = Number( col1[10] ) ; 
	SAVE_toIx_row            = Number( col1[11] ) ; 
	
	let firstRowOfGroup  = col1.slice(12).join(",")
	
	if ((x_rowGrIndex != id_gruppi_selected_groupIndex) || 
			(x_id_gruppi_iBegNum_html_rowGroup_beginNum != id_gruppi_SelectedBegin ) || 
			(x_rowGrIndex != x_rG_ixSelGrOption)) {
		console.log("Errore2 in js_go_gotIxRowFromGroup (gostr1=", gostr1 , 
			"id_gruppi_selected_groupIndex oppure id_gruppi_SelectedBegin sono cambiati" ,
			" id_gruppi_selected_groupIndex =", id_gruppi_selected_groupIndex , " nuovo=", x_rowGrIndex, 
			" id_gruppi_SelectedBegin=", id_gruppi_SelectedBegin , " nuovo=", x_id_gruppi_iBegNum_html_rowGroup_beginNum,
			" x_rowGrIndex = ", x_rowGrIndex  , "  x_rG_ixSelGrOption=", x_rG_ixSelGrOption 			
			) 
		return; 		
	}	
	
	//console.log("  2 ... gotIxRowFromGroup") 
	
	getById("id_gruppi_sel"     ).selectedIndex = id_gruppi_selected_groupIndex;
		
	getById("id_gruppo_numTotRow1" ).innerHTML = Number(x_rG_lastIxRowOfGr) -  Number(x_rG_firstIxRowOfGr) + 1; 
	
			
	getById("id_gruppi_iNumRows").value      = id_gruppi_SelectedNumRow   
		
	//console.log("  3 ... gotIxRowFromGroup") 
	/**
	let beg1 = id_gruppi_SelectedBegin + x_rG_lastIxRowOfGr - 1; 
	let end1 = id_gruppi_SelectedBegin + rG.rG_firstIxRowOfGr - 1 +	id_gruppi_SelectedNumRow - 1; 	
	logColor("%%blue", "js_go_gotIxRowFromGroup ", "%%black", "  SAVE_fromIx_row=",  SAVE_fromIx_row, " XX  differenza ", " get_ixRowBeg=", beg1 )
	logColor("%%blue", "js_go_gotIxRowFromGroup ", "%%black", "  SAVE_toIx_row  =",  SAVE_toIx_row,   " XX  differenza ", " get_ixRowEnd=", end1 )
	console.log("  4 ... gotIxRowFromGroup") 
	**/
	
} // end of js_go_gotIxRowFromGroup 

//--------------------------------------------------
function get_ixRowBeg() { 	
	let beg1 = id_gruppi_SelectedBegin + x_rG_lastIxRowOfGr - 1; 
	return beg1;
}
function get_ixRowEnd() {
	let end1 = id_gruppi_SelectedBegin + rG.rG_firstIxRowOfGr - 1 +	id_gruppi_SelectedNumRow - 1; 	
	return end1 
}

//--------------

function go_run_js_build_rowGruppi( gruppi_option) {
	// run just after the reading of the input text 
	/*
	<option>13 file: soloParoleGoetheLista_A1.csv   (783 righe)</option>
	<option>14 Esempi da dizionario</option>
	<option>15 140 verbi irregolari</option>
	*/
	let optLine = gruppi_option.split("<option>");
	let opt1, optVV;
	listaGruppiTesto = ",";
	for(let v=0; v < optLine.length; v++) {
		opt1 = (""+optLine[v]).trim();
		if (opt1 == "") continue
		optVV = opt1.split(" ");
		if (optVV.length < 1 ) {continue;}
		listaGruppiTesto = listaGruppiTesto + trimLeftZero( optVV[0] ) +","
	}	
	//console.log("js_go_build_rowGruppi ", gruppi_option, "\nlistGruppiTesto=" + listaGruppiTesto);
		
	id_gruppi_selected_groupIndex = 0;  // will be updated from last file values 
	
	getById("id_gruppi_sel").innerHTML     = gruppi_option;  
	getById("id_gruppi_sel").selectedIndex = id_gruppi_selected_groupIndex;
	 
	//setLastValuesOfExtrRowChanged("js_go_build_rowGruppi");  // 2
	
} // end of  go_run_js_build_rowGruppi
//---------------------------------------------------
function go_run_js_read_valueFromLastRun( gostr1 ) {
	/*
	la funzione viene eseguita all'inizio, usando i dati del file last_mainpage_values2.txt letto dalla funzione GO   
		*** esempio di file letto **
		rG_ixSelGrOption=2, rG_group=3, rG_firstIxRowOfGr=13, rG_lastIxRowOfGr=23, 
		id_gruppi_selected_groupIndex=2, id_gruppi_SelectedBegin=1, id_gruppi_SelectedNumRow=11, 
		ixRowBeg=13, ixRowEnd=23, 
		w_fromWord=1, w_numWords=9999, 
		sel_extrRow=anyRow
	*/
	console.log("  go_read_valueFromLastRun ") 
	
	/**
	esempio di contenuto di gostr1 
	0:rG_ixSelGrOption, 1:rG_group, 2:rG_firstIxRowOfGr, 3:rG_lastIxRowOfGr,
	html	
	5:id_gruppi_selected_groupIndex, 6:id_gruppi_SelectedBegin, 7:id_gruppi_SelectedNumRow, 
	ix
	9:ixRowBeg, 10:ixRowEnd, 
	w
	12:w_fromWord, 13:w_numWords, 
	14:sel_extrRow=anyRow
	:row=
	16rwS
	**/	
	let col1 = gostr1.split(",")
	if (col1.length < 16) {
		console.log("errore1 in go_read_valueFromLastRun il numero di valori tra virgola (", col1.length,") < 16" ,   " gostr1=" + gostr1 );
		return	;	
	}	
	//logColor("%%blue", "go_read_valueFromLastRun =>", "%%black", gostr1)  ;
	
	//  1,2,0,0,html,1,2,28,ix,0,0,w,1,774,extrRow, :row=,Die Elemente
	
	x_rG_ixSelGrOption     	    	= Number( col1[ 0 ] )
	x_rG_group               		= Number( col1[ 1 ] ) 	
	x_rG_firstIxRowOfGr 			= Number( col1[ 2 ] )     
	x_rG_lastIxRowOfGr  			= Number( col1[ 3 ] )   
	// html col1[4]			
	id_gruppi_selected_groupIndex 	= Number( col1[ 5 ] ) 
	id_gruppi_SelectedBegin   		= Number( col1[ 6 ] )
	id_gruppi_SelectedNumRow  		= Number( col1[ 7 ] ) 			
	// x col1[8]
	SAVE_fromIx_row         		= Number( col1[ 9 ] )    
	SAVE_toIx_row              		= Number( col1[10 ] )       
	// w col1[11]	
	let word_fromWord          		= Number( col1[12 ] ) 	
	let word_numWords          		= Number( col1[13 ] )			
	let sel_extrRow            		= col1[ 14 ] 	
	let row                    		= col1.slice(16).join(",")
	/**
	console.log("  3 go_read_valueFromLastRun ") 
	console.log("id_gruppi_SelectedBegin = ", id_gruppi_SelectedBegin, " typeof ", typeof id_gruppi_SelectedBegin, " type number ",  typeof Number(id_gruppi_SelectedBegin))
	
	logColor("%%blue", "    ... id_gruppi_SelectedBegin=", " %%black",id_gruppi_SelectedBegin,"%%blue", " id_gruppi_SelectedNumRow=", "%%black", id_gruppi_SelectedNumRow,
			"%%blue"," SAVE_fromIx_row=", "%%black",SAVE_fromIx_row,  
			"%%blue"," SAVE_toIx_row=",   "%%black",SAVE_toIx_row ,  " XX num row=", (1 + SAVE_toIx_row - SAVE_fromIx_row) )  ;
	**/
	getById("id_gruppi_iNumRows"   ).value     		= id_gruppi_SelectedNumRow
	getById("id_gruppi_sel"        ).selectedIndex  = id_gruppi_selected_groupIndex;  // il contenuto dovrebbe rimanere invariato	
	getById("id_gruppi_iBegNum"    ).value     		= id_gruppi_SelectedBegin ;    // relativo all'inizio del gruppo // contenuto cambia solo se è variato il gruppo
	getById("id_gruppo_numTotRow1" ).innerHTML 		= x_rG_lastIxRowOfGr -  x_rG_firstIxRowOfGr + 1; // nella pagina html è il numero tra parentesi che si trova sotto la scelta del gruppo 
	getById( sel_extrRow           ).selected 		= "true"; 	
	getById("id_inpMaxNumWords"    ).value 			= word_numWords ; 
    getById("id_inpBegFreqWList"   ).value 			= word_fromWord ;
	//--------
	//console.log("  2 go_read_valueFromLastRun ") 
	
	let beg1 = id_gruppi_SelectedBegin + x_rG_firstIxRowOfGr - 1;    // indice della prima riga richiesta
	let end1 = id_gruppi_SelectedBegin + x_rG_firstIxRowOfGr - 1 +	id_gruppi_SelectedNumRow - 1; 	
	if (end1 > x_rG_lastIxRowOfGr) { end1 = x_rG_lastIxRowOfGr; id_gruppi_SelectedNumRow = end1 - beg1 + 1; }
	
	/**
	console.log("  2.2 go_read_valueFromLastRun ") 
	logColor("%%blue", "js_go_gotIxRowFromGroup ", "%%black", "  SAVE_fromIx_row=",  SAVE_fromIx_row, " XX  differenza ", " get_ixRowBeg=", beg1 )
		console.log("  2.3 go_read_valueFromLastRun ") 
	logColor("%%blue", "js_go_gotIxRowFromGroup ", "%%black", "  SAVE_toIx_row  =",  SAVE_toIx_row,   " XX  differenza ", " get_ixRowEnd=", end1 )
	**/
	//---
	
	//console.log("  3 ...go_read_valueFromLastRun") 
	
	onchange_rowGroupSelectChange(false,10); 
	
	//setLastValuesOfExtrRowChanged("js_go_valueFromLastValue"); // 1
		
} // end of js_go_valueFromLastValue()

//--------------------------------------------------
function onchange_rowGroupSelectChange(sw_newGr,where) {
	
	// <select ele_gruppi ></select>	
	let ele_gruppi = getById("id_gruppi_sel"      )
	let ele_begNum = getById("id_gruppi_iBegNum"  )
	let ele_numRow = getById("id_gruppi_iNumRows" )

	if (ele_gruppi.selectedIndex < 0) { ele_gruppi.selectedIndex = 0; }
		
	id_gruppi_selected_groupIndex     	= ele_gruppi.selectedIndex ; // indice gruppo   ( html id="id_gruppi_sel" ) 	
	id_gruppi_SelectedBegin   			= Number( ele_begNum.value);  // indica da dove si deve iniziare a leggere il gruppo  ( html id="id_gruppi_iBegNum" ) 	
	id_gruppi_SelectedNumRow  			= Number( ele_numRow.value);  // numero di righe richieste ( html id="id_gruppi_iNumRows" ) 	
	if (id_gruppi_SelectedBegin  < 1)	{id_gruppi_SelectedBegin  = 1; ele_begNum.value = 1; }
	if (id_gruppi_SelectedNumRow < 1) 	{id_gruppi_SelectedNumRow = 1; ele_numRow.value = 1; }
	
	//----------
	// if the group  changes, reset beginning and number of rows 
	if (sw_newGr) {
		id_gruppi_SelectedBegin  = 1;      // relativo all'inizio del gruppo 
		id_gruppi_SelectedNumRow = 200;
		ele_begNum.value 		 = id_gruppi_SelectedBegin
		ele_numRow.value 		 = id_gruppi_SelectedNumRow
	}
	if (last_id_gruppi_sel_html_rowGroup_index_gr     == "") {last_id_gruppi_sel_html_rowGroup_index_gr     = id_gruppi_selected_groupIndex; } 
	if (last_id_gruppi_iBegNum_html_rowGroup_beginNum == "") {last_id_gruppi_iBegNum_html_rowGroup_beginNum = id_gruppi_SelectedBegin;       } 
	if (last_html_rowGroup_numRows                    == "") {last_html_rowGroup_numRows                    = id_gruppi_SelectedNumRow;      }  
	
	let x2 				= getById("id_sel_2_extrRow");
    let i				= x2.selectedIndex;	
	html_sel_extrRow 	= x2.options[i].id; 
	last_sel_extrRow_freqWord_list  = html_sel_extrRow  ;
	
	go_passToJs_getIxRowFromGroup( ""+id_gruppi_selected_groupIndex,  ""+id_gruppi_SelectedBegin, ""+id_gruppi_SelectedNumRow, "js_go_gotIxRowFromGroup", js_parm, js_caller );
	
} // end of onchange_rowGroupSelectChange 


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
	let eleDiv = this1.parentElement;
	let eleTD  = eleDiv.parentElement; 
	let newDiv;
	if (eleTD.children.length >= 2) { 		// elimina div che permette variazione/immissione traduzione della riga 
		newDiv = eleTD.children[1]; 	
		newDiv.remove();
		return; 
	} 	
	let ele_tran = eleDiv.children[1];		// aggiunge div che permette variazione/immissione traduzione della riga 
	newDiv = document.createElement("div");
	newDiv.style.textAlign = "left";
	eleTD.appendChild(newDiv);
	
	let newInn=""
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
	
	
	
	let eleDiv0  = this1.parentElement;
	let eleDiv1  = eleDiv0.parentElement; 
	let eleTD    = eleDiv1.parentElement; 
	
	let eleLemmaTD   = eleTD.nextElementSibling
	let eleLemDiv1   = eleLemmaTD.children[0]
	let eleLemmaSpan = eleLemDiv1.children[0]
	if (eleLemmaSpan.innerHTML == "") return; 
	
	
	//if (eleTD.children.length < 1) { return; } 
	
	//console.log("%conclickDoubleWordTran", "color:red;"); console.log(" 1 eleTD=", eleTD.outerHTML)
	
	//if (eleTD.children.length >= 2) { return; } 
	if (eleDiv1.children.length >= 2) { return; } 
	
	let ele_tran = eleDiv0.children[1];
	
	const newDiv = document.createElement("div");
	newDiv.style.textAlign = "left";
	//eleTD.appendChild(newDiv);
	eleDiv1.appendChild(newDiv);
	//console.log(red("medio eleTD="), eleTD.outerHTML) ; 	
	let newInn=""
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
	let class01 = "c_size_1_line";
	let classNN = "c_size_nn_line";
	let divToResize = this1.parentElement.parentElement.parentElement
	let swClass = ( divToResize.classList.contains( class01 )	)
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
	let wLemmaList, wTranList, wLevelList, wParaList, wExampleList,  wIxLemmaList;	
	
	// =================
	//console.log("1 onclick_saveNewWordTran"); 
	if (this1 == null) return; 
	
	vertResizeWord(this1);
	
	let eleTR = this1.parentElement; 
	let eleTD;
	let numCellWrd = -1; 
	for(let z1=0; z1 < 10; z1++) {
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
		console.log("    eleTR=", eleTR.outerHTML); 		return;	
	}	
	
	//console.log("2 onclick_saveNewWordTran", " cellIndex=", numCellWrd, " TR outerHTML=", eleTR.outerHTML ); 
	
	let elePareTr = eleTR.parentElement; // tbody
	if (elePareTr == null) return; 
	
	let eleTr2;

	let word1, ixW2StudyLs, ix1, ixLemma; 
	
	let oldDivTran, eleOldTran, oldTranslation; 
	let newDivTran, eleNewTran, newTranslation; 
	let swChg=false;
	//let listNewTranIx = [];
	//let listNewTransla= [];
	let eleTD_0, eleTD_5; let eleTD0_val;
	let eleTD_6, ele_details, ele_summ, ele_tranLemma; 
	
	let newUp = 0; 
	
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
	
	for(let t1=0; t1 < elePareTr.children.length; t1++) {
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
	
		let eleTD_5div = eleTD_5.children[0];
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
				
				let eleTD_6div =  eleTD_6.children[0];  
				let ele_divSu = eleTD_6div.children[1]
				//console.log("  ele_divSu=", ele_divSu.outerHTML)
				if (ele_divSu) {
					let ele_tranLemma = ele_divSu.children[3]
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
	let eleDiv1 = this1.parentElement;  
	let eleDiv2 = eleDiv1.parentElement; 	
	let eleTDs  = eleDiv2.parentElement;
	let eleTR   = eleTDs.parentElement; 
	/*
	let eleTR = this1.parentElement; 
	for(let z1=0; z1 < 10; z1++) {
		if (eleTR.tagName == "TR") { break; } 
		eleTR = eleTR.parentElement; 
	} 
	**/
	if (eleTR.tagName != "TR") { return; } 	
	let elePareTr = eleTR.parentElement; // tbody
	let eleTr2;


	let word1, ixW2StudyLs, ix1, ixLemma; 
	
	let oldDivTran, eleOldTran, oldTranslation; 
	let newDivTran, eleNewTran, newTranslation; 
	let swChg=false;
	//let listNewTranIx = [];
	//let listNewTransla= [];
	let eleTD_0, eleTD_5; let eleTD0_val;
	let newUp = 0; 
	let ixRow2StudyLs, nfile, idRow, ixRow, origRow, tranRow, gruppoW; 
	let row0, cols;
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
	for(let t1=0; t1 < elePareTr.children.length; t1++) {
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
			[nfile, idRow, ixRow, origRow, tranRow, gruppoW] = cols.slice(0,6); 	
		} catch(e1) {	
			continue
		}
		if (oldTranslation.indexOf(tranRow)>=0) {  // non sono esattamente eguali, oldTranslation termina con <br>
			tranRow = newTranslation;
			rowToStudy_list[ixRow2StudyLs] = nfile + "|" + idRow + "|" + ixRow + "|" + origRow + "|" + tranRow + "|" + gruppoW;  
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
	//console.log('getById("id_inpWordFra") =' , getById("id_inpWordFra").outerHTML) 
    let aWord ="";  
	if (type==2) {
		aWord = word1; 	
	} else {		
		aWord     = getById("id_inpWordFra").value.trim();  
	}	
	if (aWord == "") {
		ele_wordList.innerHTML ='<span style="color:red;">manca la parola da cercare</span>';
		return;
	}	
	myPage01.style.display = "none"; 
	go_passToJs_thisWordRowList(aWord, ""+maxNumRow5, "js_go_showWrdRowList"); 
	
	**/
	
	let maxNumRow5 = 100; 
	
	let wordLista1 = "", wordLista2=""; 
	let ele_lista1 , ele_lista2;  	
	
	
	let id_wordLista1 = "idwS1_" + id1; 
	let id_wordLista2 = "idwS2_" + id1; 
	let id_maxNum     = "idwS3_" + id1; 
	if (getById(id_maxNum)) {
		let numVal=	Number( getById(id_maxNum).value);
		if (numVal > 0) maxNumRow5 = numVal; 	
	}
	
	ele_lista1 = getById(id_wordLista1);  
	ele_lista2 = getById(id_wordLista2);  
	if (ele_lista1) wordLista1 = ele_lista1.value;
	if (ele_lista2) wordLista2 = ele_lista2.value;	
	
	logColor("%%blue","onclickSelectWord2 --> CERCA PAROLE go_passToJs_someWordsRowList ", "%%black", " wordLista1=", wordLista1, " wordLista2=", wordLista2, " maxNumRow5=", maxNumRow5,  ); 
	//console.log("%cCERCA  PAROLA "+ wordLista1 + " "+ wordLista2 , "color: green;") 
		
	//get_first_row_tr_visible();  // memorizza la prima TR visibile delle frasi in cui si trova questa funzione 
	
	//myPage01.style.display = "none"; 
	go_passToJs_someWordsRowList(wordLista1, wordLista2, ""+maxNumRow5, "js_go_showWrdRowList", js_parm, js_caller); 	
	
	
	
} // end of onclickSelectWord2

//=========================================================
function back_from_listaRighe(   page1, page2,page3, page4) {
	 
	get_first_row_tr_visible() 
	 
	onclick_jumpFromTo1_2Page( page1, page2,page3, page4)
	
} // end of back_fromn_listaRighe

//-----------------------------------

function get_first_row_tr_visible() {	

	//logColor("%%red", "get_first_row_tr_visible ", "%%black", " sw_rowListFrom_onclick=",  sw_rowListFrom_onclick)
	
	if (sw_rowListFrom_onclick == false) return;    // aggiorna "id_gruppi_iBegNum" solo se la richiesta arriva da onclick_rowList
	
	//let divConTab1 = getById("id_div_tabSub"); 
	let container  = getById("id_div_tabSub");     // elemento scrollabile che contiene la tabella
	let eleTabBody = getById("id_tabSub_tbody");
	let eleTrList  = eleTabBody.children;          //righe tabella che possono scomparire/apparire qusndo  il cursore sposta la vista 

	
	const containerTop = container.scrollTop;
	const containerBottom = containerTop + container.clientHeight;
	let first_visible_tr_id = -1;
	for (let i = 0; i < eleTrList.length; i++) {
		let ele = eleTrList[i]
		const eleTop = ele.offsetTop;
		const eleBottom = eleTop + ele.clientHeight;
		if (eleTop >= containerTop && eleBottom <= containerBottom) { // The element is fully visible in the container
			 first_visible_tr_id = i;
			 //console.log("%cget_first_row_tr_visible ","color:red;"); console.log("first_visible_tr_id =i=", i, " ele=\n",ele.innerHTML) 
			 break; 
		}
	} // end for i 
	if (first_visible_tr_id < 0) {
		//console.log("%cnessuna tr è visibile", "color:red;")
	} else {
		//console.log("%cLa prima tr visibile è " + first_visible_tr_id, "color:blue;")
		let eleFromNum=getById("id_gruppi_iBegNum");  
		//console.log('get_first_row_tr_visible 1 getById("id_gruppi_iBegNum")=', getById("id_gruppi_iBegNum").value) 
		let newSt = parseInt(eleFromNum.value) + first_visible_tr_id - 1;  	
		if (newSt < 0) newSt=0; 
		eleFromNum.value = newSt;               // set on first page
		
		//console.log('get_first_row_tr_visible 2 getById("id_gruppi_iBegNum")=', getById("id_gruppi_iBegNum").value) 

		onchange_rowGroupSelectChange(false,11);
	}
	
	
} // end of get_first_row_tr_visible	

//-----------------------------------
function back_from_listaParole(  page1, page2,page3, page4) {
	 
	get_first_word_tr_visible(); 
	 
	onclick_jumpFromTo1_2Page( page1, page2,page3, page4)
	
} // end of back_fromn_listaRighe
//-------------------------
/**
// ask before exiting 
window.onbeforeunload = function(){	
	return 'Are you sure you want to leave?';
};
**/

//-----------------------------------------
let wordButton = 0;  
let word_sw_all = false;
//--------------------------------------------
function get_first_word_tr_visible() {	

	if (word_sw_all == false) return;
	
	//logColor("%%blue"," get_first_word_tr_visible", "%%black", "id_lastWSort=",  getById("id_lastWSort").innerHTML)
	
	if (getById("id_lastWSort").innerHTML != "1") return; 
	// only if the page has not been sorted we can set the number of word to skip
 	
	let container  = getById("id_div_word_tabSub");  // elemento scrollabile
	let eleTabBody = getById("idTableWordList_tbody");
	let eleTrList  = eleTabBody.children;      //righe tabella che possono scomparire/apparire qusndo  il cursore sposta la vista 
	
	
	const containerTop = container.scrollTop;
	const containerBottom = containerTop + container.clientHeight;
	let first_visible_tr_id = -1;
	for (let i = 0; i < eleTrList.length; i++) {
		let ele = eleTrList[i]
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
			/**
				<td class="fixWordTdWidth4" style="text-align:center;font-size:0.8em;font-weight:100;">
					<span onmouseover="mouseOverWord(this)" onmouseout="mouseOutWord(this)"  >§one-numTR§</span>					
					<span style="display:none;">
						<span>§one-word1§</span>    
						<span>§ixW2StudyLs§</span>
						<span>§one-ix1§</span>
						<span>§one-ixixLemma§</span>	
						<span>§one-ixLemma§</span>	
					</span>	
			</td>  
			**/
		let ele = eleTrList[first_visible_tr_id]
		let eleTD0 = ele.children[0];
		let eleNumWord = eleTD0.children[1].children[2]
		let numWord
		if (eleNumWord) {
			numWord= Number(eleNumWord.innerHTML)
			//logColor("%%blue","    numero parola=", "%%black",numWord, " =>", eleTD0.innerHTML)  		
			
			let eleInpBeg = getById("id_inpBegFreqWList")
			
			//console.log("     prima eleInpBeg.value=",  eleInpBeg.value, "   eleNum.value=", 	getById("id_inpMaxNumWords").value ) 
			
			eleInpBeg.value = Number(eleInpBeg.value) + numWord - 1;  
			
			let eleNum 	  =	getById("id_inpMaxNumWords" )
			eleNum.value  = Number( eleNum.value ) - numWord + 1;	 // the number of row can be greater than the number of the required words (a word can appear more than on time)
			//eleNum.value  = eleTrList.length - numWord + 1	
			//console.log("     dopo eleInpBeg.value=",  eleInpBeg.value, "   eleNum.value=", eleNum.value ) 
		}
	}

} // end of get_first_word_tr_visible	
//--------------------------------------------
function onclick_word_known2(sIxWord, this1) {
	let ixWord = 0, yes_not_len1;
    try {
        ixWord = parseInt(sIxWord);
    } catch (err) {}	
	
	let sw_yesNo = (this1.innerHTML != YES1)
  
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
	getById("id_buttLearnNumW").innerHTML = numWordsKnownChanged;
	
	go_passToJs_word_known2(""+ixWord, yes_not_len1, "js_go_word_known2", js_parm, js_caller); // ask 'go' to update yes/no word known ctr  
	
} // end of onclick_word_know_yes		

//----------------------------------------
function js_go_word_known2(str1) {
	//console.log("js_go_word=", str1);   
	
	if (numWordsKnownChanged == 1) {		
		getById("id_buttLearnDiv").style.display = "block";
	} 	
	if (numWordsKnownChanged > MAX_NUM_WORD_LEARN) {		
		onclick_write_words_to_learn();
	} 
	
}// end of js_go_word_known 	

//-------------------------------------------------------------
function js_go_file_words_to_learn_written( str1 ) {
	//getById("id_w_to_learn_written").innerHTML = str1 ;	
	numWordsKnownChanged = 0; 
	getById("id_buttLearnNumW").innerHTML = numWordsKnownChanged;
	getById("id_buttLearnDiv").style.display = "none";
}
//-----------------------------------------------------------------

function onclick_hideShowWordTran(this1) {	

	let swEle = this1.previousElementSibling;
	
	//console.log("%conclick_hideShowWordTran", "color:red;"); console.log(" swEle=", swEle)

	// nasconde o mostra la traduzione di tutte le parole 
	let eleBody = getById("idTableWordList_tbody")
	let rows1 = eleBody.rows; 	
	//--------
	let lemmaCell, eleTran, visib;		
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
	for (let g=0; g < rows1.length; g++ ) {	
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

let hideShowTranExample_H1 = "mostra traduzione ed esempi"; 
let hideShowTranExample_H2 = "nascondi traduzione ed esempi"; 
//----------------
function onclick_hideShowTranExample(this1) {
	//console.log("onclick_hideShowTranExample")
	let display1;
	let lemmaCell, eleTran;
	let buttHeader = this1.innerHTML; 
	if (buttHeader == hideShowTranExample_H1) {
		display1="block";
		this1.innerHTML = hideShowTranExample_H2
	} else {
		display1="none";
		this1.innerHTML = hideShowTranExample_H1
	}	
	let eleBody = getById("idTableWordList_tbody")
	let rows1 = eleBody.rows; 
	//--------------------------
	for (let g=0; g < rows1.length; g++ ) {	
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
	let eleBody = getById("idTableWordList_tbody")
	let rows1 = eleBody.rows; 		
	let class01 = "c_size_1_line";
	let classNN = "c_size_nn_line";
	if (rows1.length < 1) return
	let lemmaCell = rows1[0].cells[NUM_CELL_LEMMA].children[0]
	let swClass = ( lemmaCell.classList.contains( class01 )	)
	//----------------------
	if (swClass) {		
		this1.innerHTML = "mostra solo 2 righe"
		for (let g=0; g < rows1.length; g++ ) {	
			lemmaCell = rows1[g].cells[NUM_CELL_LEMMA].children[0]
			lemmaCell.classList.remove( class01 );
			lemmaCell.classList.add(    classNN );		
		} // end for g
	} else {		
		this1.innerHTML = "mostra tutte le righe"
		for (let g=0; g < rows1.length; g++ ) {	
			lemmaCell = rows1[g].cells[NUM_CELL_LEMMA].children[0]
			lemmaCell.classList.remove( classNN );
			lemmaCell.classList.add(    class01 );				
		}
	} 	
} // end of TOGLIonclick_vertResizeLemma 

//-----------------------------------------------------------------

function onclick_listNoParadigmaLemma(this1) {	

	let downfilename= "parole_senza_paradigma.txt" 
	let outText = "Lista Parole senza paradigma " + "\n\n" ;
	
	
	let eleBody = getById("idTableWordList_tbody")
	let rows1 = eleBody.rows; 	
	//--------
	let lemmaCell, eleTran, elePara, eleDivSup;	
	//----------------------
	for (let g=0; g < rows1.length; g++ ) {	
		try {
			lemmaCell = rows1[g].cells[NUM_CELL_LEMMA]
			eleDivSup = lemmaCell.children[0].children[1];
			elePara = eleDivSup.children[1]; 
			
			if (elePara.innerHTML == "") {
				outText += eleDivSup.children[0].innerHTML.replaceAll("<b>","").replaceAll("</b>","") + "  | \n";				
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
	let g;
	let downfilename= "Lemma_senza_traduzione.txt" 
	let outText = "Lista Lemma senza Traduzione"+ "\n" ;
	if (swNoTran == false) {
		downfilename= "lista_tutti_lemma.txt" 
	    outText = "Lista di tutti i lemma"+ "\n" ;
	} 
	
	let eleBody = getById("idTableWordList_tbody")
	let rows1 = eleBody.rows; 	
	//--------
	let lemmaCell, eleTran, elePara, eleDivSup;	
	//----------------------
	let noTranL = []
	let tranS;
	for (let g=0; g < rows1.length; g++ ) {	
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
	let preW="", wo="";
	let nn=0;
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
	
	let downfilename= "lista_parole.txt" 
	let outText = "Lista Parole"+ "\n" ;	
	let eleBody = getById("idTableWordList_tbody")
	let rows1 = eleBody.rows; 	
	//--------
	let wordCell, lemmaCell, eleTran, elePara, eleDivSup, eleWrd;	
	let lemmaT, wordT;
	//----------------------
	let noTranL = []
	let tranS;
	for (let g=0; g < rows1.length; g++ ) {	
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
	let preW="", wo="";
	let nn=0;
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

    let element = document.createElement('a');

    element.setAttribute('href', 'data:text/plain;charset=utf-8,' +
        encodeURIComponent(text));
   
    element.setAttribute('download', filename);

    element.style.display = "none";
    document.body.appendChild(element);
    element.click();
    document.body.removeChild(element);
}

//--------------------------------------------------
function onclick_split_newText(this1) {
		let eleTxt = getById("id_newText");
		let str1 = (""+eleTxt.value).trim();
		if (str1.length == 0) return;
		/**
		console.log("punto ==> ",  str1.replaceAll(".",".\n") )
		console.log("punto? ==> ",  str1.replaceAll("?","?\n") )
			console.log("punto! ==> ",  str1.replaceAll("!","!\n") )		
		**/
		str1 = str1.replaceAll(".",".\n").replaceAll("?","?\n").replaceAll("!","!\n").replaceAll(";",";\n").replaceAll("\n\n","\n");
		
		//console.log("tutti ==> ",  str1)
			
		eleTxt.value = str1;	
		this1.nextElementSibling.style.display = "inline-block"

		
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
	for(let g=0; g < 5; g++) {   
		if (num1.length < 2) { break;}
		if (num1.substr(0,1) == "0") num1 = num1.substring(1); 
	}
	return num1;
}	
//--------------
function onclick_get_newText(this1) { 			
	let eleTxt = getById("id_newText");
	let str1 = (""+eleTxt.value).trim();
	if (str1.length == 0) return;			
	let ele_id1 = getById("id_newTxt_id")
	let ele_title1 = getById("id_newTxt_title"); 
	let id1 = ele_id1.value.trim();
	let title1 = ele_title1.value;	
	
	let msg1="";
	if (id1=="") {msg1=" manca identificativo"; }
	if (title1=="") { msg1+=" manca titolo del testo"; }
	if (id1 != "") {
		id1 = trimLeftZero( id1 );
		if (listaGruppiTesto.indexOf( "," + id1 + ",") >=0 ) {
			msg1 += "gruppo " + id1 + " non può essere utilizzato, è già esistente";   
		}	
	} 	
	if (msg1 != "") msg1="<br>errore: " + msg1;	
		
	getById("id_newTxtMsg1").innerHTML = msg1;
	if (msg1 != "") return;
	
	let righe=str1.split("\n");
	let riga;
	let outText = id1+"_0|O|file: " + title1 ;
	outText += "\n" + id1+"_0|T|file: " + title1 ;
	for (let f=0; f < righe.length; f++) {
		riga = "\n" + id1+ "_" + (f+1) + "|O|" + righe[f].trim();
		outText += riga;
	}	
	eleTxt.value = outText;
	this1.style.display = "none";
	//console.log(" onclick_get_newText ", " --> go_write_new_row_dictionary ", outText)
	
	go_write_new_row_dictionary( outText, "js_go_new_row_written","","")
	
	
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
	let eleTxt = getById("id_newText").value = ""
	let ele_id1 = getById("id_newTxt_id")
	let ele_title1 = getById("id_newTxt_title"); 
	str1 = "il nuovo testo con identificativo " + ele_id1.value + " è stato accettato";
	//let msg1 = "chiudi e riesegui l'applicazione"
	getById("id_newTxtMsg0").innerHTML = "<br>" + str1;
	getById("id_newTxtMsg1").innerHTML = "";
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
	let id1 = this1.id;
	if (id1 == undefined) return 
	let idNum = id1.replaceAll("idc_","idt_"); 
	let eleTran = getById(idNum); 
	if (eleTran == undefined) return;
	eleTran.style.display = "block";			 
} // end of onmouseOverRow
//-----------------------------------------------------------	
function onmouseOutRow(this1) {   // allontanando il mouse dalla riga originale, nasconde la traduzione (a meno che non sia attivo il tasto show T.  
	if (sw_mouseoverActiv == false) return;	
	let id1 = this1.id
	if (id1 == undefined) return 
	let idNum = id1.replaceAll("idc_","idt_"); 
	let eleTran = getById(idNum); 
	if (eleTran == undefined) return;
	// sarebbe naturale rimettere display none, ma fintanto che esistono i pulsanti di show/hide transl. forzo lo stato dettato da questi 
	let idNumButtT = id1.replaceAll("idc_","idbT_");
	if (idNumButtT == undefined) return;
	let eleTranButt = getById(idNumButtT); 
	let eleTbutCh   = eleTranButt.children[0]; 
	if (eleTbutCh == undefined) return;
	eleTran.style.display = eleTbutCh.style.display;			 
} // end of onmouseOutRow
//---------------------------------------------------------------------		

function onmouseOverRow2( num1 ) {  // toccando la riga originale, rende visibile la traduzione  
	
	let idNum = "idt_" + num1; 
	let eleTran = getById(idNum); 
	if (eleTran == undefined) return;
	eleTran.style.display = "block";	
	
} // end of onmouseOverRow
//-----------------------------------------------------------	
function onmouseOutRow2( num1 ) {   // allontanando il mouse dalla riga originale, nasconde la traduzione (a meno che non sia attivo il tasto show T.  
	let idNum = "idt_" + num1; 
	let eleTran = getById(idNum); 
	if (eleTran == undefined) return;
	eleTran.style.display = "none";	 
	
} // end of onmouseOutRow
//---------------------------------------------------------------------		

//----------------------------------------------------------
function onclick_changeTraduzione(this1 ) {	
	
	let eleTD0  = this1.parentElement;
	
	let eleLemmaTD = eleTD0.nextElementSibling;
	
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
		
	let eleWordTD  = eleTD0.previousElementSibling;
		
	let eleTD    =  eleWordTD; 
		
	let eleLemDiv1   = eleLemmaTD.children[0]
	let eleLemmaSpan = eleLemDiv1.children[0]
	if (eleLemmaSpan.innerHTML == "") return; 	
		
	if (eleLemmaTD.children.length >= 2) { return; } 
	
	let ele_oldTran = eleLemmaTD.children[0].children[1].children[1];
	const newDiv = document.createElement("div");
	newDiv.style.textAlign = "left";
	
	eleLemmaTD.appendChild(newDiv);
	let newInn=""
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
	
	
	let wLemmaList, wTranList, wLevelList, wParaList, wExampleList,  wIxLemmaList;	
	let wordx, ix12, nrow, totExtrRow2,  wLemma1, wordTran, uLearnedYN ;
	let word1, ixW2StudyLs, ix1, ixLemma; 	
	let oldDivTran, eleOldTran, oldTranslation; 
	let newDivTran, eleNewTran, newTranslation; 
	let swChg=false;
	let eleTD_0, eleTD_5; let eleTD0_val;
	let eleTD_6, ele_details, ele_summ, ele_tranLemma; 
	
	let newUp = 0; 
	
	let div_onsave00 = this1.parentElement; 	
	let divAdded00   = div_onsave00.parentElement;  
	let eleTdLemma00 = divAdded00.parentElement; 
	if (eleTdLemma00.tagName != "TD") { 
		console.log("%cERRORE eleTdLemma.tagName not equal TD =" + eleTdLemma00.tagName );
		return;
	} 	
	let eleTR = eleTdLemma00.parentElement; 	
	if (eleTR.tagName != "TR") { return; }
	let elePareTr = eleTR.parentElement; // tbody
	if (elePareTr == null) return; 
	
	//----------------------
	// cerca tutti i lemma con traduzione pendente  
	
	let tranToUpdList = [];
	let bodyChild = elePareTr.children
	for (let p1=0; p1 < bodyChild.length; p1++) {
		let eleTr0 = bodyChild[p1];
		let oneTdLemma = eleTr0.children[NUM_CELL_LEMMA]; 		
		if (oneTdLemma.children.length < 2) { continue; }	
		let oneNewLemTran = oneTdLemma.children[1]; 
		let oneNewLemTranInner = oneNewLemTran.children[1].innerHTML;
		if (oneNewLemTranInner == "") { continue; }	
		tranToUpdList.push( [ oneTdLemma.children[0].children[0].innerHTML, oneNewLemTranInner, oneNewLemTran ] ); 		
	} // end for p1	
	//-------------------
	for(let t2=0; t2 < tranToUpdList.length; t2++) {
		let toUpdLemma   = tranToUpdList[t2][0];	
		let toUpdTran    = tranToUpdList[t2][1];	
		let addedLivTran = tranToUpdList[t2][2];
		for (let p2=0; p2 < bodyChild.length; p2++) {
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
		let eleTdLemma = eleTr2.children[ NUM_CELL_LEMMA] ;
		let ele_target_lemmaName = eleTdLemma.children[0].children[0].innerHTML; 			
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

		let ele_target_lemTran = eleTdLemma.children[0].children[1].children[1]; 		
		let ele_target_tran2   = eleTdLemma.children[0].children[2].children[0].children[0]; 
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
	let class01 = "c_size_1_line";
	let classNN = "c_size_nn_line";
	let divToResize = eleTdLemma.children[0] ; 
	let swClass = ( divToResize.classList.contains( class01 )	)
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

	let wordTR = prototype_word_tr_tts.replaceAll("§1§", z3).replaceAll("§4txt§", parola1).replaceAll("§4maxNumRow§", "" + maxNumRow).
	replaceAll("§ttsWtxt§", paro_tts).replaceAll("§8numfrasi§", paro_nFrasi).replaceAll("§6tran§", trad1).
	replaceAll("§6alllemmaWord§", all_lemma_for_thisWord);
	return wordTR;

} // end of getWord_tr2 

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

function primaLeRigheGruppoRichiesto(inpBegRow, inpEndRow, rowToStudy_list) {	
	
	let nfile, idRow, ixRow, ixRowS, origRow ; 			
	let tranRow; 	
	let ixRow2StudyLs;
	
	let newRowList1 =[];
	let newRowList2 =[];
	let row0, cols;
	//------------------------
	for (let i1 = 0; i1 < rowToStudy_list.length; i1++) {	
		
		row0 = rowToStudy_list[i1].trim();	
		//if (i1 < 10) console.log("primaLeRighe... i1=", i1, " row0=", row0)	
		if (row0=="") 	  { continue;}	
		if (row0 == "\n") { continue; }
		
		row0 = i1 + "|" + row0 + "|||||"; 
		
		newRowList1.push(row0);
		
		/**
		continue;
		
		cols = row0.split("|");		
 		try{ 
			[ixRow2StudyLs, nfile, idRow, ixRowS, origRow, tranRow] = cols.slice(0,6); 
		} catch(e1) {
			console.log("showRowsAndTranButton () o=", i1, " rowToStudy_list[i1]=", rowToStudy_list[i1],
				" row0=", row0, " cols=", cols, "\n\t XXX  error ", e1); 			
		}	
		
		ixRow = parseInt(ixRowS)	
		if ((ixRow >= inpBegRow) && (ixRow <= inpEndRow)) {		
			newRowList1.push(row0);				
			if (i1 < 10) console.log("  ixRow="+ ixRow + ", 1 caricato row")    
		} else {
			newRowList2.push(row0);			
			if (i1 < 10) console.log("  ixRow="+ ixRow + ", 2 caricato row")    
		}	
		**/
    }  // end for i1
	
	return newRowList1.concat(newRowList2);

} // end of primaLeRigheGruppoRichiesto	
//----------------------------------------------
function add_otherGrafToWordList(word_to_underline_list) {
	// duplicare la parola con ortografia alternativa, es. ue= u con umlaut;  bisogna spostare questa logica in 'go'  
	let newList=[], word1, word2;
	for(let hx=0; hx < word_to_underline_list.length; hx++) {  
		word1 = word_to_underline_list;
		newList.push(word1)	
	}
	return newList
}
//---------------------------------------
function showRowsAndTranButton(wh) {	
 
	//console.log("showRowsAndTranButton  (wh=",wh)	
	let showList = ''    ;    
	let riga;		
	let row0; 
	let nfile, idRow, ixRow, origRow, gruppo; 
	
	string_tr_xx = "\n" + prototype_tr_tts; 
		
	//--------------	
	let txt1p, text_tts, tranRow; 
	let first= -1, last=-1;
	let visib;
	let ixRow2StudyLs;
	let indexGroup= Number( getById("id_gruppi_sel").selectedIndex );


    //console.log("myPage01=", myPage01.style.display,   ",  myPage04=", myPage04.style.display);   
	
	let inpBegRow = SAVE_fromIx_row;	
	let numRows   = id_gruppi_SelectedNumRow
	let inpEndRow = inpBegRow+numRows-1;
	
	//console.log("showRowsAndT 1"); 
	
	//------------------------------------
	let newRowList = primaLeRigheGruppoRichiesto(inpBegRow, inpEndRow, rowToStudy_list);
	//------------------------------
	let iNumTr =0;
	let idRow1, idRow2; 
	let PREF_MARKER = ":PREF:";
	//----------------------------------------------------
	//console.log("showRowsAndT 2"); 
	for (let i = 0; i < newRowList.length; i++) {
		one_newRowList(i);	
	} // end for i
	
	//console.log("showRowsAndT 3"); 
	//---------------------------------------------------
	eleTabSub_tbody.innerHTML = showList;		
	if ( (last - first) > 0) {
		let eleF = getById("b1_" + first);
		let eleT = getById("b2_" + last);
		onclick_tts_arrowFromIx(eleF, first, 5);
		onclick_tts_arrowToIx(  eleT, last , "3showRowsAndTranButton" );		
	}	
	onclick_jumpFromToPage( myPage04,0, myPage05);  
		
	eleTabSub_tbody.parentElement.parentElement.scrollTop = 0;
	
	//-----------------------------------------------------------------
	function one_newRowList(i) { 	
		iNumTr = i+1;
		row0 = newRowList[i].trim() + "||||||"; 		
		//if (i < 10) { console.log( " one_newRowList[]=", row0)  } 
		let cols = row0.split("|");
 		
		[ixRow2StudyLs, nfile, idRow, ixRow, origRow, tranRow, gruppo] = cols.slice(0,7); 
		
		
		if (origRow == ""       ) { return; } 
		if (origRow == undefined) { return; } 	
		
		origRow = origRow.trim(); 
		txt1p = origRow;
		
		let unaparola; 
		let class_targList = [ "c_wordTarg", "c_wordTarg2" ];
		let type=0;
		//-------
		for(let hx=0; hx < word_to_underline_list.length; hx++) {  
			unaparola = word_to_underline_list[hx].trim();
			if (unaparola == "") { continue; }
			if (unaparola == PREF_MARKER) {type=1;  continue; } 			
			txt1p = evidenzia(unaparola, class_targList[type], txt1p); 			
		}
		//---	
		txt1p = txt1p.replaceAll("§§", "");  // have beewn addded in function evidenzia
				
		text_tts = "";
		visib=""; 
		if (origRow == "") {
			visib = "visibility: hidden;"; 			
		} 	
		
		
		if (indexGroup == gruppo) {    // if ((ixRow >= inpBegRow) && (ixRow <= inpEndRow)) {
			nfile = 1;
		} else {
			nfile = 2;	
		} 
		let idro1 =idRow.split("(");
		if (idro1.length<2) {
			idRow1 = idRow; idRow2 = "";
		} else {
			idRow1 = idro1[0];  idRow2 = idro1[1]; 
		}
		let numBeg = inpBegRow + i; 
		//-------------	
		let txt1p_n   =   txt1p;
		let tranRow_n = tranRow;
        riga = string_tr_xx.replaceAll("§1§", iNumTr).
			replaceAll("§ixRow2StudyLs§"  , ""+ixRow2StudyLs).
			replaceAll("§4txt§"  , txt1p_n).
			replaceAll("§5txt§"  , tranRow_n).
			replaceAll("§ttstxt§", text_tts).
			replaceAll("§6id§"   , idRow1.replace(" "," - ") ).
			replaceAll("§6id2§"   ,idRow2).			
			replaceAll("§1beg§"  , numBeg).
			replaceAll("§6ix§"   , ixRow).
			replaceAll("§nfile§" , nfile). 
			replaceAll("§visib§" , visib);
		
		if (first < 0) first = iNumTr;
		last = iNumTr;		
		showList    += riga + "\n";	
		
		//if (i < 10) console.log("           one_newRowList ", riga) 
		
	} // end of one_newRowList
	//---------------------------
  
	
} // end of showRowsAndTranButton

//------------------------------------------------
// var a1        = JSON.stringify([12,34,56]);
// var a12_array = fromJsonStringToArray(a1)
//-------------------------------------------------
function fromJsonStringToArray( str1 ) {
	if (str1 == undefined) return [""]; 
	str1 = ""+str1;
	if (str1.length < 1) return [""]
	if ( str1.substr(0,1) == "[" ) {
		return JSON.parse( str1);
	}	
	return [ str1 ]
}
//--------------------------------------
// gestione_lemma
//----------------------------------------------------------	
	
function js_go_showWordList_lev2(wordListStr00, json_parmStr, js_caller, goFunc) {
	/**
	logColor("%%red;font-size:1.2em;font-weight:bold;", "js_go_showWordList_lev2 ", 
		"%%black", " wordListStr00.length=", wordListStr00.length, " json_parmStr=", json_parmStr, " js_caller=", js_caller)
	**/
	// parametri da pagina html ('anyRow','toBeLearned'),  ('anyRow','allWords') 
	
	/*
												row:= xWordAlpha.uWord2 + separ1 + xWordAlpha.uWord0 + separ1 +
		g04_bind_go_between_wordList_V3.go => 				"ix" + separ1 + 
		g04_bind_go_between_wordList_V3.go => 				strconv.Itoa(xWordAlpha.uIxUnW_fr) + separ1 + 
		g04_bind_go_between_wordList_V3.go => 				strconv.Itoa(xWordAlpha.uTotRow)   + separ1 +
		g04_bind_go_between_wordList_V3.go => 				lastLemma.leLemma            	+ separ1 + 
		g04_bind_go_between_wordList_V3.go => 				lastLemma.leTran 				+ separ1 +  
		g04_bind_go_between_wordList_V3.go => 				separ1 							+  
		g04_bind_go_between_wordList_V3.go => 				lastLemma.lePara	 			+ separ1 +  
		g04_bind_go_between_wordList_V3.go => 				lastLemma.leExample	 			+ separ1 +  
		g04_bind_go_between_wordList_V3.go => 				strconv.Itoa(totNumRow) 		+ separ1 +  			
		g04_bind_go_between_wordList_V3.go => 				xWordAlpha.uLearnedYN              + separ1 + 			
		g04_bind_go_between_wordList_V3.go => 				"ixLemma" + separ1 + strconv.Itoa(ix2) + separ1 +  	
		g04_bind_go_between_wordList_V3.go => 				endOfLine 	
		g04_bind_go_between_wordList_V3.go => 		//fmt.Println(green("word_to_row "), row )		
	*/	
	
	let js_parmArray = fromJsonStringToArray(json_parmStr)
	
	//console.log("json_parmStr=", json_parmStr, "  js_parmArray=",  js_parmArray,  " len=",   js_parmArray.length)
	
	let numButton   = ""; 
	let anyrow      = ""; 
	let toBeLearned = ""; 
	let tagSel = "";
	let pp=999;
	for(let p=0; p < js_parmArray.length; p++) {
		//console.log( "	js_parmArray =", js_parmArray[p], "   tagSel=", tagSel)
		if (js_parmArray[p] == "tag") { pp=p; if ((p+1) < js_parmArray.length) tagSel = js_parmArray[p+1]; break; }
	}
	if (pp > 0) numButton   = Number(js_parmArray[0]) ;  
	if (pp > 1) anyrow      = js_parmArray[1];  
	if (pp > 2) toBeLearned = js_parmArray[2]; 
	if (numButton == "") {
		if (tagSel == "RW3") numButton = 2;   
	}
	
	let swAllWords = (toBeLearned == 'allWords'); 
	
	wordButton = numButton; 
	word_sw_all = swAllWords;
	
	//logColor("js_go_showWordList_lev2", "%%green", "js_go_showWordList_lev2", "%%black", " json_parmStr=", json_parmStr, " js_caller=", js_caller, " goFunc=", goFunc)
	
	//console.log("    ", "wordListStr.length=",wordListStr00.length  ," numButton=", numButton) ;
	
	//let le2=200; if (wordListStr00.length > le2) console.log("   wordListStr=",wordListStr00.substring(0,le2) + " ..."); else  console.log("   wordListStr=",wordListStr00 );   

	// numButton=1 default ==> from onclick most frequent word list  
	// numButton=2         ==> from onclick BetweenWordList or prefix wordlist   
	// numButton=3         ==> from onclick Lemma word list   
	// numButton=5         ==> from onclick Lemma list   
	// numButton=0         ==> from word list from word, lemma, ?   
	if (numButton==1 ) {
		sw_somethingChanged = false; 
	} 
	
	word_to_underline_list = []
	let wordListStr = wordListStr00.trim();
	let len = wordListStr.length	

	if (wordListStr.substring(len-1) == ";") { len = len - 1; wordListStr = wordListStr.substring(0, len) }
	if (wordListStr.substring(len-1) == ";") { len = len - 1; wordListStr = wordListStr.substring(0, len) }
		
    // triggered by go func (  go _ passToJs_wordList )
    if (wordListStr == undefined) {
        //console.log("js_showWordList: parameter is undefined");
		onclick_jumpFromToPage( myPage02,myPage03,  myPage01); 
        return;
    }
	
	//console.log(" 2 function js_go_showWordList_lev2 2 ")
	
    if (wordListStr == "") {
        //console.log("js_showWordList: parameter is empty");
		if (numButton==1) { errorNoWord1()}
		onclick_jumpFromToPage( myPage02,myPage03,  myPage01); 
        return;
    }
	
	myPage01.style.display = "none"; 
	
	//console.log("3 function js_go_showWordList_lev2 3 ")

	let wLemmaListU, wTranListU, wLevelListU,	wParaListU, wExampleListU, wIxLemmaListU;   
	let wLemmaList,  wTranList,  wLevelList,    wParaList,  wExampleList , wIxLemmaList ;
	let word2, word0, ixUnW2, totRow2, totExtrRow2 
	let uLearnedYN;
	let chk_ix, chk_ixLemma;
	
    let wordToStudy_listStr = wordListStr.split( endOfLine );	
	
	//console.log("3 function js_go_showWordList_lev2 3 ")	
	/*
	wordListStr=
						 0                 1         2     3     4        5              6    7   8   9  10   11 
	genannt.genannt      ;.  genannt    ;.505;.    145;. 123;. 123;. nennen  ;.  nome    ;.   ;.  ;.  ;.  0;.  ;					 
	\ngen.gen            ;.  gen        ;.7548;.     2;. 123;. 123;.  gen    ;.  gen     ;.   ;.  ;.  ;.  0;.  ;;
	\ngenannte.genannte  ;.  genannte   ;.13717;.    1;. 123;. 123;. genannt ;.  chiamato;.   ;.  ;.  ;.  0;.  ;;

	*/	
	wordToStudy_list = []
	let ixNumPlus; 
	let z;
	
	
	let listKey=[]; let keyS, keyIx;
	
	if (numButton == 1) {	
			//console.log("function js_go_showWordList_lev2 4 button 1 ")
			listKey = sortWordFreqFirst(wordToStudy_listStr) ;
			for (let x=0; x < listKey.length; x++ ) {
				[keyS, keyIx] = listKey[x].split(":") 
				oneElemToStudy(keyIx, x)
			} 		 
	}
	//-------------------------------------- 
	if (numButton == 5) {
		//console.log("function js_go_showWordList_lev2 5 button 5 ")
		for (let g=0; g < wordToStudy_listStr.length; g++) {
			let wordLineZ =	wordToStudy_listStr[g]	
			console.log( "%c   wordToStudy g" + g + " =>" + wordLineZ, "color:green;" )
			if (wordLineZ == "") return; 
		
			let ww0 = ((wordLineZ + ";.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.").split(";.") ).slice(0,15);
			
			[word2, word0, chk_ix, ixUnW2, totRow2, wLemmaList, wTranList,	wLevelList,	wParaList, wExampleList,
						   totExtrRow2, uLearnedYN,   chk_ixLemma, wIxLemmaList] = ww0;  	
						   
			console.log("     ", word2, "   ", word0 )			   
			/**		
			wordToStudy_list.push(  [word2, ixUnW2, totRow2, [wLemmaList], [wTranList], [wLevelList], [wParaList], [wExampleList], 
									totExtrRow2, uLearnedYN, [wIxLemmaList], numButton] ); 
			**/
			wordToStudy_list.push(  [word0, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, 
									totExtrRow2, uLearnedYN, wIxLemmaList, numButton] ); 					
									
									
		}
	}
	//------------------	
	if ((numButton > 1) && (numButton < 5)) {
		//console.log("function js_go_showWordList_lev2 6 button 1>1 e <5 ")
		for (let x=0; x < wordToStudy_listStr.length; x++ ) {
			oneElemToStudy(x)
		}		
		/**
		listKey = sortAlpha(wordToStudy_listStr)   // cod 
		for (let x=0; x < listKey.length; x++ ) {
				[keyS, keyIx] = listKey[x].split(":") 
				oneElemToStudy(keyIx)
		} 
		**/
	}
	//console.log("function js_go_showWordList_lev2 7 ")
	//--------------------
	
	function oneElemToStudy(z, x0) {		
		if (z < 0) return;
		let wordLineZ = wordToStudy_listStr[z].trim();  
		if (wordLineZ == "") return; 
		
		//[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList]
		
		let ww0 = ((wordLineZ + ";.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.").split(";.") ).slice(0,15);
		/**
		  abschied;.abschied;.ix;.4569;.1;.abscheiden§abschied;.separato§addio§;.A1|A1§A1|A1;.abscheiden|abscheiden§abschied|der abschied;.
								Die kranken Tiere von den gesunden abscheiden|in Frieden abscheiden§|der erste Abschied von zu Hause fiel ihm sehr schwer;.
								0;.y;.ixLemma;.306§308§;.;;

			uWord2				abschied;.abschied;.
								ix;.
			uIxUnW 				4569;.
			uTotRow				1;.
			uLemmaL   *lista	abscheiden§abschied;.          
			uTranL    *lista 	separato§addio§;.                
			uLevel    *lista	A1|A1§A1|A1;.                  
			uPara     *lista	abscheiden|abscheiden§abschied|der abschied;.             
			uExample  *lista	Die kranken Tiere von den gesunden abscheiden|in Frieden abscheiden§|der erste Abschied von zu Hause fiel ihm sehr schwer;.        
			uTotExtRow 			0;.
			uLearned   			y;.
								ixLemma;.
			uIxLemmaL *lista  	306§308§;.
			;;
			----
			                    // ;. separa i diversi campi della struttura, 
								// §  separa lemma diversi   nello stesso campo
								// |  separa level, paradigma ed esempi con lo stesso lemma  
			const wSep = "§";  									   
			const endOfLine = ";;\n"; 
		**/	
		
		[ word2, word0, chk_ix, ixUnW2, totRow2, wLemmaList, wTranList,	wLevelList,	wParaList, wExampleList,
					   totExtrRow2, 
					   uLearnedYN,   chk_ixLemma, wIxLemmaList ] = ww0; 
		
		wordToStudy_list.push(  [word0, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, 
									totExtrRow2, uLearnedYN, wIxLemmaList, numButton ] ); 							
									
			
	}  // end of oneElemToStudy
	///---------------------------------
	newTran = [];
	for ( z=0; z < wordToStudy_list.length; z++) {
		newTran.push( 0 )
	}
	//------------------------
	
	fun_showWordList("3") 
	
	//console.log("function js_go_showWordList_lev2 9")
	
} // end of js_go_showWordList_lev2
//--------------------------------------------------------------

function fun_showWordList(wh, ix1=-1) {	
		
	//console.log("function fun_showWordList  wh=", wh, " ix1=", ix1)	
		
	let numNoTran = 0; // -1 
	let word2, ixUnW2, totRow2, totExtrRow2, wLemmaList, wTranList , wLevelList, wParaList, wExampleList, uLearnedYN, wIxLemmaList  ; 
	let words_to_translate_str = wordTTBegin   //  ; 
	
	for (let z=0; z < wordToStudy_list.length; z++) {			
		[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, uLearnedYN, wIxLemmaList ] = wordToStudy_list[z] ; 	
		
		for(let f = 0; f < wLemmaList.length; f++) {	
			if (wLemmaList[f] == "") { continue }			
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
	//---------------------
	if (numNoTran < 1) {		
		showWordsAndTranButton("2")
		return
	}	
		
	ele_wordsToTranslate.value = words_to_translate_str; //  .replaceAll("\n"," "); 
	ele_wordsTranslated.value = ""; 
	getById("id_notTranNum").innerHTML = numNoTran; 
	
	onclick_jumpFromToPage( myPage01,0,myPage02);  
	

} // end of fun_showWordList

//-------------------------------------------
function showWordsAndTranButton(wh) {		
	//console.log("%cfunction  showWordsAndTranButton(" + wh + ")",  "color:blue;")
	
	let showList = prototype_tableWordList_Header;  
		
	let x2 = getById("id_sel_2_extrRow");
    let i = x2.selectedIndex;
	let sel_extrRow = x2.options[i].id;
	is_selected_row_only = ( i == index_onlySelRowsWanted);  // 2showWordsAndTranButton(wh) 
	
	//console.log("    1 showWordsAndTranButton  wordToStudy_list.length=", wordToStudy_list.length)
	
	fun_selRowsWanted_changed();
	
	//console.log("    2 showWordsAndTranButton  wordToStudy_list.length=", wordToStudy_list.length)
    
	let word2, ix1,ixUnW2, totRow2, nrow, totExtrRow2, wLemma1;
	let riga;	
	let wordOrig2, wordTran2;
	let wIxLemmaList, wLemmaList, wTranList , wLevelList, wParaList, wExampleList;  
	let uLearnedYN;	
	let wordTran = ""; 
	let wLemma3, wTran3;
	let nSpanV, spanV;	
	let clas1;
	let hig=1
	let swP; 
	let col1; 
	numeroWord_TR = 0;
	let numButton;
	
	//-------------------------------------------------------------------------
	try {
		for (let ixW2StudyLs = 0; ixW2StudyLs < wordToStudy_list.length; ixW2StudyLs++) {

			[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, 
					totExtrRow2, uLearnedYN, wIxLemmaList, numButton ] = wordToStudy_list[ixW2StudyLs]; 
					
			/**
			console.log("%c    LOOP wordToStudy_list " +  ixW2StudyLs , " color:red;")
			console.log("               word2=", word2, "wLemmaList type=", typeof wLemmaList, " =>", 
				wLemmaList, " wLemmaList.length=", wLemmaList.length)
			**/
			//------------
			showList += oneTR_lemma(ixW2StudyLs, ixW2StudyLs, "", word2, ixUnW2, totRow2, 
						wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, uLearnedYN, wIxLemmaList, numButton); 
		
			//------------------------	
			/**
			for(let ixixLemma = 0 ; ixixLemma < wLemmaList.length; ixixLemma++) {	
				showList += oneTR_lemma(ixW2StudyLs, ixixLemma, "", word2, ixUnW2, totRow2, 
						wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, uLearnedYN, wIxLemmaList, numButton); 
			}
			***/
		}
	} catch(e1) {
		logColor("%%red",e1)
		console.log("%c ERRORE in showWordsAndTranButton","color:red;")
		console.log("ixW2StudyLs=",ixW2StudyLs)
		console.log("wordToStudy_list[ixW2StudyLs]=", 	wordToStudy_list[ixW2StudyLs] )
		console.log("  wParaList=", wParaList , " type=", typeof wParaList )
		
	}
	//---------------------------------------------------
	//console.log("    3 showWordsAndTranButton")
	
	showList += '   </tbody>  \n' +
		'</table> \n'; 	
    ele_wordList.innerHTML = showList;
	//nascondi_celleEgualiPrecedenti() 	
	
	//allinea_colGroupTabWord();  
   	
	onclick_jumpFromToPage( myPage01,myPage02, myPage03);  
	let ele0, swEle, swChild, numChild0;
	[ele0, swEle, swChild, numChild0] = getById_children( "idTableWordList_thead");
	if ( (swEle == false) || (swChild == false) || (numChild0 < 1) )  return; 
	cellWord_TrTH =  getById("idTableWordList_thead").children[0].children[4]; 	
	
    [ele0, swEle, swChild, numChild0] = getById_children( "idTableWordList_tbody");
	if ( (swEle == false) || (swChild == false) || (numChild0 < 1) )  return; 
	
	//console.log("ele0=", ele0, " swEle=", swEle, "  swChild=", swChild, " numChild0=", numChild0)
	
	cellWord_TrTD = getById("idTableWordList_tbody").children[0].children[4]; 
	

	resize1.observe(cellWord_TrTD)
	//resize2.observe(cellWord_TrTH)		
			
	/**
	console.log("%cfunction  showWordsAndTranButton(" + wh + ")",  "color:blue;")		
	let eleX1 = getById("idTableWordList_tbody"); 
	let eleX2; 
	for(let mio1=0; mio1 < 20; mio1++) {
		console.log( mio1 + " tag=" + eleX1.tagName + " id=" + eleX1.id  + 
			" width=" + eleX1.style.width + " offw=" + eleX1.offsetWidth + "px" + " border=", eleX1.style.border) 	;
		if (eleX1.tagName == "HTML") break;
		eleX2 = eleX1.parentElement; 
		eleX1 = eleX2; 
	} 
	**/
	
} // end of showWordsAndTranButton
//------------------------------
let cellWord_TrTD;
let cellWord_TrTH; 			
function resizeTd_wordFunc(e) {
	if (e[0].target) {
		//console.log("e=", e[0] )
		//cellWord_TrTH.style.width = e[0].target.clientWidth + "px";
		cellWord_TrTH.style.width = e[0].contentRect.width + "px"; 
		//cellWord_TrTH.style.width = e[0].contentRect.width + "px"; 

		
		//cellWord_TrTD.style.width = e[0].target.offsetWidth + "px";
	} 
}	
				
let resize1 = new ResizeObserver(resizeTd_wordFunc)
//let resize2	= new ResizeObserver(resizeTd_wordFunc)
			
			//----------------
		
//---------------------------

function nascondi_celleEgualiPrecedenti() {	
	//console.log("%cfunction  nascondi_celleEgualiPrecedenti", "color:red;")
	let IX_WORD = 4; 
	let IX_LEMMA = IX_WORD+1
	
	let eleTab = getById("idTableWordList_tbody");	
	if (eleTab == null) return;	
	if (eleTab.tagName != "TBODY") { return; }
	
	let nRighe   = eleTab.rows.length;
	let eleRighe = eleTab.rows;	

	//console.log(" numrighe=", eleTab.rows.length)
	
	let preWord = ""; let preLemma=""
	let word1 = "";   let lemma1=""
	//---------------------
	for(let g=0; g < nRighe; g++) {
		let thisRiga = eleRighe[g]
		let cells = thisRiga.cells;
		if (cells.length < IX_WORD) {word1="_"; lemma1="_"} 
		word1  = cells[IX_WORD ].innerHTML; 
		lemma1 = cells[IX_LEMMA].innerHTML; 
		if (g > 0) {  
			if (word1  == preWord ) cells[IX_WORD ].children[0].className ="c_size_1_line"
			if (lemma1 == preLemma) cells[IX_LEMMA].children[0].className ="c_size_1_line"  
		}
		preWord = word1; preLemma = lemma1
	}	
} // end of nascondi_celleEgualiPrecedenti

//-------------------------------------------

function oneTR_lemma( ixW2StudyLs, ixLemma, clas1, word1, ix1, nrow, f_lemma, f_tran, f_level, f_para, f_example, 
					totExtrRow2, uLearnedYN, f_ixLemma, numButton) {	
			//console.log("%c        oneTR_lemma(" +ixW2StudyLs + " wor1=" + word1 + " lemma=" + f_lemma, "color:blue;") 		
			let wLemma1;
			let riga;
			let wordOrig2, wordTran2;			
			let wordTran = ""; 
			let wLemma3, wTran3;
			let nSpanV, spanV;			
			let showList = ""; 		
			
			
			
			let nn_para    = f_para.split("|")	
			let nn_example = f_example.split("|")	
			
			
			
			//----------------------------------------------------------------------
			// §  separa i lemma nello stesso campo 
			// |  separa level, paradigma ed esempi con lo stesso lemma  		
			//----------------------------------------------------------------------
			/*
			uLemmaL   *lista	abscheiden§abschied;.          
			uTranL    *lista 	separato§addio§;.                
			uLevel    *lista	A1|A1§A1|A1;.                  
			uPara     *lista	abscheiden|abscheiden§abschied|der abschied;.             
			uExample  *lista	Die kranken Tiere von den gesunden abscheiden|in Frieden abscheiden§|der erste Abschied von zu Hause fiel ihm sehr schwer;.   
			uIxLemmaL *lista  	306§308§;.
			ixLemma = 0 
					f_lemma = abscheiden, f_ixLemma = 306, f_tran = separato, f_level = A1|A2, f_para = abscheiden|abscheiden, f_example= Die kranken Tiere von den gesunden abscheiden|in Frieden abscheiden
						           nn_level   = [A1, A2], 
								   nn_para    = [abscheiden, abscheiden], 
								   nn_example = [Die kranken Tiere von den gesunden abscheiden, in Frieden abscheiden]									
			*/
			//--------------------------------------------------------------------
				
			let x_level="", x_para="", x_example=""; 
			let x_para1="", x_para2="", x_example1="", x_example2=""; 	
			let jbr; 			
			
			let riga00 = ""	
			
			let hig = 1.2 
						
			let XixW2StudyLs = ""+ixW2StudyLs;  
			let Xix1   = ""+ix1;  
			let Xnrow  = ""+nrow; 	
			let XnExtrRow = "";
			if (totExtrRow2) { if (totExtrRow2 > 0) { XnExtrRow = ""+totExtrRow2; } }  	
				
			let Xword1 = word1;
			
			let Xf_lemma = f_lemma;
			let Xf_para  = f_para; 
			let Xf_tran  = f_tran.replaceAll("|", "<br>") ; 
			
			
			
			//-----------------------	
			let key="",	pKey=""
			let numLev = nn_para.length;
			//let num1 = numLev
			
			key="";
			// num > 1, significa che ci sono diverse righe
			
			//prototype_lemmaTD   //   §tdLemma§ somma di tutti i lemma 
			/***
			//c_displayNone
			let prototype_lemmaTD = `
				<div style="display:none;">§one-f_lemma§</div>
				<div class="hpad top left1 §topBorder§"  style="resize:vertical; overflow-y:auto; width:100%;">										
					<span §displayNoneL§ onmouseover='mouseOverWord(this,3)' onmouseout='mouseOutWord(this,3)'><b>§one-f_lemma§</b></span>
					<span  class="c_paradigma" onmouseover='mouseOverWord(this,3)' onmouseout='mouseOutWord(this,3)'>§one-x_para1§</span>						
					<br   class="c_paradigma"  >	
					<span class="c_wordTran" style="display:none;">§one-f_tran§</span>		
					<div   class="c_example"  >§one-x_example1§</div>
				</div> \n					
		`; 	// end of prototype_lemmaTD

			***/
			let displayNone  = "";
			let displayNone1 = "";
			let displayNoneL = "";
			let topBorder = "";
			
			/**
			console.log("f_para=", f_para, "  nn_para=", nn_para)			
			console.log("            numLev=", numLev, "   nn_para=", nn_para, 		
				"  nn_example=", nn_example)
			**/
			
		
			//------------------------------------------------------
			let tdLemmaList = ""; 
			let last_para = ""
			for (let m=0; m < numLev; m++) {	
				/**
				key = x_level + " " + x_para + " " + x_example 
				if (key == pKey) {continue; }	
				pKey = key
				**/
				let newTdLemma = prototype_lemmaTD;
				
				//x_level   = nn_level[m]; 
				x_para    = sentenceOneRow( nn_para[m] );
				x_example = sentenceOneRow( nn_example[m] ); 

				if (f_para == "") x_para = f_lemma; 	
				
				//console.log( m, " x_para=", x_para, " xexample=", x_example)
				
				if (m > 0)	topBorder = "c_topBorder"; else topBorder="";
				
				
				/**
				if (x_para != "") {
					displayNoneL = 'style="display:none;"' ;
					if (x_para == last_para) {
						newTdLemma = newTdLemma.replaceAll("c_paradigma", "c_displayNone").replaceAll("c_wordTran", "c_displayNone")
						topBorder="";
					} 
				} else {
					displayNoneL = "";
				}	
				
				last_para = x_para	
				**/
				
				newTdLemma = newTdLemma.	
					replaceAll("§topBorder§"     , topBorder ). 
					replaceAll("§displayNone§"   , displayNone ).  
					replaceAll("§displayNone1§"  , displayNone1 ).  
					replaceAll("§displayNoneL§"  , displayNoneL ). 	
					replaceAll("§one-f_lemma§"   , ""+f_lemma  ). 			 
					replaceAll("§one-x_level§"   , ""+x_level  ).
					replaceAll("§one-f_tran§"    , ""+Xf_tran  ).   
					replaceAll("§one-x_para1§"   , ""+x_para  ). 
					replaceAll("§one-x_example1§", ""+x_example ) 
				;  	
			
				tdLemmaList += newTdLemma ;
				//console.log("                             tdLemmaList=", tdLemmaList)
			}	
			//----------			
			numeroWord_TR++;				
			let newTr = newTr_from_prototype_manyLev( numeroWord_TR, clas1, XixW2StudyLs, ""+f_ixLemma, ixLemma, Xnrow, 
					Xix1, Xword1, Xf_lemma, x_level, Xf_para, Xf_tran, x_para1, x_example1,"", 0, 
					XnExtrRow, uLearnedYN, numButton, tdLemmaList); 
			/*
			console.log("        newTr = newTr_from_prototype_manyLev: ", " Xword1=", Xword1, " Xf_lemma=", Xf_lemma, " Xf_para=", Xf_para,
				" tdLemmaList.length=", tdLemmaList.length	)
			*/	
			
			showList    += newTr ;	
			
			
		return showList
		
}  // end of oneTR_lemma
//----------------------------------------

function sentenceOneRow( str1 ) {
	if (str1 == undefined)  { return str1; }
	let str2 = (""+str1).replaceAll("\n"," ")
	str2 = str2.replaceAll(". ",".<br>").replaceAll("? ","?<br>").replaceAll("! ","!<br>"); 
	
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
	
	return str2	
}

//-------------------------------------------
//--------------------------------------------------------------------

	function newTr_from_prototype_manyLev( numeroTR, clas1, ixW2StudyLs, ixLemma, ixixLemma, nrow, ix1, 
					word1, 
					f_lemma, x_level, 
					f_para, f_tran, x_para1,x_example1, showAltre, m, 
					n_extr_row1, 
					uLearnedYN, numButton,
					tdLemmaList
					) {  
		//console.log("newTr_from_prototype_manyLev ", " word1=", word1, " lemma=", f_lemma, "\n\ttdLemmaList=",tdLemmaList)    
		let displayNone  = "";
		let displayNone1 = "";
		let displayNoneL = "";
		let dyNoneTD234  = ""; 
		
		
		let newTr = prototype_oneTR_lemma.trim(); 
		
		//if (x_level != "")    { x_level    = "(lev." + x_level + ")"; }
		if (x_para1 != "")    {
			displayNoneL = 'style="display:none;"'
		}		
		if (x_example1 == undefined) console.log("x_example1=", x_example1);
		if (x_example1 != "") { x_example1 = x_example1 ; }
		let summarystyle = ' style="list-style-position: outside;" ' 
		if ((x_level == "") && ( x_para1=="") && (x_example1=="") ) {
			summarystyle= ' style="display:block;" '  // in <detail><summary></summary> other staff </details> if other stuff is empty hide the arrow (default is display:list-item) 
		}
		/**
		if (nrow == 0) {
			displayNone  = 'style="display:none;"';
			displayNone1 = 'style="display:none;"';
		} 
		**/
		
		let wordvisib = "";
		if (numButton == "5") {  
			//wordvisib = ' style="visibility:hidden;" ' ;
			//anto4agosto word1 = "";
			//anto4agosto displayNone1 = 'style="display:none;"';
		} else if (numButton == "2") {  
			n_extr_row1 = "";
		} else if (numButton == "3") {  
			n_extr_row1 = "";
		}
	
		
		/**
		if (word1 == "$lemmalist$") {
			word1 = buttonWordList( f_lemma )
		}
		**/
		let yesOrNot;
		let learnedClass;
		if (uLearnedYN == "y") {
			yesOrNot = YES1; 
			learnedClass = '';	
		} else { 
			yesOrNot = NOT_YET1;
			learnedClass = ' class="notYetLearnedButt" '
		}
		
		if (ix1 < 0) {			
			displayNone  = 'style="display:none;"';
			displayNone1 = 'style="display:none;"';
			dyNoneTD234  = 'style="display:none;"';
		} 
		if (word1 == "") {displayNone1 = 'style="display:none;"';}		
		
		//console.log(word1 , "  nrow=", nrow,  " displayNone=>" + displayNone, "<==")
		newTr = newTr.
			replaceAll("§summarystyle§"  , summarystyle).
			replaceAll(" §wordvisib§"    , wordvisib   ). 
			replaceAll("§displayNone§"   , displayNone ).  
			replaceAll("§displayNone1§"  , displayNone1 ).  
			replaceAll("§displayNoneL§"  , displayNoneL ).  
			replaceAll("§dyNoneTD234§"   , dyNoneTD234  ).
			replaceAll("§ixW2StudyLs§"      , ""+ixW2StudyLs ). 
			replaceAll("§one-ixLemma§"   , ""+ixLemma   ). 
			replaceAll("§one-ixixLemma§" , ""+ixixLemma ). 
			replaceAll("§one-numTR§"     , ""+numeroTR ). 
			replaceAll("§one-nrow§"      , ""+nrow     ). 
			replaceAll("§one-ix1§"       , ""+  ix1    ). 
			replaceAll("§one-word1§"     , ""+word1    ). 			
			replaceAll("§one-f_lemma§"   , ""+f_lemma  ). 			 
			replaceAll("§one-x_level§"   , ""+x_level  ).
			replaceAll("§one-f_tran§"    , ""+f_tran   ).   
			replaceAll("§one-x_para1§"   , ""+x_para1  ). 
			replaceAll("§one-x_example1§", ""+x_example1 ). 
			replaceAll("§one-showAltre§",  ""+ showAltre). 	
			replaceAll("§one-n_extr_row§", ""+ n_extr_row1).	
			replaceAll("§YESorNOT§"      , ""+yesOrNot).  		
			replaceAll("§learnedClass§"  , learnedClass).
			replaceAll("§tdLemma§"       , tdLemmaList )  	
			;  
		return newTr;	
		
	}  // end of newTr_from_prototype
//---------------------------------------	




