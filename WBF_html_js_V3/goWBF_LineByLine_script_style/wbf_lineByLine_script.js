"use strict";
/*  
LineByLine: A tool to practice language comprehension
Antonio Cigna 2023
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//----------------------------------------------
let ele_select_first_language = document.getElementById("id_selectFirstLanguage");

let word_to_underline = "1234"; 
var word_to_underline_list = ["11","22"]; 

var currScript = document.currentScript.src; var bar1 = currScript.lastIndexOf("\\");var bar2 = currScript.lastIndexOf("/"); 
//console.log("LOADED file SCRIPT " + currScript.substring( 1+Math.max(bar1,bar2) )) ;	
//----------------------------------------------------------------------------------------
let swdebug = false;
let sw_word_script=false; 
let word_last_BoldRow;
let last_row_bold_ele = null;
let ele_voxDisplay = document.getElementById("id_voxDisplay"); 
let ele_voxLangDisplay = document.getElementById("id_voxLangDisplay");

const NO_VIDEO_AUDIO_FILE = "noVIDEO_noAUDIO".toLowerCase();
let sw_is_no_videoaudio = false;

let clock_timer_symb = "&#x23f1;";
let playLoop_symb = "&infin;";
let speakinghead_symb = "&#128483;";
let magnifyingGlass_symb = "&#128270;";
let pause_symb = "&#x23F8;";
let play_symb = "&#x23F5;";
let word_pause_symb = "&#x1d110;"; 

let openbook_symb = "&#128214;";
let closedbook_symb = "&#128213;";
let left_arrow_symb = "&#8592;";
let right_arrow_symb = "&#8594;";
let breakwords_symb = "/|/";

let Clip_startTime = 0;
let Clip_stopTime = 0;
let clipFromRow = 0;
let clipFromRow_min = 0;
let clipToRow = 0;
let clip_play_time_interrupt = 0;

let sw_the_are_no_subtitles = false;
const NO_TEXT_NO_SUBTITLES = "NO_TEXT_NO_SUBTITLES";

let sw_CLIP_play = false;

let CLIP_Loop_StartTime = 0;
let CLIP_Loop_StopTime = 0;

let sw_CLIP_row_play = false;



let sw_active_show_lastSpokenLineTranslation = false;

let CLIP_Row_StartTime = 0;
let CLIP_Row_StopTime = 0;
let CLIP_Row_StopIx = -1;

let last_ixClip = -1;

let ele_last_tran_line;

//let ele_tts = document.getElementById("id_tts"); // synthetic voce  <div >...</div>
let sw_tts = false;
let sw_tts2;

//
let ONEDAY = 24 * 60 * 60 * 1000;
let start_milliseconds = Date.now();

let num_day_today = Math.floor(start_milliseconds / ONEDAY);

let sw_audio = false;


let title1 = "";

//--------------------------------------------------
let INITIAL_hex_BG_color = "#C8BEBC"; //  "#FFFFF";
let INITIAL_hex_FG_color = "#000000";
let sw_show_time_in_subtitle = false;
let sw_show_clip_num = true;
let numOrig = 0;
let numTran = 0;


//-------------------------------------------------------
let vid_duration = 0;

let ele_video_speed;
let video_native_height;
let video_native_width;
let newVidH, newVidW;
let newVidHperc, newVidWperc;
let inp_text_orig;
let inp_text_tran;

let ele_mask_dragsub = document.getElementById("id_mask_dragsub");
let eleTabSub = document.getElementById("id_tabSub");
let eleTabSub_tbody = document.getElementById("id_tabSub_tbody");

let selected_voice_ix                    = 0 ;     // eg. 65 	 
let selected_voiceName                   = "";     // eg. Microsoft David - English (United States)"; 	
let	selected_voiceLangRegion             = "";     // eg. en-us	
let	selected_voiceLang2                  = "";     // eg. en

let isVoiceSelected = false;  

let selected_numVoices = 0;
let maxNumVoices = 9999; // 9

let last_pitch = 1;
let last_rate  = 1;
let last_volume = 1;

let ele_playNextVa_from_hhmmss_value;
let ele_replayVa_from_hhmmss_value;
let ele_replayVa_to_hhmmss_value;

let radio_type1_SECONDS = "a";
let radio_type2_SECS_END_DIALOG = "b";
let radio_type4_LINES = "d";

let CLIP_dur_line_TYPES = radio_type4_LINES;

let LS_voice_index = -1;



//----------------------------------------------------------------

let cbc_LOCALSTOR_key = "";

//  the LS_... variables here after have their values stored in window.localStorage so that they can be retrieved in the next sessions 
//  all this value are put in a list and saved in one variable the name of which contains the title of the page (each page has its own values)  

let LS_clip_secs0b_value = 3;
let LS_clip_secs1d_value = "3-false-true";
let LS_clip_secs2c_value = 3;
let LS_clip_lines_value = 3;

let LS_clip_checked_sw_type = radio_type2_SECS_END_DIALOG; // intial default
let LS_clip_checked_type_value = 3;


let LS_stor_ext_time = 1.5;
let LS_stor_playnext_replay = "00:00:00,000;;00:00:00,000;;00:00:00,000;;";
let LS_subDeltaTime = 0;
let LS_colorBG = "#f3c89d";
let LS_colorTx = "black";
let LS_sub_beg_delta = 0;
let LS_sub_end_delta = 0;
let LS_sub_force_visible = true;


let sw_end_dialog_sentence = true;
let sw_end_dialog_overlap = true;


let cbc_LOCALSTOR_ctrkey = "";
let LS_CTR_1run_num = 0; // number of runs of this video since  building   
let LS_CTR_2run_days = 0; // number of days of runs  
let LS_CTR_3run_num_play = 0;
let LS_CTR_4run_elapsed = 0; // total elapsed time of the runs  
let LS_CTR_5run_videoRunTime = 0; // total of playback elapsed time 
let LS_CTR_6logon_num_play = 0; // number of play enter since logon   
let LS_CTR_7logon_elapsed = 0; // total elapsed time of the runs  
let LS_CTR_8logon_videoRunTime = 0; // total of playback elapsed time 
let LS_CTR_9numday_today = 0; // num.day  starting 1970 
let LS_CTR_10run_num_lines = 0; // num.of subtitle lines 
let LS_CTR_11run_num_words = 0; // num.of subtitle words 
let LS_CTR_12logon_num_lines = 0; // num.of subtitle lines 
let LS_CTR_13logon_num_words = 0; // num.of subtitle words 

//-------------------------------------------------------------
let ele_last_play;
let LAST_time_change_ix = 1;
let LAST_time_change_secs = 0;
let PLAYCLIP_TYPE = 0;
let PLAYCLIP_FROM_TIME = 0;
let PLAYCLIP_TO_TIME = 0;
let PLAYCLIP_FROM_LINE = 0;
let PLAYCLIP_TO_LINE = 0;
let CLIP_ENDTIME_PLUS = 0;
let CLIP_ENDTIME_MINUS = 0;
let SW_CLIP_ENDED = false;
//--------------------------------------------------------

let playTRIGGER_clip = 1;
let playTRIGGER_row = 2;
let playTRIGGER_emul = 3;

let lastPlayTrigger = playTRIGGER_emul;
let lastPlayRowFromIx, lastPlayRowToIx, lastPlayRowFromTime, lastPlayRowToTime, lastPlayRowLoop;
let lastPlayListFrom = [];
let lastPlayListIx = -1;
let last_BoldRow;
//----------------------------------------
let TO_TIME_TOLERANCE = 0.5;


let ele_replay_from_secs_innerHTML;

let ele_replayVa_from_row_value;

let ele_replay2_from_row_innerHTML;

let ele_replay_to_secs_innerHTML;

let ele_replayVa_to_row_value;

let ele_replay2_to_row_innerHTML;



let ele_clip_subtext;
let html_parms_queryString = "";
//------------------------



let ele_dragSubT;
let ele_dragSubT_anchor;

let wScreen;
let hScreen;
let subtitles_beg_delta_time;

let src1;
let vid;
let MAX999;
let path1;
let sayNODIALOG = "-NODIA-";
let f1;
let f2;
let f3;
let barra;


let lastClipTimeBegin;
let lastClipTimeEnd;
let ele_time_video;
let ele_sub_filler;
let ele_subOrigText2;
let ele_subTranText2;
let ele_showOrigText2_open;
let ele_showOrigText2_close;
let ele_showTranText2_open;
let ele_showTranText2_close;
let ele_subOrigSilent;
let ele_subOrigSilentH;
let ele_main_subt = document.getElementById("id_main_subt");

let list_elemSub = ["", ""];
let list_elemSub_display = [false, false];

let line_list_o_number_of_elements = 0;
let line_list_t_number_of_elements = 0;

let sw_sub_onfile = false;
let sw_sub_orig = false;
let sw_sub_tran = false;
let sw_no_subtitle = false; // no subtitles ( neither inside the video, neither in any file apart

let LIMIT_MIN_TIME_CLIP;
let MIN_ixClip = 0;
let MAX_ixClip = MAX999;

let inp_row_orig = [];
let inp_row_tran = [];

let number_of_subtitle_endsentence = 0;
let number_of_subtitle_time_overlap = 0;

let line_list_o_from00 = [];
let line_list_o_to00 = [];
let line_list_o_maxto00 = [];

let line_list_o_from1 = [];
let line_list_o_to1 = [];
let line_list_o_maxto1 = [];
let line_list_orig_text = [];
let line_list_orig_nFileR = []; 
let line_list_orig_tts = [];
let line_list_o_tran_ixmin = [];
let line_list_o_tran_ixmax = [];

let line_list_t_from00 = [];
let line_list_t_to00 = [];
let line_list_t_maxto00 = [];

let line_list_t_from1 = [];
let line_list_t_to1 = [];
let line_list_t_maxto1 = [];
let line_list_tran_text = [];

//orig_and_tran_in_one_line();
let LAST_SEARCH_o_time = -1;
let LAST_SEARCH_o_ix;
let LAST_SEARCH_o_fromto;

let LAST_SEARCH_t_time = -1;
let LAST_SEARCH_t_ix;
let LAST_SEARCH_t_fromto;
//------------------------------------



let ele_ctl_playpause = document.getElementById("id_ctl_playpause");
ele_ctl_playpause.children[0].innerHTML = ele_ctl_playpause.children[0].innerHTML.replace("§play_symb§", play_symb);
ele_ctl_playpause.children[1].innerHTML = ele_ctl_playpause.children[1].innerHTML.replace("§pause_symb§", pause_symb);
let ele_ctl_slider = document.getElementById("id_ctl_slider");
let ele_ctl_value = document.getElementById("id_ctl_value");
let ctl_slider_maxValue_hhmm = 0;
let myLang = navigator.language; // eg.  it-IT
let decimal_point = (0.123).toLocaleString(myLang).toString().substr(1, 1);


//-----------------------------

let word_fromIxToIxLimit = [-1, -1];
let word_fromIxToIxButtonElement = [null, null];


//----------------------------------------------

let clip_reset_BG_color = "white";
let clip_somerow_BG_color = "lightgrey";


let begix, endix;
let fromIxToIxLimit = [-1, -1];
let fromIxToIxButtonElement = [null, null];

let save_last_oneOnlyRow = "";
let save_last_oneOnly_idtr = "";
save_last_oneOnlyRow = "";
save_last_oneOnly_idtr = "";

let hide_translation_symb = '<span style="font-weight:bold;min-width:4em;">t?</span></span>';
let show_translation_symb = '<span style="font-weight:bold;">T</span></span>';


let note_arrow1 = '<span style="font-size:2em;width:auto;height:1.4em;">' + right_arrow_symb + '</span>';
let note_arrow2 = '<span style="font-size:2em;width:auto;height:1.4em;">' + left_arrow_symb + '</span>';
let note_speaking = '<span style="font-size:2em;width:auto;height:1.4em;">' + speakinghead_symb + '</span>';

let note_magnifyingGlass_symb = '<span style="font-size:2em;width:auto;height:1.4em;">' + magnifyingGlass_symb + '</span>';

let note_loop_speaking = '<span style="font-size:2em;width:auto;height:1.4em;">' + playLoop_symb + '</span>';
let note_hide_sub = '<span style="font-size:2em;width:auto;height:1.4em;">' + openbook_symb + '</span>';
let note_show_sub = '<span style="font-size:2em;width:auto;height:1.4em;">' + closedbook_symb + '</span>';
let note_show_tran = '<span style="font-size:2em;width:auto;height:1.4em;">' + show_translation_symb + '</span>';
let note_hide_tran = '<span style="font-size:2em;width:auto;height:1.4em;">' + hide_translation_symb + '</span>';
let note_breakwords = '<span style="font-size:2em;width:auto;height:1.4em;">' + breakwords_symb + '</span>';
let note_clock_timer_symb = '<span style="font-size:2em;width:auto;height:1.4em;">' + clock_timer_symb + '</span>';


//--------------------------------

//js2___________________  js2


let last_ele_analWords_id; 
let last_ele_analWords_tr;
let last_ele_analWords_height;
//js3___________________  	

//js4___________________  	

let startTime;
let txt_length; 
let sw_pause = false; 
let sw_cancel = false;
let time_limit = 15; // in seconds  //  it seems that  an utterance can't last more     
let ELAPSED_TIME_SPEECH_LIMIT = 1000 * ( time_limit - 1) ;  

let tot_norm_time = 0;
let tot_txt_len =0;	
let tot_norm_mill_char = 0; 	  
let tot_norm_str_leng_limit = 0;	 
let TXT_SPEECH_LENGTH_LIMIT = 80; // initial value is updated according to the actual duration runs
//----
let TTS_LOOP_begix=-1;
var	TTS_LOOP_endix=-1; 
var	TTS_LOOP_swLoop=false; 	
let TTS_LOOP_elem;   
//--------------------------
let voice_toUpdate_speech;
let speech_volume = 1;
let speech_rate = 1;
let speech_pitch = 1;
let utteranceList = [];
//-----------------
let textLines = [];

//---------------------

//------------
let lastRow = -1;
let lastCol = -1;
let lastBoldCell;
let last_blue_cell;


//let ele_tab =document.getElementById("tab1"); 
/**
let myVoice;
let voices  = [];   // all voices from synth.. getVoices() 
**/
let listVox = [];   // selected voices only   
let voiceList2=[]; 
let totNumMyLangVoices=0; 
let lastNumVoice = 0; 
let lastWordNumVoice=0;
//---------------------

let eleTabSub_diff_clientW, eleTabSub_diff_clientH; 

let eleTabSub_save_clientWidth, eleTabSub_save_clientHeight;  
let sw_eleTabSub_widthInit = true; 
let t_swMove = true;
let language_parameters=["","",""]; // from Builder parameters
//---------------------------

//let speech = new SpeechSynthesisUtterance();
let synth  = window.speechSynthesis;

//------

