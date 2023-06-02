//components
import { Login } from "./Components/login.js";
import { Footer } from "./Components/footer.js";
import { Register } from "./Components/Register.js";
import {menu} from "./Components/mainPage.js";
import { DateCheck } from "./function/DateCheck.js";
// functions
import { Sup } from "./function/Sup.js"
import { Create_cookie , get_token } from "./function/Cookies.js"
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
//MainPage()

const RegisterPage = () => {
    Register(); 
    document.getElementById('RegisterButton').addEventListener('click', () => {
        console.log('oui')
        CreateAccount(document.getElementById('email').value,document.getElementById('psw').value,document.getElementById('psw2').value,document.getElementById('date').value , document.getElementById("pseudo").value)
    })
    Footer()
}
const Connnexion = (email,psw) => {

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


const CreateAccount = (email,password1,password2,date,pseudo) => {
    console.log(email,password1,password2,date,pseudo)
    if(password1 === password2 && password1 !== '' && DateCheck(date) && /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(document.getElementById("email").value) && /^[A-Za-z0-9_-]{3,16}$/.test(document.getElementById("pseudo").value)){
        console.log('ok')
        const xhr = new XMLHttpRequest();

// Data to be sent in the request
        const data = {
            email: email,
            password: password1,
            date: date,
            pseudo: pseudo
        };

    xhr.open('POST', "http://localhost:8080/", true);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest'); 
    xhr.setRequestHeader('Access-Control-Allow-Origin', '*');
    xhr.onreadystatechange = function () {
        console.log(this);
        if (xhr.readyState === 4 && xhr.status === 200) {
            console.log(xhr.responseText);
        } else {
            console.log(xhr.responseText);
        }
    };
    console.log(JSON.stringify(data))
    xhr.send(JSON.stringify(data)); 
    }
    
}
LoginPage() // We lunch the page

/*if (get_token() === null){
    LoginPage()
} else {
    MainPage()
}*/