package wbfSubPack

import (  
	"fmt"
	"os"
    "bufio"		
    "io"	
)
//--------------------------------------

func test_folder_exist( myDir string) {
    _, err := os.Open( myDir )
    if err != nil {
		msg0:= `la cartella <span style="color:blue;">` + myDir + "</span> non esiste"			
		msg1:= `la cartella ` + myDir + "</span> non esiste"				
			
		msg2:= "read_dictionary_folder"
		errorMSG = `<br><br> <span style="color:red;">` + msg0 + `</span>` +  
			`<br><span style="font-size:0.7em;">(func ` + msg2 	 + ")" + `</span>` 		
		showErrMsg(errorMSG, msg1, msg2 )	
		
		sw_stop = true 
		return		
    }
} // end of test_folder_exist	
//------------------

//------------------
func getFileByteSize( path1 string,   fileName string) int {
	path2:=""
	if path1 != "" {
		path2 = path1 + string(os.PathSeparator) 
	} 	
	fileN := path2 + fileName 
	fileInfo, _ := os.Stat( fileN )  
	
	//fmt.Println("getFileByteSize fileN=", fileN, " fileInfo = ", fileInfo) 
	if fileInfo == nil { return 0 }
	return int( fileInfo.Size() )
} // end of getFileByteSize

//--------------
func myOpenRead( path1 string,   fileName string,   descr string,  func1    string) (*os.File, int) {
	path2:="";
	path10:=""
	if path1 != "" {
		path10 = " in " + cyan(path1)
		path2 = path1 + string(os.PathSeparator) 
	} 	
	fileN := path2 + fileName 
	
	fmt.Println("\n" + yellow("open file"),  green(fileName) , path10 )
	
	sizeByte:= getFileByteSize(path1,fileName)
	readFile, err := os.Open( fileN )  
    if err == nil {				
		fmt.Println( "\t", "size: ", sizeByte, " bytes" )	
		return readFile, sizeByte
	}
	msg1_Js:= `il file "` + fileN + `" (` + descr + " " + func1 + ")" + " non esiste"
		
	errorMSG = `<br><br>il file ` + 
				`<span style="font-size:0.7em;color:black;">(`	+ descr + `)</span>` +
				`<br><span style="color:blue;" >` + fileName + `</span>`	+ 				
				`<br><span style="font-size:0.7em;color:red;">`	+ "non esiste" 	+ `</span>` +				
				`<br><span style="font-size:0.7em; color:black;">nella cartella ` + path2    + `</span>` 
				
	showErrMsg2(errorMSG, msg1_Js)	
	
	return readFile, 0		
	
} // end of myOpenRead


//----------------------------------

func rowListFromFile( path1 string, fileName string, descr string, func1 string, bytesPerRow int ) []string { 
	
	file, sizeByte := myOpenRead( path1, fileName, descr, func1 )  
	if file == nil {			
		sw_stop = true 
		return nil
	} 
	numEleMax:= int( sizeByte / bytesPerRow ); 
	if numEleMax < 10 {numEleMax=10}
	
	fmt.Println("    allocate for a maximum of ", numEleMax, " rows (assumed ", bytesPerRow, " bytes per row as average)" )
	
	retRowL := make( [] string, 0, numEleMax)	
	
	r := bufio.NewReader(file)
	for {
	  line, _, err := r.ReadLine()
	  if err != nil {
		if err == io.EOF {
			break
		}
		break
	  }	 
	  retRowL = append( retRowL, string(line) ) 	
	}
	defer file.Close()
	
	fmt.Println("letto file " , fileName, "  num lines=", len(retRowL) )
	
	return retRowL 
	
} // end of rowListFromFile	

//-------------------------------------