let x1_line = 0;
//===================================
"use strict";
/*  
LineByLine: A tool to practice language comprehension
Antonio Cigna 2023
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//----------------------------------------------
var currScript = document.currentScript.src; var bar1 = currScript.lastIndexOf("\\");var bar2 = currScript.lastIndexOf("/"); 
//console.log("LOADED file SCRIPT " + currScript.substring( 1+Math.max(bar1,bar2) )) ;	
//----------------------------------------------------------------------------------------

function TOGLI1_onclick_tts_getOrigTxt() {
	
	var ele_inp = document.getElementById("inp_txtOrig00");	
	var input_text = ele_inp.value;
	
	if (is_srt(input_text) ) {
		var rigaFrom, rigaTo, rigaTxt, idsrt;
		[rigaFrom, rigaTo, rigaTxt, idsrt] = get_subtitle_strdata( input_text );
	    input_text = rigaTxt.join("\n").trim(); 
	} 
	
	var output_text;
	var maxRowLen = document.getElementById("id_maxLeng").value; 
	if (maxRowLen < 999) { 			
		output_text = tts_4_split_into_rows(input_text, maxRowLen);
	} else {
		output_text = input_text; 
	}
	
	//document.getElementById("inp_txtOrig00").value = output_text;
	document.getElementById("txt_pagOrig").value = output_text.trim();
	document.getElementById("page0").style.display = "none"; 
	document.getElementById("page1").style.display = "block"; 
	
	function is_srt(input_text1) {
		var numsrt = ( input_text1.replace("- ->", "-->").replace(" -> ", " --> ").replace("-- >", "-->") ).indexOf(" --> "); 
		return ( numsrt > 0); 
	}	
	
} // end of onclick_tts_getOrigTxt()

//---------------------------------------
function onchange_tts_get_oneLangVoice(this1) {
	
	var ix=-1;
	if ( typeof this1 == "object" ) { 
		ix = this1.value; 
	} else {
		ix = this1 
	}	
	
	//console.log("\t ix=", ix, "(listVox.length =" ,  listVox.length );
		
	if (listVox.length > 0) {  // if this is not the first language setting    
		//console.log("\nx\nx\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\nfill the voices again \nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n") 
		tts_2_fill_the_voices( "3onclick_tts_get_oneLangVoice" );  //  tts_9_toBeRunAfterGotVoicesPLAYER();
	}
	
	//console.log("\t voices[ix].lang=", voices[ix].lang );
	
	var langRegion ="";
	try{
		langRegion = get_languageName( voices[ix].lang ) ;
	} catch(e1) {
		return		
	}	
	
	
	var langname = langRegion.split("-")[0]; 
	document.getElementById("id_ext_language").innerHTML = voices[ix].lang + " " + langRegion; // in common/cbc_MESSAGE_manager.js 
	
	//document.getElementById("id_laName").innerHTML  =  langname ;
	document.getElementById("id_laName2").innerHTML =  langname ;
	//document.getElementById("m016-1").innerHTML  =  document.getElementById("m016").innerHTML.replace("§1§", langname);
	document.getElementById("m138-1").innerHTML  =  document.getElementById("m138").innerHTML.replace("§1§", langname);
	
	myVoice = voices[ix].lang + " " + voices[ix].name;  
	document.getElementById("id_myLang").innerHTML = myVoice;  
	selected_voice_ix        = ix; 	
	selected_voiceLangRegion =  voices[ix].lang	
	selected_voiceLang2      =  selected_voiceLangRegion.substr(0,2);
	selected_voiceName       =  voices[ix].name
	writeLanguageChoise()   // in wordsByFrequence.js 
	tts_3_show_speakingVoiceFromVoiceLangName(selected_voiceLangRegion, selected_voiceName);
		
	//console.log("onchange_tts_get_oneLangVoice() " + " selected_voice_ix=" + selected_voice_ix + "  id_my_lang = myVoice =" + myVoice);  
	
	isVoiceSelected = true; 
	tts_1_get_languagePlayer(ix); 
	   
} // 

//------------------------------------------------------------------------

function onclick_tts_getInput(elepage) {

  if (tts_1_join_orig_trad() < 0) {
	  return;
  }

  var elepage1 = document.getElementById("page1");
  var elepage2 = document.getElementById("page2");
  if (elepage == elepage1) {
	  elepage2.style.display = "block";
  } else {
	  elepage1.style.display = "block";
  }
  elepage.style.display = "none";
}
    
//-------------------------------------------------

function insert_missing_TR_m1_m2(eleTab, ele_tr_0, z3, Xprototype_tr_m1_tts, Xprototype_tr_m2_tts) {
	
	//console.log(" tabId=", eleTab.id , " tr=", ele_tr_0.tagName, " id=", ele_tr_0.id);  
	
	var idtr_0  = ele_tr_0.id; 
	var idtr_m1 = idtr_0 + "_m1";
	var idtr_m2 = idtr_0 + "_m2";
	
	//console.log("XXXXXXX  insert_missing_TR_m1_m2(eleTab, ele_tr_0, z3) eleTab.id=", eleTab.id, "  ele_tr_0.id=", ele_tr_0.id, "  z3=", z3); 
	//console.log("XXXXXXX  idtr_0 =" + idtr_0 + ",  m1=" + idtr_m1 + ", m2=" + idtr_m2); 
	
	var ele_tr_m1= document.getElementById(idtr_m1); 
	var ele_tr_m2= document.getElementById(idtr_m2); 
	
	if ( ele_tr_m1) { // TR m1 exists 
		return [true, ele_tr_0, ele_tr_m1, ele_tr_m2];  	
	} 	
	var msgEr = "error in 'insert_missing_TR_m1_m2(eleTab, ele_tr_0, z3)' "  + 'id="' + idtr_0 + '" id="' + idtr_m1 + '" id="' + idtr_m2+ '" ' ;
	
	var tr0_ix = ele_tr_0.rowIndex;
	
	//console.log("XXXXXXX  index tr0_ix =", tr0_ix); 
	
	//---------------------------------------------------------------------------------------
	var new_ele_tr_m1 = eleTab.insertRow(tr0_ix);
	ele_tr_m1 = new_ele_tr_m1;
	ele_tr_m1.id = idtr_m1;                                // TR m1 
	ele_tr_m1.style.display = "none"; 
	ele_tr_m1.style.borderBottom = "1px solid black";
	
	//console.log("XXXXXXX prima tr_1 = ", ele_tr_m1.outerHTML  ); 
	
	var newTr_m1_txt = Xprototype_tr_m1_tts.replaceAll("§1§",z3) ;
	
	// outerHTML su TR  non funziona, bisogna farlo cella per cella  	
	var ixS1 = newTr_m1_txt.indexOf("<td") 
	if (ixS1 < 0) {
		console.log("error in the prototype m1 ",newTr_m1_txt  , " ( <td not found");    
		return [false,,,]; 
	}
	var trM1_td = newTr_m1_txt.substr(  ixS1 ).replace("</tr>","").trim(); 
	var tdList1   = trM1_td.split("<td")
	for (var t=0; t< tdList1.length; t++) {
		var tdOuter = '<td' + tdList1[t];  
		var cell1 = ele_tr_m1.insertCell(-1);
		cell1.outerHTML = tdOuter	; 
	}
	
	//console.log("XXXXXXX dopo cells  TR M1  (outer) = ", ele_tr_m1.outerHTML  ); 
	
	if (!ele_tr_m1) {
		console.log(msgEr + ' <TR id="' + idtr_m1 + '"' + " just build, not found" ) 	
		return [false, ele_tr_0, ele_tr_m1, ele_tr_m2];  
	}
	//--------------------------------------------------------------------------------
	var tr_m1_ix = ele_tr_m1.rowIndex;
	var new_ele_tr_m2 = eleTab.insertRow(tr_m1_ix);
	ele_tr_m2 = new_ele_tr_m2; 
	ele_tr_m2.id = idtr_m2;                                // TR m1 
	ele_tr_m2.style.display = "none"; 
	
	if (z3 > 1) {
		ele_tr_m2.style.borderTop =  "1em solid grey";
	}
	//-------------------
	var newTr_m2_txt = Xprototype_tr_m2_tts.replaceAll("§1§",z3) 
	
	// outerHTML su TR  non funziona, bisogna farlo cella per cella  	
	var ixS2 = newTr_m2_txt.indexOf("<td") 
	if (ixS2 < 0) {
		console.log("error in the prototype m2 ",newTr_m2_txt  , " ( <td not found");    
		return [false,,,]; 
	}
	var trM2_td = newTr_m2_txt.substr(  ixS2 ).replace("</tr>","").trim(); 
	var tdList2   = trM2_td.split("<td")
	for (var t=0; t< tdList2.length; t++) {
		var tdOuter = '<td' + tdList2[t];  
		var cell1 = ele_tr_m2.insertCell(-1);
		cell1.outerHTML = tdOuter	; 
	}
	
	//console.log("XXXXXXX dopo cells  TR M2  (outer) = ", ele_tr_m2.outerHTML  ); 
	
	if (!ele_tr_m2) {
		console.log(msgEr + idtr_m2 + '"' + " just build, not found" ) 	
		return [false, ele_tr_0, ele_tr_m1, ele_tr_m2]; 
	} 
	return [true, ele_tr_0, ele_tr_m1, ele_tr_m2];
	
} // end of insert_missing_TR_M1

//-------------------------------------------------

function onclick_tts_arrowFromIx( this1, z3, wh) {
	//------------------------------------------
	// button from ( --> ) has been clicked 
	//------------------------------------------

	//console.log(" onclick_tts_arrowFromIx( this1, z3=" +z3 + ", wh=" + wh); 
	
	tts_5_removeLastBold(); 
	tts_5_fun_invisible_prev_fromto(-1);
	
	fromIxToIxLimit = [z3 ,-1]; 
	[begix, endix] =  fromIxToIxLimit;
	
	
	
	var isOk, eleTR_0, new_eleTR_0, eleTR_m1, eleTR_m2; 
	var eleTD  = this1.parentElement;
	var eleTR_0  = eleTD.parentElement;	
	
	var eleTable = eleTR_0.parentElement.parentElement;  
	
	//console.log("     eleTD=", eleTD.tagName, eleTD.id ,  " eleTR_0=", eleTR_0.tagName, eleTR_0.id,   " z3=", z3); 

	[isOk, new_eleTR_0, eleTR_m1, eleTR_m2] = insert_missing_TR_m1_m2(eleTable,eleTR_0, z3, prototype_tr_m1_tts, prototype_tr_m2_tts); 
	
    if (isOk == false) {
		console.log("error in 'onclick_tts_arrowFromIx() ... z3=" + z3 + " " + " wh="+wh + "  'insert_missing_TR_m1_m2()' " ) ;
		return; 
	}		
	//----------
	//tts_5_fun_copy_openClose_to_tr_m1(z3) ;  //  copy open/Close book style from this line idtr_xx to the upper idtr_xx_m1
	
	fromIxToIxButtonElement=[this1, null];
	this1.style.backgroundColor = "green";	

} // end of  onclick_tts_arrowFromIx()  

//---------------------------------------

function onclick_tts_arrowToIx( this1, z3, wh ) {
	//------------------------------------------
	// button to( <-- ) has been clicked 
	//------------------------------------------	
		
	var isOk, eleTR_0, new_eleTR_0, eleTR_m1, eleTR_m2; 
	var eleTD    = this1.parentElement;
	var eleTR_0  = eleTD.parentElement;
	
	var eleTable = eleTR_0.parentElement.parentElement;

	var idtr_0  = eleTR_0.id; 
	var ixPref = idtr_0.indexOf("_"); 
	var prefId = idtr_0.substr(0,ixPref+1); 
	var next_z3 = (z3+1);
	var next_id = prefId + next_z3; 
		
	var ele_nextTR = document.getElementById( next_id );	
	if (ele_nextTR) { 
		//console.log("   ele_nextTR=", ele_nextTR); 			
		var isOk, eleTR_0, new_eleTR_0, eleTR_m1, eleTR_m2; 		
		[isOk, new_eleTR_0, eleTR_m1, eleTR_m2] = insert_missing_TR_m1_m2(eleTable,ele_nextTR, next_z3, prototype_tr_m1_tts, prototype_tr_m2_tts); 

		if (isOk == false) {
			console.log("error in 'onclick_tts_arrowToIx() ... z3=" + z3 + " " + " wh="+wh + "  'insert_missing_TR_m1_m2()' " ) ;
			return; 
		}
	}
	
	//-----------------------------
	
	//----------
	//reset previous ..._ToIx  button     
		
	if (fromIxToIxButtonElement[1]) {
		fromIxToIxButtonElement[1].style.backgroundColor = null;
		var endix2 = fromIxToIxLimit[1]; 
		if (endix2 > 0)  {
			var id_post_tr_end_space2 = "idtr_" + (endix2+1) + "_m2" ; 
			if (document.getElementById(id_post_tr_end_space2  ) ) { document.getElementById(id_post_tr_end_space2 ).style.display = "none"; }
		}
	}
	//-------------------------------------------
	// new ... _ToIx
	
	fromIxToIxLimit[1]   = z3;  
	fromIxToIxButtonElement[1] = this1;
	
	[begix, endix] = fromIxToIxLimit;	
	
	/*
	this:  from id="b1_§1§" fromIx(§1§,this) -  to id="b2_§1§" 
		<tr id="idtr_§1§_m2"  ...>
		<tr id="idtr_§1§_m1"  ...>
	**/
	
	var id_pre_tr_beg_space  = "idtr_" + begix + "_m2" ; 
	var id_pre_tr_head       = "idtr_" + begix + "_m1" ; 
	var id_post_tr_end_space = "idtr_" + (endix+1) + "_m2" ; 
	
	tts_5_fun_copyHeaderSelected() ; 
	
	
	if (document.getElementById(id_pre_tr_beg_space ) ) {  
		document.getElementById(id_pre_tr_beg_space).style.display = "table-row";  
		if (begix < 2) 
			document.getElementById(id_pre_tr_beg_space).style.display = "none";  
	}
	//else {console.log(" manca "+id_pre_tr_beg_space ) }	
	if (document.getElementById(id_pre_tr_head      ) ) {  document.getElementById(id_pre_tr_head     ).style.display = "table-row"; }	
	//else {console.log(" manca "+id_pre_tr_head      ) }	
	if (document.getElementById(id_post_tr_end_space) ) { document.getElementById(id_post_tr_end_space).style.display = "table-row";}
	//else {console.log(" manca "+id_post_tr_end_space ) }	

	this1.style.backgroundColor="red";
	
	//c_onsole.log("onclick_tts_arrowToIx() calls 'fun_reset_clip_all_sel()'" )  ; 
	
	//21 novembre // fun_reset_clip_all_sel();
	
	
	
} // end of onclick_tts_arrowToIx()

//------------------------------------------

function onclick_tts_playSynthVoice_row2(this1, ixTD123, swPause,swNewVoice) {
	
	if (this1 == false) { return; }	
		

		//console.log("1 onclick_tts_playSynthVoice_row2()  "  , " totNumMyLangVoices=", totNumMyLangVoices  ); 
		if (tts_3_play_or_cancel(this1) < 0) {
				return;
			}	
			
		var ele_bookShow = document.getElementById("idb_" + ixTD123);	
		var showHideStyle = ele_bookShow.children[0].style.display;	
		tts_5_show_hideORIG2(ixTD123, showHideStyle) ;
	
		if (swNewVoice) {
			lastNumVoice++; 
			if (lastNumVoice >= totNumMyLangVoices) lastNumVoice = 0; // to change voice on each cycle
		} else {
			lastNumVoice = ele_select_first_language.selectedIndex; 
		}		
	    if (listVox.length == 0) {
			console.log("onclick_tts_playSynthVoice_row2()	listVox è vuoto")		
		}  
		//console.log("2 onclick_tts_playSynthVoice_row2() lastNumVoice=" + lastNumVoice,  "  listVox=", listVox);  
	    
		voice_toUpdate_speech = listVox[lastNumVoice][1]  ;  
		 
		 var td1 = this1.parentElement;
	     var tr1 = td1.parentElement;
		 
	     var txt1 = "",
	         txt2 = "";

		//console.log("3 onclick_tts_playSynthVoice_row2()"); 	
		
		 if (tr1.id.substr(0,5) == "idtr_") {
				var numId= tr1.id.substring(5);
				var ele_txt = document.getElementById("idc_"+ numId); 
				
				//console.log("??anto numId=", numId,  " ele_txt=" , ele_txt, " \n\tele_txtx.innerHTML=" , ele_txt.innerHTML );  
				if (ele_txt == null) return;
				var ele_tts = document.getElementById("idtts" + numId);
				txt1 = ele_txt.innerHTML; // cell[ index 5] = text 	; 
				txt2 = ele_tts.innerHTML; // text to speak				
				
				//console.log("??anto txt1=", txt1,  "   txt2=", txt2);  
				
				tts_3_boldCell(tr1, this1, 0, "onclick_tts_playSynthVoice_row");
				
		 }  else {
			    txt1 = ele_orig_line1.innerHTML;
				txt2 = ele_tts_line1.innerHTML;
		 } 
	     var txt3;
	     if ((txt2 == "") || (txt2 == "undefined"))  txt3 = txt1;
	     else txt3 = txt2;
		 txt3 =  html_to_txt( txt3 );
		 if (swPause) {		
			var ww1 =  (txt3+" ").replaceAll("–"," ").replaceAll("-"," ").
					replaceAll(", "," ").replaceAll(" ."," ").replaceAll(". "," ").replaceAll("..."," ").
					replaceAll("? "," ").replaceAll("! "," ").split(" "); 
			txt3 = "";
			for(var g=0; g < ww1.length; g++) {	
				var ww2 = ww1[g].trim(); 
				if (ww2 != "") txt3 += ww2 + ". "; 
			} 
		 }		

		/**
			if (numId == 5491) {console.log("4 onclick_tts_playSynthVoice_row2() txt1=" + txt1 + "\ntxt3="  + txt3 ); }
		**/	
		
		onclick_tts_text_to_speech2(txt3, 3);
		 
 } // end of onclick_tts_playSynthVoice_row

//-----------------------------------------------
function html_to_txt( ihtml ) {
	
	// eg. ithmn  = Thorium<span class="c_wordTarg"> eignet</span> sich auch für die Verwendung in Kernkraftwerken.
	// eg. return = Thorium eignet sich auch für die Verwendung in Kernkraftwerken.
	var len1= ihtml.length;
	var j1=0, j2=0, j3;
	var txt3 = ihtml;
	for(var z1=0; z1 < len1; z1++) {
		j1 = txt3.indexOf("<span");
		if (j1 < 0) { break;}	
		j2 = txt3.indexOf(">",j1);
		if (j2 < 0) { break;}
		j3 = txt3.indexOf("</span>",j2);
		if (j3 < 0) { break;}	
		txt3 = txt3.substring(0,j1) + txt3.substring(j2+1, j3) + txt3.substring(j3+7);	
	} 
 	return txt3.replaceAll("&nbsp;"," "); 
	//return txt3; 
	
} // end of html_to_text  

//-------------------------------------------------- 

function onclick_tts_playSynthVoice_m1_row2(this1, ixTD, swPause,swNewVoice) {

	//call_boldCell_ix(this1, ixVoice, "TTS_onclick_tts_ClipSub_Play3");

	var begix, endix; 
	[begix, endix] = fromIxToIxLimit;	
	
	var id_pref = "idc_" ;
	//------------	
	var id_pref_sp = id_pref; 
	var id_pref_idb = id_pref.replace("idc","idb");  
	
	//--------
	if (last_row_bold_ele) {
		last_row_bold_ele.classList.remove("boldLine"); 
        last_row_bold_ele.style.backgroundColor = null;
		last_row_bold_ele = null	
	}
	//------
		
	if (id_pref=="widc_") {
		id_pref_sp = "widtts_";
	} {
		if (id_pref == "idc_")  id_pref_sp = "idtts";  
	}
	
	TTS_LOOP_begix=begix;
	TTS_LOOP_endix=endix; 
	TTS_LOOP_swLoop=false; 	
	
    var txt3 = "";
    for (var t1 = begix; t1 <= endix; t1++) {
		var txtO = document.getElementById(id_pref + t1).innerHTML;
        var txtSp = document.getElementById(id_pref_sp + t1).innerHTML.trim();
		if (txtSp == "") {
			txtSp = txtO.trim(); 
		}
        txt3 += txtSp + " ";  // separatore di parole = " "    
    }
		   
		 if (swPause) {		
			var ww1 =  (txt3+" ").replaceAll("–"," ").replaceAll("-"," ").
					replaceAll(", "," ").replaceAll(" ."," ").replaceAll(". "," ").replaceAll("..."," ").
					replaceAll("? "," ").replaceAll("! "," ").split(" "); 
			txt3 = "";
			for(var g=0; g < ww1.length; g++) {	
				var ww2 = ww1[g].trim(); 
				if (ww2 != "") txt3 += ww2 + ". "; 
			} 
		 }
		 
		if (swNewVoice) {
			lastNumVoice++; 
			if (lastNumVoice >= totNumMyLangVoices) lastNumVoice = 0; // to change voice on each cycle
		} else {
			lastNumVoice = 0; 
		}
	
    onclick_tts_text_to_speech2(txt3, 5);  	

} // end of onclick_tts_playSynthVoice_m1_row2()

//------------------------------------------

function onclick_tts_playSynthVoice_m1_row3(this1, ixTD, swPause,swNewVoice) {
	var begix, endix; 
	
	var td1 = this1.parentElement;
	var tr1 = td1.parentElement;

	var id_tr0 = tr1.id.split("_");
	var numId = parseInt(id_tr0[1]); 	
	var pre_idtr= id_tr0[0]; 
	
	
	//call_boldCell_ix(this1, ixVoice, "TTS_onclick_tts_ClipSub_Play3");

	var pre_idtr1 = pre_idtr.substr(0,1); 	

	if (pre_idtr1 == "w") {
		[begix, endix] = word_fromIxToIxLimit
	} else {
		pre_idtr1 = ""; 
		[begix, endix] = fromIxToIxLimit;	
	}

	var txt1 = "",
	txt2 = "";

	var id_pref     = pre_idtr1 + "idc_"  ;
	var id_pref_idb = pre_idtr1 + "idb_"  ;  
	var id_pref_sp  = pre_idtr1 + "idtts_";
	
	
	TTS_LOOP_begix=begix;
	TTS_LOOP_endix=endix; 
	TTS_LOOP_swLoop=false; 	
	
	var txt3 = "";
    for (var t1 = begix; t1 <= endix; t1++) {
		var txtO  = document.getElementById(id_pref    + t1).innerHTML;
        var txtSp = document.getElementById(id_pref_sp + t1).innerHTML.trim();
		if (txtSp == "") {
			txtSp = txtO.trim(); 
		}
        txt3 += txtSp + " ";  // separatore di parole = " "    
    }

	if (swPause) {		
		var ww1 =  (txt3+" ").replaceAll("–"," ").replaceAll("-"," ").
		replaceAll(", "," ").replaceAll(" ."," ").replaceAll(". "," ").replaceAll("..."," ").
		replaceAll("? "," ").replaceAll("! "," ").split(" "); 
		txt3 = "";
		for(var g=0; g < ww1.length; g++) {	
			var ww2 = ww1[g].trim(); 
			if (ww2 != "") 
				txt3 += ww2 + ". "; 
		} 
	}

	if (swNewVoice) {
		lastNumVoice++; 
		if (lastNumVoice >= totNumMyLangVoices) 
			lastNumVoice = 0; // to change voice on each cycle
	} else {
		lastNumVoice = 0; 
	}

	onclick_tts_text_to_speech2(txt3, 6);  	

} // end of onclick_tts_playSynthVoice_m1_row3()
//--------------------------------

function onclick_tts_playSynthVoice_word3Freq(this1, wordToSay, swPause,swNewVoice) {
	
	if (tts_3_play_or_cancel(this1) < 0) {
		return;
	}	
    if (swNewVoice) {
		lastNumVoice++; 
		if (lastNumVoice >= totNumMyLangVoices) 
			lastNumVoice = 0; // to change voice on each cycle
	} else {
		lastNumVoice = 0; 
	}	

	var txt3 = wordToSay; 

	if (swPause) {	
		txt3 = tts_3_breakTextToPause( txt3, "word"); 
	}
	
	onclick_tts_text_to_speech2(txt3, 7);
		 
 } // end of onclick_tts_playSynthVoice_word3Freq

//-------------------------------------------------


//--------------------------------

