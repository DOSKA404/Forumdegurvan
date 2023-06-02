import { DateCheck } from "../function/DateCheck.js";
const passwordCheck = () => {
    /*
    * This function will be called every time the user change a letter in a password entry
    */
    const p = document.getElementById("pswError")
    if(document.getElementById('psw1').value === document.getElementById('psw2').value) {
        p.innerText = ""
        p.classList.remove('register-Error-on')
        p.classList.add('register-Error-off')
    } else {
        p.innerText = "passwords are not the same"
        p.classList.add('register-Error-on')
        p.classList.remove('register-Error-off')
    }
}



document.getElementById('psw1').addEventListener( 'change' , passwordCheck)
document.getElementById('psw2').addEventListener( 'change' , passwordCheck)
document.getElementById('date').addEventListener('click' , () => {
    document.getElementById('dateError').innerText = ''
    document.getElementById('dateError').classList.remove('register-Error-on')
    document.getElementById('dateError').classList.add('register-Error-off')
    
})
document.getElementById('email').addEventListener('click' , () => {
    document.getElementById('emailError').classList.remove('register-Error-on')
    document.getElementById('emailError').classList.add('register-Error-off')
    document.getElementById('emailError').innerText = ""
})
document.getElementById('pseudo').addEventListener('click' , () => {
    document.getElementById('pseudoError').classList.remove('register-Error-on')
    document.getElementById('pseudoError').classList.add('register-Error-off')
    document.getElementById('pseudoError').innerText = ""
})

document.getElementById('date').addEventListener('change' , (e) => {
    document.getElementById('dateError').innerText = ''
    const Check = DateCheck(e.target.value)
    if( Check !== true) {
        dateError.innerText = Check
        dateError.classList.add('register-Error-on')
        dateError.classList.remove('register-Error-off')
    }
})
document.getElementById('email').addEventListener('change' , (e) => {
    if(/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(e.target.value) === false){
        document.getElementById('emailError').classList.add('register-Error-on')
        document.getElementById('emailError').classList.remove('register-Error-off')
        document.getElementById('emailError').innerText = "The email is not valid"
    }
})
document.getElementById('pseudo').addEventListener('change' , (e) => {
    if(/^[A-Za-z0-9_-]{3,16}$/.test(e.target.value) === false) {
        document.getElementById('pseudoError').classList.add('register-Error-on')   
        document.getElementById('pseudoError').classList.remove('register-Error-off')
        document.getElementById('pseudoError').innerText = "The pseudo must be betwin 3 and 16 caracters , only '-' and '_' are autorized has special caracters"
    }
})

document.getElementById('RegisterButton').addEventListener( 'click' , (e) => {
    const passwordError = document.getElementById('pswError')
    if(document.getElementById('psw1').value === "") {
        passwordError.classList.add('register-Error-on')
        passwordError.classList.remove('register-Error-off')
        passwordError.innerText = "the password is empty"
    } else {
        document.getElementById('post').submit();
    }

})