package wbfSubPack

	import (
		"fmt"
		"os"
		"strings"
		"strconv"
	)
//----------------------------------
var sw_stop bool = false
var errorMSG = ""; 
//------------------------------
func begin() { 	
	//swProva := true
	fmt.Println("func begin"); 
	
	g1_read_all_files() 
	
	if sw_stop { endBegin("1"); return }
	/**
	if swProva {return} 
	
	build_and_elab_word_list()
	if sw_stop { endBegin("2"); return }
	**/
	
	stat_useWord();	
	if sw_stop { endBegin("3"); return }	
	
	/**
	//if sw_rewrite_wordLemma_dict { rewrite_word_lemma_dictionary() }
	
	if sw_stop { endBegin("4"); return }	
	**/
	//read_wordsToLearn()	
	
	endBegin("6")
	
	//writeUnaTantumNuovoFile()
	
	//-------------------------------------------
	
	numberOfRows = len(inputTextRowSlice)
	
	g32_buildStatistics()
	
	mainNum := strconv.Itoa(numberOfUniqueWords) +";" + strconv.Itoa(numberOfWords) + ";" + strconv.Itoa(numberOfRows) +	
		   ";))"	
		   //";" + "level " + msgLevelStat + "))" 		

	go_exec_js_function("js_go_showReadFile", mainNum + showReadFile);  
			
	if sw_stop { 
				fmt.Println("UI is ready ( run stopped because of some error)")
	} else {
		//go_exec_js_function("js_go_ready", prevRunLanguage + ":mainpage_value=" + last_mainpage_valueString)  // +"<file>" + prevRunListFile); 
		
		go_exec_js_function("js_go_ready", prevRunLanguage )
		//log.Println("UI is ready")
		fmt.Println("UI is ready")
	}	
	fmt.Println("\nEND of begin \n") 
	fmt.Println(cyan("\nREADY"), "\n") 

	
}// end of begin	

//----------------------------------
func g01_build_word_db() {

	fmt.Println("\n", cyan("BUILD WORD LIST") )
	
	wordSliceAlpha    = nil 
	//wordSliceFreq     = nil
	uniqueWordByFreq  = nil
	uniqueWordByAlpha = nil
	
	g11_buildWordList() 	   	
	
	//g17_build_lemma_word_ix()
	
	
	//elabWordList() 
	
	
} // end of build_and_elab_word_list()
//------------------------------
func build_and_elab_word_list() {
	fmt.Println( red("build_and_elab_word_list  vuota x prova"))
} 
//--------------------------------

func getPgmArgs( key0, key1 , key2 , key3, key4 string) (string, string, bool, int, string) {  
	
	//  getPgmArgs("-html", "-input" , "-countNumLines" ,  "-maxNumLinesToWrite")	

	args1    :=  os.Args[1:]		
	
	
	var val0, val1, val2, val3, val4 string
	for a:=0; a < (len(args1)-1); a++ {
		switch args1[a] {
			case key0 :   val0 = args1[a+1]
			case key1 :   val1 = args1[a+1]
			case key2 :   val2 = args1[a+1]
			case key3 :   val3 = args1[a+1]
			case key4 :   val4 = args1[a+1]
		}
	}  
	var isCount = false;
	if strings.TrimSpace(val2) == "true" {
		isCount = true
	}
	var num=0; 
	num, err := strconv.Atoi( strings.TrimSpace(val3) )
	if err != nil {
		num=0
	}

	//fmt.Println("args=", args1,  " val0=", val0, " val1=", val1, " val2=", val2 , " val3=", val3, " num=", num, " val4=", val4)   
	
	return val0, val1, isCount, num, val4
	
} // end of getPgmArgs
//-------------------------------
func g1_read_all_files() { 
	//swProva:= true
	fmt.Println( "func ", green("read_all_files") )
	
	g28_read_control_file()
	if sw_stop { endBegin("1"); return }
	test_all_folder()
	if sw_stop {return}
	
	g31_read_languageFile(  FOLDER_INPUT, FILE_inputLanguage)
	if sw_stop { return }
	
	fmt.Println( green("read_dictRow_Orig_and_Tran_file --> build inputTextRowSlice") )
	g25_1read_dictRow_Orig_and_Tran_file( FOLDER_IO_lastTRAN,  FILE_last_updated_dict_rows)	
	if sw_stop { return }	
	
	fmt.Println( green("read_lemma_file --> build listAllLemmaFromFile, listAllLemmaFromFile") )
	g30_read_wordLemma_file( FOLDER_I_lemma, FILE_inputWordLemma, FILE_inputWordLemmaPlus )
	if sw_stop { return }
	
	g25_2read_ParadigmaFile( FOLDER_I_paradigma, FILE_inpParadigma ) ;		
	if sw_stop { return }
	
	g25_3read_dictLemmaTran_file( FOLDER_IO_lastTRAN, FILE_last_updated_dict_words ) 
	if sw_stop { return }	
	
	fmt.Println( green("build word_db") )
	g01_build_word_db()
	fmt.Println( red("finito build_word_db"), "\n\n\n")
	
	g34_load_direct_and_inverse_lemma()
	
	/**
	RIMETTI
	if swProva { return }
	
	elabWordList_11()
	
	if swProva { return }
	
	fmt.Println( green("read_all_files read_lemma_file") )
	
	
	//g10_build_extrLemmaPair()
	
	if sw_stop { return } 
	if sw_stop == false { return } 
	
	
	read_dictLang_file( FOLDER_INPUT_OUTPUT, FILE_language );	 
	if sw_stop { return }
	
	read_dictLemmaTran_file( FOLDER_IO_lastTRAN, FILE_last_updated_dict_words ) 
	if sw_stop { return }	
	
	//read_wordToLearnFile() 
	if sw_stop { return }	
	**/
	read_lastValueSets2()
	
} // end of read_all_files
//-----------------------------------

func test_all_folder() {	
	//test_folder_exist( FOLDER_INPUT         ); if sw_stop { return } 	
	test_folder_exist( FOLDER_OUTPUT        ); if sw_stop { return } 	
	test_folder_exist( FOLDER_INPUT_OUTPUT  ); if sw_stop { return } 
	/**
	test_folder_exist( FOLDER_I_lemma )     ;  if sw_stop { return } 	
	test_folder_exist( FOLDER_I_paradigma)  ;  if sw_stop { return } 	
	**/
	test_folder_exist( FOLDER_IO_lastTRAN ) ;  if sw_stop { return } 	
	
	test_folder_exist( FOLDER_O_ARCHIVE        ) ;  if sw_stop { return } 	
	test_folder_exist( FOLDER_O_arc_TRAN_rows  ) ;  if sw_stop { return } 		
	test_folder_exist( FOLDER_O_arc_TRAN_words ) ;  if sw_stop { return } 	
	test_folder_exist( FOLDER_O_arc_TO_learn   ) ;  if sw_stop { return } 	
} // end of test_all_folders

 