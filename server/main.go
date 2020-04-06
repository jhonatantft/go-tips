package main
// Importamos 4 libs
// 1. “net/http” para acessar as funcionalidades http
// 2. “fmt” para formatar nosso texto
// 3. “html/template” uma lib que nos permite integarir com arquivos html
// 4. "time" - uma lib pra trabalhar com data e hora
import (
   "net/http"
   "fmt"
   "time"
   "html/template"
)

// Cria um struct que tem os dados que serão mostrados no html
type Welcome struct {
   Name string
   Time string
}

func main() {
   // Instancia o struc welcome e passa algumas informações aleatórias
   //We shall get the name of the user as a query parameter from the URL

   welcome := Welcome{"Mundo", time.Now().Format(time.Stamp)}

	// Dizemos ao Go onde encontraremos nosso arquivo html. Pedimos ao Go para parsear esse arquivo
   // E colocamos esse retorno rentro de template.Must() que já lida com erros se tiver algum
   
   templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

   //Our HTML comes with CSS that go needs to provide when we run the app. Here we tell go to create
   //Nosso html vem com css que o Go precivsa para prover o app. E aqui dizemos para criar
   // um handle para olhar na pasta static. Go usa static como url para o html refereciar nosso
   // arquivo css
   
   http.Handle("/static/",
      http.StripPrefix("/static/",
         http.FileServer(http.Dir("static"))))

   // Esse método recebe o caminho da URL e uma função que receve a response e a request
   http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {

	  // Recebe o nome da query da URL - ex: ?name=jhonatan - que foi passado como parâmetro. E seta welcome.Name = Jhonatan 
      if name := r.FormValue("name"); name != "" {
         welcome.Name = name;
      }
	  // Se tiver erro, mostra o erro interno do server
	  // Passo o struct welcome para o welcome-template.html
      if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
      }
   })

   // Starto o web server, seto a porta para 8080, sem um path será por default localhost
   // Printa qualquer erro 
   fmt.Println("Listening");
   fmt.Println(http.ListenAndServe(":8089", nil));
}