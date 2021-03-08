package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"weather/law"
	glocate "weather/law"
)

var s string

var tmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html>
	<head><title>weather</title></head>
	<body>

	<form action="/draw">
	<h1  style="text-align:center" class="mt-5">現在の天気</h1>
	<p style="text-align:center" >天気を調べたい場所を入れてみてね　　例:ハチ公前</p>
	       
			<p  style="text-align:center"  ><input type="search" minlength="1" name="locate" required="required"></p>
			
			<p style="text-align:center"><input style="color:文字色;background-color:背景色;font-size:15;width:150px;height:30px" type="submit" value="送信する" >
			<input type="reset" value="入力内容をリセットする">
		</form>
		
		<ol style="text-align :center">
		{{.}}
		</ol>
	</main>
	</body>
	</html>`))

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, s)

	})
	http.HandleFunc("/draw", func(w http.ResponseWriter, r *http.Request) {
		run(r.FormValue("locate"))
		http.Redirect(w, r, "/", http.StatusFound)
	})

	http.ListenAndServe(":8080", nil)
}
func run(st string) {

	data := law.HttpGet(("https://www.geocoding.jp/api?" + "q=" + st))

	v := law.XmltoD(data) //構造体のポインタを返している
	x := v.Cor.Lat
	y := v.Cor.Lng

	data2 := glocate.HttpGet("http://api.openweathermap.org/data/2.5/weather?lat=" + x + "&lon=" + y + "&APPID=7dacbd0e200a5c44b76844aa279c7aa7")

	w := law.WtoD(data2)

	m := law.Getadd(y, x)
	if strings.Contains(m.Err, "Cities") {
		print("えら")
		os.Exit(1)
	}

	var str = fmt.Sprintln(int(w.Main.Temp)-273, "℃ で", law.Honyaku(w.Weather[0].Main), " in ", m.Location[0].Prefecture, "の", m.Location[0].City /*, m.Location[0].Town*/, "\n")
	s = str
}

/*package main

import (
	"fmt"
	"weather/law"
	glocate "weather/law"
)

func main() {
	for {
		var a string
		print("天気を知りたい場所を入力してください>")
		fmt.Scan(&a)
		data := law.HttpGet("https://www.geocoding.jp/api?" + "q=" + a)

		v := law.XmltoD(data) //構造体のポインタを返している
		x := v.Cor.Lat
		y := v.Cor.Lng

		data2 := glocate.HttpGet("http://api.openweathermap.org/data/2.5/weather?lat=" + x + "&lon=" + y + "&APPID=7dacbd0e200a5c44b76844aa279c7aa7")

		w := law.WtoD(data2)

		m := law.Getadd(y, x)
		fmt.Print(int(w.Main.Temp)-273, "℃ で", law.Honyaku(w.Weather[0].Main), " in ")
		fmt.Println(m.Location[0].Prefecture, "の", m.Location[0].City /*, m.Location[0].Town, "\n")

	}

}*/
