//components
import { Login } from "./Components/login.js";
import { Footer } from "./Components/footer.js";
import {menu} from "./Components/mainPage.js";
import { Sup } from "./function/Sup.js"
const LoginPage = (fail) => {
    Login(fail)
    document.getElementById('LoginButton').addEventListener('click', () => {
        Connnexion(document.getElementById('email').value,document.getElementById('psw').value)
    })
    Footer()
}
const MainPage =()=> {
    menu()
    Footer()
}
MainPage()

const Connnexion = (email,psw) => {
    console.log(email , psw )
    /*fetch('http://localhost:6969/login')
    .then(response => response.json())
    .then(data => console.log(data))*/
    //Todo link to db
    const Respond = false

    if(Respond){
        console.log('connect')
        //todo connect
    } else {
        Sup()
        LoginPage(true)
    }
}

LoginPage()
