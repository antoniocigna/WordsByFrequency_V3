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
function onclick_print_selected_row(this1, swTran)	{
	/*
	
	<tr id="idtr_1_m2" style="display: none;">
	         <td></td>
            <td></td>
            <td></td>
            <td></td>
            <td class="borderVert"></td>
            <td></td>
            <td></td>
            <td></td>
            <td></td>
			<td></td>
	</tr>
	<tr id="idtr_1_m1" style="display: table-row; border-bottom: 1px solid black;">
			<td></td>
            <td></td>
            <td></td>
            <td></td>
            <td><td>
            <td class="borderVert" style="color:black;font-size:0.5em;text-align:center;vertical-align:middle;">
				doppio click sulla frase per inserire la traduzione o variare quella esistente 
				<div>
					<button onclick="onclick_print_selected_row(this)">Stampa le righe selezionate</button>  
				</div>
			</td>
			...
		</tr>	
		
		<tr id="idtr_1" > 
			 <td class="arrow12"><button class="buttonFromToIx" id="b1_1" onclick="onclick_tts_arrowFromIx(this, 1, 1)" style="background-color: green;">
               <span style="font-size:1em;height:1.4em;">→</span></button>
            </td>
		    ...
		</tr>	
		 <tr id="idtr_2" class="seleRow1" style="width:100%;">
            <td class="arrow12"><button class="buttonFromToIx" id="b1_2" onclick="onclick_tts_arrowFromIx(this, 2, 1)">
               <span style="font-size:1em;height:1.4em;">→</span></button>
            </td>
			...
		</tr> 	
		
		ultimo
		
        <tr id="idtr_89" class="seleRow2" style="width:100%;">
            <td class="arrow12"><button class="buttonFromToIx" id="b1_89" onclick="onclick_tts_arrowFromIx(this, 89, 1)">
               <span style="font-size:1em;height:1.4em;">→</span></button>
            </td>
            <td class="arrow12"><button class="buttonFromToIx" id="b2_89" onclick="onclick_tts_arrowToIx(  this, 89, 1)" style="background-color: red;">
               <span style="font-size:1em;height:1.4em;">←</span></button>
            </td>
	
</tbody>
                            </table>
                            <!-- end of id_tabSub --> 
							
                        </div>
						
		
	*/
	
	var fromColor = "green"
	var toColor = "red"
	var eleTD1   = this1.parentElement.parentElement
	var eleTR1   = eleTD1.parentElement 
	var eleTBODY = eleTR1.parentElement
	var numRow   = eleTBODY.children.length	
	var eleTrX, eleTD_from, eleTD_to; 
	var eleTD_row, eleTD_rowOrig, eleTD_rowTran;
	
	var ixFrom=-1, ixTo = -1 
	
	var outStr = "<html>\n  <body>\n"   
	outStr += "  <style> \n"  
	outStr += "     td { font-size:1.2em;} \n"        
	outStr += "     .c_wordTarg { font-weight:bold;} \n " 
	outStr += "     .suboLine { font-weight:bold;}   \n " 	
	outStr += "     .tranLine { font-size:0.9em; }   \n " 
	outStr += "  </style> \n" 
	outStr += "  <table>\n" ; 
	
	for (var j=eleTR1.rowIndex; j < numRow; j++) {
		eleTrX = eleTBODY.children[j]
		if (eleTrX.id.indexOf("_m") > 0 ) { continue} 
		eleTD_from = eleTrX.children[0]	
		eleTD_to   = eleTrX.children[1]			
		if (eleTD_from.innerHTML.indexOf( fromColor ) > 0 ) { ixFrom = j} 
		
		eleTD_row = eleTrX.children[5]
		
		eleTD_rowOrig = eleTD_row.children[0].children[0]
		eleTD_rowTran = eleTD_row.children[0].children[1]
		if (swTran) {
			outStr += '      <tr class="suboLine"><td>' + eleTD_rowOrig.innerHTML  + "</b></td></tr> \n"; 
			outStr += '      <tr class="tranLine"><td>' + eleTD_rowTran.innerHTML.replace("display:none;","")  + "</td></tr> \n"; 
		} else {
			outStr += '      <tr><td>' + eleTD_rowOrig.innerHTML  + "</td></tr> \n"; 
		}
		if (eleTD_to.innerHTML.indexOf( toColor ) > 0 ) { ixTo = j; break} 
	} 
	outStr += "    </table>\n  </body>\n</html> \n"	
	//console.log( outStr) 
	var myWindow = window.open("", "", "width=500,height=800");
	myWindow.document.write( outStr ) 
	
	
} // end of onclick_print_selected_row