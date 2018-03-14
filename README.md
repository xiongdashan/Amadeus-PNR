Amadeus PNR 文本解析
=======================
通过PNR的规则，对文件进行Split和Regexp操作，格式化输出乘客、行程信息

输入PNR编码如下:
```html
--- TST RLR SFP ---                                                             
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
	 38 FV PAX HU/S6-7/P4
```

输出如下：
```json
{
    "code": "OUHEXE",
    "flight_section": [{
        "index": 1,
        "marketing_airline": "HU",
        "flight_number": "7955",
        "class_avail": "I",
        "departure_date": "2018-07-13",
        "arrival_date": "2018-07-13",
        "week": "5",
        "departure_city_code": "PVG",
        "departure_city": "",
        "arrival_city_code": "SEA",
        "arrival_city": "",
        "departure_time": "13:55",
        "arrival_time": "09:35",
        "big_code": "PB6814"
    }, {
        "index": 2,
        "marketing_airline": "HU",
        "flight_number": "7956",
        "class_avail": "I",
        "departure_date": "2018-07-31",
        "arrival_date": "2018-07-30",
        "week": "1",
        "departure_city_code": "SEA",
        "departure_city": "",
        "arrival_city_code": "PVG",
        "arrival_city": "",
        "departure_time": "12:00",
        "arrival_time": "15:00",
        "big_code": "PB6814"
    }],
    "passenger": [{
        "first_name": "CHEN",
        "last_name": "LIXIN",
        "type": "ADT",
        "gender": "F",
        "id_card_type": "P",
        "id_card_no": "",
        "birthday": "1938-09-23",
        "nationality": "",
        "id_issue_country": "",
        "id_expireDate": "",
        "ticket_number": "880-70000000003",
        "index": 1
    }, {
        "first_name": "CHEN",
        "last_name": "OUXING",
        "type": "ADT",
        "gender": "F",
        "id_card_type": "P",
        "id_card_no": "",
        "birthday": "1969-09-10",
        "nationality": "",
        "id_issue_country": "",
        "id_expireDate": "",
        "ticket_number": "880-0000000003",
        "index": 2
    }, {
        "first_name": "MA",
        "last_name": "GERRY YEN",
        "type": "ADT",
        "gender": "M",
        "id_card_type": "P",
        "id_card_no": "",
        "birthday": "2004-04-28",
        "nationality": "",
        "id_issue_country": "",
        "id_expireDate": "",
        "ticket_number": "880-00000000035",
        "index": 3
    }, {
        "first_name": "MA",
        "last_name": "JASON",
        "type": "CHD",
        "gender": "M",
        "id_card_type": "P",
        "id_card_no": "",
        "birthday": "2010-06-21",
        "nationality": "",
        "id_issue_country": "",
        "id_expireDate": "",
        "ticket_number": "880-70000000003",
        "index": 4
    }, {
        "first_name": "MA",
        "last_name": "TIANYI",
        "type": "ADT",
        "gender": "M",
        "id_card_type": "P",
        "id_card_no": "",
        "birthday": "1968-11-20",
        "nationality": "",
        "id_issue_country": "",
        "id_expireDate": "",
        "ticket_number": "880-00000000036",
        "index": 5
    }]
}
```

# 匹配规则
* 第一行必须以---开头
* 第二行预订信息，取最后六位数为PNR编码
* 第三行操作信息
* 乘客信息必须按名称前面的数字做好排序，否则匹配证件信息会错误
* 保持航班信息前的数字的排序
* 每行如果非数字开头，会追加到上一匹配的Item内容中
* **乘客信息中的性别、出生日期先取数字开头项的内容，如果包含DOCS项，会替换成DOCS项里的信息**