function onclick_tts_playSynthVoice_word3(this1, ixTD123, swPause,swNewVoice) {

	if (tts_3_play_or_cancel(this1) < 0) {
		return;
	}	
    if (swNewVoice) {
		lastNumVoice++; 
		if (lastNumVoice >= totNumMyLangVoices) 
			lastNumVoice = 0; // to change voice on each cycle
	} else {
		lastNumVoice = 0; 
	}	

	var td1 = this1.parentElement;
	var tr1 = td1.parentElement;

	var txt1 = "",
	txt2 = "";

	var id_tr0 = tr1.id.split("_");
	var numId = parseInt(id_tr0[1]); 	
	var pre_idtr= id_tr0[0]; 
	var pre_idtr1 = pre_idtr.substr(0,1); 
	
	if (pre_idtr1 != "w") pre_idtr1 = "";
	
	var ele_txt = document.getElementById(pre_idtr1 + "idc_"+ numId); 
	
	if (ele_txt == null) return;
	
	var ele_tts = document.getElementById(pre_idtr1 + "idtts_" + numId);
	

	txt1 = ele_txt.innerHTML; // cell[ index 5] = text 	; 
	txt2 = ele_tts.innerHTML; // text to speak


	var txt3;
	if (txt2 == "") txt3 = txt1;
	else txt3 = txt2;

	if (swPause) {	
		txt3 = tts_3_breakTextToPause(txt3, pre_idtr1 ); 
	}
	
	onclick_tts_text_to_speech2(txt3, 7);
		 
 } // end of onclick_tts_playSynthVoice_word3

//-------------------------------------------------


function onclick_tts_text_to_speech2(txt1, wh) {	
	//console.log( "onclick_tts_text_to_speech2(", txt1 , ")" )
	sw_pause = false; 	
	startTime = new Date();	
	txt_length = txt1.length; 
	
	utteranceList = []; 
	x1_line=0; 
	txt1 = tts_3_removeBold_and_Font(txt1); 
    
	var newLine = tts_3_break_text(txt1, TXT_SPEECH_LENGTH_LIMIT, false);  // 3rd param true =  sw_hold_existing_endOfLine

	newLine = newLine.replaceAll("\n","\n\n");
	
    textLines = newLine.split("\n");
	
    var objtxt_to_speak;
    //var riga;
	//------------
    for (var v1 = 0; v1 < textLines.length; v1++) {		
		objtxt_to_speak = new SpeechSynthesisUtterance( textLines[v1]);		
        utteranceList.push(objtxt_to_speak);
    }
    objtxt_to_speak = utteranceList[0];
	
	//speech = objtxt_to_speak; // 1
  
	objtxt_to_speak.onend = tts_3_speech_end_fun;
	

	
	tts_3_speak_a_line(objtxt_to_speak,"1." + wh);	
	

} // end of onclick_tts_text_to_speech2() 	
//----------------------------------




//------------------------------------------------

function onclick_tts_change_voice(this1) {
	
    var vindex = this1.value;
	voice_toUpdate_speech = voiceList[vindex];	
	
	
	LS_voice_index = vindex; 	
	
	fun_set_localStorage_item_from_vars(); 	
	
    console.log("XXXX  voice: " + 
		" index=" + LS_voice_index +
		 " lang=" + voice_toUpdate_speech.lang +
		 " name=" + voice_toUpdate_speech.name ); 
	
} // end of  onclick_tts_change_voice(); 

//-------------------------------------------------

	
         function onclick_tts_changeRate(this1) {
			var rate = parseFloat(this1.value);
			if (rate < 0.30) rate = 0.30; 
			last_rate = rate; 
			this1.value = rate; 
			speech_rate = rate; 
			if (last_objtxt_to_speak) last_objtxt_to_speak.rate = rate; 
			var thisTD = this1.parentElement; 
			var preTD  = thisTD.parentElement.children[1]; 
			preTD.innerHTML = rate;  
         }
         //---------------------------------------
         function onclick_tts_changePitch(this1) {
         	var pitch = this1.value;
			if (pitch < 0.1) pitch = 0.1; 
			this1.value  = pitch;
         	speech_pitch = pitch;			
			if (last_objtxt_to_speak) last_objtxt_to_speak.rate = pitch; 
			var thisTD = this1.parentElement; 
			var preTD  = thisTD.parentElement.children[1]; 
			preTD.innerHTML = pitch;  
         } 
         //----------------------------------
		function onclick_tts_speech_pause() {
			//console.log("pause"); 
			if (synth.speaking) { 
				sw_pause = true; 
				synth.pause(); 
			}
		}
		//---------------------
		function onclick_tts_speech_resume() {
			//console.log("resume"); 
			sw_pause = false; 
			window.speechSynthesis.resume();
			synth.resume(); 
		}
				 
		 //----------------------
		 
         function onclick_tts_speech_cancel() {
			//console.log("*** cancel ***"); 
			sw_cancel = true; 
			if (synth.speaking) { 
				synth.cancel();
			}
         }                                                                       
         //------------------------------------

function onclick_tts_change_synth_rate(this1) {
	speech_rate = parseFloat(this1.value);
	if (speech_rate < 0.25) { 
		speech_rate = 0.25; 
	}
	document.getElementById("id_syncRate1").value = speech_rate;
	document.getElementById("id_setSpeedy").value = speech_rate;
}

//------------------------------------------------------

function onclick_tts_text_to_speech_ix(id_pref, ixWord, swLoop, this1) {
	
	var ele1 = document.getElementById(id_pref + ixWord);
	var txt1; 
	if (ele1) { 
		txt1 = ele1.innerHTML;
	} else { 
		return; 
	} 
	var tts_txt1 = txt1; 
	//  id="widc_4">coeli</div> <div style="display:none;" id="widtts_4">celi</div>
	
	
	
	if (id_pref=="widc_") {
		var eleTTS = document.getElementById("widtts_" + ixWord);
		if (eleTTS) {
			tts_txt1 = eleTTS.innerHTML;
		}	
	}   
	
	var id_pref_idb = id_pref.replace("idc","idb");  
	
	tts_showHide_if_book_opened(id_pref,id_pref_idb, ixWord) ;
	
	TTS_LOOP_begix   = ixWord;
	TTS_LOOP_endix   = ixWord; 
	TTS_LOOP_swLoop  = swLoop; 	
	TTS_LOOP_elem    = this1;
	
    onclick_tts_text_to_speech2(tts_txt1 ,1 );

} // end of onclick_tts_text_to_speech_ix()

//----------------------------------

function onclick_tts_show_rowOrig(this1, z3) {
	
	if (this1 == false) { return; }	
	
	if (this1.children[0].style.display == "none") {  // no openbook   
		this1.children[0].style.display = "block";                         // show opened book image  
		this1.children[1].style.display = "none";	
	} else {
		this1.children[0].style.display = "none";                          //  hide opened book image  
		this1.children[1].style.display = "block";						  //  show closed book image 	
	}	
    // when the book is open the text is visibile otherwise it's hided  
	// last text made visible is highlited in a yellow background 	
	var showHideStyle = this1.children[0].style.display
	
	tts_5_show_hideORIG2(z3, showHideStyle);
	
} // end of onclick_tts_show_rowOrig	
//----------------------
function tts_5_show_hideORIG2(z3, showHideStyle) {	
	let ele1 = document.getElementById("idc_" + z3); // element of original text to show/hide	
	ele1.style.display      = showHideStyle;
	
    if (last_row_bold_ele) {
		last_row_bold_ele.classList.remove("boldLine"); 
        last_row_bold_ele.style.backgroundColor = null;
		last_row_bold_ele = null	
	}
	if (ele1.style.display == "block") { // openbook ==> show 
		ele1.classList.add("boldLine");
		ele1.style.backgroundColor = "yellow";	
		last_row_bold_ele = ele1; 
    } else {
		ele1.classList.remove("boldLine"); 
        ele1.style.backgroundColor = null;
		last_row_bold_ele = null	
	}
	
} // end of tts_5_show_hideORIG2
//-------------------------------------------------

function onclick_tts_show_rowTran(this1, z3) {
	
	if (this1 == false) { return; }	
	
	if (this1.children[0].style.display == "none") {  // no openbook   
		this1.children[0].style.display = "block";                         // show T ( see translation  
		this1.children[1].style.display = "none";	
	} else {
		this1.children[0].style.display = "none";                          // 
		this1.children[1].style.display = "block";						   // show t? hide translation   	
	}
	// when the image shown is the capital letter T,  the text translation is visibile otherwise it's hided  	
	let ele1 = document.getElementById("idt_" + z3); // element of translated text 	      
	ele1.style.display      =  this1.children[0].style.display;
	
} // end of onclick_tts_show_rowTran
	
//-------------------------------------------------------------------------

function onclick_tts_word_show_one_row(this1, z3) {
	
	if (this1 == false) { return; }	
	
	if (this1.children[0].style.display == "none") {  // no openbook   
		this1.children[0].style.display = "block";                         // show opened book image  
		this1.children[1].style.display = "none";						  // hide closed book image 	
	} else {
		this1.children[0].style.display = "none";                          //  hide opened book image  
		this1.children[1].style.display = "block";						  //  show closed book image 	
	}
	// when the book is open the text is visibile otherwise it's hided  
	// last text made visible is highlited in a yellow background 	
		
	let ele1 = document.getElementById("widc_" + z3); // element of original text to show/hide	
	ele1.style.display      =  this1.children[0].style.display;
	
	
} // end of onclick_tts_word_show_one_row

//----------------------------

function onclick_showHide_orig_row_group( this1 ) {
	
	if (this1 == false) { return; }
		
	if (this1.children[0].style.display == "none") {  // no openbook   
		this1.children[0].style.display = "block";                         // show opened book image  
		this1.children[1].style.display = "none";	
	} else {
		this1.children[0].style.display = "none";                          //  hide opened book image  
		this1.children[1].style.display = "block";						  //  show closed book image 	
	}	
    // when the book is open the text is visibile otherwise it's hided  
	
	var showHideStyle = this1.children[0].style.display
	var begix, endix; 
	[begix, endix] = fromIxToIxLimit;
	if ((begix < 0) || (endix < begix)) { return ; } 
	
	var style0 = this1.children[0].style.display
	var style1 = this1.children[1].style.display
	var ele_idb, ele_idc;
	for(var g=begix; g <= endix; g++) {
		ele_idb = document.getElementById("idb_" + g);      // book opened/closed
		ele_idb.children[0].style.display = style0;  	
		ele_idb.children[1].style.display = style1;  
		ele_idc = document.getElementById("idc_" + g);      // row (orig) visible or hided  
		ele_idc.style.display = style0;  		
        ele_idc.style.backgroundColor = null;
		ele_idc.classList.remove("boldLine");
	} 
	//   antox tts_5_show_hideORIG2(z3, showHideStyle);
	
} // end of onclick_showHide_orig_row_group
//-----------------------------------------------------

function onclick_showHide_tran_row_group( this1 ) {
	
	if (this1 == false) { return; }
		
	if (this1.children[0].style.display == "none") {  // no openbook   
		this1.children[0].style.display = "block";                        // show T image   ( show translation )
		this1.children[1].style.display = "none";	
	} else {
		this1.children[0].style.display = "none";                         // 
		this1.children[1].style.display = "block";						  //  show t? image ( hide translation ) 	
	}	
    // when the T is shown the translation is visibile otherwise it's hided  
	
	var showHideStyle = this1.children[0].style.display
	var begix, endix; 
	[begix, endix] = fromIxToIxLimit;
	if ((begix < 0) || (endix < begix)) { return ; } 
	/**
	                             "idbT_ " 
	<button class="buttonTD2" id="idbT_3" onclick="onclick_tts_show_rowTran(this, 3)">
				   <span style="display:none;font-size:2em;height:1.4em; "><span style="font-weight:bold;"><span style="font-weight:bold;">T</span></span></span>
				   <span style="display:block;font-size:2em;height:1.4em;padding:0 0.1em;"><span style="font-weight:bold;min-width:4em;"><span style="font-weight:bold;min-width:4em;">t?</span></span></span>
	</button>
	--- 
	                                                "idt_ "
	<div class="tranLine" style="display: none;" id="idt_3">Per scoprire come vivevano le persone in quel periodo,<br></div>	
	**/
	var style0 = this1.children[0].style.display
	var style1 = this1.children[1].style.display
	var ele_idb, ele_idt;
	for(var g=begix; g <= endix; g++) {
		ele_idb = document.getElementById("idbT_" + g);      // T / t?
		ele_idb.children[0].style.display = style0;  	
		ele_idb.children[1].style.display = style1;  
		ele_idt = document.getElementById("idt_" + g);      // translation visible (T) or hided (t?) 
		ele_idt.style.display = style0;  	
	} 
	//   antox tts_5_show_hideORIG2(z3, showHideStyle);
	
} // end of onclick_showHide_tran_row_group

//----------------------------------------------

function onclick_showHide_orig_word_group( this1 ) {
	
	if (this1 == false) { return; }
		
	if (this1.children[0].style.display == "none") {  // no openbook   
		this1.children[0].style.display = "block";                         // show opened book image  
		this1.children[1].style.display = "none";	
	} else {
		this1.children[0].style.display = "none";                          //  hide opened book image  
		this1.children[1].style.display = "block";						  //  show closed book image 	
	}	
    // when the book is open the text is visibile otherwise it's hided  
	
	var showHideStyle = this1.children[0].style.display
	var begix, endix; 
	[begix, endix] = word_fromIxToIxLimit; 
	if ((begix < 0) || (endix < begix)) { return ; } 
	
	var style0 = this1.children[0].style.display
	var style1 = this1.children[1].style.display
	var ele_idb, ele_idc;
	for(var g=begix; g <= endix; g++) {
		ele_idb = document.getElementById("widb_" + g);      // book opened/closed
		ele_idb.children[0].style.display = style0;  	
		ele_idb.children[1].style.display = style1;  
		ele_idc = document.getElementById("widc_" + g);      // row (orig) visible or hided  
		ele_idc.style.display = style0;  		
        ele_idc.style.backgroundColor = null;
		ele_idc.classList.remove("boldLine");
	} 
	//   antox tts_5_show_hideORIG2(z3, showHideStyle);
	
} // end of onclick_showHide_orig_word_group

//---------------------------------------------------

function onclick_tts_word_arrowFromIx(ele_td_arrow, z3, isWord, is_m1) {
    //------------------------------------------
    // button from ( --> ) has been clicked 
    //------------------------------------------
	
	// se TR m1 e m2 mancano vengono inseriti
	var isOk, eleTR_0, new_eleTR_0, eleTR_m1, eleTR_m2; 
	var eleTD  = ele_td_arrow
	if (eleTD.tagName != "TD") { eleTD = ele_td_arrow.parentElement;}
	var eleTR_0  = eleTD.parentElement;	
	
	
	var eleTable = eleTR_0.parentElement.parentElement;  
	
	//console.log("     eleTD=", eleTD.tagName, eleTD.id ,  " eleTR_0=", eleTR_0.tagName, eleTR_0.id,   " z3=", z3); 
	[isOk, new_eleTR_0, eleTR_m1, eleTR_m2] = insert_missing_TR_m1_m2(eleTable,eleTR_0, z3, prototype_word_tr_m1_tts, prototype_word_tr_m2_tts); 
	
    if (isOk == false) {
		console.log("error in 'onclick_tts_word_arrowFromIx() ... z3=" + z3 + " " + " wh="+wh + "  'insert_missing_TR_m1_m2()' " ) ;
		return; 
	}		
	//----------
	
    tts_3_word_removeLastBold(isWord);
    tts_3_word_fun_invisible_prev_fromto(-1, isWord, is_m1 );
	
	//--------------------
	
		
	if (isWord) {
		word_fromIxToIxLimit = [z3, -1];
		[begix, endix] = word_fromIxToIxLimit; 
		tts_3_word_fun_copy_openClose_to_tr_m1(z3, isWord, is_m1); //  copy open/Close book style from this line idtr_xx to the upper idtr_xx_m1
		word_fromIxToIxButtonElement = [ele_td_arrow, null];
		ele_td_arrow.style.backgroundColor = "green";
	} else {
		fromIxToIxLimit = [z3, -1];
		[begix, endix] = fromIxToIxLimit;
		tts_3_word_fun_copy_openClose_to_tr_m1(z3, isWord, is_m1); //  copy open/Close book style from this line idtr_xx to the upper idtr_xx_m1
		fromIxToIxButtonElement = [ele_td_arrow, null];
		ele_td_arrow.style.backgroundColor = "green";
	}

} // end of onclick_tts_word_arrowFromIx  

//------------------------------------------

function onclick_tts_word_arrowToIx(ele_td_arrow, z3, isWord, is_m1) {
	//------------------------------------------
	// button to( <-- ) has been clicked 
	//------------------------------------------	
	// se TR m1 e m2 mancano vengono inseriti
	var isOk, eleTR_0, new_eleTR_0, eleTR_m1, eleTR_m2; 
	var eleTD  = ele_td_arrow; 
	if (eleTD.tagName != "TD") { eleTD = ele_td_arrow.parentElement;}
	var eleTR_0  = eleTD.parentElement;		
	
	var eleTable = eleTR_0.parentElement.parentElement;  

	var idtr_0  = eleTR_0.id; 
	var ixPref = idtr_0.indexOf("_"); 
	var prefId = idtr_0.substr(0,ixPref+1); 
	var next_z3 = (z3+1);
	var next_id = prefId + next_z3; 
		
	var ele_nextTR = document.getElementById( next_id );	
	if (ele_nextTR) { 
		//console.log("   ele_nextTR=", ele_nextTR); 			
		var isOk, eleTR_0, new_eleTR_0, eleTR_m1, eleTR_m2; 
		[isOk, new_eleTR_0, eleTR_m1, eleTR_m2] = insert_missing_TR_m1_m2(eleTable,ele_nextTR, next_z3, prototype_word_tr_m1_tts, prototype_word_tr_m2_tts); 

		if (isOk == false) {
			console.log("error in 'onclick_tts_arrowToIx() ... z3=" + z3 + " " + " wh="+wh + "  'insert_missing_TR_m1_m2()' " ) ;
			return; 
		}
	}	
	//-----
	//---------------------------------------------------
    var endix2; var id_post_tr_end_space2;
	var id_pre_tr_beg_space;
	var id_pre_tr_head ;
	var id_post_tr_end_space ;	
	
	if (isWord) {
		//reset previous ..._ToIx  button     
		if (word_fromIxToIxButtonElement[1]) {
			word_fromIxToIxButtonElement[1].style.backgroundColor = null;
			endix2 = word_fromIxToIxLimit[1];
			if (endix2 > 0) {
				id_post_tr_end_space2 = "widtr_" + (endix2 + 1) + "_m2";
				if (document.getElementById(id_post_tr_end_space2)) {
					document.getElementById(id_post_tr_end_space2).style.display = "none";
				}
			}
		}
		//--------- new ... _ToIx ----------------------------------
		if (z3 <= word_fromIxToIxLimit[0]) {
			// this set the arrows 
			var eleFromArrow = word_fromIxToIxButtonElement[0];
			var eleTr = eleFromArrow.parentElement.parentElement; 
			console.log("eleTR " + eleTr.tagName + " id="+ eleTr.id); 
			var preEleTr = eleTr.previousElementSibling;  
			var preEleTr2 = preEleTr.previousElementSibling;  
			console.log("preEleTR " + preEleTr.tagName + " id="+ preEleTr.id); 
			eleFromArrow.style.backgroundColor = null;
			preEleTr.style.display = "none";
			preEleTr2.style.display = "none";
			return; 
		}
		
		word_fromIxToIxLimit[1] = z3;
		word_fromIxToIxButtonElement[1] = ele_td_arrow;
		[begix, endix] = word_fromIxToIxLimit;
		id_pre_tr_beg_space  = "widtr_" + begix + "_m2";
		id_pre_tr_head       = "widtr_" + begix + "_m1";
		id_post_tr_end_space = "widtr_" + (endix + 1) + "_m2";
		tts_3_word_fun_copyHeaderSelected(begix, endix);
	} else {
		//reset previous ..._ToIx  button     
		if (fromIxToIxButtonElement[1]) {
			fromIxToIxButtonElement[1].style.backgroundColor = null;
			endix2 = word_fromIxToIxLimit[1];
			if (endix2 > 0) {
				id_post_tr_end_space2 = ele_tr_idPref + (endix2 + 1) + "_m2";
				if (document.getElementById(id_post_tr_end_space2)) {
					document.getElementById(id_post_tr_end_space2).style.display = "none";
				}
			}
		}
		//--------- new ... _ToIx -----
		fromIxToIxLimit[1] = z3;
		fromIxToIxButtonElement[1] = ele_td_arrow;
		[begix, endix] = word_fromIxToIxLimit;
		id_pre_tr_beg_space = "widtr_" + begix + "_m2";
		id_pre_tr_head = "widtr_" + begix + "_m1";
		id_post_tr_end_space = "widtr_" + (endix + 1) + "_m2";
		fun_copyHeaderSelected(begix, endix);
	}
	
    if (document.getElementById(id_pre_tr_beg_space)) {
        document.getElementById(id_pre_tr_beg_space).style.display = "table-row";
    }   
    if (document.getElementById(id_pre_tr_head)) {
        document.getElementById(id_pre_tr_head).style.display = "table-row";
    }
    if (document.getElementById(id_post_tr_end_space)) {
        document.getElementById(id_post_tr_end_space).style.display = "table-row";
    }   
    ele_td_arrow.style.backgroundColor = "red";
   
} // end of onclick_tts_arrowToIx()

