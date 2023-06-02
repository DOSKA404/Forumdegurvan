import { CreateInputContainer } from "../balises/CreateInputContainer.js";
import {CreateButton } from "../balises/button.js";
import { DateCheck } from "../function/DateCheck.js";
export const Register = (EmailErrorAPI , usernameErrorAPI) => {
    const linkcss = document.createElement('link'); linkcss.rel = 'stylesheet';linkcss.href = 'http://localhost:8080/assets/css/register.css'; linkcss.id = "";document.head.appendChild(linkcss);
    const section = document.createElement('section'); section.classList.add('register-section');
    section.appendChild(CreateInputContainer("Enter your email" , 'exemple@gmail.com','email',"http://localhost:8080/assets/img/email.svg"))
    const emailError = document.createElement("p") ; emailError.classList.add('register-Error-on') ; emailError.id = 'emailError';section.appendChild(emailError)
    section.appendChild(CreateInputContainer("Choose a password" , '********','psw',"http://localhost:8080/assets/img/password.svg",true))
    section.appendChild(CreateInputContainer("Confirm your password" , '********','psw2',"http://localhost:8080/assets/img/password.svg",true,true))
    const passwordError = document.createElement("p") ; passwordError.classList.add('register-Error-off') ; passwordError.id = 'passwordError';section.appendChild(passwordError)
    section.appendChild(CreateInputContainer("Choose a username" , 'xx-intercalaire-xx','pseudo',"http://localhost:8080/assets/img/user.svg"))
    const pseudoError = document.createElement("p") ; pseudoError.classList.add('register-Error-on') ; pseudoError.id = 'pseudoError';section.appendChild(pseudoError)
    const date = document.createElement('input'); date.type = 'date'; date.id = 'date'; date.classList.add('register-date');section.appendChild(date)
    const dateError = document.createElement("p") ; dateError.classList.add('register-Error-on') ; dateError.id = 'dateError';section.appendChild(dateError)
    CreateButton("Register","RegisterButton","button2",section)
    document.body.appendChild(section);
    /**
     * Error part 
     */
    if(EmailErrorAPI) {
        emailError.innerText = 'this email is already use by another account'
    }
    if(usernameErrorAPI) {
        pseudoError.innerText = 'this pseudo is already use by another account'
    }
    /**
     * listenners part
     */
    document.getElementById('psw').addEventListener( 'change' , passwordCheck)
    document.getElementById('psw2').addEventListener( 'change' , passwordCheck)
    document.getElementById('RegisterButton').addEventListener( 'click' , () => {
        const Check = DateCheck(document.getElementById("date").value)
        if( Check !== true) {
            dateError.innerText = Check
        }
        if(/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(document.getElementById("email").value) === false){
            emailError.innerText = "invalid email"
        }
        if(/^[A-Za-z0-9_-]{3,16}$/.test(document.getElementById("pseudo").value) === false) {
            pseudoError.innerText = "the pseudo must be betwin 3 and 16 caracters , only '-' and '_' are autorized has special caracters"
        }
        if(document.getElementById('psw').value === "" ) {
            passwordError.classList.add('register-Error-on')
            passwordError.classList.remove('register-Error-off')
            passwordError.innerText = "the password is empty"
        } 
    })
    date.addEventListener('click' , () => {
        document.getElementById('dateError').innerText = ''
    })
    document.getElementById('email').addEventListener('click' , () => {
        emailError.innerText = ''
    })
    document.getElementById('pseudo').addEventListener('click' , () => {
        pseudoError.innerText = ''
    })
}

const passwordCheck = () => {
    /*
    * This function will be called every time the user change a letter in a password entry
    */
    const p = document.getElementById("passwordError")
    if(document.getElementById('psw').value === document.getElementById('psw2').value) {
        p.innerText = ""
        p.classList.remove('register-Error-on')
        p.classList.add('register-Error-off')
    } else {
        p.innerText = "passwords are not the same"
        p.classList.add('register-Error-on')
        p.classList.remove('register-Error-off')
    }
}