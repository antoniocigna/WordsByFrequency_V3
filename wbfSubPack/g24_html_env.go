package wbfSubPack

	import (
		"fmt"
		"os"
		"strings"		
		"github.com/zserge/lorca"	
		"github.com/lxn/win"	
	)
//--------------------------------------------------------
var scrX, scrY int = getScreenXY();


//------------------------------

var ui, err = lorca.New("", "", scrX, scrY); // crea ambiente html e javascript  // if height and width set to more then maximum (eg. 2000, 2000), it seems it works  


//---------------------
func getScreenXY() (int, int) {
	
	// use ==>  var x, y int = getScreenXY();
	
	var width  int = int(win.GetSystemMetrics(win.SM_CXSCREEN));
	var height int = int(win.GetSystemMetrics(win.SM_CYSCREEN));
	if width == 0 || height == 0 {
		//fmt.Println( "errore" )
		return 2000, 2000; 
	}	
	width  = width  - 20;  // subtraction to make room for any decorations 
	height = height - 40;  // subtraction to make room for any decorations 
	
	return width, height
}
//----------------------------------	

func begin_GO_HTML_Talk() { 	
	fmt.Println("func begin_GO_HTML_Talk"); 
	setHtmlEnv();	
}
//---------------

//------------------------------------
func setHtmlEnv() {	
	fmt.Println("func setHtmlEnv:  start load html")
    // load file html 	
	
	html_path = getCompleteHtmlPath( parameter_path_html ) 
	            
	fmt.Println("path html        = " + html_path)
	
	ui.Load("file:///" + html_path + string(os.PathSeparator) + "wordsByFrequency.html" ); 
	
	fmt.Println("\n", "func setHtmlEnv: wait for html ( javascript function js_call_go()", "\n")  
	
} // end of setHtmlEnv
//--------------------------------------------------------
//-------------------------
func getCompleteHtmlPath( path_html string) string {
	
	//curDir    := "D:/ANTONIO/K_L_M_N/LINGUAGGI/GO/_WORDS_BY_FREQUENCE/WbF_prova1_input_piccolo
	 
	curDir, err := os.Getwd()
    if err != nil {
		fmt.Println("setHtmlEnv() 3 err=", err )
        //log.Fatal(err)
    }	
				
	fmt.Println("curDir           = " + curDir ); 
	
	curDirBack  := curDir
	k1:= strings.LastIndex(curDir, "/") 
	k2:= strings.LastIndex(curDir, "\\") 
	if k2 > k1 { k1 = k2 } 
	curDirBack = curDir[0:k1] 	
	
	var newPath string = ""
	if strings.Index(path_html,":") > 0 {
		newPath = path_html
	} else if path_html[0:2] == ".." {
		newPath = curDirBack  + path_html[2:] 
	} else {
		newPath = curDir + path_html
	}
	return newPath 
} 
//------------------------
func putFileError( msg1, inpFile string) {
	err1:= `document.getElementById("id_startwait").innerHTML = '<br><br> <span style="color:red;">§msg1§</span> <span style="color:blue;">§inpFile§</span>';` ; 		
	err1 = strings.ReplaceAll( err1, "§msg1§", msg1 ); 	 
	err1 = strings.ReplaceAll( err1, "§inpFile§", inpFile); 	
	ui.Eval( err1 );	
}   

//-----------------------------------
