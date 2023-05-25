//components
import { Login } from "./Components/login.js";
import { Footer } from "./Components/footer.js";
<<<<<<< HEAD
import { Register } from "./Components/Register.js";
// functions
=======
import {menu} from "./Components/mainPage.js";
>>>>>>> 9f42f243ebc81e5b4e7f987c12b367bcede615ed
import { Sup } from "./function/Sup.js"
const LoginPage = (fail) => {
    Login(fail)
    document.getElementById('LoginButton').addEventListener('click', () => {
        Connnexion(document.getElementById('email').value,document.getElementById('psw').value)
    })
    document.getElementById('tocreate').addEventListener('click', () => {
        Sup()
        RegisterPage()
    })
    Footer()
}
const MainPage =()=> {
    menu()
    Footer()
}
MainPage()

const RegisterPage = () => {
    Register(); 
    document.getElementById('RegisterButton').addEventListener('click', () => {
        console.log("hi")
        CreateAccount(document.getElementById('email').value,document.getElementById('psw').value,document.getElementById('psw2').value,document.getElementById('date').value , document.getElementById("pseudo").value)
    })
    Footer()
}
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

<<<<<<< HEAD
const CreateAccount = (email,password1,password2,date,pseudo) => {
    if(password1 === password2){
        //todo check if email is not already used
        //todo create account
        /*fetch("http://localhost:6969/createUser", { method : "POST" ,body : JSON.stringify({email : email , password : password1 , dateOfBirth : date, username : pseudo}), headers: {'Content-type': 'application/json; charset=UTF-8' }})
        .then(function (response) {
            console.log(response)
            });*/
        console.log({email : email , password : password1 , dateOfBirth : date, username : pseudo})
    }
    
}
LoginPage() // We lunch the page
RegisterPage()
=======
LoginPage()
>>>>>>> 9f42f243ebc81e5b4e7f987c12b367bcede615ed
