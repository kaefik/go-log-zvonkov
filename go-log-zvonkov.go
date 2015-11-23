package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/tealeg/xlsx"
	"github.com/headzoo/surf"
)

var (
	d1 string // начальная дата выгрузки
	d2 string // конечная дата выгрузки
)

func parse_args() bool {
	flag.StringVar(&d1, "d1", "", "Начальная дата выгрузки лога звонков: YYYY-MM-DD")
	flag.StringVar(&d2, "d2", "", "Конечная дата выгрузки лога звонков: YYYY-MM-DD")
	flag.Parse()
	if d1 == "" {
		fmt.Println("Не задан параметр -d1 . Будет использована текущая системная дата", d1)
	}
	if d2 == "" {
		fmt.Println("Не задан параметр -d2 . Будет использована текущая системная дата", d2)
	}
	return true
}

func parse_date(s string) (string, string) { // разбивают дату YYYY-MM-DD на 2 части: (YYYY-MM,DD)
	s1 := s[0:7]
	s2 := s[8:10]
	return s1, s2
}

func sec_to_hour(ss int) int {
	return ss / 3600
}

func sec_to_min(ss int) int {
	return ss / 60
}

func sec_to_s(s int) string {
	hh := sec_to_hour(s)
	mm := sec_to_min(s - hh*3600)
	ss := s - mm*60 - hh*3600
	return strconv.Itoa(hh) + ":" + strconv.Itoa(mm) + ":" + strconv.Itoa(ss)
}

// чтение файла с именем namefи возвращение содержимое файла, иначе текст ошибки
func readfilecsv(namef string) string {
	file, err := os.Open(namef)
	if err != nil {
		return "handle the error here"
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return "error here"
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return "error here"
	}
	return string(bs)
}

func readcfgs(namef string) map[string]DataTelMans {
	str := readfilecsv(namef)
	vv := strings.Split(str, "\n")
	//xx:=make(map[string]string)
	//xx["key"] = "10"
	s_inputdata := make(map[string]DataTelMans)
	for i := 0; i < len(vv)-1; i++ {
		vv1 := strings.Split(vv[i], ";")
		s_inputdata[vv1[0]] = DataTelMans{vv1[2], vv1[1], 0, 0, 0, 0,0}

	}
	return s_inputdata
}

func savetoxlsx(namef string, datas map[string]DataTelMans) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("лог звонков")
	if err != nil {
		fmt.Println(err.Error())
	}

	//заголовок таблицы
	row = sheet.AddRow() // добавить строку
	titletab := []string{"ФИО РГ",
		"номер телефона",
		"ФИО менеджера",
		"всего продолжит-ть",
		"всего кол-во звонков",
		"кол-во уникальных телефонов",
		"кол-во результ. звонков",
		"продолжительность уникальных"}
	for i := 0; i < len(titletab); i++ {
		cell = row.AddCell() // добавить ячейку в текущей строке
		cell.Value = titletab[i]
	}

	for key, _ := range datas {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = datas[key].fio_rg
		cell = row.AddCell()
		cell.Value = key
		cell = row.AddCell()
		cell.Value = datas[key].fio_man
		cell = row.AddCell()
		cell.Value = sec_to_s(datas[key].totalsec)
		cell = row.AddCell()
		cell.Value = strconv.Itoa(datas[key].totalzv)
		cell = row.AddCell()
		cell.Value = strconv.Itoa(datas[key].kolunik)
		cell = row.AddCell()
		cell.Value = strconv.Itoa(datas[key].kolresult)
		cell = row.AddCell()
		cell.Value = sec_to_s(datas[key].secresult)

	}

	err = file.Save(namef)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func savetohtml(namef string, datas map[string]DataTelMans) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("лог звонков")
	if err != nil {
		fmt.Println(err.Error())
	}

	//заголовок таблицы
	row = sheet.AddRow() // добавить строку
	titletab := []string{"ФИО РГ",
		"номер телефона",
		"ФИО менеджера",
		"всего продолжит-ть",
		"кол-во уникальных телефонов",
		"кол-во результ. звонков",
		"продолжительность уникальных"}
	for i := 0; i < len(titletab); i++ {
		cell = row.AddCell() // добавить ячейку в текущей строке
		cell.Value = titletab[i]
	}

	for key, _ := range datas {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = datas[key].fio_rg
		cell = row.AddCell()
		cell.Value = key
		cell = row.AddCell()
		cell.Value = datas[key].fio_man
		cell = row.AddCell()
		cell.Value = sec_to_s(datas[key].totalsec)
		cell = row.AddCell()
		cell.Value = strconv.Itoa(datas[key].kolunik)
		cell = row.AddCell()
		cell.Value = strconv.Itoa(datas[key].kolresult)
		cell = row.AddCell()
		cell.Value = sec_to_s(datas[key].secresult)

	}

	err = file.Save(namef)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// - функции генерации html page