//------------------------------------------

function onclick_tts_word_OneClipRow_showHide_sub( ele_idb, sw_allSel, swAllAll, isWord, is_m1) {	
	
	if (ele_idb == false) { return; }		
	
	//console.log("onclick_tts_word_OneClipRow_showHide_sub() 1 ele_idb ", ele_idb.id); 
	
	let id1;
	let inBeg, inEnd; 
	inBeg      = begix;
	inEnd      = endix; 
	if (isWord == false) { 
		if (swAllAll) {
			inBeg= clipFromRow_min;
			inEnd= line_list_o_from1.length-1; 
		}
	}
	if (begix > endix) {
		inBeg  = endix;
		inEnd  = begix; 		
	} 	 
	var style0 , style1; 	
	if (isWord) {
		word_fun_oneRow00(); 	
	} else {
		if (swAllAll) {
			if (ele_idb.children[0].style.display == "block") {  // openbook   
				style0 = "block";                         // show opened book image  
				style1 = "none";						  // hide closed book image 	
			} else {
				style0 = "none";                          //  hide opened book image  
				style1 = "block";						  //  show closed book image 	
			}
		} else {
			word_fun_oneRow00(); 	
		}
	}
	//console.log("onclick_tts_word_OneClipRow_showHide_sub() 2 ele_idb ", ele_idb.id); 
	word_fun_oneRow22(1); 
	
	if (sw_allSel) {	
		for(var g=inBeg; g <= inEnd; g++) {
			if (isWord) id1 = "widb_" + g;   
			else id1 = "idb_" + g;   			
			ele_idb = document.getElementById(id1); 
			//console.log("onclick_tts_word_OneClipRow_showHide_sub() 2 ", " id1=", id1 ); console.log("\t ele_idb ", ele_idb.id); 
			word_fun_oneRow22(1+inEnd-inBeg); 	
		} 
	} 	
	//--------------
	function word_fun_oneRow00() {
		if (ele_idb == false) { return; }
		
		if (ele_idb.children[0].style.display == "none") {  // no openbook   
			style0 = "block";                         // show opened book image  
			style1 = "none";						  // hide closed book image 	
		} else {
			style0 = "none";                          //  hide opened book image  
			style1 = "block";						  //  show closed book image 	
		}
		
	}	
	//-------------------  
	function word_fun_oneRow22(nn) {	// 2 onclick_tts_word_OneClipRow_showHide_sub
		
		
		if (ele_idb == false) { return; }
		
		ele_idb.children[0].style.display = style0;         // show/hide  opened book image  
		ele_idb.children[1].style.display = style1; 		  // show/hide closed book image 
		let subid = ele_idb.id.replace("idb","idc"); 		
		//if (is_m1) return;  
		let ele_idc = document.getElementById( subid );
		if (style0 == "block") {
			tts_3_word_fun_makeTextVisible(ele_idc);  
		} else {		
			tts_3_word_fun_makeTextInvisible(ele_idc);
		}
	} // end of word_fun_oneRow22()
	//-------------------------
	
} // end of onclick_tts_word_OneClipRow_showHide_sub()  // 2 

//-----------------------------

//--------------------------------------------------------------
function onclick_tts_playSynthVoice_m1_row(this1, numVoice) {
	TTS_onclick_tts_ClipSub_Play3(this1,numVoice) ;
}
//-----------------------------------------------------
function onclick_tts_ClipSub_Play3(this1,numVoice) {
	
	var ixVoice = numVoice - 1;
		
	voice_toUpdate_speech = listVox[ixVoice][1]; 
	
	call_boldCell_ix(this1, ixVoice, "TTS_onclick_tts_ClipSub_Play3");



	var begix, endix; 
	[begix, endix] = fromIxToIxLimit;	
	
	console.log("\t??anto? file ...existing.. file 2 TTS_onclick_tts_ClipSub_Play3() begix=" + begix + " endix=" + endix); 
	
	
	onclick_tts_text_to_speech_from_to("idc_",begix, endix, false, "2 TTS_on...Play3");
	console.log("uscito da onclick_tts_text_to_speech_from_to()"); 
	
} // end of TTS_onclick_tts_ClipSub_Play3()

//----------------------
function onclick_tts_playSynthVoice_m1_word( this1,numVoice) {
		onclick_tts_word_ClipSub_Play3(      this1, false, true, true, numVoice); 
}
//------------------------
function onclick_tts_word_ClipSub_Play3(this1, swLoop, isWord, is_m1, numVoice) {
	var ixVoice = numVoice - 1;
	
	if (ixVoice) {
		voice_toUpdate_speech = listVox[ixVoice][1] ; 
	}
	
    var sw_tts2 = (sw_tts || isWord);  
    var begix, endix;
	if (isWord) {	
		[begix, endix] = word_fromIxToIxLimit;
		if (sw_tts2) { // TTS 
			if (TTS_LOOP_swLoop) {
				TTS_LOOP_swLoop = false;
				return;
			}
			if (word_play_or_cancel(this1) < 0) {
				return;
			}
			onclick_tts_text_to_speech_from_to("widc_", begix, endix, swLoop, "3 word_onclip...Play3");
			return;
		}		
    } else { 	
		[begix, endix] = fromIxToIxLimit;
		if (sw_tts2) { // TTS 
			if (TTS_LOOP_swLoop) {
				TTS_LOOP_swLoop = false;
				return;
			}
			if (word_play_or_cancel(this1) < 0) {
				return;
			}
			onclick_tts_text_to_speech_from_to("idc_", begix, endix, swLoop, "4 word_onclip...Play3");
			return;
		}		
	}

} // end of onclick_tts_word_words1_Play3()

//-------------------------------------------------

function onclick_tts_langSelected() {

	document.getElementById("page0Lang").style.display = "none"; 
	document.getElementById("page0"    ).style.display = "block"; 

} // end of onclick_tts_langSelected()

