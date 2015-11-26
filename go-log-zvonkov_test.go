// go-log-zvonkov_test
package main

import (
		"testing"
	//	"flag"
	//	"os"
)

//func TestMain(t *testing.T) {
//	flag.Parse()
//	os.Exit(t.Run())
//}

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

// тест: выборка только тех даннх которые входят в массив телефонов keys
func TestFilternumtelst (t *testing.T) { 
	keys:=[]string{"15204","15999","10245"}
	numtel:=[]InputDataTel{ InputDataTel{"2015-11-26","15204",10,"+71111111"},InputDataTel{"2015-11-25","15224",7,"+778548545"},
							InputDataTel{"2015-11-26","14999",10,"+71115811"}, InputDataTel{"2015-11-02","10245",10,"+71111111"},InputDataTel{"2015-11-25","15224",7,"+778548545"}}
	result:=[]InputDataTel{ InputDataTel{"2015-11-26","15204",10,"+71111111"}, InputDataTel{"2015-11-02","10245",10,"+71111111"}}
	res:=filternumtels(numtel,keys)	
	if (len(res)!=2) {
		t.Error("Ошибка: ваш результат  ",res)
	} else{
		if (res[0]!=result[0]) || (res[1]!=result[1])  {
			t.Error("Ошибка: ваш результат  ",res)
		}
	}
}

 
