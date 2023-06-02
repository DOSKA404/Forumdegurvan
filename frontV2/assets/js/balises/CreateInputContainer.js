export const CreateInputContainer = ( title, placeholder, id  , imgPath , password , secondPassword) => {
    const emailcontainer = document.createElement('div'); emailcontainer.classList.add('login-container');
    const emailh2 = document.createElement('h2'); emailh2.innerText = title;
    const emaildivinput = document.createElement('div');
    const emailimg = document.createElement('img'); emailimg.src = imgPath; emailimg.classList.add('login-insideImage');
    const emailinput = document.createElement('input'); emailinput.type = 'email'; emailinput.placeholder = placeholder;emailinput.id = id;
    emaildivinput.appendChild(emailimg); emaildivinput.appendChild(emailinput);
    emailcontainer.appendChild(emailh2); emailcontainer.appendChild(emaildivinput);
    if (password) {
      if(secondPassword){
        const eye = document.createElement('img'); eye.src = 'http://localhost:8080/assets/img/eye_on.svg'; eye.classList.add('login-eye'); eye.id = 'eye2'; 
       eye.addEventListener('click', () => {
              if (document.getElementById('psw2').type == 'password') {
                document.getElementById('psw2').type = 'text';
                document.getElementById('eye2').src = 'http://localhost:8080/assets/img/eye_off.svg';
              } else {
                document.getElementById('psw2').type = 'password';
                document.getElementById('eye2').src = 'http://localhost:8080/assets/img/eye_on.svg';
              }
       });
       emaildivinput.appendChild(eye);

      } else {
        const eye = document.createElement('img'); eye.src = 'http://localhost:8080/assets/img/eye_on.svg'; eye.classList.add('login-eye'); eye.id = 'eye'; 
        eye.addEventListener('click', () => {
              if (document.getElementById('psw').type == 'password') {
                document.getElementById('psw').type = 'text';
                document.getElementById('eye').src = 'http://localhost:8080/assets/img/eye_off.svg';
              } else {
                document.getElementById('psw').type = 'password';
                document.getElementById('eye').src = 'http://localhost:8080/assets/img/eye_on.svg';
              }
       });
       emaildivinput.appendChild(eye);

      }
       
       emailinput.type = 'password';
    }
    return emailcontainer;
}