//-------------------------------------------------
"use strict";
/*  
LineByLine: A tool to practice language comprehension
Antonio Cigna 2023
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//----------------------------------------------
var currScript = document.currentScript.src; var bar1 = currScript.lastIndexOf("\\");var bar2 = currScript.lastIndexOf("/"); 
//console.log("LOADED file SCRIPT " + currScript.substring( 1+Math.max(bar1,bar2) )) ;	
//----------------------------------------------------------------------------------------

  var page1 = document.getElementById("page1");
  var page2 = document.getElementById("page2");

 
//---------------------------------	
  var builder_orig_subtitles_string = "";
  var builder_tran_subtitles_string = "";

  var sw_inp_sub_orig_builder = false;
  var sw_inp_sub_tran_builder = false;

  const TRANSLATION_NOT_WANTED = "...";
 
  var sw_translation_not_wanted = false;
  var selected_voice_language = "XX,0,0"; //  language id, number of voices eg. en-GB Microsoft ...  , index of chosen voice 
  
  //===============================
  var myVoice;
  let voices;
  //----------------------


  //fcommon_load_all_voices();  // at end calls tts_1_toBeRunAfterGotVoices()

  // WARNING: the above function contains asynchronous code.  
  // 			Any statement after this line is executed immediately without waiting its end


 // test voice with previous run
 function testPrevVoice() {

	if (prev_voiceName == "") { return -1; }
	
	for (var z1 = 0; z1 < voices.length; z1++) {
		if (voices[z1].name == prev_voiceName) return z1;	
	}	
	for (var z1 = 0; z1 < voices.length; z1++) {
		if ( voices[z1].lang == prev_voiceLangRegion) return z1	;
	}		
	for (var z1 = 0; z1 < voices.length; z1++) {
		if ( voices[z1].lang.substring(0,2) == prev_voiceLang2) return z1;	
	}	
	return -1;
 }
 
 //-------------------------------------------------
  function tts_1_toBeRunAfterGotVoices() {
	  //console.log("tts_1_toBeRunAfterGotVoices() 1 ")
	  //console.log("tts_1_toBeRunAfterGotVoices() 2 " , " voices.length=", voices.length)
      if (voices.length < 1) return;

      voices.sort(
          function(a, b) {
              return ((a.lang + a.name) > (b.lang + b.name)) ? 1 : -1;
          }
      );
	  
	  //------------------	   
      var pLang2 = "??";
      var pLang4;
      var numL2 = 0;
      var numL4 = 0;
	  var langName;
      var ixSele = -1;
	  var sele_str = "";
	  var selected_yes="selected"; 
	  var swSele=false; 
	  var voice_name =""; 

	  var ixPrevFound = -1 
	  
	  
	  if (prev_voiceName != "") {  
			ixPrevFound = testPrevVoice() 
	  }
	  
	  /**
	  if (ixPrevFound < 0) {
			console.log("tts_1_toBeRunAfterGotVoices()  ixPrevFound not found"); 
	  } else {  
			console.log("tts_1_toBeRunAfterGotVoices()  ixPrevFound=" + ixPrevFound, "\nlastRunLanguage=" +  lastRunLanguage);
	  }
	  **/
	  
	  //sele_str += '      <option value="' + 9999+ '" selected>' + '--' + " " + "----------"  +  '</option> \n';
	  
      for (var z1 = 0; z1 < voices.length; z1++) {
		  swSele = false; 
          var lang = voices[z1].lang;
          var lang2 = lang.substr(0, 2);
          if (lang != pLang4) {
              //console.log(lang);
              numL4++;
              if (lang2 != pLang2) {
				  langName = get_languageName(lang).split("-")[0].trim(); 
                  numL2++;
                  sele_str += '   </optgroup> \n';
                  sele_str += '   <optgroup label="' + lang2 + ' ' + langName + '"> \n';
              }
          }
		  /**
		  if (lang == "NONEen-GB") {
				if (ixPrevFound < 0) {
					if (ixSele < 0) {
						ixSele=z1;
					}	
				}
		  }
		  **/
		  
		  voice_name = voices[z1].name;
	  	 
		  if ((ixSele == z1 ) || (ixPrevFound == z1)) {
				sele_str += '      <option value="' + z1 + '" selected>' + lang + " " + voice_name +  '</option> \n';
		  } else {
			  sele_str += '      <option value="' + z1 + '">' + lang + " " + voice_name +  '</option> \n';
		  }	
		 
          pLang4 = lang;
          pLang2 = lang2;
      }	  
	  
	  pLang4 = lang;
      sele_str += '   </optgroup> \n';
	  
      let voiceSelect = document.getElementById("id_voices");
      voiceSelect.innerHTML = sele_str;
	  	
	  for(var g1=0; g1 < voiceSelect.children.length; g1++) {
		  var gr1 = voiceSelect.children[g1]; 
		  //console.log("voiceSelect " + g1 + " gr1= " + gr1.outerHTML);    
	  }
	  
	  if (ixPrevFound < 0) {
		   ixPrevFound=0;	  
		   
	  }	 
	  voiceSelect.selectedIndex = ixPrevFound   
	  
	  
		selected_voice_ix        = ixPrevFound   ; 	
		selected_voiceLangRegion =  voices[ixPrevFound] .lang	
		selected_voiceLang2      =  selected_voiceLangRegion.substr(0,2);
		selected_voiceName       =  voices[ixPrevFound].name
	 
	  onchange_tts_get_oneLangVoice(  voiceSelect );	 
	  /**
	  if (ixPrevFound < 0) {
		  onchange_tts_get_oneLangVoice(  voiceSelect );	  
	  } else {
		  document.getElementById("id_divVoice").style.display = "none"; 
		  onchange_tts_get_oneLangVoice(  ixPrevFound );	  
	  }	  	  
   	  **/
	  
  } // end of toBeRunAfterGotVoices()

  //===============
  //--------------------------------------------------
 function tts_1_get_orig_subtitle2() {
      /**
	
         console.log("get_orig" + "\n\t value=" +  document.getElementById("txt_pagOrig").value +
         						"\n\t fromText_to_srt=" +  fromText_to_srt(   document.getElementById("txt_pagOrig").value.trim() )  ); 
         **/
      var msgerr1 = "";
    
      builder_orig_subtitles_string = document.getElementById("txt_pagOrig").value.trim();
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

  function tts_1_get_tran_subtitle2(msgerrOrig) {
      /**
      	console.log("get_tran" + "\n\t value=" +  document.getElementById("txt_pagTrad").value +
      						"\n\t fromText_to_srt=" +  fromText_to_srt(    document.getElementById("txt_pagTrad").value.trim() )  ); 
      **/
      var msgerr1 = "";

      //builder_tran_subtitles_string = fromText_to_srt(  document.getElementById("txt_pagTrad").value.trim() );
      builder_tran_subtitles_string = document.getElementById("txt_pagTrad").value.trim();
      sw_inp_sub_tran_builder = false;
      if (builder_tran_subtitles_string != "") {
          sw_inp_sub_tran_builder = true;
      } else {
          sw_inp_sub_tran_builder = false;
          if (sw_translation_not_wanted == false) {
			  //console.log("anto1   msgerrOrig=" + msgerrOrig)
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
  function tts_1_join_orig_trad() {
	  
	  console.log("tts_1_join_orig_trad()") 
	  
      document.getElementById("id_msg16").innerHTML = "";

      sw_translation_not_wanted = false;

      var msgerr0 = "";
      var msgerr1 = tts_1_get_orig_subtitle2(); // get orig. text /srt
      var msgerr2 = tts_1_get_tran_subtitle2(msgerr1); // get tran. text/srt	

      msgerr0 += msgerr1 + msgerr2;
      var msgerr3 = "";

      msgerr0 += msgerr3;

      if (msgerr0 != "") {
          tts_1_putMsgerr(msgerr0);
          document.getElementById("id_msg16").style.color = "red";
          return -1;
      }
      document.getElementById("id_msg16").style.color = null;

      //console.log("ORIG =" + builder_orig_subtitles_string);
      //console.log("TRAN =" + builder_tran_subtitles_string);
	  
	  //tts_1_get_wantedVoices_X();  	
	  
	  tts_9_toBeRunAfterGotVoicesPLAYER(); 

      return 0;

  } // end of tts_1_join_orig_trad()

  //----------------------

  function tts_1_putMsgerr(msgerr1) {
      if (msgerr1 == "") {
          document.getElementById("id_msg16").style.display = "none";
      } else {
          if (msgerr1.substring(0, 4) == "<br>") {
              msgerr1 = msgerr1.substring(4);
          }
          document.getElementById("id_msg16").innerHTML = msgerr1;
          document.getElementById("id_msg16").style.backgroundColor = "white";
          document.getElementById("id_msg16").style.display = "block";
      }

  } // end of putMsgerr()

  //--------------------------------

  function tts_1_getMsgId(id1) {  
	var ele1 = document.getElementById(id1);  
	if (ele1 == null) {		
		console.log("msg " + id1 + "  tts_1_getMsgId() non trovato"  );
		return "";
	}	
    return ele1.innerHTML;
  }
//-------------------------------------
/*
	let selected_voice_ix                    = 0 ;     // eg. 65 	 
	let selected_voiceName                   = "";     // eg. Microsoft David - English (United States)"; 	
	let	selected_voiceLangRegion             = "";     // eg. en-us	
	let	selected_voiceLang2      
*/  
//-------------------------

function tts_1_get_wantedVoices_X() { 
 
	if (selected_voice_language == "null") {
		language_parameters = ["xx","0","0"];  
		selected_lang_id = "xx"; 
		selected_numVoices=0;  	
	} else { 	
		language_parameters = (selected_voce_language+",,,").split(","); 
		selected_lang_id    = (language_parameters[0]+"   ").trim(); // .substr(0,2); 
		selected_numVoices  = parseInt("0" + language_parameters[1].trim() )  ; 		
			
	}
	if (selected_numVoices > maxNumVoices)  selected_numVoices = maxNumVoices;
	
	console.log("parameters from Builder: language '" + selected_lang_id  + "' num.Voices=" + selected_numVoices); 
	
} // end of  get_wantedVoices_X()

//-------------------------------------
function tts_1_get_languagePlayer(ix) {

	var aVoice = document.getElementById("id_myLang").innerHTML;
	var voi = aVoice.split(" ");
	var voice2 = voi[0]; 	
	var voiceL = voice2.length;
	var nn=0;
	for(var g=0; g < voices.length; g++) {
		var myVoice = voices[g].lang;
		if (myVoice.substr(0,2) == voice2.substr(0,2) ) {			
			nn++; 
		}  
	}
	selected_voice_language = voice2 + "," + nn + "," + ix; 
	var lan1 = document.getElementById("m002").innerHTML ;
	var lan2 = document.getElementById("m003").innerHTML
	document.getElementById("id_lang2").innerHTML = "  (" + lan1 + " " + voice2 + ", " + nn + " " + lan2 + ")" ;
	
	//console.log("tts_1_get_languagePlayer(ix) " + " voices.length=" + voices.length + 
	//	"\n\t  selected_voice_language=" + selected_voice_language + "\n\t id_lang2 = " +document.getElementById("id_lang2").innerHTML ); 
	
	writeLanguageChoise(); // in file "wordsByFrequence.js"  	
		
		
} // end of get_language()

//--------------
"use strict";
/*  
LineByLine: A tool to practice language comprehension
Antonio Cigna 2023
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//----------------------------------------------
var currScript = document.currentScript.src; var bar1 = currScript.lastIndexOf("\\");var bar2 = currScript.lastIndexOf("/"); 
//console.log("LOADED file SCRIPT " + currScript.substring( 1+Math.max(bar1,bar2) )) ;	
//----------------------------------------------------------------------------------------
/*
let selected_voice_ix                    = 0 ;     // eg. 65 	 
let selected_voiceName                   = "";     // eg. Microsoft David - English (United States)"; 	
let	selected_voiceLangRegion             = "";     // eg. en-us	
let	selected_voiceLang2                  = "";     // eg. en
*/
//======================================	
function tts_2_fill_the_voices(wh) { 
	
	//console.log("\nxxxxxxxx\nxx tts_2_fill_the_voices (", wh, ")\nxxxxxxx");
	
	var numTotVoices = voices.length
	/**
	console.log("voices.length=" + numTotVoices); 
	console.log("voices[0]=" + voices[0].name); 
	console.log("voices["+(numTotVoices-1)+"]=" + voices[(numTotVoices-1)].name); 
	**/
	for(var ix=0; ix < numTotVoices; ix++) {
		if (selected_voiceName == voices[ix].name) {
			selected_voice_ix        = ix; 	
			selected_voiceLangRegion = voices[ix].lang	
			selected_voiceLang2      = selected_voiceLangRegion.substr(0,2);
			//selected_voiceName     = voices[ix].name
			break;
		}
		
	}
	var vox;
	listVox = [];
	
	//console.log("voices.length=" + numTotVoices, " selected_voice_ix=", selected_voice_ix); 
	
	// if (selected_voice_ix == 0) { return; } ???
	
	//firstly the chosen language-voice
	vox = voices[selected_voice_ix];
	 
	listVox.push( [vox.lang , vox] );     //1push voice
	if (selected_voiceLangRegion != vox.lang) {
		console.log("tts_2_fill_the_voices () " + "\n\tselected_voice_ix=" + selected_voice_ix + 
			"\n\tselected_voiceName = " + selected_voiceName +
			"\n\tselected_voiceLangRegion = " + selected_voiceLangRegion +
			"\n\tselected_voiceLang2 = " + selected_voiceLang2); 
		console.log("ERROR vox.lang (from voices[selected_voice_ix]) = " + vox.lang  + 
				" vs " + "selected_voiceLangRegion=" +selected_voiceLangRegion);  
		console.log(signalError)		
	}
	
	//------------------------------------------
	// secondly the same language-region 
	for(var v2=0; v2 < voices.length; v2++) {
		vox = voices[v2];
		if (v2 == selected_voice_ix) continue; 
		
		if (selected_voiceLangRegion != vox.lang ) continue;	
		
		listVox.push( [vox.lang , vox] );  	//2push voice	
	}
	//---------------------------------	
	// thirdly the same language
	for(var v2=0; v2 < voices.length; v2++) {
		vox = voices[v2];
		if (selected_voiceLangRegion == vox.lang ) continue;	
		if (selected_voiceLang2 != vox.lang.substr(0,2) ) continue;				
		listVox.push( [vox.lang , vox] );  //3push voice 
	}
	//---------------------------------	
	var str_option="";
	for(v3=0; v3 < listVox.length; v3++) {		
		var vv1, vv2; 
		[vv1,vv2] = listVox[v3]
		//console.log("tts_2_fill_the_voices () ",  "listVox[" +v3 + "] = " + vv1 + " " + vv2.name);
		str_option += "   <option>" + vv2.name + "</option> \n";
	}
	ele_select_first_language.innerHTML = str_option; 
	ele_select_first_language.selectedIndex = 0; 
	//----------------	
	var chosenIxVox=0;
	//-----------
	if (listVox.length == 0) {
		return; 
	}
	//console.log("listVox length=" + listVox.length); 
	voice_toUpdate_speech = listVox[0][1] ;	

	
	var voxLang;
	var pVox = ""; var xbr; 
	var vv3=0; var v3;
	var idhvox, idh2, eleH; 
	totNumMyLangVoices = listVox.length;
	
 	
} // end of fill_the_voices()

//--------------------------
function test_theVoice(lang2, myVoice) {
	if (lang2 == myVoice.lang.substr(0,2) ) return true;  
	
} // end of test_theVoice()

//====================================== "use strict";
/*  
LineByLine: A tool to practice language comprehension
Antonio Cigna 2023
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//----------------------------------------------
var currScript = document.currentScript.src;
var bar1 = currScript.lastIndexOf("\\");
var bar2 = currScript.lastIndexOf("/");
//console.log("LOADED file SCRIPT " + currScript.substring(1 + Math.max(bar1, bar2)));
//----------------------------------------------------------------------------------------

let last_objtxt_to_speak;

//=============================================

function tts_3_getUrlData3() {

    console.log("PLAYER  tts_3_getUrlData3() ");

    sw_is_no_videoaudio = true;


    //console.log("\nXXXXXXXXXXXXXXXXXXXXXXXXXX\nX getUrlData3() \nXXXXXXXXXXXXXXXXXXXXXXXXXX\n");

    /**
	document.getElementById("id_voice_language").innerHTML = decodeURI( urlParams.get("p_sel_voicelang" ) );   
	
	var selected_language_fromBuilder = document.getElementById("id_voice_language").innerHTML ; // eg.  en,6
	language_parameters = (selected_language_fromBuilder+",,,").split(",");  
    selected_lang_id = (language_parameters[0]+"   ").trim().substr(0,2); 
	selected_numVoices = parseInt("0" + language_parameters[1].trim() )  ; 	
	if (selected_numVoices > maxNumVoices)  selected_numVoices = maxNumVoices;
	***/


    sw_is_no_videoaudio = true;


    sw_is_no_videoaudio = true;

    cbc_LOCALSTOR_key = "ClipByClip_player_" + encodeURI(title1);


    tts_3_FAKE_onloaded_fun();

} // end of f3_tts_getUrlData3()

//---------------------------------------------
function tts_3_word_build_td_voices() {
    var selected_numVoices2 = 1; // selected_numVoices; // 1;
    var str1 = "";


    var fun1_m1 = '"onclick_tts_playSynthVoice_m1_word(this,';


    var fun2 = '"onclick_tts_word_playSynthVoice_row2(this1, §1§,'; //   continue -->   swLoop, isWord, is_m1,swNewVoice) {


    var oneVoice1_1 = '\n<td class="c_voice" style="';
    var oneVoice1_2 = 'background-color:lightblue;';
    var oneVoice1_31_m1 = '"  ><button class="buttonWhite" onclick=' + fun1_m1;
    var oneVoice1_32 = '"  ><button class="buttonWhite" onclick=' + fun2;

    var oneVoice1_4 = ')">';
    var oneVoice2 = `
		<span style="font-size:2em;height:1.4em;">${speakinghead_symb}</span></button></td>
				`; // do not replace backtick(` `) with quote(") or single_quotation_mark (')
    str1 = "";
    var n;
    var swB = false;


    var str_m2 = "";
    var str_m1 = "";
    for (n = 1; n <= selected_numVoices2; n++) {
        if (swB) {
            swB = false;
            str_m1 += oneVoice1_1.replace("c_voice", "c_voice_m1") + oneVoice1_31_m1 + n +
                oneVoice1_4 + n + oneVoice2.trim();
        } else {
            swB = true;
            str_m1 += oneVoice1_1.replace("c_voice", "c_voice_m1") + oneVoice1_2 + oneVoice1_31_m1 + n +
                oneVoice1_4 + n + oneVoice2.trim();
            //document.getElementById("hvox" + n).style.backgroundColor = "lightblue" ;
        }
        //document.getElementById("hvox" + n).style.fontSize = "0.5em"; 
        str_m2 += "<td></td>";
    }

    var str2 = "";
    for (n = 1; n <= selected_numVoices2; n++) {
        if (swB) {
            swB = false;
            str2 += oneVoice1_1 + oneVoice1_32 + n + oneVoice1_4 + n + oneVoice2.trim();
        } else {
            swB = true;
            str2 += oneVoice1_1 + oneVoice1_2 + oneVoice1_32 + n + oneVoice1_4 + n + oneVoice2.trim();
            //document.getElementById("hvox" + n).style.backgroundColor = "lightblue" ;
        }
        //document.getElementById("hvox" + n).style.fontSize = "0.5em"; 
    }

    //document.getElementById("col_voce1").span = selected_numVoices; 
    /**
    for(n=selected_numVoices+1; n <=6; n++) {  // all of them are defined as display table-column 
    	document.getElementById("hvox" + n).style.display="none"; 
    }
    **/

    //str1 =  voice_col_tr_td;


    return [str_m2, str_m1, str2];

} // end of word_build_td_voices()


// ===================================================================
function tts_3_spezzaRiga2(orig_riga, tts_riga) {



    var endix2 = -1;



    var orig_riga2 = orig_riga.replaceAll(". ", ". §").replaceAll("! ", "! §").replaceAll("? ", "? §").replaceAll("; ", "; §").
    replaceAll(": ", ": §").replaceAll(", ", ", §").replaceAll(" ", " §");

    if (tts_riga == "") tts_riga = orig_riga;

    var tts_riga2 = tts_riga.replaceAll(". ", ". §").replaceAll("! ", "! §").replaceAll("? ", "? §").replaceAll("; ", "; §").
    replaceAll(": ", ": §").replaceAll(", ", ", §").replaceAll(" ", " §");

    //console.log("\tspezzaRiga2( orig_riga=" + orig_riga + "\n\t\ttts_riga=" + tts_riga); 

    var listaParole = [];
    var listaParo_tts = [];
    var list1 = orig_riga2.split("§");
    var list2 = tts_riga2.split("§");

    //console.log( list1.length + " " + list2.length); 


    if (list2.length != list1.length) list2 = orig_riga2.split("§");

    for (var f = 0; f < list1.length; f++) {
        if (list1[f].trim() != "") {
            listaParole.push(list1[f]);
            listaParo_tts.push(list2[f]);
        }
    }


    var parola1, paro_tts;

    var frase_showTxt = '<table style="border:0px solid red;width:100%;margin-top:1em;"> \n';
  
    endix2 = listaParole.length;

    for (let z3 = 0; z3 < listaParole.length; z3++) {
        parola1 = listaParole[z3];
        paro_tts = listaParo_tts[z3];
        let rowclip = word_tr_allclip.replaceAll("§1§", z3).replaceAll("§4txt§", parola1).replaceAll("§ttsWtxt§", paro_tts);
        frase_showTxt += rowclip + "\n";
    } // end of for z3

    return frase_showTxt += '</table>\n';

} //  end of  spezzaRiga2()

//------------------------------------------

function tts_3_call_boldCell_ix(this1, ixVoice, wh) {
    var td1 = this1.parentElement;
    var tr1 = td1.parentElement;
    tts_3_boldCell(tr1, this1, ixVoice, " call_boldCell_ix() " + wh);
}
//-------------------------------------

function tts_3_show_speakingVoiceFromVoiceLangName(voice_lang, voice_name) {

    var msg1 = document.getElementById("m120").innerHTML; //  spoken in
    var lang0 = get_languageName(voice_lang).split("-");

    var msg = "<b>" + lang0[0] + "</b><br><small>" +
        document.getElementById("m120").innerHTML + //  spoken in
        "</small><br><b>" + lang0[1] + "</b><br>";

    ele_voxLangDisplay.innerHTML = msg + "<br><small>" + voice_name + "</small>";

} // 

//-------------
function tts_3_show_speakingVoice(objtxt_to_speak) {
    tts_3_show_speakingVoiceFromVoiceLangName(objtxt_to_speak.voice.lang, objtxt_to_speak.voice.name);
} // 

//------------------------------------
function tts_3_set_speech_Parms(objtxt_to_speak) {

	//console.log("tts_3_set_speech_Parms() lastNumVoice=" + lastNumVoice + " lang=" + listVox[lastNumVoice][0]);
	var myVoice = listVox[lastNumVoice][1];
	var voice_lang2 = myVoice.lang.substr(0,2); 
	
	//voice_lang2 = "mioerrore"; // forzo errore per testare
	
	if (voice_lang2 != selected_voiceLang2) {		
		console.log("error: voice to set (" + myVoice.name + " is not of the selected language (" +  voice_lang2 + " vs " + selected_voiceLang2); 
		console.log("lastNumVoice=" + lastNumVoice  + " listVox.length=" + listVox.length); 
		//console.log(signalerror); 
		tts_2_fill_the_voices("1tts_3_set_speech_Parms"); 
		if (lastNumVoice >= listVox.length) lastNumVoice = 0;  
		myVoice = listVox[lastNumVoice][1];
	}
	
    last_objtxt_to_speak = objtxt_to_speak;
	
    objtxt_to_speak.voice = myVoice;

    tts_3_show_speakingVoice(objtxt_to_speak);

    objtxt_to_speak.rate = speech_rate;
    objtxt_to_speak.pitch = speech_pitch;


} // tts_3_set_speech_Parms()

//---------------------------------------
function tts_3_word_fun_invisible_prev_fromto(interX, isWord, is_m1) {
    // eliminate bold of the  previous group of lines, unless this is a line in them  

    var id_pre_tr_beg_space, id_pre_tr_head, id_post_tr_end_space;
    if (isWord) {
        [begix, endix] = word_fromIxToIxLimit; // previously set 
        if ((interX >= begix) && (interX <= endix)) return;
        id_pre_tr_beg_space = "widtr_" + begix + "_m2";
        id_pre_tr_head = "widtr_" + begix + "_m1";
        id_post_tr_end_space = "widtr_" + (endix + 1) + "_m2";
        if (word_fromIxToIxButtonElement[0]) {
            word_fromIxToIxButtonElement[0].style.backgroundColor = null;
            if (document.getElementById(id_pre_tr_beg_space)) {
                document.getElementById(id_pre_tr_beg_space).style.display = "none";
            }
            if (document.getElementById(id_pre_tr_head)) {
                document.getElementById(id_pre_tr_head).style.display = "none";
            }
        }
        if (word_fromIxToIxButtonElement[1]) {
            word_fromIxToIxButtonElement[1].style.backgroundColor = null;
            if (document.getElementById(id_post_tr_end_space)) {
                document.getElementById(id_post_tr_end_space).style.display = "none";
            }
        }
    } else {
        [begix, endix] = fromIxToIxLimit;
        if ((interX >= begix) && (interX <= endix)) return;
        id_pre_tr_beg_space = "idtr_" + begix + "_m2";
        id_pre_tr_head = "idtr_" + begix + "_m1";
        id_post_tr_end_space = "idtr_" + (endix + 1) + "_m2";
        if (fromIxToIxButtonElement[0]) {
            fromIxToIxButtonElement[0].style.backgroundColor = null;
            if (document.getElementById(id_pre_tr_beg_space)) {
                document.getElementById(id_pre_tr_beg_space).style.display = "none";
            }
            if (document.getElementById(id_pre_tr_head)) {
                document.getElementById(id_pre_tr_head).style.display = "none";
            }
        }
        if (fromIxToIxButtonElement[1]) {
            fromIxToIxButtonElement[1].style.backgroundColor = null;
            if (document.getElementById(id_post_tr_end_space)) {
                document.getElementById(id_post_tr_end_space).style.display = "none";
            }
        }
    }
} // end of word_fun_invisible_prev_fromto()




//------------------------------------------

function tts_3_word_fun_copy_openClose_to_tr_m1(z3, isWord, is_m1) {
    var i_eleSubO, o_eleSubO;
    if (isWord) {
        i_eleSubO = document.getElementById("widb_" + z3);
        o_eleSubO = document.getElementById("widb_" + z3 + "_m");
    } else {
        i_eleSubO = document.getElementById("idb_" + z3);
        o_eleSubO = document.getElementById("idb_" + z3 + "_m");
    }
    try {
        o_eleSubO.children[0].style.display = i_eleSubO.children[0].style.display; // ${openbook_symb}
        o_eleSubO.children[1].style.display = i_eleSubO.children[1].style.display; // ${closedbook_symb}
    } catch (e1) {
        console.log("error in 'word_fun_copy_openClose_to_tr_m1(z3=" + z3 + ")'");
        console.log(e1);
    }
} // end of word_fun_copy_openClose_to_tr_m1() 
//---------------------------------------


//-----------------------------------------------
function tts_3_word_fun_copyHeaderSelected(begix, endix) {

    let id1;
    let inBeg, inEnd;

    id1 = "widb_" + begix + "_m";
    var ele_idb = document.getElementById(id1);
    if (ele_idb == false) {
        return;
    }

    inBeg = begix;
    inEnd = endix;
    if (begix > endix) {
        inBeg = endix;
        inEnd = begix;
    }
    var style0, style1;

    tts_3_fun_oneRowZZ1();

    for (var g = inBeg; g <= inEnd; g++) {
        id1 = "widb_" + g;
        ele_idb = document.getElementById(id1);
        tts_3_word_fun_oneRow11(ele_idb);
    }

    //--------------
    function tts_3_fun_oneRowZZ1() {
        if (ele_idb == false) {
            return;
        }
        if (ele_idb.children[0].style.display == "none") { // no openbook   
            style0 = "none"; //  hide opened book image  
            style1 = "block"; //  show closed book image 
        } else {
            style0 = "block"; // show opened book image  
            style1 = "none"; // hide closed book image 	
        }
    }
    //--------------
    function tts_3_word_fun_oneRow11(ele_idb) {
        if (ele_idb == false) {
            return;
        }
        ele_idb.children[0].style.display = style0; // show/hide opened book image  
        ele_idb.children[1].style.display = style1; // show/hide closed book image 
        let subid_idc = ele_idb.id.replace("idb", "idc");
        if (subid_idc.substring(subid_idc.length - 2) == "_m") {
            return;
        }
        let ele_idc = document.getElementById(subid_idc);
        if (ele_idc == false) return;
        if (style0 == "block") {
            tts_3_word_fun_makeTextVisible(ele_idc);
        } else {
            tts_3_word_fun_makeTextInvisible(ele_idc);
        }
    } // end of tts_3_word_fun_oneRow11) //  
    //--------------

} // end of word_fun_copyHeaderSelected()

//-----------------------------
//--------------
function tts_3_word_play_or_cancel(this1) {

    this1.style.backgroundColor = "red";
    ele_last_play = this1;
    return 0;
} // end of  word_play_or_cancel

//----------------------------

function tts_3_word_fun_makeTextInvisible(element) {
    if (element == null) {
        return;
    }
    //element.style.visibility = "hidden"; 
    element.style.display = "none";	
	
	let id1 = element.id
	/**
	if (id1.indexOf("widc_") < 0) return;
	let id2 = id1.replace("widc_", "widcT_");
	document.getElementById(id2).style.display = "none";	
	**/
	
}
//----------------------------

function tts_3_word_fun_makeTextVisible(element) {
    if (element == null) {
        return;
    }
    element.style.display = "block";
	/**
	let id1 = element.id
	if (id1.indexOf("widc_") < 0) return;	
	let id2 = id1.replace("widc_", "widcT_");
	document.getElementById(id2).style.display = "block";	
	**/
}

//------------------------
function tts_3_word_removeLastBold(isWord) {

    var id_pref = "idc_";
    if (isWord) id_pref = "widc_";

    if (word_last_BoldRow) {
        word_last_BoldRow.classList.remove("boldLine");
        word_last_BoldRow.style.backgroundColor = null;
        word_last_BoldRow.parentElement.style.border = null;
        var last_ele1_tr = word_last_BoldRow.parentElement.parentElement;
        last_ele1_tr.style.backgroundColor = "lightgrey";
    }
    if (sw_tts2) tts_remove_last_bold(id_pref, isWord);
}
//-------------------------------------------


var last_id_Bold = ["",""] 


//-----------------------------------

function tts_3_word_fun_oneClipRow_showHide_ORIG_if_book_opened(ele1, ele_to_test, z3, isWord) {

    tts_3_word_removeLastBold(isWord);
    if (ele1 == null) {
        return;
    }
    if (isWord) word_last_BoldRow = ele1;
    else word_last_BoldRow = ele1;

    if (ele_to_test.children[0].style.display == "block") { // openbook ==> show 
        //ele1.style.visibility = "visible"; 
        ele1.style.display = "block";
        ele1.classList.add("boldLine");
        ele1.style.backgroundColor = "yellow";
        ele1.parentElement.style.border = null;
        //feb if (sw_is_no_videoaudio == false) ele1_tr.style.backgroundColor = "yellow";			
        //ele1_tr.style.backgroundColor = "yellow"; //feb 		
    } else { // closebook  ==> hide 
        //ele1.style.visibility = "hidden"; 	
        ele1.style.display = "none";
        ele1.classList.remove("boldLine");
        ele1.style.backgroundColor = null;
        ele1.parentElement.style.border = "1px solid red";
        //feb if (sw_is_no_videoaudio == false) ele1_tr.style.backgroundColor = "yellow";		
        //ele1_tr.style.backgroundColor = "yellow";	//feb	  
    }

} // end of fun_oneClipRow_showHide_ORIG_if_book_opened()


//-----------------------------------------------
function tts_3_removeLastEle(id1) {
    //console.log("\tremoveLastEle(id1=" + id1);
    if (document.getElementById(id1)) document.getElementById(id1).remove();
}




//js3_____________________________
//----------------------------------------------------------
function tts_3_break_priority(txt1, maxLen, sw_let_old_newline) {

    if (sw_let_old_newline == false) {
        txt1 = txt1.replaceAll("\n", " "); // eliminate existing newline  
    }

    // break text into lines always for: end of sentence (.), question mark(?), exclamation(!), semicolon(;)  

    var row1 = txt1.replaceAll("<br>", "\n").
    replaceAll(". ", ".\n").
    replaceAll("? ", "?\n").
    replaceAll("! ", "!\n").
    replaceAll("; ", ";\n");

    return row1.split("\n");
} // end of tts_3_break_priority

//-----------------------------------------

function tts_3_break_text(txt1, maxLen, sw_let_old_newline) {

    // break for: new line, end of sentence (.), exclamantion(!)  and question mark(?), and semicolon(";")     
    // break for too long line ( last comma or last blank before reaching the maximum length   

    var rtxt2 = tts_3_break_priority(txt1, maxLen, sw_let_old_newline);

    //console.log("break_text(txt1=" + txt1 + " \n\trtxt2=" + rtxt2);   

    //-------------------------------	
    function tts_3_tooLongLine(oneLine) {
        // break in strings with their length not > maxlen  ( firstly try to find colon(:), then comma(,) and lastily space(" ")  
        var txt3 = oneLine.trim();
        var newLine2 = "";
        var len1;
        var txt3a;
        var u, u1, u2;
        //-------------
        for (var h = 0; h < txt3.length; h++) {
            len1 = txt3.length;
            if (len1 < 1) {
                break;
            }
            if (len1 <= maxLen) {
                newLine2 += txt3 + "\n";
                break;
            }

            txt3a = txt3.substring(0, maxLen);
            u1 = txt3a.lastIndexOf(": ");
            u2 = txt3a.lastIndexOf(", ");
            u = Math.max(u1, u2);
            if (u >= 0) {
                u++;
            } else {
                u = txt3a.lastIndexOf(" ");
            }
            if (u < 0) {
                u = txt3.indexOf(" "); // find next forward 
                if (u < 0) {
                    u = txt3.length; // take all string 
                }
            }
            newLine2 += txt3.substring(0, u) + "\n";

            txt3 = txt3.substring(u).trim();
        }
        return newLine2;

    } // end of tooLongLine();

    //-------------------------------------------

    var newLine = "";
    for (var g = 0; g < rtxt2.length; g++) {
        newLine += tts_3_tooLongLine(rtxt2[g]);
    }


    return newLine;

} // end of break_text() 

//-----------------------------------------
function tts_3_test_break_text(txt1) {

    console.log("old=" + "\n" + txt1 + "\n-------------------\n");

    var newLine = break_text(txt1, TXT_SPEECH_LENGTH_LIMIT, sw_let_old_newline);
    var lines = newLine.split("\n");
    for (var v1 = 0; v1 < lines.length; v1++) {
        console.log(v1 + "  " + lines[v1]);
    }
}
//-----------------------------

//js4________________________

//----------------

function tts_3_speak_a_line(objtxt_to_speak, wh) {
    if (sw_cancel) {
        TTS_LOOP_swLoop = false;
        sw_cancel = false;
        return;
    }
    //speech_rate  = last_rate;

    tts_3_set_speech_Parms(objtxt_to_speak);
    
	
    synth.speak(objtxt_to_speak);

} // end of tts_3_speak_a_line()
//-------------------------------------------


//---------------------
function tts_3_speech_end_fun() {

    x1_line++;

    if ((x1_line) >= textLines.length) {
        //console.log("tts_3_speech_end_fun()1");
        tts_3_end_speech(); // defined in the caller   **   Feb 11, 2023	
        return;
    }
    var objtxt_to_speak;
    //rigout += textLines[x1_line] + "<br>";



    if ((x1_line + 1) >= textLines.length) {
        //console.log("tts_3_speech_end_fun()2");
        tts_3_end_speech(); // defined in the caller   **   Feb 11, 2023	
        return;
    }
  	
    objtxt_to_speak = utteranceList[x1_line + 1];
    //speech = objtxt_to_speak; //2  

   
    objtxt_to_speak.onend = tts_3_speech_end_fun;

    //objtxt_to_speak.onend = (event) =>  {tts_3_speech_end_fun(event) } ;
 
    tts_3_speak_a_line(objtxt_to_speak, 2);

} // end of tts_3_speech_end_fun

//----------------------------------------------------------


//-------------------------------------------
var pLastBold_ix1 = -1; // phrase
var pLastBold_ix2 = -1;
var wLastBold_ix1 = -1; // word
var wLastBold_ix2 = -1;
//---------------------

function tts_3_remove_last_bold(id_pref, isWord) {

    var lastBold_ix1, lastBold_ix2;
    if (isWord) {
        lastBold_ix1 = wLastBold_ix1;
        lastBold_ix2 = wLastBold_ix2;
    } else {
        lastBold_ix1 = pLastBold_ix1;
        lastBold_ix2 = pLastBold_ix2;
    }
    if (lastBold_ix2 < 0) return;

    for (var v = lastBold_ix1; v <= lastBold_ix2; v++) {
        var ele1 = document.getElementById(id_pref + v);
        if (ele1 == false) continue;
        var ele1_tr = ele1.parentElement.parentElement;
        ele1.classList.remove("boldLine");
        ele1.style.backgroundColor = null;
        ele1.parentElement.style.border = null; // "1px solid red"; 
        ele1_tr.style.backgroundColor = "lightgrey"; // "yellow";	//feb	  
    }
    lastBold_ix2 = -1;
    if (isWord) {
        wLastBold_ix1 = lastBold_ix1;
        wLastBold_ix2 = lastBold_ix2;
    } else {
        pLastBold_ix1 = lastBold_ix1;
        pLastBold_ix2 = lastBold_ix2;
    }
} // end of tts_remove_last_bold()


//----------------------------------
function tts_3_showHide_if_book_opened(id_pref, id_pref_idb, z3) {
    var isWord = (id_pref.substr(0, 1) == "w");
    tts_3_word_removeLastBold(isWord);

    //console.log("1tts showHide " + z3 );

    var ele1 = document.getElementById(id_pref + z3);
    var ele_to_test = document.getElementById(id_pref_idb + z3);

    var ele1_tr = ele1.parentElement.parentElement;



    if (ele1 == null) {
        return;
    }
    if (last_blue_cell) {
        last_blue_cell.style.border = null;
    }
    if (ele_to_test.children[0].style.display == "block") { // openbook ==> show 
        //ele1.style.visibility = "visible"; 
        ele1.style.display = "block";
        ele1.classList.add("boldLine");
        ele1.style.backgroundColor = "yellow";
        ele1.parentElement.style.border = null;
        //feb if (sw_is_no_videoaudio == false) ele1_tr.style.backgroundColor = "yellow";			
        ele1_tr.style.backgroundColor = "yellow"; //feb 	
    } else { // closebook  ==> hide 
        //ele1.style.visibility = "hidden"; 	
        le1.style.display = "none";
        ele1.classList.remove("boldLine");
        ele1.style.backgroundColor = null;
        ele1.parentElement.style.border = "1px solid blue"; //red
        last_blue_cell = ele1.parentElement;
        //feb if (sw_is_no_videoaudio == false) ele1_tr.style.backgroundColor = "yellow";		
        //ele1_tr.style.backgroundColor = "yellow";	//feb	  
    }
    if (isWord) {
        wLastBold_ix1 = z3;
        wLastBold_ix2 = z3;
    } else {
        pLastBold_ix1 = z3;
        pLastBold_ix2 = z3;
    }

} // end of tts_showHide_ORIG_if_book_opened() 

//--------------------------------------------

function tts_3_end_speech_calculation() {

    // called by tts_3_end_speech() in cbc_player_script   and cbc_player_word_script

    var endTime = new Date();
    var timeDiff = endTime - startTime; //actual elapsed time in ms 

    if (txt_length < 1) {
        return;
    }


    var normal_time = timeDiff * speech_rate;

    //-----------
    tot_norm_time += normal_time;
    tot_txt_len += txt_length;
    tot_norm_mill_char = tot_norm_time / tot_txt_len;
    tot_norm_str_leng_limit = ELAPSED_TIME_SPEECH_LIMIT / tot_norm_mill_char;
    TXT_SPEECH_LENGTH_LIMIT = parseInt(tot_norm_str_leng_limit * speech_rate);

    tts_3_loopManager();

} // end of tts_3_end_speech_calculation() 
//-----------------------------------------

function tts_3_end_speech() {

    if (ele_last_play) {
        if (TTS_LOOP_swLoop == false) {
            ele_last_play.style.backgroundColor = null;
        }
    }
    tts_3_end_speech_calculation();
} // end of tts_3_end_speech()

//-------------------------------------------

function tts_3_loopManager() {

    if (sw_cancel) {
        TTS_LOOP_swLoop = false;
        sw_cancel = false;
    }
    if (TTS_LOOP_swLoop == false) return;

    lastNumVoice++;
    if (lastNumVoice >= totNumMyLangVoices) lastNumVoice = 0; // to change voice on each cycle

    var id_pref = "idc_";

    // wait 1 second and then start again  
    setTimeout(function() {
            if (TTS_LOOP_begix == TTS_LOOP_endix) {
                onclick_tts_text_to_speech_ix(id_pref, TTS_LOOP_begix, TTS_LOOP_swLoop, TTS_LOOP_elem);
            } else {
                onclick_tts_text_to_speech_from_to(id_pref, TTS_LOOP_begix, TTS_LOOP_endix, TTS_LOOP_swLoop, "1 loop_manager");
            }
        },
        1000);
} // end of tts_3_loopManager()


//--------------------------
function tts_3_boldCell(tr1, this1, ixVoice, wh) {
    if (lastBoldCell) {
        lastBoldCell.style.backgroundColor = null;
    }
    var selected_numVoices2 = 1;

    var td1 = this1.parentElement;
    var ixTab_TD = td1.cellIndex;

    var ixDiff = ixTab_TD - ixVoice;
    var ixTD_0 = 0;

    var num_tr_cells = tr1.cells.length;

    for (var n = 0; n < selected_numVoices2; n++) {
        ixTD_0 = n + ixDiff;
        if (n == ixVoice) {
            lastBoldCell = tr1.cells[ixTD_0];
            lastBoldCell.style.backgroundColor = "green";
        } else {
            tr1.cells[ixTD_0].style.backgroundColor = null;
        }
    }

} // end of tts_3_boldCell()

//--------

//---------------------------

function tts_3_breakTextToPause(txt3, pre_idtr1) {
	/**
	break down the text in characters   
	works even with multybyte chars:  eg.  日本語  is broken down into  日, 本, 語, 
	**/
    var g, ww2, txt4 = "";
    if (pre_idtr1 == "") {
        var ww1 = (txt3 + " ").replaceAll("–", " ").replaceAll("-", " ").
        replaceAll(", ", " ").replaceAll(" .", " ").replaceAll(". ", " ").replaceAll("...", " ").
        replaceAll("? ", " ").replaceAll("! ", " ").split(" ");

        for (g = 0; g < ww1.length; g++) {
            ww2 = ww1[g].trim();
            if (ww2 != "") txt4 += ww2 + ". ";
        }
    } else {
        var txt5 = (txt3 + " ").replaceAll("–", " ").replaceAll("-", " ").
        replaceAll(", ", " ").replaceAll(" .", " ").replaceAll(". ", " ").replaceAll("...", " ").
        replaceAll("? ", " ").replaceAll("! ", " ").trim();

        for (g = 0; g < txt5.length; g++) {
            ww2 = txt5.charAt(g);
            if (ww2 != "") txt4 += ww2 + ", ";
        }
		//console.log("tts_3_breakTextToPause() " +  txt3 + "\n" + txt4.replaceAll(",", "\n\t"));
    }
    return txt4;

} // end of breakToPause()	


//---------------------------

function tts_3_removeBold_and_Font(txt0) {
    // I'm sorry<font color="#E5E5E5"> Val I'm frightfully tired but</font>
    //I&apos;m sorry&lt;font color=&quot;#E5E5E5&quot;&gt; Val I&apos;m frightfully tired but&lt;/font&gt;<br>
    let txt1 = txt0.trim().toLowerCase();
    txt1 = txt1.replaceAll("&lt;", "<").replaceAll("&gt;", ">").replaceAll("&apos;", "'").replaceAll("&quot;", '"');

    let j1 = -1;
    for (let x1 = 0; x1 < txt1.length; x1++) {
        j1 = txt1.indexOf("<font ");
        if (j1 < 0) {
            break;
        }
        if (txt1.indexOf("</font>", j1) < 0) {
            return txt1;
        }
        let j2 = txt1.indexOf(">", j1);
        if (j2 < 0) {
            break;
        }
        txt1 = txt1.substring(0, j1) + txt1.substring(j2 + 1);
        txt1 = txt1.replace("</font>", "");
    }
    txt1 = txt1.replace("< ", "<").replace(" >", ">").
    replaceAll("<em>", "").replaceAll("</em>", "").
    replaceAll("<strong>", "").replaceAll("</strong>", "").
    replaceAll("<b>", "").replaceAll("</b>", "").replaceAll("<i>", "").replaceAll("</i>", "").
    replaceAll("<B>", "").replaceAll("</B>", "").replaceAll("<I>", "").replaceAll("</I>", "");
    return txt1;

} // end of removeBold_and_Font()

//--------------------

function tts_3_play_or_cancel(this1) {
    if (synth.speaking) {
        if (ele_last_play) ele_last_play.style.backgroundColor = null;

        onclick_tts_speech_cancel();
        tts_3_end_speech();
        if (this1 == ele_last_play) { // click on the same line which is running ==> it means ==> I just want to stop it   
            return -1;
        }
        // click not on the sameline running ==> I wanted to stop the last line e start a new one 

    }

    this1.style.backgroundColor = "red";
    ele_last_play = this1;
    return 0;
}

//-------------------------------------

//-----------------------------------------------------------------
function tts_3_FAKE_onloaded_fun() {
    console.log("\nXXXXXXXXXXXXXXXXXXXXXXXXXX\nX tts_3_FAKE_onloaded_fun() \nXXXXXXXXXXXXXXXXXXXXXXXXXX\n");


    tts_3_fun_player_beginning();



} // end of tts_3_FAKE_onloaded_fun() 

//----------------------------------------------

//---------------------------------------------

function tts_3_fun_player_beginning() {

    console.log("\nXXXXXXXXXXXXXXXXXXXXXXXXXX\nX tts_3_fun_player_beginning() \nXXXXXXXXXXXXXXXXXXXXXXXXXX\n");


    ele_clip_subtext = document.getElementById("id_tabSub");

    MAX_ixClip = line_list_o_number_of_elements - 1;
    MIN_ixClip = 1;


    ele_dragSubT = document.getElementById("id_dragSub");
    ele_dragSubT_anchor = document.getElementById("id_dragSub_anchor");



    wScreen = screen.availWidth;
    hScreen = screen.availHeight;

    subtitles_beg_delta_time = 0;
    //src1 = document.getElementById("myVideo").src;

    lastClipTimeBegin = 0;
    lastClipTimeEnd = 0;

    MAX999 = 999999;


    sw_sub_orig = sw_inp_sub_orig_builder;
    sw_sub_tran = sw_inp_sub_tran_builder;
    sw_sub_onfile = (sw_sub_orig || sw_sub_tran);

    sw_no_subtitle = ((sw_sub_onfile == false)); // no subtitles ( neither inside the video, neither in any file apart

    if (sw_no_subtitle) {
        eleTabSub.style.display = "none";
    }
    //-------------------------------------------------------       
    // let lev2 
    path1 = window.location.pathname;
    f1 = path1.lastIndexOf("/");
    f2 = path1.lastIndexOf("\\");
    f3 = -1;
    barra = "/";
    if (f1 > f2) {
        f3 = f1;
        barra = "/";
    } else {
        f3 = f2;
        barra = "\\";
    }


    lastClipTimeBegin = 0;
    lastClipTimeEnd = 0;

    tts_5_get_tran_subtitle_text(); // must stay before 'get_orig_subtitle_text()'	
    tts_5_get_orig_subtitle_text();


    line_list_o_number_of_elements = line_list_o_from1.length;
    line_list_t_number_of_elements = line_list_t_from1.length;

    if ((sw_sub_orig == false)) {
        document.getElementById("id_td_suborig2").style.display = "none";
    }

    if (sw_sub_tran == false) {
        document.getElementById("id_td_subtra2").style.display = "none";
    }



    LIMIT_MIN_TIME_CLIP = 0.100; // if time too near ( difference <  LIMIT_MIN_TIME_CLIP) to toTimeClip than use next clip number 

    tts_5_fun_update_html_with_last_session_values();

    //fun_set_next_div_via_cliptype(); 


    //-------------------------------

    tts_5_fun_setMinMaxIxClip();


    if (sw_is_no_videoaudio == false) {
        tts_5_fun_replace_video_src();
    }

	/**
    console.log("\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" +
        "\nsw_no_subtitle   = " + sw_no_subtitle +
        "\nsw_sub_onfile    = " + sw_sub_onfile +
        "\nsw_sub_orig      = " + sw_sub_orig + "   num." + line_list_orig_text.length +
        "\nsw_sub_tran      = " + sw_sub_tran + "   num." + line_list_tran_text.length +
        "\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx");
	**/

    if (sw_sub_onfile == false) {
        document.getElementById("id_labshowclip").style.display = "none";
        document.getElementById("id_inpshowclip").style.display = "none";

        document.getElementById("id_playNext2_line_row").style.display = "none";
        document.getElementById("id_replay2_line_row").style.display = "none";
    }



    tts_5_former_onclick_tts_ClipSub_All(); //  ...ClipSub_All(document.getElementById("id_inpshowclipA") ); 
    /**
    //subtitle_row_length2();
    onclick_tts_scroll_right( document.getElementById("id_bRigthLeft") ) ;
    onclick_tts_scroll_right( document.getElementById("id_bRigthLeft") ) ;
    onclick_tts_scroll_right( document.getElementById("id_bRigthLeft") ) ;
    **/

} // end of player script beginning 

//------------------------------------------------------------------------"use strict";
/*  
LineByLine: A tool to practice language comprehension
Antonio Cigna 2023
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//----------------------------------------------
var currScript = document.currentScript.src; var bar1 = currScript.lastIndexOf("\\");var bar2 = currScript.lastIndexOf("/"); 
//console.log("LOADED file SCRIPT " + currScript.substring( 1+Math.max(bar1,bar2) )) ;	
//----------------------------------------------------------------------------------------

function tts_4_split_into_rows(inpTxt, maxRowLen) {

    var newText = inpTxt.trim().replaceAll("  ", " ").replaceAll("  ", " ").replaceAll("  ", " ").replaceAll("  ", " ").replaceAll("  ", " ").
	replaceAll("?","?\n").replaceAll("!","!\n").
    replaceAll(". ", ". \n").replaceAll(" \n", "\n").replaceAll("\n\n", "\n");
    var rows = newText.split("\n");

	var txt2 = "";	
	var txtOut=""
    for (var z1 = 0; z1 < rows.length; z1++) {
	
		txt2 = tts_4_words_from_one_line(rows[z1], "1 "+ z1).trim();
		
		if ((txt2 == "") || (txt2 == "\n"))  continue;
		txtOut += "\n" + txt2;
    }	
	
	return txtOut.substring(1); 
	
	//----------------------------
	
    function tts_4_words_from_one_line(rowOne, wh) {
		
		var textOut = "";       
		var woFS = rowOne.split(". ");
		var wo, wo00, wo1, wo1X, zx, z2, j1;
		
		wo00 = rowOne.trim();
		
		for (zx = 0; zx < woFS.length; zx++) {		
			wo00 = woFS[zx].trim();
			if (wo00 == "") break; 
						
			for(var b=0; b < wo00.length; b++) { 
				//if (b > 50) break;
				wo1 = wo00.trim(); 
				
				if (wo1 == "") break; 
				if (wo1.length <= maxRowLen) {
					textOut+="\n" + wo1;					
					break;					
				} 
				wo1X = wo1.substring(0,maxRowLen); 

				j1 = wo1X.lastIndexOf(";");
				if (j1 < 0) {j1 = wo1X.lastIndexOf(","); }
				if (j1 >=0) j1++;  				
				if (j1 < 0) {j1 = wo1X.lastIndexOf(" "); }
				if (j1 < 0) { // no comma, neither space backward, try forward 
					j1 = (wo1+" ").indexOf(" "); 
				}
				
				wo00 = wo1.substring(j1);
				textOut+="\n" + wo1.substring(0,j1); 	
			}
		}		
		return textOut;
    } // end of  words_from_one_line()   

} // end of  split_into_rows()

//=================================================="use strict";
/*  
LineByLine: A tool to practice language comprehension
Antonio Cigna 2023
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//----------------------------------------------
var currScript = document.currentScript.src;
var bar1 = currScript.lastIndexOf("\\");
var bar2 = currScript.lastIndexOf("/");
//console.log("LOADED file SCRIPT " + currScript.substring(1 + Math.max(bar1, bar2)));
//----------------------------------------------------------------------------------------

function tts_5_get_orig_subtitle_text() {

    let i, row, from1, to1, max_to1, rowcol;

    line_list_o_from00 = [];
    line_list_o_to00 = []; //
    line_list_o_maxto00 = [];
    line_list_o_from1 = []; //
    line_list_o_to1 = []; //
    line_list_o_maxto1 = []; //
    line_list_orig_text = [];
	line_list_orig_nFileR = []; 
    line_list_orig_tts = [];
    line_list_o_tran_ixmin = [];
    line_list_o_tran_ixmax = [];

    line_list_o_from00.push(-1);
    line_list_o_to00.push(-1);
    line_list_o_maxto00.push(-1);


    line_list_o_from1.push(-1);
    line_list_o_to1.push(-1);
    line_list_o_maxto1.push(-1);

    line_list_orig_text.push("");
	line_list_orig_nFileR.push(""); 
    line_list_orig_tts.push("");
    line_list_o_tran_ixmin.push(0);
    line_list_o_tran_ixmax.push(0);


    if (sw_sub_orig == false) {
        return;
    }

    max_to1 = 0;
    to1 = 0;
    let ixmin = 0;
    let ixmax = 0;
    let tran_len = line_list_t_from1.length;
    //-------------------------
    let pre_from = 0;
    let pre_to1 = 0;

    let lastix = inp_row_orig.length - 1;

    for (i = inp_row_orig.length - 1; i > (inp_row_orig.length - 10); i--) {
        if (inp_row_orig[i].trim() == "") {
            continue;
        }
        row = inp_row_orig[i].trim();
        lastix = i;
        break;
    }
    from1 = 0;
    to1 = -1;
    let x1 = 0;
    let x2 = 0;

    number_of_subtitle_endsentence = 0;
    number_of_subtitle_time_overlap = 0;

    var rowtext, rowTTS;
	var cols = [];
	var nFileRow;
	//var nFileRow, nFile, nRows;
	var row123;
	
    //---------------------
    for (i = 0; i <= lastix; i++) {
        row = inp_row_orig[i].trim();
        if (row == "") {
            continue;
        }
		//if (row.substring(0,1) == ";") {row = row.substring(1)}
        pre_from = from1;
        pre_to1 = to1;

		// wbf:      ;;1;59;; BC DEF    ;;nFile;n.row;; row 
        cols     = (row + ";;;;;;").split(";;");
		nFileRow = cols[1].replace(";","_"); 
		//nFileRow = (cols[1]+";;;").split(";");
		//nFile    = nFileRow[0]; 
		//nRows    = nFileRow[1];
		row123   = cols[2]	  ; 
		//console.log( "??ANTO  row=" + row + "\n\tcols=" , cols, "\n\tnFileRow=" , nFileRow , "\n\tnFile="+nFile + " nRows=" + nRows + "  row123=" + row123 + "<==");  	
		rowcol   = [i, (i + 1), row123.trim()];
		
        //rowcol = [i, (i + 1), row];

        from1 = parseFloat(rowcol[0]);

        if (from1 < pre_from) {
            //     original language subtitle line ignored because not in sequence of time:  pre=" + pre_from + " now=" + row 
            continue;
        }

        if (from1 == pre_to1) { // to avoid that in the same moment there be 2 subtitles lines actives when actualy one is the following of the other  (upd 2022_03_16)

            from1 += 0.0004;
        }
        to1 = parseFloat(rowcol[1]);

        if (from1 < pre_to1) {
            number_of_subtitle_time_overlap++; // there is a time overlap ( the speaker begin to speak before the previous one have finished)
        }

        max_to1 = Math.max(max_to1, to1);


        line_list_o_from00.push(from1);
        line_list_o_to00.push(to1);
        line_list_o_maxto00.push(max_to1);

        line_list_o_from1.push(from1);
        line_list_o_to1.push(to1);
        line_list_o_maxto1.push(max_to1);


        //let rowtext = rowcol[2].trim();

        rowtext = rowcol[2].trim();
        rowTTS = tts_8_REPLACE(rowtext);

        //[rowtext, rowTTS] = splitTTS(  rowcol[2] ); 


        let lastchar = rowtext.substring(rowtext.length - 1);
        if ((".?!").indexOf(lastchar) >= 0) {
            number_of_subtitle_endsentence++;
        }


        line_list_orig_text.push(rowtext);
        line_list_orig_tts.push(rowTTS.trim());
		line_list_orig_nFileR.push( nFileRow );  

        x1 = ixmin;

        for (x1 = x1; x1 <= tran_len; x1++) {
            if (from1 >= line_list_t_from1[x1]) {
                ixmin = x1;
                continue;
            }
            ixmax = ixmin;
            for (x2 = ixmin; x2 <= tran_len; x2++) {
                if (to1 >= line_list_t_to1[x2]) {
                    ixmax = x2;
                    continue;
                }
                break;
            }
            break;
        }
        line_list_o_tran_ixmin.push(ixmin);
        line_list_o_tran_ixmax.push(ixmax);
    }

} // end of get_orig_subtitle_text() 
//---------------------------------------


function tts_5_get_tran_subtitle_text() {

    let i, row, from1, to1, max_to1, rowcol;
    line_list_t_from00 = [];
    line_list_t_to00 = []; //
    line_list_t_maxto00 = [];
    line_list_t_from1 = []; //
    line_list_t_to1 = []; //
    line_list_t_maxto1 = []; //
    line_list_tran_text = [];

    line_list_t_from00.push(-1);
    line_list_t_to00.push(-1);
    line_list_t_maxto00.push(-1);

    line_list_t_from1.push(-1);
    line_list_t_to1.push(-1);
    line_list_t_maxto1.push(-1);
    line_list_tran_text.push("");

    if (sw_sub_tran == false) {
        return;
    }
    max_to1 = 0;

    let lastix = inp_row_tran.length - 1;


    for (i = inp_row_tran.length - 1; i > (inp_row_tran.length - 10); i--) {
        if (inp_row_tran[i].trim() == "") {
            continue;
        }
        row = inp_row_tran[i].trim();
        lastix = i;
        break;
    }
	var cols = [];
	//var nFileRow, nFile, nRows, 
	var row123;
    let pre_from, pre_to1;
    from1 = 0;
    to1 = -1;
    //-------------------------
    for (i = 0; i <= lastix; i++) {
        row = inp_row_tran[i].trim();
        if (row == "") {
            continue;
        }
		//if (row.substring(0,1) == ";") {row = row.substring(1)}
        pre_from = from1;
        pre_to1 = to1;
        //rowcol = [i, (i + 1), row];
		//------------------------
		// wbf:      ;;1;59;; BC DEF    ;;nFile;n.row;; row 
        cols     = (row + ";;;;;;").split(";;");
		/**
		nFileRow = (cols[1]+";;;").split(";");
		nFile    = nFileRow[0]; 
		nRows    = nFileRow[1];
		**/
		row123   = cols[2]	  ; 
		//console.log( "??ANTO  row=" + row + "\n\tcols=" , cols, "\n\tnFileRow=" , nFileRow , "\n\tnFile="+nFile + " nRows=" + nRows + "  row123=" + row123 + "<==");  	
		rowcol   = [i, (i + 1), row123.trim()];
		//------------	
		
        from1 = parseFloat(rowcol[0]);

        if (from1 < pre_from) {
            //   translated language subtitle line ignored because not in sequence of time:  pre=" + pre_from + " now=" + row );
            continue;
        }
        if (from1 == pre_to1) { // to avoid that in the same moment there be 2 subtitles lines actives when actualy one is the following of the other  (upd 2022_03_16)
            from1 += 0.0004;
        }

        to1 = parseFloat(rowcol[1]);
        max_to1 = Math.max(max_to1, to1);


        line_list_t_from00.push(from1);
        line_list_t_to00.push(to1);
        line_list_t_maxto00.push(max_to1);

        line_list_t_from1.push(from1);
        line_list_t_to1.push(to1);
        line_list_t_maxto1.push(max_to1);

        line_list_tran_text.push(rowcol[2]);
    }

} // end of get_tran_subtitle_text() 


//-----------------------------------------------------------
function tts_5_fun_update_html_with_last_session_values() {

    if ((sw_no_subtitle) || (sw_sub_onfile == false)) {
        LS_clip_checked_sw_type = radio_type1_SECONDS; // only clip in seconds 

    }
    if ((typeof LS_clip_checked_sw_type === "undefined") || (LS_clip_checked_sw_type == "")) {
        LS_clip_checked_sw_type = radio_type1_SECONDS; // only clip in seconds 

    }


    let p_list = LS_stor_playnext_replay.split(";;");


    ele_playNextVa_from_hhmmss_value = p_list[0];
    ele_replayVa_from_hhmmss_value = p_list[1];
    ele_replayVa_to_hhmmss_value = p_list[2];
    if (ele_playNextVa_from_hhmmss_value == 0) {
        ele_playNextVa_from_hhmmss_value = "00:00:00.000";
    }
    if (ele_replayVa_from_hhmmss_value == 0) {
        ele_replayVa_from_hhmmss_value = "00:00:00.000";
    }
    if (ele_replayVa_to_hhmmss_value == 0) {
        ele_replayVa_to_hhmmss_value = "00:00:00.000";
    }


    if (ele_replayVa_from_hhmmss_value < ele_playNextVa_from_hhmmss_value) {
        ele_playNextVa_from_hhmmss_value = ele_replayVa_from_hhmmss_value;
    }

    let secs4 = tts_5_fun_N_hhmmss_to_secs(ele_replayVa_to_hhmmss_value);


    let hhmmss4 = tts_5_fun_N_secs_to_hhmmssmmm(secs4 - 1.0);
    ele_replayVa_to_hhmmss_value = hhmmss4;

    if (sw_is_no_videoaudio == false) {
        if (vid) {
            vid.currentTime = secs4;
        }
    }

    let secs1d = ("0" + LS_clip_secs1d_value + "-false-true").split("-");

    let end_dialog_sentence = secs1d[1].trim();
    let end_dialog_overlap = secs1d[2].trim();

    //-----------
    end_dialog_sentence = "false";
    end_dialog_overlap = "false";
    sw_end_dialog_sentence = false;
    sw_end_dialog_overlap = false;

    if ((number_of_subtitle_endsentence == 0) && (number_of_subtitle_time_overlap == 0)) {
        //document.getElementById("irad_r1").style.display = "none"; 		
    } else {
        if (number_of_subtitle_time_overlap > 0) {
            end_dialog_overlap = "true";
            sw_end_dialog_overlap = true;
        }
        if (number_of_subtitle_endsentence > 0) {
            end_dialog_sentence = "true";
            sw_end_dialog_sentence = true;
        }
    }


    tts_5_fun_update_LS_sub_endbeg_delta();



} // end of fun_update_html_with_last_session_values()

//-------------------------------------------------------------------------------	

function tts_5_fun_N_hhmmss_to_secs(x) {
    let col00 = ("0:0:0:" + x.toString().replace(",", ".")).split(":");
    let col = col00.slice(-3);

    let hh1 = col[0] * 3600;
    let mm1 = col[1] * 60;
    let ss1 = col[2] * 1;

    let secs = Math.round((hh1 + mm1 + ss1) * 1000) / 1000;
    return secs;
}

//--------------------------------------------------

function tts_5_fun_N_secs_to_hhmmssmmm(from1) {
    // surely a film lasts less than 10 hour

    if ((isNaN(from1)) || (from1 == 0)) {
        return "00:00:00.000";
    }
    let secs = 3600 * 10 + 1 * from1;

    let hhmmss2 = new Date(secs * 1000).toISOString().substr(11, 12);
    let hhmmss = "0" + hhmmss2.substring(1);

    return hhmmss;
}

//------------------------------

function tts_5_fun_update_LS_sub_endbeg_delta() {

    if (isNaN(LS_sub_beg_delta)) {
        LS_sub_beg_delta = parseFloat("0" + LS_sub_beg_delta);
    }
    if (isNaN(LS_sub_end_delta)) {
        LS_sub_end_delta = parseFloat("0" + LS_sub_end_delta);
    }

    let beg_val = LS_sub_beg_delta; // can be positive or negative 
    let end_val = LS_sub_end_delta; // can be positive or negative 
  
    let beg_PM = "plus";
    let end_PM = "plus";
    if (beg_val < 0) {
        beg_PM = "minus";
    }
    if (end_val < 0) {
        end_PM = "minus";
    }
   
}

//--------------------------------------------

function tts_5_fun_setMinMaxIxClip() {

    MIN_ixClip = 0;
    MAX_ixClip = 0;

    if (line_list_o_from1.length > 3) {
        MAX_ixClip = line_list_o_to1.length - 1;
        return;
    }
    return;
}

//-------------------------------------------

function tts_5_former_onclick_tts_ClipSub_All(this1) {

    var this1_checked = true;

    var sw_all = true;
  
    if (sw_sub_onfile == false) {
        return;
    }
    if (sw_sub_orig == false) {
        return;
    }

    save_last_oneOnlyRow = "";
    save_last_oneOnly_idtr = "";

    ele_last_tran_line = null;

    if (this1_checked == false) {
        fun_checked_false_goBack(this1);
        console.log("\tonclick_tts_ClipSub_All() return ");
        return;
    }



    tts_5_fun_build_all_clip();

   

    if (sw_all) {
        clipFromRow = clipFromRow_min;
        clipToRow = line_list_orig_text.length - 1;
        sw_CLIP_play = true;
        Clip_startTime = 0;
        if (sw_is_no_videoaudio == false) {
            Clip_stopTime = vid.duration;
        }
    } else {
        clipFromRow = parseInt(ele_replayVa_from_row_value);
        clipToRow = parseInt(ele_replayVa_to_row_value);
        sw_CLIP_play = true;
        Clip_startTime = parseFloat(ele_replay_from_secs_innerHTML);
        Clip_stopTime = parseFloat(ele_replay_to_secs_innerHTML);
    }

    //  xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx  now html elements for clip section are ready xxxxxxxxxxxxxxxxxxxxxxxxxxxx

    let eleF = document.getElementById("b1_" + clipFromRow);
    let eleT = document.getElementById("b2_" + clipToRow);

    //onclick_tts_arrowFromIx(eleF, clipFromRow, "2 tts_5_former_onclick_tts_ClipSub_All()" );
    //onclick_tts_arrowToIx(eleT, clipToRow);



    ele_clip_subtext.style.display = "block";


    //-------------------
    //set new dimension to be the almost the body dimension 

 
    let new_eleTabSub_clientWidth = eleTabSub.clientWidth;
    let new_eleTabSub_clientHeight = eleTabSub.clientHeight;
    //--- translate to style width/height
    eleTabSub.style.width = (new_eleTabSub_clientWidth - eleTabSub_diff_clientW) + "px"; // set  new dimension			
    eleTabSub.style.height = (new_eleTabSub_clientHeight - eleTabSub_diff_clientH) + "px"; // set  new dimension
    //----------------

    var elediv20 = document.getElementById("id_div20");

    clip_reset_BG_color = tts_5_get_backgroundColor(elediv20);
    eleTabSub.style.backgroundColor = clip_reset_BG_color;
    ele_clip_subtext.style.backgroundColor = clip_reset_BG_color;

} // end of former_onclick_tts_ClipSub_All()

//---------------------------------------------


//-----------------

function tts_5_fun_build_all_clip() {
   
    if (sw_sub_onfile == false) {
        return;
    }
    if (sw_sub_orig == false) {
        return;
    }

    save_last_oneOnlyRow = "";
    save_last_oneOnly_idtr = "";

    ele_last_tran_line = null;
   

    let clipSub_showTxt, txt1;
   

    clipFromRow = 0;
    clipToRow = line_list_orig_text.length - 1;


    sw_CLIP_play = true;
    Clip_startTime = 0;


    //--------------------------	
    var z3Beg = 0;
    for (let z3 = clipFromRow; z3 <= clipToRow; z3++) {
        txt1 = line_list_orig_text[z3].trim();
        if (txt1 == "") {
            continue;
        }
        if (txt1 == "_.") {
            continue;
        }
        z3Beg = z3;
        break;
    }
    if (z3Beg > clipFromRow) {
        clipFromRow = z3Beg;
        clipFromRow_min = clipFromRow;
    }

    begix = clipFromRow;
    endix = clipToRow;
    var text_tts;
    clipSub_showTxt = "\n";

	
	var word3="", ixUnW3="", totRow3="", wLemma3="", wTran3="";
	
	if (word_to_underline != "") {
		[word3, ixUnW3, totRow3, wLemma3, wTran3] = word_to_underline.split(",")
	}	
	if ((wLemma3 == "-") || (wLemma3 == "")) { wLemma3 = word3; }     
	
	var wordToBold = word3; 
	var wordLemma  = wLemma3;
	var wordTran   = wTran3; 
	var head1 = '<span style="color:blue;">' + wordLemma + '</span>' + 
		    '<br><span style="font-size:0.8em;">' + wordTran + '</span>';  
	document.getElementById("id_headWord").innerHTML = head1; 
	if ((word3=="") || (word3 == undefined)) {head1="";} 
	document.getElementById("id_headWord").style.textAlign = "center";
	var nFileR ; var nfile_zero
	
    for (let z3 = clipFromRow; z3 <= clipToRow; z3++) {

        txt1 = line_list_orig_text[z3].trim();
		if (txt1 == "") {
			continue;              //2 luglio 
		}
        text_tts = line_list_orig_tts[z3].trim();
		
        let trantxt1 = "";
		let txt1p=""; 

        for (let g = line_list_o_tran_ixmin[z3]; g <= line_list_o_tran_ixmax[z3]; g++) {
            let txt2 = line_list_tran_text[g];
            if (txt2.indexOf(sayNODIALOG) >= 0) {
                txt2 = "...";
            }
            trantxt1 += "<br>" + txt2;
        }
        if (trantxt1.substr(0, 4) == "<br>") {
            trantxt1 = trantxt1.substring(4);
        }
		txt1p = evidenzia( wordToBold , txt1);      // function defined  in "wordsByFrequence.js" file
		
		nFileR = line_list_orig_nFileR[z3] 
		if (nFileR.substr(0,1) == "0") { nfile_zero = "0";} else {nfile_zero="1";}
		
		
        let rowclip = string_tr_xx.replaceAll("§1§", z3).
			replaceAll("§4txt§", txt1).replaceAll("§5txt§", trantxt1).
			replaceAll("§4ptxt§", txt1p).
			replaceAll("§ttstxt§", text_tts).
			replaceAll("§6§", nFileR ).
			replaceAll("§nfile§",nfile_zero)	;		
        
		clipSub_showTxt += rowclip + "\n";
		
			
    } // end of for z3
  

    eleTabSub_tbody.innerHTML = clipSub_showTxt;



    sw_tts = true;



    //  xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx  now html elements for clip section are ready xxxxxxxxxxxxxxxxxxxxxxxxxxxx

    let eleF = document.getElementById("b1_" + clipFromRow);
    let eleT = document.getElementById("b2_" + clipToRow);


    onclick_tts_arrowFromIx(eleF, clipFromRow, "1.2 tts_5_fun_build_all_clip()"  );
    onclick_tts_arrowToIx(eleT, clipToRow, 2);



    ele_clip_subtext.style.display = "block";

    //-------------------
    //set new dimension to be the almost the body dimension 

    let new_eleTabSub_clientWidth = eleTabSub.clientWidth;
    let new_eleTabSub_clientHeight = eleTabSub.clientHeight;
    //--- translate to style width/height
    eleTabSub.style.width = (new_eleTabSub_clientWidth - eleTabSub_diff_clientW) + "px"; // set  new dimension			
    eleTabSub.style.height = (new_eleTabSub_clientHeight - eleTabSub_diff_clientH) + "px"; // set  new dimension
    //----------------

    var elediv20 = document.getElementById("id_div20");

    clip_reset_BG_color = tts_5_get_backgroundColor(elediv20);
    eleTabSub.style.backgroundColor = clip_reset_BG_color;
    ele_clip_subtext.style.backgroundColor = clip_reset_BG_color;

} // end of fun_build_all_clip()  

//----------------------------------------------------

//------------------------------------------------
function tts_5_highlightRow(ele1) {
    tts_5_removeLastBold();

    ele1.classList.add("boldLine");
    ele1.style.backgroundColor = "yellow";
    ele1.parentElement.style.border = "1px solid black";

    last_BoldRow = ele1;
}
//-------------------------
function tts_5_removeLastBold() {
    if (last_BoldRow) {
        //console.log("removeLastBold() " + last_BoldRow.id);
        var epar1 = last_BoldRow.parentElement;
        if (epar1) {
            //console.log("\tlast_BoldRow.parentElement.id=" + epar1.id);
            var epar2 = epar1.parentElement;
            if (epar2) {
                //console.log("\tlast_BoldRow.parentElement.parentElement.id=" + epar2.id);
            } else {
                //console.log("\tlast_BoldRow.parentElement.parentElement missing");
            }
        } else {
            //console.log("\tlast_BoldRow.parentElement missing");
        }


        last_BoldRow.classList.remove("boldLine");
        last_BoldRow.style.backgroundColor = null;
        last_BoldRow.parentElement.style.border = null;
        var last_ele1_tr = last_BoldRow.parentElement.parentElement;
        last_ele1_tr.style.backgroundColor = "lightgrey";
    }
    //if (sw_tts) tts_3_remove_last_bold("idc_", false);
	if (sw_tts) tts_3_remove_last_bold("idc_", false);                 //  variaz. 21 giugno
}

//---------------------------------------

function tts_5_fun_invisible_prev_fromto(interX) {
    // eliminate bold of the  previous group of lines, unless this is a line in them  
    [begix, endix] = fromIxToIxLimit; // previously set 
    if ((interX >= begix) && (interX <= endix)) {
        return;
    }
    var id_pre_tr_beg_space = "idtr_" + begix + "_m2";
    var id_pre_tr_head = "idtr_" + begix + "_m1";
    var id_post_tr_end_space = "idtr_" + (endix + 1) + "_m2";


    if (fromIxToIxButtonElement[0]) {
        fromIxToIxButtonElement[0].style.backgroundColor = null;
        if (document.getElementById(id_pre_tr_beg_space)) {
            document.getElementById(id_pre_tr_beg_space).style.display = "none";
        }
        if (document.getElementById(id_pre_tr_head)) {
            document.getElementById(id_pre_tr_head).style.display = "none";
        }
    }
    if (fromIxToIxButtonElement[1]) {
        fromIxToIxButtonElement[1].style.backgroundColor = null;
        if (document.getElementById(id_post_tr_end_space)) {
            document.getElementById(id_post_tr_end_space).style.display = "none";
        }
    }

} // end of fun_invisible_prev_fromto()

//------------------------------------------

function tts_5_fun_copyHeaderSelected() {

    let id1;
    let inBeg, inEnd;

    id1 = "idb_" + begix + "_m";
    var thisX = document.getElementById(id1);
    if (thisX == false) {
        return;
    }

    inBeg = begix;
    inEnd = endix;
    if (begix > endix) {
        inBeg = endix;
        inEnd = begix;
    }

    var style0, style1;

    tts_5_fun_oneRowZZ();

    for (var g = inBeg; g <= inEnd; g++) {
        id1 = "idb_" + g;
        thisX = document.getElementById(id1);
        tts_5_fun_oneRow11H();
    }

    //--------------
    function tts_5_fun_oneRowZZ() {
        if (thisX == false) {
            return;
        }
		if ( thisX.children.length < 1) { return; }
        if (thisX.children[0].style.display == "none") { // no openbook   
            style0 = "none"; //  hide opened book image  
            style1 = "block"; //  show closed book image 
        } else {
            style0 = "block"; // show opened book image  
            style1 = "none"; // hide closed book image 	
        }
    }
    //--------------
    function tts_5_fun_oneRow11H() {
        if (thisX == false) {
            return;
        }
		if ( thisX.children.length < 1) { return; }
        thisX.children[0].style.display = style0; // show/hide  opened book image  
        thisX.children[1].style.display = style1; // show/hide closed book image 
        let subid = thisX.id.replace("idb", "idc");

        if (subid.substring(subid.length - 2) == "_m") {
            return;
        }
        let ele1 = document.getElementById(subid);
        if (style0 == "block") {
            tts_5_fun_makeTextVisible(ele1);
        } else {
            tts_5_fun_makeTextInvisible(ele1);
        }
    }
} // end of fun_copyHeaderSelected()

//-----------------------------
//----------------------------
function tts_5_fun_makeTextInvisible(element) {
    if (element) element.style.display = "none";
}
//----------------------------
function tts_5_fun_makeTextVisible(element) {
    if (element) element.style.display = "block";
}

//----------------------------------------------

//------------------------------------------
function tts_5_get_backgroundColor(ele0) {
    // since the backgroundColor is not inherited the 'getComputedStyle()' cannot get it from the parent  
    var ele1 = ele0;
    var eleP = ele1;
    for (var i = 0; i < 99; i++) {
        ele1 = eleP;
        if (ele1 == null) {
            return "";
        }
        var bgColor = window.getComputedStyle(ele1).getPropertyValue("background-color");
        var bgc1 = bgColor.split("(");
        var bgc2 = bgc1[1].split(")");
        var bgx = (bgc2[0].replaceAll("  ", "").replaceAll(" ", "").replaceAll(",", " ").trim()).substring(0, 7);
        if (bgx == "0 0 0 0") {
            eleP = ele1.parentElement;
        } else {
            //console.log(" get_backgroundColor(ele1=" + ele1.id + ")  ==> bgColor=" + bgColor); 
            return bgColor;
        }
    }
    console.log(" get_backgroundColor(ele1)  ==> bgColor=" + "  empty2");
    return "";
}



//---------------------------------------

function tts_5_fun_scroll_tr_toTop( this1, swDeb ){

	var ele_tr_target = this1; // .parentElement.parentElement ; 
	var numid= parseInt(ele_tr_target.id.substring(5).split("_")[0] );  // id="idtr_xx"   id="idtr_xx_m1"	
	
	if (swDeb) console.log("fun_scroll_tr_toTop() 1 numid=" + numid); 
	
	var ele_tr_nearest, diff_off=0;
	if (numid > 1) {
		ele_tr_nearest = document.getElementById(  "idtr_" +(numid-1)); 	
		if (ele_tr_nearest) {
			diff_off = ele_tr_target.offsetTop - ele_tr_nearest.offsetTop; 	
		}
	} else {
		ele_tr_nearest = document.getElementById(  "idtr_" +(numid+1)); 			
		if (ele_tr_nearest) {
			diff_off = 0 - (ele_tr_target.offsetTop - ele_tr_nearest.offsetTop); 	
		}
	}	
	var ele_container = document.getElementById("id_section_row");
	if (swDeb) { 
		var compHeight = window.getComputedStyle(ele_container).getPropertyValue("height");	
	
		console.log("fun_scroll_tr_toTop() 2  ele_tr_target.offsetTop - diff_off=" +  ele_tr_target.offsetTop + " - " + diff_off );
	    console.log("   ele_container  (id_tabSub).id = " + ele_container.id +
			" offsetHeight=" + ele_container.offsetHeight + " offsetTop=" + ele_container.offsetTop + 
			" style.height=" + ele_container.style.height + " computed height=" + compHeight);  
	} 
	if (numid < 2) {
		ele_container.scrollTop = 0; 
		return; 
	}
	try{
		ele_container.scrollTop = ele_tr_target.offsetTop - diff_off;
	} catch(e1) {
		console.log("fun_scroll_tr_toTop(this1)" + " this1.id=" + this1.id ); 
		console.log("  \t ele_tr_target.id=" + ele_tr_target.id); 
		ele_container.scrollTop = ele_tr_target.offsetTop - diff_off;
	}
} // end of fun_scroll_tr_toTop() 
//-------------------------------------------"use strict";
/*  
LineByLine: A tool to practice language comprehension
Antonio Cigna 2023
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//----------------------------------------------
var currScript = document.currentScript.src; var bar1 = currScript.lastIndexOf("\\");var bar2 = currScript.lastIndexOf("/"); 
//console.log("LOADED file SCRIPT " + currScript.substring( 1+Math.max(bar1,bar2) )) ;	
//----------------------------------------------------------------------------------------
/***
	may be the case that a synthetic voice for a language does not exist ( eg. in case of a dead language as latin )
    this function try to solve this problem 
			by changing the text so that it read by an other language voice can sounds like the original one
			
	for instance: latin word 'coeli' modified in 'celi' to be read by an Italian voice sounds as ecclesiastical Latin.   
	
**/
//----------------------------------------------------------------------------------------
function tts_8_REPLACE( txt1 ) {  
	// this function is called for each original language line text
   	// if tts text is equal to the original then this function should return "" 
	//
	// this function should be driven by the array 'language_parameters'; 
	//                         language_parameters[0] language id of the voice
	//                         language_parameters[1] number of voices 
	//						   language_parameters[2,...] routine to use and other parameters if needed
	//
	//-----------------------------------------	
	// eg. var  tts_txt = txt1.replaceAll("oe","e");  return tts_txt;  
	
	return ""; 
	
} // end of add_text_to_be_spoken_line()
//----------------------------------------------------------
"use strict";
/*  
LineByLine: A tool to practice language comprehension
Antonio Cigna 2023
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//----------------------------------------------
var currScript = document.currentScript.src; var bar1 = currScript.lastIndexOf("\\");var bar2 = currScript.lastIndexOf("/"); 
//console.log("LOADED file SCRIPT " + currScript.substring( 1+Math.max(bar1,bar2) )) ;	
//----------------------------------------------------------------------------------------

function NOT_USED_tts_9_get_wantedVoices_from_URL_parameters() { 

	// eg. file:///D:/ANTONIO/ClipByClip_V2/cbc_base/cbc_player/cbc_PLAYER.html?p_title=OneLesson
	
	// this function is called in script_9... file 
	
	
	//console.log("\nXXXXXXXXXXXXXXXXXXXXXXXXXX\nX get_html_URL_parameters() \nXXXXXXXXXXXXXXXXXXXXXXXXXX\n"); 
	
	html_parms_queryString = window.location.search;
	console.log("window.location.search = html_parms_queryString="  +  html_parms_queryString);
	
	if ((html_parms_queryString == undefined ) || (html_parms_queryString.trim() == "") ) {	
		return; 	
	} 	
	const urlParams = new URLSearchParams(html_parms_queryString);		
	var selected_language_fromBuilder = decodeURI( urlParams.get("p_sel_voicelang" ) );    // eg.  en,6
	console.log("selected_language_fromBuilder = " + selected_language_fromBuilder);  
	if (selected_language_fromBuilder == "null") {
		language_parameters = ["xx","0"];  
		selected_lang_id = "xx"; 
		selected_numVoices=0;  	
	} else { 	
		language_parameters = (selected_language_fromBuilder+",,,").split(","); 
		selected_lang_id    = (language_parameters[0]+"   ").trim(); // .substr(0,2); 
		selected_numVoices  = parseInt("0" + language_parameters[1].trim() )  ; 		 
	}
	if (selected_numVoices > maxNumVoices)  selected_numVoices = maxNumVoices;
	
	console.log("parameters from Builder: language '" + selected_lang_id  + "' num.Voices=" + selected_numVoices); 
	
} // end of  get_wantedVoices_from_URL_parameters()

//------------------------------------
function tts_9_toBeRunAfterGotVoicesPLAYER() { 
	console.log(" tts_9_toBeRunAfterGotVoicesPLAYER()")
	//
	// called by asynchronous code in fcommon_load_all_voices()  (if numVoices > 0) 
	//
	//-------------------------------------
	//console.log("script_9 toBeRunAfterGotVoices()"); 	
	//console.log("script_9 voices.length=" + voices.length);
	
	tts_2_fill_the_voices( "2tts_9_toBeRunAfterGotVoicesPLAYER" );  //  in script_2... file  
	
	string_tr_xx = "\n" + prototype_tr_m2_tts + "\n" + prototype_tr_m1_tts + "\n" + prototype_tr_tts; 
	word_tr_allclip =  "\n" + prototype_word_tr_m2_tts + "\n" + prototype_word_tr_m1_tts + "\n" + prototype_word_tr_tts; 
		
	tts_3_getUrlData3(); 	
	
	//--------------------
} // end of toBeRunAfterGotVoices() 	

//=============================================
"use strict";
/*  
LineByLine: A tool to practice language comprehension
Antonio Cigna 2023
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//----------------------------------------------
var currScript = document.currentScript.src; var bar1 = currScript.lastIndexOf("\\");var bar2 = currScript.lastIndexOf("/"); 
//console.log("LOADED file SCRIPT " + currScript.substring( 1+Math.max(bar1,bar2) )) ;	
//----------------------------------------------------------------------------------------
// the functions in this file are copied from the "builder" part of "ClipByClip" and "RaS" function
//----------------------------------------------------------------------------------------------------

function get_subtitle_strdata(inpsub) {
    //---------------------------------------
    // get subtitles in srt or vtt format 
    // and return a string with a line for each subtitle group     
    //---------------------------------------
    /* 
		sample of a srt file
			14
			00:01:08,000 --> 00:01:10,000      ixtime[13] = 52
			one line of text
			other lines of text
			 
			15
			00:01:10,000 --> 00:01:15,000      ixtime[14] = 57
			lines
			of texts   
		---------------------------------------------
		sample of a vtt file 
			WEBVTT                             (this is the first line which identify a webvtt subtitle file  		
			STYLE
			::cue(.S1) {
			color: white;
			...
			}
			REGION
			id:R1
			...

			1
			00:00:01.000 --> 00:00:02.040 region:R1 align:center
			<c.S1>one line</c>

			2
			00:00:02.600 --> 00:00:03.640 region:R1 align:center
			<c.S1>another line</c>

    */
    //-------------------------------------------------------


    var subline = inpsub.trim().split("\n"); // <===  INPUT  

    //----------------------------
    var i, txt1;



    var list_time_from = [0];
    var list_time_to = [0];
    var list_text = [""];
    var idsrt = [0];
    //------------------------------
    subline.push("");

    // look at what follows the time line,  stop when there is a blank line ( but might be missing ) 

    //list_start=[];  
    var line = "";
    var preline = "";

    var list_ix_beg = [0];
    var list_ix_end = [0];

    list_time_from = [0];
    list_time_to = [0];
    list_text = [""];
    idsrt = [""];
    var num = 0;
    /***
    subline.push(""); 
    subline.push("99999999"); 
    subline.push("99:99:99,999 --> 99:99:99,999"); 
    **/

    var isTime;
    var time_err;
    var timehh_fromX;
    var timehh_toX;

    //---
    function close_previous_group() {
        // the previous line should be a number 
        num = list_ix_end.length - 2;
        if (list_ix_beg[num] >= 0) {
            if (isNaN(preline)) {
                list_ix_end[num] = i;
            } else {
                list_ix_end[num] = i - 1;
            }
        }
    } //end of  close_previous_group()

    //-----------------------------------
    for (i = 0; i < subline.length; i++) {
        preline = line;
        line = subline[i];
        [isTime, time_err, timehh_fromX, timehh_toX] = isTimeLine(line);
        if (isTime == false) {
            continue;
        }
        if (time_err) {
            // error in time line
            list_ix_beg.push(-1);
            list_ix_end.push(-1);
            list_time_from.push(0);
            list_time_to.push(0);
            list_text.push("");
            idsrt.push("");
            close_previous_group();

            continue;
        }

        var timehh_from = timehh_fromX[1];
        var timehh_to = timehh_toX[1];

        var time_secs_from = timehh_fromX[0];
        var time_secs_to = timehh_toX[0];

        list_ix_beg.push(i);
        list_ix_end.push(i);


        list_time_from.push(time_secs_from);
        list_time_to.push(time_secs_to);

        list_text.push("");

        idsrt.push(timehh_from + " " + timehh_to);
        close_previous_group();

    }

    list_ix_end[list_ix_end.length - 1] = subline.length;


    var txt00;
    for (var k = 1; k < list_ix_end.length; k++) {
        txt1 = "";
        var ixfrom = list_ix_beg[k];
        if (ixfrom < 0) {
            continue;
        }
        var ixto = list_ix_end[k];
        for (var z = ixfrom + 1; z < ixto; z++) {
            txt00 = subline[z].trim();

            //console.log("z=" + z + " " + subline[z] + "<==");

            if (txt00 != "") txt1 += "\n" + txt00 + " ";
        }
        list_text[k] = txt1.substring(1);
    }

    return [list_time_from, list_time_to, list_text, idsrt];


    // add_nodialog_clips2( list_time_from, list_time_to, list_text ,idsrt );	

} // end of get_subtitle_strdata()
//----------------------------------

