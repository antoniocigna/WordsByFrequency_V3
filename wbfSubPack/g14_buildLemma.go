package wbfSubPack

import (  
	"fmt"
	"strconv"
    //"strings"
	//"sort"
)
//--------------------------------------------------------------------------

func g14_appendOneLemma( xLemma string, fromIx int, toIx int, numLemmaOrig int, numLemmaAdded int	) (int, int) {
	
	
	
	var leV lemmaStruct; 
	iixLem:=0	
	leV.leLemma    = xLemma
	leV.leNumWords = 0 
	leV.leFromIxLW = fromIx 
	leV.leToIxLW   = toIx  
	leV.leTran     = ""
	//leV.leLevel    = ""   
	leV.lePara     = ""   
	leV.leExample  = "" 
	//-----------			
	/**
	ixTra := lookForAllTran( xLemma,0 ) 
	if ixTra >= 0 { 
		leV.leTran = dictLemmaTran[ixTra].dL_tran 
	} 		
	***/
	numLemmaOrig++
	lemmaSlice = append(lemmaSlice, leV ) 
	iixLem = len(lemmaSlice) -1 
	for h:=fromIx; h<= toIx; h++ {
		/**
		if xLemma == "pünktlich" {
			fmt.Println( red("lemma "), xLemma, " wordLemmaPair[", h , "]=" , 	wordLemmaPair[h]); 	
		}
		**/
		wordLemmaPair[h].lIxLemma = iixLem    
	}	
	return numLemmaOrig, numLemmaAdded			
} // end of g14_appendOneLemma 

//---------------------

func g14_updateLemmaWithTranPara() {
	
	/*
	aggiorna la prima lista (lemmaSlice) con i dati delle altre due 
	tutte le liste devono essere già in sequenza di codice (alfanumerico)
	---
	nella lemmaSlice ogni elemento ha struttura: [codice sequenza, dati vari, [ dati lista2], [ dati lista3]  ]
	le altre due liste ogni elemento ha la struttura: [codice sequenza, dati vari]
	*/
	//---------------------------------------
	var maxNumErr = 10;	
	
	numOutSeq := 0;
	len1 := len(lemmaSlice) 
	len2 := len(dictLemmaTran   )
	len3 := len(lemma_para_list   )
	loopMax := len1 + len2 + len3;
	
	j1:=0; j2:=0; j3:=0; 
	min1:= ""; 
	type1:=0; 	
	//----------------
	t0 :=-1;
	preMin := ""
	index_ix1 := -1
	var cod1, cod2, cod3 string;
	var pcod1, pcod2, pcod3 string;
	numIgn2 :=0; numIgn3 :=0;
	//-----------------------
	
	for t0=0; t0 < loopMax; t0++ {
		if (j1 < len1) {cod1 = "1" + lemmaSlice[j1].leLemma } else {cod1 = "9" }
		if (j2 < len2) {cod2 = "1" + dictLemmaTran[j2].dL_lemma   } else {cod2 = "9" }
		if (j3 < len3) {cod3 = "1" + lemma_para_list[j3].p_lemma    } else {cod3 = "9" }
		if (cod1 <= cod2) {
			min1 = cod1; type1=1; 
			if (cod1 < pcod1) {	_ = g14_outSeqErr("g14_updateLemmaWithTranPara 1 ", type1, pcod1, cod1, numOutSeq, maxNumErr); return ; 	}
			pcod1 = cod1
		}  else {
			min1 = cod2; type1=2;
			if (cod2 < pcod2) {	_ = g14_outSeqErr("g14_updateLemmaWithTranPara 2 ", type1, pcod2, cod2, numOutSeq, maxNumErr); return ; 	}
			pcod2 = cod2
		}
		if (cod3 < min1) { 
			min1 = cod3; type1=3; 
			if (cod3 < pcod3) {	_ = g14_outSeqErr("g14_updateLemmaWithTranPara 3 ", type1, pcod3, cod3, numOutSeq, maxNumErr); return ; 	}
			pcod3 = cod3
		} 
		
		if (min1[0:1] == "9") {break}
		
		if (min1 > preMin) {
			if (preMin != "") { seqCodeChange(index_ix1, lemmaSlice)}			
			index_ix1 = -1
			preMin    = min1 
		} else {
			if (min1 < preMin) { 
				numOutSeq = g14_outSeqErr( "g14_updateLemmaWithTranPara 4 ", type1, preMin, min1, numOutSeq, maxNumErr)
				if numOutSeq < 0 { return }
				continue; 
				/**
				fmt.Println("XXXXXXXX  ERRORE di Sequenza XXXXXXXX " + 
					" codice precedente=" + preMin[1:] + 
					" codice attuale=" + min1[1:] + " tipo=" + strconv.Itoa(type1) ) 
				numOutSeq++	
				if (numOutSeq > maxNumErr) {return}	
				**/				
			}
		}
		/**
			leTran  = dL_tran     
			lePara  = p_para  
		**/
		if (type1 == 1) {
			index_ix1 = j1  
			j1++;
		} else {
			if (type1 == 2) {
				if (index_ix1 < 0) { 
					//fmt.Println("               tipo2 senza master, ignorato " , dictLemmaTran[j2] )
					numIgn2++
				} else {
					lemmaSlice[index_ix1].leTran = dictLemmaTran[j2].dL_tran  
				}
				j2++;
			} else {
				if (index_ix1 < 0) { 
					fmt.Println("               tipo3 senza master, ignorato " , lemma_para_list[j3] )
					numIgn3++
				} else {
					lemmaSlice[index_ix1].lePara    = lemma_para_list[j3].p_para  
					lemmaSlice[index_ix1].leExample = lemma_para_list[j3].p_example  
				}
				j3++;				
			}
		}		
		
	} // end for t0
	//-----------------
	
	if (preMin != "") {seqCodeChange(index_ix1, lemmaSlice)}
	
	fmt.Println("  in lemmaSlice, trovati ", numIgn2, " tipo2 (traduz.) senza lemma, " , numIgn3, " tipo3 (paradigma) senza lemma, ",   
		len(lemmaSlice), " righe in lemmaSlice, \n\t",  len(dictLemmaTran), " tipo2 letti, di cui ", (len(dictLemmaTran)-numIgn2), " usati",
		",\n\t",  len(lemma_para_list), " tipo3 letti, di cui ", (len(lemma_para_list)-numIgn3), " usati " )


} // end of g14_updateLemmaWithTranPara
//-------------------------------------------

func seqCodeChange( index_ix1 int, lemmaSlice []lemmaStruct) {		
	/**
	if (index_ix1 < 0) {  
		//fmt.Println(preMin, " SENZA MASTER IGNORA") ;
	} else {
		fmt.Println( lemmaSlice[index_ix1] )
	}
	***/
} // end of seqCodeChange
//-----------------------------
func g14_outSeqErr(descr string,  type1 int, preCod string, cod string, numOutSeq int, maxNumErr int) int {		
	
	fmt.Println(red("XXXXXXXX  ERRORE di Sequenza XXXXXXXX ") + descr,   
					" codice precedente=" + preCod[1:] + 
					" codice attuale=" + cod[1:] + " tipo=" + strconv.Itoa(type1) ) 
				numOutSeq++	
				if (numOutSeq > maxNumErr) { return -1 }	
				return numOutSeq; 
} // end of seqCodeChange

//-------------------------------------------
