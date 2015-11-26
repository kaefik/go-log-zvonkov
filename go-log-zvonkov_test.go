// go-log-zvonkov_test
package main

import (
		"testing"
)

// тест разбивают дату YYYY-MM-DD на 2 части: (YYYY-MM,DD)
func TestParse_date(t *testing.T) { 
	var s string
	// проверка на правильные входные данные
	s="2015-12-29"
	syear:="2015-12"
	sday:="29"
	syearmonth, sdays:=parse_date(s)
	if (syearmonth!=syear) || (sdays!=sday){
		t.Error("неправильный разбор даты: ",syearmonth)
	}
	//проверка на неправильные данные
	sno:="29-12-2015"
	syearmonth, sday=parse_date(sno)
	if (syearmonth!="") || (sday!=""){
		t.Error("не прошел тест некорректных входных данных: ",sno)
	}		
	sno=""
	syearmonth, sday=parse_date(sno)
	if (syearmonth!="") || (sday!=""){
		t.Error("не прошел тест некорректных пустых входных данных: ",sno)
	}	
}

func TestSec_to_hour(t *testing.T) { 
	var ssec int = 3600
	res:=sec_to_hour(ssec)
	if res!=1{
		t.Error("Ошибка, должно быть 1 час : ",res)
	}
}