function get_timehhmmss(str0) {
    /*
    	try to manager time in srt or vtt  even when non correctly written ( maybe it's out of automatic translation)

    	expected ==>    00:12:45,123                   ( hh:mm:ss,mmmm ) 
    					00:12:45.123 other staff vtt   ( hh:mm:ss.mmmm )
    					00:12:45                       ( hh:mm:ss ) 	
    	can manage	1: 2: 45,123      transformed to 00:02:45.123
    				12:45,123         transformed to 00:12:45.123
    				45,123            transformed to 00:00:45.123					
    */

    var str1 = (str0 + "").trim().replace(",", ".").replaceAll(": ", ":").replaceAll(": ", ":");

    var tt1 = str1.split(":");
    var len1 = tt1.length;
    if (len1 < 3) {
        str1 = "00:" + str1;
        if (len1 < 2) {
            str1 = "00:" + str1;
        }
        tt1 = str1.split(":");
    }
    var tHH = tt1[0].trim();
    var tMM = tt1[1].trim();
    var tSS = tt1[2].trim().split(" ")[0]; // ignore whatever follows ( that is the case of vtt subtitiles) 


    if (isNaN(tHH)) {
        return;
    }
    if (isNaN(tMM)) {
        return;
    }
    if (isNaN(tSS)) {
        return;
    }
    var nHH = parseInt(tHH);
    var nMM = parseInt(tMM);
    var nSS = parseFloat(tSS);


    if ((nHH >= 60) || (nHH < 0) || (nMM >= 60) || (nMM < 0) || (nSS >= 60) || (nSS < 0)) {
        return;
    }

    var seconds = nHH * 3600 + nMM * 60 + nSS;

    tHH = (100 + nHH).toString().substring(1);
    tMM = (100 + nMM).toString().substring(1);
    tSS = (100.0000001 + nSS).toString().substr(1, 6);

    return [seconds, tHH + ":" + tMM + ":" + tSS];

} // end of get_timehhmmss(); 

//----------------------------------------

function isTimeLine(str1) {

    var line = str1.trim().
    replace("- ->", "-->").
    replace(" -> ", " --> ").
    replace("-- >", "-->"); // replace used to avoid random error from automatic translation  
    if (line.indexOf("-->") < 0) {
        return [false, true, "", ""];
    }
    var part = line.split("-->");
    var timehh_fromX = get_timehhmmss(part[0]);
    var timehh_toX = get_timehhmmss(part[1]);

    if ((timehh_fromX === undefined) || (timehh_toX === undefined)) {
        console.log("error in " + line);
        return [true, true, part[0], part[1]];
    }

    return [true, false, timehh_fromX, timehh_toX];
}
//==============================================================================