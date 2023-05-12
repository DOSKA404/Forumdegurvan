export const Login = () => {
    const linkcss = document.createElement('link'); linkcss.rel = 'stylesheet';linkcss.href = 'http://localhost:8080/assets/css/login.css'; linkcss.id = "";document.head.appendChild(linkcss);
    const section = document.createElement('section'); section.classList.add('login-section');
    const emailcontainer = document.createElement('div'); emailcontainer.classList.add('login-container');
    const emailh2 = document.createElement('h2'); emailh2.innerText = 'Enter your Email';
    const emaildivinput = document.createElement('div');
    const emailimg = document.createElement('img'); emailimg.src = 'http://localhost:8080/assets/img/email.svg';
    const emailinput = document.createElement('input'); emailinput.type = 'email'; emailinput.placeholder = 'exemple.gmail.com';
    emaildivinput.appendChild(emailimg); emaildivinput.appendChild(emailinput);
    emailcontainer.appendChild(emailh2); emailcontainer.appendChild(emaildivinput);
    section.appendChild(emailcontainer);
    document.body.appendChild(section);
}