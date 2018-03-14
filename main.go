package main

import (
	"Amadeus-PNR/pnrorder"
	"encoding/json"
	"fmt"
)

func main() {
	pnr := `--- TST RLR SFP ---                                                             
	RP/ONOOOOOOO /ONOOOOOOO            B0/GS   1MAR18/1224Z   OUHEXE                 
	ONOOOOOOO/2045LW/1MAR18                                                         
	  1.CHEN/LIXIN   2.CHEN/OUXING   3.MA/GERRY YEN                            
	  4.MA/JASON MSTR(CHD/21JUN10)   5.MA/TIANYI                            
	  6  HU7955 I 13JUL 5 PVGSEA HK5  1355 0935  13JUL  E  CA/PB6814                
	  7  HU7956 I 30JUL 1 SEAPVG HK5  1200 1500  31JUL  E  CA/PB6814                
	  8 APE OP@XXXXX.CN                                                 
	  9 TK PAX OK01MAR/ONT1S212G//ETHU/S6-7/P1-3,5                                  
	 10 TK OK01MAR/ONT1S212G//ETHU                                                  
	 11 SSR CHLD HU HK1 21JUN10/P4                                                  
	 12 SSR ADPI 1A KK1 HU7955 REQ SEC FLT PSGR DATA 72 HBD FOR ALL                 
		   PSGRS                                                                    
	 13 SSR ADPI 1A HK1 HU7956 REQ SEC FLT PSGR DATA 72 HBD FOR ALL                 
		   PSGRS                                                                    
	 14 SSR OTHS 1A PLS REMOVING HU HX SEGMENTS 24 HOURS BEFORE                     
		   DEPARTURETO AVOID ADM PENALTY                                            
	 15 SSR ADTK 1A BY ONT11MAR18/0300 OR CXL HU7955 I13JUL                         
	 16 SSR DOCS HU HK1 ////21JUN10/M//MA/JASON/P4                           
	 17 SSR DOCS HU HK1 ////10SEP69/F//15APR19/CHEN/OUXING/P2                      
	 18 SSR DOCS HU HK1 ////20NOV68/M//16MAY26/MA/TINYI/P5                         
	 19 SSR DOCS HU HK1 ////23SEP38/F//11DEC18/CHEN/LIXIN/P1                                                                              
	 20 SSR DOCS HU HK1 ////28APR04/M//28DEC22/MA/GERRYYEN/P3                    
	 21 OSI HU CTCT 15232971626                                                     
	 22 OSI HU CTCM 18686122446/P1/2/3/4/5                                          
	 23 RC ONT1S212G-W/CHG1500CNY CXL2000CNY 2PC                                    
	 24 RM NOTIFY PASSENGER PRIOR TO TICKET PURCHASE & CHECK-IN:                    
		   FEDERAL LAWS FORBID THE CARRIAGE OF HAZARDOUS MATERIALS -                
		   GGAMAUSHAZ/S6-7                                                          
	 25 FA PAX 880-70000000003/ETHU/USD0000.87/01MAR18/ONT1S212G/0550                
		   2501/S6-7/P1                                                             
	 26 FA PAX 880-74000000003/ETHU/USD0000.87/01MAR18/ONT1S212G/0550                                                                         
		   2501/S6-7/P1                                                             
	 26 FA PAX 880-0000000003/ETHU/USD0000.87/01MAR18/ONT1S212G/0550                
		   2501/S6-7/P2                                                             
	 27 FA PAX 880-00000000035/ETHU/USD0000.87/01MAR18/ONT1S212G/0550                
		   2501/S6-7/P3                                                             
	 28 FA PAX 880-00000000036/ETHU/USD9999.87/01MAR18/ONT1S212G/0550                
		   2501/S6-7/P5                                                             
	 29 FA PAX 880-70000000003/ETHU/USD99999.61/01MAR18/ONT1S212G/0550                
		   2501/S6-7/P4                                                             
	 30 FB PAX 0000000000 TTP/RT OK ETICKET - USD16485.09/S6-7/P1-3                 
		   ,5                                                                       
	 31 FB PAX 0000000001 TTP/RT OK ETICKET - USD00900.09/S6-7/P4                   
	 32 FE PAX Q/NON-END/PENALTY APPLY -BG:HU/S6-7/P1-3,5                           
	 33 FE PAX Q/NON-END/PENALTY APPLY -BG:HU/S6-7/P4                               
	 34 FM *M*3                                                                     
	 35 FP CASH                                                                     
	 36 FT *CCA1AM4095                                                              
	 37 FV PAX HU/S6-7/P1-3,5                                                       
	 38 FV PAX HU/S6-7/P4  `

	order := pnrorder.NewPNR(pnr)
	order.Analysis()
	data := order.Ouput()
	buffer, _ := json.Marshal(data)
	fmt.Println(string(buffer))
}