//-- генерация ячейки таблицы в html
func gentablecell(str string) string {
	return "<TD>" + str + "</TD>" + "\n"
}

//-- генерация строки таблицы в html
func gentablestroka(str []string) string {
	res0 := ""
	for i := 0; i < len(str); i++ {
		res0 += gentablecell(str[i])
	}
	return "<TR>" + "\n" + res0 + "</TR>" + "\n"
}

func makestrfromarray(dd DataTelMans) []string {
	res := []string{dd.fio_rg, dd.fio_man}
	return res
}

//-- генерация таблицы в html: первый параметр это заголовок таблицы, второй параметр [[],[],...] - строки таблицы
func genhtmltable(datas map[string]DataTelMans, zagol string) string {
	res := ""
	//res = map gentablestroka str

	titletab := []string{"ФИО РГ",
		"номер телефона",
		"ФИО менеджера",
		"всего продолжит-ть",
		"всего кол-во звонков",
		"кол-во уникальных телефонов",
		"кол-во результ. звонков",
		"продолжительность уникальных"}
	tabletitle := gentablestroka(titletab)

	tabledata := ""
	for key, _ := range datas {
		str := []string{
			datas[key].fio_rg,
			key,
			datas[key].fio_man,
			sec_to_s(datas[key].totalsec),
			strconv.Itoa(datas[key].totalzv),
			strconv.Itoa(datas[key].kolunik),
			strconv.Itoa(datas[key].kolresult),
			sec_to_s(datas[key].secresult)}

		tabledata += gentablestroka(str)
	}

	zagolovok := "<CAPTION>" + zagol + "</CAPTION>\n"
	tablehtml := zagolovok + tabletitle + tabledata
	return "<TABLE>" + "\n" + "<TABLE BORDER>\n" + tablehtml + res + "</TABLE>"
}

func genhtmlpage(datas map[string]DataTelMans, zagol string) string {
	begstr := "<html>\n <head>\n <meta charset='utf-8'>\n <title>" + zagol + "</title>\n </head>\n <body>\n"
	bodystr := genhtmltable(datas, zagol)
	endstr := "</body>\n" + "</html>"
	return begstr + bodystr + endstr
}

func savestrtofile(namef string, str string) int {
	file, err := os.Create(namef)
	if err != nil {
		// handle the error here
		return -1
	}
	defer file.Close()

	file.WriteString(str)
	return 0
}

// - end функции генерации html page

// сохраняет в файл csv результат запроса в файл с именем namef
func savehttptocsv(namef string, suri string, suri2 string) int {
	// Create a new browser and open reddit.
    bow := surf.NewBrowser()
    err := bow.Open(suri)
    if err != nil {
        panic(err)
    }

    //bow = surf.NewBrowser()
    err = bow.Open(suri2)
    if err != nil {
        panic(err)
    }
	rescsv:=bow.Body()
	
	savestrtofile("report.csv",rescsv)
	return 0
}

// структура входящих данных
type InputDataTel struct {
	datacall  string // время и дата звонка
	telsource string // источник звонка (кто звонил)
	secs      int    // продолжительность в сек
	teldest   string // куда звонил источник

}

// структура справочника телефонов менеджеров
type DataTelMans struct {
	fio_rg    string // ФИО РГ
	fio_man   string // ФИО менеджера
	totalsec  int    // общая продолжительность звонков (в сек)
	kolunik   int    //кол-во уникальных телефонных номеров
	kolresult int    //кол-во результативных звоноков
	secresult int    // продолжительность результативных звонков (в сек)
	totalzv   int  // общее кол-во звоноков
}



func num_mes(m time.Month) int { //переводит из типа time.Month в число
	res := 0
	switch m {
	case 1:
		res = 1
	case 2:
		res = 2
	case 3:
		res = 3
	case 4:
		res = 4
	case 5:
		res = 5
	case 6:
		res = 6
	case 7:
		res = 7
	case 8:
		res = 8
	case 9:
		res = 9
	case 10:
		res = 10
	case 11:
		res = 11
	case 12:
		res = 12

	}
	return res

}

