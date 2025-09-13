package wbfSubPack

import (  
	"fmt"
    "strings"
	"runtime"
)
	
//----------------------------------------------------------------
func go_exec_js_functionPlus(js_function string, inpstr string, js_parm string, js_caller string) {
	var goFunc string 
 	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		goFunc = strings.ReplaceAll(details.Name(), "main.","")
	} else {
		goFunc=""
	}	
	k1:= strings.Index(js_function, "(") 
	if k1 > 0 {	
		js_function = strings.TrimSpace(js_function[0:k1] )
	}	
	k1 = strings.Index(js_function, ",") 
	if k1 > 0 {	
		js_function = strings.TrimSpace(js_function[0:k1] )
	}	
	js_function = strings.TrimSpace(js_function)
	
	/**
	fmt.Println("  js_function=" + js_function + "<==")
	fmt.Println("    js_parm  =" + js_parm )
	fmt.Println("    js_caller=" + js_caller)
	fmt.Println("    goFunc   =" + goFunc)
	**/
	
	/*
	This function executes a javascript eval command 
	which must execute a function by passing string constant to it. 
	Should this string contain some new line, e syntax error would occur in eval the statement.
	
	To avoid this kind of error, the string argument (inpstr) of the javascript function (js_function) 
	is forced to be always enclosed in back ticks trasforming it in "template literal".  
	Just in case back ticks and dollars are in the string, they are replaced by " "   	
	*/
	
	inpstr = strings.ReplaceAll( inpstr, "`", " "   ); 	   	 
	inpstr = strings.ReplaceAll( inpstr, "$", "&dollar;"); 	
	
	//fmt.Println(cyan("go_exec_js_function "), "js_function=", js_function, "js_parm=", js_parm, " jsInpFunction=", jsInpFunction, " gofunc=", "go=" + goFunc)
	
	//js_function = "miaProva"; inpstr = "MIA prova di inpstr"; fmt.Println("go_exec_js_functionPlus ", "	inpstr=>" + inpstr + "<==")
	
	evalStr := fmt.Sprintf( "%s(`%s`,`%s`,`%s`,`%s`);",  js_function, strings.TrimSpace( inpstr ), js_parm, js_caller, goFunc ) ; 
	
	//fmt.Println(cyan("go_exec_js_function "), "evalStr=" + evalStr); 
	
	ui.Eval(evalStr)
	
} // end of go_exec_js_functionPlus

//----------------------------------------------------------------
//----------------------------------------------------------------
func go_exec_js_function(js_function0 string, inpstr string) {
	var goFunc string 
 	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		goFunc = strings.ReplaceAll(details.Name(), "main.","")
	} else {
		goFunc=""
	}
	js_fun        := strings.Split( (js_function0 + ",,,,") ,",") 	
	js_function   := strings.TrimSpace( js_fun[0] )
	if js_function == ""  { return }
	jsInpFunction := strings.TrimSpace( js_fun[1] )
	
	js_parm:=""
	k1:= strings.Index(js_function, "(") 
	if k1 > 0 {
		js_parm     = strings.ReplaceAll(  js_function[k1+1:], ")","")			
		js_function = strings.TrimSpace(js_function[0:k1] )
	}	
	/**
	fmt.Println("  js_function=" + js_function)
	fmt.Println("      js_parm=" + js_parm)
	fmt.Println("jsInpFunction=" + jsInpFunction)
	fmt.Println("       goFunc=" + goFunc)
	**/
	
	/*
	This function executes a javascript eval command 
	which must execute a function by passing string constant to it. 
	Should this string contain some new line, e syntax error would occur in eval the statement.
	
	To avoid this kind of error, the string argument (inpstr) of the javascript function (js_function) 
	is forced to be always enclosed in back ticks trasforming it in "template literal".  
	Just in case back ticks and dollars are in the string, they are replaced by " "   	
	*/
	inpstr = strings.ReplaceAll( inpstr, "`", " "   ); 	   	 
	inpstr = strings.ReplaceAll( inpstr, "$", "&dollar;"); 
	
	//fmt.Println(cyan("go_exec_js_function "), "js_function=", js_function, "js_parm=", js_parm, " jsInpFunction=", jsInpFunction, " gofunc=", "go=" + goFunc)
	//fmt.Println("go_exec_js_function ", "	inpstr=", inpstr)
	
	evalStr := fmt.Sprintf( "%s(`%s`,`%s`,`%s`,`%s`);",  js_function, inpstr, js_parm, "js=" + jsInpFunction, "go=" + goFunc ) ; 
	
	//fmt.Println(cyan("go_exec_js_function "), "evalStr=" + evalStr); 
	
	ui.Eval(evalStr)
	
} // end of go_exec_js_function
//----------------------------------------------------------------

func go_exec_js_functionX(js_function string, inpstr string ) {
	/*
	This function executes a javascript eval command 
	which must execute a function by passing string constant to it. 
	Should this string contain some new line, e syntax error would occur in eval the statement.
	
	To avoid this kind of error, the string argument (inpstr) of the javascript function (js_function) 
	is forced to be always enclosed in back ticks trasforming it in "template literal".  
	Just in case back ticks and dollars are in the string, they are replaced by their html symbols.   	
	*/
	inpstr = strings.ReplaceAll( inpstr, "`", " "   ); 	  
	//inpstr = strings.ReplaceAll( inpstr, "`", "&#96;"   ); 	 
	inpstr = strings.ReplaceAll( inpstr, "$", "&dollar;"); 
	
	evalStr := fmt.Sprintf( "%s(`%s`);",  js_function,  inpstr ) ; 
	
	ui.Eval(evalStr)
	
} // end of go_exec_js_function

//--------------------------------
