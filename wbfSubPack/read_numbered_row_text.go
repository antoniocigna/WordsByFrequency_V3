package wbfSubPack

import (  
	"fmt"
    "strings"
	"strconv"
    //"sort"
)
//------------------------------------------------

func read_dictRow_Orig_and_Tran_file( path1 string, inpRowFile string) {
	bytesPerRow:= 10
    lineD := rowListFromFile( path1, inpRowFile, "righe orig/tran", "read_dictRow_Orig_and_Tran_file", bytesPerRow)  
	if sw_stop { return }
	
	lineZ := ""
	/*		
	3|O|Ein Lebewesen tauscht aber auch Energie und Dinge, Stoffe mit seiner Umwelt aus.
	3|T|Ma un essere vivente scambia anche energia, cose e sostanze con il suo ambiente.
	*/
	
	
	prevRunLanguage = ""
	//prevRunListFile = ""
	//var rowDict rDictStruct
	var rS1 rowStruct
	var pre_id_key=""
	
	showReadFile = ""
	numLines:=0
	
	inputTextRowSlice = nil 
	isUsedArray       = nil 
	
	//rowLineIxList    = nil
	
	//var gRix rowIxStruct
	var indice = 0
	numAll :=0; numO :=0; numT:=0 ; numT_err:=0; num_oth :=0
	//--
	pre_id_group     := ""
	pre_id_keyFirst  := ""
	pre_id_keyLast   := ""
	pre_id_row       := ""
	id_group         := ""
	gruppi_option    := ""	
	
	num_O_ix   := 0 
	firstIxRow :=0
	lastIxRow  :=0
	//----------

	group_zero := "."; 
	
	lineZero := []string{ group_zero + "_0|O|"}  // empty row
	
	lineD = append( lineZero, lineD...)  // insertion of an empty row to avoid using index 0  
	ngr:=0
	
	var rG rowGroupStruct ; 
	
	//----
	for z:=0; z< len(lineD); z++ { 
		lineZ = strings.TrimSpace(lineD[z]) + "||||"	
		field  := strings.Split( lineZ, "|" )
		id_key := strings.TrimSpace( field[0] )
		ty     := strings.TrimSpace( field[1] ) 		
		row    := strings.TrimSpace( field[2] )  	 
		
		
		if ((row == "") || (id_key == "") || (ty=="")) { continue }
		
		//fmt.Println("carica riga z=", z, " lineZ=", lineZ, "\n\t fields: 0=", field[0], "  1=", field[1], " 2=", field[2], "  3=", field[3]   )
		
		id_group = get_rowid2(id_key)
		
		//fmt.Println("\t id_key=", id_key, " id_group=", id_group)
		
		if id_group != pre_id_group {
			if pre_id_group != "" {
				// elabora fine gruppo precedente {
				
				//fmt.Println("id_key=", pre_id_group, " from=", pre_id_keyFirst, " to=", pre_id_keyLast, " first_ixRow=", firstIxRow, " lastIxRow=", lastIxRow)  
				
				if pre_id_group != group_zero { 
					gruppi_option += "<option>" +  fmt.Sprintf("%s %s", pre_id_group,  pre_id_row)  + "</option>\n"
					rG.rG_ixSelGrOption   = ngr    // group number = index of group selection 
					rG.rG_group           = pre_id_group
					rG.rG_firstIxRowOfGr  = firstIxRow 
					rG.rG_lastIxRowOfGr   = lastIxRow						
					lista_gruppiSelectRow = append( lista_gruppiSelectRow, rG )    
					ngr++
					
					fmt.Println("Group=", pre_id_group, " key: from=", pre_id_keyFirst, " to=", pre_id_keyLast, " ixRow: first_ixRow=", firstIxRow, " lastIxRow=", lastIxRow)  
					
				}			
			}
			// elabora inizio nuovo gruppo 
			pre_id_group     = id_group 
			pre_id_keyFirst  = id_key
			pre_id_row       = row
			num_O_ix     = 0  
			if ty == "O" {
				num_O_ix = 0  
			} else {
				fmt.Println( red("errore "), " in read_dictRow_Orig_and_Tran_file" , " id_key=", pre_id_group, " from=", pre_id_keyFirst, " to=", pre_id_keyLast , 
					red("row type " + ty + " not equal O"), "\n\t", lineZ) 
				continue	
			}
			//newNum=0  ; // ?anto rinumera  
			pre_id_key  = id_key
			
			//fmt.Println("\nNUOVO Group=", id_group)
			
		}
		//--
		
		pre_id_keyLast   = id_key
		
		switch ty {
			case "O" :				
				pre_id_key  = id_key			
				numO++				
				num_O_ix++  
				rS1.rIdRow 		 = id_key
				rS1.rRow1  		 = row
				rS1.rTran1		 = ""
				rS1.rixGroup     = ngr      // indice del gruppo 
				rS1.rixBaseGroup = num_O_ix // posizione del row nel gruppo ( si inzia dal num.1 )  
		    case "T" : 
				if id_key != pre_id_key {  // error 
					numT_err++
					continue
				}		
				inputTextRowSlice[ numLines-1].rTran1  = row
				numT++
				continue 
			default  :  
				num_oth++
			    continue
		}
	
		numAll++
		numLines++
		inputTextRowSlice = append(inputTextRowSlice, rS1);	
		isUsedArray       = append(isUsedArray, false)  
		
		indice   = len(inputTextRowSlice)-1
		
		if num_O_ix==1 {
			firstIxRow = indice
		}  
		lastIxRow = indice
		
	} // end of for z 
	//-------------

	if pre_id_group != group_zero { 
		gruppi_option += "<option>" +  fmt.Sprintf("%s %s", pre_id_group,  pre_id_row)  + "</option>\n"
		rG.rG_ixSelGrOption   = ngr    // group number = index of group selection 
		rG.rG_group           = pre_id_group
		rG.rG_firstIxRowOfGr  = firstIxRow 
		rG.rG_lastIxRowOfGr   = lastIxRow						
		lista_gruppiSelectRow = append( lista_gruppiSelectRow, rG )    
		ngr++
		fmt.Println("Group=", pre_id_group, " key: from=", pre_id_keyFirst, " to=", pre_id_keyLast, " ixRow: first_ixRow=", firstIxRow, " lastIxRow=", lastIxRow)  		
	}
	//-------------
	
	go_exec_js_function( "js_go_build_rowGruppi", gruppi_option); 	
	
	//-------
	//fmt.Println("read_dictRow_Orig_and_Tran_file "  , " numAll=", numAll, " numO=", numO, " numT=", numT, " numT_err=", numT_err, " num_oth=", num_oth)
	/**
	for z:=0; z< len(inputTextRowSlice); z++ {
		fmt.Println("caricata inputTextRowSlice[",z,"]=", inputTextRowSlice[z] )
	}
	**/
	//-------------------
	
	showReadFile = showReadFile + strconv.Itoa(numLines) + "<file>" + inpRowFile + ";" ; 
	
	fmt.Println( "read ", len(lineD) , " lines of file ", inpRowFile, " which contains ",  numLines, " text rows" );  	
	
	//testRowGroup()  // TEST
	//-------------------
	
	
} // end of  read_dictRow_file


//--------------------------------

func get_rowid2(id string) string {
	// es.  10_5
	var id_pref string
	k1:= strings.Index(id,"_")
	if k1 >= 0 { 
		id_pref   =  strings.TrimSpace(  id[:k1] ) 
	} else {
		id_pref   = id  
	} 	
	return id_pref
	
} // end of get_rowid2	
//--------------------------------