"use strict";
/*  
Words By Frequence: A tool to practice language comprehension
Antonio Cigna 2023
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//----------------------------------------------
/***
function js_call_go() {
				 goStart(); 
			}
***/			
			
function js_call_go() {  // called by  html page body onload
	var msg1 = "html loaded"; 
	//console.log("html function js_call_go: ", msg1, "\n") 
	
	go_passToJs_html_is_ready(msg1,  "") ;  //  "js_go_go_is_ready"); 

}
//-------------------------------
/**
function js_go_go_is_ready(msg1) {
	
	var msg2 = msg1 + " " + "html ready"

	//console.log("html function js_go_go_is_ready: ", msg2, "\n" )
	
	go_passToJs_html_and_go_ready( msg2, "")
}
***/
//----------------------------------------------	