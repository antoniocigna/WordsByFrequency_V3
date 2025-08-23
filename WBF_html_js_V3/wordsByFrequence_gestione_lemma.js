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

//----------------------
function js_go_showWordList_lev2(wordListStr00, numButton=1, jsFunc="",goFunc="") {
	/*
	console.log("js_go_showWordList_lev2  nel file wordsByFrequence_gestione_lemma.js"); 	
	console.log("%cfunction js_go_showWordList_lev2 ","color:green;"); 
	console.log("    ", "wordListStr.length=",wordListStr00.length  ," numButton=", numButton, " <-- " + goFunc + " <-- " + jsFunc) ;
	*/
	//console.log("js_go_showWordList_lev2    ", "wordListStr.length=",wordListStr00.length  ," numButton=", numButton, " <-- " + goFunc + " <-- " + jsFunc) ;
	//console.log("    ", wordListStr00);

	// numButton=1 default ==> from onclick most frequent word list  
	// numButton=2         ==> from onclick BetweenWordList or prefix wordlist   
	// numButton=3         ==> from onclick Lemma word list   
	// numButton=5         ==> from onclick Lemma list   
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
	
	//console.log("function js_go_showWordList_lev2 2 ")
	
    if (wordListStr == "") {
        //console.log("js_showWordList: parameter is empty");
		if (numButton==1) { errorNoWord1()}
		onclick_jumpFromToPage( myPage02,myPage03,  myPage01); 
        return;
    }
	
	myPage01.style.display = "none"; 
	
	//console.log("function js_go_showWordList_lev2 3 ")

	var wLemmaListU, wTranListU, wLevelListU,	wParaListU, wExampleListU, wIxLemmaListU;   
	var wLemmaList,  wTranList,  wLevelList,    wParaList,  wExampleList , wIxLemmaList ;
	var word2, ixUnW2, totRow2, totExtrRow2 
	var uLearnedYN;
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
			//console.log("function js_go_showWordList_lev2 4 button 1 ")
			listKey = sortWordFreqFirst(wordToStudy_listStr) ;
			for (var x=0; x < listKey.length; x++ ) {
				[keyS, keyIx] = listKey[x].split(":") 
				oneElemToStudy(keyIx, x)
			} 		 
	}
	//-------------------------------------- 
	if (numButton == 5) {
		//console.log("function js_go_showWordList_lev2 5 button 5 ")
		for (var g=0; g < wordToStudy_listStr.length; g++) {
			var wordLineZ =	wordToStudy_listStr[g]	
			//console.log( "%c   wordToStudy g" + g + " =>" + wordLineZ, "color:green;" )
			if (wordLineZ == "") return; 
		
			var ww0 = ((wordLineZ + ";.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.").split(";.") ).slice(0,15);
			
			[wordCod, word2, chk_ix, ixUnW2, totRow2, wLemmaList, wTranList,	wLevelList,	wParaList, wExampleList,
						   totExtrRow2, uLearnedYN,   chk_ixLemma, wIxLemmaList] = ww0;  	
			/**		
			wordToStudy_list.push(  [word2, ixUnW2, totRow2, [wLemmaList], [wTranList], [wLevelList], [wParaList], [wExampleList], 
									totExtrRow2, uLearnedYN, [wIxLemmaList], numButton] ); 
			**/
			wordToStudy_list.push(  [word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, 
									totExtrRow2, uLearnedYN, wIxLemmaList, numButton] ); 					
									
									
		}
	}
	//------------------	
	if ((numButton > 1) && (numButton < 5)) {
		//console.log("function js_go_showWordList_lev2 6 button 1>1 e <5 ")
		listKey = sortAlpha(wordToStudy_listStr)   // cod 
		
		for (var x=0; x < listKey.length; x++ ) {
				[keyS, keyIx] = listKey[x].split(":") 
				oneElemToStudy(keyIx)
		} 
	}
	//console.log("function js_go_showWordList_lev2 7 ")
	//--------------------
	
	function oneElemToStudy(z, x0) {		
		if (z < 0) return;
		var wordLineZ = wordToStudy_listStr[z].trim();  
		if (wordLineZ == "") return; 
		
		//[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList]
		
		var ww0 = ((wordLineZ + ";.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.;.").split(";.") ).slice(0,15);
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
		/**
		[wordCod, word2, chk_ix, ixUnW2, totRow2, wLemmaListU, wTranListU,	wLevelListU,	wParaListU, wExampleListU,
					   totExtrRow2, 
					   uLearnedYN,   chk_ixLemma, wIxLemmaListU] = ww0;  
		
		wLemmaList   = wLemmaListU.split(   wSep ) 			   
		wTranList    = wTranListU.split(    wSep ) 		
		wLevelList   = wLevelListU.split(   wSep ) 		
		wParaList    = wParaListU.split(    wSep ) 		
		wExampleList = wExampleListU.split( wSep ) 		   
		wIxLemmaList = wIxLemmaListU.split( wSep ) 		

		if ((word2 == "") || (word2 == "...") ) return ; 
		if (word2.indexOf("…") >= 0) return;  
		var trattini = "-_*."; 
		if (trattini.indexOf( word2.substring(0,1) ) >=0) return;
		
		wordToStudy_list.push(  [word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, 
									totExtrRow2, uLearnedYN, wIxLemmaList, numButton ] ); 
		**/
		[wordCod, word2, chk_ix, ixUnW2, totRow2, wLemmaList, wTranList,	wLevelList,	wParaList, wExampleList,
					   totExtrRow2, 
					   uLearnedYN,   chk_ixLemma, wIxLemmaList ] = ww0; 
		
		wordToStudy_list.push(  [word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, 
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
		
	var numNoTran = 0; // -1 
	var word2, ixUnW2, totRow2, totExtrRow2, wLemmaList, wTranList , wLevelList, wParaList, wExampleList, uLearnedYN, wIxLemmaList  ; 
	var words_to_translate_str = wordTTBegin   //  ; 
	
	for (var z=0; z < wordToStudy_list.length; z++) {			
		[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, uLearnedYN, wIxLemmaList ] = wordToStudy_list[z] ; 	
		
		for(var f = 0; f < wLemmaList.length; f++) {	
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
	document.getElementById("id_notTranNum").innerHTML = numNoTran; 
	
	onclick_jumpFromToPage( myPage01,0,myPage02);  
	

} // end of fun_showWordList

//-------------------------------------------
function showWordsAndTranButton(wh) {		
	//console.log("%cfunction  showWordsAndTranButton(" + wh + ")",  "color:blue;")
	var showList = prototype_tableWordList_Header;  
		
	var x2 = document.getElementById("id_sel_2_extrRow");
    var i = x2.selectedIndex;
	var sel_extrRow = x2.options[i].id;
	is_selected_row_only = ( i == index_onlySelRowsWanted);  // 2showWordsAndTranButton(wh) 
	
	//console.log("    1 showWordsAndTranButton  wordToStudy_list.length=", wordToStudy_list.length)
	
	fun_selRowsWanted_changed();
	
	//console.log("    2 showWordsAndTranButton  wordToStudy_list.length=", wordToStudy_list.length)
    
	var word2, ix1,ixUnW2, totRow2, nrow, totExtrRow2, wLemma1;
	var riga;	
	var wordOrig2, wordTran2;
	var wIxLemmaList, wLemmaList, wTranList , wLevelList, wParaList, wExampleList;  
	var uLearnedYN;	
	var wordTran = ""; 
	var wLemma3, wTran3;
	var nSpanV, spanV;	
	var clas1;
	var hig=1
	var swP; 
	var col1; 
	numeroWord_TR = 0;
	var numButton;
	
	//-------------------------------------------------------------------------
	try {
		for (var ixW2StudyLs = 0; ixW2StudyLs < wordToStudy_list.length; ixW2StudyLs++) {

			[word2, ixUnW2, totRow2, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, 
					totExtrRow2, uLearnedYN, wIxLemmaList, numButton ] = wordToStudy_list[ixW2StudyLs]; 
					
			//nSpanV =  wLemmaList.length
			//console.log("%c    LOOP wordToStudy_list " +  ixW2StudyLs , " color:red;")
			/*
			console.log("               word2=", word2, "wLemmaList type=", typeof wLemmaList, " =>", 
				wLemmaList, " wLemmaList.length=", wLemmaList.length)
			*/
			//------------
			showList += oneTR_lemma(ixW2StudyLs, ixW2StudyLs, "", word2, ixUnW2, totRow2, 
						wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, uLearnedYN, wIxLemmaList, numButton); 
		
			//------------------------	
			/**
			for(var ixixLemma = 0 ; ixixLemma < wLemmaList.length; ixixLemma++) {	
				showList += oneTR_lemma(ixW2StudyLs, ixixLemma, "", word2, ixUnW2, totRow2, 
						wLemmaList, wTranList, wLevelList, wParaList, wExampleList, totExtrRow2, uLearnedYN, wIxLemmaList, numButton); 
			}
			***/
		}
	} catch(e1) {
		console.log("%c ERRORE in showWordsAndTranButton","color:red;")
		console.log("ixW2StudyLs=",ixW2StudyLs)
		console.log("wordToStudy_list[ixW2StudyLs]=", 	wordToStudy_list[ixW2StudyLs] )
		console.log("  wParaList=", wParaList , " type=", typeof wParaList )
		console.log(e1)
	}
	//---------------------------------------------------
	//console.log("    3 showWordsAndTranButton")
	
	showList += '   </tbody>  \n' +
		'</table> \n'; 	
    ele_wordList.innerHTML = showList;
	//nascondi_celleEgualiPrecedenti() 	
	
	//allinea_colGroupTabWord();  
   	
	onclick_jumpFromToPage( myPage01,myPage02, myPage03);  
	
	
   cellWord_TrTH =  document.getElementById("idTableWordList_thead").children[0].children[4]; 	
		
	
	cellWord_TrTD = document.getElementById("idTableWordList_tbody").children[0].children[4]; 

	resize1.observe(cellWord_TrTD)
	//resize2.observe(cellWord_TrTH)		
			
			//----------------
	
	
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
				
var resize1 = new ResizeObserver(resizeTd_wordFunc)
//var resize2	= new ResizeObserver(resizeTd_wordFunc)
			
			//----------------
		
//---------------------------

function nascondi_celleEgualiPrecedenti() {	
	//console.log("%cfunction  nascondi_celleEgualiPrecedenti", "color:red;")
	let IX_WORD = 4; 
	let IX_LEMMA = IX_WORD+1
	
	let eleTab = document.getElementById("idTableWordList_tbody");	
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
/**
function oneTR_lemma( ixW2StudyLs, ixLemma, clas1, word1, ix1, nrow, wLemmaList, wTranList, wLevelList, wParaList, wExampleList, 
					totExtrRow2, uLearnedYN, wIxLemmaList, numButton) {	
***/
function oneTR_lemma( ixW2StudyLs, ixLemma, clas1, word1, ix1, nrow, f_lemma, f_tran, f_level, f_para, f_example, 
					totExtrRow2, uLearnedYN, f_ixLemma, numButton) {	
			//console.log("%c        oneTR_lemma(" +ixW2StudyLs + " wor1=" + word1 + " lemma=" + f_lemma, "color:blue;") 		
			var wLemma1;
			var riga;
			var wordOrig2, wordTran2;			
			var wordTran = ""; 
			var wLemma3, wTran3;
			var nSpanV, spanV;			
			var showList = ""; 		
			//if (wExampleList.length == 0)  wExampleList[0] = "";
			//------------------------------------------------
			/**
			var f_lemma   = wLemmaList[  ixLemma ];
			var f_ixLemma = wIxLemmaList[ixLemma ];
			var f_tran    = wTranList[   ixLemma ]; 			
			var f_para    = wParaList[   ixLemma ]; 
			var f_example = wExampleList[ixLemma ]; 
			
			//console.log("oneTR_lemma ixLemma=", ixLemma, " wParaList=",  wParaList)
			if (f_lemma   == undefined) f_lemma = "";
			if (f_ixLemma == undefined) f_ixLemma = "";
			if (f_tran    == undefined) f_tran = "";		
			if (f_para    == undefined) f_para    = "";	
			if (f_example == undefined) f_example = "";
			//console.log("   oneTR_lemma  f_para=", f_para)
			***/
			//------------------------------------------------------
			//var nn_level   = f_level.split("|")	
			
			//console.log("f_para=", f_para,  " type=", typeof f_para)
			
			
			var nn_para    = f_para.split("|")	
			var nn_example = f_example.split("|")	
			
			
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
				
			var x_level="", x_para="", x_example=""; 
			var x_para1="", x_para2="", x_example1="", x_example2=""; 	
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
			var numLev = nn_para.length;
			//var num1 = numLev
			
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
			
			//console.log("            numLev=", numLev, "   nn_para=", nn_para, "  nn_example=", nn_example)
			
			//------------------------------------------------------
			let tdLemmaList = ""; 
			let last_para = ""
			for (var m=0; m < numLev; m++) {	
				key = x_level + " " + x_para + " " + x_example 
				if (key == pKey) {continue; }	
				pKey = key
				let newTdLemma = prototype_lemmaTD	
				
				//x_level   = nn_level[m]; 
				x_para    = sentenceOneRow( nn_para[m] );
				x_example = sentenceOneRow( nn_example[m] ); 	
				
				if (m > 0)	topBorder = "c_topBorder"; else topBorder="";
				
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
			var newTr = newTr_from_prototype_manyLev( numeroWord_TR, clas1, XixW2StudyLs, ""+f_ixLemma, ixLemma, Xnrow, 
					Xix1, Xword1, Xf_lemma, x_level, Xf_para, Xf_tran, x_para1, x_example1,"", m, 
					XnExtrRow, uLearnedYN, numButton, tdLemmaList); 
			/*
			console.log("        newTr = newTr_from_prototype_manyLev: ", " Xword1=", Xword1, " Xf_lemma=", Xf_lemma, " Xf_para=", Xf_para,
				" tdLemmaList.length=", tdLemmaList.length	)
			*/	
			
			showList    += newTr ;	
			
			
		return showList
		
}  // end of oneTR_lemma
//----------------------------------------
//--------------------------------------------------------------------

	function newTr_from_prototype_manyLev( numeroTR, clas1, ixW2StudyLs, ixLemma, ixixLemma, nrow, ix1, word1, f_lemma, x_level, 
					f_para, f_tran, x_para1,x_example1, showAltre, m, 
					n_extr_row1, 
					uLearnedYN, numButton,
					tdLemmaList
					) {  
	   
		var displayNone  = "";
		var displayNone1 = "";
		var displayNoneL = "";
		var dyNoneTD234  = ""; 
		
		
		var newTr = prototype_oneTR_lemma.trim(); 
		
		//if (x_level != "")    { x_level    = "(lev." + x_level + ")"; }
		if (x_para1 != "")    {
			displayNoneL = 'style="display:none;"'
		}		
		if (x_example1 == undefined) console.log("x_example1=", x_example1);
		if (x_example1 != "") { x_example1 = x_example1 ; }
		var summarystyle = ' style="list-style-position: outside;" ' 
		if ((x_level == "") && ( x_para1=="") && (x_example1=="") ) {
			summarystyle= ' style="display:block;" '  // in <detail><summary></summary> other staff </details> if other stuff is empty hide the arrow (default is display:list-item) 
		}
		/**
		if (nrow == 0) {
			displayNone  = 'style="display:none;"';
			displayNone1 = 'style="display:none;"';
		} 
		**/
		
		var wordvisib = "";
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
		var yesOrNot;
		var learnedClass;
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


