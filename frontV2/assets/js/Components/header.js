let open = false;
document.getElementById("openAccountOptions").addEventListener("click", () => {
    if(open){
        document.getElementById("openAccountOptions").src = "assets/img/user.svg"
    } else {
        document.getElementById("openAccountOptions").src = "assets/img/close.svg"
    }
    
    open = !open;
    document.getElementById("AccountOption").classList.toggle("hidden");
    document.getElementById("AccountOption").classList.toggle("AccountOption");
    
    document.getElementById("leftside").classList.toggle("margin-right");
});