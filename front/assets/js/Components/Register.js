import { CreateInputContainer } from "../balises/CreateInputContainer.js";
import {CreateButton } from "../balises/button.js";
export const Register = () => {
    const linkcss = document.createElement('link'); linkcss.rel = 'stylesheet';linkcss.href = 'http://localhost:8080/assets/css/register.css'; linkcss.id = "";document.head.appendChild(linkcss);
    const section = document.createElement('section'); section.classList.add('register-section');
    section.appendChild(CreateInputContainer("Enter your email" , 'exemple@gmail.com','email',"http://localhost:8080/assets/img/email.svg"))
    section.appendChild(CreateInputContainer("Choose a password" , '********','psw',"http://localhost:8080/assets/img/password.svg",true))
    section.appendChild(CreateInputContainer("Confirm your password" , '********','psw2',"http://localhost:8080/assets/img/password.svg",true,true))
    section.appendChild(CreateInputContainer("Choose a username" , 'xx-intercalaire-xx','pseudo',"http://localhost:8080/assets/img/user.svg"))
    const date = document.createElement('input'); date.type = 'date'; date.id = 'date'; date.classList.add('register-date') ;section.appendChild(date)
    CreateButton("Register","RegisterButton","button2",section)
    document.body.appendChild(section);
}