func main() {
	namef := "Report.csv"
	nameFlog := "list-num-tel.csv"
	namefresult := "log-zvonkov"	
	res_sec := 20 // маркер результативности звонка менеджера (в сек)

//----------------------------------------------
	if !parse_args() {
	   return
 	}
	
	var(
		begyearmonth, begday, endyearmonth, endday string
	)
	
	curdate := time.Now()
	
	if (d1!="") {
		begyearmonth,begday=parse_date(d1)		
	}
	if (d2!="") {
			endyearmonth,endday=parse_date(d2)		
			}else {		
				tekyear, tekmonth, tekday := time.Now().Date()
				begyearmonth=strconv.Itoa(tekyear) + "-" + strconv.Itoa(num_mes(tekmonth))
				endyearmonth=strconv.Itoa(tekyear) + "-" + strconv.Itoa(num_mes(tekmonth))
				fmt.Println(tekday)
				begday=strconv.Itoa(tekday)
				endday=strconv.Itoa(tekday)
			}		
//----------------------------------------------

	tekdate := begyearmonth+"-"+begday
	fmt.Println(tekdate)
	
	suri := "http://voip.2gis.local/cisco-stat/cdr.php?s=1&t=&order=dateTimeOrigination&sens=DESC&current_page=0&posted=1&current_page=0&fromstatsmonth=" + begyearmonth + "&tostatsmonth=" + endyearmonth + "&Period=Day&fromday=true&fromstatsday_sday=" + begday + "&fromstatsmonth_sday=" + begyearmonth + "&today=true&tostatsday_sday=" + endday + "&tostatsmonth_sday=" + endyearmonth + "&callingPartyNumber=&callingPartyNumbertype=2&originalCalledPartyNumber=%2B7&originalCalledPartyNumbertype=2&origDeviceName=&origDeviceNametype=1&destDeviceName=&destDeviceNametype=1&resulttype=min&image16.x=28&image16.y=8"
	fmt.Println(suri)
	suri2 := "http://voip.2gis.local/cisco-stat/export_csv.php"
	fmt.Println(suri2)
	
	savehttptocsv(namef,suri,suri2)

	str := readfilecsv(namef)
	strnumtel := readcfgs(nameFlog)

	//загрузка конфига справочника
	// ВЫБОРКА НУЖНЫХ ПОЛЕЙ: дата,источник звонка, продолжительность звонка,номер куда звонили
	vv := strings.Split(str, "\n")
	var vv1 []string
	s_inputdata := make([]InputDataTel, 0)
	for i := 0; i < len(vv)-1; i++ {
		vv1 = strings.Split(vv[i], ";")
		isec, _ := strconv.Atoi(vv1[10]) //конвертация из string в int
		s_inputdata = append(s_inputdata, InputDataTel{vv1[0], vv1[1], isec, vv1[2]})		
	}
	
    fmt.Println(s_inputdata)



	var buf_telunik map[string]int // буфер уникальных номеров для текущего внутр номера - длина этого map будет кол-во уникальных номеров
	ss := make([]InputDataTel, 0)
	kolres := 0
	totressec := 0
	totsec := 0
	totkol:=0 // общее кол-во звонков
	for key, _ := range strnumtel {
		numtel := key
		buf_telunik = make(map[string]int)
		totkol=0 // общее кол-во звонков
		kolres = 0    // счетчик кол-ва результативных звонков
		totressec = 0 // счетчик продолжительности результативных звонков
		totsec = 0    // счетчик общей продолжительности звонков
		// фильтрация по номеру телефона который указан в последовательности numtel
		for i := 0; i < len(s_inputdata)-1; i++ {
			if strings.Contains(s_inputdata[i].telsource, numtel) {
				ss = append(ss, s_inputdata[i])
				buf_telunik[s_inputdata[i].teldest] += 1
				totsec += s_inputdata[i].secs
				totkol+=1
				if s_inputdata[i].secs >= res_sec { // фильтрация по условию результирующего звонка
					kolres += 1
					totressec += s_inputdata[i].secs
				}
			}			
		}
		tm := strnumtel[key]
		strnumtel[key] = DataTelMans{tm.fio_rg, tm.fio_man, totsec, len(buf_telunik), kolres, totressec,totkol}
	}

	 fmt.Println(strnumtel)

	savetoxlsx(namefresult+".xlsx", strnumtel)
	str_title := "Лог звонков:  с \n" + begyearmonth + "-" + begday + " по " + endyearmonth + "-" + endday + ". Выгружено: " + curdate.String()
	htmlresult := genhtmlpage(strnumtel, str_title)
	savestrtofile(namefresult+".html", htmlresult)
	
	fmt.Println("The end....")

	//savetopdf("лог звонков.pdf",strnumtel)

}
