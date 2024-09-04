package wbfSubPack

	import (
		"fmt"
		"os"
		"os/signal"
		"strings"	
		"runtime"	
	)

//------------------------------------------------------
func MyMain() {

	fmt.Println("\n======================\n         My Main()  INIZIO di mainPack \n===============================\n")
	fmt.Println(  red("WordsByFrequence - wbfMain") )
	
	fmt.Println( "\ncolori:", red("rosso"), green("verde"), yellow("giallo"),  magenta("magenta"), cyan("ciano") , "\n"  )  
	
	//---------------
	val0, val1, val2, val3, val4 := getPgmArgs("-html",  "-input" , "-countNumLines" , "-maxNumLinesToWrite", "-lemmaformat")	
	//----------
	countNumLines      = val2
	maxNumLinesToWrite = val3
	lemmaFormat        = val4
	
	sw_lemma_word = (lemmaFormat == "lemma-word")
	
	if val0 != "" { parameter_path_html = strings.ReplaceAll(val0,"\\","/")  } 
	if val1 != "" { file_inputControl   = strings.ReplaceAll(val1,"\\","/")  }  		
	
	fmt.Println("\n"+ "parameter_path_html =" + parameter_path_html + "\n" +  "input = " + file_inputControl )
	fmt.Println("countNumLines       = " ,  countNumLines)
	//fmt.Println("maxNumLinesToWrite = " , maxNumLinesToWrite) 
	
	fmt.Println("lemmaformat = ", lemmaFormat, " sw_lemma_word =", sw_lemma_word)
	fmt.Println("\n----------------\n")	
	
	//-----------------------------------
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	//  err := lorca.New("", "", 480, 320, args...) moved out of main so that ui is available outside main()
	if err != nil {
		fmt.Println( red( "errore in lorca "), err )  //  //log.Fatal(err)
	}
	defer ui.Close()
	
	get_all_binds()  //  binds inside are executed asynchronously after calling from js function (html/js are ready) 
	
	begin_GO_HTML_Talk();  // this function is  executed firstily before html/js is ready  
	
	// the following in main() is executed at the end when the browser is close 
	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
		case <-sigc:
		case <-ui.Done():
	}
	fmt.Println("exiting") // log.Println("exiting...")
}
//-----------------------------------------
func endBegin(wh string) {
	//fmt.Println("func endBegin (", wh,")")
	if sw_stop { 
		fmt.Println("\nXXXXXXXX  error found XXXXXXXXXXXXXX\n"); 
	}	
	sw_begin_ended = true 		
}
//--------------------------------

func read_dictLang_file( path1 string, inpFile string) {
	bytesPerRow:= 10
    lineD := rowListFromFile( path1, inpFile, "scelta lingua e voce", "read_dictLang_file", bytesPerRow)  
	if sw_stop { return }
	
	lineZ := ""
	prevRunLanguage = ""	
	for z:=0; z< len(lineD); z++ { 
		lineZ = strings.TrimSpace(lineD[z]) 
		if lineZ == "" { continue }
		if lineZ[0:9] == "language=" {			
			prevRunLanguage = lineZ[9:] 			
		}
	}	
}// end of read_dictLang_file		

//-------------------------------------

func check(e error) {
    if e != nil {
        panic(e)
    }
}

//--------------------------------

