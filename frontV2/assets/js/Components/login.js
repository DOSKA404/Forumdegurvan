import { CreateButton } from "../balises/button.js";
import { CreateInputContainer } from "../balises/CreateInputContainer.js";
export const Login = (fail) => {
    const linkcss = document.createElement('link'); linkcss.rel = 'stylesheet';linkcss.href = 'http://localhost:8080/assets/css/login.css'; linkcss.id = "";document.head.appendChild(linkcss);
    const section = document.createElement('section'); section.classList.add('login-section');
    if(fail){
        const fail = document.createElement('p'); fail.innerText = "Email or password incorrect"; fail.classList.add('login-fail'); section.appendChild(fail);
    }
    section.appendChild(CreateInputContainer('Enter your email','exemple@gmail.com' , "email" , "http://localhost:8080/assets/img/email.svg"));
    section.appendChild(CreateInputContainer('Enter your password','********' , "psw" , "http://localhost:8080/assets/img/password.svg" , true));
    CreateButton('Login', 'LoginButton', 'Button1', section);
    const toCreate = document.createElement('p'); toCreate.innerText = "Don't have an account ?"; toCreate.classList.add('login-toCreate'); toCreate.id = "tocreate"; section.appendChild(toCreate);
    document.body.appendChild(section);
}