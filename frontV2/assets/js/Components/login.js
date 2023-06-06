document.getElementById('eye').addEventListener('click' , () => {
    if(document.getElementById('psw').type === "password") {
       document.getElementById('eye').src = "/assets/img/eye_off.svg" 
        document.getElementById('psw').type = "text"
    } else {
        document.getElementById('eye').src = "/assets/img/eye_on.svg" 
        document.getElementById('psw').type = "password"
    }
})

document.getElementById('LoginButton').addEventListener('click', (e) => {
    document.getElementById('post').submit();
})